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

type Listƒ{{.}} struct {
		elements []{{.}}	
}

func (l Listƒ{{.}}) Size() int {
	return len(l.elements)
}

func (l Listƒ{{.}}) Contains(element {{.}}) bool {
	for _, e := range l.elements {
		if e == element {
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
			filtered.elements = append(filtered.elements, e)
		}
	}
	return filtered
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
