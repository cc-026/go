package collections

type BitArray interface {
	RemoveAll()
	AddAll()
	AddFlag(flag int)
	RemoveFlag(flag int)
	HasFlag(flag int) bool
	IsEmpty() bool
	Size() int
	AddFlagFromBoolArray(array []bool)
	AddFlagFromBitArray(array BitArray)
	Clone() BitArray
}

type bitArray32 uint32

func (b *bitArray32) RemoveAll() {
	*b = 0
}

func (b *bitArray32) AddAll() {
	var zero bitArray32 = 0
	*b = ^zero
}

func (b *bitArray32) AddFlag(flag int) {
	if flag >= b.Size() {
		return
	}
	*b |= 1 << flag
}

func (b *bitArray32) RemoveFlag(flag int) {
	if flag >= b.Size() {
		return
	}
	*b &= ^(1 << flag)
}

func (b *bitArray32) HasFlag(flag int) bool {
	if flag >= b.Size() {
		return false
	}
	return (*b & (1 << flag)) != 0

}

func (b *bitArray32) IsEmpty() bool {
	return *b == 0
}

func (b *bitArray32) Size() int {
	return 32
}

func (b *bitArray32) AddFlagFromBoolArray(array []bool) {
	if nil == array {
		return
	}

	for i := 0; i < b.Size() && i < len(array); i++ {
		if array[i] {
			b.AddFlag(i)
		}
	}
}

func (b *bitArray32) AddFlagFromBitArray(array BitArray) {
	if nil == array {
		return
	}

	for i := 0; i < b.Size() && i < array.Size(); i++ {
		if array.HasFlag(i) {
			b.AddFlag(i)
		}
	}
}

func (b *bitArray32) Clone() BitArray {
	clone := *b
	return &clone
}

type bitArray64 uint64

func (b *bitArray64) RemoveAll() {
	*b = 0
}

func (b *bitArray64) AddAll() {
	var zero bitArray64 = 0
	*b = ^zero
}

func (b *bitArray64) AddFlag(flag int) {
	if flag >= b.Size() {
		return
	}
	*b |= 1 << flag
}

func (b *bitArray64) RemoveFlag(flag int) {
	if flag >= b.Size() {
		return
	}
	*b &= ^(1 << flag)
}

func (b *bitArray64) HasFlag(flag int) bool {
	if flag >= b.Size() {
		return false
	}
	return (*b & (1 << flag)) != 0

}

func (b *bitArray64) IsEmpty() bool {
	return *b == 0
}

func (b *bitArray64) Size() int {
	return 32
}

func (b *bitArray64) AddFlagFromBoolArray(array []bool) {
	if nil == array {
		return
	}

	for i := 0; i < b.Size() && i < len(array); i++ {
		if array[i] {
			b.AddFlag(i)
		}
	}
}

func (b *bitArray64) AddFlagFromBitArray(array BitArray) {
	if nil == array {
		return
	}

	for i := 0; i < b.Size() && i < array.Size(); i++ {
		if array.HasFlag(i) {
			b.AddFlag(i)
		}
	}
}

func (b *bitArray64) Clone() BitArray {
	clone := *b
	return &clone
}

type bitArray struct {
	array []uint64
}

func (b *bitArray) RemoveAll() {
	for i := 0; i < len(b.array); i++ {
		b.array[i] = 0
	}
}

func (b *bitArray) AddAll() {
	var zero uint64 = 0
	for i := 0; i < len(b.array); i++ {
		b.array[i] = ^zero
	}
}

func (b *bitArray) AddFlag(flag int) {
	idx := flag / 64
	if idx >= len(b.array) {
		return
	}
	b.array[idx] |= 1 << (flag % 64)
}

func (b *bitArray) RemoveFlag(flag int) {
	idx := flag / 64
	if idx >= len(b.array) {
		return
	}
	b.array[idx] &= ^(1 << (flag % 64))
}

func (b *bitArray) HasFlag(flag int) bool {
	idx := flag / 64
	if idx >= len(b.array) {
		return false
	}

	return (b.array[idx] & (1 << (flag % 64))) != 0
}

func (b *bitArray) IsEmpty() bool {
	for i := 0; i < len(b.array); i++ {
		if b.array[i] != 0 {
			return false
		}
	}

	return true
}

func (b *bitArray) Size() int {
	return len(b.array) * 64
}

func (b *bitArray) AddFlagFromBoolArray(array []bool) {
	if nil == array {
		return
	}

	for i := 0; i < b.Size() && i < len(array); i++ {
		if array[i] {
			b.AddFlag(i)
		}
	}
}

func (b *bitArray) AddFlagFromBitArray(array BitArray) {
	if nil == array {
		return
	}

	for i := 0; i < b.Size() && i < array.Size(); i++ {
		if array.HasFlag(i) {
			b.AddFlag(i)
		}
	}
}

func (b *bitArray) Clone() BitArray {
	c := newBitArray(b.Size())
	for i := 0; i < len(b.array); i++ {
		c.array[i] = b.array[i]
	}

	return c
}

func NewBitArray(size int) BitArray {
	if size <= 32 {
		return new(bitArray32)
	} else if size <= 64 {
		return new(bitArray64)
	}

	return newBitArray(size)
}

func newBitArray(size int) *bitArray {
	l := size / 64
	if size%64 != 0 {
		l++
	}
	return &bitArray{
		array: make([]uint64, l),
	}
}
