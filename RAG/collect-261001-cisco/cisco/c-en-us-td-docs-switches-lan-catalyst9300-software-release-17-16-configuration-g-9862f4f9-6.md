---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9-6
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9.md
source_anchor: ""
source_lines: [513, 712]
sha256: 2841ce9f85c13da919900107f812c439256e61d7e1fe7053bca0631caa435e94
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9

LACP system priority
System ID (the device MAC address)
LACP port priority
Port number
In priority comparisons, numerically lower values have higher priority. The priority decides which ports should be put in
standby mode when there is a hardware limitation that prevents all compatible ports from aggregating.
Determining which ports are active and which are hot standby is a two-step procedure. First the system with a numerically
lower system priority and system ID is placed in charge of the decision. Next, that system decides which ports are active
and which are hot standby, based on its values for port priority and port number. The port priority and port number values
for the other system are not used.
You can change the default values of the LACP system priority and the LACP port priority to affect how the software selects
active and standby links.
(Optional) Configuring the Link Aggregation Control Protocol Max Bundle
When you specify the maximum number of bundled LACP ports allowed in a port channel, the remaining ports in the port channel
are designated as hot-standby ports.
Beginning in privileged EXEC mode, follow these steps to configure the maximum number of LACP ports in a port channel.
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
interface port-channelchannel-number
Example:
Device(config)# interface port-channel 2
Enters interface configuration mode for a port channel.
For channel-number, the range is 1 to 128.
Step 4
lacp max-bundlemax-bundle-number
Example:
Device(config-if)# lacp max-bundle 3
Specifies the maximum number of LACP ports in the port-channel bundle.
The range is 1 to 8.
Step 5
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
Configuring Link Aggregation Control Protocol Port-Channel Standalone Disable
To disable the standalone EtherChannel member port state on a port channel, perform this task on the port channel interface:
Configuring the Link Aggregation Control Protocol Port Channel Min-Links
You can specify the minimum number of active ports that must be in the link-up state and bundled in an EtherChannel for the
port channel interface to transition to the link-up state. Using EtherChannel min-links, you can prevent low-bandwidth LACP
EtherChannels from becoming active. Port channel min-links also cause LACP EtherChannels to become inactive if they have too
few active member ports to supply the required minimum bandwidth.
To configure the minimum number of links that are required for a port channel. Perform the following tasks.
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
configure terminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
interface port-channelchannel-number
Example:
Device(config)# interface port-channel 2
Enters interface configuration mode for a port-channel.
For channel-number, the range is 1 to 128.
Step 4
port-channel min-linksmin-links-number
Example:
Device(config-if)# port-channel min-links 3
Specifies the minimum number of member ports that must be in the link-up state and bundled in the EtherChannel for the port
channel interface to transition to the link-up state.
For min-links-number, the range is 2 to 8.
Step 5
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
(Optional) Configuring the Link Aggregation Control Protocol System Priority
You can configure the system priority for all the EtherChannels that are enabled for LACP by using the lacp system-priority command in global configuration mode. You cannot configure a system priority for each LACP-configured channel. By changing
this value from the default, you can affect how the software selects active and standby links.
You can use the show etherchannel summary command in privileged EXEC mode to see which ports are in the hot-standby mode (denoted with an H port-state flag).
Follow these steps to configure the LACP system priority.
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
lacp system-prioritypriority
Example:
Device(config)# lacp system-priority 32000
Configures the LACP system priority.
The range is 1 to 65535. The default is 32768.
The lower the value, the higher the system priority.
Step 4
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
(Optional) Configuring the Link Aggregation Control Protocol Port Priority
By default, all ports use the same port priority. If the local system has a lower value for the system priority and the system
ID than the remote system, you can affect which of the hot-standby links become active first by changing the port priority
of LACP EtherChannel ports to a lower value than the default. The hot-standby ports that have lower port numbers become active
in the channel first. You can use the show etherchannel summary privileged EXEC command to see which ports are in the hot-standby mode (denoted with an H port-state flag).
Note
If LACP is not able to aggregate all the ports that are compatible (for example, the remote system might have more restrictive
hardware limitations), all the ports that cannot be actively included in the EtherChannel are put in the hot-standby state
and are used only if one of the channeled ports fails.
Follow these steps to configure the LACP port priority.
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
configure terminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
interface interface-id
Example:
Device(config)# interface gigabitethernet 1/0/2
Specifies the port to be configured, and enters interface configuration mode.
Step 4
lacp port-prioritypriority
Example:
Device(config-if)# lacp port-priority 32000
Configures the LACP port priority.
The range is 1 to 65535. The default is 32768. The lower the value, the more likely that the port will be used for LACP transmission.
Step 5
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
Configuring Link Aggregation Control Protocol 1:1 Redundancy
Note
LACP 1:1 redundancy must be enabled at both ends of the LACP EtherChannel.
For the LACP 1:1 Redundancy feature to work, the lacp max-bundle 1 command must be configured along with the lacp fast-switchover command.
For the LACP 1:1 Hot Standby Dampening feature to work, the lacp max-bundle 1 and lacp fast-switchover commands must be configured before the lacp fast-switchover dampening command is configured.
To configure LACP 1:1 redundancy, perform this procedure:
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password, if prompted.
Step 2
configure terminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
interface port-channelgroup_number
Example:
Device(config)# interface port-channel 40
Selects an LACP port channel interface and enters interface configuration mode.
Step 4
lacp fast-switchover
Example:
Device(config-if)# lacp fast-switchover
Enables the LACP 1:1 Redundancy feature on the EtherChannel.
Step 5
lacp max-bundle 1
Example:
Device(config-if)# lacp max-bundle 1
Sets the maximum number of active member ports to be one. The only value that is supported with LACP 1:1 redundancy is 1.
(Optional) Enables the LACP 1:1 Hot Standby Dampening feature for this EtherChannel. The range for the time parameter is from
30 to 180 seconds.
Step 7
end
Example:
Device(config-if)# end
Exits interface configuration mode and returns to privileged EXEC mode.
