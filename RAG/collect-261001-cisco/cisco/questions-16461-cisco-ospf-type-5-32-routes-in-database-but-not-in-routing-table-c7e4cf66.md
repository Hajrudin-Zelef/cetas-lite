---
id: collect-261001-cisco/cisco/questions-16461-cisco-ospf-type-5-32-routes-in-database-but-not-in-routing-table-c7e4cf66
title: "questions-16461-cisco-ospf-type-5-32-routes-in-database-but-not-in-routing-table-c7e4cf66"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/questions-16461-cisco-ospf-type-5-32-routes-in-database-but-not-in-routing-table-c7e4cf66.md
source_anchor: ""
source_lines: [1, 74]
sha256: cbff6a208d06c0c54255504624f4e29107fb374e62c1cb50ca69ce9e9ed5403d
---

# questions-16461-cisco-ospf-type-5-32-routes-in-database-but-not-in-routing-table-c7e4cf66

Our network runs OSPF for IPv4 on Cisco hardware. The following image shows the relevant topology:
Topology
Problem Statement
We have a Cisco ASA firewall, which also is the VPN gateway for incoming VPN connections. As said, we are running OSPF as internal routing protocol. The ASA performs reverse route injection (RRI) for the VPN clients, which means it creates a static /32 route for every VPN connection. Those /32 routes are then redistributed into the OSPF domain.
The Router on the right is our perimeter router running BGP. In between there is a Multilayer Switch (Cat6500 w/ Sup720-3B and IPServices) also running OSPF. The problem is, that those /32 routers are visible in the OSPF database both at the MLS and the Router, but are not inserted into the routing table on the MLS. At the router, they however are in the routing table.
Debugging Output
There are two such subnets (two different VPN profiles), which are xx.xx.37.0/24 and xx.xx.202.0/24. As those /32 routers are redistributed into OSPF, they appear as LSA type 5 in the rest of the routing domain, however, they are not injected into the routing table at the MLS.
We can see those LSAs at the database:
show ip ospf database external | i xx.xx.202.
  Link State ID: xx.xx.202.5 (External Network Number )
  Link State ID: xx.xx.202.6 (External Network Number )
  Link State ID: xx.xx.202.8 (External Network Number )
If we look at the specific LSA, we also can see that the routing bit is set, hence the LSA is eligible for OSPF routing:
show ip ospf database external xx.xx.202.5
            OSPF Router with ID (xx.xx.207.4) (Process ID 2)
                Type-5 AS External Link States
  Routing Bit Set on this LSA
  LS age: 1887
  Options: (No TOS-capability, DC)
  LS Type: AS External Link
  Link State ID: xx.xx.202.5 (External Network Number )
  Advertising Router: xx.xx.192.20
  LS Seq Number: 80000001
  Checksum: 0x8147
  Length: 36
  Network Mask: /32
        Metric Type: 2 (Larger than any link state path)
        TOS: 0
        Metric: 20
        Forward Address: xx.xx.192.21
        External Route Tag: 0
Further, debugging the SPF of the OSPF process shows, that the partial update is received and processed, and no error can be found (debug ip ospf spf):
Feb  2 16:17:40.537: OSPF: Schedule partial SPF - type 5 id xx.xx.202.5 adv rtr xx.xx.192.20
Feb  2 16:17:40.537: OSPF: Service partial SPF 0/1/0
Feb  2 16:17:40.537: OSPF process partial spfQ entry
Feb  2 16:17:40.537: OSPF process partial spfQ entry
Feb  2 16:17:40.537: OSPF process partial spfQ LSA id xx.xx.202.5: mask 255.255.255.255, type 5 adv_rtr xx.xx.192.20, age 1, seq 0x80000001 (Area dummy area)
Feb  2 16:17:40.537: OSPF: Start partial processing Type 5 External LSA xx.xx.202.5, mask 255.255.255.255, adv xx.xx.192.20, age 1, seq 0x80000001, metric 20, metric-type 2
Feb  2 16:17:40.537:    Add path: next-hop xx.xx.192.21, interface Vlan511
Feb  2 16:17:40.537: OSPF: insert route list LS ID xx.xx.202.5, type 5, adv rtr xx.xx.192.20
Feb  2 16:17:40.537: OSPF process partial spfQ entry
However, the route is not inserted into the routing table. As a work around, we added a static route using the /24 subnet. This however is no reason why the /32 subnet is not placed into the table. While the AD is lower (1 static route vs 110 OSPF external), the longest prefix match rule is evaluated first, hence the /32 subnet should be added.
Interestingly, at the Router behind the MLS, the /32 route is added to the routing table as expected.
Further Information
MLS is a Cisco 6500 Chassis w/ Sup720-3B (MSFC3) SUP running "s72033-ipservicesk9-mz.122-33.SXI14.bin" software. The Router is a Cisco ASR1001-F router, running "asr1000rp1-advipservicesk9.03.04.06.S.151-3.S6.bin" software.
There are no routing filters in place, the OSPF config of the devices is pretty basic. OSPF adjacencies are OK, all other routes or OK.
OSPF configuration of MLS:
router ospf 2
 router-id xx.xx.207.4
 log-adjacency-changes
 auto-cost reference-bandwidth 100000
 area 0 authentication message-digest
 passive-interface default
 no passive-interface ...
 network 0.0.0.0 255.255.255.255 area 0
OSPF configuration Router:
router ospf 2
 router-id xx.xx.207.1
 auto-cost reference-bandwidth 100000
 area 0 authentication message-digest
 traffic-share min across-interfaces
 passive-interface default
 no passive-interface ...
 default-information originate always metric-type 1
New Information
After some further troubleshooting, I may have found the issue, however, have no idea how to solve it. Looking at the firewall, the generated static route looks like this:
S xx.xx.202.13 255.255.255.255 [1/0] via xx.xx.192.21, outside
The default behaviour of RRI is to set the next-hop of the generated static route to the default gateway of the ASA interface where the crypto map is applied (which is outside in our case). In our case, it is xx.xx.192.21, which is the interface of the MLS!
According to Cisco Documentation, a non-zero forwarding address is set, if the following conditions are met:
  OSPF is enabled on the ASBR’s next hop interface AND ASBR’s next hop interface is non-passive under OSPF AND ASBR’s next hop interface is not point-to-point AND ASBR’s next hop interface is not point-to-multipoint AND ASBR’s next hop interface address falls under the network range specified in the router ospf command.
All these conditions are met, as we need the outside interface (which is next hop) enabled in OSPF, as it is the interface pointing to the MLS.
As written above, the LSA at the MLS contains the forwarding address of xx.xx.192.21 which is the address of the MLS in the subnet between MPLS and ASA. According to the documentation linked above, the forwarding address needs to be known as internal OSPF route. This is given at the Router, and the routing entries are added to the routing table. This however is not given at the MLS.
Question
Why are the routes not added to the routing table? I appreciate any idea.
