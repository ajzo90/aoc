package main

import (
	. "aoc2022/pkg/aoc"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Node struct {
	Child []*Node
	Name  string
	Size  int
	IsDir bool
	Path  []string
}

func (n *Node) Add(o Node) {
	if !n.IsDir {
		panic("not a dir")
	}
	for _, ch := range n.Child {
		if ch.Name == o.Name {
			return
		}
	}
	o.Path = append(o.Path, n.Path...)
	o.Path = append(o.Path, o.Name)
	n.Child = append(n.Child, &o)
}

func (n *Node) DirSize() int {
	if !n.IsDir {
		return n.Size
	}
	var size int
	for _, ch := range n.Child {
		size += ch.DirSize()
	}
	return size
}

func (n *Node) Find(path []string) *Node {
	for len(path) > 0 {
		var ok bool
		for _, ch := range n.Child {
			if ch.Name == path[0] {
				n = ch
				path = path[1:]
				ok = true
				break
			}
		}
		if !ok {
			panic(fmt.Sprintf("not found: %v from %v", path, n.Path))
		}
	}
	return n
}

func (n *Node) Walk(f func(*Node)) {
	f(n)
	for _, ch := range n.Child {
		ch.Walk(f)
	}
}

func part1(isPart2 bool) func(text string, rows []string) string {
	return func(text string, rows []string) string {

		var wd []string
		var tree = Node{IsDir: true, Name: "/"}

		for len(rows) > 0 {
			v := rows[0]
			rows = rows[1:]

			if v[0] == '$' {
				log.Println("CMD", v)
				parts := strings.Split(v[2:], " ")
				cmd := parts[0]
				args := parts[1:]
				switch cmd {
				case "ls":
					for len(rows) > 0 && rows[0][0] != '$' {
						n := tree.Find(wd)
						args = strings.Split(rows[0], " ")
						log.Println("ls", args)

						switch args[0] {
						case "dir":
							n.Add(Node{Name: args[1], IsDir: true})
						default:
							n.Add(Node{Name: args[1], IsDir: false, Size: Int(args[0])})
						}
						rows = rows[1:]
					}

				case "cd":
					switch args[0] {
					case "..":
						wd = wd[:len(wd)-1]
					case "/":
						wd = nil
					default:
						tree.Find(wd).Add(Node{Name: args[0], IsDir: true})
						wd = append(wd, args[0])
					}
				}
			}
		}
		if isPart2 {
			sz := tree.DirSize()
			avail := 70000000 - 30000000
			diff := sz - avail

			var min = math.MaxInt64

			tree.Walk(func(n *Node) {
				log.Println("SIZE", n.DirSize())
				if sz := n.DirSize(); sz >= diff && n.IsDir && sz < min {
					min = sz
				}
			})
			return strconv.Itoa(min)
		}

		var sum = 0
		tree.Walk(func(n *Node) {
			if sz := n.DirSize(); sz <= 100000 && n.IsDir {
				sum += sz
			}
		})

		return strconv.Itoa(sum)
	}
}

func main() {

	var aoc = New()

	var input = aoc.Input()

	Assert(example, "95437", part1(false))

	aoc.Submit(1, part1(false)(input, nil))

	Assert(example, "24933642", part1(true))

	aoc.Submit(2, part1(false)(input, nil))

}

const example = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
