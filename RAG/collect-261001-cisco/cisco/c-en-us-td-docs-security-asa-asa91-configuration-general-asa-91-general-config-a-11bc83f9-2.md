---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9-2
title: "c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9.md
source_anchor: ""
source_lines: [7, 83]
sha256: eed56e3c1de448eaae4ac5be5d13c60d75b2aa5023f3b2886f81c9b8f0ac3fe0
---

# c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9

  This chapter describes how to configure reusable named objects and groups for use in your configuration, and it includes the following sections:
Information About Objects
Objects are reusable components for use in your configuration. They can be defined and used in ASA configurations in the place of inline IP addresses, services, names, and so on. Objects make it easy to maintain your configurations because you can modify an object in one place and have it be reflected in all other places that are referencing it. Without objects you would have to modify the parameters for every feature when required, instead of just once. For example, if a network object defines an IP address and subnet mask, and you want to change the address, you only need to change it in the object definition, not in every feature that refers to that IP address.
Licensing Requirements for Objects
Guidelines and Limitations
Supported in single and multiple context mode.
Supported in routed and transparent firewall mode.
- Supports IPv6.
- The ASA does not support IPv6 nested network object groups, so you cannot group an object with IPv6 entries under another IPv6 object group.
- You can mix IPv4 and IPv6 entries in a network object group; you cannot use a mixed object group for NAT.
Additional Guidelines and Limitations
- Object must have unique names. While you might want to create a network object group named “Engineering” and a service object group named “Engineering,” you need to add an identifier (or “tag”) to the end of at least one object group name to make it unique. For example, you can use the names “Engineering_admins” and “Engineering_hosts” to make the object group names unique and to aid in identification.
- Objects and object groups share the same name space.
- You cannot remove an object or make an object empty if it is used in a command.
Configuring Objects
Configuring Network Objects and Groups
This section describes how to configure network objects and groups, and it includes the following topics:
Configuring a Network Object
A network object can contain a host, a network IP address, or a range of IP addresses, a fully qualified domain name (FQDN). You can also enable NAT rules on the object (excepting FQDN objects). (See Chapter 4, “Configuring Network Object NAT,” in the firewall configuration guide for more information.)
Detailed Steps
|  |  |  | 
|---|---|---|
| Step 1 |  | Creates a new network object. The obj_name is a text string up to 64 characters in length and can be any combination of letters, digits, and the following characters: The prompt changes to network object configuration mode. | 
| Step 2 | { host ip_addr \| subnet net_addr net_mask \| range ip_addr_1 ip_addr_2 \| fqdn fully_qualified_domain_name } | Assigns the IP address or FQDN to the named object. Note You cannot configure NAT for an FQDN object. | 
| Step 3 |  | Adds a description to the object. | 
Examples
To create a network object, enter the following commands:
Configuring a Network Object Group
Network object groups can contain multiple network objects as well as inline networks. Network object groups can support a mix of both IPv4 and IPv6 addresses.
Restrictions
You cannot use a mixed IPv4 and IPv6 object group for NAT, or object groups that include FQDN objects.
Detailed Steps
|  |  |  | 
|---|---|---|
| Step 1 |  | Adds a network group. The grp_id is a text string up to 64 characters in length and can be any combination of letters, digits, and the following characters: The prompt changes to protocol configuration mode. | 
| Step 2 |  | (Optional) Adds a description. The description can be up to 200 characters. | 
| Step 3 | Add one or more of the following group members: |  | 
|  |  | Adds an object to the network object group. | 
|  |  | Adds a host or network inline, either IPv4 or IPv6. | 
|  |  | Adds an existing object group under this object group. The nested group must be of the same type. | 
Example
To create a network group that includes the IP addresses of three administrators, enter the following commands:
Create network object groups for privileged users from various departments by entering the following commands:
You then nest all three groups together as follows:
Configuring Service Objects and Service Groups
Service objects and groups identify protocols and ports. This section describes how to configure service objects, service groups, TCP and UDP port service groups, protocol groups, and ICMP groups, and it includes the following topics:
Configuring a Service Object
The service object can contain a protocol, ICMP, ICMPv6, TCP or UDP port or port ranges.
Detailed Steps
|  |  |  | 
|---|---|---|
| Step 1 |  | Creates a new service object. The obj_name is a text string up to 64 characters in length and can be any combination of letters, digits, and the following characters: The prompt changes to service object configuration mode. | 
| Step 2 |  | Adds a service definition to the object. The protocol argument specifies an IP protocol name or number. If you specify the icmp , icmp6 , tcp , or udp protocols, you can specify additional properties; for other protocols, there are no additional properties. For icmp and icmp6 (for ICMP version 6), you can optionally include the ICMP type, such as “echo” or the type number. If you specify a type, you can optionally include an ICMP code, between 1 and 255. For TCP and UDP, you can optionally specify the source and/or destination ports, between 0 and 65535. For a list of supported names, see the CLI help. The operator can be: | 
Example
To create a service object, enter the following commands:
Configuring a Service Group
A service object group includes a mix of protocols, if desired, including optional source and destination ports for TCP or UDP.
Detailed Steps
|  |  |  | 
|---|---|---|
| Step 1 |  | Adds a service group. The grp_id is a text string up to 64 characters in length and can be any combination of letters, digits, and the following characters: The prompt changes to service configuration mode. | 
| Step 2 | Add one or more of the following group members: |  | 
|  |  | Identifies the protocol name or number, between 0 and 255. | 
|  |  | You can specify the source and/or destination ports, between 0 and 65535, for the TCP and UDP protocols; tcp-udp matches both protocols. For a list of supported names, see the CLI help. The operator can be: | 
|  |  | Specifies that the service type is for ICMP or ICMPv6 connections. You can optionally specify the ICMP type by name or number, between 0 and 255. The optional icmp_code specifies an ICMP code, between 1 and 255. | 
|  |  | Specifies a service object name, created with the object service command. | 
|  |  | Adds an existing object group under this object group. The nested group must be of the same type. | 
| Step 3 |  | (Optional) Adds a description. The description can be up to 200 characters. | 
Examples
The following example shows how to add both TCP and UDP services to a service object group:
The following example shows how to add multiple service objects to a service object group:
Configuring a TCP or UDP Port Service Group
A TCP or UDP service group includes a group of ports for a specific protocol (TCP, UDP, or TCP-UDP).
|  |  |  | 
|---|---|---|
| Step 1 |  | Adds a service group. The object keyword adds an additional object to the service object group. The grp_id is a text string up to 64 characters in length and can be any combination of letters, digits, and the following characters: Specifies the protocol for the services (ports) you want to add with either the tcp , udp , or tcp-udp keywords. Enter the tcp-udp keyword if your service uses both TCP and UDP with the same port number, for example, DNS (port53). The prompt changes to service configuration mode. | 
| Step 2 | Add one or more of the following group members: |  | 
