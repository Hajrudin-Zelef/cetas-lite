---
id: collect-261001-huawei/huawei/knowledge-base-design-and-configure-lldp-configuguration-example-huawei-6c21c10f
title: "knowledge-base-design-and-configure-lldp-configuguration-example-huawei-6c21c10f"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/knowledge-base-design-and-configure-lldp-configuguration-example-huawei-6c21c10f.md
source_anchor: ""
source_lines: [1, 32]
sha256: 055b0d17ceaae4d7be00fbd9d1fe66acb798ec5265e0c358734040fb34277575
---

# knowledge-base-design-and-configure-lldp-configuguration-example-huawei-6c21c10f

Technology: Switching
Area: Neighbor Discovery
Vendor: Huawei
Software: eNSP
Platform: Quidway switches
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
Note: The default delay in sending LLDP packets is 2 seconds.
To configure the hold-multiplier of device stored on neignbors Run:
lldp message-transmission hold-multiplier hold
Note: The default hold time multiplier is 4
Config verification:
display lldp local[ interface interface-type interface-number ] command to view LLDP local information on a specified interface or all interfaces.display lldp neighbor[ interface interface-type interface-number ] command to view neighbor information in the system or on an interface.display lldp neighbor briefcommand to view brief information about neighbors.
In routine maintenance, you can run the following commands in any view to check the LLDP status.
[SwitchA] display lldp neighbor brief
Local Intf   Neighbor Dev             Neighbor Intf             Exptime(s)
GE1/0/1      SwitchB                  GE1/0/1                   101
[SwitchB] display lldp neighbor brief
Local Intf   Neighbor Dev             Neighbor Intf             Exptime(s)
GE1/0/1      SwitchA                  GE1/0/1                   101
