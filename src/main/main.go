package main

import (
	"fmt"

	"github.com/rebeccaskinner/golang-collections/src/list"
)

func main() {
	f := func(i interface{}) {
		fmt.Println(i)
	}
	l := list.Reverse(list.Cons(3, list.Cons(2, list.Cons(1, list.Return(0)))))
	list.MapM(f, l)
}
