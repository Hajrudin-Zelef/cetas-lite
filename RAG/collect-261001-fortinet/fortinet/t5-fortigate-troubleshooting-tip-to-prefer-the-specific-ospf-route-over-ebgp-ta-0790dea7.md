---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-to-prefer-the-specific-ospf-route-over-ebgp-ta-0790dea7
title: "t5-fortigate-troubleshooting-tip-to-prefer-the-specific-ospf-route-over-ebgp-ta--0790dea7"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-to-prefer-the-specific-ospf-route-over-ebgp-ta--0790dea7.md
source_anchor: ""
source_lines: [1, 145]
sha256: dbdbfde98882ec969e17b101932a8d3c19f9c6f5ad706146d7f9f2910717b3a2
---

# t5-fortigate-troubleshooting-tip-to-prefer-the-specific-ospf-route-over-ebgp-ta--0790dea7

Troubleshooting Tip: To prefer the specific OSPF route over EBGP
Description
This article describes how to prefer specific OSPF route over BGP.
Useful links:
Fortinet Documentation
Fortinet KB – https://kb.fortinet.com/kb/documentLink.do?externalID=FD30086
https://help.fortinet.com/cli/fos60hlp/60/index.htm#FortiOS/fortiOS-cli-ref/config/router/bgp.htm
Solution
To achieve the requirement that only default route should be preferred via OSPF and rest two routes 10.174.0.0/20 & 10.175.0.0/20 should be preferred over EBGP, on FORTIGATE1, increase EBGP Administrative Distance more than 110 only for 0.0.0.0/0 using below steps.
-Create the Access-list only for Default-Route with exact-match enable
#config router access-list
-Increase Distance for default route to 120 under admin-distance for EBGP neighbor 10.171.12.5
#config router bgp
-Soft Clear the BGP neighbor
#execute router clear bgp ip 10.171.12.5 soft
RESULTS:
-As per routing table only default route is preferred via OSPF and rest two routes 10.174.0.0/20 & 10.175.0.0/20 are preferred via EBGP.
 
FORTIGATE1 # get router info routing-table all
Routing table for VRF=0
Codes:
K - kernel, C - connected, S - static, R - RIP, B - BGP
O - OSPF, IA - OSPF inter area
N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
E1 - OSPF external type 1, E2 - OSPF external type 2
i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
* - candidate default
O*E2    0.0.0.0/0 [110/10] via 10.172.12.49, port3, 01:45:25
- Routing database is showing the default route via EBGP with modified Distance 120 and via OSPF with default distance 110. 
FORTIGATE1 # get router info routing-table database
Routing table for VRF=0
Codes:
K - kernel, C - connected, S - static, R - RIP, B - BGP
O - OSPF, IA - OSPF inter area
N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
E1 - OSPF external type 1, E2 - OSPF external type 2
i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
> - selected route, * - FIB route, p - stale info
B       0.0.0.0/0 [120/0] via 10.171.12.5, 00:03:36
This article describes how to prefer specific OSPF route over BGP.
Useful links:
Fortinet Documentation
Fortinet KB – https://kb.fortinet.com/kb/documentLink.do?externalID=FD30086
https://help.fortinet.com/cli/fos60hlp/60/index.htm#FortiOS/fortiOS-cli-ref/config/router/bgp.htm
- FORTIGATE2 is advertising 0.0.0.0/0, 10.174.0.0/20 and 10.175.0.0/20 routes to FORTIGATE1 via EBGP. 
- Also, FORTIGATE3 is advertising the same routes 0.0.0.0/0, 10.174.0.0/20 and 10.175.0.0/20 to FORTIGATE1 via OSPF.  
- FORTIGATE1 will install all 3 routes received from EBGP as EBGP has administrative distance of 20 which is less than administrative distance of OSPF which is 110.
- However, the requirement is that only default route should be preferred via OSPF and rest two routes 10.174.0.0/20 & 10.175.0.0/20 should be preferred via EBGP.
- FORTIGATE1 will install all 3 routes received from EBGP as EBGP has administrative distance of 20 which is less than administrative distance of OSPF which is 110.
- However, the requirement is that only default route should be preferred via OSPF and rest two routes 10.174.0.0/20 & 10.175.0.0/20 should be preferred via EBGP.
- All BGP routes are active in the routing table.
FORTIGATE1 # get router info routing-table all
FORTIGATE1 # get router info routing-table all
Routing table for VRF=0
Codes: 
K  - kernel, C - connected, S - static, R - RIP, B - BGP
O - OSPF, IA - OSPF inter area
N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
E1 - OSPF external type 1, E2 - OSPF external type 2
i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
* - candidate default
B*      0.0.0.0/0 [20/0] via 10.171.12.5, port1, 01:06:39
O - OSPF, IA - OSPF inter area
N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
E1 - OSPF external type 1, E2 - OSPF external type 2
i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
* - candidate default
B*      0.0.0.0/0 [20/0] via 10.171.12.5, port1, 01:06:39
B 10.174.0.0/20 [20/0] via 10.171.12.5, port1, 01:13:00
B 10.175.0.0/20 [20/0] via 10.171.12.5, port1, 01:12:31
-As per routing database, all OSPF routes are inactive
FORTIGATE1 # get router info routing-table database
Routing table for VRF=0
FORTIGATE1 # get router info routing-table database
Routing table for VRF=0
Codes: 
K  - kernel, C - connected, S - static, R - RIP, B - BGP
O - OSPF, IA - OSPF inter area
N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
E1 - OSPF external type 1, E2 - OSPF external type 2
i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
> - selected route, * - FIB route, p - stale info
O - OSPF, IA - OSPF inter area
N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
E1 - OSPF external type 1, E2 - OSPF external type 2
i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
> - selected route, * - FIB route, p - stale info
B *> 0.0.0.0/0 [20/0] via 10.171.12.5, port1, 01:06:49
O E2 0.0.0.0/0 [110/10] via 10.172.12.49, port3, 01:08:13
O E2 10.174.0.0/20 [110/6] via 10.172.12.49, port3, 01:11:33
B *> 10.174.0.0/20 [20/0] via 10.171.12.5, port1, 01:13:10
O E2 10.175.0.0/20 [110/6] via 10.172.12.49, port3, 01:11:33
B *> 10.175.0.0/20 [20/0] via 10.171.12.5, port1, 01:12:41
Solution
To achieve the requirement that only default route should be preferred via OSPF and rest two routes 10.174.0.0/20 & 10.175.0.0/20 should be preferred over EBGP, on FORTIGATE1, increase EBGP Administrative Distance more than 110 only for 0.0.0.0/0 using below steps.
-Create the Access-list only for Default-Route with exact-match enable
#config router access-list
edit "DEFAULT_ROUTE"
set comments "MATCH ONLY DEFAULT ROUTE"
config rule
edit 1
set prefix 0.0.0.0 0.0.0.0
set exact-match enable
next
end
next
end
#config router bgp
config admin-distance
edit 1
set neighbour-prefix 10.171.12.5 255.255.255.255
set route-list "DEFAULT_ROUTE"
set distance 120
next
end
end
#execute router clear bgp ip 10.171.12.5 soft
-As per routing table only default route is preferred via OSPF and rest two routes 10.174.0.0/20 & 10.175.0.0/20 are preferred via EBGP.
FORTIGATE1 # get router info routing-table all
Routing table for VRF=0
Codes:
K - kernel, C - connected, S - static, R - RIP, B - BGP
O - OSPF, IA - OSPF inter area
N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
E1 - OSPF external type 1, E2 - OSPF external type 2
i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
* - candidate default
O*E2    0.0.0.0/0 [110/10] via 10.172.12.49, port3, 01:45:25
B 10.174.0.0/20 [20/0] via 10.171.12.5, port1, 01:50:22
B 10.175.0.0/20 [20/0] via 10.171.12.5, port1, 01:49:53
FORTIGATE1 # get router info routing-table database
Routing table for VRF=0
Codes:
K - kernel, C - connected, S - static, R - RIP, B - BGP
O - OSPF, IA - OSPF inter area
N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
E1 - OSPF external type 1, E2 - OSPF external type 2
i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
> - selected route, * - FIB route, p - stale info
B       0.0.0.0/0 [120/0] via 10.171.12.5, 00:03:36
O E2 *> 0.0.0.0/0 [110/10] via 10.172.12.49, port3, 01:49:00
O E2 10.174.0.0/20 [110/6] via 10.172.12.49, port3, 01:52:20
B *> 10.174.0.0/20 [20/0] via 10.171.12.5, port1, 01:53:57
O E2 10.175.0.0/20 [110/6] via 10.172.12.49, port3, 01:52:20
B *> 10.175.0.0/20 [20/0] via 10.171.12.5, port1, 01:53:28
