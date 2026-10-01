---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59-8
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59.md
source_anchor: ""
source_lines: [1045, 1208]
sha256: 74c5e020b83d70468994ded8130754fc914f2e20622aa294f7386d4657918131
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59

interoperability with devices where the graceful failover defaults may delay
the time taken for a disabled port to be brought down or cause traffic from the
peer to be lost, you can disable convergence. If the downstream access switch
is not a Cisco Nexus device, disable the LACP graceful convergence option.
Note
The port channel
has to be in the administratively down state before the command can be run.
LACP sets a port to
the suspended state if it does not receive an LACP PDU from the peer. This
process can cause some servers to fail to boot up as they require LACP to
logically bring up the port.
Note
You should only enter the lacp suspend-individual command on edge ports. The port channel has to be in the administratively down state before you can use this command.
The delayed LACP feature enables one port channel member, the delayed LACP port, to come up first as a member of a regular
port channel before LACP PDUs are received. You configure the delayed LACP feature using the lacp mode delaycommand on a port channel followed by configuring the LACP port priority on a one member port of the port channel.
Note
For vPC, you must
enable the delayed LACP on both vPC switches.
Note
For vPC, when
the delayed LACP port is on the primary switch and the primary switch fails to
boot, you need to remove the vPC configuration on the delayed LACP port-channel
of the acting primary switch and flap the port-channel for a new port to be
chosen as the delayed LACP port on the existing port-channel.
SUMMARY STEPS
configure terminal
interface port-channelnumber
lacp mode delay
DETAILED STEPS
Command or Action
Purpose
Step 1
configure terminal
Enters global
configuration mode.
Step 2
interface port-channelnumber
Specifies the
port channel interface to configure and enters the interface configuration
mode.
Step 3
lacp mode delay
Enables delayed
LACP.
Note
To disable
delayed LACP, use the
no lacp mode
delay command.
Complete the
configuration of the delayed LACP by configuring the LACP port priority. See
the "Configuring the LACP Port Priority" section for details.
The priority
of a LACP port determines the election of the delayed LACP port. The port with
the lowest numerical priority is elected.
When two or
more ports have the same best priority, the VDC system MAC is used to determine
which vPC is used. Then within a non-vPC switch or the elected vPC switch, the
smallest of the ethernet port names is used.
When the
delayed LACP feature is configured and made effective with a port channel flap,
the delayed LACP port operates as a member of a regular port channel, allowing
data to be exchanged between the server and switch. After receiving the first
LACP PDU, the delayed LACP port transitions from a regular port member to a
LACP port member.
Note
The election
of the delayed LACP port is not complete or effective until the port channel
flaps on the switch or at a remote server.
switch# config terminal
switch(config)# interface po 1
switch(config-if)# no lacp mode delay
Configuring Port
Channel Hash Distribution
Cisco NX-OS supports the adaptive and fixed hash distribution
configuration for both global and port-channel levels. This option minimizes
traffic disruption by minimizing Result Bundle Hash (RBH) distribution changes
when members come up or go down so that flows that are mapped to unchange RBH
values continue to flow through the same links. The port-channel level
configuration overrules the global configuration. The default configuration is
adaptive globally, and there is no configuration for each port channel, so
there is no change during an ISSU. No ports are flapped when the command is
applied, and the configuration takes effect at the next member link change
event. Both modes work with RBH module or non-module schemes.
During an ISSD to a lower version that does not support this feature,
you must disable this feature if the fixed mode command is being used globally
or if there is a port-channel level configuration.
(Optional)
Copies the running configuration to the startup configuration.
Example
This example shows
how to configure hash distribution as a global-level command:
switch# configure terminal
switch(config)# no port-channel hash-distribution fixed
Enabling ECMP Resilient Hashing
Resilient ECMP ensures minimal impact to the existing flows when members are deleted from an ECMP group. This is achieved
by replicating the existing members in a round-robin fashion at the indices that were previously occupied by the deleted members.
SUMMARY STEPS
configure terminal
hardware profile ecmp resilient
copy running-config startup-config
reload
DETAILED STEPS
Command or Action
Purpose
Step 1
configure terminal
Example:
switch# configure terminal
Enters global configuration mode.
Step 2
hardware profile ecmp resilient
Example:
switch(config)# hardware profile ecmp resilient
Enables ECMP resilient hashing and displays the following: Warning: The command will take effect after next reload.
Note
This command is not supported on Cisco Nexus 9808 platform switches.
Copies the running configuration to the startup configuration.
Step 4
reload
Example:
switch(config)# reload
Reboots the switch.
Configuring ECMP Load Balancing
To configure the ECMP load-sharing algorithm, use the following command in global configuration mode:
Before you begin
SUMMARY STEPS
ip load-sharing address {destination port destination | source-destination [port source-destination | gre | gtpu | ipv6-flowlabel | ttl | udfoffsetoffsetlengthlength
| symmetricinnerallgreheader]} [universal-idseed] [rotaterotate] [concatenation]
(Optional)
show ip load-sharing
DETAILED STEPS
Command or Action
Purpose
Step 1
ip load-sharing address {destination port destination | source-destination [port source-destination | gre | gtpu | ipv6-flowlabel | ttl | udfoffsetoffsetlengthlength
| symmetricinnerallgreheader]} [universal-idseed] [rotaterotate] [concatenation]
Example:
ip load-sharing address source-destination
Example:
switch(config)# ip load-sharing address source-destination ipv6-flowlabel
Example:
switch(config)# ip load-sharing address source-destination ttl
Example:
switch(config)# ip load-sharing address source-destination udf offset 8 length 8
Example:
switch(config)# [no] ip load-sharing address source-destination port source-destination symmetric
Example:
switch(config)# ip load-sharing address source-destination port source-destination inner [all|greheader]
Configures the ECMP load-sharing algorithm for data traffic.
The gre option specifies the source-destination value for the Generic Routing Encapsulation (GRE) key.
The gtpu option specifies the GPRS Tunneling Protocol (GTP) tunnel endpoint identifier (TEID) value for the port source-destination.
The ipv6-flowlabel option includes the IPv6 flow label for computing ECMP hashing. It ensures that traffic flows are distributed on all links
based on different flow label values. Enabling or disabling this option also enables or disables it for port-channel load-balancing
if Layer 4 parameters are enabled using the port-channel load-balance command. Only the following devices support this option:
Cisco Nexus 9364C and 9300-EX/FX/FX2 platform switches
Cisco Nexus 9500 platform switches with X9700-EX/FX line cards and FM-E2 fabric modules in all routing modes
Cisco Nexus 9500 platform switches with X9700-EX/FX line cards and FM-E fabric modules in non-hierarchical routing modes where
IPv6 routes are programmed in the line card
Beginning with Cisco NX-OS Release 9.3(5), Cisco Nexus N9K-C9316D-GX, N9K-C93600CD-GX, N9K-C9364C-GX switches support this
option.
The ttl option includes time-to-live information for computing ECMP hashing. It ensures that traffic flows are distributed on all
links based on different TTL values. For IPv4 flows, it is based on ttl values. For IPv6 flows, it is based on hop limit.
Enabling or disabling this option also enables or disables it for port-channel load-balancing if Layer 4 parameters are enabled
