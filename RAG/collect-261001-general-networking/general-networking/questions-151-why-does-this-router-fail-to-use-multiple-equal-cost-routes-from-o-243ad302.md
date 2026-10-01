---
id: collect-261001-general-networking/general-networking/questions-151-why-does-this-router-fail-to-use-multiple-equal-cost-routes-from-o-243ad302
title: "questions-151-why-does-this-router-fail-to-use-multiple-equal-cost-routes-from-o-243ad302"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/questions-151-why-does-this-router-fail-to-use-multiple-equal-cost-routes-from-o-243ad302.md
source_anchor: ""
source_lines: [1, 153]
sha256: ab7743019c151f6dca32ff1a473f0512fd67c12e7c3895848166e5241dc5505a
---

# questions-151-why-does-this-router-fail-to-use-multiple-equal-cost-routes-from-o-243ad302

A Brocade router running NetIron 5.2, and a Dynamips emulation running IOS 12.4 both fail to dependably use equal cost paths in what seems to be a very simple configuration. I say "dependably" because I did observe very briefly the expected behavior from the Brocade, but it was after a longish period of inactivity and vanished as soon as I perturbed the system.
In case it is relevant, I am setting up an anycast destination for a stateless network service. This will be a site-specific configuration, with little or no network path length diversity, so the purpose of the anycast capability is mainly for fail-over redundancy and partly for load sharing. My question concerns the load sharing function.
My prototype network looks like this.
R1 <--N1--> Rcore <--N2--> R2
R1 and R2 advertise the anycast destination via OSPF. If Rcore has identical cost to on its N1 and N2 interfaces, and the anycast destination has identical cost at each of the other routers, I would expect to find two equal cost paths in the Rcore routing table. I don't, at least not always, and not even predictably. Why not?
Among other, potentially relevant details are the following.
- Rcore is an ABR connectingN1 andN2 , which are in separate NSSA areas, to the backbone (not shown).
- Inspection of the Rcore LSDB confirms that the expected LSAs fromR1 andR2 are correctly installed, and contain the anycast destination with the correct metric.
I have considered three possible explanations.
- I don't understand how to configure and operate ECMP on these platforms. This is likely, and I hope to find enlightenment here that I could not find in the vendor docs.
- I don't understand the interaction among OSPF Areas, ECMP, anycast, etc. I don't think this is very likely, but I am prepared to learn something anyway.
- These two vendors implementations suffer from the same bug, or mis-feature, or something. I suppose it could happen.
Edit Added example config and output for Dynamips/Dynagen implementation.
Configs
Router R1
!
version 12.4
!
hostname R1
!
interface Loopback0
 ip address 1.1.1.1 255.255.255.255
 ip ospf cost 1
!
interface FastEthernet0/0
 ip address 10.1.0.2 255.255.255.252
 ip ospf cost 2
!
router ospf 1
 router-id 10.1.0.2
 log-adjacency-changes
 area 0.0.0.1 nssa no-summary
 passive-interface Loopback0
 network 1.1.1.1 0.0.0.0 area 0.0.0.1
 network 10.1.0.0 0.0.0.3 area 0.0.0.1
!
line console 0
 exec-timeout 0 0
!
end
Router R2
!
version 12.4
!
hostname R2
!
interface Loopback0
 ip address 1.1.1.1 255.255.255.255
 ip ospf cost 1
!
interface FastEthernet0/0
 ip address 10.2.0.2 255.255.255.252
 ip ospf cost 2
!
router ospf 1
 router-id 10.2.0.2
 log-adjacency-changes
 area 0.0.0.2 nssa no-summary
 passive-interface Loopback0
 network 1.1.1.1 0.0.0.0 area 0.0.0.2
 network 10.2.0.0 0.0.0.3 area 0.0.0.2
!
line console 0
 exec-timeout 0 0
!
end
Router Rcore
!
version 12.4
!
hostname Rcore
!
interface Loopback0
 ip address 10.0.0.1 255.255.255.255
 ip ospf cost 4
!
interface FastEthernet0/0
 ip address 10.1.0.1 255.255.255.252
 ip ospf cost 2
!
interface FastEthernet0/1
 ip address 10.2.0.1 255.255.255.252
 ip ospf cost 2
!
router ospf 1
 log-adjacency-changes
 area 0.0.0.1 nssa no-summary
 area 0.0.0.2 nssa no-summary
 passive-interface Loopback0
 network 10.0.0.1 0.0.0.0 area 0.0.0.0
 network 10.1.0.0 0.0.0.3 area 0.0.0.1
 network 10.2.0.0 0.0.0.3 area 0.0.0.2
!
line console 0
 exec-timeout 0 0
!
end
Output
Router Rcore
Router LSA from R1
Rcore#sh ip ospf database router 10.1.0.2
            OSPF Router with ID (10.0.0.1) (Process ID 1)
        Router Link States (Area 0.0.0.1)
  LS age: 1618
  Options: (No TOS-capability, DC)
  LS Type: Router Links
  Link State ID: 10.1.0.2
  Advertising Router: 10.1.0.2
  LS Seq Number: 80000002
  Checksum: 0x726F
  Length: 48
  Number of Links: 2
    Link connected to: a Stub Network
     (Link ID) Network/subnet number: 1.1.1.1
     (Link Data) Network Mask: 255.255.255.255
      Number of TOS metrics: 0
       TOS 0 Metrics: 1
    Link connected to: a Transit Network
     (Link ID) Designated Router address: 10.1.0.2
     (Link Data) Router Interface address: 10.1.0.2
      Number of TOS metrics: 0
       TOS 0 Metrics: 2
Router LSA from R2
Rcore#sh ip ospf database router 10.2.0.2
            OSPF Router with ID (10.0.0.1) (Process ID 1)
        Router Link States (Area 0.0.0.2)
  LS age: 1766
  Options: (No TOS-capability, DC)
  LS Type: Router Links
  Link State ID: 10.2.0.2
  Advertising Router: 10.2.0.2
  LS Seq Number: 80000002
  Checksum: 0x8A53
  Length: 48
  Number of Links: 2
    Link connected to: a Stub Network
     (Link ID) Network/subnet number: 1.1.1.1
     (Link Data) Network Mask: 255.255.255.255
      Number of TOS metrics: 0
       TOS 0 Metrics: 1
    Link connected to: a Transit Network
     (Link ID) Designated Router address: 10.2.0.2
     (Link Data) Router Interface address: 10.2.0.2
      Number of TOS metrics: 0
       TOS 0 Metrics: 2
Routing table excerpt
Rcore#sh ip route 1.1.1.1
Routing entry for 1.1.1.1/32
  Known via "ospf 1", distance 110, metric 3, type intra area
  Last update from 10.1.0.2 on FastEthernet0/0, 00:30:54 ago
  Routing Descriptor Blocks:
  * 10.1.0.2, from 10.1.0.2, 00:30:54 ago, via FastEthernet0/0
      Route metric is 3, traffic share count is 1
