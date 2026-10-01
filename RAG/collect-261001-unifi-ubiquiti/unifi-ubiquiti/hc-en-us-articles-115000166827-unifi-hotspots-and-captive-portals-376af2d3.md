---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-115000166827-unifi-hotspots-and-captive-portals-376af2d3
title: "hc-en-us-articles-115000166827-unifi-hotspots-and-captive-portals-376af2d3"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Stripe"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-115000166827-unifi-hotspots-and-captive-portals-376af2d3.md
source_anchor: ""
source_lines: [1, 25]
sha256: e2914dbd5b6830b3a14c6552a99b7decacbf6daa1524e85ec1acd2aaf543b1e3
---

# hc-en-us-articles-115000166827-unifi-hotspots-and-captive-portals-376af2d3

UniFi Hotspots and Captive Portals
UniFi’s Hotspot Portal allows you to create a professional, custom-branded landing page with flexible authentication options, enabling secure guest connections to your network.
Enabling a Hotspot and Captive Portal
A Hotspot isolates connected clients from the rest of your network, ensuring security and segmentation. You can configure a Hotspot for either a WiFi SSID or an entire network (VLAN):
WiFi Only
- Go to UniFi Network > Settings > WiFi.
- Select or create a WiFi SSID.
- Enable Hotspot Portal > Captive Portal.
Entire Network (VLAN)
- (Optional) Create a new Network in UniFi Network > Settings > Networks.
- Navigate to Settings > Security > Firewall.
  - Note: This requires Zone-Based Firewalling, as part of Network v9.0 and later.
- Select the Hotspot Zone.
- Select the desired Hotspot Network.
Once enabled, go to Insights > Hotspot and click on Landing Page to configure a Captive Portal with branding options, including welcome text, logo, and color scheme.
Authentication
Easily select one or multiple ways for your guests to connect, including:
- Facebook: Guests can authenticate using a Facebook account.
- Password: Guests must enter a password to connect.
- Payment: Guests must pay to use the WiFi. This is currently supported by Stripe.
- Vouchers: Provide guests with vouchers that can be used to authenticate. Customize vouchers to support various expiration times, bandwidth limits, or data consumption quotas.
- RADIUS (advanced): Preconfigure guest authentication via a RADIUS server.
- External Portal Server (advanced): Integrate with a third-party portal server.
Configuring Public Guest WiFi
To learn more about public WiFi best practices, including how to enhance security and optimize performance, click here.
