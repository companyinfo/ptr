package ptr

import "time"

// StringSlice converts a slice of strings to a slice of string pointers.
func StringSlice(vs []string) []*string {
	return ToSlice(vs)
}

// ToStringSlice converts a slice of string pointers to a slice of strings.
// Nil pointers are converted to empty strings.
func ToStringSlice(vs []*string) []string {
	return FromSlice(vs)
}

// IntSlice converts a slice of ints to a slice of int pointers.
func IntSlice(vs []int) []*int {
	return ToSlice(vs)
}

// ToIntSlice converts a slice of int pointers to a slice of ints.
// Nil pointers are converted to 0.
func ToIntSlice(vs []*int) []int {
	return FromSlice(vs)
}

// Int8Slice converts a slice of int8s to a slice of int8 pointers.
func Int8Slice(vs []int8) []*int8 {
	return ToSlice(vs)
}

// ToInt8Slice converts a slice of int8 pointers to a slice of int8s.
// Nil pointers are converted to 0.
func ToInt8Slice(vs []*int8) []int8 {
	return FromSlice(vs)
}

// Int16Slice converts a slice of int16s to a slice of int16 pointers.
func Int16Slice(vs []int16) []*int16 {
	return ToSlice(vs)
}

// ToInt16Slice converts a slice of int16 pointers to a slice of int16s.
// Nil pointers are converted to 0.
func ToInt16Slice(vs []*int16) []int16 {
	return FromSlice(vs)
}

// Int32Slice converts a slice of int32s to a slice of int32 pointers.
func Int32Slice(vs []int32) []*int32 {
	return ToSlice(vs)
}

// ToInt32Slice converts a slice of int32 pointers to a slice of int32s.
// Nil pointers are converted to 0.
func ToInt32Slice(vs []*int32) []int32 {
	return FromSlice(vs)
}

// Int64Slice converts a slice of int64s to a slice of int64 pointers.
func Int64Slice(vs []int64) []*int64 {
	return ToSlice(vs)
}

// ToInt64Slice converts a slice of int64 pointers to a slice of int64s.
// Nil pointers are converted to 0.
func ToInt64Slice(vs []*int64) []int64 {
	return FromSlice(vs)
}

// UintSlice converts a slice of uints to a slice of uint pointers.
func UintSlice(vs []uint) []*uint {
	return ToSlice(vs)
}

// ToUintSlice converts a slice of uint pointers to a slice of uints.
// Nil pointers are converted to 0.
func ToUintSlice(vs []*uint) []uint {
	return FromSlice(vs)
}

// Uint8Slice converts a slice of uint8s to a slice of uint8 pointers.
func Uint8Slice(vs []uint8) []*uint8 {
	return ToSlice(vs)
}

// ToUint8Slice converts a slice of uint8 pointers to a slice of uint8s.
// Nil pointers are converted to 0.
func ToUint8Slice(vs []*uint8) []uint8 {
	return FromSlice(vs)
}

// Uint16Slice converts a slice of uint16s to a slice of uint16 pointers.
func Uint16Slice(vs []uint16) []*uint16 {
	return ToSlice(vs)
}

// ToUint16Slice converts a slice of uint16 pointers to a slice of uint16s.
// Nil pointers are converted to 0.
func ToUint16Slice(vs []*uint16) []uint16 {
	return FromSlice(vs)
}

// Uint32Slice converts a slice of uint32s to a slice of uint32 pointers.
func Uint32Slice(vs []uint32) []*uint32 {
	return ToSlice(vs)
}

// ToUint32Slice converts a slice of uint32 pointers to a slice of uint32s.
// Nil pointers are converted to 0.
func ToUint32Slice(vs []*uint32) []uint32 {
	return FromSlice(vs)
}

// Uint64Slice converts a slice of uint64s to a slice of uint64 pointers.
func Uint64Slice(vs []uint64) []*uint64 {
	return ToSlice(vs)
}

// ToUint64Slice converts a slice of uint64 pointers to a slice of uint64s.
// Nil pointers are converted to 0.
func ToUint64Slice(vs []*uint64) []uint64 {
	return FromSlice(vs)
}

// Float32Slice converts a slice of float32s to a slice of float32 pointers.
func Float32Slice(vs []float32) []*float32 {
	return ToSlice(vs)
}

// ToFloat32Slice converts a slice of float32 pointers to a slice of float32s.
// Nil pointers are converted to 0.0.
func ToFloat32Slice(vs []*float32) []float32 {
	return FromSlice(vs)
}

// Float64Slice converts a slice of float64s to a slice of float64 pointers.
func Float64Slice(vs []float64) []*float64 {
	return ToSlice(vs)
}

// ToFloat64Slice converts a slice of float64 pointers to a slice of float64s.
// Nil pointers are converted to 0.0.
func ToFloat64Slice(vs []*float64) []float64 {
	return FromSlice(vs)
}

// BoolSlice converts a slice of bools to a slice of bool pointers.
func BoolSlice(vs []bool) []*bool {
	return ToSlice(vs)
}

// ToBoolSlice converts a slice of bool pointers to a slice of bools.
// Nil pointers are converted to false.
func ToBoolSlice(vs []*bool) []bool {
	return FromSlice(vs)
}

// ByteSlice converts a slice of bytes to a slice of byte pointers.
func ByteSlice(vs []byte) []*byte {
	return ToSlice(vs)
}

// ToByteSlice converts a slice of byte pointers to a slice of bytes.
// Nil pointers are converted to 0.
func ToByteSlice(vs []*byte) []byte {
	return FromSlice(vs)
}

// TimeSlice converts a slice of time.Time values to a slice of time.Time pointers.
func TimeSlice(vs []time.Time) []*time.Time {
	return ToSlice(vs)
}

// ToTimeSlice converts a slice of time.Time pointers to a slice of time.Time values.
// Nil pointers are converted to zero time.
func ToTimeSlice(vs []*time.Time) []time.Time {
	return FromSlice(vs)
}

// DurationSlice converts a slice of time.Duration values to a slice of time.Duration pointers.
func DurationSlice(vs []time.Duration) []*time.Duration {
	return ToSlice(vs)
}

// ToDurationSlice converts a slice of time.Duration pointers to a slice of time.Duration values.
// Nil pointers are converted to 0.
func ToDurationSlice(vs []*time.Duration) []time.Duration {
	return FromSlice(vs)
}
