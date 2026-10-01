---
id: collect-261001-cisco/cisco/questions-610274-bgp-multi-homed-cisco-routers-configure-for-resilience-but-sepa-67f04691
title: "questions-610274-bgp-multi-homed-cisco-routers-configure-for-resilience-but-sepa-67f04691"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-610274-bgp-multi-homed-cisco-routers-configure-for-resilience-but-sepa-67f04691.md
source_anchor: ""
source_lines: [1, 18]
sha256: dfacddde6fca899ccd4457d40ec04b733f848a09b95339c40d6e78424d7c9475
---

# questions-610274-bgp-multi-homed-cisco-routers-configure-for-resilience-but-sepa-67f04691

I've got 2 Cisco routers, each with a single Internet feed, providing BGP with default route only, which allows us to advertise certain routes out each feed. We're not using the full DFZ because although we usually run all Cisco 3925s, we currently have a Cisco 2911 in temporarily and it hasn't enough RAM to process the DFZ.
The routers share a Layer2 LAN, using a Campus vLAN (Router-On-A-Stick) configuration and we'd really like to take advantage of this shared LAN connectivity and gain some Internet feed resilience, using BGP to not only advertise the primary route via its natural home, but also the route that belongs to the other platform's router, but with that secondary route AS-Path Prepended, to make it less desirable than the route advertised via its natural home router & vice versa.
We also have a separate (currently non-live) site with 2 routers and a (currently) spare pair of /24s which I'm using to experiment with, but I've been unsuccessful getting a primary route and an AS-path prepended secondary route, advertised at the same time, via the same BGP neighbour. My testing is on AS39152 with routes 91.192.234.0/24 and 91.192.235.0/24, advertising to AS29550, using this to show the results of my testing: -
% telnet route-server.as3257.net             
route-server.as3257.net> sho ip bgp regexp 29550 39152
BGP table version is 25589080, local router ID is 213.200.87.253
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal
Origin codes: i - IGP, e - EGP, ? - incomplete
   Network          Next Hop            Metric LocPrf Weight Path
*> 91.192.234.0/24  213.200.64.93           18             0 3257 5089 29550 39152 i
*> 91.192.235.0/24  213.200.64.93           18             0 3257 2914 29550 39152 i
My goal is to get results like this, from a single router's BGP advertisement (you'll almost definitely need to scroll to the right, to see my point): -
   Network          Next Hop            Metric LocPrf Weight Path
*> 91.192.234.0/24  213.200.64.93           18             0 3257 5089 29550 39152 i
*> 91.192.235.0/24  213.200.64.93           18             0 3257 2914 29550 39152 39152 39152 39152 i
I read through http://blog.ipspace.net/2008/02/bgp-essentials-as-path-prepending.html but what I'm looking for is a way to mix prefixes with and without AS-path prepending and I haven't yet found what I'm after. This may be because it's not possible, or more likely (I hope), that I just don't know what search terminology to use.
I've not put any of my config in the question, as it's simply not working and don't know if it's even possible - I can share it if it helps.
Is what I'm trying to achieve actually possible and if so, please could you share an example config?
