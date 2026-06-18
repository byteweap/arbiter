# Security Policy

## Supported Versions

Security fixes are provided for the latest released major version.

| Version | Supported |
| ------- | --------- |
| 1.x     | Yes       |
| < 1.0   | No        |

## Reporting a Vulnerability

Please do not report security vulnerabilities through public GitHub issues.

To report a vulnerability, contact the maintainer privately with:

- A clear description of the issue.
- A minimal reproduction or proof of concept, if available.
- The affected version or commit.
- Any known impact and mitigation.

If this repository does not yet publish a dedicated security contact, use GitHub's private vulnerability reporting feature when available, or contact the repository owner through their public GitHub profile.

You should receive an initial response within 7 days. Confirmed issues will be triaged based on severity, fixed in a private branch when appropriate, and disclosed after a patched release is available.

## Security-Oriented Validation Rules

Arbiter includes rules such as `XSS` and `SQLInjection` that detect common suspicious input patterns. These rules are convenience validators only. They are not complete security controls and must not replace:

- Context-aware output escaping.
- HTML/content sanitization.
- Content Security Policy.
- Parameterized SQL queries.
- Least-privilege database permissions.
- Framework and infrastructure security controls.

Applications should treat these rules as one layer in a broader defense strategy.
