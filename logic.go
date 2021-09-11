package lep

import (
	"strings"
)

func parseExpressions(items ...interface{}) (expr []Expression) {
	for _, item := range items {
		if e, ok := item.(Expression); ok {
			expr = append(expr, e)
		} else if slice, ok := item.([]interface{}); ok {
			expr = append(expr, parseExpressions(slice...)...)
		}
	}
	return expr
}

type AndX struct {
	conjuncts []Expression
}

var _ Expression = (*AndX)(nil)

func And(expr ...Expression) *AndX {
	var conjuncts []Expression
	for _, e := range expr {
		if exprAndX, ok := e.(*AndX); ok {
			conjuncts = append(conjuncts, exprAndX.conjuncts...)
		} else {
			conjuncts = append(conjuncts, e)
		}
	}
	return &AndX{conjuncts: conjuncts}
}

func (e AndX) Equals(other Expression) bool {
	otherAndX, ok := other.(*AndX)
	if !ok {
		return false
	}

	leftConjuncts := e.conjuncts
	rightConjuncts := make([]Expression, len(otherAndX.conjuncts))
	copy(rightConjuncts, otherAndX.conjuncts)

	var leftFound bool
	for _, leftConjunct := range leftConjuncts {
		leftFound = false
		for i, rightConjunct := range rightConjuncts {
			if leftConjunct.Equals(rightConjunct) {
				rightConjuncts = append(rightConjuncts[:i], rightConjuncts[i+1:]...)
				leftFound = true
				break
			}
		}
		if !leftFound {
			return false
		}
	}
	return len(rightConjuncts) == 0
}

func (e AndX) String() string {
	var items []string
	for _, conjunct := range e.conjuncts {
		if _, ok := conjunct.(*OrX); ok {
			items = append(items, "("+conjunct.String()+")")
		} else {
			items = append(items, conjunct.String())
		}
	}
	return strings.Join(items, " && ")
}

func parseAnd(elements ...interface{}) (*AndX, error) {
	expr := parseExpressions(elements...)
	return And(expr...), nil
}

type OrX struct {
	disjunctions []Expression
}

var _ Expression = (*OrX)(nil)

func Or(expr ...Expression) *OrX {
	var disjunctions []Expression
	for _, e := range expr {
		if exprOrX, ok := e.(*OrX); ok {
			disjunctions = append(disjunctions, exprOrX.disjunctions...)
		} else {
			disjunctions = append(disjunctions, e)
		}
	}
	return &OrX{disjunctions: disjunctions}
}

func (e OrX) Equals(other Expression) bool {
	otherOrX, ok := other.(*OrX)
	if !ok {
		return false
	}

	leftDisjunctions := e.disjunctions
	rightDisjunctions := make([]Expression, len(otherOrX.disjunctions))
	copy(rightDisjunctions, otherOrX.disjunctions)

	var leftFound bool
	for _, leftDisjunction := range leftDisjunctions {
		leftFound = false
		for i, rightDisjunction := range rightDisjunctions {
			if leftDisjunction.Equals(rightDisjunction) {
				rightDisjunctions = append(rightDisjunctions[:i], rightDisjunctions[i+1:]...)
				leftFound = true
				break
			}
		}
		if !leftFound {
			return false
		}
	}
	return len(rightDisjunctions) == 0
}

func (e OrX) String() string {
	var items []string
	for _, disjunction := range e.disjunctions {
		items = append(items, disjunction.String())
	}
	return strings.Join(items, " || ")
}

func parseOr(elements ...interface{}) (*OrX, error) {
	expr := parseExpressions(elements...)
	return Or(expr...), nil
}
