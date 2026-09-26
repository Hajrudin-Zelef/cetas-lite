---
id: collect-260926-mikrotik/mikrotik/equal-bandwidth-distribution-pcq-vs-untouched-2
title: "equal-bandwidth-distribution-pcq-vs-untouched"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "parameters", "research", "throughput"]
source: docs/RAG/lot-mikrotik/forum/qos/equal-bandwidth-distribution-pcq-vs-untouched.md
source_anchor: ""
source_lines: [195, 256]
sha256: 9b3c7779481477af8f92fe845f66afe70cf62a0f9517e44c52770ddd2475660b
---

# equal-bandwidth-distribution-pcq-vs-untouched

** at this time there is no advantage in mangling as all to-from-lan is marked with “pcq-connection”, still a to-do for future?*

Yes!

** “0.0.0.0/0” means “anything”, internal or external*

Ok.

*interface queue or global queue have both their strengths and weaknesses[…]*

Thanks for the explanation! I changed the parent queues to interface:




```
/queue type
add kind=pcq name=pcq-down pcq-classifier=dst-address pcq-dst-address6-mask=64 pcq-src-address6-mask=64 pcq-total-limit=12650KiB
add kind=pcq name=pcq-up pcq-classifier=src-address pcq-dst-address6-mask=64 pcq-src-address6-mask=64 pcq-total-limit=12650KiB
/queue tree
add max-limit=85M name=down packet-mark=pcq-down parent=ether2 queue=pcq-down
add max-limit=8500k name=up packet-mark=pcq-up parent=ether1 queue=pcq-up
```

** “pcq-total-limit=12650KiB” that’s a lot of buffer[…]*

Well, I didn’t quite understand these parameters yet. I saw some examples on this MUM document (pages 26-27): https://mum.mikrotik.com/presentations/US08/janism.pdf

Based on those examples, I thought the values I set meant 253 users (/24) could use the queues simultaneously (although very unlikely to happen).

Currently, values on both queue types are set to:

Limit: 50

Total limit: 12650

Could you please give me some directions about these parameters and suggest something? I think it’s the only missing piece to my setup now!

             
            
           
          
            
            
              If it was up to me, the pcq-total-limit shouldn’t be much larger than 1/10 s of max transmission on upload side: suppose you have a gamer, that would have latency of 100ms… he wouldn’t be happy.

On download the queue is only there to account for and spread the bandwidth. It’s an artificial bottleneck to control throughput. So there is no need to have a large buffer. => default values will do.

Large queues are good to fill the pipe, but bad for QoS.

             
            
           
          
            
            
              
The bandwidth I provided on the post didn’t correspond to the actual scenario (I thought using a round number would make it more simple to understand), so not that many users will be queued at the same time.

I did some research (even on this forum) and I thought there’s not much straightforward information on this topic (pcq-limit and pcq-total-limit). I ran some tests on my own and set those parameters to values that I judged to be balanced. Not worthy sharing/explaining the values that I’ve chosen because it’s NOT based on deep reading/lots of experience, so it could be misleading. But one thing that I noticed and should note is that “pcq-total-limit” doesn’t affect users individually, since they are being limited by “pcq-limit” (single sub-stream size), so despite how big “pcq-total-limit” is, single sub-streams are never gonna be bigger than “pcq-limit”.

Thank you very much for your help.
