---
id: collect-260926-mikrotik/mikrotik/questions-815433-limit-bandwidth-on-a-mikrotik-router-by-port-number-2ac8349a
title: "questions-815433-limit-bandwidth-on-a-mikrotik-router-by-port-number-2ac8349a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-815433-limit-bandwidth-on-a-mikrotik-router-by-port-number-2ac8349a.md
source_anchor: ""
source_lines: [1, 25]
sha256: 61174321729dff2e2fc71176abc3ffd5fbee106921c3f87e67f3d88aa585ee43
---

# questions-815433-limit-bandwidth-on-a-mikrotik-router-by-port-number-2ac8349a

I have this one Mikrotik router (RB951G-2HnD) that serves as a main router with port forwarding for HTTP/HTTPS and Plex to a server behind NAT. Now since there isn't much upload bandwidth (15Mbps), I'd love to limit the bandwith for HTTP/HTTPS on the WAN interface so that there's some bandwidth left for Plex. I've been looking in /queue but haven't seen anything that limits by port number.
                    
                        Add a comment
                    
                 | 
            
                
            
        
         
    2 Answers 2
This is described here
- http://wiki.mikrotik.com/wiki/Traffic_Priortization,_RouterOS_QoS_Implemetation
or here
- http://forum.mikrotik.com/viewtopic.php?t=73214
First you need to mark your packets. And to mark you can use port number
/ip firewall mangle
add action=mark-packet chain=forward dst-port=80 new-packet-mark=http passthrough=no protocol=tcp
add action=mark-packet chain=forward dst-port=443 new-packet-mark=https passthrough=no protocol=tcp
And then you can do traffic shaping on those marked packets
Below is also a very good and simple example
- 
        I added a queue and apply it to the WAN interface: /queue simple add name="wan-https-limit" target=pppoe-out1 max-limit=10M/10MManuth Chek– Manuth Chek2016-11-17 12:57:11 +00:00Commented Nov 17, 2016 at 12:57
You need to create a mangle rule in /ip firewall that will match outgoing packets to ports 80 and 443 and mark them (action=mark-packet).
Then you can create a queue and use the packet mark you defined on the mangle rule to match and limit those packets.
