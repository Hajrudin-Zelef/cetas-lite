---
id: collect-261001-cisco/cisco/github-sdiehl-infosec-gns3-network-lab-gns3-network-topology-lab-featuring-nat-cisco-3725-
title: "github-sdiehl-infosec-gns3-network-lab-gns3-network-topology-lab-featuring-nat-cisco-3725-router-swi"
domain: cisco
role: reference
task: reference
actors: ["Oracle"]
dates: ["2026-04"]
keywords: ["cybersecurity", "ethernet"]
source: docs/RAG/collect-261001-cisco/github-sdiehl-infosec-gns3-network-lab-gns3-network-topology-lab-featuring-nat-cisco-3725-router-swi.md
source_anchor: ""
source_lines: [1, 65]
sha256: beaaf2426a870130caa78ed5160304f62b8a1f895b3a8e0851b2dd6ed9aea2d2
---

# github-sdiehl-infosec-gns3-network-lab-gns3-network-topology-lab-featuring-nat-cisco-3725-router-swi

This project demonstrates the design, configuration, and testing of a simulated enterprise network using GNS3 (Graphical Network Simulator 3). The topology includes a NAT device, a Cisco 3725 router, two Ethernet switches, and six VPCS (Virtual PC Simulator) end devices across two separate subnets.

This lab was completed as part of a Routing and Switching course and represents hands-on application of networking concepts including inter-network routing, Network Address Translation (NAT), and live packet analysis using Wireshark.

**Work in Progress** — Topology diagram with IP annotations and
supporting documentation files are being added. Last updated: April 2026.

⚠️ 

- **GNS3** v2.2.57 — Network emulation platform
- **Oracle VirtualBox** — Hypervisor for GNS3 VM
- **Cisco IOS** 12.4 on c3725 — Real router operating system
- **Wireshark** — Live packet capture and analysis
- **VPCS** — Virtual PC Simulator for end devices
- **Solar-PuTTY** — Terminal emulator for console access

```
NAT1
 |
R1 (Cisco 3725)
├── Switch1
│    ├── PC1 (192.168.1.2)
│    ├── PC2 (192.168.1.3)
│    └── PC3 (192.168.1.4)
└── Switch2
     ├── PC4 (192.168.2.2)
     ├── PC5 (192.168.2.3)
     └── PC6 (192.168.2.4)
```
| Device | IP Address | Subnet Mask | Gateway | 
|---|---|---|---|
| R1 Fa0/1 | 192.168.1.1 | 255.255.255.0 | N/A | 
| R1 Fa1/0 | 192.168.2.1 | 255.255.255.0 | N/A | 
| PC1 | 192.168.1.2 | 255.255.255.0 | 192.168.1.1 | 
| PC2 | 192.168.1.3 | 255.255.255.0 | 192.168.1.1 | 
| PC3 | 192.168.1.4 | 255.255.255.0 | 192.168.1.1 | 
| PC4 | 192.168.2.2 | 255.255.255.0 | 192.168.2.1 | 
| PC5 | 192.168.2.3 | 255.255.255.0 | 192.168.2.1 | 
| PC6 | 192.168.2.4 | 255.255.255.0 | 192.168.2.1 | 

- **VirtualBox** setup and configuration for GNS3 VM
- **Cisco IOS router interfaces** — IP addressing and interface activation
- **Inter-network routing** — traffic routed between two separate subnets
- **Network Address Translation (NAT)** — PAT configured to allow internal devices to reach external networks including verified connectivity to 8.8.8.8
- **Wireshark packet capture** — live analysis of Layer 2 (MAC/ARP) and Layer 3 (IP/ICMP) traffic

- Difference between **network simulation** (Cisco Packet Tracer) and**network emulation** (GNS3)
- **Layer 2 switching** — MAC address based forwarding within a subnet
- **Layer 3 routing** — IP address based forwarding between subnets
- **NAT/PAT** — private to public IP address translation
- **ARP** — address resolution between IP and MAC addresses
- **Wireshark analysis** — distinguishing between Layer 2 and Layer 3 traffic in live captures

The following documents were produced alongside this lab:

| Document | Description | 
|---|---|
| GNS3 VM Setup Guide | VirtualBox configuration with default vs required settings comparison | 
| GNS3 Lab Configuration Steps | Step by step topology build and router configuration | 
| GNS3 NAT Configuration Guide | Full NAT setup with troubleshooting tips | 
| GNS3 Troubleshooting Guide | Common issues and fixes | 

`Cisco IOS` `Network Routing` `NAT/PAT` `Wireshark` `GNS3` `VirtualBox` `TCP/IP` `Network Design` `Technical Documentation` `CLI Configuration`

This project was completed as part of an Associate's Degree in Networking and Cybersecurity. It reflects hands-on experience with enterprise networking tools and concepts directly applicable to roles in network engineering, cybersecurity, and IT infrastructure.
