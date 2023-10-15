package family

import "fmt"

type Child struct {
	Name     string
	Children *[]Child
}

type Tree struct {
	children []Child
}

func NewTree() *Tree {
	return &Tree{
		children: []Child{},
	}
}

func (t *Tree) AddChild(newChild Child) {
	t.children = append(t.children, newChild)
}

func (t *Tree) FindChildrenRecursively() {
	t.findChildrenRecursive(&t.children)
}

func (t *Tree) findChildrenRecursive(children *[]Child) {
	if children == nil {
		return
	}

	for _, child := range *children {
		fmt.Println(child.Name)
		t.findChildrenRecursive(child.Children)
	}
}
