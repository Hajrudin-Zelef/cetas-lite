---
id: collect-261001-cisco/cisco/ayoub-lah-enterprise-network-security-lab-4e893e59-2
title: "1. Import EVE-NG OVA into VirtualBox"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/collect-261001-cisco/ayoub-lah-enterprise-network-security-lab-4e893e59.md
source_anchor: ""
source_lines: [181, 313]
sha256: eaa856485d1757984347ef349057d9e6f3b81e0ae2381b38beb4c0e1db473e59
---

# 1. Import EVE-NG OVA into VirtualBox

    set status enable
    set description "FGT1-Master Enterprise Lab"
    set contact-info "admin@lab.local"
    set location "EVE-NG Lab"
end
config system snmp community
    edit 1
        set name "LabSNMP"
        set status enable
        config hosts
            edit 1
                set ip 192.168.56.0 255.255.255.0
            next
        end
        set events cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down
    next
end
| ID | Name | Source | Destination | Service | Action | NAT | IPS | 
|---|---|---|---|---|---|---|---|
| 1 | Users-to-Internet | vlan10 | port1 | ALL | ACCEPT | ✅ | IPS-Lab | 
| 2 | Users-to-DMZ | vlan10 | vlan50 | HTTP/HTTPS/PING | ACCEPT | ❌ | IPS-Lab | 
| 3 | DMZ-block-LAN | vlan50 | vlan10 | ALL | DENY | ❌ | ❌ | 
| 4 | Guest-Internet-only | vlan30 | port1 | ALL | ACCEPT | ✅ | ❌ | 
| 5 | LAN-to-SiteA | vlan10 | VPN-SiteA | ALL | ACCEPT | ❌ | ❌ | 
| 6 | SiteA-to-LAN | VPN-SiteA | vlan10 | ALL | ACCEPT | ❌ | ❌ | 
| 7 | LAN-to-SiteB | vlan10 | VPN-SiteB | ALL | ACCEPT | ❌ | ❌ | 
| 8 | SiteB-to-LAN | VPN-SiteB | vlan10 | ALL | ACCEPT | ❌ | ❌ | 
| 9 | SSLVPN-to-LAN | ssl.root | vlan10 | ALL | ACCEPT | ❌ | IPS-Lab | 
# Test 1 — DHCP Verification (from PC7 console in EVE-NG)
VPCS> dhcp
DDORA IP 192.168.10.100/24 GW 192.168.10.1  ✅
# Test 2 — Gateway Ping
VPCS> ping 192.168.10.1
84 bytes from 192.168.10.1 icmp_seq=1 ttl=255 time=14.532 ms  ✅
# Test 3 — IPsec VPN to Site A
VPCS> ping 172.16.10.1
84 bytes from 172.16.10.1 icmp_seq=1 ttl=253 time=25.1 ms  ✅
# Test 4 — IPsec VPN to Site B
VPCS> ping 172.16.20.1
84 bytes from 172.16.20.1 icmp_seq=1 ttl=253 time=22.4 ms  ✅# Test 5 — VLAN Isolation (PC7 → PC10, should be BLOCKED)
VPCS> ping 192.168.30.100
192.168.30.100 icmp_seq=1 timeout  ✅ BLOCKED
# Test 6 — DMZ to LAN (PC9 → PC7, should be BLOCKED)
VPCS> ping 192.168.10.100
192.168.10.100 icmp_seq=1 timeout  ✅ BLOCKED
# Test 7 — Port Scan (from Windows host via nmap)
nmap -sS -Pn -p 1-1000 192.168.56.200
PORT    STATE  SERVICE
22/tcp  open   ssh
80/tcp  open   http
443/tcp open   https  ✅
# Test 8 — IPS Port Filtering
nmap -sT -Pn --max-rtt-timeout 100ms -p 22,23,80,443 192.168.10.100
PORT    STATE    SERVICE
22/tcp  filtered ssh
80/tcp  filtered http    ✅ Filtered by IPS-Lab
# Test 9 — SNMP (from Windows host)
nmap -sU -p 161 --script snmp-info --script-args snmpcommunity=LabSNMP 192.168.56.200
161/udp open|filtered snmp  ✅# Verify HA Status on FGT1
FGT1-Master # get system ha status
Model: FortiGate-VM64-KVM
Mode: a-p
Group: 0
Debug: 0
ses_pickup: disable, ses_pickup_delay=disable
Master:
  FGT1-Master(0) HA cluster index = 0, priority = 200, override = disable
Slave:
  FGT2-Backup(1) HA cluster index = 1, priority = 100, override = disable
number of vcluster: 1
vcluster 1: work 192.168.56.200# Create SSH tunnel (Windows CMD)
ssh -L 8443:192.168.56.200:10443 root@192.168.56.110
# Access SSL VPN portal
https://localhost:8443
# Login: vpnuser / VPNuser123!  ✅
| Test | Description | Expected | Result | 
|---|---|---|---|
| 01 | DHCP VLAN10 PC7 | 192.168.10.100 | ✅ | 
| 02 | DHCP VLAN10 PC8 | 192.168.10.101 | ✅ | 
| 03 | DHCP VLAN30 PC10 | 192.168.30.100 | ✅ | 
| 04 | DHCP VLAN50 PC9 | 192.168.50.100 | ✅ | 
| 05 | PC7 → Gateway ping | OK | ✅ | 
| 06 | PC7 → PC10 VLAN isolation | BLOCKED | ✅ | 
| 07 | PC7 → PC9 DMZ access | ALLOWED | ✅ | 
| 08 | PC9 → PC7 DMZ→LAN block | BLOCKED | ✅ | 
| 09 | FortiGate HA sync | In-Sync | ✅ | 
| 10 | R3 → FGT1 ping | OK | ✅ | 
| 11 | IPsec VPN Site A | Tunnel active | ✅ | 
| 12 | IPsec VPN Site B | Tunnel active | ✅ | 
| 13 | PC7 → Site A (172.16.10.1) | OK via VPN | ✅ | 
| 14 | PC7 → Site B (172.16.20.1) | OK via VPN | ✅ | 
| 15 | SSL VPN Web Mode | Portal accessible | ✅ | 
| 16 | IPS Port Scan detection | Ports filtered | ✅ | 
| 17 | SNMP port 161 | Open | ✅ | 
| Metric | Value | 
|---|---|
| FortiGate vCPUs | 1/1 (100%) | 
| Allocated RAM | 2GB/2GB (98%) | 
| Firmware | v7.0.9 build0444 (Mature) | 
| HA Sync Status | In-Sync ✅ | 
| Active IPsec Tunnels | 2 (SiteA + SiteB) | 
| Firewall Policies | 9 active rules | 
| IPS Signatures | 32 (SCAN category) | 
⚠️ These credentials are for lab/educational purposes only. Never use these in production.
| Device | Username | Password | 
|---|---|---|
| EVE-NG SSH | root | eve | 
| EVE-NG Web | admin | eve | 
| FGT1-Master | admin | Admin1234! | 
| FGT2-Backup | admin | Admin1234! | 
| SSL VPN User | vpnuser | VPNuser123! | 
- 
Web Filtering — FortiGuard category-based filtering requires an active FortiGuard subscription. In this lab, the evaluation license expired after 15 days, limiting FortiGuard cloud services.
- 
SSL VPN Browser Access — FortiGate v7.0.9 KVM uses older TLS/DH parameters incompatible with modern browsers (Chrome, Firefox, Edge). Workaround: SSH tunnel ( ssh -L 8443:192.168.56.200:10443 root@192.168.56.110 ).
- 
VPCS Limitations — VPCS clients cannot initiate inter-VLAN traffic tests directly in some scenarios. Tests validated via FortiGate CLI pings.
- 
FortiClient Compatibility — FortiClient 7.4.x is not compatible with FortiGate 7.0.x for SSL VPN tunnel mode. Web mode via SSH tunnel was used as an alternative.
- ✅ Network Design — Enterprise topology with redundancy
- ✅ Firewall Administration — FortiGate NGFW configuration
- ✅ High Availability — Active-Passive HA cluster setup
- ✅ VPN Engineering — IPsec site-to-site + SSL remote access
- ✅ Network Segmentation — VLAN design and inter-VLAN routing
- ✅ Security Policies — Stateful firewall rules with UTM profiles
- ✅ Intrusion Prevention — IPS signature-based detection/blocking
- ✅ Network Monitoring — SNMPv2c configuration and testing
- ✅ Troubleshooting — Systematic diagnosis with packet captures
- ✅ Documentation — Technical reports, diagrams, and GitHub
Ayoub Lahlaibi
Master's Student — IT Security & Big Data (SIT & Big Data) Faculté des Sciences et Techniques de Tanger Université Abdelmalek Essaâdi, Morocco
⭐ If this project was helpful, please consider giving it a star! ⭐
Built with ❤️ for learning and portfolio purposes
