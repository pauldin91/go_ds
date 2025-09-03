package go_ds

import (
	rbtree "github.com/pauldin91/go_ds/src/rbtree"
	"golang.org/x/exp/constraints"
)

type RBTree[T constraints.Ordered] rbtree.RBTree[T]
type RBNode[T constraints.Ordered] rbtree.RBNode[T]
