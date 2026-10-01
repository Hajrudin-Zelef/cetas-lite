---
id: collect-261001-meraki/meraki/t5-security-sd-wan-migration-to-full-meraki-solution-m-p-44290-da9cc0f0
title: "t5-security-sd-wan-migration-to-full-meraki-solution-m-p-44290-da9cc0f0"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "research"]
source: docs/RAG/collect-261001-meraki/t5-security-sd-wan-migration-to-full-meraki-solution-m-p-44290-da9cc0f0.md
source_anchor: ""
source_lines: [1, 81]
sha256: cdd41e21cb91f918673c55740cab7ea14b24b84edc1f7e598420126872512a74
---

# t5-security-sd-wan-migration-to-full-meraki-solution-m-p-44290-da9cc0f0

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-18-2019 03:47 AM
Hi everyone , 
We are going to deploy a Meraki solution for an organisation that has 7 branch offices and one central office . They have existing network using other devices and ongoing services they use 24/7 . They have two providers at each site , and my first plan was to use the failover link only for the Meraki deployment , one branch at a time , testing the connectivity and services , before fully migrating the whole network . Any experience and tips one this matter ?  
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-18-2019 06:37 AM
I would say that the way you presented would be a good way to do it. That way you're slowly migrating over while checking that all your connections/services are working correctly. Kudos! Only thing I would say is to make sure to keep your old equipment ready just in case you need to go back and check things or if something doesn't go right.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-18-2019 06:37 AM
I would say that the way you presented would be a good way to do it. That way you're slowly migrating over while checking that all your connections/services are working correctly. Kudos! Only thing I would say is to make sure to keep your old equipment ready just in case you need to go back and check things or if something doesn't go right.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-18-2019 08:30 AM
When my old company migrated away from sonicwalls we stood up an MX in our HQ, added a few users, added the routes we needed for access and tested. Then went to a site, stood up a MX on the backup circuit as you described and moved a few users over just as we did at HQ. Tested a good 3 weeks and then moved the whole HQ over, then the 1st site. Once we got the bugs worked out, we went site by site swapping the whole sites network out in a day, and tested for a week.
Biggest thing we learned was you have to put actual business traffic across the test network. Web browsing is not business traffic (not for us anyway) and then only true test is to do the cut over and see what doesn't work. fortunately we were able to do all the initial testing after hours for HQ and at a small site with little traffic.
Hope your project goes well and you enjoy meraki as much much as we did. It really is almost always simple with meraki. The challenges come when you try to integrate 3rd parties
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-18-2019 09:03 AM
Hi,
I would use the main link - and leave the backup for the existing, willing to be decommissioned Branch/HQ when Migration comes to play, especially that should prove also the backups are working, so when your migration will move for the backup links to move, it i already proven.
Would do a phased and detailed migration Plane what corresponding with your current Technologies and Sections and the future ones, like:
"Existing config compare with the future Solution / one by one per each technology or protocol"
WAN
Cisco Routers - Meraki MX - Main Link move / adding 4G stick for check SD-WAN
DNS / Vlans / DHCP / ACL / Access restrictions / L3-L7 "policies" / Blacklisting / Whitelisting
then
VPN what allowed and what not
then moving towards down the Stack to Switches
Cisco Switches - Meraki MS
DNS/ Vlans/ ACL
then to AP`s / Cameras / to Endpoints.
All good ? - Backup link move / full decommission of existing
PPDIOO with Meraki , will cost more time on research for you first, but as every Engineer knows: measure twice , do once, stand back and optimize until its perfect.
Happiness !
kajcsu
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-19-2019 12:06 AM
Thanks for your advices , will let you know how it goes
