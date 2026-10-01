---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-150
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [21568, 21703]
sha256: f8ce892fc2e5bf12c5e2d4856a451e62f7f83dcc5b4070c8b1f2ce9178e36e42
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The command output shows that the tunnel is up. The preceding information is
                 only a part of the command output. The ... symbol is an ellipsis, indicating that
                 some information is not presented.
         Step 7 Configure an SRLG.
                 # Add the links PE1 -> P1 and PE1 -> P4 to SRLG1.
                 [PE1] interface vlanif 100
                 [PE1-Vlanif100] mpls te srlg 1
                 [PE1-Vlanif100] quit
                 [PE1] interface vlanif 300
                 [PE1-Vlanif300] mpls te srlg 1
                 [PE1-Vlanif300] quit

                 # Configure an SRLG-based path calculation mode on the ingress PE1 of the
                 tunnel.
                 [PE1] mpls
                 [PE1-mpls] mpls te srlg path-calculation strict
                 [PE1-mpls] quit

                 # After the configuration is complete, check SRLG information and SRLG member
                 interfaces. The following example uses the command output on P1.
                 [P1] display mpls te srlg all
                 Total SRLG supported : 1024
                 Total SRLG configured : 1


                  SRLG   1:           Vlanif100           Vlanif200

                 # Check the SRLGs to which an interface belongs. The following example uses the
                 command output on PE1.
                 [PE1] display mpls te link-administration srlg-information

                  SRLGs on Vlanif100:
                          1

                  SRLGs on Vlanif300:
                          1

                 # Check SRLG TEDB information. The following example uses the command
                 output on PE1.
                 [PE1] display mpls te cspf tedb srlg 1
                 Interface-Address IGP-Type Area
                 10.1.1.1        ISIS    Level-1
                 10.3.1.1        ISIS    Level-1
                 10.1.1.1        ISIS    Level-2
                 10.3.1.1        ISIS    Level-2

         Step 8 Configure CR-LSP hot standby.
                 # Configure PE1.
                 [PE1] interface tunnel1
                 [PE1-Tunnel1] mpls te backup hot-standby
                 [PE1-Tunnel1] quit

                 ----End

Verifying the Configuration
                 # Disable VLANIF800 on PE1.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                       361
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                 [PE1] interface vlanif 800
                 [PE1-Vlanif800] shutdown
                 [PE1-Vlanif800] quit

                 # Run the display mpls te hot-standby state interface tunnel1 command on
                 PE1 again. The command output shows that the hot-standby CR-LSP index is 0x0,
                 indicating that no hot-standby CR-LSP is established and the path in the same
                 SRLG is excluded.

Configuration Scripts
                 ●      PE1
                        #
                        sysname PE1
                        #
                        vlan batch 100 300 800
                        #
                        mpls lsr-id 5.5.5.5
                        #
                        mpls
                         mpls te
                         mpls te srlg path-calculation strict
                         mpls te cspf
                         mpls rsvp-te
                        #
                        explicit-path main
                         next hop 10.3.1.2
                         next hop 10.6.1.2
                         next hop 6.6.6.6
                        #
                        isis 1
                         cost-style wide
                         network-entity 10.0000.0000.0005.00
                         traffic-eng level-1-2
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.252
                         mpls
                         mpls te
                         mpls te srlg 1
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.3.1.1 255.255.255.252
                         mpls
                         mpls te
                         mpls te srlg 1
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface Vlanif800
                         ip address 10.8.1.1 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/3
                         port link-type trunk


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      362
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

