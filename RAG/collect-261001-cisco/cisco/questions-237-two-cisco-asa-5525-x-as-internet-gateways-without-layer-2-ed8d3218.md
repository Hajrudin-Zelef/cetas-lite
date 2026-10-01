---
id: collect-261001-cisco/cisco/questions-237-two-cisco-asa-5525-x-as-internet-gateways-without-layer-2-ed8d3218
title: "questions-237-two-cisco-asa-5525-x-as-internet-gateways-without-layer-2-ed8d3218"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "dci", "research"]
source: docs/RAG/collect-261001-cisco/questions-237-two-cisco-asa-5525-x-as-internet-gateways-without-layer-2-ed8d3218.md
source_anchor: ""
source_lines: [1, 19]
sha256: 94609eb16ba1c94a71bc72471e9d11a6de23caefc6aee1bea9cab979b0e53042
---

# questions-237-two-cisco-asa-5525-x-as-internet-gateways-without-layer-2-ed8d3218

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
11
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Add another reason to hate NAT to the list. I'm bringing up two Internet egress points in our corporate network. The edge devices will be ASA 5525-X firewalls. Traditionally you would put these into some sort of cluster, but this requires L2 connectivity. Since these devices will be in separate parts of my network, L2 connectivity is not an easy option.
My current running solution is to bring them both up as independent firewalls and advertise a default route from each. Any ECMP should have the same hash for each flow and push it towards the "correct" egress firewall.
My question is this:
Is there a way to cluster two ASAs without needing a L2 link?
I want a second/third/hundred pair of eyes on my current solution assuming "No" is the answer to #1.
Designate one Internet circuit as the primary and the other as failover
Implement "NAT outside" (public space) routing between the sites with the firewalls
The first option ensures traffic is always going through either one firewall or the other so that NAT doesn't break.
The second option allows you to load balance across both circuits: One equal-cost default route from each circuit, with your local public prefix(es) advertised out both circuits. (This option ignores how connectivity between the sites is accomplished.)
Don't have much experience with ASA's, so I can't really answer question #1.
Be careful with your assumptions about ECMP, however. Different equipment handles ECMP differently. I've seen ECMP implementations from the granularity of "per-destination-prefix" load-balancing (which almost isn't ECMP at all), to "per-packet" load-balancing.
You're going to have to get a bit more sophisticated with your routing handling to make this work. Look for the DCI interconnect information that ioshints has published, that should help you figure out how you might design your network to handle this. Sorry I don't have a URL close to hand on this.
Personally, I wouldn't even consider long distance clustering. I believe Cisco's recommendation is direct cable connect. ECMP may be workable, but it is vital to do per-destination (the Cisco default AFAIK) and not per-packet. Consider the impact on a passive ftp transfer that requires dual outbound connections.
