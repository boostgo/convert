package convert

// BoolFromString convert string to bool
func BoolFromString(x string) bool {
	return x == "true" || x == "TRUE"
}

// BoolFromInt convert any integer type to bool
func BoolFromInt[T Integer](x T) bool {
	return x == 1
}
