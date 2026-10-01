---
id: collect-261001-huawei/huawei/questions-22949-e12187cd
title: "questions-22949-e12187cd"
domain: huawei
role: reference
task: reference
actors: []
dates: ["2017-08-06"]
keywords: []
source: docs/RAG/collect-261001-huawei/questions-22949-e12187cd.md
source_anchor: ""
source_lines: [1, 16]
sha256: 18e653697ba82b9e094a86e866ec9b492b9786348087e9171889b15b893c96f9
---

# questions-22949-e12187cd

Since there aren't many resources out there regarding virtual rendezvous protocol, my basic understanding of it seems very similar to BGP. Which basically had me questioning why the protocol even exists. Can anyone please explain VRP, and how it differs from BGP?
- 
        Maybe you could add some more context. From the little reading I've done, VRP just describes how to hosts can find each other. Which is nothing like BGP, which enables two routers to exchange routes to known networks and a plethora of additional information about each possible destination. Maybe you could describe how you think they are similar, or even interchangeable?Eddie– Eddie2015-10-01 03:29:21 +00:00Commented Oct 1, 2015 at 3:29
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you can provide your own answer and accept it.Ron Maupin– Ron Maupin ♦2017-08-06 23:57:45 +00:00Commented Aug 6, 2017 at 23:57
1 Answer 1
actually there is no relation between them , BGP is a routing protocol which in general will be used to exchange routing information (routing table for example) between multi hopped routers , but VRP is sort of communication establish mechanism (between two peers) SIP and so on . so simply BGP work mainly in layer 3 and VRP work in layer 7
- 
        Correct apart from VRP does not need to be at layer 7, vrp protocols for multicast for example are still at L3 however this should be taken as the correct answer bgp is a routing protocol and vrp is a form of communication for 'neighbour' discovery and not sharing routes. For example in a multicast network vrp will advertise a ip and priority to others listening on the same multicast group and the 'winner' will be used as a mroute as the next hoppolihrono_crepes– polihrono_crepes2015-10-04 09:22:09 +00:00Commented Oct 4, 2015 at 9:22
- 
            
            
- 
        
- 
        so this is used mainly for peering with directly connected devices?user2883071– user28830712015-10-13 19:52:51 +00:00Commented Oct 13, 2015 at 19:52
