---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59-5
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59.md
source_anchor: ""
source_lines: [416, 555]
sha256: 706f563da5a01022b1219f6ca6e583767c3abccdb693af10c00793d3fb4662b0
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59

contains the port channel and member ports. You can use the numbers from 1 to 4096 in each VDC to number the port channels.
All ports in one port channel must be in the same VDC. When you are using LACP, all possible 8 active ports and all possible
8 standby ports must be in the same VDC.
Note
You must configure load balancing using port channels in the default VDC. See the “Load Balancing Using Port Channels” section
for more information about load balancing.
High Availability
Port channels provide high availability by load balancing traffic across multiple ports. If a physical port fails, the port
channel is still operational if there is an active member in the port channel. You can bundle ports from different modules
and create a port channel that remains operational even if a module fails because the settings are common across the module.
Port channels support stateful and stateless restarts. A stateful restart occurs on a supervisor switchover. After the switchover,
the Cisco NX-OS software applies the runtime configuration after the switchover.
The port channel goes down if the operational ports fall below the configured minimum links number.
Note
See the Cisco Nexus 9000 Series NX-OS High Availability and Redundancy Guide for complete information about high-availability features.
Prerequisites for Port Channeling
Port channeling has the following prerequisites:
You must be logged onto the device.
All ports for a single port channel must be either Layer 2 or Layer 3 ports.
All ports for a single port channel must meet the compatibility requirements. See the “Compatibility Requirements” section
for more information about the compatibility requirements.
You must configure load balancing from the default VDC.
Guidelines and Limitations
Port channeling has the following configuration guidelines and limitations:
For scaled port-channel deployments on Cisco Nexus 9516 switch with Gen 1 line cards, you need to use the port-channel scale-fanout command followed by copy run start and reload commands.
show commands with the internal keyword are not supported.
The LACP port-channel minimum links and maxbundle feature is not supported for host interface port channels.
Enable LACP before you can use that feature.
You can configure multiple port channels on a device.
Do not put shared and dedicated ports into the same port channel. (See the “Configuring Basic Interface Parameters” chapter
for information about shared and dedicated ports.)
For Layer 2 port channels, ports with different STP port path costs can form a port channel if they are compatibly configured
with each other. See the “Compatibility Requirements” section for more information about the compatibility requirements.
When L2 ePBR is configured between the L3 port channel interface, port channel will not come up as the LACP packet drops at
the ePBR device.
In STP, the port-channel cost is based on the aggregated bandwidth of the port members.
After you configure a port channel, the configuration that you apply to the port channel interface affects the port channel
member ports. The configuration that you apply to the member ports affects only the member port where you apply the configuration.
LACP does not support half-duplex mode. Half-duplex ports in LACP port channels are put in the suspended state.
Do not configure ports that belong to a port channel group as private VLAN ports. While a port is part of the private VLAN
configuration, the port channel configuration becomes inactive.
Channel member ports cannot be a source or destination SPAN port.
Port-channels are not supported on generation 1 100G line cards (N9K-X9408PC-CFP2) or generic expansion modules (N9K-M4PC-CFP2).
Port-channels are supported on devices with generation 2 (and later) 100G interfaces.
Resilient hashing (port-channel load-balancing resiliency) and VXLAN configurations are not compatible with VTEPs using ALE
uplink ports.
Note
Resilient hashing is disabled by default.
The maximum number of subinterfaces for a satellite/FEX port is 63.
On a Cisco Nexus 92300YC switch, the first 24 ports that are part of the same quadrant. All the ports in the same quadrant
must have same speed. Having different speed on ports in a quadrant is not supported. Following are the first 24 ports on
the Cisco Nexus 92300YC switch that share same quadrant:
1,4,7,10
2,5,8,11
3,6,9,12
13,16,19,22
14,17,20,23
15,18,21,24
On a Cisco Nexus 9500 switch with a X96136YC-R line card, the ports 17–48 are part of the same quadrant. Ports in the same
quadrant must have same speed (1/10G or 25G) on all ports. Having different speed on ports in a quadrant is not supported.
If you set different speed in any of the ports in a quadrant, the ports go into error disable state. Interfaces in same quadrant
are:
17–20
21–24
25–28
29–32
33–36
37–40
41–44
45–48
Resilient hashing is supported on Cisco Nexus 9500 Series switches with N9K-X9636C-R, N9K-X9636Q-R, N9K-X9636C-RX, and N9K-X96136YC-R
line cards.
Port-channel symmetric hashing is supported on Cisco Nexus 9200, 9300-EX, 9300-FX/FX2, and 9300-GX platform switches and Cisco
Nexus 9500 platform switches with N9K-X9732C-EX, N9K-X9736C-EX, N9K-X9736C-FX, and N9K-X9732C-FX line cards.
ECMP symmetric hashing is supported on Cisco Nexus 9200, 9300-EX, and 9300-FX/FX2 platform switches and Cisco Nexus 9500 platform
switches with N9K-X9732C-EX, N9K-X9736C-EX, N9K-X9736C-FX, and N9K-X9732C-FX line cards.
GRE inner headers are supported on the following switches:
Cisco Nexus 9364C platform switches
Cisco Nexus 9336C-FX2, 9348GC-FXP, 93108TC-FX, 93180YC-FX, and 93240YC-FX2 platform switches
Cisco Nexus 9300-GX platform switches.
Cisco Nexus 9500 platform switches with N9K-X9736C-FX line cards
Beginning with Cisco NX-OS Release 9.3(6), Cisco Nexus 9300-FX2 platform switches support the coexistence of VXLAN and IP-in-IP
tunneling. For more information, including limitations, see the VXLAN and IP-in-IP Tunneling section in the Cisco Nexus 9000 Series NX-OS VXLAN Configuration Guide, Release 9.3(x).
For FEX interfaces using LACP, all DME oper/runtime properties for the FEX interfaces does not get updated. All runtime updates
for FEX ports happens from FEX LACP process context and are not communicated to the parent switch.This is a day-1 behaviour.
Beginning with Cisco NX-OS Release 10.3(1)F, the hashing based on src/dst ip and src/dst L4 port number is supported on Cisco Nexus 9808 platform switches.
Default
Settings
The following table lists the default
settings for port-channel parameters.
Table 4. Default Port-Channel Parameters
Parameters
Default
Port channel
Admin up
Load balancing method for Layer 3 interfaces
Source and destination IP address
Load balancing method for Layer 2 interfaces
Source and destination MAC address
Load balancing per module
Disabled
LACP
Disabled
Channel mode
on
LACP system priority
32768
LACP port priority
32768
Minimum links for LACP
1
Maxbundle
32
Minimum links for FEX fabric port channel
1
Configuring Port
Channels
Note
See the "Configuring Basic Interface Parameters” chapter for
information about configuring the maximum transmission unit (MTU) for the
port-channel interface. See the “Configuring Layer 3 Interfaces” chapter for
information about configuring IPv4 and IPv6 addresses on the port-channel
interface.
Note
If you are familiar with the Cisco IOS CLI, be aware that the Cisco
NX-OS commands for this feature might differ from the Cisco IOS commands that
you would use.
You can create a
port channel before you create a channel group. The software automatically
creates the associated channel group.
Note
When the port
channel is created before the channel group, the port channel should be
configured with all of the interface attributes that the member interfaces are
configured with. Use the
switchport mode trunk {allowed vlanvlan-id |
nativevlan-id}
command to configure the
members.
This is required
