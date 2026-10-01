---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-bgp-allowas-in-enable-or-as-override-when-local-as-ta-b693e214-1
title: "t5-fortigate-technical-tip-bgp-allowas-in-enable-or-as-override-when-local-as-ta-b693e214"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-bgp-allowas-in-enable-or-as-override-when-local-as-ta-b693e214.md
source_anchor: ""
source_lines: [1, 147]
sha256: c27734b2e27569df3c57c67bddaaf12b5487db6a43eb6fbfb52bb7a884e725d5
---

# t5-fortigate-technical-tip-bgp-allowas-in-enable-or-as-override-when-local-as-ta-b693e214

Technical Tip: BGP : 'allowas-in-enable' or 'as-override' when local AS in the AS-path conflicts with the received prefix
Description
This article explains how to configure 'allowas-in-enable' or 'as-override' when using MPLS with the same AS in different locations 
to avoid routing loops.
For example: Using AS 65001 at locations A and B.
If the route advertised by Location A is rejected by location B because of the AS path the route from location B will be rejected by location A.
This can be avoided by configuring 'allowas-in-enable' or 'as-override'.
Solution
BGP Config and route advertisement:
Location A:
config router bgp
    set as 65001
    set router-id 10.109.19.146
        config neighbor
            edit "10.109.16.172"                             <----- ISP IP.
                set soft-reconfiguration enable
                set remote-as 1111
            next
        end
        config redistribute "connected"
            set status enable
end
Route advertise by Location A to ISP:
Location A # get  router info  bgp neighbors 10.109.16.172 advertised-routes
BGP table version is 2, local router ID is 10.109.19.146
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal
Origin codes: i - IGP, e - EGP, ? - incomplete
   Network          Next Hop            Metric LocPrf Weight RouteTag Path
*> 10.109.16.0/20   10.109.19.146                      32768        0 ?
*> 10.201.0.0/20    10.109.19.146                      32768        0 ?
*> 10.205.0.0/20    10.109.19.146                      32768        0 ?
ISP:
config router bgp
    set as 1111
        config neighbor
            edit "10.109.19.146"                             <----- Location A.
                set soft-reconfiguration enable
                set remote-as 65001
            next 
            edit "10.109.16.73"                              <----- Location B.
                set soft-reconfiguration enable
                set remote-as 65001
            next
        end
end
Here, the ISP is advertising the route to Location B.
ISP # get router info bgp  neighbors 10.109.16.73 advertised-routes
BGP table version is 4, local router ID is 10.201.0.172
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal
Origin codes: i - IGP, e - EGP, ? - incomplete
   Network          Next Hop            Metric LocPrf Weight RouteTag Path
*> 10.109.16.0/20   10.109.19.146                          0        0 65001 ?
*> 10.201.0.0/20    10.109.19.146                          0        0 65001 ?
*> 10.205.0.0/20    10.109.19.146                          0        0 65001 ?
Total number of prefixes 3:
Location B.
config router bgp
    set as 65001
        config neighbor
            edit "10.109.16.172"                             <----- ISP.
                set soft-reconfiguration enable
                set remote-as 1111
            next
        end
end
The route advertise is visible by the ISP, but the route is not installed in the routing table of location B.
Location B #  get router info bgp neighbors 10.109.16.172 received-route
BGP table version is 1, local router ID is 192.168.100.99
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal
Origin codes: i - IGP, e - EGP, ? - incomplete
   Network          Next Hop            Metric LocPrf Weight RouteTag Path
*> 10.109.16.0/20   10.109.19.146                          0        0 1111 65001 ?
*> 10.201.0.0/20    10.109.19.146                          0        0 1111 65001 ?
*> 10.205.0.0/20    10.109.19.146                          0        0 1111 65001 
Total number of prefixes is 3.
In the routing table of location B, the route is not installed.
Location B # get router info routing-table details
Routing table for VRF=0
Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP
       O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
       * - candidate default
S* 0.0.0.0/0 [10/0] via 10.109.31.254, port1
                  [10/0] via 10.109.63.254, mgmt
C       10.109.16.0/20 is directly connected, port1
Solution:
There are two solutions to overcome this problem: 
- Configure the 'allowas-in-enable' in the BGP config of Location B.
Location B.
config router bgp
    set as 65001
        config neighbor
            edit "10.109.16.172"                            <----- ISP.
                set allowas-in-enable enable    
                set soft-reconfiguration enable
                set remote-as 1111
            next
        end
end
FortiGate by default allows three same AS with the command 'allowas-in-enable', to allow more than three AS, then use the command 'allowas-in <number>'.
Location B # get router info routing-table details
Routing table for VRF=0
Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP
       O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
       * - candidate default
S*      0.0.0.0/0 [10/0] via 10.109.31.254, port1
                  [10/0] via 10.109.63.254, mgmt
C       10.109.16.0/20 is directly connected, port1
B       10.201.0.0/20 [20/0] via 10.109.19.146, port1, 00:00:38
B       10.205.0.0/20 [20/0] via 10.109.19.146, port1, 00:00:38
Location B # get router info bgp network 10.205.0.0/20
BGP routing table entry for 10.205.0.0/20
Paths: (1 available, best #1, table Default-IP-Routing-Table)
  Not advertised to any peer
  1111 65001
    10.109.19.146 from 10.109.16.172 (10.201.0.172)
      Origin incomplete metric 0, localpref 100, valid, external, best
      Last update: Sun Oct  6 13:13:19 2019
-  Enable the 'as-override' in the BGP config of the ISP. For IPv6, use the command 'set as-override6 enable' instead. This command will replace the AS number of the receiving neighbor only, and the rest of the AS number remains intact as below. This means that the conflict AS will be replaced with the ISP AS number.
ISP:
config router bgp
    set as 1111
        config neighbor
            edit "10.109.19.146"
                set soft-reconfiguration enable
                set as-override enable
                set remote-as 65001
            next
            edit "10.109.16.73"
                set soft-reconfiguration enable
                set as-override enable
                set remote-as 65001
            next
        end
end
After enabling the as-override, the routes are available in the routing table of Location B.
Location B # get  router  info routing-table details
Routing table for VRF=0
Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP
       O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
