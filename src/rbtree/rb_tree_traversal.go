package ds

import (
	"golang.org/x/exp/constraints"
)

type RBNode[T constraints.Ordered] struct {
	Data  T
	Color NodeColor
}

func newRBNode[T constraints.Ordered](d T, c NodeColor) RBNode[T] {
	return RBNode[T]{
		Data:  d,
		Color: c,
	}
}

func (t *rbTreeNode[T]) inOrder(list *[]RBNode[T]) {
	if t.child[LEFT] != nil {
		t.child[LEFT].inOrder(list)
	}
	var node = newRBNode(t.data, t.color)
	*list = append(*list, node)
	if t.child[RIGHT] != nil {
		t.child[RIGHT].inOrder(list)
	}
}

func (t *rbTreeNode[T]) postOrder(list *[]RBNode[T]) {
	if t.child[LEFT] != nil {
		t.child[LEFT].postOrder(list)
	}
	if t.child[RIGHT] != nil {
		t.child[RIGHT].postOrder(list)
	}

	var node = newRBNode(t.data, t.color)
	*list = append(*list, node)
}

func (t *rbTreeNode[T]) preOrder(list *[]RBNode[T]) {

	var node = newRBNode(t.data, t.color)
	*list = append(*list, node)
	if t.child[LEFT] != nil {
		t.child[LEFT].preOrder(list)
	}
	if t.child[RIGHT] != nil {
		t.child[RIGHT].preOrder(list)
	}
}
