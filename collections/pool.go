package collections

import "sync"

type IPoolObject interface {
	IsInPool() bool

	setInPool(bool)
}

type PoolObject struct {
	isInPool bool
}

func (p PoolObject) IsInPool() bool {
	return p.isInPool
}

func (p PoolObject) setInPool(b bool) {
	p.isInPool = b
}

type Pool[T IPoolObject] interface {
	Get() T
	Put(t T)

	create(create func() T, onGet func(T), onPut func(T) bool) *pool[T]
}

type pool[T IPoolObject] struct {
	//initOnce sync.Once
	pool  sync.Pool
	onGet func(T)
	onPut func(T) bool
}

func (p *pool[T]) Get() T {
	t := p.pool.Get().(T)
	t.setInPool(false)
	if nil != p.onGet {
		p.onGet(t)
	}
	return t
}

func (p *pool[T]) Put(t T) {
	if nil != p.onPut && false == p.onPut(t) {
		return
	}

	t.setInPool(true)
	p.pool.Put(t)
}

func (p *pool[T]) create(create func() T, onGet func(T), onPut func(T) bool) *pool[T] {
	//var once bool
	//p.initOnce.Do(func() {
	//	once = true
	p.pool.New = func() interface{} {
		return create()
	}
	p.onGet = onGet
	p.onPut = onPut
	//})
	//if once {
	//} else {
	//}
	return p
}

func NewPool[T IPoolObject](create func() T, onGet func(T), onPut func(T) bool) Pool[T] {
	return new(pool[T]).create(create, onGet, onPut)
}
