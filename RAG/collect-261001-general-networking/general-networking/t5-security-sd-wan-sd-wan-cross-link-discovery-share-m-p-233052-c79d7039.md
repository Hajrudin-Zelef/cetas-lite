---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-sd-wan-cross-link-discovery-share-m-p-233052-c79d7039
title: "t5-security-sd-wan-sd-wan-cross-link-discovery-share-m-p-233052-c79d7039"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-sd-wan-cross-link-discovery-share-m-p-233052-c79d7039.md
source_anchor: ""
source_lines: [1, 97]
sha256: 103669ba03c5af8b8dcd6870b62a24670935e492391ef24196e560bc7c45dc9e
---

# t5-security-sd-wan-sd-wan-cross-link-discovery-share-m-p-233052-c79d7039

SD WAN cross link discovery share
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-21-2024 02:08 AM
I accidentally found some undocumented MX behavior while troubleshooting a customer by using iPerf.
Situation:
You have SD-WAN tunnels (hub to hub in this case but it would work the same in hub to spoke) between 2 MX peers both using their 2 WAN ports.  So in this case you have 4 actual tunnels.  WAN-1 to WAN-1, WAN-2 to WAN-2, WAN-1 to WAN-2 and WAN-2 to WAN-1.
The idea is that you can only influence outbound traffic in your SD-WAN rules where you can set WAN-1 or WAN-2 as exit interface for a certain application. But you don't really have control if your traffic leaves WAN-1 on 1 appliance it would be received on WAN-1 on the other appliance. In case you are using private links or a private ISP WAN you would want this to be the behavior. The problem is that the documentation does not really reveal how in some cases traffic could be crossed between WANs.
According to my troubleshooting test it seems that the SD-WAN rules are not just outbound but bi-directional. So you can actually influence on what WAN your MX receives certain traffic.
The test:
I had an iperf server behind the remote MX and an iperf client behind this one.
The networks on one side are all subnets of 10.105.0.0/16 and on the other side are 10.106.0.0/16.
There are rules on each MX that force all private network traffic over WAN2 (by using custom policy network:10.105.0.0/16 and 10.106.0.0/16).  So as expected when monitoring the uplink page on both MX'es you could see both WAN-2 traffic going way up during the test.
Then I added another SD-WAN rule on one side where I matched all traffic going to or coming from: TCP/5201 (default iperf port) to use WAN-1 instead. When I did the outbound test as expected both WAN-1 links were being used for that transfer. However when I did the inbound test WAN-2 was going up while WAN-1!! was going up on this side. So the remote MX peer was only matching the general WAN2 network rule while this side was matching the TCP/5201 rule and traffic was actually crossing.
Background:
I usually don't have customers that use datacenters to link up to Azure but rather direct internet connections.  So I was not used to have many SD-WAN rules to be able to test this and the only real rules I can use is actual server usage between SD-WAN peers.  The iperf test was a nice eye opener.
Takeaways:
So in complex SD-WAN environments where applications are used over the SD-WAN tunnels, be careful with your policies.  If you don't want traffic crossing between your upstream providers which can cause extra delay or bandwidth issues be sure to double check both directions of your SD-WAN policies.
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
04-21-2024 01:05 PM
Interesting. My prior understanding was AutoVPN was meant to be stateful. If traffic goes out WAN1 then the reply traffic should come back via WAN1.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2024 12:11 PM
I believe SD WAN is treated as a single link, so statefulness wouldn't matter between both links.
Of course when you start an iPerf session upstream and after that one another downstream you are technically in another TCP stream.  I haven't used a capture (since that would have been costly over that line) to verify if another source port was used each time.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-22-2024 12:30 AM
Thanks for sharing this! Highly interesting find!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2024 12:11 PM
No problem. I enjoy sharing accidental discoveries.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-22-2024 01:17 AM
Thanks @joey.debra, our WAN is sort of like this and definitely performs the same as yours. We have two MPLS networks with remote sites having one WAN on each. At the central site with terminate the WANs in an MS355 and the MX is connected in single ended mode. I was sure I was seeing the same as you describe, but hadn't had a chance to fully test it. Thank you!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2024 12:13 PM
I would enjoy some pure lab testing time each week to discover intricaties However I neither have the amount of devices required nor the time to fully test that stuff
