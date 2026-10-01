---
id: collect-261001-cisco/cisco/questions-7201-cisco-reliable-fast-bgp-failover-when-mixing-bgp-nht-and-bfd-0a1d19da
title: "questions-7201-cisco-reliable-fast-bgp-failover-when-mixing-bgp-nht-and-bfd-0a1d19da"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-cisco/questions-7201-cisco-reliable-fast-bgp-failover-when-mixing-bgp-nht-and-bfd-0a1d19da.md
source_anchor: ""
source_lines: [1, 14]
sha256: 4f84d21121cf93127887ac73d8cd8594a3da3e967a712bf3b7db76abcd40e457
---

# questions-7201-cisco-reliable-fast-bgp-failover-when-mixing-bgp-nht-and-bfd-0a1d19da

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
7
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
When running BFD on interfaces between two directly connected iBGP and IGP neighbours, is it more stable/reliable to use BFD for fast IGP reaction times and BGP Next Hop Tracking for faster BGP reaction times (rather than BFD for both)?
The reason I ask this is that if BFD is used for IGP and BGP link failure notification, an interface that goes down causes the BGP session to be torn down the IGP to run SPF. If the IGP runs SPF and completes in 200ms for example using a tuned IGP deployment, reachability to that iBGP neighbour may be restored and so the iBGP session could have remained.
I'm suggesting the following (because this is how I am interpeting the Cisco docs);
If a link goes down BFD signals this to the IGP and the IGP will recompute a new path to the iBGP neighbour (assuming no ECMP or Un-ECMP or IP Fast Re-Route becasue in those scenarios a link failure will not disrupt the iBGP session I believe?). In this scenario when a new path is computed by the IGP the iBGP peering continues without interuption. When it can't be computed (no other path is available) setting a lower NHT value like bgp nexthop trigger delay 1 to match the fast IGP convergence would trigger the BGP session to be torn down now that there is definately no path to the iBGP neighbour. So does it make sense to not use BFD for iBGP neigbors and use BGP NHT instead?
Well, actually BGP NHT can use BFD to speed up neighbor problems detection, and it's usually the best scaling mechanism, because BFD is often offloaded in hardware.
Running BFD for both IGP and BGP may be good idea, depending on the scaling capabilities of your gear and hardware acceleration/offload. It's the configuration that has to make sense from logical standpoint.
So, your decision boils down to this: if you wish to tune both IGP and EGP protocols to try tu run at the same time (bgp nexthop trigger delay 0 and BGP BFD NHT), which may lead to problems, oscillations and so on, or you'd like to focus on achieving fast IGP convergence and with presence of multiple forwarding paths, let BGP linger a moment longer just to make sure your IGP converges (bgp nexthop trigger delay 1 with or without BGP BFD NHT).
That last approach usually makes more sense, and that's why NHT delay by default is set to 5 seconds. If your IGP converges in 200ms (or generally, under second) tuning BGP NHT to 1 second may make sense.
