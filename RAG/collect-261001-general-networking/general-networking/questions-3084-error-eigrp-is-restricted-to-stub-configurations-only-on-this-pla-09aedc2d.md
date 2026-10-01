---
id: collect-261001-general-networking/general-networking/questions-3084-error-eigrp-is-restricted-to-stub-configurations-only-on-this-pla-09aedc2d
title: "questions-3084-error-eigrp-is-restricted-to-stub-configurations-only-on-this-pla-09aedc2d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "license", "research"]
source: docs/RAG/collect-261001-general-networking/questions-3084-error-eigrp-is-restricted-to-stub-configurations-only-on-this-pla-09aedc2d.md
source_anchor: ""
source_lines: [1, 28]
sha256: c377be7934f21a5b906f47c2f8ad880a4677057911dc48aa47e4658319ba467b
---

# questions-3084-error-eigrp-is-restricted-to-stub-configurations-only-on-this-pla-09aedc2d

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
12
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We just bought some new Cat4507R+E distribution switches. When we bought the switches, we asked for an Advanced Services K9 image, but our reseller told us they only come with a Universal K9 image. I didn't know better, so I said OK.
However, when I tried to bring the switches online in production, I had a problem with a huge number of unreachable subnets. After sweating it for a very long fifteen minutes (while my boss insulted me for not testing the switch), I finally tracked the problem down to eigrp stub connected in the EIGRP process.
I tried remove the "eigrp stub" statement but I couldn't. When I tried, the switch errors out with EIGRP is restricted to stub configurations only on this platform.
I had to roll back to our original distributions and my boss got even hotter about it.
My google searches blame this problem on the ip base license, which only supports EIGRP stub routing. Our reseller says it will take at least 24 hours and thousands of dollars to get the Advanced Services K9 upgrade. We didn't budget for the extra expense and we have other business units waiting on the upgrade. Oh yeah, and my boss isn't happy.
Before we borrow money from another department, is there anything else we can do to get this working now?
I really don't like how Cisco crippled EIGRP in the ipbase image... at a minimum, you should get some kind of warning that your configuration won't work, but they haven't done that either.
As you discovered, eigrp stub prevents EIGRP from advertising downstream routes from the switches connected to your distribution.
You can trick EIGRP into advertising routes it learned from other rotuers (even as a stub router) by using a stub leak-map...
ip prefix-list p~MATCH_ANY seq 5 permit 0.0.0.0/0 le 32
!
route-map r~MATCH_ANY permit 10
match ip address prefix-list p~MATCH_ANY
!
router eigrp 100
no auto-summary
network 172.16.0.0
eigrp log-neighbor-warnings
eigrp log-neighbor-changes
eigrp stub connected summary static redistributed leak-map r~MATCH_ANY
Technically, you can get away with using eigrp stub leak-map indefinitely, without upgrading your ipbase image.
Your 4507 comes with an evaluation license that will let you run enterprise services for 60 days. You can activate the eval with the "license boot level entservices" command. Note this requires a reload of the switch. This will buy you some time while you purchase a permanent license.
