---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-64
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [8915, 9099]
sha256: 46056bbd6b8204cea2dd516c0b16c3f8bfe1cda48a77ee0a59a00d3c65950c16
---

                 # Repeat this step for PE2. For configuration details, see Configuration Scripts in
                 this section.
         Step 3 Configure LDP session protection.
                 # Configure PE1.
                 [PE1] mpls ldp
                 [PE1-mpls-ldp] session protection duration infinite
                 [PE1-mpls-ldp] quit

                 # Configure PE2.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                        149
MPLS Configuration
MPLS Configuration                                                                                 3 MPLS LDP Configuration

                 [PE2] mpls ldp
                 [PE2-mpls-ldp] session protection duration infinite
                 [PE2-mpls-ldp] quit

         Step 4 Verify the configuration.

                 # Shut down VLANIF200 on PE1 to simulate a link fault. Run the display mpls ldp
                 remote-peer command. The command output shows that LDP session protection
                 has taken effect.
                 [PE1] display mpls ldp remote-peer
                                    LDP Remote Entity Information
                  ------------------------------------------------------------------------------
                  Remote Peer Name : pe2
                  Description        : ----
                  Remote Peer IP : 3.3.3.3                 LDP ID        : 1.1.1.1:0
                  Transport Address : 1.1.1.1              Entity Status : Active

                  Configured Keepalive Hold Timer : 45 Sec
                  Configured Keepalive Send Timer : ----
                  Configured Hello Hold Timer            : 45 Sec
                  Negotiated Hello Hold Timer             : 45 Sec
                  Configured Hello Send Timer            : ----
                  Configured Delay Timer               : 10 Sec
                  Hello Packet sent/received          : 91/86
                  Label Advertisement Mode               : Downstream Unsolicited
                  Auto-config                   : Session-Protect
                  Manual-config                   : effective
                  Session-Protect effect           : YES
                  Session-Protect Duration           : infinite
                  Session-Protect Remain              : ----
                  ------------------------------------------------------------------------------
                  TOTAL: 1 Remote-Peer(s) Found.

                 ----End

Configuration Scripts
                 ●      PE1
                        #
                        sysname PE1
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                        #
                        mpls ldp
                         session protection duration infinite
                         #
                         ipv4-family
                        #
                        isis 1
                         is-level level-2
                         network-entity 10.0000.0000.0001.00
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.252
                         isis enable 1
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.252
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           150
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack0
                         ip address 1.1.1.1 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      PE3
                        #
                        sysname PE3
                        #
                        vlan batch 100 200
                        #
                        isis 1
                         is-level level-2
                         network-entity 10.0000.0000.0003.00
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.252
                         isis enable 1
                        #
                        interface Vlanif200
                         ip address 10.1.3.1 255.255.255.252
                         isis enable 1
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack0
                         ip address 2.2.2.2 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                        #
                        mpls ldp
                         session protection duration infinite
                         #
                         ipv4-family
                        #
                        isis 1
                         is-level level-2
                         network-entity 10.0000.0000.0002.00
                        #
                        interface Vlanif100
                         ip address 10.1.3.2 255.255.255.252
                         isis enable 1
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.252


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                       151
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration

                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack0
                         ip address 3.3.3.3 255.255.255.255
                         isis enable 1
                        #
                        return



3.18 Configuring LDP-IGP Synchronization

