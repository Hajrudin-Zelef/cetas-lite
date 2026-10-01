---
id: collect-261001-general-networking/general-networking/t5-wireless-issue-with-wireless-layer-3-roaming-and-mtu-workaround-and-m-p-15124-463002a3
title: "t5-wireless-issue-with-wireless-layer-3-roaming-and-mtu-workaround-and-m-p-15124-463002a3"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-wireless-issue-with-wireless-layer-3-roaming-and-mtu-workaround-and-m-p-15124-463002a3.md
source_anchor: ""
source_lines: [1, 104]
sha256: 3620216c03d282c1b1fa8c2ae8ad646170d17dfa93a5608d583cbe50779a1930
---

# t5-wireless-issue-with-wireless-layer-3-roaming-and-mtu-workaround-and-m-p-15124-463002a3

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-01-2022 12:27 PM
UPDATE: Fixed code levels confirmed. See below.
There is a known issue with Meraki MR wireless code with the MTU changing when L3 Roaming. Using Ping on macOS, I am seeing an MTU of 1448 on WiFi. It should be 1500 to match the interface MTU on the client machines. I'm posting this to save others time and aggravation.
Symptom Of The Issue
Pinging on a SSID that uses Layer 3 Roaming. The WAP is actually telling us its MTU.
Generate 1500 byte Ping packets - look at the 3rd line below
% ping -c 2 -D -s 1472 www.boeing.com 
PING www.lb.boeing.com  (130.76.184.142): 1472 data bytes
556 bytes from 10.34.130.34: frag needed and DF set (MTU 1448)
Vr HL TOS Len ID Flg off TTL Pro cks Src Dst
4 5 00 dc05 c58d 0 0000 40 01 9294 10.34.152.2 130.76.184.142
Request timeout for icmp_seq 0
556 bytes from 10.34.130.34: frag needed and DF set (MTU 1448)
Vr HL TOS Len ID Flg off TTL Pro cks Src Dst
4 5 00 dc05 a50c 0 0000 40 01 b315 10.34.152.2 130.76.184.142
--- www.lb.boeing.com  ping statistics ---
2 packets transmitted, 0 packets received, 100.0% packet loss
Lowering Ping packet size to 1448 - and it now works
% ping -c 2 -D -s 1420 www.boeing.com 
PING www.lb.boeing.com  (130.76.184.142): 1420 data bytes
1428 bytes from 130.76.184.142: icmp_seq=0 ttl=49 time=45.630 ms
1428 bytes from 130.76.184.142: icmp_seq=1 ttl=49 time=43.216 ms
--- www.lb.boeing.com  ping statistics ---
2 packets transmitted, 2 packets received, 0.0% packet loss
Note: For the pings above, ICMP has a 28 byte overhead. (1472 + 28 = 1500 bytes.)
Code levels impacted: Issue was found in MR 27 code and is still present in MR 28.6.1. Meraki Engineering is aware of this issue.
Fix Incoming: The next 28.X firmware release is likely (but not yet confirmed) to have that issue resolved. No ETA for the next firmware release. Thanks to the efforts of our SE, fix is confirmed in 28.7 and 29.X beta. Both should be available in the July timeframe.
Workaround from Meraki Support: Use Bridge Mode for now if it is not absolutely necessary to have L3 roaming.
If you are having this issue, please open a case with Meraki Support so that they know you are impacted by the issue.
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
06-04-2022 12:09 PM
Just to make sure readers see this: Thanks to the efforts of our SE, fix for Layer 3 Roaming is confirmed in 28.7 and 29.X beta. Both should be available in the July timeframe.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-01-2022 05:18 PM
To your knowledge , does this affect all the APs under 28.X code ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-02-2022 12:01 AM
I know it impacts MR56s, but I do not know if it impacts other MR WAPs. 
Speculation based on incomplete info: Probably platform-independant.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-02-2022 10:30 AM
Good work Jason!
I see you also are feeling the "fun" of this 28.6 code. Hope you and the family are doing well.
I think in the other posts you mentioned the trigger for this bug was a new Internet connection. Was it a Ziply connection by chance?
Thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-02-2022 11:11 AM
Long time no see! Look me up - my cell number hasn't changed since we last talked.
It was not Ziply. The ISP in question got the circuit in for us in a hurry, so I'm leaving them unnamed.
That said, jumping from 27.7.1 to 28.6 on MR33s made a HUGE difference in network performance with a Layer 2 Bridging configuration. (From almost unusable to pretty good.)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-04-2022 12:09 PM
Just to make sure readers see this: Thanks to the efforts of our SE, fix for Layer 3 Roaming is confirmed in 28.7 and 29.X beta. Both should be available in the July timeframe.
