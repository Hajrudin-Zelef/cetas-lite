---
id: collect-261001-cisco/cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac-1
title: "en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac.md
source_anchor: ""
source_lines: [1, 60]
sha256: 8f73333ab64bac02dbf42b73a927ad35b512b89961754c25554e9ea66232a0a5
---

# en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac

Configuring Network Object NAT
All NAT rules that are configured as a parameter of a network object are considered to be network object NAT rules. Network object NAT is a quick and easy way to configure NAT for a single IP address, a range of addresses, or a subnet. After you configure the network object, you can then identify the mapped address for that object.
This chapter describes how to configure network object NAT, and it includes the following sections:
Note For detailed information about how NAT works, see Chapter3, “Information About NAT”
Information About Network Object NAT
When a packet enters the ASA, both the source and destination IP addresses are checked against the network object NAT rules. The source and destination address in the packet can be translated by separate rules if separate matches are made. These rules are not tied to each other; different combinations of rules can be used depending on the traffic.
Because the rules are never paired, you cannot specify that a source address should be translated to A when going to destination X, but be translated to B when going to destination Y. Use twice NAT for that kind of functionality (twice NAT lets you identify the source and destination address in a single rule).
For detailed information about the differences between twice NAT and network object NAT, see the “How NAT is Implemented” section.
Network object NAT rules are added to section 2 of the NAT rules table. For more information about NAT ordering, see the “NAT Rule Order” section.
Licensing Requirements for Network Object NAT
The following table shows the licensing requirements for this feature:
Prerequisites for Network Object NAT
Depending on the configuration, you can configure the mapped address inline if desired or you can create a separate network object or network object group for the mapped address (the object network or object-group network command). Network object groups are particularly useful for creating a mapped address pool with discontinous IP address ranges or multiple hosts or subnets. To create a network object or group, see the general operations configuration guide.
For specific guidelines for objects and groups, see the configuration section for the NAT type you want to configure. See also the “Guidelines and Limitations” section.
Guidelines and Limitations
Supported in single and multiple context mode.
- Supported in routed and transparent firewall mode.
- In transparent mode, you must specify the real and mapped interfaces; you cannot use any .
- In transparent mode, you cannot configure interface PAT, because the transparent mode interfaces do not have IP addresses. You also cannot use the management IP address as a mapped address.
- In transparent mode, translating between IPv4 and IPv6 networks is not supported. Translating between two IPv6 networks, or between two IPv4 networks is supported.
- Supports IPv6. See also the “NAT and IPv6” section.
- For routed mode, you can also translate between IPv4 and IPv6.
- For transparent mode, translating between IPv4 and IPv6 networks is not supported. Translating between two IPv6 networks, or between two IPv4 networks is supported.
- For transparent mode, a PAT pool is not supported for IPv6.
- For static NAT, you can specify an IPv6 subnet up to /64. Larger subnets are not supported.
- When using FTP with NAT46, when an IPv4 FTP client connects to an IPv6 FTP server, the client must use either the extended passive mode (EPSV) or extended port mode (EPRT); PASV and PORT commands are not supported with IPv6.
- You can only define a single NAT rule for a given object; if you want to configure multiple NAT rules for an object, you need to create multiple objects with different names that specify the same IP address, for example, object network obj-10.10.10.1-01 , object network obj-10.10.10.1-02 , and so on.
- If you change the NAT configuration, and you do not want to wait for existing translations to time out before the new NAT configuration is used, you can clear the translation table using the clear xlate command. However, clearing the translation table disconnects all current connections that use translations.
Note If you remove a dynamic NAT or PAT rule, and then add a new rule with mapped addresses that overlap the addresses in the removed rule, then the new rule will not be used until all connections associated with the removed rule time out or are cleared using the clear xlate command. This safeguard ensures that the same address is not assigned to multiple hosts.
- Objects and object groups used in NAT cannot be undefined; they must include IP addresses.
- You cannot use an object group with both IPv4 and IPv6 addresses; the object group must include only one type of address.
- You can use the same mapped object or group in multiple NAT rules.
- The mapped IP address pool cannot include:
– The mapped interface IP address. If you specify any interface for the rule, then all interface IP addresses are disallowed. For interface PAT (routed mode only), use the interface keyword instead of the IP address.
– (Transparent mode) The management IP address.
– (Dynamic NAT) The standby interface IP address when VPN is enabled.
– Existing VPN pool addresses.
- For application inspection limitations with NAT or PAT, see the “Default Settings and NAT Limitations” section in Chapter9, “Getting Started with Application Layer Protocol Inspection”
Default Settings
- (Routed mode) The default real and mapped interface is Any, which applies the rule to all interfaces.
- The default behavior for identity NAT has proxy ARP enabled, matching other static NAT rules. You can disable proxy ARP if desired. See the “Routing NAT Packets” section for more information.
- If you specify an optional interface, then the ASA uses the NAT configuration to determine the egress interface, but you have the option to always use a route lookup instead. See the “Routing NAT Packets” section for more information.
Configuring Network Object NAT
This section describes how to configure network object NAT and includes the following topics:
Adding Network Objects for Mapped Addresses
For dynamic NAT, you must use an object or group for the mapped addresses. Other NAT types have the option of using inline addresses, or you can create an object or group according to this section. For more information about configuring a network object or group, see the general operations configuration guide.
Guidelines
- A network object group can contain objects and/or inline addresses of either IPv4 or IPv6 addresses. The group cannot contain both IPv4 and IPv6 addresses; it must contain one type only.
- See the “Guidelines and Limitations” section for information about disallowed mapped IP addresses.
- Dynamic NAT:
– You cannot use an inline address; you must configure a network object or group.
– The object or group cannot contain a subnet; the object must define a range; the group can include hosts and ranges.
– If a mapped network object contains both ranges and host IP addresses, then the ranges are used for dynamic NAT, and then the host IP addresses are used as a PAT fallback.
– Instead of using an object, you can optionally configure an inline host address or specify the interface address.
– If you use an object, the object or group cannot contain a subnet; the object must define a host, or for a PAT pool, a range; the group (for a PAT pool) can include hosts and ranges.
– Instead of using an object, you can configure an inline address or specify the interface address (for static NAT-with-port-translation).
– If you use an object, the object or group can contain a host, range, or subnet.
– Instead of using an object, you can configure an inline address.
– If you use an object, the object must match the real addresses you want to translate.
Detailed Steps
