---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/filipowm-terraform-provider-unifi-blob-head-docs-resources-setting-guest-access-b5b16825-2
title: "Configure guest access settings for your UniFi network"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google", "Stripe"]
dates: []
keywords: ["sandbox"]
source: docs/RAG/collect-261001-unifi-ubiquiti/filipowm-terraform-provider-unifi-blob-head-docs-resources-setting-guest-access--b5b16825.md
source_anchor: ""
source_lines: [94, 223]
sha256: ec0593143811d2e2092fc6dd0740159d94b6b66ed997185e5a4ccebb1f6d0c1d
---

    # Languages supported
    languages = ["PL"]
  }
}
- allowed_subnet (String) Subnet allowed for guest access.
- auth (String) Authentication method for guest access. Valid values are:
- none - No authentication required
- hotspot - Password authentication
- facebook_wifi - Facebook auth entication
- custom - Custom authentication
For password authentication, set auth to hotspot and password_enabled to true.
For voucher authentication, set auth to hotspot and voucher_enabled to true.
For payment authentication, set auth to hotspot and payment_enabled to true.
- auth_url (String) URL for authentication. Must be a valid URL including the protocol.
- authorize (Attributes) Authorize.net payment settings. (see below for nested schema)
- custom_ip (String) Custom IP address. Must be a valid IPv4 address (e.g.,192.168.1.1 ).
- ec_enabled (Boolean) Enable enterprise controller functionality.
- expire (Number) Expiration time for guest access.
- expire_number (Number) Number value for the expiration time.
- expire_unit (Number) Unit for the expiration time. Valid values are:
- 1 - Minute
- 60 - Hour
- 1440 - Day
- 10080 - Week
- facebook (Attributes) Facebook authentication settings. (see below for nested schema)
- facebook_wifi (Attributes) Facebook WiFi authentication settings. (see below for nested schema)
- google (Attributes) Google authentication settings. (see below for nested schema)
- ippay (Attributes) IPpay Payments settings. (see below for nested schema)
- merchant_warrior (Attributes) MerchantWarrior payment settings. (see below for nested schema)
- password (String, Sensitive) Password for guest access.
- payment_gateway (String) Payment gateway. Valid values are:
- paypal - PayPal
- stripe - Stripe
- authorize - Authorize.net
- quickpay - QuickPay
- merchantwarrior - Merchant Warrior
- ippay - IP Payments
- paypal (Attributes) PayPal payment settings. (see below for nested schema)
- portal_customization (Attributes) Portal customization settings. (see below for nested schema)
- portal_enabled (Boolean) Enable the guest portal.
- portal_hostname (String) Hostname to use for the captive portal.
- portal_use_hostname (Boolean) Use a custom hostname for the portal.
- quickpay (Attributes) QuickPay payment settings. (see below for nested schema)
- radius (Attributes) RADIUS authentication settings. (see below for nested schema)
- redirect (Attributes) Redirect after authentication settings. (see below for nested schema)
- restricted_dns_servers (List of String) List of restricted DNS servers for guest networks. Each value must be a valid IPv4 address.
- restricted_subnet (String) Subnet for restricted guest access.
- site (String) The name of the UniFi site where this resource should be applied. If not specified, the default site will be used.
- stripe (Attributes) Stripe payment settings. (see below for nested schema)
- template_engine (String) Template engine for the portal. Valid values are:jsp ,angular .
- voucher_customized (Boolean) Whether vouchers are customized.
- voucher_enabled (Boolean) Enable voucher-based authentication for guest access.
- wechat (Attributes) WeChat authentication settings. (see below for nested schema)
- facebook_enabled (Boolean) Whether Facebook authentication for guest access is enabled.
- google_enabled (Boolean) Whether Google authentication for guest access is enabled.
- id (String) The unique identifier of this resource.
- password_enabled (Boolean) Enable password authentication for guest access.
- payment_enabled (Boolean) Enable payment for guest access.
- radius_enabled (Boolean) Whether RADIUS authentication for guest access is enabled.
- redirect_enabled (Boolean) Whether redirect after authentication is enabled.
- restricted_dns_enabled (Boolean) Whether restricted DNS servers for guest networks are enabled.
- wechat_enabled (Boolean) Whether WeChat authentication for guest access is enabled.
Required:
- login_id (String) Authorize.net login ID for authentication.
- transaction_key (String) Authorize.net transaction key for authentication.
Optional:
- use_sandbox (Boolean) Use sandbox mode for Authorize.net payments.
Required:
- app_id (String) Facebook application ID for authentication.
- app_secret (String, Sensitive) Facebook application secret for authentication.
Optional:
- scope_email (Boolean) Request email scope for Facebook authentication.
Required:
- gateway_id (String) Facebook WiFi gateway ID.
- gateway_name (String) Facebook WiFi gateway name.
- gateway_secret (String, Sensitive) Facebook WiFi gateway secret.
Optional:
- block_https (Boolean) Mode HTTPS for Facebook WiFi.
Required:
- client_id (String) Google client ID for authentication.
- client_secret (String) Google client secret for authentication.
Optional:
- domain (String) Restrict Google authentication to specific domain.
- scope_email (Boolean) Request email scope for Google authentication.
Required:
- terminal_id (String, Sensitive) Terminal ID for IP Payments.
Optional:
- use_sandbox (Boolean) Whether to use sandbox mode for IPPay payments.
Required:
- api_key (String, Sensitive) MerchantWarrior API key.
- api_passphrase (String, Sensitive) MerchantWarrior API passphrase.
- merchant_uuid (String, Sensitive) MerchantWarrior merchant UUID.
Optional:
- use_sandbox (Boolean) Whether to use sandbox mode for MerchantWarrior payments.
Required:
- password (String, Sensitive) PayPal password.
- signature (String, Sensitive) PayPal signature.
- username (String, Sensitive) PayPal username. Must be a valid email address.
Optional:
- use_sandbox (Boolean) Whether to use sandbox mode for PayPal payments.
Optional:
- authentication_text (String) Custom authentication text for the portal.
- bg_color (String) Background color for the custom portal. Must be a valid hex color code (e.g., #FFF or #FFFFFF).
- bg_image_file_id (String) ID of the background image portal file. File must exist in controller, useunifi_portal_file to manage it.
- bg_image_tile (Boolean) Tile the background image.
- bg_type (String) Type of portal background. Valid values are:
- color - Solid color background
- image - (not yet supported!) Custom image background
- gallery - Image from Unsplash gallery
- box_color (String) Color of the login box in the portal. Must be a valid hex color code (e.g., #FFF or #FFFFFF).
- box_link_color (String) Color of links in the login box. Must be a valid hex color code (e.g., #FFF or #FFFFFF).
- box_opacity (Number) Opacity of the login box (0-100).
- box_radius (Number) Border radius of the login box in pixels.
- box_text_color (String) Text color in the login box. Must be a valid hex color code (e.g., #FFF or #FFFFFF).
- button_color (String) Button color in the portal. Must be a valid hex color code (e.g., #FFF or #FFFFFF).
- button_text (String) Custom text for the login button.
- button_text_color (String) Button text color. Must be a valid hex color code (e.g., #FFF or #FFFFFF).
- customized (Boolean) Whether the portal is customized.
- languages (List of String) List of enabled languages for the portal.
- link_color (String) Color for links in the portal. Must be a valid hex color code (e.g., #FFF or #FFFFFF).
- logo_file_id (String) ID of the logo image portal file. File must exist in controller, useunifi_portal_file to manage it.
- logo_position (String) Position of the logo in the portal. Valid values are: left, center, right.
- logo_size (Number) Size of the logo in pixels.
- success_text (String) Text displayed after successful authentication.
- text_color (String) Main text color for the portal. Must be a valid hex color code (e.g., #FFF or #FFFFFF).
- title (String) Title of the portal page.
- tos (String) Terms of service text.
- tos_enabled (Boolean) Enable terms of service acceptance requirement.
- unsplash_author_name (String) Name of the Unsplash author for gallery background.
- unsplash_author_username (String) Username of the Unsplash author for gallery background.
