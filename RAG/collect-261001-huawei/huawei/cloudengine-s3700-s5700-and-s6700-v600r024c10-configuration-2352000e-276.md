---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-276
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [40562, 40716]
sha256: ce654392234c45129b9b5dc220f51b15f8ea8c755af1755d75d5f44feb035eb1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                                 ● If the remote PE is configured to accept only VLAN tagged packets, run the
                                   mpls l2vpn default vlan command to configure the default VLAN of the
                                   main interface before binding the local Ethernet interface to the VSI.
                                 ● If the remote PE is configured to accept only double-tagged packets, run the
                                   mpls l2vpn vlan-stacking stack-vlan command to configure the stacked
                                   VLAN of the main interface before binding the local Ethernet interface to the
                                   VSI.
                                 ● After a main interface is bound to a VSI, the mpls l2vpn default vlan or
                                   mpls l2vpn vlan-stacking stack-vlan command configuration cannot be
                                   modified. To modify the configuration, you must unbind the main interface
                                   from the VSI first. This operation will interrupt VPLS services on the main
                                   interface. Exercise caution when performing this operation.
                        e.   Bind the interface to a VSI.
                             l2 binding vsi vsi-name

                    ●   Bind an Ethernet sub-interface to a VSI.

                        Perform the following steps on the PEs at both ends of a PW.

                        a.   Enter the system view.
                             system-view

                        b.   Enter the Ethernet sub-interface view.
                             interface interface-type interface-number.subinterface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Specify the VLAN to be associated with the Ethernet sub-interface and
                             set the VLAN encapsulation type.
                             dot1q termination vid low-pe-vid


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   651
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration


                        e.   Bind the sub-interface to a VSI.
                             l2 binding vsi vsi-name

                    ●   Bind a VLANIF interface to a VSI.
                        Perform the following steps on the PEs at both ends of a PW.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the VLANIF interface view.
                             interface vlanif vlan-id

                        c.   Bind the VLANIF interface to the VSI.
                             l2 binding vsi vsi-name

                    ●   Bind an Eth-Trunk interface to a VSI.
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
                             An Ethernet interface can be added to only one Eth-Trunk interface.
                             Before adding an Ethernet interface to an Eth-Trunk interface, ensure that
                             the Ethernet interface does not belong to any Eth-Trunk interface.
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

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                      652
VPN Configuration
VPN Configuration                                                                            6 VPLS Configuration


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

