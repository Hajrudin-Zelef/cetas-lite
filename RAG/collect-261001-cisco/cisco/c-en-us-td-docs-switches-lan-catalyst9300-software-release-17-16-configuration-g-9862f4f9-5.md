---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9-5
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9.md
source_anchor: ""
source_lines: [365, 512]
sha256: f9b39c858bcf37dc075c8ee911d3271c45a397fcffec14c3da8ea67203c5b4f4
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9

(Optional) If you configure the port as a static-access port, assign it to only one VLAN. The range is 1 to 4094.
Step 6
channel-groupchannel-group-numbermode {auto [non-silent] | desirable [non-silent ] | on } | { active | passive}
Example:
Device(config-if)# channel-group 5 mode auto
Assigns the port to a channel group, and specifies the PAgP or the LACP mode.
For mode, select one of these keywords:
auto— Enables PAgP only if a PAgP device is detected. It places the port into a passive negotiating state, in which the port responds
to PAgP packets it receives but does not start PAgP packet negotiation.
desirable— Unconditionally enables PAgP. It places the port into an active negotiating state, in which the port starts negotiations with
other ports by sending PAgP packets.
on— Forces the port to channel without PAgP or LACP. In the on mode, an EtherChannel exists only when a port group in the on mode is connected to another port group in the on mode.
non-silent— (Optional) If your device is connected to a partner that is PAgP-capable, configures the device port for nonsilent operation
when the port is in the auto or desirable mode. If you do not specify non-silent, silent is assumed. The silent setting is for connections to file servers or packet analyzers. This setting allows PAgP to
operate, to attach the port to a channel group, and to use the port for transmission.
active—Enables LACP only if a LACP device is detected. It places the port into an active negotiating state in which the port starts
negotiations with other ports by sending LACP packets.
passive— Enables LACP on the port and places it into a passive negotiating state in which the port responds to LACP packets that it
receives, but does not start LACP packet negotiation.
Step 7
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
Configuring Layer 3 EtherChannels
Follow these steps to assign an Ethernet port to a Layer 3 EtherChannel. This procedure is required.
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
interfaceinterface-id
Example:
Device(config)# interface gigabitethernet 1/0/2
Specifies a physical port, and enters interface configuration mode.
Valid interfaces include physical ports.
For a PAgP EtherChannel, you can configure up to eight ports of the same type and speed for the same group.
For a LACP EtherChannel, you can configure up to 16 Ethernet ports of the same type. Up to eight ports can be active, and
up to eight ports can be in standby mode.
Step 4
no ip address
Example:
Device(config-if)# no ip address
Ensures that there is no IP address assigned to the physical port.
Step 5
noswitchport
Example:
Device(config-if)# no switchport
Puts the port into Layer 3 mode.
Step 6
channel-groupchannel-group-numbermode{ auto [ non-silent] | desirable [ non-silent] | on} | { active | passive}
Example:
Device(config-if)# channel-group 5 mode auto
Assigns the port to a channel group, and specifies the PAgP or the LACP mode.
For mode, select one of these keywords:
auto—Enables PAgP only if a PAgP device is detected. It places the port into a passive negotiating state, in which the port responds
to PAgP packets it receives but does not start PAgP packet negotiation. This keyword is not supported when EtherChannel members
are from different switches in the switch stack.
desirable—Unconditionally enables PAgP. It places the port into an active negotiating state, in which the port starts negotiations
with other ports by sending PAgP packets. This keyword is not supported when EtherChannel members are from different switches
in the switch stack.
on—Forces the port to channel without PAgP or LACP. In the on mode, an EtherChannel exists only when a port group in the on mode is connected to another port group in the on mode.
non-silent—(Optional) If your device is connected to a partner that is PAgP capable, configures the device port for nonsilent operation
when the port is in the auto or desirable mode. If you do not specify non-silent, silent is assumed. The silent setting is for connections to file servers or packet analyzers. This setting allows PAgP to
operate, to attach the port to a channel group, and to use the port for transmission.
active—Enables LACP only if a LACP device is detected. It places the port into an active negotiating state in which the port starts
negotiations with other ports by sending LACP packets.
passive— Enables LACP on the port and places it into a passive negotiating state in which the port responds to LACP packets that it
receives, but does not start LACP packet negotiation.
Configures an EtherChannel extended load-balancing method.
The default is src-mac.
Select one of these load-distribution methods:
dst-ip—Specifies destination-host IP address.
dst-mac—Specifies the destination-host MAC address of the incoming packet.
dst-port—Specifies the destination TCP/UDP port.
ipv6-label—Specifies the IPv6 flow label.
l3-proto—Specifies the Layer 3 protocol.
src-ip—Specifies the source host IP address.
src-mac—Specifies the source MAC address of the incoming packet.
src-port—Specifies the source TCP/UDP port.
Step 4
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
(Optional) Configuring the Port Aggregation Protocol Learn Method and Priority
To configure the PAgP learn method and priority, perform this procedure:
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
interfaceinterface-id
Example:
Device(config)# interface gigabitethernet 1/0/2
Specifies the port for transmission, and enters interface configuration mode.
Step 4
pagp learn-methodphysical-port
Example:
Device(config-if)# pagp learn-method physical port
Selects the PAgP learning method.
By default, aggregation-port learning is selected, which means the device sends packets to the source by using any of the ports in the EtherChannel. With aggregate-port
learning, it is not important on which physical port the packet arrives.
Selects physical-port to connect with another device that is a physical learner.
Make sure to configure the port-channel load-balance global configuration command to src-mac.
The learning method must be configured the same at both ends of the link.
Step 5
pagp port-prioritypriority
Example:
Device(config-if)# pagp port-priority 200
Assigns a priority so that the selected port is chosen for packet transmission.
For priority, the range is 0 to 255. The default is 128. The higher the priority, the more likely that the port will be used for PAgP
transmission.
Step 6
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
Configuring Link Aggregation Control Protocol Hot-Standby Ports
When LACP is enabled, the software, by default, tries to configure the maximum number of LACP-compatible ports in a channel,
up to a maximum of 16 ports. Only eight LACP links can be active at one time; the remaining eight links are placed in hot-standby
mode. If one of the active links becomes inactive, a link that is in the hot-standby mode becomes active in its place.
You can override the default behavior by specifying the maximum number of active ports in a channel, in which case, the remaining
ports become hot-standby ports. For example, if you specify a maximum of five ports in a channel, up to 11 ports become hot-standby
ports.
If you configure more than eight links for an EtherChannel group, the software automatically decides which of the hot-standby
ports to make active based on the LACP priority. To every link between systems that operate LACP, the software assigns a unique
priority that is made up of these elements (in priority order):
