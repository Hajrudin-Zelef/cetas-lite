---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9.md
source_anchor: ""
source_lines: [69, 144]
sha256: cb62d4981462ac5656980bb4d680afe7d5463c8743b8d2a907b364ea6dc902e7
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9

also must set the load-distribution method to source-based distribution, so that any given source MAC address is always sent
on the same physical port.
You also can configure a single port within the group for all transmissions and use other ports for hot-standby. The unused
ports in the group can be swapped into operation in just a few seconds if the selected single port loses hardware-signal detection.
You can configure which port is always selected for packet transmission by changing its priority with the pagp port-priority interface configuration command. The higher the priority, the more likely that the port will be selected.
Note
The device supports address learning only on aggregate ports even though the physical-port keyword is provided in the CLI. The pagp learn-method command and the pagp port-priority command have no effect on the device hardware, but they are required for PAgP interoperability with devices that only support
address learning by physical ports, such as the Catalyst 1900 switch.
When the link partner of the device is a physical learner, we recommend that you configure the device as a physical-port learner
by using the pagp learn-method physical-port interface configuration command. Set the load-distribution method based on the source MAC address by using the port-channel load-balance src-mac global configuration command. The device then sends packets to the physical learner using the same port in the EtherChannel
from which it learned the source address. Only use the pagp learn-method command in this situation.
Port Aggregation Protocol Interaction with Other Features
The Dynamic Trunking Protocol (DTP) and the Cisco Discovery Protocol (CDP) send and receive packets over the physical ports
in the EtherChannel. Trunk ports send and receive PAgP protocol data units (PDUs) on the lowest numbered VLAN.
In Layer 2 EtherChannels, the first port in the channel that comes up provides its MAC address to the EtherChannel. If this
port is removed from the bundle, one of the remaining ports in the bundle provides its MAC address to the EtherChannel. For Layer 3 EtherChannels, the MAC address is allocated by the active device as soon as the interface is created (through
the interface port-channel global configuration command).
PAgP sends and receives PAgP PDUs only from ports that are up and have PAgP enabled for the auto or desirable mode.
Link Aggregation Control Protocol
The LACP is defined in IEEE 802.3ad and enables Cisco devices to manage Ethernet channels between devices that conform to
the IEEE 802.3ad protocol. LACP facilitates the automatic creation of EtherChannels by exchanging LACP packets between Ethernet
ports.
By using LACP, the switch or switch stack learns the identity of partners capable of supporting LACP and the capabilities
of each port. It then dynamically groups similarly configured ports into a single logical link (channel or aggregate port).
Similarly configured ports are grouped based on hardware, administrative, and port parameter constraints. For example, LACP
groups the ports with the same speed, duplex mode, native VLAN, VLAN range, and trunking status and type. After grouping the
links into an EtherChannel, LACP adds the group to the spanning tree as a single device port.
The independent mode behavior of ports in a port channel is changed. With CSCtn96950, by default, standalone mode is enabled.
When no response is received from an LACP peer, ports in the port channel are moved to suspended state.
LACP modes specify whether a port can send LACP packets or only receive LACP packets.
Table 2. EtherChannel LACP Modes
Mode
Description
active
Places a port into an active negotiating state in which the port starts negotiations with other ports by sending LACP packets.
passive
Places a port into a passive negotiating state in which the port responds to LACP packets that it receives, but does not start
LACP packet negotiation. This setting minimizes the transmission of LACP packets.
Both the active and passive LACP modes enable ports to negotiate with partner ports to an EtherChannel based on criteria such as port speed, and for Layer 2
EtherChannels, based on trunk state and VLAN numbers.
Ports can form an EtherChannel when they are in different LACP modes as long as the modes are compatible. For example:
A port in the active mode can form an EtherChannel with another port that is in the active or passive mode.
A port in the passive mode cannot form an EtherChannel with another port that is also in the passive mode because neither port starts LACP negotiation.
Link Aggregation Control Protocol Standalone Mode on Ethernet Channel
When one end of an EtherChannel has more members than the other, the unmatched ports enter the standalone state. The standalone
mode is also called the independent mode. In the standalone mode the port is not bundled in an EtherChannel. The port functions
as a standalone data port and it can send and receive BPDUs and data traffic.
In a topology that is not protected from Layer 2 loops by the spanning tree protocol
(STP), a port in the standalone state can cause significant network errors. You can
enter the port-channel standalone-disable command in the
interface configuration mode to put ports into the suspended state instead of the
standalone state.
The standalone mode is particularly relevant when a port (A) in a Layer 2 LACP EtherChannel is connected to an unresponsive
port (B) on the peer. When LACP standalone is disabled on the EtherChannel, all traffic arriving on A is blocked (the default
behavior on a switch). In some scenarios, you might want to allow management traffic on such ports. You can do this by enabling
LACP standalone (or independent) mode. To enable the standalone mode on a Layer 2 LACP Etherchannel, use the no port-channel standalone disable command in the interface configuration mode. To disable the Standalone mode and revert to the default use the port-channel standalone disable command in the interface configuration mode.
Note
LACP standalone mode is disabled by default.
Starting with the Cisco IOS XE Dublin 17.10.1 release, you can configure the LACP standalone mode on a Layer 3 EtherChannel. To configure the standalone mode use the no port-channel standalone disable command in the interface configuration mode. To disable the Standalone mode and revert to the default use the port-channel standalone disable command in the interface configuration mode.
Link Aggregation Control Protocol and Link Redundancy
LACP port-channel operation, bandwidth availability, and link redundancy can be further refined with the LACP port-channel
min-links and the LACP max-bundle features.
The LACP port-channel min-links feature:
Configures the minimum number of ports that must be linked up and bundled in the LACP port channel.
Prevents a low-bandwidth LACP port channel from becoming active.
Causes an LACP port channel to become inactive if there are too few active members ports to supply the required minimum bandwidth.
The LACP max-bundle feature:
Defines an upper limit on the number of bundled ports in an LACP port channel.
Allows hot-standby ports with fewer bundled ports. For example, in an LACP port channel with five ports, you can specify
a max-bundle of three, and the two remaining ports are designated as hot-standby ports.
Link Aggregation Control Protocol Interaction with Other Features
The DTP and the CDP send and receive packets over the physical ports in the EtherChannel. Trunk ports send and receive LACP
PDUs on the lowest numbered VLAN.
In Layer 2 EtherChannels, the first port in the channel that comes up provides its MAC address to the EtherChannel. If this
port is removed from the bundle, one of the remaining ports in the bundle provides its MAC address to the EtherChannel. For Layer 3 EtherChannels, the MAC address is allocated by the active device as soon as the interface is created through the
interface port-channel global configuration command.
