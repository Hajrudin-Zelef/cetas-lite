---
id: collect-261001-fortinet/fortinet/questions-77331-fortigate-ipsec-peer-to-peer-and-two-remote-peers-806dc54f
title: "questions-77331-fortigate-ipsec-peer-to-peer-and-two-remote-peers-806dc54f"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-77331-fortigate-ipsec-peer-to-peer-and-two-remote-peers-806dc54f.md
source_anchor: ""
source_lines: [1, 10]
sha256: f7e761eb9111a7a9a311cc3390cbaaf1cee186837184565578a362f9e6d856f9
---

# questions-77331-fortigate-ipsec-peer-to-peer-and-two-remote-peers-806dc54f

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a FortiGate firewall in the main office and some branches with Cisco ASA. An IPsec tunnel runs between the main office and each branch. Some branches have two ISP - main and reserve. For example, building a tunnel between Cisco ASA with one public address and remote Cisco ASA with two public address is a simple task: we can set two remote peers in a crypto map for the device in main office. But we can not to do the same in FortiGate IPSec settings of Phase 2.
This can easily be done by using route-based tunnels and throwing BGP on top so that you can peer with both of the ASAs connections at the same time and use one of the methods that BGP supports for route preference, called AS-path prepending, to “weight” one of the routes so it’s less preferred than the other but will already be visible as a backup route in the event the preferred route disappears (like if the connection is severed).
FortiGates use route-based tunnels by default, though you can enable policy-based tunnels via the Feature Visibility screen. For the ASA side, you will need to run 9.7 or newer versions of ASA OS in order to support VTIs (virtual tunnel interfaces) and to be able to create route-based tunnels.
For examples and details on how to create route-based tunnels and VTIs with BGP on an ASA, you can follow the guide here
