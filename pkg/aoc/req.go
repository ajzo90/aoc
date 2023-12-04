package aoc

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var session string

func buildReq(url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	var cookie = &http.Cookie{Name: "session", Value: session}
	req.AddCookie(cookie)
	return req, nil
}

func Req(url string) (*http.Response, error) {
	req, err := buildReq(url, nil)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Second)
	return http.DefaultClient.Do(req)
}

func Submit(year, day int, part int, answer string) error {

	name := fmt.Sprintf("aoc_answer_%d_%d_%d_%sX", year, day, part, answer)
	return withCachedFile(name, func() (*http.Response, error) {
		Url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)

		var v = url.Values{}
		v.Add("level", strconv.Itoa(part))
		v.Add("answer", answer)

		enc := v.Encode()

		req, err := buildReq(Url, strings.NewReader(enc))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		//panic("send req..")

		return http.DefaultClient.Do(req)
	}, func(f *os.File) error {

		data, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		if bytes.Contains(data, []byte("That's not the right answer.")) {
			//idx := bytes.Index(data, []byte("please wait"))
			log.Println(string(data))
			panic("incoreect")
			//p, _, _ := bytes.Cut(data[idx:], []byte("\n"))
			//log.Fatal(string(p))
		} else if bytes.Contains(data, []byte("You gave an answer too recently; you have to wait after submitting an answer before trying again.")) {
			os.Remove(f.Name())
			panic("wait")
		} else if bytes.Contains(data, []byte("That's the right answer!")) {
			fmt.Println("Correct answer!", answer)
		} else {
			io.Copy(os.Stdout, bytes.NewReader(data))
		}
		return nil
	})
}

func Assert(input string, expected string, f func(string, []string) string) {
	got := f(input, strings.Split(strings.Trim(input, "\n"), "\n"))
	if expected != got {
		panic(fmt.Sprintf("expected '%s', got '%s'", expected, got))
	}
	fmt.Println("Assert OK")
}

func (aoc *aoc) Input() (string, []string) {
	data := Input(aoc.year, aoc.day)
	return data, strings.Split(strings.Trim(data, "\n"), "\n")
}

func (aoc *aoc) Submit(part int, answer string) {
	log.Println("submit", aoc.year, aoc.day, part, answer)
	if err := Submit(aoc.year, aoc.day, part, answer); err != nil {
		panic(err)
	}
}

func Input(year, day int) string {
	var data string

	err := withCachedFile(fmt.Sprintf("aoc_input_%d_%d_", year, day), func() (*http.Response, error) {
		return Req(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))
	}, func(f *os.File) error {
		data = string(Must(io.ReadAll(f)))
		return nil
	})
	if err != nil {
		panic(err)
	}
	return data
}

func RunPart(year int, day int, part int, fn func(f io.Reader) string) {

	var p = fmt.Sprintf("aoc_input_%d_%d", year, day)

	var cb = func(f *os.File) error {

		fmt.Println("run part", part)
		fmt.Println()
		got := fn(f)

		if len(got) > 0 {
			return Submit(year, day, part, got)
		}
		return nil
	}

	err := withCachedFile(p, func() (*http.Response, error) {
		return Req(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))
	}, cb)
	if err != nil {
		panic(err)
	}
}

type aoc struct {
	year    int
	day     int
	example string
}

func (a *aoc) Part(p int, fn func(r io.Reader) string) *aoc {
	RunPart(a.year, a.day, p, func(f io.Reader) string {
		return fn(f)
	})
	return a
}

func (a *aoc) part(part int, cb func(string, []string) string) *aoc {
	return a.Part(part, func(r io.Reader) string {
		all := string(bytes.Trim(Must(io.ReadAll(r)), "\n"))
		rows := strings.Split(all, "\n")
		return cb(all, rows)
	})
}

func (a *aoc) Part2(cb func(string, []string) string) *aoc {
	return a.part(2, cb)
}

func New() *aoc {

	_, file, _, _ := runtime.Caller(1)
	_, after, _ := strings.Cut(file, "cmd/")
	parts := strings.Split(after, "/")

	year, day := parts[0], parts[1]

	return &aoc{year: Int(year), day: Int(day)}
}

func withCachedFile(p string, cb func() (*http.Response, error), h func(f *os.File) error) error {
	p = os.TempDir() + p
	log.Println(p)
	f, err := os.Open(p)
	if os.IsNotExist(err) {
		f, err = os.CreateTemp("", "")
		if err != nil {
			return err
		}
		resp, err := cb()
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		log.Println("copy file", p)
		if _, err := io.Copy(f, resp.Body); err != nil {
			return err
		} else if f.Close(); err != nil {
			return err
		} else if err := os.Rename(f.Name(), p); err != nil {
			return err
		} else if err := resp.Body.Close(); err != nil {
			return err
		}
		f, err = os.Open(p)
	}
	if err != nil {
		return err
	}
	defer f.Close()
	return h(f)
}
