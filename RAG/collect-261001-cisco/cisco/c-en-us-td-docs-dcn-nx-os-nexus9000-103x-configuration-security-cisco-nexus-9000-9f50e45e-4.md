---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e-4
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "throughput"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e.md
source_anchor: ""
source_lines: [326, 390]
sha256: cb88fd80acc50bfd54487a91913faae41f71c7ba4753a508050cb64b9e7122a8
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e

                                 The verified maximum number of ACEs is 100 for IPv4 traffic and an additional 100 for IPv6 traffic. Throughput may be impacted if more than this scale is applied.
Statistics and ACLs
The device can maintain global statistics for each rule that you configure in IPv4, IPv6, and MAC ACLs. If an ACL is applied to multiple interfaces, the maintained rule statistics are the sum of packet matches (hits) on all the interfaces on which that ACL is applied.
| Note | The device does not support interface-level ACL statistics. | 
For each ACL that you configure, you can specify whether the device maintains statistics for that ACL, which allows you to turn ACL statistics on or off as needed to monitor traffic filtered by an ACL or to help troubleshoot the configuration of an ACL.
The device does not maintain statistics for implicit rules in an ACL. For example, the device does not maintain a count of packets that match the implicit deny ip any any rule at the end of all IPv4 ACLs. If you want to maintain statistics for implicit rules, you must explicitly configure the ACL with rules that are identical to the implicit rules.
Atomic ACL Updates
By default, when a supervisor module of a Cisco Nexus 9000 Series device updates an I/O module with changes to an ACL, it performs an atomic ACL update. An atomic update does not disrupt traffic that the updated ACL applies to; however, an atomic update requires that an I/O module that receives an ACL update has enough available resources to store each updated ACL entry in addition to all pre-existing entries in the affected ACL. After the update occurs, the additional resources used for the update are freed. If the I/O module lacks the required resources, the device generates an error message and the ACL update to the I/O module fails.
If an I/O module lacks the resources required for an atomic update, you can disable atomic updates by using the no hardware access-list update atomic command; however, during the brief time required for the device to remove the preexisting ACL and implement the updated ACL, traffic that the ACL applies to is dropped by default.
If you want to permit all traffic that an ACL applies to while it receives a nonatomic update, use the hardware access-list update default-result permit command.
This example shows how to disable atomic updates to ACLs:
switch# config t 
switch(config)# no hardware access-list update atomic
This example shows how to permit affected traffic during a nonatomic ACL update:
switch# config t 
switch(config)# hardware access-list update default-result permit
This example shows how to revert to the atomic update method:
switch# config t 
switch(config)# no hardware access-list update default-result permit 
switch(config)# hardware access-list update atomic
Session Manager Support for IP ACLs
Session Manager supports the configuration of IP and MAC ACLs. This feature allows you to verify ACL configuration and confirm that the resources required by the configuration are available prior to committing them to the running configuration.
ACL TCAM Regions
You can change the size of the ACL ternary content addressable memory (TCAM) regions in the hardware.
On Cisco Nexus 9300 and 9500 Series switches, the egress TCAM size is 1K, divided into four 256 entries. On other Cisco Nexus 9300 and 9500 Series switches, the ingress TCAM size is 4K, divided into eight 256 slices and four 512 slices. A slice is the unit of allocation. A slice can be allocated to one region only. For example, a 512-size slice cannot be used to configure two features of size 256 each. Similarly, a 256-size slice cannot be used to configure two features of size 128 each. The IPv4 TCAM regions are single wide. The IPv6, QoS, MAC, CoPP, and system TCAM regions are double wide and consume double the physical TCAM entries. For example, a logical region size of 256 entries actually consumes 512 physical TCAM entries.
You can create IPv6, port ACLs, VLAN ACLs, and router ACLs, and you can match IPv6 and MAC addresses for QoS. However, Cisco NX-OS cannot support all of them simultaneously. You must remove or reduce the size of the existing TCAM regions (TCAM carving) to enable the IPv6, MAC, or other desired TCAM regions. For every TCAM region configuration command, the system evaluates if the new change can be fit in the TCAM. If not, it reports an error, and the command is rejected. You must remove or reduce the size of existing TCAM regions to make room for new requirements.
On Cisco Nexus 9200 Series switches, the egress TCAM size is 2K, and the ingress TCAM size is 4K. The concepts of TCAM slices and single- and double-wide regions do not apply to these switches. For example, the ing-ifacl region can host IPv4, IPv6, or MAC type entries. IPv4 and MAC types occupy one TCAM entry whereas IPv6 types occupy two TCAM entries.
For N9K-X9636C-RX, when PACL uses external TCAM region, the internal TCAM needs to take 2K for ifacl and the ingress RACL-IPv4 can take upto 2044. Additional four entries are required when egress PACL external TCAM region is used.
ACL TCAM region sizes have the following guidelines and limitations:
- 
                                 				
                                 To enable RACL or PACL on existing TCAM regions, you must carve the TCAM region beyond 12, 288.
- 
                                 				
                                 On Cisco Nexus 9300 Series switches, the X9536PQ, X9564PX, and X9564TX line cards are used to enforce the QoS classification policies applied on 40G ports. It has 768 TCAM entries available for carving in 256-entry granularity. These region names are prefixed with "ns-".
- 
                                 				
                                 For the X9536PQ, X9564PX, and X9564TX line cards, only the IPv6 TCAM regions consume double-wide entries. The rest of the TCAM regions consume single-wide entries.
- 
                                 				
                                 When a VACL region is configured, it is configured with the same size in both the ingress and egress directions. If the region size cannot fit in either direction, the configuration is rejected.
- 
                                 				
                                 				
                                 On Cisco Nexus 9200 Series switches, the ing-sup region occupies a minimum size of 512 entries, and the egr-sup region occupies a minimum size of 256 entries. These regions cannot be configured to lesser values. Any region size can be carved with a value only in multiples of 256 entries (with the exception of the span region, which can be carved only in multiples of 512 entries).
- 
                                 				
                                 RACL v6, CoPP, and multicast have default TCAM sizes and these TCAM sizes must be non-zero on the following Cisco Nexus 9504 and Cisco Nexus 9508 line cards to avoid line card failure during reload: 
  - 
                                       						
                                       N9K-X96136YC-R
  - 
                                       						
                                       N9K-X9636C-RX
  - 
                                       						
                                       N9K-X9636Q-R
  - 
                                       						
                                       N9K-X9636C-R
- 
                                       						
                                       
- 
                                 				
