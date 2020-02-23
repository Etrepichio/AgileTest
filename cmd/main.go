package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	var (
		searchedElement = flag.String("searchedElement", "make-everything-ok-button", "element ID to found")
		originalHTML    = flag.String("originalHTML", "sample-0-origin.html", "sample HTML address in which to find an element")
		diffCaseHTML    = flag.String("diffCaseHTML", "sample-1-evil-gemini.html", "HTML file address to compare")
	)
	flag.Parse()

	//First, we read the originalHTML searching for the Element we want to use to search for similars on the other document.
	input, err := ioutil.ReadFile(*originalHTML)
	if err != nil {
		panic("Failure at reading original HTML")
	}

	doc, err := html.Parse(strings.NewReader(string(input)))
	if err != nil {
		panic("Failure at parsing original HTML")
	}
	var originpath []string
	node := searchElement(doc, "id", *searchedElement, &originpath)
	if node == nil {
		panic("Desired element wasnt found in original HTML")
	}

	//Now we read the target file
	input2, err := ioutil.ReadFile(*diffCaseHTML)
	if err != nil {
		panic("Failure at reading diffCase HTML")
	}

	doc2, err := html.Parse(strings.NewReader(string(input2)))
	if err != nil {
		panic("Failure at parsing diffCase HTML")
	}

	//Using the information of the original element (all his attributes) we will search for any similar element(the ones that share a key-value data) in the target file
	result := foundSimilar(node, doc2)
	if result != nil {
		for _, p := range result {
			fmt.Println(p.node)
			fmt.Println(p.path)
		}
	}

}

//getAttribute obtains any of the attributes of an element by its key name
func getAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

//checkAttribute checks the value of the key is the same as the wanted
func checkAttribute(n *html.Node, key, value string) bool {
	if n.Type == html.ElementNode {
		s, ok := getAttribute(n, key)
		if ok && s == value {
			return true
		}
	}
	return false
}

//searchElement searches recursively for a key-value pair in each node (element) of the html document
func searchElement(n *html.Node, key, value string, path *[]string) *html.Node {
	if checkAttribute(n, key, value) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := searchElement(c, key, value, path)
		if result != nil {
			*path = append(*path, n.Data)
			return result
		}
	}

	return nil
}

//foundSimilar looks for every node(element) similar as the one labeled as origin in the node labeled as target.
func foundSimilar(origin, target *html.Node) (result []Result) {
	for _, attr := range origin.Attr {
		var p []string
		simil := searchElement(target, attr.Key, attr.Val, &p)
		if simil != nil {
			res := Result{
				node: simil,
				path: p,
			}
			result = append(result, res)
		}
	}

	return
}

//Result contains the information of each match: The node itself and the path to it
type Result struct {
	node *html.Node
	path []string
}
