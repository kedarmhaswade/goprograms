package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("doc.FirstChild = %v\n", doc.FirstChild.Data)
	fmt.Printf("doc.Type = %#v\n", doc.Type)
	outline(nil, doc)
}
func outline(stack []string, node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Println(stack)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(append(stack, c.Data), c)
	}
}
