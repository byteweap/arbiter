# Changelog

All notable changes to this project are documented in this file.

This project follows Semantic Versioning for public releases.

## Unreleased

### Added

- Added security policy and vulnerability reporting guidance.
- Added contribution guidelines for development, testing, and pull requests.
- Documented semantic versioning and public API compatibility expectations.

### Changed

- Updated README examples to match the current public API.
- Clarified that XSS and SQL injection rules are heuristic validators, not complete security controls.
- Updated GoDoc examples that referenced older rule names and field APIs.

## v1.0.0 - 2026-05-26

### Added

- Initial stable release of Arbiter.
- Generic value validation through `Validate` and `ValidateAll`.
- Struct field validation through `ValidateStruct`, `Field`, `NestedField`, and `SliceField`.
- Built-in validation rules for required values, nil checks, length, ranges, numeric constraints, strings, regexes, network values, files, time values, and security-oriented checks.
- CI, lint, test, coverage, and release workflows.

### License

- Distributed under the MIT License.
