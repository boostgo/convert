package convert

import (
	"fmt"
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

func String(value any) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case []byte:
		if v == nil {
			return ""
		}

		return StringFromBytes(v)
	case string:
		return v
	case *string:
		if v == nil {
			return ""
		}

		return *v
	case fmt.Stringer:
		if v == nil {
			return ""
		}

		return v.String()
	case error:
		if v == nil {
			return ""
		}

		return v.Error()
	case int8:
		return StringFromInt8(v)
	case *int8:
		if v == nil {
			return ""
		}

		return StringFromInt8(*v)
	case int16:
		return StringFromInt16(v)
	case *int16:
		if v == nil {
			return ""
		}

		return StringFromInt16(*v)
	case int32:
		return StringFromInt32(v)
	case *int32:
		if v == nil {
			return ""
		}

		return StringFromInt32(*v)
	case int64:
		return StringFromInt64(v)
	case *int64:
		if v == nil {
			return ""
		}

		return StringFromInt64(*v)
	case int:
		return StringFromInt(v)
	case *int:
		if v == nil {
			return ""
		}

		return StringFromInt(*v)
	case uint8:
		return StringFromUint8(v)
	case *uint8:
		if v == nil {
			return ""
		}

		return StringFromUint8(*v)
	case uint16:
		return StringFromUint16(v)
	case *uint16:
		if v == nil {
			return ""
		}

		return StringFromUint16(*v)
	case uint32:
		return StringFromUint32(v)
	case *uint32:
		if v == nil {
			return ""
		}

		return StringFromUint32(*v)
	case uint64:
		return StringFromUint64(v)
	case *uint64:
		if v == nil {
			return ""
		}

		return StringFromUint64(*v)
	case uint:
		return StringFromUint(v)
	case *uint:
		if v == nil {
			return ""
		}

		return StringFromUint(*v)
	case float32:
		return StringFromFloat32(v)
	case *float32:
		if v == nil {
			return ""
		}

		return StringFromFloat32(*v)
	case float64:
		return StringFromFloat64(v)
	case *float64:
		if v == nil {
			return ""
		}

		return StringFromFloat64(*v)
	case bool:
		return StringFromBool(v)
	case *bool:
		if v == nil {
			return ""
		}

		return StringFromBool(*v)
	default:
		if value == nil {
			return ""
		}

		return fmt.Sprintf("%v", value)
	}
}
