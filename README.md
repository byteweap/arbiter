<p align="center">
  <img src="arbiter.png" alt="Arbiter Logo" width="200"/>
</p>

# Arbiter

[![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go)](https://go.dev)
[![CI](https://github.com/byteweap/arbiter/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/byteweap/arbiter/actions)
[![codecov](https://codecov.io/gh/byteweap/arbiter/branch/master/graph/badge.svg)](https://codecov.io/gh/byteweap/arbiter)
[![Go Report Card](https://goreportcard.com/badge/github.com/byteweap/arbiter)](https://goreportcard.com/report/github.com/byteweap/arbiter)
[![Release](https://img.shields.io/github/v/release/byteweap/arbiter)](https://github.com/byteweap/arbiter/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

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
- MIT licensed and suitable for commercial use

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
        arbiter.Field(&person.Name,
            rule.Len[string](2, 50).Errf("name must be 2-50 characters"),
            rule.Required[string]().Errf("name is required"),
        ),
        arbiter.Field(&person.Age,
            rule.Min(0),
            rule.Max(120),
        ),
        arbiter.Field(&person.Email,
            rule.Required[string]().Errf("required"),
            rule.IsEmail().Errf("invalid email"),
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
err := arbiter.Validate("hello",
    rule.Len[string](3, 10).Errf("Invalid string"),
)

// ValidateAll collects all validation errors
errs := arbiter.ValidateAll("hello",
    rule.Required[string]().Errf("required"),
    rule.Len[string](3, 10).Errf("invalid string"),
)

// ValidateStruct validates a struct and its fields
err := arbiter.ValidateStruct(person, "Person cannot be nil",
    arbiter.Field(&person.Name, rule.Required[string]()),
    arbiter.Field(&person.Age, rule.Min(0)),
)
```

### 2. Field Validator

For validating struct fields:

```go
// Create a field validation rule
nameRule := arbiter.Field(&person.Name,
    rule.Len[string](2, 50).Errf("Name must be 2-50 characters"),
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
- ...
  
#### Time Rules
- `Before`: Earlier than specified time
- `After`: Later than specified time
- `Between`: Time range validation
- ...
  
#### File Rules
- `Size`: File size validation
- `Extension`: File extension validation
- ...
  
#### Network Rules
- `IP`: IP address validation
- `URL`: URL validation
- ...

### Regex Rules
- `IsEmail`: email validation
- `IsPhone`: phone number validation
- `Regex`: custom regex validation
- ...

### Security-Oriented Rules
- `PasswordStrength`: password strength validation
- `PasswordComplex`: password complexity validation
- `XSS`: heuristic checks for common XSS patterns
- `SQLInjection`: heuristic checks for common SQL injection patterns

`XSS` and `SQLInjection` are input validation helpers, not security boundaries. They do not replace HTML escaping, content sanitization, Content Security Policy, parameterized SQL queries, least-privilege database access, or other application security controls.
  
## Best Practices

### 1. Error Handling

```go
// Stop at first error
err := arbiter.Validate(value,
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
errs := arbiter.ValidateAll(value,
    rule1,
    rule2,
    rule3,
)
if len(errs) > 0 {
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
    return arbiter.ValidateStruct(u, "User cannot be nil",
        arbiter.Field(&u.Username,
            rule.Len[string](3, 20).Errf("Username must be 3-20 characters"),
            rule.Required[string]().Errf("Username is required"),
        ),
        arbiter.Field(&u.Password,
            rule.Len[string](8, 50),
            rule.SpecialChars(true).Errf("Password must contain special characters"),
        ),
        arbiter.Field(&u.Email,
            rule.IsEmail().Errf("Invalid email"),
        ),
    )
}
```

### 3. Rule Composition

```go
// Combine rules with AND
stringRule := rule.And[string](
    rule.Required[string]().Errf("required"),
    rule.Len[string](3, 10).Errf("Invalid string"),
)

// Combine rules with OR
contactRule := rule.Or[string](
    rule.IsEmail(),
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
customRule := &CustomRule{err: errors.New("custom error")}
err := arbiter.Validate("", customRule)
```

## Contributing

Contributions are welcome. Please read [CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request.

For security issues, do not open a public issue. Please follow [SECURITY.md](SECURITY.md).

## Stability and Compatibility

Arbiter follows Semantic Versioning for public releases:

- Patch releases fix bugs and documentation without breaking public APIs.
- Minor releases may add new rules or capabilities while keeping existing public APIs compatible.
- Major releases may include breaking changes and will document migration notes in [CHANGELOG.md](CHANGELOG.md).

The public compatibility surface includes exported packages, types, functions, methods, errors, and documented behavior in `github.com/byteweap/arbiter` and `github.com/byteweap/arbiter/rule`.

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

Distributed under the MIT License. See `LICENSE` for more information.

## Version History

See [CHANGELOG.md](CHANGELOG.md) for release history.

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
