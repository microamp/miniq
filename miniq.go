package miniq

import (
	"io"

	"golang.org/x/net/html"
)

type predicate func(*html.Node) bool

func search(in chan *html.Node, p predicate) chan *html.Node {
	out := make(chan *html.Node)

	go func() {
		defer close(out)

		var fn func(n *html.Node)
		fn = func(n *html.Node) {
			if p(n) {
				out <- n
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				fn(c)
			}
		}

		for n := range in {
			fn(n)
		}
	}()

	return out
}

func Q(r io.Reader, qs string) (chan *html.Node, error) {
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	ps := preds(qs)

	first := make(chan *html.Node)
	in := first

	var out chan *html.Node
	for _, p := range ps {
		out = search(in, p)
		in = out
	}

	go func() {
		first <- root
		close(first)
	}()

	return out, nil
}

func QURL(url, qs string) (chan *html.Node, error) {
	r, err := fetch(url)
	if err != nil {
		return nil, err
	}
	return Q(r, qs)
}
