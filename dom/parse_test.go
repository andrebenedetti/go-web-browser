package dom

import (
	"fmt"
	"testing"
)

func TestAddUnclosedTag(t *testing.T) {
	tree := NewTree()
	tree.AddTag("html")

	fmt.Println(tree.unfinished)
	if tree.unfinished[0].text != "html" {
		t.Fatal("Error adding first open tag to the tree")
	}
}

func TestAddClosedEmptyTag(t *testing.T) {
	tree := NewTree()
	tree.AddTag("html")
	tree.AddTag("/html")
	// tree.Finish()
	fmt.Println(tree.unfinished)
	if tree.Root.text != "html" {
		t.Fatal("html tag not found on root of tree")
	}

	if len(tree.Root.children) != 0 {
		t.Fatal("Root's children should be empty")
	}
}

func TestAddClosedTagWithText(t *testing.T) {
	tree := NewTree()
	tree.AddTag("p")
	tree.AddText("Lorem ipsum")
	tree.AddTag("/p")
	// tree.Finish()
	fmt.Println(tree.unfinished)
	if tree.Root.text != "p" {
		t.Fatal("p tag not found on root of tree")
	}

	if len(tree.Root.children) != 1 {
		t.Fatal("Tag should have 1 child")
	}

	if tree.Root.children[0].text != "Lorem ipsum" {
		t.Fatal("Tag's text was not properly parsed")
	}
}

func TestAddNestedTags(t *testing.T) {
	tree := NewTree()
	tree.AddTag("html")
	tree.AddTag("p")
	tree.AddText("Lorem ipsum")
	tree.AddTag("/p")
	tree.AddTag("/html")

	if tree.Root.text != "html" {
		t.Fatal("Root should be html")
	}

	if tree.Root.children[0].text != "p" {
		t.Fatal("Root html should have 'p' as child")
	}

	if tree.Root.children[0].children[0].text != "Lorem ipsum" {
		t.Fatal("'p' should have 'Lorem ipsum' as child")
	}

}
