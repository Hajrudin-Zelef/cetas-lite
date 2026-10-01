---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59-4
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59.md
source_anchor: ""
source_lines: [303, 415]
sha256: 553c0c4b51b034912fc73f416fa2dc9063097970bd8f8de623d5fbdff3f2d60c
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59

Two devices can form an LACP port channel when their ports are in different LACP modes if the modes are compatible as in the
following example:
Table 2. Channel Modes Compatibility
Device 1 > Port-1
Device 2 > Port-2
Result
Active
Active
Can form a port channel.
Active
Passive
Can form a port channel.
Passive
Passive
Cannot form a port channel because no ports can initiate negotiation.
On
Active
Cannot form a port channel because LACP is enabled only on one side.
On
Passive
Cannot form a port channel because LACP is not enabled.
Each system that runs LACP has an LACP system priority value. You can accept the default value of 32768 for this parameter,
or you can configure a value between 1 and 65535. LACP uses the system priority with the MAC address to form the system ID
and also uses the system priority during negotiation with other devices. A higher system priority value means a lower priority.
Note
The LACP system ID is the combination of the LACP system priority value and the MAC address.
LACP Port Priority
Each port that is configured to use LACP has an LACP port priority. You can accept the default value of 32768 for the LACP
port priority, or you can configure a value between 1 and 65535. LACP uses the port priority with the port number to form
the port identifier.
LACP uses the port priority to decide which ports should be put in standby mode when there is a limitation that prevents all
compatible ports from aggregating and which ports should be put into active mode. A higher port priority value means a lower
priority for LACP. You can configure the port priority so that specified ports have a lower priority for LACP and are most
likely to be chosen as active links, rather than hot-standby links.
LACP Administrative Key
LACP automatically configures an administrative key value equal to the channel-group number on each port configured to use
LACP. The administrative key defines the ability of a port to aggregate with other ports.
A port’s ability to aggregate with other ports is determined by these factors:
Port physical characteristics, such as the data rate and the duplex capability
Configuration restrictions that you establish
LACP Marker Responders
You can dynamically redistribute the data traffic by using port channels. This redistribution might result from a removed
or added link or a change in the load-balancing scheme. Traffic redistribution that occurs in the middle of a traffic flow
can cause misordered frames.
LACP uses the Marker Protocol to ensure that frames are not duplicated or reordered due to this redistribution. The Marker
Protocol detects when all the frames of a given traffic flow are successfully received at the remote end. LACP sends Marker
PDUs on each of the port-channel links. The remote system responds to the Marker PDU once it receives all the frames received
on this link prior to the Marker PDU. The remote system then sends a Marker Responder. Once the Marker Responders are received
by the local system on all member links of the port channel, the local system can redistribute the frames in the traffic flow
with no chance of misordering. The software supports only Marker Responders.
LACP-Enabled and Static Port Channels Differences
The following table summarizes the major differences between port channels with LACP enabled and static port channels.
Table 3. Port Channels with LACP Enabled and Static Port Channels
Configurations
Port Channels with LACP Enabled
Static Port Channels
Protocol applied
Enable globally
Not applicable
Channel mode of links
Can be either:
Active
Passive
Can only be On
Maximum number of links in channel
32
32
LACP Compatibility Enhancements
When a Cisco Nexus 9000 Series device is connected to a non-Nexus peer, its graceful failover defaults may delay the time
that is taken to bring down a disabled port or cause traffic from the peer to be lost. To address these conditions, the lacp graceful-convergence command was added.
By default, LACP sets a port to suspended state if it does not receive an LACP PDU from the peer. lacp suspend-individual is a default configuration on Cisco Nexus 9000 series switches. This command puts the port in suspended state if it does
not receive any LACP PDUs. In some cases, although this feature helps in preventing loops created due to misconfigurations,
it can cause servers fail to boot up because they require LACP to logically bring up the port. You can put a port into an
individual state by using the no lacp suspend-individual. Port in individual sate takes attributes of the individual port based on the port configuration.
LACP port-channels exchange LACP PDUs for quick bundling of links when connecting a server and a switch. However, the links
go into suspended state when the PDUs are not received.
The delayed LACP feature enables one port-channel member, the delayed-LACP port, to come up first as a member of a regular port-channel before
LACP PDUs are received. After it is connected in LACP mode, other members, the auxiliary LACP ports, are brought up. This
avoids having the links becoming suspended when PDUs are not received.
Which port in the port-channel comes up first depends on the port-priority value of the ports. A member link in a port channel
with lowest priority value, will come come up first as a LACP delayed port. Regardless of the operational status of the links,
the configured priority of a LACP port is used to select the delayed-lacp port
Guidelines and Limitations
This feature supports Layer 2 port-channels with or without VPC running in spanning-tree port type trunk mode. These guidelines
and limitations apply to LACP:
Using no lacp suspend-individual and lacp mode delay on a same port channel is not recommended because it can put non-lacp delayed ports in individual state. As a best practice,
you must avoid combining these two configurations.
Not supported on Layer 3 port-channels.
Not supported on Nexus 9000 switches on the FEX NIF fabric port-channel or FEX HIF host port-channels
LACP Port-Channel Minimum Links and LACP MaxBundle
A port channel aggregates similar ports to provide increased bandwidth in a single manageable interface.
The introduction of the minimum links and LACP MaxBundle feature further refines LACP port-channel operation and provides
increased bandwidth in one manageable interface.
The LACP port-channel minimum links feature does the following:
Configures the minimum number of ports that must be linked up and bundled in the LACP port channel.
Prevents the low-bandwidth LACP port channel from becoming active.
Causes the LACP port channel to become inactive if there are few active members ports to supply the required minimum bandwidth.
The LACP MaxBundle defines the maximum number of bundled ports allowed in a LACP port channel.
The LACP MaxBundle feature does the following:
Defines an upper limit on the number of bundled ports in an LACP port channel.
Allows hot-standby ports with fewer bundled ports. (For example, in an LACP port channel with five ports, you can designate
two of those ports as hot-standby ports.)
Note
The minimum links and LACP MaxBundle features only work with LACP port-channels. The switch allows you to configure these
features on non-LACP port-channels, but the features are not operational.
LACP Fast Timers
You can change the LACP timer rate to modify the duration of the LACP timeout. Use the lacp rate command to set the rate at
which LACP control packets are sent to an LACP-supported interface. You can change the timeout rate from the default rate
(30 seconds) to the fast rate (1 second). This command is supported only on LACP-enabled interfaces. To configure the LACP
fast time rate, see the “Configuring the LACP Fast Timer Rate” section.
ISSU and ungraceful switchovers are not supported with LACP fast timers.
Virtualization Support
You must configure the member ports and other port channel-related configuration from the virtual device context (VDC) that
