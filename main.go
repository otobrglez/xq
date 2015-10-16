package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
)

func main() {
	xpathString := ""
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Missing second argument, XPATH!")
		os.Exit(2)
	}   else {
		xpathString = os.Args[1]
	}

	page, _ := ioutil.ReadAll(os.Stdin)
	doc, err := gokogiri.ParseHtml(page); if err != nil {
		fmt.Fprintln(os.Stderr, "Problem parsing document.")
	}
	defer doc.Free()

	xps := xpath.Compile(xpathString)
	defer xps.Free()

	search, err := doc.Search(xps); if err == nil {
		for _, s := range search {
			fmt.Println(s.Content())
		}
	} else {
		fmt.Fprintln(os.Stderr, "Sorry. Got error.")
	}
}