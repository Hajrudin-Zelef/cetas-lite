---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9-4
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9.md
source_anchor: ""
source_lines: [220, 364]
sha256: 9ac72b855f10570a1fe918a2e28b243a36b894e8bd2ccfd459b127a7dba22d16
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9

With LACP, the system ID uses the stack MAC address from the active switch. When an active switch fails or leaves the stack
and the standby switch becomes the new active switch, the LACP system ID is unchanged. By default, the LACP configuration
is not affected after the active switch changes.
Default EtherChannel Configuration
The default EtherChannel configuration is described in this table.
Table 3. Default EtherChannel Configuration
Feature
Default Setting
Channel groups
None assigned.
Port-channel logical interface
None defined.
PAgP mode
No default.
PAgP learn method
Aggregate-port learning on all ports.
PAgP priority
128 on all ports.
LACP mode
No default.
LACP learn method
Aggregate-port learning on all ports.
LACP port priority
32768 on all ports.
LACP system priority
32768.
LACP system ID
LACP system priority and the switch or stack MAC address.
Load-balancing
Load distribution on the switch is based on the source-MAC address of the incoming packet.
The source-MAC address is src-mac.
EtherChannel Configuration Guidelines
If improperly configured, some EtherChannel ports are automatically disabled to avoid network loops and other problems. Follow
these guidelines to avoid configuration problems:
A maximum of 128 EtherChannels are supported on a switch or switch stack.
Configure all ports in an EtherChannel to operate at the same speeds and duplex modes.
Enable all ports in an EtherChannel. A port in an EtherChannel that is disabled by using the shutdown interface configuration command is treated as a link failure, and its traffic is transferred to one of the remaining ports
in the EtherChannel.
When a group is first created, all ports follow the parameters set for the first port to be added to the group. If you change
the configuration of one of these parameters, you must also make the changes to all ports in the group:
Allowed-VLAN list
Spanning-tree path cost for each VLAN
Spanning-tree port priority for each VLAN
Spanning-tree Port Fast setting
Do not configure a port to be a member of more than one EtherChannel group.
Do not configure an EtherChannel in both the PAgP and LACP modes. EtherChannel groups running PAgP and LACP can coexist on
the same switch or on different switches in the stack. Individual EtherChannel groups can run either PAgP or LACP, but they
cannot interoperate.
Do not configure a port that is an active or a not-yet-active member of an EtherChannel as an IEEE 802.1x port. If you try
to enable IEEE 802.1x on an EtherChannel port, an error message appears, and IEEE 802.1x is not enabled.
If EtherChannels are configured on device interfaces, remove the EtherChannel configuration from the interfaces before globally
enabling IEEE 802.1x on a device by using the dot1x system-auth-control global configuration command.
When configuring Layer 2 EtherChannels, follow these guidelines:
Assign all ports in the EtherChannel to the same VLAN, or configure them as trunks. Ports with different native VLANs cannot
form an EtherChannel.
An EtherChannel supports the same allowed range of VLANs on all the ports in a trunking Layer 2 EtherChannel. If the allowed
range of VLANs is not the same, the ports do not form an EtherChannel even when PAgP is set to the auto or desirable mode.
Ports with different spanning-tree path costs can form an EtherChannel if they are otherwise compatibly configured. Setting
different spanning-tree path costs does not, by itself, make ports incompatible for the formation of an EtherChannel.
Layer 3 EtherChannel Configuration Guidelines
For Layer 3 EtherChannels, assign the Layer 3 address to the port-channel logical interface, not to the physical ports in
the channel.
Auto-LAG
The auto-LAG feature provides the ability to auto create EtherChannels on ports that are connected to a switch. By default,
auto-LAG is disabled globally and is enabled on all port interfaces. The auto-LAG applies to a switch only when it is enabled
globally.
On enabling auto-LAG globally, the following scenarios are possible:
All port interfaces participate in creation of auto EtherChannels provided the partner port interfaces have EtherChannel
configured on them. For more information, see the "The supported auto-LAG configurations between the actor and partner devices" table below.
Ports that are already part of manual EtherChannels cannot participate in creation of auto EtherChannels.
When auto-LAG is disabled on a port interface that is already a part of an auto created EtherChannel, the port interface unbundles
from the auto EtherChannel.
The following table shows the supported auto-LAG configurations between the actor and partner devices:
Table 4. The supported auto-LAG configurations between the actor and partner devices
Actor/Partner
Active
Passive
Auto
Active
Yes
Yes
Yes
Passive
Yes
No
Yes
Auto
Yes
Yes
Yes
On disabling auto-LAG globally, all auto created Etherchannels become manual EtherChannels.
You cannot add any configurations in an existing auto created EtherChannel. To add, you should first convert it into a manual
EtherChannel by executing the port-channel<channel-number>persistent.
Note
Auto-LAG uses the LACP protocol to create auto EtherChannel. Only one EtherChannel can be automatically created with the
unique partner devices.
Follow these guidelines when configuring the auto-LAG feature.
When auto-LAG is enabled globally and on the port interface, and if you do not want the port interface to become a member
of the auto EtherChannel, disable the auto-LAG on the port interface.
A port interface will not bundle to an auto EtherChannel when it is already a member of a manual EtherChannel. To allow it
to bundle with the auto EtherChannel, first unbundle the manual EtherChannel on the port interface.
When auto-LAG is enabled and auto EtherChannel is created, you can create multiple EtherChannels manually with the same partner
device. But by default, the port tries to create auto EtherChannel with the partner device.
The auto-LAG is supported only on Layer 2 EtherChannel. It is not supported on Layer 3 interface and Layer 3 EtherChannel.
The auto-LAG is supported on cross-stack EtherChannel.
How to Configure EtherChannels
After you configure an EtherChannel, configuration changes applied to the port-channel interface apply to all the physical
ports assigned to the port-channel interface, and configuration changes that are applied to the physical port affect only
the port where you apply the configuration. When bundling physical ports into an EtherChannel, configurine settings on the EtherChannel itself, rather than directly on
its member ports.
The following sections provide various configuration information for EtherChannels:
Configure Layer 2 EtherChannels by assigning ports to a channel group with the channel-group command in interface configuration mode. This command automatically creates the port-channel logical interface.
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
interface interface-id
Example:
Device(config)# interface gigabitethernet 1/0/1
Specifies a physical port, and enters interface configuration mode.
Valid interfaces are physical ports.
For a PAgP EtherChannel, you can configure up to eight ports of the same type and speed for the same group.
For a LACP EtherChannel, you can configure up to 16 Ethernet ports of the same type. Up to eight ports can be active, and
up to eight ports can be in standby mode.
Step 4
switchport mode {access | trunk}
Example:
Device(config-if)# switchport mode access
Assigns all ports as static-access ports in the same VLAN, or configure them as trunks.
If you configure the port as a static-access port, assign it to only one VLAN. The range is 1 to 4094.
Step 5
switchport access vlanvlan-id
Example:
Device(config-if)# switchport access vlan 22
