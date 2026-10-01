---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721-9
title: "c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["preemption"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721.md
source_anchor: ""
source_lines: [639, 716]
sha256: 83a9f751f7c037d3d7d5bf8514ea8148a1134f90e6ce310e8286a689143e4292
---

# c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721

                                       establish traffic flow responsibilities between paired units in Active/Standby failover configurations.
Primary and secondary unit differences
The main differences between the two units in a failover pair are related to which unit is active and which unit is standby, namely which IP addresses to use and which unit actively passes traffic.
However, a few differences exist between the units depending on which unit is primary (as specified in the configuration) and which unit is secondary:
- 
                                       
                                       The primary unit always becomes the active unit if both units start up at the same time (and are of equal operational health).
- 
                                       
                                       The primary unit MAC addresses are always coupled with the active IP addresses. The exception to this rule occurs when the secondary unit becomes active and cannot obtain the primary unit MAC addresses over the failover link. In this case, the secondary unit MAC addresses are used.
Active unit determination at startup
Active unit determination at startup is a high availability process that
- 
                                       
                                       sets a unit to standby if it boots and detects a peer already running as active,
- 
                                       
                                       sets a unit to active if it boots and does not detect a peer, and
- 
                                       
                                       sets the primary unit as active and secondary unit as standby if both units boot simultaneously.
Failover events
A failover event is a system condition that
- 
                                       
                                       triggers an active unit to transfer control to a standby unit in Active/Standby failover configurations
- 
                                       
                                       occurs on a unit basis and cannot fail over individual or groups of contexts even on systems running in multiple context mode, and
- 
                                       
                                       follows specific failover policies that determine whether failover occurs and what actions each unit takes.
Failover event policies and actions
This table shows the failover action for each failure event. For each failure event, the table shows the failover policy (failover or no failover), the action taken by the active unit, the action taken by the standby unit, and any special notes about the failover condition and actions.
| Table 4. Failover events |  |  |  |  | 
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
                                    		  
                                    The primary unit provides the running configuration to the pair when they boot simultaneously.
-  
                                    		  
                                    Each failover group in the configuration is configured with a primary or secondary unit preference. When used with preemption, this preference ensures that the failover group runs on the correct unit after it starts up. Without preemption, both groups run on the first unit to boot up.
Active Unit Determination for Failover Groups at Startup
The unit on which a failover group becomes active is determined as follows:
- 
                                    		  
                                    When a unit boots while the peer unit is not available, both failover groups become active on the unit.
- 
                                    		  
                                    When a unit boots while the peer unit is active (with both failover groups in the active state), the failover groups remain in the active state on the active unit regardless of the primary or secondary preference of the failover group until one of the following occurs: 
  - 
                                          				
                                          A failover occurs.
  - 
                                          				
                                          A failover is manually forced.
  - 
                                          				
