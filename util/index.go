package util

import (
	"github.com/STRockefeller/readonly"
)

type IIndexGetter interface {
	GetIndex() uint32
}

type Index struct {
	idx readonly.Final[uint32]
}

func (g *Index) Set(idx uint32) {
	g.idx.Set(idx)
}

func (g *Index) GetIndex() uint32 {
	return g.idx.Get()
}
