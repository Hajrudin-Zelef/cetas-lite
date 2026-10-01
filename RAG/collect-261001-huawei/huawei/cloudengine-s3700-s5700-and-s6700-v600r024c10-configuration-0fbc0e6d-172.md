---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-172
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [24954, 25125]
sha256: f773f9442bba423be35d92fdd8b57d5881647f84caa4f273f14a5a8dc5089250
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             414
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


                 Figure 4-39 BFD for CR-LSP before and after a link fault occurs




4.25.2 Enabling BFD Globally

Prerequisites
                 Before configuring dynamic BFD for CR-LSP, complete one of the following tasks:

                 ●      Configure a dynamic MPLS TE tunnel.
                 ●      Configure CR-LSP backup.


Context
                 Before configuring dynamic BFD for CR-LSP, enable BFD globally on the ingress
                 and egress of a tunnel.

                         NOTE

                        In BFD for LSP, the forwarding modes for the forward and return paths can differ. For
                        example, packets may travel to the destination over an LSP, but to the source over an IP
                        path. However, the forward and return paths must traverse the same link. If not,
                        pinpointing the exact path where a fault occurs becomes challenging. Before deploying
                        BFD, ensure that the forward and reverse paths are over the same link so that BFD can
                        correctly identify the faulty path.

                 Perform the following configuration on the ingress and egress of an MPLS TE
                 tunnel.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         415
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enable BFD globally.
                 bfd

                 You can perform BFD-related configuration only after enabling BFD globally using
                 the bfd command.

                 ----End

4.25.3 Configuring the Function of Dynamically Creating BFD
Sessions on an Ingress
Context
                 You can enable the function of dynamically creating BFD sessions for TE tunnels
                 using either of the following methods:
                 ●      Enable the function of dynamically creating BFD sessions globally.
                        This method is recommended when the ingress needs to automatically create
                        BFD sessions for most TE tunnels.
                 ●      Enable the function of dynamically creating BFD sessions on specific
                        tunnel interfaces.
                        This method is recommended when the ingress needs to automatically create
                        BFD sessions for a small number of TE tunnels.
                 Perform the following configuration on the ingress of an MPLS TE tunnel.

Procedure
                 ●      Enable the function of dynamically creating BFD sessions globally.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the MPLS view.
                             mpls
                        c.   Enable the function of dynamically creating BFD sessions for MPLS TE
                             tunnels.
                             mpls te bfd enable

                             After this command is run in the MPLS view, dynamic BFD for CR-LSP is
                             enabled on all tunnel interfaces, excluding the interfaces on which
                             dynamic BFD for CR-LSP is blocked.
                        d.   (Optional) Block BFD for CR-LSP.
                             quit
                             interface tunnel interface-number
                             mpls te bfd block

                             If certain TE tunnels do not require BFD for CR-LSP, block this function on
                             the tunnel interfaces of these tunnels.
                 ●      Enable the function of dynamically creating BFD sessions on specific tunnel
                        interfaces.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                          416
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS TE tunnel interface view.
                             interface tunnel interface-number

                        c.   Enable the function of dynamically creating BFD sessions for MPLS TE
                             tunnels.
                             mpls te bfd enable

                             This command run in the tunnel interface view takes effect only on the
                             current tunnel interface.

                 ----End

4.25.4 Configuring the Function of Passively Creating BFD
Sessions on an Egress

Context
                 An LSP is unidirectional. Once the ingress of an LSP establishes a BFD session, it
                 sends an LSP ping packet to the egress. The egress can only automatically initiate
                 a BFD session upon receiving this LSP ping packet. Throughout the BFD session
                 establishment process, the role of the egress is passive.

                 Perform the following configuration on the egress of an MPLS TE tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the BFD view.
                 bfd

         Step 3 Enable the function of passively creating BFD sessions.
                 mpls-passive

                 After this command is run, a BFD session will be established only after the egress
                 receives an LSP ping request packet that carries a BFD TLV from the ingress.

         Step 4 (Optional) Change the destination UDP port number of a specified passive BFD
                session.
                 passive-session udp-port 3784 peer peer-ip

                 The default destination UDP port number is 4784 for a passive BFD session.

         Step 5 (Optional) Change the detection multiplier of a passive BFD session.
                 passive-session detect-multiplier multiplier-value peer peerip-value

                 The default detection multiplier of a passive BFD session is 3.

                 ----End



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 417
MPLS Configuration
MPLS Configuration                                                                             4 MPLS TE Configuration


4.25.5 (Optional) Adjusting BFD Parameters on an Ingress

Context
                 You can adjust BFD parameters on a tunnel ingress using either of the following
                 methods:

                 ●      Adjust BFD parameters globally.
                        Use this method when most TE tunnels on the ingress use the same BFD
                        parameters.
                 ●      Adjust BFD parameters on specific tunnel interfaces.
                        Use this method if some TE tunnels on the ingress need to use BFD
                        parameters that are different from global BFD parameters.
                         NOTE

