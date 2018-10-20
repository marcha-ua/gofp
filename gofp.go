// gofp is an Owl Functional Parser.
// gofp reads OWL-Functional input into Golang structures, for further processing.
package gofp

// The names and the package structure of gofp resemble the OWL Quick Reference Guide, found here:
// https://www.w3.org/2007/OWL/refcard
// For example, the gofp package "axioms" resembles the Guides section "2.5 Axioms".

// Some things in gofp are, surely, made wrong.
// At least, the following statements, found in:
//   https://www.w3.org/TR/owl2-syntax/#Appendix:_Complete_Grammar_.28Normative.29
// were not considered yet:

// - OWL functional-style Syntax documents may have the strings "Prefix" or "Ontology" (case dependent) near the beginning of the document.
//  Remark: what means "near" ?

// - Sets written in one of the exchange syntaxes (e.g., XML or RDF/XML) are not necessarily expected to be duplicate free. Duplicates SHOULD be eliminated when ontology documents written in such syntaxes are converted into instances of the UML classes of the structural specification.
// 	An ontology written in functional-style syntax can contain the following class expression:
// 	ObjectUnionOf( a:Person a:Animal a:Animal )
// 	During parsing, this expression should be "flattened" to the following expression:
// 	ObjectUnionOf( a:Person a:Animal )
//  Remark: Gofp does not yet "flatten" that.

// - A functional-style syntax ontology document SHOULD use the UTF-8 encoding [RFC 3629].
//  Remark: For gofp, it MUST be UTF-8

import (
	"io"

	"reifenberg.de/gofp/owlfunctional/ontologies"
	"reifenberg.de/gofp/owlfunctional/parser"
)

// OntologyFromReader parses an owl-functional file contents into an Ontology struct.
// r is the OWL-Functional file contents.
// sourceName: see parser.NewParser()
func OntologyFromReader(r io.Reader, sourceName string) (ontology *ontologies.Ontology, err error) {
	var prefixes map[string]string

	p := parser.NewParser(r, sourceName)
	// parser.TokenLog = true
	prefixes = map[string]string{}
	for {
		tok, lit, pos := p.ScanIgnoreWSAndComment()
		switch tok {
		case parser.Prefix:
			p.Unscan()
			if err = parsePrefixTo(prefixes, lit, p); err != nil {
				err = pos.Errorf("Parsing prefix raised:%v", err)
				return
			}
		case parser.Ontology:
			p.Unscan()
			ontology = ontologies.NewOntology(prefixes)
			if err = ontology.Parse(p); err != nil {
				return
			}
		case parser.EOF:
			return
		default:
			err = pos.ErrorfUnexpectedToken(tok, lit, "Prefix or Ontology")
			return
		}
	}

}

// parsePrefixTo parses the next Prefix expression and
// fills the given prefixes map.
func parsePrefixTo(prefixes map[string]string, lit string, p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.Prefix, parser.B1); err != nil {
		return err
	}
	tok, prefix, pos := p.ScanIgnoreWSAndComment()

	if tok == parser.COLON {
		// empty Prefix(:=...)
		p.Unscan()
		prefix = ""
	} else {
		// Prefix(IDENT:=...)
		if tok != parser.IDENT {
			return pos.Errorf("unexpected %v when parsing prefix, need IDENT", prefix)
		}
	}
	if err = p.ConsumeTokens(parser.COLON, parser.EQUALS); err != nil {
		return err
	}
	tok, prefixVal, pos := p.ScanIgnoreWSAndComment()
	if tok != parser.IRI {
		return pos.Errorf("unexpected %v when parsing prefix, need IRI", prefixVal)
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return err
	}
	if _, ok := prefixes[prefix]; ok {
		return pos.Errorf(`second occurrence of prefix "%v"`, prefix)
	}
	prefixes[prefix] = prefixVal
	return
}