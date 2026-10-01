---
id: collect-261001-general-networking/general-networking/questions-815286-how-to-configure-hp-msm720-with-multiple-vsc-from-which-one-has-e66a1c3d
title: "questions-815286-how-to-configure-hp-msm720-with-multiple-vsc-from-which-one-has-e66a1c3d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/questions-815286-how-to-configure-hp-msm720-with-multiple-vsc-from-which-one-has-e66a1c3d.md
source_anchor: ""
source_lines: [1, 6]
sha256: 85f58e8fd75233baf40f9c50a1ceb3c019c0ad4658e228f00736c0e7969e3c6d
---

# questions-815286-how-to-configure-hp-msm720-with-multiple-vsc-from-which-one-has-e66a1c3d

I have one really odd project to pull out. I am replacing entire infrastructure on one of our customers. They requirements are quite easy until they decided to put something else. 3 WAN addresses that are assigned like this: 1x WAN for plain internet connection which is distributed via HP E-MSM720 controller and HP MSM460 access points (default VLAN and VSC), now the fun part. 2x WAN to be assigned for one server inside clients network VLAN 6 subnet 192.168.x.x/23. This is also quite simple to pull. What I am having trouble, most likely because I have never played with HP controllers, is to set additional VSC with VLAN 6 and stream content from that server via Wireless. Now I must assume that on Meraki switch I need 2 ports as assign VLAN 6 on them (untagged). On Controller same thing. I need configure new VLAN 6 using "Network"->"Network Profiles", and under "Network"-> "VLANs" select new created profile and assign port as untagged. Since from some strange reason I cannot use "Automated Workflow" as I am getting an error, I need to create new VSC and bind it with that VLAN/Network profile that I would like to create. Creating profile should be easy by selecting "Controller"->"VSCs"-> "Add New VSC Profile". There I would set all the parameters such as SSID, Authentication, broadcasting, hiding SSID, etc, etc.
Can somebody confirm if the binding of VSC with new VLAN 6 using "VSC bindings" interfere with default VSC or disrupt current wireless link which is already configured? Manual procedure that I would use would be: - Select new "Controller"->"VSC"->"VSC Bindings" -> "Add New Binding"-> Check option "Egress Network - Network Profile: VLAN 6" and Save.
Will this in any way drop connection for current default VSC which uses default VLAN?
Also one additional question which I don't know if anyone knows, and I definitely need clarification:
If the VLAN is going to be set as IGMP (multicast), how can you set VSC to support IGMP? There is an option of IGMP proxy but is related to an interface. I assume that this interface that is facing switches/access points, if configured with IGMP proxy, might disrupt wireless internet connection for default VSC and default VLAN correct? (normal internet connection is unicast, while video stream is using multicast)...
Please refer from the schematic below, how it should look like. Clients network schematic
