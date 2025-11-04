package avltree

import (
	"golang.org/x/exp/constraints"
)

type Direction int

const (
	LEFT  Direction = 0
	RIGHT Direction = 1
)

type avlTreeNode[T constraints.Ordered] struct {
	data   T
	child  []*avlTreeNode[T]
	height int
}

func (t *avlTreeNode[T]) search(data T) *avlTreeNode[T] {
	if t == nil || t.data == data {
		return t
	} else {
		var dir Direction = LEFT
		if data > t.data {
			dir = 1 - dir
		}
		return t.child[dir].search(data)
	}
}

func (x *avlTreeNode[T]) rotate(dir Direction) *avlTreeNode[T] {
	y := x.child[1-dir]
	x.child[1-dir] = y.child[dir]
	y.child[dir] = x
	return y
}

func (x *avlTreeNode[T]) doubleRotate(dir Direction) *avlTreeNode[T] {
	x.child[1-dir] = x.child[1-dir].rotate(1 - dir)
	return x.rotate(dir)
}

func (t *avlTreeNode[T]) edge(path *[]*avlTreeNode[T], dir Direction) {
	if t.child[dir] == nil {
		*path = append(*path, t)
	} else {
		t.child[dir].edge(path, dir)
	}
}
