---
id: collect-261001-general-networking/general-networking/t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4-2
title: "t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4.md
source_anchor: ""
source_lines: [23, 108]
sha256: 8a1c68f1e36a4c664c5c8f525dccbb5911c953427e88c9f87ac653a52993cdf6
---

# t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4

			Other Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-22-2014 12:12 PM
Hi Abhisar,
Regarding your individual decisions: Manually configuring the Root Bridge is a natural thing to do. You should never leave your network just pick up a root switch based on default switch settings.
On end ports, using PortFast and BPDU Guard is a must especially if you are running Rapid PVST+ or MSTP.
Regarding the Root Guard on ports to other switches - this is something I do not recommend. The Root Guard is a protective mechanism in situations when your network and the network of your customer need to form a single STP domain, yet you want to have the STP Root Bridge in your network part and you do not want your customer to take over this root switch selection. In these cases, you would put the Root Guard on ports toward the customer. However, inside your own network, using Root Guard is a questionable practice. Your network can be considered trustworthy and there is no rogue root switch to protect against. Using Root Guard in your own network could cause your network to be unable to converge on a new workable spanning tree if any of the primary links failed, and it would also prevent your network from converging to a secondary root switch if the primary root switch failed entirely. Therefore, I personally see no reason to use Root Guard inside your own network - on the contrary, I am concerned that it would basically remove the possibility of your network to actually utilize the redundant links and switches.
Regarding EtherChannels - yes, you are right, using the on mode can, under circumstances, lead to permanent switching loops. EtherChannel is one of few technologies in which I wholeheartedly recommend on relying on a signalling protocol to set it up, as opposed to configuring it manually. The active mode is my preferred mode, as it utilizes the open LACP to signal the creation of an EtherChannel, and setting both ends of a link to active helps to bring up the EtherChannel somewhat faster.
If you are using fiber links between switches, I recommend running UDLD on them to be protected against issues caused by uni-directional links. UDLD is not helpful on copper ports and is not recommended to be run on them. However, I strongly recommend running Loop Guard configured globally with the spanning-tree loopguard default. Loop Guard can, and should, be run regardless of UDLD, and they can be used both as they nicely complement each other.
My $0.02...
Best regards,
Peter
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-22-2014 12:12 PM
Hi Abhisar,
Regarding your individual decisions: Manually configuring the Root Bridge is a natural thing to do. You should never leave your network just pick up a root switch based on default switch settings.
On end ports, using PortFast and BPDU Guard is a must especially if you are running Rapid PVST+ or MSTP.
Regarding the Root Guard on ports to other switches - this is something I do not recommend. The Root Guard is a protective mechanism in situations when your network and the network of your customer need to form a single STP domain, yet you want to have the STP Root Bridge in your network part and you do not want your customer to take over this root switch selection. In these cases, you would put the Root Guard on ports toward the customer. However, inside your own network, using Root Guard is a questionable practice. Your network can be considered trustworthy and there is no rogue root switch to protect against. Using Root Guard in your own network could cause your network to be unable to converge on a new workable spanning tree if any of the primary links failed, and it would also prevent your network from converging to a secondary root switch if the primary root switch failed entirely. Therefore, I personally see no reason to use Root Guard inside your own network - on the contrary, I am concerned that it would basically remove the possibility of your network to actually utilize the redundant links and switches.
Regarding EtherChannels - yes, you are right, using the on mode can, under circumstances, lead to permanent switching loops. EtherChannel is one of few technologies in which I wholeheartedly recommend on relying on a signalling protocol to set it up, as opposed to configuring it manually. The active mode is my preferred mode, as it utilizes the open LACP to signal the creation of an EtherChannel, and setting both ends of a link to active helps to bring up the EtherChannel somewhat faster.
If you are using fiber links between switches, I recommend running UDLD on them to be protected against issues caused by uni-directional links. UDLD is not helpful on copper ports and is not recommended to be run on them. However, I strongly recommend running Loop Guard configured globally with the spanning-tree loopguard default. Loop Guard can, and should, be run regardless of UDLD, and they can be used both as they nicely complement each other.
My $0.02...
Best regards,
Peter
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-23-2014 02:15 AM
Dear Peter,
Thank you for the valuable suggestions, I will consider during configuration.
Just one clarification,
Last week I faced small outage while creating portchannel, the setup as follows
Cisco 6506 VSS (Root Bridge) connected to other switches in the network.
Parallely, I was working on creating one more separate setup with
Cisco 6807 VSS (Root Bridge) connected to 7 stacks of each two 3850s.
I connected trunk link between 6506 and 6807XL and link came up with 6506 as Root. But when creating Etherchannel (mode on) on one of the Stack of 3850s, outage came on whole network. I got following syslog on 6807
*Nov 18 09:23:42.545: %PM-SW1-4-ERR_DISABLE: channel-misconfig error detected on                                                                                         Po11, putting Te1/5/5 in err-disable state
*Nov 18 09:23:42.621: %PM-SW1-4-ERR_DISABLE: channel-misconfig error detected on                                                                                         Po11, putting Te2/5/5 in err-disable state
*Nov 18 09:23:42.689: %PM-SW1-4-ERR_DISABLE: channel-misconfig error detected on Po11, putting Po11 in err-disable state
*Nov 18 09:23:42.689: %PM-SW2_STBY-4-ERR_DISABLE: channel-misconfig error detected on Te1/5/5, putting Te1/5/5 in err-disable state
*Nov 18 09:23:42.781: %PM-SW2_STBY-4-ERR_DISABLE: channel-misconfig error detected on Te2/5/5, putting Te2/5/5 in err-disable state
*Nov 18 09:23:42.981: %PM-SW2_STBY-4-ERR_DISABLE: channel-misconfig error detected on Po11, putting Te1/5/5 in err-disable state
*Nov 18 09:23:42.981: %PM-SW2_STBY-4-ERR_DISABLE: channel-misconfig error detected on Po11, putting Te2/5/5 in err-disable state
*Nov 18 09:23:42.981: %PM-SW2_STBY-4-ERR_DISABLE: channel-misconfig error detected on Po11, putting Po11 in err-disable state
Can you please help me to find out the reason behind it, as I want to connect the trunk link back.
Thank You,
Abhisar.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-23-2014 07:57 AM
Hi Abhisar it seems like the configuration on the portchannel and physical interface does not match
te1/5/5
te2/5/5
po11
this will cause the interfaces to transition to err-disable state. If you could kindly make sure configs are same, or paste them here if you have doubt, we can peer review for you.
hope this helps
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-23-2014 10:47 AM
Dear Bilal,
