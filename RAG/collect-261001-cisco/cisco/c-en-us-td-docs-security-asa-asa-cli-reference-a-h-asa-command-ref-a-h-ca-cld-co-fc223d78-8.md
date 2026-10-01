---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-8
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [953, 1069]
sha256: e95e2c0b0dc2c27ade6cc676b774f5821c3e06ebc482d4882f33eee322af833d
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

                              You can precede the capture  command with cluster exec  to issue the capture  command on one unit and run the command in all the other units at the same time. After you have performed cluster-wide capture,
                              to copy the same capture file from all units in the cluster at the same time to a TFTP server, enter the cluster exec copy  command on the master unit.
                           
                           
ciscoasa# cluster exec capture 
capture_name arguments
ciscoasa# cluster exec copy
 /pcap capture
: cap_name
 tftp
://location
/path
/filename
.pcap
                           
                              Multiple PCAP files, one from each unit, are copied to the TFTP server. The destination capture file name is automatically
                              attached with the unit name, such as filename_A.pcap, filename_B.pcap, and so on. In this example, A and B are cluster unit
                              names.
                           
                           
                           
                              When you capture traces on cluster units, they are persistent on each cluster node until you manually clear them from the
                              buffer. Decrypted IPsec packets are captured once they enter ASA. The captured packet includes both normal and decapsulated
                              traffic.
                           
                           
                           
                              
                                 |  Note |  A different destination name is generated if you add the unit name at the end of the filename.  | 
                           
                                 Limitations
                           
                           
                              The following are some of the limitations of the capture feature. Most of the limitations are caused by the distributed nature
                              of the ASA architecture and by the hardware accelerators that are being used in the ASA. 
                           
                           
                           
                              
                              - 
                                 
                                    You can configure captures on the cluster control link within a context; only the packet that is associated with the context
                                    sent in the cluster control link is captured.
                                 
- 
                                  For a shared VLAN, the following guidelines apply:  
                                    
                                    
  - 
                                       
                                        You can only configure one capture for the VLAN; if you configure a capture in multiple contexts on the shared VLAN, then
                                          only the last capture that was configured is used. 
                                       
  - 
                                       
                                        If you remove the last-configured (active) capture, no captures become active, even if you have previously configured a capture
                                          in another context; you must remove and readd the capture to make it active. 
                                       
  - 
                                       
                                        All traffic that enters the interface to which the capture is attached (and that matches the capture access list) is captured,
                                          including traffic to other contexts on the shared VLAN. 
                                       
  - 
                                       
                                        Therefore, if you enable a capture in Context A for a VLAN that is also used by Context B, both Context A and Context B ingress
                                          traffic are captured. 
                                       
- 
                                 
                                    For egress traffic, only the traffic of the context with the active capture is captured. The only exception is when you do
                                    not enable the ICMP inspection (therefore the ICMP traffic does not have a session in the accelerated path). In this case,
                                    both ingress and egress ICMP traffic for all contexts on the shared VLAN is captured. 
                                 
- 
                                 
                                    Configuring a capture typically involves configuring an access list that matches the traffic that needs to be captured. After
                                    an access list that matches the traffic pattern is configured, then you need to define a capture and associate this access
                                    list to the capture, along with the interface on which the capture needs to be configured. Note that a capture only works
                                    if an access list and an interface are associated with a capture for capturing IPv4 traffic. The access list is not required
                                    for IPv6 traffic.
                                 
- 
                                 
                                    For the ASA CX module traffic, captured packets contain an additional AFBP header that your PCAP viewer might not understand;
                                    be sure to use the appropriate plugin to view these packets.
                                 
- 
                                 
                                    For inline SGT tagged packets, captured packets contain an additional CMD header that your PCAP viewer might not understand.
                                 
- 
                                 
                                 When performing a switch packet capture on Secure Firewall 
                                    SGT-tagged packets are not captured if you apply a filter for SGT-tagged packets. However, SGT traffic can be captured on
                                    an interface if no filtering is applied.
                                 
- 
                                 
                                    If there is no ingress interface and therefore no global interface, packets sent on the backplane are treated as control packets
                                    in the system context. These packets bypass the access list check and are always captured. This behavior applies in both single
                                    mode and multiple context mode.
                                 
- 
                                 
                                    The show capture  command shows the correct reason when capturing a specific asp-drop. However, the show capture  command does not show the correct reason when capturing all asp-drops.
                                 
- 
                                 
                                 The switch packet capture feature for Secure Firewall 6100 series has these limitations:
- 
                                 
                                 The switch packet capture feature for Secure Firewall 4200 and Secure Firewall 6100 series has these limitations: 
                                    
                                    
  - 
                                       
