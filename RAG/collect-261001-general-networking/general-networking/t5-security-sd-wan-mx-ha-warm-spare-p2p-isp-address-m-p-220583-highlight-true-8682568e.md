---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-mx-ha-warm-spare-p2p-isp-address-m-p-220583-highlight-true-8682568e
title: "t5-security-sd-wan-mx-ha-warm-spare-p2p-isp-address-m-p-220583-highlight-true-8682568e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-mx-ha-warm-spare-p2p-isp-address-m-p-220583-highlight-true-8682568e.md
source_anchor: ""
source_lines: [1, 82]
sha256: f705b30566f9071c17f1be46b19c4a477307d781ed72df1be820b9fb250acceb
---

# t5-security-sd-wan-mx-ha-warm-spare-p2p-isp-address-m-p-220583-highlight-true-8682568e

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-03-2024 02:53 PM
Hi,
We recently purchased 2 MX units and some MS switches and wireless units. I am interested in connecting the MX in HA or WARM SPARE. The problem is that my ISP gives me one exit port from the modem and P2P addresses between us. I try to connect through my MS switches but the WARM SPARE does not synchronize, I have a small drawing of the network (until I create everything in Visio). Did I connect correctly? It should work with a P2P address? Is there another better configuration?
ISP address 5.29.116.169/31
MX address 5.29.116.168/31
Thanks
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
01-03-2024 03:28 PM
As mentioned MX HA requires unique WAN IPs and reachability to dashboard from each MX.
So you could get multiple IPs from one provider and connect each MX to the provider equipment or use multiple providers with one IP from each and each connected to only one MX. Either way supports HA.
Of course with the second example of two providers only connecting individually to each MX you cannot use a virtual IP for the WAN interface(s). And failover isn't as seamless as failing from the primary to spare MX would mean a new WAN IP is in use.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-03-2024 03:02 PM
Each MX in a HA configuration requires a unique WAN IP. It appears your ISP needs to give you more IPs if they only allow it to work like that with P2P IPs instead of just handing out private IPs on the inside LAN of that modem.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-03-2024 03:16 PM
So in this way the HA will not work. Do I need more IP addresses? If I purchase another line And connect each line to the MX, the HA will work ?
Thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-03-2024 03:28 PM
As mentioned MX HA requires unique WAN IPs and reachability to dashboard from each MX.
So you could get multiple IPs from one provider and connect each MX to the provider equipment or use multiple providers with one IP from each and each connected to only one MX. Either way supports HA.
Of course with the second example of two providers only connecting individually to each MX you cannot use a virtual IP for the WAN interface(s). And failover isn't as seamless as failing from the primary to spare MX would mean a new WAN IP is in use.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-04-2024 12:32 PM
Hi,
Thanks for your help, but I want to understand something else, we have another branch with a similar network configuration with FORTI equipment and everything works, the FORTI is connected in HA ACTIVE PASSIVE. The MX doesn't know how to work in a similar way to FORTI?
Thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-04-2024 01:09 PM
@shlomoi different vendors use different HA methods. Meraki need at least two IPs, one for each device, on the WAN side.
