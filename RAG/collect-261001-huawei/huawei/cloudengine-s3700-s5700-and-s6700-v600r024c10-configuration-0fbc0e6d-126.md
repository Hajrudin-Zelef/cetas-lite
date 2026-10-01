---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-126
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-04-08", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [17816, 17946]
sha256: 24d2516a124fc1f19cc0054c2994d45ad42c97c0130caa860a07f43ef4bd0a9c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        To avoid loops in this scenario, ensure that all connected interfaces have STP disabled and
                        are removed from VLAN1. If STP is enabled and VLANIF interfaces of switches are used to
                        construct a Layer 3 ring network, an interface on the network will be blocked. As a result,
                        Layer 3 services on the network cannot run properly.


                 Figure 4-18 Network diagram for configuring RSVP-TE authentication (manual TE
                 FRR)




Configuration Roadmap
                 The configuration roadmap is as follows:

                 1.     Configure manual TE FRR.
                 2.     Configure RSVP-TE authentication on LSR2 and LSR3 to prevent forged RSVP-
                        TE resource reservation requests from occupying network resources.


Procedure
         Step 1 Configure manual TE FRR.

                 Configure the primary tunnel and bypass tunnel according to 4.26.5 Example for
                 Configuring Manual MPLS TE FRR and bind them.

         Step 2 Configure RSVP-TE authentication on LSR2 and LSR3.

                 To check whether the authentication function is successfully configured, configure
                 the RSVP-TE handshake mechanism and a local password.



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                       300
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


                        NOTE

                 Because the LSR ID of the peer device is used as the neighbor address, CSPF must be enabled on
                 the device that needs to be configured with RSVP-TE authentication.

                 # Configure RSVP-TE authentication on LSR2.
                 [LSR2] mpls
                 [LSR2-mpls] mpls rsvp-te peer 3.3.3.9
                 [LSR2-mpls-rsvp-te-peer-3.3.3.9] mpls rsvp-te authentication cipher YsHsjx_202206
                 [LSR2-mpls-rsvp-te-peer-3.3.3.9] mpls rsvp-te authentication handshake
                 [LSR2-mpls-rsvp-te-peer-3.3.3.9] quit

                 # Configure RSVP-TE authentication on LSR3.
                 [LSR3] mpls
                 [LSR3-mpls] mpls te cspf
                 [LSR3-mpls] mpls rsvp-te peer 2.2.2.9
                 [LSR3-mpls-rsvp-te-peer-2.2.2.9] mpls rsvp-te authentication cipher YsHsjx_202206
                 [LSR3-mpls-rsvp-te-peer-2.2.2.9] mpls rsvp-te authentication handshake
                 [LSR3-mpls-rsvp-te-peer-2.2.2.9] quit

                 ----End

Verifying the Configuration
                 # Check the status of the authentication function on LSR2.
                 [LSR2] display mpls rsvp-te statistics global
                  LSR ID: 2.2.2.9             LSP Count: 2
                  PSB Count: 2                 RSB Count: 2
                  RFSB Count: 1

                 Total Statistics Information:
                  PSB CleanupTimeOutCounter: 0       RSB CleanupTimeOutCounter: 1
                  SendPacketCounter: 81         RecPacketCounter: 82
                  SendCreatePathCounter: 12       RecCreatePathCounter: 16
                  SendRefreshPathCounter: 41       RecRefreshPathCounter: 12
                  SendCreateResvCounter: 3        RecCreateResvCounter: 6
                  SendRefreshResvCounter: 11       RecRefreshResvCounter: 26
                  SendResvConfCounter: 0         RecResvConfCounter: 0
                  SendHelloCounter: 0          RecHelloCounter: 0
                  SendAckCounter: 0            RecAckCounter: 0
                  SendPathErrCounter: 0         RecPathErrCounter: 0
                  SendResvErrCounter: 0         RecResvErrCounter: 0
                  SendPathTearCounter: 7         RecPathTearCounter: 5
                  SendResvTearCounter: 1         RecResvTearCounter: 1
                  SendSrefreshCounter: 3        RecSrefreshCounter: 6
                  SendAckMsgCounter: 3           RecAckMsgCounter: 3
                  SendChallengeMsgCounter: 1       RecChallengeMsgCounter: 1
                  SendResponseMsgCounter: 1        RecResponseMsgCounter: 1
                  SendErrMsgCounter: 0          RecErrMsgCounter: 0
                  SendRecoveryPathMsgCounter: 0      RecRecoveryPathMsgCounter: 0
                  SendGRPathMsgCounter: 0          RecGRPathMsgCounter: 0
                  ResourceReqFaultCounter: 0      RecGRPathMsgFromLSPMCounter: 0
                  Bfd neighbor count: 3        Bfd session count: 0

                 The values of the SendChallengeMsgCounter, RecChallengeMsgCounter,
                 SendResponseMsgCounter, and RecResponseMsgCounter fields are not 0,
                 indicating that the handshake between the PLR and MP is successful. This means
                 that the authentication function is configured successfully.
                 # Shut down the protected outbound interface on the PLR (LSR2).
                 [LSR2] interface vlanif 200
                 [LSR2-Vlanif200] shutdown
                 [LSR2-Vlanif200] quit


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                    301
MPLS Configuration
MPLS Configuration                                                                          4 MPLS TE Configuration


                 # Run the display interface tunnel 1 command on LSR1. The command output
                 shows that Tunnel1 is still up.
                 [LSR1] display interface tunnel 1
                 Tunnel1 current state : UP (ifindex: 28)
                 Line protocol current state : UP
                 Last line protocol up time : 2024-04-08 06:20:17
                 Description:
                 ...

                 # Run the tracert lsp te tunnel 1 command on LSR1 to check the path of the
                 tunnel.
                 [LSR1] tracert lsp te tunnel 1
                  LSP Trace Route FEC: TE TUNNEL IPV4 SESSION QUERY Tunnel1 , press CTRL_C to break.
                  TTL Replier           Time Type       Downstream
                  0                       Ingress 10.1.1.2/[17 ]
                  1   10.1.1.2        2 ms Transit 10.1.4.2/[16 ]
                  2   10.1.4.2        2 ms Transit 10.1.5.2/[3 ]
                  3   10.1.5.2        1 ms Transit 10.1.3.2/[3 ]
                  4   4.4.4.9         11     Egress

                 The preceding information indicates that the link has been switched to the bypass
                 tunnel (bypass CR-LSP).

