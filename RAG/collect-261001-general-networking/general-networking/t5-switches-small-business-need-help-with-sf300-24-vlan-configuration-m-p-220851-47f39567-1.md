---
id: collect-261001-general-networking/general-networking/t5-switches-small-business-need-help-with-sf300-24-vlan-configuration-m-p-220851-47f39567-1
title: "t5-switches-small-business-need-help-with-sf300-24-vlan-configuration-m-p-220851-47f39567"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switches-small-business-need-help-with-sf300-24-vlan-configuration-m-p-220851-47f39567.md
source_anchor: ""
source_lines: [1, 163]
sha256: 4bad4a6b2f41d524254e1abdfc9c9a8c874c48f9d6557eee4e96cacb5404d643
---

# t5-switches-small-business-need-help-with-sf300-24-vlan-configuration-m-p-220851-47f39567

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-13-2013 11:18 AM
Hi,
Let me preface this with the fact that networking is most definitely not my strong point, so any help here is greatly appreciated.
I am trying to segment out some virtual desktop infrastructure onto its own VLAN using a Cisco SF300-24 Layer 3 switch. I can get the switch to connect to the network with VLAN 1 assigned an IP on the network's subnet (192.168.16.X), but I can't get anything set up on VLAN 20 (192.168.20.X subnet) to connect past VLAN 20's gateway IP (192.168.20.254). The ports assigned to VLAN 20 are set to Access mode if that matters.
Here is a diagram to illustrate what it all looks like, since there is another switch (L2) involved.
So I'm not really sure what I'm missing here since all the settings seem pretty straight forward.
Solved! Go to Solution.
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-18-2013 08:25 AM
Hi Simon, I'd recommend you remove any active directory server and basically remove all security factors. This will give the idea of where to begin.
If you take a very fundamental set up, 2 window 7 desktops without a window firewall enabled, they will both work as expected.
You have to remember that in the firewalls, even if the devices are able to respond to ICMP if the request is coming from a different subnet they will not as it is recognized as a foreign network. You have to make this network known to these computers or make it so the computer doesn't care.
You may be able to accomplish this by simply adding additional subnets on the advance configuration of the network card (if this doesn't take up too much address space) as an example.
Or, as you have discovered you can add routes which is a bit cumbersome and inconvienent but effective.
-Tom 
Please mark answered for helpful posts
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-13-2013 01:37 PM
Hi,
Not familiar with SF300 or Dell switch, but computer in vlan 20 has the correct default gateway of 192.168.20.254?
If yes, computer in vlan 20 can ping 192.168.16.6?
Is the link between the SF300 and Dell switch a trunk link or an access link. If it is an acess link, you need to change it to trunk link using dot.1q to carry both vlan 1 and 20.
HTH
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-13-2013 03:25 PM
Sorry, looks like I left some info out.
I will have to double check the computer that is connected to the switch when I get to the office, but it is set to automatically get the DHCP settings, which it correctly does and is able to ping the switch at 192.168.20.254, so I would think the gateway settings are correct. The problem is that no computer in VLAN 20 can ping anything past the VLAN 20 gateway.
The link between SF300 and Dell switch is set to trunk (default, I didn't make any changes to it)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-14-2013 04:19 AM
Wrong forum post in "small business - switches". You can move your post using the actions panel on the right.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-14-2013 04:48 PM
Hi Simon, what are you trying to access? If your computer is trying to get to that exchange, that exchange needs to have a default gateway set for the 16.1. But it may also need the firewall disabled or a rule enabled to allow communication outside of its subnet.
That dell switch, is it basically acting as an unmanaged switch? Does it have any ingress filters?
-Tom 
Please mark answered for helpful posts
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-14-2013 06:33 PM
At this point I'd be happy with pinging anything past VLAN 20's gateway.
The Dell switch isn't doing anything other than acting as a gateway. In fact, I tried taking that part of the network out of the equation and hooked up the SF300-24 directly to the T1 router and still couldn't reach beyond VLAN 20's gateway.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-15-2013 05:37 PM
Hi Simon, please consider this scenario.
Host A = 192.168.1.10, /24 with gateway 192.168.1.1
Host B = 192.168.2.10, /24 with gateway 192.168.2.1
Vlan 1 = 192.168.1.1
Vlan 2 = 192.168.2.1
Host A connects to port 1 of the switch while Host B connects to port 2 of the switch.
The configuration is as such-
config t
vlan database
vlan 2
interface vlan 1
ip address 192.168.1.1/ /24
interface vlan 2
ip address 192.168.2.1 /24
interface fa01
switchport mode access
switchport access vlan 1
interface fa2
switchport access vlan 2
Connect the respective host to their ports, if you perform a show ip route you should see the routes dynamically built. Assuming the devices have the capability to respond to ICMP you should be able to ping to the switch SVI of each subnet the the respective host.
I would highly recommend to disable any firewall or security settings of the host stations. These features break things like ICMP, netbios, etc.
A simple point to point set up like this works every time and it is the foundation to make your problems go away.
-Tom 
Please mark answered for helpful posts
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-16-2013 06:05 AM
Hi Tom,
Thanks for the reply. From what I can see, I already have things configured like you described. The only thing that I just tried was deleting the static routes to see if they'd be dynamically created. They were successfully created dynamically as you outlined.
The computer connected to VLAN 20 can successfully ping VLAN 1's gateway (192.168.16.6) and VLAN 20's gateway (192.168.16.254), but cannot ping any part of the network past that. Anything on VLAN 1 (192.168.16.X) can ping anything on that subnet, but cannot ping any part of VLAN 20.
Yesterday a Cisco support rep was remoted into my system and spent several hours trying to get this to work without success, so I don't feel too dumb about not getting this to work. He said they would try to replicate the setup in their lab and get back to me today, so we'll see.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-17-2013 08:13 AM
Hi Simon if your exact set up is as I describe, it is an external factor (such as wireless connection) or the devices/computers you're using. This has worked since the inception of the switches. From the best I can recall, since public release this is one functionality that has not had a bug.
-Tom 
Please mark answered for helpful posts
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-18-2013 06:44 AM
Tom,
I suspect you are correct.
I configured Routing and Remote access on a Windows server with a static route 192.168.20.0/24 --> 192.168.16.6 (VLAN 1 IP)
From that server (only) I can ping VLAN 20 computers and VLAN 20 computers can ping that server (only).
Not sure where to go from here though.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-18-2013 08:25 AM
Hi Simon, I'd recommend you remove any active directory server and basically remove all security factors. This will give the idea of where to begin.
If you take a very fundamental set up, 2 window 7 desktops without a window firewall enabled, they will both work as expected.
