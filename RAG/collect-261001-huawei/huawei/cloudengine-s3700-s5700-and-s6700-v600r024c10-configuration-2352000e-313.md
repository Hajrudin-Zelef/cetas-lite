---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-313
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [46508, 46640]
sha256: 2d7a0d5eafea5d968a0e822f9dc7b4b46555bb8e42c68397532c9c28e97b2bb0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    8.   PC1 receives the ARP Reply packet of PC2. MAC address learning is then
                         complete.
                    9.   When PE1 broadcasts the ARP Request packet over PW1, it also sends the
                         packet to PE3 over PW2. After PE3 receives the ARP Request packet, PE3 adds
                         the MAC address of PC1 to its MAC address table, as shown in the blue
                         section of the MAC address entry. Based on split horizon, PE3 sends the ARP
                         Request packet only to PC3. Since PC3 is not the destination of the ARP
                         Request packet, it does not send any ARP Reply packet.

6.14.2 Configuring MAC Address Learning
Prerequisites
                    Before configuring MAC address learning, you have completed the following task:
                    ●    Configure LDP VPLS

Context
                    In VPLS, packets are forwarded according to MAC address entries. In most cases,
                    MAC addresses are learned automatically. To defend against attacks and resolve
                    faults, the device provides a mechanism to manage VSI-based MAC addresses.
                    A physical interface can belong to multiple VLANs, and multiple VLANIF interfaces
                    can be bound to the same VSI. That is, a VSI associated with a physical interface
                    can be bound to multiple VLANs, and a VSI bound to a VLANIF interface can be
                    associated with multiple physical interfaces. Given this, when configuring static
                    MAC address entries or blackhole entries for the VSI bound to a VLANIF interface,
                    you must specify both the physical interface and VLANIF interface.
                    Perform the following steps on the PEs at both ends of a PW.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure a static MAC address entry.
                    mac-address static mac-address { interface-type interface-number | interface-name } { vlanif-type vlanif-
                    number | vlanif-name } vsi vsi-name

                    Static MAC address entries will not be aged out after being created. When
                    receiving a frame with a specific MAC address, the device forwards the frame
                    through the corresponding outbound interface directly. Static MAC address entries
                    will not be lost even if the device is reset or an interface board on the device is hot
                    swapped.
         Step 3 Configure the aging time of dynamic MAC address entries.
                    mac-address aging-time aging-time

         Step 4 Configure global blackhole MAC address entries.
                    mac-address blackhole mac-address { vlan vlan-id | vsi vsi-name }

                    Global blackhole MAC address entries will not be aged out after being created. If
                    the destination or source MAC address of a packet matches a global blackhole
                    MAC address entry, the packet is discarded. This prevents unnecessary MAC

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                              747
VPN Configuration
VPN Configuration                                                                                  6 VPLS Configuration


                    address entries, such as MAC address entries of unauthorized users, from
                    occupying the MAC address table space and protects the network from attacks
                    that use MAC addresses. Global blackhole MAC address entries will not be lost
                    even if the device is reset or a board on the device is hot swapped.
         Step 5 Enter the VSI view.
                    vsi vsi-name [ static | auto ]

                    The names of different VSIs on a device cannot be the same.
         Step 6 Configure LDP as the signaling protocol of the VSI and enter the VSI-LDP view.
                    pwsignal ldp

         Step 7 Configure a VSI ID.
                    vsi-id vsi-id

         Step 8 Return to the VSI view.
                    quit

         Step 9 Enable MAC address learning.
                    mac-learning { enable | disable }

        Step 10 Configure a MAC address learning limit rule.
                    mac-address limit { maximum maxValue | action { discard | forward } | alarm { enable | disable } } *

                            NOTE

                           If the VSI has learned some MAC addresses, run the undo mac-address dynamic command
                           to clear the learned MAC addresses. This ensures that the number of MAC addresses that
                           can be learned is accurately limited.
                           When running the mac-address limit command for the first time, you must run the mac-
                           address limit maximum maxValue command before configuring action and alarm. There
                           is no such requirement if you have run the mac-address limit command.

                    ----End

6.14.3 Verifying the Configuration
Procedure
                    ●      Run the display mac-address aging-time vsi [ vsi-name ] command to check
                           the aging time of MAC address entries.
                    ●      Run the display mac-address [ mac-address ] [ vlan vlan-id | vsi vsi-name ]
                           [ verbose ] command to check information about MAC address entries.
                    ●      Run the display mac-address static [ vsi vsi-name ] command to check static
                           MAC address entries.
                    ●      Run the display mac-address dynamic [ vsi vsi-name ] command to check
                           dynamic MAC address entries.
                    ●      Run the display mac-address blackhole [ vsi vsi-name ] [ verbose ]
                           command to check static blackhole MAC address entries.
                    ●      Run the display mac-address [ vsi vsi-name ] [ interface-type interface-
                           number ] command to check the MAC address learning limit.
                    ●      Run the display vsi [ name vsi-name ] [ verbose ] command to check the
                           MAC address learning mode.
                    ----End

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                              748
VPN Configuration
VPN Configuration                                                                    6 VPLS Configuration




6.15 Configuring VPLS Service Isolation

6.15.1 Understanding VPLS Service Isolation
                    Users of different services can be isolated using different VSIs. Users in the same
                    VSI also need to be isolated.

Service Isolation Modes
                    VPLS networks use a full mesh of PWs and split horizon to prevent loops. Split
                    horizon requires that if a packet is received over a PW of a VSI, the packet is not
                    forwarded over other PWs associated with the VSI. VPLS supports either the hub
                    or spoke service isolation mode. In hub mode, traffic forwarding must comply with
                    split horizon rules. Whereas, in spoke mode, traffic forwarding does not need to
                    comply with split horizon rules. As described in Table 6-9, traffic cannot be
                    exchanged between hub AC interfaces or between hub PWs in a VSI. ("T" indicates
                    that traffic can be exchanged between AC interfaces or between PWs, and "F"
                    indicates that traffic cannot be exchanged between AC interfaces or between
                    PWs.)

