---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-36
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [5001, 5178]
sha256: 3972567a16818f01fc106a2ad57bdded3be5df00cb548d98161dc38c3c15a0e1
---

                    # Configure 10GE 1/0/1, 10GE 1/0/2, 10GE 1/0/3, and 10GE 1/0/4 to trust outer
                    802.1p values in packets. By default, a Layer 2 interface trusts outer 802.1p values.
                    In this case, skip this step.
                    [DeviceD] interface 10ge 1/0/1
                    [DeviceD-10GE1/0/1] trust 8021p outer
                    [DeviceD-10GE1/0/1] quit
                    [DeviceD] interface 10ge 1/0/2
                    [DeviceD-10GE1/0/2] trust 8021p outer
                    [DeviceD-10GE1/0/2] quit
                    [DeviceD] interface 10ge 1/0/3
                    [DeviceD-10GE1/0/3] trust 8021p outer
                    [DeviceD-10GE1/0/3] quit
                    [DeviceD] interface 10ge 1/0/4
                    [DeviceD-10GE1/0/4] trust 8021p outer
                    [DeviceD-10GE1/0/4] quit

         Step 3 Create DiffServ domains, configure priority mapping in the DiffServ domains, and
                bind the DiffServ domains to interfaces.
                    # Create DiffServ domains ds1 and ds2 on DeviceD and map the 802.1p values of
                    packets sent from Host1 and Host2 to different internal priorities.
                    [DeviceD] diffserv domain ds1
                    [DeviceD-dsdomain-ds1] 8021p-inbound 0 phb af4 green
                    [DeviceD-dsdomain-ds1] quit
                    [DeviceD] diffserv domain ds2
                    [DeviceD-dsdomain-ds2] 8021p-inbound 0 phb af2 green
                    [DeviceD-dsdomain-ds2] quit

                    # Bind DiffServ domains ds1 and ds2 to 10GE 1/0/1 and 10GE 1/0/2, respectively.
                    [DeviceD] interface 10ge 1/0/1
                    [DeviceD-10GE1/0/1] trust upstream ds1
                    [DeviceD-10GE1/0/1] quit
                    [DeviceD] interface 10ge 1/0/2
                    [DeviceD-10GE1/0/2] trust upstream ds2
                    [DeviceD-10GE1/0/2] quit

                    ----End

Verifying the Configuration
                    # Display queue statistics on the outbound interface. If there are packet statistics
                    in queue 2 corresponding to the internal priority AF2, and in queue 4
                    corresponding to the internal priority AF4, priority mapping is configured
                    successfully.

Configuration Scripts
                    DeviceD
                    #
                    sysname DeviceD
                    #
                    vlan batch 100 200
                    #
                    diffserv domain ds1
                     8021p-inbound 0 phb af4 green
                    #
                    diffserv domain ds2
                     8021p-inbound 0 phb af2 green
                    #
                    interface 10GE1/0/1
                     port link-type access


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                             89
QoS Configuration
QoS Configuration                                                                8 Priority Mapping Configuration

                     port default vlan 100
                     trust upstream ds1
                    #
                    interface 10GE1/0/2
                     port link-type access
                     port default vlan 200
                     trust upstream ds2
                    #
                    interface 10GE1/0/3
                     port link-type trunk
                     port trunk allow-pass vlan 100 200
                    #
                    interface 10GE1/0/4
                     port link-type trunk
                     port trunk allow-pass vlan 100 200
                    #
                    return



8.6 Configuring an Interface Priority
Context
                    The device provides differentiated services based on the interface priority in the
                    following scenarios:
                    ●     The trust upstream none command is configured on an interface.
                    ●     An interface receives untagged packets.
                                NOTE

                               ● The interface priority cannot be specified for Eth-Trunk member interfaces.
                               ● If no DiffServ domain is applied to an inbound interface and the interface priority
                                 is configured, the device adds a VLAN tag to received packets, uses the interface
                                 priority as the external priority, and sends packets to queues based on the interface
                                 priority.
                               ● When an Ethernet interface works in Layer 3 mode, the interface priority is 0 and
                                 cannot be configured.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the interface view.
                    interface interface-type interface-number

         Step 3 Configure an interface priority.
                    port priority priority-value

                    ----End




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                      90
QoS Configuration
QoS Configuration                                                           8 Priority Mapping Configuration




8.7 Configuring the Internal Priority of Protocol
Packets Sent by the Local Device
Context
                    A device processes packets based on their internal priorities. You can change the
                    internal priority of protocol packets sent by the local device to preferentially
                    process packets with a high internal priority on the local device.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure the internal priority of protocol packets sent by the local device.
                    set priority traffic-class traffic-class-value

                    By default, the internal priority is not configured for protocol packets sent by the
                    local device, and the local device sends all protocol packets based on their original
                    internal priorities.

                    ----End


8.8 Configuring the Mapping Between Internal
Priorities and Queues
Context
                    A device schedules packets based on queues. After external priorities are mapped
                    to internal priorities (PHBs), the device sends packets to queues based on the
                    mapping between internal priorities and queue indexes. By doing this, the device
                    provides differentiated services.
                    By default, internal priorities and port queues are mapped in one-to-one mode. In
                    practice, packets of two or more internal priorities are scheduled in the same
                    queue to conserve buffer resources on the device. The device sends packets to
                    different port queues based on the internal priority, and performs traffic shaping,
                    congestion avoidance, and queue scheduling for the queues.

                           NOTE

                         Only the S6730E-H-V2, S6730-H-V2, S6750E-S, S6750-S, S5755E-H, S5755-H, and S5732-H-
                         V2 series support this function.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure the mappings between internal priorities and queues.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                             91
QoS Configuration
QoS Configuration                                                              8 Priority Mapping Configuration

                    qos local-precedence-queue-map service-class queue-index

