package ptr

import (
	"strings"
	"testing"
	"time"
)

// Benchmark generic To function
func BenchmarkTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = To(42)
	}
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = To("hello world")
	}
}

func BenchmarkToStruct(b *testing.B) {
	type Data struct {
		Name string
		Age  int
	}
	v := Data{Name: "Alice", Age: 30}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = To(v)
	}
}

// Benchmark generic From function
func BenchmarkFrom(b *testing.B) {
	p := To(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = From(p)
	}
}

func BenchmarkFromNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = From[int](nil)
	}
}

// Benchmark FromOr function
func BenchmarkFromOr(b *testing.B) {
	p := To(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FromOr(p, 100)
	}
}

func BenchmarkFromOrNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FromOr[int](nil, 100)
	}
}

// Benchmark Equal function
func BenchmarkEqual(b *testing.B) {
	a := To(42)
	c := To(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Equal(a, c)
	}
}

func BenchmarkEqualNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Equal[int](nil, nil)
	}
}

// Benchmark Copy function
func BenchmarkCopy(b *testing.B) {
	p := To(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Copy(p)
	}
}

func BenchmarkCopyString(b *testing.B) {
	p := To("hello world, this is a longer string to benchmark")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Copy(p)
	}
}

// Benchmark IsNil function
func BenchmarkIsNil(b *testing.B) {
	p := To(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsNil(p)
	}
}

func BenchmarkIsNilWithNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsNil[int](nil)
	}
}

// Benchmark type-specific functions
func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = String("hello")
	}
}

func BenchmarkToString_TypeSpecific(b *testing.B) {
	p := String("hello")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToString(p)
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Int(42)
	}
}

func BenchmarkToInt_TypeSpecific(b *testing.B) {
	p := Int(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInt(p)
	}
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Bool(true)
	}
}

func BenchmarkToBool_TypeSpecific(b *testing.B) {
	p := Bool(true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToBool(p)
	}
}

func BenchmarkFloat64_TypeSpecific(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Float64(3.14159)
	}
}

func BenchmarkToFloat64_TypeSpecific(b *testing.B) {
	p := Float64(3.14159)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToFloat64(p)
	}
}

// Benchmark Coalesce function
func BenchmarkCoalesce(b *testing.B) {
	p1 := To(42)
	p2 := To(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Coalesce[int](nil, p1, p2)
	}
}

func BenchmarkCoalesceAllNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Coalesce[int](nil, nil, nil)
	}
}

// Benchmark Set function
func BenchmarkSet(b *testing.B) {
	p := To(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Set(p, 100)
	}
}

func BenchmarkSetNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Set[int](nil, 100)
	}
}

// Benchmark Map function
func BenchmarkMap(b *testing.B) {
	p := To("hello world")
	fn := func(s string) int { return len(s) }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Map(p, fn)
	}
}

func BenchmarkMapComplex(b *testing.B) {
	p := To("hello world, this is a longer string for benchmarking")
	fn := func(s string) string { return strings.ToUpper(s) }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Map(p, fn)
	}
}

func BenchmarkMapNil(b *testing.B) {
	fn := func(s string) int { return len(s) }
	for i := 0; i < b.N; i++ {
		_ = Map[string, int](nil, fn)
	}
}

// Benchmark ToSlice function
func BenchmarkToSlice(b *testing.B) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToSlice(values)
	}
}

func BenchmarkToSliceLarge(b *testing.B) {
	values := make([]int, 1000)
	for i := range values {
		values[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToSlice(values)
	}
}

// Benchmark FromSlice function
func BenchmarkFromSlice(b *testing.B) {
	ptrs := []*int{To(1), To(2), To(3), To(4), To(5), To(6), To(7), To(8), To(9), To(10)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FromSlice(ptrs)
	}
}

func BenchmarkFromSliceWithNils(b *testing.B) {
	ptrs := []*int{To(1), nil, To(3), nil, To(5), To(6), nil, To(8), To(9), nil}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FromSlice(ptrs)
	}
}

func BenchmarkFromSliceLarge(b *testing.B) {
	ptrs := make([]*int, 1000)
	for i := range ptrs {
		v := i
		ptrs[i] = &v
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FromSlice(ptrs)
	}
}

// Benchmark Time functions
func BenchmarkTime(b *testing.B) {
	v := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Time(v)
	}
}

func BenchmarkToTime(b *testing.B) {
	v := time.Now()
	p := &v
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToTime(p)
	}
}

// Benchmark Duration functions
func BenchmarkDuration(b *testing.B) {
	v := time.Second * 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Duration(v)
	}
}

func BenchmarkToDuration(b *testing.B) {
	v := time.Second * 30
	p := &v
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToDuration(p)
	}
}

// Benchmark Complex functions
func BenchmarkComplex64(b *testing.B) {
	v := complex64(3 + 4i)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Complex64(v)
	}
}

func BenchmarkToComplex64(b *testing.B) {
	v := complex64(3 + 4i)
	p := &v
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToComplex64(p)
	}
}

func BenchmarkComplex128(b *testing.B) {
	v := complex(3, 4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Complex128(v)
	}
}

func BenchmarkToComplex128(b *testing.B) {
	v := complex(3, 4)
	p := &v
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToComplex128(p)
	}
}

// Benchmark Uintptr functions
func BenchmarkUintptr(b *testing.B) {
	v := uintptr(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Uintptr(v)
	}
}

func BenchmarkToUintptr(b *testing.B) {
	v := uintptr(42)
	p := &v
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToUintptr(p)
	}
}

// Benchmark new functional programming functions

func BenchmarkOr(b *testing.B) {
	a := To(42)
	fallback := To(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Or(a, fallback)
	}
}

func BenchmarkOrNil(b *testing.B) {
	fallback := To(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Or[int](nil, fallback)
	}
}

func BenchmarkFilter(b *testing.B) {
	p := To(42)
	predicate := func(v int) bool { return v > 40 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Filter(p, predicate)
	}
}

func BenchmarkFilterNil(b *testing.B) {
	predicate := func(v int) bool { return v > 40 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Filter[int](nil, predicate)
	}
}

func BenchmarkFlatMap(b *testing.B) {
	p := To(42)
	transform := func(v int) *string {
		s := strings.Repeat("x", v)
		return &s
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FlatMap(p, transform)
	}
}

func BenchmarkFlatMapNil(b *testing.B) {
	transform := func(v int) *string {
		s := "result"
		return &s
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FlatMap[int, string](nil, transform)
	}
}

func BenchmarkApply(b *testing.B) {
	p := To(42)
	fn := func(v int) {}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Apply(p, fn)
	}
}

func BenchmarkApplyNil(b *testing.B) {
	fn := func(v int) {}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Apply[int](nil, fn)
	}
}

func BenchmarkModify(b *testing.B) {
	p := To(100)
	transform := func(v int) int { return v * 2 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Modify(p, transform)
	}
}

func BenchmarkModifyNil(b *testing.B) {
	transform := func(v int) int { return v * 2 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Modify[int](nil, transform)
	}
}

func BenchmarkNonZero(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NonZero(42)
	}
}

func BenchmarkNonZeroZero(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NonZero(0)
	}
}

func BenchmarkIsZero(b *testing.B) {
	p := To(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsZero(p)
	}
}

func BenchmarkIsZeroNil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsZero[int](nil)
	}
}

func BenchmarkIsZeroZeroValue(b *testing.B) {
	p := To(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsZero(p)
	}
}

func BenchmarkSwap(b *testing.B) {
	a := To(1)
	bVal := To(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Swap(a, bVal)
	}
}

func BenchmarkSwapNil(b *testing.B) {
	bVal := To(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Swap[int](nil, bVal)
	}
}
