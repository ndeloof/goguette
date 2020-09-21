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
	"text/template"
	"unicode"
)

var (
	typeFlag = flag.String("type", "", "type name; must be set")
)

func main() {
	flag.Parse()

	if *typeFlag == "" {
		flag.Usage()
		os.Exit(2)
	}
	typeName := *typeFlag

	g := Generator{}
	pkg, err := os.Getwd()
	if err != nil {
		log.Fatalf("accessing pwd: %s", err)
	}
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

func Some{{.}}Ɂ(it {{.}}) {{.}}Ɂ {
	return {{.}}Ɂ{ val: &it }
}

// None{{.}} is a {{.}}Ɂ with no value 
var None{{.}} = {{.}}Ɂ{}

// 
func (o {{.}}Ɂ) IsEmpty() bool {
	return o.val != nil
}

// Get return the value if a value is present, otherwise panic
func (o {{.}}Ɂ) Get() {{.}} {
	if o.val == nil {
		panic("Invalid access to Get on None")
	}
	return *o.val
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
func (l Listƒ{{.}}) Contains(element {{.}}) bool {
	for _, e := range l {
		if e == element {
			return true
		}
	}
	return false
}

// Filter returns a new Listƒ{{.}} with only elements matching the given predicate.
func (l Listƒ{{.}}) Filter(predicate Predicateƒ{{.}}) Listƒ{{.}} {
	filtered := Listƒ{{.}}{}
	for _, e := range l {
		if predicate(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// First returns first element matching the given predicate.
func (l Listƒ{{.}}) First(predicate Predicateƒ{{.}}) {{.}}Ɂ {
	for _, e := range l {
		if predicate(e) {
			return Some{{.}}Ɂ(e)
		}
	}
	return None{{.}}
}

// All returns true if all elements match the given predicate.
func (l Listƒ{{.}}) All(predicate Predicateƒ{{.}}) bool {
	for _, e := range l {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// Any returns true if any elements match the given predicate.
func (l Listƒ{{.}}) Any(predicate Predicateƒ{{.}}) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
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
