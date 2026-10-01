---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-37
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [5179, 5305]
sha256: 8813b7b2d79f1dfc08918c20674d99e502a74be61811c25cafcc08505edc4629
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    For the default mappings between internal priorities and queues, see 8.4 Default
                    Settings for Priority Mapping.

                    ----End


8.9 Configuring a Packet Priority
Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure the 802.1p value of packets.
                    set priority 8021p 8021p-value

                    By default, the 802.1p value is not configured for packets.
         Step 3 Configure the DSCP value of packets.
                    set priority dscp dscp-value

                    By default, the DSCP value is not configured for packets.

                    ----End


8.10 Troubleshooting Priority Mapping

8.10.1 Packets Enter Incorrect Queues
Fault Symptoms
                    Packets enter incorrect queues.

Possible Causes
                    Possible causes are as follows:
                    ●    Priority mappings configured in the DiffServ domain bound to the inbound
                         interface are incorrect.
                    ●    There are configurations affecting packet queuing on the inbound interface.
                    ●    There are configurations affecting packet queuing in the VLAN to which the
                         packets belong.
                    ●    There are configurations affecting packet queuing in the system.

Procedure
         Step 1 Check whether priority mappings are correct.
                    Run the display this command in the inbound interface view to check the
                    configuration of the trust upstream command. (If the trust upstream command
                    is not configured, the DiffServ domain default is applied to an interface.) Then

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                               92
QoS Configuration
QoS Configuration                                                        8 Priority Mapping Configuration


                    run the display diffserv domain [ name ds-domain-name | brief ] command to
                    check whether the priority mappings configured in the trusted DiffServ domain
                    are correct.
                    ●   If the priority mappings are incorrect, run the ip-dscp-inbound or 8021p-
                        inbound command to correctly configure priority mappings.
                    ●   If the priority mappings are correct, go to step 2.
         Step 2 Check whether any configurations are affecting packet queuing on the inbound
                interface.
                    The following configurations affect the queues that packets enter on the inbound
                    interface:
                    ●   If the packets match a traffic policy that is applied to the inbound direction
                        (using the traffic-policy command) and contains the action of remark local-
                        precedence, the device sends the packets to queues based on the re-marked
                        internal priorities.
                    ●   If the trust upstream none command is configured, the device does not
                        perform priority mapping for incoming packets on the interface. Instead, the
                        device places packets into queues based on interface priorities.
                    Check whether any of the preceding configurations exist on the inbound interface.
                    ●   If such configurations are found, delete or modify them.
                    ●   If none of the configurations is found, go to step 3.
         Step 3 Check whether any configurations are affecting packet queuing in the VLAN to
                which the packets belong.
                    The following configurations affect packet queuing in a VLAN:
                    ●   If the packets match a traffic policy that is applied to the inbound direction
                        (using the traffic-policy command) and contains the action of remark local-
                        precedence, the device sends the packets to queues based on the re-marked
                        internal priorities.
                    ●   If the packets match a traffic policy that is applied to the inbound direction
                        (using the traffic-policy command) and contains the action of remark
                        8021p, the device maps the re-marked priorities of packets to internal
                        priorities and sends the packets to queues based on the mapped internal
                        priorities.
                    Check whether any of the preceding configurations exist in the VLAN.
                    ●   If such configurations are found, delete or modify them.
                    ●   If none of the configurations is found, go to step 4.
         Step 4 Check whether any configurations are affecting packet queuing in the system.
                    The following configurations affect packet queuing in the system:
                    ●   If the qos local-precedence-queue-map command is configured, the device
                        sends packets to queues based on the mapping between internal priorities
                        and queues specified by this command.
                    ●   If the packets match a global policy that is applied to the inbound direction
                        (using the traffic-policy global command) and contains the action of remark
                        local-precedence, the device sends packets to queues based on the re-
                        marked internal priorities.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                93
QoS Configuration
QoS Configuration                                                          8 Priority Mapping Configuration


                    ●   If the packets match a global policy that is applied to the inbound direction
                        (using the traffic-policy global command) and contains the action of remark
                        8021p, the device maps the re-marked priorities of packets to internal
                        priorities and sends the packets to queues based on the mapped internal
                        priorities.
                    Run the display current-configuration command to check for the preceding
                    configurations in the system. If such configurations are found, delete or modify
                    them.

                    ----End

8.10.2 Priority Mapping Results Are Incorrect
Fault Symptoms
                    Priority mapping results are incorrect.

Possible Causes
                    Possible causes are as follows:
                    ●   On the outbound interface, packets do not enter queues mapped to external
                        priorities.
                    ●   The priority types trusted by the inbound and outbound interfaces are
                        incorrect.
                    ●   The priority mappings configured in the DiffServ domains bound to the
                        inbound and outbound interfaces are incorrect.
                    ●   There are configurations affecting priority mapping on the inbound and
                        outbound interfaces.

