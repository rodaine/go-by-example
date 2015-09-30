package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var p = fmt.Println

func main() {
	pattern := "p([a-z]+)ch"
	s := "peach punch pinch"

	match, _ := regexp.MatchString(pattern, s)
	p(match)

	r, _ := regexp.Compile(pattern)

	p(r.MatchString(s))
	p(r.FindString(s))
	p(r.FindStringSubmatch(s))
	p(r.FindStringSubmatchIndex(s))

	p(r.FindAllString(s, -1))
	p(r.FindAllString(s, 2))
	p(r.FindAllStringSubmatchIndex(s, -1))

	p(r.Match([]byte("peach")))

	r = regexp.MustCompile(pattern)
	p(r)

	p(r.ReplaceAllString(s, "<fruit>"))

	in := []byte(s)
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	p(string(out))
}
