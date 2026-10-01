---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-218
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [32160, 32251]
sha256: e1a25344a9c812ab5de7f770aae7bdb44f9ef912956f531fc23607ca600f9bb2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                           ● The IDs of VCs using the same encapsulation type must be unique.
                           ● The raw and tagged parameters are available in this command only for Ethernet links.
                           ● If an Ethernet sub-interface is used, run the dot1q termination vid low-pe-vid
                             command in the interface view to configure the encapsulation type and VLAN ID of the
                             Ethernet sub-interface.
                           ● The default tunnel policy uses LDP tunnels for VPWS connections. To use other types of
                             tunnels, you can configure a tunnel policy and set the tunnel-policy policy-name
                             parameter to reference the tunnel policy. For details about how to configure a tunnel
                             policy, see Configuring a Tunnel Policy.

         Step 7 Configure the primary LDP VPWS connection.
                    mpls l2vc { ip-address | pw-template pw-template-name } * vc-id [ [ control-word | no-control-word ] |
                    [ raw | tagged ] | tunnel-policy policy-name | ignore-standby-state ] *

         Step 8 Configure the secondary LDP VPWS connection.
                    mpls l2vc { ip-address | pw-template pw-template-name } * vc-id [ [ control-word | no-control-word ] |
                    [ raw | tagged ] | tunnel-policy policy-name | ignore-standby-state | secondary ] *

                    When configuring the secondary VPWS connection, ensure that the parameter
                    settings for the primary, secondary, and bypass PWs are consistent. Parameter
                    setting inconsistencies may result in the secondary or bypass PW failing to take
                    over traffic if the primary PW fails, leading to service interruption.




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                           514
VPN Configuration
VPN Configuration                                                                               5 VPWS Configuration


                          NOTE

                         In the networking where CEs are asymmetrically connected to PEs:
                         ● Only the PE to which a CE is single-homed requires the configuration of both the
                           primary and secondary PWs. The PEs to which a CE is dual-homed require only the
                           configuration of the primary PW.
                         ● The singled-homed CE interface connected to the corresponding PE requires both
                           primary and secondary IP addresses. When the primary path is available, the CE uses the
                           primary IP address to communicate with the remote CE. When a fault occurs on the
                           primary path, this CE uses the secondary IP address to communicate with the remote
                           CE.
                         In the networking where CEs are asymmetrically connected to PEs, the secondary PW does
                         not transmit data when the primary and backup paths both work properly. If the AC
                         interface on the secondary PW borrows the IP address of the AC interface on the primary
                         PW, the following situations arise:
                         ● The policy of no switchback cannot be configured.
                         ● The local CE has two equal-cost direct routes to the remote CE, with both routes having
                           the same destination address and same next hop. The route that passes through the
                           secondary PW is actually an invalid route.
                         ● If CEs exchange routing information using routing protocols, you need to change the
                           cost or metric of the AC interface on the backup path to be greater than that of the AC
                           interface on the primary path. In this case, the local CE may fail to communicate with
                           the remote CE, but can communicate with other user devices on the remote end.
                         ● If CEs use static routes to exchange routing information, ensure that the preference of
                           the backup route is lower than that of the primary route (a larger value indicates a
                           lower preference). To do so, run the ip route-static dest-ip-address mask out-interface
                           preference preference-value command.

         Step 9 (Optional) Configure physical-layer fault notification.
                    After physical-layer fault notification is enabled, the AC interface can quickly
                    detect physical layer faults and trigger fast switchover of upper-layer applications.
                    1.   Enter the AC interface view.
                         interface interface-type interface-number
                    2.   Switch the interface working mode to Layer 3.
                         undo portswitch
                         Determine whether to perform this step based on the current interface
                         working mode.
                    3.   Enable physical-layer fault notification.
                         mpls l2vpn trigger if-down
                         By default, physical-layer fault notification is disabled.
                    4.   Return to the system view.
                         quit

        Step 10 (Optional) Configure a switchback policy.
                    1.   Enter the AC interface view.
                         interface interface-type interface-number
                    2.   Switch the interface working mode to Layer 3.
                         undo portswitch
                         Determine whether to perform this step based on the current interface
                         working mode.
                    3.   Configure a switchback policy.
                         mpls l2vpn reroute { { delay delayTime | immediately } [ resume resumeTime ] | never }
                         By default, delayed switchback is used in FRR or master/slave mode.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      515
VPN Configuration
VPN Configuration                                                                    5 VPWS Configuration


