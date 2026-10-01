---
id: collect-261001-fortinet/fortinet/t5-support-forum-does-fgt-sdwan-support-configure-auto-discovery-receiver-and-m-df5deec0
title: "t5-support-forum-does-fgt-sdwan-support-configure-auto-discovery-receiver-and-m--df5deec0"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/collect-261001-fortinet/t5-support-forum-does-fgt-sdwan-support-configure-auto-discovery-receiver-and-m--df5deec0.md
source_anchor: ""
source_lines: [1, 41]
sha256: bab7d14aef7fc4f6c98b4b6ab4f743bd3fc5ebae03f1255b8045880b9fa4aadd
---

# t5-support-forum-does-fgt-sdwan-support-configure-auto-discovery-receiver-and-m--df5deec0

HI,


Would like to know does FGT sdwan support configure auto-discovery-receiver and auto-discovery-sender in same devices (Region)? Let me share some background here, i would like configure 3 tier SDWAN connection like below. may i know is it feasible to configure ADVPN "**auto-discovery-receiver**" in "**Region device**" as a **spoke** device talk to **Hub (HQ)**, as well as configure ADVPN **"auto-discovery-sender"** as a hub for **(Branch) spoke** ?


Hub (HQ) <------ **Spoke (Region) /Hub** <------- spoke (Branch)


Thanks


Hi,

Yes, you can have both auto-discovery-receiver and auto-discovery-sender enabled in the same device with SD-WAN configured.

Here are some considerations and insights regarding the setup and requirements:

In essence, your network structure involves a central Hub at the Headquarters (HQ), overseeing multiple Regions. Each Region, in turn, acts as a hub for its respective branches. The goal here is to establish seamless communication: Regions forming shortcut tunnels with other Regions linked to the HQ-Hub, and branches creating shortcut tunnels with other branches connected to the same Region-Hub.

Achieving the desired setup encounters a challenge when attempting to establish a single tunnel on the Region-Hub side. The reason is rooted in the configuration disparity between the Hub and the spokes: the Hub utilizes Dial-Up tunnels, while the spokes rely on Static tunnels. The initiation of VPN tunnels consistently originates from the spoke end due to the Hub's lack of knowledge about the VPN gateway addresses of the spokes.

In light of this, the Region-HUB is configured with a Dial-Up tunnel (for branches to connect) and the HQ-HUB with another Dial-Up tunnel (for regions to connect). However, these two Dial-Up tunnels—one from the HQ-HUB and the other from the Region-HUB—cannot directly communicate or form a tunnel between them. To address this, specific configurations are needed on the Hubs.

So, that the communication between the Region-HUBs can flow via shortcuts between them. At the same time, the communication between the branches from the same location can flow via the shortcuts formed between them. Now if the communication between the branch from different regions, would flow via the shortcuts between the Region-HUBs.

If you require, shortcuts to be formed between the branches of different regions, then additional configuration is required.

Cheers,


HI @rarumugam ,


It's quit useful for the information. May i know is it a standard way to achieve above design as i don't see any Fortinet document have related information. Just found out Multi region sample without any sample configuration/parameter from fortinet document.


Regarding the additional configuration for shortcuts to be formed between the branches of different regions, could share it out the sample or parameter?


The Fortinet Security Fabric brings together the concepts of convergence and consolidation to provide comprehensive cybersecurity protection for all users, devices, and applications and across all network edges.
