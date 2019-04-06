package main

import (
	"fmt"

	"github.com/microamp/miniq"
)

func main() {
	url := "http://www.nzkoreapost.com/bbs/board.php?bo_table=market_buynsell&sca=%EC%9C%A0%EC%95%84%EC%9A%A9%ED%92%88"
	qs := "div.list-pc td.list-subject a strong"

	fmt.Println("Listing baby products on sale from The Korea Post...")

	ns, err := miniq.QURL(url, qs)
	if err != nil {
		panic(err)
	}

	for n := range ns {
		fmt.Println(n.FirstChild.Data)
	}
}
