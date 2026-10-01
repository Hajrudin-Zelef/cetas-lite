---
id: collect-261001-cisco/cisco/enterprise-en-typical-stack-configuration-of-fixed-switches-thread-6672253202033-bad4b4f6-1
title: "enterprise-en-typical-stack-configuration-of-fixed-switches-thread-6672253202033-bad4b4f6"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-typical-stack-configuration-of-fixed-switches-thread-6672253202033-bad4b4f6.md
source_anchor: ""
source_lines: [1, 54]
sha256: b2050fea7e029a48224cb9e58109b098af6e73099f130466938bd89f2ad1ba18
---

# enterprise-en-typical-stack-configuration-of-fixed-switches-thread-6672253202033-bad4b4f6

Hello, everyone! This post describes the stack deployment method and recommendations.
Recommended Stack Deployment Scenarios
Determining the Stack Topology
Stack Configuration and Deployment Recommendations
This is the most common scenario when aggregation switches set up a stack system, as shown in Figure 4-1.
The following switch models can set up a stack system in this scenario: S6700EI, S6720S-EI, S6720EI, S6720HI, S5700HI, S5710HI, S5710EI, S5700EI, S5700SI, S5720EI, S5720HI, and S5730HI.
In this scenario, each switch in a stack connects to a core device through Eth-Trunk. The stack system simplifies management of aggregation devices and improves uplink reliability of access devices.
Figure 4-1  Stack system operating on aggregation switches 
This is the most common scenario when Layer 2 access switches set up a stack system, as shown in Figure 4-2.
The following switch models can set up a stack system in this scenario: S2720EI, S2750EI, S5700LI, S5700EI, S5710-C-LI, S5710-X-LI, S5720LI, S5720S-LI, S5700SI, S5720SI, S5720S-SI, S5720I-SI, S5700S-LI, S5730SI, S5730S-EI, S6720LI, S6720S-LI, S6720SI, and S6720S-SI.
In this scenario, each switch in a stack connects to an aggregation device through Eth-Trunk. The stack system simplifies management of and improves uplink reliability of access devices.
Figure 4-2  Stack system operating on access switches 
This scenario rarely occurs. Figure 4-3 shows the networking of this scenario.
In this scenario, multiple stack systems form a ring through Eth-Trunk, and one stack system connects to aggregation switches through Eth-Trunk. This scenario reduces the number of management IP addresses of access devices.
Figure 4-3  Stack system operating on an access ring 
NOTE:
The following recommendations are provided based on the positioning of fixed switch models. If customers have special requirements, it is recommended to deploy high-end devices at a lower network layer; it is not recommended to deploy low-end devices at a higher network layer. For example, it is recommended to deploy aggregation switches at the access layer rather than to deploy access switches at the aggregation layer.
To ensure stack reliability and bandwidth, you are advised to do as follows:
Ensure that each member device connects to the core device through an uplink port. This connection prevents upstream traffic forwarding from being affected when any member device fails.
When using multiple devices to set up a stack, ensure the same stack bandwidth between any two devices. Otherwise, the bandwidth of the stack system is the minimum stack bandwidth.
Table 4-1 Scenario recommendations
| Model | Scenario 1 | Scenario 2 | Scenario 3 | 
|---|---|---|---|
| S5700HI, S5710HI, S5710EI, S6700EI | First preferred | Second preferred | Not recommended | 
| S5720EI, S5720HI, S5730HI, S6720EI, S6720S-EI | First preferred | Second preferred | Second preferred | 
| S5700EI, S5700SI | First preferred | First preferred | Second preferred | 
| S5720SI, S5720S-SI, S5720I-SI | Second preferred | First preferred | First preferred | 
| S2720EI, S2750EI, S5700LI, S5720LI, S5720S-LI, S5730SI, S5730S-EI, S6720LI, S6720S-LI, S6720SI, S6720S-SI | Not recommended | First preferred | Second preferred | 
| S5700S-LI, S5710-C-LI, S5710-X-LI | Not recommended | First preferred | First preferred | 
A stack can be connected in a chain or ring topology depending on the stack connection mode, as shown in Figure 4-4. Table 4-2 compares the two stack topologies in terms of reliability, link bandwidth utilization, and convenience of cable connections.Figure 4-4  Stack topologies 
Table 4-2  Comparison between stack topologies
| Stack Topology | Advantages | Disadvantages | Applicable Scenario | 
|---|---|---|---|
| Chain topology | Applicable to long-distance stacking because the first and last member switches do not need to be connected by a physical link. |  | Member devices are far from one another and a ring topology is difficult to deploy. | 
| Ring topology |  | The first and last member switches need to be connected by a physical link, so this topology is not applicable to long-distance stacking. | Member switches are located near one another. | 
Two devices can set up a stack in a chain topology, as shown in Figure 4-5. In this topology, only one logical stack port exists between the two devices and no loop exists in the stack.Figure 4-5  Only one logical stack port between two member devices 
Two devices can set up a stack with back-to-back networking, as shown in Figure 4-6. In this networking, two logical stack ports exist between the two devices, and one loop exists in the stack, which will be automatically eliminated by the system.Figure 4-6  Two logical stack ports between two member devices 
When using two devices to set up a stack, you are advised to do as follows:
If the devices provide no more than 28 ports, use the networking with only one logical stack port. Otherwise, use the back-to-back networking.
If more member devices need to be added to the stack in the future, use the back-to-back networking, which will require minimum modification to the existing system.
Connect at least two stack cables between the two devices to ensure reliability.
Stack Setup Guidelines
A stack contains a maximum of five, eight, or nine stack members depending on switch models. To ensure high forwarding performance and reliability, ensure that the number of stack members of each switch series does not exceed the recommended value. For example, the recommended number of S5700LI switches in a stack is 2 to 5. For the recommended value of each switch series, see "Service Port Stacking Support" or "Stack Card Stacking Support".
When setting up a stack containing a mix of different switch models, follow the following rules:
A stack cannot be set up among different switch series. For example, a stack cannot be set up among the S5700 and S6700 switches.
A stack cannot be set up among different switch models of the same switch series. For example, a stack cannot be set up among the S5720EI and S5720HI switches.
A stack can be set up among only certain different switch models of the same switch series. For example, a stack can be set up among the S5720SI and S5720S-SI switches but not among the S5720LI and S5720S-LI switches. For details about whether a stack can be set up among different switch models, see stack support of different switch models in Stacking Support and Version Requirements. As described in Remarksof S6720LI and S6720S-LI Service Port Stacking Support, a stack can be set up among all S6720LI models, among all S6720S-LI models, but not among the S6720LI and S6720S-LI models.
Version restrictions:
When multiple switches set up a stack, member switches will synchronize the running version of the master switch. If a member switch does not support this running version, it will restart repeatedly.
In V200R009C00, if MPLS-incapable S5720EIs exist in a stack, this stack cannot have MPLS enabled. If member devices in a stack are running MPLS services, adding MPLS-incapable S5720EIs to the stack is not allowed.
An S5720HI supports the stacking function since V200R009C00. When a member device in a stack is faulty and fails to restart for three consecutive times, the device attempts to roll back to a version earlier than V200R009C00 for restart. When the device restarts successfully after rolling back to a version earlier than V200R009C00, a multi-active situation may occur because the version earlier than V200R009C00 does not support the stacking function. To prevent this situation, you are advised to delete the system software earlier than V200R009C00 from member devices when using S5720HIs to set up a stack.
MAD specifications:
You can configure a maximum of eight direct detection links for each member switch in a stack.
You can configure the relay mode on a maximum of four Eth-Trunks in a stack.
