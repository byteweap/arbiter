// Package arbiter provides validation functionality for various data types.
// This file contains types and functions for validating struct fields.
package arbiter

import "github.com/byteweap/arbiter/rule"

// IFieldRule is an interface that defines the contract for field validation rules.
// Any type that implements this interface can be used with ValidateStruct.
//
// Example:
//
//	// Custom field rule implementation
//	type CustomFieldRule struct {
//	    field *string
//	    rules []rule.Rule[string]
//	}
//
//	func (f *CustomFieldRule) validate() error {
//	    for _, rule := range f.rules {
//	        if err := rule.Validate(*f.field); err != nil {
//	            return err
//	        }
//	    }
//	    return nil
//	}
type IFieldRule interface {
	validate() error
}

// FieldRule is a generic type that implements IFieldRule for validating a field
// of any type with a set of validation rules.
//
// Example:
//
//	type Person struct {
//	    Name  string
//	    Age   int
//	    Email string
//	}
//
//	person := &Person{
//	    Name:  "John",
//	    Age:   30,
//	    Email: "john@example.com",
//	}
//
//	// Create field rules
//	nameRule := Field(&person.Name,
//	    rule.Length(2, 50),
//	    rule.String().Errf("Name is required"),
//	)
//
//	ageRule := Field(&person.Age,
//	    rule.Min(0),
//	    rule.Max(120),
//	)
//
//	emailRule := Field(&person.Email,
//	    rule.Email(),
//	    rule.String().Errf("Invalid email"),
//	)
//
//	// Use with ValidateStruct
//	err := ValidateStruct(person, "Person cannot be nil",
//	    nameRule, ageRule, emailRule,
//	)
type FieldRule[T any] struct {
	field *T
	rules []rule.Rule[T]
}

// Field creates a new field validation rule for a field of any type.
// The field parameter is a pointer to the field to validate.
// The rules parameter is a list of validation rules to apply to the field.
//
// Example:
//
//	type User struct {
//	    Username string
//	    Password string
//	}
//
//	user := &User{
//	    Username: "johndoe",
//	    Password: "secret123",
//	}
//
//	// Create field rules
//	usernameRule := Field(&user.Username,
//	    rule.Length(3, 20),
//	    rule.String().Errf("Username is required"),
//	)
//
//	passwordRule := Field(&user.Password,
//	    rule.Length(8, 50),
//	    rule.String().Errf("Password is required"),
//	)
//
//	// Use with ValidateStruct
//	err := ValidateStruct(user, "User cannot be nil",
//	    usernameRule, passwordRule,
//	)
func Field[T any](field *T, rules ...rule.Rule[T]) *FieldRule[T] {
	return &FieldRule[T]{field: field, rules: rules}
}

// validate applies all validation rules to the field.
// It returns nil if all rules pass, or the first error encountered.
//
// Example:
//
//	type Product struct {
//	    Name  string
//	    Price float64
//	}
//
//	product := &Product{
//	    Name:  "Widget",
//	    Price: 19.99,
//	}
//
//	// Create field rules
//	nameRule := Field(&product.Name,
//	    rule.Length(1, 100),
//	    rule.String().Errf("Name is required"),
//	)
//
//	priceRule := Field(&product.Price,
//	    rule.Min(0.0),
//	    rule.Precision(2),
//	)
//
//	// Validate fields directly
//	err := nameRule.validate()  // returns nil
//	err = priceRule.validate()  // returns nil
//
//	// Or use with ValidateStruct
//	err = ValidateStruct(product, "Product cannot be nil",
//	    nameRule, priceRule,
//	)
func (f *FieldRule[T]) validate() error {
	for _, r := range f.rules {
		if err := r.Validate(*f.field); err != nil {
			return err
		}
	}
	return nil
}

// NestedFieldRule validates a nested struct field by applying a list of sub-field rules.
type NestedFieldRule struct {
	fields []IFieldRule
}

// NestedField creates a validation rule for a nested struct field.
// The field parameter is a pointer to the nested struct.
// The sub-fields parameter is a list of field rules to apply to the nested struct.
//
// Example:
//
//	type Address struct {
//	    City   string
//	    Street string
//	}
//	type User struct {
//	    Name    string
//	    Address Address
//	}
//
//	err := arbiter.ValidateStruct(user, "User cannot be nil",
//	    arbiter.Field(&user.Name, rule.Required[string]()),
//	    arbiter.NestedField(&user.Address,
//	        arbiter.Field(&user.Address.City, rule.Required[string]()),
//	        arbiter.Field(&user.Address.Street, rule.Len[string](1, 100)),
//	    ),
//	)
func NestedField(_ any, fields ...IFieldRule) *NestedFieldRule {
	return &NestedFieldRule{fields: fields}
}

// validate applies all sub-field rules to the nested struct.
// Returns nil if all rules pass, or the first error encountered.
func (n *NestedFieldRule) validate() error {
	for _, field := range n.fields {
		if err := field.validate(); err != nil {
			return err
		}
	}
	return nil
}

// SliceFieldRule validates each element in a slice by applying rules generated from a callback.
type SliceFieldRule[T any] struct {
	field *[]T
	fn    func(*T) IFieldRule
}

// SliceField creates a validation rule for a slice field.
// The field parameter is a pointer to the slice.
// The fn parameter is a callback that receives a pointer to each element and returns validation rules.
//
// Example:
//
//	type User struct {
//	    Tags []string
//	}
//
//	err := arbiter.ValidateStruct(user, "User cannot be nil",
//	    arbiter.SliceField(&user.Tags, func(tag *string) arbiter.IFieldRule {
//	        return arbiter.Field(tag, rule.Len[string](1, 20))
//	    }),
//	)
func SliceField[T any](field *[]T, fn func(*T) IFieldRule) *SliceFieldRule[T] {
	return &SliceFieldRule[T]{field: field, fn: fn}
}

// validate iterates over each element in the slice and applies the rules from the callback.
// Returns nil if all elements pass, or the first error encountered.
func (s *SliceFieldRule[T]) validate() error {
	if s.fn == nil || s.field == nil {
		return nil
	}
	for i := range *s.field {
		f := s.fn(&(*s.field)[i])
		if err := f.validate(); err != nil {
			return err
		}
	}
	return nil
}
