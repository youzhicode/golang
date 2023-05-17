package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %v\n", url, err)
	}
	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/html" && strings.HasPrefix(contentType, "text/html") {
		return fmt.Errorf("%s document type unsuport %s\n", url, contentType)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	visitAll := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitAll, nil)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		forEachNode(c, pre, nil)
	}
}

func main() {

	for _, url := range os.Args[1:] {
		if err := title(url); err != nil {
			fmt.Printf("main %s have error %v\n", url, err)
		}
	}

}
