package collections

import (
	"cc_026_fzx/util"
)

const (
	constSparsePageSize = 64
)

type SparseSets[T any] interface {
	Clear()
	Insert(t T) error
	Erase(t T) error
	EraseAt(idx uint32) error
	Contains(t T) bool
	ContainsAt(idx uint32) bool
	IsEmpty() bool
	Size() int
	At(idx uint32) (T, bool)
	getId(t T) uint32
	page(idx uint32) int
	offset(idx uint32) int
	index(idx uint32) *Optional[uint32]
	assure(t T)
}

type page struct {
	array [constSparsePageSize]Optional[uint32]
}

type sparseSets[T any] struct {
	density []T
	sparse  []*page
}

func (ss *sparseSets[T]) Clear() {
	ss.density = ss.density[:0]
	ss.sparse = ss.sparse[:0]
}

func (ss *sparseSets[T]) Insert(t T) error {
	if ss.Contains(t) {
		return util.Err("already contains:", t)
	}

	ss.density = append(ss.density, t)
	ss.assure(t)
	ss.index(ss.getId(t)).Set(uint32(len(ss.density) - 1))

	return nil
}

func (ss *sparseSets[T]) Erase(t T) error {
	return ss.EraseAt(ss.getId(t))
}

func (ss *sparseSets[T]) EraseAt(idx uint32) error {
	if false == ss.ContainsAt(idx) {
		return util.Err("no element:", idx)
	}

	spIdx := ss.index(idx)
	id, _ := spIdx.Get()
	if uint32(len(ss.density)-1) == id {
		spIdx.Clear()
		ss.density = ss.density[:len(ss.density)-1]
	} else {
		last := ss.density[len(ss.density)-1]
		ss.index(ss.getId(last)).Set(id)
		ss.density[id] = last
		spIdx.Clear()
		ss.density = ss.density[:len(ss.density)-1]
	}

	return nil
}

func (ss *sparseSets[T]) Contains(t T) bool {
	return ss.ContainsAt(ss.getId(t))
}

func (ss *sparseSets[T]) ContainsAt(idx uint32) bool {
	p := ss.page(idx)
	if p >= len(ss.sparse) || nil == ss.sparse[p] {
		return false
	}

	spIdx := ss.index(idx)
	return spIdx.IsSet()
}

func (ss *sparseSets[T]) IsEmpty() bool {
	return 0 == len(ss.density)
}

func (ss *sparseSets[T]) Size() int {
	return len(ss.density)
}

func (ss *sparseSets[T]) At(idx uint32) (T, bool) {
	if false == ss.ContainsAt(idx) {
		var zero T
		return zero, false
	}

	spIdx := ss.index(idx)
	id, _ := spIdx.Get()
	return ss.density[id], true
}

func (ss *sparseSets[T]) getId(t T) uint32 {
	var inter interface{} = t
	switch v := inter.(type) {
	case util.IIndexGetter:
		if nil != v {
			return v.GetIndex()
		}
	case uint32:
		return v
	}

	panic("invalid type for sparseSets[T]")
}

func (ss *sparseSets[T]) page(idx uint32) int {
	return int(idx / constSparsePageSize)
}

func (ss *sparseSets[T]) offset(idx uint32) int {
	return int(idx % constSparsePageSize)
}

func (ss *sparseSets[T]) index(idx uint32) *Optional[uint32] {
	p := ss.sparse[ss.page(idx)]
	return &p.array[ss.offset(idx)]
}

func (ss *sparseSets[T]) assure(t T) {
	p := ss.page(ss.getId(t))
	if p >= len(ss.sparse) {
		for i := len(ss.sparse); i < p; i++ {
			ss.sparse = append(ss.sparse, nil)
		}
		ss.sparse = append(ss.sparse, &page{})
	} else if nil == ss.sparse[p] {
		ss.sparse[p] = &page{}
	}
}

func NewSparseSetsIndexGetter[T util.IIndexGetter]() SparseSets[T] {
	return &sparseSets[T]{}
}

func NewSparseSetsUint32() SparseSets[uint32] {
	return &sparseSets[uint32]{}
}
