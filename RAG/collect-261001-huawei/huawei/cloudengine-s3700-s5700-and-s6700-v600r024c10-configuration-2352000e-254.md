---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-254
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [37289, 37394]
sha256: 5efabc9f666a636116ba40a91648f517113f9fc62b3897cf0583e53cff44754a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPLS Access Modes
                    ●   VLANIF interface in switching or routing mode
                        There are two types of VLANIF interfaces:
                        –   A VLANIF interface in routing mode is multiplexed from a physical
                            interface. For example, a GE interface can be divided into multiple sub-
                            interfaces, with each sub-interface acting as a VLANIF interface in routing
                            mode.
                        –   A VLANIF interface in switching mode is a logical interface, but not the
                            sub-interface of a physical interface. A VLANIF interface in switching
                            mode can contain multiple physical interfaces and receive VLAN packets
                            from these physical interfaces.
                        A physical interface in a VLANIF interface in switching mode can send VLAN
                        packets in the following modes:
                        –   Access mode: allows only VLAN packets with the default VLAN ID to pass
                            through.
                        –   Trunk mode: allows only VLAN packets with the VLAN ID of the local
                            VLANIF interface to pass through.
                    ●   CE-to-PE access mode
                        A CE can access a PE in the following modes:
                        –   Through an access port: An access port allows only default VLAN packets
                            of this port to pass through. The VLAN packets on this physical port are
                            untagged.
                            You can assign multiple access ports of the PE to a VLAN for user access.
                        –   Through a trunk port: A trunk port allows the packets of multiple VLANs
                            to pass through. Packets of the default VLAN (one of these VLANs) are
                            untagged, whereas packets of other VLANs are tagged. You can connect
                            the trunk port of the PE to the Ethernet switch to allow the access of
                            multiple VLAN users.

6.2.2 PW Signaling Protocols
PW Signaling Protocols and VPLS Implementation Modes
                    Typically, LDP and BGP are used as the PW signaling protocols. VPLS implemented
                    based on signaling protocols include LDP VPLS, BGP VPLS, and BGP AD VPLS.
                    Table 6-4 compares the VPLS implementation modes.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          595
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration


                    Table 6-4 Comparison between VPLS implementation modes
                    Type      Description      Characteristics                       Application
                                                                                     Scenario

                    LDP       VPLS that uses   ● Protocol implementation is          LDP VPLS
                    VPLS      LDP as the         simple, and PEs do not need to      applies to
                              signaling          deliver high performance, but       networks that
                              protocol is        VPN member discovery must           have a small
                              also called        be manually configured.             number of sites
                              LDP VPLS.        ● After a PE is added, PWs            or do not need
                                                 between the newly added PE          to span ASs,
                                                 and existing PEs need to be         especially when
                                                 established.                        PEs do not run
                                                                                     BGP.
                                               ● An LDP session needs to be
                                                 established between every two
                                                 PEs. The number of LDP
                                                 sessions is directly proportional
                                                 to the number of PEs squared.
                                               ● Labels are distributed to PEs
                                                 on demand, ensuring high
                                                 label utilization.
                                               ● If a VPLS network spans more
                                                 than one autonomous system
                                                 (AS), VSIs in each AS must use
                                                 the same VSI ID range.

                    BGP       VPLS that uses   ● PEs must run BGP and deliver        BGP VPLS
                    VPLS      BGP as the         high performance. Automatic         applies to the
                              signaling          VPN member discovery is             core layer of
                              protocol is        supported, simplifying user         large networks
                              also called        operations.                         where PEs run
                              BGP VPLS.        ● After a PE is added,                BGP and inter-
                                                 configurations on existing PEs      AS
                                                 do not need to be modified, as      communication
                                                 long as the total number of         is required.
                                                 PEs does not exceed the
                                                 number allowed by the label
                                                 block.
                                               ● Route reflectors (RRs) are used
                                                 to reduce the number of BGP
                                                 connections, increasing
                                                 network scalability.
                                               ● Usage of label blocks wastes
                                                 label resources to some extent.
                                               ● VPN targets are used to
                                                 identify VPN member
                                                 relationships, allowing a VPLS
                                                 network to span multiple ASs.




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                            596
VPN Configuration
VPN Configuration                                                                6 VPLS Configuration


                     Type        Description      Characteristics                   Application
                                                                                    Scenario

