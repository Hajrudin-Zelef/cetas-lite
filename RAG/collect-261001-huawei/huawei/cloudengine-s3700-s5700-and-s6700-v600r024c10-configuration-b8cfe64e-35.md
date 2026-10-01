---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-35
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "latency"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [4835, 5000]
sha256: ed4c70ed82766887e591a28f140d7d5a371e39b8dc652db4b0603c650377bb58
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

QoS Configuration
QoS Configuration                                                        8 Priority Mapping Configuration


8.5.2 Configuring a DiffServ Domain
Context
                    When the device is used as an edge node connecting the DiffServ domain and
                    another network, you need to configure mappings between internal and external
                    priorities.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a DiffServ domain and enter the DiffServ domain view, or enter the view of
                an existing DiffServ domain.
                    diffserv domain ds-domain-name

                    The domain default defines default mappings between external priorities and
                    internal or drop priorities. You can modify the mappings defined in the domain
                    default, but you cannot delete this domain.
         Step 3 Run the following commands as required.
                     Operation                                  Command

                     Map the 802.1p value of an incoming        8021p-inbound 8021p-value phb
                     VLAN packet to an internal priority        service-class [ color ]
                     and color the packet.

                     Map the internal priority or drop          8021p-outbound service-class color
                     priority to the 802.1p value of VLAN       map 8021p-value
                     packets in the outbound direction.

                     Map DSCP values of incoming IP             ip-dscp-inbound dscp-value phb
                     packets to internal priorities and color   service-class [ color ]
                     the packets.

                     Map the internal priority or drop          ip-dscp-outbound service-class color
                     priority to the DSCP value of IP           map dscp-value
                     packets in the outbound direction.




                    For details about the default priority mappings, see 8.4 Default Settings for
                    Priority Mapping.

                    ----End

8.5.3 Applying the DiffServ Domain
Context
                    You can bind a DiffServ domain to an inbound or outbound interface of packets so
                    that the device can implement mapping between external priorities and internal
                    priorities/colors according to the mappings defined in the DiffServ domain.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            86
QoS Configuration
QoS Configuration                                                         8 Priority Mapping Configuration


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the interface view.
                    interface interface-type interface-number

         Step 3 Apply a DiffServ domain.
                    trust upstream { ds-domain-name | none }

                    By default, the domain default is applied to an interface.
                    If trust upstream none is configured on an interface, the device performs no
                    priority mapping on incoming packets.

                          NOTE

                        The trust upstream none command cannot be configured together with the qos phb
                        marking 8021p disable and qos phb marking dscp enable commands on the device.

         Step 4 (Optional) Enable the mapping between PHBs and DSCP values for outgoing
                packets.
                    qos phb marking dscp enable

         Step 5 (Optional) Disable the mapping between PHBs and 802.1p values for outgoing
                packets.
                    qos phb marking 8021p disable

                    ----End

8.5.4 Verifying the Configuration
Procedure
                    ●    Run the display diffserv domain [ brief | name ds-domain-name ] command
                         to check the DiffServ domain configuration.
                    ●    Run the display qos queue statistics { interface interface-name | interface
                         interface-type interface-number | slot slotid } command to check queue-
                         based traffic statistics.
                    ●    Run the display qos configuration interface [ { interface-type interface-
                         number | interface-name } ] command to check all QoS configurations on an
                         interface.
                    ----End

8.5.5 Example for Configuring Priority Mapping
Networking Requirements
                    In Figure 8-4, Host1 and Host2 are connected to DeviceB and DeviceC through
                    DeviceD, and then access the network through DeviceA.
                    The 802.1p values of packets sent from Host1 and Host2 are both 0. To ensure
                    normal running of low-latency services on Host1, the customer wants to configure
                    the CoS of Host1 to be higher than that of Host2, offering differentiated services
                    for Host1 and Host2.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                             87
QoS Configuration
QoS Configuration                                                                8 Priority Mapping Configuration


                    Figure 8-4 Network diagram of priority mapping
                          NOTE

                        In this example, interface 1, interface 2, interface 3, and interface 4 represent 10GE 1/0/1,
                        10GE 1/0/2, 10GE 1/0/3, and 10GE 1/0/4, respectively.




Procedure
         Step 1 Create VLANs and configure interfaces so that DeviceD can communicate with
                Host1, Host2, DeviceB, and DeviceC.

                    # Create VLAN 100 and VLAN 200.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceD
                    [DeviceD] vlan batch 100 200

                    # Set the link type of 10GE 1/0/3 and 10GE 1/0/4 on DeviceD to trunk, and add
                    10GE 1/0/1 to VLAN 100, 10GE 1/0/2 to VLAN 200, and 10GE 1/0/3 and 10GE
                    1/0/4 to both VLAN 100 and VLAN 200.
                    [DeviceD] interface 10ge 1/0/1
                    [DeviceD-10GE1/0/1] portswitch
                    [DeviceD-10GE1/0/1] port link-type access
                    [DeviceD-10GE1/0/1] port default vlan 100
                    [DeviceD-10GE1/0/1] quit
                    [DeviceD] interface 10ge 1/0/2
                    [DeviceD-10GE1/0/2] portswitch
                    [DeviceD-10GE1/0/2] port link-type access
                    [DeviceD-10GE1/0/2] port default vlan 200
                    [DeviceD-10GE1/0/2] quit
                    [DeviceD] interface 10ge 1/0/3
                    [DeviceD-10GE1/0/3] portswitch
                    [DeviceD-10GE1/0/3] port link-type trunk
                    [DeviceD-10GE1/0/3] port trunk allow-pass vlan 100 200
                    [DeviceD-10GE1/0/3] quit
                    [DeviceD] interface 10ge 1/0/4
                    [DeviceD-10GE1/0/4] portswitch
                    [DeviceD-10GE1/0/4] port link-type trunk
                    [DeviceD-10GE1/0/4] port trunk allow-pass vlan 100 200
                    [DeviceD-10GE1/0/4] quit

         Step 2 Configure interfaces to trust outer 802.1p values in packets.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         88
QoS Configuration
QoS Configuration                                                          8 Priority Mapping Configuration


