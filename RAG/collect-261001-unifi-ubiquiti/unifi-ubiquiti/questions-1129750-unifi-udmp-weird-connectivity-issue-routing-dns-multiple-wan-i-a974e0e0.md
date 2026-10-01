---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1129750-unifi-udmp-weird-connectivity-issue-routing-dns-multiple-wan-i-a974e0e0
title: "questions-1129750-unifi-udmp-weird-connectivity-issue-routing-dns-multiple-wan-i-a974e0e0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1129750-unifi-udmp-weird-connectivity-issue-routing-dns-multiple-wan-i-a974e0e0.md
source_anchor: ""
source_lines: [1, 16]
sha256: ca9da8988080c8749893514a25426e4d92504fd1c0e0376afcb0c024b88ad4f5
---

# questions-1129750-unifi-udmp-weird-connectivity-issue-routing-dns-multiple-wan-i-a974e0e0

We are experiencing an odd issue, seemingly related to routing or DNS.
We have a "hub and spoke" topology using Unifi equipment (UDMP's). Each site connects via IPSEC tunnel to an AWS EC2 instance running VyOS to handle core routing between sites and other infrastructure in AWS.
In the past, when we had more of a hybrid topology with some on-prem servers, each site had another IPSEC tunnel connecting to the main office, required for the old VoIP server, and we had a few on-prem DNS servers.
We have since moved all infrastructure into AWS, and these second IPSEC tunnels to the main office are no longer needed. I have taken most of the site's tunnels connecting to the main office down, and everything works fine for those other sites. I have one site left (site3) that is giving me problems whenever I take their tunnel down.
The Issue: Whenever I take down the IPSEC tunnel between "site 3" and the main office, things work for maybe 10 minutes before people start complaining that they "have no internet". I determined they were probably still using the old on-prem DNS servers, so I switched their primary DNS servers to the DNS servers in AWS, with google dns as a backup. Fine, no problem, everything working. I take the tunnel down again, and I start getting calls. This time users say they lost their mapped drives (the file server in AWS).
What is weird is that everything works fine (site 3's connectivity to aws) when their IPSEC tunnel to the main office is up. When I take it down, things work for maybe 10 minutes or so, then it stops working. You would think their site is routing through the tunnel to the main office then up to AWS, but this is not the case. A traceroute from a client machine at site3 shows 3 hops to connect to EC2 instances: out their WAN, to VyOS IP, to server IP. A look at the routing table on client machine at site3 shows no entry for the AWS network, thus traffic is sent to 0.0.0.0, their UDMP gateway. A look at the routing table on the site3 UDMP shows 1 entry for the aws VPC network, 172.30.0.0/16, with the next hop being the VyOS router.
1 interesting detail is that even though everything is set to allow ICMP/respond to ping, neither the UDMP nor the vyos router can ping each other or ec2 instances... however clients on site3 network can ping everything.
I checked the security rules for the EC2 instances, and all required networks and WAN IPs are included.
I am fresh out of ideas when I noticed that site3 udmp is configured with a static WAN IP, but also has configuration settings set for "router", and additional IP addresses. These are the details:
WAN IP=108.x.69.250
subnet mask: 255.255.255.248
Router: 108.x.69.249
Additional IP addresses: 108.x.69.251/32, 108.x.69.252/32, 108.x.69.253/32, 108.x.69.254/32, 108.x.69.255/32
A look in the security rules for AWS/EC2 showed that while 108.x.69.250/32 is allowed, none of the other IPs in the subnet are included (next hop ISP router, or additional IPS). I changed the AWS security allowed entry to 108.x.69.248/29, however this is a hail mary. I'm not too confident this will be the fix.
Anybody have any thoughts or ideas? I can't test again until after hours but I thought I might get someone else's take on the situation. Anyone have experience working with UDMP with static WAN but also with these additional fields configured for router and additional IPs?
I've included a beautiful diagram of the topology for your reading pleasure!
