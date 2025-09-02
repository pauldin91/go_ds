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
		tree.Insert(i)
		assert.Equal(t, actual[i].Color, result[i].Color, "unexpected expected")
		assert.Equal(t, actual[i].Data, result[i].Data, "unexpected expected")
	}
}

func TestDelete(t *testing.T) {
	var tree = ds.New[int]()
	tree.Insert(12)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(13)
	tree.Insert(14)

}

func TestInOrder(t *testing.T) {

	var tree = ds.New[int]()
	tree.Insert(12)
	tree.Insert(6)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(16)
	tree.Insert(14)
	tree.Insert(18)
	var expected = []int{4, 6, 8, 12, 14, 16, 18}
	var actual = tree.InOrder()
	for i := range actual {
		assert.Equal(t, actual[i].Data, expected[i], "fail")
	}
}
