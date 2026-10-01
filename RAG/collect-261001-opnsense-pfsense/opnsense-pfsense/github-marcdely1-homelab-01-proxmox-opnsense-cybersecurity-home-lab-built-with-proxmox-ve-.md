---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-marcdely1-homelab-01-proxmox-opnsense-cybersecurity-home-lab-built-with-proxmox-ve-
title: "github-marcdely1-homelab-01-proxmox-opnsense-cybersecurity-home-lab-built-with-proxmox-ve-and-opnsen"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["AWS", "Intel"]
dates: []
keywords: ["cybersecurity", "aws", "consumer", "ethernet", "intel", "license", "mit license", "open source"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-marcdely1-homelab-01-proxmox-opnsense-cybersecurity-home-lab-built-with-proxmox-ve-and-opnsen.md
source_anchor: ""
source_lines: [1, 122]
sha256: 4b44dfcef5639edd54010794bf2a0b8a4576c902cdd7dae83424c4028f5c1d11
---

# github-marcdely1-homelab-01-proxmox-opnsense-cybersecurity-home-lab-built-with-proxmox-ve-and-opnsen

A fully virtualized cybersecurity home lab running OPNsense firewall as a VM on Proxmox VE, built on Dell Optiplex 3060 Micro hardware.


**Author:** Marc Dely

**Location:** Boston, MA

**LinkedIn:** linkedin.com/in/marc-dely-439916186

**Certification:** CompTIA Security+

This project documents the design, build, and configuration of a cybersecurity home lab using enterprise-grade open source software on consumer hardware. The lab serves as a foundation for hands-on practice in network security, firewall administration, intrusion detection, and cloud security concepts.

**Core goal:** Build a production-like security environment to
develop skills toward a career as a Cybersecurity Analyst and
Cloud Security Architect.

| Component | Details | 
|---|---|
| Proxmox Host | Dell Optiplex 3060 Micro | 
| CPU | Intel Core i5-8500T (8th Gen) | 
| RAM | 16GB DDR4 | 
| Storage | 256GB NVMe SSD | 
| WAN NIC | TP-Link USB 3.0 Gigabit Ethernet (UE300) | 
| GUI Workstation | Dell Optiplex 3060 Micro | 
| Switch | Netgear FS105 5-Port Fast Ethernet | 
| Rack | 10-inch Mini Rack | 

| Software | Version | Role | 
|---|---|---|
| Proxmox VE | 9.1.1 | Hypervisor | 
| OPNsense | 26.1 | Firewall / Router | 
| FreeBSD | HardenedBSD | OPNsense base OS | 
| Debian Linux | 12 | Proxmox base OS | 

ISP Modem │ └── TP-Link USB NIC (ue0) ← WAN interface │ OPNsense VM (FreeBSD) │ vtnet1 (LAN) → 192.1xx.x.x/xx │ Proxmox vmbr0 bridge │ nic0 (built-in NIC) │ Netgear Switch │ ├── Proxmox Host (192.1xx.x.x) └── Workstation / Optiplex 2 (192.1xx.x.x)

| Device | IP Address | Role | 
|---|---|---|
| OPNsense LAN | 192.16x.x.x | Default gateway | 
| Proxmox Host | 192.16x.x.x | Hypervisor management | 
| DHCP Pool | 192.16x.x.x00–x00 | Client devices | 

Running OPNsense as a VM enables instant snapshots before any config change, fast rollback if something breaks, and the ability to run additional VMs (vulnerable machines, SIEM, etc.) on the same host.

The TP-Link USB NIC is passed directly to the OPNsense VM rather than bridged. This gives OPNsense exclusive ownership of the WAN interface and keeps Proxmox off the internet-facing network entirely.

The WAN bridge (vmbr1) has no IP address assigned on the Proxmox host. This keeps Proxmox invisible on the WAN segment — only OPNsense can speak on that network, preventing direct internet exposure of the hypervisor.

- Install Proxmox VE on Dell Optiplex 3060 Micro
- Configure static IP on vmbr0 (192.16x.x.x/xx)
- Create vmbr1 as WAN bridge with no IP

Edit `/etc/network/interfaces`:

```
auto vmbr0
iface vmbr0 inet static
        address 192.16x.x.x/xx
        gateway 192.16x.x.x
        bridge-ports nic0
        bridge-stp off
        bridge-fd 0
auto vmbr1
iface vmbr1 inet manual
        bridge-ports none
        bridge-stp off
        bridge-fd 0
```
- Machine: q35
- BIOS: OVMF (UEFI)
- CPU: 2 cores, x86-64-v2-AES
- RAM: 4GB
- Disk: 64GB
- Network: vtnet0 → vmbr1 (WAN), vtnet1 → vmbr0 (LAN)
- USB passthrough: TP-Link NIC (0bda:8153)

- Boot from ISO in Proxmox noVNC console
- Install to da0 (64GB virtual disk)
- Assign interfaces: WAN → ue0, LAN → vtnet1
- LAN IP: 192.16x.x.x/xx

- DNS: 8.8.8.8 / 8.8.4.4
- Timezone: America/New_York
- DHCP range: 192.16x.x.x00–x00
- Web GUI: HTTPS on port 443

Proxmox sets `bridge-vlan-aware yes` by default. This caused
untagged traffic to be silently dropped between OPNsense and LAN
clients. Fix: remove `bridge-vlan-aware yes` and
`bridge-vids 2-4094` from vmbr0 in `/etc/network/interfaces`.

OPNsense had `192.16x.x.x` on both vtnet0 and vtnet1 due to
leftover config from interface reassignment. This caused OPNsense
to receive packets on vtnet1 but attempt responses via vtnet0.
Fix: `ifconfig vtnet0 0.0.0.0` and `arp -d 192.16x.x.x00`.

The TP-Link USB NIC appears as `ue0` in FreeBSD (OPNsense), not
`vtnet0` as expected. Always verify interface names with
`ifconfig -a` before assigning WAN/LAN.

OPNsense boots into live mode by default from ISO. Interface
assignments don't persist until properly installed to disk using
the `installer` login. Always verify with `df -h` and
`ls /conf/` that config.xml exists on disk.

- Configure Suricata IDS/IPS for intrusion detection
- Set up WireGuard VPN for remote lab access
- Deploy CrowdSec for threat intelligence
- Add managed switch for VLAN segmentation
- Deploy vulnerable VMs (Metasploitable, DVWA)
- Set up SIEM (Wazuh or Splunk)
- Integrate with cloud environment (AWS/Azure security labs)

- Hypervisor deployment and VM management (Proxmox VE)
- Network bridge configuration (Linux networking)
- Firewall deployment and configuration (OPNsense/FreeBSD)
- Network troubleshooting (tcpdump, ARP, pfctl)
- Linux and BSD command line administration
- Security architecture design (network segmentation, DMZ concepts)

MIT License — see LICENSE for details.
