---
id: collect-261001-huawei/huawei/questions-83086-182bfca6
title: "questions-83086-182bfca6"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-83086-182bfca6.md
source_anchor: ""
source_lines: [1, 38]
sha256: 7073e32fe1a0f22b02ae94f64b413334891363a0f64eef9332dc904cf657c9eb
---

# questions-83086-182bfca6

Currently, I am trying to configure a network to transport multiple VLANs for a client who has points in different locations (Side A and Side B). My idea was to use QinQ for this. The network topology is represented in the figure below:
The configuration of switches A and B is as follows:
Switch A
#
interface XGigabitEthernet0/0/44
 description Clientx - Side A
 port link-type hybrid
 qinq vlan-translation enable
 port hybrid tagged vlan 668
 port hybrid untagged vlan 667
 port vlan-stacking vlan 100 stack-vlan 667
 port vlan-stacking vlan 200 stack-vlan 667
 port vlan-stacking vlan 499 stack-vlan 667
#
#
interface Vlanif667
 l2 binding vsi clientex
#
Switch B
#
interface XGigabitEthernet0/0/24
 description ClientX - Side B
 port link-type hybrid
 qinq vlan-translation enable
 port hybrid tagged vlan 668
 port hybrid untagged vlan 667
 port vlan-stacking vlan 100 stack-vlan 667
 port vlan-stacking vlan 200 stack-vlan 667
 port vlan-stacking vlan 499 stack-vlan 667
#
#
interface Vlanif667
 l2 binding vsi clientex
#
The VSI is active because there is an MPLS session established between switch A and switch B. However, when I try to ping from switch A to B or vice-versa, using the /30 IP addresses added to VLANs 100, 200, and 499, there is no communication. I tried changing from VSI to MPLS L2VC, but without success. Communication only works when I pass the VLAN 667 directly to the interfaces of Switch A, X and B (without using MPLS transport). Does anyone know what the problem might be?
Notes:
- I cannot maintain the configuration by passing the VLANs on the interfaces because the actual topology is larger than this example, and some interfaces have disabled port switches.
- I cannot choose to configure VSI directly on the physical interfaces and remove QinQ because I need to pass VLAN 668 for management of the customer equipment connected to Switch A.
