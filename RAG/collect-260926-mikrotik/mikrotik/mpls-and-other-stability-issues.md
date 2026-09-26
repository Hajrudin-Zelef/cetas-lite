---
id: collect-260926-mikrotik/mikrotik/mpls-and-other-stability-issues
title: "mpls-and-other-stability-issues"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "ethernet", "memory", "training"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mpls-and-other-stability-issues.md
source_anchor: ""
source_lines: [1, 21]
sha256: f3854a606243aa6febfe391111a56004b7c08313521077258bc1d9c35e5ddfe4
---

# mpls-and-other-stability-issues

Hi Everyone:

I run a small AS as a (Mikrotik-) Training vehicle. The core is a CHR and a x86 as main BGP routers in two data centers, with 2x Transit and a few peering sessions. All external interfaces are via physical Mikrotik routers, connected via MPLS PW to the CHR/x86 core routers. There are also off-site routers connected via GRE/IPsec and L2TP/IPsec, and through MPLS PWs i can move any L2 interface to any location in the network, which has a total of about 10 routers. All in all, a mix of professional-grade links (1G Ethernet between data centers, Transit, Peering) and consumer-grade DSL links carrying L2TP and GRE. The network allows for some experimentation, while of course BGP advertisement and peers should be stable.

The network was reliable over a few years; throughout the last 12 months I was replacing ROS V6 with ROS V7. I am thankful to Mikrotik for supporting older routers with current firmware – it is not only generous, it also allows to run a whole network on a single ROS version. Having said that, I continue to struggle with some stability issues:

**1) BGP itself seems to be stable at 7.16** with session uptimes of 100 days or more. Two external sessions with full-table (approx 980k routes), plus peering, both IPv4 and IPv6. The setup is simple (full mesh / no RR).

**2) MPLS headache**: I was unable to get MPLS/LDP to work reliably in V7 - a few weeks ago I replaced all MPLS PW with EoIP, and since then the stability issues are gone. One issue might be that I have no intent to use MPLS for IP traffic (only for PW): For this reason, I use LDP Advertise Filters (also Accept Filters) so that only Router Loopbacks (all in a /24) are propagated. While this works seemingly at least initially, later on suddenly a router becomes unreachable. After debugging OSPF LSPs and routing tables on the whole path forth and back, it becomes apparent that the routes are correct. Turn off MPLS, and the router becomes reachable. Seems that MPLS directs some traffic into LSPs, but the LSPs lead to nirvana. Contributing to the issue, in my opinion, is the rather dynamic nature of my network: L2TPs come and go, causing relatively frequent (= every few days) topology changes. It seems to me that there are race conditions in the code – between processes, or also between CPUs – with the result that some things usually work, but then statistically fail the 30th or 50th time. I also had OSPF in once case not adapting to a topology change – due to lack of time I had to reboot the CHR, and voila it worked. I am also aware of tips here in the forum relating to MPLS: To exclusively assign a label range per router, and to use the same Hello intervals for LDP and OSPF. However, if this helps, then it confirms issues - and such solutions only lower the probability, i.e. one then has to “less often” restart half the network to make it operate again.

**3) Testing:** My real question is how I can manage to run into stability issues with only a few hours time for Mikrotik every month? Shouldn’t Mikrotik catch such issues? Given the (sometimes) very few hours between when a ROS version is compiled and then it is published, I can only suspect that not that much quality assurance is happening at Mikrotik. My suggestion would be a) that Mikrotik publicly documents the tests conducted, including the configurations. Such “reference networks” would make us users understand whether we are in uncharted territorry, or whether we are using a feature in a known way. It would also complement the documentation. And please use scripting to artificially create unstable networks in the lab; then let the lab run for weeks to understand whether software is stable, and whether memory leaks exist (you can get back to me for more details).

1. Please tell what you know **does not work** . There are issues in Winbox “everyone know about”, for example. In aviation, an aircraft does not need to have all systems operational – regulations allow to fly with some features inoperable, as long as minimum conditions are met (see Minimum Equipment List etc.). In Mikrotik world, the pilots would at least like to know what is broken, then we could avoid those landings in fog at night when MPLS is still not stable. Point 4 here, you may have noticed, is the opposite of point 3 above (where the claim of ¨it works" should be backed by relasing corresponding config).

While Mikrotiks are complex products, and while **in my own experience 80% of the errors in the end were my own faults**, I am still uneasy about some features of the platform. MPLS no way at the moment, BGP.. hmm… I don´t dare to touch it sometimes. IPsec … hmm… OSPF… hmmm…

Any input from other users about the use of MPLS PWs and/or OSPF stability?

regards,

azg
