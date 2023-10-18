package tree

type BinaryNode[T any] struct {
	value int
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func postOrderWalk[T any](curr *BinaryNode[T], path *[]int) []int {
	if curr == nil {
		return *path
	}

	postOrderWalk(curr.left, path)
	postOrderWalk(curr.right, path)
	*path = append(*path, curr.value)

	return *path
}

func inOrderWalk[T any](curr *BinaryNode[T], path *[]int) []int {
	if curr == nil {
		return *path
	}

	inOrderWalk(curr.left, path)
	*path = append(*path, curr.value)
	inOrderWalk(curr.right, path)

	return *path
}

func preOrderWalk[T any](curr *BinaryNode[T], path *[]int) []int {
	if curr == nil {
		return *path
	}

	*path = append(*path, curr.value)

	preOrderWalk(curr.left, path)
	preOrderWalk(curr.right, path)

	return *path
}

func (bn *BinaryNode[T]) PreOrderSearch() []int {
	return preOrderWalk(bn, &[]int{})
}

func (bn *BinaryNode[T]) InOrderSearch() []int {
	return inOrderWalk(bn, &[]int{})
}

func (bn *BinaryNode[T]) PostOrderSearch() []int {
	return postOrderWalk(bn, &[]int{})
}
