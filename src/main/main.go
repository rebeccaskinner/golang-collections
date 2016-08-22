package main

import (
	"fmt"
	"time"

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

func demoLazy() {
	slowNum := func(i int) func() interface{} {
		return func() interface{} {
			fmt.Printf("sleepily generating %v\n", i)
			time.Sleep(3 * time.Second)
			return i
		}
	}
	mapFunc := func(i interface{}) interface{} {
		num := i.(int)
		fmt.Println("in map func...")
		time.Sleep(2 * time.Second)
		return (num + 5)
	}
	l := list.Mzero()
	l = list.Consf(slowNum(5), l)
	l = list.Consf(slowNum(4), l)
	l = list.Consf(slowNum(3), l)
	l = list.Consf(slowNum(2), l)
	l = list.Consf(slowNum(1), l)
	l = list.Map(mapFunc, l)
	fmt.Println(list.Head(l))
}

func main() {
	demoConcat()
	demoLazy()
}
