// DO NOT EDIT, GENERATED CODE

package example

//go generate goguette -type=Bar

// PredicateƒBar check a condition on Bar
type PredicateƒBar func(it Bar) bool

// --- Some or None

// BarɁ hold an optional Bar value
type BarɁ struct {
	val *Bar
}

func SomeBarɁ(it Bar) BarɁ {
	return BarɁ{val: &it}
}

// NoneBar is a BarɁ with no value
var NoneBar = BarɁ{}

//
func (o BarɁ) IsEmpty() bool {
	return o.val != nil
}

// Get return the value if a value is present, otherwise panic
func (o BarɁ) Get() Bar {
	if o.val == nil {
		panic("Invalid access to Get on None")
	}
	return *o.val
}

// OrElse return the value if present, otherwise return other.
func (o BarɁ) OrElse(other Bar) Bar {
	if o.val == nil {
		return other
	}
	return *o.val
}

// Filter return an BarɁ describing the value is it matches the predicate, otherwise return None
func (o BarɁ) Filter(predicate PredicateƒBar) BarɁ {
	if o.val == nil {
		return NoneBar
	}
	if predicate(*o.val) {
		return o
	}
	return NoneBar
}

// --- List

// ListƒBar is an ordered collection of Bar
type ListƒBar []Bar

// Size returns the size of the collection.
func (l ListƒBar) Size() int {
	return len(l)
}

// Contains checks if all elements in the specified collection are contained in this collection.
func (l ListƒBar) Contains(element Bar) bool {
	for _, e := range l {
		if e == element {
			return true
		}
	}
	return false
}

// Filter returns a new ListƒBar with only elements matching the given predicate.
func (l ListƒBar) Filter(predicate PredicateƒBar) ListƒBar {
	filtered := ListƒBar{}
	for _, e := range l {
		if predicate(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// First returns first element matching the given predicate.
func (l ListƒBar) First(predicate PredicateƒBar) BarɁ {
	for _, e := range l {
		if predicate(e) {
			return SomeBarɁ(e)
		}
	}
	return NoneBar
}

// All returns true if all elements match the given predicate.
func (l ListƒBar) All(predicate PredicateƒBar) bool {
	for _, e := range l {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// Any returns true if any elements match the given predicate.
func (l ListƒBar) Any(predicate PredicateƒBar) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}
