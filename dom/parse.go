package dom

import (
	"fmt"
	"strings"
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
	if tag[0] == '/' || tag[len(tag)-1] == '/' {
		fmt.Printf("Adding tag: %s\n", tag)
		// pop last node

		var node *Element
		// self closing tag needs further parsing
		if tag[len(tag)-1] == '/' {
			tag = strings.Split(tag, " ")[0]
			parent := t.unfinished[len(t.unfinished)-1]
			node = &Element{text: tag, parent: parent, children: make([]*Element, 0, 1)}
		} else {
			node = t.unfinished[len(t.unfinished)-1]
		}

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
