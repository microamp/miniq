package main

import (
	"fmt"

	"github.com/microamp/miniq"
)

func main() {
	url := "https://github.com/trending"
	qs := "ol.repo-list li.d-block div.d-inline-block h3 a"

	fmt.Println("Listing trending repositories from GitHub...")

	ns, err := miniq.QURL(url, qs)
	if err != nil {
		panic(err)
	}

	rank := 1
	for n := range ns {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Printf("#%02d: https://github.com/%s\n", rank, a.Val)
				continue
			}
		}
		rank++
	}
}
