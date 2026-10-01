---
id: collect-261001-general-networking/general-networking/t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4-3
title: "t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-general-networking/t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4.md
source_anchor: ""
source_lines: [109, 203]
sha256: 648703a8d82a2cb65322f83ebeeea1ef5ebdf9958944bd8842d3dcf5ed68a3f1
---

# t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4

I dont think its a config issue. I had applied the config on 6807 but did not applied on 3850s and suddenly error occured.
interface TenGigabitEthernet1/5/5
 switchport
 switchport mode trunk
 channel-group 11 mode on
interface TenGigabitEthernet2/5/5
 switchport
 switchport mode trunk
 channel-group 11 mode on
interface Port-channel11
 switchport
 switchport mode trunk
As I checked about the error, I found that if you config mode on on one side and dont config other side withing some time then there are chances of loops.
So I removed mode on and configured as active. Also, my VSL portchannel are on mode on and etherchannel is up, this will create any issue?
Thank You,
Abhisar.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-24-2014 08:16 AM
Hello Abhisar, your VSL requires mode ON, this is by design to my recollection and by default. Please do not change this otherwise your VSS will break.
hth
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-24-2014 11:58 AM
Hi Bilal,
You are correct in that a VSL link should use the unconditional EtherChannel mode on configuration because the links run a different protocol called Link Management Protocol:
http://www.cisco.com/c/en/us/td/docs/solutions/Enterprise/Campus/VSS30dg/campusVSS_DG/VSS-dg_ch2.html#wp1056015
However, Abhisar has not indicated whether the EtherChannel he was creating was related to the VSL link, and according to the outputs and descriptions he has provided, it does not seem that the issue was related to the VSL.
In general, what Abhisar's outputs show is a simple EtherChannel Misconfiguration Guard kicking into action. This mechanism is triggered whenever physical ports of a single Port-channel receive BPDUs sourced from different MAC addresses. In most cases, this is an indication that while this switch is already considering the physical ports to be bundled in a single Port-channel, the switch at the other end of these links is still treating the ports as individual ports. In fact, I am not surprised that the EtherChannel Misconfig Guard has been triggered in Abhisar's case: he has configured the Port-channel interface on one of the switches into the static mode on and before he could configure the other end for the mode on as well, the EtherChannel Misconfig Guard has brought down the entire Port-channel as it was receiving BPDUs from the other switch sourced from different MAC addresses.
There is a general rule for configuring the mode on Port-channels: first, switch off the ports you want to bundle on both switches, then configure both switches for the mode on, and only then turn the ports back on. Never configure the Port-channel in such a way that would allow one switch to treat the links as already bundled while the other is still unconfigured, no matter how short the time would be.
Best regards,
Peter
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-24-2014 11:43 PM
Dear Guys,
Thank you for your comments. As per Peter, yes it happened due to the Etherchannel mode on.
I was configuring etherchannel between Cisco 6807 and Stack of Two 3850, and this Stack of 3850 was root and when I configured mode on on 6807 both links went to forwarding and got looped. Now, I made 6807 as root and Etherchannel is on LACP.
Guys, request if you can help to sort out one more issue.
Now, I have two separate setup.
(Attached file)
One is 6506(VSS) connected to other access switches(5 4506 and 2 3750s) and this is working fine. The VTP domain is CISCO.
Other is Cisco 6807(VSS) connected to other switches(7 stacks of two 3850s and 3 4506) and this connected to each other via etherchannel. There is no traffic on this setup. There is no VTP domain configured. I guess this should work.
Now, when I connected the trunk link (Etherchannel LACP of 2 1G SFP)between both VSS the network gets little slow, specially browsing traffic as ping to some servers is 1ms. Why this has happened? After connecting both the setup Cisco 6506 is root as priority defined for this is 24576 and for 6807 is 28672.
Is this because 6807 bacame secondary root? Or if I keep 6807 priority to default then Stack of 3850s will become root and will connect the link between both VSS then 6506 will become as root, this can help?
If I configure VTP domain on new setup as CISCO and all switches as client, does this affect anything on network?
Can you please help what might be reason for slow.
Thank You,
Abhisar.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2014 01:06 AM
Hi Abhisar,
Let me go over some of your statements as I need more clarification.
Other is Cisco 6807(VSS) connected to other switches(7 stacks of two 3850s and 3 4506) and this connected to each other via etherchannel. There is no traffic on this setup. There is no VTP domain configured. I guess this should work.
Regarding VTP, is this part of network configured for VTP Transparent or VTP Off mode, or are the VTP settings left on default values? I am asking because leaving the VTP settings entirely unconfigured and then connecting the 6807 via a trunk link to an existing switch that runs VTP may cause the 6807 to adopt the VTP domain name and download the VLAN database, and subsequently propagate it to the attached 3850s and 4506s. This will certainly happen if the VTP settings on the 6087 were left at their default values, and you are not using VTP password protection in your existing VTP domain.
Is this because 6807 bacame secondary root? Or if I keep 6807 priority to default then Stack of 3850s will become root and will connect the link between both VSS then 6506 will become as root, this can help?
This is certainly not an issue of the secondary root bridge. In fact, there is no such role in STP. The name "secondary root bridge" is just a saying that tells which switch is going to become the root switch if - and only if - the current root switch fails. However, until that happens, the switch configured as a "secondary root bridge" is just an ordinary switch having no special role with respect to STP.
In fact, I am surprised that by connecting your two VSS pods, the latency in your networks gets increased. I am afraid that this requires further investigation. I personally suspect some kind of flooding present in your network that potentially results into increased delays. However, to investigate this issue, you should try monitoring the traffic and the traffic loads on all ports to see if there is indeed any kind of unexpected load present. With the information we have up to now, no conclusion can be drawn yet.
Best regards,
Peter
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2014 01:47 AM
Dear Peter,
Thank for the comments.
I will only connect the Cisco 6807 to 6506. I will not connect any stack connected to 6807, I will shutdown the all the etherchannels.
I am not sure about the traffic, I will troubleshoot if something is wrong.
Thank You,
Abhisar.
