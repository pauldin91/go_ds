package go_ds

import (
	ds "github.com/pauldin91/go_ds/src/rbtree"
	"golang.org/x/exp/constraints"
)

type RBTree[T constraints.Ordered] ds.RBTree[T]
type RBNode[T constraints.Ordered] ds.RBNode[T]
