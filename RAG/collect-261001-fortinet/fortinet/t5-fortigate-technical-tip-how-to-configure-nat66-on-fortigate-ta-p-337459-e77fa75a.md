---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-how-to-configure-nat66-on-fortigate-ta-p-337459-e77fa75a
title: "t5-fortigate-technical-tip-how-to-configure-nat66-on-fortigate-ta-p-337459-e77fa75a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-how-to-configure-nat66-on-fortigate-ta-p-337459-e77fa75a.md
source_anchor: ""
source_lines: [1, 4]
sha256: 23aeb8d767c167602d7b5b6d857357088e96b67ea0859ea00e63dcf6eefa2d7d
---

# t5-fortigate-technical-tip-how-to-configure-nat66-on-fortigate-ta-p-337459-e77fa75a

Technical Tip: How to configure NAT66 on FortiGate
| Description | This article describes the steps to configure NAT66 on a FortiGate device, including the necessary firewall policies and configuration steps along with troubleshooting commands. | 
| Scope | FortiGate. | 
| Solution | NAT66 (Network Address Translation for IPv6) allows the translation of one IPv6 address to another, similar to how NAT is implemented for IPv4 (NAT44) using VIP. This can be useful when you need to translate internal IPv6 addresses to external IPv6 addresses in specific scenarios, such as network security, load balancing, or to meet routing requirements.  Prerequisites:  Step 1: Create the IPv6 VIP:   GUI:    CLI:  config firewall vip6 edit "example-vip6" set extip 2a02:xx::xx set mappedip 2001:xx::xx set nat66 enable end  Step 2: Apply the IPv6 VIP in a Firewall Policy:    GUI:    CLI:  config firewall policy edit 'ID' set name "NAT66" set srcintf "Internet_WAN" set dstintf "DMZ" set action accept set srcaddr6 "srcaddr6" set dstaddr6 "example-vip6" set schedule "always" set service "ALL" set nat enable next end  Troubleshooting commands:  Routing & Neighbor solicitation list commands:  get sys status get router info6 routing-table connected get router info6 routing-table static diagnose ipv6 neighbor-cache list diagnose ipv6 addr list get router info6 kernel  IPv6 traffic debug Commands:  diagnose debug reset diagnose debug flow filter6 clear diagnose debug flow filter6 addr xxxx::xx <----- Replace xxxx::xx with users source IPv6 address. diagnose debug flow show function-name enable diagnose debug flow trace start6 1000 diagnose debug enable  To stop the debug:  diagnose debug disable  IPv6 session list:  diagnose sys session6 filter clear diagnose sys session6 filter src xxxx::xx <----- Replace xxxx::xx with users source IPv6 address. diagnose sys session6 list  Sniffer Commands:  diagnose sniffer packet any "host xxxx::xx" 6 0 l <----- Replace xxxx::xx with users source IPv6 address. Related article: Technical Tip: How to use debug flow and sniffer to capture IPv6 traffic |
