package example

import (
	"fmt"
	"strings"
)

type Foo struct {
	val string
}

type Bar struct {
	val string
}

func demo() {
	list := ListƒFoo{Foo{val: "one"}, Foo{val: "two"}, Foo{val: "three"}}
	maybeTwo := list.
		Filter(func(it Foo) bool { return len(it.val) == 3 }).
		Find(func(it Foo) bool { return it.val[1] == 'w' })
	fmt.Println(maybeTwo.
		Filter(func(it Foo) bool { return it.val[2] == 'o' }).
		MapToBar(func(it Foo) Bar { return Bar{val: strings.ToUpper(it.val)} }).
		Get())
}

// Sorry, this one you have to create manually
func (o FooɁ) MapToBar(transformation func(Foo) Bar) BarɁ {
	if o.IsEmpty() {
		return NoneBar
	} else {
		return SomeBar(transformation(o.Get()))
	}
}
