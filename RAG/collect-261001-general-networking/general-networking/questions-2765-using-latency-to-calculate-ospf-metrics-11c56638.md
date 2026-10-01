---
id: collect-261001-general-networking/general-networking/questions-2765-using-latency-to-calculate-ospf-metrics-11c56638
title: "questions-2765-using-latency-to-calculate-ospf-metrics-11c56638"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "cost", "ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-2765-using-latency-to-calculate-ospf-metrics-11c56638.md
source_anchor: ""
source_lines: [1, 17]
sha256: 124ddc7f4e2c3533d8c0f43543c8a14d3bb046e0ec72aaff673d863aa99aeae3
---

# questions-2765-using-latency-to-calculate-ospf-metrics-11c56638

Is it possible to configure OSPF to use latency in its link cost/metric calculations. The idea being to enable performance based routing whereby the fastest and least congested links in the network will be chosen dynamically, rather than relying on OSPF's default behaviour which results in 100Mbps links having the same cost as 1Gbps links.
- 
        1if you are using cisco gear, look up auto cost reference bancdwidth to make automatic distinctions between fast ethernet and gig e costs.Mike Pennington– Mike Pennington2013-08-11 13:59:44 +00:00Commented Aug 11, 2013 at 13:59
3 Answers 3
Is it possible to configure OSPF to use latency in its link cost/metric calculations. The idea being to enable performance based routing whereby the fastest and least congested links in the network will be chosen dynamically, rather than relying on OSPF's default behaviour which results in 100Mbps links having the same cost as 1Gbps links.
Short answer: No, not with OSPF alone
Long answer:
The only way for OSPF to dynamically calculate paths based on latency / congestion is to use MPLS Traffic Engineering with offline optimizations of MPLS TE costs based on your criteria; MPLS TE uses OSPF LSAs to carry information about the label switched paths. However, MPLS Traffic Engineering is a heavy hammer and many network operations can't deal with the additional workflow introduced into provisioning or troubleshooting MPLS TE.
Another answer suggests that you should not adjust link costs based on bandwidth, and to use the role of a node for costs. I can't speak for his network, but this guidance is unnecessary in many cases since the lowest cost path in a well-designed topology automatically follows through the core of the network. I wouldn't try to adjust an inefficient topology with link costs... just make traffic flow through the core naturally and ensure that OSPF sees a 1GE as a better path than a FastEthernet link. This will naturally happen if you lay out the topology well, and use auto-cost reference-bandwidth under the OSPF process. Be sure you use this on all OSPF routers so they understand the link costs in the same way.
You shouldn't rely on default metrics/cost, not in IGP or STP. My one question review if LAN was designed is to ask which link is blocking, if engineer can't answer without looking, it's strong indication the network was not designed.
You should design how you want traffic to flow under normal operation and under fault scenarios and then design metric standard with satisfies those requirements. I prefer to do this with pen and paper add coffee if so inclined.
I greatly prefer role based metric standard, where P-P, P-PE, PE-PE etc each have same metric value, regardless of bandwidth or latency. That is, I know where I want my traffic to go, to me it does not matter what the latency is or what the capacity is, if I don't have enough capacity where I want the traffic to flow, I'll upgrade the links.
Role based certainly isn't only way to do it, if you want to base your metric standard on link speed or measured latency, go ahead.
Regardless how you design your metric standard it will be fully static and unaware of actual link conditions such as packet loss, latency, jitter, capacity etc. If you want network to more dynamically respond (demand which I tend to think often rises from poor planning) to changing network conditions, you need to look into RSVP or even SDN.
Another option if to use PfR or OER that is routers use active probes or netflow data to select a path dynamically. Or of course the very simple option is to hard code OSPF cost values.
The other consideration is are you trying to code a static metric or do you want dynamic behavior (latency changes with traffic load).
Finally as was mentioned above, are you sure you want the complexity, if you look at EIGRP it has a load metric but no one ever uses because it adds operational complexity and you can get traffic oscillation.
