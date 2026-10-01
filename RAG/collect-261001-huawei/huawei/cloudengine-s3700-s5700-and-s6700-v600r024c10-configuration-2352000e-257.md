---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-257
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [37707, 37864]
sha256: b2ba8f5b404095a9979c4f5ee0c83824011f78f5f3de174ef78b1487dbc550b6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                602
VPN Configuration
VPN Configuration                                                                               6 VPLS Configuration


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

6.5.3 Verifying the Configuration

Procedure
                    ●   Run the display vsi [ name vsi-name ] [ verbose ] command to check VSI
                        information.
                    ●   Run the display vsi services { all | vsi-name | interface interface-type
                        interface-number | vlan vlan-id | interface interface-name } command to
                        check information about the AC interface associated with the VSI.
                    ●   Run the display mpls label-stack vpls vsi vsi-name peer peer-id vc-id vc-id
                        command to check label stack information in VPLS scenarios.

                    ----End


6.6 Configuring LDP VPLS
Prerequisites
                    Before configuring LDP VPLS, you have completed the following tasks:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     603
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                    ●   Configure IP addresses and an IGP on PEs and Ps.
                    ●   Configure LSR IDs, enable MPLS and MPLS LDP, and establish LDP sessions on
                        PEs and Ps.
                    ●   (Optional) Establish remote MPLS LDP sessions if PEs are not directly
                        connected.
                    ●   Enable MPLS L2VPN on PEs.
                    ●   Establish tunnels between PEs to carry L2VPN services.

6.6.1 Understanding LDP VPLS
Context
                    LDP VPLS uses a static discovery mechanism to discover VPLS members using LDP
                    signaling. VPLS information is carried in extended type-length-value (TLV) fields
                    (type 128 and type 129 FEC TLVs) of LDP signaling messages. During the
                    establishment of a PW, the label distribution mode is downstream unsolicited
                    (DU) and the label retention mode is liberal.

Related Concepts
                    LDP VPLS involves the following concepts:
                    ●   FEC: A set of data flows with certain similarities. Data flows in the same FEC
                        are processed by LSRs in the same way. FECs can be classified based on
                        addresses, service types, or QoS policies, among others.
                    ●   TLV: A highly efficient and expansible coding mode for protocol packets. To
                        support new features, you only need to add new types of TLVs to carry
                        information required by the features.
                    ●   DU: A label distribution mode in which an LSR distributes labels to FECs
                        without having to receive Label Request messages from its upstream LSR.
                    ●   Liberal: A label retention mode in which an LSR retains the label mapping
                        received from a neighboring LSR, regardless of whether the neighboring LSR
                        is its next hop. In liberal label retention mode, an LSR can use the labels sent
                        from neighboring LSRs that are not at the next hop to re-establish a label
                        switched path (LSP). This mode requires more memory and label space than
                        the conservative mode.

Implementation
                    Figure 6-6 shows the process for establishing a PW using LDP signaling.

                    Figure 6-6 Establishing a PW using LDP signaling




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            604

