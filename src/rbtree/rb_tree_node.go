package rbtree

import "golang.org/x/exp/constraints"

type rbTreeNode[T constraints.Ordered] struct {
	data  T
	child []*rbTreeNode[T]
	color NodeColor
}

func new[T constraints.Ordered](data T) *rbTreeNode[T] {
	return &rbTreeNode[T]{
		data:  data,
		child: make([]*rbTreeNode[T], 2),
		color: RED,
	}
}

func (x *rbTreeNode[T]) isRed() bool {
	return x != nil && x.color == RED
}

func (t *rbTreeNode[T]) insert(data T) *rbTreeNode[T] {
	if t == nil {
		return new(data)
	}
	var dir Direction = LEFT
	if data > t.data {
		dir = RIGHT
	}
	t.child[dir] = t.child[dir].insert(data)
	return t.insertRebalance(dir)
}

func (t *rbTreeNode[T]) insertRebalance(dir Direction) *rbTreeNode[T] {
	if t.child[dir].isRed() {
		//siblings are RED
		if t.child[1-dir].isRed() {
			//exists grandchild RED
			if t.child[dir].child[dir].isRed() ||
				t.child[dir].child[1-dir].isRed() {
				t.colorFlip()
			}
		} else {
			if t.child[dir].child[dir].isRed() {
				t = t.rotate(1 - dir)
			} else if t.child[dir].child[1-dir].isRed() {
				t = t.doubleRotate(1 - dir)
			}
		}
	}
	return t
}

func (t *rbTreeNode[T]) search(data T) *rbTreeNode[T] {
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

func (t *rbTreeNode[T]) delete(data T, ok *bool) *rbTreeNode[T] {
	if t == nil {
		*ok = true
		return t
	}
	if t.data == data {
		if t.child[LEFT] == nil || t.child[RIGHT] == nil {

			var temp *rbTreeNode[T] = nil
			if t.child[LEFT] != nil {
				temp = t.child[LEFT]
			}
			if t.child[RIGHT] != nil {
				temp = t.child[RIGHT]
			}
			if t.isRed() {
				*ok = true
			} else if temp.isRed() {
				temp.color = BLACK
				*ok = true
			}
			t = nil
			return temp
		} else {
			var tt []*rbTreeNode[T] = make([]*rbTreeNode[T], 0)
			t.child[LEFT].edge(&tt, RIGHT)
			var temp *rbTreeNode[T] = tt[len(tt)-1]
			t.data = temp.data
			data = temp.data
		}
	}

	var dir Direction = LEFT
	if data > t.data {
		dir = RIGHT
	}
	t.child[dir] = t.child[dir].delete(data, ok)

	if !*ok {
		return t.deleteRebalance(dir, ok)
	}

	return t
}

func (t *rbTreeNode[T]) deleteRebalance(dir Direction, ok *bool) *rbTreeNode[T] {
	var parent *rbTreeNode[T] = t
	var sibling *rbTreeNode[T] = t.child[1-dir]

	if sibling.isRed() {
		t = t.rotate(dir)
		sibling = parent.child[1-dir]
	}

	if sibling != nil {
		if !sibling.child[LEFT].isRed() && !sibling.child[RIGHT].isRed() {
			if parent.isRed() {
				*ok = true
			}
			parent.color = BLACK
			sibling.color = RED

		} else {

			var parentColor NodeColor = parent.color
			var isRedSiblingReduction bool = !(t == parent)

			if sibling.child[1-dir].isRed() {
				parent = parent.rotate(dir)
			} else {
				parent = parent.doubleRotate(dir)
			}

			parent.color = parentColor
			parent.child[LEFT].color = BLACK
			parent.child[RIGHT].color = BLACK

			if isRedSiblingReduction {
				t.child[dir] = parent
			} else {
				t = parent
			}
			*ok = true
		}
	}

	return t

}

func (t *rbTreeNode[T]) edge(path *[]*rbTreeNode[T], dir Direction) {
	if t.child[dir] == nil {
		*path = append(*path, t)
	} else {
		t.child[dir].edge(path, dir)
	}
}

func (x *rbTreeNode[T]) rotate(dir Direction) *rbTreeNode[T] {
	y := x.child[1-dir]
	x.child[1-dir] = y.child[dir]
	y.child[dir] = x
	y.color = x.color
	x.color = RED
	return y
}

func (x *rbTreeNode[T]) doubleRotate(dir Direction) *rbTreeNode[T] {
	x.child[1-dir] = x.child[1-dir].rotate(1 - dir)
	return x.rotate(dir)
}

func (x *rbTreeNode[T]) colorFlip() {
	x.color = 1 - x.color
	x.child[LEFT].color = 1 - x.child[LEFT].color
	x.child[RIGHT].color = 1 - x.child[RIGHT].color
}
