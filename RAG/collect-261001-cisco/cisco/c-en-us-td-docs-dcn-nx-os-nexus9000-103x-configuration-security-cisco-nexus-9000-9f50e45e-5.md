---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e-5
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e.md
source_anchor: ""
source_lines: [391, 472]
sha256: 1258f6717cfb5c7cc140d9f7b08a879b3656a8c1790b92943ef78b788ca4676e
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e

                                 When the egress RACL is beyond 4K, the TCAM carving configuration has to be ingress RACL (RACL) + egress RACL (e-racl) summing to 20480. See the following TCAM carving example: hardware access-list tcam region ifacl 0
hardware access-list tcam region ipv6-ifacl 0
hardware access-list tcam region mac-ifacl 0
hardware access-list tcam region racl 0
hardware access-list tcam region ipv6-racl 0
hardware access-list tcam region span 0
hardware access-list tcam region redirect_v4 0
hardware access-list tcam region redirect_v6 0
hardware access-list tcam region e-racl 20480
- 
                                 				
                                 You can partially use IPv6 RACL with IPv6 IFCAL. This is applicable to Cisco Nexus N9K-C9508 and N9K-C9504 with N9K-X96136YC-R, N9K-X9636C-R, N9K-X9636Q-R, and N9K-X9636C-RX line cards.
- 
                                 				
                                 The N9K-X9636C-R and N9K-X9636Q-R line cards support a maximum TCAM region size of 12K. If you configure a greater number, the TCAM region is set to 12K.
- 
                                 				
                                 The N9K-X96136YC-R and N9K-X9636C-R line cards support egress RACL of 2K.
- 
                                 				
                                 The N9K-X9636C-RX line card supports a TCAM region size beyond 12K. If you configure the RACL IPv4 TCAM region to 100K, the TCAM region is set to 12K for the N9K-X9636C-R and N9K-X9636Q-R line cards and to 100K for the N9K-X9636C-RX line card, provided you have set all of the other TCAM regions and made space for the N9K-X9636C-R and N9K-X9636Q-R line cards to accommodate 12K.
- 
                                 				
                                 Beginning with Cisco NX-OS Release 10.2(2)F, The N9K-X9636C-R and N9K-X9636Q-R line cards support a maximum TCAM region size of 20K. If you configure a greater number, the TCAM region is re-set to 20K.
- 
                                 				
                                 In addition to the internal TCAM, an external TCAM of 128K is available on the N9K-X9636C-RX line card.
The following table summarizes the regions that need to be configured for a given feature to work. The region sizes should be selected based on the scale requirements of a given feature.
| Table 2. Features per ACL                                     		TCAM Region |  | 
|---|---|
| Feature Name | Region Name | 
|---|---|
| Port ACL | ifacl: For IPv4 port ACLs ifacl-udf: For UDFs on IPv4 port ACLs ing-ifacl: For ingress IPv4, IPv6, and MAC port ACLs ing-ifacl: For ingress IPv4, IPv6, MAC port ACLs, and MAC port ACLs with UDF ipv6-ifacl: For IPv6 port ACLs mac-ifacl: For MAC port ACLs | 
| Port QoS (QoS classification policy applied on Layer 2 ports or port channels) | qos, qos-lite, rp-qos, rp-qos-lite, ns-qos, e-qos, or e-qos-lite: For classifying IPv4 packets ing-l2-qos: For classifying ingress Layer 2 packets ipv6-qos, rp-ipv6-qos, ns-ipv6-qos, or e-ipv6-qos: For classifying IPv6 packets mac-qos, rp-mac-qos, ns-mac-qos, or e-mac-qos: For classifying non-IP packets | 
| VACL | vacl: For IPv4 packets ipv6-vacl: For IPv6 packets mac-vacl: For non-IP packets | 
| VLAN QoS (QoS classification policy applied on a VLAN) | vqos or ns-vqos: For classifying IPv4 packets ipv6-vqos or ns-ipv6-vqos: For classifying IPv6 packets ing-l3-vlan-qos: For classifying ingress Layer 3, VLAN, and SVI QoS packets mac-vqos or ns-mac-vqos: For classifying non-IP packets | 
| RACL | egr-racl: For egress IPv4 and IPv6 RACLs e-racl: For egress IPv4 RACLs e-ipv6-racl: For egress IPv6 RACLs ing-racl: For ingress IPv4 and IPv6 RACLs racl: For IPv4 RACLs racl-lite: For IPv4 RACLs racl-udf: For UDFs on IPv4 RACLs ipv6-racl: For IPv6 RACLs | 
| Layer 3 QoS (QoS classification policy applied on Layer 3 ports or port channels) | l3qos, l3qos-lite, or ns-l3qos: For classifying IPv4 packets ipv6-l3qos or ns-ipv6-l3qos: For classifying IPv6 packets | 
| VLAN source or VLAN filter SPAN (for Cisco Nexus 9500 or 9300 Series switches) Rx SPAN on 40G ports (for Cisco Nexus 9300 Series switches only) | span | 
| SPAN filters | ifacl: For filtering IPv4 traffic on Layer 2 (switch port) source interfaces. ifacl-udf: For UDFs on IPv4 port ACLs ipv6-ifacl: For filtering IPv6 traffic on Layer 2 (switch port) source interfaces. mac-ifacl: For filtering Layer 2 traffic on Layer 2 (switch port) source interfaces. racl-udf: For UDFs on IPv4 RACLs vacl: For filtering IPv4 traffic on VLAN sources. ipv6-vacl: For filtering IPv6 traffic on VLAN sources. mac-vacl: For filtering Layer 2 traffic on VLAN sources. racl: For filtering IPv4 traffic on Layer 3 interfaces. ipv6-racl: For filtering IPv6 traffic on Layer 3 interfaces. ing-l2-span-filter: For filtering ingress Layer 2 SPAN traffic ing-l3-span-filter: For filtering ingress Layer 3 and VLAN SPAN traffic | 
| SVI counters | svi | 
| BFD, DHCP relay, or DHCPv6 relay | redirect | 
| CoPP | copp | 
| System-managed ACLs | system | 
| vPC convergence | vpc-convergence | 
| Fabric extender (FEX) | fex-ifacl, fex-ipv6-ifacl, fex-ipv6-qos, fex-mac-ifacl, fex-mac-qos, fex-qos, fex-qos-lite | 
| Dynamic ARP inspection (DAI) | arp-ether | 
| IP source guard (IPSG) | ipsg | 
| Multicast PIM Bidir | mcast_bidir | 
| Static MPLS | mpls | 
| Network address translation (NAT) | nat | 
| NetFlow | ing-netflow | 
| OpenFlow | openflow | 
| sFlow | sflow | 
| Supervisor modules | egr-sup: Egress supervisor ing-sup: Ingress supervisor | 
| Policy-Based Routing (PBR) | ing-racl: For matching ingress L3 traffic for PBR. | 
| Layer 2 Intelligent Traffic Director (ITD) | vacl: Programs L2 redirect ACLs at the VLAN level. | 
| Layer 3 Intelligent Traffic Director (ITD) | ing-racl: Programs L3 redirect ACLs for ITD. | 
| Enhanced Policy-Based Redirect at L2 (ePBR) | ing-ifacl: Programs L2 redirect ACLs for ePBR L2. | 
| Enhanced Policy-Based Redirect at L3 (ePBR) | ing-racl: Programs L3 redirect ACLs for ePBR L3. | 
| Note |  | 
| Note |  | 
| Note |  | 
| Note |  | 
| Note |  | 
| Note |  | 
| Note |  | 
| Note |  | 
| Note |  | 
Maximum Label Sizes Supported for ACL Types
Cisco NX-OS switches support the following label sizes for the corresponding ACL types:
| Table 3. ACL Types and Maximum Label Sizes |  |  |  | 
|---|---|---|---|
| ACL Types | Direction | Label | Label Type | 
|---|---|---|---|
| RACL/PBR/VACL/ L3-VLAN QoS/L3-VLAN SPAN ACL | Ingress | 62 | BD | 
| PACL/L2 QoS/L2 SPAN ACL | Ingress |  | IF | 
| RACL/VACL/L3-VLAN QoS | Egress | 254 | BD | 
| L2 QoS | Egress | 31 | IF | 
The label size can be increased to 62 when you enter the hardware access-list tcam label ing-ifacl 6 command and reload the switch.
Beginning with Cisco NX-OS Release 9.3(6), the hardware access-list tcam label ing-ifacl 6 command is introduced and is applicable only for Cisco Nexus 9300-FX platform switches.
Beginning with Cisco NX-OS Release 10.1(2), the hardware access-list tcam label ing-ifacl 6 command is also supported on Cisco Nexus 9300-FX2 platform switches.
