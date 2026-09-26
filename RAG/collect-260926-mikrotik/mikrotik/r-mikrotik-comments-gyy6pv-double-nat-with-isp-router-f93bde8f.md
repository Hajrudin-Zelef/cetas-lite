---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-gyy6pv-double-nat-with-isp-router-f93bde8f
title: "r-mikrotik-comments-gyy6pv-double-nat-with-isp-router-f93bde8f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-gyy6pv-double-nat-with-isp-router-f93bde8f.md
source_anchor: ""
source_lines: [1, 53]
sha256: 479522a1536b356620cef2dcff771639a3d4aeef105562151063d6b45bc16876
---

# r-mikrotik-comments-gyy6pv-double-nat-with-isp-router-f93bde8f

Double NAT with ISP Router 
        
        
        
    
    
    Hello there!
So I just recently switched to a fiber connection at home and it comes with TV and VoIP so I can’t put the ISP’s router in bridge mode or I’ll lose access to the other services.
I have a few servers running at home which are connected to a Mikrotik router currently acting as a dumb switch, however I would like to use it as a router/firewall because I do not trust the ISP’s firewall and they don’t offer a whole lot in terms of customization.
My network would then look something like this:
      ISP Router (10.0.0.1/24) serving as WiFi AP
      LAN1 (10.0.0.2): Mikrotik (172.16.0.1/24) where servers are connected
However devices on 10.0.0.0/24 are unable to access devices on 172.16.0.0/24 unless I configure DST-NAT, which is expected since technically 10.0.0.0/24 is WAN from the Mikrotik’s POV.
The problem with that is that then all the servers share the same IP address (10.0.0.2) and I would like to instead have both networks talk to each other while keeping the Mikrotik’s firewall rules for the actual WAN.
I understand I need to configure routing and firewall rules between these two routers but I can’t find documentation on how to actually do it so I would really appreciate any help from you guys!
From the ISP router I am able to add an IP Alias, which as I understand would allow the devices on 10.0.0.0/24 to talk directly to 172.16.0.0/24, or am I misunderstanding it?
Section des commentaires
In ISP router add route to 172.16.0.0/24 via 10.0.0.2
Mikrotik already defaults all route to 10.0.0.1 so you dont need to add anything.
Make sure Forward is allow and you have no masquerade rules in the Mikrotik for this interface, that way the packets are not gonna do NAT, just routing, the 10.0.0.0/24 and the 172.16.0.0/24 networks would see each other
Thank you! In that scenario, would I still be able to use firewall rules on the Mikrotik or would it be basically a DMZ?
Ok cool this worked but I need to disable the last firewall rule (drop all from WAN not DSTNATed) in order to get 10.0.0.0/24 to access servers on 172.16.0.0/24, I tried a few different rules to set before that but none of them worked. Is there a way to add an exception to this rule for a specific address list?
Add an accept Forward rule before this drop rule, that allows forward from src 10.0.0.0/24 dst 172.16.0.0/24 then the forward rule will hit before the drop
You would be essentially CG natting yourself.
Yes. No kink shaming please lmao
Writing from memory, but this (in theory) should work:
Configure LAN connection between ISP and MT using
/30subnet (e.g set ISP LAN interface to10.0.0.1/30and MT iface to10.0.0.0.2/30),
On ISP router, add static route to LAN via MT interface (i.e.
172.16.0.0/24 gw 10.0.0.2),
On MT, add static default route to ISP router (i.e.
0.0.0.0 gw 10.0.0.1),
(potentially, not sure w/o testing), on MT WAN interface - you may need to configure arp mode=
proxy-arp,
Move Wireless network to MT.
Will give that a try! However I need to leave WiFi on the ISP router since the only other AP I have laying around is quite old and doesn’t cover the hole house. Would that be a problem?
You can replace the
/30step with/24. The main purpose was to create point-to-point network, but on second thought - depending on how TV/VoIP are provided this could be not be feasible (and if you want to keep WiFi on, then you'll need a wider subnet to allocate wireless clients).
It would also be worth to sniff around that ISP traffic to see if you could "unpack" VoIP/TV traffic directly on MT and therefore have MT acting as public IP gateway.
What ISP you are on?
It would be a hell of a lot less complicated if you could get another routed IP to assign to the MT and also do the Wifi through the MT if feasible. IF you're worried about the security on the ISP firewall, when you allow everything on 10. to freely talk to everything on 172. then technically the MT isn't protecting 172. because you are allowing everything between the two networks.
Residential link won’t allow for more than one dynamic IP so I’m out of luck there...
I don’t want to allow everything between the two subnets, I just want them to be able to see each other as their own individual addresses but with firewall rules in between.
Say I have a webserver on 172.16.0.5, I would then only allow port 80 and 443 to it but anything on 10.x/24 should be able to connect directly to 17216.0.5:80/443 instead of 10.0.0.2:80/443 (Mikrotik’s “WAN” IP with port forward rules). Does that make sense? Lol it’s a bit hard to explain
I have a similar mikrotik (Lab) setup where I’m using rfc1918 Subnets for both Wan and Lan and all I needed to add was a srcnat masquerade rule under IP/Firewall/Nat. Use the Source nat chain with the Lan interface as the outgoing interface and then choose masquerade as the action and you’re ready to go.
I don’t think I follow. Would that be two Mikrotik routers connected to each other?
Single mikrotik that I’m using for labs. It sits behind a Fortinet that I’m using for internet. I used to have the traffic routed between the MT and Fortinet but I ended up just Natting because I wanted to add a few more vlans/subnets on the MT and just couldn’t be bothered adding static routes on the Fortinet every time. I also didn’t want to turn on dynamic routing for this part of my homelab😀.
Long story short, here is the firewall nat snippet:
1 chain=srcnat action=masquerade out-interface=LAN log=no log-prefix=""
However I do agree with the other commenters that you should look to move the wifi network to the MT for better firewalling capabilities.
My provider has something similar with voip and tv, i dont use those services but im 99.9% confident that they use vlans for tv/voip separation. Maybeee, you can ask them for vlan id, and turn isp router in bridge mode and configure vlans on mt for voip/tv.
They do and I know it’s possible and it would work for the living room TV since that STB is connected over ethernet to the router, but the other STBs are all coax (HPNA) so I don’t really have a choice.
Yes, you can control whatever comes through it with the mikrotik firewall
