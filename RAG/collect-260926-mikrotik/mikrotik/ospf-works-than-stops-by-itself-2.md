---
id: collect-260926-mikrotik/mikrotik/ospf-works-than-stops-by-itself-2
title: "ospf-works-than-stops-by-itself"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/ospf-works-than-stops-by-itself.md
source_anchor: ""
source_lines: [246, 285]
sha256: 30ce2b377f749834944abf9fdcebebee52eae37790451485e5f61bd1a88f9388
---

# ospf-works-than-stops-by-itself

           
          
            
            
              K. 'll try that today, asap.

             
            
           
          
            
            
              

It works !



Add /32 on bridge, all ok.

             
            
           
          
            
            
              
Not sure if my recent issue was directly OSPF or a bug with 4.17 using virtual AP,

Each AP’s are using OSPF servers to clients, on a routed network, works OK but when i setup a Virtual AP to any of the sector AP’s, on a ping test beyond the AP after 10pings it would drop for 4 and then resume - OK for 10 and drop for 4,etc,

There no dropouts when pinging from client to VAP,



The Ap’s using 4.17 (433AH) and connected to a 493AH (also was 4.17) using a /28

I simply downgraded the 493AH from 4.17 to 3.30 and problem solved, so for now any MT device not using NV2 is downgraded to 3.30 and works much stable for now.

I am going against the logic of *"…using the same* routeros versions and loopbacks are less likely to trigger or manifest the bug(s)… "
