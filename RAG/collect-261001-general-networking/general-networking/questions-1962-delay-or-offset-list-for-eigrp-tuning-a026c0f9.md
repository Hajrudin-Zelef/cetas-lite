---
id: collect-261001-general-networking/general-networking/questions-1962-delay-or-offset-list-for-eigrp-tuning-a026c0f9
title: "questions-1962-delay-or-offset-list-for-eigrp-tuning-a026c0f9"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2017-04-19"]
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1962-delay-or-offset-list-for-eigrp-tuning-a026c0f9.md
source_anchor: ""
source_lines: [1, 23]
sha256: 2fb08b873ef87dda2bdcbc43f1727c11312a60628476980f5551d28fd66e0cd5
---

# questions-1962-delay-or-offset-list-for-eigrp-tuning-a026c0f9

Studying for CCNP-ROUTE - everything I read says it's preferable to use delay and offset-lists for metric tuning vs. say bandwidth in EIGRP. In practice what is more commonly used out of those two options (delay/Offset-lists)? How common is low level metric tuning for EIGRP?
4 Answers 4
Any metric manipulation I've ever seen has been via offset-lists. This allows very fine-grained control over the metric, and removes the need to calculate metric differences based on different interface delays.
Note that offset-lists have the added benefit of being applied (optionally) per-interface. So you can essentially do everything that a delay change does with an offset-list.
In the 'wild', the only thing my org uses offset-lists for is to force traffic to take a particular backup WAN link when the primary is having issues that don't actually take it down. That is -- I've seen very little EIGRP metric manipulation in production networks. I'm sure someone else will almost immediately disagree. =)
  In practice what is more commonly used out of those two options (delay/Offset-lists)
Use whatever suits your requirements.
- delay is an interface-level command and gives you control over the delay on a specific interface
- offset-list is a protocol-level command that could change metrics for the whole EIGRP process.
Thus, if it's me and I want to change metrics for an individual link, I use an interface delay... it's easy to understand what is happening when you look at the interface.
If I want to affect all metrics for the whole EIGRP process, I would use a offset-list...
For the most part, it's personal preference how you do it... whatever you do, be consistent throughout the network. In all probability, five years from now, you will be at another company and the "next guy" will have to understand what you do (maybe at 3am in the morning)... consistency is valuable.
  How common is low level metric tuning for EIGRP?
Depends on the company... I have seen it used for stuff like firewall bypass links (i.e. if they have to take the firewall out of service, they bypass the FW with a higher delay links "just in case"). It's most commonly used for a dedicated backup path.
- 
        Cool distinction on when to use each - thanks for your insight!A L– A L2013-06-25 22:19:10 +00:00Commented Jun 25, 2013 at 22:19
Slightly different tools that work best in different places.
The offset-list allows adjustment to routes as they propagate into a given part of the network without altering the metrics of the originating network. In the case of two networks interconnecting at multiple points via EIGRP this allows a given path to be biased upward or downward without playing with a metric that might cause different (likely undesirable) behavior in the existing topology.
Delay is generally the preferred method for tweaking path selection on a universal basis. Bandwidth can work, of course, but always remember that EIGRP considers the lowest bandwidth link in a given path - which means dropping your 1G link to 900M doesn't mean much if it ultimately goes through a 100M connection. Delay, in contrast, is directly additive on a per-hop basis and, as such, has a much greater effect on the route ultimately calculated.
The other K values (reliability, utilization, MTU, etc) should basically -never- be touched. These function are either better accomplished through other mechanisms, will create rolling oscillation issued in your network or just shouldn't realistically be a point of consideration in the first place. The largest EIGRP deployments in the world do fine without touching the other K values. If it is suggested that these values be set I would take a harder look at the requirements and tools available.
offset-list actually alters the delay property for a given prefix matching the ACL. Lab it. you'll see.
- 
        Thank you for participating, but you should really expand your answer to make it more useful. an example of this would help a lot, and a better explanation of why, and perhaps some backup documentation would be helpful, too.2017-04-19 12:59:13 +00:00Commented Apr 19, 2017 at 12:59
