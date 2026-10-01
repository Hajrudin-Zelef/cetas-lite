---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9-7
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9.md
source_anchor: ""
source_lines: [713, 897]
sha256: 40e27e9cbf8ba6b216469ad594b465e3399f6ae6c13a29ce0aee44c90be6e3ad
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9

Configuring Link Aggregation Control Protocol 1:1 Redundancy Fast Rate Timer
You can change the LACP timer rate to modify the duration of the LACP timeout. Use the lacp rate command to set the rate at which LACP control packets are received by an LACP-supported interface. You can change the timeout
rate from the default rate (30 seconds) to the fast rate (1 second). This command is supported only on LACP-enabled interfaces.
To configure LACP 1:1 redundancy fast rate timer, perform this procedure:
Configures an interface and enters interface configuration mode.
Step 4
lacp rate{ normal | fast}
Example:
Device(config-if)# lacp rate fast
Configures the rate at which LACP control packets are received by an LACP-supported interface.
To reset the timeout rate to its default, use the no lacp rate command.
Step 5
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
Step 6
show lacp internal
Example:
Device# show lacp internal
Device# show lacp counters
Verifies your configuration.
Configuring Auto-LAG Globally
To configure Auto-LAG globally, perform this procedure:
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
[no] port-channel auto
Example:
Device(config)# port-channel auto
Enables the auto-LAG feature on a switch globally. Use the no form of this command to disable the auto-LAG feature on the
switch globally.
Note
By default, the auto-LAG feature is enabled on the port.
Step 4
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
Step 5
show etherchannel auto
Example:
Device# show etherchannel auto
Displays that EtherChannel is created automatically.
Configuring Auto-LAG on a Port Interface
To configure Auto-LAG on a port interface, perform this procedure:
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
Device(config)# interface gigabitethernet 1/0/1
Specifies the port interface to be enabled for auto-LAG, and enters interface configuration mode.
Step 4
[no] channel-group auto
Example:
Device(config-if)# channel-group auto
(Optional) Enables auto-LAG feature on individual port interface. Use the no form of this command to disable the auto-LAG
feature on individual port interface.
Note
By default, the auto-LAG feature is enabled on the port.
Step 5
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
Step 6
show etherchannel auto
Example:
Device# show etherchannel auto
Displays that EtherChannel is created automatically.
Configuring Persistence with Auto-LAG
You use the persistence command to convert the auto created EtherChannel into a manual one and allow you to add configuration
on the existing EtherChannel.
To configure persistence with Auto-LAG, perform this procedure:
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
port-channelchannel-numberpersistent
Example:
Device# port-channel 1 persistent
Converts the auto created EtherChannel into a manual one and allows you to add configuration on the EtherChannel.
Step 3
show etherchannel summary
Example:
Device# show etherchannel summary
Displays the EtherChannel information.
Monitoring EtherChannel, Port Aggregation Protocol, and Link Aggregation Control Protocol Status
You can display EtherChannel, PAgP, and LACP status using the commands listed in this table.
Table 5. Commands for Monitoring EtherChannel, PAgP, and LACP Status
Clears PAgP channel-group information and traffic counters.
show etherchannel[ channel-group-number { detail | load-balance | port | port-channel | protocol| summary}] [ detail | load-balance | port | port-channel | protocol | auto | summary]
Displays EtherChannel information in a brief, detailed, and one-line summary form. Also displays the load-balance or frame-distribution
scheme, port, port-channel, protocol, and Auto-LAG information.
show pagp [ channel-group-number] { counters | internal | neighbor}
Displays PAgP information such as traffic information, the internal PAgP configuration, and neighbor information.
show pagp [channel-group-number] dual-active
Displays the dual-active detection status.
show lacp [channel-group-number] {counters | internal | neighbor | sys-id}
Displays LACP information such as traffic information, the internal LACP configuration, and neighbor information.
show running-config
Verifies your configuration entries.
show etherchannel load-balance
Displays the load balance or frame distribution scheme among ports in the port channel.
Configuration Examples for EtherChannels
The following sections provide various configuration examples for EtherChannels:
This example shows how to configure an EtherChannel on a single switch in the stack. It assigns two ports as static-access ports in VLAN 10 to channel 5 with the PAgP mode desirable:
This example shows how to configure an EtherChannel on a single switch in the stack. It assigns two ports as static-access ports in VLAN 10 to channel 5 with the LACP mode active:
Device# configure terminal
Device(config)# interface range gigabitethernet2/0/1 -2
Device(config-if-range)# switchport mode access
Device(config-if-range)# switchport access vlan 10
Device(config-if-range)# channel-group 5 mode active
Device(config-if-range)# end
This example shows how to configure a cross-stack EtherChannel. It uses LACP passive mode and assigns two ports on stack member
1 and one port on stack member 2 as static-access ports in VLAN 10 to channel 5:
PoE or LACP negotiation errors may occur if you configure two ports from switch to the access point (AP). This scenario can
be avoided if the port channel configuration is on the switch side. For more details, see the following example:
If the port reports LACP errors on port flap, you should include the following command as well: no errdisable detect cause pagp-flap
Example: Configuring Layer 3 EtherChannels
This example shows how to configure a Layer 3 EtherChannel. It assigns two ports to channel 5 with the LACP mode active:
Device# configure terminal
Device(config)# interface range gigabitethernet2/0/1 -2
Device(config-if-range)# no ip address
Device(config-if-range)# no switchport
Device(config-if-range)# channel-group 5 mode active
Device(config-if-range)# end
This example shows how to configure a cross-stack Layer 3 EtherChannel. It assigns two ports on stack member 2 and one port
on stack member 3 to channel 7 using LACP active mode:
Device# configure terminal
Device(config)# interface range gigabitethernet2/0/4 -5
Device(config-if-range)# no ip address
Device(config-if-range)# no switchport
Device(config-if-range)# channel-group 7 mode active
Device(config-if-range)# exit
Device(config)# interface gigabitethernet3/0/3
Device(config-if)# no ip address
Device(config-if)# no switchport
Device(config-if)# channel-group 7 mode active
Device(config-if)# exit
Example: Configuring Link Aggregation Control Protocol Hot-Standby Ports
This example shows how to configure an EtherChannel (port channel 2) that will be active when there are at least three active
ports, will comprise up to seven active ports and the remaining ports (up to nine) as hot-standby ports:
This is a sample output from the show lacp internal command:
Device# show lacp 1 internal
Flags: S - Device is requesting Slow LACPDUs
F - Device is requesting Fast LACPDUs
A - Device is in Active mode
P - Device is in Passive mode
Channel group 1,[146 s left to exit dampening state]
LACP port Admin Oper Port Port
Port Flags State Priority Key Key Number State
Fa1/1 FA hot-sby 30000* 0x1 0x1 0x103 0x7
Fa1/2 SA bndl 32768 0x1 0x1 0x102 0x3D
