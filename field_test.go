// Package arbiter_test provides tests for the arbiter package.
// These tests verify the functionality of field validation rules.
package arbiter_test

import (
	"testing"

	"github.com/byteweap/arbiter"
	"github.com/byteweap/arbiter/rule"
)

type testUser struct {
	Username string
	Age      int
	Website  string
	Password string
	Chinese  string
}

func TestFieldValid(t *testing.T) {
	user := &testUser{
		Username: "johndoe",
		Age:      25,
		Website:  "example.com",
		Password: "StrongP@ssw0rd",
		Chinese:  "你好世界",
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username, rule.HalfWidthOnly()),
	)
	if err != nil {
		t.Errorf("Expected no error for valid username, got %v", err)
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Age, rule.Min[int](0), rule.Max[int](120)),
	)
	if err != nil {
		t.Errorf("Expected no error for valid age, got %v", err)
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Website, rule.Domain()),
	)
	if err != nil {
		t.Errorf("Expected no error for valid website, got %v", err)
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Password, rule.PasswordStrength()),
	)
	if err != nil {
		t.Errorf("Expected no error for valid password, got %v", err)
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Chinese, rule.ChineseOnly()),
	)
	if err != nil {
		t.Errorf("Expected no error for valid Chinese text, got %v", err)
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username, rule.HalfWidthOnly()),
		arbiter.Field(&user.Age, rule.Min[int](0), rule.Max[int](120)),
		arbiter.Field(&user.Website, rule.Domain()),
		arbiter.Field(&user.Password, rule.PasswordStrength()),
		arbiter.Field(&user.Chinese, rule.ChineseOnly()),
	)
	if err != nil {
		t.Errorf("Expected no error for valid user, got %v", err)
	}
}

func TestFieldInvalid(t *testing.T) {
	user := &testUser{
		Username: "Ｊｏｈｎ",
		Age:      -1,
		Website:  "invalid",
		Password: "weak",
		Chinese:  "Hello",
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username, rule.HalfWidthOnly()),
	)
	if err == nil {
		t.Error("Expected error for full-width username, got nil")
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Age, rule.Min[int](0), rule.Max[int](120)),
	)
	if err == nil {
		t.Error("Expected error for negative age, got nil")
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Website, rule.Domain()),
	)
	if err == nil {
		t.Error("Expected error for invalid website domain, got nil")
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Password, rule.PasswordStrength()),
	)
	if err == nil {
		t.Error("Expected error for weak password, got nil")
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Chinese, rule.ChineseOnly()),
	)
	if err == nil {
		t.Error("Expected error for non-Chinese text, got nil")
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username, rule.HalfWidthOnly()),
		arbiter.Field(&user.Age, rule.Min(0), rule.Max(120)),
		arbiter.Field(&user.Website, rule.Domain()),
		arbiter.Field(&user.Password, rule.PasswordStrength()),
		arbiter.Field(&user.Chinese, rule.ChineseOnly()),
	)
	if err == nil {
		t.Error("Expected error for invalid user, got nil")
	}
}

func TestFieldNilPointer(t *testing.T) {
	err := arbiter.ValidateStruct(nil, "User cannot be nil")
	if err == nil || err.Error() != "User cannot be nil" {
		t.Errorf("Expected error 'User cannot be nil', got %v", err)
	}
}

func TestFieldMultipleRules(t *testing.T) {
	user := &testUser{
		Username: "johndoe",
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username,
			rule.HalfWidthOnly(),
			rule.LowerCaseOnly(),
		),
	)
	if err != nil {
		t.Errorf("Expected no error for valid username with multiple rules, got %v", err)
	}

	user.Username = "JOHN123"
	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username,
			rule.HalfWidthOnly(),
			rule.LowerCaseOnly(),
		),
	)
	if err == nil {
		t.Error("Expected error for uppercase username, got nil")
	}
}

func TestFieldCustomError(t *testing.T) {
	user := &testUser{
		Age: -1,
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Age,
			rule.Min[int](0).Errf("Age cannot be negative"),
		),
	)
	if err == nil || err.Error() != "Age cannot be negative" {
		t.Errorf("Expected custom error message, got %v", err)
	}
}

// Nested struct tests

type testAddress struct {
	City   string
	Street string
	Zip    string
}

type testPersonWithAddress struct {
	Name    string
	Age     int
	Address testAddress
}

func TestNestedFieldValid(t *testing.T) {
	person := &testPersonWithAddress{
		Name: "John",
		Age:  30,
		Address: testAddress{
			City:   "Beijing",
			Street: "Main St",
			Zip:    "100000",
		},
	}

	err := arbiter.ValidateStruct(person, "Person cannot be nil",
		arbiter.Field(&person.Name, rule.Required[string]()),
		arbiter.NestedField(&person.Address,
			arbiter.Field(&person.Address.City, rule.Required[string]()),
			arbiter.Field(&person.Address.Street, rule.Len[string](1, 100)),
		),
	)
	if err != nil {
		t.Errorf("Expected no error for valid nested struct, got %v", err)
	}
}

func TestNestedFieldInvalid(t *testing.T) {
	person := &testPersonWithAddress{
		Name: "John",
		Age:  30,
		Address: testAddress{
			City:   "",
			Street: "Main St",
		},
	}

	err := arbiter.ValidateStruct(person, "Person cannot be nil",
		arbiter.Field(&person.Name, rule.Required[string]()),
		arbiter.NestedField(&person.Address,
			arbiter.Field(&person.Address.City, rule.Required[string]()),
			arbiter.Field(&person.Address.Street, rule.Len[string](1, 100)),
		),
	)
	if err == nil {
		t.Error("Expected error for empty city in nested struct, got nil")
	}
}

func TestNestedFieldMultipleLevels(t *testing.T) {
	type Inner struct {
		Value string
	}
	type Middle struct {
		Inner Inner
	}
	type Outer struct {
		Middle Middle
	}

	outer := &Outer{
		Middle: Middle{
			Inner: Inner{Value: ""},
		},
	}

	err := arbiter.ValidateStruct(outer, "Outer cannot be nil",
		arbiter.NestedField(&outer.Middle,
			arbiter.NestedField(&outer.Middle.Inner,
				arbiter.Field(&outer.Middle.Inner.Value, rule.Required[string]()),
			),
		),
	)
	if err == nil {
		t.Error("Expected error for empty value in deeply nested struct, got nil")
	}

	outer.Middle.Inner.Value = "hello"
	err = arbiter.ValidateStruct(outer, "Outer cannot be nil",
		arbiter.NestedField(&outer.Middle,
			arbiter.NestedField(&outer.Middle.Inner,
				arbiter.Field(&outer.Middle.Inner.Value, rule.Required[string]()),
			),
		),
	)
	if err != nil {
		t.Errorf("Expected no error for valid deeply nested struct, got %v", err)
	}
}

// Slice field tests

type testUserWithTags struct {
	Name string
	Tags []string
}

func TestSliceFieldValid(t *testing.T) {
	user := &testUserWithTags{
		Name: "John",
		Tags: []string{"go", "rust", "python"},
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Name, rule.Required[string]()),
		arbiter.SliceField(&user.Tags, func(tag *string) arbiter.IFieldRule {
			return arbiter.Field(tag, rule.Len[string](1, 20))
		}),
	)
	if err != nil {
		t.Errorf("Expected no error for valid slice, got %v", err)
	}
}

func TestSliceFieldInvalid(t *testing.T) {
	user := &testUserWithTags{
		Name: "John",
		Tags: []string{"go", "", "python"},
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Name, rule.Required[string]()),
		arbiter.SliceField(&user.Tags, func(tag *string) arbiter.IFieldRule {
			return arbiter.Field(tag, rule.Required[string]())
		}),
	)
	if err == nil {
		t.Error("Expected error for empty tag in slice, got nil")
	}
}

func TestSliceFieldEmpty(t *testing.T) {
	user := &testUserWithTags{
		Name: "John",
		Tags: []string{},
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.SliceField(&user.Tags, func(tag *string) arbiter.IFieldRule {
			return arbiter.Field(tag, rule.Required[string]())
		}),
	)
	if err != nil {
		t.Errorf("Expected no error for empty slice, got %v", err)
	}
}

func TestSliceFieldIntegers(t *testing.T) {
	type Config struct {
		Scores []int
	}

	cfg := &Config{Scores: []int{90, 85, 70, 110}}
	err := arbiter.ValidateStruct(cfg, "Config cannot be nil",
		arbiter.SliceField(&cfg.Scores, func(score *int) arbiter.IFieldRule {
			return arbiter.Field(score, rule.Max[int](100))
		}),
	)
	if err == nil {
		t.Error("Expected error for score > 100, got nil")
	}

	cfg.Scores = []int{90, 85, 70}
	err = arbiter.ValidateStruct(cfg, "Config cannot be nil",
		arbiter.SliceField(&cfg.Scores, func(score *int) arbiter.IFieldRule {
			return arbiter.Field(score, rule.Max[int](100))
		}),
	)
	if err != nil {
		t.Errorf("Expected no error for valid scores, got %v", err)
	}
}

func TestSliceFieldStructElements(t *testing.T) {
	type Item struct {
		Name  string
		Price float64
	}
	type Order struct {
		Items []Item
	}

	order := &Order{
		Items: []Item{
			{Name: "Widget", Price: 9.99},
			{Name: "", Price: 19.99},
		},
	}

	err := arbiter.ValidateStruct(order, "Order cannot be nil",
		arbiter.SliceField(&order.Items, func(item *Item) arbiter.IFieldRule {
			return arbiter.NestedField(item,
				arbiter.Field(&item.Name, rule.Required[string]()),
				arbiter.Field(&item.Price, rule.Min[float64](0)),
			)
		}),
	)
	if err == nil {
		t.Error("Expected error for empty item name, got nil")
	}

	order.Items[1].Name = "Gadget"
	err = arbiter.ValidateStruct(order, "Order cannot be nil",
		arbiter.SliceField(&order.Items, func(item *Item) arbiter.IFieldRule {
			return arbiter.NestedField(item,
				arbiter.Field(&item.Name, rule.Required[string]()),
				arbiter.Field(&item.Price, rule.Min[float64](0)),
			)
		}),
	)
	if err != nil {
		t.Errorf("Expected no error for valid order, got %v", err)
	}
}

func TestSliceFieldNilFn(t *testing.T) {
	user := &testUserWithTags{
		Name: "John",
		Tags: []string{"go"},
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.SliceField[string](&user.Tags, nil),
	)
	if err != nil {
		t.Errorf("Expected no error for nil callback, got %v", err)
	}
}
