package main

import (
	"fmt"
	"os"
)

var pl = fmt.Println
var pf = fmt.Printf

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	i := 123
	f := 12.3
	d := 4.56
	s := "\"string\""

	pf("%v\n", p)
	pf("%+v\n", p)
	pf("%#v\n", p)
	pf("%T\n", p)
	pf("%p\n", &p)
	pl()
	pf("%t\n", true)
	pl()
	pf("%d\n", i)
	pf("%b\n", i)
	pf("%c\n", i)
	pf("%x\n", i)
	pl()
	pf("%f\n", f)
	pf("%e\n", f)
	pf("%E\n", f)
	pl()
	pf("%s\n", s)
	pf("%q\n", s)
	pf("%x\n", s)
	pl()
	pf("|%6d|\n", i)
	pf("|%6.2f|%6.2f|\n", f, d)
	pf("|%-6.2f|%-6.2f|\n", f, d)
	pl()
	pf("|%6s|%-6s|\n", "foo", "bar")

	st := fmt.Sprintf("a %s", "string")
	pl(st)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
