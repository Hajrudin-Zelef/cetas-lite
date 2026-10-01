---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/oriolrius-unifi-cli-5c4660d1
title: "Run directly with uvx (no installation needed)"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/oriolrius-unifi-cli-5c4660d1.md
source_anchor: ""
source_lines: [1, 66]
sha256: 0ae96184d5780ac3ca0ce53e2874303d64157f0c29c0774e846e8063e298de1b
---

# Run directly with uvx (no installation needed)

CLI tool for managing UniFi UDM Pro networks. Built on aiounifi with custom SSO + TOTP MFA authentication.
- List clients with online/offline status, IP, MAC, network, vendor info
- List network devices (APs, switches, gateways) with firmware and client counts
- List wireless networks (WLANs)
- List connect/disconnect sessions over a time window (diagnose WiFi reconnect storms)
- Create/clear DHCP fixed-IP reservations and force a client to reconnect (kick)
- Filter by status (online/offline), network name, and sort by any field
- Output as rich table, JSON, or CSV
- Auth token caching (avoids repeated TOTP prompts)
- Raw API endpoint access for advanced queries (GET or POST with a JSON body)
# Run directly with uvx (no installation needed)
uvx --from "git+https://github.com/oriolrius/unifi-cli.git" unifi-cli configure
uvx --from "git+https://github.com/oriolrius/unifi-cli.git" unifi-cli clients
# Or install as a tool
uv tool install "git+https://github.com/oriolrius/unifi-cli.git"
unifi-cli configure
unifi-cli clients# Interactive configuration
unifi-cli configure
# Or set environment variables
export UNIFI_HOST=192.168.1.1
export UNIFI_PORT=443
export UNIFI_SITE=default
export UNIFI_USERNAME=admin
export UNIFI_PASSWORD=secret
export UNIFI_TOTP_SECRET=BASE32SECRET
Config is stored in ~/.config/unifi-cli/config.json (mode 600).
# List all clients (online + offline)
unifi-cli clients
# Online clients only
unifi-cli clients --status online
# Filter by network and sort by hostname
unifi-cli clients --network IoT --sort hostname
# JSON output
unifi-cli --format json clients --status online
# CSV output
unifi-cli --format csv clients
# List network devices
unifi-cli devices
# List WLANs
unifi-cli networks
# Connect/disconnect sessions in the last 72h (all clients)
unifi-cli sessions
# Sessions for one device (by MAC or by name/hostname), last 24h
unifi-cli sessions --name Wallbox --hours 24
unifi-cli sessions --mac 24:cd:8d:85:d6:ae --sort duration
# DHCP fixed-IP reservation (by MAC or name), then force the client onto it
unifi-cli reserve --name madoka2mqtt --ip 10.2.8.204 --label madoka2mqtt
unifi-cli reconnect --name madoka2mqtt
unifi-cli reserve --mac b8:27:eb:02:f7:3a --clear   # remove the reservation
# Raw API query (GET)
unifi-cli raw stat/health
unifi-cli raw stat/device
# Raw API query (POST with a JSON body)
unifi-cli raw -X post -d '{"type":"all","within":72,"mac":"24:cd:8d:85:d6:ae"}' stat/session
This tool handles the UDM Pro's Ubiquiti SSO + TOTP MFA authentication automatically:
- Sends credentials to get an MFA challenge
- Generates a TOTP code from your secret
- Completes authentication with the MFA cookie
- Caches the session token for ~23 hours
If the cached token expires, it automatically re-authenticates.
Note: aiounifi v90 has built-in SSO MFA support, but has a cookie-handling bug where the MFA cookie is set without URL context, causing aiohttp to never send it. This tool includes a workaround.
- Python >= 3.13
- UniFi OS controller (UDM, UDM Pro, UDM SE, Cloud Key Gen2+)
- Ubiquiti account with TOTP MFA enabled
This tool is built on aiounifi by Kane610. The library's async architecture and typed data models made it a pleasure to work with. Special thanks to the aiounifi team for maintaining this project and to the Home Assistant community for the extensive documentation of the UniFi API quirks.
MIT
