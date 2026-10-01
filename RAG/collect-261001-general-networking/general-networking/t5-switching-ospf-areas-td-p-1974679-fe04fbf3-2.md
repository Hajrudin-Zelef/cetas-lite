---
id: collect-261001-general-networking/general-networking/t5-switching-ospf-areas-td-p-1974679-fe04fbf3-2
title: "t5-switching-ospf-areas-td-p-1974679-fe04fbf3"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-ospf-areas-td-p-1974679-fe04fbf3.md
source_anchor: ""
source_lines: [18, 126]
sha256: c247b4fc02ab17087554a18c2f799b2720ca22ed15cd2f2ef12715b1b7d77c21
---

# t5-switching-ospf-areas-td-p-1974679-fe04fbf3

			Other Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-29-2012 05:48 AM
Hello Mark,
Areas in OSPF can not be used arbitrarily. If you decide to use multiple areas in OSPF, one of these areas must be Area 0 working as the backbone area. All other areas must be physically attached to Area 0 - in other words, all other areas must include a router that has a physical connectivity (at least one interface) to the Area 0. The Area 0 must be only one and it must be continuous. It is not allowed in OSPF to daisy-chain areas (there are certain exceptions but let's not make things too complicated at this point).
Your network fails all these requirements. It does not have the Area 0 at all and it uses areas in a daisy-chain fashion. All these facts contribute to the fact that the routers do not exchange routing information.
Your network, quite frankly, is quite small to test multiple areas, but this is what you could do to test the area behavior:
- Merge areas 10, 20 and 30 into Area 0. In other words, all links between routers will be placed into backbone Area 0.
- Put R1 Fa0/1 into Area 1.
- Put R5 Fa0/0 into Area 2.
Then you should be able to see all routes on all routers, however, the networks 192.168.0.0/28 and 172.16.178.0/28 should be identified as OIA - inter-area OSPF routes.
Best regards,
Peter
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-29-2012 06:18 AM
Hello Mark,
the short answer is that there are too many areas in your lab setup that miss a connection with backbone area.
A correct OSPF multi-area design requires that all non-zero areas have a direct connection with area 0 that act as backbone of OSPF domain.
Each non-zero area connects through one or more ABR ( Area Border Router) to area 0, This means that these ABR are member of area 0 and at least a non-zero area at the same time.
OSPF route propagation is perfomed by ABR: they translate LSA that are internal to an area ( LSA type 1, LSA type 2) in LSA type 3 that are injected into area 0. Other ABRs can take these LSA type 3 and can generate other LSA type 3 that represent the same IP subnet to inject them to the non-backbone area(s) connected to them.
Route calculation in inter-area scenario is performed not in a link state manner: there is no knoweledge of topology outside the area and route is calculated using LSA type 3 and the route to the ABR that has (re)generated it
For this reason area 0 has to be in the middle of communication and routing between area x and area y.
In your lab the only working ABR nodes are those with a connection to area 0, other nodes do not perform translation between area x and area y even if they are member of both areas
Follow Peter's suggestions to fix your lab setup
Edit:
there are two answers because I thought the first was gone ...
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-29-2012 05:48 AM
Hello Mark,
Areas in OSPF can not be used arbitrarily. If you decide to use multiple areas in OSPF, one of these areas must be Area 0 working as the backbone area. All other areas must be physically attached to Area 0 - in other words, all other areas must include a router that has a physical connectivity (at least one interface) to the Area 0. The Area 0 must be only one and it must be continuous. It is not allowed in OSPF to daisy-chain areas (there are certain exceptions but let's not make things too complicated at this point).
Your network fails all these requirements. It does not have the Area 0 at all and it uses areas in a daisy-chain fashion. All these facts contribute to the fact that the routers do not exchange routing information.
Your network, quite frankly, is quite small to test multiple areas, but this is what you could do to test the area behavior:
- Merge areas 10, 20 and 30 into Area 0. In other words, all links between routers will be placed into backbone Area 0.
- Put R1 Fa0/1 into Area 1.
- Put R5 Fa0/0 into Area 2.
Then you should be able to see all routes on all routers, however, the networks 192.168.0.0/28 and 172.16.178.0/28 should be identified as OIA - inter-area OSPF routes.
Best regards,
Peter
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-29-2012 05:58 AM
Hello Mark,
correct OSPF multi area design requires that all areas are connected to the backbone 0 area with one or more ABR.
In your setup you have multiple areas, but not all areas have a direct connection to the backbone.
The normal flow of LSA is:
from original non zero area the LSAs are translated in LSA type 3 by ABR and inejcted in area 0.
Other ABR nodes can take these LSA type 3 that have single area scope ( area 0) and they can translate them in LSA type 3 that are injected in other non zero areas connected to the ABR. The new LSA type 3 will be flooded in the non zero area.
This is wanted and part of OSPF design because route calculation in inter -area scenario is not totally link state: it is performed by considering the LSA type 3 metric plus the intra area path to the ABR that originated it.
There is no topological information about what is outside the single area, so this calculation is more like distance vector.
For safety it is necessary to go through area 0 to allow communication and routing between two areas.
So the ABR with a connection to area 0 are working correctly, those without a connection with area 0 should keep databases of each area separated without performing translation in LSA type 3 between area x and area y even if they are connected ( member of ) both areas.
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-29-2012 06:18 AM
Hello Mark,
the short answer is that there are too many areas in your lab setup that miss a connection with backbone area.
A correct OSPF multi-area design requires that all non-zero areas have a direct connection with area 0 that act as backbone of OSPF domain.
Each non-zero area connects through one or more ABR ( Area Border Router) to area 0, This means that these ABR are member of area 0 and at least a non-zero area at the same time.
OSPF route propagation is perfomed by ABR: they translate LSA that are internal to an area ( LSA type 1, LSA type 2) in LSA type 3 that are injected into area 0. Other ABRs can take these LSA type 3 and can generate other LSA type 3 that represent the same IP subnet to inject them to the non-backbone area(s) connected to them.
Route calculation in inter-area scenario is performed not in a link state manner: there is no knoweledge of topology outside the area and route is calculated using LSA type 3 and the route to the ABR that has (re)generated it
For this reason area 0 has to be in the middle of communication and routing between area x and area y.
In your lab the only working ABR nodes are those with a connection to area 0, other nodes do not perform translation between area x and area y even if they are member of both areas
Follow Peter's suggestions to fix your lab setup
Edit:
there are two answers because I thought the first was gone ...
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-29-2012 06:37 AM
Thank you Giuseppe & Peter.
