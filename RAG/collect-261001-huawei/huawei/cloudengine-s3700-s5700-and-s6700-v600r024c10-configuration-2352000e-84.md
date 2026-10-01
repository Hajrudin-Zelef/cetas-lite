---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-84
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [11613, 11741]
sha256: cd9d9c11d048b0330be63ac28728d11df435733d2619d8d4aac1f26d8c158c8a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           184
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


                    5.   After receiving the VPNv4 route, the UPE converts the route into an IPv4 route
                         and imports it into its VPN IPv4 routing table if the route's next hop is
                         reachable.
                    6.   The UPE advertises the IPv4 route to CE1 using the IP protocol.

                    Figure 3-35 Route advertisement from Device1 to CE1 on an HoVPN




Route Advertisement from Device1 to CE1 on an H-VPN
                    Figure 3-36 shows route advertisement from Device1 to CE1 on an H-VPN.
                    1.   Device1 advertises an IPv4 route to the NPE using the IP protocol.
                    2.   The NPE applies for label L3 for the received IPv4 route and converts it into a
                         VPNv4 route. Then, the NPE sets itself as the next hop of the route and
                         advertises it to the SPE.
                    3.   After receiving the VPNv4 route, the SPE saves label L3 locally and applies for
                         label L4 for the route. Then, the SPE sets itself as the next hop of the route
                         and advertises it to the UPE.
                    4.   After receiving the VPNv4 route, the UPE converts the route into an IPv4 route
                         and imports it into its VPN IPv4 routing table if the route's next hop is
                         reachable.
                    5.   The UPE advertises the IPv4 route to CE1 using the IP protocol.

                    Figure 3-36 Route advertisement from Device1 to CE1 on an H-VPN




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           185
VPN Configuration
VPN Configuration                                                            3 IPv4 L3VPN Configuration


Packet Forwarding from Device1 to CE1 on an HoVPN or H-VPN
                    Figure 3-37 shows packet forwarding from Device1 to CE1 on an HoVPN or H-
                    VPN.
                    1.   Device1 sends a VPN packet to the NPE.
                    2.   After receiving the packet, the NPE searches its VPN forwarding table for a
                         tunnel to forward the packet based on the destination address of the packet.
                         Then, the NPE adds an inner label L2 and an outer label Lu to the packet and
                         sends the packet to the SPE over the found tunnel.
                    3.   After receiving the packet, the SPE replaces the outer label Lu with Lv and the
                         inner label L2 with L1. Then, the SPE sends the packet to the UPE over the
                         same tunnel.
                    4.   After receiving the packet, the UPE removes the outer label Lv, searches for a
                         VPN instance corresponding to the packet based on the inner label L1, and
                         removes the inner label L1 after the VPN instance is found. Finally, the UPE
                         sends the packet through this outbound interface to CE1. At this time, the
                         packet is a native IP packet.

                    Figure 3-37 Packet forwarding from Device1 to CE1 on an HoVPN or H-VPN




Packet Forwarding from CE1 to Device1 on an HoVPN
                    Figure 3-38 shows packet forwarding from CE1 to Device1 on an HoVPN.
                    1.   CE1 sends a VPN packet to the UPE.
                    2.   After receiving the packet, the UPE searches its VRF table for a tunnel to
                         forward the packet based on the destination address of the packet (the UPE
                         does so by matching the destination address of the packet against the
                         forwarding entry for the default or summary route). Then, the UPE adds an
                         inner label L4 and an outer label Lv to the packet and sends the packet to the
                         SPE over the found tunnel.
                    3.   After receiving the packet, the SPE removes the outer label Lv and searches
                         for the VPN instance corresponding to the packet based on the inner label L4.
                         Then, the SPE removes the inner label L4 and searches the VPN forwarding
                         table of the found VPN instance for a tunnel to forward the packet based on
                         the destination address of the packet. Finally, the SPE adds an inner label L3
                         and an outer label Lu to the packet and sends the packet to the NPE over the
                         found tunnel.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                          186
VPN Configuration
VPN Configuration                                                            3 IPv4 L3VPN Configuration


                    4.   After receiving the packet, the NPE removes the outer label Lu, searches for a
                         VPN instance corresponding to the packet based on the inner label L3, and
                         removes the inner label L3 after the VPN instance is found. Finally, the NPE
                         sends the packet through this outbound interface to Device1. At this time, the
                         packet is a native IP packet.

                    Figure 3-38 Packet forwarding from CE1 to Device1 on an HoVPN




Packet Forwarding from CE1 to Device1 on an H-VPN
                    Figure 3-39 shows packet forwarding from CE1 to Device1 on an H-VPN.
                    1.   CE1 sends a VPN packet to the UPE.
                    2.   After receiving the packet, the UPE searches its VPN forwarding table for a
                         tunnel to forward the packet based on the destination address of the packet
                         (the UPE does so by matching the destination address of the packet against
                         the forwarding entries for specific routes received from the SPE). Then, the
                         UPE adds an inner label L4 and an outer label Lv to the packet and sends the
                         packet to the SPE over the found tunnel.
                    3.   After receiving the packet, the SPE replaces the outer label Lv with Lu and the
                         inner label L2 with L3. Then, the SPE sends the packet to the NPE over the
                         same tunnel.
                    4.   After receiving the packet, the NPE removes the outer label Lu, searches for a
                         VPN instance corresponding to the packet based on the inner label L3, and
                         removes the inner label L3 after the VPN instance is found. Finally, the NPE
                         sends the packet through this outbound interface to Device1. At this time, the
                         packet is a native IP packet.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                          187
VPN Configuration
VPN Configuration                                                                    3 IPv4 L3VPN Configuration


                    Figure 3-39 Packet forwarding from CE1 to Device1 on an H-VPN




