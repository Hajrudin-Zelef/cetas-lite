---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-247
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [36252, 36442]
sha256: ac8fa9edc4c8f16986298b29114a95a1fb321dcb57deed03a500442e86d389bd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    A BFD session can be either static or dynamic. Determine which one to use as
                    needed.


Procedure
                    ●    Deploy BFD on PEs. For configuration details, see 5.11 Configuring BFD for
                         VPWS.

                    ----End

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            574
VPN Configuration
VPN Configuration                                                                                 5 VPWS Configuration


5.12.5 (Optional) Binding a Service PW to the mPW

Context
                    To accelerate PW fault detection, BFD is typically used. If there are a large number
                    of service PWs with the same source and same destination, you can configure an
                    mPW with the same source and same destination as those of the service PWs and
                    associate these service PWs with the mPW. By tracking the status of the mPW,
                    BFD can quickly detect faults on service PWs associated with the mPW. This
                    method does not require BFD to be configured for service PWs, thereby reducing
                    the number of BFD sessions and conserving both system resources and public
                    network link bandwidth.

Procedure
         Step 1 Create a loopback interface and enter the loopback interface view.
                    interface loopback loopback-number

         Step 2 Create an mPW.
                    mpls l2vc [ instance-name instance-name ] { ip-address | pw-template pw-template-name } * vc-id
                    { [ control-word | tunnel-policy policy-name ] * admin | admin [ control-word | tunnel-policy policy-
                    name ] * | tunnel-policy policy-name admin control-word | control-word admin tunnel-policy policy-
                    name }

         Step 3 Return to the system view.
                    quit

         Step 4 Enter the view of the interface on which the service PW resides.
                    interface interface-type interface-number

         Step 5 Switch the interface working mode to Layer 3.
                    undo portswitch

                    Determine whether to perform this step based on the current interface working
                    mode.

         Step 6 Bind the service PW to the mPW.
                    mpls l2vc [ secondary | bypass ] track admin-vc interface interface-type interface-number

                    ----End

5.12.6 Verifying the Configuration

Procedure
                    ●      Run the display mpls l2vc [ vc-id | interface interface-type interface-
                           number ] command on a PE to check local VPWS connection information.
                    ●      Run the display mpls l2vc track admin-vrrp [ interface interface-type
                           interface-number vrid virtual-router-id ] command to check information
                           about the PW bound to an mVRRP group.
                    ●      Run the display mpls l2vc track admin-vc command to check information
                           about the service PW bound to an mPW.

                    ----End

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                575
VPN Configuration
VPN Configuration                                                                     5 VPWS Configuration




5.13 Maintaining VPWS

5.13.1 Collecting VPWS Traffic Statistics

Context
                    To monitor network operating status and locate faults in an easier way on a VPWS
                    network, configure the traffic statistics collection function.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the AC interface view.
                    interface interface-type interface-number

         Step 3 Switch the interface working mode to Layer 3.
                    undo portswitch

                    Determine whether to perform this step based on the current interface working
                    mode.

         Step 4 Configure VPWS PW traffic statistics collection.
                    mpls l2vpn pw traffic-statistics enable [ secondary | bypass ]

                    ●    If you do not specify secondary and bypass, the traffic statistics collection
                         function applies to the primary VPWS PW.
                    ●    If you specify secondary, the traffic statistics collection function applies to the
                         secondary VPWS PW.
                    ●    If you specify bypass, the traffic statistics collection function applies to the
                         bypass VPWS PW.

                    ----End

5.13.2 Resetting VPWS PWs

Context
                    To reset VPWS information, perform the following operations in the user view.



                        NOTICE

                    Resetting a PW template interrupts the services carried by PWs or the PWs that
                    use the PW template. Exercise caution when using the commands.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                576
VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration


                    Table 5-7 Resetting VPWS PWs
                     Operation                 Command

                     Reset a PW or PW          reset pw [ peer-address peer-address ] pw-id
                     template.                 { ethernet | vlan }

                     Clear VPWS PW traffic     reset traffic-statistics l2vpn pw { all | interface
                     statistics.               interface-type interface-number [ secondary |
                                               bypass ] }




5.13.3 Monitoring the VPWS Operating Status
Context
                    In routine maintenance, you can run the following commands in any view to learn
                    the VPWS operating status.
                    Table 5-8 lists the operations for monitoring the VPWS operating status.

                    Table 5-8 Monitoring the VPWS operating status
                     Operation                Command

                     Check CCC                display vll ccc [ ccc-name | type local ]
                     connection
                     information.

                     Check information        display l2vpn ccc-interface vc-type { all | ldp-vc | ccc |
                     about the interfaces     vpls-vc | static-vc | bgp-vc } [ up | down ]
                     used by L2VPN
                     connections.

                     Check local VPWS         display mpls l2vc [ vc-id | interface interface-type
                     connection               interface-number ]
                     information on PEs.

                     Check remote VPWS        display mpls l2vc remote-info [ vc-id | unmatch |
                     connection               verbose ]
                     information on PEs.

                     Check information        display mpls static-l2vc [ vc-id | interface interface-
                     about static VPWS        type interface-number | state { down | up } | brief ]
                     connections on PEs.

                     Check PW template        display pw-template [ pw-template-name ]
                     information.

                     Check information        display l2vpn error discard
                     about signaling
                     messages discarded
                     by the L2VPN
                     component.



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              577
VPN Configuration
VPN Configuration                                                                5 VPWS Configuration


