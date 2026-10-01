---
id: collect-261001-general-networking/general-networking/questions-7395-which-vlans-are-allowed-on-this-trunk-62833580
title: "questions-7395-which-vlans-are-allowed-on-this-trunk-62833580"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "pruning"]
source: docs/RAG/collect-261001-general-networking/questions-7395-which-vlans-are-allowed-on-this-trunk-62833580.md
source_anchor: ""
source_lines: [1, 21]
sha256: 453233f6dce28905dcabd363dace3f48cc1bc5b87e4a063f8e6f0c9250c86e64
---

# questions-7395-which-vlans-are-allowed-on-this-trunk-62833580

When you configure a trunk link with the command:
(config-if)#switchport mode trunk
all the vlans will be allowed in the trunk link.
If you want to control which VLANs are allowed in a trunk link use the following commands:
(config-if)#switchport trunk allowed vlan
  WORD    VLAN IDs of the allowed VLANs when this port is in trunking mode
  add     add VLANs to the current list
  all     all VLANs
  except  all VLANs except the following
  none    no VLANs
  remove  remove VLANs from the current list
For example if you want to allow VLANs 1,10,20,30,99 in a trunk link you have to type:
(config-if)#switchport trunk allowed vlan 1,10,20,30,99
Pay attention that if you want to add a vlan on a trunk link (for example VLAN 100) you have to use the add statement (this is really a commond mistake), for example:
(config-if)#switchport trunk allowed vlan add 100
When you read "Vlans allowed and active" it simply means that just the active VLANs (the ones you have defined) are allowed in that trunk link.
Reference here: How to define the VLANs allowed on a trunk link
EDIT: Analyzing the output of the show interface trunk command it's easy to see that it is organized in 4 sections; 
- The first one is a list of all the trunk interfaces togheter with the encapsulation and the native VLAN.
- For what concern the second and the third lines the important point to clarify is the difference between allowed and active. Allowed simply stays for "It is allowed to pass" meanwhile active means "this VLAN is up and running in the swtich". Note that active does not simply mean "it is present in the database" since you can shutdown a vlan with (config-vlan)#shutdown .
- The last one deserves a little note; when dealing with PVST (Per-VLAN Spanning Tree) the switch creates a loop-free topology for each L2 broadcast domain (VLAN). Since it is not said that each VLAN share the same phisical topology there could be the case in which, even if a vlan is active and allowed, that vlan will not be forwarded across the trunk interface. This is because that interface it is not in the Forwarding state for that particular PVST process. Is it also possible that a VLAN is not forwarded across a trunk if it has been pruned by the VTP pruning process.
