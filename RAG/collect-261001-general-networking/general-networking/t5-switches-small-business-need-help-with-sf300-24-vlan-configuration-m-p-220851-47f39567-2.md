---
id: collect-261001-general-networking/general-networking/t5-switches-small-business-need-help-with-sf300-24-vlan-configuration-m-p-220851-47f39567-2
title: "t5-switches-small-business-need-help-with-sf300-24-vlan-configuration-m-p-220851-47f39567"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switches-small-business-need-help-with-sf300-24-vlan-configuration-m-p-220851-47f39567.md
source_anchor: ""
source_lines: [164, 223]
sha256: 322b6af4428bdc62f939ea20717641c7bf4f9db3ec2682f78bc2afa0f1542d56
---

# t5-switches-small-business-need-help-with-sf300-24-vlan-configuration-m-p-220851-47f39567

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
07-19-2013 06:12 AM
Thanks Tom,
You put me on the right track, adding a default gateway to the adapter let me start to route data across the network/VLANs.
Thanks for your help.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-24-2013 11:52 AM
Ok, I've almost got everything working correctly now, I just can't access the internet from anything other than VLAN 1's subnet.
Here is my setup that works as far as VLAN routing:
T1 Router (192.168.16.1) --> Cisco SF300 VLAN 1 (192.168.16.254) --> VLAN 20/30/40/50/60
I can ping all over the network between VLANs, DNS works as it should.
What I cannot do is ping any internet addresses from anything other than VLAN 1 (192.168.16.X). This happens whether I specify an IP address or name, such as www.google.com. DNS works, since it pulls a valid IP, but it times out for pinging. The SF300 has no issue pinging internet addresses. A tracert from a computer on VLAN 50 shows the first hop to the gateway IP (192.168.50.254), then times out.
I have one static route defined (0.0.0.0 0.0.0.0 192.168.16.1) and the others are discovered dynamically, as they should.
The ports on VLAN 1 are set to trunk/untagged. Ports on the other VLANs are set to access/untagged.
Any ideas?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-25-2013 07:30 AM
Hi Simon, does your router support the vlans? Otherwise the router will require a static route pointing back to the SVI of the switch.
Here is an example post-
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
07-25-2013 10:20 AM
Tom Watts wrote:
Hi Simon, does your router support the vlans? Otherwise the router will require a static route pointing back to the SVI of the switch.
Here is an example post-
https://supportforums.cisco.com/thread/2123434
-Tom 
Please mark answered for helpful posts
Tom, once again, I think you are correct. The router is a Cisco 2600 series (not sure of the exact model) and is the only place I can think of where the problem could occur. Unfortunately I'll have to wait until tomorrow to verify since that's when the console cable I need is being delivered
