---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046-5
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046.md
source_anchor: ""
source_lines: [478, 629]
sha256: 84e54624478348541f75cdc3a997df3e8861426d2538ec70930672b21ef04644
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046

time, and maximum-age time for a network of that diameter, which can significantly reduce the convergence time. You can use
the hello keyword to override the automatically calculated hello time.
To configure the root device, perform this procedure:
Configures a device to become the root for the specified VLAN.
For vlan-id, you can specify a single VLAN identified by VLAN ID number, a range of VLANs separated by a hyphen, or a series of VLANs
separated by a comma. The range is 1 to 4094.
(Optional) For diameternet-diameter, specify the maximum number of devices between any two end stations. The range is 2 to 7.
Step 4
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
What to do next
After configuring the switch as the root switch, we recommend that you avoid manually configuring the hello time, forward-delay
time, and maximum-age time through the spanning-tree vlanvlan-idhello-time, spanning-tree vlanvlan-idforward-time, and the spanning-tree vlanvlan-idmax-age global configuration commands.
(Optional) Configuring a Secondary Root Device
When you configure a switch as the secondary root, the switch priority is modified from the default value (32768) to 28672.
With this priority, the switch is likely to become the root switch for the specified VLAN if the primary root switch fails.
This is assuming that the other network switches use the default switch priority of 32768, and therefore, are unlikely to
become the root switch.
You can execute this command on more than one switch to configure multiple backup root switches. Use the same network diameter
and hello-time values that you used when you configured the primary root switch with the spanning-tree vlanvlan-idroot primary global configuration command.
To configure a secondary root device, perform this procedure:
Configures a device to become the secondary root for the specified VLAN.
For vlan-id, you can specify a single VLAN identified by VLAN ID number, a range of VLANs separated by a hyphen, or a series of VLANs
separated by a comma. The range is 1 to 4094.
(Optional) For diameternet-diameter, specify the maximum number of devices between any two end stations. The range is 2 to 7.
Use the same network diameter value that you used when configuring the primary root switch.
Step 4
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
(Optional) Configuring Port Priority
To configure port priority, perform this procedure:
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
Specifies an interface to configure, and enters interface configuration mode.
Valid interfaces include physical ports and port-channel logical interfaces (port-channelport-channel-number).
Step 4
spanning-tree port-prioritypriority
Example:
Device(config-if)# spanning-tree port-priority 0
Configures the port priority for an interface.
For priority, the range is 0 to 240, in increments of 16; the default is 128. Valid values are 0, 16, 32, 48, 64, 80, 96, 112, 128, 144,
160, 176, 192, 208, 224, and 240. All other values are rejected. The lower the number, the higher the priority.
For vlan-id, you can specify a single VLAN identified by VLAN ID number, a range of VLANs separated by a hyphen, or a series of VLANs
separated by a comma. The range is 1 to 4094.
For priority, the range is 0 to 240, in increments of 16; the default is 128. Valid values are 0, 16, 32, 48, 64, 80, 96, 112, 128, 144,
160, 176, 192, 208, 224, and 240. All other values are rejected. The lower the number, the higher the priority.
Step 6
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
(Optional) Configuring Path Cost
To configure path cost, perform this procedure:
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
Specifies an interface to configure, and enters interface configuration mode. Valid interfaces include physical ports and
port-channel logical interfaces (port-channelport-channel-number).
Step 4
spanning-tree costcost
Example:
Device(config-if)# spanning-tree cost 250
Configures the cost for an interface.
If a loop occurs, spanning tree uses the path cost when selecting an interface to place into the forwarding state. A lower
path cost represents higher-speed transmission.
For cost, the range is 1 to 200000000; the default value is derived from the media speed of the interface.
If a loop occurs, spanning tree uses the path cost when selecting an interface to place into the forwarding state. A lower
path cost represents higher-speed transmission.
For vlan-id, you can specify a single VLAN identified by VLAN ID number, a range of VLANs separated by a hyphen, or a series of VLANs
separated by a comma. The range is 1 to 4094.
For cost, the range is 1 to 200000000; the default value is derived from the media speed of the interface.
Step 6
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
The show spanning-treeinterfaceinterface-id privileged EXEC command displays information only for ports that are in a link-up operative state. Otherwise, you can use
the show running-config privileged EXEC command to confirm the configuration.
(Optional) Configuring the Device Priority of a VLAN
You can configure the switch priority and make it more likely that a standalone switch or a switch in the stack will be chosen
as the root switch.
Note
Exercise care when using this command. For most situations, we recommend that you use the spanning-tree vlanvlan-idroot primary and the spanning-tree vlanvlan-idroot secondary global configuration commands to modify the switch priority.
To configure device priority of a VLAN, perform this procedure:
For vlan-id, you can specify a single VLAN identified by VLAN ID number, a range of VLANs separated by a hyphen, or a series of VLANs
separated by a comma. The range is 1 to 4094.
For priority, the range is 0 to 61440 in increments of 4096; the default is 32768. The lower the number, the more likely the switch will
be chosen as the root switch.
Valid priority values are 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, 57344,
and 61440. All other values are rejected.
Step 4
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
(Optional) Configuring the Hello Time
The hello time is the time interval between configuration messages that are generated and sent by the root switch.
To configure the hello time, perform this procedure:
Configures the hello time of a VLAN. The hello time is the time interval between configuration messages that are generated
and sent by the root switch. These messages mean that the switch is alive.
For vlan-id, you can specify a single VLAN identified by VLAN ID number, a range of VLANs separated by a hyphen, or a series of VLANs
separated by a comma. The range is 1 to 4094.
For seconds, the range is 1 to 10; the default is 2.
Step 3
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
(Optional) Configuring the Forwarding-Delay Time for a VLAN
To configure the forwarding-delay time for a VLAN, perform this procedure:
Configures the forward time of a VLAN. The forwarding delay is the number of seconds an interface waits before changing from
its spanning-tree learning and listening states to the forwarding state.
For vlan-id, you can specify a single VLAN identified by VLAN ID number, a range of VLANs separated by a hyphen, or a series of VLANs
separated by a comma. The range is 1 to 4094.
For seconds, the range is 4 to 30; the default is 15.
Step 4
end
