---
id: collect-261001-cisco/cisco/enterprise-en-huawei-switch-basic-configurations-with-example-5-examples-thread-1829e1ea
title: "enterprise-en-huawei-switch-basic-configurations-with-example-5-examples-thread--1829e1ea"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet", "voice"]
source: docs/RAG/collect-261001-cisco/enterprise-en-huawei-switch-basic-configurations-with-example-5-examples-thread--1829e1ea.md
source_anchor: ""
source_lines: [1, 59]
sha256: 41b8902380b365a0977d0a191d43ccc263218e093759a16fac7afec09c5da512
---

# enterprise-en-huawei-switch-basic-configurations-with-example-5-examples-thread--1829e1ea

Access port config:
An access port belongs to and carries the traffic of only one VLAN within 2-layer switch. Traffic is both received and sent in native formats with no VLAN tagging whatsoever. Anything arriving on an access port is simply assumed to belong to the VLAN assigned to the port.
To configure port in vlan 10:
<HUAWEI> system-view[HUAWEI] interface GigabitEthernet 0/0/15[HUAWEI-GigabitEthernet 0/0/15]port link-type access[HUAWEI-GigabitEthernet 0/0/15]port default vlan 10[HUAWEI-GigabitEthernet 0/0/15]quit
Trunk port config:
A trunk interface often connects to a switch, router, AP, or voice terminal that can receive and send tagged and untagged frames simultaneously. It allows tagged frames from multiple VLANs and untagged frames from only one VLAN called native vlan.
To configure ethernet port as a trunk port and allow to pass all vlans:
<HUAWEI> system-view
[HUAWEI] interface GigabitEthernet 0/0/13
[HUAWEI-GigabitEthernet 0/0/13]port link-type trunk
[HUAWEI-GigabitEthernet 0/0/13]port trunk allow-pass vlan all
[HUAWEI-GigabitEthernet 0/0/13]quit
How to configure VLAN on Huawei:
Virtual Local Area Network (VLAN) technology divides a physical LAN into multiple broadcast domains, each of which is called a VLAN. Hosts within a VLAN can communicate with each other but cannot communicate directly with hosts in other VLANs. Consequently, broadcast packets are confined to within a single VLAN.
To create vlan:
<Huawei> system-view[Huawei] sysname RouterA[RouterA] vlan batch 2 3
To assign interface to specific vlan:
[RouterA] interface ethernet 2/0/1[RouterA-Ethernet2/0/1] port link-type access[RouterA-Ethernet2/0/1] port default vlan 2[RouterA-Ethernet2/0/1] quit
LACP configuration:
Static LACP mode:
Devices that are directly connected through an Eth-Trunk can support LACP, by running the mode lacp-static command you can configure the Eth-Trunk to work in static LACP mode. This mode can implement both load balancing and redundancy.
<HUAWEI> system-view[HUAWEI] interface Eth-Trunk 1[HUAWEI- interface Eth-Trunk 1]mode lacp-static[HUAWEI- interface Eth-Trunk 1]quit
LACP Dynamic mode:
When a server is connected to a file server over the device that has an Eth-Trunk configured, one is able to run the mode lacp-dynamic command on the device to configure its Eth-Trunk to work in dynamic mode.
<HUAWEI> system-view[HUAWEI] interface Eth-Trunk 1[HUAWEI- interface Eth-Trunk 1]mode lacp-dynamic[HUAWEI- interface Eth-Trunk 1]quit
LLDP configuration example:
The Link Layer Discovery Protocol (LLDP) is a standard Layer 2 topology discovery protocol defined in IEEE 802.1ab. LLDP allows a device to send local management information such as management IP address, device ID, and port ID to neighbors. LLDP provides a standard link-layer discovery method. Layer 2 information obtained through LLDP allows the NMS (Network Management System) to detect the topology of neighboring devices. Details, such as device capabilities or device identity can be advertised using this protocol.
Let’s take a look at LLDP configuration example for Huawei.
Enabling global LLDP on Switch:
<Switch>system-view
 [Switch]lldp enable
Interval for sending LLDP packets- When the LLDP status of a device keeps unchanged, the device sends LLDP packets to its neighbors at a certain interval. The value of interval ranges from 5 to 32768. Increasing the value of interval is not restricted by the value of delay.
Delay in sending LLDP packets- If the device status changes frequently, a delay is required before the device sends an LLDP packet to its neighbors. The value of delay ranges from 1 to 8192. Decreasing the value of delay is not restricted by the value of interval.
Hold time multiplier of device information on neighbors- The hold time multiplier is used to calculate the Time to Live (TTL), which determines how long device information can be saved on the neighbors. You can specify the hold time of device information on the neighbors. After receiving an LLDP packet, a neighbor updates the aging time of the device information from the sender based on the TTL.
The storage time calculation formula is: TTL = Min (65535, (interval x hold)).
To configure the interval for sending LLDP packets Run:
lldp message-transmission interval interval
Note: The default interval for sending LLDP packets is 30 seconds.
To delay in sending LLDP packets in Run:
lldp message-transmission delay delay
To delay in sending LLDP packets in Run:
lldp message-transmission delay delay
Note: The default delay in sending LLDP packets is 2 seconds.
To configure the hold-multiplier of device stored on neighbors Run:
lldp message-transmission hold-multiplier hold
Note: The default hold time multiplier is 4
Config verification:
- Run the display lldp local[ interface interface-type interface-number ] command to view LLDP local information on a specified interface or all interfaces.
- Run the display lldp neighbor[ interface interface-type interface-number ] command to view neighbor information in the system or on an interface.
- Run the display lldp neighbor brief command to view brief information about neighbors.
In routine maintenance, you can run the following commands in any view to check the LLDP status.
- Run the display lldp statistics[ interface interface-type interface-number ] command to view statistics about sent and received LLDP packets in the system or on an interface.
- Run the display cdp statistics[ interface interface-type interface-number ] command to view statistics about sent and received CDP packets in the system or on an interface.
[SwitchA] display lldp neighbor brief
 Local Intf   Neighbor Dev             Neighbor Intf             Exptime(s)
 GE1/0/1      SwitchB                  GE1/0/1                   101
 [SwitchB] display lldp neighbor brief
 Local Intf   Neighbor Dev             Neighbor Intf             Exptime(s)
 GE1/0/1      SwitchA                  GE1/0/1                   101
