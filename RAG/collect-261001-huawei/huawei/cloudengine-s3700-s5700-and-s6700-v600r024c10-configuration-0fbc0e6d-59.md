---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-59
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [8193, 8343]
sha256: 28f782690c7f18535abf08f58d4dc3746cbf60dfe838eb57ffadf8944d4d0a65
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                                    After this command is run, the egress does not create a BFD session immediately.
                                    Instead, the egress waits for an LSP ping request carrying the BFD TLV before
                                    creating a BFD session.

                 ----End

3.16.4 Configuring a Policy for Triggering Dynamic BFD for
LDP LSP
                 Configure a policy for dynamically establishing a BFD session to monitor an LDP
                 LSP.

Context
                 The establishment of BFD sessions for LDP LSPs can be triggered in either of
                 following modes:
                 ●      Host mode: applicable when all host addresses can be used to establish BFD
                        sessions. In this mode, you can specify next hops and outbound interfaces to
                        specify the LSPs for which BFD sessions can be established.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      137
MPLS Configuration
MPLS Configuration                                                                         3 MPLS LDP Configuration


                 ●      FEC list mode: applicable when only some host addresses can be used to
                        establish BFD sessions.
                 Perform the following steps on the ingress of an LSP to be monitored.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 (Optional) To use the FEC list mode, perform the following operations:
                 1.     Create a FEC list and enter the FEC list view.
                        fec-list list-name

                 2.     Add a node to the FEC list.
                        fec-node fec-node-address [ nexthop nexthop-address | outgoing-interface interface-type interface-
                        number ]
                 3.     Return to the system view.
                        quit

         Step 3 Enter the MPLS view.
                 mpls

         Step 4 Configure a policy for triggering dynamic BFD for LDP LSP.
                 mpls bfd-trigger { host [ nexthop next-hop-address | outgoing-interface interface-type interface-
                 number ] * | fec-list list-name } [ option-tlv ]

                 After this command is run, the device starts to create a BFD session.

                 ----End

3.16.5 (Optional) Adjusting BFD Parameters
                 If default BFD settings do not meet actual requirements, you can adjust them,
                 including the minimum interval at which BFD packets are sent, the minimum
                 interval at which BFD packets are received, the detection period, and the BFD
                 detection multiplier.

Context
                 Perform the following steps on the ingress.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the BFD view.
                 bfd

         Step 3 Adjust the interval at which LSP ping packets are sent.
                 mpls ping interval interval

         Step 4 Exit the BFD view.
                 quit

         Step 5 Enter the MPLS view.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                            138
MPLS Configuration
MPLS Configuration                                                                          3 MPLS LDP Configuration

                 mpls

         Step 6 (Optional) Set BFD detection time parameters.
                 mpls bfd { min-tx-interval min-tx-interval-value | min-rx-interval min-rx-interval-value | detect-
                 multiplier detect-multiplier-value }*

                 Effective local interval at which BFD packets are sent = max{Locally configured
                 interval at which BFD packets are sent, Remotely configured interval at which BFD
                 packets are received}
                 Effective local interval at which BFD packets are received = max{Remotely
                 configured interval at which BFD packets are sent, Locally configured interval at
                 which BFD packets are received}
                 Local detection period = Local interval at which BFD packets are received x
                 Remote BFD detection multiplier
                 By default, the minimum interval at which BFD packets are sent and the minimum
                 interval at which BFD packets are received are both 10 ms, and the detection
                 multiplier is 3 on the ingress. The minimum interval at which BFD packets are sent
                 and the minimum interval at which BFD packets are received are 10 ms, and the
                 detection multiplier is 3 on the egress.
                 Therefore, you can adjust the minimum interval at which BFD packets are sent,
                 the minimum interval at which BFD packets are received, and the detection
                 multiplier only on the ingress to update BFD detection time parameters on both
                 the ingress and egress.

                 ----End

3.16.6 Verifying the Configuration
                 After configuring dynamic BFD for LDP LSP, you can check BFD configurations and
                 session information on the ingress and egress of a specified LDP LSP.

Procedure
                 ●      Run the display mpls bfd session [ fec ip-address | nexthop ip-address |
                        outgoing-interface interface-type interface-number | protocol { rsvp-te |
                        ldp }] [ verbose ] command to check BFD session information.
                 ●      Run the display bfd session all verbose command to check detailed BFD
                        session information on the ingress.
                 ●      Run the display bfd session passive-dynamic verbose command to check
                        detailed BFD session information on the egress.
                 ----End

3.16.7 Example for Configuring Dynamic BFD for LDP LSP
                 This section provides an example for configuring dynamic BFD for LDP LSP. The
                 configuration involves enabling MPLS and MPLS LDP globally and for specific
                 interfaces and enabling BFD on two ends of a link to be monitored.

Networking Requirements
                 The proliferation of MPLS LDP applications drives the increasing demand for
                 network reliability. To meet the reliability requirement, BFD for LDP can be used to

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                            139
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration


                 rapidly detect faults and trigger a primary/backup LSP switchover. BFD for LDP is
                 often used together with LDP FRR.
                 On the network shown in Figure 3-27, PE1, P1, P2, and PE2 are in the same MPLS
                 domain. Establish primary and backup LDP LSPs from PE1 to PE2. Configure
                 dynamic BFD to check LDP LSP connectivity.

                 Figure 3-27 Networking diagram of dynamic BFD for LDP LSP
                         NOTE

                        In this example, interfaces 1 and 2 represent VLANIF100 and VLANIF200, respectively.




