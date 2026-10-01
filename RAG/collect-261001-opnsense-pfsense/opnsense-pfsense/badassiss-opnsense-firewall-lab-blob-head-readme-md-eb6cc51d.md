---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/badassiss-opnsense-firewall-lab-blob-head-readme-md-eb6cc51d
title: "OPNsense Firewall Lab — Network Segmentation & IDS/IPS"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cybersecurity", "exploit"]
source: docs/RAG/collect-261001-opnsense-pfsense/badassiss-opnsense-firewall-lab-blob-head-readme-md-eb6cc51d.md
source_anchor: ""
source_lines: [1, 145]
sha256: 4c8da2df28ac770f69a3024aadcd24499fc873f0869c65b1a3f5145be4ac34e5
---

# OPNsense Firewall Lab — Network Segmentation & IDS/IPS

A fully configured software firewall lab built on VMware, demonstrating enterprise-grade network segmentation, firewall rule management, and real-time intrusion detection with Suricata.



This lab simulates a real-world network security architecture using OPNsense 26.1.6 as the core firewall and router. The environment is segmented into three isolated zones — Admin, User, and DMZ — each with dedicated firewall rules and monitored by Suricata IDS/IPS across all interfaces.



```
                        Internet (WAN)
                              |
                         em0 / WAN
                              |
                    ┌─────────────────┐
                    │   OPNsense VM   │
                    │  Firewall/Router │
                    │  Suricata IDS   │
                    └─────────────────┘
                              |
                         em1 trunk
                              |
          ┌───────────────────┼───────────────────┐
          |                   |                   |
   VLAN 10 (Admin)     VLAN 20 (User)      VLAN 30 (DMZ)
   192.168.10.0/24    192.168.20.0/24    192.168.30.0/24
          |                   |                   |
     Admin VM            User VM           Web Server VM
  192.168.10.180      192.168.20.100      192.168.30.100
```


| Technology | Version | Role | 
| OPNsense | 26.1.6 | Firewall, Router, DHCP | 
| Suricata | Built-in | IDS/IPS Engine | 
| VMware Workstation | Latest | Hypervisor | 
| TinyCore Linux | 14.x | Client VMs | 
| Kea DHCP | Built-in | DHCP Server | 
| Emerging Threats | ET Open | IDS Rulesets | 


| VM | OS | RAM | Network | IP | 
| OPNsense | FreeBSD | 2GB | Bridged + VMnet1/2/3 | Gateway | 
| Admin-VM | TinyCore Linux | 256MB | VMnet1 | 192.168.10.180 | 
| User-VM | TinyCore Linux | 256MB | VMnet2 | 192.168.20.100 | 
| DMZ-VM | TinyCore Linux | 256MB | VMnet3 | 192.168.30.100 | 

| Interface | Device | Subnet | Zone | 
| WAN | em0 | DHCP | Internet | 
| LAN | em1 | 192.168.10.0/24 | Admin | 
| USER | em2 | 192.168.20.0/24 | Users | 
| DMZ | em3 | 192.168.30.0/24 | DMZ | 



### Admin Zone (VLAN 10) — Full Access

```
Pass | LAN net → any | Admin has unrestricted access
```

### User Zone (VLAN 20) — Internet Only

```
Block | USER net → LAN net   | Cannot reach Admin zone
Block | USER net → DMZ net   | Cannot reach DMZ zone
Pass  | USER net → any       | Internet access allowed
```

### DMZ Zone (VLAN 30) — Isolated

```
Block | DMZ net → LAN net    | Cannot reach Admin zone
Block | DMZ net → USER net   | Cannot reach User zone
Pass  | DMZ net → any        | Limited internet for updates
```


## Suricata IDS/IPS Configuration

| Parameter | Value | 
| Mode | PCAP live (IDS) | 
| Interfaces | WAN, LAN, USER, DMZ | 
| Promiscuous mode | Enabled | 
| Auto-update schedule | Daily 03:00 | 

- ET Open Emerging Malware
- ET Open Emerging Exploit
- ET Open Emerging Scan
- ET Open Emerging DOS
- ET Open Emerging Attack Response
- ET Open Emerging Web Server
- ET Open Emerging DNS
- ET Open Emerging Misc
- ET Open Emerging Info


## VLAN Isolation Test Results

| Test | From | To | Expected | Result | 
| Admin → Internet | Admin-VM | 8.8.8.8 | ✅ Pass | ✅ Pass | 
| Admin → User | Admin-VM | 192.168.20.1 | ✅ Pass | ✅ Pass | 
| Admin → DMZ | Admin-VM | 192.168.30.1 | ✅ Pass | ✅ Pass | 
| User → Admin | User-VM | 192.168.10.1 | ❌ Block | ❌ Blocked | 
| User → DMZ | User-VM | 192.168.30.1 | ❌ Block | ❌ Blocked | 
| DMZ → Admin | DMZ-VM | 192.168.10.1 | ❌ Block | ❌ Blocked | 
| DMZ → User | DMZ-VM | 192.168.20.1 | ❌ Block | ❌ Blocked | 
| Suricata Alert | Windows Host | testmynids.org | ✅ Alert | ✅ Detected | 


## Security Hardening Applied

- ✅ WebGUI access restricted to LAN interface only
- ✅ SSH authentication hardened
- ✅ Block RFC1918 private networks on WAN
- ✅ Block bogon networks on WAN
- ✅ Firewall logging enabled on all block rules
- ✅ Suricata daily rule auto-updates scheduled
- ✅ Configuration backup saved


- Software firewall deployment and configuration
- Network segmentation using VLANs
- Firewall rule writing and policy enforcement
- Intrusion Detection System (IDS) deployment
- DHCP server configuration (Kea DHCP)
- Network traffic analysis with tcpdump
- VMware virtual network design
- Linux command line administration
- Security hardening best practices


- How OPNsense processes firewall rules (first-match, top-down)
- The difference between IDS (detect) and IPS (block) modes
- How Suricata inspects traffic across multiple interfaces
- Why DMZ zones must never have direct access to internal networks
- How Kea DHCP serves multiple subnets across VLANs
- How NAT enables internal VMs to reach the internet through a single WAN IP
- Diagnosing network issues using tcpdump and pfctl



Built as a hands-on cybersecurity/network administration portfolio project.
Demonstrates practical skills in firewall management, network segmentation, and intrusion detection.
