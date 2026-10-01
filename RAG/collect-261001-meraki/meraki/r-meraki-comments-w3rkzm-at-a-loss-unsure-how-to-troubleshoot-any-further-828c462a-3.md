---
id: collect-261001-meraki/meraki/r-meraki-comments-w3rkzm-at-a-loss-unsure-how-to-troubleshoot-any-further-828c462a-3
title: "r-meraki-comments-w3rkzm-at-a-loss-unsure-how-to-troubleshoot-any-further-828c462a"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-w3rkzm-at-a-loss-unsure-how-to-troubleshoot-any-further-828c462a.md
source_anchor: ""
source_lines: [57, 85]
sha256: 3410169a8940559551b4fb7116ceacc2cfca191d27c09583f2d74a5a5d77cdd0
---

# r-meraki-comments-w3rkzm-at-a-loss-unsure-how-to-troubleshoot-any-further-828c462a

Secondly, another area to check would be your SD-WAN & Traffic shaping. Just for kicks, in the uplink configuration, if that number is set much higher than your ISP WAN speed, then the extra packets can get lost. Meraki Support wont discover those two. Unless you get someone on the line that is OG Meraki and ISP. That was a painful lesson to learn, fortunately in our case, our Fiber ISP has some great engineers to work with. Tweaking that is really helpful.
Id check that and see if that helps. Of course, you would need to restart all connected devices to get the new DNS info piped back out to them.
I have questions regarding the VLAN and LAN connections, which I will read up on the rest of the threads first... sounds like a peculiar issue.
Had something similar happen on a T1 years ago and it was the “clocking” on the T1 circuit which had to be set to External/Internal on the T1 Routers, Cisco by the way. Was this all working before that update? I recently had to roll back a switch update because it destroyed the Trunking on the network.
Unfortunately this location has "always" been a little hit or miss, but I will say that 15.x was by far the most stable build.
I just have such a hard time pointing the finger at the ISP here -- rock solid connection, and I can swap it out with an 881 and have zero issues -- but then the Meraki gets dumped in place and issues appear.
I think u/Able-Stretch9223 was leading me towards MTU issues (if it's not always DNS, it's MTU) -- but I would suspect the Meraki magic happens and figures that out automatically/would report in the logs if there was an MTU issue. Additionally, the other office with the exact same service has (basically) zero issues -- so if it were an MTU issue, I would expect to experience this.
What kind of link is it? Ethernet? MPLS? I did have s bad MTU issue on an MPLS link years ago there was an MTU mismatch upstream.
Have you opened a ticket with the ISP? Could it be a routing issue with the ISP and how it’s handling their Meraki cloud traffic?
No. I just have such a difficult time pointing the finger at the ISP (even with it being Rogers). It doesn't stand to reason to me that the Cisco 881 would perform perfectly fine, but a Meraki doesn't.
The Meraki also shows perfect uptime/connectivity to the cloud (remember, I don't lose pings to the modem or the Meraki WAN interface -- just internal tunnels).
AFAIK, the Meraki registration servers come into play to establish/maintain a tunnel, but once established the connection is direct -- so even if the ISP was taking a sub-optimal route to Meraki, that process should be fairly straight forward.
Ever figure anything out with this? I'm having the same issue with two satellite offices in an eerily similar configuration to yours. Specifically a Rogers business static IP with a Hitron CODA-4680 bridged to a Meraki Z3. Has been a major issue for these offices since...June-ish?
I'm actually running a Hitron CODA-4582 -- but eerily similar indeed.
So -- there's multiple ways I can answer this. The most problematic site (the one I referenced in this scenario) started largely having these issues after a power failure/APC battery depletion/failure. I went back on my e-mail chain with APC about it, and it looks like it was back in March of 2022.
Before then, once in a blue moon there would be issues -- but largely stable. After, it was a complete crapshoot. Some days/weeks -- it'd be perfectly fine, others not.
Oddly enough, it (and the other office) have been stable since I posted this issue.
Just based on experience, I would say the following items are what seem to have "fixed" it:
Upgrade the hub to 17.8 beta (had to anyway due to work around a WAN1/WAN2 locking up until one of the WAN connections is physically pulled -- other issue I've ranted about on here)
Upgrade the spokes to 17.8 beta
Factory reset spokes -- pin hole push, manually reconfigure static IP's, etc.
Added a laptop cooling pad for both the modem + MX67W -- this one, to be exact: https://www.amazon.ca/gp/product/B08B38LQZ6/ref=ppx_yo_dt_b_asin_title_o00_s00?ie=UTF8&psc=1
The MX67W has a USB port, so I used that for the power. The one site annoyed me, so I left the RGB on, just because.
Is your modem static IP and WAN interface IP staying up though? When I first had Rogers install the lines -- the Hitrons were complete trash. I ended up having to have them downgrade me to a specific version that seemed most stable.
I haven't been able to dig into it as deep as you have. What I know is this:
2 locations with Rogers Hitron modems had similar symptoms that started a few weeks before the giant outage, possibly coincidental
at first thought it was an internal equipment issue at one of the locations, so swapped out an Aruba switch as well as the MX at that location - but symptoms still continue
VPN will be fine for a random amount of time, but then will see flapping and dropped packets in our monitoring, a power cycle of the Rogers equipment sometimes fixes but have also had to pinhole reset the Rogers modem to fix for a longer stretch
We do have other locations with Rogers connections that aren't seeing this sort of behaviour, with identical Meraki equipment. I can't be sure of what Rogers modem they're using, however.
