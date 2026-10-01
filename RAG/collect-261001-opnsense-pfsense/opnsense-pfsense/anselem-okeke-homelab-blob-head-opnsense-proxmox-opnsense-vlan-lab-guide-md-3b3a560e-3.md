---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e-3
title: "anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e.md
source_anchor: ""
source_lines: [935, 986]
sha256: 0451741f79e691d5d4bb1a072147969025aba8758af922834af06cfc710cc1e1
---

# anselem-okeke-homelab-blob-head-opnsense-proxmox-opnsense-vlan-lab-guide-md-3b3a560e

```
SERVERS/CLIENTS/GUEST → MGMT
```
Replace broad rules like:

```
Allow CLIENTS to any
Allow SERVERS to any
```
with stricter rules.

Example for CLIENTS to Windows Server:

```
CLIENTS net → 192.168.20.162 → DNS 53
CLIENTS net → 192.168.20.162 → Kerberos 88
CLIENTS net → 192.168.20.162 → LDAP 389
CLIENTS net → 192.168.20.162 → SMB 445
CLIENTS net → 192.168.20.162 → RDP 3389 if needed
```
For learning, broad rules are okay first. For enterprise practice, tighten rules after verification.

SERVERS VLAN:

```
Gateway: 192.168.20.1
Windows Server: 192.168.20.x
DHCP working
```
CLIENTS VLAN:

```
Gateway: 192.168.30.1
Windows 11: 192.168.30.x
DHCP working
```
Next required step:

```
Add firewall rules:
CLIENTS network → any
SERVERS network → any
```
Then test:

```
ping gateway
ping 8.8.8.8
nslookup google.com
browse internet
access OPNsense UI using the VLAN gateway IP
```
