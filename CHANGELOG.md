# Changelog

All notable changes to this project will be documented in this file.

## [1.0.0] - 2025-05-26

### Features

- Generic support for validating any data type
- 69 built-in validation rules (string, numeric, time, file, network, regex, etc.)
- `Validate` - quick validation, stops at first error
- `ValidateAll` - collects all validation errors
- `ValidateStruct` - struct field validation
- Chainable API with custom error messages (`Errf`)
- Conditional validation (`And`, `Or`)
- Dependency validation (`Dependency`, `MutualExclude`)
- Struct field validation with `Field` helper

### String Rules

- `StartWith` / `EndWith` - prefix and suffix validation
- `ChineseOnly` / `FullWidthOnly` / `HalfWidthOnly` - character type validation
- `UpperCaseOnly` / `LowerCaseOnly` - case validation
- `SpecialChars` / `Contains` / `NotContains` - substring validation

### Numeric Rules

- `Min` / `Max` / `Between` - range validation
- `Positive` / `Negative` / `Even` / `Odd` - number property validation
- `Precision` / `DivisibleBy` / `MultipleOf` - precision validation

### Time Rules

- `Before` / `After` / `TimeBetween` - time range validation
- `DateFormat` / `TimeFormat` / `DateTimeFormat` - format validation
- `Weekend` / `Workday` / `Holiday` - date type validation

### Network Rules

- `IP` / `IPv4` / `IPv6` - IP address validation
- `URL` / `Domain` / `Port` - network address validation
- `MACAddress` / `SubnetMask` - network config validation

### File Rules

- `FileSize` / `FileType` / `FileExtension` / `FileMimeType` - file attribute validation

### Regex Rules

- `IsEmail` / `IsPhone` / `IsIDCard` - common format validation
- `IsBankCard` / `IsPassport` / `IsTaxNumber` - document validation
- `IsSocialCredit` / `Regex` - custom regex validation

### Other Rules

- `Required` / `NonZero` / `Zero` / `Nil` - null value validation
- `Len` / `In` / `NotIn` - length and collection validation
