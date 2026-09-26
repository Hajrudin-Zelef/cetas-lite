---
id: collect-260926-mikrotik/mikrotik/v7-1-mpls-vpls-question-2
title: "feb/09/2022 14:25:54 by RouterOS 7.2rc3"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/v7-1-mpls-vpls-question.md
source_anchor: ""
source_lines: [277, 366]
sha256: 75aa1b55b2d69435cbfd723b8197abc147d7bf86c0f87a796f921abb7ae5dedd
---

# feb/09/2022 14:25:54 by RouterOS 7.2rc3

If anybody has a worcing config, please share with me, beacause I slowly rip my hair out  (just kidding…I’m bald  )


             
            
           
          
            
            
              Some things to consider.  Max phy MTU on the rb2011 is limited much lower than you appear to have set on the 1100.

RB2011 series 	ether1-ether5:4074; ether6-ether10:2028; sfp1:4074

RB1100AH 	ether1-ether10:9498; ether11:9500, ether12-ether13:9116

https://wiki.mikrotik.com/wiki/Manual:Maximum_Transmission_Unit_on_RouterBoards



I cant remember if this is necessary or not but set your mpls interface input=yes

/mpls interface set 0 input=yes

You also have your mpls MTU set to 1600 and your interface MTU appears to be default 1500.

I have not tested with as new of a version as you are using though, but it *should* still be working.

             
            
           
          
            
            
              
The MTUs are all ok, but that “/mpls interface set 0 input=yes” saved my day. I set it on both routers and it worked like charm. Thank you very much!

             
            
           
          
            
            
              I was missing the input=yes in my config. Adding it improved things - instead of the VPLS tunnel not doing anything and then the devices freezing, I get exactly one ping through and then the devices freeze and have to be power cycled.

             
            
           
          
            
            
              
I have the same problem with 7.1.3 - VPLS is still not working correctly. Actually, sometimes it starts working after I disable/enable the VPSL in config (I can see packets sent/receive), just then after a few seconds my router gets frozen, and the only method to restart it is to power cycle the router.

In my case, I’m trying to establish VPLS between bridges on RB3011 UiAS and hAP AC2. It used to work perfectly with RouterOS v6, just I’d like to migrate to v7 due to WireGuard.

             
            
           
          
            
            
              Unfortunately, the VPLS does not work with an MTU greater than 1500. When I set the PW L2MTU higher, it does nothing. L2MTU is not displayed at all 

             
            
           
          
            
            
              
Yes, I can confirm that issue.

```
v7.1.3 and 7.2rc4:
ping 172.16.70.1 do-not-fragment size=1501
  SEQ HOST                                     SIZE TTL TIME       STATUS                                                
    0 172.16.70.1                                                  timeout                                               
    1 172.16.70.1                                                  timeout                                               
    2 172.16.70.1                                                  timeout                                               
    sent=3 received=0 packet-loss=100% 
v6.48.6:
ping 172.16.70.1 do-not-fragment size=1501
  SEQ HOST                                     SIZE TTL TIME       STATUS                                                
    0 172.16.70.1                              1501  64 1ms594us  
    1 172.16.70.1                              1501  64 1ms601us  
    2 172.16.70.1                              1501  64 1ms634us  
    sent=3 received=3 packet-loss=0% min-rtt=1ms594us avg-rtt=1ms609us max-rtt=1ms634us
```

Mikrotik when it will be fixed?
