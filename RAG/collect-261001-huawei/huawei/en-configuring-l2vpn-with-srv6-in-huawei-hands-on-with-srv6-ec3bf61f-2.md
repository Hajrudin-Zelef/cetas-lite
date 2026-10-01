---
id: collect-261001-huawei/huawei/en-configuring-l2vpn-with-srv6-in-huawei-hands-on-with-srv6-ec3bf61f-2
title: "en-configuring-l2vpn-with-srv6-in-huawei-hands-on-with-srv6-ec3bf61f"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/en-configuring-l2vpn-with-srv6-in-huawei-hands-on-with-srv6-ec3bf61f.md
source_anchor: ""
source_lines: [220, 234]
sha256: fdcb5de7ea495ac1e0148c4b4edf550e2dbd098200a438a775ac801a659ec205
---

# en-configuring-l2vpn-with-srv6-in-huawei-hands-on-with-srv6-ec3bf61f

Interfaces of CEs 1 and 2, with an IPv4 address for communication.
ARP table from CE1 and a ping to CE2, confirming connectivity between the CEs.
Neighbors LLDP in CE1, showing the path being “Transparent” from the ECs point of view.
8.7: While CE1 is exchanging “pings” with CE2, a packet capture on the R1 router’s interface for traffic destined for the rest of the SRv6 network produces the following output:
EVPN packets are encapsulated, and when they are forwarded to the SRv6 network, the “destination address” becomes 2001:db8:6:6::A, which is the “SID” END.DX2.  
 
The most interesting thing about SRv6 is that the package is “IPv6”. If our environment contained only SRv6-supporting PEs, the rest of the network would know how to forward traffic without any problems  
Summary
In this lab, we explore the configuration of an L2VPN tunnel using SRv6, showing the basic configurations of each device.  
SRv6, with its advanced capabilities, is proving to be a promising technology for the networks of the future. To learn more about implementing SRv6 in your network, turn to Made4it, an expert in the field. We can assist you in the migration process from MPLS to SRv6 by ensuring the coexistence of these protocols.    
Complete configurations
Download this complete lab, with topologies, configurations and roadmap here:  
Authors:  
Gabriel Henrique, Network and Project Analyst at Made4it. 
Rafael Ganascim, Co-Founder of Made4it.
