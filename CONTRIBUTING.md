# Contributing to Arbiter

Thanks for helping improve Arbiter. This project aims to stay small, predictable, and easy to adopt in production Go code.

## Development Setup

Use Go 1.23 or newer.

```bash
go test ./...
go vet ./...
golangci-lint run ./...
```

Before sending a pull request, please run the full local check set above.

## Pull Requests

Good pull requests are focused and easy to review:

- Keep behavior changes small and intentional.
- Add or update tests for user-visible behavior.
- Update README, GoDoc, or CHANGELOG entries when public behavior changes.
- Avoid unrelated refactors in feature or bug-fix pull requests.
- Preserve existing API compatibility unless the change is explicitly planned for a major release.

## Public API Compatibility

Arbiter follows Semantic Versioning:

- Patch releases should not break public APIs or documented behavior.
- Minor releases may add public APIs while preserving compatibility.
- Major releases may introduce breaking changes with migration notes.

The public compatibility surface includes exported packages, types, functions, methods, errors, and documented behavior in `github.com/byteweap/arbiter` and `github.com/byteweap/arbiter/rule`.

## Adding Validation Rules

When adding a rule:

- Implement the `rule.Rule[T]` interface.
- Provide a default exported error value when useful.
- Add `Errf` for custom error messages, matching existing rule style.
- Treat empty strings as valid for optional string rules unless the rule is specifically about required values.
- Add table-driven tests for valid inputs, invalid inputs, custom errors, and edge cases.
- Document security limitations if the rule detects security-sensitive input.

## Reporting Security Issues

Please follow [SECURITY.md](SECURITY.md) and do not open public issues for vulnerabilities.
