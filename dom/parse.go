package dom

import (
	"fmt"
	"strings"
)

type Element struct {
	elementType string // "element" or "text"
	text        string
	children    []*Element
	parent      *Element
	attributes  map[string]string
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

type ParsedTag struct {
	isSelfClosing bool
	// TODO: self-closing tags are not defined by the ending slash '/', as it is actually optional
	// Some elements are by definition self closing, so we can't rely on the ending slash to parse that.
	// isClosing means it ends with "/>" and is NOT self-closing
	isClosing  bool
	tag        string
	attributes map[string]string
}

func parseTag(tag string) ParsedTag {
	t := ParsedTag{}
	t.isSelfClosing = tag[len(tag)-1] == '/'
	t.isClosing = tag[0] == '/'

	// Opening and self-closing tags may have attributes
	// like <meta charset="UTF-8" /> or <p class="title">
	// Also, parse attributes such as <button disabled> in which
	// we don't have the form attr=val
	if t.isSelfClosing || !t.isClosing {
		split := strings.Split(tag, " ")
		t.tag = split[0]

		t.attributes = make(map[string]string)
		for i := 1; i < len(split); i++ {
			if strings.Contains(split[i], "=") {
				kv := strings.Split(split[i], "=")
				// Remove quotes if present
				if strings.Contains(kv[1], "'") || strings.Contains(kv[1], "\"") {
					t.attributes[kv[0]] = kv[1][1 : len(kv[1])-1]
				} else {
					t.attributes[kv[0]] = kv[1]
				}
			} else {
				// attributes without '=' default to an empty string
				t.attributes[split[i]] = ""
			}
		}
	}

	return t
}

func (t *tree) AddTag(tag string) {
	if tag == "!doctype html" {
		return
	}

	newTag := parseTag(tag)

	var node *Element
	if newTag.isSelfClosing {
		parent := t.unfinished[len(t.unfinished)-1]
		node = &Element{text: newTag.tag, parent: parent, children: make([]*Element, 0, 1), attributes: newTag.attributes}
		parent.children = append(parent.children, node)
		return
	}

	if newTag.isClosing {
		// self closing tag needs further parsing
		node = t.unfinished[len(t.unfinished)-1]

		if len(t.unfinished) == 1 {
			t.Root = node
			return
		}

		t.unfinished = t.unfinished[:len(t.unfinished)-1]

		parent := t.unfinished[len(t.unfinished)-1]
		parent.children = append(parent.children, node)
	} else {
		if len(t.unfinished) > 0 {
			parent := t.unfinished[len(t.unfinished)-1]
			node := Element{text: newTag.tag, parent: parent, children: make([]*Element, 0, 1)}
			t.unfinished = append(t.unfinished, &node)
		} else {
			node := Element{text: tag, children: make([]*Element, 0, 1), attributes: newTag.attributes}
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
	fmt.Println(root.attributes)

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
