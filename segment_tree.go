package segment

import (
	"errors"
	"math"
)

type Tree interface {
	Update(idx, value int) error
	QuerySum(begin, end int) (int, error)
}

type TreeImpl struct {
	init bool
	data []int
	len  int
	mh   int
	t    []int
}

// NewSegmentTree build a segment tree from an existing int array, exp: []int{1, 3, 5, 7, 9}
func NewSegmentTree(data []int) *TreeImpl {
	st := &TreeImpl{data: data, len: len(data)}

	st.mh = int(math.Log2(float64(st.len))) + 2
	st.t = make([]int, int(math.Pow(2, float64(st.mh))-1))
	st.build(0, 0, st.len-1, -1, 0)
	st.init = true

	return st
}

func (st *TreeImpl) build(node, left, right int, idx int, value int) int {
	if left == right {
		if idx < 0 {
			st.t[node] = st.data[left]
		} else {
			st.t[node] = value
			st.data[idx] = value
		}
	} else {
		mid := (left + right) / 2
		st.t[node] = st.build(node*2+1, left, mid, idx, value) + st.build(node*2+2, mid+1, right, idx, value)
	}

	return st.t[node]
}

// Update changes original array value, and rebuild tree
func (st *TreeImpl) Update(idx, value int) error {
	if idx < 0 {
		return errors.New("idx should greater than or equal to zero")
	}

	st.update(idx, value, 0, 0, st.len-1)

	return nil
}

func (st *TreeImpl) update(idx, value, node, left, right int) {
	if left == right {
		st.data[idx] = value
		st.t[node] = value
	} else {
		mid := (left + right) / 2
		if idx >= left && idx <= mid {
			st.update(idx, value, node*2+1, left, mid)
		} else {
			st.update(idx, value, node*2+2, mid+1, right)
		}

		st.t[node] = st.t[node*2+1] + st.t[node*2+2]
	}
}

// QuerySum returns sum of data[begin:end]
func (st *TreeImpl) QuerySum(begin, end int) (int, error) {
	return st.sum(0, 0, st.len-1, begin, end), nil
}

func (st *TreeImpl) sum(node, l, r, begin, end int) int {
	if r < begin || l > end {
		return 0
	}

	if l >= begin && r <= end {
		return st.t[node]
	}

	mid := (l + r) / 2
	return st.sum(2*node+1, l, mid, begin, end) + st.sum(2*node+2, mid+1, r, begin, end)
}
