---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide-e2ffcac9-4
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9.md
source_anchor: ""
source_lines: [236, 380]
sha256: b19c1f371384ee0b163cafe4628e837481a3e525200822fbc734741d5f22a9ff
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9

                                       Blocking—The Layer 2 LAN port does not participate in frame forwarding.
-  
                                       			 
                                       Learning—The Layer 2 LAN port prepares to participate in frame forwarding.
-  
                                       			 
                                       Forwarding—The Layer 2 LAN port forwards frames.
-  
                                       			 
                                       Disabled—The Layer 2 LAN port does not participate in STP and is not forwarding frames.
When you enable Rapid PVST+, every port in the device, VLAN, and network goes through the blocking state and the transitory states of learning at power up. If properly configured, each Layer 2 LAN port stabilizes to the forwarding or blocking state.
When the STP algorithm places a Layer 2 LAN port in the forwarding state, the following process occurs:
-  
                                       			 
                                       The Layer 2 LAN port is put into the blocking state while it waits for protocol information that suggests it should go to the learning state.
-  
                                       			 
                                       The Layer 2 LAN port waits for the forward delay timer to expire, moves the Layer 2 LAN port to the learning state, and restarts the forward delay timer.
-  
                                       			 
                                       In the learning state, the Layer 2 LAN port continues to block frame forwarding as it learns the end station location information for the forwarding database.
-  
                                       			 
                                       The Layer 2 LAN port waits for the forward delay timer to expire and then moves the Layer 2 LAN port to the forwarding state, where both learning and frame forwarding are enabled.
Blocking State
A Layer 2 LAN port in the blocking state does not participate in frame forwarding.
A Layer 2 LAN port in the blocking state performs as follows:
- 
                                          
                                          
                                          
                                          Discards frames received from the attached segment.
- 
                                          
                                          
                                          
                                          Discards frames switched from another port for forwarding.
- 
                                          
                                          
                                          
                                          Does not incorporate the end station location into its address database. (There is no learning on a blocking Layer 2 LAN port, so there is no address database update.)
- 
                                          
                                          
                                          
                                          Receives BPDUs and directs them to the system module.
- 
                                          
                                          
                                          
                                          Receives, processes, and transmits BPDUs received from the system module.
- 
                                          
                                          
                                          
                                          Receives and responds to control plane messages.
Learning State
A Layer 2 LAN port in the learning state prepares to participate in frame forwarding by learning the MAC addresses for the frames. The Layer 2 LAN port enters the learning state from the blocking state.
A Layer 2 LAN port in the learning state performs as follows:
- 
                                          
                                          
                                          
                                          Discards frames received from the attached segment.
- 
                                          
                                          
                                          
                                          Discards frames switched from another port for forwarding.
- 
                                          
                                          
                                          
                                          Incorporates the end station location into its address database.
- 
                                          
                                          
                                          
                                          Receives BPDUs and directs them to the system module.
- 
                                          
                                          
                                          
                                          Receives, processes, and transmits BPDUs received from the system module.
- 
                                          
                                          
                                          
                                          Receives and responds to control plane messages.
Forwarding State
A Layer 2 LAN port in the forwarding state forwards frames. The Layer 2 LAN port enters the forwarding state from the learning state.
A Layer 2 LAN port in the forwarding state performs as follows:
- 
                                          
                                          
                                          
                                          Forwards frames received from the attached segment.
- 
                                          
                                          
                                          
                                          Forwards frames switched from another port for forwarding.
- 
                                          
                                          
                                          
                                          Incorporates the end station location information into its address database.
- 
                                          
                                          
                                          
                                          Receives BPDUs and directs them to the system module.
- 
                                          
                                          
                                          
                                          Processes BPDUs received from the system module.
- 
                                          
                                          
                                          
                                          Receives and responds to control plane messages.
Disabled State
A Layer 2 LAN port in the disabled state does not participate in frame forwarding or STP. A Layer 2 LAN port in the disabled state is virtually nonoperational.
A disabled Layer 2 LAN port performs as follows:
- 
                                          
                                          
                                          
                                          Discards frames received from the attached segment.
- 
                                          
                                          
                                          
                                          Discards frames switched from another port for forwarding.
- 
                                          
                                          
                                          
                                          Does not incorporate the end station location into its address database. (There is no learning, so there is no address database update.)
- 
                                          
                                          
                                          
