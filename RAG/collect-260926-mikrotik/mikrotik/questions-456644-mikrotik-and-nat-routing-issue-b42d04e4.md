---
id: collect-260926-mikrotik/mikrotik/questions-456644-mikrotik-and-nat-routing-issue-b42d04e4
title: "questions-456644-mikrotik-and-nat-routing-issue-b42d04e4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-456644-mikrotik-and-nat-routing-issue-b42d04e4.md
source_anchor: ""
source_lines: [1, 43]
sha256: 104d3a67a2ed3ef3b3eb2b55a7d80a46944faca4a0955e6a774e557195d3955e
---

# questions-456644-mikrotik-and-nat-routing-issue-b42d04e4

I have basic NAT/Routing problem with Mikrotik RB750 that I've been unable to solve over the past days. From our ISP we have 26 IP addresses: 10.10.10.192/27, with 10.10.10.193 being the gateway and 10.10.10.194 the first available IP.
What I need is that everything connected to ether2 gets a public IP from the DHCP server, and everything connected to ether3 gets a local IP from another DHCP (192.168.100.0/24). All clients should have internet access (I'll figure out bandwidth throttling later) and optimally just 'see' each other (all boxes are Win7, I guess this can ultimately be handled with VPN).
Here is my setup: ether1 (10.10.10.194) is connected directly to ISP.
20 clients connected to ether2(10.10.10.195), and another 20 to ether3(10.10.10.196) (both through same 24 port switches).
This is my setup, which doesn't work, all 20 clients from ether2 can access the internet, though all comm. seems to come from 10.10.10.194 (is this due to the masquerade on ether1?), and ether3 can't access the internet at all.
I think that I need to masquerade ether3, and SNAT/DNAT or NETMAP ether2, but that doesn't work either, I guess that I need to somehow 'wire' both ether2+3 to ether1.
Address list:
 #   ADDRESS            NETWORK         INTERFACE                                                          
 0   ;;; public
     10.10.10.194/32  10.10.10.192  ether1-gateway
 1   ;;; inner DHCP
     192.168.100.0/24   192.168.100.0   ether3-private
 2   ;;; public
     10.10.10.195/32  10.10.10.192  ether2-pub
 3   ;;; public
     10.10.10.196/32  10.10.10.192  ether3-private
NAT
 0   ;;; ether3 nat
     chain=srcnat action=src-nat to-addresses=10.10.10.196 
     src-address=192.168.100.0/24 out-interface=ether3-private 
 1   ;;; ether3 nat
     chain=dstnat action=dst-nat to-addresses=192.168.100.0/24 
     in-interface=ether3-private 
 2   ;;; ether1 masquerade
     chain=srcnat action=masquerade to-addresses=10.10.10.194 
     out-interface=ether1-gateway 
Routes:
 #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
 0 A S  0.0.0.0/0                          ether1-gateway            1
 2 A S  10.10.10.192/27  10.10.10.195  ether2-pub                1
 3 ADC  10.10.10.192/32  10.10.10.195  ether2-pub                0
                                           ether1-gateway    
                                           ether3-private    
 4 ADC  192.168.100.0/24   192.168.100.0   ether3-private            0
IP Pools:
 # NAME             RANGES                         
 0 public-pool     10.10.10.201-10.10.10.220  
 1 private-pool    192.168.100.2-192.168.100.254
DHCP configs:
 #   NAME               INTERFACE              RELAY           ADDRESS-POOL              LEASE-TIME ADD-ARP
 0   public-dhcp        ether2-pub                             public-pool               3d        
 1   private-dhcp       ether3-private                         private-pool              3d
Thanks!
