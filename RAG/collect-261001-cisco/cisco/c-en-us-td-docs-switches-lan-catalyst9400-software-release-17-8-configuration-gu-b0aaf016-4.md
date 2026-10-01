---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016-4
title: "c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "throughput"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016.md
source_anchor: ""
source_lines: [199, 274]
sha256: 7c06e6d22b0d5610348da553c27185e7de5f1eaf996ca61b9f6f8d26a73c066d
---

# c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016

                                          Packets that require both inside and outside translation.
- 
                                          
                                          
- 
                                    
                                    The maximum number of sessions that can be translated and forwarded in the hardware in an ideal setting is limited to 7000. Additional flows that require translation are handled in the software data plane at a reduced throughput. Note 
 Each translation consumes two entries in TCAM. 
- 
                                    
                                    A configured NAT rule might fail to get programmed into the hardware owing to resource constraint. This could result in packets that correspond to the given rule to get forwarded without translation.
- 
                                    
                                    ALG support is currently limited to FTP, TFTP, and ICMP protocols. Also, although TCP SYN, TCP FIN, and TCP RST are not part of ALG traffic, they are processed as part of ALG traffic.
- 
                                    
                                    Dynamically created NAT flows age out after a period of inactivity.
- 
                                    
                                    
                                    Policy Based Routing (PBR) and NAT are not supported on the same interface. PBR and NAT work together only if they are configured on different interfaces.
- 
                                    
                                    Port Channel is not supported in NAT configuration.
- 
                                    
                                    NAT does not support translation of fragmented packets.
- 
                                    
                                    Bidirectional Forwarding Detection (BFD) sessions may fail if they are configured to operate using the same address that is used for dynamic NAT. To avoid a conflict that arises when both BFD and Dynamic NAT are configured on the device, use an address that does not overlap with NAT. If you must configure BFD and dynamic NAT overloading on the same interface, deploy a pool-based dynamic NAT overload configuration. Ensure that you do not use the chosen NAT pool address for BFD even in this scenario.
- 
                                    
                                    
                                    
                                    Equal-cost multi-path routing (ECMP) is not supported with NAT.
- 
                                    
                                    
                                    When Flexible NetFlow and Network Address Translation (NAT) are configured on an interface: 
  - 
                                          
                                          Flexible NetFlow will display and export the actual flow details; but not the translated flow details. Application-level gateway (ALG) flow details are not part of the actual flow details that are exported.
  - 
                                          
                                          If the ALG traffic gets translated through the CPU, Flexible NetFlow will display and export the translated flow details for the ALG traffic.
- 
                                          
                                          
- 
                                    
                                    
                                    Explicit deny access control entry (ACE) in NAT ACL is not supported. Only explicit permit ACE is supported.
- 
                                    
                                    
                                    NETCONF configuration fails if it is configured to operate using the same IP address that is used for configuring NAT using interface overload.
- 
                                    
                                    NAT is not supported over GRE tunnels.
- 
                                    
                                    When both the ingress and egress NAT interfaces are on the same switch, hardware NAT entries in TCAM may go out of sync on the remote or standby switch after stateful switchover (SSO). When AOT is enabled, this can occur on the remote switch after NAT timeout.
Performance and Scale Numbers for NAT
NAT module is capable of performing translation and forwarding in the hardware at line-rate, by programming the relevant hardware tables with the forwarding and rewrite information. You can configure a NAT-focused resource allocation scheme to obtain increased NAT throughput.
Configure SDM template NAT to achieve better performance and scale number. Refer Configuring Switch Database Management (SDM) Template
The maximum number of TCAM flows that are available in the hardware is 2000. However, for C9400-SUP-1XL, you can enhance the scale by configuring SDM template to achieve better performance and scale numbers, with which a maximum scale of 14000 can be achieved.
| Note | Using Address Only Translation optimizes the handling of flows and enhances the scale of the NAT feature. | 
Address Only Translation
Address only Translation (AOT) functionality can be employed in situations that require only the address fields to be translated and not the transport ports. In such settings, enabling AOT functionality significantly increases the number of flows that can be translated and forwarded in the hardware at line-rate. This improvement is brought about by optimizing the usage of various hardware resources associated with translation and forwarding. A typical NAT focused resource allocation scheme sets aside 14000 TCAM entries for performing hardware translation. This places a strict upper limit on the number of flows that can be translated and forwarded at line-rate. Under AOT scheme, the usage of TCAM resource is highly optimized thereby enabling the accommodation of more number of flows in the TCAM tables and this provides a significant improvement in the hardware translation and forwarding scale. AOT can be very effective in situations where majority of the flows are destined to a single or a small set of destinations. Under such favourable conditions, AOT can potentially enable line-rate translation and forwarding of all the flows originating from the given end-point(s). AOT functionality is disabled by default. It can be enabled using the no ip nat create flow-entries command. The existing dynamic flow can be cleared using the clear ip nat translation command. The AOT feature can be disabled using the ip nat create flow-entries command.
Restrictions for Address Only Translation
- 
                                 
                                 
                                 AOT feature is expected to function correctly only in translation scenarios corresponding to simple inside static and inside dynamic rules. The simple static rule must be of the type ip nat inside source static local-ip global-ip , and the dynamic rule must be of the type ip nat inside source list access-list pool name .
- 
                                 
