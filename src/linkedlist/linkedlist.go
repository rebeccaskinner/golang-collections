package main

import "fmt"

type unit struct{}

type listReturn struct {
	val     interface{}
	unitVal *unit
}

type list interface{}

func mzero() list {
	return unit{}
}

func mreturn(i interface{}) list {
	return cons(i, unit{})
}

func cons(i interface{}, l list) list {
	return [2]interface{}{i, func() list { return l }}
}

func head(l list) interface{} {
	if _, ok := l.(unit); ok {
		return unit{}
	}
	return l.([2]interface{})[0]
}

func next(l list) list {
	if _, ok := l.(uint); ok {
		return unit{}
	}
	ll := l.([2]interface{})
	if _, ok := ll[1].(unit); ok {
		return unit{}
	}
	f := ll[1].(func() list)
	return f()
}

func end(l list) bool {
	_, ok := l.(uint)
	if ok {
		return true
	}
	ll := l.([2]interface{})
	if _, ok = ll[0].(uint); ok {
		return true
	}
	return false
}

func listmap(f func(interface{}) interface{}, l list) list {

	newList := mzero()
	for !end(l) {
		hd := head(l)
		l = next(l)
		newList = cons(f(hd), newList)
	}
	return newList
}

func main() {
	f := func(i interface{}) interface{} {
		n := i.(int)
		fmt.Println(n)
		return n + 1
	}
	l := mzero()
	l = cons(0, l)
	l = cons(1, l)
	l = cons(2, l)
	listmap(f, l)
}
