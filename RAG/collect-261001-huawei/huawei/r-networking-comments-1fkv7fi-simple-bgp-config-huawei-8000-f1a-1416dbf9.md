---
id: collect-261001-huawei/huawei/r-networking-comments-1fkv7fi-simple-bgp-config-huawei-8000-f1a-1416dbf9
title: "r-networking-comments-1fkv7fi-simple-bgp-config-huawei-8000-f1a-1416dbf9"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/r-networking-comments-1fkv7fi-simple-bgp-config-huawei-8000-f1a-1416dbf9.md
source_anchor: ""
source_lines: [1, 72]
sha256: b8cb4282578643c9111d8ab0fb54bbe98686d5c1f76bbeb7697c2ed06865112d
---

# r-networking-comments-1fkv7fi-simple-bgp-config-huawei-8000-f1a-1416dbf9

Simple BGP config (Huawei 8000 F1A) 
        
        
        
    
    
    Hi all,
      this is the situation we have [it's my first experience with BGP]:
Two routers (with loopback0 10.0.2.1 and 10.0.2.2), each with an eBGP connection to the ISP. An iBGP sessione in between.
    
I want to avoid to become AS transit
This config on one of the routers doesn't announce the route we got from RIPE
[~8KF1A-02]dis curr conf bgp
      bgp XX9402
router-id AA.YY.130.46
peer 10.0.2.1 as-number XX9402
peer 10.0.2.1 connect-interface LoopBack0
peer AA.BB.CC.185 as-number XXX74
peer AA.BB.CC.185 ebgp-max-hop 3
peer AA.BB.CC.185 connect-interface GigabitEthernet0/1/0
    
      ipv4-family unicast
undo synchronization
aggregate XXX.YYY.250.0 255.255.255.0 as-set
network XXX.YYY.250.0 255.255.255.0
peer 10.0.2.1 enable
peer 10.0.2.1 next-hop-local
peer AA.BB.CC.185 enable
peer AA.BB.CC.185 as-path-filter FROM_WIND export
    
[~8KF1A-02_FASTWEB]dis ip int brief | exc unass
      Info: It will take a long time if the content you search is too much or the string you input is too long, you can press CTRL_C to break.
*down: administratively down
!down: FIB overload down
^down: standby
(l): loopback
(s): spoofing
(d): Dampening Suppressed
(E): E-Trunk down
(td): transceiver unmatch down
The number of interface that is UP in Physical is 8
The number of interface that is DOWN in Physical is 53
The number of interface that is UP in Protocol is 8
The number of interface that is DOWN in Protocol is 53
    
      Interface                         IP Address/Mask      Physical   Protocol VPN
25GE0/1/28(100M)                  172.16.3.253/23down       down     --
Eth-Trunk1                        172.16.31.2/30up         up       --
GigabitEthernet0/1/0(10G)         AA.YY.130.46/29      up         up       --
LoopBack0                         10.0.2.2/32up         up(s)    --
LoopBack1                         XXX.YYY.250.1/32     up         up(s)    -- <<<<<<<<<<<
LoopBack1023                      128.21.245.83/16up         up(s)    l3vpn
    
      [~8KF1A-02]dis bgp routing-table peer AA.BB.CC.18513.156.51.185 advertised-routes
[~8KF1A-02]
    
I don't know where I'm doing wrong.
Would you have any hint for me, please?
Panatism
Section des commentaires
Which route? Can you explain more here?
What does the config for FROM_WIND look like?
Are you actually redistributing any local routes into the BGP RIB? You typically use a bgp network command or redistribute connected to get local routes into the BGP RIB. Once they are in the RIB, the routes can be advertised to BGP peers.
A simple config to avoid becoming a transit AS is to use a regex AS-path filter that looks like this: "^$". This filter only accepts routes where the AS-path is empty, which means the route originated in your own AS. That means that any route you receive from another AS will not be sent to whichever peer the AS-path filter is applied to.
Sorry , you should have read "prefix from RIPE" instead of "route rom RIPE".
We get the FRT from both the ISPs, and each router must announce it to the other peer.
So each router should have the routing table mainly populated by eBGP and then iBGP from the other router. The thing is that with the config I came to rge router itself showed that no prefixes were announced.
It is weird.
Panatism
The trick is to only announce your own ASN and your customer ASNs (although I imagine you don't have any customer ASNs in this setup).
You are telling the router to export as-paths that match the FROM_WIND filter. Check that.
You also want to make sure your upstreams set a max-prefix for your sessions, to ensure you don't send them too many prefixes. If you only have 1-2 prefixes, ask them to set a max-prefix of 10. It doesn't need to be tight and precise, it needs to protect you from typos and accidents.
