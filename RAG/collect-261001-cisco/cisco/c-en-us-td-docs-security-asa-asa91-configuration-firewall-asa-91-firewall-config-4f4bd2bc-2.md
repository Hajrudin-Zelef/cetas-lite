---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-4f4bd2bc-2
title: "c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-4f4bd2bc"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-4f4bd2bc.md
source_anchor: ""
source_lines: [7, 67]
sha256: 85f2c7c9bcd038180cf36d7552c9ae8de0bfe6a45ef9c5b0b9fd37067a1a3359
---

# c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-4f4bd2bc

  This chapter describes how to control network access through the ASA using access rules and includes the following sections:
Note You use access rules to control network access in both routed and transparent firewall modes. In transparent mode, you can use both access rules (for Layer 3 traffic) and EtherType rules (for Layer 2 traffic).
To access the ASA interface for management access, you do not also need an access rule allowing the host IP address. You only need to configure management access according to the general operations configuration guide.
Information About Access Rules
You create an access rule by applying an extended or EtherType ACL to an interface or globally for all interfaces.You can use access rules in routed and transparent firewall mode to control IP traffic. An access rule permits or denies traffic based on the protocol, a source and destination IP address or network, and optionally the source and destination ports.
For transparent mode only, an EtherType rule controls network access for non-IP traffic. An EtherType rule permits or denies traffic based on the EtherType.
This section includes the following topics:
General Information About Rules
This section describes information for both access rules and EtherType rules, and it includes the following topics:
Implicit Permits
For routed mode, the following types of traffic are allowed through by default:
- Unicast IPv4 traffic from a higher security interface to a lower security interface.
- Unicast IPv6 traffic from a higher security interface to a lower security interface.
For transparent mode, the following types of traffic are allowed through by default:
- Unicast IPv4 traffic from a higher security interface to a lower security interface.
- Unicast IPv6 traffic from a higher security interface to a lower security interface.
- ARPs in both directions.
Note ARP traffic can be controlled by ARP inspection, but cannot be controlled by an access rule.
For other traffic, you need to use either an extended access rule (IPv4 and IPv6) or an EtherType rule (non-IPv4/IPv6).
Information About Interface Access Rules and Global Access Rules
You can apply an access rule to a specific interface, or you can apply an access rule globally to all interfaces. You can configure global access rules in conjunction with interface access rules, in which case, the specific interface access rules are always processed before the general global access rules.
Note Global access rules apply only to inbound traffic. See the “Inbound and Outbound Rules” section.
Using Access Rules and EtherType Rules on the Same Interface
You can apply one access rule and one EtherType rule to each direction of an interface.
Implicit Deny
ACLs have an implicit deny at the end of the list, so unless you explicitly permit it, traffic cannot pass. For example, if you want to allow all users to access a network through the ASA except for particular addresses, then you need to deny the particular addresses and then permit all others.
For EtherType ACLs, the implicit deny at the end of the ACL does not affect IP traffic or ARPs; for example, if you allow EtherType 8037, the implicit deny at the end of the ACL does not now block any IP traffic that you previously allowed with an extended ACL (or implicitly allowed from a high security interface to a low security interface). However, if you explicitly deny all traffic with an EtherType ACE, then IP and ARP traffic is denied.
If you configure a global access rule, then the implicit deny comes after the global rule is processed. See the following order of operations:
Inbound and Outbound Rules
The ASA supports two types of ACLs:
- Inbound—Inbound access rules apply to traffic as it enters an interface. Global access rules are always inbound.
- Outbound—Outbound ACLs apply to traffic as it exits an interface.
Note “Inbound” and “outbound” refer to the application of an ACL on an interface, either to traffic entering the ASA on an interface or traffic exiting the ASA on an interface. These terms do not refer to the movement of traffic from a lower security interface to a higher security interface, commonly known as inbound, or from a higher to lower interface, commonly known as outbound.
An outbound ACL is useful, for example, if you want to allow only certain hosts on the inside networks to access a web server on the outside network. Rather than creating multiple inbound ACLs to restrict access, you can create a single outbound ACL that allows only the specified hosts. (See Figure 6-1.) The outbound ACL prevents any other hosts from reaching the outside network.
See the following commands for this example:
Transactional-Commit Model
The ASA rule-engine supports a new feature for rule updation called the Transactional-Commit Model. When this feature is enabled, a rule update is applied after the rule compilation is completed; without affecting the rule matching performance. With the legacy model, rule updates take effect immediately but rule matching slows down during the rule compilation period. This feature is useful to prevent potential packet drops during large compilation of rules under high traffic conditions. This feature is also useful to reduce the rule compilation time under two specific patterns of configurations:
Guidelines and Limitations
Supported in single and multiple context mode.
Supported in routed and transparent firewall mode.
Additional Guidelines and Limitations
Evaluate the following alternatives before using the transactional commit model:
- While using large rules, try to optimize the number of rules by using the Object Group Search setting in Advanced Access Rule Configuration settings.
- Perform an incremental rule update instead of a bulk rule update. If a bulk update is necessary perform the bulk update during the maintenance window, when traffic is low.
Information About Extended Access Rules
This section describes information about extended access rules and includes the following topics:
Access Rules for Returning Traffic
For TCP and UDP connections for both routed and transparent mode, you do not need an access rule to allow returning traffic because the ASA allows all returning traffic for established, bidirectional connections.
For connectionless protocols such as ICMP, however, the ASA establishes unidirectional sessions, so you either need access rules to allow ICMP in both directions (by applying ACLs to the source and destination interfaces), or you need to enable the ICMP inspection engine. The ICMP inspection engine treats ICMP sessions as bidirectional connections. To control ping, specify echo-reply (0) (ASA to host) or echo (8) (host to ASA).
Allowing Broadcast and Multicast Traffic through the Transparent Firewall Using Access Rules
In routed firewall mode, broadcast and multicast traffic is blocked even if you allow it in an access rule, including unsupported dynamic routing protocols and DHCP (unless you configure DHCP relay). Transparent firewall mode can allow any IP traffic through.
Note Because these special types of traffic are connectionless, you need to apply an access rule to both interfaces, so returning traffic is allowed through.
Table 6-1 lists common traffic types that you can allow through the transparent firewall.
|  |  |  | 
|---|---|---|
| DHCP | UDP ports 67 and 68 | If you enable the DHCP server, then the ASA does not pass DHCP packets. | 
| EIGRP | Protocol 88 | — | 
| OSPF | Protocol 89 | — | 
| Multicast streams | The UDP ports vary depending on the application. | Multicast streams are always destined to a Class D address (224.0.0.0 to 239.x.x.x). | 
| RIP (v1 or v2) | UDP port 520 | — | 
Management Access Rules
