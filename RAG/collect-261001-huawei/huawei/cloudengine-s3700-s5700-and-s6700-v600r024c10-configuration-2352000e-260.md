---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-260
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [38148, 38281]
sha256: 644b7ca50e1cac4ed51be0e31d82a691a03fac80dcd9a76b5e46809d09d80d93
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                             The types of interfaces on both ends of an Eth-Trunk member link must
                             be the same.
                        g.   Return to the system view.
                             quit

                        h.   Enter the Eth-Trunk interface view.
                             interface eth-trunk trunk-id

                        i.   Bind the Eth-Trunk interface to a VSI.
                             l2 binding vsi vsi-name

                    ●   Bind an Eth-Trunk sub-interface to a VSI.

                        Perform the following steps on the PEs at both ends of a PW.

                        a.   Enter the system view.
                             system-view

                        b.   Create an Eth-Trunk interface.
                             interface eth-trunk trunk-id

                        c.   Return to the system view.
                             quit

                        d.   Enter the view of the interface that needs to be added to the Eth-Trunk
                             interface.
                             interface interface-type interface-number

                        e.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        f.   Add the interface to the Eth-Trunk interface.
                             eth-trunk trunk-id

                             Before adding a member interface to an Eth-Trunk, ensure that the
                             member interface does not have Layer 3 configurations such as an IP
                             address and has no services configured.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                      610
VPN Configuration
VPN Configuration                                                                              6 VPLS Configuration


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


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                  611
VPN Configuration
VPN Configuration                                                                               6 VPLS Configuration


                         f.   Bind the sub-interface to a VSI.
                              l2 binding vsi vsi-name
                    ●    Bind a QinQ stacking sub-interface to a VSI.
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
                         f.   Bind the sub-interface to a VSI.
                              l2 binding vsi vsi-name

                    ----End

