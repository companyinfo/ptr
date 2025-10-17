// Package ptr provides helper functions to create and safely dereference pointers.
//
// This package solves common Go problems with pointers:
//   - Go doesn't allow taking the address of literals: &"hello" is invalid
//   - Dereferencing nil pointers causes panics
//   - Working with optional fields in structs and JSON requires verbose code
//
// The package provides both generic functions that work with any type using Go 1.18+
// generics, and type-specific convenience functions for common types (string, int, bool, etc.)
// for better IDE autocomplete and ergonomics.
//
// Basic usage:
//
//	// Create pointers easily
//	name := ptr.String("Alice")
//	age := ptr.Int(30)
//
//	// Safely dereference with zero-value fallback
//	fmt.Println(ptr.ToString(name))  // "Alice"
//	fmt.Println(ptr.ToInt(nil))      // 0
//
//	// Generic functions work with any type
//	user := ptr.To(User{Name: "Bob", Age: 25})
//	fmt.Println(ptr.From(user))
package ptr

import "time"

// To returns a pointer to the provided value.
// This is useful for creating pointers to literals or values in a single expression.
//
// Example:
//
//	s := ptr.To("hello")  // *string
//	i := ptr.To(42)       // *int
func To[T any](v T) *T {
	return &v
}

// From dereferences the pointer and returns its value.
// If the pointer is nil, it returns the zero value of type T.
//
// Example:
//
//	s := ptr.To("hello")
//	v := ptr.From(s)  // "hello"
//	v = ptr.From[string](nil)  // ""
func From[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// FromOr dereferences the pointer and returns its value.
// If the pointer is nil, it returns the provided default value.
//
// Example:
//
//	s := ptr.To("hello")
//	v := ptr.FromOr(s, "default")  // "hello"
//	v = ptr.FromOr[string](nil, "default")  // "default"
func FromOr[T any](p *T, defaultValue T) T {
	if p == nil {
		return defaultValue
	}
	return *p
}

// Equal returns true if both pointers point to equal values.
// Returns true if both pointers are nil.
// Returns false if only one pointer is nil.
//
// Example:
//
//	a := ptr.To(42)
//	b := ptr.To(42)
//	ptr.Equal(a, b)  // true
//	ptr.Equal[int](nil, nil)  // true
func Equal[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// Copy creates a new pointer with a shallow copy of the value.
// For types containing pointers, slices, or maps, only the top-level
// value is copied; nested pointers still reference the same memory.
// Returns nil if the input pointer is nil.
//
// Example:
//
//	original := ptr.To(42)
//	copy := ptr.Copy(original)
//	// copy points to a different memory location
func Copy[T any](p *T) *T {
	if p == nil {
		return nil
	}
	v := *p
	return &v
}

// IsNil returns true if the pointer is nil.
//
// Example:
//
//	s := ptr.To("hello")
//	ptr.IsNil(s)  // false
//	ptr.IsNil[string](nil)  // true
func IsNil[T any](p *T) bool {
	return p == nil
}

// MustFrom dereferences the pointer and returns its value.
// Panics if the pointer is nil. Use this only when nil is a programming error.
//
// Example:
//
//	s := ptr.To("hello")
//	v := ptr.MustFrom(s)  // "hello"
//	v = ptr.MustFrom[string](nil)  // panics
func MustFrom[T any](p *T) T {
	if p == nil {
		panic("ptr: nil pointer passed to MustFrom")
	}
	return *p
}

// Coalesce returns the first non-nil pointer from the provided list.
// Returns nil if all pointers are nil.
//
// Example:
//
//	a := ptr.Coalesce(nil, nil, ptr.To(42), ptr.To(100))  // returns pointer to 42
//	b := ptr.Coalesce[int](nil, nil)  // returns nil
func Coalesce[T any](ptrs ...*T) *T {
	for _, p := range ptrs {
		if p != nil {
			return p
		}
	}
	return nil
}

// Set sets the value of the pointer. If the pointer is nil, it's a no-op.
// Returns true if the value was set, false if the pointer was nil.
//
// Example:
//
//	p := ptr.To(42)
//	ptr.Set(p, 100)  // *p is now 100, returns true
//	ptr.Set[int](nil, 100)  // returns false
func Set[T any](p *T, value T) bool {
	if p == nil {
		return false
	}
	*p = value
	return true
}

// Map applies a transformation function to the pointer value.
// Returns nil if the input pointer is nil.
//
// Example:
//
//	s := ptr.To("hello")
//	length := ptr.Map(s, func(s string) int { return len(s) })  // pointer to 5
//	ptr.Map[string, int](nil, func(s string) int { return len(s) })  // nil
func Map[T, R any](p *T, fn func(T) R) *R {
	if p == nil {
		return nil
	}
	result := fn(*p)
	return &result
}

// ToSlice converts a slice of values to a slice of pointers.
// Returns nil if the input slice is nil.
//
// Example:
//
//	values := []int{1, 2, 3}
//	ptrs := ptr.ToSlice(values)  // []*int with pointers to 1, 2, 3
func ToSlice[T any](values []T) []*T {
	if values == nil {
		return nil
	}
	result := make([]*T, len(values))
	for i := range values {
		result[i] = &values[i]
	}
	return result
}

// FromSlice converts a slice of pointers to a slice of values.
// Nil pointers are converted to zero values.
// Returns nil if the input slice is nil.
//
// Example:
//
//	ptrs := []*int{ptr.To(1), nil, ptr.To(3)}
//	values := ptr.FromSlice(ptrs)  // []int{1, 0, 3}
func FromSlice[T any](ptrs []*T) []T {
	if ptrs == nil {
		return nil
	}
	result := make([]T, len(ptrs))
	for i, p := range ptrs {
		result[i] = From(p)
	}
	return result
}

// String returns a pointer to the provided string value.
func String(v string) *string {
	return To(v)
}

// ToString dereferences a string pointer and returns its value.
// Returns empty string if the pointer is nil.
func ToString(p *string) string {
	return From(p)
}

// Int returns a pointer to the provided int value.
func Int(v int) *int {
	return To(v)
}

// ToInt dereferences an int pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToInt(p *int) int {
	return From(p)
}

// Int64 returns a pointer to the provided int64 value.
func Int64(v int64) *int64 {
	return To(v)
}

// ToInt64 dereferences an int64 pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToInt64(p *int64) int64 {
	return From(p)
}

// Bool returns a pointer to the provided bool value.
func Bool(v bool) *bool {
	return To(v)
}

// ToBool dereferences a bool pointer and returns its value.
// Returns false if the pointer is nil.
func ToBool(p *bool) bool {
	return From(p)
}

// Float64 returns a pointer to the provided float64 value.
func Float64(v float64) *float64 {
	return To(v)
}

// ToFloat64 dereferences a float64 pointer and returns its value.
// Returns 0.0 if the pointer is nil.
func ToFloat64(p *float64) float64 {
	return From(p)
}

// Int32 returns a pointer to the provided int32 value.
func Int32(v int32) *int32 {
	return To(v)
}

// ToInt32 dereferences an int32 pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToInt32(p *int32) int32 {
	return From(p)
}

// Int8 returns a pointer to the provided int8 value.
func Int8(v int8) *int8 {
	return To(v)
}

// ToInt8 dereferences an int8 pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToInt8(p *int8) int8 {
	return From(p)
}

// Int16 returns a pointer to the provided int16 value.
func Int16(v int16) *int16 {
	return To(v)
}

// ToInt16 dereferences an int16 pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToInt16(p *int16) int16 {
	return From(p)
}

// Uint returns a pointer to the provided uint value.
func Uint(v uint) *uint {
	return To(v)
}

// ToUint dereferences a uint pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToUint(p *uint) uint {
	return From(p)
}

// Uint8 returns a pointer to the provided uint8 value.
func Uint8(v uint8) *uint8 {
	return To(v)
}

// ToUint8 dereferences a uint8 pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToUint8(p *uint8) uint8 {
	return From(p)
}

// Uint16 returns a pointer to the provided uint16 value.
func Uint16(v uint16) *uint16 {
	return To(v)
}

// ToUint16 dereferences a uint16 pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToUint16(p *uint16) uint16 {
	return From(p)
}

// Uint32 returns a pointer to the provided uint32 value.
func Uint32(v uint32) *uint32 {
	return To(v)
}

// ToUint32 dereferences a uint32 pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToUint32(p *uint32) uint32 {
	return From(p)
}

// Uint64 returns a pointer to the provided uint64 value.
func Uint64(v uint64) *uint64 {
	return To(v)
}

// ToUint64 dereferences a uint64 pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToUint64(p *uint64) uint64 {
	return From(p)
}

// Float32 returns a pointer to the provided float32 value.
func Float32(v float32) *float32 {
	return To(v)
}

// ToFloat32 dereferences a float32 pointer and returns its value.
// Returns 0.0 if the pointer is nil.
func ToFloat32(p *float32) float32 {
	return From(p)
}

// Byte returns a pointer to the provided byte value.
func Byte(v byte) *byte {
	return To(v)
}

// ToByte dereferences a byte pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToByte(p *byte) byte {
	return From(p)
}

// Rune returns a pointer to the provided rune value.
func Rune(v rune) *rune {
	return To(v)
}

// ToRune dereferences a rune pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToRune(p *rune) rune {
	return From(p)
}

// Time returns a pointer to the provided time.Time value.
func Time(v time.Time) *time.Time {
	return To(v)
}

// ToTime dereferences a time.Time pointer and returns its value.
// Returns zero time if the pointer is nil.
func ToTime(p *time.Time) time.Time {
	return From(p)
}

// Duration returns a pointer to the provided time.Duration value.
func Duration(v time.Duration) *time.Duration {
	return To(v)
}

// ToDuration dereferences a time.Duration pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToDuration(p *time.Duration) time.Duration {
	return From(p)
}

// Complex64 returns a pointer to the provided complex64 value.
func Complex64(v complex64) *complex64 {
	return To(v)
}

// ToComplex64 dereferences a complex64 pointer and returns its value.
// Returns 0+0i if the pointer is nil.
func ToComplex64(p *complex64) complex64 {
	return From(p)
}

// Complex128 returns a pointer to the provided complex128 value.
func Complex128(v complex128) *complex128 {
	return To(v)
}

// ToComplex128 dereferences a complex128 pointer and returns its value.
// Returns 0+0i if the pointer is nil.
func ToComplex128(p *complex128) complex128 {
	return From(p)
}

// Uintptr returns a pointer to the provided uintptr value.
func Uintptr(v uintptr) *uintptr {
	return To(v)
}

// ToUintptr dereferences a uintptr pointer and returns its value.
// Returns 0 if the pointer is nil.
func ToUintptr(p *uintptr) uintptr {
	return From(p)
}
