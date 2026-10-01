---
id: collect-261001-cisco/cisco/questions-32544-basic-catalyst-3560-egress-shaping-c12716a7
title: "questions-32544-basic-catalyst-3560-egress-shaping-c12716a7"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/questions-32544-basic-catalyst-3560-egress-shaping-c12716a7.md
source_anchor: ""
source_lines: [1, 29]
sha256: bbbe351f13ae5e5af9f170cf84c789c93a6679ae07fc5a13aed1c64b8f373f0a
---

# questions-32544-basic-catalyst-3560-egress-shaping-c12716a7

We have a service provider (and we can't change providers) who is giving us a "metro ethernet" style connection between two of our locations. On each end, we plug into an ethernet port on a provider's switch and they ship frames back and forth. We get a certain bandwidth from them and they are dropping packets that burst over the bandwidth.
I'm pretty sure that a good way for us to not burst over their limit and avoid dropped packets is for us to shape our traffic to fit under the limit. I think I'm very close to understanding how to do this, but it's pretty complicated. We have a Cisco Catalyst 3560X on each side of the connection.
If I want to shape traffic down to 50 Mbps across the tunnel, it looks like the right (maybe only?) way to do it is to use shaping on the egress queues of the ports used for the link on each of our 3560s. We do not need to mark or classify any traffic, we just want to shape everything down to 50 Mbps. Here's an example port config right now:
interface GigabitEthernet0/1
 speed auto 10 100
 spanning-tree portfast disable
I know I'll want to do mls qos in global config mode. Then I should see something like this:
[Switch name]# show mls qos int gig0/1 queueing
GigabitEthernet0/1 
Egress Priority Queue : disabled
Shaped queue weights (absolute) :  25 0 0 0
Shared queue weights  :  25 25 25 25
The port bandwidth limit : 100  (Operational Bandwidth:100.0)
The port is mapped to qset : 1
My understanding so far is the following, feel free to correct me:
- All traffic will be CoS 0/unmarked so will go into egress queue 2 by default.
- Egress queue 2 is sharing the bandwidth equally with queue 3 and 4, and queue 1's weight is ignored.
- Egress queue 1 is shaped to 1/25 of the interface bandwidth, so 4 Mbps in this case.
So I get that queues 2 - 4 are each guaranteed 33% of the bandwidth (33 Mbps, right?) and queue 1 is shaped to 4 Mbps. My first question is:
  With this default configuration, if only queue 2 is used, how much bandwidth will it get? 100 Mbps? And if all queues were fully utilized, queue 1 would have 4 Mbps and queues 2 - 4 would each have 32 Mbps (100 - 4 = 96/3 = 32)?
And now the real question:
  To shape all unclassified egress traffic to fit into 50 Mbps, can I
  just enter
  srr-queue bandwidth shape 0 2 0 0 on the interface in question and be done?
It seems like the queue sharing and shaping limits aren't guaranteed, so I might need to shape down to a nominal 45 Mbps on the egress queue if any burst over 50 Mbps is to be avoided. Can I do that by just running srr-queue bandwidth limit 90 combined with the above shaping? Would it be the same to instead use:
srr-queue bandwidth shape 0 1 0 0
srr-queue bandwidth limit 45
Would that shape queue 2 to 45 Mbps (on a 100 Mbps interface)?
Once I understand that, I'm guessing my next stop is sorting out buffer allocations and thresholds so my shaping is dropping as few packets as possible, right? That can be a separate question if necessary, but actually that seems to make a lot more sense so far.
