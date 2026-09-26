---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-82
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-82.md
source_anchor: ""
source_lines: [1, 68]
sha256: 43aa39340936d5fbd20b70936134c56070812c2e21a8ebc856114735ddba7676
---

# Overview

The following steps are a recommendation on how to additionally protect your device with already configured strong firewall rules.

# RouterOS version

Start by upgrading your RouterOS version. Some older releases have had certain weaknesses or vulnerabilities, that have been fixed. Keep your device up to date, to be sure it is secure. Click "check for updates" in WinBox or WebFig, to upgrade. We suggest you follow announcements on our security announcement blog to be informed about any new security issues.

# Access to a router

## Access username

Change the default username *admin* to a different name. A custom name helps to protect access to your router if anybody has direct access to your router: 

## Access password

MikroTik routers require password configuration, we suggest using a password generator tool to create secure and non-repeating passwords. With a secure password, we mean:

- Minimum 12 characters;
- Include numbers, Symbols, Capital and lowercase letters;
- Is not a Dictionary Word or a Combination of Dictionary Words;
- Note that quotes in the password require escaping;

## Securing access to the device

To prevent remote access to your device, there is a pre-configured firewall that blocks WAN (internet side) connections. This is intentional, please do not remove these rules unless you're absolutely certain that the connection is secure.

If you intend to open remote access to your device, we recommend securing the connection using a Virtual Private Network (VPN) such as WireGuard.

A configuration guide for WireGuard VPN is available here.

## RouterOS MAC-access

RouterOS has built-in options for easy management access to network devices. The particular services should be shut down on production networks: **MAC-Telnet, MAC-WinBox,** and **MAC-Ping:**

## Neighbor Discovery

MikroTik Neighbor discovery protocol is used to show and recognize other MikroTik routers in the network, and disable neighbor discovery on all interfaces:

## Bandwidth server

A bandwidth server is used to test throughput between two MikroTik routers. Disable it in the production environment:

## DNS cache

A router might have DNS cache enabled, which decreases the resolving time for DNS requests from clients to remote servers. In case DNS cache is not required on your router or another router is used for such purposes, disable it:

## Other clients services

RouterOS might have other services enabled (they are disabled by default RouterOS configuration). MikroTik caching proxy, socks, UPnP, and cloud services:

More Secure SSH access

It is possible to enable more strict SSH settings (add aes-128-ctr and disallow hmac sha1 and groups with sha1) with this command:

# Router interface

## Ethernet/SFP interfaces

It is good practice to disable all unused interfaces on your router, to decrease unauthorized access to your router:

Where **X** numbers of unused interfaces.

## LCD

Some RouterBOARDs have an LCD module for informational purposes, set a pin:

or disable it:
