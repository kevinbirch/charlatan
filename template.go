package main

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"
)

const charlatanTemplate = `// generated by "{{.CommandLine}}".  DO NOT EDIT.

package {{.PackageName}}

import (
	"reflect"
	"testing"
{{range .Imports}}	{{if .Alias}}{{.Alias}} {{end}}{{.Path}}
{{end}}
)
{{range .Interfaces}}{{range .Methods}}
// {{.Name}}Invocation represents a single call of Fake{{.Interface}}.{{.Name}}
type {{.Name}}Invocation struct {
{{if .Parameters}}	Parameters struct {
{{range .Parameters}}	{{.FieldFormat}}
{{end}}
	}{{end}}
{{if .Results}}	Results struct {
{{range .Results}}	{{.FieldFormat}}
{{end}}
	}{{end}}
}
{{end}}

/*
Fake{{.Name}} is a mock implementation of {{.Name}} for testing.
{{if .Methods}}{{with $m := index .Methods 0}}Use it in your tests as in this example:

	package example

	func TestWith{{$m.Interface}}(t *testing.T) {
		f := &{{$.PackageName}}.Fake{{$m.Interface}}{
			{{$m.Name}}Hook: func({{$m.ParametersDeclaration}}) ({{$m.ResultsDeclaration}}) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			}
		}

		// test code goes here ...

		// assert state of Fake{{.Name}} ...
		f.Assert{{$m.Name}}CalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a painc if any
unexpected calls are made to Fake{{.Name}}.
{{end}}{{end}}*/
type Fake{{.Name}} struct {
{{range .Methods}} {{.Name}}Hook func({{.ParametersSignature}}) ({{.ResultsSignature}})
{{end}}
{{range .Methods}} {{.Name}}Calls []*{{.Name}}Invocation
{{end}}}

{{range $m := .Methods}}
{{with $f := gensym}}func ({{$f}} *Fake{{$m.Interface}}) {{$m.Name}}({{$m.ParametersDeclaration}}) ({{$m.ResultsDeclaration}}) {
	invocation := new({{$m.Name}}Invocation)

{{if $m.Parameters}}{{range $m.Parameters}} invocation.Parameters.{{.TitleCase}} = {{.Name}}
{{end}}{{end}}
{{if $m.Results}} {{$m.ResultsReference}} = {{$f}}.{{$m.Name}}Hook({{$m.ParametersReference}})
{{else}} {{$f}}.{{$m.Name}}Hook({{$m.ParametersReference}})
{{end}}
{{if $m.Results}}{{range $m.Results}}invocation.Results.{{.TitleCase}} = {{.Name}}
{{end}}{{end}}
	{{$f}}.{{$m.Name}}Calls = append({{$f}}.{{$m.Name}}Calls, invocation)

	return
}{{end}}

// {{.Name}}Called returns true if Fake{{.Interface}}.{{.Name}} was called
func (f *Fake{{.Interface}}) {{.Name}}Called() bool {
	return len(f.{{.Name}}Calls) != 0
}

// Assert{{.Name}}Called calls t.Error if Fake{{.Interface}}.{{.Name}} was not called
func (f *Fake{{.Interface}}) Assert{{.Name}}Called(t *testing.T) {
	t.Helper()
	if len(f.{{.Name}}Calls) == 0 {
		t.Error("Fake{{.Interface}}.{{.Name}} not called, expected at least one")
	}
}

// {{.Name}}NotCalled returns true if Fake{{.Interface}}.{{.Name}} was not called
func (f *Fake{{.Interface}}) {{.Name}}NotCalled() bool {
	return len(f.{{.Name}}Calls) == 0
}

// Assert{{.Name}}NotCalled calls t.Error if Fake{{.Interface}}.{{.Name}} was called
func (f *Fake{{.Interface}}) Assert{{.Name}}NotCalled(t *testing.T) {
	t.Helper()
	if len(f.{{.Name}}Calls) != 0 {
		t.Error("Fake{{.Interface}}.{{.Name}} called, expected none")
	}
}

// {{.Name}}CalledOnce returns true if Fake{{.Interface}}.{{.Name}} was called exactly once
func (f *Fake{{.Interface}}) {{.Name}}CalledOnce() bool {
	return len(f.{{.Name}}Calls) == 1
}

// Assert{{.Name}}CalledOnce calls t.Error if Fake{{.Interface}}.{{.Name}} was not called exactly once
func (f *Fake{{.Interface}}) Assert{{.Name}}CalledOnce(t *testing.T) {
	t.Helper()
	if len(f.{{.Name}}Calls) != 1 {
		t.Errorf("Fake{{.Interface}}.{{.Name}} called %d times, expected 1", len(f.{{.Name}}Calls))
	}
}

// {{.Name}}CalledN returns true if Fake{{.Interface}}.{{.Name}} was called at least n times
func (f *Fake{{.Interface}}) {{.Name}}CalledN(n int) bool {
	return len(f.{{.Name}}Calls) >= n
}

// Assert{{.Name}}CalledN calls t.Error if Fake{{.Interface}}.{{.Name}} was called less than n times
func (f *Fake{{.Interface}}) Assert{{.Name}}CalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.{{.Name}}Calls) < n {
		t.Errorf("Fake{{.Interface}}.{{.Name}} called %d times, expected >= %d", len(f.{{.Name}}Calls), n)
	}
}

{{if .Parameters}}// {{.Name}}CalledWith returns true if Fake{{.Interface}}.{{.Name}} was called with the given values
{{with $f := gensym}}func ({{$f}} *Fake{{$m.Interface}}) {{$m.Name}}CalledWith({{$m.ParametersDeclaration}}) bool {
	var found bool
	for _, call := range {{$f}}.{{$m.Name}}Calls {
		if {{range $i, $p := $m.Parameters}}{{if $i}} && {{end}}reflect.DeepEqual(call.Parameters.{{$p.TitleCase}}, {{$p.Name}}){{end}} {
			found = true
			break
		}
	}

	return found
}{{end}}

// Assert{{.Name}}CalledWith calls t.Error if Fake{{.Interface}}.{{.Name}} was not called with the given values
{{with $f := gensym}}func ({{$f}} *Fake{{$m.Interface}}) Assert{{$m.Name}}CalledWith(t *testing.T, {{$m.ParametersDeclaration}}) {
	t.Helper()
	var found bool
	for _, call := range {{$f}}.{{$m.Name}}Calls {
		if {{range $i, $p := $m.Parameters}}{{if $i}} && {{end}}reflect.DeepEqual(call.Parameters.{{$p.TitleCase}}, {{$p.Name}}){{end}} {
			found = true
			break
		}
	}

	if !found {
		t.Error("Fake{{$m.Interface}}.{{$m.Name}} not called with expected parameters")
	}
}{{end}}

// {{.Name}}CalledOnceWith returns true if Fake{{.Interface}}.{{.Name}} was called exactly once with the given values
{{with $f := gensym}}func ({{$f}} *Fake{{$m.Interface}}) {{$m.Name}}CalledOnceWith({{$m.ParametersDeclaration}}) bool {
	var count int
	for _, call := range {{$f}}.{{$m.Name}}Calls {
		if {{range $i, $p := $m.Parameters}}{{if $i}} && {{end}}reflect.DeepEqual(call.Parameters.{{$p.TitleCase}}, {{$p.Name}}){{end}} {
			count++
		}
	}

	return count == 1
}{{end}}

// Assert{{.Name}}CalledOnceWith calls t.Error if Fake{{.Interface}}.{{.Name}} was not called exactly once with the given values
{{with $f := gensym}}func ({{$f}} *Fake{{$m.Interface}}) Assert{{$m.Name}}OnceCalledWith(t *testing.T, {{$m.ParametersDeclaration}}) {
	t.Helper()
	var count int
	for _, call := range {{$f}}.{{$m.Name}}Calls {
		if {{range $i, $p := $m.Parameters}}{{if $i}} && {{end}}reflect.DeepEqual(call.Parameters.{{$p.TitleCase}}, {{$p.Name}}){{end}} {
			count++
		}
	}

	if count != 1 {
		t.Errorf("Fake{{$m.Interface}}.{{$m.Name}} called %d times with expected parameters, expected one", count)
	}
}{{end}}

// {{.Name}}ResultsForCall returns the result values for the first call to Fake{{.Interface}}.{{.Name}} with the given values
{{with $f := gensym}}func ({{$f}} *Fake{{$m.Interface}}) {{$m.Name}}ResultsForCall({{$m.ParametersDeclaration}}) ({{$m.ResultsDeclaration}}, found bool) {
	for _, call := range {{$f}}.{{$m.Name}}Calls {
		if {{range $i, $p := $m.Parameters}}{{if $i}} && {{end}}reflect.DeepEqual(call.Parameters.{{$p.TitleCase}}, {{$p.Name}}){{end}} {
			{{range $m.Results}}{{.Name}} = call.Results.{{.TitleCase}}
			{{end}}found = true
			break
		}
	}

	return
}{{end}}
{{end}}{{/* end if .Parameters */}}
{{end}}{{/* end range $m := .Methods */}}
{{end}}{{/* end range .Interfaces */}}
`

var (
	symGen         = SymbolGenerator{Prefix: "_f"}
	funky          = template.FuncMap{"gensym": func() string { return symGen.Next() }}
	charlatanTempl = template.Must(template.New("charlatan").Funcs(funky).Parse(charlatanTemplate))
)

type Template struct {
	CommandLine string
	PackageName string
	Imports     []*Import
	Interfaces  []*Interface
}

func (t *Template) Execute() ([]byte, error) {
	var buf bytes.Buffer
	if err := charlatanTempl.Execute(&buf, t); err != nil {
		return nil, err
	}

p	src, err := format.Source(buf.Bytes())
	if err != nil {
		// Should not happen except when developing this code.
		// The user can compile the output to see the error.
		return buf.Bytes(), fmt.Errorf("internal error: invalid code generated: %s", err)
	}

	return src, nil
}
