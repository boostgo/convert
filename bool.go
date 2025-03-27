package convert

// BoolFromString convert string to bool
func BoolFromString(input string) bool {
	return input == "true" || input == "TRUE"
}

// BoolFromInt convert any integer type to bool
func BoolFromInt[T Integer](input T) bool {
	return input == 1
}
