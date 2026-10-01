---
id: collect-261001-cisco/cisco/ayoub-lah-enterprise-network-security-lab-4e893e59-1
title: "1. Import EVE-NG OVA into VirtualBox"
domain: cisco
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-cisco/ayoub-lah-enterprise-network-security-lab-4e893e59.md
source_anchor: ""
source_lines: [1, 180]
sha256: 6807a9798f072f6e1a9489afb966b11239624f0a5362156f87d95c29f038fb7c
---

# 1. Import EVE-NG OVA into VirtualBox

A complete enterprise-grade network security infrastructure simulation built on EVE-NG
FortiGate HA Cluster · IPsec Site-to-Site VPN · SSL VPN · IPS/IDS · VLAN Segmentation · SNMP Monitoring
Features · Architecture · Technologies · Lab Setup · Configuration · Testing · Results
This project simulates a production-ready enterprise network infrastructure using industry-standard tools and technologies. The lab demonstrates real-world network engineering and security concepts including firewall high availability, encrypted tunnels, VLAN segmentation, intrusion prevention, and infrastructure monitoring.
| Feature | Technology | Status | 
|---|---|---|
| Network Simulation | EVE-NG Community 6.2.0-4 | ✅ | 
| Next-Gen Firewall | FortiGate v7.0.9 | ✅ | 
| High Availability | FortiGate Active-Passive HA | ✅ | 
| WAN Routing | Cisco vIOS Router (R3-Siege) | ✅ | 
| VLAN Segmentation | 6 VLANs (Users, Servers, Guest, DMZ, WAN) | ✅ | 
| DHCP Services | FortiGate DHCP per VLAN | ✅ | 
| IPsec VPN | Site-to-Site (2 remote sites) | ✅ | 
| SSL VPN | Web Mode + Tunnel Mode | ✅ | 
| Intrusion Prevention | IPS-Lab Profile (32 signatures) | ✅ | 
| Web Filtering | WebFilter-Lab Profile |  | 
| SNMP Monitoring | SNMPv2c Community | ✅ | 
| Firewall Policies | 9 rules with NAT & IPS | ✅ | 
| Attack/Defense Scenarios | Port scan, VLAN isolation tests | ✅ | 
⚠️ Web Filtering requires active FortiGuard license. Configuration documented but FortiGuard categories unavailable in evaluation environment.
                        ┌─────────────┐
                        │   Internet  │
                        └──────┬──────┘
                               │
                        ┌──────┴──────┐
                        │  R3-Siege   │  WAN Router
                        │ 10.0.0.1/30 │
                        │ 10.0.0.5/30 │
                        └──┬──────┬───┘
                           │      │
              WAN1          │      │         WAN2
         10.0.0.0/30       │      │    10.0.0.4/30
                    ┌──────┘      └──────┐
                    │                    │
             ┌──────┴──────┐    ┌────────┴──────┐
             │ FGT1-Master │◄──►│  FGT2-Backup  │
             │  Priority:  │ HA │   Priority:   │
             │     200     │    │      100      │
             └──────┬──────┘    └───────┬───────┘
                    │   Trunk dot1q      │
                    └────────┬───────────┘
                             │
                      ┌──────┴──────┐
                      │   SW-Core   │  Cisco vIOS-L2
                      │  L2 Switch  │
                      └──┬──┬──┬──┬─┘
                         │  │  │  │
               ┌─────────┘  │  │  └─────────┐
               │            │  │             │
          VLAN10        VLAN30  VLAN50    VLAN100/200
          Users         Guest    DMZ       WAN Sites
        192.168.10     .30.0   .50.0    ──► R4, R5
          PC7,PC8      PC10     PC9
| Device | Interface | Connected To | Interface | 
|---|---|---|---|
| R3 | Gi0/0 | FGT1 | port1 | 
| R3 | Gi0/1 | FGT2 | port1 | 
| R3 | Gi0/2 | SW | Gi2/0 (VLAN100) | 
| R3 | Gi0/3 | SW | Gi2/1 (VLAN200) | 
| FGT1 | port2 | SW | Gi0/0 (trunk) | 
| FGT2 | port3 | SW | Gi0/1 (trunk) | 
| FGT1 | port4 | FGT2 | port4 (HA) | 
| FGT1 | port3 | Cloud0 | pnet0 (Mgmt) | 
| SW | Gi0/2 | R4 | Gi0/0 (VLAN100) | 
| SW | Gi0/3 | R5 | Gi0/0 (VLAN200) | 
| SW | Gi1/0 | PC7 | eth0 (VLAN10) | 
| SW | Gi1/1 | PC8 | eth0 (VLAN10) | 
| SW | Gi1/2 | PC9 | eth0 (VLAN50) | 
| SW | Gi1/3 | PC10 | eth0 (VLAN30) | 
| Device | Model | Role | 
|---|---|---|
| FGT1 | FortiGate-VM64-KVM | Primary Firewall (Active) | 
| FGT2 | FortiGate-VM64-KVM | Backup Firewall (Passive) | 
| R3 | Cisco vIOS 15.9-3.M6 | WAN Router / Siege | 
| R4 | Cisco vIOS 15.9-3.M6 | Remote Site A Router | 
| R5 | Cisco vIOS 15.9-3.M6 | Remote Site B Router | 
| SW | Cisco vIOS-L2 | Core L2 Switch | 
| PC7-PC10 | VPCS | End-user Clients | 
| Tool | Version | Purpose | 
|---|---|---|
| EVE-NG Community | 6.2.0-4 | Network Emulation Platform | 
| VirtualBox | 6.x | Hypervisor (Windows Host) | 
| FortiOS | v7.0.9 build0444 | Firewall OS | 
| PuTTY | Latest | SSH/Console Access | 
| Nmap | 7.95 | Network Scanning Tests | 
| Link | Network | R3 IP | FGT IP | 
|---|---|---|---|
| WAN1 (FGT1) | 10.0.0.0/30 | 10.0.0.1 | 10.0.0.2 | 
| WAN2 (FGT2) | 10.0.0.4/30 | 10.0.0.5 | 10.0.0.6 | 
| Link | Network | R3 IP | Router IP | 
|---|---|---|---|
| To Site A (R4) | 10.1.0.0/24 | 10.1.0.1 | 10.1.0.2 | 
| To Site B (R5) | 10.2.0.0/24 | 10.2.0.1 | 10.2.0.2 | 
| VLAN | Name | Network | Gateway | DHCP Pool | 
|---|---|---|---|---|
| 10 | Users | 192.168.10.0/24 | 192.168.10.1 | .100 – .200 | 
| 20 | Servers | 192.168.20.0/24 | 192.168.20.1 | .100 – .200 | 
| 30 | Guest | 192.168.30.0/24 | 192.168.30.1 | .100 – .200 | 
| 50 | DMZ | 192.168.50.0/24 | 192.168.50.1 | .100 – .200 | 
| 100 | WAN-SiteA | 10.1.0.0/24 | — | — | 
| 200 | WAN-SiteB | 10.2.0.0/24 | — | — | 
| Site | Network | Gateway | 
|---|---|---|
| Site A (R4) | 172.16.10.0/24 | 172.16.10.1 | 
| Site B (R5) | 172.16.20.0/24 | 172.16.20.1 | 
| Type | Range | 
|---|---|
| SSL VPN Pool | 10.10.10.10 – 10.10.10.100 | 
| Management | 192.168.56.200/24 | 
- Windows 10/11 host machine (minimum 20GB RAM recommended)
- VirtualBox 6.x or higher
- EVE-NG Community Edition OVA
- FortiGate QEMU image (fortinet-FGT-v7.0.9)
- Cisco vIOS image (vios-adventerprisek9-m.SPA.159-3.M6)
- Cisco vIOS-L2 image (viosl2-adventerprisek9-m.ssa.high_iron_20200929)
# 1. Import EVE-NG OVA into VirtualBox
# Network Adapter 1: NAT
# Network Adapter 2: Host-Only (Promiscuous Mode: Allow All)  ← CRITICAL
# 2. Access EVE-NG
Web UI:  http://192.168.56.110  (admin/eve)
SSH:     ssh root@192.168.56.110  (root/eve)# SSH into EVE-NG
ssh root@192.168.56.110
# Upload and install FortiGate image
cd /opt/unetlab/addons/qemu/
mkdir fortinet-FGT-v7.0.9
# Upload virtioa.qcow2 to this directory
# Fix permissions
/opt/unetlab/wrappers/unl_wrapper -a fixpermissions
⚠️ Critical: Set Promiscuous Mode to Allow All on the Host-Only adapter of the EVE-NG VM. Without this, management traffic to FortiGate (192.168.56.200) will not reach the host.
FGT1-Master:
config system ha
    set group-name "FGT-HA"
    set mode a-p
    set password "ha-secret"
    set priority 200
    set hbdev "port4" 50
end
FGT2-Backup:
config system ha
    set group-name "FGT-HA"
    set mode a-p
    set password "ha-secret"
    set priority 100
    set hbdev "port4" 50
end
config vpn ipsec phase1-interface
    edit "VPN-SiteA"
        set interface "port1"
        set ike-version 1
        set peertype any
        set proposal des-sha256
        set dhgrp 14
        set remote-gw 10.1.0.2
        set psksecret "VPN-SiteA-Secret123"
    next
end
config vpn ssl settings
    set servercert "Fortinet_Factory"
    set tunnel-ip-pools "SSLVPN_TUNNEL_ADDR1"
    set source-interface "port3"
    set source-address "all"
    set default-portal "full-access"
    set port 10443
    set port-precedence disable
end
config ips sensor
    edit "IPS-Lab"
        set comment "IPS profile for enterprise lab"
        config entries
            edit 1
                set rule SCAN
                set action block
                set log enable
                set packet-log enable
            next
        end
        set block-malicious-url enable
    next
end
config system snmp sysinfo
