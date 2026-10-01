---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-287
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [42267, 42408]
sha256: 5b1fde34aa912300358914bc003bfc0e0c4a9e6538471441f9669026b17cf098
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Application Scenarios of Inter-AS VPLS Options
                    Option A is easy to configure, does not require MPLS to run between ASBRs, and
                    does not require any special configuration for inter-AS communication. This mode,
                    however, has poor scalability and high requirements on PEs. It applies to the early
                    service deployment phase during which there is a small number of inter-AS VPNs.

6.10.2 Configuring Inter-AS LDP VPLS Option A
Prerequisites
                    Before configuring inter-AS LDP VPLS Option A, you have completed the following
                    tasks:

                    ●   Configure an IGP on the MPLS backbone network of each AS to enable IP
                        connectivity of the backbone network within each AS.
                    ●   Configure basic MPLS functions on the MPLS backbone network of each AS.
                    ●   Configure MPLS LDP and establish LDP LSPs on the MPLS backbone network
                        of each AS.
                    ●   Establish a tunnel between the PE and ASBR in the same AS (in Option A
                        mode).

Context
                    The following describes the configuration of inter-AS LDP VPLS Option A:

                    ●   LDP VPLS is configured in each AS.
                    ●   Each ASBR regards its peer ASBR as a local CE.
                    ●   No inter-AS configuration needs to be performed on ASBRs.
                    ●   No IP address needs to be configured for directly connected interfaces
                        between ASBRs.

                    Configure LDP VPLS on the PEs and ASBRs in each AS. For configuration details,
                    see Configuring LDP VPLS.

                         NOTE

                        In Option A, an ASBR must reserve an interface as an AC interface for each inter-AS VC.
                        Option A can be used when the number of inter-AS VCs is small. Compared with L3VPN,
                        inter-AS L2VPN Option A consumes more resources and involves more configurations.
                        Therefore, inter-AS L2VPN Option A is not recommended.


6.10.3 Configuring Inter-AS BGP VPLS Option A
Prerequisites
                    Before configuring inter-AS BGP VPLS Option A, you have completed the following
                    tasks:

                    ●   Configure an IGP on the MPLS backbone network of each AS to enable IP
                        connectivity of the backbone network within each AS.
                    ●   Configure basic MPLS functions on the MPLS backbone network of each AS.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    679
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration


                    ●   Establish a tunnel between the PE and ASBR in the same AS (in Option A
                        mode).


Context
                    The following describes the configuration of inter-AS BGP VPLS Option A:

                    ●   BGP VPLS is configured in each AS.
                    ●   Each ASBR regards its peer ASBR as a local CE.
                    ●   PEs and ASBRs have VSIs configured, and their AC interfaces are bound to
                        these VSIs. The VSI for PEs is used for connecting to CEs, and the VSI for
                        ASBRs is used for connecting to the peer ASBR.

                    Configure BGP VPLS on the PEs and ASBRs in each AS. For configuration details,
                    see 6.7 Configuring BGP VPLS.

                         NOTE

                        In Option A, an ASBR must reserve an interface as an AC interface for each inter-AS VC.
                        Option A can be used when the number of inter-AS VCs is small. Compared with L3VPN,
                        inter-AS L2VPN Option A consumes more resources and involves more configurations.
                        Therefore, inter-AS L2VPN Option A is not recommended.


6.10.4 Verifying the Configuration

Procedure
                    ●   Run the display vsi [ name vsi-name ] [ verbose ] command to check VSI
                        information.
                    ●   Run the display vsi remote ldp [ [ router-id ip-address ] [ pw-id pw-id ] |
                        [ verbose ] | [ unmatch ] ] command to check remote VSI information.
                    ●   Run the display vsi remote bgp [ nexthop nexthop-address { export-vpn-
                        target vpn-target ] | route-distinguisher | route-distinguisher } command
                        to check remote VSI information.
                    ●   Run the display vpls connection [ bgp | ldp | vsi vsi-name ] [ down | up ]
                        [ verbose ] command to check VPLS connections.

                    ----End

6.10.5 Example for Configuring Inter-AS LDP VPLS Option A

Networking Requirements
                    On the enterprise network shown in Figure 6-24, Site1 connects to the VPLS
                    domain of AS100 by connecting CE1 to PE1, and Site2 connects to the VPLS
                    domain of AS200 by connecting CE2 to PE2. The network environments of the
                    branch sites are stable. AS100 and AS200 communicate with each other through
                    ASBR_PE1 and ASBR_PE2. IS-IS is used as the IGP on the MPLS backbone network
                    in an AS. Users at Site1 and Site2 need to communicate at Layer 2, and user
                    information needs to be retained in Layer 2 packets when the packets are
                    transmitted over the backbone network.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    680
VPN Configuration
VPN Configuration                                                                            6 VPLS Configuration


                    Figure 6-24 Network diagram of configuring inter-AS LDP VPLS Option A
                          NOTE

                         In this example, interface1 and interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure VPLS to transparently transmit Layer 2 packets on the backbone
                         network to implement Layer 2 communication between Site1 and Site2 and
                         to retain user information in Layer 2 packets when the packets are
                         transmitted over the backbone network.
                    2.   Use LDP VPLS to implement Layer 2 communication between CEs because the
                         network environments of the branch sites are stable.
                    3.   Configure an IGP on the backbone network to implement communication
                         between devices within an AS on the public network.
                    4.   Configure basic MPLS functions and LDP on PEs on the backbone network to
                         implement VPLS.
                    5.   Establish tunnels for transmitting data between PEs within an AS to prevent
                         data from being accessed by the public network. Establish a dynamic LSP
                         between the PE and ASBR_PE in the same AS. If the PE and ASBR_PE are not
                         directly connected, establish a remote LDP session.
                    6.   Enable MPLS L2VPN on PEs to implement VPLS.
                    7.   Create VSIs on PEs in the same AS, configure LDP as the signaling protocol,
                         and bind VSIs to AC interfaces to implement LDP VPLS.
                    8.   Configure the peer ASBR as a CE on ASBR_PEs, and bind VSIs to peer
                         interfaces to implement inter-AS VPLS OptionA.

