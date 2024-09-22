package util

import "math"

const (
	INVALID_SEQ_ID SeqID = 0
	MAX_SEQ_ID           = math.MaxUint32
)

type SeqID uint32

func (sid SeqID) IsValid() bool {
	return sid != INVALID_SEQ_ID
}

func (sid SeqID) IsInvalid() bool {
	return sid == INVALID_SEQ_ID
}

type SeqIDGen interface {
	Reset()
	Get() SeqID

	create() *seqIDGen
}

type seqIDGen struct {
	cur SeqID
}

func (idGen *seqIDGen) Reset() {
	idGen.cur = INVALID_SEQ_ID
}

func (idGen *seqIDGen) Get() SeqID {
	if idGen.cur == MAX_SEQ_ID {
		idGen.cur = INVALID_SEQ_ID
	}
	idGen.cur++
	return idGen.cur
}

func (idGen *seqIDGen) create() *seqIDGen {
	idGen.Reset()
	return idGen
}

func NewSeqIDGen() SeqIDGen {
	return new(seqIDGen).create()
}
