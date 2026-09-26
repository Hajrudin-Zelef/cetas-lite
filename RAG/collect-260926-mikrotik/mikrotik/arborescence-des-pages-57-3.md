---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-57-3
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-57.md
source_anchor: ""
source_lines: [247, 273]
sha256: 4cdfb65a0e8fda88e20757093dbb33c473aa95a32427f4a19c22ef15bdeb8f9b
---

# Introduction

Second mangle rule marks connections which is further used to mark routing and also ensures that mark stays on connection after IP address is gone from the address list and tunnel is established.

Third mangle rule is used to force the packet to use the correct routing table for the second WAN interface.

Last NAT rule is required to ensure that packet is sent out with the correct source IP, as it is not adjusted by the mangle rules and if not implemented packet could have different source IP depending on the setup. 

As you can see rules are duplicated for WAN3, to ensure that WAN3 interface is also usable with WireGuard.

This rule set, ensures that WireGuard tunnel is established on the interface that recived the incoming handshake.

# 2FA setup

HotSpot services can be bound to WireGuard interfaces in RouterOSv7. This allows administrators to combine WireGuard’s high-performance encryption with the HotSpot’s captive portal features, supporting additional user-level authentication via local database, User Manager, or RADIUS, and OTP feature allows to strengthen security even more.

## Configuration example

Make sure Wireguard interface is configured on your device:

Setup HotSpot on Wireguard interface:

After HotSpot setup, once Wireguard peer is connected, additional login to portal will be required to proceed. Regular HotSpot username/password authorization could be used. For better security it is possible to utilize /ip hotspot user otp-password or /user-manager/user totp-password field. Add user with totp-secret:

You can find multiple open-source totp generator tools, that can provide you with appropriate key and necessary options to import it to your used "authenticator" app.

Once setup is complete, launch Wireguard on your device, and after successful authorization Wireguard/HotSpot server, HotSpot login page will be prompted and you will need to enter HotSpot username and totp 6 symbol password (that is changing once in 30 seconds) from "authenticator" app.

It is strongly recommended to use HTTPs login page for HotSpot, and make sure Wireguard peer has an option to verify login page certificate.
