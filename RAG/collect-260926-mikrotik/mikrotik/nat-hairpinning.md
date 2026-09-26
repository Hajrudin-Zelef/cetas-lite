---
id: collect-260926-mikrotik/mikrotik/nat-hairpinning
title: "nat-hairpinning"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/nat-hairpinning.md
source_anchor: ""
source_lines: [1, 90]
sha256: 4c7a1503f0db97fe77134e269de534516907d64733aeb7b1a9a6541bb3699a7b
---

# nat-hairpinning

Well, schucks, I seem to be stupid here:

I need to make NAT hairpinning work and I’m no networking expert. All my past routers just had a checkbox, here I have guides, forum posts, even stack exchange questions. None of that helped me achieve the goal.

For example, the very command suggested in the official guide:

```
/ip firewall nat
add action=masquerade chain=srcnat comment="NAT hairpin" dst-address=192.168.254.20 out-interface-list=LAN protocol=tcp src-address=192.168.254.0/24
```

note how I replaced “out-interface=LAN” with “out-interface-list=LAN” because “out-interface” does not accept “LAN” as a value. I also tried to set “lan-interface=bridge” without success.

I also replaced the IPs with my actual ones. 192.168.254.20 is the server in question. It works when I connect to it from outside my home (e.g. 4G connection)

How exactly does one go about setting up hairpin NAT?

             
            
           
          
            
            
              Yes loopback checkboxes were convenient but one didnt learn squat about traffic flow/networking.

All you need is here. –  https://forum.mikrotik.com/viewtopic.php?t=179343

             
            
           
          
            
            
              I tried the first approach:

```
ip firewall nat add chain=srcnat action=masquerade dst-address=192.168.254.0/24 src-address=192.168.254.0/24
```

It doesn’t have any effect. Even if I add logging to the rule, no logs are generated. I moved it to the top of NAT rules, moved it to the bottom, no effect.

What’s even worse: this router has nigh same configuration than the office router and in office, hairpinning works without any intervention (there are no such rules defined), while at home I just can’t get it to work.

             
            
           
          
            
            
              All the information you require is there, if  you think you only need to do one thing, you clearly didnt read or understand.

             
            
           
          
            
            
              
I think You need to watch this: https://www.youtube.com/watch?v=_kw_bQyX-3U

It is perfect if You want manual port forwardings, and not using dynamic rules, like UPnP.

             
            
           
          
            
            
              Here is a good video from Mikrotik on this topic - https://www.youtube.com/watch?v=1I5FywY6opQ

             
            
           
          
            
            
              
Ah. So if I want this to work, actually, I still need to implement the (A) point (I have a fixed static IP). Sorry, I took the “All solutions work” sentence too eagerly and didn’t understand that solution 1 wasn’t just that one entry.

I will try this, just for the heck of it, but ultimately I think I’m just going to create the DNS entry for the server in question. Looks easier and more straight-forward. It’s only one server.

And I also think I know why everything works in the office: we have a telecom router in between and it is likely the one thing doing the hairpinning…

             
            
           
          
            
            
              Changed the wording slightly to avoid this for others,  thanks for the heads up!
