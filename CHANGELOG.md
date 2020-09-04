# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Clean json marshalling for internal configurations.
- Improve test coverage and go docs.

## [v0.2.0] - 2020-09-02
### Added
- Mark as first official release.
- QR support to export key URI.
- Support to format OTP configuration as authenticator key URI.
- Coverage report with GitHub Actions and Coveralls. 

### Changed
- Improve usage instructions in the README.
- Generate random keys without padding.
- Documentation references pkg.go.dev instead of godoc.
- Internal structure refactor.

### Fixed
- Handle key padding correctly when generating OTPs.

## [v0.1.0] - 2020-08-28
### Added
- Generation and validation of tokens.
- Simple standalone usage examples.
- GitHub Actions configuration for continuous testing.
- HMAC-Based and Time-Based OTP types.
- OTP calculation algorithm [full spec](https://tools.ietf.org/html/rfc4226#section-5).
- List of supported algorithms.
- Basic go module configuration, README, badges.

[Unreleased]: https://github.com/jltorresm/otpgo/compare/v0.2.0...HEAD
[v0.2.0]: https://github.com/jltorresm/otpgo/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/jltorresm/otpgo/compare/5130d24...v0.1.0