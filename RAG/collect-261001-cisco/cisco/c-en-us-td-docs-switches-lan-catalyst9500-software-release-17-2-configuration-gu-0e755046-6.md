---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046-6
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046.md
source_anchor: ""
source_lines: [630, 724]
sha256: 9169bafa5a01c5814563257008b3009c1b1ed69c3b97194986498378beed81b6
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046

Example:
Device(config)# end
Returns to privileged EXEC mode.
(Optional) Configuring the Maximum-Aging Time for a VLAN
To configure the maximum-aging time for a VLAN, perform this procedure:
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
spanning-tree vlanvlan-idmax-ageseconds
Example:
Device(config)# spanning-tree vlan 20 max-age 30
Configures the maximum-aging time of a VLAN. The maximum-aging time is the number of seconds a switch waits without receiving
spanning-tree configuration messages before attempting a reconfiguration.
For vlan-id, you can specify a single VLAN identified by VLAN ID number, a range of VLANs separated by a hyphen, or a series of VLANs
separated by a comma. The range is 1 to 4094.
For seconds, the range is 6 to 40; the default is 20.
Step 4
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
(Optional) Configuring the Transmit Hold-Count
You can configure the BPDU burst size by changing the transmit hold count value.
Note
Changing this parameter to a higher value can have a significant impact on CPU utilization, especially in Rapid PVST+ mode.
Lowering this value can slow down convergence in certain scenarios. We recommend that you maintain the default setting.
To configure the transmit hold-count, perform this procedure:
Configures the number of BPDUs that can be sent before pausing for 1 second.
For value, the range is 1 to 20; the default is 6.
Step 4
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
Monitoring Spanning Tree Protocol Configuration Status
Table 4. Commands for Displaying STP Configuration Status
show spanning-tree active
Displays STP configuration information on active interfaces only.
show spanning-tree detail
Displays a detailed summary of interface information.
show spanning-treevlanvlan-id
Displays STP configuration information for the specified VLAN.
show spanning-treeinterfaceinterface-id
Displays STP configuration information for the specified interface.
show spanning-treeinterfaceinterface-idportfast
Displays STP portfast information for the specified interface.
show spanning-treesummary [totals]
Displays a summary of interface states or displays the total lines of the STP state section.
To clear STP counters, use the clear spanning-tree [interface interface-id] privileged EXEC command.
Additional References for Spanning Tree Protocol
Related Documents
Related Topic
Document Title
For complete syntax and usage information for the commands used in this chapter.
See the Layer 2/3 Commands section of the Command Reference (Catalyst 9500 Series Switches)
Feature History for Spanning Tree Protocol
This table provides release and related information for features explained in this module.
These features are available on all releases subsequent to the one they were introduced in, unless noted otherwise.
Table 5. New Feature History
Release
Feature
Feature Information
Cisco IOS XE Everest 16.5.1a
Spanning Tree Protocol
STP is a Layer 2 link management protocol that provides path redundancy while preventing loops in the network.
Support for this feature was introduced only on the C9500-12Q, C9500-16X, C9500-24Q, C9500-40X models of the Cisco Catalyst
9500 Series Switches.
Cisco IOS XE Fuji 16.8.1a
Spanning Tree Protocol
Support for this feature was introduced only on the C9500-32C, C9500-32QC, C9500-48Y4C, and C9500-24Y4C models of the Cisco
Catalyst 9500 Series Switches.
Cisco IOS XE Gibraltar 16.11.1
Spanning Tree Instances
The number of supported spanning tree instances was increased to 256.
Support for this feature was introduced only on the C9500-12Q, C9500-16X, C9500-24Q, C9500-40X models of the Cisco Catalyst
9500 Series Switches.
Cisco IOS XE Amsterdam 17.2.1
Spanning Tree Instances
The number of supported spanning tree instances was increased to 300.
Support for this feature was introduced only on the C9500-12Q, C9500-16X, C9500-24Q, C9500-40X models of the Cisco Catalyst
9500 Series Switches.
Use Cisco Feature Navigator to find information about platform and software image support. To access Cisco Feature Navigator,
go to http://www.cisco.com/go/cfn.
