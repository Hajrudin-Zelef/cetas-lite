---
id: collect-261001-cisco/cisco/enterprise-en-bfd-for-is-is-thread-485001-863-a8f16171
title: "enterprise-en-bfd-for-is-is-thread-485001-863-a8f16171"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-bfd-for-is-is-thread-485001-863-a8f16171.md
source_anchor: ""
source_lines: [1, 6]
sha256: 370d838e3f311774a1ab0db05111bd321ef9968a54c6786167b7c38a829627fa
---

# enterprise-en-bfd-for-is-is-thread-485001-863-a8f16171

Generally, the interval at which the Intermediate System to Intermediate System (IS-IS) protocol sends Hello messages is 10 seconds. If a device does not receive any Hello message from its neighbor within three Hello intervals, the device deletes the neighbor. Therefore, it takes a device a number of seconds to detect that a neighbor is Down. This leads to the loss of a large number of packets in a high-speed network.
In BFD for IS-IS, the establishment of a BFD session is dynamically triggered by IS-IS but not configured manually. When detecting a fault, the BFD session notifies IS-IS of the fault through the Routing Management Module (RM). IS-IS processes the neighbor-Down event and quickly sends the link state PDU (LSP), and performs the partial route calculation (PRC). In this manner, IS-IS routes fast converge.
The BFD fault detection interval is at the millisecond level. Instead of replacing the IS-IS Hello mechanism, BFD works with IS-IS to detect the adjacency fault more quickly. In addition, BFD instructs IS-IS to recalculate routes, ensuring correct packet forwarding.
The RM allows IS-IS and BFD to interact with each other. Through the RM, IS-IS instructs BFD to dynamically set up or delete BFD sessions. The BFD event messages are also delivered to IS-IS through the RM.
Figure 1 BFD for IS-IS networking diagram 
After BFD is enabled on Switch A, Switch B, and Switch C, the BFD session can quickly detect faults on the link between Switch A and Switch B, and notify IS-IS through the RM. Then, IS-IS sets the neighbor status to Down to trigger the IS-IS topology calculation. In addition, IS-IS updates LSPs to ensure that Switch C (Switch B's neighbor) can receive the updated LSPs from Switch B in time. This implements fast network topology convergence.
