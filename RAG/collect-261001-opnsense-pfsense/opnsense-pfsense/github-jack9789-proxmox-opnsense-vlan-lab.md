---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-jack9789-proxmox-opnsense-vlan-lab
title: "github-jack9789-proxmox-opnsense-vlan-lab"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-jack9789-proxmox-opnsense-vlan-lab.md
source_anchor: ""
source_lines: [1, 65]
sha256: 3d93fac2c0aa763b9667ee8f7b4475720229cb0b11ae777c09a557bd0d831c8f
---

# github-jack9789-proxmox-opnsense-vlan-lab

A isolated, VLAN-segmented enterprise-style network lab built with production-like security practices. Designed to refresh and extend skills in networking, system administration, cybersecurity, and IT support.

**GitHub:** https://github.com/jack9789/proxmox-opnsense-vlan-lab.
**Author:** Sheng-You Chen (Adelaide, Australia)

- Demonstrate hands-on experience with modern infrastructure tools (Proxmox VE, OPNsense, Cisco switching, Active Directory).
- Practice secure network segmentation using 802.1Q VLANs and firewall policies.
- Build and manage a multi-VLAN Active Directory environment.
- Simulate real-world IT support and cybersecurity scenarios.
- Document configurations for portfolio and resume use.

- **Proxmox Host** : ASUS NUC14 Pro (single NIC)
- **Firewall/Router** : OPNsense on dedicated AOOSTAR N150 mini PC (2 NICs)
- **Switch** : Cisco SG300-28 (managed, Layer 2+ with 802.1Q support)
- **ISP Router** : Home modem/router (double-NAT for full lab isolation)
- **Proxmox Backup Server** : Acer Aspire 3 (single NIC)
- **Guest WIFI** : GL-BE3600 wireless router

| VLAN | Purpose | Subnet | Key Devices / VMs | 
|---|---|---|---|
| 10 | Management | 192.168.10.0/24 | Proxmox host (192.168.10.5) OPNsense (192.168.10.1) | 
| 20 | Servers | 192.168.20.0/24 | Windows Server 2022 AD DC (192.168.20.10) AlmaLinux Docker host (192.168.20.20) – Portainer, Nextcloud, Zammad, SearXNG, etc. | 
| 30 | Trusted Clients | 192.168.30.0/24 | Windows 11 clients (DHCP from OPNsense) Ubuntu desktop client (DHCP) | 
| 40 | Attacker / DMZ | 192.168.40.0/24 | Kali Linux (192.168.40.10 – static) | 
| 50 | Guest WIFI | 192.168.50.0/24 | GL-BE3600 wireless router (Planning) | 

- **Trunk ports** on Cisco SG300 carry tagged VLANs 10, 20, 30, 40 to both OPNsense and Proxmox.
- **Double-NAT isolation** : Lab sits behind home ISP router – zero impact on daily devices.
- **Inter-VLAN routing & firewalling** handled by OPNsense with strict rules (e.g., block Attacker → Servers by default).

- **Virtualization** : Proxmox VE – VMs and Containers across VLANs
- **Directory Services** : Active Directory Domain Controller (DNS, DHCP scopes for trusted VLANs, GPOs)
- **Docker Services** (on AlmaLinux host):
  - Portainer (management)
  - NetData (monitoring)
  - Zammad (helpdesk ticketing)
  - SearXNG (private search)
  - AnythingLLM (self-hosted AI)
  - Nginx Proxy Manager(Reverse Proxy)
- **Security** :
  - Strict firewall rules between VLANs
  - Double-NAT isolation: Lab network sits behind the home ISP router (first NAT) and OPNsense firewall (second NAT)
  - Isolated attacker VLAN for safe testing
- **Monitoring & Backup** : Proxmox Backup Server on vms

- Cisco IOS VLAN/trunk configuration
- OPNsense VLAN sub-interfaces, DHCP relay, firewall rules
- Proxmox single-NIC VLAN-aware bridge setup
- Windows Server AD DS, DNS, DHCP, multi-subnet management
- Linux system administration (AlmaLinux, Ubuntu domain join via SSSD)
- Docker container deployment and management
- Network troubleshooting across segmented environment

- Sanitized configuration exports (OPNsense config.xml, Cisco running-config, Proxmox network interfaces)
- Topology diagram

- Add VLAN 50 for guest wifi
- Implement Guacamole for browser-based remote desktop (IT support simulation)
- More advanced GPOs and monitoring

**This project is actively maintained and used for continuous learning.**

Sheng-You Chen

jack9789@gmail.com
