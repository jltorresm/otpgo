# otpgo :: WIP
HMAC-Based and Time-Based One-Time Password (HOTP and TOTP) library for Go. 
Implements [RFC 4226][rfc4226] and [RFC 6238][rfc6238].

[![License][licenseBadge]][licenseLink]
[![Go Report Card][goReportBadge]][goReportLink]
[![Test Status][testStatusBadge]][testStatusLink]
[![GoDoc][goDocBadge]][goDocLink]

# Contents
- [Supported Operations](#supported-operations)
- [Planned Functionality](#planned-functionality)
- [Reading Material](#reading-material)
- [Usage](#usage)
    - [Generating Codes](#generating-codes)
    - [Verifying Codes](#verifying-codes)
    - [Registering with Authenticator App](#registering-with-authenticator-apps)
- [Defaults](#defaults)
    - [HOTP Parameters](#hotp-parameters)
    - [TOTP Parameters](#totp-parameters)

## Supported Operations
- Generate HOTP and TOTP codes.

## Planned Functionality
- Verify HOTP an TOTP codes.
- Generate QR code image to register secrets in authenticator apps.

## Reading Material
- [HOTP: An HMAC-Based One-Time Password Algorithm][rfc4226]
- [TOTP: Time-Based One-Time Password Algorithm][rfc6238]
- [Google Authenticator Key URI Format][googleURI]
- [Browser Authenticator Demo][debugger]

## Usage

### Generating Codes
The simplest way to generate codes is to create the HOTP/TOTP struct and call 
`Generate()`

```
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
h := otpgo.HOTP{}
token, _ := h.Generate()
```

Each type allows customization. For HMAC-Based tokens you can specify:
- Key: Secret string, base32 encoded
- Counter: Unsigned int
- Algorithm: One of `HmacSHA1`, `HmacSHA256` or `HmacSHA512`
- Length: `Length1` up to `Length8`

### Verifying Codes
TBD

### Registering With Authenticator Apps
TBD

## Defaults
If caller doesn't provide a custom configuration when generating OTPs. The 
library will ensure the following default values (any empty value will be 
filled).

### HOTP Parameters
|Parameter|Default Value                      |
|:-------:|:---------------------------------:|
|Hash     |`SHA1`                             |
|Length   |`6`                                |
|Key      |`64` random bytes `base32` encoded |

### TOTP Parameters
|Parameter|Default Value                      |
|:-------:|:---------------------------------:|
|Period   |`30` seconds                       |
|Hash     |`SHA1`                             |
|Length   |`6`                                |
|Key      |`64` random bytes `base32` encoded |

[licenseBadge]: https://img.shields.io/github/license/jltorresm/otpgo
[licenseLink]: https://github.com/jltorresm/otpgo/blob/main/LICENSE
[goReportBadge]: https://goreportcard.com/badge/github.com/jltorresm/otpgo
[goReportLink]: https://goreportcard.com/report/github.com/jltorresm/otpgo
[testStatusBadge]: https://img.shields.io/github/workflow/status/jltorresm/otpgo/test?label=test&logo=github
[testStatusLink]: https://github.com/jltorresm/otpgo/actions?query=workflow%3Atest
[goDocBadge]: https://godoc.org/github.com/jltorresm/otpgo?status.svg
[goDocLink]: https://godoc.org/github.com/jltorresm/otpgo

[latest]: https://github.com/kilico-travel/kilico-api/releases/latest
[rfc4226]: https://tools.ietf.org/html/rfc4226
[rfc6238]: https://tools.ietf.org/html/rfc6238
[googleURI]: https://github.com/google/google-authenticator/wiki/Key-Uri-Format
[debugger]: https://rootprojects.org/authenticator/
