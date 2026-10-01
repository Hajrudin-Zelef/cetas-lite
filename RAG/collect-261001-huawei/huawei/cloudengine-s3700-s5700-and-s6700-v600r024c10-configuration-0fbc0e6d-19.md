---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-19
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [1986, 2136]
sha256: 83c86ddfc00e9b38196ded918f382c355f324cba7360f418ea33a1e143abdcd3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 When LDP is enabled to distribute labels to all peers, P2 sends a Label Mapping
                 message associated with the route to P1 after receiving the Label Mapping
                 message from P1, which allows LDP to generate a liberal LSP on P1. If the link
                 between P1 and P3 becomes faulty, the route from PE1 to PE3 is switched from
                 PE1 -> P1 -> P3 -> PE3 to PE1 -> P1 -> P2 -> P4 -> P3 -> PE3, P2 becomes the
                 downstream of P1, and the liberal LSP changes to a normal LSP, resulting in faster
                 LSP convergence.


                 Figure 3-4 Networking topology of distributing labels to all peers by LDP




                 To configure some LSRs to distribute labels only to upstream peers, configure split
                 horizon.


3.3 Configuration Precautions for MPLS LDP

3.4 Default Settings for MPLS LDP
                 Table 3-1 describes the default settings for MPLS LDP.


Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                            34
MPLS Configuration
MPLS Configuration                                                         3 MPLS LDP Configuration


                 Table 3-1 Default settings for MPLS LDP

                  Parameter                                 Default Setting

                  Global MPLS capability                    Disabled

                  Global MPLS LDP capability                Disabled

                  Link Hello send timer                     5s

                  Link Hello hold timer                     15s

                  Target Hello send timer                   15s

                  Target Hello hold timer                   45s

                  Keepalive send timer                      15s

                  Keepalive hold timer                      45s

                  Exponential backoff timer                 Initial value: 15s; maximum value:
                                                            120s

                  LDP-OSPF synchronization                  Disabled

                  LDP-IS-IS synchronization                 Disabled

                  LDP GR                                    Disabled




3.5 Configuring Static LSPs

3.5.1 Understanding Static LSPs
                 Static LSPs can only be manually configured by administrators, but cannot be
                 dynamically established using LDP. As such, static LSPs are applicable to networks
                 with simple and stable topologies.

                 The path that an IP packet passes through on an MPLS network is called an LSP.
                 An LSP can be statically established through manual configuration or dynamically
                 established through a label distribution protocol.


Functions of Static LSPs
                 Generally, LDP is used to establish LSPs on an MPLS network. LDP establishes LSPs
                 based on routing information. If LDP fails, MPLS traffic may be lost. To protect key
                 data or important services, static LSPs can be configured.

                 The establishment of static LSPs does not require a label distribution protocol or
                 exchange of control packets, consuming fewer resources. However, static LSPs
                 cannot be dynamically adjusted according to network topology changes and
                 require administrator intervention. Therefore, they are applicable to small
                 networks with a simple and stable topology.



Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                               35
MPLS Configuration
MPLS Configuration                                                                          3 MPLS LDP Configuration


                 When configuring a static LSP, the administrator needs to manually allocate labels
                 for each LSR. The outgoing label value of the previous node must be equal to the
                 incoming label value of the current node.
                 In real-world applications, you can establish static LSPs on the backbone network
                 to facilitate L2VPN and L3VPN service deployment.

3.5.2 Configuring Static LSPs
                 To configure a static LSP, perform manual configurations on the ingress, transit
                 nodes, and egress of the static LSP.

Context
                 The establishment of static LSPs does not require a label distribution protocol or
                 exchange of control packets. As such, static LSPs consume fewer resources and are
                 applicable to small networks with a simple and stable topology. Static LSPs cannot
                 dynamically adapt to network topology changes. Once the network topology
                 changes, an administrator must modify configurations on each LSR of each
                 involved LSP so that the LSPs can work properly.
                 Static LSPs and static CR-LSPs share the same label space (16–1023). When
                 configuring a static LSP, ensure that the outgoing label on the previous node is the
                 same as the incoming label on the next hop.

Procedure
                 ●      Perform the following configurations on the ingress.
                        a.   Enter the system view.
                             system-view
                        b.   Configure the local node as the ingress of the LSP.
                             static-lsp ingress lsp-name destination ip-address { mask-length | mask } { nexthop next-hop-
                             address | outgoing-interface interface-type interface-number } out-label out-label

                             To modify the destination destination-address, nexthop next-hop-
                             address, outgoing-interface interface-type interface-number, and out-
                             label out-label parameter settings, run the static-lsp ingress command
                             to set new values directly, not requiring you to clear previous settings
                             using the undo static-lsp ingress command.

                                   NOTE

                                  You are advised to specify a next hop for a static LSP. Ensure that the local
                                  routing table contains a routing entry that exactly matches the specified
                                  destination IP address and next-hop IP address.
                                  If an Ethernet interface is used as an outbound interface of an LSP, you must
                                  specify the nexthop next-hop-address parameter to ensure normal traffic
                                  forwarding on the LSP.
                 ●      Perform the following configurations on each transit node.
                        a.   Enter the system view.
                             system-view
                        b.   Configure the local node as the transit node of the LSP.
                             static-lsp transit lsp-name [ incoming-interface in-interface-type in-interface-number ] in-
                             label in-label { nexthop next-hop-address | outgoing-interface out-interface-type out-
                             interface-number } out-label out-label


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                                  36
MPLS Configuration
MPLS Configuration                                                                          3 MPLS LDP Configuration


