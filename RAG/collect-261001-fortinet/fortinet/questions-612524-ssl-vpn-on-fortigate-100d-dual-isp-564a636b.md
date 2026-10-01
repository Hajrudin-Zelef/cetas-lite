---
id: collect-261001-fortinet/fortinet/questions-612524-ssl-vpn-on-fortigate-100d-dual-isp-564a636b
title: "questions-612524-ssl-vpn-on-fortigate-100d-dual-isp-564a636b"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "research"]
source: docs/RAG/collect-261001-fortinet/questions-612524-ssl-vpn-on-fortigate-100d-dual-isp-564a636b.md
source_anchor: ""
source_lines: [1, 14]
sha256: a6d29f94d6a5fce359c5d62b2946a832b6e747f3cc596e9d6b5fa9e38b378b76
---

# questions-612524-ssl-vpn-on-fortigate-100d-dual-isp-564a636b

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a Fortigate 100D and have been using it with a single internet connection for some time without issue and have also been using SSL VPN to connect into the network. The SSL VPN uses 2 factor authentication (Fortitoken).
I have added a second ISP connection and configured Equal Cost Multi Path (ECMP) Routing. This is configured that if an internet IP can not be reached the path will be marked as down.
The problem I have is that I can't get the SSL VPN to work on WAN2 and I'm wondering what the best way to design this is: if I get the SSL VPN working on WAN2, should I create a DNS Name which has two A records for the two public IPs of the connections so if either is up the clients will be able to get a connection to the firewall to authenticate?
Should I create a DNS Name which has two A records for the two public IPs of the connections so if either is up the clients will be able to get a connection to the firewall to authenticate?
If you do, the clients will get one IP address from DNS unpredictably, and then DNS caching will ensure they get stuck with it. If that connection fails, the client will not try the other address, or know anything about it.
Round-robin DNS is a (bad) load balancer, it's not a failover solution.
If the Fortigate SSL VPN client can support a primary and secondary connection address, make a DNS name for each one and put them both in. If it cannot, then you might just have to tell everyone both and get them to try one then the other.
I can't see anything in the Fortigate SSL VPN documentation that mentions failover, or multi-wan or backup or secondary addresses at all, so I guess it has no way to handle it automatically.
