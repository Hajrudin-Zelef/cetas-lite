---
id: collect-261001-cisco/cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b-3
title: "en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b.md
source_anchor: ""
source_lines: [97, 133]
sha256: b0f80fec8f79a452d91737259fa89223979d3b09e8a92033ada8c5b2e7b4ff73
---

# en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b

ciscoasa(config)# object network SERVERS_1
ciscoasa(config-network-object)# subnet 209.165.201.0 255.255.255.224
ciscoasa(config)# object network SERVERS_2
ciscoasa(config-network-object)# subnet 203.0.113.0 255.255.255.0
ciscoasa(config)# nat (inside,outside) source dynamic INSIDE_NW MAPPED_1 destination static SERVERS_1 SERVERS_1
ciscoasa(config)# nat (inside,outside) source dynamic INSIDE_NW MAPPED_2 destination static SERVERS_2 SERVERS_2
The following example configures dynamic NAT for an IPv6 inside network 2001:DB8:AAAA::/96 when accessing servers on the IPv4 209.165.201.1/27 network as well as servers on the 203.0.113.0/24 network:
ciscoasa(config)# object network INSIDE_NW
ciscoasa(config-network-object)# subnet 2001:DB8:AAAA::/96
ciscoasa(config)# object network MAPPED_1
ciscoasa(config-network-object)# range 209.165.200.225 209.165.200.254
ciscoasa(config)# object network MAPPED_2
ciscoasa(config-network-object)# range 209.165.202.129 209.165.200.158
ciscoasa(config)# object network SERVERS_1
ciscoasa(config-network-object)# subnet 209.165.201.0 255.255.255.224
ciscoasa(config)# object network SERVERS_2
ciscoasa(config-network-object)# subnet 203.0.113.0 255.255.255.0
ciscoasa(config)# nat (inside,outside) source dynamic INSIDE_NW MAPPED_1 destination static SERVERS_1 SERVERS_1
ciscoasa(config)# nat (inside,outside) source dynamic INSIDE_NW MAPPED_2 destination static SERVERS_2 SERVERS_2
Configuring Dynamic PAT (Hide)
This section describes how to configure twice NAT for dynamic PAT (hide). For more information, see the “Dynamic PAT” section.
Guidelines
- If available, the real source port number is used for the mapped port. However, if the real port is not available, by default the mapped ports are chosen from the same range of ports as the real port number: 0 to 511, 512 to 1023, and 1024 to 65535. Therefore, ports below 1024 have only a small PAT pool that can be used. (8.4(3) and later, not including 8.5(1) or 8.6(1)) If you have a lot of traffic that uses the lower port ranges, you can now specify a flat range of ports to be used instead of the three unequal-sized tiers: either 1024 to 65535, or 1 to 65535.
- If you use the same PAT pool object in two separate rules, then be sure to specify the same options for each rule. For example, if one rule specifies extended PAT and a flat range, then the other rule must also specify extended PAT and a flat range.
For extended PAT for a PAT pool:
- Many application inspections do not support extended PAT. See the “Default Settings and NAT Limitations” section in “Getting Started with Application Layer Protocol Inspection,” for a complete list of unsupported inspections.
- If you enable extended PAT for a dynamic PAT rule, then you cannot also use an address in the PAT pool as the PAT address in a separate static NAT-with-port-translation rule. For example, if the PAT pool includes 10.1.1.1, then you cannot create a static NAT-with-port-translation rule using 10.1.1.1 as the PAT address.
- If you use a PAT pool and specify an interface for fallback, you cannot specify extended PAT.
- For VoIP deployments that use ICE or TURN, do not use extended PAT. ICE and TURN rely on the PAT binding to be the same for all destinations.
For round robin for a PAT pool:
- If a host has an existing connection, then subsequent connections from that host will use the same PAT IP address if ports are available. Note : This “stickiness” does not survive a failover. If the ASA fails over, then subsequent connections from a host may not use the initial IP address.
- Round robin, especially when combined with extended PAT, can consume a large amount of memory. Because NAT pools are created for every mapped protocol/IP address/port range, round robin results in a large number of concurrent NAT pools, which use memory. Extended PAT results in an even larger number of concurrent NAT pools.
Detailed Steps
| Step 1 | Create network objects or groups for the: | See the “Adding Network Objects for Real and Mapped Addresses” section. If you want to translate all source traffic, you can skip adding an object for the source real addresses, and instead specify the any keyword in the nat command. If you want to use the interface address as the mapped address, you can skip adding an object for the source mapped addresses, and instead specify the interface keyword in the nat command. If you want to configure destination static interface NAT with port translation only, you can skip adding an object for the destination mapped addresses, and instead specify the interface keyword in the nat command. | 
| Step 2 | (Optional) Create service objects for the: | See the “(Optional) Adding Service Objects for Real and Mapped Ports” section. | 
| Step 3 | nat [ ( real_ifc , mapped_ifc ) ] [ line \| { after-auto [ line ]}] source dynamic { real-obj \| any } { mapped_obj [ interface [ ipv6 ]] \| [ pat-pool mapped_obj [ round-robin ] [ extended ] [ flat [ include-reserve ]] [ interface [ ipv6 ]] \| interface [ ipv6 ]} [ destination static { mapped_obj \| interface [ ipv6 ]} real_obj ] [ service mapped_dest_svc_obj real_dest_svc_obj ] [ dns ] [ unidirectional ] [ inactive ] [ description desc ] ciscoasa(config)# nat (inside,outside) source dynamic MyInsNet interface destination static Server1 Server1 description Interface PAT for inside addresses when going to server 1 | Configures dynamic PAT (hide) . See the following guidelines:  – Real—Specify a network object, group, or the any keyword. Use the any keyword if you want to translate all traffic from the real interface to the mapped interface. – Mapped—Configure one of the following: - Network object—Specify a network object that contains a host address. - pat-pool —Specify the pat-pool keyword and a network object or group that contains multiple addresses. - interface —(Routed mode only) Specify the interface keyword alone to only use interface PAT. If you specify ipv6 , then the IPv6 address of the interface is used. When specified with a PAT pool or network object, the interface keyword enables interface PAT fallback. After the PAT IP addresses are used up, then the IP address of the mapped interface is used. For this option, you must configure a specific interface for the mapped_ifc . (continued) | 
|  |  | (continued) For a PAT pool, you can specify one or more of the following options: -- Round robin—The round-robin keyword enables round-robin address allocation for a PAT pool. Without round robin, by default all ports for a PAT address will be allocated before the next PAT address is used. The round-robin method assigns an address/port from each PAT address in the pool before returning to use the first address again, and then the second address, and so on. -- Extended PAT—The extended keyword enables extended PAT. Extended PAT uses 65535 ports per service , as opposed to per IP address, by including the destination address and port in the translation information. Normally, the destination port and address are not considered when creating PAT translations, so you are limited to 65535 ports per PAT address. For example, with extended PAT, you can create a translation of 10.1.1.1:1027 when going to 192.168.1.7:23 as well as a translation of 10.1.1.1:1027 when going to 192.168.1.7:80. -- Flat range—The flat keyword enables use of the entire 1024 to 65535 port range when allocating ports. When choosing the mapped port number for a translation, the ASA uses the real source port number if it is available. However, without this option, if the real port is not available, by default the mapped ports are chosen from the same range of ports as the real port number: 1 to 511, 512 to 1023, and 1024 to 65535. To avoid running out of ports at the low ranges, configure this setting. To use the entire range of 1 to 65535, also specify the include-reserve keyword. (continued) | 
