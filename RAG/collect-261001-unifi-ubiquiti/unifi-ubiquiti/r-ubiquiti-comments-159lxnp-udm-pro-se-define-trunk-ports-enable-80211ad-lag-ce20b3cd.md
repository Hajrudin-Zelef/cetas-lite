---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-159lxnp-udm-pro-se-define-trunk-ports-enable-80211ad-lag-ce20b3cd
title: "r-ubiquiti-comments-159lxnp-udm-pro-se-define-trunk-ports-enable-80211ad-lag-ce20b3cd"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-159lxnp-udm-pro-se-define-trunk-ports-enable-80211ad-lag-ce20b3cd.md
source_anchor: ""
source_lines: [1, 29]
sha256: 0d285fc9d5f7677562c1d5c4792238748f771f5c76634ca7c65b9e04e99db52f
---

# r-ubiquiti-comments-159lxnp-udm-pro-se-define-trunk-ports-enable-80211ad-lag-ce20b3cd

UDM Pro SE - Define Trunk Ports + Enable 802.11ad LAG
Hello,
Former Edgerouter 12 and current UDM Pro SE Owner as of 2 days ago.
I'm trying to find within the UI where to enable and assign trunk ports and enable aggregation.
I updated to the latest version and using the current UI. I've watched a few youtube videos and none of them look like my interface. This is a bit confusing
My goal is to assign a Native VLAN 1 on the SFP Lan Interface with Trunks for Vlan 10 and 20. Looking like this:
10.0.1.1/24 - Native MGMT (vlan1?)
10.0.10.1/24 - VLAN 10
10.0.20.1/24 VLAN 20
I would then like to enable a port channel to Cisco Catalyst 9800.
Any help would be apperciated!
Thanks
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Can the 2 SFP ports be uplink(s) to another switch or WLC?
No, they can’t. I was also disappointed when I figured that out a few weeks ago.
Yikes - do you have a document for that ? Seems like it would be simple to implement.
The UDMP/UDMSE don’t have LAG/portchannel because it would be useless. Take a look at the block diagram here, pay attention to what’s between the 8 LAN ports and the rest: https://ubntwiki.com/products/unifi/unifi_dream_machine_pro
The UDMSE replaces the RJ45 WAN with 2.5Gb but the internal link to the switch chip is still 1Gb.
This is effectively a 4-port router. Guess how they made the UXGP?
If you create the VLANs (Networks) and leave the LAN port in the default state, it will be the same as “switchport mode trunk” “switchport trunk native vlan 1” in Cisco terms - meaning it will have all VLANs other than 1 tagged.
Useless is subjective it's for lab testing. In addition an argument could be made for redudancy. I'm looking to aggregate the SFP ports which appear to have more throughput.
Ubiquiti hasn’t implemented it. Go ask on their forum if you want it implemented.
The intended deployment if you care about performance is to connect a better switch with SFP+ to have 10Gb/s. The builtin switch is intended for small deployments that don’t need more than that number of ports. It also doesn’t have STP or other loop detection.
Not a feature in the gateways
