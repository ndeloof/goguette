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
)

var (
	typeName = flag.String("type", "", "type name; must be set")
)

func main() {
	flag.Parse()

	g := Generator{}
	pkg, err := os.Getwd()
	if err != nil {
		log.Fatalf("accessing pwd: %s", err)
	}
	g.Printf("package %s", filepath.Base(pkg))
	g.Println()
	g.Printf("//go generate goguette -type=%s", *typeName)
	g.Println()

	tmpl, err := template.New("list").Parse(
		`type Listƒ{{.}} struct {
		elements []{{.}}	
}

func (l Listƒ{{.}}) Size() int {
	return len(l.elements)
}

func (l Listƒ{{.}}) Contains(elements {{.}}) bool {
	for _, e := range l.elements {
		if e == el {
			return true
		}
	}
	return false
}

type Predicateƒ{{.}} func(it {{.}}) bool

func (l Listƒ{{.}}) All(predicate Predicateƒ{{.}}) Listƒ{{.}} {
	filtered := Listƒ{{.}}{}
	for _, e := range l.elements {
		if predicate(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

`)
	if err != nil {
		log.Fatalf("error in template: %s", err)
	}
	err = tmpl.Execute(&g.buf, *typeName)
	if err != nil {
		log.Fatalf("error applying template: %s", err)
	}

	out := fmt.Sprintf("%s_goguette.go", *typeName)
	src := g.format()

	err = ioutil.WriteFile(out, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
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
