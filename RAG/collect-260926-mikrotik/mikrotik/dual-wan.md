---
id: collect-260926-mikrotik/mikrotik/dual-wan
title: "dual-wan"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/dual-wan.md
source_anchor: ""
source_lines: [1, 76]
sha256: 89b67ea39a1cbbe3bf321daa6ad9b0fa11ffb224970da7e04de8a4cf214b069e
---

# dual-wan

Hello, I need some help to configure a dual WAN failover which switches to second WAN when ISP1 GW is up, but there is no internet. I checked a lot of topics and different configs from the forum, but they all work the same way

For example:


Here we have ISP1 and ISP2 connecting to **ether1** and **ether2** of the MikroTik with static IPs.

All the setups I’ve tried only switches to ISP2 when I disconnect the cable(blue arrow) from **ether1**, because in this case there is no ping to ISP1 GW and Google DNS[8.8.8.8], but in case were I plug out the cable(green arrow) from ISP1 router or there is just a problem with the provider there is still ping to ISP1 GW, but there is no ping to Google DNS[8.8.8.8] respectively there is no internet to the router and its clients, but it does not switch to ISP2. Is there a way to configure it to watch DNS not the GW?

             
            
           
          
            
            
              First of all your IP addresses on two different WAN connections are the same in the diagram,  WHY?

Search for recursive routing…

The basic setup for failover is as follows… an correctly that  doesnt help if one can connect to the ISP but the ISP to the WWW is broken.

Note by separating the routes by distance, the router will always choose hte Primary ISP, when available.

*/ip route
add check-gateway=ping comment=Primary ISP distance**=5** dst-address=0.0.0.0/0 gateway=Primary-gatewayIP
add comment=SecondaryISP distance=**10** dst-address=0.0.0.0/0 gateway=Secondary-gatewayIP*

Thus we use recursive routing… and in this case a well known DNS server address to check connectivity to the WWW.   One  has to start being cognizant of scope and target scope when configuring the recursive routes!

*/ip route
add check-gateway=ping distance=5 dst-address=0.0.0.0/0 gateway=1.0.0.1 scope=10 target-scope=12
add distance=5 dst-address=1.0.0.1/32 gateway=PrimaryISP-gatewayIP scope=11 target-scope=11
add comment=SecondaryISP distance=10 dst-address=0.0.0.0/0 gateway=SecondaryISP-gatewayIP scope=10 target-scope=30*

FY Reading - **PARA I.**   - https://forum.mikrotik.com/viewtopic.php?t=182373

             
            
           
          
            
            
              I could have sworn they were both 10.10.10.2  when I first looked,  must be tired.

             
            
           
          
            
            
              Sorry for the mild necro of this post…

I’m dealing with the same setup at a client right now and absolutely recursive routing is the ticket!

Also very important to mention that if your still learning to set up a mikrotik router (As I am) then make sure you have strong firewall rules set up as well, and test them!

I’m in love with the absolute control a MikroTik device gives you but it’s a double edged sword, it lets you do almost anything… especially make mistakes!

             
            
           
          
            
            
              The basic default setup of the MT is very good but once you start creating networks and groups of users some slight modifications are required.

There is no need to over do it on the firewall side.

             
            
           
          
            
            
              
Hello, sorry for late reply.  With only those 3 settings in ip route it works fine. Thanks.
