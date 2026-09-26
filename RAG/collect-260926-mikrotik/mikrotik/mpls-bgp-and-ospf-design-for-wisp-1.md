---
id: collect-260926-mikrotik/mikrotik/mpls-bgp-and-ospf-design-for-wisp-1
title: "mpls-bgp-and-ospf-design-for-wisp"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mpls-bgp-and-ospf-design-for-wisp.md
source_anchor: ""
source_lines: [1, 195]
sha256: 6b88099684c91e1ad891795ad4c8e28ed5fd89ebdaaf62c29809be2c3e151aa4
---

# mpls-bgp-and-ospf-design-for-wisp

I am working on a new design for our WISP. We currently have around 800 customers and around 50 towers. We are getting ready to order 2 Mikrotik CCR1072’s for our core routers and upgrade from a bridged network to MPLS.

Our needs are as follows

Clients get a static NATed IPv4 address or static Public

IPv6 support

Ability to create tunnels for clients with multiple sites

Redundant connections for backhauls at different points of entry

Redundant upstream providers at different locations. Our main 10G connections comes from 1 locations while we have a couple 200M connections at different locations. redundancy will need to be set on a site (tower) level.

I would like to run iBGP for our public/nated ip addresses on the internal side and we will be using eBGP on the upstream side. We have plans to get 2 10G connections dropped to our cabinet and use eBGP to create failover between the CCR1072’s.

I have created a test lab using 4 Mikrotik routers joined as a ring using MPLS and OSPF. I configured a unique loopback address per router and /30s between the routers. Failover when breaking on of the links takes anywhere from 1 second to 30 seconds. Currently iBGP is not working in the lab, each peer is stuck at open sent. I am not sure where I went wrong on the setup. Any help or suggestions would be welcomed.


```
Core Router
/interface bridge
add name=Loopback
/interface ethernet
set [ find default-name=combo1 ] mtu=1526
set [ find default-name=ether1 ] l2mtu=1800 mtu=1800
set [ find default-name=ether2 ] l2mtu=1800 mtu=1800
/interface vlan
add interface=combo1 mtu=1526 name=vlan205 vlan-id=205
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/routing bgp instance
set default as=100 router-id=10.255.0.1
/routing ospf instance
set [ find default=yes ] distribute-default=always-as-type-1 mpls-te-area=backbone mpls-te-router-id=Loopback router-id=10.255.0.1
/ip neighbor discovery-settings
set discover-interface-list=!dynamic
/ip address
add address=192.168.33.1/30 interface=ether1 network=192.168.33.0
add address=192.168.33.14/30 interface=ether2 network=192.168.33.12
add address=10.255.0.1 interface=Loopback network=10.255.0.1
add address=192.168.205.200/24 interface=vlan205 network=192.168.205.0
/ip route
add distance=1 gateway=192.168.205.1
/mpls ldp
set enabled=yes lsr-id=10.255.0.1 transport-address=10.255.0.1
/mpls ldp interface
add interface=ether1
add interface=ether2
/routing bgp peer
add name=Core-R2 remote-address=10.255.0.2 remote-as=100 ttl=default
add name=Core-R3 remote-address=10.255.0.3 remote-as=100 ttl=default
add name=Core-R4 remote-address=10.255.0.4 remote-as=100 ttl=default
/routing ospf interface
add interface=ether1 network-type=broadcast use-bfd=yes
add interface=ether2 network-type=broadcast use-bfd=yes
/routing ospf network
add area=backbone network=192.168.33.0/24
add area=backbone network=10.255.0.1/32
/system identity
set name=Core
Router 2
/interface bridge
add name=Loopback
/interface ethernet
set [ find default-name=ether1 ] l2mtu=1800 mtu=1800
set [ find default-name=ether2 ] l2mtu=1800 mtu=1800
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/routing bgp instance
set default as=100 router-id=10.255.0.2
/routing ospf instance
set [ find default=yes ] mpls-te-area=backbone mpls-te-router-id=Loopback router-id=10.255.0.2
/ip neighbor discovery-settings
set discover-interface-list=!dynamic
/ip address
add address=192.168.33.2/30 interface=ether1 network=192.168.33.0
add address=192.168.33.5/30 interface=ether2 network=192.168.33.4
add address=10.255.0.2 interface=Loopback network=10.255.0.2
/mpls ldp
set enabled=yes lsr-id=10.255.0.2 transport-address=10.255.0.2
/mpls ldp interface
add interface=ether1
add interface=ether2
/routing bgp peer
add name=R2-Core remote-address=10.255.0.1 remote-as=100 ttl=default
add name=R2-R3 remote-address=10.255.0.3 remote-as=100 ttl=default
add name=R2-R4 remote-address=10.255.0.4 remote-as=100 ttl=default
/routing ospf interface
add interface=ether1 network-type=broadcast use-bfd=yes
add interface=ether2 network-type=broadcast use-bfd=yes
/routing ospf network
add area=backbone network=192.168.33.0/24
add area=backbone network=10.255.0.2/32
/system identity
set name=R2
Router 3
/interface bridge
add name=Loopback
/interface ethernet
set [ find default-name=ether1 ] l2mtu=1800 mtu=1800
set [ find default-name=ether2 ] l2mtu=1800 mtu=1800
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/routing bgp instance
set default as=100 router-id=10.255.0.3
/routing ospf instance
set [ find default=yes ] mpls-te-area=backbone mpls-te-router-id=Loopback redistribute-connected=as-type-1 router-id=10.255.0.3
/ip neighbor discovery-settings
set discover-interface-list=!dynamic
/ip address
add address=192.168.33.6/30 interface=ether1 network=192.168.33.4
add address=192.168.33.9/30 interface=ether2 network=192.168.33.8
add address=10.255.0.3 interface=Loopback network=10.255.0.3
add address=192.168.10.1/24 interface=ether3 network=192.168.10.0
/mpls ldp
set enabled=yes lsr-id=10.255.0.3 transport-address=10.255.0.3
/mpls ldp interface
add interface=ether1
add interface=ether2
/routing bgp network
add network=192.168.10.0/24
/routing bgp peer
add name=R3-Core remote-address=10.255.0.1 remote-as=100 ttl=default
add name=R3-R2 remote-address=10.255.0.2 remote-as=100 ttl=default
add name=R3-R4 remote-address=10.255.0.4 remote-as=100 ttl=default
/routing ospf interface
add interface=ether1 network-type=broadcast use-bfd=yes
add interface=ether2 network-type=broadcast use-bfd=yes
/routing ospf network
add area=backbone network=192.168.33.0/24
add area=backbone network=10.255.0.3/32
/system identity
set name=R3
Router 4
/interface bridge
add name=Loopback
/interface ethernet
set [ find default-name=ether1 ] l2mtu=1800 mtu=1800
set [ find default-name=ether2 ] l2mtu=1800 mtu=1800
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/routing bgp instance
set default as=100 router-id=10.255.0.4
/routing ospf instance
set [ find default=yes ] mpls-te-area=backbone mpls-te-router-id=Loopback redistribute-connected=as-type-1 router-id=10.255.0.4
/ip neighbor discovery-settings
set discover-interface-list=!dynamic
/ip address
add address=192.168.33.10/30 interface=ether1 network=192.168.33.8
add address=192.168.33.13/30 interface=ether2 network=192.168.33.12
add address=10.255.0.4 interface=Loopback network=10.255.0.4
/mpls ldp
set enabled=yes lsr-id=10.255.0.4 transport-address=10.255.0.4
/mpls ldp interface
add interface=ether1
add interface=ether2
/routing bgp peer
add name=R4-Core remote-address=10.255.0.1 remote-as=100 ttl=default
add name=R4-R2 remote-address=10.255.0.2 remote-as=100 ttl=default
add name=R4-R3 remote-address=10.255.0.3 remote-as=100 ttl=default
/routing ospf interface
add interface=ether1 network-type=broadcast use-bfd=yes
add interface=ether2 network-type=broadcast use-bfd=yes
/routing ospf network
add area=backbone network=192.168.33.0/24
add area=backbone network=10.255.0.4/32
/system identity
set name=R4
```

             
            
           
          
            
            
              Your BGP peer remote-addresses are wrong.  These should be the peer’s address on the /30 links.

             
            
           
          
            
            
              Here are a couple of presentations I’ve done at different MUMs that may help you out with design for this type of network. Hope this helps!

**BGP as an IGP for Carrier/Enterprise Networks**

https://mum.mikrotik.com//presentations/US13/kevin.pdf

**ISP Architecture – MPLS Overview, Design and Implementation for WISPs.**

https://mum.mikrotik.com//presentations/US16/presentation_3327_1462279781.pdf

