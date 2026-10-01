---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-305
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [45177, 45332]
sha256: b1970c0da89eb3c70fb961ce543acd3671c06bddbf1a711535ef2e9b828bbeb2
---

                    # Check LSP information on SPE1.
                    [SPE1] display mpls lsp
                     Flag after Out IF: (I) - LSP Is Only Iterated by RLFA
                    -------------------------------------------------------------------------------
                                 LSP Information: LDP LSP
                    -------------------------------------------------------------------------------
                    FEC              In/Out Label In/Out IF                      Vrf Name
                    1.1.1.9/32         3/NULL         -/-
                    2.2.2.9/32         NULL/3         -/Vlanif30
                    2.2.2.9/32         4096/3        -/Vlanif30
                    3.3.3.9/32         NULL/4097        -/Vlanif30
                    3.3.3.9/32         4097/4097       -/Vlanif30

         Step 4 Establish a remote LDP session.

Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       726
VPN Configuration
VPN Configuration                                                                                    6 VPLS Configuration


                    # Configure SPE1.
                    [SPE1] mpls ldp remote-peer 3.3.3.9
                    [SPE1-mpls-ldp-remote-3.3.3.9] remote-ip 3.3.3.9
                    [SPE1-mpls-ldp-remote-3.3.3.9] quit

                    # Configure SPE2.
                    [SPE2] mpls ldp remote-peer 1.1.1.9
                    [SPE2-mpls-ldp-remote-1.1.1.9] remote-ip 1.1.1.9
                    [SPE2-mpls-ldp-remote-1.1.1.9] quit

                    # Check LDP session information on SPE1.
                    [SPE1] display mpls ldp session

                    LDP Session(s) in Public Network
                    Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                    A '*' before a session means the session is being deleted.
                    ------------------------------------------------------------------------------
                    PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                    ------------------------------------------------------------------------------
                    2.2.2.9:0        Operational DU Passive 0000:00:02 10/10
                    3.3.3.9:0        Operational DU Passive 0000:00:01 5/5
                    ------------------------------------------------------------------------------
                    TOTAL: 2 session(s) Found.

                    After the configurations are complete, run the display mpls ldp session command
                    on SPE1 and SPE2. The status of the peer relationship between SPE1 and SPE2 is
                    Operational, indicating that the peer relationship has been established.
         Step 5 Establish static LSPs between UPEs and SPEs.
                    # Configure UPE1.
                    [UPE1] mpls lsr-id 4.4.4.9
                    [UPE1] mpls
                    [UPE1-mpls] quit
                    [UPE1] interface vlanif 20
                    [UPE1-Vlanif20] mpls
                    [UPE1-Vlanif20] quit
                    [UPE1] static-lsp ingress UPE1toSPE1 destination 1.1.1.9 32 nexthop 3.1.1.1 out-label 20
                    [UPE1] static-lsp egress SPE1toUPE1 incoming-interface vlanif 20 in-label 30

                    # Configure UPE2.
                    [UPE2] mpls lsr-id 5.5.5.9
                    [UPE2] mpls
                    [UPE2-mpls] quit
                    [UPE2] interface vlanif 50
                    [UPE2-Vlanif50] mpls
                    [UPE2-Vlanif50] quit
                    [UPE2] static-lsp ingress UPE2toSPE2 destination 3.3.3.9 32 nexthop 4.1.1.1 out-label 40
                    [UPE2] static-lsp egress SPE2toUPE2 incoming-interface vlanif 50 in-label 50

                    # Configure SPE1.
                    [SPE1] static-lsp ingress SPE1toUPE1 destination 4.4.4.9 32 nexthop 3.1.1.2 out-label 30
                    [SPE1] static-lsp egress UPE1toSPE1 incoming-interface vlanif 20 in-label 20

                    # Configure SPE2.
                    [SPE2] static-lsp ingress SPE2toUPE2 destination 5.5.5.9 32 nexthop 4.1.1.2 out-label 50
                    [SPE2] static-lsp egress UPE2toSPE2 incoming-interface vlanif 50 in-label 40

         Step 6 Enable MPLS L2VPN on UPEs and configure UPEs to access SPEs through static
                VPWS.
                    # Configure UPE1.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                       727
VPN Configuration
VPN Configuration                                                                                 6 VPLS Configuration

                    [UPE1] mpls l2vpn
                    [UPE1-l2vpn] quit
                    [UPE1] interface vlanif 10
                    [UPE1-Vlanif10] mpls static-l2vc destination 1.1.1.9 transmit-vpn-label 100 receive-vpn-label 100
                    [UPE1-Vlanif10] quit

                    # Configure UPE2.
                    [UPE2] mpls l2vpn
                    [UPE2-l2vpn] quit
                    [UPE2] interface vlanif 60
                    [UPE2-Vlanif60] mpls static-l2vc destination 3.3.3.9 transmit-vpn-label 100 receive-vpn-label 100
                    [UPE2-Vlanif60] quit

         Step 7 Enable MPLS L2VPN and configure a VSI on SPEs.

                    # Configure SPE1.
                    [SPE1] mpls l2vpn
                    [SPE1-l2vpn] quit
                    [SPE1] vsi V100 static
                    [SPE1-vsi-V100] pwsignal ldp
                    [SPE1-vsi-V100-ldp] vsi-id 100
                    [SPE1-vsi-V100-ldp] mac-withdraw enable
                    [SPE1-vsi-V100-ldp] peer 3.3.3.9
                    [SPE1-vsi-V100-ldp] peer 4.4.4.9 static-upe trans 100 recv 100
                    [SPE1-vsi-V100-ldp] quit
                    [SPE1-vsi-V100] quit

                    # Configure SPE2.
                    [SPE2] mpls l2vpn
                    [SPE2-l2vpn] quit
                    [SPE2] vsi V100 static
                    [SPE2-vsi-V100] pwsignal ldp
                    [SPE2-vsi-V100-ldp] vsi-id 100
                    [SPE2-vsi-V100-ldp] mac-withdraw enable
                    [SPE2-vsi-V100-ldp] peer 1.1.1.9
                    [SPE2-vsi-V100-ldp] peer 5.5.5.9 static-upe trans 100 recv 100
                    [SPE2-vsi-V100-ldp] quit
                    [SPE2-vsi-V100] quit

                    ----End

Verifying the Configuration
                    # Check static VC information on UPE1.
                    [UPE1] display mpls static-l2vc interface vlanif 10
                     *Client Interface     : Vlanif10 is up
                      AC Status          : up
                      VC State          : up
                      VC ID            :0
                      VC Type           : VLAN
                      Destination         : 1.1.1.9
                      Transmit VC Label : 100
                      Receive VC Label        : 100
                      Label Status        :0
                      Token Status          :0
                      Control Word           : Disable
                      VCCV Capability         : alert ttl lsp-ping bfd
                      active state      : active
                      TTL Value           :1
                      Link State        : up
                      Tunnel Policy        : --
                      PW Template Name            : --
                      Main or Secondary : Main
                      load balance type : flow


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            728
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration

