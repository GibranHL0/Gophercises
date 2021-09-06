package Link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a Link in an HTML document
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a slice of links parsed
// from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	// 1.  Find <a> nodes in document
	nodes := linkNodes(doc)

	// 2. Create the links slice
	links := make([]Link, 0)

	// 3.  For each link node
	for _, node := range nodes {
		// 3.1 Build a Link
		links = append(links, buildLink(node))
	}

	// 4.  Return the Links
	return links, nil
}

func buildLink(n *html.Node) Link {
	var link Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}

	link.Text = getText(n)

	return link
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var text string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c)
	}

	return strings.Join(strings.Fields(text), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var nodes []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}

	return nodes
}
