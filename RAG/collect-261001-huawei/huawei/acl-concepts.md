---
id: collect-261001-huawei/huawei/acl-concepts
title: "acl-concepts"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/acl-concepts.md
source_anchor: ""
source_lines: [1, 116]
sha256: 61b1e6c8591ce75862ce3d6a979db96c8645d9b9a7b9421c5ef21c898afcae0d
---

# acl-concepts

## Types of IPv4 ACLs (4.4)

This section compares IPv4 standard and extended ACLs.

### Standard and Extended ACLs (4.4.1)

The previous sections describe the purpose of ACLs as well as guidelines for ACL creation. This section covers standard and extended ACLs and named and *numbered ACLs*, and it provides examples of placement of these ACLs.

There are two types of IPv4 ACLs:

- *Standard ACLs***:** These ACLs permit or deny packets based only on the source IPv4 address.
- *Extended ACLs***:** These ACLs permit or deny packets based on the source IPv4 address and destination IPv4 address, protocol type, source and destination TCP or UDP ports, and more.

For example, Example 4-3 shows how to create a standard ACL. In this example, ACL 10 permits hosts on the source network 192.168.10.0/24. Because of the implied “deny any” at the end, all traffic except for traffic coming from the 192.168.10.0/24 network is blocked with this ACL.

#### **Example 4-3** Standard ACL Example

`R1(config)# **access-list 10 permit 192.168.10.0 0.0.0.255**
R1(config)#`

In Example 4-4, the extended ACL 100 permits traffic originating from any host on the 192.168.10.0/24 network to any IPv4 network if the destination host port is 80 (HTTP).

#### **Example 4-4** Extended ACL Example

`R1(config)# **access-list 100 permit tcp 192.168.10.0 0.0.0.255 any eq www**
R1(config)#`

Notice that the standard ACL 10 is only capable of filtering by source address, while the extended ACL 100 is filtering on the source and destination Layer 3 and Layer 4 protocol (for example, TCP) information.

### Numbered and Named ACLs (4.4.2)

For IPv4, there are both numbered and *named ACLs*.

#### Numbered ACLs

ACLs 1 to 99 and 1300 to 1999 are standard ACLs, while ACLs 100 to 199 and 2000 to 2699 are extended ACLs, as shown in Example 4-5.

#### **Example 4-5** Available ACL Numbers

`R1(config)# **access-list ?**
  <1-99>       IP standard access list
  <100-199>     IP extended access list
  <1100-1199>   Extended 48-bit MAC address access list
  <1300-1999>   IP standard access list (expanded range)
  <200-299>     Protocol type-code access list
  <2000-2699>   IP extended access list (expanded range)
  <700-799>     48-bit MAC address access list
  rate-limit    Simple rate-limit specific access list
  template      Enable IP template acls
Router(config)# **access-list**`

#### Named ACLs

Using named ACLs is the preferred method when configuring ACLs. You can name standard and extended ACLs to provide information about the purpose of each ACL. For example, the extended ACL name FTP-FILTER is far easier to identify than the ACL number 100.

The **ip access-list** global configuration command is used to create a named ACL, as shown in Example 4-6.

#### **Example 4-6** Example of a Named ACL

`R1(config)# **ip access-list extended FTP-FILTER**
R1(config-ext-nacl)# **permit tcp 192.168.10.0 0.0.0.255 any eq ftp**
R1(config-ext-nacl)# **permit tcp 192.168.10.0 0.0.0.255 any eq ftp-data**
R1(config-ext-nacl)#`

The following are the general rules to follow for named ACLs:

- Assign a name to identify the purpose of the ACL.
- Names can contain alphanumeric characters.
- Names cannot contain spaces or punctuation.
- It is suggested that a name be written in CAPITAL LETTERS.
- Entries can be added or deleted within an ACL.

### Where to Place ACLs (4.4.3)

Every ACL should be placed where it has the greatest impact on efficiency.

Figure 4-5 illustrates where standard and extended ACLs should be located in an enterprise network.

Say that the objective is to prevent traffic originating in the 192.168.10.0/24 network from reaching the 192.168.30.0/24 network. Extended ACLs should be located as close as possible to the source of the traffic to be filtered. This way, undesirable traffic is denied close to the source network, without crossing the network infrastructure.

Standard ACLs should be located as close to the destination as possible. If a standard ACL were placed at the source of the traffic, the “permit” or “deny” would occur based on the given source address, regardless of the traffic destination.

Placement of an ACL and, therefore, the type of ACL used, may also depend on a variety of factors, as listed in Table 4-11.

#### **Table 4-11** ACL Placement Factors

| **Factors Influencing ACL Placement** | **Explanation** | 
| The extent of organizational control |  | 
| Bandwidth of the networks involved |  | 
| Ease of configuration |  | 

### Standard ACL Placement Example (4.4.4)

Following the guidelines for ACL placement, standard ACLs should be located as close to the destination as possible.

In Figure 4-6, the administrator wants to prevent traffic originating in the 192.168.10.0/24 network from reaching the 192.168.30.0/24 network.

Following the basic placement guidelines, the administrator would place a standard ACL on router R3. There are two possible interfaces on R3 to which to apply the standard ACL:

- **R3 S0/1/1 interface (inbound):** The standard ACL can be applied inbound on the R3 S0/1/1 interface to deny traffic from the .10 network. However, it would also filter .10 traffic to the 192.168.31.0/24 (.31 in this example) network. Therefore, the standard ACL should not be applied to this interface.
- **R3 G0/0 interface (outbound):** The standard ACL can be applied outbound on the R3 G0/0/0 interface. This will not affect other networks that are reachable by R3. Packets from the .10 network will still be able to reach the .31 network. This is the best interface to place the standard ACL to meet the traffic requirements.

### Extended ACL Placement Example (4.4.5)

Extended ACLs should be located as close to the source as possible to prevent unwanted traffic from being sent across multiple networks only to be denied when it reaches its destination.

However, an organization can only place ACLs on devices that it controls. Therefore, the extended ACL placement must be determined in the context of where organizational control extends.

In Figure 4-7, for example, Company A wants to deny Telnet and FTP traffic to Company B’s 192.168.30.0/24 network from its 192.168.11.0/24 network while permitting all other traffic.

There are several ways to accomplish these goals. An extended ACL on R3 would accomplish the task, but the administrator does not control R3. In addition, this solution would allow unwanted traffic to cross the entire network, only to be blocked at the destination, which would affect overall network efficiency.

The solution is to place on R1 an extended ACL that specifies both source and destination addresses. There are two possible interfaces on R1 to apply the extended ACL:

- **R1 S0/1/0 interface (outbound):** The extended ACL can be applied outbound on the S0/1/0 interface. However, this solution would process all packets leaving R1, including packets from 192.168.10.0/24.
- **R1 G0/0/1 interface (inbound):** The extended ACL can be applied inbound on the G0/0/1, and only packets from the 192.168.11.0/24 network are subject to ACL processing on R1. Because the filter is to be limited to only those packets leaving the 192.168.11.0/24 network, applying the extended ACL to G0/1 is the best solution.
