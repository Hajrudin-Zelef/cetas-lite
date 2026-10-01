---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-120
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [17017, 17157]
sha256: 5c43c1d5f0c48af10c71ea2582e865dad6a37cee1b1ca85deb2ceb5060b0af5d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      Neighbor interface address authentication
                        Neighbor address authentication is performed based on the addresses of
                        interfaces between two directly or indirectly connected LSRs. It applies to
                        MPLS TE FRR inter-domain scenarios. For example, the interfaces between the
                        PLR and MP can adopt this authentication mode.
                        –   In inter-domain scenarios where MPLS TE FRR is configured, neighbor
                            interface address authentication is recommended.
                        –   In inter-domain scenarios where MPLS TE FRR is not configured, neighbor
                            interface address authentication or interface authentication is
                            recommended.
                 ●      Neighboring node authentication (LSR authentication)
                        Neighboring node authentication is performed between two directly or
                        indirectly connected LSRs. In this authentication mode, global RSVP-TE key
                        authentication takes effect. For example, neighboring node authentication is
                        usually configured between the PLR and MP based on LSR IDs.
                        Neighboring node authentication is recommended for all non-inter-domain
                        scenarios.
                 ●      Interface authentication
                        Interface authentication is performed between the directly connected
                        interfaces of two LSRs, for example, the interfaces between LSR1 and LSR2.
                        Neighbor interface address authentication or interface authentication is
                        recommended for all inter-domain scenarios where MPLS TE FRR is not
                        configured.
                 The keys configured for the interfaces of two neighbors must be the same;
                 otherwise, authentication fails and the received RSVP-TE packets are discarded.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           286
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 If interface authentication, neighboring node authentication, and neighbor
                 interface address authentication are all configured for the same neighbor,
                 neighbor interface address authentication preferentially takes effect, followed by
                 neighboring node authentication and interface authentication, in descending
                 order.
                 You can select different RSVP authentication modes as required. Table 4-12 lists
                 corresponding selection rules.

                 Table 4-12 Rules for selecting RSVP-TE key authentication modes
                  RSVP-TE        Neighbor                                Interface Authentication
                  Key            Interface         Neighboring
                  Authentica     Address           Node
                  tion           Authentication    Authentication

                  Authenticati   ● RSVP-TE         ● RSVP-TE             ● RSVP-TE neighbor
                  on mode          neighbor-         neighbor-             interface-based
                                   based             based                 authentication
                                   authenticati      authenticatio       ● A key is configured for
                                   on                n                     the specified interface.
                                 ● A key is        ● A key is
                                   configured        configured
                                   based on an       based on the
                                   interface         LSR ID of the
                                   address of        neighboring
                                   the               node.
                                   neighboring
                                   node.

                  Application    ● Incoming        Packets received      Packets received from or
                  scope            packets,        from and sent to      sent to an interface
                                   whose           a neighboring
                                   source or       node
                                   next-hop
                                   address is
                                   the same as
                                   the
                                   configured
                                   one
                                 ● Outgoing
                                   packets,
                                   whose
                                   destination
                                   or next-hop
                                   address is
                                   the same as
                                   the
                                   configured
                                   one

                  Priority       High              Medium                Low




Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                                287
MPLS Configuration
MPLS Configuration                                                                  4 MPLS TE Configuration


                  RSVP-TE           Neighbor                                 Interface Authentication
                  Key               Interface          Neighboring
                  Authentica        Address            Node
                  tion              Authentication     Authentication

                  Application       All scenarios      Non-inter-            Scenarios except those
                  scenarios                            domain                where MPLS TE FRR is
                                                       scenarios             configured and the primary
                                                                             CR-LSP is in the FRR-in-use
                                                                             state

                  Advantage         None               Simple                None
                                                       configuration




4.10.2 Configuring RSVP-TE Authentication
Prerequisites
                 Before configuring RSVP-TE authentication, you have completed the following
                 task:
                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 To improve network security, configure RSVP-TE authentication between RSVP-TE
                 neighbors. The keys configured for both ends must be the same. Otherwise,
                 authentication fails, and packets received by the RSVP neighbor or interface are
                 discarded.
                 You can configure RSVP-TE key authentication in the interface view or MPLS RSVP-
                 TE neighbor view.
                 ●      RSVP-TE key authentication configured in the interface view applies to two
                        directly connected nodes.
                 ●      RSVP-TE key authentication configured in the RSVP-TE neighbor view applies
                        to any two nodes that are configured as neighbors of each other. This
                        configuration is recommended.
                 Perform the following configuration on each node of an MPLS TE tunnel.


                        NOTICE

                 The configuration must be completed on the two directly connected interfaces
                 within three refresh periods. Otherwise, the involved session goes down.


Procedure
         Step 1 Enter the system view.
                 system-view


