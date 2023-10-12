package dom

import (
	"fmt"
)

// elementType may be "Element" or "Text"
type Element struct {
	elementType string
	text        string
	children    []*Element
	parent      *Element
}

type tree struct {
	Root       *Element
	unfinished []*Element
}

func NewTree() *tree {
	return &tree{unfinished: make([]*Element, 0)}
}

func (t *tree) AddText(text string) {
	fmt.Printf("Adding text: %s\n", text)
	parent := t.unfinished[len(t.unfinished)-1]
	node := Element{text: text, parent: parent, elementType: "text"}
	parent.children = append(parent.children, &node)
}

func (t *tree) AddTag(tag string) {
	if tag == "!doctype html" {
		return
	}
	if tag[0] == '/' {
		fmt.Printf("Adding tag: %s\n", tag)
		// pop last node
		node := t.unfinished[len(t.unfinished)-1]
		if len(t.unfinished) == 1 {
			t.Root = node
			return
		}

		newUnifished := make([]*Element, len(t.unfinished)-1)
		copy(newUnifished, t.unfinished[:len(t.unfinished)-1])
		t.unfinished = newUnifished

		parent := t.unfinished[len(t.unfinished)-1]
		parent.children = append(parent.children, node)
	} else {
		if len(t.unfinished) > 0 {
			parent := t.unfinished[len(t.unfinished)-1]
			node := Element{text: tag, parent: parent, children: make([]*Element, 0, 1)}
			t.unfinished = append(t.unfinished, &node)
		} else {
			node := Element{text: tag, children: make([]*Element, 0, 1)}
			t.unfinished = append(t.unfinished, &node)
		}
	}
}

// func (t *tree) Finish() *Element {
// 	fmt.Println("Finishing")
// 	fmt.Println(t.unfinished)
// 	fmt.Println(len(t.unfinished))
// 	if len(t.unfinished) == 0 {
// 		t.AddTag("html")
// 	}
// 	for len(t.unfinished) > 1 {
// 		node := t.unfinished[len(t.unfinished)-1]
// 		t.unfinished = t.unfinished[:len(t.unfinished)-1]
// 		parent := t.unfinished[len(t.unfinished)-1]
// 		parent.children = append(parent.children, node)
// 	}

// 	res := t.unfinished[len(t.unfinished)-1]
// 	t.unfinished = t.unfinished[:len(t.unfinished)-1]
// 	return res
// }

// func findHtmlTag(document string) (int, error) {
// 	inTag := false
// 	tag := ""
// 	for i, c := range document {

// 		if c == '<' {
// 			inTag = true
// 		} else if c == '>' {
// 			inTag = false
// 			if tag == "body" {
// 				return i - 5, nil
// 			} else {
// 				tag = ""
// 			}
// 		} else if inTag {
// 			tag += string(c)
// 		}
// 	}

// 	return -1, errors.New("tag <html> was not found in document")
// }

// Parse picks up a complete document, which might include
// <head>, <script>, <css> and other things, separates the html
// body and builds the DOM tree with it.
// func Parse(document string, tree *tree) error {
// 	// var htmlBody string

// 	// htmlTagIndex, err := findHtmlTag(document)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(document[htmlTagIndex:])
// 	parseBody(document[htmlTagIndex:], tree)

// 	return nil
// }

func PrintTree(root *Element, indent int) {
	if root == nil {
		return
	}
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Println(root.text)

	for _, child := range root.children {
		PrintTree(child, indent+2)
	}
}

// TODO: we need a "finish" function to handle the case when we have text (unclosed, of course)
// together with the finishing tag ("html")
func ParseBody(body string, tree *tree) {
	text := ""
	inTag := false
	for _, c := range body {
		if c == '<' {
			inTag = true
			if text != "" {
				tree.AddText(text)
				text = ""
			}
		} else if c == '>' {
			inTag = false
			tree.AddTag(text)
			text = ""
		} else {
			text = text + string(c)
		}
	}

	if !inTag && text != "" {
		tree.AddText(text)
	}

}
