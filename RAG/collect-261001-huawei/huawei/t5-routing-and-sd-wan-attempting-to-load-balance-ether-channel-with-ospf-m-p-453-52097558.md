---
id: collect-261001-huawei/huawei/t5-routing-and-sd-wan-attempting-to-load-balance-ether-channel-with-ospf-m-p-453-52097558
title: "t5-routing-and-sd-wan-attempting-to-load-balance-ether-channel-with-ospf-m-p-453-52097558"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-huawei/t5-routing-and-sd-wan-attempting-to-load-balance-ether-channel-with-ospf-m-p-453-52097558.md
source_anchor: ""
source_lines: [1, 163]
sha256: c67b90108509908bc5fc7985c907fe8537ef1dd858f04004dfe1ec91e13b0c51
---

# t5-routing-and-sd-wan-attempting-to-load-balance-ether-channel-with-ospf-m-p-453-52097558

Attempting to Load Balance Ether-channel with OSPF
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-14-2022 07:26 AM
Hi All,
I have an ASR 9XX box which I have an OSPF interface configured on a 20Gbps etherchannel/port-channel (port-channel A).
I have a local handoff to a customer POP where I provide l2vpn terminations on another 20Gbps etherchannel on the same router. This port-channel (port-channel B) is load-balancing fine without any issues.
I am seeing just the receive traffic load-balanced on the port-channel A which hosts the OSPF connection. The transmit traffic is not load balanced.
I am using the default src-mac setting for load balacing algorithm. My device does not have the ability to use the experimental bits, etc.
Does anyone have any idea or have seen this behavior on their devices before?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-14-2022 07:46 AM
Just another note: 
I can seemingly load balance straight layer2 traffic (from the customer POP to my router) on the port-channel B, but the IP load balancing does not seem to be working on the port-channel A which is hosting the OSPF link. 
Here is what I see when I do the show int port-channel <A> etherchannel command:
show interfaces port-channel <A> etherchannel
 All IDBs List contains 2 configured interfaces
  Port: TenGigabitEthernet0/1/1 (index: 0)
  Port: TenGigabitEthernet0/1/2 (index: 1)
 Active Member List contains 2 interfaces
  Port: TenGigabitEthernet0/1/1
    LACP Mode: Active
  Port: TenGigabitEthernet0/1/2
    LACP Mode: Active
 Passive Member List contains 0 interfaces
 Load-Balancing method applied: flow-based
 Bucket Information for Flow-Based LB:
 Interface:                                Buckets
    TenGigabitEthernet0/1/1:
                           Bucket 0 , Bucket 1 , Bucket 2 , Bucket 3
                           Bucket 4 , Bucket 5 , Bucket 6 , Bucket 7
    TenGigabitEthernet0/1/2:
                           Bucket 8 , Bucket 9 , Bucket 10, Bucket 11
                           Bucket 12, Bucket 13, Bucket 14, Bucket 15
					
				
			
			
				
			
			
				
			
			
			
			
			
			
		
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-18-2022 12:48 PM
Here is another possibly useful link that I found https://networkengineering.stackexchange.com/questions/71939/how-can-you-determine-the-result-of-a-load-balancing-hashing-algorithm-such-as
I am looking for a syntax that is compatible with my ASR 9XX devices.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-14-2022 07:54 AM
Cannot speak for an ASR 9xx, but on many prior gen ISRs, they didn't, as far as I know, support multiple LB options (as often the case with Cisco switches). Possibly, because, on a router, many of Cisco's IGP support L3 load balancing; Cisco's OSPF does (if ECMP).
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-14-2022 08:02 AM
So I suppose the expected method would be to instead use two 10G OSPF uplinks instead that would be ECMP instead of a single LACP/Port-channel with one OSPF link. If OSPF properly load balances it then it shouldn't be an issue, and there should be no need for the larger 20Gbps OSPF link.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-14-2022 08:48 AM
Yes, dual 10G OSPF links, using ECMP, might provide a solution.
However, there's possible "gotchas" in using such an approach too.
First, both Cisco L2 Etherchannel, and, by default, OSPF ECMP, will direct a single flow across just one link. I.e. a single flow will not be able to obtain more than a single link's bandwidth. The exception to this is CEF packet-by-packet distribution; which would allow a single flow to obtain all the links' aggregate bandwidth, per flow. However, because of possible packet re-ordering issues, this is generally not a recommended solution.
Second, believe CEF based OSPF doesn't round-robin flows, believe same flow attributes will always be directed to same link. Unlike Cisco switches which often offer multiple LB choices, your router might not offer any alternatives (although, I recall, some platforms do offer CEF hashing choices - mainly to avoid same path bias, for same flow, on every transit device). However, later PfR variants (unknown if supported on ASR 9xx) support dynamic LBing that integrates with IGPs, such as, I believe, OSPF. Such LBing can attempt to have an equal loading of each link, including using links of different bandwidths.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-14-2022 08:55 AM
Thanks for the response. I am also testing some of the cef settings after doing some reading on that. I found this article/page: https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/mp_l3_vpns/configuration/xe-3s/asr903/17-1-1/b-mpls-l3-vpns-xe-17-1-asr900/b-mpls-l3-vpns-xe-17-1-asr900_chapter_0101.html#reference_DA6510217ADB4930BED473BDAB419E71
as well as this: https://www.ccexpert.us/mpls-network/load-balancing-in-cef.html. 
So far it doesn't look my routers support anything aside from the default ip load-sharing per-destination. Here is where I tried to look for other options using the auto-complete:
ip load-sharing ?
  per-destination  Deterministic distribution
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-14-2022 09:19 AM
Don't know if it would help you, but here's more information about a way to deal with what I called CEF bias, or what Cisco calls CEF Polarization (info). Again, feature might not be available on your platform.
The prior feature intended for multiple, consecutive, transit devices, all doing CEF, per-destination, LBing, but just using different hashing attributes can, sometimes, make a huge difference to your LBing results.
Any command, like IP CEF LOAD-SHARING?
Again, for optimal LBing, hard to beat PfR dynamic LBing (which avoid the issue of one link being hammered, while another doing nothing [assuming there are multiple flows]). Used to be PfR provided as part of non-basic IOS images. Now, I believe, it's often a separately licensed feature. Also again, don't know whether it's available on your platform.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-14-2022 11:32 AM
I do have some of the ip cef load-sharing commands:
#ip cef load-sharing ?
  algorithm  Per-destination load sharing algorithm selection
#ip cef load-sharing algorithm ?
  include-ports  Algorithm that includes layer 4 ports
  original       Original algorithm
  tunnel         Algorithm for use in tunnel only environments
  universal      Algorithm for use in most environments
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-15-2022 10:27 AM
That's possibly good. Among those, one might work better for you than the others, if you pursue OSPF ECMP.
