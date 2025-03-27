# `github.com/boostgo/convert`

# Get started

```go
package main

import (
	"fmt"

	"github.com/boostgo/convert"
)

func main() {
	// string
	fmt.Println(convert.StringFromBool(true))                   // "true"
	fmt.Println(convert.StringFromBool(false))                  // "false"
	fmt.Println(convert.StringFromInt(123))                     // "123"
	fmt.Println(convert.StringFromInt64(777))                   // "777"
	fmt.Println(convert.StringFromFloat32(1.23))                // "1.230000"
	fmt.Println(convert.StringFromFloat64(1.23))                // "1.23"
	fmt.Println(convert.StringFromBytes([]byte("hello world"))) // "hello world"

	// bytes
	fmt.Println(convert.BytesFromString("hello world")) // [104 101 108 108 111 32 119 111 114 108 100]
	fmt.Println(convert.Bytes(nil))                     // []
	fmt.Println(convert.StringFromBytes(convert.Bytes(map[string]any{
		"key1": "value1",
		"key2": "value2",
	}))) // {"key1":"value1","key2":"value2"}

	// int
	fmt.Println(convert.Int(true))    // 1
	fmt.Println(convert.Int(1.23))    // 1
	fmt.Println(convert.Int(1.63))    // 1
	fmt.Println(convert.Int("777"))   // 777
	fmt.Println(convert.Int64("777")) // 777

	// float
	fmt.Println(convert.Float32FromString("1.23")) // 1.23
	fmt.Println(convert.Float64FromString("1.23")) // 1.23
	fmt.Println(convert.Float32(999))              // 999

	// bool
	fmt.Println(convert.BoolFromString("true")) // true
	fmt.Println(convert.BoolFromString("TRUE")) // true
	fmt.Println(convert.BoolFromString("TRue")) // false
	fmt.Println(convert.BoolFromInt(1))         // true
	fmt.Println(convert.BoolFromInt(2))         // false
	fmt.Println(convert.BoolFromInt(0))         // false
}

```

### String

Convert to string from:
- **bytes** - StringFromBytes(x []byte) string
- **int8** - StringFromInt8(x int8) string
- **int16** - StringFromInt16(x int16) string
- **int32** - StringFromInt32(x int32) string
- **int64** - StringFromInt64(x int64) string
- **int** - StringFromInt(x int) string
- **uint8** - StringFromUint8(x uint8) string
- **uint16** - StringFromUint16(x uint16) string
- **uint32** - StringFromUint32(x uint32) string
- **uint64** - StringFromUint64(x uint64) string
- **uint** - StringFromUint(x uint) string
- **float32** - StringFromFloat64(x float32) string
- **float64** - StringFromFloat32(x float64) string
- **bool** - StringFromBool(x bool) string
- **any** - String(x any) string

### Bytes

Convert to bytes from:
- **string** - BytesFromString(x string) []byte
- **any** - Bytes(x any) []byte

### Int

Convert to any integer type from:
- **bool** - IntFromBool(x bool) int
- **bool** - Int8FromBool(x bool) int8
- **bool** - Int16FromBool(x bool) int16
- **bool** - Int32FromBool(x bool) int32
- **bool** - Int64FromBool(x bool) int64
- **bool** - UintFromBool(x bool) uint
- **bool** - Uint8FromBool(x bool) uint8
- **bool** - Uint16FromBool(x bool) uint16
- **bool** - Uint32FromBool(x bool) uint32
- **bool** - Uint64FromBool(x bool) uint64
- **string** - IntFromString(x string) int
- **string** - Int8FromString(x string) int8
- **string** - Int16FromString(x string) int16
- **string** - Int32FromString(x string) int32
- **string** - Int64FromString(x string) int64
- **string** - UintFromString(x string) uint
- **string** - Uint8FromString(x string) uint8
- **string** - Uint16FromString(x string) uint16
- **string** - Uint32FromString(x string) uint32
- **string** - Uint64FromString(x string) uint64
- **any** - Int(x any) int
- **any** - Int8(x any) int8
- **any** - Int16(x any) int16
- **any** - Int32(x any) int32
- **any** - Int64(x any) int64
- **any** - Uint(x any) uint
- **any** - Uint8(x any) uint8
- **any** - Uint16(x any) uint16
- **any** - Uint32(x any) uint32
- **any** - Uint64(x any) uint64

### Float

Convert to float32/float64 from:
- **[any integer]** - Float32FromInt(x Integer) float32
- **[any integer]** - Float64FromInt(x Integer) float64
- **string** - Float32FromString(x string) float32
- **string** - Float64FromString(x string) float64
- **any** - Float32(x any) float32
- **any** - Float64(x any) float64

### Bool

Convert to bool from:
- **string** - BoolFromString(x string) bool
- **[any integer]** - BoolFromInt(x Integer) bool
- **any** - Bool(x any) bool
