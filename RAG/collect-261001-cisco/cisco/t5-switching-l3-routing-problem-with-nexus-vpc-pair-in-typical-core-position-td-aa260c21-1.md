---
id: collect-261001-cisco/cisco/t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td-aa260c21-1
title: "t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td--aa260c21"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td--aa260c21.md
source_anchor: ""
source_lines: [1, 46]
sha256: 1b35c9f4b3208907ce25eec2d360d602346c52e4b6a54cf7312b0f940b21c26e
---

# t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td--aa260c21

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-18-2022 02:45 PM
This is a standard configuration, with 2 nexus 9k running VPC between them, and catalyst access switch trunked at L2 redundantly to the two Nexus. All L3 SVI lives on Nexus core, and enumerated with HSRP on each core box.
N9K-1------|
|| CAT3k-1 (vlans 100/200)
N9K-2------|
CAT3k switch is port-channeled to the 2 x N9K, and trunked for vlans 100/200. VPC peer link carries vlans 100/200 and runs on the same native vlan as the CAT switch port channel native vlan.
I noticed an odd problem in that not all traceroutes go through primary HSRP interfaces (N9K-1) when I go back and forth between vlan100 and 200. So essentially N9K-1 has higher HSRP and VPC priority but traffic goes through N9K-2 when routing to vlan 100:
N9K-1:
int vlan100
ip addr 192.168.100.2
hsrp ver 2
hsrp 1
preempt
priority 150
ip addr 192.168.100.1
N9K-2:
int vlan100
ip addr 192.168.100.3
hsrp ver 2
hsrp 1
preempt
priority 125
ip addr 192.168.100.1
Here is the VPC config:
vpc domain 1
peer-switch
role 200 / 300 (primary / secondary)
peer-keepalive dest 1.1.1.1 source 1.1.1.2 vrf vpckeepalive
peer-gateway
fast-convergence
ip arp synchronize
When I traceroute between vlans 100 and 200, sometimes the default gateway of N9K-2 responds, even though HSRP and VPC domain priority is higher. When I shut one one of the HSRP interfaces (100 or 200) on only one of the N9K's, there is weird on/off connectivity between the subnets. So I suspect traffic is coming on on the CAT3k uplink to N9K1 for example, but N9K1's vlan100 interface is down. However, the VPC peer link is not carrying the traffic over to the vlan100 interface of the N9K2. There is no routing between N9K1 and 2 over the VPC peer link. All vlans are L2 with the L3 as SVI on each N9K. Vlans 100/200 are on the VPC peer link.
Anyone have a clue? I suspect its related to the peer-gateway or peer-switch command. Essentially to make it work, I have to shut all SVI's on one or the other N9K. My port-channel hash algorithm is ip-src-dst load balance.
Solved! Go to Solution.
- Labels:
- 
						
							
		
