---
id: collect-261001-general-networking/general-networking/manual-two-factor-html-45abea1a
title: "manual-two-factor-html-45abea1a"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-two-factor-html-45abea1a.md
source_anchor: ""
source_lines: [1, 17]
sha256: 22cc029a11ab51a4529a133b945ea5a2d279008baf8f69135ad92795b8b3294e
---

# manual-two-factor-html-45abea1a

Two-factor authentication
Two-factor authentication also known as 2FA or 2-Step Verification is an authentication method that requires two components, such as a pin/password + a token.
OPNsense (version >=16.1.14) offers support for Two-factor authentication throughout the entire system, with one exception being console/ssh access.
Supported services are:
- OPNsense Graphical User Interface
- Captive Portal
- Virtual Private Networking - OpenVPN & IPsec
- Caching Proxy
New authentication servers can be added via , which supports both local users and users synchronised via ldap.
Time-based One-time Password
TOTP is an algorithm that computes a one-time password from a shared secret key and the current time. OPNsense supports RFC 6238.
Google Authenticator
OPNsense fully supports the use of Google’s Authenticator application. This application can generate tokens on Android, iOS and BlackBerry OS. The usage of this application is free and very simple to setup using OPNsense.
Other TOTP tokens
The 2FA feature can be used with any time based one-time password token, although it may be necessary to convert the tokens seed to the used format (base32).
Configuration & Setup
To setup see: Configure 2FA TOTP & Google Authenticator.
