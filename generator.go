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

// Listƒ{{.}} is an ordered collection of {{.}}
type Listƒ{{.}} interface {
	// Size returns the size of the collection.
	Size() int

	// Contains checks if all elements in the specified collection are contained in this collection.
	Contains(element {{.}}) bool

	// Filter returns a new Listƒ{{.}} with only elements matching the given predicate.
	Filter(predicate Predicateƒ{{.}}) Listƒ{{.}}

	// First returns first element matching the given predicate.
	First(predicate Predicateƒ{{.}}) *{{.}}
	
	// All returns true if all elements match the given predicate.
	All(predicate Predicateƒ{{.}}) bool

	// Any returns true if any elements match the given predicate.
	Any(predicate Predicateƒ{{.}}) bool
}

// Predicateƒ{{.}} check a condition on {{.}}
type Predicateƒ{{.}} func(it {{.}}) bool


// NewListƒ{{.}} is constructor for a Listƒ{{.}}
func NewListƒ{{.}}(elements ... {{.}}) Listƒ{{.}} {
	return &listƒ{{.}}{ 
		elements: elements,
	}
}

type listƒ{{.}} struct {
		elements []{{.}}	
}

func (l *listƒ{{.}}) Size() int {
	return len(l.elements)
}

func (l *listƒ{{.}}) Contains(element {{.}}) bool {
	for _, e := range l.elements {
		if e == element {
			return true
		}
	}
	return false
}

func (l *listƒ{{.}}) Filter(predicate Predicateƒ{{.}}) Listƒ{{.}} {
	filtered := Listƒ{{.}}{}
	for _, e := range l.elements {
		if predicate(e) {
			filtered.elements = append(filtered.elements, e)
		}
	}
	return filtered
}

func (l *listƒ{{.}}) First(predicate Predicateƒ{{.}}) *{{.}} {
	filtered := Listƒ{{.}}{}
	for _, e := range l.elements {
		if predicate(e) {
			return &e
		}
	}
	return nil
}

func (l *listƒ{{.}}) All(predicate Predicateƒ{{.}}) bool {
	for _, e := range l.elements {
		if !predicate(e) {
			return false
		}
	}
	return true
}

func (l *listƒ{{.}}) Any(predicate Predicateƒ{{.}}) bool {
	for _, e := range l.elements {
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
