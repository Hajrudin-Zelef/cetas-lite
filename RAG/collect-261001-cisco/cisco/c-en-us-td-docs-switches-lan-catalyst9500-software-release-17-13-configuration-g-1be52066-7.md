---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066-7
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "license", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066.md
source_anchor: ""
source_lines: [609, 803]
sha256: 001aa829b467080e554cff6c2492e759a9f901063f593964bf2e7ee4e3e5d250
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066

This software version supports only Smart Licensing as the software licensing mechanism.
PLEASE READ THE FOLLOWING TERMS CAREFULLY. INSTALLING THE LICENSE OR
LICENSE KEY PROVIDED FOR ANY CISCO SOFTWARE PRODUCT, PRODUCT FEATURE,
AND/OR SUBSEQUENTLY PROVIDED SOFTWARE FEATURES (COLLECTIVELY, THE
"SOFTWARE"), AND/OR USING SUCH SOFTWARE CONSTITUTES YOUR FULL
ACCEPTANCE OF THE FOLLOWING TERMS. YOU MUST NOT PROCEED FURTHER IF YOU
ARE NOT WILLING TO BE BOUND BY ALL THE TERMS SET FORTH HEREIN.
Your use of the Software is subject to the Cisco End User License Agreement
(EULA) and any relevant supplemental terms (SEULA) found at
http://www.cisco.com/c/en/us/about/legal/cloud-and-software/software-terms.html.
You hereby acknowledge and agree that certain Software and/or features are
licensed for a particular term, that the license to such Software and/or
features is valid only for the applicable term and that such Software and/or
features may be shut down or otherwise terminated by Cisco after expiration
of the applicable license term (e.g., 90-day trial period). Cisco reserves
the right to terminate any such Software feature electronically or by any
other means available. While Cisco may provide alerts, it is your sole
responsibility to monitor your usage of any such term Software feature to
ensure that your systems and networks are prepared for a shutdown of the
Software feature.
*Sep 24 04:10:35.378: %IOSXE_OIR-6-ONLINECARD: Card (fp) online in slot F0
*Sep 24 04:10:35.468: %SYS-5-CONFIG_P: Configured programmatically by process IOSD ipc task from console as vty2
FIPS: Crimson DB Key Check : Key Not Found, FIPS Mode Not Enabled
cisco C9606R (X86) processor (revision V01) with 6029940K/6147K bytes of memory.
Processor board ID FXS2418Q1V9
0 Virtual Ethernet interface
78 Forty/Hundred Gigabit Ethernet interfaces
136 Ten/TwentyFive/Fifty Gigabit Ethernet interfaces
2 Forty/Hundred/TwoHundred Gigabit Ethernet interfaces
4 Forty/Hundred/FourHundred Gigabit Ethernet interfaces
96 Ten Gigabit Ethernet interfaces
32768K bytes of non-volatile configuration memory.
33554432K bytes of physical memory.
11161600K bytes of Bootflash at bootflash:.
1638400K bytes of Crash Files at crashinfo:.
234430023K bytes of SATA hard disk at disk0:.
11161600K bytes of Bootflash at bootflash-2-0:.
1638400K bytes of Crash Files at crashinfo-2-0:.
234430023K bytes of SATA hard disk at disk0-2-0:.
Base Ethernet MAC Address : 3c:57:31:04:81:c0
Motherboard Assembly Number : 4DB9
Motherboard Serial Number : FXS241400M8
Model Revision Number : V02
Motherboard Revision Number : 6
Model Number : C9606R
System Serial Number : FXS2418Q1V9
<<followed by the regular IOS-XE CLI prompt>>
Configuring Secure StackWise Virtual
Before you begin
Note
Ensure that the devices are in a standalone mode.
Disable FIPS mode using the no fips authorization-key command before configuring the Secure StackWise Virtual authorization key.
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode. Enter your password, if prompted.
Configures the Secure StackWise Virtual authorization key.
Step 4
exit
Example:
Device(config)# exit
Returns to privileged EXEC mode.
Step 5
reload
Example:
Device# reload
Restarts the switch and the configuration of Secure StackWise Virtual takes effect.
Configuring BUM Traffic Optimization
To configure BUM traffic optimization globally, perform the following procedure:
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
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
svl l2bum optimization
Example:
Device(config)# svl l2bum optimization
Enables the BUM traffic optimization within StackWise Virtual setup globally. This feature is enabled by default.
Use the no form of this command to disable this feature.
Step 4
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
Step 5
show platform pm l2bum-status vlanvlan-id
Example:
Device# show platform pm l2bum-status vlan 1
Displays the number of forwarding ports in VLAN. number of physical ports count in forwarding state
Step 6
show platform software fed switch ac fss bum-opt summary
Example:
Device# show platform software fed switch ac fss bum-opt summary
Displays the final state of optimization.
Configuring StackWise Virtual Fast Hello Dual-Active-Detection Link
To configure StackWise Virtual Fast Hello DAD link, perform the following procedure. This procedure is optional.
Note
Dynamic addition and removal of DAD links are supported on Cisco Catalyst 9500X Series Switches, and therefore, a reload is not required for adding or removing the DAD links when the device is already operating in SVL
mode.
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
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
Perform one the of the following depending on the switch that you are configuring.
If you are configuring a Cisco Catalyst 9500 Series Switch, use interface { TenGigabitEthernet| FortyGigabitEthernet} <interface>
If you are configuring a Cisco Catalyst 9500 Series High Performance Switch, use interface { HundredGigE | FortyGigabitEthernet | TwentyFiveGigE} <interface>
Associates the interface with StackWise Virtual dual-active-detection.
Note
This command will not be visible on the device after the configuration, but will continue to function.
Step 5
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
Step 6
write memory
Example:
Device# write memory
Saves the running-configuration which resides in the system RAM and updates the ROMMON variables. If you do not save the changes,
the changes will no longer be part of the startup configuration when the switch reloads. Note that the configuration for stackwise-virtual dual-active-detection is saved only in the running-configuration and not the startup-configuration.
Note
On the Cisco Catalyst 9500X Series Switches, the system database is updated instead of ROMMON variables.
Step 7
reload
Example:
Device# reload
Restarts the switch and configuration takes effect.
Enabling ePAgP Dual-Active-Detection
To enable ePAgP dual-active-detection on a switch port, perform the following procedure. This procedure is optional.
Enables dual-active detection trust mode on channel-group with the configured ID.
Step 12
exit
Example:
Device(config-stackwise-virtual)# exit
Exits the StackWise-Virtual configuration mode.
Step 13
interface port-channelportchannel
Example:
Device(config)# interface port-channel 1
Configured port-channel on the switch.
Step 14
no shutdown
Example:
Device(config-if)# no shutdown
Enables the configured port-channel on the switch.
Step 15
end
Example:
Device(config-if)# end
Exits interface configuration.
Step 16
write memory
Example:
Device# write memory
Saves the running-configuration which resides in the system RAM and updates the ROMMON variables. If you do not save the changes,
the changes will no longer be part of the startup configuration when the switch reloads. Note that the configuration for dual-active detection pagp trust channel-group channel-group id is saved to the running-configuration and the startup-configuration after the reload.
Note
On the Cisco Catalyst 9500X Series Switches, the system database is updated instead of ROMMON variables.
Disabling Recovery Reload
After recovering from StackWise Virtual link failure, the switch in recovery mode performs a recovery action by automatically
reloading the switch. This is the default behaviour in the event of a link failure. In order to retain a switch in recovery
mode and prevent the switch from reloading automatically, you must perform the following steps.
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
