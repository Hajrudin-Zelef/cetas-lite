---
id: collect-261001-fortinet/fortinet/questions-25235-fortigate-remembers-bad-routes-afaeb359
title: "questions-25235-fortigate-remembers-bad-routes-afaeb359"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["asic", "parameters"]
source: docs/RAG/collect-261001-fortinet/questions-25235-fortigate-remembers-bad-routes-afaeb359.md
source_anchor: ""
source_lines: [1, 6]
sha256: e718f4c7cff41fc4bbdd14c9057a4c2e97c7e70817f2537aebc7724061fe0447
---

# questions-25235-fortigate-remembers-bad-routes-afaeb359

First off, this behavior is as expected.  
The Fortigate (as a stateful firewall) will create a session from the information of the first packet arriving. It will determine the route to apply and whether forwarding is permitted or not. After these decisions, subsequent traffic belonging to the same session is forwarded without any further decisions to make. (As a point aside, in Fortinet firewalls established sessions are indeed offloaded from the CPU to a network processing ASIC at that point - which is incapable of looking up routes or policies. The principle holds true even if traffic is not offloaded, though.)    
Now, when the tunnel goes down, the route across is deleted from the Routing Table and VPN sessions are deleted. The next best route to a private network usually is the default route, so for arriving traffic destined to the network behind the tunnel new sessions are created and routed out the WAN interface (where they are discarded at the next router, hopefully).  
When the tunnel is coming up again, new sessions are correctly routed and forwarded across the tunnel because their first packet trigger the decisions I mentioned above. Existing sessions are not affected - this would imply that with every change of the routing table, the FGT would have to re-evaluate each session if there was a better route than before. This is impractible at least. Traffic will follow the "wrong" route until the session is timed out.  
There is an easy fix for this kind of scenario: create static blackhole routes on the FGT for all private, RFC1918 networks you might use behind VPN tunnels, with 192.168.x.x/24, 172.16.x.y/16 and 10.x.y.z/8 being the most popular. You need to set the distance parameters for these blackhole routes to 254 to keep them inactive as long as other, intended routes exist. The trick is that for blackhole routes, the FGT will not create sessions. So when the tunnel is down and the tunnel route discarded, the blackhole route is used - packets will be discarded immediately. Each packet arriving will trigger a session setup and a routing decision and eventually will be forwarded across the tunnel right after it comes up again.  
In the Fortinet support forums you will find a couple of threads explaining this. This post offers a bulk command file ready to import all RFC1918 blackhole routes: https://forum.fortinet.com/FindPost/120872
