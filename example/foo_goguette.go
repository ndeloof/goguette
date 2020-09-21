// DO NOT EDIT, GENERATED CODE

package example

//go generate goguette -type=Foo

type ListƒFoo struct {
	elements []Foo
}

func (l ListƒFoo) Size() int {
	return len(l.elements)
}

func (l ListƒFoo) Contains(element Foo) bool {
	for _, e := range l.elements {
		if e == element {
			return true
		}
	}
	return false
}

type PredicateƒFoo func(it Foo) bool

func (l ListƒFoo) All(predicate PredicateƒFoo) ListƒFoo {
	filtered := ListƒFoo{}
	for _, e := range l.elements {
		if predicate(e) {
			filtered.elements = append(filtered.elements, e)
		}
	}
	return filtered
}
