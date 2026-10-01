---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e-1
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e.md
source_anchor: ""
source_lines: [1, 127]
sha256: 85e01fb5a5dd2b535c77bdb46925f142e216e3a5147ef5ac131426b7728b41de
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e

About ACLs
An ACL is an ordered set of rules that you can use to filter traffic. Each rule specifies a set of conditions that a packet must satisfy to match the rule. When the device determines that an ACL applies to a packet, it tests the packet against the conditions of all rules. The first matching rule determines whether the packet is permitted or denied. If there is no match, the device applies the applicable implicit rule. The device continues processing packets that are permitted and drops packets that are denied.
You can use ACLs to protect networks and specific hosts from unnecessary or unwanted traffic. For example, you could use ACLs to disallow HTTP traffic from a high-security network to the Internet. You could also use ACLs to allow HTTP traffic but only to specific sites, using the IP address of the site to identify it in an IP ACL.
ACL Types and Applications
The device supports the following types of ACLs for security traffic filtering:
- IPv4 ACLs
- The device applies IPv4 ACLs only to IPv4 traffic.
- IPv6 ACLs
- The device applies IPv6 ACLs only to IPv6 traffic.
- MAC ACLs
- The device applies MAC ACLs only to non-IP traffic.
IP and MAC ACLs have the following types of applications:
- Port ACL
- Filters Layer 2 traffic
- MAC ACL with UDF-based match
- Filters MAC ACLs with UDF-based match
- Router ACL
- Filters Layer 3 traffic
- VLAN ACL
- Filters VLAN traffic
- VTY ACL
- Filters virtual teletype (VTY) traffic
This table summarizes the applications for security ACLs.
| Table 1. Security ACL                                     		Applications |  |  | 
|---|---|---|
| Application | Supported Interfaces | Types of ACLs Supported | 
|---|---|---|
| Port ACL |  When a port ACL is applied to a trunk port, the ACL filters traffic on all VLANs on the trunk port. |  | 
| Router ACL |  |  | 
| VLAN ACL |  |  | 
| VTY ACL |  |  | 
| Note | You must enable VLAN interfaces globally before you can configure a VLAN interface. | 
| Note | MAC ACLs are supported on Layer 3 interfaces only if you enable MAC packet classification. | 
Order of ACL Application
When the device processes a packet, it determines the forwarding path of the packet. The path determines which ACLs that the device applies to the traffic. The device applies the ACLs in the following order:
-  
                                 		  
                                 Port ACL
-  
                                 		  
                                 Ingress VACL
-  
                                 		  
                                 Ingress router ACL
-  
                                 		  
                                 Ingress VTY ACL
-  
                                 		  
                                 Egress VTY ACL
-  
                                 		  
                                 Egress router ACL
-  
                                 		  
                                 Egress VACL
If the packet is bridged within the ingress VLAN, the device does not apply router ACLs.
About Rules
Rules are what you create, modify, and remove when you configure how an ACL filters network traffic. Rules appear in the running configuration. When you apply an ACL to an interface or change a rule within an ACL that is already applied to an interface, the supervisor module creates ACL entries from the rules in the running configuration and sends those ACL entries to the applicable I/O module. Depending upon how you configure the ACL, there may be more ACL entries than rules, especially if you implement policy-based ACLs by using object groups when you configure rules.
You can create rules in access-list configuration mode by using the permit or deny command. The device allows traffic that matches the criteria in a permit rule and blocks traffic that matches the criteria in a deny rule. You have many options for configuring the criteria that traffic must meet in order to match the rule.
This section describes some of the options that you can use when you configure a rule.
Protocols for IP ACLs and MAC ACLs
IPv4, IPv6, and MAC ACLs allow you to identify traffic by protocol. For your convenience, you can specify some protocols by name. For example, in an IPv4 or IPv6 ACL, you can specify ICMP by name.
You can specify any protocol by number. In MAC ACLs, you can specify protocols by the EtherType number of the protocol, which is a hexadecimal number. For example, you can use 0x0800 to specify IP traffic in a MAC ACL rule.
In IPv4 and IPv6 ACLs, you can specify protocols by the integer that represents the Internet protocol number.
Source and Destination
In each rule, you specify the source and the destination of the traffic that matches the rule. You can specify both the source and destination as a specific host, a network or group of hosts, or any host. How you specify the source and destination depends on whether you are configuring IPv4 ACLs, IPv6 ACLs, or MAC ACLs.
Implicit Rules for IP and MAC ACLs
IP and MAC ACLs have implicit rules, which means that although these rules do not appear in the running configuration, the device applies them to traffic when no other rules in an ACL match. When you configure the device to maintain per-rule statistics for an ACL, the device does not maintain statistics for implicit rules.
All IPv4 ACLs include the following implicit rule:
deny ip any any
This implicit rule ensures that the device denies unmatched IP traffic.
All IPv6 ACLs include the following implicit rule:
deny ipv6 any any
This implicit rule ensures that the device denies unmatched IPv6 traffic.
| Note |  | 
All MAC ACLs include the following implicit rule:
deny any any protocol This implicit rule ensures that the device denies the unmatched traffic, regardless of the protocol specified in the Layer 2 header of the traffic.
Additional Filtering Options
You can identify traffic by using additional options. These options differ by ACL type. The following list includes most but not all additional filtering options:
-  
                                    		  
                                    IPv4 ACLs support the following additional filtering options: 
  -  
                                          				
                                          Layer 4 protocol
  -  
                                          				
                                          TCP and UDP ports
  -  
                                          				
                                          ICMP types and codes
  -  
                                          				
                                          IGMP types
  -  
                                          				
                                          Precedence level
  -  
                                          				
                                          Differentiated Services Code Point (DSCP) value
  -  
                                          				
                                          TCP packets with the ACK, FIN, PSH, RST, SYN, or URG bit set
  -  
                                          				
                                          Established TCP connections
  -  
                                          				
                                          Packet length
-  
                                          				
                                          
-  
                                    		  
                                    IPv6 ACLs support the following additional filtering options: 
  -  
                                          				
                                          Layer 4 protocol
  -  
                                          				
                                          Encapsulating Security Payload
  -  
                                          				
                                          Payload Compression Protocol
  -  
                                          				
