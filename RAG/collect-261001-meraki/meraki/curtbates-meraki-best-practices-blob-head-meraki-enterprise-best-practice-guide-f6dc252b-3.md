---
id: collect-261001-meraki/meraki/curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide-f6dc252b-3
title: "curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide--f6dc252b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput", "voice"]
source: docs/RAG/collect-261001-meraki/curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide--f6dc252b.md
source_anchor: ""
source_lines: [116, 152]
sha256: dc0d90e80ad24e5b46f9924b898eeba58fd36b57c5c2fd23576d2262f82de15c
---

# curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide--f6dc252b

| MS (switching) | Stacking (preferred) with cross-stack LACP uplinks on top+bottom switches; warm spare (VRRP) for L3 gateways where stacking isn't used. | 
| MR (wireless) | Bridge mode + common VLAN per roaming domain for seamless failover; Campus Gateway for large multi-subnet roaming. | 
| WAN/SD-WAN | Auto VPN auto-builds tunnels on all reachable uplinks for transparent path failover. | 
Foundation
- Org boundary chosen around the intended admin + Auto VPN domain
- VLAN plan: employee / voice / guest / IoT / management, no overlaps
- Standard /24 (or /23) user subnets; dedicated management VLAN
- Firmware standardized across stacks and HA pairs
Wired (MS)
- RSTP on; root forced (priority 0) on a stable core switch; diameter < 7
- BPDU Guard on access ports; Loop/Root Guard on appropriate trunks
- Trunks pruned; native/allowed VLANs matched both ends; LACP pre-configured in dashboard
- UDLD alert-only on fiber; auto-negotiation on Meraki links
- L3 interfaces on data VLANs only; mgmt subnet ≠ any L3 subnet (non-MS390 stacks)
- OSPF: Area 0 attached, MD5 auth, transit-VLAN design, sensible passive interfaces
- Allowed DHCP servers set; dashboard permitted in ACLs (≤128 ACEs)
- ≤400 switches / <8000 ports per network; Cat-6a for multi-gig
Wireless (MR)
- Sized by capacity (throughput and client count), ~25/radio, ~50/AP
- Active survey + spectrum analysis, ≥25 dB SNR both bands
- ≤3 SSIDs (≤5 max), one per auth type; bridge mode (not NAT) for voice
- RF profiles per area: 20 MHz width, min bitrate ≥12 Mbps, DFS on, client balancing on
- RX-SOP tuned and tested; 802.11k/i/OKC on; 802.11r on for voice
- Voice/video shaping rule set to ignore per-client limit
SD-WAN (MX)
- Uplink bandwidth limits match ISP rates; multiple test IPs; hourly list updates
- Load balancing / flow preferences configured (e.g., guest → secondary WAN)
- Performance policies per app class; default traffic-shaping rules enabled
- Auto VPN hub-and-spoke with HQ/DC as hubs
- Client VPN via Systems Manager policy (+ split tunnel where appropriate)
Security
- Correct MX mode (NAT vs concentrator) for the topology
- Inter-VLAN firewall rules (guest isolated); ordered top-down
- Granular L7 rules; country blocks only when justified
- Port-forwarding/NAT scoped narrowly, no "Any" remote IPs
- AMP + IDS/IPS (Prevention) enabled; spoofing protection set to Block
Reference: Cisco Meraki Best Practice Design — General MX Best Practices, General MS Best Practices, and High-Density Wi-Fi Deployments (documentation.meraki.com). Confirm model-specific behavior, licensing, and firmware caveats against current Meraki documentation and your Meraki SE.
