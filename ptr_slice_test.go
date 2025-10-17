package ptr

import (
	"testing"
	"time"
)

func TestStringSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []string{}, 0},
		{"single element", []string{"hello"}, 1},
		{"multiple elements", []string{"hello", "world", "foo"}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringSlice(tt.input)
			if tt.input == nil {
				if result != nil {
					t.Errorf("StringSlice(%v) = %v, want nil", tt.input, result)
				}
				return
			}
			if len(result) != tt.want {
				t.Errorf("StringSlice(%v) length = %d, want %d", tt.input, len(result), tt.want)
			}
			for i, p := range result {
				if p == nil {
					t.Errorf("StringSlice(%v)[%d] is nil", tt.input, i)
				} else if *p != tt.input[i] {
					t.Errorf("StringSlice(%v)[%d] = %v, want %v", tt.input, i, *p, tt.input[i])
				}
			}
		})
	}
}

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []*string
		want  []string
	}{
		{"nil slice", nil, nil},
		{"empty slice", []*string{}, []string{}},
		{"all non-nil", []*string{String("a"), String("b")}, []string{"a", "b"}},
		{"with nil", []*string{String("a"), nil, String("c")}, []string{"a", "", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToStringSlice(tt.input)
			if tt.want == nil && result != nil {
				t.Errorf("ToStringSlice(%v) = %v, want nil", tt.input, result)
				return
			}
			if len(result) != len(tt.want) {
				t.Errorf("ToStringSlice(%v) length = %d, want %d", tt.input, len(result), len(tt.want))
				return
			}
			for i, v := range result {
				if v != tt.want[i] {
					t.Errorf("ToStringSlice(%v)[%d] = %v, want %v", tt.input, i, v, tt.want[i])
				}
			}
		})
	}
}

func TestIntSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := IntSlice(input)

	if len(result) != len(input) {
		t.Fatalf("IntSlice length = %d, want %d", len(result), len(input))
	}

	for i, p := range result {
		if p == nil {
			t.Errorf("IntSlice[%d] is nil", i)
		} else if *p != input[i] {
			t.Errorf("IntSlice[%d] = %d, want %d", i, *p, input[i])
		}
	}
}

func TestToIntSlice(t *testing.T) {
	input := []*int{Int(1), nil, Int(3)}
	want := []int{1, 0, 3}
	result := ToIntSlice(input)

	if len(result) != len(want) {
		t.Fatalf("ToIntSlice length = %d, want %d", len(result), len(want))
	}

	for i, v := range result {
		if v != want[i] {
			t.Errorf("ToIntSlice[%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestBoolSlice(t *testing.T) {
	input := []bool{true, false, true}
	result := BoolSlice(input)

	if len(result) != len(input) {
		t.Fatalf("BoolSlice length = %d, want %d", len(result), len(input))
	}

	for i, p := range result {
		if p == nil {
			t.Errorf("BoolSlice[%d] is nil", i)
		} else if *p != input[i] {
			t.Errorf("BoolSlice[%d] = %v, want %v", i, *p, input[i])
		}
	}
}

func TestToBoolSlice(t *testing.T) {
	input := []*bool{Bool(true), nil, Bool(false)}
	want := []bool{true, false, false}
	result := ToBoolSlice(input)

	if len(result) != len(want) {
		t.Fatalf("ToBoolSlice length = %d, want %d", len(result), len(want))
	}

	for i, v := range result {
		if v != want[i] {
			t.Errorf("ToBoolSlice[%d] = %v, want %v", i, v, want[i])
		}
	}
}

func TestFloat64Slice(t *testing.T) {
	input := []float64{1.1, 2.2, 3.3}
	result := Float64Slice(input)

	if len(result) != len(input) {
		t.Fatalf("Float64Slice length = %d, want %d", len(result), len(input))
	}

	for i, p := range result {
		if p == nil {
			t.Errorf("Float64Slice[%d] is nil", i)
		} else if *p != input[i] {
			t.Errorf("Float64Slice[%d] = %f, want %f", i, *p, input[i])
		}
	}
}

func TestToFloat64Slice(t *testing.T) {
	input := []*float64{Float64(1.1), nil, Float64(3.3)}
	want := []float64{1.1, 0.0, 3.3}
	result := ToFloat64Slice(input)

	if len(result) != len(want) {
		t.Fatalf("ToFloat64Slice length = %d, want %d", len(result), len(want))
	}

	for i, v := range result {
		if v != want[i] {
			t.Errorf("ToFloat64Slice[%d] = %f, want %f", i, v, want[i])
		}
	}
}

func TestTimeSlice(t *testing.T) {
	now := time.Now()
	input := []time.Time{now, now.Add(time.Hour)}
	result := TimeSlice(input)

	if len(result) != len(input) {
		t.Fatalf("TimeSlice length = %d, want %d", len(result), len(input))
	}

	for i, p := range result {
		if p == nil {
			t.Errorf("TimeSlice[%d] is nil", i)
		} else if !p.Equal(input[i]) {
			t.Errorf("TimeSlice[%d] = %v, want %v", i, *p, input[i])
		}
	}
}

func TestToTimeSlice(t *testing.T) {
	now := time.Now()
	zero := time.Time{}
	input := []*time.Time{Time(now), nil}
	result := ToTimeSlice(input)

	if len(result) != len(input) {
		t.Fatalf("ToTimeSlice length = %d, want %d", len(result), len(input))
	}

	if !result[0].Equal(now) {
		t.Errorf("ToTimeSlice[0] = %v, want %v", result[0], now)
	}
	if !result[1].Equal(zero) {
		t.Errorf("ToTimeSlice[1] = %v, want %v", result[1], zero)
	}
}

func TestDurationSlice(t *testing.T) {
	input := []time.Duration{time.Second, time.Minute}
	result := DurationSlice(input)

	if len(result) != len(input) {
		t.Fatalf("DurationSlice length = %d, want %d", len(result), len(input))
	}

	for i, p := range result {
		if p == nil {
			t.Errorf("DurationSlice[%d] is nil", i)
		} else if *p != input[i] {
			t.Errorf("DurationSlice[%d] = %v, want %v", i, *p, input[i])
		}
	}
}

func TestToDurationSlice(t *testing.T) {
	input := []*time.Duration{Duration(time.Second), nil, Duration(time.Minute)}
	want := []time.Duration{time.Second, 0, time.Minute}
	result := ToDurationSlice(input)

	if len(result) != len(want) {
		t.Fatalf("ToDurationSlice length = %d, want %d", len(result), len(want))
	}

	for i, v := range result {
		if v != want[i] {
			t.Errorf("ToDurationSlice[%d] = %v, want %v", i, v, want[i])
		}
	}
}

func TestAllIntegerSlices(t *testing.T) {
	// Test Int8
	int8Input := []int8{1, 2, 3}
	int8Result := Int8Slice(int8Input)
	if len(int8Result) != 3 || *int8Result[0] != 1 {
		t.Error("Int8Slice failed")
	}
	int8Back := ToInt8Slice(int8Result)
	if len(int8Back) != 3 || int8Back[0] != 1 {
		t.Error("ToInt8Slice failed")
	}

	// Test Int16
	int16Input := []int16{10, 20, 30}
	int16Result := Int16Slice(int16Input)
	if len(int16Result) != 3 || *int16Result[1] != 20 {
		t.Error("Int16Slice failed")
	}
	int16Back := ToInt16Slice(int16Result)
	if len(int16Back) != 3 || int16Back[1] != 20 {
		t.Error("ToInt16Slice failed")
	}

	// Test Int32
	int32Input := []int32{100, 200, 300}
	int32Result := Int32Slice(int32Input)
	if len(int32Result) != 3 || *int32Result[2] != 300 {
		t.Error("Int32Slice failed")
	}
	int32Back := ToInt32Slice(int32Result)
	if len(int32Back) != 3 || int32Back[2] != 300 {
		t.Error("ToInt32Slice failed")
	}

	// Test Int64
	int64Input := []int64{1000, 2000, 3000}
	int64Result := Int64Slice(int64Input)
	if len(int64Result) != 3 || *int64Result[0] != 1000 {
		t.Error("Int64Slice failed")
	}
	int64Back := ToInt64Slice(int64Result)
	if len(int64Back) != 3 || int64Back[0] != 1000 {
		t.Error("ToInt64Slice failed")
	}
}

func TestAllUnsignedIntegerSlices(t *testing.T) {
	// Test Uint
	uintInput := []uint{1, 2, 3}
	uintResult := UintSlice(uintInput)
	if len(uintResult) != 3 || *uintResult[0] != 1 {
		t.Error("UintSlice failed")
	}
	uintBack := ToUintSlice(uintResult)
	if len(uintBack) != 3 || uintBack[0] != 1 {
		t.Error("ToUintSlice failed")
	}

	// Test Uint8
	uint8Input := []uint8{10, 20, 30}
	uint8Result := Uint8Slice(uint8Input)
	if len(uint8Result) != 3 || *uint8Result[1] != 20 {
		t.Error("Uint8Slice failed")
	}
	uint8Back := ToUint8Slice(uint8Result)
	if len(uint8Back) != 3 || uint8Back[1] != 20 {
		t.Error("ToUint8Slice failed")
	}

	// Test Uint16
	uint16Input := []uint16{100, 200, 300}
	uint16Result := Uint16Slice(uint16Input)
	if len(uint16Result) != 3 || *uint16Result[2] != 300 {
		t.Error("Uint16Slice failed")
	}
	uint16Back := ToUint16Slice(uint16Result)
	if len(uint16Back) != 3 || uint16Back[2] != 300 {
		t.Error("ToUint16Slice failed")
	}

	// Test Uint32
	uint32Input := []uint32{1000, 2000, 3000}
	uint32Result := Uint32Slice(uint32Input)
	if len(uint32Result) != 3 || *uint32Result[0] != 1000 {
		t.Error("Uint32Slice failed")
	}
	uint32Back := ToUint32Slice(uint32Result)
	if len(uint32Back) != 3 || uint32Back[0] != 1000 {
		t.Error("ToUint32Slice failed")
	}

	// Test Uint64
	uint64Input := []uint64{10000, 20000, 30000}
	uint64Result := Uint64Slice(uint64Input)
	if len(uint64Result) != 3 || *uint64Result[1] != 20000 {
		t.Error("Uint64Slice failed")
	}
	uint64Back := ToUint64Slice(uint64Result)
	if len(uint64Back) != 3 || uint64Back[1] != 20000 {
		t.Error("ToUint64Slice failed")
	}
}

func TestFloat32Slice(t *testing.T) {
	input := []float32{1.1, 2.2, 3.3}
	result := Float32Slice(input)

	if len(result) != len(input) {
		t.Fatalf("Float32Slice length = %d, want %d", len(result), len(input))
	}

	for i, p := range result {
		if p == nil {
			t.Errorf("Float32Slice[%d] is nil", i)
		} else if *p != input[i] {
			t.Errorf("Float32Slice[%d] = %f, want %f", i, *p, input[i])
		}
	}
}

func TestToFloat32Slice(t *testing.T) {
	input := []*float32{Float32(1.1), nil, Float32(3.3)}
	want := []float32{1.1, 0.0, 3.3}
	result := ToFloat32Slice(input)

	if len(result) != len(want) {
		t.Fatalf("ToFloat32Slice length = %d, want %d", len(result), len(want))
	}

	for i, v := range result {
		if v != want[i] {
			t.Errorf("ToFloat32Slice[%d] = %f, want %f", i, v, want[i])
		}
	}
}

func TestByteSlice(t *testing.T) {
	input := []byte{1, 2, 3}
	result := ByteSlice(input)

	if len(result) != len(input) {
		t.Fatalf("ByteSlice length = %d, want %d", len(result), len(input))
	}

	for i, p := range result {
		if p == nil {
			t.Errorf("ByteSlice[%d] is nil", i)
		} else if *p != input[i] {
			t.Errorf("ByteSlice[%d] = %d, want %d", i, *p, input[i])
		}
	}
}

func TestToByteSlice(t *testing.T) {
	input := []*byte{Byte(1), nil, Byte(3)}
	want := []byte{1, 0, 3}
	result := ToByteSlice(input)

	if len(result) != len(want) {
		t.Fatalf("ToByteSlice length = %d, want %d", len(result), len(want))
	}

	for i, v := range result {
		if v != want[i] {
			t.Errorf("ToByteSlice[%d] = %d, want %d", i, v, want[i])
		}
	}
}
