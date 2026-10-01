---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-802-1x-authentication-failure-with-workstations-on-windows-11-df763b5a
title: "t5-security-sd-wan-802-1x-authentication-failure-with-workstations-on-windows-11-df763b5a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-802-1x-authentication-failure-with-workstations-on-windows-11-df763b5a.md
source_anchor: ""
source_lines: [1, 58]
sha256: b42200950f60eca8a49ea8bb6056c672f53b5458b77584e3a41f2d0abf44303a
---

# t5-security-sd-wan-802-1x-authentication-failure-with-workstations-on-windows-11-df763b5a

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-14-2023 01:23 PM
Currently we are noticing some users that upgraded to Windows 11 are now unable to connect to our switches using 802.1X authentication.
Is there any documentation that is out there to fix this issue. I had to change my switch ports from 802.1X authentication policy to "Open". Not the most secure when doing that.
Also the Windows 11 users are unable to connect to the 802.1X authentication to our corporate WIFI.
We have 175 users on windows 10 all work and connect using 802.1X authentication and the only 5 Windows 11 users in our environment cannot connect using 802.1X authentication.
RADIUS server authentication using Active Directory credentials works fine.
Windows 11 is doing something to affect the 802.1X authentication. That is what we need to find out or is there any documentation fixing this issue.
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-10-2024 08:16 AM
The issue that I found was that Device guard was enabled on windows 11. Disabled device guard via GPO and the issue was resolved. PC had to be rebooted after the port was changed back to 802.1x.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-10-2024 06:08 AM
the solution found in our organization was to insert the certificate of the RADIUS servers via GPO for wired clients, and in the connection profile, mark both servers as trusted
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-11-2025 08:59 AM
Moderator commentary: Looks like there are multiple potential solutions for this problem. I'm going to highlight two of them here, to hopefully help those of you who land on this page from a search engine.
from @DeanN1:
The issue that I found was that Device guard was enabled on windows 11. Disabled device guard via GPO and the issue was resolved. PC had to be rebooted after the port was changed back to 802.1x.
from @andersoninhudes:
the solution found in our organization was to insert the certificate of the RADIUS servers via GPO for wired clients, and in the connection profile, mark both servers as trusted
- « Previous
- 
						
  - 1
  - 2
- Next »
