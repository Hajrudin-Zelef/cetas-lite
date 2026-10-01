---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide-e2ffcac9-6
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9.md
source_anchor: ""
source_lines: [405, 454]
sha256: 48536bce3028953c86b93292fabd4e0041069cda5b6b50ac34c374166b68f0f8
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9

                                       It is an edge port (a port configured to be at the edge of the network).
If a designated port is in the forwarding state and is not configured as an edge port, it transitions to the blocking state when the Rapid PVST+ forces it to synchronize with new root information. In general, when the Rapid PVST+ forces a port to synchronize with root information and the port does not satisfy any of the above conditions, its port state is set to blocking.
After ensuring that all of the ports are synchronized, the device sends an agreement message to the designated device that corresponds to its root port. When the devices connected by a point-to-point link are in agreement about their port roles, Rapid PVST+ immediately transitions the port states to the forwarding state.
Processing Superior BPDU Information
A superior BPDU is a BPDU with root information (such as a lower switch ID or lower path cost) that is superior to what is currently stored for the port.
If a port receives a superior BPDU, Rapid PVST+ triggers a reconfiguration. If the port is proposed and is selected as the new root port, Rapid PVST+ forces all the designated, nonedge ports to synchronize.
If the received BPDU is a Rapid PVST+ BPDU with the proposal flag set, the device sends an agreement message after all of the other ports are synchronized. The new root port transitions to the forwarding state as soon as the previous port reaches the blocking state.
If the superior information received on the port causes the port to become a backup port or an alternate port, Rapid PVST+ sets the port to the blocking state and sends an agreement message. The designated port continues sending BPDUs with the proposal flag set until the forward-delay timer expires. At that time, the port transitions to the forwarding state.
Processing Inferior BPDU Information
An inferior BPDU is a BPDU with root information (such as a higher switch ID or higher path cost) that is inferior to what is currently stored for the port.
If a designated port receives an inferior BPDU, it immediately replies with its own information.
Detecting Unidirectional Link Failure:Rapid PVST+
The software checks the consistency of the port role and state in the received BPDUs to detect unidirectional link failures that could cause bridging loops using the Unidirectional Link Detection (UDLD) feature. This feature is based on the dispute mechanism.
See the Cisco Nexus 9000 Series NX-OS Interfaces Configuration Guide, for information on UDLD.
When a designated port detects a conflict, it keeps its role, but reverts to a discarding state because disrupting connectivity in case of inconsistency is preferable to opening a bridging loop.
Port Cost
| Note | Rapid PVST+ uses the short (16-bit) path-cost method to calculate the cost by default. With the short path-cost method, you can assign any value in the range of 1 to 65535. However, you can configure the device to use the long (32-bit) path-cost method, which allows you to assign any value in the range of 1 to 200,000,000. You configure the path-cost calculation method globally. | 
This table shows how the STP port path-cost default value is determined from the media speed and path-cost calculation method of a LAN interface.
| Table 4. Default Port                                           		  Cost |  |  | 
|---|---|---|
| Bandwidth | Short Path-Cost Method of Port Cost | Long Path-Cost Method of Port Cost | 
|---|---|---|
| 10 Mbps | 100 | 2,000,000 | 
| 100 Mbps | 19 | 200,000 | 
| 1 Gbps | 4 | 20,000 | 
| 10 Gbps | 2 | 2,000 | 
| 40 Gbps | 1 | 500 | 
| 100 Gbps | 1 | 200 | 
| 400 Gbps | 1 | 50 | 
If a loop occurs, STP considers the port cost when selecting a LAN interface to put into the forwarding state.
You can assign the lower cost values to LAN interfaces that you want STP to select first and higher cost values to LAN interfaces that you want STP to select last. If all LAN interfaces have the same cost value, STP puts the LAN interface with the lowest LAN interface number in the forwarding state and blocks other LAN interfaces.
On access ports, you assign the port cost by the port. On trunk ports, you assign the port cost by the VLAN; you can configure the same port cost to all the VLANs on a trunk port.
Port Priority
If a redundant path occurs and multiple ports have the same path cost, Rapid PVST+ considers the port priority when selecting which LAN port to put into the forwarding state. You can assign lower priority values to LAN ports that you want Rapid PVST+ to select first and higher priority values to LAN ports that you want Rapid PVST+ to select last.
If all LAN ports have the same priority value, Rapid PVST+ puts the LAN port with the lowest LAN port number in the forwarding state and blocks other LAN ports. The possible priority range is from 0 through 224 (the default is 128), configurable in increments of 32. The device uses the port priority value when the LAN port is configured as an access port and uses the VLAN port priority values when the LAN port is configured as a trunk port.
Rapid PVST+ and IEEE 802.1Q Trunks
The 802.1Q trunks impose some limitations on the STP strategy for a network. In a network of Cisco network devices connected through 802.1Q trunks, the network devices maintain one instance of STP for each VLAN allowed on the trunks. However, non-Cisco 802.1Q network devices maintain only one instance of STP for all VLANs allowed on the trunks, which is the Common Spanning Tree (CST).
When you connect a Cisco network device to a non-Cisco device through an 802.1Q trunk, the Cisco network device combines the STP instance of the 802.1Q VLAN of the trunk with the STP instance of the non-Cisco 802.1Q network device. However, all per-VLAN STP information that is maintained by Cisco network devices is separated by a cloud of non-Cisco 802.1Q network devices. The non-Cisco 802.1Q cloud that separates the Cisco network devices is treated as a single trunk link between the network devices.
For more information on 802.1Q trunks, see the Cisco Nexus 9000 Series NX-OS Interfaces Configuration Guide.
Rapid PVST+ Interoperation with Legacy 802.1D STP
Rapid PVST+ can interoperate with devices that are running the legacy 802.1D protocol. The device knows that it is interoperating with equipment running 802.1D when it receives a BPDU version 0. The BPDUs for Rapid PVST+ are version 2. If the BPDU received is an 802.1w BPDU version 2 with the proposal flag set, the device sends an agreement message after all of the other ports are synchronized. If the BPDU is an 802.1D BPDU version 0, the device does not set the proposal flag and starts the forward-delay timer for the port. The new root port requires twice the forward-delay time to transition to the forwarding state.
The device interoperates with legacy 802.1D devices as follows:
-  
                                    			 
                                    Notification—Unlike 802.1D BPDUs, 802.1w does not use TCN BPDUs. However, for interoperability with 802.1D devices, the device processes and generates TCN BPDUs.
-  
                                    			 
                                    Acknowledgment—When an 802.1w device receives a TCN message on a designated port from an 802.1D device, it replies with an 802.1D configuration BPDU with the TCA bit set. However, if the TC-while timer (the same as the TC timer in 802.1D) is active on a root port connected to an 802.1D device and a configuration BPDU with the TCA set is received, the TC-while timer is reset. This method of operation is required only for 802.1D devices. The 802.1w BPDUs do not have the TCA bit set.
-  
                                    			 
