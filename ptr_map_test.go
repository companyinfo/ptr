package ptr

import (
	"testing"
	"time"
)

func TestToMap(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
	}{
		{"nil map", nil},
		{"empty map", map[string]int{}},
		{"single entry", map[string]int{"a": 1}},
		{"multiple entries", map[string]int{"a": 1, "b": 2, "c": 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMap(tt.input)
			if tt.input == nil {
				if result != nil {
					t.Errorf("ToMap(%v) = %v, want nil", tt.input, result)
				}
				return
			}
			if len(result) != len(tt.input) {
				t.Errorf("ToMap(%v) length = %d, want %d", tt.input, len(result), len(tt.input))
			}
			for k, v := range tt.input {
				if p, ok := result[k]; !ok {
					t.Errorf("ToMap(%v) missing key %q", tt.input, k)
				} else if p == nil {
					t.Errorf("ToMap(%v)[%q] is nil", tt.input, k)
				} else if *p != v {
					t.Errorf("ToMap(%v)[%q] = %v, want %v", tt.input, k, *p, v)
				}
			}
		})
	}
}

func TestFromMap(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]*int
		want  map[string]int
	}{
		{"nil map", nil, nil},
		{"empty map", map[string]*int{}, map[string]int{}},
		{"all non-nil", map[string]*int{"a": Int(1), "b": Int(2)}, map[string]int{"a": 1, "b": 2}},
		{"with nil", map[string]*int{"a": Int(1), "b": nil, "c": Int(3)}, map[string]int{"a": 1, "b": 0, "c": 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromMap(tt.input)
			if tt.want == nil && result != nil {
				t.Errorf("FromMap(%v) = %v, want nil", tt.input, result)
				return
			}
			if len(result) != len(tt.want) {
				t.Errorf("FromMap(%v) length = %d, want %d", tt.input, len(result), len(tt.want))
				return
			}
			for k, v := range tt.want {
				if got, ok := result[k]; !ok {
					t.Errorf("FromMap(%v) missing key %q", tt.input, k)
				} else if got != v {
					t.Errorf("FromMap(%v)[%q] = %v, want %v", tt.input, k, got, v)
				}
			}
		})
	}
}

func TestStringMap(t *testing.T) {
	input := map[string]string{"a": "hello", "b": "world"}
	result := StringMap(input)

	if len(result) != len(input) {
		t.Fatalf("StringMap length = %d, want %d", len(result), len(input))
	}

	for k, v := range input {
		if p, ok := result[k]; !ok {
			t.Errorf("StringMap missing key %q", k)
		} else if p == nil {
			t.Errorf("StringMap[%q] is nil", k)
		} else if *p != v {
			t.Errorf("StringMap[%q] = %q, want %q", k, *p, v)
		}
	}
}

func TestToStringMap(t *testing.T) {
	input := map[string]*string{"a": String("hello"), "b": nil, "c": String("world")}
	want := map[string]string{"a": "hello", "b": "", "c": "world"}
	result := ToStringMap(input)

	if len(result) != len(want) {
		t.Fatalf("ToStringMap length = %d, want %d", len(result), len(want))
	}

	for k, v := range want {
		if got, ok := result[k]; !ok {
			t.Errorf("ToStringMap missing key %q", k)
		} else if got != v {
			t.Errorf("ToStringMap[%q] = %q, want %q", k, got, v)
		}
	}
}

func TestIntMap(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	result := IntMap(input)

	if len(result) != len(input) {
		t.Fatalf("IntMap length = %d, want %d", len(result), len(input))
	}

	for k, v := range input {
		if p, ok := result[k]; !ok {
			t.Errorf("IntMap missing key %q", k)
		} else if p == nil {
			t.Errorf("IntMap[%q] is nil", k)
		} else if *p != v {
			t.Errorf("IntMap[%q] = %d, want %d", k, *p, v)
		}
	}
}

func TestToIntMap(t *testing.T) {
	input := map[string]*int{"a": Int(1), "b": nil, "c": Int(3)}
	want := map[string]int{"a": 1, "b": 0, "c": 3}
	result := ToIntMap(input)

	if len(result) != len(want) {
		t.Fatalf("ToIntMap length = %d, want %d", len(result), len(want))
	}

	for k, v := range want {
		if got, ok := result[k]; !ok {
			t.Errorf("ToIntMap missing key %q", k)
		} else if got != v {
			t.Errorf("ToIntMap[%q] = %d, want %d", k, got, v)
		}
	}
}

func TestBoolMap(t *testing.T) {
	input := map[string]bool{"a": true, "b": false, "c": true}
	result := BoolMap(input)

	if len(result) != len(input) {
		t.Fatalf("BoolMap length = %d, want %d", len(result), len(input))
	}

	for k, v := range input {
		if p, ok := result[k]; !ok {
			t.Errorf("BoolMap missing key %q", k)
		} else if p == nil {
			t.Errorf("BoolMap[%q] is nil", k)
		} else if *p != v {
			t.Errorf("BoolMap[%q] = %v, want %v", k, *p, v)
		}
	}
}

func TestToBoolMap(t *testing.T) {
	input := map[string]*bool{"a": Bool(true), "b": nil, "c": Bool(false)}
	want := map[string]bool{"a": true, "b": false, "c": false}
	result := ToBoolMap(input)

	if len(result) != len(want) {
		t.Fatalf("ToBoolMap length = %d, want %d", len(result), len(want))
	}

	for k, v := range want {
		if got, ok := result[k]; !ok {
			t.Errorf("ToBoolMap missing key %q", k)
		} else if got != v {
			t.Errorf("ToBoolMap[%q] = %v, want %v", k, got, v)
		}
	}
}

func TestFloat64Map(t *testing.T) {
	input := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	result := Float64Map(input)

	if len(result) != len(input) {
		t.Fatalf("Float64Map length = %d, want %d", len(result), len(input))
	}

	for k, v := range input {
		if p, ok := result[k]; !ok {
			t.Errorf("Float64Map missing key %q", k)
		} else if p == nil {
			t.Errorf("Float64Map[%q] is nil", k)
		} else if *p != v {
			t.Errorf("Float64Map[%q] = %f, want %f", k, *p, v)
		}
	}
}

func TestToFloat64Map(t *testing.T) {
	input := map[string]*float64{"a": Float64(1.1), "b": nil, "c": Float64(3.3)}
	want := map[string]float64{"a": 1.1, "b": 0.0, "c": 3.3}
	result := ToFloat64Map(input)

	if len(result) != len(want) {
		t.Fatalf("ToFloat64Map length = %d, want %d", len(result), len(want))
	}

	for k, v := range want {
		if got, ok := result[k]; !ok {
			t.Errorf("ToFloat64Map missing key %q", k)
		} else if got != v {
			t.Errorf("ToFloat64Map[%q] = %f, want %f", k, got, v)
		}
	}
}

func TestTimeMap(t *testing.T) {
	now := time.Now()
	input := map[string]time.Time{"a": now, "b": now.Add(time.Hour)}
	result := TimeMap(input)

	if len(result) != len(input) {
		t.Fatalf("TimeMap length = %d, want %d", len(result), len(input))
	}

	for k, v := range input {
		if p, ok := result[k]; !ok {
			t.Errorf("TimeMap missing key %q", k)
		} else if p == nil {
			t.Errorf("TimeMap[%q] is nil", k)
		} else if !p.Equal(v) {
			t.Errorf("TimeMap[%q] = %v, want %v", k, *p, v)
		}
	}
}

func TestToTimeMap(t *testing.T) {
	now := time.Now()
	zero := time.Time{}
	input := map[string]*time.Time{"a": Time(now), "b": nil}
	result := ToTimeMap(input)

	if len(result) != len(input) {
		t.Fatalf("ToTimeMap length = %d, want %d", len(result), len(input))
	}

	if !result["a"].Equal(now) {
		t.Errorf("ToTimeMap[a] = %v, want %v", result["a"], now)
	}
	if !result["b"].Equal(zero) {
		t.Errorf("ToTimeMap[b] = %v, want %v", result["b"], zero)
	}
}

func TestDurationMap(t *testing.T) {
	input := map[string]time.Duration{"a": time.Second, "b": time.Minute}
	result := DurationMap(input)

	if len(result) != len(input) {
		t.Fatalf("DurationMap length = %d, want %d", len(result), len(input))
	}

	for k, v := range input {
		if p, ok := result[k]; !ok {
			t.Errorf("DurationMap missing key %q", k)
		} else if p == nil {
			t.Errorf("DurationMap[%q] is nil", k)
		} else if *p != v {
			t.Errorf("DurationMap[%q] = %v, want %v", k, *p, v)
		}
	}
}

func TestToDurationMap(t *testing.T) {
	input := map[string]*time.Duration{"a": Duration(time.Second), "b": nil, "c": Duration(time.Minute)}
	want := map[string]time.Duration{"a": time.Second, "b": 0, "c": time.Minute}
	result := ToDurationMap(input)

	if len(result) != len(want) {
		t.Fatalf("ToDurationMap length = %d, want %d", len(result), len(want))
	}

	for k, v := range want {
		if got, ok := result[k]; !ok {
			t.Errorf("ToDurationMap missing key %q", k)
		} else if got != v {
			t.Errorf("ToDurationMap[%q] = %v, want %v", k, got, v)
		}
	}
}

func TestAllIntegerMaps(t *testing.T) {
	// Test Int8
	int8Input := map[string]int8{"a": 1, "b": 2}
	int8Result := Int8Map(int8Input)
	if len(int8Result) != 2 || *int8Result["a"] != 1 {
		t.Error("Int8Map failed")
	}
	int8Back := ToInt8Map(int8Result)
	if len(int8Back) != 2 || int8Back["a"] != 1 {
		t.Error("ToInt8Map failed")
	}

	// Test Int16
	int16Input := map[string]int16{"a": 10, "b": 20}
	int16Result := Int16Map(int16Input)
	if len(int16Result) != 2 || *int16Result["b"] != 20 {
		t.Error("Int16Map failed")
	}
	int16Back := ToInt16Map(int16Result)
	if len(int16Back) != 2 || int16Back["b"] != 20 {
		t.Error("ToInt16Map failed")
	}

	// Test Int32
	int32Input := map[string]int32{"a": 100, "b": 200}
	int32Result := Int32Map(int32Input)
	if len(int32Result) != 2 || *int32Result["a"] != 100 {
		t.Error("Int32Map failed")
	}
	int32Back := ToInt32Map(int32Result)
	if len(int32Back) != 2 || int32Back["a"] != 100 {
		t.Error("ToInt32Map failed")
	}

	// Test Int64
	int64Input := map[string]int64{"a": 1000, "b": 2000}
	int64Result := Int64Map(int64Input)
	if len(int64Result) != 2 || *int64Result["b"] != 2000 {
		t.Error("Int64Map failed")
	}
	int64Back := ToInt64Map(int64Result)
	if len(int64Back) != 2 || int64Back["b"] != 2000 {
		t.Error("ToInt64Map failed")
	}
}

func TestAllUnsignedIntegerMaps(t *testing.T) {
	// Test Uint
	uintInput := map[string]uint{"a": 1, "b": 2}
	uintResult := UintMap(uintInput)
	if len(uintResult) != 2 || *uintResult["a"] != 1 {
		t.Error("UintMap failed")
	}
	uintBack := ToUintMap(uintResult)
	if len(uintBack) != 2 || uintBack["a"] != 1 {
		t.Error("ToUintMap failed")
	}

	// Test Uint8
	uint8Input := map[string]uint8{"a": 10, "b": 20}
	uint8Result := Uint8Map(uint8Input)
	if len(uint8Result) != 2 || *uint8Result["b"] != 20 {
		t.Error("Uint8Map failed")
	}
	uint8Back := ToUint8Map(uint8Result)
	if len(uint8Back) != 2 || uint8Back["b"] != 20 {
		t.Error("ToUint8Map failed")
	}

	// Test Uint16
	uint16Input := map[string]uint16{"a": 100, "b": 200}
	uint16Result := Uint16Map(uint16Input)
	if len(uint16Result) != 2 || *uint16Result["a"] != 100 {
		t.Error("Uint16Map failed")
	}
	uint16Back := ToUint16Map(uint16Result)
	if len(uint16Back) != 2 || uint16Back["a"] != 100 {
		t.Error("ToUint16Map failed")
	}

	// Test Uint32
	uint32Input := map[string]uint32{"a": 1000, "b": 2000}
	uint32Result := Uint32Map(uint32Input)
	if len(uint32Result) != 2 || *uint32Result["b"] != 2000 {
		t.Error("Uint32Map failed")
	}
	uint32Back := ToUint32Map(uint32Result)
	if len(uint32Back) != 2 || uint32Back["b"] != 2000 {
		t.Error("ToUint32Map failed")
	}

	// Test Uint64
	uint64Input := map[string]uint64{"a": 10000, "b": 20000}
	uint64Result := Uint64Map(uint64Input)
	if len(uint64Result) != 2 || *uint64Result["a"] != 10000 {
		t.Error("Uint64Map failed")
	}
	uint64Back := ToUint64Map(uint64Result)
	if len(uint64Back) != 2 || uint64Back["a"] != 10000 {
		t.Error("ToUint64Map failed")
	}
}

func TestFloat32Map(t *testing.T) {
	input := map[string]float32{"a": 1.1, "b": 2.2}
	result := Float32Map(input)

	if len(result) != len(input) {
		t.Fatalf("Float32Map length = %d, want %d", len(result), len(input))
	}

	for k, v := range input {
		if p, ok := result[k]; !ok {
			t.Errorf("Float32Map missing key %q", k)
		} else if p == nil {
			t.Errorf("Float32Map[%q] is nil", k)
		} else if *p != v {
			t.Errorf("Float32Map[%q] = %f, want %f", k, *p, v)
		}
	}
}

func TestToFloat32Map(t *testing.T) {
	input := map[string]*float32{"a": Float32(1.1), "b": nil}
	want := map[string]float32{"a": 1.1, "b": 0.0}
	result := ToFloat32Map(input)

	if len(result) != len(want) {
		t.Fatalf("ToFloat32Map length = %d, want %d", len(result), len(want))
	}

	for k, v := range want {
		if got, ok := result[k]; !ok {
			t.Errorf("ToFloat32Map missing key %q", k)
		} else if got != v {
			t.Errorf("ToFloat32Map[%q] = %f, want %f", k, got, v)
		}
	}
}

func TestByteMap(t *testing.T) {
	input := map[string]byte{"a": 1, "b": 2}
	result := ByteMap(input)

	if len(result) != len(input) {
		t.Fatalf("ByteMap length = %d, want %d", len(result), len(input))
	}

	for k, v := range input {
		if p, ok := result[k]; !ok {
			t.Errorf("ByteMap missing key %q", k)
		} else if p == nil {
			t.Errorf("ByteMap[%q] is nil", k)
		} else if *p != v {
			t.Errorf("ByteMap[%q] = %d, want %d", k, *p, v)
		}
	}
}

func TestToByteMap(t *testing.T) {
	input := map[string]*byte{"a": Byte(1), "b": nil}
	want := map[string]byte{"a": 1, "b": 0}
	result := ToByteMap(input)

	if len(result) != len(want) {
		t.Fatalf("ToByteMap length = %d, want %d", len(result), len(want))
	}

	for k, v := range want {
		if got, ok := result[k]; !ok {
			t.Errorf("ToByteMap missing key %q", k)
		} else if got != v {
			t.Errorf("ToByteMap[%q] = %d, want %d", k, got, v)
		}
	}
}
