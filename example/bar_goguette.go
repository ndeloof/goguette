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

func SomeBar(it Bar) BarɁ {
	return BarɁ{val: &it}
}

// NoneBar is a BarɁ with no value
var NoneBar = BarɁ{}

// IsEmpty return true if value is not set, i.e. this is None
func (o BarɁ) IsEmpty() bool {
	return o.val != nil
}

func TryBarɁ(fn func() (Bar, error)) BarɁ {
	val, err := fn()
	if error != nil {
		return NoneBar
	}
	return SomeBar(val)
}

// Get return the value if a value is present, otherwise panic
func (o BarɁ) Get() Bar {
	return *o.val
}

// OrElse return the value if present, otherwise return an error.
func (o BarɁ) OrError(message string, args ...interface{}) (Bar, error) {
	if o.val == nil {
		return Bar{}, fmt.Errorf(message, args...)
	}
	return *o.val, nil
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
func (l ListƒBar) Contains(elements ...Bar) bool {
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

// Filter returns a new ListƒBar with only elements matching the given predicates.
func (l ListƒBar) Filter(predicates ...PredicateƒBar) ListƒBar {
	filtered := ListƒBar{}
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
func (l ListƒBar) Find(predicates ...PredicateƒBar) BarɁ {
LOOP:
	for _, e := range l {
		for _, predicate := range predicates {
			if !predicate(e) {
				continue LOOP
			}
		}
		return SomeBar(e)
	}
	return NoneBar
}

// All returns true if all elements match the given predicates.
func (l ListƒBar) All(predicates ...PredicateƒBar) bool {
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
func (l ListƒBar) Any(predicates ...PredicateƒBar) bool {
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
func (l ListƒBar) Count(predicates ...PredicateƒBar) int {
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
func (l ListƒBar) Distinct(predicates ...PredicateƒBar) ListƒBar {
	uniq := map[Bar]struct{}{}
	for _, e := range l {
		uniq[e] = struct{}{}
	}
	var distinct ListƒBar
	for e := range uniq {
		distinct = append(distinct, e)
	}
	return distinct
}

func (l ListƒBar) Intersect(other ListƒBar) ListƒBar {
	var intersect ListƒBar
	for _, e := range l {
		if other.Contains(e) {
			intersect = append(intersect, e)
		}
	}
	return intersect
}

func (l ListƒBar) MinBy(comparator func(a, b Bar) int) BarɁ {
	if len(l) == 0 {
		return NoneBar
	}
	min := l[0]
	for _, e := range l[1:] {
		if comparator(e, min) < 0 {
			min = e
		}
	}
	return SomeBar(min)
}

func (l ListƒBar) MaxBy(comparator func(a, b Bar) int) BarɁ {
	if len(l) == 0 {
		return NoneBar
	}
	max := l[0]
	for _, e := range l[1:] {
		if comparator(e, max) > 0 {
			max = e
		}
	}
	return SomeBar(max)
}

func (l ListƒBar) Partition(predicate PredicateƒBar) (ListƒBar, ListƒBar) {
	var right, left ListƒBar
	for _, e := range l {
		if predicate(e) {
			right = append(right, e)
		} else {
			left = append(left, e)
		}

	}
	return right, left
}
