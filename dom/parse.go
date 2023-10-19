package dom

import (
	"fmt"
	"strings"
)

type Element struct {
	ElementType string // "element" or "text"
	Text        string
	Children    []*Element
	Parent      *Element
}

type Tree struct {
	Root       *Element
	unfinished []*Element
}

func NewTree() *Tree {
	return &Tree{unfinished: make([]*Element, 0)}
}

func (t *Tree) addText(text string) {
	parent := t.unfinished[len(t.unfinished)-1]
	node := Element{Text: text, Parent: parent, ElementType: "text"}
	parent.Children = append(parent.Children, &node)
}

func (t *Tree) addTag(tag string) {
	if tag == "!doctype html" {
		return
	}
	if tag[0] == '/' || tag[len(tag)-1] == '/' {
		var node *Element
		// self closing tag needs further parsing
		if tag[len(tag)-1] == '/' {
			tag = strings.Split(tag, " ")[0]
			parent := t.unfinished[len(t.unfinished)-1]
			node = &Element{Text: tag, Parent: parent, Children: make([]*Element, 0, 1), ElementType: "tag"}
			parent.Children = append(parent.Children, node)
			return
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
		parent.Children = append(parent.Children, node)
	} else {
		if len(t.unfinished) > 0 {
			parent := t.unfinished[len(t.unfinished)-1]
			node := Element{Text: tag, Parent: parent, Children: make([]*Element, 0, 1), ElementType: "tag"}
			t.unfinished = append(t.unfinished, &node)
		} else {
			node := Element{Text: tag, Children: make([]*Element, 0, 1), ElementType: "tag"}
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
	fmt.Println(root.Text)

	for _, child := range root.Children {
		PrintTree(child, indent+2)
	}
}

func ParseBody(body string, tree *Tree) {
	body = strings.ReplaceAll(body, "\n", "")
	body = strings.ReplaceAll(body, "\t", "")
	text := ""
	inTag := false
	for _, c := range body {
		if c == '<' {
			inTag = true
			if text != "" {
				tree.addText(text)
				text = ""
			}
		} else if c == '>' {
			inTag = false
			tree.addTag(text)
			text = ""
		} else {
			text = text + string(c)
		}
	}

	if !inTag && text != "" {
		tree.addText(text)
	}
}
