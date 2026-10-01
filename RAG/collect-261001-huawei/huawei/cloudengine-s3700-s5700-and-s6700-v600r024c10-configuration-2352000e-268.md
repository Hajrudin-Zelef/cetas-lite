---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-268
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [39353, 39500]
sha256: 71f4c2adcb732509952e3280473bc62e60aca8635075fd7de778fe51d989afa3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                             Before adding a member interface to an Eth-Trunk, ensure that the
                             member interface does not have Layer 3 configurations such as an IP
                             address and has no services configured.
                             An Ethernet interface can be added to only one Eth-Trunk interface.
                             Before adding an Ethernet interface to an Eth-Trunk interface, ensure that
                             the Ethernet interface does not belong to any Eth-Trunk interface.
                             The types of interfaces on both ends of an Eth-Trunk member link must
                             be the same.
                        g.   Return to the system view.
                             quit

                        h.   Enter the Eth-Trunk sub-interface view.
                             interface eth-trunk trunk-id.subnumber

                        i.   Specify the VLAN to be associated with the Eth-Trunk sub-interface and
                             set the VLAN encapsulation type.
                             dot1q termination vid low-pe-vid

                        j.   Bind the Eth-Trunk sub-interface to a VSI.
                             l2 binding vsi vsi-name

                    ●   Bind a dot1q VLAN tag termination sub-interface to a VSI.
                        Perform the following steps on the PEs at both ends of a PW.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                       630
VPN Configuration
VPN Configuration                                                                              6 VPLS Configuration


                        a.   Enter the system view.
                             system-view

                        b.   Enter the sub-interface view of the AC interface connecting a PE to a CE.
                             interface interface-type interface-number.subinterface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Configure a dot1q VLAN tag termination sub-interface.
                             encapsulation dot1q-termination

                        e.   Configure the sub-interface to terminate single-tagged packets.
                             dot1q termination vid low-pe-vid

                        f.   Bind the sub-interface to a VSI.
                             l2 binding vsi vsi-name

                    ●   Bind a QinQ VLAN tag termination sub-interface to a VSI.
                        Perform the following steps on the PEs at both ends of a PW.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the sub-interface view of the AC interface connecting a PE to a CE.
                             interface interface-type interface-number.subinterface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Configure a QinQ VLAN tag termination sub-interface.
                             encapsulation qinq-termination

                        e.   Configure the sub-interface to terminate double-tagged packets.
                             qinq termination pe-vid pe-vid ce-vid ce-vid [ to high-ce-vid ]

                        f.   Bind the sub-interface to a VSI.
                             l2 binding vsi vsi-name

                    ●   Bind a QinQ stacking sub-interface to a VSI.
                        Perform the following steps on the PEs at both ends of a PW.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the sub-interface view of the AC interface connecting a PE to a CE.
                             interface interface-type interface-number.subinterface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Configure QinQ stacking on the sub-interface.
                             qinq stacking vid low-ce-vid [ to high-ce-vid ] pe-vid pe-vid

                        e.   Configure the sub-interface to terminate double-tagged packets.
                             qinq termination pe-vid pe-vid ce-vid ce-vid [ to high-ce-vid ]


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     631
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                         f.   Bind the sub-interface to a VSI.
                              l2 binding vsi vsi-name

                    ----End

6.7.6 Configuring VPLS Flow Label-based Load Balancing
Context
                    Packets of multiple data flows on the same PW carry the same VC labels which
                    are encapsulated on a PE. When these packets arrive at a P, the P forwards them
                    along only one path, regardless of whether multiple load-balancing paths exist.
                    To load-balance different data flows, configure flow label-based load balancing on
                    the PE. After the configurations are complete, the PE adds flow labels following
                    private network labels when encapsulating data packets, and the P load-balances
                    these data flows based on the flow labels.
                    On the network shown in Figure 6-14, two different data flows need to be
                    transmitted to the peer end through an L2VPN.
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

                    Figure 6-14 Typical networking of flow label-based load balancing




                    Flow label-based load balancing applies to an L2VPN on which multiple links exist
                    between Ps. Flow label-based load balancing allows data flows on the same VPN

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         632
VPN Configuration
VPN Configuration                                                                              6 VPLS Configuration


                    to be load-balanced among different paths based on flow labels for higher
                    network resource utilization.

Procedure
         Step 1 Enter the system view.
                    system-view

