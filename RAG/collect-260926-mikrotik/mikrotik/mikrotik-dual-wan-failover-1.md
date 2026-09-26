---
id: collect-260926-mikrotik/mikrotik/mikrotik-dual-wan-failover-1
title: "mikrotik-dual-wan-failover"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mikrotik-dual-wan-failover.md
source_anchor: ""
source_lines: [1, 106]
sha256: c9b5a63988dba3494cb6a0c997cc02e52c53c3e8a9f8aba051796e7990db428d
---

# mikrotik-dual-wan-failover

Hello guys,

I’ve recently bough a mikrotik router and the model I chose as let’s say my teaching router thanks to the help of some colleagues from the forum is hAP ac2.  If someone is interested of something about the need of the router - here is the thread I made http://forum.mikrotik.com/t/mikrotik-routers-question/122624/1

However, one of the critical things I needed was the dual WAN support. I need it only as a fail-over without load balance. This weekend I had few hours to play with the new router and  to try to make a simple setup.

Here is the current situation about the ISPs:

1. Main link - 100Mbps PPPoE  directly to the hAP
2. Back up link - 50 Mbps ADSL  - phone line->Modem->hAP
3. Static local IPs from both providers.

Here is what I’ve done so far:

1. 
I stepped on the default configuration of the hAP and from there I’ve tried to build up what I need. Firstly I’ve changed rotuer’s IP  192.168.0.1 and created a new DHCP server wit primary DNS 192.168.0.1 and secondary 8.8.8.8, added pool in the needed range.
2. 
Excluded the ether2 port from the bridge and added it to the list of WANs so other rules can apply to it.
3. 
I’ve created a ppoe client for port1 with the needed user and password while using peer DNS and default route / ppoe-out  was also added to the WAN list.
4. 
Static IP for the port2 - 192.168.x.x, added route with gateway (the ADSL modem)
5. 
Static DNS - same as the adsl modem
6. 
Distance of the main link is 1 and distance of the backup - 2

So in this configuration everything seems to work for now even if I don’t know if there is something missed and the test was done only by disconnecting the WAN ports. However this configurations is apart from the PCC tutorial and I’m not sure if the PCC could be applied to this case when there is one static and one pppoe? Could I use as gateway the whole pppoe-out in the PCC wiki scenario? I’ve read several posts and wikis about the dual WAN scenario but it seems that most of them are using load balance and they are mostly for static addresses. There are also the mangle rules which I haven’t had the time to study more carefully. I saw that the preferred way of dual WAN fail-over is the PCC but what about the mangle rules, they seems to use it too?  What would be the best way to configure dual WAN fail-over in my case and is my configuration by far worth a dime?

I’m using the default settings for the firewall applied to both WAN ports via the WAN list, same for NAT.

Thanks for the help in advance guys. It’s a great device with plenty of settings and I like it. Still it would be great if they’ve added something as the quick config for dual WAN options as it’s a common thing now days. This router would be used for experiments for now until I start to feel a bit more comfortable with the OS and I get my knowledge together.

             
                
            
           
          
            
            
              PCC is for load balancing, from your description, you do not need that.

Then I would also change the ADSL Modem to bridge mode and configure ADSL PPPoE on the Mikrotik.

The do not use the "Add default Gateway"in the PPPoE settings, instead create static default routes with a distance of 1 and 2, 2 for the adsl and use "check gateway " on the static routes

             
            
           
          
            
            
              
As @CZFan has said already: if so, don’t bother about PCC, as PCC is here first for load distribution, and only as a side effect it provides some kind of failover. Leaving out PCC will relieve you from having to understand the mangle rules for the moment.




Which is also the weakest point of that configuration. Even though the handover interface of WAN1 is Ethernet, there may still be something between your Ethernet port and the actual Internet which may fail without your Ethernet interface going down, and in such case the route via that interface will stay up so no failover will happen. So here I recommend this article explaining how to monitor that accesss to internet is really possible via each WAN. The Mikrotik wiki describes the same coniguration but in a much less explanatory way.




I agree it would be great - but a great waste of developers’ efforts that could be better spent on features which cannot be achieved by configuration. Every user has a different environment, so WAN1 may be anything out of (PPPoE, static IP configuration, DHCP) on anything out of (ethernet, wireless), leaving aside LTE with its two modes (serial or Ethernet emulation) and so can be the WAN2, and every user has different requirements, e.g. a mere failover in your case and load distribution in someone else’s case. Others may want one of those basic approaches for most of the traffic but some services to be accessed solely via one of the WANs. So I personally like the current approach where QuickSet is for people who have bought Mikrotik by chance and the real configuration interface is for those who have chosen it for its flexibility. Flexibility means a lot of things can be configured, and without an understanding what each setting is necessary for it is close to impossible to answer properly all what a configuration wizard would have to ask.

             
            
           
          
            
            
              
I though about something like this but if I set the ADSL modem to bridge I’d need the password so I’ll be able to create the pppoe from the mikrotik. Sadly I don’t have this information and the ISP won’t give it to me if requested. Normally I won’t even have access to their device and would be forced to manage it by their limited web but however I have access to the modem. I’d probably be able even to recover the password for the pppoe but it may be too much.  Other problem that will result directly from that is the DVR which needs some ports forwarding so I’d have to configure it too.

According to the Add default gateway - I’m a bit confused as the local address is static but in the default settings it takes the remote address. I’ll have to try this otherwise the pppoe route is still set to 1 and the adsl to 0 - that’s for the 0.0.0.0/0 dest.

Thanks for the link, I’m going to check it for sure and I’ll try to make this work. According to the quick settings - it is true that it’s a bit of a hard work to implement it. Still even if it’s a bit tricky to set things up and you’ll need a lot of reading and network knowledge I really like this product. You can learn a lot from it.

At the bottom line it turns out that I can use the distance difference and monitoring to realize the setup. It’s great as I was just preparing to read about the PCC and mangle rules for the next time I have free time.

             
            
           
          
            
            
              Just a note when keeping ADSL modem in router mode, you must not use nat / masquerade between else you will have a double NAT situation than can cause issues

             
            
           
          
            
            
              @CZFan, sorry… although it is true that multiple NAT does cause issues in rare cases (in 99,9% situations it is just as bad as a single NAT), you cannot just disable src-nat (masquerade) between Mikrotik and the ADSL modem, but you also have to add route(s) to Mikrotik’s LAN subnet(s) to the ADSL modem. Otherwise the modem would send packets for these subnets back up the WAN (which is its default gateway).

             
            
           
          
            
            
              Hello guys,

I haven’t had much time recently to play with the fail-over but today I had some time and I decided to test the fail-over scenario from the article **sindy** posted here. I think that I’m facing a problem and I’m not exactly sure where it comes from.

