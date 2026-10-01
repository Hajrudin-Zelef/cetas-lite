---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-250
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [36717, 36886]
sha256: 96b4b11b3c9505d0938a6b67bc61d23a898164ed875187754ea3877b7c4c4451
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                          <HUAWEI> tracert vc -vpn-instance vpn1 1 2 label-alert full-lsp-path
                          TTL Replier       Time Type       Downstream
                           0                    Ingress 10.2.2.2/[21505 1026 ]
                           1  10.2.2.2      60 ms Transit 10.3.3.3/[1026 ]
                           2  10.3.3.3      50 ms Transit 10.4.4.4/[3 ]
                           3  10.1.1.1      70 ms Egress

                    The preceding command output shows each node on the PW path and the
                    response time of each node.


5.14 Troubleshooting VPWS

5.14.1 An LDP VPWS PW Cannot Go Up

Fault Symptom
                    After LDP VPWS is configured, the local end cannot ping the remote end, and the
                    VC status is down.


Procedure
         Step 1 Check whether the two ends use the same encapsulation type and MTU value.

                    Run the display mpls l2vc vc-id command to check VC information.
                    <HUAWEI> display mpls l2vc 102
                     Total LDP VC : 1 0 up    1 down

                     *client interface       : Vlanif10 is up
                       Administrator PW             : no
                       session state        : down
                       AC status           : up
                       Ignore AC state          : disable
                       VC state           : down
                       Label state         :0
                       Token state           :0
                       VC ID             : 102
                       VC type            : VLAN
                       destination          : 2.2.2.2
                       local VC label        : 1032           remote VC label   :0
                       control word           : disable
                       remote control word : none
                       forwarding entry           : not exist
                       local group ID         :0
                       remote group ID             :0
                       local AC OAM State            : up
                       local PSN OAM State : up
                       local forwarding state : not forwarding
                       local status code        : 0x1
                       BFD for PW              : unavailable
                       VCCV State             : up
                       manual fault            : not set
                       active state        : inactive
                       link state        : down
                       local VC MTU              : 1500         remote VC MTU        :0
                       local VCCV            : alert ttl lsp-ping bfd
                       remote VCCV                : none
                       tunnel policy name           : --
                       PW template name                : --
                       primary or secondary : primary
                    .....


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                  583
VPN Configuration
VPN Configuration                                                                     5 VPWS Configuration


                    If the two ends use different encapsulation types or MTU values, configure the
                    same type of AC interfaces on the two ends and run the mpls mtu command to
                    set the same MTU value on the two ends.
                    If both ends use the same encapsulation type and MTU value but the fault
                    persists, go to step 2.

                           NOTE

                         A VC can go up only when both ends use the same encapsulation type and same MTU
                         value.

         Step 2 Check whether VC IDs at both ends are the same.
                    <HUAWEI> display mpls l2vc 102
                     Total LDP VC : 1 0 up    1 down

                     *client interface      : Vlanif10 is up
                       Administrator PW         : no
                       session state       : up
                       AC status          : up
                       Ignore AC state        : disable
                       VC state          : down
                       Label state        :0
                       Token state         :0
                       VC ID            : 102
                       VC type           : VLAN
                    .....

                    If the VC IDs are different, run the undo mpls l2vc command on one end to
                    delete the existing VC ID, and then run the mpls l2vc command to set the VC ID
                    to the same as that on the other end.
                    If the two ends use the same VC ID but the fault persists, go to step 3.

                           NOTE

                         A VC can go up only when both ends use the same VC ID.

         Step 3 Check whether the control words on both ends are the same.
                    <HUAWEI> display mpls l2vc 102
                     Total LDP VC : 1 0 up    1 down

                    *client interface      : Vlanif10 is up
                     Administrator PW            : no
                     session state        : up
                     AC status           : up
                     Ignore AC state          : disable
                     VC state           : down
                     Label state         :0
                     Token state           :0
                     VC ID             : 102
                     VC type            : VLAN
                     destination          : 2.2.2.2
                     local VC label        : 1032          remote VC label   : 1500
                     control word           : disable
                     remote control word : none
                     forwarding entry          : not exist
                     local group ID         :0
                     remote group ID            :0
                     local AC OAM State           : up
                     local PSN OAM State : up
                     local forwarding state : not forwarding
                     local status code        : 0x1
                     BFD for PW              : unavailable
                     VCCV State             : up


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                            584
VPN Configuration
VPN Configuration                                                               5 VPWS Configuration

                       manual fault         : not set
                       active state      : inactive
                       link state       : down
                       local VC MTU          : 1500        remote VC MTU   :0
                       local VCCV         : alert ttl lsp-ping bfd
                       remote VCCV            : none
                       tunnel policy name       : --
                       PW template name           : --
                       primary or secondary : primary
                    .....

                    A VC can go up only when both ends use the same control word. If the control
                    words on both ends are different, run the undo mpls l2vc command to delete the
                    VC connection on one end, and then run the mpls l2vc command to create a VC
                    connection and configure the same control word on both ends.

                    ----End




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                   585
VPN Configuration
VPN Configuration                                                           6 VPLS Configuration




                                                    6         VPLS Configuration


