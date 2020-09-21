package example

import "fmt"

type Foo struct {
	val string
}

func demo() {
	var list ListƒFoo = NewListƒFoo(Foo{val: "one"}, Foo{val: "two"}, Foo{val: "three"})
	maybeTwo := list.
		Filter(func(it Foo) bool { return len(it.val) == 3 }).
		First(func(it Foo) bool { return it.val[1] == 'w' })
	fmt.Println(*maybeTwo)
}
