---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-redistribute-bgp-routes-learnt-through-different-bgp-c-3ca36cd1
title: "fortigate-3-technical-tip-redistribute-bgp-routes-learnt-through-different-bgp-c-3ca36cd1"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-redistribute-bgp-routes-learnt-through-different-bgp-c-3ca36cd1.md
source_anchor: ""
source_lines: [1, 106]
sha256: ea324d10eb183f9d1596b14edb7632cf03cedc26d0282c84e0f463dd7f2e1842
---

# fortigate-3-technical-tip-redistribute-bgp-routes-learnt-through-different-bgp-c-3ca36cd1

Technical Tip: Redistribute BGP routes learnt through different BGP Communities into OSPF.
Description
This article describes how to redistribute BGP routes learned through different BGP Communities into OSPF. This solution requires FortiOS 6.4.3 or above (Bug ID# 644461).
Scope
FortiGate.
Solution
Consider the following example case:
- ROUTER 1 and FortiGate are in a BGP neighborship with ROUTER 1 advertising 192.168.233.0/24 to FortiGate using BGP-COMMUNITY: 6500:100 and 192.168.234.0/24 using BGP-COMMUNITY: 6501:101.
- FortiGate and ROUTER 2 are in OSPF Neighborship with each other.
- It is desired for only BGP routes learned using BGP-COMMUNITY: 6500:100 to be redistributed into OSPF.
- Since BGP-COMMUNITY is specific to BGP protocol, it is not possible to use BGP-COMMUNITY to filter BGP routes to be redistributed into OSPF directly, but it is possible to use tags to tag the incoming routes learned using BGP-COMMUNITIES and then use the tags to redistribute the BGP routes into OSPF.
- In the above example, verify the routes associated with the BGP community by using the following command:
get router info bgp network <route learnt>
- Next, add an inbound route map on the BGP configuration on the FortiGate to tag the routes learned through BGP-COMMUNITY '6500:100' with a value of '200' and tag the routes learned through BGP-COMMUNITY '6501:101' with a value of '100'.
- Add the community list and set the match-action to 'permit' for the BGP-COMMUNITIES.
config router community-list
edit "6500:100"
config rule
edit 1
set action permit
set match "6500:100"
next
end
next
edit "6501:101"
config rule
edit 1
set action permit
set match "6501:101"
next
end
next
end
- Create an inbound route map to tag the routes learned through their respective BGP-COMMUNITIES.
config router route-map
edit "inbound"
config rule
edit 1
set match-community "6500:100" <----- Created above.
set set-tag 200
next
edit 2
set match-community "6501:101" <----- Created above.
set set-tag 100
next
end
end
- Reference the inbound route-map in the BGP configuration of the FortiGate.
config router bgp
set as 101
set router-id 172.16.1.1
config neighbor
edit "172.16.1.102"
set interface "lan"
set remote-as 100
set route-map-in "inbound" <----- Inbound route-map.
next
end
end
- Create a route map for the routes advertised by OSPF to match-tag 200 assigned to BGP routes learned using BGP-COMMUNITY '6500:100'.
config router route-map
edit "outbound-OSPF"
config rule
edit 1
set match-tag 200 <----- Use routes with tag 200 only.
next
end
next
end
- Reference the route map for 'outbound-OSPF' in OSPF configuration of the FortiGate.
config router ospf
set router-id 172.16.88.1
config area
edit 0.0.0.0
next
end
config ospf-interface
edit "ospf"
set interface "lan"
next
end
config network
edit 1
set prefix 0.0.0.0 0.0.0.0
next
end
config redistribute "connected"
end
config redistribute "static"
end
config redistribute "rip"
end
config redistribute "bgp"
set status enable <----- Set redistribution status enable.
set routemap "outbound-OSPF" <----- Set route-map.
end
config redistribute "isis"
end
end
Verification:
The routing table of ROUTER 2 should only see the 192.168.233.0/24 learned via OSPF from FortiGate.
The following commands may be useful:
get router info ospf neighbor
get router info routing-table all
get router info ospf network
Related documents:
