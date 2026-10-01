---
id: collect-261001-fortinet/fortinet/questions-45464-fortigate-interfaces-in-policies-0c81551b
title: "questions-45464-fortigate-interfaces-in-policies-0c81551b"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-45464-fortigate-interfaces-in-policies-0c81551b.md
source_anchor: ""
source_lines: [1, 30]
sha256: fdc4b7b0e5ca82cd581b380493d2d3a1b817ea2b3ff2983c667dc1b033a021ed
---

# questions-45464-fortigate-interfaces-in-policies-0c81551b

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
What are the implications of not using interfaces when defining policies on fortigate firewalls?
i.e
Policy would be
edit 1
set srcintf "any"
set dstintf "any"
set srcaddr "ip-192.168.1.10"
set dstaddr "ip-172.16.1.254
set status enable
set schedule "always"
set service "TCP-8080"
set logtraffic disable
next
Is there any benefits of specifying the interface?
My understanding is the fortigate firewalls apply RPF checks against traffic when it enters the firewall.
For example if packet from 192.168.1.0/24 via port1, It would expect a route in the routing table to exist for that prefix if not RPF would block it.
We are looking at automating how we populate firewall rules and removing the interface logic would make it a a lot simpler to implement.
The interface is part of the policy. When you use any, the rule applies to all interfaces, ie. it doesn't matter where the packet enters or is bound to exit the firewall for the policy to be applied.
FGs apply RPF - in order to accept a source address on an interface there must be an appropriate route out of that interface. So, essentially you're not required to use interfaces in the policies at all times.
RPF can be disabled by turning on asymmetric routing (config system setting, set asymmetric enable), disabling stateful inspection on the way. A better way may be to set a route with a high metric that won't ever get used.
There is hardly ever a good or mandatory reason not to specify the interfaces in a policy. There is always a price you pay for convenience: loss of control and time while maintaining and debugging.
Try to determine the flow of traffic through, say, 80 firewall policies most of which use the 'any' interface. Which one is 'shadowing' others, and are explicit policies hit at all, if not, why? This is near to impossible. Of course you can sniff but that's in realtime. Firewalls are a lot about planning and foreseeing.
To put it differently: policies reflect the flow of data (and it's protection) in the way you have planned and foreseen it. Firewalls are built to prevent deviations from that plan. That's why there are implicit DENY policies for each interface pair in FortiOS, as an example. You do have the information about which network exists (or is routed to) behind each interface. In not using this information you prevent the firewall from attaining it's full potential.
The loss of RPF is IMHO only a side effect of the 'any' interface. Usually, if the protected networks are connected to the internet you need a default route. This route is used all the time if you use an 'any' interface, even for traffic of unknown origin. Effectively, this eliminates RPF.
