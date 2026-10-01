---
id: collect-261001-cisco/cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac-2
title: "en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac.md
source_anchor: ""
source_lines: [61, 91]
sha256: 33cda283a7ee83ea4ddc6dc80bccb864a747a853a465b9944fbc3d76f3ca0cf1
---

# en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac

| object network obj_name ciscoasa(config)# object network TEST ciscoasa(config-network-object)# range 10.1.1.1 10.1.1.70 | Adds a network object, either IPv4 or IPv6. | 
| object-group network grp_name  { network-object  { object   net_obj_name  \|  subnet_address netmask  \|  host   ip_address } \|  group-object   grp_obj_name } ciscoasa(config)# object network TEST ciscoasa(config-network-object)# range 10.1.1.1 10.1.1.70  ciscoasa(config)# object network TEST2 ciscoasa(config-network-object)# range 10.1.2.1 10.1.2.70  ciscoasa(config-network-object)# object-group network MAPPED_IPS ciscoasa(config-network)# network-object object TEST ciscoasa(config-network)# network-object object TEST2 ciscoasa(config-network)# network-object host 10.1.2.79 | Adds a network object group, either IPv4 or IPv6. | 
Configuring Dynamic NAT
This section describes how to configure network object NAT for dynamic NAT. For more information, see the “Dynamic NAT” section.
Detailed Steps
| Step 1 | Create a network object or group for the mapped addresses. | See the “Adding Network Objects for Mapped Addresses” section. | 
| Step 2 | object network obj_name ciscoasa(config)# object network my-host-obj1 | Configures a network object for which you want to configure NAT, or enters object network configuration mode for an existing network object. | 
| Step 3 | { host ip_address \| subnet subnet_address netmask \| range ip_address_1 ip_address_2 } ciscoasa(config-network-object)# subnet 10.1.1.0 255.255.255.0 | If you are creating a new network object, defines the real IP address(es) (either IPv4 or IPv6) that you want to translate. | 
| Step 4 | nat [ ( real_ifc , mapped_ifc ) ] dynamic mapped_obj [ interface [ ipv6 ]] [ dns ] ciscoasa(config-network-object)# nat (inside,outside) dynamic MAPPED_IPS interface | Configures dynamic NAT for the object IP addresses. Note You can only define a single NAT rule for a given object. See the “Additional Guidelines” section. See the following guidelines:  – An existing network object (see Step 1). – An existing network object group (see Step 1).  | 
Examples
The following example configures dynamic NAT that hides 192.168.2.0 network behind a range of outside addresses 10.2.2.1 through 10.2.2.10:
The following example configures dynamic NAT with dynamic PAT backup. Hosts on inside network 10.76.11.0 are mapped first to the nat-range1 pool (10.10.10.10-10.10.10.20). After all addresses in the nat-range1 pool are allocated, dynamic PAT is performed using the pat-ip1 address (10.10.10.21). In the unlikely event that the PAT translations are also used up, dynamic PAT is performed using the outside interface address.
The following example configures dynamic NAT with dynamic PAT backup to translate IPv6 hosts to IPv4. Hosts on inside network 2001:DB8::/96 are mapped first to the IPv4_NAT_RANGE pool (209.165.201.1 to 209.165.201.30). After all addresses in the IPv4_NAT_RANGE pool are allocated, dynamic PAT is performed using the IPv4_PAT address (209.165.201.31). In the event that the PAT translations are also used up, dynamic PAT is performed using the outside interface address.
Configuring Dynamic PAT (Hide)
This section describes how to configure network object NAT for dynamic PAT (hide). For more information, see the “Dynamic PAT” section.
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
| Step 1 | (Optional) Create a network object or group for the mapped addresses. | See the “Adding Network Objects for Mapped Addresses” section. | 
| Step 2 | object network obj_name ciscoasa(config)# object network my-host-obj1 | Configures a network object for which you want to configure NAT, or enters object network configuration mode for an existing network object. | 
| Step 3 | { host ip_address \| subnet subnet_address netmask \| range ip_address_1 ip_address_2 } ciscoasa(config-network-object)# range 10.1.1.1 10.1.1.90 | If you are creating a new network object, defines the real IP address(es) (either IPv4 or IPv6) that you want to translate. | 
| Step 4 | nat [ ( real_ifc , mapped_ifc ) ] dynamic { mapped_inline_host_ip \| mapped_obj \| pat-pool mapped_obj [ round-robin ] [ extended ] [ flat [ include-reserve ]] \| interface [ ipv6 ]} [ interface [ ipv6 ]] [ dns ] ciscoasa(config-network-object)# nat (any,outside) dynamic interface | Configures dynamic PAT for the object IP addresses. You can only define a single NAT rule for a given object. See the “Additional Guidelines” section. See the following guidelines:  – An inline host address. – An existing network object that is defined as a host address (see Step 1). – pat-pool —An existing network object or group that contains multiple addresses. – interface —(Routed mode only) The IP address of the mapped interface is used as the mapped address. If you specify ipv6 , then the IPv6 address of the interface is used. For this option, you must configure a specific interface for the mapped_ifc . You must use this keyword when you want to use the interface IP address; you cannot enter it inline or as an object. – Round robin—The round-robin keyword enables round-robin address allocation for a PAT pool. Without round robin, by default all ports for a PAT address will be allocated before the next PAT address is used. The round-robin method assigns an address/port from each PAT address in the pool before returning to use the first address again, and then the second address, and so on. (continued) | 
