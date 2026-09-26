---
id: collect-260926-mikrotik/mikrotik/filter-rules
title: "filter-rules"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/filter-rules.md
source_anchor: ""
source_lines: [1, 80]
sha256: 5506e8177b00ea37645ac22e6cf567553865d78029ec69516b1ebb857a685b65
---

# filter-rules

Looking for advice on router firewall rules.

Our fiber comes into sfp-sfpplus1. I have sfp-sfpplus1, ether2 and ether3 in a bridge. Ports ether2 and ether3 are connected to a third party firewall in HA sync, which performs all access rules and NAT policies and security is enabled. The firewall is connected to the core switch.

Fiber > MikroTik CCR (ether2, ether3) > Firewall

At minimal if the firewall performs all rules, NAT, and security- what firewall filter rules should I use on the router?

             
            
           
          
            
            
              The main question is if you use this device as a router or just as a dumb media converter and if it has it’s own IP? Does it need any rulues except guarding access via Winbox ones?

             
            
           
          
            
            
              It’s used as a router and has a static IP assigned. The third party firewall performs all of the rules.

I wasn’t sure if I needed to have any rules set on the MikroTik firewall as well.

             
            
           
          
            
            
              
Only for chain=input to protect MT router itself from possible attacks.

             
            
           
          
            
            
              Currently I have these rules. So I can remove them all?

/ip firewall filter

add action=accept chain=input comment=“allow established/related” connection-state=established,related

add action=drop chain=input comment=“drop invalid” connection-state=invalid

add action=accept chain=input comment=“allow icmp” protocol=icmp

add action=accept chain=input comment=“allow management” src-address=192.168.88.0/24

add action=drop chain=input comment=“drop all”

             
            
           
          
            
            
              These rules control access to router’s management interface. They are fine, they don’t interfere with traffic passing router in any way.

             
            
           
          
            
            
              Do I really need them? What would be best practice for only allowing management access to the router?

             
            
           
          
            
            
              Since CCR is used as router and needs IP addresses set on all relevant interfaces for routing purposes, it also needs firewall rules which limit access to its management. Again: the rules you showed are fine as long as the permitted IP subnet (192.168.88.0/24) is safe. Whether it’s safe depends on the rest of router’s config (all of it) and physical layout. What I meant by my previous post is that rules controlling management access are completely ignored for traffic passing between two router interfaces (which is forwarded traffic). As you have dedicated firewall devices in place, this router doesn’t have to touch forwarded traffic, hence there’s no need for any firewall rule in forward chain .

Out-of-band management is usually preferred as it allows settings which define allowed access paths better. It is impossible to tell if you already have it based only on the firewall rules you posted. But configuring out-of-band management access doesn’t mean that firewall rules are not necessary, they are always necessary.
