---
id: collect-261001-general-networking/general-networking/clementdmsn-homelab-writeup-ec921537-1
title: "clementdmsn-homelab-writeup-ec921537"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/clementdmsn-homelab-writeup-ec921537.md
source_anchor: ""
source_lines: [1, 187]
sha256: 3083f1ad30af05ad1879b1e17f90822cc9b8898c3e0445fc11160bbefc063ab0
---

# clementdmsn-homelab-writeup-ec921537

Technologies:

- Proxmox VE
- OPNsense firewall
- Active Directory (Windows Server)
- VLAN segmentation
- DHCP relay
- DMZ architecture

Features:

- network segmentation across four VLANs
- firewall rule isolation between zones
- Active Directory environment with core internal services (DNS, DHCP, domain time services)
- DHCP relay across VLANs
- DMZ-hosted public services
- VPN access for administration

This project implements a virtualized network with four segmented VLANs, inter-VLAN routing on OPNsense, AD-integrated DNS/DHCP, DHCP relay, WAN exposed DMZ services, and WireGuard-based remote administration.

The objective is to study fundamental principles of network architecture, segmentation, and infrastructure security through the construction of a small but realistic enterprise-style topology.

The lab focuses on:

- network segmentation
- firewall-based traffic control
- centralized identity services
- infrastructure administration networks
- service isolation through a DMZ

The infrastructure includes:

- a firewall/router
- administration hosts
- an Active Directory environment
- Windows client machines
- internal infrastructure servers
- a DMZ network hosting public services
- a VPN tunnel for remote administration

All systems are deployed inside a virtualized environment running on a single physical host.

The lab uses separated clients, servers, admin, and DMZ zones to enforce trust boundaries and restrict east-west traffic.

The main design goals were:

**Network segmentation**
The infrastructure is divided into multiple network zones in order to isolate services and reduce the attack surface.

**Least privilege access**
Each network is restricted to the minimum communication required for its function.

**Centralized identity and network services**
Active Directory provides authentication and core network services for domain-joined systems, including DNS, DHCP, and time synchronization.

**Realistic network topology**
The lab is structured around the following network zones :

- administrative networks
- server networks
- client networks
- DMZ service networks
- remote administrative access through a VPN tunnel

The entire infrastructure runs on a single machine acting as a virtualization host.

Because of this constraint, every component of the infrastructure is virtualized:

- firewall
- servers
- client machines
- administration hosts

Virtualization makes it possible to simulate a complete multi-network infrastructure using a single physical system.

The lab uses the following technologies.

**Virtualization platform**
The infrastructure is hosted on Proxmox VE, which provides virtual machines, software networking, and storage management.

**Firewall and routing**
Network routing and security policies are implemented using OPNsense.

**Operating systems**
Linux systems are used for administration and infrastructure services.

**Directory services**
A domain controller is deployed using Windows Server to provide:

- Active Directory Domain Services
- DNS
- DHCP

**WEB SERVER**
Web server is managed by nginx.

**VPN**
WireGuard is used as the VPN solution for remote administrative access and is hosted on OPNsense.

The infrastructure is divided into four isolated network zones.

Each zone represents a different trust level in the infrastructure.

**ADMIN**
The ADMIN network hosts machines are used to manage infrastructure components.
Only systems in this network are allowed to access firewall management interfaces and perform administrative tasks.

**SERVERS**
The SERVERS network hosts internal infrastructure services.

- domain controller
- DNS
- DHCP
- time synchronization
- backup services

**CLIENTS**
The CLIENTS network contains user workstations joined to the Active Directory domain.
This network has restricted access to internal services.

**DMZ**
The DMZ network hosts services intended to be reachable from outside the internal infrastructure.
Services in this zone are isolated from internal systems.

All traffic between network segments is routed through the firewall, which enforces security policies between zones.

Each network zone is implemented as a separate VLAN.

| Network | VLAN | Subnet | Gateway | 
|---|---|---|---|
| CLIENTS | 10 | 192.168.10.0/24 | 192.168.10.1 | 
| SERVERS | 20 | 192.168.20.0/24 | 192.168.20.1 | 
| ADMIN | 30 | 192.168.30.0/24 | 192.168.30.1 | 
| DMZ | 40 | 192.168.40.0/24 | 192.168.40.1 | 

Each subnet uses the firewall as its default gateway.

The VLANs are trunked through a single bridge inside the virtualization platform, allowing the firewall to route traffic between network zones while maintaining isolation.

`ADMIN` network purpose is to isolate privileged access from user and service networks.

`CLIENTS` contains the least trusted internal managed endpoints and is restricted to only required infrastructure services.

`SERVERS` hosts core internal services and should not be broadly reachable from lower-trust zones.

`DMZ` hosts externally reachable services and is isolated to prevent compromise from pivoting inward.

**Inter-zone access policy summary** :

- `ADMIN` may initiate management access to all internal zones.
- `CLIENTS` may access SERVERS only for AD-related services such as DNS, Kerberos, LDAP, SMB, RPC, and NTP.
- `SERVERS` have limited outbound access for updates, DNS, and time synchronization.
- `DMZ` is isolated from internal zones and only allowed minimal supporting services.
- Internet-originated traffic is only forwarded to explicitly published `DMZ` services.
- Management access to firewall services is restricted to `ADMIN` and VPN-originated admin access.

Traffic between networks is controlled by firewall rules.

The firewall follows a **default deny model**, meaning all traffic is blocked unless explicitly allowed.

Security policies are applied between zones to restrict communication to only what is necessary.

Administrative access is allowed either from the ADMIN network or through the WireGuard VPN tunnel. The VPN is restricted to the ADMIN segment only and does not provide broad access to the rest of the internal network.

Firewall aliases are used to simplify rule management by grouping commonly used ports.

| Network | Action | Rule | Source | Destination | Protocol | Ports | 
|---|---|---|---|---|---|---|
| `ADMIN` | PASS | `ALLOW_ADMIN_WEB_OUT` | `ADMIN net` | `*` | TCP | `WEB_PORTS` | 
| `ADMIN` | PASS | `ALLOW_ADMIN_DNS_TO_FORWARD` | `ADMIN net` | `This Firewall` | TCP / UDP | `DNS_PORT` | 
| `ADMIN` | PASS | `ALLOW_MGMT_FROM_ADMIN` | `ADMIN net` | `This Firewall` | TCP | `MGMT_PORTS` | 
| `ADMIN` | PASS | `ALLOW_NTP_TO_FIREWALL` | `ADMIN net` | `This Firewall` | UDP | `NTP_PORT` | 

| Network | Action | Rule | Source | Destination | Protocol | Ports | 
|---|---|---|---|---|---|---|
| `CLIENTS` | BLOCK | `BLOCK_MANAGEMENT_FROM_CLIENTS` | `CLIENTS net` | `This Firewall` | TCP | `MGMT_PORTS` | 
| `CLIENTS` | PASS | `ALLOW_DNS_TO_DC` | `CLIENTS net` | `DC_IP` | TCP / UDP | `DNS_PORT` | 
| `CLIENTS` | PASS | `ALLOW_CLIENTS_BROADCAST_FOR_DHCP` | `CLIENTS net` | `BROADCAST` | UDP | `DHCP_PORTS` | 
| `CLIENTS` | PASS | `ALLOW_CLIENTS_WEB_OUT` | `CLIENTS net` | `*` | TCP | `WEB_PORTS` | 
| `CLIENTS` | PASS | `ALLOW_CLIENTS_RPC_RANGE_TO_DC` | `CLIENTS net` | `DC_IP` | TCP | `RPC_PORTS_RANGE` | 
| `CLIENTS` | PASS | `ALLOW_RPC_TO_DC` | `CLIENTS net` | `DC_IP` | TCP | `RPC_PORT` | 
| `CLIENTS` | PASS | `ALLOW_KERBEROS_TO_DC` | `CLIENTS net` | `DC_IP` | TCP | `KERBEROS_PORTS` | 
| `CLIENTS` | PASS | `ALLOW_NETBIOS_TO_DC` | `CLIENTS net` | `DC_IP` | TCP / UDP | `NETBIOS_PORTS` | 
| `CLIENTS` | PASS | `ALLOW_LDAP_TO_DC` | `CLIENTS net` | `DC_IP` | TCP | `LDAP_PORT` | 
| `CLIENTS` | PASS | `ALLOW_SMB_TO_DC` | `CLIENTS net` | `DC_IP` | TCP | `SMB_PORT` | 
| `CLIENTS` | PASS | `ALLOW_NTP_TO_DC` | `CLIENTS net` | `DC_IP` | UDP | `NTP_PORT` | 

