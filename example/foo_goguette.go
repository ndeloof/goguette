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

func SomeFoo(it Foo) FooɁ {
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
func (l ListƒFoo) Contains(elements ...Foo) bool {
	for _, e := range l {
		for i, el := range elements {
			if e == el {
				k := len(elements)
				if k == 1 {
					return true
				}
				elements[k-1], elements[i] = elements[i], elements[k-1]
				elements = elements[:k-1]
			}
		}
	}
	return false
}

// Filter returns a new ListƒFoo with only elements matching the given predicates.
func (l ListƒFoo) Filter(predicates ...PredicateƒFoo) ListƒFoo {
	filtered := ListƒFoo{}
LOOP:
	for _, e := range l {
		for _, predicate := range predicates {
			if !predicate(e) {
				continue LOOP
			}
		}
		filtered = append(filtered, e)
	}
	return filtered
}

// First returns first element matching the given predicates.
func (l ListƒFoo) Find(predicates ...PredicateƒFoo) FooɁ {
LOOP:
	for _, e := range l {
		for _, predicate := range predicates {
			if !predicate(e) {
				continue LOOP
			}
		}
		return SomeFoo(e)
	}
	return NoneFoo
}

// All returns true if all elements match the given predicates.
func (l ListƒFoo) All(predicates ...PredicateƒFoo) bool {
	for _, e := range l {
		for _, predicate := range predicates {
			if !predicate(e) {
				return false
			}
		}
	}
	return true
}

// Any returns true if any elements match the given predicates.
func (l ListƒFoo) Any(predicates ...PredicateƒFoo) bool {
LOOP:
	for _, e := range l {
		for _, predicate := range predicates {
			if !predicate(e) {
				continue LOOP
			}
		}
		return true
	}
	return false
}

// Count returns the number of elements matching the given predicates.
func (l ListƒFoo) Count(predicates ...PredicateƒFoo) int {
	count := 0
LOOP:
	for _, e := range l {
		for _, predicate := range predicates {
			if !predicate(e) {
				continue LOOP
			}
		}
		count += 1
	}
	return count
}

// Distinct returns a list containing only distinct elements from the given collection.
func (l ListƒFoo) Distinct(predicates ...PredicateƒFoo) ListƒFoo {
	uniq := map[Foo]struct{}{}
	for _, e := range l {
		uniq[e] = struct{}{}
	}
	var distinct ListƒFoo
	for e := range uniq {
		distinct = append(distinct, e)
	}
	return distinct
}

func (l ListƒFoo) Intersect(other ListƒFoo) ListƒFoo {
	var intersect ListƒFoo
	for _, e := range l {
		if other.Contains(e) {
			intersect = append(intersect, e)
		}
	}
	return intersect
}

func (l ListƒFoo) MinBy(comparator func(a, b Foo) int) FooɁ {
	if len(l) == 0 {
		return NoneFoo
	}
	min := l[0]
	for _, e := range l[1:] {
		if comparator(e, min) < 0 {
			min = e
		}
	}
	return SomeFoo(min)
}

func (l ListƒFoo) MaxBy(comparator func(a, b Foo) int) FooɁ {
	if len(l) == 0 {
		return NoneFoo
	}
	max := l[0]
	for _, e := range l[1:] {
		if comparator(e, max) > 0 {
			max = e
		}
	}
	return SomeFoo(max)
}

func (l ListƒFoo) Partition(predicate PredicateƒFoo) (ListƒFoo, ListƒFoo) {
	var right, left ListƒFoo
	for _, e := range l {
		if predicate(e) {
			right = append(right, e)
		} else {
			left = append(left, e)
		}

	}
	return right, left
}
