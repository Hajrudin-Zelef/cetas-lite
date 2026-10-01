---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-205
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [30093, 30235]
sha256: aec4df9769827be7f77c75570b8717deeae017ece79ff0ba4ca14e1c51507d56
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

5.8.4 (Optional) Configuring Flow Label-based Load Balancing
Context
                    Packets of multiple data flows on the same PW carry the same VC labels, which
                    are encapsulated on a PE. As such, when these packets arrive at a P, the P
                    forwards them all along the same path, though there are multiple load-balancing
                    paths available.
                    To load-balance different data flows, configure flow-label-based load balancing on
                    the PE. After the configurations are complete, the PE adds flow labels following
                    private network labels when encapsulating data packets, and the P load-balances
                    these data flows based on flow labels.
                    On the network shown in Figure 5-26, two different data flows need to be
                    transmitted to the remote end through an L2VPN.
                    1.   PE1 calculates flow labels based on the source and destination IP addresses of
                         the two data flows. The following assumes that the flow labels are 1025 and
                         1026.
                    2.   PE1 adds the flow labels following the private network labels in the packets
                         of the two data flows.
                    3.   When the two data flows reach P1, P1 performs a hash calculation based on
                         the flow labels and maps the two data flows onto different paths. The
                         following assumes that the next hop of the data flow with flow label 1025 is
                         P2, and the next hop of the data flow with flow label 1026 is P3.
                    4.   When the two data flows reach PE2, the private network labels and flow
                         labels are sequentially removed. PE2 then forwards the two data flows to
                         their destination CEs based on their private network labels.

                    Figure 5-26 Flow-label-based load balancing




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         481
VPN Configuration
VPN Configuration                                                                              5 VPWS Configuration


                    Flow-label-based load balancing applies to an L2VPN where there are multiple
                    links between Ps. Flow-label-based load balancing allows data flows on the same
                    VPN to be load-balanced along different paths based on flow labels, which
                    improves network resource utilization.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the AC interface view.
                    interface interface-type interface-number

         Step 3 Switch the interface working mode to Layer 3.
                    undo portswitch

                    Determine whether to perform this step based on the current interface working
                    mode.
         Step 4 Enable flow label-based load balancing for PWs on the interface.
                    mpls l2vpn flow-label { both | send | receive } [ secondary ] [ static ]

                    By default, flow label-based load balancing of PWs is disabled on an interface.

                          NOTE

                         ● Flow label-based load balancing can be configured successfully only when at least one
                           of the following conditions is true:
                             –    The receive parameter is configured on the local end, and the send parameter is
                                  configured on the remote end.
                             –    The send parameter is configured on the local end, and the receive parameter is
                                  configured on the remote end.
                             –    Both the send and receive parameters are configured on the local and remote
                                  ends.
                         ● If secondary is configured, the flow label-based load balancing function takes effect
                           only for the secondary PW on the interface. If secondary is not configured, the flow
                           label-based load balancing function takes effect only for the primary PW on the
                           interface.
                         ● If static is configured, the flow label-based load balancing function is configured
                           statically. The two PEs on both ends deliver the flow label-based load balancing
                           function, regardless of whether the other end has this function enabled. For dynamic
                           PWs, if static is not configured, the flow label-based load balancing function is
                           negotiated between the local end and remote end. For static PWs, the flow label-based
                           load balancing function is statically configured, regardless of whether static is
                           configured.
                             If the static flow label-based load balancing configuration does not match on both ends,
                             the device discards packets carrying a flow label, causing packet loss.

                    ----End




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                      482
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration


5.8.5 Verifying the Configuration

Procedure
                    ●    Run the display mpls static-l2vc [ vc-id | interface interface-type interface-
                         number | state { down | up } | brief ] command to check information about a
                         static PW on the devices at both ends of the PW.

                    ----End

5.8.6 Example for Configuring SVC VPWS

Networking Requirements
                    SVC VPWS needs to be configured on CE1 and CE2. An SVC connection needs to
                    be created on PEs and VC labels need to be specified.


                    Figure 5-27 Network diagram of configuring SVC VPWS
                          NOTE

                         In this example, interface1 and interface2 represent VLANIF 10 and VLANIF 20, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure an IGP on the MPLS backbone network to implement IP
                         connectivity.
                    2.   Enable MPLS and MPLS L2VPN.
                    3.   Create static L2VCs between PEs and configure VC labels.
                               NOTE

                              By default, the Link-type Negotiation Protocol (LNP) is enabled globally on the device.
                              If a VLANIF interface is used as an AC interface for L2VPN, the configuration conflicts
                              with LNP. In this case, run the lnp disable command in the system view to disable
                              LNP.



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      483
VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration


Procedure
         Step 1 Configure the VLANs that interfaces belong to and assign IP addresses to VLANIF
                interfaces.

