---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf-7
title: "c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf.md
source_anchor: ""
source_lines: [411, 500]
sha256: a971c6570fbbe882947bf73b98281b5efda4eb423ab78a059d389ebd0ef81b99
---

# c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf

- 
                                    Local Certificate Authorities (CAs)
- 
                                    ASA images
- 
                                    ASDM images
Command Replication
After startup, commands that you enter on the active unit are immediately replicated on the standby unit. You do not have to save the active configuration to flash memory to replicate the commands.
In Active/Active failover, commands entered in the system execution space are replicated from the unit on which failover group 1 is in the active state.
Failure to enter the commands on the appropriate unit for command replication to occur causes the configurations to be out of synchronization. Those changes may be lost the next time the initial configuration synchronization occurs.
The following commands are replicated to the standby ASA:
- 
                                    All configuration commands except for mode, firewall , and failover lan unit
- 
                                    copy running-config startup-config
- 
                                    delete
- 
                                    mkdir
- 
                                    rename
- 
                                    rmdir
- 
                                    write memory
The following commands are not replicated to the standby ASA:
- 
                                    All forms of the copy command except for copy running-config startup-config
- 
                                    All forms of the write command except for write memory
- 
                                    debug
- 
                                    failover lan unit
- 
                                    firewall
- 
                                    show
- 
                                    terminal pager and pager
About Active/Standby Failover
Active/Standby failover lets you use a standby ASA to take over the functionality of a failed unit. When the active unit fails, the standby unit becomes the active unit.
| Note | For multiple context mode, the ASA can fail over the entire unit (including all contexts) but cannot fail over individual contexts separately. | 
Primary/Secondary Roles and Active/Standby Status
The main differences between the two units in a failover pair are related to which unit is active and which unit is standby, namely which IP addresses to use and which unit actively passes traffic.
However, a few differences exist between the units based on which unit is primary (as specified in the configuration) and which unit is secondary:
-  
                                    		  
                                    The primary unit always becomes the active unit if both units start up at the same time (and are of equal operational health).
-  
                                    		  
                                    The primary unit MAC addresses are always coupled with the active IP addresses. The exception to this rule occurs when the secondary unit becomes active and cannot obtain the primary unit MAC addresses over the failover link. In this case, the secondary unit MAC addresses are used.
Active Unit Determination at Startup
The active unit is determined by the following:
- 
                                    		  
                                    If a unit boots and detects a peer already running as active, it becomes the standby unit.
- 
                                    		  
                                    If a unit boots and does not detect a peer, it becomes the active unit.
- 
                                    		  
                                    If both units boot simultaneously, then the primary unit becomes the active unit, and the secondary unit becomes the standby unit.
Failover Events
In Active/Standby failover, failover occurs on a unit basis. Even on systems running in multiple context mode, you cannot fail over individual or groups of contexts.
The following table shows the failover action for each failure event. For each failure event, the table shows the failover policy (failover or no failover), the action taken by the active unit, the action taken by the standby unit, and any special notes about the failover condition and actions.
| Table 3. Failover Events |  |  |  |  | 
|---|---|---|---|---|
| Failure Event | Policy | Active Unit Action | Standby Unit Action | Notes | 
|---|---|---|---|---|
| Active unit failed (power or hardware) | Failover | n/a | Become active Mark active as failed | No hello messages are received on any monitored interface or the failover link. | 
| Formerly active unit recovers | No failover | Become standby | No action | None. | 
| Standby unit failed (power or hardware) | No failover | Mark standby as failed | n/a | When the standby unit is marked as failed, then the active unit does not attempt to fail over, even if the interface failure threshold is surpassed. | 
| Failover link failed during operation | No failover | Mark failover link as failed | Mark failover link as failed | You should restore the failover link as soon as possible because the unit cannot fail over to the standby unit while the failover link is down. | 
| Failover link failed at startup | No failover | Become active Mark failover link as failed | Become active Mark failover link as failed | If the failover link is down at startup, both units become active. | 
| State link failed | No failover | No action | No action | State information becomes out of date, and sessions are terminated if a failover occurs. | 
| Interface failure on active unit above threshold | Failover | Mark active as failed | Become active | None. | 
| Interface failure on standby unit above threshold | No failover | No action | Mark standby as failed | When the standby unit is marked as failed, then the active unit does not attempt to fail over even if the interface failure threshold is surpassed. | 
About Active/Active Failover
This section describes Active/Active failover.
Active/Active Failover Overview
In an Active/Active failover configuration, both ASAs can pass network traffic. Active/Active failover is only available to ASAs in multiple context mode. In Active/Active failover, you divide the security contexts on the ASA into a maximum of 2 failover groups.
A failover group is simply a logical group of one or more security contexts. You can assign failover group to be active on the primary ASA, and failover group 2 to be active on the secondary ASA. When a failover occurs, it occurs at the failover group level. For example, depending on interface failure patterns, it is possible for failover group 1 to fail over to the secondary ASA, and subsequently failover group 2 to fail over to the primary ASA. This event could occur if the interfaces in failover group 1 are down on the primary ASA but up on the secondary ASA, while the interfaces in failover group 2 are down on the secondary ASA but up on the primary ASA.
The admin context is always a member of failover group 1. Any unassigned security contexts are also members of failover group 1 by default. If you want Active/Active failover, but are otherwise uninterested in multiple contexts, the simplest configuration would be to add one additional context and assign it to failover group 2.
| Note | When configuring Active/Active failover, make sure that the combined traffic for both units is within the capacity of each unit. | 
| Note | You can assign both failover groups to one ASA if desired, but then you are not taking advantage of having two active ASAs. | 
Primary/Secondary Roles and Active/Standby Status for a Failover Group
As in Active/Standby failover, one unit in an Active/Active failover pair is designated the primary unit, and the other unit the secondary unit. Unlike Active/Standby failover, this designation does not indicate which unit becomes active when both units start simultaneously. Instead, the primary/secondary designation does two things:
-  
                                    		  
