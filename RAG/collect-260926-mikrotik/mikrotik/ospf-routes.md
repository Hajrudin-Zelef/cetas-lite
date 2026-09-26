---
id: collect-260926-mikrotik/mikrotik/ospf-routes
title: "DST-ADDRESS        GATEWAY             DISTANCE"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/ospf-routes.md
source_anchor: ""
source_lines: [1, 55]
sha256: 5394c89f397c7d14105fe550868edc4ddef62de0fdaeceed1d046506c862f135
---

# DST-ADDRESS        GATEWAY             DISTANCE

Hi all,

I own a little wisp, just upgraded some of my routers from 6.49 to 7.5.

Now in /ip routes I have this:

Is it normal? What means “DH”? I’m missing some configuration?

Thank you!

             
            
           
          
            
            
              Moving up a major version in production is a terrible idea, you need to do this in the lab and research what these things mean.

In winbox you can hover over the flags and it’ll generally tell you what they are in a tooltip. I can’t find anything in the docs about H personally though, but my guess would be installed in the [H]ardware route table, i.e. switch chip with L3 features so it doesn’t run through CPU?

D = Dynamic, i.e. generally learned or installed through a dynamic process like pppoe, ospf, bgp, etc.

On terminal you can also do /ip/route print and it may give you a quick rundown on the flags. (see here)

             
            
           
          
            
            
              I’m not having issues with performance or everything else, all is working.

Just curious to know what is “DH”, if I over it tells me “Dynamic - Hardware offloaded”.

If i run Ip route print from command line, it doesn’t show me the DH routes. Just the OSPF ones:

```
Flags: D - DYNAMIC; A - ACTIVE; c, s, o, y - COPY
Columns: DST-ADDRESS, GATEWAY, DISTANCE
#     DST-ADDRESS        GATEWAY             DISTANCE
0  As 0.0.0.0/0          100.70.1.81                1
  DAo 10.0.0.185/32      100.70.1.81%combo1       110
  DAo 10.0.0.189/32      100.70.1.81%combo1       110
  DAo 10.0.0.193/32      100.70.1.81%combo1       110
  DAo 10.0.0.195/32      100.70.1.81%combo1       110
  DAo 10.0.0.201/32      100.70.1.81%combo1       110
  DAo 10.0.0.202/32      100.70.1.81%combo1       110
  DAo 10.0.0.203/32      100.70.1.81%combo1       110
  DAo 10.0.0.204/32      100.70.1.81%combo1       110
  DAo 10.0.0.205/32      100.70.1.81%combo1       110
  DAo 10.0.0.207/32      100.70.1.81%combo1       110
  DAo 10.0.0.208/32      100.70.1.81%combo1       110
  DAo 10.0.0.209/32      100.70.1.81%combo1       110
  DAo 10.0.0.210/32      100.70.1.81%combo1       110
```
