---
id: collect-261001-meraki/meraki/t5-switching-meraki-ms210-and-c9300x-lacp-1-port-gets-suspended-s-m-p-223605-hig-1a12196d
title: "t5-switching-meraki-ms210-and-c9300x-lacp-1-port-gets-suspended-s-m-p-223605-hig-1a12196d"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-switching-meraki-ms210-and-c9300x-lacp-1-port-gets-suspended-s-m-p-223605-hig-1a12196d.md
source_anchor: ""
source_lines: [1, 135]
sha256: 0cc53f7c8942dac85eab95e4b9c97d2267d14b84a8987692690a95a399d0e541
---

# t5-switching-meraki-ms210-and-c9300x-lacp-1-port-gets-suspended-s-m-p-223605-hig-1a12196d

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2024 07:39 PM
Hi Experts,
Do you have any idea why 1 port of C9300X gets suspended when configured using LACP protocol?
In MS210-48, I just bundled port 49 and 50 wherein I confirmed that AGGR/0 is created.
Configuration of C9300X is pretty basic.
int range twe 1/0/1, twe 2/0/1
switchport mode trunk
channel-group 101 mode active
!
int port-channel 101
switchport mode trunk
!
Is there someone with the same concern?
TIA.
Mike
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
01-28-2024 08:00 PM
Can you 100% confirm that the ports are connected correctly?
LACP is designed to suspend the ports if they're misconfigured or miscabled.
You can look at the LACP flags or enable debugs on the C9300X for LACP events to understand where the issue is.
Or perform a packet capture on the Meraki side to confirm the source and destination of LACP packets
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2024 08:00 PM
Can you 100% confirm that the ports are connected correctly?
LACP is designed to suspend the ports if they're misconfigured or miscabled.
You can look at the LACP flags or enable debugs on the C9300X for LACP events to understand where the issue is.
Or perform a packet capture on the Meraki side to confirm the source and destination of LACP packets
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2024 09:08 PM
Hi @Brash ,
Thank you for the response.
I'm pretty sure that it is cabled and configured correctly.
That's why I'm scratching my head on this issue
Maybe I can try some debugs on C9300X tomorrow when I'm onsite.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2024 10:50 PM
I have had issues with LAG config before under certain circumstances. For example if the MS is using the C9300X as the uplink towards the internet. I'm not 100% sure why the config does not take in this scenario, but i suspect that there is a race condition in play. Chicken and egg so to speak.
Try removing the config so you are only using a single uplink, configure meraki and wait a few minutes, then configure the C9300.
Also, should go without saying, but verify cabeling and ports. Check LLDP on both switches etc.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2024 11:04 PM
Hi @martin-lystad ,
Thank you for the response.
This is the exact procedure that we made.
Have the MS connect to internet using port 52 then configure 49-50 as aggregate.
Wait for around 30 mins, then change the connection by removing port 52 and use 49-50.
During the first few minutes I can see that it's properly working but when I added and connected other IDF switches with the same settings going to C9300X, I noticed that 1 of the port gets suspended.
When I tried to do "shut / no shut" from C9300X port-channels that is previously working, 1 of the port also got suspended.
Currently, MSs can connect to internet using aggregation link 49 but with 1 port (50) suspended in C9300X.
Regards,
Mike
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2024 11:33 PM
Suspensions are only on the Catalyst side and have to do with either configuration mismatch or operational mismatch.
So if both ports are configured with the same layer 1 and 2 things (being speed/duplex for layer 1 and trunk config at layer 2) then you will also have to look at the operational state. In some cases in autonegotiation on copper if you have a 1 gig and a 10 gig in the same channel then you will have a suspended member.
I'm not a 100% sure if the suspended state can also come from a member mismatch (if the switch at the other end uses a different local MAC on each member).
It is best to enable terminal monitor on the catalyst if you have SSH'd into it and disable the port and re enable it. You may have the issue listed in the syslogs that appear when it suspends again.
There can however be some weird bugs with port channels on Catalysts where you could temporarily remove the members from the channel and re add them to fix.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-31-2024 12:19 AM
Hi @joey.debra , @Brash , @martin-lystad ,
Thank you very much for all your inputs in this case.
I managed to resolve the problem.
Basically, most of the "aggregation" that I configured in Meraki Dashboard isn't there when I checked the settings yesterday. It's just weird that the first switch that I successfully tested for link aggregation also encountered issue hours after it has been established.
Since I have less experience with Meraki, I tend to have a test first before proceeding with other equipment.
In the end I re-configured most of the Meraki switches for aggregation and from there LACP worked well.
Regards,
Mike
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-31-2024 12:26 AM
I just noticed also that Meraki login page is new, not sure when the update came in.
I hope it didn't add up to the missing aggregation configuration that I made during the time I'm configuring the switches
