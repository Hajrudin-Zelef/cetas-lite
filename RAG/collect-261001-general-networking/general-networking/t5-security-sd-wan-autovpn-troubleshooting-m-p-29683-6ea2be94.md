---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-autovpn-troubleshooting-m-p-29683-6ea2be94
title: "t5-security-sd-wan-autovpn-troubleshooting-m-p-29683-6ea2be94"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-autovpn-troubleshooting-m-p-29683-6ea2be94.md
source_anchor: ""
source_lines: [1, 147]
sha256: 6e93d252f2cdd93fd2fa32e04293fb8dd4a07b6cffc308ea4ad41a51151078c8
---

# t5-security-sd-wan-autovpn-troubleshooting-m-p-29683-6ea2be94

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-10-2018 01:45 PM
Continuing issue for us is a AutoVPN connection that literally can bounce a dozen times a day. But Meraki dashboard, even though it sends constant alerts saying VPN is down/up, says the connection is solid. Location is a DSL site. Question is, 1. can double NAT cause issues with the AutoVPN? 2. Does anyone else have experience with a MX on a DSL connection and have any issues? 3. Is there ANYTHING at all I can use to try and diagnose what is going on with this VPN connection?
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
10-13-2018 11:43 PM
Hey @Tim Bisel,
You want to set the uplink statistic up pointing to the public IP address of the remote end, not the LAN side.
That way you will monitor WAN-WAN connectivity rather than over the VPN.
Thanks!
G
Appreciate who helps and be respectful of every opinion and every solution offered.
Share the love, especially the Meraki one!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-10-2018 01:59 PM
Hey TimBisel,
Are you seeing anything in the VPN status of the MX, at the DSL site and the other end? I haven't heard of the type of connection affect VPN, as you have the correct ports open upstream (Verify your specific network's connection to the Meraki VPN registry by going to Help > Firewall Rules). To answer question 1 as well, double NAT can cause flapping in previous cases.
Finally I would recommend looking at the things listed here and see if you experience the same issues:
Cheers,
Nick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-10-2018 02:32 PM
I have a very large amount of VPN status change alerts in both sides of the vpn but no details on why. I think I might be running into issues with the onsite router (Verizon DSL) I am thinking that if it was setup with the default 192 address and still has NAT translation enabled that maybe the firewall rules were left default as well and could be causing it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-11-2018 04:21 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-11-2018 05:24 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-12-2018 07:15 AM
Hi @Tim Bisel,
I will echo @nicholaswu8 response, and I would recommend to follow the article to troubleshoot this too.
There is, however, another action you might want to take to understand if there is a break in the communication between the two sites: you can navigate in Security appliance > Traffic shaping and set up an Uplink statistic with destination the public IP address of the remote Meraki MX.
You can then see the information from Security Appliance > Appliance status > Uplink and this would greatly help you to understand if there's any drop occurring over the WAN or if the issue is just on the VPN tunnel.
If you see frequent disconnections on the event log to the VPN registry, give us a call as we might be able to take some remediation steps.
Thanks!
Giacomo
Appreciate who helps and be respectful of every opinion and every solution offered.
Share the love, especially the Meraki one!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-12-2018 12:26 PM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-13-2018 11:43 PM
Hey @Tim Bisel,
You want to set the uplink statistic up pointing to the public IP address of the remote end, not the LAN side.
That way you will monitor WAN-WAN connectivity rather than over the VPN.
Thanks!
G
Appreciate who helps and be respectful of every opinion and every solution offered.
Share the love, especially the Meraki one!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-15-2018 05:43 AM
Ok that makes more sense, thanks.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-17-2018 06:22 AM
One of out sites had some old network equipment die so this is the first I was able to look at it. But does seems to has some short periods of packet loss. But no notifications from these drops and the last alerts I got was at 3am and seems to line up with one of the peaks of packet loss. But there are some other points that have packet loss and I did not receive an alert. Is there a % threshold on when a link is considered down? Anyone know how much of an overhead VPN takes on the bandwidth, thought maybe our speeds just slow down to the point VPN can't stay connected.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-17-2018 01:27 PM
Happened twice more, looks like it is packet loss hitting ~25%. So it must be some threshold for packet loss that is sending the alarm. The ping helped, thanks everyone.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-17-2018 07:27 AM
We've got several AutoVPN sites on DSL/Cable along with several more on enterprise fiber and have seen this same VPN tunnel flapping behavior on occasion. I've always assumed it was the lifetime timer expiring and tearing down the tunnel due to an absence of interesting traffic, but never examined it any further since it didn't affect operations.
While this may not apply to a Meraki AutoVPN connection, here's a Meraki document on IPSec VPN Lifetime settings that may help out if you've got a connection to a non-Meraki peer.
