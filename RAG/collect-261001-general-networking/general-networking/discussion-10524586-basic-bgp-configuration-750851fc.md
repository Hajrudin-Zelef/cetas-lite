---
id: collect-261001-general-networking/general-networking/discussion-10524586-basic-bgp-configuration-750851fc
title: "discussion-10524586-basic-bgp-configuration-750851fc"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/discussion-10524586-basic-bgp-configuration-750851fc.md
source_anchor: ""
source_lines: [1, 196]
sha256: 5a21abbdf766fa5550f4c49f7a33621ff33bad6b624f194078bd3d3443b8a50e
---

# discussion-10524586-basic-bgp-configuration-750851fc

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2008 12:32 PM - edited 03-04-2019 12:35 AM
folks
i need to set up a basic bgp config and i've never used bgp before
my isp's router connects to the internet and i connect a hsrp pair of routers to that
i have a default route to the isp but i now need to use bgp to advertise a internal firewall dmz with some public servers (web etc)
the isp's router has an internal IP of
192.168.50.225
my router has an external IP of
192.168.50.236 (the hsrp standby is 192.168.50.238)
the networks i need to advertise are
network 192.168.50.224 mask 255.255.255.240 (my external router interface)
network 172.16.224.112 mask 225.225.255.240 (my internal router interface)
i also need to advertise a static subnet in a firewall dmz
redistribute static 172.16.225.0 225.225.254.0
is the following config ok (as you can see i substituted all the IPs)
thanks to anyone taking the time to read this or reply to it
thanks
router BGP 12345
network 192.168.50.224 mask 255.255.255.240
network 172.16.224.112 mask 225.225.255.240
neighbor 192.168.50.225 remote-as 12345
neighbor 192.168.50.227 remote-as 12345
redistribute static 172.16.225.0 225.225.254.0
static route 172.16.225.0 225.225.254.0 172.16.224.113 (an internal firewall)
thanks to anyone taking the time to read this or reply
greatly appreciated
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Routing Protocols
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2008 12:37 PM
Michael
If your BGP AS is 12345 and the remote AS is 12345 that will be IBGP you run not EBGP. Are you sure this is what you want. I would have though your ISP was using a different AS number.
You don't need the
redistribute static 172.16.225.0 225.225.255.240.0
instead under your BGP config
router BGP 12345
network 172.16.225.0 mask 255.255.255.240
this assumes you have the static route in your routing table.
Jon
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2008 12:37 PM
Michael
If your BGP AS is 12345 and the remote AS is 12345 that will be IBGP you run not EBGP. Are you sure this is what you want. I would have though your ISP was using a different AS number.
You don't need the
redistribute static 172.16.225.0 225.225.255.240.0
instead under your BGP config
router BGP 12345
network 172.16.225.0 mask 255.255.255.240
this assumes you have the static route in your routing table.
Jon
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2008 12:52 PM
folks
many thanks for your input, its greatly appreciated
from jon's post i realise i should have added that AS no 12345 is a private AS number between my kit and the ISP
thanks again to you all
i'll make sure to rate all your posts tomorrow
thanks again
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-09-2008 12:03 PM
jon
many thanks for your input
i've resolved the problem
you where quite right, i was incorrectly using the same AS no inside the router bgp config, i should have been using the ISPs AS no
also i was able to redistribute several networks deep inside my network by configuring several static routes and then using a network statement for each network
many thanks again for your time and patience
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-09-2008 12:06 PM
Michael
Glad you got it working and thanks for letting us know.
Jon
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2008 12:40 PM
I had to implement same setup several times, to me the simplest was a default route to the firewall. Advertising the firewall dmz and the internal web servers was overkill.
Hope this helps, apologizze if I missed something
Ray
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2008 12:53 PM
many thanks ray
greatly appreciated
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2008 12:43 PM
hello Michael,
the configuration can even be simpler.
BGP network command is different: it can be used to advertise a prefix if it is installed in the routing table by any means (including static routes)
so I would suggest only one change:
network 172.16.225.0 225.225.255.240.0
+
no auto-summary
note:
this
redistribute static 172.16.225.0 225.225.255.240.0
is not a correct command you should reference a route map that calls an access-list that matches the desired prefix or a prefix-list (easier to configure )
the static syntax is
ip route 172.16.225.0 225.225.225.240.0 ***.***.***.***
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2008 12:54 PM
giuslar
thanks for your contribution
i appreciate all your posts
thanks again
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2008 01:17 PM
Michael
I wonder if you really need to run BGP. In your original post you explain that: "i now need to use bgp to advertise a internal firewall dmz with some public servers (web etc)". If what you need is to have Internet visibility to those public addresses I wonder if the ISP would just static route those addresses to you and advertise them to the Internet. This would seem to achieve what you need and to avoid the complexity of running BGP.
HTH
Rick
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-03-2008 01:21 PM
rick
this is what i thought but we have a DR site on the same external network with the ISP and they're insistent we use BGP so that failover works between the two sites
i had hoped we could use a internal routing protocol and they could import the routes we advertise but .......
thanks for your interest
greatly appreciated
