---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59-6
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "full-duplex", "parameters", "throughput"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59.md
source_anchor: ""
source_lines: [556, 807]
sha256: 5b9e3c390d8c4e48ceda2cca10af6cd6aee0bacd92b3d29966f858ecfe2ceb3c
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59

only when the channel group members are Layer 2 ports (switchport) and trunks
(switchport mode trunk).
Note
Use the
no interface
port-channel command to remove the port channel and delete the
associated channel group.
Command
Purpose
no interface port-channelchannel-number
Example:
switch(config)# no interface port-channel 1
Removes the port channel and deletes the associated channel
group.
Specifies the
port-channel interface to configure, and enters the interface configuration
mode. The range is from 1 to 4096. The Cisco NX-OS software automatically
creates the channel group if it does not already exist.
Step 3
show port-channel summary
Example:
switch(config-router)# show port-channel
summary
(Optional)
Displays information about the port channel.
Step 4
no shutdown
Example:
switch# configure terminal
switch(config)# int e3/1
switch(config-if)# no shutdown
(Optional)
Clears the errors on the interfaces and VLANs where policies correspond with
hardware policies. This command allows policy programming to continue and the
port to come up. If policies do not correspond, the errors are placed in an
error-disabled policy state.
See the
“Compatibility Requirements” section for details on how the interface
configuration changes when you delete the port channel.
Adding a Layer 2
Port to a Port Channel
You can add a Layer
2 port to a new channel group or to a channel group that already contains Layer
2 ports. The software creates the port channel associated with this channel
group if the port channel does not already exist.
Note
Use the
no
channel-group command to remove the port from the channel group.
Command
Purpose
no channel-group
Example:
switch(config)# no channel-group
Removes the port from the channel group.
Before you begin
Enable LACP if you
want LACP-based port channels.
All Layer 2 member
ports must run in full-duplex mode and at the same speed
(Optional) Configures necessary parameters for a Layer 2 trunk port.
Step 6
channel-groupchannel-number [force] [mode {on |
active
|
passive}]
Example:
switch(config-if)# channel-group 5
switch(config-if)# channel-group 5 force
Configures the
port in a channel group and sets the mode. The channel-number range is from 1
to 4096. This command creates the port channel associated with this channel
group if the port channel does not already exist. All static port-channel
interfaces are set to mode
on. You must set all LACP-enabled port-channel
interfaces to
active or
passive. The default mode is
on.
(Optional)
Forces an interface with some incompatible configurations to join the channel.
The forced interface must have the same speed, duplex, and flow control
settings as the channel group.
Note
The
force option fails if the port has a QoS policy
mismatch with the other members of the port channel.
Step 7
show interfacetypeslot/port
Example:
switch# show interface port channel 5
(Optional)
Displays interface information.
Step 8
no shutdown
Example:
switch# configure terminal
switch(config)# int e3/1
switch(config-if)# no shutdown
(Optional)
Clears the errors on the interfaces and VLANs where policies correspond with
hardware policies. This command allows policy programming to continue and the
port to come up. If policies do not correspond, the errors are placed in an
error-disabled policy state.
You can add a Layer
3 port to a new channel group or to a channel group that is already configured
with Layer 3 ports. The software creates the port channel associated with this
channel group if the port channel does not already exist.
If the Layer 3 port
that you are adding has a configured IP address, the system removes that IP
address before adding the port to the port channel. After you create a Layer 3
port channel, you can assign an IP address to the port-channel interface.
Note
Use the
no
channel-group command to remove the port from the channel group.
The port reverts to its original configuration. You must reconfigure the IP
addresses for this port.
Command
Purpose
no channel-group
Example:
switch(config)# no channel-group
Removes the port from the channel group.
Before you begin
Enable LACP if you
want LACP-based port channels.
Remove any IP
addresses configured on the Layer 3 interface.
SUMMARY STEPS
configure terminal
interfacetypeslot/port
no switchport
channel-groupchannel-number [force] [mode {on |
active
|
passive}]
Specifies the
interface that you want to add to a channel group, and enters the interface
configuration mode.
Step 3
no switchport
Example:
switch(config-if)# no switchport
Configures the
interface as a Layer 3 port.
Step 4
channel-groupchannel-number [force] [mode {on |
active
|
passive}]
Example:
switch(config-if)# channel-group 5
switch(config-if)# channel-group 5 force
Configures the
port in a channel group and sets the mode. The channel-number range is from 1
to 4096. The Cisco NX-OS software creates the port channel associated with this
channel group if the port channel does not already exist.
(Optional)
Forces an interface with some incompatible configurations to join the channel.
The forced interface must have the same speed, duplex, and flow control
settings as the channel group.
Step 5
show interfacetypeslot/port
Example:
switch# show interface ethernet 1/4
(Optional)
Displays interface information.
Step 6
no shutdown
Example:
switch# configure terminal
switch(config)# int e3/1
switch(config-if)# no shutdown
(Optional)
Clears the errors on the interfaces and VLANs where policies correspond with
hardware policies. This command allows policy programming to continue and the
port to come up. If policies do not correspond, the errors are placed in an
error-disabled policy state.
Specifies the
bandwidth, which is used for informational purposes. The range is from 1 to
3,200,000,000 kbs. The default value depends on the total active interfaces in
the channel group.
Step 4
delayvalue
Example:
switch(config-if)# delay 10000
switch(config-if)#
Specifies the
throughput delay, which is used for informational purposes. The range is from 1
to 16,777,215 tens of microseconds. The default value is 10 microseconds.
Step 5
exit
Example:
switch(config-if)# exit
switch(config)#
Exits the
interface mode and returns to the configuration mode.
Step 6
show interface port-channelchannel-number
Example:
switch# show interface port-channel 2
(Optional)
Displays interface information for the specified port channel.
Shutting Down and
Restarting the Port-Channel Interface
You can shut down
and restart the port-channel interface. When you shut down a port-channel
interface, no traffic passes and the interface is administratively down.
Specifies the
port-channel interface that you want to configure, and enters the interface
mode.
Step 3
shutdown
Example:
switch(config-if)# shutdown
switch(config-if)#
Shuts down the
interface. No traffic passes and the interface displays as administratively
down. The default is no shutdown.
Note
Use the no shutdown command to open the interface.
The interface
displays as administratively up. If there are no operational problems, traffic
passes. The default is no shutdown.
Step 4
exit
Example:
switch(config-if)# exit
switch(config)#
Exits the
interface mode and returns to the configuration mode.
Step 5
show interface port-channelchannel-number
Example:
switch(config-router)# show interface port-channel 2
(Optional)
Displays interface information for the specified port channel.
Step 6
no shutdown
Example:
switch# configure terminal
switch(config)# int e3/1
switch(config-if)# no shutdown
(Optional)
Clears the errors on the interfaces and VLANs where policies correspond with
hardware policies. This command allows policy programming to continue and the
port to come up. If policies do not correspond, the errors are placed in an
error-disabled policy state.
Allows you to
add a description to the port-channel interface. You can use up to 80
characters in the description. By default, the description does not display;
