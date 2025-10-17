package ptr

import "time"

// ToMap converts a map with value type T to a map with pointer value type *T.
// Returns nil if the input map is nil.
//
// Example:
//
//	values := map[string]int{"a": 1, "b": 2}
//	ptrs := ptr.ToMap(values)  // map[string]*int
func ToMap[T any](values map[string]T) map[string]*T {
	if values == nil {
		return nil
	}
	result := make(map[string]*T, len(values))
	for k, v := range values {
		v := v // Create new variable to take address of
		result[k] = &v
	}
	return result
}

// FromMap converts a map with pointer value type *T to a map with value type T.
// Nil pointers are converted to zero values.
// Returns nil if the input map is nil.
//
// Example:
//
//	ptrs := map[string]*int{"a": ptr.To(1), "b": nil}
//	values := ptr.FromMap(ptrs)  // map[string]int{"a": 1, "b": 0}
func FromMap[T any](ptrs map[string]*T) map[string]T {
	if ptrs == nil {
		return nil
	}
	result := make(map[string]T, len(ptrs))
	for k, p := range ptrs {
		result[k] = From(p)
	}
	return result
}

// StringMap converts a map of strings to a map of string pointers.
func StringMap(vs map[string]string) map[string]*string {
	return ToMap(vs)
}

// ToStringMap converts a map of string pointers to a map of strings.
// Nil pointers are converted to empty strings.
func ToStringMap(vs map[string]*string) map[string]string {
	return FromMap(vs)
}

// IntMap converts a map of ints to a map of int pointers.
func IntMap(vs map[string]int) map[string]*int {
	return ToMap(vs)
}

// ToIntMap converts a map of int pointers to a map of ints.
// Nil pointers are converted to 0.
func ToIntMap(vs map[string]*int) map[string]int {
	return FromMap(vs)
}

// Int8Map converts a map of int8s to a map of int8 pointers.
func Int8Map(vs map[string]int8) map[string]*int8 {
	return ToMap(vs)
}

// ToInt8Map converts a map of int8 pointers to a map of int8s.
// Nil pointers are converted to 0.
func ToInt8Map(vs map[string]*int8) map[string]int8 {
	return FromMap(vs)
}

// Int16Map converts a map of int16s to a map of int16 pointers.
func Int16Map(vs map[string]int16) map[string]*int16 {
	return ToMap(vs)
}

// ToInt16Map converts a map of int16 pointers to a map of int16s.
// Nil pointers are converted to 0.
func ToInt16Map(vs map[string]*int16) map[string]int16 {
	return FromMap(vs)
}

// Int32Map converts a map of int32s to a map of int32 pointers.
func Int32Map(vs map[string]int32) map[string]*int32 {
	return ToMap(vs)
}

// ToInt32Map converts a map of int32 pointers to a map of int32s.
// Nil pointers are converted to 0.
func ToInt32Map(vs map[string]*int32) map[string]int32 {
	return FromMap(vs)
}

// Int64Map converts a map of int64s to a map of int64 pointers.
func Int64Map(vs map[string]int64) map[string]*int64 {
	return ToMap(vs)
}

// ToInt64Map converts a map of int64 pointers to a map of int64s.
// Nil pointers are converted to 0.
func ToInt64Map(vs map[string]*int64) map[string]int64 {
	return FromMap(vs)
}

// UintMap converts a map of uints to a map of uint pointers.
func UintMap(vs map[string]uint) map[string]*uint {
	return ToMap(vs)
}

// ToUintMap converts a map of uint pointers to a map of uints.
// Nil pointers are converted to 0.
func ToUintMap(vs map[string]*uint) map[string]uint {
	return FromMap(vs)
}

// Uint8Map converts a map of uint8s to a map of uint8 pointers.
func Uint8Map(vs map[string]uint8) map[string]*uint8 {
	return ToMap(vs)
}

// ToUint8Map converts a map of uint8 pointers to a map of uint8s.
// Nil pointers are converted to 0.
func ToUint8Map(vs map[string]*uint8) map[string]uint8 {
	return FromMap(vs)
}

// Uint16Map converts a map of uint16s to a map of uint16 pointers.
func Uint16Map(vs map[string]uint16) map[string]*uint16 {
	return ToMap(vs)
}

// ToUint16Map converts a map of uint16 pointers to a map of uint16s.
// Nil pointers are converted to 0.
func ToUint16Map(vs map[string]*uint16) map[string]uint16 {
	return FromMap(vs)
}

// Uint32Map converts a map of uint32s to a map of uint32 pointers.
func Uint32Map(vs map[string]uint32) map[string]*uint32 {
	return ToMap(vs)
}

// ToUint32Map converts a map of uint32 pointers to a map of uint32s.
// Nil pointers are converted to 0.
func ToUint32Map(vs map[string]*uint32) map[string]uint32 {
	return FromMap(vs)
}

// Uint64Map converts a map of uint64s to a map of uint64 pointers.
func Uint64Map(vs map[string]uint64) map[string]*uint64 {
	return ToMap(vs)
}

// ToUint64Map converts a map of uint64 pointers to a map of uint64s.
// Nil pointers are converted to 0.
func ToUint64Map(vs map[string]*uint64) map[string]uint64 {
	return FromMap(vs)
}

// Float32Map converts a map of float32s to a map of float32 pointers.
func Float32Map(vs map[string]float32) map[string]*float32 {
	return ToMap(vs)
}

// ToFloat32Map converts a map of float32 pointers to a map of float32s.
// Nil pointers are converted to 0.0.
func ToFloat32Map(vs map[string]*float32) map[string]float32 {
	return FromMap(vs)
}

// Float64Map converts a map of float64s to a map of float64 pointers.
func Float64Map(vs map[string]float64) map[string]*float64 {
	return ToMap(vs)
}

// ToFloat64Map converts a map of float64 pointers to a map of float64s.
// Nil pointers are converted to 0.0.
func ToFloat64Map(vs map[string]*float64) map[string]float64 {
	return FromMap(vs)
}

// BoolMap converts a map of bools to a map of bool pointers.
func BoolMap(vs map[string]bool) map[string]*bool {
	return ToMap(vs)
}

// ToBoolMap converts a map of bool pointers to a map of bools.
// Nil pointers are converted to false.
func ToBoolMap(vs map[string]*bool) map[string]bool {
	return FromMap(vs)
}

// ByteMap converts a map of bytes to a map of byte pointers.
func ByteMap(vs map[string]byte) map[string]*byte {
	return ToMap(vs)
}

// ToByteMap converts a map of byte pointers to a map of bytes.
// Nil pointers are converted to 0.
func ToByteMap(vs map[string]*byte) map[string]byte {
	return FromMap(vs)
}

// TimeMap converts a map of time.Time values to a map of time.Time pointers.
func TimeMap(vs map[string]time.Time) map[string]*time.Time {
	return ToMap(vs)
}

// ToTimeMap converts a map of time.Time pointers to a map of time.Time values.
// Nil pointers are converted to zero time.
func ToTimeMap(vs map[string]*time.Time) map[string]time.Time {
	return FromMap(vs)
}

// DurationMap converts a map of time.Duration values to a map of time.Duration pointers.
func DurationMap(vs map[string]time.Duration) map[string]*time.Duration {
	return ToMap(vs)
}

// ToDurationMap converts a map of time.Duration pointers to a map of time.Duration values.
// Nil pointers are converted to 0.
func ToDurationMap(vs map[string]*time.Duration) map[string]time.Duration {
	return FromMap(vs)
}
