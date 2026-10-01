---
id: collect-261001-cisco/cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac-3
title: "en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac.md
source_anchor: ""
source_lines: [92, 135]
sha256: b80bb6591a6bec9dc61294ff4492298b97b744aa0aff5bcfc0abba538b8caa1f
---

# en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac

|  |  | (continued) – Extended PAT—The extended keyword enables extended PAT. Extended PAT uses 65535 ports per service , as opposed to per IP address, by including the destination address and port in the translation information. Normally, the destination port and address are not considered when creating PAT translations, so you are limited to 65535 ports per PAT address. For example, with extended PAT, you can create a translation of 10.1.1.1:1027 when going to 192.168.1.7:23 as well as a translation of 10.1.1.1:1027 when going to 192.168.1.7:80. – Flat range—The flat keyword enables use of the entire 1024 to 65535 port range when allocating ports. When choosing the mapped port number for a translation, the ASA uses the real source port number if it is available. However, without this option, if the real port is not available, by default the mapped ports are chosen from the same range of ports as the real port number: 1 to 511, 512 to 1023, and 1024 to 65535. To avoid running out of ports at the low ranges, configure this setting. To use the entire range of 1 to 65535, also specify the include-reserve keyword.  | 
Examples
The following example configures dynamic PAT that hides the 192.168.2.0 network behind address 10.2.2.2:
The following example configures dynamic PAT that hides the 192.168.2.0 network behind the outside interface address:
The following example configures dynamic PAT with a PAT pool to translate the inside IPv6 network to an outside IPv4 network:
ciscoasa(config)# object network IPv4_POOL
ciscoasa(config-network-object)# range 203.0.113.1 203.0.113.254
Configuring Static NAT or Static NAT-with-Port-Translation
This section describes how to configure a static NAT rule using network object NAT. For more information, see the “Static NAT” section.
Detailed Steps
| Step 1 | (Optional) Create a network object or group for the mapped addresses. | See the “Adding Network Objects for Mapped Addresses” section. | 
| Step 2 | object network obj_name ciscoasa(config)# object network my-host-obj1 | Configures a network object for which you want to configure NAT, or enters object network configuration mode for an existing network object. | 
| Step 3 | { host ip_address \| subnet subnet_address netmask \| range ip_address_1 ip_address_2 } ciscoasa(config-network-object)# subnet 10.2.1.0 255.255.255.0 | If you are creating a new network object, defines the real IP address(es) (IPv4 or IPv6) that you want to translate. | 
| Step 4 | nat [ ( real_ifc , mapped_ifc ) ] static { mapped_inline_ip \| mapped_obj \| interface [ ipv6 ]} [ net-to-net ] [ dns \| service { tcp \| udp } real_port mapped_port ] [ no-proxy-arp ] ciscoasa(config-network-object)# nat (inside,outside) static MAPPED_IPS service tcp 80 8080 | Configures static NAT for the object IP addresses. You can only define a single NAT rule for a given object.  – An inline IP address. The netmask or range for the mapped network is the same as that of the real network. For example, if the real network is a host, then this address will be a host address. In the case of a range, then the mapped addresses include the same number of addresses as the real range. For example, if the real address is defined as a range from 10.1.1.1 through 10.1.1.6, and you specify 172.20.1.1 as the mapped address, then the mapped range will include 172.20.1.1 through 172.20.1.6. – An existing network object or group (see Step 1). – interface —(Static NAT-with-port-translation only; routed mode) For this option, you must configure a specific interface for the mapped_ifc . If you specify ipv6 , then the IPv6 address of the interface is used. Be sure to also configure the service keyword. Typically, you configure the same number of mapped addresses as real addresses for a one-to-one mapping. You can, however, have a mismatched number of addresses. See the “Static NAT” section.  | 
Examples
The following example configures static NAT for the real host 10.1.1.1 on the inside to 10.2.2.2 on the outside with DNS rewrite enabled.
The following example configures static NAT for the real host 10.1.1.1 on the inside to 10.2.2.2 on the outside using a mapped object.
The following example configures static NAT-with-port-translation for 10.1.1.1 at TCP port 21 to the outside interface at port 2121.
The following example maps an inside IPv4 network to an outside IPv6 network.
The following example maps an inside IPv6 network to an outside IPv6 network.
Configuring Identity NAT
This section describes how to configure an identity NAT rule using network object NAT. For more information, see the “Identity NAT” section.
Detailed Steps
| Step 1 | (Optional) Create a network object for the mapped addresses. | The object must include the same addresses that you want to translate. See the “Adding Network Objects for Mapped Addresses” section. | 
| Step 2 | object network obj_name ciscoasa(config)# object network my-host-obj1 | Configures a network object for which you want to perform identity NAT, or enters object network configuration mode for an existing network object. This network object has a different name from the mapped network object (see Step 1) even though they both contain the same IP addresses. | 
| Step 3 | { host ip_address \| subnet subnet_address netmask \| range ip_address_1 ip_address_2 } ciscoasa(config-network-object)# subnet 10.1.1.0 255.255.255.0 | If you are creating a new network object, defines the real IP address(es) (IPv4 or IPv6) to which you want to perform identity NAT. If you configured a network object for the mapped addresses in Step 1, then these addresses must match. | 
| Step 4 | nat [ ( real_ifc , mapped_ifc ) ] static { mapped_inline_ip \| mapped_obj } [ no-proxy-arp ] [ route-lookup ] ciscoasa(config-network-object)# nat (inside,outside) static MAPPED_IPS | Configures identity NAT for the object IP addresses. Note You can only define a single NAT rule for a given object. See the “Additional Guidelines” section. See the following guidelines:  – Network object—Including the same IP address as the real object (see Step 1). – Inline IP address—The netmask or range for the mapped network is the same as that of the real network. For example, if the real network is a host, then this address will be a host address. In the case of a range, then the mapped addresses include the same number of addresses as the real range. For example, if the real address is defined as a range from 10.1.1.1 through 10.1.1.6, and you specify 10.1.1.1 as the mapped address, then the mapped range will include 10.1.1.1 through 10.1.1.6.  | 
Example
The following example maps a host address to itself using an inline mapped address:
ciscoasa(config)# object network my-host-obj1
ciscoasa(config-network-object)# host 10.1.1.1
ciscoasa(config-network-object)# nat (inside,outside) static 10.1.1.1
The following example maps a host address to itself using a network object:
ciscoasa(config)# object network my-host-obj1-identity
ciscoasa(config-network-object)# host 10.1.1.1
ciscoasa(config-network-object)# object network my-host-obj1
ciscoasa(config-network-object)# host 10.1.1.1
ciscoasa(config-network-object)# nat (inside,outside) static my-host-obj1-identity
Configuring Per-Session PAT Rules
By default, all TCP PAT traffic and all UDP DNS traffic uses per-session PAT. To use multi-session PAT for traffic, you can configure per-session PAT rules: a permit rule uses per-session PAT, and a deny rule uses multi-session PAT. For more information about per-session vs. multi-session PAT, see the “Per-Session PAT vs. Multi-Session PAT” section.
Defaults
By default, the following rules are installed:
Note You cannot remove these rules, and they always exist after any manually-created rules. Because rules are evaluated in order, you can override the default rules. For example, to completely negate these rules, you could add the following:
Detailed Steps
