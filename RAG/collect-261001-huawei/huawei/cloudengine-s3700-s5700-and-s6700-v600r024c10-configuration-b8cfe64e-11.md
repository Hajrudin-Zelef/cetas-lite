---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-11
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [969, 1119]
sha256: ed51614053288ee5c7352a9ae6e27a02d09cbb36d2027942a5c40825e71e97d1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                         Inner information in   if-match srv6 transit acl { acl-number | acl-name }
                         SRv6 packets to be
                         matched against
                         ACL rules

                         Inner information in   if-match srv6 transit ipv6 acl { acl-number | acl-
                         SRv6 packets to be     name } [ loose-mode | strict-mode ]
                         matched against
                         ACL6 rules




                        Only the S6730E-H-V2, S6730-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H,
                        S5755-H, and S5732-H-V2 series support the commands for configuring SRv6
                        packet matching rules.
                    ●   Application/Application group rule
                         Matching Rule          Command

                         Application            if-match application application-name
                                                NOTE
                                                 application-name specifies the name of an application. You
                                                 can run the display application command to view
                                                 application names.




                        Only the S6730E-H-V2, S6730-H-V2, S5755E-H, S5755-H, S5732-H-V2,
                        S6750E-S, S6750-S, and S5755-S series support the command for configuring
                        an application matching rule.
                    ●   Other rules
                         Matching Rule          Command

                         Inner information in   if-match gre [ inner-source-ip source-ip-address
                         GRE packets            [ mask source-masklen ] | inner-destination-ip
                                                dest-ip-address [ mask dest-masklen ] | inner-
                                                protocol protocol | inner-source-port source-port-
                                                begin | inner-destination-port dest-port-begin ] *
                                                This function is supported only on the S6750-H and
                                                S6780-H.

                         QoS local ID           if-match qos-local-id qos-local-id




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                                  16
QoS Configuration
QoS Configuration                                                                            3 MQC Configuration


                            Matching Rule               Command

                            EXP value of MPLS           if-match mpls-exp mpls-exp-value &<1-8>
                            packets                     Only the S6780-H, S6750-S, S6730-H-V2, S6750-H,
                                                        S5732-H-V2, S6750E-S, S6730E-H-V2, S5755E-H,
                                                        S5755-S and S5755-H series support this function.


         Step 4 Exit the traffic classifier view.
                    quit

         Step 5 (Optional) Enable the device to provide nonstop services upon modification of
                MQC-based traffic classification rules.
                    traffic-policy atomic-update-mode

                    If a traffic classification rule is modified after this command is run, the system
                    delivers the new rule and then deletes the old rule to ensure that traffic is not
                    interrupted.
                    With this function enabled, if a traffic policy that has been successfully applied
                    fails to be applied again due to its configuration change, the original traffic policy
                    still takes effect.
                    By default, this function is disabled.

                    ----End


3.5 Configuring a Traffic Behavior
Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a traffic behavior and enter the traffic behavior view, or enter the view of
                an existing traffic behavior.
                    traffic behavior behavior-name

         Step 3 Define actions in the traffic behavior. You can configure multiple non-conflicting
                actions in a traffic behavior.
                            NOTE

                           For details about precautions for each action, see the corresponding commands in the
                           Command Reference.
                           For S5735-L-V2, S5735-S-V2, S5735I-L-V2, S5735I-S-V2, S5735I-H-V2, S5735R-L-V2, S3710-
                           H, S5735R-S-V2, S5735E-L-V2, S5735E-S-V2 series, the device supports a maximum of 1024
                           traffic behaviors.
                           For the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-H-V2,
                           S5755E-H, S5755-S and S5755-H series, the device supports a maximum of 2048 traffic
                           behaviors.

                     Action                  Command

                     Packet filtering        See 4 Packet Filtering Configuration.


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                     17
QoS Configuration
QoS Configuration                                                                  3 MQC Configuration


                    Action               Command

                    Traffic statistics   See 5 Traffic Statistics Collection Configuration.
                    collection

                    MQC-based re-        See 6 Re-marking Configuration.
                    marking

                    MQC-based            See 7 Redirection Configuration.
                    redirection

                    MQC-based            See 9.5.1 Configuring MQC-based Traffic Policing (Level-1
                    traffic policing     CAR) in Traffic Policing, Traffic Shaping, and Interface-
                                         based Rate Limiting Configuration.

                    Hierarchical         car car-name share
                    traffic policing     See 9.5.2 Configuring Traffic Policing (Level-2 CAR) in
                                         Traffic Policing, Traffic Shaping, and Interface-based Rate
                                         Limiting Configuration.
                                         Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-
                                         V2, S6750E-S, S6730E-H-V2, S5755E-H, S5755-S and S5755-H
                                         series support this function.

                    MQC-based            mirroring observe-port observe-port-index
                    flow mirroring       mirroring observe-port group group-id
                                         For details, see "Configuring Flow Mirroring" in CLI
                                         Configuration Guide > System Monitoring Configuration >
                                         Mirroring Configuration.

                    Disabling MAC        mac-address learning disable
                    address learning     For details, see "Disabling MAC Address Learning" in CLI
                                         Configuration Guide > Ethernet Switching Configuration >
                                         MAC Configuration.

                    MQC-based            vlan-mapping { vlan vlan-id | inner-vlan inner-vlan-id }
                    VLAN mapping         For details, see "Configuring VLAN Mapping" in CLI
                                         Configuration Guide > Ethernet Switching Configuration >
                                         VLAN Configuration.

