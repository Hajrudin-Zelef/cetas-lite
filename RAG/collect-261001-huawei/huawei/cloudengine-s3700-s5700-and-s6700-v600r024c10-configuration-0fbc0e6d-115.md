---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-115
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [16252, 16414]
sha256: 9d561f7689e10ad3ec8b3ed7c121c837569021546e58a838f60b36c5e9249a00
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                 S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                 Layer 2 mode to Layer 3 mode using the undo portswitch command.

                 Determine whether to perform this step based on the current interface mode.

         Step 9 Enable the RSVP-TE Hello extension mechanism on an interface.
                 mpls rsvp-te hello

                 After the RSVP-TE Hello extension function is enabled globally, you also need to
                 enable it on each interface that requires the function.

                 ----End

Verifying the Configuration
                 ●      Run the display mpls rsvp-te command to check the RSVP-TE configuration.
                 ●      Run the display default-parameter mpls rsvp-te command to check default
                        MPLS RSVP-TE parameters.
                 ●      Run the display mpls rsvp-te session ingress-lsr-id tunnel-id egress-lsr-id
                        command to check all information about a specified RSVP-TE session.
                 ●      Run the display mpls rsvp-te psb-content [ ingress-lsr-id tunnel-id [ lsp-id ]]
                        command to check RSVP-TE PSB information.
                 ●      Run the display mpls rsvp-te rsb-content [ ingress-lsr-id tunnel-id lsp-id ]
                        command to check RSVP-TE RSB information.
                 ●      Run the display mpls rsvp-te statistics { global | interface { interface-type
                        interface-number | interface-name } } command to check RSVP-TE statistics.
                 ●      Run the display mpls rsvp-te peer [ interface { interface-type interface-
                        number | interface-name } | peer-address ] command to check information
                        about RSVP-TE neighbors on RSVP-TE-enabled interfaces.

4.9.6 Configuring Reliable RSVP-TE Message Transmission

Prerequisites
                 Before configuring reliable RSVP-TE message transmission, complete the following
                 task:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    274
MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration


                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 On a network where BFD is not configured, you can configure reliable RSVP-TE
                 message transmission to increase the probability of detecting link faults and
                 reduce long-time traffic loss caused by intermittent link disconnections.
                 Perform the following configuration on each node of an MPLS TE tunnel.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Enable reliable RSVP-TE message transmission.
                 mpls rsvp-te reliable-delivery

                 ----End

4.9.7 Example for Configuring RSVP-TE Srefresh (Manual TE
FRR)
Networking Requirements
                 On the network shown in Figure 4-15, an RSVP-TE tunnel is established from LSR1
                 to LSR4 along the path LSR1 -> LSR2 -> LSR3 -> LSR4. Configure TE FRR to protect
                 the link LSR2 -> LSR3. Establish a bypass CR-LSP along the path LSR2 -> LSR5 ->
                 LSR3. LSR2 is a point of local repair (PLR), and LSR3 is a merge point (MP). Use
                 explicit paths to establish the primary and bypass MPLS TE tunnels.
                 Configure Srefresh on LSR2 and LSR3.

                         NOTE

                        To avoid loops in this scenario, ensure that all connected interfaces have STP disabled and
                        are removed from VLAN1. If STP is enabled and VLANIF interfaces of switches are used to
                        construct a Layer 3 ring network, an interface on the network will be blocked. As a result,
                        Layer 3 services on the network cannot run properly.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                       275
MPLS Configuration
MPLS Configuration                                                          4 MPLS TE Configuration


                 Figure 4-15 Network diagram of RSVP-TE Srefresh (manual TE FRR)




Configuration Roadmap
                 The configuration roadmap is as follows:

                 1.     Configure manual TE FRR.
                 2.     Configure RSVP-TE Srefresh on the PLR and MP of the bypass tunnel to
                        improve the reliability of RSVP-TE message transmission and resource
                        utilization.


Procedure
         Step 1 Configure manual TE FRR.

                 Configure the primary tunnel and bypass tunnel according to 4.26.5 Example for
                 Configuring Manual MPLS TE FRR and bind them.

         Step 2 Configure RSVP-TE Srefresh on LSR2 and LSR3.

                 # Configure RSVP-TE Srefresh on LSR2.
                 [LSR2] mpls
                 [LSR2-mpls] mpls rsvp-te srefresh
                 [LSR2-mpls] quit

                 # Configure RSVP-TE Srefresh on LSR3.
                 [LSR3] mpls
                 [LSR3-mpls] mpls rsvp-te srefresh
                 [LSR3-mpls] quit

                 ----End

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                      276
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


Verifying the Configuration
                 # Check the Srefresh status on LSR2.
                 [LSR2] display mpls rsvp-te statistics global
                  LSR ID: 2.2.2.9             LSP Count: 2
                  PSB Count: 2                 RSB Count: 2
                  RFSB Count: 1

                 Total Statistics Information:
                  PSB CleanupTimeOutCounter: 0       RSB CleanupTimeOutCounter: 0
                  SendPacketCounter: 122613        RecPacketCounter: 127446
                  SendCreatePathCounter: 25        RecCreatePathCounter: 260
                  SendRefreshPathCounter: 62209      RecRefreshPathCounter: 62113
                  SendCreateResvCounter: 21        RecCreateResvCounter: 31
                  SendRefreshResvCounter: 60101      RecRefreshResvCounter: 64792
                  SendResvConfCounter: 0          RecResvConfCounter: 0
                  SendHelloCounter: 0           RecHelloCounter: 0
                  SendAckCounter: 0             RecAckCounter: 0
                  SendPathErrCounter: 242         RecPathErrCounter: 0
                  SendResvErrCounter: 0          RecResvErrCounter: 0
                  SendPathTearCounter: 11         RecPathTearCounter: 8
                  SendResvTearCounter: 2          RecResvTearCounter: 0
                  SendSrefreshCounter: 1        RecSrefreshCounter: 1
                  SendAckMsgCounter: 1           RecAckMsgCounter: 1
                  SendChallengeMsgCounter: 0        RecChallengeMsgCounter: 0
                  SendResponseMsgCounter: 0         RecResponseMsgCounter: 0
                  SendErrMsgCounter: 0           RecErrMsgCounter: 0
                  SendRecoveryPathMsgCounter: 0       RecRecoveryPathMsgCounter: 0
                  SendGRPathMsgCounter: 0           RecGRPathMsgCounter: 0
                  ResourceReqFaultCounter: 0       RecGRPathMsgFromLSPMCounter: 0
                  Bfd neighbor count: 3         Bfd session count: 0

                 The values of the SendSrefreshCounter, RecSrefreshCounter,
                 SendAckMsgCounter, and RecAckMsgCounter fields are not 0, indicating that
                 Srefresh packets have been successfully transmitted.

