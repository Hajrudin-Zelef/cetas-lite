---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide-e2ffcac9-5
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9.md
source_anchor: ""
source_lines: [381, 404]
sha256: eff2e358ed15e7808b46ba24ef573d0192b761a82450d052fcc49f8bec088805
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9

                                          Does not receive BPDUs from neighbors.
- 
                                          
                                          
                                          
                                          Does not receive BPDUs for transmission from the system module.
Summary of Port States
This table lists the possible operational and Rapid PVST+ states for ports and whether the port is included in the active topology.
| Table 3.                                                                                                                                                                                                                                                                                     Port State Active Topology |  |  | 
|---|---|---|
| Operational Status | Port State | Is Port Included in the Active Topology? | 
|---|---|---|
| Enabled | Blocking | No | 
| Enabled | Learning | Yes | 
| Enabled | Forwarding | Yes | 
| Disabled | Disabled | No | 
Synchronization of Port Roles
When the device receives a proposal message on one of its ports and that port is selected as the new root port, Rapid PVST+ forces all other ports to synchronize with the new root information.
The device is synchronized with superior root information received on the root port if all other ports are synchronized. An individual port on the device is synchronized if either of the following applies:
-  
                                       			 
                                       That port is in the blocking state.
-  
                                       			 
