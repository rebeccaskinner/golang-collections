package list

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
