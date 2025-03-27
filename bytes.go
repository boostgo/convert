package convert

import (
	"reflect"
	"unsafe"
)

// BytesFromString converts string to bytes slice with no allocation
func BytesFromString(x string) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&x))).Data,
		Len:  len(x),
		Cap:  len(x),
	}))
}
