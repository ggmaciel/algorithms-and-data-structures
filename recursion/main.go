package main

import "recursion/family"

func main() {
	tree := family.NewTree()

	maria := family.Child{Name: "Maria", Children: &[]family.Child{{Name: "Gustavo", Children: nil}}}
	joao := family.Child{Name: "João", Children: &[]family.Child{maria}}
	antonio := family.Child{Name: "Antonio", Children: nil}

	tree.AddChild(joao)
	tree.AddChild(antonio)
	tree.FindChildrenRecursively()
}
