package list

import "errors"

var ErrOutOfRange = errors.New("out of range")

type unit struct{}

type listReturn struct {
	val     interface{}
	unitVal *unit
}

type list interface{}

func Mzero() list {
	return unit{}
}

func Return(i interface{}) list {
	return Returnf(func() interface{} { return i })
}

func Cons(i interface{}, l list) list {
	return Consf(func() interface{} { return i }, l)
}

func Consf(f func() interface{}, l list) list {
	if l == nil {
		l = Mzero()
	}
	return [2]interface{}{f, func() list { return l }}
}

func Returnf(f func() interface{}) list {
	return Consf(f, Mzero())
}

func Head(l list) interface{} {
	if l == nil {
		l = Mzero()
	}
	if _, ok := l.(unit); ok {
		return unit{}
	}
	lf := l.([2]interface{})[0].(func() interface{})
	return lf()
}

func Tail(l list) list {
	if l == nil {
		l = Mzero()
	}
	if _, ok := l.(uint); ok {
		return unit{}
	}
	ll := l.([2]interface{})
	f := ll[1].(func() list)
	return f()
}

func HdTail(l list) (interface{}, list) {
	return Head(l), Tail(l)
}

func IsEmpty(l list) bool {
	if l == nil {
		l = Mzero()
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

func Map(f func(interface{}) interface{}, l list) list {
	if IsEmpty(l) {
		return Mzero()
	}
	elem := l.([2]interface{})
	valFunc := elem[0].(func() interface{})
	next := elem[1].(func() list)
	mapperFunc := func() interface{} {
		return f(valFunc())
	}
	return Consf(mapperFunc, Map(f, next()))
}

func MapM(f func(interface{}), l list) {
	adapter := func(i interface{}) interface{} {
		f(i)
		return nil
	}
	Seq(Map(adapter, l))
}

func Seq(l list) {
	for !IsEmpty(l) {
		Head(l)
		l = Tail(l)
	}
}

func Foldl(f func(interface{}, interface{}) interface{}, val interface{}, l list) interface{} {
	if IsEmpty(l) {
		return val
	}
	hd, tl := HdTail(l)
	return Foldl(f, f(val, hd), tl)
}

func Foldl1(f func(interface{}, interface{}) interface{}, l list) interface{} {
	hd, tl := HdTail(l)
	return Foldl(f, hd, tl)
}

func Index(idx uint, l list) interface{} {
	for cur := uint(0); cur < idx; cur++ {
		if IsEmpty(l) {
			return Mzero()
		}
		l = Tail(l)
	}
	if IsEmpty(l) {
		return Mzero()
	}
	return Head(l)
}

func Reverse(l list) list {
	foldFunc := func(carry, elem interface{}) interface{} {
		return Cons(elem, carry)
	}
	return Foldl(foldFunc, Mzero(), l).(list)
}
