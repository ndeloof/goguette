// DO NOT EDIT, GENERATED CODE

package example

//go generate goguette -type=Foo

// ListƒFoo is an ordered collection of Foo
type ListƒFoo interface {
	// Size returns the size of the collection.
	Size() int

	// Contains checks if all elements in the specified collection are contained in this collection.
	Contains(element Foo) bool

	// Filter returns a new ListƒFoo with only elements matching the given predicate.
	Filter(predicate PredicateƒFoo) ListƒFoo

	// First returns first element matching the given predicate.
	First(predicate PredicateƒFoo) *Foo

	// All returns true if all elements match the given predicate.
	All(predicate PredicateƒFoo) bool

	// Any returns true if any elements match the given predicate.
	Any(predicate PredicateƒFoo) bool
}

// PredicateƒFoo check a condition on Foo
type PredicateƒFoo func(it Foo) bool

// NewListƒFoo is constructor for a ListƒFoo
func NewListƒFoo(elements ...Foo) ListƒFoo {
	return &listƒFoo{
		elements: elements,
	}
}

type listƒFoo struct {
	elements []Foo
}

func (l *listƒFoo) Size() int {
	return len(l.elements)
}

func (l *listƒFoo) Contains(element Foo) bool {
	for _, e := range l.elements {
		if e == element {
			return true
		}
	}
	return false
}

func (l *listƒFoo) Filter(predicate PredicateƒFoo) ListƒFoo {
	filtered := ListƒFoo{}
	for _, e := range l.elements {
		if predicate(e) {
			filtered.elements = append(filtered.elements, e)
		}
	}
	return filtered
}

func (l *listƒFoo) First(predicate PredicateƒFoo) *Foo {
	filtered := ListƒFoo{}
	for _, e := range l.elements {
		if predicate(e) {
			return &e
		}
	}
	return nil
}

func (l *listƒFoo) All(predicate PredicateƒFoo) bool {
	for _, e := range l.elements {
		if !predicate(e) {
			return false
		}
	}
	return true
}

func (l *listƒFoo) Any(predicate PredicateƒFoo) bool {
	for _, e := range l.elements {
		if predicate(e) {
			return true
		}
	}
	return false
}
