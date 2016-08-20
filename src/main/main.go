package main

import (
	"fmt"

	"github.com/rebeccaskinner/golang-collections/src/list"
)

func PrintList(l list.List) {
	printFunc := func(i interface{}) { fmt.Println(i) }
	list.MapM(printFunc, l)
}

func demoConcat() {
	listOne := list.New(1, 2, 3)
	listTwo := list.New(4, 5, 6)
	PrintList(list.Concat(listOne, listTwo))
}

func main() {
	demoConcat()
}
