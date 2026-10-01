---
id: collect-261001-meraki/meraki/r-meraki-comments-169y4cw-autovpn-local-internet-breakout-4edf7fae
title: "r-meraki-comments-169y4cw-autovpn-local-internet-breakout-4edf7fae"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-169y4cw-autovpn-local-internet-breakout-4edf7fae.md
source_anchor: ""
source_lines: [1, 28]
sha256: 6e95785f9006d845df6f7495f99a5d0f199d3d67d7eca2289956fb683c404e2a
---

# r-meraki-comments-169y4cw-autovpn-local-internet-breakout-4edf7fae

AutoVPN local internet breakout 
        
    I'd like to send all traffic from a spoke MX to hub MX except internet traffic.
I can do this already easily - under site-to-site VPN - tick IPv4 default route. Then set local internet breakout rules to include internet-bound subnets. But this is messy and ugly.
I'm trying to do this, because I cant find a way to send traffic to the hub MX if it's not advertised on AutoVPN.
Some traffic is destined for another private MPLS network, which outside of the Meraki AutoVPN and is connected to the hub MX's WAN2 uplink.
Does anyone have any good ideas on how to achieve this?
Section des commentaires
I think you're having an X/Y problem moment.
This is the default behavior on any router right out of the box. Non-local (what you called "internet") traffic gets sent out directly to the internet, and local traffic gets routed internally ("all traffic except internet"). When you set up multiple networks with Meraki and enable VPNing in either manner, this will still be the default behavior.
However, you also said:
I think your REAL problem is that you're not properly advertising internal-only routes.
Any reason you don't just form a direct VPN link to that external network rather than hair pinning it in Meraki? You CAN hairpin but it should be avoided when possible.
I've been using local internet breakout already by setting the default route to the hub MX, then "opting out" all internet traffic - however this is crazy hack, for obvious reasons.
I looking for something where I could route traffic to the hub which matched a set of subnets, but wasn't advertised already on AutoVPN.
Yup, I think I've been looking for a quick hack when the answer is I need to go the corporate/politics route - to get BGP announcements and a tunnel directly into the other private net - which is what I was hoping to avoid.
Yeah, I can imagine that, but this will work better, more reliable, and people after you will be able to know what is happening.
You could turn off the default route setting and instead advertise all local routes (192.168.0.0/16, 10.0.0.0/8 and 172.16.0.0/12) from the hub by creating static routes to your next hop from there, and enabling their advertisement through auto vpn. More specific routes will still take precedence, so your other autovpn routes will still go where they are supposed to. You will get a warning when you save specifying that it overlaps and more specific routes will be used, and then should get desired behaviour when you accept that warning and it applies.
If you have multiple hubs you will want to contact meraki support to disable route summarisation, as in that situation it breaks hub preference settings (as the hub with this setting ONLY advertises the summarised routes which are now the full local range, not the more specific ones)
This is what I have done during a transition between MPLS to SD-WAN so we can send any non-migrated sites back to our MPLS via the head office. It works for our purpose but it does break template auto ip address selection, so we template by cloning site and manually specifying ip addresses.
Does this mean multiple hubs can advertise the same 10.0.0.0/8 to the MPLS network. Do you configure this using static routes or addressing and VLANs? And does it require a one-arm concentrator, rather than a NAT-based concentrator.
Might be able to use this as a stop gap, if so.
You configure it as a static route on the addressing and VLANs page and then enable the route for VPN. NAT-mode/Concentrator mode work fine either way. You don't need local internet breakout for what you are trying to do, you just need to make sure you are advertising the correct static routes and subnets.
Sorry for long delay on reply, I’m not on reddit all that often.
Yes, you advertise via static routes, you can already do this from multiple hubs, and priority will be based on your hubs priority per spoke.
That comes to the specific reason for disabling route summarisation is if you advertise the entire 10.0.0.0/8, it will summarise all other more specific routes (for example, a /24 local to the hub) into that single /8, and then that becomes the lowest priority hub regardless of where it sits, because other hubs will have the more specific route advertised for everything (including the route to the /24 on that hub), then you get routing paths jumping via random hubs fans not your expected priority order of hubs.
Using a routed hub lets you separate out a dedicated /30 interconnect vlan to an MPLS router to send traffic.
On the MPLS side, your MPLS router would need to be advertising the return route for your SD-WAN subnets back to the SD-WAN network via the meraki interconnect IP, mine does this via static route on the MPLS router, which is then redistributed via BGP to the MPLS. If you have multiple, you would need something to specify priority within the MPLS of each point of interconnect, and something to stop redistributing the route from that specific point of interconnect if the router couldn’t reach its local meraki. I opted to just have one point of interconnect, and if that drops I would add and redistribute the static route from the secondary site, as downtime of our specific interconnect in this case is more an inconvenience than a giant business disruption, and I could get it back up easily and pretty quickly.
