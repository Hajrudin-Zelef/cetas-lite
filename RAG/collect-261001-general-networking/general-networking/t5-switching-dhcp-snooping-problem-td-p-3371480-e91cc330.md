---
id: collect-261001-general-networking/general-networking/t5-switching-dhcp-snooping-problem-td-p-3371480-e91cc330
title: "t5-switching-dhcp-snooping-problem-td-p-3371480-e91cc330"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-general-networking/t5-switching-dhcp-snooping-problem-td-p-3371480-e91cc330.md
source_anchor: ""
source_lines: [1, 189]
sha256: 841f684db291f60608201b7907dbf430df7fd08b630188b087fb44bb23f7ed57
---

# t5-switching-dhcp-snooping-problem-td-p-3371480-e91cc330

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-23-2018 02:51 PM - edited 03-08-2019 02:46 PM
I have enabled DHCP Snooping on all my new Cat 9300 access switches using the below commands
ip dhcp snooping vlan 100,110,120,130
no ip dhcp snooping information option
ip dhcp snooping
I then trusted the uplink interfaces that connect directly to the core. (ip dhcp snooping trust)
The DHCP server is connected to the core, on vlan 120, and is a MS Windows 2008 R2 server
The Voice VLAN is 110, and all clients plug their computers into the computer port on the back of the phone
Everything worked fine in my testing but today there were a ton of users on various VLANs who could not get IP addresses
To resolve, I turned off snooping for all switches, and everything was fine. I turned it back on for 1 VLAN to try to troubleshoot the issue, but I did not have a problem. It was working as expected.
Can't figure this one out. Any ideas?
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Other Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-13-2019 06:02 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-23-2018 04:19 PM
Hi,
Waht version of ios is running on the switches? Are there any error in the logs? Is this happening randomly or does it happen when the lease is being renewed?
Thanks
John
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-23-2018 05:26 PM
I didn't see anything in the logs. I enabled the feature Friday night, so nobody was really logging in until this morning.
It didn't seem to affect everyone, but it's possible the machines that were working never restarted
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-23-2018 06:21 PM
we are experiencing similar issue, except it appears to be affecting our IP phones. did you find any solution to this yet?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-23-2018 07:20 PM
are you also using Cat 9300's?
I don't have a solution yet. In my case, ip phones were already up, so they were fine, but once I tried restarting one, it would no longer get an ip address.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 02:43 PM
Anyone have any suggestions? TAC's advice is turn it back on and wait for it to happen again
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-09-2018 12:21 PM
We are having the same issue with our 9300s. When the switches are first configured DHCP snooping is configured and seems to work. Then it stops working.
Turning it off and back on doesn't see to fix it.
Has anyone found a solution yet?
Lynne
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-27-2018 06:17 PM
Hi All,
So I have been battling both of the DHCP snooping bugs and although v16.9.1 fixed the Ether-Channel issue I am still seeing the other issue that starts up every few days to a week where users are unable to get an address from the DHCP server. I opened a new case and sent them the Show tech but I have also figured out a workaround until they get a new version.
I know it's not ideal using a kron and some of you hate them, but it's what I have for now so I can have my cake and eat it too. If you have a better solution please let me know.
My other option was to disable DHCP snooping completely, and I was not about to do that because we have Software/Hardware engineers here that have to have isolated networks but be able to reach the internet so they use small home NAT routers. When we didn't have snooping enabled, we would have an outage once a week because those smart people can't differentiate the WAN from the LAN ports nor the Labels above the WALL Jacks at their desk that say LAB and CORP.
Anyway, Here is what I did for the Workaround and it seems to work for now. Basically every day at 03:00UTC it runs the command to clear all of the DHCP snooping bindings. You may need to do this more often depending on how active you network is with devices moving around but this works for me. I will definitely remove this once there is a fix for the issue and suggest you do too. You can use whatever name you want I just used "Clear-Snoop".
kron occurrence DAILY-WrMem at 3:00 recurring
 policy-list Clear-Snoop
!
kron policy-list Clear-Snoop
 cli clear ip dhcp snooping binding *
exit
!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 02:54 PM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 03:01 PM
yes, they are. I will check the bug.
Thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 03:10 PM
Thanks so much for this link. This describes my issue exactly. When I trusted the ports, I configured the trust on the port channel itself, which automatically added it to the individual interfaces.
Sounds like the safest option is to upgrade to 16.6.3
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 03:26 PM
That's good to know. I tried the latest 16.6 ios on the 3650 platform a few months back and ran into similar issues with dhcp snooping (couldn't find a specific bug for it though).
Cheers
Andy
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 03:30 PM
Cisco TAC told me we are hitting the same bug->
https://bst.cloudapps.cisco.com/bugsearch/bug/CSCvi39202/?reffering_site=dumpcr
for the fix, they told me the following->
Regarding the final fix to these bug as I mentioned there are two releases that are documented will include the fix, they are:
16.6.4 estimate date to be available is Jun 29th
16.9.1 estimate date to be available is July 18th.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 08:00 PM
The link was last updated yesterday and shows as fixed?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-26-2018 10:59 AM - edited 04-26-2018 10:59 AM
I received same update from Cisco that the bug will be fixed in 16.6.4.
Is there a way to search for bugs by model, instead of by bug ID?
Also, I was wondering where do you all store the dhcp snooping database? I left it at the default, but maybe I should store on flash?
