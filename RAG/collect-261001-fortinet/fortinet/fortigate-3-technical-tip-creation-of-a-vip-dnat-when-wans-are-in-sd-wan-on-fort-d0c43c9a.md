---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-creation-of-a-vip-dnat-when-wans-are-in-sd-wan-on-fort-d0c43c9a
title: "fortigate-3-technical-tip-creation-of-a-vip-dnat-when-wans-are-in-sd-wan-on-fort-d0c43c9a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-creation-of-a-vip-dnat-when-wans-are-in-sd-wan-on-fort-d0c43c9a.md
source_anchor: ""
source_lines: [1, 4]
sha256: 77184811f3d15fa0ecb34886789c13823c719f7af24c1f219dd216fe84c6b36f
---

# fortigate-3-technical-tip-creation-of-a-vip-dnat-when-wans-are-in-sd-wan-on-fort-d0c43c9a

Technical Tip: Creation of a VIP (DNAT) when WANs are in SD-WAN on FortiGate
| Description | This article describes how to avoid a scenario where failure occurs because the VIP is bound to the incorrect interface and there is improper routing of reply traffic. | 
| Scope | FortiGate. | 
| Solution | SD-WAN only determines the outgoing path; for incoming traffic, bind the VIP to the physical WAN(s), or use extintf any (not to the SD-WAN zone).  Option A: A single WAN (member of SD-WAN).       Example:  config firewall vip edit "vip_web_wan1" set extip 203.0.113.10 set mappedip "10.10.10.10" set extintf "wan1" set portforward enable set extport 443 set mappedport 443 next end   config firewall policy edit 100 set name "in_web_wan1" set srcintf "wan1" set dstintf "lan" set srcaddr "all" set dstaddr "vip_web_wan1" set action accept set schedule "always" set service "HTTPS" set nat disable next end  Option B: Two WANs. Create a VIP for each WAN (each with its own public IP) and set up two identical policies by changing the Incoming interface and the Destination (VIP of WAN1 / VIP of WAN2). SD-WAN does not perform inbound load balancing; if it is required public high availability, utilize two DNS records (A/AAAA) with a low TTL or a failover mechanism in the DNS/ISP.  Example:  config firewall vip edit "vip_web_wan2" set extip 198.51.100.20 set mappedip "10.10.10.10" set extintf "wan2" set portforward enable set extport 443 set mappedport 443 next end   config firewall policy edit 101 set name "in_web_wan2" set srcintf "wan2" set dstintf "lan" set srcaddr "all" set dstaddr "vip_web_wan2" set action accept set schedule "always" set service "HTTPS" set nat disable next end  Notes:   Troubleshooting. diagnose sniffer packet any 'host <public_IP> and port <port>' 4 0 l    Related documents: |
