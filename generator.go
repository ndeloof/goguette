package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

var (
	typeFlag = flag.String("type", "", "comma-separated list of type names; must be set")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of goguette:\n")
	fmt.Fprintf(os.Stderr, "\tgoguette -type T\n")
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "\thttp://github.com/ndeloof/goguette\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("goguette: ")
	flag.Usage = Usage
	flag.Parse()

	if len(*typeFlag) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	typeNames := strings.Split(*typeFlag, ",")
	pkg, err := os.Getwd()
	if err != nil {
		log.Fatalf("accessing pwd: %s", err)
	}

	for _, typeName := range typeNames {
		g := Generator{}
		g.Printf("// DO NOT EDIT, GENERATED CODE")
		g.Println()
		g.Printf("package %s", filepath.Base(pkg))
		g.Println()
		g.Printf("//go generate goguette -type=%s", typeName)
		g.Println()

		tmpl, err := template.New("list").Parse(
			`

// Predicateƒ{{.}} check a condition on {{.}}
type Predicateƒ{{.}} func(it {{.}}) bool


// --- Some or None

// {{.}}Ɂ hold an optional {{.}} value
type {{.}}Ɂ struct {
	val *{{.}}
}

func Some{{.}}(it {{.}}) {{.}}Ɂ {
	return {{.}}Ɂ{ val: &it }
}

// None{{.}} is a {{.}}Ɂ with no value 
var None{{.}} = {{.}}Ɂ{}

// IsEmpty return true if value is not set, i.e. this is None 
func (o {{.}}Ɂ) IsEmpty() bool {
	return o.val != nil
}

func Try{{.}}Ɂ (fn func() ({{.}}, error)) {{.}}Ɂ {
	val, err := fn()
	if error != nil {
		return None{{.}}
	}
	return Some{{.}}(val)
}

// Get return the value if a value is present, otherwise panic
func (o {{.}}Ɂ) Get() ({{.}}) {
	return *o.val
}


// OrElse return the value if present, otherwise return an error.
func (o {{.}}Ɂ) OrError(message string, args ... interface{}) ({{.}}, error) {
	if o.val == nil {
		return {{.}}{}, fmt.Errorf(message, args...)
	}
	return *o.val, nil
}

// OrElse return the value if present, otherwise return other.
func (o {{.}}Ɂ) OrElse(other {{.}}) {{.}} {
	if o.val == nil {
		return other
	}
	return *o.val
}

// Filter return an {{.}}Ɂ describing the value is it matches the predicate, otherwise return None
func (o {{.}}Ɂ) Filter(predicate Predicateƒ{{.}}) {{.}}Ɂ {
	if o.val == nil {
		return None{{.}}
	}
	if predicate(*o.val) {
		return o
	}
	return None{{.}}
}

// --- List

// Listƒ{{.}} is an ordered collection of {{.}}
type Listƒ{{.}} []{{.}}

// Size returns the size of the collection.
func (l Listƒ{{.}}) Size() int {
	return len(l)
}

// Contains checks if all elements in the specified collection are contained in this collection.
func (l Listƒ{{.}}) Contains(elements ... {{.}}) bool {
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

// Filter returns a new Listƒ{{.}} with only elements matching the given predicates.
func (l Listƒ{{.}}) Filter(predicates ... Predicateƒ{{.}}) Listƒ{{.}} {
	filtered := Listƒ{{.}}{}
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
func (l Listƒ{{.}}) Find(predicates ... Predicateƒ{{.}}) {{.}}Ɂ {
	LOOP:
	for _, e := range l {
		for _, predicate := range predicates {
			if !predicate(e) {
				continue LOOP
			}
		}
		return Some{{.}}(e)
	}
	return None{{.}}
}

// All returns true if all elements match the given predicates.
func (l Listƒ{{.}}) All(predicates ... Predicateƒ{{.}}) bool {
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
func (l Listƒ{{.}}) Any(predicates ... Predicateƒ{{.}}) bool {
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
func (l Listƒ{{.}}) Count(predicates ... Predicateƒ{{.}}) int {
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
func (l Listƒ{{.}}) Distinct(predicates ... Predicateƒ{{.}}) Listƒ{{.}} {
	uniq := map[{{.}}]struct{}{}
	for _, e := range l {
		uniq[e] = struct{}{}
	}
	var distinct Listƒ{{.}}
	for e := range uniq {
		distinct = append(distinct, e)
	}
	return distinct
}

func (l Listƒ{{.}}) Intersect(other Listƒ{{.}}) Listƒ{{.}} {
	var intersect Listƒ{{.}}
	for _, e := range l {
		if other.Contains(e) {
			intersect = append(intersect, e)
		}
	}
	return intersect
}

func (l Listƒ{{.}}) MinBy(comparator func(a, b {{.}}) int) {{.}}Ɂ {
	if len(l) == 0 {
		return None{{.}}
	}
	min := l[0]
	for _, e := range l[1:] {
		if comparator(e, min) < 0 {
			min = e
		}
	}
	return Some{{.}}(min)
}

func (l Listƒ{{.}}) MaxBy(comparator func(a, b {{.}}) int) {{.}}Ɂ {
	if len(l) == 0 {
		return None{{.}}
	}
	max := l[0]
	for _, e := range l[1:] {
		if comparator(e, max) > 0 {
			max = e
		}
	}
	return Some{{.}}(max)
}

func (l Listƒ{{.}}) Partition(predicate Predicateƒ{{.}}) (Listƒ{{.}}, Listƒ{{.}}) {
	var right, left Listƒ{{.}}
	for _, e := range l {
		if (predicate(e)) {
			right = append(right, e)
		} else {
			left = append(left, e)
		}

	}
	return right, left
}


`)
		if err != nil {
			log.Fatalf("error in template: %s", err)
		}
		err = tmpl.Execute(&g.buf, typeName)
		if err != nil {
			log.Fatalf("error applying template: %s", err)
		}

		out := fmt.Sprintf("%s_goguette.go", lower(typeName))
		src := g.format()

		// fmt.Println(string(src))
		err = ioutil.WriteFile(out, src, 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	}
}

func lower(s string) string {
	for i, v := range s {
		return string(unicode.ToLower(v)) + s[i+1:]
	}
	return s
}

type Generator struct {
	buf bytes.Buffer
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
	fmt.Fprintln(&g.buf)
}

func (g *Generator) Println() {
	fmt.Fprintln(&g.buf)
}

func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}
