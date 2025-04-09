package rule

// Integer represents all integer types that can be used in validation rules.
// This includes both signed and unsigned integers of various sizes.

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Ordered defines numeric types that can be compared using <, <=, >, and >= operators.
// This includes integers and floating-point numbers.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// RequiredType defines types that can be checked for required/optional status.
// This includes both value types and their pointer variants.
type RequiredType interface {
	~string | ~*string |
		~int | ~*int |
		~int8 | ~*int8 |
		~int16 | ~*int16 |
		~int32 | ~*int32 |
		~int64 | ~*int64 |
		~uint | ~*uint |
		~uint8 | ~*uint8 |
		~uint16 | ~*uint16 |
		~uint32 | ~*uint32 |
		~uint64 | ~*uint64 |
		~float32 | ~*float32 |
		~float64 | ~*float64
}

// InType defines types that can be used in set membership validation.
// This includes all basic types and their pointer variants.
type InType interface {
	~int | ~*int |
		~int8 | ~*int8 |
		~int16 | ~*int16 |
		~int32 | ~*int32 |
		~int64 | ~*int64 |
		~uint | ~*uint |
		~uint8 | ~*uint8 |
		~uint16 | ~*uint16 |
		~uint32 | ~*uint32 |
		~uint64 | ~*uint64 |
		~float32 | ~*float32 |
		~float64 | ~*float64 |
		~string | ~*string |
		~bool | ~*bool
}

// Zeroable defines types that can be checked for zero values.
// This includes all basic types and any interface type.
type Zeroable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string | ~bool |
		~complex64 | ~complex128 |
		any
}
