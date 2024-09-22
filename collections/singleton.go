package collections

import (
	"cc_026_fzx/util"
	//"cc_026_fzx/util"
	//"math/rand"
	"sync"
	"sync/atomic"
	//"time"
)

type ISingleton interface {
	Init()
	//GetInstance() interface{}

	singleton()
}

type iSingletonPtr[T any] interface {
	*T
	ISingleton
}

var (
	mutex  sync.Mutex
	insMap = make(map[interface{}]interface{})
	//insMap = make(map[uint32]interface{})
	//insSMap sync.Map
	writer atomic.Bool
)

type Singleton struct {
	_ util.NoCopy
}

//func (*Singleton) GetInstance() interface{} {
//	return nil
//}

func (*Singleton) Init() {}

func (*Singleton) singleton() {}

func Instance[T any, Ptr iSingletonPtr[T]]() *T {
	//return insOfRawGet((Ptr)(nil)).(*T)
	return insOfInnerGet[T, Ptr]().(*T)
}

//	func insOfRawGet(i ISingleton) interface{} {
//		return i.GetInstance()
//	}
//var i1 = 0
//var i2 = 0

func insOfInnerGet[T any, Ptr iSingletonPtr[T]]() interface{} {
	//i := typeFor[T]()
	//if t, has := insSMap.Load(i); has && nil != t {
	//	return t
	//} else {
	//	t = new(T)
	//	insSMap.Store(i, t)
	//	return t
	//}
	key := (*T)(nil)
	if writer.Load() {
		//util.Log().LogDebug("InWrite", i2)
		//i2++
		mutex.Lock()
		defer mutex.Unlock()
	}

	if t, has := insMap[key]; has && nil != t {
		return t
	} else {
		//if false == writer.Load() {
		mutex.Lock()
		//r := rand.New(rand.NewSource(time.Now().UnixNano()))
		//s1 := time.Duration((r.Int63()%1000)*2) * time.Millisecond
		//time.Sleep(s1)
		writer.Store(true)
		//s2 := time.Duration((r.Int63()%1000)*2) * time.Millisecond
		//time.Sleep(s2)
		//util.Log().LogDebug("Try Create", i1, s1, s2)
		//i1++

		defer func() {
			writer.Store(false)
			mutex.Unlock()
		}()
		//}

		if t, has = insMap[key]; false == has || nil == t {
			//util.Log().LogDebug("InCreate")
			t = new(T)
			t.(ISingleton).Init()
			insMap[key] = t
		}
		return t
	}
}

//func typeFor[T any]() uint32 {
//	return typeOf((*T)(nil))
//}
//
//func typeOf(i any) uint32 {
//	eface := (*emptyInterface)(unsafe.Pointer(&i))
//	return eface.typ.Hash
//}
//
//type Type struct {
//	Size_       uintptr
//	PtrBytes    uintptr // number of (prefix) bytes in the type that can contain pointers
//	Hash        uint32  // hash of type; avoids computation in hash tables
//	TFlag       uint8   // extra type information flags
//	Align_      uint8   // alignment of variable with this type
//	FieldAlign_ uint8   // alignment of struct field with this type
//	Kind_       uint8   // enumeration for C
//	// function for comparing objects of this type
//	// (ptr to object A, ptr to object B) -> ==?
//	Equal func(unsafe.Pointer, unsafe.Pointer) bool
//	// GCData stores the GC type data for the garbage collector.
//	// If the KindGCProg bit is set in kind, GCData is a GC program.
//	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
//	GCData    *byte
//	Str       int32 // string form
//	PtrToThis int32 // type for pointer to this type, may be zero
//}
//
//type emptyInterface struct {
//	typ  *Type
//	word unsafe.Pointer
//}
