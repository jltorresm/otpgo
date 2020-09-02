# otpgo :: WIP
HMAC-Based and Time-Based One-Time Password (HOTP and TOTP) library for Go. 
Implements [RFC 4226][rfc4226] and [RFC 6238][rfc6238].

[![License][licenseBadge]][licenseLink]
[![Go Report Card][goReportBadge]][goReportLink]
[![Test Status][testStatusBadge]][testStatusLink]
[![Coverage Status][coverageBadge]][coverageLink]
[![PkgGoDev][pkgGoDevBadge]][pkgGoDevLink]
[![Latest Release][releaseBadge]][releaseLink]

# Contents
- [Supported Operations](#supported-operations)
- [Planned Functionality](#planned-functionality)
- [Reading Material](#reading-material)
- [Usage](#usage)
    - [Generating Codes](#generating-codes)
    - [Verifying Codes](#verifying-codes)
    - [Registering with Authenticator App](#registering-with-authenticator-apps)
        - [QR Code](#qr-code)
        - [Manual Registration](#manual-registration)
- [Defaults](#defaults)
    - [HOTP Parameters](#hotp-parameters)
    - [TOTP Parameters](#totp-parameters)

## Supported Operations
- Generate HOTP and TOTP codes.
- Verify HOTP an TOTP codes.
- Export OTP config as a [Google Authenticator URI][googleURI].
- Export OTP config as a QR code image (used to register secrets in authenticator apps).

## Planned Functionality
- Export OTP config as a JSON.

## Reading Material
- [HOTP: An HMAC-Based One-Time Password Algorithm][rfc4226]
- [TOTP: Time-Based One-Time Password Algorithm][rfc6238]
- [Google Authenticator Key URI Format][googleURI]
- [Browser Authenticator Demo][debugger]

## Usage

### Generating Codes
The simplest way to generate codes is to create the HOTP/TOTP struct and call 
`Generate()`

```go
// 
// HMAC-Based
//

// Will use all default values, counter starts in 0
h := otpgo.HOTP{}
token, _ := h.Generate()

// Increment counter and generate next code
h.Counter++
token2, _ := h.Generate()

//
// Time-Based
//

// Will use all default values
t := otpgo.TOTP{}
token, _ := t.Generate()
```

Each type allows customization. For **HMAC-Based** tokens you can specify:
- **Key**: Secret string, base32 encoded
- **Counter**: Unsigned int
- **Leeway**: Unsigned int
- **Algorithm**: One of `HmacSHA1`, `HmacSHA256` or `HmacSHA512`
- **Length**: `Length1` up to `Length8`

For **Time-Based** tokens you can specify:
- **Key**: Secret string, base32 encoded
- **Period**: Integer, period length in seconds
- **Delay**: Integer, acceptable number of steps for validation
- **Algorithm**: One of `HmacSHA1`, `HmacSHA256` or `HmacSHA512`
- **Length**: `Length1` up to `Length8`

### Verifying Codes
Once you receive a token from the user you can verify it by specifying the 
expected parameters and calling `Validate(token string)`.

```go
// 
// HMAC-Based
//
h := otpgo.HOTP{
    Key: "my-secret-key",
    Counter: 123, // The expected counter
}
ok, _ := h.Validate("the-token")

//
// Time-Based
//
t := otpgo.TOTP{
    Key: "my-secret-key",
}
ok, _ = t.Validate("the-token")
```

When calling `HOTP.Validate()` note that the internal counter will be increased
if validation is successful, so that the next valid token will correspond to the
increased counter.

Both `HOTP` and `TOTP` will accept tokens that match the exact 
`Counter`/`Timestamp` or a token within the specified `Leeway`/`Delay`.

### Registering With Authenticator Apps
Most authenticator apps will give the user 2 options to register a new account:
scan a QR code which contains all config and secrets for the OTP generation, or 
manually enter the secret key and additional info (such as username and issuer).
The former being the preferred way because of the ease of use and the avoidance
of human error.

#### QR Code
TODO

#### Manual registration
TODO

## Defaults
If caller doesn't provide a custom configuration when generating OTPs. The 
library will ensure the following default values (any empty value will be 
filled).

### HOTP Parameters
|Parameter        |Default Value                      |
|:---------------:|:---------------------------------:|
|Leeway           |`1` counter down & up              |
|Hash / Algorithm |`SHA1`                             |
|Length           |`6`                                |
|Key              |`64` random bytes `base32` encoded |

### TOTP Parameters
|Parameter        |Default Value                      |
|:---------------:|:---------------------------------:|
|Period           |`30` seconds                       |
|Delay            |`1` period under & over            |
|Hash / Algorithm |`SHA1`                             |
|Length           |`6`                                |
|Key              |`64` random bytes `base32` encoded |

[licenseBadge]: https://img.shields.io/github/license/jltorresm/otpgo
[licenseLink]: https://github.com/jltorresm/otpgo/blob/main/LICENSE
[goReportBadge]: https://goreportcard.com/badge/github.com/jltorresm/otpgo
[goReportLink]: https://goreportcard.com/report/github.com/jltorresm/otpgo
[testStatusBadge]: https://img.shields.io/github/workflow/status/jltorresm/otpgo/test?label=test&logo=github
[testStatusLink]: https://github.com/jltorresm/otpgo/actions?query=workflow%3Atest
[pkgGoDevBadge]: https://pkg.go.dev/badge/github.com/jltorresm/otpgo
[pkgGoDevLink]: https://pkg.go.dev/github.com/jltorresm/otpgo
[releaseBadge]: https://img.shields.io/github/v/release/jltorresm/otpgo?include_prereleases
[releaseLink]: https://github.com/jltorresm/otpgo/releases/latest
[coverageBadge]: https://coveralls.io/repos/github/jltorresm/otpgo/badge.svg?branch=main
[coverageLink]: https://coveralls.io/github/jltorresm/otpgo?branch=main

[latest]: https://github.com/kilico-travel/kilico-api/releases/latest
[rfc4226]: https://tools.ietf.org/html/rfc4226
[rfc6238]: https://tools.ietf.org/html/rfc6238
[googleURI]: https://github.com/google/google-authenticator/wiki/Key-Uri-Format
[debugger]: https://rootprojects.org/authenticator/
