---
id: collect-261001-cisco/cisco/r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a-1
title: "r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a.md
source_anchor: ""
source_lines: [1, 17]
sha256: 853694d3b4ade28d447b2fb5ac39c02d7cf2ef8b6021f429bc01fad7d2a7dc5e
---

# r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a

Spines not passing traffic: VxLAN EVPN Cisco NX-OS 
        
        
        
    
    
    
      This is driving my a little nuts. I have a spine/leaf configuration using Nexus C93180YC-FX3 in all roles. I'm trying to test a failure scenario (to be fair I haven't tested the steady state yet) and my spines appear to be eating traffic. My NVE interfaces transmit but never receive and ICMP traffic ingresses to the spines but never egresses anywhere. As such, host to host communication is not working.I have 5 vPC switch pairs (10 switches). Member 1 is connected to spine 1 and member 2 is connected to spine 2 for all of them. I'm testing this weird failure state. However, again, I have not actually tested the steady state. I only have two hosts connected. One on switch pair 507/508 and one on switch pair 503/504I'm using OSPF as my underlay and I can see and route to all loopbacks. Leaf switches are using Loopback 1 for NVE source interface and have the anycast IP secondary configured.
    
I added this monitor to spine 1 (coresw501)
monitor session 1  
  source interface Ethernet1/2 both   
  source interface Ethernet1/4 both  
  destination interface sup-eth0  
  no shut  
      Then I ran ethanalyzer and when ICMP goes from host 10.55.4.9 -> 10.55.2.9
    
