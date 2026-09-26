---
id: collect-260926-mikrotik/mikrotik/questions-1130412-mikrotik-multiple-wan-incoming-traffic-02c8b62f
title: "Get/Assign WAN1 gateway interface/IP"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-1130412-mikrotik-multiple-wan-incoming-traffic-02c8b62f.md
source_anchor: ""
source_lines: [1, 28]
sha256: f1179d95f8769cda074ce5dc80e8128ebf956effaf4f3389c160c2918a23f49b
---

# Get/Assign WAN1 gateway interface/IP

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
i'm writing here to ask you if you have ever had an issue similar to mine.
I'm using a Mikrotik RB750UP for my office. The router is managing 3 different WANs. I'm not using any balancing or bonding of such lines, but I use them just for failover reasons.
Basically, the connections are always on and a distance-based routing rule is used to switch among the active connections in case of failure of a WAN.
What I need to enforce is to somehow mark the traffic which comes into a WAN and let the router to send the reply back from the same WAN interface and IP.
For instance: I send a request to WAN2_ip:port from an external device. The RB has the WAN1 on (which is also the first choice among the possible routing rules) and uses this connection to send back the reply (which is not the correct one).
I tried to mark the connections coming into a certain interface, and mark the packets belonging to that connection with a routing mark. But it does not seem to work.
I think you're on the right track with the marks. It's been a while since I've used this, but I think you need to use a separate routing table for each WAN link. Similar to the information documented here.
So, for example you'd set something up like the following for the 1st WAN connection and repeat for each additional WAN.
(replace/define ${wan...} variables with appropriate values)
Copy
# Get/Assign WAN1 gateway interface/IP
wan1_interface=eth0
wan1_gateway_ip_address="$(ifconfig br0 | grep 'inet addr:' | cut -d ':' -f2 | cut -d ' ' -f1)"# Create routing table for WAN1
ip route replace default via ${wan1_gateway_ip_address} dev ${wan1_interface} table 100
ip rule add fwmark 0x7 table 100
# Setup policy for marking WAN1 traffic
iptables -t mangle -N wan1_policy
iptables -t mangle -A wan1_policy -j MARK --set-mark 0x7
# Mark traffic coming in on WAN1 interface
iptables -t mangle -A PREROUTING -i ${wan1_interface} -j wan1_policy
I apologize if this is a bit crude or incomplete. Unfortunately I'm a Windows administrator in my profession. So, I don't tend to tinker with the networking or Linux/Unix side of things often enough to be very fluent at this level. But, I will update the answer if I can recall any more specific information from the dusty corners of my brain. :)
UPDATE: I found an old script I'd used and updated the example above with more accurate commands. I did have to tweak them for your situation (I was using them to direct traffic to a VPN connection), but I think they should be good now.
