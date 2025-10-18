package ptr_test

import (
	"encoding/json"
	"fmt"
	"time"

	"go.companyinfo.dev/ptr"
)

// Example demonstrates basic usage of the ptr package
func Example() {
	// Create pointers to values
	name := ptr.String("Alice")
	age := ptr.Int(30)
	active := ptr.Bool(true)

	fmt.Printf("Name: %s\n", ptr.ToString(name))
	fmt.Printf("Age: %d\n", ptr.ToInt(age))
	fmt.Printf("Active: %v\n", ptr.ToBool(active))

	// Safe dereferencing with nil
	var nilName *string
	fmt.Printf("Nil name: %s\n", ptr.ToString(nilName))

	// Output:
	// Name: Alice
	// Age: 30
	// Active: true
	// Nil name:
}

// Example_jsonMarshaling demonstrates using ptr for JSON with optional fields
func Example_jsonMarshaling() {
	type User struct {
		Name  string  `json:"name"`
		Email *string `json:"email,omitempty"`
		Age   *int    `json:"age,omitempty"`
	}

	// User with all fields
	user1 := User{
		Name:  "Alice",
		Email: ptr.String("alice@example.com"),
		Age:   ptr.Int(30),
	}

	// User with optional fields as nil (will be omitted from JSON)
	user2 := User{
		Name: "Bob",
	}

	json1, _ := json.Marshal(user1)
	json2, _ := json.Marshal(user2)

	fmt.Printf("User 1: %s\n", json1)
	fmt.Printf("User 2: %s\n", json2)

	// Output:
	// User 1: {"name":"Alice","email":"alice@example.com","age":30}
	// User 2: {"name":"Bob"}
}

// Example_generics demonstrates using generic functions with custom types
func Example_generics() {
	type Color struct {
		R, G, B uint8
	}

	// Generic To works with any type
	red := ptr.To(Color{255, 0, 0})
	fmt.Printf("Red: %+v\n", ptr.From(red))

	// Copy creates independent copies
	blue := ptr.To(Color{0, 0, 255})
	blueCopy := ptr.Copy(blue)
	blue.R = 128 // Modify original
	fmt.Printf("Original modified: %+v\n", *blue)
	fmt.Printf("Copy unchanged: %+v\n", *blueCopy)

	// Equal compares values
	color1 := ptr.To(Color{255, 0, 0})
	color2 := ptr.To(Color{255, 0, 0})
	fmt.Printf("Colors equal: %v\n", ptr.Equal(color1, color2))

	// Output:
	// Red: {R:255 G:0 B:0}
	// Original modified: {R:128 G:0 B:255}
	// Copy unchanged: {R:0 G:0 B:255}
	// Colors equal: true
}

// Example_fromOr demonstrates using FromOr for custom defaults
func Example_fromOr() {
	// Configuration with optional values
	type Config struct {
		Host    string
		Port    *int
		Timeout *int
	}

	config := Config{
		Host: "localhost",
		Port: ptr.Int(8080),
		// Timeout is nil
	}

	// Use FromOr to provide defaults
	host := config.Host
	port := ptr.FromOr(config.Port, 3000)
	timeout := ptr.FromOr(config.Timeout, 30) // Default 30 seconds

	fmt.Printf("Host: %s\n", host)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Timeout: %d seconds\n", timeout)

	// Output:
	// Host: localhost
	// Port: 8080
	// Timeout: 30 seconds
}

// Example_coalesce demonstrates using Coalesce to find the first non-nil value
func Example_coalesce() {
	// Simulate fetching configuration from multiple sources
	var envPort *int    // Not set in environment
	var configPort *int // Not set in config file
	defaultPort := ptr.Int(8080)

	// Use the first available value
	port := ptr.Coalesce(envPort, configPort, defaultPort)

	fmt.Printf("Port: %d\n", *port)

	// Output:
	// Port: 8080
}

// Example_set demonstrates using Set to modify pointer values
func Example_set() {
	type User struct {
		Name  string
		Email *string
	}

	user := User{
		Name:  "Alice",
		Email: ptr.String("old@example.com"),
	}

	// Update email if it exists
	if ptr.Set(user.Email, "new@example.com") {
		fmt.Println("Email updated successfully")
	}

	fmt.Printf("New email: %s\n", ptr.ToString(user.Email))

	// Output:
	// Email updated successfully
	// New email: new@example.com
}

// Example_map demonstrates using Map to transform pointer values
func Example_map() {
	// Transform a string to its length
	name := ptr.String("Alice")
	nameLength := ptr.Map(name, func(s string) int {
		return len(s)
	})

	fmt.Printf("Name length: %d\n", *nameLength)

	// Map returns nil if input is nil
	var nilName *string
	nilLength := ptr.Map(nilName, func(s string) int {
		return len(s)
	})

	fmt.Printf("Nil length is nil: %v\n", nilLength == nil)

	// Output:
	// Name length: 5
	// Nil length is nil: true
}

// Example_sliceOperations demonstrates using ToSlice and FromSlice
func Example_sliceOperations() {
	// Convert slice of values to slice of pointers
	ages := []int{25, 30, 35, 40}
	agePointers := ptr.ToSlice(ages)

	fmt.Printf("First age pointer: %d\n", *agePointers[0])

	// Convert slice of pointers back to values
	// Nil pointers become zero values
	pointers := []*int{ptr.Int(1), nil, ptr.Int(3)}
	values := ptr.FromSlice(pointers)

	fmt.Printf("Values: %v\n", values)

	// Output:
	// First age pointer: 25
	// Values: [1 0 3]
}

// Example_timeAndDuration demonstrates using Time and Duration helpers
func Example_timeAndDuration() {
	type Event struct {
		Name      string         `json:"name"`
		Timestamp *time.Time     `json:"timestamp,omitempty"`
		Timeout   *time.Duration `json:"timeout,omitempty"`
	}

	// Create event with time pointers
	event := Event{
		Name:      "API Call",
		Timestamp: ptr.Time(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
		Timeout:   ptr.Duration(30 * time.Second),
	}

	eventJSON, _ := json.Marshal(event)
	fmt.Printf("Event: %s\n", eventJSON)

	// Output:
	// Event: {"name":"API Call","timestamp":"2024-01-01T00:00:00Z","timeout":30000000000}
}

// Example_sliceHelpers demonstrates using type-specific slice conversion functions
func Example_sliceHelpers() {
	// Convert slice of values to slice of pointers
	numbers := []int{1, 2, 3, 4, 5}
	numberPtrs := ptr.IntSlice(numbers)
	fmt.Printf("Number of pointers: %d\n", len(numberPtrs))
	fmt.Printf("First value: %d\n", *numberPtrs[0])

	// Convert slice of pointers back to values
	// Nil pointers become zero values
	mixedPtrs := []*string{ptr.String("hello"), nil, ptr.String("world")}
	values := ptr.ToStringSlice(mixedPtrs)
	fmt.Printf("Values: %v\n", values)

	// Works with booleans
	flags := []bool{true, false, true}
	flagPtrs := ptr.BoolSlice(flags)
	fmt.Printf("All flags have pointers: %t\n", len(flagPtrs) == len(flags))

	// Output:
	// Number of pointers: 5
	// First value: 1
	// Values: [hello  world]
	// All flags have pointers: true
}

// Example_mapHelpers demonstrates using type-specific map conversion functions
func Example_mapHelpers() {
	// Convert map of values to map of pointers
	config := map[string]int{
		"timeout":    30,
		"retries":    3,
		"maxWorkers": 10,
	}
	configPtrs := ptr.IntMap(config)
	fmt.Printf("Timeout pointer: %d\n", *configPtrs["timeout"])

	// Convert map of pointers back to values
	// Nil pointers become zero values
	settings := map[string]*string{
		"host":   ptr.String("localhost"),
		"port":   ptr.String("8080"),
		"scheme": nil, // Will become empty string
	}
	settingValues := ptr.ToStringMap(settings)
	fmt.Printf("Host: %s, Port: %s, Scheme: %q\n",
		settingValues["host"],
		settingValues["port"],
		settingValues["scheme"])

	// Works with floats
	prices := map[string]float64{
		"apple":  1.50,
		"banana": 0.75,
		"orange": 2.00,
	}
	pricePtrs := ptr.Float64Map(prices)
	fmt.Printf("Apple price: $%.2f\n", *pricePtrs["apple"])

	// Output:
	// Timeout pointer: 30
	// Host: localhost, Port: 8080, Scheme: ""
	// Apple price: $1.50
}

// Example_bulkOperations demonstrates using Map/Slice helpers for bulk data operations
func Example_bulkOperations() {
	// Prepare prices as pointers for optional pricing
	priceValues := map[string]float64{
		"1": 19.99,
		"2": 29.99,
		"3": 39.99,
	}
	prices := ptr.Float64Map(priceValues)

	// Prepare product IDs
	ids := []int{1, 2, 3, 4, 5}
	idPtrs := ptr.IntSlice(ids)

	fmt.Printf("Created %d price pointers\n", len(prices))
	fmt.Printf("Created %d ID pointers\n", len(idPtrs))
	fmt.Printf("Sample price: $%.2f\n", *prices["1"])
	fmt.Printf("Sample ID: %d\n", *idPtrs[0])

	// Output:
	// Created 3 price pointers
	// Created 5 ID pointers
	// Sample price: $19.99
	// Sample ID: 1
}

// Example_or demonstrates using Or for fallback values
func Example_or() {
	// Configuration with optional override
	var userTheme *string // nil means use default
	defaultTheme := ptr.String("dark")

	selectedTheme := ptr.Or(userTheme, defaultTheme)
	fmt.Printf("Using theme: %s\n", *selectedTheme)

	// With user preference
	userTheme = ptr.String("light")
	selectedTheme = ptr.Or(userTheme, defaultTheme)
	fmt.Printf("Using theme: %s\n", *selectedTheme)

	// Output:
	// Using theme: dark
	// Using theme: light
}

// Example_nonZero demonstrates omitting zero values for APIs
func Example_nonZero() {
	// Creating API request with only non-zero values
	type UpdateRequest struct {
		Name  *string `json:"name,omitempty"`
		Age   *int    `json:"age,omitempty"`
		Email *string `json:"email,omitempty"`
	}

	// Only update name and email, leave age unchanged
	req := UpdateRequest{
		Name:  ptr.NonZero("Alice"),      // included
		Age:   ptr.NonZero(0),            // nil, omitted from JSON
		Email: ptr.NonZero("a@test.com"), // included
	}

	if req.Name != nil {
		fmt.Printf("Name: %s\n", *req.Name)
	}
	if req.Age != nil {
		fmt.Printf("Age: %d\n", *req.Age)
	} else {
		fmt.Println("Age: not set")
	}
	if req.Email != nil {
		fmt.Printf("Email: %s\n", *req.Email)
	}

	// Output:
	// Name: Alice
	// Age: not set
	// Email: a@test.com
}

// Example_filter demonstrates conditional pointer filtering
func Example_filter() {
	// Filter valid ages (18+)
	ages := []*int{ptr.Int(15), ptr.Int(25), ptr.Int(17), ptr.Int(30)}

	fmt.Println("Ages 18+:")
	for _, age := range ages {
		validAge := ptr.Filter(age, func(a int) bool { return a >= 18 })
		if validAge != nil {
			fmt.Printf("  %d\n", *validAge)
		}
	}

	// Output:
	// Ages 18+:
	//   25
	//   30
}

// Example_flatMap demonstrates chaining transformations
func Example_flatMap() {
	// Parse string to int, then check if even
	nums := []string{"42", "invalid", "17", "100"}

	fmt.Println("Even numbers:")
	for _, s := range nums {
		result := ptr.FlatMap(ptr.String(s), func(str string) *int {
			// Try to parse
			var val int
			if _, err := fmt.Sscanf(str, "%d", &val); err == nil {
				if val%2 == 0 {
					return ptr.Int(val)
				}
			}
			return nil
		})
		if result != nil {
			fmt.Printf("  %d\n", *result)
		}
	}

	// Output:
	// Even numbers:
	//   42
	//   100
}

// Example_modify demonstrates in-place transformations
func Example_modify() {
	// Apply discount to price
	price := ptr.Float64(100.0)

	fmt.Printf("Original price: $%.2f\n", *price)

	// Apply 20% discount
	ptr.Modify(price, func(p float64) float64 {
		return p * 0.8
	})

	fmt.Printf("Discounted price: $%.2f\n", *price)

	// Output:
	// Original price: $100.00
	// Discounted price: $80.00
}

// Example_apply demonstrates side effects on pointer values
func Example_apply() {
	// Log configuration value if set
	debugMode := ptr.Bool(true)

	ptr.Apply(debugMode, func(enabled bool) {
		if enabled {
			fmt.Println("Debug mode is enabled")
		}
	})

	// Won't print anything for nil
	var traceMode *bool
	executed := ptr.Apply(traceMode, func(enabled bool) {
		fmt.Println("This won't print")
	})

	if !executed {
		fmt.Println("Trace mode not configured")
	}

	// Output:
	// Debug mode is enabled
	// Trace mode not configured
}

// Example_isZero demonstrates checking for zero values
func Example_isZero() {
	// Check various pointer states
	var nilPtr *int
	zeroPtr := ptr.Int(0)
	valuePtr := ptr.Int(42)

	fmt.Printf("nil pointer is zero: %v\n", ptr.IsZero(nilPtr))
	fmt.Printf("pointer to 0 is zero: %v\n", ptr.IsZero(zeroPtr))
	fmt.Printf("pointer to 42 is zero: %v\n", ptr.IsZero(valuePtr))

	// Useful for validation
	emptyName := ptr.String("")
	if ptr.IsZero(emptyName) {
		fmt.Println("Name is required")
	}

	// Output:
	// nil pointer is zero: true
	// pointer to 0 is zero: true
	// pointer to 42 is zero: false
	// Name is required
}

// Example_swap demonstrates exchanging pointer values
func Example_swap() {
	// Swap two configuration values
	primary := ptr.String("server-a")
	backup := ptr.String("server-b")

	fmt.Printf("Before: primary=%s, backup=%s\n", *primary, *backup)

	ptr.Swap(primary, backup)

	fmt.Printf("After: primary=%s, backup=%s\n", *primary, *backup)

	// Output:
	// Before: primary=server-a, backup=server-b
	// After: primary=server-b, backup=server-a
}

// Example_functionalChaining demonstrates combining multiple operations
func Example_functionalChaining() {
	// Complex data transformation pipeline
	input := ptr.String("  hello world  ")

	// Chain: trim, uppercase, check length
	result := ptr.Filter(
		ptr.Map(input, func(s string) string {
			// Trim spaces
			trimmed := ""
			start, end := 0, len(s)-1
			for start <= end && s[start] == ' ' {
				start++
			}
			for end >= start && s[end] == ' ' {
				end--
			}
			if start <= end {
				trimmed = s[start : end+1]
			}
			// Convert to uppercase
			upper := ""
			for _, c := range trimmed {
				if c >= 'a' && c <= 'z' {
					upper += string(c - 32)
				} else {
					upper += string(c)
				}
			}
			return upper
		}),
		func(s string) bool { return len(s) > 5 },
	)

	if result != nil {
		fmt.Printf("Result: %s\n", *result)
	} else {
		fmt.Println("Result: filtered out")
	}

	// Output:
	// Result: HELLO WORLD
}

// Example_bind demonstrates using Bind for monadic operations
func Example_bind() {
	// Parse string to int with validation
	parseInt := func(s string) *int {
		var val int
		if _, err := fmt.Sscanf(s, "%d", &val); err == nil {
			return ptr.To(val)
		}
		return nil
	}

	result1 := ptr.Bind(ptr.To("42"), parseInt)
	result2 := ptr.Bind(ptr.To("invalid"), parseInt)

	if result1 != nil {
		fmt.Printf("Parsed: %d\n", *result1)
	}
	if result2 == nil {
		fmt.Println("Failed to parse")
	}

	// Output:
	// Parsed: 42
	// Failed to parse
}

// Example_getOr demonstrates configuration defaults
func Example_getOr() {
	type Config struct {
		Timeout    *int
		MaxRetries *int
		Debug      *bool
	}

	config := Config{
		Timeout: ptr.To(60),
		// MaxRetries not set
		Debug: ptr.To(true),
	}

	timeout := ptr.GetOr(config.Timeout, 30)
	retries := ptr.GetOr(config.MaxRetries, 3)
	debug := ptr.GetOr(config.Debug, false)

	fmt.Printf("Timeout: %d\n", timeout)
	fmt.Printf("Retries: %d\n", retries)
	fmt.Printf("Debug: %v\n", debug)

	// Output:
	// Timeout: 60
	// Retries: 3
	// Debug: true
}

// Example_tableValidation demonstrates table-driven validation patterns
func Example_tableValidation() {
	tests := []struct {
		age  *int
		want string
	}{
		{ptr.To(25), "adult"},
		{ptr.To(15), "minor"},
		{ptr.To(18), "adult"},
		{nil, "unknown"},
	}

	for _, tt := range tests {
		result := ptr.Filter(tt.age, func(v int) bool { return v >= 18 })
		status := "minor"
		if result != nil {
			status = "adult"
		}
		if tt.age == nil {
			status = "unknown"
		}
		fmt.Printf("Age %v is %s\n", ptr.From(tt.age), status)
	}

	// Output:
	// Age 25 is adult
	// Age 15 is minor
	// Age 18 is adult
	// Age 0 is unknown
}
