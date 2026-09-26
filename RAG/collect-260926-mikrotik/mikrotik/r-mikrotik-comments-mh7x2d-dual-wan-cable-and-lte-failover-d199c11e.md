---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-mh7x2d-dual-wan-cable-and-lte-failover-d199c11e
title: "r-mikrotik-comments-mh7x2d-dual-wan-cable-and-lte-failover-d199c11e"
domain: mikrotik
role: reference
task: reference
actors: ["SpaceX"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-mh7x2d-dual-wan-cable-and-lte-failover-d199c11e.md
source_anchor: ""
source_lines: [1, 39]
sha256: d7912368b5c7d0d1a70b3aaae9a17601206802854ec717b95f88f9c2dd61043d
---

# r-mikrotik-comments-mh7x2d-dual-wan-cable-and-lte-failover-d199c11e

Dual WAN (Cable and LTE) failover 
        
        
        
    
    
    Basically, default config covers most cases: when current gateway dies - mikrotik marks currently used route as down and therefore next available route (based on distance) is used, which is route from LTE interface.
However, I've reached edge case - ISP's gateway is working fine, but still 'no internet' (ISP's upstream or is dead or upstream's upstream, who knows). In this case, basic failover won't handle switching to LTE interface because default route is still up because ISP's gateway is reachable.
So how do I proceed from here? I've heard about netwatch tool but there's no clear way (if any) to specify which interface to 'watch' from, therefore if I'll start netwatching my desired host (say google dns), what I'd assume would happen:
- 
      netwatch pings 8.8.8.8, its suddenly down
- 
      netwatch runs 'down' script which basically switches default gateway to lte1 interface
- 
      netwatch pings 8.8.8.8 again, now using the only available working route (default gateway = lte1 interface) - link is up
- 
      netwatch runs 'up' script which basically switches default gateway to ether1 interface (which is still not working)
So it ends up switching interfaces back and forth.
Basically, how to achieve failover with duo WAN's assuming cable ISP is crap but it's gateway is always reachable? I'd appreciate any advice on that matter.
Section des commentaires
Recursive routing... you need to force checking 8.8.8.8 over a specific route.
Uhmm, can't you point which interface ping should use, in RouterOS?
That's just for ping, not for netwatch.
Should works in scripts, right?
not with netwatch. no.
It uses the 'best' metric route to communicate out. Recursive routing fixes this.
Great resource here, I would like to do this for LTE and Starlink connections.
Currently I have an RBM11G+LTE modem mounted within an exterior antenna enclosure, powered over POE from a CRS112 which serves as my main closet switch. I'm guessing CRS112 isn't the right device to manage failover, but would love expertise from this community.
Is there another recommended mikrotik device that I would need to place in between the two that I have now to manage the failover from Starlink?
Commentaire supprimé par un membre de l’équipe de modération
Thanks for advice, I'll test it and get back with results.
Commentaire supprimé par un membre de l’équipe de modération
Also, I found this useful presentation that explains in detail WAN failover with recursive routing, with different scenarios:
https://mum.mikrotik.com/presentations/TH18/presentation_5725_1534743837.pdf
Thank you!
That won't work with the script. Because once the first route goes down, the 2nd weight will be active.
He needs to check a different external IP per gateway he needs check is up.
Commentaire supprimé par un membre de l’équipe de modération
You should not be using 8.8.8.8 for this, especially not using icmp.
