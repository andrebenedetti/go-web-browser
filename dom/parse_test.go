package dom

import (
	"fmt"
	"testing"
)

func TestAddUnclosedTag(t *testing.T) {
	tree := NewTree()
	tree.addTag("html")

	fmt.Println(tree.unfinished)
	if tree.unfinished[0].Text != "html" {
		t.Fatal("Error adding first open tag to the tree")
	}
}

func TestAddClosedEmptyTag(t *testing.T) {
	tree := NewTree()
	tree.addTag("html")
	tree.addTag("/html")
	// tree.Finish()
	fmt.Println(tree.unfinished)
	if tree.Root.Text != "html" {
		t.Fatal("html tag not found on root of tree")
	}

	if len(tree.Root.Children) != 0 {
		t.Fatal("Root's children should be empty")
	}
}

func TestAddClosedTagWithText(t *testing.T) {
	tree := NewTree()
	tree.addTag("p")
	tree.addText("Lorem ipsum")
	tree.addTag("/p")
	fmt.Println(tree.unfinished)
	if tree.Root.Text != "p" {
		t.Fatal("p tag not found on root of tree")
	}

	if len(tree.Root.Children) != 1 {
		t.Fatal("Tag should have 1 child")
	}

	if tree.Root.Children[0].Text != "Lorem ipsum" {
		t.Fatal("Tag's text was not properly parsed")
	}
}

func TestAddSelfClosingTag(t *testing.T) {
	tree := NewTree()

	tree.addTag("html")
	tree.addTag("meta charset=\"UTF-8\" /")

	if len(tree.unfinished) > 1 || tree.unfinished[0].Text != "html" {
		t.Fatal("Self-closing tag should not be added to the unfinished slice")
	}

	if tree.unfinished[0].Children[0].Text != "meta" {
		t.Fatal("Self closing tag should be finished and added to parent")
	}
}

func TestAddNestedTags(t *testing.T) {
	tree := NewTree()
	tree.addTag("html")
	tree.addTag("p")
	tree.addText("Lorem ipsum")
	tree.addTag("/p")
	tree.addTag("/html")

	if tree.Root.Text != "html" {
		t.Fatal("Root should be html")
	}

	if tree.Root.Children[0].Text != "p" {
		t.Fatal("Root html should have 'p' as child")
	}

	if tree.Root.Children[0].Children[0].Text != "Lorem ipsum" {
		t.Fatal("'p' should have 'Lorem ipsum' as child")
	}

}

func TestParseBody(t *testing.T) {
	body := `
	<!doctype html>
	<html>
		<head>
			<title>Example Domain</title>

			<meta charset="utf-8" />
			<meta http-equiv="Content-type" content="text/html; charset=utf-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1" />
		<style type="text/css">
			body {
				background-color: #f0f0f2;
				margin: 0;
				padding: 0;
				font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;

			}
			div {
				width: 600px;
				margin: 5em auto;
				padding: 2em;
				background-color: #fdfdff;
				border-radius: 0.5em;
				box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);
			}
			a:link, a:visited {
				color: #38488f;
				text-decoration: none;
			}
			@media (max-width: 700px) {
				div {
					margin: 0 auto;
					width: auto;
				}
			}
		</style>
		</head>

		<body>
		<div>
			<h1>Example Domain</h1>
			<p>This domain is for use in illustrative examples in documents. You may use this
			domain in literature without prior coordination or asking for permission.</p>
			<p><a href="https://www.iana.org/domains/example">More information...</a></p>
		</div>
		</body>
	</html>
`
	tree := NewTree()
	ParseBody(body, tree)

	if tree.Root.Children[0].Text != "head" {
		t.Fatal("Parsed body should have head as children of root element")
	}
	if tree.Root.Children[1].Text != "body" {
		t.Fatal("Parsed body should have body as children of root element")
	}

	// TODO: test the remaining of the tree
}
