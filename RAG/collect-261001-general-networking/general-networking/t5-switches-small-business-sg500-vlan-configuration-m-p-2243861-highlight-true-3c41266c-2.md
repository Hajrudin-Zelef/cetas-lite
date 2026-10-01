---
id: collect-261001-general-networking/general-networking/t5-switches-small-business-sg500-vlan-configuration-m-p-2243861-highlight-true-3c41266c-2
title: "t5-switches-small-business-sg500-vlan-configuration-m-p-2243861-highlight-true-3c41266c"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/t5-switches-small-business-sg500-vlan-configuration-m-p-2243861-highlight-true-3c41266c.md
source_anchor: ""
source_lines: [15, 202]
sha256: e65998271ac2c34023397d014da6a077e603701763189280657553c182cbe58a
---

# t5-switches-small-business-sg500-vlan-configuration-m-p-2243861-highlight-true-3c41266c

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-09-2013 08:09 PM
Greetings.
I just recently purchased the SG500-52 Switch and is currently trying to configure the device for inter-VLAN connection and internet access.
I have set the device to L3 Mode. The gateway config IP is 192.168.9.254. The default VLAN ID is 1. The PC used for configuring the switch is plugged into the GE2 LAN port, and the internet is plugged into the GE48 LAN port.
I have 3 VLANs set up. VLAN 20 for Sales, VLAN 60 for Accounting, and VLAN 90 for IT. Port membership are as follows. All untagged
- GE1,GE2,GE25,GE26 are VLAN 20
- GE6 is VLAN 60
- GE9 is VLAN 90
- GE48 has membership of all VLANs. (1UP, 20T, 60T, 90T)
All ports are set to trunk mode, except for GE48, which has been set to General. IPs are manually configured with DHCP turned off.
From this PC (192.168.20.1), I can Ping and detect the computers within the same VLAN (VLAN 20) but computers in the different VLAN is completely inaccessible. Furthermore, I cannot access the internet.
Please help. Any suggestions would be appreciated. If you need more info, please do ask.
P.S. Sorry for my English.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-10-2013 10:05 PM
That is correct. All computers in different VLANs can ping each other. Except for the router located in VLAN 1, which can be pinged only if the computer is in VLAN 1.
Put simply, no computers can access the internet unless assigned to VLAN 1.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-10-2013 10:50 PM
Theodore,
What you are missing is a static route in the RV042 back to each VLAN on the switch. The switch VLAN interfaces know to send WAN traffic to 192.168.9.254, and the VLAN 1 interface knows that the gateway is 192.168.9.29. The problem is that 192.168.9.29 has no idea how to get back to 192.168.20.0/24, 192.168.60.0/24 and 192.168.90.0/24 so it just drops the traffic.
RV042 Config: Setup-> Advanced Routing-> Static Routing
Destination IP: 192.168.20.0
Subnet Mask: 255.255.255.0
Default Gateway: 192.168.9.254
Hop Count: 1
Interface: LAN
Destination IP: 192.168.60.0
Subnet Mask: 255.255.255.0
Default Gateway: 192.168.9.254
Hop Count: 1
Interface: LAN
Destination IP: 192.168.90.0
Subnet Mask: 255.255.255.0
Default Gateway: 192.168.9.254
Hop Count: 1
Interface: LAN
- Marty
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-10-2013 11:45 PM
Thanks a lot, Marty, for such a detailed procedure.
I've managed to set up the parameters in the RV042 router and all computers in all VLANs can now detect/ping the router. Despite this, however, internet access is still not possible. I've tried disabling the router firewall, but it still does not work.
P.S. - DNS appears to be operational though.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-11-2013 12:12 AM
Theodore,
Glad I could help you make some progress.
You should never need to disable the firewall in the router. This sounds like a problem with the default gateway of the PC.
Can the PC ping 8.8.8.8? That would indicate internet access but a DNS issue.
What do you mean "DNS appears to be operational though."? Can you resolve google.com from the PC?
What is the Default Gateway of the PC? It should be the VLAN interface of the switch for the VLAN that it is connected to. (i.e. 192.168.20.254)
Try also:
192.168.9.254
and
192.168.9.29
Have you created any Access Rules in the router? (Delete them if you did)
- Marty
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-11-2013 12:32 AM
First off, I have re-enabled the firewall router.
Regarding "DNS operational", I managed to ping Google. from the client PC (using ping www.google.com, and the URL is translated into 74.125.135.100), however, it still shows as "Request Timed Out". Pinging 8.8.8.8 does not work as well.
The configuration of the client is as follows.
IP address: 192.168.20.1 (VLAN 20)
Subnet: 255.255.255.0
Default Gateway: 192.168.20.254
DNS1: 192.168.9.29
DNS2: 192.168.9.254
I tried changing the gateway, to no avail. Furthermore, there are also no access rules defined in the router either.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-11-2013 01:43 AM
Theodore,
Try adding this to the switch config:
ip route 0.0.0.0 0.0.0.0 192.168.9.29
- Marty
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-11-2013 01:50 AM
Added accordingly. However, still unable to access the internet.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-11-2013 05:11 AM
Hi, please look at this post, as long as no one is stuck on the fact it is a sx300 (there's almost no difference if any difference).
https://supportforums.cisco.com/thread/2123434
-Tom 
Please mark answered for helpful posts
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-13-2013 02:03 AM
I did it according to the instructions, but internet connectivity is still down.
One notice though, my network does not use DHCP configuration. IPs are statically assigned to ensure compatibility with existing equipments.
From the client's PC in VLAN 20, I can now...
- Ping every computer in the same and different VLANs (VLAN 1,20,60,90)
- Access the SG500 switch from any VLANs (1,20,60,90)
- Pinged the router successfully at address 192.168.9.29 (at VLAN 1) from any other VLANs.
Except for the internet connection, which appears to be unusable. Pinging 8.8.8.8 and 4.2.2.2 returned timeout.
Please help. I think it's pretty close now but there's still something either missing or misconfigured.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-13-2013 08:14 AM
Hi Theordore, please allow me access to your network. I do not believe anyone can help you at this point as all the information required has been provided. I have tried and tested this for years at this point.
If you could please coordinate a teamviewer8 session to my email at tmw0402@hotmail.com that would be good.
-Tom 
Please mark answered for helpful posts
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-16-2013 03:05 PM
Seems like the answers you have been given should work, but are you sure you have the IPs configured correctly. I think the issue is in your IPs for your switch and your internet router/fw. In your first post you say the "gateway IP is 192.168.9.254). To me it seems like you are saying your internet router's internal IP is 192.168.9.254. So your VLAN switch is configured to 192.168.9.253 on VLAN 1? Is this correct? If so, I think the information you were given is correct above, just some IPs are off.
You want all ports (INCLUDING THE INTERNET ROUTER) to be access ports, NOT TRUNK OR GENERAL. So you would have:
interface vlan 20
description GE1,GE2,GE25,GE26
ip address 192.168.20.254 255.255.255.0
interface vlan 60
description GE6
ip address 192.168.60.254 255.255.255.0
interface vlan 90
descritpion GE9
ip address 192.168.90.254 255.255.255.0
interface vlan 1
description GE48 (internet router)
ip address 192.168.9.253 255.255.255.0  <<< THIS IS THE IP OF THE VLAN SWITCH
