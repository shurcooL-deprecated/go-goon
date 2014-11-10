// Package bypass allows bypassing reflect restrictions on accessing unexported struct fields.
package bypass

import (
	"reflect"
	"unsafe"
)

const (
	// ptrSize is the size of a pointer on the current arch.
	ptrSize = unsafe.Sizeof((*byte)(nil))
)

var (
	// offsetPtr, offsetScalar, and offsetFlag are the offsets for the
	// internal reflect.Value fields.  These values are valid before golang
	// commit ecccf07e7f9d which changed the format.  The are also valid
	// after commit 82f48826c6c7 which changed the format again to mirror
	// the original format.  Code in the init function updates these offsets
	// as necessary.
	offsetPtr    = uintptr(ptrSize)
	offsetScalar = uintptr(0)
	offsetFlag   = uintptr(ptrSize * 2)

	// flagKindWidth and flagKindShift indicate various bits that the
	// reflect package uses internally to track kind information.
	//
	// flagRO indicates whether or not the value field of a reflect.Value is
	// read-only.
	//
	// flagIndir indicates whether the value field of a reflect.Value is
	// the actual data or a pointer to the data.
	//
	// These values are valid before golang commit 90a7c3c86944 which
	// changed their positions.  Code in the init function updates these
	// flags as necessary.
	flagKindWidth = uintptr(5)
	flagKindShift = uintptr(flagKindWidth - 1)
	flagRO        = uintptr(1 << 0)
	flagIndir     = uintptr(1 << 1)
)

func init() {
	// Older versions of reflect.Value stored small integers directly in the
	// ptr field (which is named val in the older versions).  Versions
	// between commits ecccf07e7f9d and 82f48826c6c7 added a new field named
	// scalar for this purpose which unfortunately came before the flag
	// field, so the offset of the flag field is different for those
	// versions.
	//
	// This code constructs a new reflect.Value from a known small integer
	// and checks if the size of the reflect.Value struct indicates it has
	// the scalar field. When it does, the offsets are updated accordingly.
	vv := reflect.ValueOf(0xf00)
	if unsafe.Sizeof(vv) == (ptrSize * 4) {
		offsetScalar = ptrSize * 2
		offsetFlag = ptrSize * 3
	}

	// Commit 90a7c3c86944 changed the flag positions such that the low
	// order bits are the kind.  This code extracts the kind from the flags
	// field and ensures it's the correct type.  When it's not, the flag
	// order has been changed to the newer format, so the flags are updated
	// accordingly.
	upf := unsafe.Pointer(uintptr(unsafe.Pointer(&vv)) + offsetFlag)
	upfv := *(*uintptr)(upf)
	flagKindMask := uintptr((1<<flagKindWidth - 1) << flagKindShift)
	if (upfv&flagKindMask)>>flagKindShift != uintptr(reflect.Int) {
		flagKindShift = 0
		flagRO = 1 << 5
		flagIndir = 1 << 6
	}
}

// ---

// Considering using /Users/Dmitri/Dropbox/Work/2013/GoLand/src/github.com/shurcooL/play/8/main.go approach,
// just modify the reflect.Value.flag's flagRO bit.

var flagValOffset = func() uintptr {
	field, ok := reflect.TypeOf(reflect.Value{}).FieldByName("flag")
	if !ok {
		panic("reflect.Value has no flag field")
	}
	return field.Offset
}()

type flag uintptr

func bypass(v reflect.Value) reflect.Value {
	if !v.IsValid() || v.CanInterface() {
		return v
	}
	flagp := (*flag)(unsafe.Pointer(uintptr(unsafe.Pointer(&v)) + flagValOffset))
	*flagp &^= (flag)(flagRO)
	return v
}

func UnsafeReflectValueC(v reflect.Value) (rv reflect.Value) {
	return bypass(v)
}
