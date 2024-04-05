package main

import (
	"fmt"

	"github.com/rajatjindal/foo-wasip2/foo-namespace/pkg/foo"
)

func init() {}

func main() {
	fmt.Println("hello")
	x := foo.GreetExported()
	fmt.Println(*x)
}
