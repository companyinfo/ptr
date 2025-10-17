# ptr

A lightweight, production-ready Go package for pointer operations with zero dependencies and full generic type support.

[![Go Reference](https://pkg.go.dev/badge/go.companyinfo.dev/ptr.svg)](https://pkg.go.dev/go.companyinfo.dev/ptr)
[![Go Report Card](https://goreportcard.com/badge/go.companyinfo.dev/ptr)](https://goreportcard.com/report/go.companyinfo.dev/ptr)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## Overview

This package solves common pointer-related challenges in Go:

- Creating pointers from literals and constants
- Safe dereferencing without nil panics
- Optional field handling in APIs and configuration
- Batch pointer operations on slices and maps
- Type-safe conversions for all Go built-in types

**Key Features:**

- Generic functions supporting any type (Go 1.18+)
- Type-specific helpers for better IDE autocomplete
- Zero dependencies, pure Go standard library
- Sub-nanosecond performance with zero allocations
- Comprehensive test coverage and production-proven

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Core Concepts](#core-concepts)
- [API Documentation](#api-documentation)
  - [Generic Functions](#generic-functions)
  - [Slice Operations](#slice-operations)
  - [Map Operations](#map-operations)
  - [Utility Functions](#utility-functions)
  - [Type-Specific Functions](#type-specific-functions)
  - [Type-Specific Slice Functions](#type-specific-slice-functions)
  - [Type-Specific Map Functions](#type-specific-map-functions)
- [Practical Examples](#practical-examples)
- [API Reference](#api-reference)
- [Performance](#performance)
- [Best Practices](#best-practices)
- [FAQ](#faq)
- [Contributing](#contributing)
- [License](#license)

## Installation

```bash
go get go.companyinfo.dev/ptr
```

**Requirements:**

- Go 1.18 or later (for generics support)
- No external dependencies

## Quick Start

```go
package main

import (
    "fmt"
    "go.companyinfo.dev/ptr"
)

func main() {
    // Problem: Can't take address of literals
    // var name *string = &"Alice"  // ❌ Compile error
    
    // Solution: Use ptr.String()
    name := ptr.String("Alice")    // ✅ Works
    
    // Problem: Nil dereference causes panic
    var nilPtr *string
    // fmt.Println(*nilPtr)         // ❌ Panic!
    
    // Solution: Safe dereferencing
    fmt.Println(ptr.ToString(nilPtr))  // ✅ Returns "" safely
    
    // Generic support for any type
    type User struct { Name string }
    user := ptr.To(User{Name: "Bob"})
    fmt.Println(ptr.From(user).Name)   // "Bob"
}
```

## Core Concepts

### Why Pointers in Go?

Pointers are essential in Go for:

- **Optional fields**: Distinguishing between zero values and missing values
- **API design**: Making struct fields optional (e.g., `omitempty` in JSON)
- **Database nullability**: Representing NULL values from SQL databases
- **Performance**: Avoiding large struct copies

### The Pointer Problem

Go doesn't allow direct address-taking of literals or constants:

```go
// These don't work:
s := &"hello"              // ❌ Cannot take address of "hello"
i := &42                   // ❌ Cannot take address of 42
b := &true                 // ❌ Cannot take address of true

// Traditional workaround:
str := "hello"
s := &str                  // ✅ Works but verbose

// With ptr:
s := ptr.String("hello")   // ✅ Clean and simple
i := ptr.Int(42)
b := ptr.Bool(true)
```

### Safe Dereferencing

Dereferencing nil pointers causes runtime panics. This package provides safe alternatives:

```go
var s *string  // nil

// Unsafe:
// value := *s  // ❌ Panic!

// Safe with ptr:
value := ptr.From(s)              // Returns "" (zero value)
value := ptr.FromOr(s, "default") // Returns "default"
value := ptr.MustFrom(s)          // Panics with clear message (for programmer errors)
```

## API Documentation

### Generic Functions

These functions work with any type using Go generics:

#### `To[T any](v T) *T`

Create a pointer from a value:

```go
s := ptr.To("hello")      // *string
i := ptr.To(42)           // *int
b := ptr.To(true)         // *bool

type User struct {
    Name string
    Age  int
}
u := ptr.To(User{Name: "Alice", Age: 30})  // *User
```

#### `From[T any](p *T) T`

Dereference a pointer with zero-value fallback:

```go
s := ptr.To("hello")
fmt.Println(ptr.From(s))  // "hello"

var nilStr *string
fmt.Println(ptr.From(nilStr))  // "" (empty string)

var nilInt *int
fmt.Println(ptr.From(nilInt))  // 0
```

#### `FromOr[T any](p *T, defaultValue T) T`

Dereference a pointer with custom default value:

```go
s := ptr.To("hello")
fmt.Println(ptr.FromOr(s, "default"))  // "hello"

var nilStr *string
fmt.Println(ptr.FromOr(nilStr, "default"))  // "default"
```

#### `MustFrom[T any](p *T) T`

Dereference a pointer and panic if nil (use only when nil is a programming error):

```go
s := ptr.To("hello")
fmt.Println(ptr.MustFrom(s))  // "hello"

var nilStr *string
fmt.Println(ptr.MustFrom(nilStr))  // panics!
```

#### `Coalesce[T any](ptrs ...*T) *T`

Return the first non-nil pointer from a list:

```go
var envPort *int      // nil
var configPort *int   // nil
defaultPort := ptr.Int(8080)

port := ptr.Coalesce(envPort, configPort, defaultPort)
fmt.Println(*port)  // 8080
```

#### `Set[T any](p *T, value T) bool`

Safely set a pointer value with nil-check:

```go
p := ptr.To(42)
ptr.Set(p, 100)  // *p is now 100, returns true
ptr.Set[int](nil, 100)  // no-op, returns false
```

#### `Map[T, R any](p *T, fn func(T) R) *R`

Transform a pointer value with a function:

```go
s := ptr.To("hello")
length := ptr.Map(s, func(s string) int { return len(s) })
fmt.Println(*length)  // 5

// Returns nil if input is nil
var nilStr *string
result := ptr.Map(nilStr, func(s string) int { return len(s) })
// result is nil
```

### Slice Operations

#### `ToSlice[T any](values []T) []*T`

Convert a slice of values to a slice of pointers:

```go
ages := []int{25, 30, 35, 40}
agePointers := ptr.ToSlice(ages)
// []*int with pointers to each value
```

#### `FromSlice[T any](ptrs []*T) []T`

Convert a slice of pointers to a slice of values:

```go
pointers := []*int{ptr.Int(1), nil, ptr.Int(3)}
values := ptr.FromSlice(pointers)
fmt.Println(values)  // [1 0 3] - nil becomes zero value
```

### Map Operations

#### `ToMap[T any](values map[string]T) map[string]*T`

Convert a map with value type T to a map with pointer value type *T:

```go
config := map[string]int{
    "timeout": 30,
    "retries": 3,
}
configPtrs := ptr.ToMap(config)
// map[string]*int with pointer values
```

#### `FromMap[T any](ptrs map[string]*T) map[string]T`

Convert a map with pointer value type *T to a map with value type T:

```go
settings := map[string]*string{
    "host": ptr.String("localhost"),
    "port": nil,
}
values := ptr.FromMap(settings)
fmt.Println(values)  // map[host:localhost port:] - nil becomes empty string
```

### Utility Functions

#### `Equal[T comparable](a, b *T) bool`

Safely compare two pointers:

```go
a := ptr.To(42)
b := ptr.To(42)
c := ptr.To(99)

ptr.Equal(a, b)           // true (same value)
ptr.Equal(a, c)           // false (different values)
ptr.Equal[int](nil, nil)  // true (both nil)
ptr.Equal(a, nil)         // false (one nil)
```

#### `Copy[T any](p *T) *T`

Create a shallow copy of a pointer:

```go
original := ptr.To(42)
copied := ptr.Copy(original)

*original = 100
fmt.Println(*copied)  // 42 (unchanged)
```

#### `IsNil[T any](p *T) bool`

Check if a pointer is nil:

```go
s := ptr.To("hello")
ptr.IsNil(s)        // false
ptr.IsNil[string](nil)  // true
```

### Type-Specific Functions

For better IDE autocomplete and convenience, the package provides type-specific functions:

#### String

```go
name := ptr.String("Alice")
fmt.Println(ptr.ToString(name))     // "Alice"
fmt.Println(ptr.ToString(nil))      // ""
```

#### Int

```go
age := ptr.Int(30)
fmt.Println(ptr.ToInt(age))   // 30
fmt.Println(ptr.ToInt(nil))   // 0
```

#### Int64

```go
id := ptr.Int64(123456789)
fmt.Println(ptr.ToInt64(id))   // 123456789
fmt.Println(ptr.ToInt64(nil))  // 0
```

#### Bool

```go
active := ptr.Bool(true)
fmt.Println(ptr.ToBool(active))  // true
fmt.Println(ptr.ToBool(nil))     // false
```

#### Float64

```go
price := ptr.Float64(19.99)
fmt.Println(ptr.ToFloat64(price))  // 19.99
fmt.Println(ptr.ToFloat64(nil))    // 0.0
```

#### Additional Numeric Types

The package also provides helpers for all Go numeric types:

```go
// Integer types
i8 := ptr.Int8(127)
i16 := ptr.Int16(32767)
i32 := ptr.Int32(2147483647)

// Unsigned integer types
u := ptr.Uint(42)
u8 := ptr.Uint8(255)
u16 := ptr.Uint16(65535)
u32 := ptr.Uint32(4294967295)
u64 := ptr.Uint64(18446744073709551615)

// Float types
f32 := ptr.Float32(3.14)

// Character types
b := ptr.Byte('A')
r := ptr.Rune('世')

// Pointer type
uptr := ptr.Uintptr(0x1234)
```

#### Time Types

Helpers for working with time-related types:

```go
// Time helpers
now := ptr.Time(time.Now())
timestamp := ptr.ToTime(now)  // time.Time value
fmt.Println(ptr.ToTime(nil).IsZero())  // true

// Duration helpers
timeout := ptr.Duration(30 * time.Second)
d := ptr.ToDuration(timeout)  // time.Duration value
fmt.Println(ptr.ToDuration(nil))  // 0
```

#### Complex Numbers

Helpers for complex number types:

```go
// Complex64
c64 := ptr.Complex64(3 + 4i)
val64 := ptr.ToComplex64(c64)  // complex64 value

// Complex128
c128 := ptr.Complex128(3 + 4i)
val128 := ptr.ToComplex128(c128)  // complex128 value
```

### Type-Specific Slice Functions

For better IDE autocomplete and convenience, type-specific slice conversion functions are available for all common types:

#### String Slices

```go
names := []string{"Alice", "Bob", "Charlie"}
namePtrs := ptr.StringSlice(names)
// []*string with pointers to each name

mixed := []*string{ptr.String("Alice"), nil, ptr.String("Charlie")}
values := ptr.ToStringSlice(mixed)
fmt.Println(values)  // [Alice  Charlie] - nil becomes empty string
```

#### Integer Slices

```go
// Works with all integer types: Int, Int8, Int16, Int32, Int64
ages := []int{25, 30, 35}
agePtrs := ptr.IntSlice(ages)

ids := []int64{1001, 1002, 1003}
idPtrs := ptr.Int64Slice(ids)

// Convert back to values
values := ptr.ToIntSlice(agePtrs)  // [25 30 35]
```

#### Other Numeric Type Slices

```go
// Unsigned integers: Uint, Uint8, Uint16, Uint32, Uint64
counts := []uint{1, 2, 3}
countPtrs := ptr.UintSlice(counts)

// Floats: Float32, Float64
prices := []float64{19.99, 29.99, 39.99}
pricePtrs := ptr.Float64Slice(prices)

// Bytes
data := []byte{0x01, 0x02, 0x03}
dataPtrs := ptr.ByteSlice(data)
```

#### Boolean Slices

```go
flags := []bool{true, false, true}
flagPtrs := ptr.BoolSlice(flags)

values := ptr.ToBoolSlice(flagPtrs)  // [true false true]
```

#### Time Type Slices

```go
// Time slices
timestamps := []time.Time{time.Now(), time.Now().Add(time.Hour)}
timePtrs := ptr.TimeSlice(timestamps)

// Duration slices
durations := []time.Duration{time.Second, time.Minute, time.Hour}
durationPtrs := ptr.DurationSlice(durations)
```

**Available slice functions for all types:**

- `IntSlice`, `Int8Slice`, `Int16Slice`, `Int32Slice`, `Int64Slice`
- `UintSlice`, `Uint8Slice`, `Uint16Slice`, `Uint32Slice`, `Uint64Slice`
- `Float32Slice`, `Float64Slice`
- `BoolSlice`, `ByteSlice`, `StringSlice`
- `TimeSlice`, `DurationSlice`

And their corresponding `To*Slice` functions for converting back to values.

### Type-Specific Map Functions

Type-specific map conversion functions for converting between `map[string]T` and `map[string]*T`:

#### String Maps

```go
settings := map[string]string{
    "host": "localhost",
    "port": "8080",
}
settingPtrs := ptr.StringMap(settings)
// map[string]*string with pointer values

values := ptr.ToStringMap(settingPtrs)
// Back to map[string]string
```

#### Integer Maps

```go
// Works with all integer types: Int, Int8, Int16, Int32, Int64
config := map[string]int{
    "timeout":  30,
    "retries":  3,
    "maxConns": 100,
}
configPtrs := ptr.IntMap(config)

// With nil values
mixed := map[string]*int{
    "timeout": ptr.Int(30),
    "retries": nil,  // Will become 0
}
values := ptr.ToIntMap(mixed)  // map[string]int
```

#### Other Numeric Type Maps

```go
// Unsigned integers: Uint, Uint8, Uint16, Uint32, Uint64
limits := map[string]uint64{
    "maxSize":  1024000,
    "maxFiles": 100,
}
limitPtrs := ptr.Uint64Map(limits)

// Floats: Float32, Float64
prices := map[string]float64{
    "basic":    9.99,
    "premium":  19.99,
    "enterprise": 99.99,
}
pricePtrs := ptr.Float64Map(prices)
```

#### Boolean Maps

```go
features := map[string]bool{
    "caching":    true,
    "monitoring": true,
    "debug":      false,
}
featurePtrs := ptr.BoolMap(features)

values := ptr.ToBoolMap(featurePtrs)
```

#### Time Type Maps

```go
// Time maps
events := map[string]time.Time{
    "created":  time.Now(),
    "updated":  time.Now(),
}
eventPtrs := ptr.TimeMap(events)

// Duration maps
timeouts := map[string]time.Duration{
    "read":  30 * time.Second,
    "write": 10 * time.Second,
}
timeoutPtrs := ptr.DurationMap(timeouts)
```

**Available map functions for all types:**

- `IntMap`, `Int8Map`, `Int16Map`, `Int32Map`, `Int64Map`
- `UintMap`, `Uint8Map`, `Uint16Map`, `Uint32Map`, `Uint64Map`
- `Float32Map`, `Float64Map`
- `BoolMap`, `ByteMap`, `StringMap`
- `TimeMap`, `DurationMap`

And their corresponding `To*Map` functions for converting back to value maps.

## Practical Examples

### REST API with Optional Fields

Building APIs with optional query parameters and request bodies:

```go
// Request model with optional filters
type ListUsersRequest struct {
    Page     int     `json:"page"`
    PageSize int     `json:"page_size"`
    Role     *string `json:"role,omitempty"`      // Optional filter
    Active   *bool   `json:"active,omitempty"`    // Optional filter
    MinAge   *int    `json:"min_age,omitempty"`   // Optional filter
}

// Handler function
func ListUsers(req ListUsersRequest) ([]User, error) {
    query := "SELECT * FROM users WHERE 1=1"
    args := []interface{}{}
    
    // Build dynamic query based on optional fields
    if !ptr.IsNil(req.Role) {
        query += " AND role = ?"
        args = append(args, ptr.ToString(req.Role))
    }
    if !ptr.IsNil(req.Active) {
        query += " AND active = ?"
        args = append(args, ptr.ToBool(req.Active))
    }
    if !ptr.IsNil(req.MinAge) {
        query += " AND age >= ?"
        args = append(args, ptr.ToInt(req.MinAge))
    }
    
    // Execute query...
    return fetchUsers(query, args...)
}

// Client usage
req := ListUsersRequest{
    Page:     1,
    PageSize: 50,
    Role:     ptr.String("admin"),  // Filter by role
    Active:   ptr.Bool(true),       // Filter by active status
    MinAge:   nil,                  // Don't filter by age
}
```

### JSON Marshaling with Optional Fields

```go
type User struct {
    Name  string  `json:"name"`
    Email *string `json:"email,omitempty"`
    Age   *int    `json:"age,omitempty"`
}

user := User{
    Name:  "Alice",
    Email: ptr.String("alice@example.com"),
    Age:   ptr.Int(30),
}

// Email and Age will be included in JSON
// If set to nil, they will be omitted
```

### Working with API Responses

```go
type APIResponse struct {
    Data  *UserData
    Error *string
}

func processResponse(resp APIResponse) {
    if !ptr.IsNil(resp.Error) {
        log.Printf("Error: %s", ptr.ToString(resp.Error))
        return
    }
    
    // Safely access data with fallback
    name := ptr.FromOr(resp.Data.Name, "Unknown")
    fmt.Printf("User: %s\n", name)
}
```

### Function Parameters with Optional Values

```go
func CreateUser(name string, age *int, email *string) User {
    return User{
        Name:  name,
        Age:   ptr.ToInt(age),        // 0 if nil
        Email: ptr.ToString(email),   // "" if nil
    }
}

// Call with optional parameters
user1 := CreateUser("Alice", ptr.Int(30), ptr.String("alice@example.com"))
user2 := CreateUser("Bob", nil, nil)
```

### Database NULL Values

```go
type Product struct {
    ID          int
    Name        string
    Description *string  // NULL in database if nil
    Price       *float64 // NULL in database if nil
}

func getProduct(id int) Product {
    // ... fetch from database
    return Product{
        ID:          id,
        Name:        "Product Name",
        Description: ptr.String("A great product"),
        Price:       ptr.Float64(29.99),
    }
}
```

### Configuration with Fallbacks and Precedence

Real-world configuration loading with multiple sources:

```go
type AppConfig struct {
    Host       string
    Port       int
    Timeout    time.Duration
    MaxRetries int
    Debug      bool
}

// LoadConfig demonstrates precedence: env vars > config file > defaults
func LoadConfig() AppConfig {
    // Try to load from environment
    var envPort *int
    if portStr := os.Getenv("APP_PORT"); portStr != "" {
        if port, err := strconv.Atoi(portStr); err == nil {
            envPort = ptr.Int(port)
        }
    }
    
    // Try to load from config file
    var filePort *int
    if config := loadConfigFile(); config != nil {
        filePort = config.Port
    }
    
    // Define defaults
    defaultPort := ptr.Int(8080)
    defaultTimeout := ptr.Duration(30 * time.Second)
    defaultRetries := ptr.Int(3)
    defaultDebug := ptr.Bool(false)
    
    // Use Coalesce for precedence: env > file > default
    return AppConfig{
        Host:       getEnvOr("APP_HOST", "localhost"),
        Port:       ptr.ToInt(ptr.Coalesce(envPort, filePort, defaultPort)),
        Timeout:    ptr.ToDuration(ptr.Coalesce(getEnvDuration("APP_TIMEOUT"), defaultTimeout)),
        MaxRetries: ptr.ToInt(ptr.Coalesce(getEnvInt("MAX_RETRIES"), defaultRetries)),
        Debug:      ptr.ToBool(ptr.Coalesce(getEnvBool("DEBUG"), defaultDebug)),
    }
}

// Helper to get env var with default
func getEnvOr(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

### Partial Updates and PATCH Operations

Handling partial updates where only provided fields should be updated:

```go
type UpdateUserRequest struct {
    Name     *string `json:"name,omitempty"`
    Email    *string `json:"email,omitempty"`
    Age      *int    `json:"age,omitempty"`
    Active   *bool   `json:"active,omitempty"`
}

func (s *UserService) UpdateUser(id int, req UpdateUserRequest) error {
    // Only update fields that were explicitly provided
    updates := make(map[string]interface{})
    
    if !ptr.IsNil(req.Name) {
        updates["name"] = ptr.ToString(req.Name)
    }
    if !ptr.IsNil(req.Email) {
        updates["email"] = ptr.ToString(req.Email)
    }
    if !ptr.IsNil(req.Age) {
        updates["age"] = ptr.ToInt(req.Age)
    }
    if !ptr.IsNil(req.Active) {
        updates["active"] = ptr.ToBool(req.Active)
    }
    
    if len(updates) == 0 {
        return errors.New("no fields to update")
    }
    
    return s.db.UpdateUser(id, updates)
}

// Usage:
// PATCH /users/123 with {"email": "new@example.com"}
// Only email is updated, other fields remain unchanged
req := UpdateUserRequest{
    Email: ptr.String("new@example.com"),
    // Name, Age, Active are nil - won't be updated
}
```

### Batch Processing with Slices

```go
// Convert user IDs to pointers for JSON
userIDs := []int64{1001, 1002, 1003}
idPointers := ptr.Int64Slice(userIDs)

// Process results
results := []*ProcessResult{
    processUser(idPointers[0]),
    processUser(idPointers[1]),
    processUser(idPointers[2]),
}

// Extract values (nil results become zero values)
resultValues := ptr.FromSlice(results)
```

### Bulk Data Conversion

Working with bulk data using type-specific helpers:

```go
// Convert multiple prices at once
priceList := []float64{9.99, 19.99, 29.99, 39.99}
pricePtrs := ptr.Float64Slice(priceList)

// Use in API response
type ProductList struct {
    Prices []*float64 `json:"prices,omitempty"`
}

// Convert configuration maps
envVars := map[string]string{
    "DATABASE_HOST": "localhost",
    "DATABASE_PORT": "5432",
    "API_KEY":       "secret",
}
envVarPtrs := ptr.StringMap(envVars)

// Process and convert back
processedVars := processConfig(envVarPtrs)
finalConfig := ptr.ToStringMap(processedVars)
```

### API Response Transformation

```go
// Transform API response data
type APIUser struct {
    ID    int64
    Name  string
    Roles []string
}

// Convert roles to pointers for optional field handling
users := []APIUser{
    {ID: 1, Name: "Alice", Roles: []string{"admin", "user"}},
    {ID: 2, Name: "Bob", Roles: []string{"user"}},
}

// Extract IDs as pointers
ids := make([]int64, len(users))
for i, u := range users {
    ids[i] = u.ID
}
idPtrs := ptr.Int64Slice(ids)

// Use in bulk operations
results := bulkFetchUserData(idPtrs)
```

### Data Transformation

```go
// Transform pointer values without nil checks
names := []*string{
    ptr.String("alice"),
    ptr.String("bob"),
    nil,
}

// Map to uppercase
for i, name := range names {
    names[i] = ptr.Map(name, strings.ToUpper)
}
// Result: ["ALICE", "BOB", nil]
```

## API Reference

### Generic Function Reference

| Function | Description |
|----------|-------------|
| `To[T any](v T) *T` | Create a pointer from a value |
| `From[T any](p *T) T` | Dereference with zero-value fallback |
| `FromOr[T any](p *T, defaultValue T) T` | Dereference with custom default |
| `MustFrom[T any](p *T) T` | Dereference and panic if nil |
| `Coalesce[T any](ptrs ...*T) *T` | Return first non-nil pointer from list |
| `Set[T any](p *T, value T) bool` | Set pointer value with nil-check |
| `Map[T, R any](p *T, fn func(T) R) *R` | Transform pointer value with function |
| `Equal[T comparable](a, b *T) bool` | Compare two pointers safely |
| `Copy[T any](p *T) *T` | Create a shallow copy of a pointer |
| `IsNil[T any](p *T) bool` | Check if pointer is nil |

### Slice Function Reference

| Function | Description |
|----------|-------------|
| `ToSlice[T any](values []T) []*T` | Convert slice of values to slice of pointers |
| `FromSlice[T any](ptrs []*T) []T` | Convert slice of pointers to slice of values |

### Map Function Reference

| Function | Description |
|----------|-------------|
| `ToMap[T any](values map[string]T) map[string]*T` | Convert map of values to map of pointer values |
| `FromMap[T any](ptrs map[string]*T) map[string]T` | Convert map of pointer values to map of values |

### Type-Specific Function Reference

#### Common Type Functions

| Type | Create | Dereference |
|------|--------|-------------|
| string | `String(v string) *string` | `ToString(p *string) string` |
| int | `Int(v int) *int` | `ToInt(p *int) int` |
| int64 | `Int64(v int64) *int64` | `ToInt64(p *int64) int64` |
| bool | `Bool(v bool) *bool` | `ToBool(p *bool) bool` |
| float64 | `Float64(v float64) *float64` | `ToFloat64(p *float64) float64` |

#### Numeric Type Functions

| Type | Create | Dereference |
|------|--------|-------------|
| int8 | `Int8(v int8) *int8` | `ToInt8(p *int8) int8` |
| int16 | `Int16(v int16) *int16` | `ToInt16(p *int16) int16` |
| int32 | `Int32(v int32) *int32` | `ToInt32(p *int32) int32` |
| uint | `Uint(v uint) *uint` | `ToUint(p *uint) uint` |
| uint8 | `Uint8(v uint8) *uint8` | `ToUint8(p *uint8) uint8` |
| uint16 | `Uint16(v uint16) *uint16` | `ToUint16(p *uint16) uint16` |
| uint32 | `Uint32(v uint32) *uint32` | `ToUint32(p *uint32) uint32` |
| uint64 | `Uint64(v uint64) *uint64` | `ToUint64(p *uint64) uint64` |
| float32 | `Float32(v float32) *float32` | `ToFloat32(p *float32) float32` |
| byte | `Byte(v byte) *byte` | `ToByte(p *byte) byte` |
| rune | `Rune(v rune) *rune` | `ToRune(p *rune) rune` |
| uintptr | `Uintptr(v uintptr) *uintptr` | `ToUintptr(p *uintptr) uintptr` |

#### Time and Complex Type Functions

| Type | Create | Dereference |
|------|--------|-------------|
| time.Time | `Time(v time.Time) *time.Time` | `ToTime(p *time.Time) time.Time` |
| time.Duration | `Duration(v time.Duration) *time.Duration` | `ToDuration(p *time.Duration) time.Duration` |
| complex64 | `Complex64(v complex64) *complex64` | `ToComplex64(p *complex64) complex64` |
| complex128 | `Complex128(v complex128) *complex128` | `ToComplex128(p *complex128) complex128` |

### Type-Specific Slice Function Reference

For each type, both slice conversion functions are available:

| Type | To Slice | From Slice |
|------|----------|------------|
| string | `StringSlice([]string) []*string` | `ToStringSlice([]*string) []string` |
| int | `IntSlice([]int) []*int` | `ToIntSlice([]*int) []int` |
| int8 | `Int8Slice([]int8) []*int8` | `ToInt8Slice([]*int8) []int8` |
| int16 | `Int16Slice([]int16) []*int16` | `ToInt16Slice([]*int16) []int16` |
| int32 | `Int32Slice([]int32) []*int32` | `ToInt32Slice([]*int32) []int32` |
| int64 | `Int64Slice([]int64) []*int64` | `ToInt64Slice([]*int64) []int64` |
| uint | `UintSlice([]uint) []*uint` | `ToUintSlice([]*uint) []uint` |
| uint8 | `Uint8Slice([]uint8) []*uint8` | `ToUint8Slice([]*uint8) []uint8` |
| uint16 | `Uint16Slice([]uint16) []*uint16` | `ToUint16Slice([]*uint16) []uint16` |
| uint32 | `Uint32Slice([]uint32) []*uint32` | `ToUint32Slice([]*uint32) []uint32` |
| uint64 | `Uint64Slice([]uint64) []*uint64` | `ToUint64Slice([]*uint64) []uint64` |
| float32 | `Float32Slice([]float32) []*float32` | `ToFloat32Slice([]*float32) []float32` |
| float64 | `Float64Slice([]float64) []*float64` | `ToFloat64Slice([]*float64) []float64` |
| bool | `BoolSlice([]bool) []*bool` | `ToBoolSlice([]*bool) []bool` |
| byte | `ByteSlice([]byte) []*byte` | `ToByteSlice([]*byte) []byte` |
| time.Time | `TimeSlice([]time.Time) []*time.Time` | `ToTimeSlice([]*time.Time) []time.Time` |
| time.Duration | `DurationSlice([]time.Duration) []*time.Duration` | `ToDurationSlice([]*time.Duration) []time.Duration` |

### Type-Specific Map Function Reference

For each type, both map conversion functions are available (all maps use `string` keys):

| Type | To Map | From Map |
|------|--------|----------|
| string | `StringMap(map[string]string) map[string]*string` | `ToStringMap(map[string]*string) map[string]string` |
| int | `IntMap(map[string]int) map[string]*int` | `ToIntMap(map[string]*int) map[string]int` |
| int8 | `Int8Map(map[string]int8) map[string]*int8` | `ToInt8Map(map[string]*int8) map[string]int8` |
| int16 | `Int16Map(map[string]int16) map[string]*int16` | `ToInt16Map(map[string]*int16) map[string]int16` |
| int32 | `Int32Map(map[string]int32) map[string]*int32` | `ToInt32Map(map[string]*int32) map[string]int32` |
| int64 | `Int64Map(map[string]int64) map[string]*int64` | `ToInt64Map(map[string]*int64) map[string]int64` |
| uint | `UintMap(map[string]uint) map[string]*uint` | `ToUintMap(map[string]*uint) map[string]uint` |
| uint8 | `Uint8Map(map[string]uint8) map[string]*uint8` | `ToUint8Map(map[string]*uint8) map[string]uint8` |
| uint16 | `Uint16Map(map[string]uint16) map[string]*uint16` | `ToUint16Map(map[string]*uint16) map[string]uint16` |
| uint32 | `Uint32Map(map[string]uint32) map[string]*uint32` | `ToUint32Map(map[string]*uint32) map[string]uint32` |
| uint64 | `Uint64Map(map[string]uint64) map[string]*uint64` | `ToUint64Map(map[string]*uint64) map[string]uint64` |
| float32 | `Float32Map(map[string]float32) map[string]*float32` | `ToFloat32Map(map[string]*float32) map[string]float32` |
| float64 | `Float64Map(map[string]float64) map[string]*float64` | `ToFloat64Map(map[string]*float64) map[string]float64` |
| bool | `BoolMap(map[string]bool) map[string]*bool` | `ToBoolMap(map[string]*bool) map[string]bool` |
| byte | `ByteMap(map[string]byte) map[string]*byte` | `ToByteMap(map[string]*byte) map[string]byte` |
| time.Time | `TimeMap(map[string]time.Time) map[string]*time.Time` | `ToTimeMap(map[string]*time.Time) map[string]time.Time` |
| time.Duration | `DurationMap(map[string]time.Duration) map[string]*time.Duration` | `ToDurationMap(map[string]*time.Duration) map[string]time.Duration` |

## Performance

The package has **minimal overhead** with most operations optimized to near-zero cost by the Go compiler. Benchmark results:

```text
BenchmarkTo-12                  1000000000    0.12 ns/op    0 B/op    0 allocs/op
BenchmarkFrom-12                1000000000    0.15 ns/op    0 B/op    0 allocs/op
BenchmarkFromOr-12              1000000000    0.11 ns/op    0 B/op    0 allocs/op
BenchmarkEqual-12               1000000000    0.13 ns/op    0 B/op    0 allocs/op
BenchmarkCopy-12                1000000000    0.14 ns/op    0 B/op    0 allocs/op
BenchmarkIsNil-12               1000000000    0.12 ns/op    0 B/op    0 allocs/op
```

**Key Performance Characteristics:**

- Sub-nanosecond operations - All functions complete in ~0.1-0.2 ns
- Zero allocations - No heap allocations for pointer operations
- Compiler optimized - Functions are typically inlined by the Go compiler
- Production ready - Performance suitable for hot code paths

Run benchmarks yourself:

```bash
go test -bench=. -benchmem
```

## Best Practices

### When to Use Pointers

**Use pointers for:**

- Optional fields in structs (especially for JSON/API models)
- Distinguishing between "not set" and "zero value"
- Database NULL values
- Configuration with multiple precedence levels
- Large structs to avoid copying overhead

**Avoid pointers for:**

- Small, frequently-accessed values (int, bool) in hot paths
- Values that should never be nil
- Internal function parameters unless needed for mutation
- Simple data without optional semantics

### Choosing Between `From`, `FromOr`, and `MustFrom`

```go
// Use From() when zero value is acceptable
age := ptr.From(user.Age)  // 0 if nil

// Use FromOr() when you need a specific default
timeout := ptr.FromOr(config.Timeout, 30*time.Second)

// Use MustFrom() only when nil indicates a programmer error
// (not for user input or external data)
id := ptr.MustFrom(user.ID)  // Panic if nil = bug in code
```

### Generic vs Type-Specific Functions

```go
// Prefer type-specific functions for common types (better IDE support)
name := ptr.String("Alice")     // ✅ Clear and autocomplete-friendly
age := ptr.Int(30)

// Use generics for custom types or when type-agnostic code is needed
type UserID int64
id := ptr.To[UserID](12345)     // ✅ Works with custom types

func GetValue[T any](p *T, defaultVal T) T {
    return ptr.FromOr(p, defaultVal)  // ✅ Generic function
}
```

### Error Handling Patterns

```go
// Anti-pattern: Silently converting errors to nil
func getUser(id int) *User {
    user, err := db.GetUser(id)
    if err != nil {
        return nil  // ❌ Lost error information
    }
    return &user
}

// Better: Return both value and error
func getUser(id int) (*User, error) {
    user, err := db.GetUser(id)
    if err != nil {
        return nil, err  // ✅ Preserve error
    }
    return &user, nil
}

// When using ptr for optional values, consider validation
func updateUser(req UpdateUserRequest) error {
    if !ptr.IsNil(req.Email) {
        email := ptr.ToString(req.Email)
        if !isValidEmail(email) {
            return errors.New("invalid email")  // ✅ Validate before use
        }
    }
    return nil
}
```

### Memory and Performance Considerations

```go
// Each pointer adds indirection and potential cache misses
// For hot paths with small values, benchmark both approaches

// Approach 1: Pointers (flexible, optional semantics)
type Config struct {
    MaxRetries *int
    Timeout    *time.Duration
}

// Approach 2: Values with sentinel (no indirection, better cache locality)
type Config struct {
    MaxRetries int  // 0 or -1 means "not set"
    Timeout    time.Duration
}

// Choose based on your use case:
// - API/JSON models: Pointers (distinguish null from zero)
// - Internal hot paths: Values (better performance)
// - Configuration: Pointers (clear optional semantics)
```

## FAQ

### Q: When should I use this package vs standard Go?

**A:** Use this package when you need:

- Pointers to literals: `ptr.Int(42)` vs the verbose two-line alternative
- Safe nil handling: `ptr.From(x)` vs manual nil checks
- Optional API fields: Clear distinction between "not provided" and "zero value"

**Standard Go is fine when:**

- You don't need pointers at all
- You're comfortable with manual nil checks
- You prefer zero dependencies (though this package has zero deps too)

### Q: Is it safe to use in production?

**A:** Yes. The package is:

- Battle-tested in production environments
- Fully tested with comprehensive coverage
- Zero external dependencies
- Simple, focused API with no surprises
- Performance-optimized (sub-nanosecond operations)

### Q: Why use `ptr.From()` instead of checking nil manually?

**A:** Compare:

```go
// Manual nil check (verbose, repetitive)
var name string
if user.Name != nil {
    name = *user.Name
}
// name is "" if user.Name was nil

// With ptr (concise, clear intent)
name := ptr.From(user.Name)
```

The benefit increases with multiple optional fields.

### Q: What about `sql.Null*` types for database work?

**A:** Both approaches are valid:

```go
// Using sql.NullString
type User struct {
    Email sql.NullString
}
// Pros: Standard library, clear database intent
// Cons: Verbose to work with, doesn't play well with JSON

// Using *string with ptr
type User struct {
    Email *string `json:"email,omitempty"`
}
// Pros: Works with JSON, APIs, and databases
// Cons: Need to handle conversion for SQL

// You can combine both:
func (u User) ToSQL() UserDB {
    return UserDB{
        Email: sql.NullString{
            String: ptr.ToString(u.Email),
            Valid:  !ptr.IsNil(u.Email),
        },
    }
}
```

### Q: Does this work with older Go versions?

**A:** Requires Go 1.18+ for generics. For older versions:

- Use type-specific functions (they don't require generics)
- Or stick with Go 1.17 patterns (manual pointer handling)

### Q: How does this compare to similar packages?

**A:** This package focuses on:

- **Simplicity**: Small, focused API
- **Performance**: Zero allocation, compiler-optimized
- **Type safety**: Leverages Go generics
- **Zero dependencies**: Pure standard library

Similar packages may offer different trade-offs. This package prioritizes simplicity and performance.

### Q: Can I use this with reflection or JSON unmarshaling?

**A:** Yes, pointers created by this package are regular Go pointers:

```go
type User struct {
    Name *string `json:"name,omitempty"`
}

// JSON unmarshaling works normally
json.Unmarshal(data, &user)

// Creating for marshaling
user := User{Name: ptr.String("Alice")}
json.Marshal(user)  // {"name":"Alice"}

user.Name = nil
json.Marshal(user)  // {} - field omitted
```

## Stability and Versioning

This package follows [Semantic Versioning](https://semver.org/):

- **v1.x.x**: Current stable version with backward compatibility guarantee
- No breaking changes in minor or patch releases
- New features added in minor versions (v1.1.0, v1.2.0, etc.)
- Bug fixes in patch versions (v1.0.1, v1.0.2, etc.)

The public API is stable and production-ready. We're committed to maintaining backward compatibility.

## License

Apache License 2.0 - see [LICENSE](LICENSE) file for details.

## Contributing

We welcome contributions! Here's how you can help:

### Reporting Issues

- Check if the issue already exists before creating a new one
- Include Go version, operating system, and minimal reproduction code
- Describe expected vs actual behavior

### Pull Requests

Before submitting a PR:

1. **Write tests** - Ensure your changes are covered by tests
2. **Run tests** - `go test -v ./...`
3. **Run benchmarks** - `go test -bench=. -benchmem` (if performance-related)
4. **Format code** - `go fmt ./...`
5. **Update documentation** - Keep README and godoc comments current

### Development Setup

```bash
# Clone the repository
git clone https://github.com/companyinfo/ptr.git
cd ptr

# Run tests
go test -v

# Run benchmarks
go test -bench=. -benchmem

# Run with coverage
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Code Standards

- Follow Go best practices and idioms
- Maintain backward compatibility unless it's a major version bump
- Keep functions simple and focused
- Add godoc comments for public functions
- Ensure zero allocations for performance-critical functions

### What We're Looking For

- Bug fixes with test cases
- Performance improvements with benchmarks
- Documentation improvements
- Additional type-specific helpers (if commonly needed)
- Real-world use case examples

### What We're Not Looking For

- Breaking changes without discussion
- Features that add external dependencies
- Complex features that increase API surface unnecessarily

---

**[Documentation](https://pkg.go.dev/go.companyinfo.dev/ptr)** · **[Report Bug](https://github.com/companyinfo/ptr/issues)** · **[Request Feature](https://github.com/companyinfo/ptr/issues)**

Made with ❤️ by the CompanyInfo team
