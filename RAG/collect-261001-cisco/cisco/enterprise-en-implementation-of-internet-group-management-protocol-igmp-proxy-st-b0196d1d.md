---
id: collect-261001-cisco/cisco/enterprise-en-implementation-of-internet-group-management-protocol-igmp-proxy-st-b0196d1d
title: "enterprise-en-implementation-of-internet-group-management-protocol-igmp-proxy-st-b0196d1d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-implementation-of-internet-group-management-protocol-igmp-proxy-st-b0196d1d.md
source_anchor: ""
source_lines: [1, 11]
sha256: 614c738d37dc30206b64c7a56c48a04a95bb3688eaa3209238b4e99f9391cdb7
---

# enterprise-en-implementation-of-internet-group-management-protocol-igmp-proxy-st-b0196d1d

Hello everyone!
Today I share with you a technical application on how to implement Internet Group Management Protocol (IGMP) proxy!
Important Steps while implementing of IGMP Proxy
* The IGMP proxy Implementation is as follow:
- Step 1 To order a video program, an IGMP user must send an IGMP request to the IGMP Proxy for joining the multicast group corresponding to the program.
- Step 2 Upon receiving the request, the IGMP Proxy authenticates the user. If the user is authorized to watch or preview the program, the IGMP Proxy allows the user to join the multicast group. If not, IGMP Proxy rejects the user's request.
- Step 3 The IGMP Proxy sends general group query messages to all online IGMP users at regular intervals. If it fails to receive a response from a user within a certain period, it considers the user has left he multicast group, and deletes the user from the multicast group.
- Step 4 when leaving a multicast group, an IGMP user sends a leave packet to the IGMP Proxy. If the user is set with quick leave attributes, the IGMP proxy deletes the user from the multicast group directly. If not, the IGMP Proxy sends a specific group display to the multicast group. if having not received any response packets from the user within a certain period, the IGMP proxy deletes the user from the multicast group.
- Step 5 Meanwhile, when receiving a display from the upper layer router, the IGMP Proxy returns a response the current multicast state to the router.
That's all, welcome to leave a message to discuss.
Thank you!
