package parsefuncs

import (
	"testing"

	"reifenberg.de/gofp/mock"
	"reifenberg.de/gofp/owlfunctional/literal"
	"reifenberg.de/gofp/owlfunctional/parser"
)

func TestParseInt(t *testing.T) {
	var p *parser.Parser
	var err error
	var l literal.OWLLiteral
	_, prefixes := mock.NewBuilder().AddPrefixes("xsd").Get()

	p = mock.NewTestParser(`1`)
	l, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if l.Value != "1" {
		t.Fatal(l)
	}
	if l.Literaltype != "xsd:integer" {
		t.Fatal(l)
	}
	if l.LangTag != "" {
		t.Fatal(l)
	}

	p = mock.NewTestParser(`"099"^^xsd:positiveInteger`)
	l, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if l.Value != "099" {
		t.Fatal(l)
	}
	if l.Literaltype != "xsd:positiveInteger" {
		t.Fatal(l)
	}
	if l.LangTag != "" {
		t.Fatal(l)
	}
}
func TestParseFloat(t *testing.T) {
	var p *parser.Parser
	var err error
	var l literal.OWLLiteral
	_, prefixes := mock.NewBuilder().AddPrefixes("xsd").Get()

	p = mock.NewTestParser(`3.0`)
	l, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if l.Value != "3.0" {
		t.Fatal(l)
	}
	if l.Literaltype != "xsd:decimal" {
		t.Fatal(l)
	}
	if l.LangTag != "" {
		t.Fatal(l)
	}
}

func TestParseString(t *testing.T) {
	var p *parser.Parser
	var err error
	var l literal.OWLLiteral
	_, prefixes := mock.NewBuilder().AddPrefixes("xsd").Get()

	p = mock.NewTestParser(`"Hello World"`)
	l, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if l.Value != "Hello World" {
		t.Fatal(l.Value)
	}
	if l.Literaltype != "xsd:string" {
		t.Fatal(l)
	}
	if l.LangTag != "" {
		t.Fatal(l)
	}

	p = mock.NewTestParser(`"Hello Wörld"@en`)
	l, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if l.Value != "Hello Wörld" {
		t.Fatal(l.Value)
	}
	if l.Literaltype != "xsd:string" {
		t.Fatal(l)
	}
	if l.LangTag != "en" {
		t.Fatal(l)
	}

	p = mock.NewTestParser(`"123"^^xsd:string`)
	l, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if l.Value != "123" {
		t.Fatal(l.Value)
	}
	if l.Literaltype != "xsd:string" {
		t.Fatal(l)
	}
	if l.LangTag != "" {
		t.Fatal(l)
	}

	p = mock.NewTestParser(`"0.0"@LongLangTäg^^xsd:string`)
	l, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if l.Value != "0.0" {
		t.Fatal(l.Value)
	}
	if l.Literaltype != "xsd:string" {
		t.Fatal(l)
	}
	if l.LangTag != "LongLangTäg" {
		t.Fatal(l)
	}
}

func TestParseBool(t *testing.T) {
	var p *parser.Parser
	var err error
	var l literal.OWLLiteral
	_, prefixes := mock.NewBuilder().AddPrefixes("xsd").Get()

	p = mock.NewTestParser(`true`)
	l, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if l.Value != "true" {
		t.Fatal(l.Value)
	}
	if l.Literaltype != "xsd:boolean" {
		t.Fatal(l)
	}
	if l.LangTag != "" {
		t.Fatal(l)
	}
}

func TestParseMismatches(t *testing.T) {
	var p *parser.Parser
	var err error
	_, prefixes := mock.NewBuilder().AddPrefixes("xsd").Get()

	p = mock.NewTestParser(`1^^xsd:string`)
	_, err = ParseOWLLiteral(p, prefixes)
	if err == nil {
		t.Fatal("mismatch expected")
	}
	p = mock.NewTestParser(`"1A"^^xsd:int`)
	_, err = ParseOWLLiteral(p, prefixes)
	if err == nil {
		t.Fatal("mismatch expected")
	}
}

func TestParseCustomDatatypeName(t *testing.T) {
	var p *parser.Parser
	var err error
	var l literal.OWLLiteral
	_, prefixes := mock.NewBuilder().AddPrefixes("", "xsd").Get()

	p = mock.NewTestParser(`"Da5id"^^:blacksun`)
	l, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if l.Value != "Da5id" {
		t.Fatal(l)
	}
	if l.Literaltype != ":blacksun" {
		t.Fatal(l)
	}
	if l.LangTag != "" {
		t.Fatal(l)
	}

	p = mock.NewTestParser(`"Da5id"^^unknown:blacksun`)
	l, err = ParseOWLLiteral(p, prefixes)
	if err == nil {
		t.Fatal("prefix error expected")
	}
}