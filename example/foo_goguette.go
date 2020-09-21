// DO NOT EDIT, GENERATED CODE

package example

//go generate goguette -type=Foo

// PredicateƒFoo check a condition on Foo
type PredicateƒFoo func(it Foo) bool

// --- Some or None

// FooɁ hold an optional Foo value
type FooɁ struct {
	val *Foo
}

func SomeFooɁ(it Foo) FooɁ {
	return FooɁ{val: &it}
}

// NoneFoo is a FooɁ with no value
var NoneFoo = FooɁ{}

//
func (o FooɁ) IsEmpty() bool {
	return o.val != nil
}

// Get return the value if a value is present, otherwise panic
func (o FooɁ) Get() Foo {
	if o.val == nil {
		panic("Invalid access to Get on None")
	}
	return *o.val
}

// OrElse return the value if present, otherwise return other.
func (o FooɁ) OrElse(other Foo) Foo {
	if o.val == nil {
		return other
	}
	return *o.val
}

// Filter return an FooɁ describing the value is it matches the predicate, otherwise return None
func (o FooɁ) Filter(predicate PredicateƒFoo) FooɁ {
	if o.val == nil {
		return NoneFoo
	}
	if predicate(*o.val) {
		return o
	}
	return NoneFoo
}

// --- List

// ListƒFoo is an ordered collection of Foo
type ListƒFoo []Foo

// Size returns the size of the collection.
func (l ListƒFoo) Size() int {
	return len(l)
}

// Contains checks if all elements in the specified collection are contained in this collection.
func (l ListƒFoo) Contains(element Foo) bool {
	for _, e := range l {
		if e == element {
			return true
		}
	}
	return false
}

// Filter returns a new ListƒFoo with only elements matching the given predicate.
func (l ListƒFoo) Filter(predicate PredicateƒFoo) ListƒFoo {
	filtered := ListƒFoo{}
	for _, e := range l {
		if predicate(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// First returns first element matching the given predicate.
func (l ListƒFoo) First(predicate PredicateƒFoo) FooɁ {
	for _, e := range l {
		if predicate(e) {
			return SomeFooɁ(e)
		}
	}
	return NoneFoo
}

// All returns true if all elements match the given predicate.
func (l ListƒFoo) All(predicate PredicateƒFoo) bool {
	for _, e := range l {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// Any returns true if any elements match the given predicate.
func (l ListƒFoo) Any(predicate PredicateƒFoo) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}
