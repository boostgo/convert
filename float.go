package convert

// Float32FromInt convert any integer type value to float32
func Float32FromInt[T Integer](x T) float32 {
	return float32(x)
}

// Float64FromInt convert any integer type value to float64
func Float64FromInt[T Integer](x T) float64 {
	return float64(x)
}

// Float32 convert any value to float32.
//
// If value is nil return 0.
func Float32(x any) float32 {
	switch v := x.(type) {
	case float32:
		return v
	case *float32:
		if v == nil {
			return 0
		}

		return *v
	case float64:
		return float32(v)
	case *float64:
		if v == nil {
			return 0
		}

		return float32(*v)
	case int:
		return float32(v)
	case *int:
		if v == nil {
			return 0
		}

		return float32(*v)
	case int8:
		return float32(v)
	case *int8:
		if v == nil {
			return 0
		}

		return float32(*v)
	case int16:
		return float32(v)
	case *int16:
		if v == nil {
			return 0
		}

		return float32(*v)
	case int32:
		return float32(v)
	case *int32:
		if v == nil {
			return 0
		}

		return float32(*v)
	case int64:
		return float32(v)
	case *int64:
		if v == nil {
			return 0
		}

		return float32(*v)
	case uint:
		return float32(v)
	case *uint:
		if v == nil {
			return 0
		}

		return float32(*v)
	case uint8:
		return float32(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return float32(*v)
	case uint16:
		return float32(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return float32(*v)
	case uint32:
		return float32(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return float32(*v)
	case uint64:
		return float32(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return float32(*v)
	default:
		return 0
	}
}

// Float64 convert any value to float32.
//
// If value is nil return 0.
func Float64(x any) float64 {
	switch v := x.(type) {
	case float32:
		return float64(v)
	case *float32:
		if v == nil {
			return 0
		}

		return float64(*v)
	case float64:
		return v
	case *float64:
		if v == nil {
			return 0
		}

		return *v
	case int:
		return float64(v)
	case *int:
		if v == nil {
			return 0
		}

		return float64(*v)
	case int8:
		return float64(v)
	case *int8:
		if v == nil {
			return 0
		}

		return float64(*v)
	case int16:
		return float64(v)
	case *int16:
		if v == nil {
			return 0
		}

		return float64(*v)
	case int32:
		return float64(v)
	case *int32:
		if v == nil {
			return 0
		}

		return float64(*v)
	case int64:
		return float64(v)
	case *int64:
		if v == nil {
			return 0
		}

		return float64(*v)
	case uint:
		return float64(v)
	case *uint:
		if v == nil {
			return 0
		}

		return float64(*v)
	case uint8:
		return float64(v)
	case *uint8:
		if v == nil {
			return 0
		}

		return float64(*v)
	case uint16:
		return float64(v)
	case *uint16:
		if v == nil {
			return 0
		}

		return float64(*v)
	case uint32:
		return float64(v)
	case *uint32:
		if v == nil {
			return 0
		}

		return float64(*v)
	case uint64:
		return float64(v)
	case *uint64:
		if v == nil {
			return 0
		}

		return float64(*v)
	default:
		return 0
	}
}
