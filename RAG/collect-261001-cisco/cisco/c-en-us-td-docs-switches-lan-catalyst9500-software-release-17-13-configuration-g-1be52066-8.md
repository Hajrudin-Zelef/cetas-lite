---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066-8
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066.md
source_anchor: ""
source_lines: [804, 950]
sha256: 2e8887af57f8b534279f6465e0ce728ea910bf9de3c9cf708e200f549af7f3d7
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066

Enter your password, if prompted.
Step 2
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
stackwise-virtual
Example:
Device(config)# stackwise-virtual
Enables Cisco StackWise Virtual and enters stackwise-virtual mode.
Note that the configuration for dual-activerecovery-reload-disable is saved only in the running-configuration and not the startup-configuration.
Step 5
end
Example:
Device(config-stackwise-virtual)# end
Returns to privileged EXEC mode.
Disabling Cisco StackWise Virtual
To disable Cisco StackWise Virtual on a switch, perform the following procedure:
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
The following is a sample configuration for configuring SVL on a switch.
On Switch 1:
Device>enable
Device#configure terminal
Device(config)#interface FortyGigabitEthernet1/1/1
Device(config-if)#stackwise-virtual link 1
WARNING: All the extraneous configurations will be removed for FortyGigabitEthernet1/1/1 on reboot
INFO: Upon reboot, the config will be part of running config but not part of start up config.
Device(config-if)#end
Device#write memory
Device#reload
On Switch 2:
Device>enable
Device#configure terminal
Device(config)#interface FortyGigabitEthernet1/1/1
Device(config-if)#stackwise-virtual link 1
WARNING: All the extraneous configurations will be removed for FortyGigabitEthernet1/1/1 on reboot
INFO: Upon reboot, the config will be part of running config but not part of start up config.
Device(config-if)#end
Device#write memory
Device#reload
The following example show how to configure SVL on the C9500X-28C8D model of Cisco Catalyst 9500X Series Switches:
Device> enable
Device# configure terminal
Device(config)# interface FiftyGigabitEthernet1/0/5
Device(config-if)# stackwise-virtual link 1
WARNING: AUTO-LC-SHUTDOWN is disabled for chassis 1. Suggested to be enabled and configured with highest priority for line cards with stackwise virtual links
WARNING: All the extraneous configurations will be removed for FiftyGigE1/0/5
*Sep 29 09:33:25.572: %SYS-5-CONFIG_P: Configured programmatically by process Exec from console as console
Example: Configuring Secure StackWise Virtual
The following is a sample configuration for configuring Secure StackWise Virtual.
Example: Displaying Secure StackWise Virtual Authorization Key and Status
The following is an example displaying the Secure StackWise Virtual authorization key.
Device# show secure-stackwise-virtual authorization-key
Secure-stackwise-virtual: Stored key (16) : 12345678901234567890123456789012
The following is an example displaying the Secure StackWise Virtual authorization key status.
Device# show secure-stackwise-virtual status
Switch is running in SECURE-SVL mode
Example: Disabling Secure StackWise Virtual
The following is an example of Secure StackWise Virtual authorization key zeroization.
Device(config)# secure-stackwise-virtual zeroize sha1-key
**Critical Warning** - This command is irreversible
and will zeroize the Secure-SVL-VPK by Deleting
the IOS image and config files, please use extreme
caution and confirm with Yes on each of three
iterations to complete. The system will reboot
after the command executes successfully
Do you want to proceed ?? (yes/[no]):
Example: Configuring StackWise Virtual Fast Hello Dual-Active-Detection Link
The following example shows how to configure a StackWise Virtual Fast Hello dual-active detection link on a Switch 1 and Switch
2. You cannot configure StackWise Virtual Fast Hello dual-active-detection links on ports that are already configured as StackWise
Virtual link ports:
On Switch 1:
Device>enable
Device#configure terminal
Device(config)#interface FortyGigabitEthernet1/0/3
Device(config-if)#stackwise-virtual dual-active-detection
Please reload the switch for Stackwise Virtual configuration to take effect
Upon reboot, the config will be part of running config but not part of start up config.
Device(config-if)#exit
On Switch 2:
Device(config)#interface FortyGigabitEthernet1/0/3
Device(config-if)#stackwise-virtual dual-active-detection
Please reload the switch for Stackwise Virtual configuration to take effect
Upon reboot, the config will be part of running config but not part of start up config.
Device(config-if)#end
On both the switches:
Device#write memory
Device#reload
The following is a sample configuration for configuring a StackWise Virtual Fast Hello dual-active-detection link on the C9500X-28C8D model of Cisco Catalyst 9500X Series Switches.
Device> enable
Device# configure terminal
Device(config)# interface FiftyGigabitEthernet2/5/0/41
Device(config-if)# stackwise-virtual dual-active-detection
WARNING: All the extraneous configurations will be removed for FiftyGigE2/5/0/41.
*Sep 29 09:38:01.035: %SYS-5-CONFIG_P: Configured programmatically by process Exec from console as console
*Sep 29 09:38:01.063: %LINEPROTO-5-UPDOWN: Line protocol on Interface FiftyGigE2/5/0/41, changed state to down
Device(config-if)#
*Sep 29 09:38:02.067: %LINEPROTO-5-UPDOWN: Line protocol on Interface FiftyGigE1/6/0/41, changed state to down
Device(config-if)#
Device(config-if)#
*Sep 29 09:38:03.067: %LINK-3-UPDOWN: Interface FiftyGigE1/6/0/41, changed state to down
*Sep 29 09:38:03.080: %LINK-3-UPDOWN: Interface FiftyGigE2/5/0/41, changed state to downint fif
*Sep 29 09:38:07.544: %CLUSTERMGR-6-DUAL_ACTIVE_CFG_MSG: Chassis 2 B0/0: clustermgr: Dual Active Detection link is available now1/6/0
*Sep 29 09:38:09.525: %LINK-3-UPDOWN: Interface FiftyGigE1/6/0/41, changed state to up
*Sep 29 09:38:09.544: %LINK-3-UPDOWN: Interface FiftyGigE2/5/0/41, changed state to up
*Sep 29 09:38:10.525: %LINEPROTO-5-UPDOWN: Line protocol on Interface FiftyGigE1/6/0/41, changed state to up/41
Device(config-if)#
Device(config-if)#
Device(config-if)#
*Sep 29 09:38:10.544: %LINEPROTO-5-UPDOWN: Line protocol on Interface FiftyGigE2/5/0/41, changed state to upstackwise-virtual dual-active-detection
Device(config-if)#stackwise-virtual dual-active-detection
WARNING: All the extraneous configurations will be removed for FiftyGigE1/6/0/41.
*Sep 29 09:38:14.108: %SYS-5-CONFIG_P: Configured programmatically by process Exec from console as console
*Sep 29 09:38:14.141: %LINEPROTO-5-UPDOWN: Line protocol on Interface FiftyGigE1/6/0/41, changed state to down
Device(config-if)#
*Sep 29 09:38:14.144: %CLUSTERMGR-6-DUAL_ACTIVE_CFG_MSG: Chassis 2 B0/0: clustermgr: Dual Active Detection links are not available anymore
Device#
*Sep 29 09:38:16.213: %LINK-3-UPDOWN: Interface FiftyGigE1/6/0/41, changed state to down
*Sep 29 09:38:16.216: %LINK-3-UPDOWN: Interface FiftyGigE2/5/0/41, changed state to down
*Sep 29 09:38:17.206: %SYS-5-CONFIG_I: Configured from console by console
*Sep 29 09:38:17.217: %LINEPROTO-5-UPDOWN: Line protocol on Interface FiftyGigE2/5/0/41, changed state to down
*Sep 29 09:38:17.254: %CLUSTERMGR-6-DUAL_ACTIVE_CFG_MSG: Chassis 2 B0/0: clustermgr: Dual Active Detection link is available now
*Sep 29 09:38:17.255: %CLUSTERMGR-6-DUAL_ACTIVE_CFG_MSG: Chassis 1 B0/0: clustermgr: Dual Active Detection link is available now
Device#
*Sep 29 09:38:19.252: %LINK-3-UPDOWN: Interface FiftyGigE2/5/0/41, changed state to up
*Sep 29 09:38:19.256: %LINK-3-UPDOWN: Interface FiftyGigE1/6/0/41, changed state to up
*Sep 29 09:38:20.252: %LINEPROTO-5-UPDOWN: Line protocol on Interface FiftyGigE2/5/0/41, changed state to up
Device#
Device#
