---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf-8
title: "c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["preemption"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf.md
source_anchor: ""
source_lines: [501, 539]
sha256: 50293ea8e579c45f7a591af39f833529dfdee5dcf95813cf82b6bafdfd88dbe5
---

# c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf

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
                                          				
                                          A preemption for the failover group is configured, which causes the failover group to automatically become active on the preferred unit when the unit becomes available.
- 
                                          				
                                          
Failover Events
In an Active/Active failover configuration, failover occurs on a failover group basis, not a system basis. For example, if you designate both failover groups as Active on the primary unit, and failover group 1 fails, then failover group 2 remains Active on the primary unit while failover group 1 becomes active on the secondary unit.
Because a failover group can contain multiple contexts, and each context can contain multiple interfaces, it is possible for all interfaces in a single context to fail without causing the associated failover group to fail.
The following table shows the failover action for each failure event. For each failure event, the policy (whether or not failover occurs), actions for the active failover group, and actions for the standby failover group are given.
| Table 4. Failover Events |  |  |  |  | 
|---|---|---|---|---|
| Failure Event | Policy | Active Group Action | Standby Group Action | Notes | 
|---|---|---|---|---|
| A unit experiences a power or software failure | Failover | Become standby Mark as failed | Become active Mark active as failed | When a unit in a failover pair fails, any active failover groups on that unit are marked as failed and become active on the peer unit. | 
| Interface failure on active failover group above threshold | Failover | Mark active group as failed | Become active | None. | 
| Interface failure on standby failover group above threshold | No failover | No action | Mark standby group as failed | When the standby failover group is marked as failed, the active failover group does not attempt to fail over, even if the interface failure threshold is surpassed. | 
| Formerly active failover group recovers | No failover | No action | No action | Unless failover group preemption is configured, the failover groups remain active on their current unit. | 
| Failover link failed at startup | No failover | Become active | Become active | If the failover link is down at startup, both failover groups on both units become active. | 
| State link failed | No failover | No action | No action | State information becomes out of date, and sessions are terminated if a failover occurs. | 
| Failover link failed during operation | No failover | n/a | n/a | Each unit marks the failover link as failed. You should restore the failover link as soon as possible because the unit cannot fail over to the standby unit while the failover link is down. |
