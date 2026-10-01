---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-314
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [46641, 46803]
sha256: f2ba08165022bb5ce99b15f6db4415da55ddd3edc7d646ad6810f6eda6682a5c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Table 6-9 Interworking by default
                     Item             Hub AC            Spoke AC         Hub PW           Spoke PW

                     Hub AC           F                 T                T                T

                     Spoke AC         T                 T                T                T

                     Hub PW           T                 T                F                T

                     Spoke PW         T                 T                T                T




Isolation of Users in Different VSIs
                    If PE resources are sufficient and the network structure is clear, you can use
                    different VSIs to isolate traffic of different users. In this way, users are grouped
                    and allocated to different VSIs. Users in a VSI cannot communicate with users in
                    another VSI.
                    In Figure 6-38, CE1, CE2, CE3, CE4, and CE5 use the same type of service. It is
                    required that CE1, CE3, and CE5 communicate with one another, CE2 and CE4
                    communicate with each other, and CE1, CE3, and CE5 not communicate with CE2
                    and CE4. To meet the requirements, different VSIs can be configured to isolate
                    user traffic.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              749
VPN Configuration
VPN Configuration                                                                    6 VPLS Configuration


                    Figure 6-38 Network diagram of using VSIs to isolate users




                    This method has the following advantages:
                    ●   Ensures that the logical network structure is clear, facilitating management
                        and control.
                    ●   Reduces the MAC addresses learned by and resources used by different VSIs.
                    ●   Facilitates fault locating and maintenance.
                    The disadvantage is that adjusting mutual access requirements has a great impact.

Isolation of Users of the Same Service in the Same VSI
                    Service isolation requirements of a VSI are classified into the following types:
                    ●   Local access users in the same VSI are isolated as needed.
                    ●   Local access users and remote access users in the same VSI are isolated.
                    In a common VPLS scenario, the default attribute of an AC interface is spoke, and
                    that of a PW is hub.
                    In Figure 6-39, CE1, CE2, CE3, CE4, and CE5 belong to the same VPN. The local
                    CEs (CE1, CE2, and CE3) connected to PE1 can communicate with one another and
                    with the remote CEs, CE4 (connected to PE2) and CE5 (connected to PE3).
                    However, CE4 and CE5 cannot communicate with each other due to their VSI PW
                    attribute being set to hub.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              750
VPN Configuration
VPN Configuration                                                                    6 VPLS Configuration


                    Figure 6-39 Isolation of common VPLS services




                    To prevent all local users on PE1 from communicating with each other, a VSI is
                    configured on PE1 and bound to PE1's AC interface, and traffic forwarding in
                    spoke mode is disabled. As shown in Table 6-10, services on spoke ACs are
                    isolated from one another.

                    Table 6-10 Interworking after traffic forwarding in spoke mode is disabled in VPLS
                     Item             Hub AC           Spoke AC          Hub PW           Spoke PW

                     Hub AC           F                T                 T                T

                     Spoke AC         T                F                 T                F

                     Hub PW           T                T                 F                T

                     Spoke PW         T                F                 T                F




                    If the AC interface attribute of the VSI is changed from spoke to hub and traffic
                    exchange between the hub AC and hub PW is disabled, communication between
                    some local users on PE1 and between local users on PE1 and remote users is
                    isolated, effectively isolating different users of the same service in the same VSI.
                    In an HVPLS scenario, the default attributes of AC interfaces and PWs between
                    SPEs and UPEs are spoke, and the default attribute of PWs between SPEs is hub.
                    In Figure 6-40, when the SPE designates the UPEs as peers, the attribute of the
                    PWs between the SPE and UPEs changes to spoke. Consequently, all local CEs
                    (CE1, CE2, and CE3) connected to the SPE can communicate with one another, as
                    well as with the remote CEs, CE4 (connected to UPE1) and CE5 (connected to
                    UPE2). Additionally, CE4 and CE5 can communicate with each other. Disabling
                    traffic exchange in spoke mode disables traffic exchange between spoke ACs,
                    between spoke ACs and UPE PWs, and between UPE PWs.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             751
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                    Figure 6-40 HVPLS service isolation




6.15.2 Configuring VPLS Service Isolation
Prerequisites
                    Before configuring VPLS service isolation, you have completed the following tasks:
                    ●      For LDP VPLS, run the pwsignal ldp and vsi-id vsi-id commands.
                    ●      For BGP VPLS, run the pwsignal bgp and route-distinguisher route-
                           distinguisher commands.
                    ●      For BGP AD VPLS, run the bgp-ad and vpls-id vpls-id commands.

Context
                    When multiple users are bound to the same VSI, you can prevent them from
                    communicating with each other by running the isolate spoken command. This
                    enables forwarding isolation between AC interfaces, resulting in the VSI attribute
                    of all AC interfaces in the VSI being set to spoke.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the created VSI.
                    vsi vsi-name

         Step 3 Enable forwarding isolation between AC interfaces in the VSI.
                    isolate spoken

                    By default, forwarding isolation is disabled between AC interfaces in a VSI.
         Step 4 Return to the system view.
                    quit

         Step 5 (Optional) Set the VSI attribute of an AC interface to hub so that this AC interface
                can communicate with other AC interfaces in the same VSI.
                    1.     Enter the AC interface view.
                           interface interface-type interface-number
                    2.     Switch the interface working mode to Layer 3.
                           undo portswitch

                           Determine whether to perform this step based on the current interface
                           working mode.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                      752
VPN Configuration
VPN Configuration                                                                               6 VPLS Configuration


