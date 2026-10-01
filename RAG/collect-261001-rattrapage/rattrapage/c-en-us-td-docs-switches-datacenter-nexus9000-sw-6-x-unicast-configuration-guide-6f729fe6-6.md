---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6-6
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "preemption"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6.md
source_anchor: ""
source_lines: [247, 298]
sha256: 821b1f49b37625e484109c4fad6c15981726f174d425ddcc0071dde86a2e12a6
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6

This example shows how to configure HSRP object tracking on Ethernet interface 1/2:
switch(config-track)# track 2 ip route 192.0.2.0/8 reachability
Configuring the HSRP Priority
You can configure the priority of an HSRP group. HSRP uses the priority to determine which HSRP group member acts as the active router. If you configure HSRP on a vPC-enabled interface, you can optionally configure the upper and lower threshold values to control when to fail over to the vPC trunk. If the standby router priority falls below the lower threshold, HSRP sends all standby router traffic across the vPC trunk to forward through the active HSRP router. HSRP maintains this scenario until the standby HSRP router priority increases above the upper threshold.
For IPv6 HSRP groups, if all group members have the same priority, HSRP selects the active router based on the IPv6 link-local address.
To configure the HSRP priority, use the following command in the HSRP group configuration mode:
|  |  | 
|---|---|
| priority level [ forwarding-threshold lower lower-value upper upper-value]  Example: switch(config-if-hsrp)# priority 60 forwarding-threshold lower 40 upper 50 | Sets the priority level used to select the active router in an HSRP group. The level range is from 0 to 255. The default is 100. Optionally, this command sets the upper and lower threshold values used by vPC to determine when to fail over to the vPC trunk. The lower-value range is from 1 to 255. The default is 1. The upper-value range is from 1 to 255. The default is 255. | 
Customizing HSRP
You can optionally customize the behavior of HSRP. Be aware that as soon as you enable an HSRP group by configuring a virtual IP address, that group is now operational. If you first enable an HSRP group before customizing HSRP, the router could take control over the group and become the active router before you finish customizing the feature. If you plan to customize HSRP, you should do so before you enable the HSRP group. To customize HSRP, use the following commands in HSRP configuration mode:
|  |  | 
|---|---|
| name string  Example: switch(config-if-hsrp)# name HSRP-1 | Specifies the IP redundancy name for an HSRP group. The string is from 1 to 255 characters. The default string has the following format: hsrp-interface short-name group-id. For example, hsrp-Eth2/1-1. | 
| preempt [ delay [ minimum seconds ] [ reload seconds ] [ sync seconds ]]  Example: switch(config-if-hsrp)# preempt delay minimum 60 | Configures the router to take over as an active router for an HSRP group if it has a higher priority than the current active router. This command is disabled by default. Optionally, a delay can be configured that delays the HSRP group preemption by the configured time. The range is from 0 to 3600 seconds. | 
| timers [ msec ] hellotime [ msec ] holdtime  Example: switch(config-if-hsrp)# timers 5 18 | Configures the hello and hold time for this HSRP member as follows:  The optional msec keyword specifies that the argument is expressed in milliseconds instead of the default seconds. The timer ranges for milliseconds are as follows: | 
To customize HSRP, use the following commands in interface configuration mode:
|  |  | 
|---|---|
| hsrp delay minimum seconds  Example: switch(config-if)# hsrp delay minimum 30 | Specifies the minimum amount of time that HSRP waits after a group is enabled before participating in the group. The range is from 0 to 10000 seconds. The default is 0. | 
| hsrp delay reload seconds  Example: switch(config-if)# hsrp delay reload 30 | Specifies the minimum amount of time that HSRP waits after a reload and before participating in the group. The range is from 0 to 10000 seconds. The default is 0. | 
Configuring Extended Hold Timers for HSRP
You can configure HSRP to use extended hold timers to support extended NSF during a controlled (graceful) switchover. You should configure extended hold timers on all HSRP routers (see the “High Availability and Extended Nonstop Forwarding” section).
Note You must configure extended hold timers on all HSRP routers if you configure extended hold timers. If you configure a nondefault hold timer, you should configure the same value on all HSRP routers when you configure HSRP extended hold timers.
Note HSRP extended hold timers are not applied if you configure millisecond hello and hold timers for HSRPv1. This statement does not apply to HSRPv2.
To configure HSRP extended hold timers, use the following command in global configuration mode:
Use the show hsrp command or the show running-config hsrp command to display the extended hold time.
Verifying the HSRP Configuration
To display HSRP configuration information, perform one of the following tasks:
|  |  | 
|---|---|
| show hsrp [group group-number] | Displays the HSRP status for all groups or one group. | 
| show hsrp delay [interface interface-type slot/port] | Displays the HSRP delay value for all interfaces or one interface. | 
| show hsrp [interface interface-type slot/port] | Displays the HSRP status for an interface. | 
| show hsrp [group group-number] [interface interface-type slot/port] [active] [all] [init] [learn] [listen] [speak] [standby] | Displays the HSRP status for a group or interface for virtual forwarders in the active, init, learn, listen, or standby state. Use the all keyword to see all states, including disabled. | 
| show hsrp [group group-number] [interface interface-type slot/port] active] [all] [init] [learn] [listen] [speak] [standby] brief | Displays a brief summary of the HSRP status for a group or interface for virtual forwarders in the active, init, learn, listen, or standby state. Use the all keyword to see all states, including disabled. | 
Configuration Examples for HSRP
This example shows how to enable HSRP on an interface with MD5 authentication and interface tracking:
track 2 interface ethernet 2/2 ip
authenticate md5 key-chain hsrp-keys
This example shows how to configure the HSRP priority on an interface:
Additional References
For additional information related to implementing HSRP, see the following sections:
Related Documents
|  |  | 
|---|---|
| Configuring the Virtual Router Redundancy Protocol | Chapter 18, “Configuring VRRP” | 
| Configuring high availability | Cisco Nexus 9000 Series NX-OS High Availability and Redundancy Guide | 
MIBs
|  |  | 
|---|---|
| MIBs related to HSRP | To locate and download supported MIBs, go to the following URL: ftp://ftp.cisco.com/pub/mibs/supportlists/nexus9000/Nexus9000MIBSupportList.html |
