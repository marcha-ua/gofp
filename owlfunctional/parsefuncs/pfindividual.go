package parsefuncs

import (
	"reifenberg.de/gofp/owlfunctional/individual"
	"reifenberg.de/gofp/owlfunctional/parser"
	"reifenberg.de/gofp/parsehelper"
	"reifenberg.de/gofp/tech"
)

func ParseIndividual(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (a individual.Individual, err error) {
	var prefix, name string
	pos := p.Pos()
	prefix, name, err = parsehelper.ParsePrefixedName(p)
	if err != nil {
		err = pos.Errorf("parsing individual:%v", err)
		return
	}
	a = individual.Individual{Name: parser.FmtPrefixedName(prefix, name)}
	return
}

// ParseIndividualsUntilB2 parses all Individuals until ")" is found
// The closing ")" is not consumed.
func ParseIndividualsUntilB2(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (as []individual.Individual, err error) {

	var tok parser.Token
	var a individual.Individual

	for {
		tok, _, _ = p.ScanIgnoreWSAndComment()
		p.Unscan()
		if tok == parser.B2 {
			break
		}

		a, err = ParseIndividual(p, decls, prefixes)
		if err != nil {
			return
		}
		as = append(as, a)
	}

	return
}