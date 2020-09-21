package goguette

//go generate goguette -type=string

type Listƒstring struct {
	elements []string
}

func (l Listƒstring) Size() int {
	return len(l.elements)
}

func (l Listƒstring) Contains(elements string) bool {
	for _, e := range l.elements {
		if e == el {
			return true
		}
	}
	return false
}

type Predicateƒstring func(it string) bool

func (l Listƒstring) All(predicate Predicateƒstring) Listƒstring {
	filtered := Listƒstring{}
	for _, e := range l.elements {
		if predicate(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
