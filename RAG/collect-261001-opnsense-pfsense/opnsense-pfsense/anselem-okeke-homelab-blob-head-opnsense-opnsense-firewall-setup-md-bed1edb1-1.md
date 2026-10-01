---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/anselem-okeke-homelab-blob-head-opnsense-opnsense-firewall-setup-md-bed1edb1-1
title: "anselem-okeke-homelab-blob-head-opnsense-opnsense-firewall-setup-md-bed1edb1"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/anselem-okeke-homelab-blob-head-opnsense-opnsense-firewall-setup-md-bed1edb1.md
source_anchor: ""
source_lines: [1, 433]
sha256: efc97ec9bdf87e433d78db1c5ec2faa7e674aba500950ae321843ba2357d2c8e
---

# anselem-okeke-homelab-blob-head-opnsense-opnsense-firewall-setup-md-bed1edb1

This document describes the step-by-step setup of **OPNsense** as the firewall/router layer for the homelab environment.

The goal is to provide:

- Central firewall and routing control
- LAN/WAN separation
- DHCP and DNS services
- VLAN-ready network design
- Secure management access
- Foundation for monitoring and future automation

```
                    Internet / ISP Router
                            |
                            |
                         WAN NIC
                      +-------------+
                      |  OPNsense   |
                      |  Firewall   |
                      +-------------+
                         LAN NIC
                            |
                            |
                      Homelab Switch
                            |
        ------------------------------------------------
        |                    |                         |
     Jumpbox              Proxmox                 Kubernetes Nodes
  192.168.30.x        192.168.30.x              192.168.30.x
```
This guide assumes:

- OPNsense is already installed
- OPNsense has at least two interfaces:
  - `WAN`
  - `LAN`
- The LAN network uses:

```
LAN subnet: 192.168.30.0/24
OPNsense LAN IP: 192.168.30.1
```
Adjust the IP addresses if your environment uses a different subnet.

After installation, access the OPNsense web UI from a machine on the LAN:

```
https://192.168.30.1
```
Default login:

```
Username: root
Password: opnsense
```
Change the default password immediately after first login.


Go to:

```
System → Access → Users
```
Edit the `root` user and set a strong password.

Recommended password rules:

- Minimum 16 characters
- Uppercase and lowercase letters
- Numbers
- Special characters
- Store in a password manager

Go to:

```
Interfaces → LAN
```
Recommended LAN settings:

```
Enable interface: Yes
IPv4 Configuration Type: Static IPv4
IPv4 Address: 192.168.30.1/24
IPv6 Configuration Type: None or DHCPv6 if required
```
Save and apply changes.

Go to:

```
Interfaces → WAN
```
Common WAN options:

```
IPv4 Configuration Type: DHCP
```
Use this if OPNsense receives an IP address from your home router.

```
IPv4 Configuration Type: Static IPv4
IPv4 Address: <WAN-IP>
Gateway: <ISP-GATEWAY>
```
Use this only if your ISP or upstream router provides static addressing.

Save and apply changes.

Go to:

```
Services → ISC DHCPv4 → LAN
```
Enable DHCP on LAN:

```
Enable DHCP server on LAN interface: Yes
```
Example DHCP range:

```
From: 192.168.30.100
To:   192.168.30.200
```
Recommended static/manual IP range:

```
192.168.30.2   - 192.168.30.49    Infrastructure
192.168.30.50  - 192.168.30.99    Servers / Kubernetes / Proxmox
192.168.30.100 - 192.168.30.200   DHCP clients
192.168.30.201 - 192.168.30.254   Reserved / future use
```
Save and apply.

| Component | Example IP | Notes | 
|---|---|---|
| OPNsense LAN | `192.168.30.1` | Default gateway | 
| Proxmox host | `192.168.30.10` | Virtualization host | 
| Jumpbox | `192.168.30.20` | Admin machine | 
| Talos control plane 1 | `192.168.30.31` | Kubernetes control plane | 
| Talos control plane 2 | `192.168.30.32` | Kubernetes control plane | 
| Talos control plane 3 | `192.168.30.33` | Kubernetes control plane | 
| Talos worker 1 | `192.168.30.41` | Kubernetes worker | 
| Talos worker 2 | `192.168.30.42` | Kubernetes worker | 

Go to:

```
Services → Unbound DNS → General
```
Enable Unbound:

```
Enable Unbound: Yes
Listen Port: 53
Network Interfaces: LAN
```
Recommended options:

```
Register DHCP leases: Yes
Register DHCP static mappings: Yes
```
This allows DHCP hostnames to be resolved locally.

Example:

```
ping proxmox.local
ping jumpbox.local
```
Go to:

```
System → Settings → General
```
Recommended DNS servers:

```
1.1.1.1
8.8.8.8
```
Or use privacy-focused resolvers:

```
9.9.9.9
1.1.1.1
```
Recommended setting:

```
Allow DNS server list to be overridden by DHCP/PPP on WAN: No
```
Go to:

```
Firewall → Rules → LAN
```
Default LAN rule usually allows LAN clients to access any destination.

Example default LAN rule:

```
Action: Pass
Interface: LAN
Protocol: Any
Source: LAN net
Destination: Any
Description: Allow LAN to any
```
This is acceptable for the initial homelab setup.

Later, this can be hardened into:

- Admin VLAN
- Server VLAN
- Kubernetes VLAN
- Guest VLAN
- IoT VLAN

Go to:

```
Firewall → Rules → WAN
```
By default, inbound WAN traffic should be blocked.

Recommended baseline:

```
No public inbound access unless explicitly required.
```
Avoid exposing:

- OPNsense Web UI
- Proxmox UI
- Kubernetes API
- SSH
- Grafana
- Longhorn UI

If remote access is required, use VPN instead.

Go to:

```
System → Settings → Administration
```
Enable SSH only if needed:

```
Enable Secure Shell: Yes
Permit root user login: No
Password login: No, if using SSH keys
Listen Interfaces: LAN only
```
Recommended:

```
Allow SSH only from jumpbox/admin machine.
```
Example firewall rule:

```
Action: Pass
Interface: LAN
Protocol: TCP
Source: Jumpbox IP
Destination: This Firewall
Destination Port: 22
Description: Allow SSH from Jumpbox to OPNsense
```
Go to:

```
System → Settings → Administration
```
Recommended:

```
Protocol: HTTPS
TCP Port: 443
Listen Interfaces: LAN
```
Do not expose the web UI on WAN.

Optional hardening:

```
Disable web GUI redirect rule: Yes
Session timeout: 15-30 minutes
```
Go to:

```
Services → ISC DHCPv4 → LAN
```
Under DHCP leases, add static mappings for important systems.

Recommended static mappings:

| Host | IP | 
|---|---|
| `proxmox` | `192.168.30.10` | 
| `jumpbox` | `192.168.30.20` | 
| `talos-cp1` | `192.168.30.31` | 
| `talos-cp2` | `192.168.30.32` | 
| `talos-cp3` | `192.168.30.33` | 
| `talos-w1` | `192.168.30.41` | 
| `talos-w2` | `192.168.30.42` | 

This makes the network easier to document, troubleshoot, and monitor.

From the jumpbox or any LAN machine:

`ping -c 4 192.168.30.1`
Expected result:

```
Replies from OPNsense LAN IP
```
`ping -c 4 1.1.1.1`
Expected result:

```
Internet routing works
```
`nslookup google.com`
Or:

`dig google.com`
Expected result:

```
DNS resolves successfully
```
`ip route`
Expected default route:

```
default via 192.168.30.1 dev <interface>
```
Example:

```
default via 192.168.30.1 dev ens18
```
From the OPNsense web UI:

```
Interfaces → Overview
```
Verify:

| Interface | Expected Status | 
|---|---|
| WAN | Up | 
| LAN | Up | 
| WAN IP | Assigned | 
| LAN IP | `192.168.30.1/24` | 

Go to:

```
Services → ISC DHCPv4 → Leases
```
Check that LAN devices receive IP addresses from the configured DHCP range.

Expected:

```
Clients receive 192.168.30.100 - 192.168.30.200
Gateway is 192.168.30.1
DNS is 192.168.30.1 or configured resolver
```
Go to:

```
Firewall → Log Files → Live View
```
Useful filters:

```
Source IP: 192.168.30.x
Destination port: 53
Destination port: 443
Action: Block
```
Use this during troubleshooting to verify whether traffic is being passed or blocked.

After basic connectivity works, begin tightening access.

```
Source: LAN net
Destination: Any
Protocol: Any
Action: Pass
```
```
Source: Jumpbox IP
Destination: This Firewall
Ports: 22, 443
Action: Pass
```
```
Source: LAN net
Destination: This Firewall
Ports: 22, 443
Action: Block
```
Place allow rules above block rules.

Future VLAN design:

| VLAN | Name | Example Subnet | Purpose | 
|---|---|---|---|
| 10 | Management | `192.168.10.0/24` | Proxmox, OPNsense, switches | 
| 20 | Servers | `192.168.20.0/24` | Linux servers | 
| 30 | Kubernetes | `192.168.30.0/24` | Talos/Kubernetes nodes | 
| 40 | Monitoring | `192.168.40.0/24` | Grafana, Prometheus, Loki | 
| 50 | Guest | `192.168.50.0/24` | Guest devices | 
| 60 | IoT | `192.168.60.0/24` | Cameras, smart devices | 

Recommended security direction:

```
Management VLAN can access all required systems.
Server/Kubernetes VLANs have restricted access.
Guest and IoT VLANs cannot access management systems.
```
Go to:

