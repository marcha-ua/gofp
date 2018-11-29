package parsehelper

import (
	"net/url"
	"strconv"
	"strings"

	"reifenberg.de/gofp/owlfunctional/parser"
)

func ParsePrefixedName(p *parser.Parser) (prefix, name string, err error) {

	tok, lit, pos := p.ScanIgnoreWSAndComment()
	if tok == parser.IDENT {
		// prefix:classname
		prefix = lit
		if err = p.ConsumeTokens(parser.COLON); err != nil {
			return
		}
	} else if tok == parser.COLON {
		// :classname
		prefix = ""
	} else {
		err = pos.Errorf("unexpected %v, need prefixed name", lit)
		return
	}

	tok, name, pos = p.ScanIgnoreWSAndComment()
	if tok != parser.IDENT {
		err = pos.Errorf("unexpected %v, need identifier in prefixed name", lit)
	}
	return
}

// ParseUnprefixedIRI parses an IRI which is not shortened with a prefix. Instead, it must look like "<.*>"
func ParseUnprefixedIRI(p *parser.Parser) (iri string, err error) {
	pos := p.Pos()
	tok, iri, pos := p.ScanIgnoreWSAndComment()
	if tok == parser.IRI {
		if !(strings.HasPrefix(iri, "<") && strings.HasSuffix(iri, ">")) {
			err = pos.Errorf("expected IRI, but missing < and > on the ends (found:%v)", iri)
		}
	} else {
		err = pos.Errorf("expected IRI, but found:%v", parser.DescribeToklit(tok, iri))
	}
	return
}

// ParseUnprefixedIRI parses an IRI which must be surrounded with "<" ">" and must have a fragment, separated with #.
func ParseIRIWithFragment(p *parser.Parser) (prefix, fragment string, err error) {
	pos := p.Pos()
	tok, iri, pos := p.ScanIgnoreWSAndComment()
	if tok == parser.IRI {
		if !(strings.HasPrefix(iri, "<") && strings.HasSuffix(iri, ">")) {
			err = pos.Errorf("expected IRI, but missing < and > on the ends (found:%v)", iri)
			return
		}
		var u *url.URL
		u, err = url.Parse(iri[1 : len(iri)-1])
		fragment = u.Fragment
		if len(fragment) == 0 {
			err = pos.Errorf("expected IRI with fragment, but missing (found:%v)", iri)
			return
		}
		prefix = iri[1 : len(iri)-1-len(fragment)-1] // everything before, and excluding, the fragments "#"
	} else {
		err = pos.Errorf("expected IRI, but found:%v", parser.DescribeToklit(tok, iri))
	}
	return
}

func ParseNonNegativeInteger(p *parser.Parser) (res int, err error) {
	tok, lit, pos := p.ScanIgnoreWSAndComment()
	if tok != parser.INTLIT {
		err = pos.Errorf("int literal needed, found %v", lit)
		return
	}
	res, err = strconv.Atoi(lit)
	if res < 0 {
		err = pos.Errorf("nonnegative integer needed, found %v", lit)
		return
	}
	return
}
