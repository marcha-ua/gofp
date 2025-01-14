package axioms

import (
	"github.com/shful/gofp/owlfunctional/individual"
	"github.com/shful/gofp/owlfunctional/literal"
	"github.com/shful/gofp/owlfunctional/meta"
)

type DataPropertyDomain struct {
	R meta.DataProperty
	C meta.ClassExpression
}

type DataPropertyRange struct {
	R meta.DataProperty
	D meta.DataRange
}

// SubClassOf states C1 is subclass of C2
type SubClassOf struct {
	C1 meta.ClassExpression
	C2 meta.ClassExpression
}

type EquivalentClasses struct {
	EquivalentClasses []meta.ClassExpression //todo is there a min len in OWL ?
}

type DifferentIndividuals struct {
	As []individual.Individual
}

type DisjointClasses struct {
	DisjointClasses []meta.ClassExpression //todo is there a min len in OWL ?
}

// SubDataPropertyOf defines P1 subPropertyOf P2
type SubDataPropertyOf struct {
	P1 meta.DataProperty
	P2 meta.DataProperty
}

// SubObjectPropertyOf defines P1 subPropertyOf P2
type SubObjectPropertyOf struct {
	P1 meta.ObjectPropertyExpression
	P2 meta.ObjectPropertyExpression
}

// InverseObjectProperties defines P1 and P2 are inverse.
// InverseObjectProperties(P1,P2) implies InverseObjectProperties(P2,P1)
type InverseObjectProperties struct {
	P1 meta.ObjectPropertyExpression
	P2 meta.ObjectPropertyExpression
}

type ObjectPropertyDomain struct {
	P meta.ObjectPropertyExpression
	C meta.ClassExpression
}

type ObjectPropertyRange struct {
	P meta.ObjectPropertyExpression
	C meta.ClassExpression
}

type ClassAssertion struct {
	C meta.ClassExpression
	A individual.Individual
}

type DataPropertyAssertion struct {
	R meta.DataProperty
	A individual.Individual
	V literal.OWLLiteral
}
