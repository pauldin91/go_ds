package ds_tests

import (
	"testing"

	ds "github.com/pauldin91/go_ds/src/rbtree"
	"gotest.tools/v3/assert"
)

func TestSearch(t *testing.T) {
	var tree = ds.New[int]()
	tree.Insert(12)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(13)
	tree.Insert(14)

	var exists = tree.Search(2)
	assert.Equal(t, exists, true, "fail")
	exists = tree.Search(5)
	assert.Equal(t, exists, false, "fail")
}

func TestInsert(t *testing.T) {
	var tree = ds.New[int]()
	var actual = make([]ds.RBNode[int], 0)
	for i := 1; i <= 10; i++ {
		tree.Insert(i)
		actual = append(actual, ds.RBNode[int]{Data: i, Color: ds.BLACK})
	}
	actual[7].Color = ds.RED
	actual[9].Color = ds.RED
	var result = tree.InOrder()

	for i := 0; i < 10; i++ {
		assert.Equal(t, actual[i].Color, result[i].Color, "unexpected expected")
		assert.Equal(t, actual[i].Data, result[i].Data, "unexpected expected")
	}
}

func TestDelete(t *testing.T) {
	var tree = ds.New[int]()
	var actual = make([]ds.RBNode[int], 0)
	for i := 1; i <= 10; i++ {
		tree.Insert(i)
	}
	tree.Remove(6)
	tree.Remove(4)
	tree.Remove(10)
	tree.Remove(8)

	actual = append(actual, ds.RBNode[int]{Data: 1, Color: ds.RED})
	actual = append(actual, ds.RBNode[int]{Data: 2, Color: ds.BLACK})
	actual = append(actual, ds.RBNode[int]{Data: 3, Color: ds.BLACK})
	actual = append(actual, ds.RBNode[int]{Data: 5, Color: ds.BLACK})
	actual = append(actual, ds.RBNode[int]{Data: 7, Color: ds.RED})
	actual = append(actual, ds.RBNode[int]{Data: 9, Color: ds.BLACK})

	var result = tree.InOrder()

	for i := 0; i < 6; i++ {
		assert.Equal(t, actual[i].Color, result[i].Color, "unexpected expected")
		assert.Equal(t, actual[i].Data, result[i].Data, "unexpected expected")
	}

}
