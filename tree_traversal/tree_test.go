package tree

import (
	"reflect"
	"testing"
)

func TestTreeTraversal(t *testing.T) {
	tree := &BinaryNode[int]{
		value: 20,
		right: &BinaryNode[int]{
			value: 50,
			right: &BinaryNode[int]{
				value: 100,
			},
			left: &BinaryNode[int]{
				value: 30,
				right: &BinaryNode[int]{
					value: 45,
				},
				left: &BinaryNode[int]{
					value: 29,
				},
			},
		},
		left: &BinaryNode[int]{
			value: 10,
			right: &BinaryNode[int]{
				value: 15,
			},
			left: &BinaryNode[int]{
				value: 5,
				right: &BinaryNode[int]{
					value: 7,
				},
			},
		},
	}

	expectedResult := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	preOrderSearchResult := tree.PreOrderSearch()

	if !reflect.DeepEqual(preOrderSearchResult, expectedResult) {
		t.Errorf("Pre-order traversal result is incorrect. Got %v, expected %v", preOrderSearchResult, expectedResult)
	}

	expectedResult = []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	inOrderSearchResult := tree.InOrderSearch()

	if !reflect.DeepEqual(inOrderSearchResult, expectedResult) {
		t.Errorf("In-order traversal result is incorrect. Got %v, expected %v", inOrderSearchResult, expectedResult)
	}

	expectedResult = []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	postOrderSearchResult := tree.PostOrderSearch()

	if !reflect.DeepEqual(postOrderSearchResult, expectedResult) {
		t.Errorf("Post-order traversal result is incorrect. Got %v, expected %v", postOrderSearchResult, expectedResult)
	}
}
