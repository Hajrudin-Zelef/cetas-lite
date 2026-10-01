---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-277
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [40717, 40844]
sha256: 8a6f3f3174f00c127452db667fc4a12220d4192e28b17d0588a2d72e9505c550
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        b.   Enter the sub-interface view of the AC interface connecting a PE to a CE.
                             interface interface-type interface-number.subinterface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Configure a dot1q VLAN tag termination sub-interface.
                             encapsulation dot1q-termination


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                653
VPN Configuration
VPN Configuration                                                                               6 VPLS Configuration


                        e.    Configure the sub-interface to terminate single-tagged packets.
                              dot1q termination vid low-pe-vid

                        f.    Bind the sub-interface to a VSI.
                              l2 binding vsi vsi-name

                    ●   Bind a QinQ VLAN tag termination sub-interface to a VSI.
                        Perform the following steps on the PEs at both ends of a PW.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the sub-interface view of the AC interface connecting a PE to a CE.
                              interface interface-type interface-number.subinterface-number

                        c.    Switch the interface working mode to Layer 3.
                              undo portswitch

                              Determine whether to perform this step based on the current interface
                              working mode.
                        d.    Configure a QinQ VLAN tag termination sub-interface.
                              encapsulation qinq-termination

                        e.    Configure the sub-interface to terminate double-tagged packets.
                              qinq termination pe-vid pe-vid ce-vid ce-vid [ to high-ce-vid ]

                        f.    Bind the sub-interface to a VSI.
                              l2 binding vsi vsi-name

                    ●   Bind a QinQ stacking sub-interface to a VSI.
                        Perform the following steps on the PEs at both ends of a PW.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the sub-interface view of the AC interface connecting a PE to a CE.
                              interface interface-type interface-number.subinterface-number

                        c.    Switch the interface working mode to Layer 3.
                              undo portswitch

                              Determine whether to perform this step based on the current interface
                              working mode.
                        d.    Configure QinQ stacking on the sub-interface.
                              qinq stacking vid low-ce-vid [ to high-ce-vid ] pe-vid pe-vid

                        e.    Configure the sub-interface to terminate double-tagged packets.
                              qinq termination pe-vid pe-vid ce-vid ce-vid [ to high-ce-vid ]

                        f.    Bind the sub-interface to a VSI.
                              l2 binding vsi vsi-name

                    ----End

6.8.5 Verifying the Configuration
Procedure
                    ●   Run the display vsi [ name vsi-name ] [ verbose ] command to check VSI
                        information.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     654
VPN Configuration
VPN Configuration                                                                            6 VPLS Configuration


                    ●   Run the display vsi bgp-ad { import-vt | export-vt | remote-export-vt }
                        command to check information about the VPN targets of the local and
                        remote devices on a BGP AD VPLS network.
                    ●   Run the display vsi bgp-ad remote vpls-id vpls-id command to check
                        information about a specified remote PE on a BGP AD VPLS network.
                    ●   Run the display vpls connection [ bgp-ad | vsi vsi-name ] [ down | up ]
                        [ verbose ] command to check information about BGP AD VPLS connections.
                    ●   Run the display bgp l2vpn-ad routing-table vpls-ad command to check BGP
                        L2VPN AD routing information.
                    ----End

6.8.6 Example for Configuring BGP AD VPLS
Networking Requirements
                    Figure 6-19 shows a backbone network built by an enterprise. The enterprise has
                    a large number of branch sites, only three of which are shown in this example.
                    The network environment often changes. Site1 connects to the backbone network
                    by connecting CE1 to PE1, Site2 connects to the backbone network by connecting
                    CE2 to PE2, and Site3 connects to the backbone network by connecting CE3 to
                    PE3. Users at Site1, Site2, and Site3 need to communicate at Layer 2, and user
                    information needs to be retained in Layer 2 packets when the packets are
                    transmitted over the backbone network.

                         NOTE

                        To avoid loops in this scenario, ensure that all connected interfaces have STP disabled and
                        are removed from VLAN 1. If STP is enabled and VLANIF interfaces of switches are used to
                        construct a Layer 3 ring network, a specific connected interface between the PEs will be
                        blocked. As a result, Layer 3 services on the network cannot run properly.


                    Figure 6-19 Network diagram of configuring BGP AD VPLS
                         NOTE

                        In this example, interface1, interface2, and interface3 represent 10GE1/0/1, 10GE1/0/2, and
                        10GE1/0/3, respectively.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      655
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration




