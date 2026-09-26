---
id: collect-260926-mikrotik/mikrotik/github-whilcayangyang-mikrotik-homelab-config-a-production-grade-routeros-7-configuration--5
title: "Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["copyright", "license"]
source: docs/RAG/lot-mikrotik/RouterOS/github-whilcayangyang-mikrotik-homelab-config-a-production-grade-routeros-7-configuration-for-a-segm.md
source_anchor: ""
source_lines: [468, 491]
sha256: 95c514c040dc5f0a0071c8bf1b11ce9e94d61b7faaf946e9773993ed84a101f5
---

# Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:

1. Back up the current configuration first: `/export file=backup-before-changes`
2. Replace every `<placeholder>` above with your real values (VLAN subnets, device MACs, ProtonVPN keys/endpoints, admin IPs, SMTP credentials, timezone, NTP pool).
3. Apply in RouterOS CLI/Winbox terminal, in the order presented (interfaces → addressing → firewall → services) — firewall rules depend on address-lists and interface-lists existing first.
4. Verify VPN kill-switch behavior: disable a WireGuard peer temporarily and confirm a device in its address-list **loses connectivity** rather than falling back to plain WAN.
5. Verify brute-force lockout: attempt 3 failed logins from a non-whitelisted IP within 5 minutes and confirm the source lands in `admin-blacklist` .

- **Never publish an unredacted `/export`** — it can include stored secrets (SMTP, RADIUS shared secrets, occasionally cached credentials) depending on RouterOS version and`hide-sensitive` flag.
- **WireGuard private keys** are never shown in this README and should never be committed anywhere, including private repos, without encryption at rest.
- **Static DHCP leases + MAC-based address lists** are the backbone of this VPN-routing/ACL scheme — keep that mapping current when devices are replaced.
- Review `admin-blacklist` /`admin-stage1` /`admin-stage2` address-lists periodically (`/ip firewall address-list print where list~"admin-"` ) to catch active brute-force attempts against your admin ports.

Licensed under the GNU General Public License v3.0.

```
Copyright (C) 2026  <your-name-or-handle>
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.
```
