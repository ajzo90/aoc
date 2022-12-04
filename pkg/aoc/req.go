package aoc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var session string

func Req(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	var cookie = &http.Cookie{Name: "session", Value: session}
	req.AddCookie(cookie)

	return http.DefaultClient.Do(req)
}

func Assert(data string, fn func(f io.Reader) string, expected string) {
	got := fn(strings.NewReader(data))
	if got != expected {
		log.Fatalf("unexpected result, got %s, expected %s", got, expected)
	}
}

func WithData(day int, fn func(f io.Reader) string, expected string) {
	if err := RunDay2022(day, fn, expected); err != nil {
		panic(err)
	}
}
func RunDay2022(day int, fn func(f io.Reader) string, expected string) error {
	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day)

	var p = fmt.Sprintf("%saoc2022_%d.input", os.TempDir(), day)
	log.Println(p)

	f, err := os.Open(p)
	if os.IsNotExist(err) {
		f, err = os.CreateTemp("", "")
		if err != nil {
			return err
		}
		log.Println("GET DATA")
		resp, err := Req(url)
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

	got := fn(f)
	if len(expected) > 0 && got != expected {
		log.Fatalf("unexpected result, got %s, expected %s", got, expected)
	}
	fmt.Println("result", got)

	return nil
}
