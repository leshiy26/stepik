package main

import (
	"fmt"
	"stepik/morestrings"

	"github.com/google/go-cmp/cmp"
	"github.com/leshiy26/stepik/module2"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(module2.TestModule2())
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
