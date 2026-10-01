---
id: collect-261001-cisco/cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b-2
title: "en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b.md
source_anchor: ""
source_lines: [55, 96]
sha256: 31b095bb24e662786265d8dc9b4d387d0b7d9a9007506a8f5c33e922d5992ec5
---

# en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b

– The mapped object or group cannot contain a subnet; a network object must define a host, or for a PAT pool, a range; a network object group (for a PAT pool) can include hosts and ranges.
– The mapped object or group can contain a host, range, or subnet.
– The static mapping is typically one-to-one, so the real addresses have the same quantity as the mapped addresses. You can, however, have different quantities if desired. For more information, see the “Static NAT” section.
– The real and mapped objects must match; you can use the same object for both, or you can create separate objects that contain the same IP addresses.
- Destination Static NAT or Static NAT with port translation (the destination translation is always static):
– Although the main feature of twice NAT is the inclusion of the destination IP address, the destination address is optional. If you do specify the destination address, you can configure static translation for that address or just use identity NAT for it. You might want to configure twice NAT without a destination address to take advantage of some of the other qualities of twice NAT, including the use of network object groups for real addresses, or manually ordering of rules. For more information, see the “Main Differences Between Network Object NAT and Twice NAT” section.
– For identity NAT, the real and mapped objects must match; you can use the same object for both, or you can create separate objects that contain the same IP addresses.
– The static mapping is typically one-to-one, so the real addresses have the same quantity as the mapped addresses. You can, however, have different quantities if desired. For more information, see the “Static NAT” section.
– For static interface NAT with port translation (routed mode only), you can specify the interface keyword instead of a network object/group for the mapped address. For more information, see the “Static Interface NAT with Port Translation” section.
Detailed Steps
| object network obj_name ciscoasa(config)# object network MyInsNet ciscoasa(config-network-object)# subnet 10.1.1.0 255.255.255.0 | Adds a network object, either IPv4 or IPv6. | 
| object-group network grp_name  { network-object  { object   net_obj_name  \|  subnet_address netmask  \|  host   ip_address } \|  group-object   grp_obj_name } ciscoasa(config)# object network TEST ciscoasa(config-network-object)# range 10.1.1.1 10.1.1.70  ciscoasa(config)# object network TEST2 ciscoasa(config-network-object)# range 10.1.2.1 10.1.2.70  ciscoasa(config-network-object)# object-group network MAPPED_IPS ciscoasa(config-network)# network-object object TEST ciscoasa(config-network)# network-object object TEST2 ciscoasa(config-network)# network-object host 10.1.2.79 | Adds a network object group, either IPv4 or IPv6. | 
(Optional) Adding Service Objects for Real and Mapped Ports
Configure service objects for:
- Source real port (Static only) or Destination real port
- Source mapped port (Static only) or Destination mapped port
For more information about configuring a service object, see the general operations configuration guide.
Guidelines
- NAT only supports TCP or UDP. When translating a port, be sure the protocols in the real and mapped service objects are identical (both TCP or both UDP).
- The “not equal” ( neq ) operator is not supported.
- For identity port translation, you can use the same service object for both the real and mapped ports.
- Source Dynamic NAT—Source Dynamic NAT does not support port translation.
- Source Dynamic PAT (Hide)—Source Dynamic PAT does not support port translation.
- Source Static NAT or Static NAT with port translation—A service object can contain both a source and destination port; however, you should specify either the source or the destination port for both service objects. You should only specify both the source and destination ports if your application uses a fixed source port (such as some DNS servers); but fixed source ports are rare. For example, if you want to translate the port for the source host, then configure the source service.
- Source Identity NAT—A service object can contain both a source and destination port; however, you should specify either the source or the destination port for both service objects. You should only specify both the source and destination ports if your application uses a fixed source port (such as some DNS servers); but fixed source ports are rare. For example, if you want to translate the port for the source host, then configure the source service.
- Destination Static NAT or Static NAT with port translation (the destination translation is always static)—For non-static source NAT, you can only perform port translation on the destination. A service object can contain both a source and destination port, but only the destination port is used in this case. If you specify the source port, it will be ignored.
Detailed Steps
Configuring Dynamic NAT
This section describes how to configure twice NAT for dynamic NAT. For more information, see the “Dynamic NAT” section.
Detailed Steps
| Step 1 | Create network objects or groups for the: | See the “Adding Network Objects for Real and Mapped Addresses” section. If you want to translate all source traffic, you can skip adding an object for the source real addresses, and instead specify the any keyword in the nat command. If you want to configure destination static interface NAT with port translation only, you can skip adding an object for the destination mapped addresses, and instead specify the interface keyword in the nat command. | 
| Step 2 | (Optional) Create service objects for the: | See the “(Optional) Adding Service Objects for Real and Mapped Ports” section. | 
| Step 3 | nat [ ( real_ifc , mapped_ifc ) ] [ line \| { after-auto [ line ]}] source dynamic { real_obj \| any } { mapped_obj [ interface [ ipv6 ]]} [ destination static { mapped_obj \| interface [ ipv6 ]} real_obj ] [ service mapped_dest_svc_obj real_dest_svc_obj ] [ dns ] [ unidirectional ] [ inactive ] [ description desc ] ciscoasa(config)# nat (inside,outside) source dynamic MyInsNet NAT_POOL destination static Server1_mapped Server1 service MAPPED_SVC REAL_SVC | Configure dynamic NAT . See the following guidelines:  – Real—Specify a network object, group, or the any keyword. – Mapped—Specify a different network object or group. You can optionally configure the following fallback method: Interface PAT fallback—(Routed mode only) The interface keyword enables interface PAT fallback. If you specify ipv6 , then the IPv6 address of the interface is used. After the mapped IP addresses are used up, then the IP address of the mapped interface is used. For this option, you must configure a specific interface for the mapped_ifc . | 
|  |  | (Continued) – Mapped—Specify a network object or group, or for static interface NAT with port translation only, specify the interface keyword. If you specify ipv6 , then the IPv6 address of the interface is used. If you specify interface , be sure to also configure the service keyword. For this option, you must configure a specific interface for the real_ifc . See the “Static Interface NAT with Port Translation” section for more information. – Real—Specify a network object or group. For identity NAT, simply use the same object or group for both the real and mapped addresses.  | 
Examples
The following example configures dynamic NAT for inside network 10.1.1.0/24 when accessing servers on the 209.165.201.1/27 network as well as servers on the 203.0.113.0/24 network:
ciscoasa(config)# object network INSIDE_NW
ciscoasa(config-network-object)# subnet 10.1.1.0 255.255.255.0
ciscoasa(config)# object network MAPPED_1
ciscoasa(config-network-object)# range 209.165.200.225 209.165.200.254
ciscoasa(config)# object network MAPPED_2
ciscoasa(config-network-object)# range 209.165.202.129 209.165.200.158
