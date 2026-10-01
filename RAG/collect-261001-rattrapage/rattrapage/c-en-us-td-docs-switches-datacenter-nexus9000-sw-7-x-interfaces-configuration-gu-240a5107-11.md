---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107-11
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107.md
source_anchor: ""
source_lines: [671, 691]
sha256: 08a14f1e1a10f75ecf4294eb659a2d3f279bd8943852cdfdb3ec4c1fd35c66c2
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107

                                       vPC Role Election and Sticky-bit By default, the Cisco NX-OS software uses the lowest MAC address to elect the primary device. However, if the role priority is set, then the device with the lowest priority will be elected as the primary device. When the primary device is reloaded, the system comes back online and connectivity to the vPC secondary device (now the operational primary) is restored. The operational role of the secondary device (operational primary) does not change (to avoid unnecessary disruptions). This behavior is achieved with a sticky-bit, where the sticky information is not saved in the startup configuration. This method makes the device that is up and running win over the reloaded device. Hence, the vPC primary becomes the vPC operational secondary. Sticky-bit is also set when a vPC node comes up with vPC Peer-Link and peer-keepalive down and it becomes primary after the auto recovery period.
- 
                                       						
                                       vPC Delay Restore The delay restore timer is used to delay the vPC from coming up on the restored vPC peer device after a reload when the peer adjacency is already established. To delay the VLAN interfaces on the restored vPC peer device from coming up, use the interfaces-vlan option of the delay restore command.
- 
                                       						
                                       vPC Auto-Recovery During a data center power outage when both vPC peer switches go down, if only one switch is restored, the auto-recovery feature allows that switch to assume the role of the primary switch and the vPC links come up after the auto-recovery time period. The default auto-recovery period is 240 seconds.
| Table 1. Migration Steps and Expected Behavior |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
|  | Migration Step | Expected Behavior | Node1 Configured Role (Ex: role priority 100) | Node1 Operational Role | Node2 Configured Role (Ex: role priority 200) | Node2 Operational Role | 
|---|---|---|---|---|---|---|
| 1 | Initial state | Traffic is forwarded by both vPC peers – Node1 and Node2. Node1 is primary and Node2 is secondary. | primary | Primary Sticky bit: False | secondary | Secondary Sticky bit: False | 
| 2 | Node2 replacement – Shut all vPCs and uplinks on Node2. vPC Peer-Link and vPC peer-keepalive are in administrative up state. | Traffic converged on Primary vPC peer Node1. | primary | Primary Sticky bit: False | secondary | Secondary Sticky bit: False | 
| 3 | Remove Node2. | Node1 will continue to forward traffic. | primary | Primary Sticky bit: False | n/a | n/a | 
| 4 | Configure New_Node2. Copy the configuration to startup config. vPC Peer-Link and peer-keepalive in administrative up state. Power off New_Node2. Make all connections. Power on New_Node2. | New_Node2 will come up as secondary. Node1 continue to be primary. Traffic will continue to be forwarded on Node1. | primary | Primary Sticky bit: False | secondary | Secondary Sticky bit: False | 
| 5 | Bring up all vPCs and uplink ports on New_Node2. | Traffic will be forwarded by both Node1 and New_Node2. | primary | Primary Sticky bit: False | secondary | Secondary Sticky bit: False | 
| 6 | Node1 replacement - Shut vPCs and uplinks on Node1. | Traffic will converge on New_Node2. | primary | Primary Sticky bit: False | secondary | Secondary Sticky bit: False | 
| 7 | Remove Node1. | New_Node2 will become secondary, operational primary and sticky bit will be set to True. | n/a | n/a | secondary | Primary Sticky bit: True | 
| 8 | Configure New_Node1. Copy running to startup. Power off the new Node1. Make all connections. Power on New_Node1. | New_Node1 will come up as primary, operational secondary. | primary | Secondary Sticky bit: False | secondary | Primary Sticky bit: True | 
| 9 | Bring up all vPCs and uplink ports on New_Node1. | Traffic will be forwarded by both New_Node1 and New_Node2. | primary | Secondary Sticky bit: False | secondary | Primary Sticky bit: True | 
| Note | If you prefer to have the configured secondary node as the operational secondary and the configured primary as the operational primary, then Node2 can be reloaded at the end of the migration. This is optional and does not have any functional impact. |
