package main

import (
	"fmt"
	"time"
)

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
	return mreturnf(func() interface{} { return i })
}

func cons(i interface{}, l list) list {
	return consf(func() interface{} { return i }, l)
}

func consf(f func() interface{}, l list) list {
	if l == nil {
		l = mzero()
	}
	return [2]interface{}{f, func() list { return l }}
}

func mreturnf(f func() interface{}) list {
	return consf(f, mzero())
}

func head(l list) interface{} {
	if l == nil {
		l = mzero()
	}
	if _, ok := l.(unit); ok {
		return unit{}
	}
	lf := l.([2]interface{})[0].(func() interface{})
	return lf()
}

func next(l list) list {
	if l == nil {
		l = mzero()
	}
	if _, ok := l.(uint); ok {
		return unit{}
	}
	ll := l.([2]interface{})
	f := ll[1].(func() list)
	return f()
}

func end(l list) bool {
	if l == nil {
		l = mzero()
	}
	_, ok := l.(unit)
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

func listmapM(f func(interface{}), l list) {
	adapter := func(i interface{}) interface{} {
		f(i)
		return nil
	}
	listmap(adapter, l)
}

func main() {
	f := func(i interface{}) {
		fmt.Println(i)
	}

	slowValue := func() interface{} {
		time.Sleep(5 * time.Second)
		return 1
	}

	l := cons(3, cons(2, consf(slowValue, mreturn(0))))
	listmapM(f, l)
}
