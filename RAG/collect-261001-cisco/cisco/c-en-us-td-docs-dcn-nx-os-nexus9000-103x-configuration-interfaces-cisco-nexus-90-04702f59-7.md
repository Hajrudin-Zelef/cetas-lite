---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59-7
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59.md
source_anchor: ""
source_lines: [808, 1044]
sha256: 79b62efdc74276775a37ac3715608fb1e7ab17ec887d61ca2113fbc285ccf660
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59

you must configure this parameter before the description displays in the
output.
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
switch# show interface port-channel 2
(Optional)
Displays interface information for the specified port channel.
You can configure
the load-balancing algorithm for port channels that applies to the entire
device.
Note
Use the
no port-channel
load-balance command to restore the default load-balancing
algorithm of source-dest-mac for non-IP traffic and source-dest-ip for IP
traffic.
Specifies the load-balancing algorithm for the device. The range depends on the device. The default for Layer 3 is src-dst ip-l4port for both IPv4 and IPv6, and the default for non-IP is src-dst mac.
Note
GRE inner IP headers supports source, destination and source-destination.
Note
Only the following load-balancing algorithms support symmetric hashing:
src-dst ip
src-dst ip-l4port
Step 3
show port-channel load-balance
Example:
switch(config-router)# show port-channel
load-balance
(Optional) Displays the port-channel load-balancing algorithm.
(Optional) Copies the running configuration to the startup configuration.
Enabling
LACP
LACP is disabled by
default; you must enable LACP before you begin LACP configuration. You cannot
disable LACP while any LACP configuration is present.
LACP learns the
capabilities of LAN port groups dynamically and informs the other LAN ports.
Once LACP identifies correctly matched Ethernet links, it group the links into
a port channel. The port channel is then added to the spanning tree as a single
bridge port.
To configure LACP,
you must do the following:
Enable LACP
globally by using the
feature
lacp command.
You can use
different modes for different interfaces within the same LACP-enabled port
channel. You can change the mode between
active and
passive for an interface only if it is the only
interface that is designated to the specified channel group.
After you enable
LACP, you can configure the channel mode for each individual link in the LACP
port channel as
active or
passive. This channel configuration mode allows the
link to operate with LACP.
When you configure
port channels with no associated aggregation protocol, all interfaces on both
sides of the link remain in the
on
channel mode.
You can configure
the LACP minimum links feature. Although minimum links and maxbundles work only
in LACP, you can enter the CLI commands for these features for non-LACP port
channels, but these commands are nonoperational.
Note
Use the no lacp min-links command to restore the default
port-channel minimum links configuration.
Command
Purpose
no lacp min-links
Example:
switch(config)# no lacp min-links
Restores
the default port-channel minimum links configuration.
Before you begin
Ensure that you are
in the correct port-channel interface.
You can configure
the LACP maxbundle feature. Although minimum links and maxbundles work only in
LACP, you can enter the CLI commands for these features for non-LACP port
channels, but these commands are nonoperational.
Note
Use the
no lacp
max-bundle command to restore the default port-channel max-bundle
configuration.
Command
Purpose
no lacp max-bundle
Example:
switch(config)# no lacp max-bundle
Restores the default port-channel max-bundle configuration.
Before you begin
Ensure that you are
in the correct port-channel interface.
Specifies the
interface to configure, and enters the interface configuration mode.
Step 3
lacp max-bundlenumber
Example:
switch(config-if)# lacp max-bundle
Specifies the
port-channel interface to configure max-bundle.
The default
value for the port-channel max-bundle is 16. The allowed range is from 1 to 32.
Note
Even if the
default value is 16, the number of active members in a port channel is the
minimum of the pc_max_links_config and pc_max_active_members that is allowed in
the port channel.
Step 4
show running-config interface
port-channelnumber
Example:
switch(config-if)# show running-config interface port-channel 3
(Optional)
Displays the port-channel max-bundle configuration.
Example
This example shows
how to configure the port channel interface max-bundle:
You can change the
LACP timer rate to modify the duration of the LACP timeout. Use the
lacp rate
command to set the rate at which LACP
control packets are sent to an LACP-supported interface. You can change the
timeout rate from the default rate (30 seconds) to the fast rate (1 second).
This command is supported only on LACP-enabled interfaces.
Note
We do not
recommend changing the LACP timer rate. HA and SSO are not supported when the
LACP fast rate timer is configured.
Note
Configuring lacp rate fast is not recommended on the vPC Peer-Links. When lacp rate fast is configured on the vPC Peer-Link member interfaces, an alert is displayed in the syslog messages only when the LACP logging
level is set to 5.
This example shows
how to restore the LACP default rate (30 seconds) on Ethernet interface 1/4.
switch# configure terminal
switch (config)# interface ethernet 1/4
switch(config-if)# no lacp rate fast
Configuring the LACP
System Priority
The LACP system ID
is the combination of the LACP system priority value and the MAC address.
Before you begin
Enable LACP.
SUMMARY STEPS
configure terminal
lacp system-prioritypriority
show lacp
system-identifier
copy running-config startup-config
DETAILED STEPS
Command or Action
Purpose
Step 1
configure terminal
Example:
switch# configure terminal
switch(config)#
Enters global
configuration mode.
Step 2
lacp system-prioritypriority
Example:
switch(config)# lacp system-priority 40000
Configures the
system priority for use with LACP. Valid values are from 1 through 65535, and
higher numbers have a lower priority. The default value is 32768.
Note
Each VDC has
a different LACP system ID because the software adds the MAC address to this
configured value.
Specifies the
interface that you want to add to a channel group, and enters the interface
configuration mode.
Step 3
lacp port-prioritypriority
Example:
switch(config-if)# lacp port-priority
40000
Configures the
port priority for use with LACP. Valid values are from 1 through 65535, and
higher numbers have a lower priority. The default value is 32768.
You can configure the MAC address used by the LACP for protocol exchanges and the optional role. By default, the LACP uses
the VDC MAC address. By default, the role is primary.
Use the no lacp system-mac command to make LACP use the default (VDC) MAC address and default role.
This procedure is supported on the Cisco Nexus 9336C-FX2, 93300YC-FX2, and 93240YC-FX2-Z switches.
Before you begin
LACP must be enabled.
SUMMARY STEPS
configure terminal
lacp system-macmac-addressrolerole-value
(Optional) show lacp system-identifier
copy running-config startup-config
DETAILED STEPS
Command or Action
Purpose
Step 1
configure terminal
Example:
switch# configure terminal
Enter global configuration mode.
Step 2
lacp system-macmac-addressrolerole-value
Example:
switch(config)# lacp system-mac 000a.000b.000c role primary
switch(config)# lacp system-mac 000a.000b.000c role secondary
Specifies the MAC address to use in the LACP protocol exchanges. The role is optional. Primary is the default.
Copies the running configuration to the startup configuration.
Example
The following example shows how to configure the role of a switch as primary.
Switch1# sh lacp system-identifier
32768,0-b-0-b-0-b
Switch1# sh run | grep lacp
feature lacp
lacp system-mac 000b.000b.000b role primary
The following example shows how to configure the role of a switch as secondary.
Switch2# sh lacp system-identifier
32768,0-b-0-b-0-b
Switch2# sh run | grep lacp
feature lacp
lacp system-mac 000b.000b.000b role secondary
Disabling LACP
Graceful Convergence
By default, LACP
graceful convergence is enabled. In situations where you need to support LACP
