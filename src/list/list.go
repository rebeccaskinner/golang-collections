package list

import "errors"

var ErrOutOfRange = errors.New("out of range")

type unit struct{}

type listReturn struct {
	val     interface{}
	unitVal *unit
}

type List interface{}

func Mzero() List {
	return unit{}
}

func Return(i interface{}) List {
	return Returnf(func() interface{} { return i })
}

func Cons(i interface{}, l List) List {
	return Consf(func() interface{} { return i }, l)
}

func Consf(f func() interface{}, l List) List {
	if l == nil {
		l = Mzero()
	}
	return [2]interface{}{f, func() List { return l }}
}

func Returnf(f func() interface{}) List {
	return Consf(f, Mzero())
}

func Head(l List) interface{} {
	if l == nil {
		l = Mzero()
	}
	if _, ok := l.(unit); ok {
		return unit{}
	}
	lf := l.([2]interface{})[0].(func() interface{})
	return lf()
}

func Tail(l List) List {
	if l == nil {
		l = Mzero()
	}
	if _, ok := l.(uint); ok {
		return unit{}
	}
	ll := l.([2]interface{})
	f := ll[1].(func() List)
	return f()
}

func HdTail(l List) (interface{}, List) {
	return Head(l), Tail(l)
}

func IsEmpty(l List) bool {
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

func Map(f func(interface{}) interface{}, l List) List {
	if IsEmpty(l) {
		return Mzero()
	}
	elem := l.([2]interface{})
	valFunc := elem[0].(func() interface{})
	next := elem[1].(func() List)
	mapperFunc := func() interface{} {
		return f(valFunc())
	}
	return Consf(mapperFunc, Map(f, next()))
}

func MapM(f func(interface{}), l List) {
	adapter := func(i interface{}) interface{} {
		f(i)
		return nil
	}
	Seq(Map(adapter, l))
}

func Seq(l List) {
	for !IsEmpty(l) {
		Head(l)
		l = Tail(l)
	}
}

func Foldl(f func(interface{}, interface{}) interface{}, val interface{}, l List) interface{} {
	if IsEmpty(l) {
		return val
	}
	hd, tl := HdTail(l)
	return Foldl(f, f(val, hd), tl)
}

func Foldr(f func(interface{}, interface{}) interface{}, val interface{}, l List) interface{} {
	if IsEmpty(l) {
		return val
	}
	hd, tl := HdTail(l)
	return f(Foldr(f, val, tl), hd)
}

func Foldl1(f func(interface{}, interface{}) interface{}, l List) interface{} {
	hd, tl := HdTail(l)
	return Foldl(f, hd, tl)
}

func Index(idx uint, l List) interface{} {
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

func Reverse(l List) List {
	foldFunc := func(carry, elem interface{}) interface{} {
		return Cons(elem, carry)
	}
	return Foldl(foldFunc, Mzero(), l).(List)
}

func Append(i interface{}, l List) List {
	return Reverse(Cons(i, Reverse(l)))
}

func Concat(back, front List) List {
	foldFunc := func(carry, elem interface{}) interface{} {
		return Cons(elem, carry)
	}
	return Foldr(foldFunc, front, back).(List)
}

func New(elems ...interface{}) List {
	l := Mzero()
	for _, elem := range elems {
		l = Cons(elem, l)
	}
	return Reverse(l)
}
