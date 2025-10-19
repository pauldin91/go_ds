package rbtree

import "golang.org/x/exp/constraints"

type RBTree[T constraints.Ordered] struct {
	root *rbTreeNode[T]
}

func New[T constraints.Ordered]() *RBTree[T] {
	return &RBTree[T]{
		root: nil,
	}
}

func (t *RBTree[T]) Insert(data T) {
	t.root = t.root.insert(data)
	t.root.color = BLACK
}

func (t *RBTree[T]) Remove(data T) {
	var ok bool = false
	t.root = t.root.delete(data, &ok)
	t.root.color = BLACK
}

func (t *RBTree[T]) Search(data T) bool {
	return t.root.search(data) != nil
}

func (t *RBTree[T]) InOrder() []RBNode[T] {
	var result = make([]RBNode[T], 0)
	t.root.inOrder(&result)
	return result
}

func (t *RBTree[T]) PreOrder() []RBNode[T] {
	var result = make([]RBNode[T], 0)
	t.root.preOrder(&result)
	return result
}

func (t *RBTree[T]) PostOrder() []RBNode[T] {
	var result = make([]RBNode[T], 0)
	t.root.postOrder(&result)
	return result
}
