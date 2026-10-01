---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-mx-105-bgp-deployment-as-edge-y-device-m-p-224535-highlight-t-982d954f
title: "t5-security-sd-wan-mx-105-bgp-deployment-as-edge-y-device-m-p-224535-highlight-t-982d954f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "regulation"]
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-mx-105-bgp-deployment-as-edge-y-device-m-p-224535-highlight-t-982d954f.md
source_anchor: ""
source_lines: [1, 113]
sha256: 221586e8ecd8125ec1f238fd92751950428daa2e934f5a8c460af7735b1f652c
---

# t5-security-sd-wan-mx-105-bgp-deployment-as-edge-y-device-m-p-224535-highlight-t-982d954f

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-02-2024 08:49 AM
Hi,
I work for a state agency where we must procure our WAN connections via another state agency that handles a lot of services for the state. We also aren't allowed to roll our own vpn or SD-WAN due to this regulation which seems like it may be a sticking point for using the MX. The rest of my network infrastructure is comprised of Meraki switches so I looked into the MX when I needed to replace the cisco layer 3 switch providing our connection to their network and our WAN connection. CDW assured me it would work for my situation and so I purchased some for our different office locations. It looks like routed mode would be best for my situation but I read that this would cause the MX to NAT traffic to it's WAN connection interface which wouldn't be good. I'm a bit lost as how to proceed and will be opening a ticket with Meraki but thought I would ask here as well.
My topology will be
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
02-02-2024 01:52 PM
Routed mode BGP is new in 18.2 and does support peering on the WAN and LAN. Still basically in beta and requires Support to enable it today.
Some design examples are shown here https://documentation.meraki.com/MX/Networks_and_Routing/Border_Gateway_Protocol_(BGP)#Scenario_4:_Routed_mode_-_eBGP_peer_on_WAN_of_spoke
I'd recommend you talk through the design requirements/goals with your Meraki SE. From what I gather from this thread I'm not sure you were given proper guidance from the partner or perhaps they didn't understand your requirements.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-02-2024 08:53 AM
Here is some information.
- Mixed operation of MX14 and MX15/16 networks is NOT supported
- IPv6 is currently only supported in One-Armed Concentrator mode
- To enable routed mode BGP there are two requirements:
  - MX must be running MX 18.205 or higher firmware,
  - You will also require the MX to be put into No-NAT mode, please reach out to support to have this enabled.
https://documentation.meraki.com/MX/Networks_and_Routing/Border_Gateway_Protocol_(BGP)
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-02-2024 09:34 AM
Thanks so much for the info on the No-NAT mode! That is very helpful I appreciate it. Luckily not running ipv6 so that isn't an issue for this deployment. I'll be running one MX-105 at our main office and MX-85s at our satellite offices. We also have two data centers and Azure which we connect to as well just for context. Any ideas on how to configure this on the MX to work as I wish?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-02-2024 09:35 AM
Specifically I was wondering about the Auto-VPN portion that needs to enabled in order to use BGP. Is there a way to leverage this without actually routing any traffic to the vpn? Also if the vpn tunnels goes down will BGP also go down for the rest of the routing? Thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-02-2024 09:40 AM
The documentation has some scenarios, have you checked them out?
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-02-2024 09:52 AM
Yes I've been over them and all the documentation I could find but none of the scenarios seem that close to mine. I see that the routed mode scenario 4 is likely the closest to my scenario but I won't be using the vpn at all. My LAN will be connected via one of the LAN ports and all that traffic needs to route through the MX device. All the LAN traffic will be routed to the MX via our Palo Alto on site. The WAN connection will be connected to 4 BGP peers for our uplink/WAN. So really all the MX needs to do is have a static route to the Palo Alto for all 10.2.0.0/16 (LAN space) and handle route distribution on the WAN connection. I may introduce more features later on if the MX handles the traffic as I expect it should but I need to get it's main function operational.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-02-2024 10:02 AM
To be honest, as far as I know, just like OSPF, BGP is just for routing traffic within SD-WAN, so I believe that MX will not meet what you want.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-02-2024 01:52 PM
Routed mode BGP is new in 18.2 and does support peering on the WAN and LAN. Still basically in beta and requires Support to enable it today.
Some design examples are shown here https://documentation.meraki.com/MX/Networks_and_Routing/Border_Gateway_Protocol_(BGP)#Scenario_4:_Routed_mode_-_eBGP_peer_on_WAN_of_spoke
I'd recommend you talk through the design requirements/goals with your Meraki SE. From what I gather from this thread I'm not sure you were given proper guidance from the partner or perhaps they didn't understand your requirements.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-07-2024 12:43 PM
Wanted to clarify one bit. BGP on the LAN side of the MX doesn't require Support involvement. With firmware 18.2 on a supported hardware model this BGP can be enabled. BGP on the WAN side is what requires Support enablement at the moment.
