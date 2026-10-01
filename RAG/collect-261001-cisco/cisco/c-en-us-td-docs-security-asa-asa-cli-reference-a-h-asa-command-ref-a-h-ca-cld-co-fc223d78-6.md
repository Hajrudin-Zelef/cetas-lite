---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-6
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [728, 821]
sha256: 883e1cec207bae8ded376aa1591a3474e787e71ad3f243e6c0eb4ace6c29739e
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

                                    |  7.0(8) |  Added the all option to capture all packets that the ASA drops. | 
                                 
                                 
                                    
                                    |  7.2(1) |  This command was modified to include the following options: trace                                               trace_count, match                                            prot, real-time,                                                                                          host                                            ip, any ,                                               mask,  and operator.                                                                                    | 
                                 
                                 
                                    
                                    |  8.0(2) |  This command was modified to update the path to capture contents. | 
                                 
                                 
                                    
                                    |  8.4(1) |  The new type keywords ikev1 and ikev2 were added. | 
                                 
                                 
                                    
                                    |  8.4(2) |  Additional detail was added to the output for IDS. | 
                                 
                                 
                                    
                                    |  8.4(4.1) |  The asa_dataplane  option was added to support traffic over the backplane to the ASA CX module.                                         | 
                                 
                                 
                                    
                                    |  9.0(1) |  The cluster,                                              cluster                                              exec,                                               and reinject-hide  keywords were added. The new                                              type  option lacp  was added. Support for multiple-context mode was added for ISAKMP.                                         | 
                                 
                                 
                                    
                                    |  9.1(3) |  Supports filtering of packets captured on the ASA CX backplane with the asa_dataplane  option.                                         | 
                                 
                                 
                                    
                                    |  9.2(1) |  The asa_dataplane  option was extended to support the ASA FirePOWER module.                                         | 
                                 
                                 
                                    
                                    |  9.3(1) |  The inline-tag                                            tag keyword-argument pair was added to support the SGT plus Ethernet Tagging feature.                                         | 
                                 
                                 
                                    
                                    |  9.6(2) |  Packet capture of type                                              asp-drop                                               supports ACL and match filtering.                                         | 
                                 
                                 
                                    
                                    |  9.7(1) |  Added the stop  keyword to manually stop and start the packet capture.                                         | 
                                 
                                 
                                    
                                    |  9.8(1) |  This command was updated to store the contents of all the active captures to files on flash or disks at the time of box crash. | 
                                 
                                 
                                    
                                    |  9.9(1) |  Support for capturing clustering persistent tracing and decrypted packets. New options were added: persist  and include-decrypted .                                          In addition, the ethernet-type                                              ipx  was removed, because IPX corresponds to 3 separate ethernet-types. Instead, use the hexadecimal value of the IPX type you                                           want to capture.                                         | 
                                 
                                 
                                    
                                    |  9.10(1) |  Added the any4 and any6 keywords for the match  option to capture IPv4 and IPv6 network traffic respectively.                                         | 
                                 
                                 
                                    
                                    |  9.12(1) |  Added cp-cluster                                               to capture control packets on cluster interface.                                         | 
                                 
                                 
                                    
                                    | 9.18(1) |  Included real-time                                                keyword to enable real-time switch packet capture.                                         | 
                                 
                                 
                                    
                                    | 9.20(1) | The direction keyword was added to capture switch traffic that flows in egress, ingress, or both directions. This keyword is applicable only for Secure Firewall 4200 model devices.  | 
                                 
                                 
                              
                           
Usage Guidelines
                           
                           
                              Capturing packets is useful when troubleshooting connectivity problems or monitoring suspicious activity. You can create multiple
                              captures. The capture  command is not saved to the running configuration, and is not copied to the standby unit during failover.
                           
                           
                           
                              The ASA is capable of tracking all IP traffic that flows across it and of capturing all the IP traffic that is destined to
                              it, including all the management traffic (such as SSH and Telnet traffic). 
                           
                           
                           
                              The ASA architecture consists of three different sets of processors for packet processing; this architecture poses certain
                              restrictions on the capability of the capture feature. Typically most of the packet forwarding functionality in the ASA is
                              handled by the two front-end network processors, and packets are sent to the control-plane general-purpose processor only
                              if they need application inspection. The packets are sent to the session management path network processor only if there is
                              a session miss in the accelerated path processor. 
                           
                           
                           
