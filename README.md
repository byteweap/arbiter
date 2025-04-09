<p align="center">
  <img src="arbiter.png" alt="Arbiter Logo" width="200"/>
</p>

# Arbiter

A powerful and flexible data validation framework for Go.

English | [简体中文](README-CN.md)

## Overview

Arbiter is a comprehensive data validation framework written in Go that provides rich validation rules and flexible validation mechanisms. It supports validation of basic data types, strings, numbers, time, files, and more, with built-in support for struct field validation.

## Features

- Generic support for validating any data type
- Rich set of built-in validation rules
- Custom validation rule support
- Struct field validation
- Chainable API
- Custom error messages
- Conditional validation
- Dependency validation

## Installation

```bash
go get github.com/byteweap/arbiter
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/byteweap/arbiter"
    "github.com/byteweap/arbiter/rule"
)

type Person struct {
    Name  string
    Age   int
    Email string
}

func main() {
    person := &Person{
        Name:  "John",
        Age:   30,
        Email: "john@example.com",
    }

    err := arbiter.ValidateStruct(person, "Person cannot be nil",
        rule.Field(&person.Name,
            rule.Length(2, 50),
            rule.String().Errf("Name is required"),
        ),
        rule.Field(&person.Age,
            rule.Min(0),
            rule.Max(120),
        ),
        rule.Field(&person.Email,
            rule.Email(),
            rule.String().Errf("Invalid email"),
        ),
    )

    if err != nil {
        fmt.Printf("Validation error: %v\n", err)
    }
}
```

## Core Components

### 1. Validator (Arbiter)

The main validation functions:

```go
// Validate applies multiple rules to a single value
err := Validate("hello",
    rule.Length(3, 10),
    rule.String().Errf("Invalid string"),
)

// ValidateWithErrs collects all validation errors
err := ValidateWithErrs("hello",
    rule.Length(3, 10),
    rule.String().Errf("Invalid string"),
)

// ValidateStruct validates a struct and its fields
err := ValidateStruct(person, "Person cannot be nil",
    rule.Field(&person.Name, ...),
    rule.Field(&person.Age, ...),
)
```

### 2. Field Validator

For validating struct fields:

```go
// Create a field validation rule
nameRule := rule.Field(&person.Name,
    rule.Length(2, 50),
    rule.String().Errf("Name is required"),
)
```

### 3. Validation Rules

#### String Rules
- `StartWith`: Validate string prefix
- `EndWith`: Validate string suffix
- `ChineseOnly`: Validate Chinese characters
- `FullWidthOnly`: Validate full-width characters
- `HalfWidthOnly`: Validate half-width characters
- `UpperCaseOnly`: Validate uppercase letters
- `LowerCaseOnly`: Validate lowercase letters
- `SpecialChars`: Validate special characters
- `Contains`: Validate substring presence
- `NotContains`: Validate substring absence

#### Numeric Rules
- `Min`: Minimum value
- `Max`: Maximum value
- `Between`: Range validation
- `Positive`: Positive number validation
- `Negative`: Negative number validation
- `Even`: Even number validation
- `Odd`: Odd number validation
- `Precision`: Decimal precision validation

#### Time Rules
- `Before`: Earlier than specified time
- `After`: Later than specified time
- `Between`: Time range validation

#### File Rules
- `Exists`: File existence validation
- `Size`: File size validation
- `Extension`: File extension validation

#### Network Rules
- `IP`: IP address validation
- `URL`: URL validation
- `Email`: Email validation

## Best Practices

### 1. Error Handling

```go
// Collect all validation errors
err := Validate(value,
    rule1,
    rule2,
    rule3,
)
if err != nil {
    // Handle first error
}
```

```go
// Collect all validation errors
errs := ValidateWithErrs(value,
    rule1,
    rule2,
    rule3,
)
if len(errs) > 0  {
    // Handle multiple errors
}
```

### 2. Struct Validation

```go
type User struct {
    Username string
    Password string
    Email    string
}

func (u *User) Validate() error {
    return ValidateStruct(user, "User cannot be nil",
        rule.Field(&user.Username,
            rule.Length(3, 20),
            rule.String().Errf("Username is required"),
        ),
        rule.Field(&user.Password,
            rule.Length(8, 50),
            rule.SpecialChars(true).Errf("Password must contain special characters"),
        ),
        rule.Field(&user.Email,
            rule.Email(),
            rule.String().Errf("Invalid email"),
        ),
    )
}


```

### 3. Rule Composition

```go
// Combine rules with AND
rule := rule.And(
    rule.Length(3, 10),
    rule.String().Errf("Invalid string"),
)

// Combine rules with OR
rule := rule.Or(
    rule.Email(),
    rule.URL(),
)
```

### 4. Custom Rules

```go
// Create a custom rule
type CustomRule struct {
    err error
}

func (r *CustomRule) Validate(value string) error {
    if value == "" {
        return r.err
    }
    return nil
}

// Use the custom rule
rule := &CustomRule{err: errors.New("custom error")}
err := Validate("", rule)
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Testing

Run the test suite:

```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -v -cover ./...
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## License

[License details to be added]

## Version History

[Version history to be added] 


## Why Arbiter?

Arbiter, meaning "judge" or "arbitrator", is designed to be your code's validation authority. We chose this name because:

1. **Comprehensive Validation**: Like a judge who evaluates evidence, Arbiter thoroughly validates your data against defined rules.
2. **Type Safety**: Built with Go's type system in mind, providing compile-time type checking and generic support.
3. **Flexible Rules**: Supports both built-in and custom validation rules, allowing you to define exactly what constitutes valid data.
4. **Performance**: Optimized for high-performance validation with minimal allocations.
5. **Developer Experience**: Intuitive API design with chainable rules and clear error messages.

### Advantages Over Tag-based Validation

While tag-based validation libraries are popular, Arbiter takes a different approach with several key advantages:

1. **Type Safety and IDE Support**
   - Tag-based: Validation rules defined in string tags, lacking type checking and IDE support
   - Arbiter: Rules defined using Go code, providing full type safety and IDE features (autocomplete, refactoring, etc.)

2. **Runtime Performance**
   - Tag-based: Requires reflection to parse tags and validate at runtime
   - Arbiter: Direct function calls with minimal reflection, resulting in better performance

3. **Flexibility and Maintainability**
   - Tag-based: Complex validation rules become hard to read and maintain in tags
   - Arbiter: Rules are regular Go code, allowing for better organization and reuse

4. **Debugging and Testing**
   - Tag-based: Errors in validation rules are only discovered at runtime
   - Arbiter: Validation logic can be unit tested and debugged like normal code

5. **Custom Rules**
   - Tag-based: Custom validators often require registration and reflection
   - Arbiter: Custom rules are just Go interfaces, simple to implement and use

6. **Conditional Validation**
   - Tag-based: Complex conditions are difficult to express in tags
   - Arbiter: Full power of Go for conditional logic

When building modern applications, data validation is crucial. Arbiter provides a robust, type-safe, and extensible solution that grows with your application.
