package aoc

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
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
	return http.DefaultClient.Do(req)
}

func Submit(year, day int, part int, answer string) error {

	name := fmt.Sprintf("aoc_answer_%d_%d_%d_%s", year, day, part, answer)
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
			idx := bytes.Index(data, []byte("please wait"))
			p, _, _ := bytes.Cut(data[idx:], []byte("\n"))
			log.Fatal(string(p))
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

func Assert(data string, fn func(f io.Reader) string, expected string) {
	got := fn(strings.NewReader(data))
	if got != expected {
		log.Fatalf("unexpected result, got '%s', expected '%s'", got, expected)
	}
}

func WithData(year int, day int, part int, fn func(f io.Reader) string, expected string) {
	if err := Run(year, day, part, fn, expected); err != nil {
		panic(err)
	}
}
func RunPart(year int, day int, part int, exData string, exExpect string, finalRes string, fn func(f io.Reader) string) {
	fmt.Println("example part", part)
	Assert(exData, fn, exExpect)

	fmt.Println("run part", part)
	WithData(year, day, part, fn, finalRes)
}

type aoc struct {
	year    int
	day     int
	example string
}

func (a *aoc) Part(p int, exampleRes string, fn func(r io.Reader) int) *aoc {
	RunPart(a.year, a.day, p, a.example, exampleRes, "", func(f io.Reader) string {
		return strconv.Itoa(fn(f))
	})
	return a
}

func New(year int, day int, exampleData string) *aoc {
	return &aoc{year: year, day: day, example: exampleData}
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

func Run(year, day, part int, fn func(f io.Reader) string, expected string) error {

	var p = fmt.Sprintf("aoc_input_%d_%d", year, day)

	return withCachedFile(p, func() (*http.Response, error) {
		return Req(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))
	}, func(f *os.File) error {
		got := fn(f)
		if len(expected) > 0 && got != expected {
			log.Fatalf("unexpected result, got '%s', expected '%s'", got, expected)
		}
		if len(expected) == 0 {
			return Submit(year, day, part, got)
		}
		return nil
	})
}
