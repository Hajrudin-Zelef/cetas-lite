---
id: collect-261001-general-networking/general-networking/infrastructure-security-and-segmentation-5
title: "infrastructure-security-and-segmentation"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/infrastructure-security-and-segmentation.md
source_anchor: ""
source_lines: [292, 378]
sha256: 58e8a0dc36c4de05d80d4ef7c6c6d72ca0f10b55f1fb399079d48d94e544d7ca
---

# infrastructure-security-and-segmentation

A PACL is applied to an interface with the **ip access-group** *access-list* **in** command. Example 2-48 shows a PACL applied to interface Gi0/5 to block RDP and Telnet traffic.

#### **Example 2-48** *Applying a PACL*

`SW1(config)#**ip access-list extended pacl-5**
SW1(config-ext-nacl)#**deny tcp any any eq 3389**
SW1(config-ext-nacl)#**deny udp any any eq 3389**
SW1(config-ext-nacl)#**deny tcp any any eq 23**
SW1(config-ext-nacl)#**permit ip any any**
SW1(config-ext-nacl)#**exit**
SW1(config)#**int Gi0/5**
SW1(config-if)#**ip access-group pacl-5 in**`

A PACL can be further augmented with the IP Source Guard (IPSG) feature, which uses information from DHCP snooping to dynamically configure a port such that traffic is allowed only if it is sourced from an IP address bound to the interface. This is an effective method for blocking traffic with spoofed IP addresses. IPSG can be enabled with the **ip verify source** command on the interface and verified with the **show ip verify source** exec mode command. Example 2-49 shows IPSG enabled on the Gi0/9 interface, and only packets with source IP address 192.168.1.10 will be allowed out.

#### **Example 2-49** *Verifying IP Source Guard*

`SW1#**show ip verify source**
Interface   Filter-type   Filter-mode   IP-address    Mac-address   Vlan  Log
---------   -----------   -----------   ----------    -----------   ----  ---
Gi0/9      ip             active       192.168.1.10                 1     disabled`

The second method for filtering traffic at Layer 2 is to use VACLs. VACLs filter traffic that enters the VLAN from any source, including hosts in the VLAN. This makes it an effective tool for filtering traffic between hosts in the same VLAN as well as traffic being received from outside.

VACLs are configured using a VLAN access map. An access map is a series of **match** and **action** sets that define interesting traffic and action to be taken on them. Interesting traffic is defined by matching an IPv4, IPv6, or MAC access list. For each set of matched traffic, two actions can be defined: **forward** or **drop**. Optionally, the **log** keyword can be used with the **drop** action.

Access maps are defined with the **vlan access-map** *name sequence* command. Each map can have multiple sequences, with each sequence defining a **match** and **action** set. The access map can be applied to VLANs with the **vlan filter** *map-name* **vlan-list** *vlan-list* command, where *vlan-list* can be a single VLAN, a range of VLANs, or multiple VLANs as a comma-separated list.

Example 2-50 shows a VLAN access map applied for VLAN 1 to drop RDP and Telnet traffic. Note that the named ACLs, **rdp-traffic** and **telnet-traffic**, include a **permit** statement for the interesting traffic. Interesting traffic is always defined with a **permit** statement in the ACL so that it matches a VACL sequence. The VACL itself, though, is configured to drop the matched traffic.

#### **Example 2-50** *Creating and Applying a VACL*

`SW1(config)#**ip access-list extended rdp-traffic**
SW1(config-ext-nacl)#**permit tcp any any eq 3389**
SW1(config-ext-nacl)#**permit udp any any eq 3389**
SW1(config-ext-nacl)#**exit**
SW1(config)#**ip access-list extended telnet-traffic**
SW1(config-ext-nacl)#**permit tcp any any eq 23**
SW1(config-ext-nacl)#**exit**
SW1(config)#**ip access-list extended other-traffic**
SW1(config-ext-nacl)#**permit ip any any**
SW1(config-ext-nacl)#**exit**
SW1(config)#**vlan access-map vacl1 10**
SW1(config-access-map)#**match ip address rdp-traffic**
SW1(config-access-map)#**action drop log**
SW1(config-access-map)#**exit**
SW1(config)#**vlan access-map vacl1 20**
SW1(config-access-map)#**match ip address telnet-traffic**
SW1(config-access-map)#**action drop log**
SW1(config-access-map)#**exit**
SW1(config)#**vlan access-map vacl1 30**
SW1(config-access-map)#**match ip address other-traffic**
SW1(config-access-map)#**action forward**
SW1(config)#**vlan filter vacl1 vlan-list 1**`

### Security at the Layer 3 Data Plane

The data plane of a Layer 3 device uses information learned from the control plane protocols to route traffic between subnets. It uses the IP headers to determine where the intended destination is and routes the packet to the next hop. Given that a router works on a subnet level, it is easy to apply broad controls such as filtering and QoS. This section looks at some of the most common security features applied at the data layer of a Layer 3 device.

#### Traffic Filtering at Layer 3

The primary method of filtering traffic at Layer 3 is using access control lists (ACLs). ACLs are the Swiss Army knife of security with various uses. From broad traffic filtering based on source or destination address to granular filtering based on ports, protocol characteristics, or time, ACLs can be used in various ways. As mentioned before, they are even used to classify traffic for other security and non-security features.

ACLs are sequential lists of **permit** or **deny** statements, called access control entries (ACEs), that packets are evaluated against until the first match. When a packet matches an ACE, the specified action is taken. Using ACLs consists of two steps:

- **Step 1. Creating ACLs:** The first step in using ACLs is to create them. The type of ACL and its content determine the steps required to create it. While there are many variations of access lists, sometimes based on their usage, the five most common types—standard, extended, named, time-based, and reflexive—are discussed in the following sections.
- **Step 2. Applying ACLs:** ACLs need to be applied to interfaces, in the path of traffic, before they can be used. In addition to the interface, the direction in which the ACL needs to be applied has to be specified. Cisco routers allow one ACL per interface per direction.

Before looking into specific ACL types, it is important to know that ACLs on Cisco routers use something called inverse masks or wildcard masks instead of subnet masks to define source and destination traffic. As the name implies, an inverse mask is an inversed subnet mask. When broken down into binary numbers, each 0 bit in an inverse mask indicates that the corresponding address bit has to match exactly, while a 1 indicates that the corresponding address bit can be anything. For example, an IP address of 10.1.2.0 with an inverse mask of 0.0.0.255 means all bits of the first three octets must match exactly, while all bits of the last octet can be anything. Any IP address from 10.1.2.0 to 10.1.2.255 will match such an inverse mask.

An easy way to determine the inverse mask for a given subnet mask is to subtract it from 255.255.255.255. Example 2-51 shows a few examples of this.

#### **Example 2-51** *Finding the Inverse Mask from a Subnet Mask*

`255.255.255.255 - 255.255.255.0 = 0.0.0.255
255.255.255.255 - 255.255.255.128 = 0.0.0.127
255.255.255.255 - 255.255.240.0 = 0.0.15.255
255.255.255.255 - 255.128.0.0 = 0.127.255.255`

One important thing to remember about ACLs is that each of them has an implicit deny at the end. Traffic that is not permitted by any ACE in the ACL will be denied by the implicit deny at the end.

#### Standard ACLs

Standard ACLs are the simplest form of ACLs, and they filter based on only the source IP address of a packet. These ACLs can be numbered between 1 and 99 or 1300 and 1999. The source address can be a single host address, a subnet defined with an inverse mask, or simply all hosts defined with the *any* keyword. The syntax for creating a standard ACL is as follows:

`**access-list** *access-list-number* {**permit**|**deny**}{**host** *source*|*source inverse-mask*|**any**}[**log**]`

