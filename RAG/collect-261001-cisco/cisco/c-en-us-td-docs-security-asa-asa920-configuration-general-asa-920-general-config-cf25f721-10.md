---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721-10
title: "c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["preemption"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721.md
source_anchor: ""
source_lines: [717, 735]
sha256: 531c8208512a56375734394c4fce8d1a11530754d4defe4c904dc5edb080456e
---

# c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721

                                          A preemption for the failover group is configured, which causes the failover group to automatically become active on the preferred unit when the unit becomes available.
- 
                                          				
                                          
Failover Events
In an Active/Active failover configuration, failover occurs on a failover group basis, not a system basis. For example, if you designate both failover groups as Active on the primary unit, and failover group 1 fails, then failover group 2 remains Active on the primary unit while failover group 1 becomes active on the secondary unit.
Because a failover group can contain multiple contexts, and each context can contain multiple interfaces, it is possible for all interfaces in a single context to fail without causing the associated failover group to fail.
The following table shows the failover action for each failure event. For each failure event, the policy (whether or not failover occurs), actions for the active failover group, and actions for the standby failover group are given.
| Table 5. Failover Events |  |  |  |  | 
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
