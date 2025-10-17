package ptr_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/companyinfo/ptr"
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
