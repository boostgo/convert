package convert

import (
	"encoding/json"
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

func Bytes(x any) []byte {
	if x == nil {
		return nil
	}

	switch v := x.(type) {
	case []byte:
		return v
	case string:
		return BytesFromString(v)
	case *string:
		if v == nil {
			return nil
		}

		return BytesFromString(*v)
	default:
		blob, err := json.Marshal(x)
		if err != nil {
			return nil
		}

		return blob
	}
}
