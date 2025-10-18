package ptr

import (
	"testing"
	"time"
)

// Test generic To function
func TestTo(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		v := "hello"
		p := To(v)
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != v {
			t.Errorf("expected %q, got %q", v, *p)
		}
	})

	t.Run("int", func(t *testing.T) {
		v := 42
		p := To(v)
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != v {
			t.Errorf("expected %d, got %d", v, *p)
		}
	})

	t.Run("bool", func(t *testing.T) {
		v := true
		p := To(v)
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != v {
			t.Errorf("expected %v, got %v", v, *p)
		}
	})

	t.Run("struct", func(t *testing.T) {
		type testStruct struct {
			Name string
			Age  int
		}
		v := testStruct{Name: "Alice", Age: 30}
		p := To(v)
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != v {
			t.Errorf("expected %+v, got %+v", v, *p)
		}
	})
}

// Test generic From function
func TestFrom(t *testing.T) {
	t.Run("string non-nil", func(t *testing.T) {
		v := "hello"
		p := &v
		result := From(p)
		if result != v {
			t.Errorf("expected %q, got %q", v, result)
		}
	})

	t.Run("string nil", func(t *testing.T) {
		result := From[string](nil)
		if result != "" {
			t.Errorf("expected empty string, got %q", result)
		}
	})

	t.Run("int non-nil", func(t *testing.T) {
		v := 42
		p := &v
		result := From(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("int nil", func(t *testing.T) {
		result := From[int](nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})

	t.Run("bool non-nil", func(t *testing.T) {
		v := true
		p := &v
		result := From(p)
		if result != v {
			t.Errorf("expected %v, got %v", v, result)
		}
	})

	t.Run("bool nil", func(t *testing.T) {
		result := From[bool](nil)
		if result != false {
			t.Errorf("expected false, got %v", result)
		}
	})
}

// Test generic FromOr function
func TestFromOr(t *testing.T) {
	t.Run("string non-nil", func(t *testing.T) {
		v := "hello"
		p := &v
		result := FromOr(p, "default")
		if result != v {
			t.Errorf("expected %q, got %q", v, result)
		}
	})

	t.Run("string nil", func(t *testing.T) {
		defaultVal := "default"
		result := FromOr[string](nil, defaultVal)
		if result != defaultVal {
			t.Errorf("expected %q, got %q", defaultVal, result)
		}
	})

	t.Run("int non-nil", func(t *testing.T) {
		v := 42
		p := &v
		result := FromOr(p, 100)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("int nil", func(t *testing.T) {
		defaultVal := 100
		result := FromOr[int](nil, defaultVal)
		if result != defaultVal {
			t.Errorf("expected %d, got %d", defaultVal, result)
		}
	})
}

// Test Equal function
func TestEqual(t *testing.T) {
	t.Run("both nil", func(t *testing.T) {
		if !Equal[string](nil, nil) {
			t.Error("expected true for two nil pointers")
		}
	})

	t.Run("first nil", func(t *testing.T) {
		v := "hello"
		if Equal[string](nil, &v) {
			t.Error("expected false when first is nil")
		}
	})

	t.Run("second nil", func(t *testing.T) {
		v := "hello"
		if Equal(&v, nil) {
			t.Error("expected false when second is nil")
		}
	})

	t.Run("equal values", func(t *testing.T) {
		v1 := "hello"
		v2 := "hello"
		if !Equal(&v1, &v2) {
			t.Error("expected true for equal values")
		}
	})

	t.Run("different values", func(t *testing.T) {
		v1 := "hello"
		v2 := "world"
		if Equal(&v1, &v2) {
			t.Error("expected false for different values")
		}
	})

	t.Run("same pointer", func(t *testing.T) {
		v := "hello"
		p := &v
		if !Equal(p, p) {
			t.Error("expected true for same pointer")
		}
	})
}

// Test Copy function
func TestCopy(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		result := Copy[string](nil)
		if result != nil {
			t.Error("expected nil for nil input")
		}
	})

	t.Run("string copy", func(t *testing.T) {
		v := "hello"
		original := &v
		copied := Copy(original)

		if copied == nil {
			t.Fatal("expected non-nil copy")
		}
		if *copied != *original {
			t.Errorf("expected %q, got %q", *original, *copied)
		}
		if copied == original {
			t.Error("expected different pointer addresses")
		}
	})

	t.Run("int copy", func(t *testing.T) {
		v := 42
		original := &v
		copied := Copy(original)

		if copied == nil {
			t.Fatal("expected non-nil copy")
		}
		if *copied != *original {
			t.Errorf("expected %d, got %d", *original, *copied)
		}
		if copied == original {
			t.Error("expected different pointer addresses")
		}
	})

	t.Run("modification independence", func(t *testing.T) {
		v := 42
		original := &v
		copied := Copy(original)

		*original = 100
		if *copied == *original {
			t.Error("copy should not be affected by original modification")
		}
		if *copied != 42 {
			t.Errorf("expected copied value to remain 42, got %d", *copied)
		}
	})
}

// Test IsNil function
func TestIsNil(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		if !IsNil[string](nil) {
			t.Error("expected true for nil pointer")
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		v := "hello"
		if IsNil(&v) {
			t.Error("expected false for non-nil pointer")
		}
	})
}

// Test type-specific String functions
func TestString(t *testing.T) {
	v := "hello"
	p := String(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %q, got %q", v, *p)
	}
}

func TestToString(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := "hello"
		p := &v
		result := ToString(p)
		if result != v {
			t.Errorf("expected %q, got %q", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToString(nil)
		if result != "" {
			t.Errorf("expected empty string, got %q", result)
		}
	})
}

// Test type-specific Int functions
func TestInt(t *testing.T) {
	v := 42
	p := Int(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToInt(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := 42
		p := &v
		result := ToInt(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToInt(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Int64 functions
func TestInt64(t *testing.T) {
	v := int64(42)
	p := Int64(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToInt64(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := int64(42)
		p := &v
		result := ToInt64(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToInt64(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Bool functions
func TestBool(t *testing.T) {
	v := true
	p := Bool(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %v, got %v", v, *p)
	}
}

func TestToBool(t *testing.T) {
	t.Run("non-nil true", func(t *testing.T) {
		v := true
		p := &v
		result := ToBool(p)
		if result != v {
			t.Errorf("expected %v, got %v", v, result)
		}
	})

	t.Run("non-nil false", func(t *testing.T) {
		v := false
		p := &v
		result := ToBool(p)
		if result != v {
			t.Errorf("expected %v, got %v", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToBool(nil)
		if result != false {
			t.Errorf("expected false, got %v", result)
		}
	})
}

// Test type-specific Float64 functions
func TestFloat64(t *testing.T) {
	v := 3.14
	p := Float64(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %f, got %f", v, *p)
	}
}

func TestToFloat64(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := 3.14
		p := &v
		result := ToFloat64(p)
		if result != v {
			t.Errorf("expected %f, got %f", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToFloat64(nil)
		if result != 0.0 {
			t.Errorf("expected 0.0, got %f", result)
		}
	})
}

// Example tests for documentation

func ExampleTo() {
	// Create pointers to literals
	s := To("hello")
	i := To(42)
	b := To(true)

	println(*s) // hello
	println(*i) // 42
	println(*b) // true
}

func ExampleFrom() {
	// Dereference with zero-value fallback
	s := To("hello")
	println(From(s)) // hello

	var nilStr *string
	println(From(nilStr)) // "" (empty string)
}

func ExampleFromOr() {
	// Dereference with custom default
	s := To("hello")
	println(FromOr(s, "default")) // hello

	var nilStr *string
	println(FromOr(nilStr, "default")) // default
}

func ExampleEqual() {
	a := To(42)
	b := To(42)
	c := To(99)

	println(Equal(a, b))          // true
	println(Equal(a, c))          // false
	println(Equal[int](nil, nil)) // true
}

func ExampleCopy() {
	original := To(42)
	copied := Copy(original)

	*original = 100
	println(*copied) // 42 (unchanged)
}

func ExampleIsNil() {
	s := To("hello")
	var nilStr *string

	println(IsNil(s))      // false
	println(IsNil(nilStr)) // true
}

func ExampleMustFrom() {
	s := To("hello")
	println(MustFrom(s)) // hello

	// MustFrom will panic if passed nil
	// var nilStr *string
	// println(MustFrom(nilStr)) // panics!
}

func ExampleString() {
	// Convenient way to create string pointers
	name := String("Alice")
	println(*name) // Alice
}

func ExampleToString() {
	name := String("Alice")
	println(ToString(name)) // Alice

	var nilName *string
	println(ToString(nilName)) // "" (empty string)
}

func ExampleInt() {
	age := Int(30)
	println(*age) // 30
}

func ExampleToInt() {
	age := Int(30)
	println(ToInt(age)) // 30

	var nilAge *int
	println(ToInt(nilAge)) // 0
}

func ExampleBool() {
	active := Bool(true)
	println(*active) // true
}

func ExampleToBool() {
	active := Bool(true)
	println(ToBool(active)) // true

	var nilBool *bool
	println(ToBool(nilBool)) // false
}

// Test MustFrom function
func TestMustFrom(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := "hello"
		p := &v
		result := MustFrom(p)
		if result != v {
			t.Errorf("expected %q, got %q", v, result)
		}
	})

	t.Run("nil panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for nil pointer")
			}
		}()
		_ = MustFrom[string](nil)
	})

	t.Run("int non-nil", func(t *testing.T) {
		v := 42
		p := &v
		result := MustFrom(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})
}

// Test type-specific Int8 functions
func TestInt8(t *testing.T) {
	v := int8(42)
	p := Int8(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToInt8(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := int8(42)
		p := &v
		result := ToInt8(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToInt8(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Int16 functions
func TestInt16(t *testing.T) {
	v := int16(42)
	p := Int16(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToInt16(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := int16(42)
		p := &v
		result := ToInt16(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToInt16(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Int32 functions
func TestInt32(t *testing.T) {
	v := int32(42)
	p := Int32(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToInt32(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := int32(42)
		p := &v
		result := ToInt32(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToInt32(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Uint8 functions
func TestUint8(t *testing.T) {
	v := uint8(42)
	p := Uint8(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToUint8(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := uint8(42)
		p := &v
		result := ToUint8(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToUint8(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Uint16 functions
func TestUint16(t *testing.T) {
	v := uint16(42)
	p := Uint16(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToUint16(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := uint16(42)
		p := &v
		result := ToUint16(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToUint16(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Uint32 functions
func TestUint32(t *testing.T) {
	v := uint32(42)
	p := Uint32(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToUint32(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := uint32(42)
		p := &v
		result := ToUint32(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToUint32(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Uint functions
func TestUint(t *testing.T) {
	v := uint(42)
	p := Uint(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToUint(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := uint(42)
		p := &v
		result := ToUint(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToUint(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Uint64 functions
func TestUint64(t *testing.T) {
	v := uint64(42)
	p := Uint64(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToUint64(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := uint64(42)
		p := &v
		result := ToUint64(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToUint64(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Float32 functions
func TestFloat32(t *testing.T) {
	v := float32(3.14)
	p := Float32(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %f, got %f", v, *p)
	}
}

func TestToFloat32(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := float32(3.14)
		p := &v
		result := ToFloat32(p)
		if result != v {
			t.Errorf("expected %f, got %f", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToFloat32(nil)
		if result != 0.0 {
			t.Errorf("expected 0.0, got %f", result)
		}
	})
}

// Test type-specific Byte functions
func TestByte(t *testing.T) {
	v := byte('A')
	p := Byte(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToByte(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := byte('A')
		p := &v
		result := ToByte(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToByte(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test type-specific Rune functions
func TestRune(t *testing.T) {
	v := rune('世')
	p := Rune(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToRune(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := rune('世')
		p := &v
		result := ToRune(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToRune(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test Coalesce function
func TestCoalesce(t *testing.T) {
	t.Run("all nil", func(t *testing.T) {
		result := Coalesce[int](nil, nil, nil)
		if result != nil {
			t.Error("expected nil for all nil pointers")
		}
	})

	t.Run("first non-nil", func(t *testing.T) {
		v := 42
		result := Coalesce(&v, nil, nil)
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if *result != v {
			t.Errorf("expected %d, got %d", v, *result)
		}
	})

	t.Run("middle non-nil", func(t *testing.T) {
		v := 42
		result := Coalesce[int](nil, &v, To(100))
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if *result != v {
			t.Errorf("expected %d, got %d", v, *result)
		}
	})

	t.Run("multiple non-nil returns first", func(t *testing.T) {
		v1 := 42
		v2 := 100
		result := Coalesce(&v1, &v2)
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if *result != v1 {
			t.Errorf("expected %d, got %d", v1, *result)
		}
	})

	t.Run("empty variadic", func(t *testing.T) {
		result := Coalesce[int]()
		if result != nil {
			t.Error("expected nil for empty variadic")
		}
	})
}

// Test Set function
func TestSet(t *testing.T) {
	t.Run("set non-nil", func(t *testing.T) {
		v := 42
		p := &v
		result := Set(p, 100)
		if !result {
			t.Error("expected true for successful set")
		}
		if *p != 100 {
			t.Errorf("expected 100, got %d", *p)
		}
	})

	t.Run("set nil", func(t *testing.T) {
		result := Set[int](nil, 100)
		if result {
			t.Error("expected false for nil pointer")
		}
	})

	t.Run("set string", func(t *testing.T) {
		v := "hello"
		p := &v
		result := Set(p, "world")
		if !result {
			t.Error("expected true for successful set")
		}
		if *p != "world" {
			t.Errorf("expected 'world', got %q", *p)
		}
	})
}

// Test Map function
func TestMap(t *testing.T) {
	t.Run("map non-nil", func(t *testing.T) {
		s := "hello"
		p := &s
		result := Map(p, func(s string) int { return len(s) })
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if *result != 5 {
			t.Errorf("expected 5, got %d", *result)
		}
	})

	t.Run("map nil", func(t *testing.T) {
		result := Map[string, int](nil, func(s string) int { return len(s) })
		if result != nil {
			t.Error("expected nil for nil input")
		}
	})

	t.Run("map int to string", func(t *testing.T) {
		v := 42
		p := &v
		result := Map(p, func(i int) string { return "hello" })
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if *result != "hello" {
			t.Errorf("expected 'hello', got %q", *result)
		}
	})

	t.Run("map to same type", func(t *testing.T) {
		v := 10
		p := &v
		result := Map(p, func(i int) int { return i * 2 })
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if *result != 20 {
			t.Errorf("expected 20, got %d", *result)
		}
	})
}

// Test ToSlice function
func TestToSlice(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		result := ToSlice[int](nil)
		if result != nil {
			t.Error("expected nil for nil input")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		result := ToSlice([]int{})
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if len(result) != 0 {
			t.Errorf("expected length 0, got %d", len(result))
		}
	})

	t.Run("int slice", func(t *testing.T) {
		values := []int{1, 2, 3}
		result := ToSlice(values)
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if len(result) != 3 {
			t.Fatalf("expected length 3, got %d", len(result))
		}
		for i, p := range result {
			if p == nil {
				t.Errorf("expected non-nil pointer at index %d", i)
				continue
			}
			if *p != values[i] {
				t.Errorf("expected %d at index %d, got %d", values[i], i, *p)
			}
		}
	})

	t.Run("string slice", func(t *testing.T) {
		values := []string{"a", "b", "c"}
		result := ToSlice(values)
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if len(result) != 3 {
			t.Fatalf("expected length 3, got %d", len(result))
		}
		for i, p := range result {
			if p == nil {
				t.Errorf("expected non-nil pointer at index %d", i)
				continue
			}
			if *p != values[i] {
				t.Errorf("expected %q at index %d, got %q", values[i], i, *p)
			}
		}
	})
}

// Test FromSlice function
func TestFromSlice(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		result := FromSlice[int](nil)
		if result != nil {
			t.Error("expected nil for nil input")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		result := FromSlice([]*int{})
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if len(result) != 0 {
			t.Errorf("expected length 0, got %d", len(result))
		}
	})

	t.Run("all non-nil", func(t *testing.T) {
		ptrs := []*int{To(1), To(2), To(3)}
		result := FromSlice(ptrs)
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if len(result) != 3 {
			t.Fatalf("expected length 3, got %d", len(result))
		}
		expected := []int{1, 2, 3}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("expected %d at index %d, got %d", expected[i], i, v)
			}
		}
	})

	t.Run("with nil pointers", func(t *testing.T) {
		ptrs := []*int{To(1), nil, To(3)}
		result := FromSlice(ptrs)
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if len(result) != 3 {
			t.Fatalf("expected length 3, got %d", len(result))
		}
		expected := []int{1, 0, 3}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("expected %d at index %d, got %d", expected[i], i, v)
			}
		}
	})

	t.Run("string slice", func(t *testing.T) {
		ptrs := []*string{To("a"), nil, To("c")}
		result := FromSlice(ptrs)
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if len(result) != 3 {
			t.Fatalf("expected length 3, got %d", len(result))
		}
		expected := []string{"a", "", "c"}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("expected %q at index %d, got %q", expected[i], i, v)
			}
		}
	})
}

// Test type-specific Time functions
func TestTime(t *testing.T) {
	v := time.Now()
	p := Time(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if !p.Equal(v) {
		t.Errorf("expected %v, got %v", v, *p)
	}
}

func TestToTime(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := time.Now()
		p := &v
		result := ToTime(p)
		if !result.Equal(v) {
			t.Errorf("expected %v, got %v", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToTime(nil)
		if !result.IsZero() {
			t.Errorf("expected zero time, got %v", result)
		}
	})
}

// Test type-specific Duration functions
func TestDuration(t *testing.T) {
	v := time.Second * 5
	p := Duration(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %v, got %v", v, *p)
	}
}

func TestToDuration(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := time.Minute * 30
		p := &v
		result := ToDuration(p)
		if result != v {
			t.Errorf("expected %v, got %v", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToDuration(nil)
		if result != 0 {
			t.Errorf("expected 0, got %v", result)
		}
	})
}

// Test type-specific Complex64 functions
func TestComplex64(t *testing.T) {
	v := complex64(3 + 4i)
	p := Complex64(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %v, got %v", v, *p)
	}
}

func TestToComplex64(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := complex64(3 + 4i)
		p := &v
		result := ToComplex64(p)
		if result != v {
			t.Errorf("expected %v, got %v", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToComplex64(nil)
		if result != 0 {
			t.Errorf("expected 0+0i, got %v", result)
		}
	})
}

// Test type-specific Complex128 functions
func TestComplex128(t *testing.T) {
	v := complex(3, 4)
	p := Complex128(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %v, got %v", v, *p)
	}
}

func TestToComplex128(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := complex(3, 4)
		p := &v
		result := ToComplex128(p)
		if result != v {
			t.Errorf("expected %v, got %v", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToComplex128(nil)
		if result != 0 {
			t.Errorf("expected 0+0i, got %v", result)
		}
	})
}

// Test type-specific Uintptr functions
func TestUintptr(t *testing.T) {
	v := uintptr(42)
	p := Uintptr(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestToUintptr(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		v := uintptr(42)
		p := &v
		result := ToUintptr(p)
		if result != v {
			t.Errorf("expected %d, got %d", v, result)
		}
	})

	t.Run("nil", func(t *testing.T) {
		result := ToUintptr(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})
}

// Test Or function
func TestOr(t *testing.T) {
	t.Run("first non-nil", func(t *testing.T) {
		a := To(42)
		b := To(100)
		result := Or(a, b)
		if result != a {
			t.Error("expected first pointer")
		}
		if *result != 42 {
			t.Errorf("expected 42, got %d", *result)
		}
	})

	t.Run("first nil, second non-nil", func(t *testing.T) {
		var a *int
		b := To(100)
		result := Or(a, b)
		if result != b {
			t.Error("expected second pointer")
		}
		if *result != 100 {
			t.Errorf("expected 100, got %d", *result)
		}
	})

	t.Run("both nil", func(t *testing.T) {
		var a, b *int
		result := Or(a, b)
		if result != nil {
			t.Error("expected nil")
		}
	})

	t.Run("both non-nil, returns first", func(t *testing.T) {
		a := To("first")
		b := To("second")
		result := Or(a, b)
		if result != a {
			t.Error("expected first pointer")
		}
		if *result != "first" {
			t.Errorf("expected 'first', got %q", *result)
		}
	})
}

// Test Filter function
func TestFilter(t *testing.T) {
	t.Run("predicate true", func(t *testing.T) {
		p := To(42)
		result := Filter(p, func(v int) bool { return v > 40 })
		if result != p {
			t.Error("expected same pointer")
		}
		if *result != 42 {
			t.Errorf("expected 42, got %d", *result)
		}
	})

	t.Run("predicate false", func(t *testing.T) {
		p := To(42)
		result := Filter(p, func(v int) bool { return v > 50 })
		if result != nil {
			t.Error("expected nil")
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		var p *int
		called := false
		result := Filter(p, func(v int) bool {
			called = true
			return true
		})
		if result != nil {
			t.Error("expected nil")
		}
		if called {
			t.Error("predicate should not be called for nil pointer")
		}
	})

	t.Run("string filter", func(t *testing.T) {
		p := To("hello")
		result := Filter(p, func(s string) bool { return len(s) > 3 })
		if result == nil {
			t.Fatal("expected non-nil")
		}
		if *result != "hello" {
			t.Errorf("expected 'hello', got %q", *result)
		}
	})
}

// Test FlatMap function
func TestFlatMap(t *testing.T) {
	t.Run("successful transformation", func(t *testing.T) {
		s := To("42")
		result := FlatMap(s, func(str string) *int {
			if str == "42" {
				return To(42)
			}
			return nil
		})
		if result == nil {
			t.Fatal("expected non-nil")
		}
		if *result != 42 {
			t.Errorf("expected 42, got %d", *result)
		}
	})

	t.Run("transformation returns nil", func(t *testing.T) {
		s := To("invalid")
		result := FlatMap(s, func(str string) *int {
			return nil
		})
		if result != nil {
			t.Error("expected nil")
		}
	})

	t.Run("nil input", func(t *testing.T) {
		var s *string
		called := false
		result := FlatMap(s, func(str string) *int {
			called = true
			return To(42)
		})
		if result != nil {
			t.Error("expected nil")
		}
		if called {
			t.Error("function should not be called for nil pointer")
		}
	})

	t.Run("chaining transformations", func(t *testing.T) {
		s := To("hello")
		result := FlatMap(s, func(str string) *int {
			return To(len(str))
		})
		if result == nil {
			t.Fatal("expected non-nil")
		}
		if *result != 5 {
			t.Errorf("expected 5, got %d", *result)
		}
	})
}

// Test Apply function
func TestApply(t *testing.T) {
	t.Run("non-nil pointer", func(t *testing.T) {
		p := To(42)
		executed := false
		var value int
		result := Apply(p, func(v int) {
			executed = true
			value = v
		})
		if !result {
			t.Error("expected true")
		}
		if !executed {
			t.Error("function should have been executed")
		}
		if value != 42 {
			t.Errorf("expected 42, got %d", value)
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		var p *int
		executed := false
		result := Apply(p, func(v int) {
			executed = true
		})
		if result {
			t.Error("expected false")
		}
		if executed {
			t.Error("function should not have been executed")
		}
	})

	t.Run("side effects", func(t *testing.T) {
		p := To("hello")
		output := ""
		Apply(p, func(s string) {
			output = s + " world"
		})
		if output != "hello world" {
			t.Errorf("expected 'hello world', got %q", output)
		}
	})
}

// Test Modify function
func TestModify(t *testing.T) {
	t.Run("non-nil pointer", func(t *testing.T) {
		p := To(5)
		result := Modify(p, func(v int) int { return v * 2 })
		if !result {
			t.Error("expected true")
		}
		if *p != 10 {
			t.Errorf("expected 10, got %d", *p)
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		var p *int
		result := Modify(p, func(v int) int { return v * 2 })
		if result {
			t.Error("expected false")
		}
	})

	t.Run("string transformation", func(t *testing.T) {
		p := To("hello")
		Modify(p, func(s string) string { return s + " world" })
		if *p != "hello world" {
			t.Errorf("expected 'hello world', got %q", *p)
		}
	})

	t.Run("complex transformation", func(t *testing.T) {
		type person struct {
			Name string
			Age  int
		}
		p := To(person{Name: "Alice", Age: 30})
		Modify(p, func(per person) person {
			per.Age += 1
			return per
		})
		if p.Age != 31 {
			t.Errorf("expected Age 31, got %d", p.Age)
		}
	})
}

// Test NonZero function
func TestNonZero(t *testing.T) {
	t.Run("non-zero int", func(t *testing.T) {
		p := NonZero(42)
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != 42 {
			t.Errorf("expected 42, got %d", *p)
		}
	})

	t.Run("zero int", func(t *testing.T) {
		p := NonZero(0)
		if p != nil {
			t.Error("expected nil for zero value")
		}
	})

	t.Run("non-zero string", func(t *testing.T) {
		p := NonZero("hello")
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != "hello" {
			t.Errorf("expected 'hello', got %q", *p)
		}
	})

	t.Run("zero string", func(t *testing.T) {
		p := NonZero("")
		if p != nil {
			t.Error("expected nil for empty string")
		}
	})

	t.Run("non-zero bool", func(t *testing.T) {
		p := NonZero(true)
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != true {
			t.Error("expected true")
		}
	})

	t.Run("zero bool", func(t *testing.T) {
		p := NonZero(false)
		if p != nil {
			t.Error("expected nil for false")
		}
	})

	t.Run("non-zero float", func(t *testing.T) {
		p := NonZero(3.14)
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != 3.14 {
			t.Errorf("expected 3.14, got %f", *p)
		}
	})

	t.Run("zero float", func(t *testing.T) {
		p := NonZero(0.0)
		if p != nil {
			t.Error("expected nil for zero float")
		}
	})
}

// Test IsZero function
func TestIsZero(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		var p *int
		if !IsZero(p) {
			t.Error("expected true for nil pointer")
		}
	})

	t.Run("zero value", func(t *testing.T) {
		p := To(0)
		if !IsZero(p) {
			t.Error("expected true for zero value")
		}
	})

	t.Run("non-zero value", func(t *testing.T) {
		p := To(42)
		if IsZero(p) {
			t.Error("expected false for non-zero value")
		}
	})

	t.Run("empty string", func(t *testing.T) {
		p := To("")
		if !IsZero(p) {
			t.Error("expected true for empty string")
		}
	})

	t.Run("non-empty string", func(t *testing.T) {
		p := To("hello")
		if IsZero(p) {
			t.Error("expected false for non-empty string")
		}
	})

	t.Run("false bool", func(t *testing.T) {
		p := To(false)
		if !IsZero(p) {
			t.Error("expected true for false")
		}
	})

	t.Run("true bool", func(t *testing.T) {
		p := To(true)
		if IsZero(p) {
			t.Error("expected false for true")
		}
	})
}

// Test Swap function
func TestSwap(t *testing.T) {
	t.Run("swap two ints", func(t *testing.T) {
		a := To(1)
		b := To(2)
		Swap(a, b)
		if *a != 2 {
			t.Errorf("expected a=2, got %d", *a)
		}
		if *b != 1 {
			t.Errorf("expected b=1, got %d", *b)
		}
	})

	t.Run("swap two strings", func(t *testing.T) {
		a := To("first")
		b := To("second")
		Swap(a, b)
		if *a != "second" {
			t.Errorf("expected a='second', got %q", *a)
		}
		if *b != "first" {
			t.Errorf("expected b='first', got %q", *b)
		}
	})

	t.Run("first nil", func(t *testing.T) {
		var a *int
		b := To(2)
		originalB := *b
		Swap(a, b)
		if *b != originalB {
			t.Error("b should not have changed")
		}
	})

	t.Run("second nil", func(t *testing.T) {
		a := To(1)
		var b *int
		originalA := *a
		Swap(a, b)
		if *a != originalA {
			t.Error("a should not have changed")
		}
	})

	t.Run("both nil", func(t *testing.T) {
		var a, b *int
		Swap(a, b) // should not panic
		// Both remain nil after swap, which is expected behavior
	})

	t.Run("swap structs", func(t *testing.T) {
		type data struct {
			Value int
		}
		a := To(data{Value: 10})
		b := To(data{Value: 20})
		Swap(a, b)
		if a.Value != 20 {
			t.Errorf("expected a.Value=20, got %d", a.Value)
		}
		if b.Value != 10 {
			t.Errorf("expected b.Value=10, got %d", b.Value)
		}
	})
}

// Test Bind function
func TestBind(t *testing.T) {
	t.Run("successful bind", func(t *testing.T) {
		s := To("42")
		result := Bind(s, func(str string) *int {
			if str == "42" {
				return To(42)
			}
			return nil
		})
		if result == nil {
			t.Fatal("expected non-nil")
		}
		if *result != 42 {
			t.Errorf("expected 42, got %d", *result)
		}
	})

	t.Run("bind returns nil", func(t *testing.T) {
		s := To("invalid")
		result := Bind(s, func(str string) *int {
			return nil
		})
		if result != nil {
			t.Error("expected nil")
		}
	})

	t.Run("nil input", func(t *testing.T) {
		var s *string
		result := Bind(s, func(str string) *int {
			return To(42)
		})
		if result != nil {
			t.Error("expected nil")
		}
	})
}

// Test GetOr function
func TestGetOr(t *testing.T) {
	t.Run("non-nil value", func(t *testing.T) {
		s := To("hello")
		result := GetOr(s, "default")
		if result != "hello" {
			t.Errorf("expected 'hello', got %q", result)
		}
	})

	t.Run("nil value", func(t *testing.T) {
		var s *string
		result := GetOr(s, "default")
		if result != "default" {
			t.Errorf("expected 'default', got %q", result)
		}
	})

	t.Run("int type", func(t *testing.T) {
		i := To(42)
		result := GetOr(i, 100)
		if result != 42 {
			t.Errorf("expected 42, got %d", result)
		}
	})
}

// Test Must variant functions
func TestMustString(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		s := To("hello")
		result := MustString(s)
		if result != "hello" {
			t.Errorf("expected 'hello', got %q", result)
		}
	})

	t.Run("nil panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		MustString(nil)
	})
}

func TestMustInt(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		i := To(42)
		result := MustInt(i)
		if result != 42 {
			t.Errorf("expected 42, got %d", result)
		}
	})

	t.Run("nil panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		MustInt(nil)
	})
}

func TestMustInt64(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		i := To(int64(42))
		result := MustInt64(i)
		if result != 42 {
			t.Errorf("expected 42, got %d", result)
		}
	})

	t.Run("nil panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		MustInt64(nil)
	})
}

func TestMustBool(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		b := To(true)
		result := MustBool(b)
		if !result {
			t.Error("expected true")
		}
	})

	t.Run("nil panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		MustBool(nil)
	})
}

func TestMustFloat64(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		f := To(3.14)
		result := MustFloat64(f)
		if result != 3.14 {
			t.Errorf("expected 3.14, got %f", result)
		}
	})

	t.Run("nil panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		MustFloat64(nil)
	})
}
