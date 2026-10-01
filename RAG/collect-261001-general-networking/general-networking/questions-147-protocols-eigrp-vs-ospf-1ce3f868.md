---
id: collect-261001-general-networking/general-networking/questions-147-protocols-eigrp-vs-ospf-1ce3f868
title: "questions-147-protocols-eigrp-vs-ospf-1ce3f868"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-147-protocols-eigrp-vs-ospf-1ce3f868.md
source_anchor: ""
source_lines: [1, 25]
sha256: 5275656e927760984de0fbe5146243c22a948a0e08f835f75407e6bf2fbc1f07
---

# questions-147-protocols-eigrp-vs-ospf-1ce3f868

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
23
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
EIGRP and OSPF are both IGP protocols, the former is a mostly Cisco protocol and the latter is a well established open standard. What are the benefits of one over the other?
Put another way, when deploying a network, why choose one over the other? If you have mixed devices the choice would obviously be OSPF, but what if you are running a Cisco only shop? Are there any features where EIGRP excels compared to OSPF that would make it feasible to only deploy EIGRP?
If we look at EIGRP with default settings and OSPF with default settings and there are multiple loop free paths to a destination then EIGRP will converge much faster because it keeps what are called feasible successors in it's topology database. These are basically loop free alternatives to the best path. EIGRP also has summarization at any point in the network. It also has stub feature which is useful when you don't want to use a router for transit. Commonly deployed in DMVPNS. EIGRP is also less confusing than OSPF because it does not have different network types and EIGRP is easier to deploy in hub and spoke scenarios.
EIGRP uses a flat network without areas, this can both be an advantage and disadvantage.
OSPF is obviously an open standard so it's the logical choice if you have multiple vendors. It can perform well but it requires that you tweak SPF timers because by default in IOS there is a 5 second wait before running the SPF algorithm. OSPF uses areas which means you can segment the network more logically. OSPF can only summarize between areas. OSPF is link state so it has a better view of the entire network than EIGRP before it runs the SPF algorithm. Network administrators will usually be more comfortable with OSPF because it's more commonly deployed.
Both protocols have advantages and disadvantages. But the common answer that EIGRP should be discarded because being proprietary is not entirely true any longer.
You can read about the finer workings of these protocols for yourself, they are thoroughly documented on the Internet and it's a doddle to find information on them.
From a practicle perspective I would say that in the case of EIGRP vs OSPF, OSPF always wins for the following reasons:
Convergence Speed:
Everyone always mentioned that EIGRP is faster than OSPF using default settings. If you deploy either protocol without reading about them and use their default settings, then you clearly don't know what you are doing in my opinion. Why would you deploy default settings without knowing what they are, and when you do realise what they are you realise that OSPF supports BFD and becomes lightening quick (as does ISIS).
Traffic Engineering:
Because OSPF like ISIS is based on TLV values, it has been extended quite a lot. It has support for extensions like MPLS-TE and GMPLS.
Continual Expansion
As I mentioned above, OSPF and ISIS have been extended quite a lot and extension drafts are being written fairly regularly and will continue to be. EIGRP doesn't have many of the advanced options these two do.
Scalability
OSPF scales better than EIGRP with its use of areas however, I don't think this really matters either (like the convergence time aregument, due to BFD). Not many people are running 10k routes in one area in OSPF. Typically I would use an IGP for fast routing within a given part of a network, but ultimately iBGP carries all the internal routes. Every single router doesn't need every internal route in its RIB sourced via OSPF if you have hundreds or thousands or routers, some of them are so far away (topologically speaking) it's worthless knowing about them.
Interoperability
Lastly there is the obviously reason that EIGRP is/was a Cisco proprietary technology. Although this has recently been submitted into a draft for other software vendors to start incorporating, it's too late (I believe). No currently running network is going to waste huge sums of money switching from some other IGP to EIGRP, and I don't know why a new network would consider it (if you are going to be mixing Cisco equipment with non-Cisco). Simply because non-Cisco equipment that supports OSPF has been doing so for years. The code is tried and test, many bugs fixed, oodles of information on line etc. It will take years for EIGRP to catch up (if it isn't too late already!).
I would suggest the right answer is it depends on the topology of the network. OSPF requires area boundaries to do summarization, if you are doing a green field network or your topology is condusive to drawing areas out then by all means use it. If your network requires spokes to connect to multiple hubs, then OSPF is harder to do, when I had this requirement on my frame relay network I migrated remote sites to BGP. I wanted to use EIGRP and stub routers but cisco mentioned that more resources were being spent on BGP-OSPF interoperability vs EIGRP-OSPF interoperability so I went BGP on that basis. Another way to put it is that EIGRP with its stub routers and the ability to summarize wherever you want will scale better in a 'messy' topology.
