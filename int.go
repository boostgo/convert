package convert

// Int convert any value to int.
//
// If x is nil returns 0
func Int(x any) int {
	switch v := x.(type) {
	case int:
		return v
	case *int:
		if v == nil {
			return 0
		}

		return *v
	case int8:
		return int(v)
	case *int8:
		if v == nil {
			return 0
		}

		return int(*v)
	case int16:
		return int(v)
	case *int16:
		if v == nil {
			return 0
		}

		return int(*v)
	case int32:
		return int(v)
	case *int32:
		if v == nil {
			return 0
		}

		return int(*v)
	case int64:
		return int(v)
	case *int64:
		if v == nil {
			return 0
		}

		return int(*v)
	case uint:
		return int(v)
	case *uint:
		if v == nil {
			return 0
		}

		return int(*v)
	case uint8:
		return int(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return int(*v)
	case uint16:
		return int(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return int(*v)
	case uint32:
		return int(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return int(*v)
	case uint64:
		return int(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return int(*v)
	case float32:
		return int(v)
	case *float32:
		if v == nil {
			return 0
		}

		return int(*v)
	case float64:
		return int(v)
	case *float64:
		if v == nil {
			return 0
		}

		return int(*v)
	default:
		return 0
	}
}

// Int8 convert any value to int8.
//
// If x is nil returns 0
func Int8(x any) int8 {
	switch v := x.(type) {
	case int8:
		return v
	case *int8:
		if v == nil {
			return 0
		}

		return *v
	case int:
		return int8(v)
	case *int:
		if v == nil {
			return 0
		}

		return int8(*v)
	case int16:
		return int8(v)
	case *int16:
		if v == nil {
			return 0
		}

		return int8(*v)
	case int32:
		return int8(v)
	case *int32:
		if v == nil {
			return 0
		}

		return int8(*v)
	case int64:
		return int8(v)
	case *int64:
		if v == nil {
			return 0
		}

		return int8(*v)
	case uint:
		return int8(v)
	case *uint:
		if v == nil {
			return 0
		}

		return int8(*v)
	case uint8:
		return int8(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return int8(*v)
	case uint16:
		return int8(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return int8(*v)
	case uint32:
		return int8(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return int8(*v)
	case uint64:
		return int8(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return int8(*v)
	case float32:
		return int8(v)
	case *float32:
		if v == nil {
			return 0
		}

		return int8(*v)
	case float64:
		return int8(v)
	case *float64:
		if v == nil {
			return 0
		}

		return int8(*v)
	default:
		return 0
	}
}

// Int16 convert any value to int16.
//
// If x is nil returns 0
func Int16(x any) int16 {
	switch v := x.(type) {
	case int16:
		return v
	case *int16:
		if v == nil {
			return 0
		}

		return *v
	case int8:
		return int16(v)
	case *int8:
		if v == nil {
			return 0
		}

		return int16(*v)
	case int:
		return int16(v)
	case *int:
		if v == nil {
			return 0
		}

		return int16(*v)
	case int32:
		return int16(v)
	case *int32:
		if v == nil {
			return 0
		}

		return int16(*v)
	case int64:
		return int16(v)
	case *int64:
		if v == nil {
			return 0
		}

		return int16(*v)
	case uint:
		return int16(v)
	case *uint:
		if v == nil {
			return 0
		}

		return int16(*v)
	case uint8:
		return int16(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return int16(*v)
	case uint16:
		return int16(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return int16(*v)
	case uint32:
		return int16(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return int16(*v)
	case uint64:
		return int16(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return int16(*v)
	case float32:
		return int16(v)
	case *float32:
		if v == nil {
			return 0
		}

		return int16(*v)
	case float64:
		return int16(v)
	case *float64:
		if v == nil {
			return 0
		}

		return int16(*v)
	default:
		return 0
	}
}

// Int32 convert any value to int32.
//
// If x is nil returns 0
func Int32(x any) int32 {
	switch v := x.(type) {
	case int8:
		return int32(v)
	case *int8:
		if v == nil {
			return 0
		}

		return int32(*v)
	case int16:
		return int32(v)
	case *int16:
		if v == nil {
			return 0
		}

		return int32(*v)
	case int32:
		return v
	case *int32:
		if v == nil {
			return 0
		}

		return *v
	case int64:
		return int32(v)
	case *int64:
		if v == nil {
			return 0
		}

		return int32(*v)
	case uint:
		return int32(v)
	case *uint:
		if v == nil {
			return 0
		}

		return int32(*v)
	case uint8:
		return int32(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return int32(*v)
	case uint16:
		return int32(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return int32(*v)
	case uint32:
		return int32(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return int32(*v)
	case uint64:
		return int32(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return int32(*v)
	case float32:
		return int32(v)
	case *float32:
		if v == nil {
			return 0
		}

		return int32(*v)
	case float64:
		return int32(v)
	case *float64:
		if v == nil {
			return 0
		}

		return int32(*v)
	default:
		return 0
	}
}

// Int64 convert any value to int64.
//
// If x is nil returns 0
func Int64(x any) int64 {
	switch v := x.(type) {
	case int8:
		return int64(v)
	case *int8:
		if v == nil {
			return 0
		}

		return int64(*v)
	case int16:
		return int64(v)
	case *int16:
		if v == nil {
			return 0
		}

		return int64(*v)
	case int32:
		return int64(v)
	case *int32:
		if v == nil {
			return 0
		}

		return int64(*v)
	case int64:
		return v
	case *int64:
		if v == nil {
			return 0
		}

		return int64(*v)
	case uint:
		return int64(v)
	case *uint:
		if v == nil {
			return 0
		}

		return int64(*v)
	case uint8:
		return int64(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return int64(*v)
	case uint16:
		return int64(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return int64(*v)
	case uint32:
		return int64(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return int64(*v)
	case uint64:
		return int64(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return int64(*v)
	case float32:
		return int64(v)
	case *float32:
		if v == nil {
			return 0
		}

		return int64(*v)
	case float64:
		return int64(v)
	case *float64:
		if v == nil {
			return 0
		}

		return int64(*v)
	default:
		return 0
	}
}

// Uint convert any value to uint.
//
// If x is nil returns 0
func Uint(x any) uint {
	switch v := x.(type) {
	case int8:
		return uint(v)
	case *int8:
		if v == nil {
			return 0
		}

		return uint(*v)
	case int16:
		return uint(v)
	case *int16:
		if v == nil {
			return 0
		}

		return uint(*v)
	case int32:
		return uint(v)
	case *int32:
		if v == nil {
			return 0
		}

		return uint(*v)
	case int64:
		return uint(v)
	case *int64:
		if v == nil {
			return 0
		}

		return uint(*v)
	case uint:
		return uint(v)
	case *uint:
		if v == nil {
			return 0
		}

		return uint(*v)
	case uint8:
		return uint(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return uint(*v)
	case uint16:
		return uint(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return uint(*v)
	case uint32:
		return uint(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return uint(*v)
	case uint64:
		return uint(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return uint(*v)
	case float32:
		return uint(v)
	case *float32:
		if v == nil {
			return 0
		}

		return uint(*v)
	case float64:
		return uint(v)
	case *float64:
		if v == nil {
			return 0
		}

		return uint(*v)
	default:
		return 0
	}
}

// Uint8 convert any value to uint8.
//
// If x is nil returns 0
func Uint8(x any) uint8 {
	switch v := x.(type) {
	case int8:
		return uint8(v)
	case *int8:
		if v == nil {
			return 0
		}

		return uint8(*v)
	case int16:
		return uint8(v)
	case *int16:
		if v == nil {
			return 0
		}

		return uint8(*v)
	case int32:
		return uint8(v)
	case *int32:
		if v == nil {
			return 0
		}

		return uint8(*v)
	case int64:
		return uint8(v)
	case *int64:
		if v == nil {
			return 0
		}

		return uint8(*v)
	case uint:
		return uint8(v)
	case *uint:
		if v == nil {
			return 0
		}

		return uint8(*v)
	case uint8:
		return uint8(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return uint8(*v)
	case uint16:
		return uint8(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return uint8(*v)
	case uint32:
		return uint8(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return uint8(*v)
	case uint64:
		return uint8(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return uint8(*v)
	case float32:
		return uint8(v)
	case *float32:
		if v == nil {
			return 0
		}

		return uint8(*v)
	case float64:
		return uint8(v)
	case *float64:
		if v == nil {
			return 0
		}

		return uint8(*v)
	default:
		return 0
	}
}

// Uint16 convert any value to uint16.
//
// If x is nil returns 0
func Uint16(x any) uint16 {
	switch v := x.(type) {
	case int8:
		return uint16(v)
	case *int8:
		if v == nil {
			return 0
		}

		return uint16(*v)
	case int16:
		return uint16(v)
	case *int16:
		if v == nil {
			return 0
		}

		return uint16(*v)
	case int32:
		return uint16(v)
	case *int32:
		if v == nil {
			return 0
		}

		return uint16(*v)
	case int64:
		return uint16(v)
	case *int64:
		if v == nil {
			return 0
		}

		return uint16(*v)
	case uint:
		return uint16(v)
	case *uint:
		if v == nil {
			return 0
		}

		return uint16(*v)
	case uint8:
		return uint16(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return uint16(*v)
	case uint16:
		return v
	case *uint16:
		if v == nil {
			return 0
		}

		return *v
	case uint32:
		return uint16(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return uint16(*v)
	case uint64:
		return uint16(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return uint16(*v)
	case float32:
		return uint16(v)
	case *float32:
		if v == nil {
			return 0
		}

		return uint16(*v)
	case float64:
		return uint16(v)
	case *float64:
		if v == nil {
			return 0
		}

		return uint16(*v)
	default:
		return 0
	}
}

// Uint32 convert any value to uint32.
//
// If x is nil returns 0
func Uint32(x any) uint32 {
	switch v := x.(type) {
	case int8:
		return uint32(v)
	case *int8:
		if v == nil {
			return 0
		}

		return uint32(*v)
	case int16:
		return uint32(v)
	case *int16:
		if v == nil {
			return 0
		}

		return uint32(*v)
	case int32:
		return uint32(v)
	case *int32:
		if v == nil {
			return 0
		}

		return uint32(*v)
	case int64:
		return uint32(v)
	case *int64:
		if v == nil {
			return 0
		}

		return uint32(*v)
	case uint:
		return uint32(v)
	case *uint:
		if v == nil {
			return 0
		}

		return uint32(*v)
	case uint8:
		return uint32(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return uint32(*v)
	case uint16:
		return uint32(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return uint32(*v)
	case uint32:
		return v
	case *uint32:
		if v == nil {
			return 0
		}

		return uint32(*v)
	case uint64:
		return uint32(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return uint32(*v)
	case float32:
		return uint32(v)
	case *float32:
		if v == nil {
			return 0
		}

		return uint32(*v)
	case float64:
		return uint32(v)
	case *float64:
		if v == nil {
			return 0
		}

		return uint32(*v)
	default:
		return 0
	}
}

// Uint64 convert any value to uint64.
//
// If x is nil returns 0
func Uint64(x any) uint64 {
	switch v := x.(type) {
	case int8:
		return uint64(v)
	case *int8:
		if v == nil {
			return 0
		}

		return uint64(*v)
	case int16:
		return uint64(v)
	case *int16:
		if v == nil {
			return 0
		}

		return uint64(*v)
	case int32:
		return uint64(v)
	case *int32:
		if v == nil {
			return 0
		}

		return uint64(*v)
	case int64:
		return uint64(v)
	case *int64:
		if v == nil {
			return 0
		}

		return uint64(*v)
	case uint:
		return uint64(v)
	case *uint:
		if v == nil {
			return 0
		}

		return uint64(*v)
	case uint8:
		return uint64(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return uint64(*v)
	case uint16:
		return uint64(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return uint64(*v)
	case uint32:
		return uint64(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return uint64(*v)
	case uint64:
		return v
	case *uint64:
		if v == nil {
			return 0
		}

		return *v
	case float32:
		return uint64(v)
	case *float32:
		if v == nil {
			return 0
		}

		return uint64(*v)
	case float64:
		return uint64(v)
	case *float64:
		if v == nil {
			return 0
		}

		return uint64(*v)
	default:
		return 0
	}
}
