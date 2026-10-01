---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/filipowm-terraform-provider-unifi-blob-head-docs-resources-setting-guest-access-b5b16825-3
title: "Configure guest access settings for your UniFi network"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Stripe"]
dates: []
keywords: ["sandbox"]
source: docs/RAG/collect-261001-unifi-ubiquiti/filipowm-terraform-provider-unifi-blob-head-docs-resources-setting-guest-access--b5b16825.md
source_anchor: ""
source_lines: [224, 251]
sha256: 7338dee5148559c911c51ac2f8ecfef0d69464ffe1a33244f4d6152cf2a3fa01
---

# Configure guest access settings for your UniFi network

- welcome_text (String) Welcome text displayed on the portal.
- welcome_text_enabled (Boolean) Enable welcome text display.
- welcome_text_position (String) Position of the welcome text. Valid values are:under_logo ,above_boxes .
Required:
- agreement_id (String, Sensitive) QuickPay agreement ID.
- api_key (String, Sensitive) QuickPay API key.
- merchant_id (String, Sensitive) QuickPay merchant ID.
Optional:
- use_sandbox (Boolean) Enable sandbox mode for QuickPay payments.
Required:
- auth_type (String) RADIUS authentication type. Valid values are:chap ,mschapv2 .
- profile_id (String) ID of the RADIUS profile to use.
Optional:
- disconnect_enabled (Boolean) Enable RADIUS disconnect messages.
- disconnect_port (Number) Port for RADIUS disconnect messages.
Required:
- url (String) URL to redirect to after authentication. Must be a valid URL.
Optional:
- to_https (Boolean) Redirect HTTP requests to HTTPS.
- use_https (Boolean) Use HTTPS for the redirect URL.
Required:
- api_key (String, Sensitive) Stripe API key.
Required:
- app_id (String) WeChat App ID for social authentication.
- app_secret (String, Sensitive) WeChat App secret.
- secret_key (String, Sensitive) WeChat secret key.
Optional:
- shop_id (String) WeChat Shop ID for payments.
