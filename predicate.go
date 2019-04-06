package miniq

import (
	"strings"

	"golang.org/x/net/html"
)

const (
	symbolClass = "."
	symbolID    = "#"
)

type token struct {
	raw string
}

func (t *token) hasSymbol(symbol string) bool {
	return strings.Contains(t.raw, symbol)
}

func (t *token) getAttrVal(symbol string) string {
	return strings.Split(t.raw, symbol)[1]
}

func (t *token) hasClass() bool {
	return t.hasSymbol(symbolClass)
}

func (t *token) getClassVal() string {
	return t.getAttrVal(symbolClass)
}

func (t *token) hasID() bool {
	return t.hasSymbol(symbolID)
}

func (t *token) getIDVal() string {
	return t.getAttrVal(symbolID)
}

func (t *token) getData() string {
	if t.hasClass() {
		return strings.Split(t.raw, symbolClass)[0]
	}
	if t.hasID() {
		return strings.Split(t.raw, symbolID)[0]
	}
	return t.raw
}

func tokenise(qs string) []*token {
	var ts []*token
	for _, s := range strings.Split(qs, " ") {
		ts = append(ts, &token{raw: s})
	}
	return ts
}

func pred(t *token) predicate {
	return func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == t.getData() {
			if !t.hasClass() && !t.hasID() {
				return true
			} else {
				classVal := t.getClassVal()
				for _, a := range n.Attr {
					if a.Key == "class" {
						for _, val := range strings.Split(a.Val, " ") {
							if val == classVal {
								return true
							}
						}
					}
				}
				return false
			}
		}
		return false
	}
}

func preds(qs string) []predicate {
	var ps []predicate
	for _, t := range tokenise(qs) {
		ps = append(ps, pred(t))
	}
	return ps
}
