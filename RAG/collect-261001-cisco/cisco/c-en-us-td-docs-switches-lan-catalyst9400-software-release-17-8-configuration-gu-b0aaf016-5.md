---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016-5
title: "c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016.md
source_anchor: ""
source_lines: [275, 316]
sha256: ad3adbac3185ad72f3afc650c7a9abacdbddabf2b00ad27e63a1658982995886
---

# c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016

                                 When AOT is enabled, the show ip nat translations command will not give visibility into all the NAT flows being translated and forwarded.
Using Application-Level Gateways with NAT
NAT performs translation services on any TCP/UDP traffic that does not carry source and destination IP addresses in the application data stream. Protocols that do not carry the source and destination IP addresses include HTTP, TFTP, telnet, archie, finger, Network Time Protocol (NTP), Network File System (NFS), remote login (rlogin), remote shell (rsh) protocol, and remote copy (rcp).
NAT Application-Level Gateway (ALG) enables certain applications that carry address/port information in their payloads to function correctly across NAT domains. In addition to the usual translation of address/ports in the packet headers, ALGs take care of translating the address/ports present in the payload and setting up temporary mappings.
Best Practices for NAT Configuration
- 
                                    
                                    In cases where two static NAT rules overlap with each other, such as static subnet translation rule overlapping with a corresponding fully qualified static rule, then the more specific rule should be configured ahead of the other.
- 
                                    
                                    In cases where both static and dynamic rules are configured, ensure that the local addresses specified in the rules do not overlap. If such an overlap is possible, then the ACL associated with the dynamic rule should exclude the corresponding addresses used by the static rule. Similarly, there must not be any overlap between the global addresses as this could lead to undesired behavior.
- 
                                    
                                    VRF to Global translation functionality considers the NAT outside interface to be associated with default or global VRF. Therefore, placing the NAT outside interface in a non-default VRF is not recommended while performing VRF aware NAT.
- 
                                    
                                    Do not employ loose filtering such as permit ip any any in an ACL associated with NAT rule as this could result in unwanted packets being translated.
- 
                                    
                                    Do not share an address pool across multiple NAT rules.
- 
                                    
                                    Do not define the same inside global address in Static NAT and Dynamic Pool. This action can lead to undesirable results.
- 
                                    
                                    Exercise caution while modifying the default timeout values associated with NAT. Small timeout values could result in high CPU usage.
- 
                                    
                                    Exercise caution while manually clearing the translation entries as this could result in the disruption of application sessions.
- 
                                    
                                    Follow these steps before you make NAT configuration changes during active translations. 
  - 
                                          
                                          Stop the ingress and egress of traffic matching the given configuration. This may require applying an appropriate ACL filter or shutting down the given interfaces.
  - 
                                          
                                          Clear any existing translation entries that correspond to the given configuration.
  - 
                                          
                                          Make the desired configuration change and re-enable the stopped traffic.
-
