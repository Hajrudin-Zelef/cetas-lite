---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066-6
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066"
domain: cisco
role: reference
task: reference
actors: ["California"]
dates: []
keywords: ["copyright", "disclosure", "ethernet", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066.md
source_anchor: ""
source_lines: [462, 608]
sha256: 0c5fa8ebf78763333e310fa53d05a75453e7e24d34f3d58962665a6b9f534a39
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066

(Optional) Specifies the Cisco StackWise Virtual domain ID.
The domain ID range is from 1 to 255. The default value is one.
Step 7
end
Example:
Device(config-stackwise-virtual)# end
Returns to privileged EXEC mode.
Step 8
show stackwise-virtual
Example:
Device# show stackwise-virtual
Step 9
write memory
Example:
Device# write memory
Saves the running-configuration which resides in the system RAM and updates the ROMmon variables. If you do not save the changes,
the changes will no longer be part of the startup configuration when the switch reloads. Note that the configurations for
stackwise-virtual and domain are saved to the running-configuration and the startup-configuration after the reload.
Note
On the Cisco Catalyst 9500X Series Switches, the system database is updated instead of ROMMON variables.
Step 10
reload
Example:
Device# reload
Restarts the switch and forms the stack.
Configuring Cisco StackWise Virtual Link
Note
Depending on the switch model, SVL is supported on all 10G interfaces and 40G interfaces of the Cisco Catalyst 9500 Series
switches and on all the 400G, 100G, 40G, 25G and 10G interfaces of the Cisco Catalyst 9500 Series high performance switches.
However, a combination of different interface speeds is not supported.
Dynamic addition and removal of SVL links are supported on the Cisco Catalyst 9500X Series Switches, and therefore, a reload is not required for adding or removing the SVL links when the device is already operating in SVL
mode.
If a user tries to remove the last active SVL link, they will be notified of a stack split through a syslog message.
To configure a switch port as an SVL port, perform the following procedure on both the switches:
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
Perform one the of the following actions depending on the switch that you are configuring.
If you are configuring a Cisco Catalyst 9500 Series Switch, use interface { TenGigabitEthernet| FortyGigabitEthernet} <interface>
If you are configuring a Cisco Catalyst 9500 Series high-performance Switch, use interface { HundredGigE | FortyGigabitEthernet | TwentyFiveGigE} <interface>
If you are configuring a Cisco Catalyst 9500X-28C8D switch, use interface { HundredGigE | FourHundredGigE } <interface>
If you are configuring a Cisco Catalyst 9500X-60L4D switch, use interface { FiftyGigE | FourHundredGigE } <interface>
Example:
Device(config)# interface TenGigabitEthernet1/0/2
Enters Ethernet interface configuration mode.
Step 4
stackwise-virtual link link value
Example:
Device(config-if)# stackwise-virtual link 1
Associates the interface with configured SVL.
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
the changes will no longer be part of the startup configuration when the switch reloads. Note that the configuration for stackwise-virtual link link value is saved only in the running-configuration and not the startup-configuration.
Note
On the Cisco Catalyst 9500X Series Switches, the system database is updated instead of ROMMON variables.
Step 7
reload
Example:
Device# reload
Restarts the switch.
Note: When converting a Cisco Catalyst 9500 Series High Performance switch from standalone mode to SVL mode for the first time,
one of the switches boots up or resets, for resolving the switch number conflict and sets the SWITCH_NUMBER environment variable
to 2. The following message appears at the console prompt indicating this:
Waiting for remote chassis to join
#######################################################################
Chassis number is 2
All chassis in the stack have been discovered. Accelerating discovery
Chassis is reloading, reason: Configured Switch num conflicts with peer,
Changing local switch number to 2 and reloading to take effect
On the Cisco Catalyst 9500X Series Switches, the following message appears on the console prompt:
%CLUSTERMGR-1-RELOAD: B0/0: clustermgr: Reloading due to reason Chassis is reloading;
switch num conflicts with peer, changing local switch number to 2 and reloading to take effect
%PMAN-5-EXITACTION: B0/0: pvp: Process manager is exiting: process exit with reload fru code
%PMAN-5-EXITACTION: R0/0: pvp: Process manager is exiting: reload fru action requested
While restarting the C9500X-28C8D model of Cisco Catalyst 9500 Series Switches to enable the SVL, an initial prompt corresponding to the IOSd-BP is seen:
** Note ** Please note the MAC address in this prompt corresponds to the BIA on the GigabitEthernet 0/0 which is the Management interface on the switch
This prompt is a transitionary prompt during the bringup phase before the IOS-XE prompt appears on the console,please do not issue any show commands on this prompt
None of the standard show commands are supported on IOSd-BP prompt.
Once the LACP and ISIS protocol negotiations are completed for the SVL and DAD links, IOSd-RP prompt will be launched which is the regular IOS-XE CLI prompt
Initializing Hardware......
System Bootstrap, Version 17.11.1[Int-Alpha], RELEASE SOFTWARE (P)
Compiled Fri Aug 26 10:18:40 2022 by rel
Current ROMMON image : Primary Rommon Image
Last reset cause:CPU Reset
C9600X-SUP-2 platform with 33554432 Kbytes of main memory
Preparing to autoboot. [Press Ctrl-C to interrupt] 0
boot: attempting to boot from [bootflash:packages.conf]
boot: reading file packages.conf
<output truncated>
The following is a sample output from the bootup in SVL mode:
<<Beginning of the IOSd-BP prompt >>>>>>
*Sep 24 04:09:28.926: %LINK-3-UPDOWN: Interface CEOBC, changed state to up
sw.3C57310481C0-bp>enable
sw.3C57310481C0-bp#
sw.3C57310481C0-bp#
*Sep 24 04:09:29.926: %LINEPROTO-5-UPDOWN: Line protocol on Interface CEOBC, changed state to up
sw.3C57310481C0-bp#
sw.3C57310481C0-bp#
*Sep 24 04:09:38.960: %PKI-2-NON_AUTHORITATIVE_CLOCK: PKI functions can not be initialized until an authoritative time source, like NTP, can be obtained.
sw.3C57310481C0-bp#
sw.3C57310481C0-bp#
*Sep 24 04:09:51.642: %SPA_OIR-6-ONLINECARD: SPA (C9600-LC-48YL) online in subslot 2/0
*Sep 24 04:09:52.068: %SPA_OIR-6-ONLINECARD: SPA (C9600-LC-24C) online in subslot 5/0
sw.3C57310481C0-bp#
sw.3C57310481C0-bp#
*Sep 24 04:10:12.810: %SPA_OIR-6-ONLINECARD: SPA (C9600-LC-48TX) online in subslot 6/0
sw.3C57310481C0-bp#
*Sep 24 04:10:18.494: %LINK-3-UPDOWN: Interface HundredGigE5/0/17, changed state to up
*Sep 24 04:10:20.512: %LINK-3-UPDOWN: Interface Port-channel241, changed state to up
*Sep 24 04:10:21.512: %LINEPROTO-5-UPDOWN: Line protocol on Interface HundredGigE5/0/17, changed state to up
*Sep 24 04:10:21.512: %LINEPROTO-5-UPDOWN: Line protocol on Interface Port-channel241, changed state to up
sw.3C57310481C0-bp#
*Sep 24 04:10:23.515: %CLNS-5-ADJCHANGE: ISIS: Adjacency to 0490.0415.8000 (Port-channel241) Up, new adjacency
*Sep 24 04:10:30.611: %CLNS-5-ADJCHANGE: ISIS: Adjacency to sw.3C5731049E00 (Port-channel241) Up, new adjacency
Restricted Rights Legend
Use, duplication, or disclosure by the Government is
subject to restrictions as set forth in subparagraph
(c) of the Commercial Computer Software - Restricted
Rights clause at FAR sec. 52.227-19 and subparagraph
(c) (1) (ii) of the Rights in Technical Data and Computer
Software clause at DFARS sec. 252.227-7013.
Cisco Systems, Inc.
170 West Tasman Drive
San Jose, California 95134-1706
Cisco IOS Software [Dublin], Catalyst L3 Switch Software (CAT9K_IOSXE), Experimental Version 17.10.20220921:075136 [BLD_V1710_THROTTLE_LATEST_20220921_071938:/nobackup/mcpre/s2c-build-ws 101]
Copyright (c) 1986-2022 by Cisco Systems, Inc.
Compiled Wed 21-Sep-22 00:52 by mcpre
