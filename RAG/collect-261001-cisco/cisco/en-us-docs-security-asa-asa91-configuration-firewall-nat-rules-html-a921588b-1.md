---
id: collect-261001-cisco/cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b-1
title: "en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b.md
source_anchor: ""
source_lines: [1, 54]
sha256: 58bde2aafd79f96a7c21b34bb0512d68d184a269511ccd7c14ce82a43354a020
---

# en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b

Configuring Twice NAT
Twice NAT lets you identify both the source and destination address in a single rule. This chapter shows you how to configure twice NAT and includes the following sections:
Note For detailed information about how NAT works, see Chapter3, “Information About NAT”
Information About Twice NAT
Twice NAT lets you identify both the source and destination address in a single rule. Specifying both the source and destination addresses lets you specify that a source address should be translated to A when going to destination X, but be translated to B when going to destination Y, for example.
Note For static NAT, the rule is bidirectional, so be aware that “source” and “destination” are used in commands and descriptions throughout this guide even though a given connection might originate at the “destination” address. For example, if you configure static NAT with port address translation, and specify the source address as a Telnet server, and you want all traffic going to that Telnet server to have the port translated from 2323 to 23, then in the command, you must specify the source ports to be translated (real: 23, mapped: 2323). You specify the source ports because you specified the Telnet server address as the source address.
The destination address is optional. If you specify the destination address, you can either map it to itself (identity NAT), or you can map it to a different address. The destination mapping is always a static mapping.
Twice NAT also lets you use service objects for static NAT-with-port-translation; network object NAT only accepts inline definition.
For detailed information about the differences between twice NAT and network object NAT, see the “How NAT is Implemented” section.
Twice NAT rules are added to section 1 of the NAT rules table, or if specified, section 3. For more information about NAT ordering, see the “NAT Rule Order” section.
Prerequisites for Twice NAT
- For both the real and mapped addresses, configure network objects or network object groups (the object network or object-group network command). Network object groups are particularly useful for creating a mapped address pool with discontinuous IP address ranges or multiple hosts or subnets. To create a network object or group, see the general operations configuration guide.
- For static NAT-with-port-translation, configure TCP or UDP service objects (the object service command). To create a service object, see the general operations configuration guide.
For specific guidelines for objects and groups, see the configuration section for the NAT type you want to configure. See also the “Guidelines and Limitations” section.
Guidelines and Limitations
This section includes the guidelines and limitations for this feature.
Supported in single and multiple context mode.
- Supported in routed and transparent firewall mode.
- In transparent mode, you must specify the real and mapped interfaces; you cannot use any .
- In transparent mode, you cannot configure interface PAT, because the transparent mode interfaces do not have IP addresses. You also cannot use the management IP address as a mapped address.
- In transparent mode, translating between IPv4 and IPv6 networks is not supported. Translating between two IPv6 networks, or between two IPv4 networks is supported.
- Supports IPv6.
- For routed mode, you can also translate between IPv4 and IPv6.
- For transparent mode, translating between IPv4 and IPv6 networks is not supported. Translating between two IPv6 networks, or between two IPv4 networks is supported.
- For transparent mode, a PAT pool is not supported for IPv6.
- For static NAT, you can specify an IPv6 subnet up to /64. Larger subnets are not supported.
- When using FTP with NAT46, when an IPv4 FTP client connects to an IPv6 FTP server, the client must use either the extended passive mode (EPSV) or extended port mode (EPRT); PASV and PORT commands are not supported with IPv6.
- (This limitation is for 9.1.0 to 9.1.5; this limitation was removed in 9.1.6 and following maintenance releases.) You cannot configure FTP destination port translation when the source IP address is a subnet (or any other application that uses a secondary connection); the FTP data channel establishment does not succeed. For example, the following configuration does not work:
- If you change the NAT configuration, and you do not want to wait for existing translations to time out before the new NAT information is used, you can clear the translation table using the clear xlate command. However, clearing the translation table disconnects all current connections that use translations.
Note If you remove a dynamic NAT or PAT rule, and then add a new rule with mapped addresses that overlap the addresses in the removed rule, then the new rule will not be used until all connections associated with the removed rule time out or are cleared using the clear xlate command. This safeguard ensures that the same address is not assigned to multiple hosts.
- You cannot use an object group with both IPv4 and IPv6 addresses; the object group must include only one type of address.
- When using the any keyword in a NAT rule, the definition of “any” traffic (IPv4 vs. IPv6) depends on the rule. Before the ASA performs NAT on a packet, the packet must be IPv6-to-IPv6 or IPv4-to-IPv4; with this prerequisite, the ASA can determine the value of any in a NAT rule. For example, if you configure a rule from “any” to an IPv6 server, and that server was mapped from an IPv4 address, then any means “any IPv6 traffic.” If you configure a rule from “any” to “any,” and you map the source to the interface IPv4 address, then any means “any IPv4 traffic” because the mapped interface address implies that the destination is also IPv4.
- Objects and object groups used in NAT cannot be undefined; they must include IP addresses.
- You can use the same objects in multiple rules.
- The mapped IP address pool cannot include:
– The mapped interface IP address. If you specify any interface for the rule, then all interface IP addresses are disallowed. For interface PAT (routed mode only), use the interface keyword instead of the IP address.
– (Transparent mode) The management IP address.
– (Dynamic NAT) The standby interface IP address when VPN is enabled.
Default Settings
- By default, the rule is added to the end of section 1 of the NAT table.
- (Routed mode) The default real and mapped interface is Any, which applies the rule to all interfaces.
- If you specify an optional interface, then the ASA uses the NAT configuration to determine the egress interface, but you have the option to always use a route lookup instead.
Configuring Twice NAT
This section describes how to configure twice NAT. This section includes the following topics:
Adding Network Objects for Real and Mapped Addresses
For each NAT rule, configure up to four network objects or groups for:
Objects are required unless you specify the any keyword inline to represent all traffic, or for some types of NAT, the interface keyword to represent the interface address. For more information about configuring a network object or group, see the general operations configuration guide.
Guidelines
- A network object group can contain objects and/or inline addresses of either IPv4 or IPv6 addresses. The group cannot contain both IPv4 and IPv6 addresses; it must contain one type only.
- See the “Guidelines and Limitations” section for information about disallowed mapped IP addresses.
- Source Dynamic NAT:
– You typically configure a larger group of real addresses to be mapped to a smaller group.
– The mapped object or group cannot contain a subnet; the object must define a range; the group can include hosts and ranges.
– If a mapped network object contains both ranges and host IP addresses, then the ranges are used for dynamic NAT, and the host IP addresses are used as a PAT fallback.
