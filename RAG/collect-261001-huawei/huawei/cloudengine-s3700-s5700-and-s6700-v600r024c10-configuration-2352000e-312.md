---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-312
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [46355, 46507]
sha256: 9df5fd0eebcffa7c06075c051089678b3ab2d330d83e34968cc38545e245b9e1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.2.2.0 0.0.0.255
                          network 10.2.4.0 0.0.0.255
                        #
                        return
                    ●   UPE2
                        #
                        sysname UPE2
                        #
                        vlan batch 40 60
                        #
                        mpls lsr-id 5.5.5.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        interface Vlanif40
                         ip address 10.2.4.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif60
                         mpls l2vc 3.3.3.9 100
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 60
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface LoopBack1
                         ip address 5.5.5.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 5.5.5.9 0.0.0.0
                          network 10.2.4.0 0.0.0.255
                        #
                        return



6.14 Configuring MAC Address Learning
6.14.1 Understanding MAC Address Learning
Context
                    On an Ethernet network, a port sends unicast packets with unknown destination
                    MAC addresses, broadcast packets, and multicast packets to all the other ports on
                    the local Ethernet segment. As an Ethernet-based technology, VPLS emulates an
                    Ethernet bridge for user networks. To forward packets on a VPLS network, PEs

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                       744
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration


                    must establish MAC address tables and forward packets based on MAC addresses
                    or MAC addresses and VLAN tags.

Related Concepts
                    ●   MAC address learning: PEs create a MAC address table through dynamic MAC
                        address learning and associate destination MAC addresses with PWs.
                    ●   MAC address aging: If a MAC address entry learned by a PE is no longer used,
                        an aging mechanism is required to remove the entry. If a MAC address entry
                        is not updated within a specified period of time, this entry will be aged out.

Implementation
                    PEs establish MAC address tables through dynamic MAC address learning and
                    associate destination MAC addresses with PWs. The following table describes the
                    MAC address learning process.

                    Table 6-8 MAC address learning process
                     MAC Address             Description
                     Learning Process

                     Learning MAC            After receiving packets from a CE, a PE maps their
                     addresses from user-    source MAC addresses to AC interfaces. Figure 6-37
                     side packets            shows a mapping example with Port1.

                     Learning MAC            A PW consists of a pair of MPLS VCs in opposite
                     addresses from PW-      directions. A PW can go up only after MPLS VCs in both
                     side packets            directions are established. After a PE receives a packet
                                             with an unknown source MAC address from a PW, the
                                             PE maps the source MAC address to the PW receiving
                                             the packet.




                    Figure 6-37 shows the process for MAC address learning and flooding on a PE.
                    PC1 and PC2 both belong to VLAN 10. When PC1 pings IP address 1.1.1.2, PC1
                    does not know the MAC address corresponding to this IP address and advertises
                    an ARP Request packet.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          745
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                    Figure 6-37 MAC address learning process




                    1.   PE1 receives the ARP Request packet sent by PC1 through Port1 (belonging to
                         VLAN 10) that is connected to CE1, and adds the MAC address of PC1 to its
                         MAC address table, as shown in the blue section of the MAC address entry.
                    2.   PE1 broadcasts the ARP Request packet to its other ports (PW1 and PW2 can
                         be viewed as ports), as shown by the blue dashed line.
                    3.   PE2 receives the ARP Request packet from PW1, and adds the MAC address of
                         PC1 to its MAC address table, as shown in the blue section of the MAC
                         address entry.
                    4.   PE2 sends the ARP Request packet only to the port connected to CE2 (as
                         indicated by the blue dashed line on PE2), but not to PW1. This ensures that
                         only PC2 receives the ARP Request packet. This follows the VPLS split horizon
                         rule, which ensures that packets received from public network PWs are
                         forwarded only to private networks and not to other public network PWs.
                    5.   PC2 receives the ARP packet from PE2 and finds that it is the destination of
                         this packet. PC2 then sends an ARP Reply packet to PC1 (as indicated by the
                         orange dashed line).
                    6.   PE2 receives the ARP Reply packet sent by PC2 through Port2, and adds the
                         MAC address of PC2 to its MAC address table (as indicated by the orange
                         section of the MAC address entry). The destination MAC address of the ARP
                         Reply packet is the MAC address of PC1 (MAC A). After searching its MAC
                         address table, PE2 sends the ARP Reply packet to PE1 over PW1.
                    7.   PE1 receives the ARP Reply packet from PE2, adds the MAC address of PC2 to
                         its own MAC address table, as shown in the orange section of the MAC
                         address entry. After searching its MAC address table, PE1 sends the ARP Reply
                         packet to PC1.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          746
VPN Configuration
VPN Configuration                                                                                    6 VPLS Configuration


