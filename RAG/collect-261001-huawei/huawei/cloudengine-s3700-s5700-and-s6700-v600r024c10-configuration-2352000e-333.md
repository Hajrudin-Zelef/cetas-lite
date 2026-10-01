---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-333
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [49788, 49943]
sha256: 2486033198fe30d994439549e40d9025d95c19a288b7f28227e7ecd142dffc2a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 3 Check whether the VSI IDs or negotiation VC IDs on both ends are the same.
                    <HUAWEI> display vsi name tt verbose
                     ***VSI Name                : tt
                        Administrator VSI          : no
                        Isolate Spoken          : disable
                        VSI Index            :3
                        PW Signaling             : ldp
                        Member Discovery Style : static
                        PW MAC Learn Style             : unqualify
                        Encapsulation Type           : vlan
                        MTU                 : 1500
                        Diffserv Mode            : uniform
                        Service Class         : --
                        Color             : --
                        DomainId               : 255
                        Domain Name                  :
                        Tunnel Policy Name            : p1
                        Ignore AcState           : disable
                        P2P VSI             : disable
                        Multicast Fast Switch : disable
                        Create Time            : 2 days, 2 hours, 47 minutes, 40 seconds
                        VSI State           : up
                        Resource Status           : --

                         VSI ID           : 101
                        *Peer Router ID       : 1.1.1.15
                         primary or secondary : primary
                    ......

                    ●     If the VSI IDs or negotiation VC IDs on both ends are different, run the vsi-id
                          or peer peer-address negotiation-vc-id vc-id command in the VSI-LDP view

Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                  800
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration


                          to change the VSI ID or negotiation ID on one end to be the same as that on
                          the other end.
                    ●     If the VSI IDs or negotiation VC IDs on both ends are the same, go to step 4.
                           NOTE

                         A VSI can be up only when VSI IDs or negotiation VC IDs on both ends are the same.

         Step 4 Check whether the LDP session is up on both ends.
                    Run the display vsi name vsi-name verbose command to check whether the
                    Session field value is up.
                    <HUAWEI> display vsi name tt verbose
                     ***VSI Name                : tt
                        Administrator VSI          : no
                        Isolate Spoken          : disable
                        VSI Index            :3
                        PW Signaling             : ldp
                        Member Discovery Style : static
                        PW MAC Learn Style             : unqualify
                        Encapsulation Type           : vlan
                        MTU                 : 1500
                        Diffserv Mode            : uniform
                        Service Class         : --
                        Color             : --
                        DomainId               : 255
                        Domain Name                  :
                        Tunnel Policy Name            : p1
                        Ignore AcState           : disable
                        P2P VSI             : disable
                        Multicast Fast Swicth : disable
                        Create Time            : 2 days, 2 hours, 47 minutes, 40 seconds
                        VSI State           : up
                        Resource Status           : --

                         VSI ID           : 101
                        *Peer Router ID        : 1.1.1.15
                         primary or secondary : primary
                         ignore-standby-state : no
                         VC Label           : 187393
                         Peer Type           : dynamic
                         Session           : up
                         Tunnel ID           : 0x0000000001004c4d62
                    ......
                    ●     If the LDP session is down, see Fault Symptom to bring the LDP session up.
                    ●     If the LDP session is up, go to step 5.
                           NOTE

                         The two ends can perform L2VPN negotiation only after the LDP session is up.

         Step 5 Check whether the VSI has selected a tunnel.
                    Run the display vsi name vsi-name verbose command.
                    ●     Check whether the Tunnel ID field is empty. If the Tunnel ID field is empty,
                          the VSI does not select any tunnel.
                    ●     Check the Tunnel Policy Name field. If this field is not displayed, the VSI uses
                          an LDP LSP or has no tunnel policy configured. If this field is displayed, the
                          VSI selects an MPLS TE tunnel and a tunnel policy is configured. If the Tunnel
                          Policy Name field value indicates that a tunnel policy is applied to the VSI,
                          run the display this command in the tunnel policy view to check the
                          configuration of the tunnel policy.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                801
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration

                        [HUAWEI-tunnel-policy-p1] display this
                        #
                        tunnel-policy p1
                         tunnel select-seq cr-lsp load-balance-number 1
                        #

                         NOTE

                        If the tunnel binding destination dest-ip-address te { tunnel interface-number }
                        command is configured in the tunnel policy view, you also need to run the mpls te
                        reserved-for-binding command in the tunnel interface view.

                    If the tunnel status is down at both ends, see LDP LSP Down or An MPLS TE
                    Tunnel's State Is Down to bring the tunnel up. If the tunnel is up and the TE
                    interface is correctly configured, go to step 6.

                         NOTE

                        A VSI can be up only when the corresponding tunnel between the two ends is up.

         Step 6 Check whether the local and remote AC interfaces are both up.
                    Run the display vsi name vsi-name verbose command to check whether the
                    value of the State field corresponding to the Interface Name field is up.
                    ●   If the AC interfaces on both ends are down, see Interface Basic Configuration
                        to bring the interfaces up.
                    ●   If the AC interfaces on both ends are up, go to step 7.
                         NOTE

                        A VSI can be up only when AC interfaces on both ends are up.

                    ----End




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                 802
VPN Configuration
VPN Configuration                                                   7 Tunnel Management Configuration




            7           Tunnel Management Configuration


                    7.1 Overview of Tunnel Management
                    7.2 Understanding Tunnel Management
                    7.3 Configuration Precautions for Tunnel Management
                    7.4 Configuring a Tunnel Policy
                    7.5 Configuring a Tunnel Selector
                    7.6 Maintaining Tunnel Management


