---
id: collect-261001-general-networking/general-networking/t5-switching-manually-add-mac-to-dhcp-snooping-database-td-p-2247387-004849a4
title: "t5-switching-manually-add-mac-to-dhcp-snooping-database-td-p-2247387-004849a4"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-manually-add-mac-to-dhcp-snooping-database-td-p-2247387-004849a4.md
source_anchor: ""
source_lines: [1, 128]
sha256: 53d55a11c72481997aed3261eff29bd7e54f0dd3c84d1cfb2dde4b7776376f7e
---

# t5-switching-manually-add-mac-to-dhcp-snooping-database-td-p-2247387-004849a4

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-29-2013 03:25 AM - edited 03-07-2019 01:05 PM
Hello,
I have enabeld DHCP snooping on a VLAN and left it a few days an enabled Dynamic ARP Inspection and all is good apart from one user who has a static IP so he doesn't appear in the DHCP snooping bindings list, but his port goes into err-disabled. I guess I can trust the port, but can you add static entries, I guess soemone coudl spoof their mac and assign the correct static IP?
Thanks
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
04-29-2013 05:26 AM
Hello Andy - yes it would. Pauls example is correct.
Please see this: http://ccietobe.blogspot.co.uk/2009/01/dynamic-arp-inspection-with-non-dhcp.html
But just to add I would do this instead:
ip arp inspection filter TST vlan x 
Hope this helps.
Please rate useful posts and remember to mark any solved questions as answered. Thank you.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-29-2013 04:27 AM
Hello Andy
The DAI filter list is inspected before the snoop d/b so in case, so to answer your question, yes you can add a static mac address.
arp access-list TST
permit ip host x.x.x.x  mac host 0000.1111.2222
ip arp inspection filter-list TST
res
Paul
Please don't forget to rate any posts that have been helpful.
Thanks.
Please rate and mark as an accepted solution if you have found any of the information provided useful.
This then could assist others on these forums to find a valuable answer and broadens the community’s global network.
Kind Regards
Paul
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-29-2013 05:01 AM
Thanks Paul, so will this work along side the "
ip arp inspection vlan x" that is running?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-29-2013 05:26 AM
Hello Andy - yes it would. Pauls example is correct.
Please see this: http://ccietobe.blogspot.co.uk/2009/01/dynamic-arp-inspection-with-non-dhcp.html
But just to add I would do this instead:
ip arp inspection filter TST vlan x 
Hope this helps.
Please rate useful posts and remember to mark any solved questions as answered. Thank you.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-29-2013 07:24 AM
That worked great, the url was perfect and yes you need to use:
ip arp inspection filter TST vlan x
and not
ip arp inspection filter-list TST
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-29-2013 07:31 AM
One other thing, in the "show mac address-table" there are dynamic and static. Dynamic are learned by arp and static are manually entered. I see my port-security enabled ports are static, so I guess these are just being seen as my manually everything the mac address?
interface FastEthernet0/13
switchport mode access
switchport port-security
switchport port-security aging time 1
switchport port-security aging type inactivity
ip arp inspection trust
spanning-tree portfast
spanning-tree bpduguard enable
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-29-2013 07:49 AM
Yes, when you have port security this way, im sure i remember studying that the mac addresses are then written to the configuration of the switch. Therefor it is static.
You could try a show run and see if this is true, its been a while and I can't remember.
Hope this helps.
Please rate useful posts and remember to mark any solved questions as answered. Thank you.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-29-2013 08:13 AM
They don't appear in the running config, but I know they will if they are sticky ports and are saved to the startup config if I wr mem.
