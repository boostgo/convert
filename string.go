package convert

import (
	"strconv"
	"unsafe"
)

// StringFromBytes convert bytes array/slice to string with no allocation.
func StringFromBytes(buffer []byte) string {
	return *(*string)(unsafe.Pointer(&buffer))
}

// StringFromInt8 convert int8 to string
func StringFromInt8(integer int8) string {
	return strconv.Itoa(int(integer))
}

// StringFromInt16 convert int16 to string
func StringFromInt16(integer int16) string {
	return strconv.Itoa(int(integer))
}

// StringFromInt32 convert int32 to string
func StringFromInt32(integer int32) string {
	return strconv.Itoa(int(integer))
}

// StringFromInt64 convert int64 to string
func StringFromInt64(integer int64) string {
	return strconv.FormatInt(integer, 10)
}

// StringFromInt convert int to string
func StringFromInt(integer int) string {
	return strconv.Itoa(integer)
}

// StringFromUint8 convert uint8 to string
func StringFromUint8(integer uint8) string {
	return strconv.FormatUint(uint64(integer), 10)
}

// StringFromUint16 convert uint16 to string
func StringFromUint16(integer uint16) string {
	return strconv.FormatUint(uint64(integer), 10)
}

// StringFromUint32 convert uint32 to string
func StringFromUint32(integer uint32) string {
	return strconv.FormatUint(uint64(integer), 10)
}

// StringFromUint64 convert uint64 to string
func StringFromUint64(integer uint64) string {
	return strconv.FormatUint(integer, 10)
}

// StringFromUint convert uint to string
func StringFromUint(integer uint) string {
	return strconv.FormatUint(uint64(integer), 10)
}

// StringFromFloat32 convert float32 to string
func StringFromFloat32(integer float32) string {
	return strconv.FormatFloat(float64(integer), 'E', -1, 32)
}

// StringFromFloat64 convert float64 to string
func StringFromFloat64(integer float64) string {
	return strconv.FormatFloat(integer, 'E', -1, 32)
}

// StringFromBool convert bool to string
func StringFromBool(flag bool) string {
	const (
		t = "true"
		f = "false"
	)

	if flag {
		return t
	}

	return f
}
