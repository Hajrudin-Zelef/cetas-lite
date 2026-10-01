---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-303
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [44877, 45026]
sha256: bb98ca4a223d8eb17e88dd5d64b6fd7d3908dae943aceb3085219a5dce2f75d5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              720
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration


                    Figure 6-34 Typical HVPLS networking for VPWS accessing VPLS




6.13.2 Configuring Static VPWS Accessing VPLS

Prerequisites
                    Before configuring static VPWS accessing VPLS, you have completed the following
                    tasks:

                    ●   Configure IP addresses and an IGP on UPEs, SPEs, and Ps.
                    ●   Configure LSR IDs and enable MPLS and MPLS LDP on UPEs, SPEs, and Ps.
                    ●   Enable MPLS L2VPN on UPEs and SPEs.
                    ●   Configure static LSPs on UPEs and SPEs.
                         NOTE

                        You need to configure a remote LDP session if SPEs or a UPE and an SPE are not directly
                        connected.
                        When VPN services need to be transmitted over a specific TE tunnel or when load balancing
                        needs to be performed among multiple tunnels to fully use network resources, you also
                        need to configure tunnel policies.


Context
                    On an HVPLS network, if UPEs do not support dynamic VPWS, configure static
                    VPWS for accessing SPEs.

Procedure
                    ●   Configure static VPWS on UPEs for accessing SPEs. For configuration details,
                        see 5.8.3 Configuring an SVC VPWS Connection.
                    ●   Bind a VSI to static VPWS on each SPE.

                        To implement communication between SPEs and between SPEs and UPEs, on
                        SPEs, configure VSI peer relationships between SPEs and between SPEs and
                        UPEs.

                        a.   Enter the system view.
                             system-view


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    721
VPN Configuration
VPN Configuration                                                                                   6 VPLS Configuration


                        b.    Create a VSI and specify the static member discovery mode for the VSI.
                              vsi vsi-name static

                        c.    Configure LDP as the signaling protocol of the VSI and enter the VSI-LDP
                              view.
                              pwsignal ldp

                        d.    Configure a VSI ID.
                              vsi-id vsi-id

                              ▪     The VSI IDs of the two ends of a PW must be the same; otherwise,
                                    the VSI cannot be created. If the VSI IDs of the two ends are
                                    different, specify the negotiation-vc-id vc-id parameter in the peer
                                    command to set the VSI ID used for PW negotiation.

                              ▪     VSIs exist only on PEs. A PE can be configured with multiple VSIs, but
                                    the ID of each VSI must be unique on a PE.
                        e.    Configure VSI peer relationships between SPEs.
                              peer peer-address [ negotiation-vc-id vc-id ] [ tnl-policy policy-name ]

                        f.    Configure VSI peer relationships between SPEs and UPEs.
                              peer peer-address [ negotiation-vc-id vc-id ] [ tnl-policy policy-name ] static-upe trans
                              transmit-label recv receive-label

                    ----End

6.13.3 Configuring Dynamic VPWS Accessing VPLS

Prerequisites
                    Before configuring dynamic VPWS accessing VPLS, you have completed the
                    following tasks:

                    ●   Configure IP addresses and an IGP on UPEs, SPEs, and Ps.
                    ●   Configure LSR IDs and enable MPLS and MPLS LDP on UPEs, SPEs, and Ps.
                    ●   Enable MPLS L2VPN on UPEs and SPEs.
                    ●   Establish tunnels between SPEs and between UPEs and SPEs to transmit
                        service traffic.
                               NOTE

                              You need to configure a remote LDP session if SPEs or a UPE and an SPE are not
                              directly connected.


Procedure
                    ●   Configure dynamic VPWS on UPEs for accessing SPEs. For configuration
                        details, see 5.7.3 Configuring an LDP VPWS Connection.
                    ●   Bind a VSI to dynamic VPWS on SPEs. For configuration details, see 6.9
                        Configuring LDP HVPLS.

                    ----End



Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                              722
VPN Configuration
VPN Configuration                                                                           6 VPLS Configuration


6.13.4 Verifying the Configuration
Procedure
                    ●   Run the display vsi [ name vsi-name ] [ verbose ] command to check VPLS
                        VSI information.
                    ●   Run the display mpls static-l2vc interface { interface-name | interface-type
                        interface-number } command to check information about the static VC
                        configured on the device.
                    ●   Run the display mpls l2vc [ vc-id | interface interface-type interface-
                        number ] command to check information about LDP VPWS connections.
                    ●   Run the display vsi mac-withdraw loop-detect command to check
                        information about MAC Withdraw loop detection.
                    ----End

6.13.5 Example for Configuring Static VPWS Accessing VPLS
Networking Requirements
                    Figure 6-35 shows a backbone network built by an enterprise. UPEs do not
                    support dynamic VPWS and need to access SPEs through static VPWS. Site1
                    connects to the backbone network by connecting CE1 to UPE1, while Site2
                    connects to the backbone network by connecting CE2 to UPE2. Users at Site1 and
                    Site2 need to communicate at Layer 2, and user information needs to be retained
                    in Layer 2 packets when the packets are transmitted over the backbone network.

                    Figure 6-35 Network diagram of configuring static VPWS accessing VPLS
                         NOTE

                        In this example, interface1 and interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     723
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration




