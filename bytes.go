package convert

import (
	"reflect"
	"unsafe"
)

// BytesFromString converts string to bytes slice with no allocation
func BytesFromString(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
		Len:  len(s),
		Cap:  len(s),
	}))
}
