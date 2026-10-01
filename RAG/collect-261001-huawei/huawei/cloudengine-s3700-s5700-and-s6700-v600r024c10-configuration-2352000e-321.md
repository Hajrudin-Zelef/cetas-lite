---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-321
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [47768, 47927]
sha256: a769961235d62fc91097917441bac11184bc8c640c0b3603f6dfaa78cf98fb4a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

6.17.3 Configuring MAC Withdraw Loop Detection
Context
                    Data can be forwarded between a hub PW and a spoke PW or between spoke
                    PWs on an HVPLS or VPLS network. Therefore, a loop may be generated during
                    the forwarding of MAC Withdraw messages. As a result, the control plane is
                    unable to converge in time, leading to the loss of some packets. If the loop
                    involves a large number of MAC Withdraw messages, denial of service (DoS)
                    attacks may occur. To address this problem, configure MAC Withdraw loop
                    detection to prevent MAC Withdraw message loops.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enable L2VPN and enter the MPLS L2VPN instance view.
                    mpls l2vpn

         Step 3 Enable MAC Withdraw loop detection.
                    vpls mac-withdraw loop-detect enable

                    By default, MAC Withdraw loop detection is disabled.

                    ----End

Verifying the Configuration
                    ●    Run the display vsi [ name vsi-name ] verbose command to check the MAC
                         Withdraw configuration.
                    ●    Run the display vsi [ name vsi-name ] mac-withdraw loop-detect command
                         to check information about MAC Withdraw loop detection.

6.17.4 Configuring a VSI to Ignore the AC Status
Context
                    If the services running on a legacy network need to switch to a new network, you
                    can configure VSIs to ignore the AC status and check whether the VSIs on the new
                    network can work properly. On the network shown in Figure 6-43, if the services
                    running on the legacy VPLS network need to switch to the new network, and you
                    want to check whether the VSI on the new network can work properly before the
                    service switchover, configure the VSI to ignore the AC status on DeviceD. After the
                    configuration is complete, the VSI on DeviceD remains up before the digital
                    subscriber line access multiplexer (DSLAM) is connected to the new network.

                    AC statuses are classified into the following types:

Issue 01 (2025-03-03)                   Copyright © Huawei Technologies Co., Ltd.                      768
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                    ●   Status of a physical or logical interface bound to a VSI
                    ●   UPE PW status in a VPWS accessing VPLS scenario
                    A VSI can go up only if at least one AC interface or one UPE PW is up. If an AC
                    interface is down and the PW is up, the VSI remains up after being enabled to
                    ignore the AC status. If an AC interface is up and the PW is down, the VSI also
                    remains up after being enabled to ignore the AC status.

                    Figure 6-43 Configuring the VSI to ignore the AC status




                    Perform the following steps on a PE (DeviceD in Figure 6-43):

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Perform the following steps as required:
                    ●   To enable all VSIs on the device to ignore the AC status, perform the
                        following operations:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           769
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration


                        a.    Enter the MPLS L2VPN view.
                              mpls l2vpn

                        b.    Configure VSIs to ignore the AC status.
                              vpls ignore-ac-state

                    ●   To configure a certain VSI to ignore the AC status, perform the following
                        operations:
                        a.    Enter the VSI view.
                              vsi vsi-name static

                        b.    Configure the VSI to ignore the AC status.
                              ignore-ac-state

                    ----End

Follow-up Procedure
                    The vpls ignore-ac-state or ignore-ac-state command is used only during the
                    service switchover from a legacy VPLS network to a new one. After the service
                    switchover is complete, run the undo vpls ignore-ac-state or undo ignore-ac-
                    state command to restore the default settings.


6.18 Configuring ERPS over VPLS
                         NOTE

                        This configuration is supported only by the S6780-H, S6750-H, S6750E-S, S6750-S, S6730E-
                        H-V2, S6730-H-V2, S5755E-H, S5755-H, and S5732-H-V2 series.


6.18.1 Understanding ERPS over VPLS
                    When an ERPS-enabled interface on a switch has a sub-interface or there is a
                    corresponding VLANIF interface on the switch, and the sub-interface or VLANIF
                    interface is bound to a VSI, the sub-interface or VLANIF interface cannot promptly
                    detect the topology change of the main interface. This delays the switch from
                    instructing devices on a VPLS network to update MAC address entries. To solve
                    this problem, configure the topology change notification function on the main
                    interface.
                    Figure 6-44 shows a VPLS network where CEs are connected to PEs. A loop occurs
                    on the network, and PE3 receives two copies of traffic from the remote CEs. To
                    solve this problem, enable ERPS on PE1, CE1, CE2, and PE2 and configure interface
                    2 of CE2 as an RPL owner port to block traffic from CE1. In this way, traffic from
                    CE1 is directly transmitted to PE3 through PE1 without passing through CE2,
                    thereby preventing duplicate traffic or loops.
                    In Figure 6-44, an ERPS ring is connected to a VPLS network through Ethernet
                    sub-interfaces or VLANIF interfaces. To ensure that the VPLS network can
                    promptly detect the topology change of the ERPS ring, enable the topology
                    change notification function on the main interfaces through which PE1 and PE2
                    are connected to the ERPS ring.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  770
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                    Figure 6-44 Network diagram of configuring ERPS over VPLS for CE access




6.18.2 Configuring ERPS over VPLS

Prerequisites
                    ●    PEs on the VPLS backbone network run a routing protocol so that they can
                         communicate with each other.
                    ●    Basic MPLS functions have been configured and LDP LSPs have been
                         established on the VPLS backbone network.
                    ●    VPLS connections have been established between PEs and Ethernet sub-
                         interfaces or VLANIF interfaces have been bound to VSIs.
                    ●    Interfaces on CEs and PEs have been added to the ERPS ring.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the interface view.
                    interface interface-type interface-number

