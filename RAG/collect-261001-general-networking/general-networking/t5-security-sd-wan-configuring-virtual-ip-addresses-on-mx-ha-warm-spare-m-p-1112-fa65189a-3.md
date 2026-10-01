---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a-3
title: "t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a.md
source_anchor: ""
source_lines: [175, 200]
sha256: d503cee74b35aee3dc8a2a0ee92d291fdad24d54cef0e78c871d44aedb38e603
---

# t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a

- LAN interfaces only have a single IP address that is configured on the primary MX. This IP address is passed to the standby MX when failover occurs (so you only need one IP address on the LAN for two MXs)
- VRRP messages only occur between the primary and secondary MX appliances on LAN interfaces - but they occur on all configured VLANs (be aware of this if you have multiple connections between the MXs)
- VRRP messages occur at Layer 2, hence why you don't need a separate IP address on each of the MXs on the LAN side (this is confusing for anyone who has used HSRP)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-09-2021 08:03 AM
Thank you!. Now I get it.
LAN interfaces only have a single IP address that is configured on the primary MX. This IP address is passed to the standby MX when failover occurs (so you only need one IP address on the LAN for two MXs)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-20-2025 09:37 AM
I have a pair of MX68 in HA mode with two WAN connections to ISP routers with each router with a dedicated /28 public subnets. VRRP address configured on the MX68. We a
have a MR44 connected to port 11.
If WAN 1 is disconnected traffic doesn't failover to WAN 2. The WAP goes offline. The WAP is using the VRRP address as its public ip address. A packet capture shows any traffic with source of the VRRP failing.
When I change the setting to use physical IP address. The WAP comes online. Traffic fails over to WAN 2 with no issue.
Is is this normal behaviour or are firewall rules needed to use the VRRP address?
