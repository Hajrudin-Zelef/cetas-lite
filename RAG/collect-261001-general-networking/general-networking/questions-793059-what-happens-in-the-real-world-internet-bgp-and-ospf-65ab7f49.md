---
id: collect-261001-general-networking/general-networking/questions-793059-what-happens-in-the-real-world-internet-bgp-and-ospf-65ab7f49
title: "questions-793059-what-happens-in-the-real-world-internet-bgp-and-ospf-65ab7f49"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-793059-what-happens-in-the-real-world-internet-bgp-and-ospf-65ab7f49.md
source_anchor: ""
source_lines: [1, 11]
sha256: 4348c8e56b6c6de5766ca74f290869a7d7d922a7391d6aa0cbc9e024e6492058
---

# questions-793059-what-happens-in-the-real-world-internet-bgp-and-ospf-65ab7f49

I have never worked on the customer premises, but my assumption is that one service provider runs IGP services internally - For example - OSPF or IS-IS. Now, when one service provider has to reach another service provider to communicate server on the third service provider, it will pass/transit through another ISP.
In my example - ISP-2. My question is -
Apart from multiple OSPF area's, which is to limit 50 routers per area, is it really the way internet works. You can ignore the fact that it can be IS-IS or other igp protocols within AS. In simple words, is this a kind of real topology of internet.
Next, my context based question is if this is the typical scenario, then how I can ping 11.11.11.11 from R2.
I am saying this because -
- AS-100 has all their internal route information -- OSPF 1 -- it doesn't know how to reach 11.11.11.11
- AS -200 has all their internal route information - OSPF 2
- AS -300 has all their internal route information - OSPF 3
I have connected AS 100 with AS 200 using EBGP. Similarly, AS 200 with AS-300 . This I have done using neighbor command.
Questions -
Do I need to advertise each and every network information of ISP-2 to ISP1 and ISP-3 so that they all can communicate? If so, is it done in the same way in the real network. This will be such a tiring job to advertise each and every network information of ISPs. How can I do this automatically? I am using CISCO - GNS3 any cisco commands will help. I have explored re-distribute but I haven't understood much. Also, in the real network - if you advertise each and every routing information to other ISPs then the routing table will become so large. Do we normally do this? But, I believe if we don't do so then how can I communicate to internal network of other ISPs. So, I believe we have to do so. This means routing table will be very large. Am I correct here?
