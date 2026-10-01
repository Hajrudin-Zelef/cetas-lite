---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-ask-the-expert-configuring-and-troubleshooting-border-gate-722f8d45-3
title: "t5-routing-and-sd-wan-ask-the-expert-configuring-and-troubleshooting-border-gate-722f8d45"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-ask-the-expert-configuring-and-troubleshooting-border-gate-722f8d45.md
source_anchor: ""
source_lines: [213, 261]
sha256: 0a3fc838e17b0a53919cce3f33dc190c3c65abfa17e8a24ce6743722a0d34667
---

# t5-routing-and-sd-wan-ask-the-expert-configuring-and-troubleshooting-border-gate-722f8d45

With reference to you query it is not recommended to use BGP in this setup because as a best practice BGP is a viable solution when used in dual home scenario so here you can configure IGP with your service provider. Or if you wanted to run BGP you have to ask for 2 eBGP peering with provider.
However if you are keen to run BGP with in the specified conditions you can try a workaround of running eBGP peering on HSRP/VRRP virtual IP but it will cause the delay and only the session initiated by provider router will establish the BGP. You can minimize the delay upto some extent by changing the HSRP and BGP timers.
But apart from delays there will be one problem that your eBGP session from the standby router will be in active state and keep on probing and I think would not be acceptable. This is not a recommended solution and just a workaround.
Hope it answers your query.
Thanks & Regards
Sandeep
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-30-2013 09:31 AM
Hi Sandeep,
Simulated in lab, results are same as you mentioned, thanks for your inputs.
/SANJEEV
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-26-2013 08:45 PM
Hi,
We have a VSS domain, with 2 BGP upstream connections (to the same AS), one on each domain-switch...
In BGP we set maximum-paths 2 I'd like to know if there is a way to load-balance over both links outgoing traffic.
I do see both bgp routes in the routing table but VSS is prefering the link on the active switch (as expected I guess), is there a way to overide this behaviour and send traffic over the vsl-link to the other link?  (don't feel for manipulating bgp attributes for half of the routes).
Tnx
Josh.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-29-2013 06:54 PM
Hi Josh,
First of all in order to utilize both L3 links, you need to make sure that devices are dual-homed to both VSS chassis with Multi-Chassis EtherChannel (MEC), otherwise traffic will only be sent out from the local chassis which is an expected behavior of VSS.
I have seen the similar issue earlier where customer has the single connectivity between LAN and VSS core and soon as he connected to both VSS switches it started load balancing.
If in your case you already have the dual-homed (between VSS core and LAN).Please share the below captures
- show ip bgp (from vss) and specify any route
- show ip route (from vss) for the same route in above capture
- traceroute from your VSS switch and LAN to any IP address in outer segment ( from VSS switch, machine and switch below VSS domain in LAN)
Please feel free to contact in case you have any further query.
Thanks & Regards
Sandeep
