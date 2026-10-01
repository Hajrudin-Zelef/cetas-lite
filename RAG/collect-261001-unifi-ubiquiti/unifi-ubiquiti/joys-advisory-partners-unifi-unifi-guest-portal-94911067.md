---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/joys-advisory-partners-unifi-unifi-guest-portal-94911067
title: "joys-advisory-partners-unifi-unifi-guest-portal-94911067"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/joys-advisory-partners-unifi-unifi-guest-portal-94911067.md
source_anchor: ""
source_lines: [1, 30]
sha256: 51282b6c7c8e8a7b2ca1f3718b34627135ee918af9614d3b53e2a0713318da7a
---

# joys-advisory-partners-unifi-unifi-guest-portal-94911067

A captive portal for UniFi guest WiFi networks using Authentik OIDC authentication.
This portal replaces the built-in UniFi captive portal with a custom web app that:
- Authenticates guests via Authentik OIDC
- Supports multiple UniFi sites (JFMT, JFHR)
- Authorizes guest MAC addresses via the UniFi API
- Supports per-user session duration overrides via Authentik user attributes
Guest connects to WiFi
  → UniFi redirects to portal (/portal?site=jfmt&mac=xx:xx&...)
  → Portal initiates Authentik OIDC login
  → Guest authenticates (password + optional MFA)
  → Portal calls UniFi API to authorize guest MAC
  → Guest redirected to original URL
cp .env.example .env
Edit .env with your Authentik and UniFi credentials.
Place the following in app/static/:
- pup.jpg — the Shiba Inu photo
- jfmt-pdx-logo.svg — the JFMT-PDX logo
Create an OIDC provider in Authentik for the portal:
- Redirect URI: https://portal.jfmt-pdx.net/callback
- Note the Client ID and Client Secret for your .env
In UniFi Hotspot Portal:
- Enable External Portal Server
- Set URL to https://portal.jfmt-pdx.net/portal
docker compose up -d
By default guests get 24 hours (1440 minutes). To override for a specific user, set the following attribute on their Authentik user profile:
guest_wifi_duration_minutes: 480
- unifi-utils-python — UniFi API client
- FastAPI
- Authentik
Apache License 2.0
