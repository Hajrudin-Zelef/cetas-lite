---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1000060766-3a410bc-how-do-i-configure-the-link-type-of-an-174bba9b
title: "enterprise-en-doc-edoc1000060766-3a410bc-how-do-i-configure-the-link-type-of-an--174bba9b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1000060766-3a410bc-how-do-i-configure-the-link-type-of-an--174bba9b.md
source_anchor: ""
source_lines: [1, 32]
sha256: a57da2fa58120dca68136bd018b890db292756e28e2ec3a0cca3932cb829dcbf
---

# enterprise-en-doc-edoc1000060766-3a410bc-how-do-i-configure-the-link-type-of-an--174bba9b

Enterprise
Descriptions of link types
| Type | Usage Scenario | 
|---|---|
| Access | An access interface on a switch is used to connect a user host. The access interface sends only untagged Ethernet frames to the remote device. | 
| Trunk | A trunk interface on a switch is used to connect to another switch. The trunk interface allows Ethernet frames from multiple VLANs to pass. | 
| Hybrid | A hybrid interface on a switch can connect either to a host or to another switch. The hybrid interface allows Ethernet frames from multiple VLANs to pass, and can be configured to determine whether the Ethernet frames from the outbound interface carry tags. | 
Configurations of link types
The following provides configuration examples of three link types. By default, the link type of an interface is access.
<HUAWEI> system-view
[~HUAWEI] vlan batch 3
[*HUAWEI] interface 10ge 1/0/1
[*HUAWEI-10GE1/0/1] port link-type access  //Set the link type to access.
[*HUAWEI-10GE1/0/1] port default vlan 3  //Set the default VLAN ID of the access interface to 3.
[*HUAWEI-10GE1/0/1] quit
[*HUAWEI] commit
<HUAWEI> system-view
[~HUAWEI] vlan batch 10 to 30
[*HUAWEI] interface 10ge 1/0/1
[*HUAWEI-10GE1/0/1] port link-type trunk  //Set the link type to trunk.
[*HUAWEI-10GE1/0/1] port trunk allow-pass vlan 10 to 30  //Set the range of VLANs allowed by the trunk interface to 10-30.
[*HUAWEI-10GE1/0/1] quit
[*HUAWEI] commit
<HUAWEI> system-view
[~HUAWEI] vlan batch 3 to 9
[*HUAWEI] interface 10ge 1/0/1
[*HUAWEI-10GE1/0/1] port link-type hybrid  //Set the link type to hybrid.
[*HUAWEI-10GE1/0/1] port hybrid untagged vlan 3 to 5  //Set the range of VLANs allowed by the hybrid interface in untagged mode to 3-5.
[*HUAWEI-10GE1/0/1] port hybrid tagged vlan 6 to 9  //Set the range of VLANs allowed by the hybrid interface in tagged mode to 6-9.
[*HUAWEI-10GE1/0/1] quit
[*HUAWEI] commit
Select the content with the mouse pointer to quickly report the problem.
