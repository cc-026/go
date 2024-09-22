package collections

type Optional[T any] struct {
	isSet bool
	value T
}

func (op *Optional[T]) Set(v T) {
	op.isSet = true
	op.value = v
}

func (op *Optional[T]) Get() (T, bool) {
	return op.value, op.isSet
}

func (op *Optional[T]) Clear() {
	var zero T
	op.isSet = false
	op.value = zero
}

func (op *Optional[T]) IsSet() bool {
	return op.isSet
}
