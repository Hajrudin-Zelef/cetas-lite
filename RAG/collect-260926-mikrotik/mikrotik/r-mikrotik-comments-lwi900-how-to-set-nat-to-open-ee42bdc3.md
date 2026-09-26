---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-lwi900-how-to-set-nat-to-open-ee42bdc3
title: "r-mikrotik-comments-lwi900-how-to-set-nat-to-open-ee42bdc3"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "latency"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/r-mikrotik-comments-lwi900-how-to-set-nat-to-open-ee42bdc3.md
source_anchor: ""
source_lines: [1, 35]
sha256: 79e0517473f5787255eda3547d52a8fc928745e12ff99ec33db3bbc2f1901ad1
---

# r-mikrotik-comments-lwi900-how-to-set-nat-to-open-ee42bdc3

How to set "NAT" to "Open" 
        
    I was playing Grand Theft Auto Online and I was getting errors that I should set my "NAT" from "Strict" to "Open" for better latency. Internet suggests that I log in to my router to change this. I'm greeted with RouterOS v7.0. I didn't set up the router so I have no idea how to change the "NAT". Can someone explain how to change it like I'm 5?
Edit:
I've tried lots of things and the error still comes up. I've come to the conclusion that GTA Online sucks. Nothing new
Section des commentaires
'open' nat is pretty much just a certain set of public ports dedicated to your devices private ip. it's one step below a DMZ. it's abit different from each game and platform. The way I solved it was to find all the ports i could on forums and stuff about the game/platform in question, and make a dst-nat rule for each (tcp and udp), sending them all to the console.
"Nice" tips from that game devs:
Forward TCP Ports: 80, 443
Disable any firewall or other network filtering
Bypass the router entirely and connect your PC directly to the modem using a wired Ethernet cable (not wi-fi).
This is fking ridiculous. I am speechless. The next should be: put all your money into the box and drag it to the street.
I'd be amazed if a game was listening on port 80+443 - doesn't windows prevent normal applications from listening on ports below 1024?
Your conclusion is probably wrong. Try enable UPNP.
Tried that
Disable all firewall rules. If it works, enable them one by one and keep logging on. I'm pretty sure playstation and such uses nothing but upnp.
RouterOS 7.x is still in Beta (7.0 is old, latest is 7.1b4), it probably isn't a great idea to run this on equipment where you require stability, or aren't familiar with RouterOS.
That aside, the list of ports used by GTA Online is here: https://support.rockstargames.com/articles/206210548/How-to-Resolve-Errors-in-GTA-Online-about-Strict-NAT-Type
Before you go down the path of forwarding a heap of ports, have you looked at enabling UPnP as a potential fix? IP -> UPnP; Enable and Apply, and ensure you've set the Internal and External Interfaces correctly (Interfaces button).
If you portforward to a device, keep in mind that is for that device alone. The rest will get strict NAT.
The painless option is UPnP. This goes with a few Filter rules as well. It will allow at least to variably open the ports that the games need in a given moment.
That is if the software is able to talk to UPnP, and it poses a few threats.
Some worms and cracking tools do talk UPnP and ask the router to forward ports directly to the machine ip and voilá, an external attacker is inside your network.
UPnP is a nice idea, but is considered too risky.
With routerOS there is no easy way to set your "NAT" to open. You're way better off using port forwarding to forward the ports you need to the device and ensuring your firewall rules allow open and related connections.
port forwarding is literary a technology behind oversimplified term "Open NAT"
edit: didn't think it through
Except we are not enabling upnp so it's specific ports only.
Just set the nat for no or any IP and then set the firewall for any IP to that port. MT would happy to nat whatever with your permission
just use upnp for automatic port forwarding if the app supports it. or manually forward ports dst-nat. or dmz(forward all ports) for the whole network or a specific IP.
of course it wont make any difference if youre behind ISP nat, resulting in double nat or even more. in that case, you should tell ISP to assign you a public IP, make it the next hop after your router and DMZ you on their side of network.
First off, I'm 99% sure you're behind a CGNAT.
Second of all, read this: https://serverfault.com/a/1053113/616716
This of all, I'm 100% sure multicast routing/UPnP configuration is broken on your router.
Ask your mother to enable the whorecunt option. That will open it up...
