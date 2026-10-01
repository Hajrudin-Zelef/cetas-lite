---
id: collect-261001-cisco/cisco/questions-410591-cisco-asa-and-multiple-vlans-55c49300
title: "questions-410591-cisco-asa-and-multiple-vlans-55c49300"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research", "throughput"]
source: docs/RAG/collect-261001-cisco/questions-410591-cisco-asa-and-multiple-vlans-55c49300.md
source_anchor: ""
source_lines: [1, 28]
sha256: 2714c74b54fe17ca2c4c729339ba0fde6abb9689f847f8e9bdd41684e250a278
---

# questions-410591-cisco-asa-and-multiple-vlans-55c49300

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
9
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I currently manage 6 Cisco ASA devices (2 pairs of 5510s and 1 pair of 5550s). They all work quite nicely and are stable so this is more of a best-practice advice question rather then "OMG it's broken help me fix it".
My network is split into multiple VLANs. Pretty much each service role has its own VLAN so DB servers would have their own VLAN, APP servers, Cassandra nodes.
Traffic is being managed on an allow only specific, deny rest basics (so default policy is to drop all traffic). I do this by creating two ACLs per network interface, eg.:
access-list dc2-850-db-in ACL that's being applied to dc2-850-db interface in "in" direction
access-list dc2-850-db-out ACL that's being applied to dc2-850-db interface in "out" direction
It's all pretty tight and works as expected, however I was wondering whether this is the best possible way to go? At the moment I got to a point where I have over 30 VLANs and I must say it's becoming slightly confusing at some points to manage those.
Probably something like common/shared ACLs would help here which I could inherit from other ACLs but AFAIK there's no such thing...
For you having Cisco ASA devices (2 pairs of 5510s and 1 pair of 5550s). This means you are moving away from packet filtering with acls and moving to firewall zone based techniques in ASAs.
Create class-maps ,policy-maps and service-policies.
These techniques were made for it to be less confusing as the areas increase.
There is a book ,You might want to the read.
The accidental adminitrator -It really helped me.
Have a look at it and move from the acls in in two different directions.
With ASAs you should have no problem.
In the past ,I made 800 series ip inspect and ZBF ,then compared there advantages and they used the same technique in the ASAs moving away from packet filtering to advanced ip inspect.
One very simple (and, admittedly, a bit of a cheat) solution would be to assign each VLAN interface a security-level consistent with the traffic it needs to allow.
You can then set same-security-traffic permit inter-interface, thus obviating the need to specifically route and secure the same VLAN across multiple devices.
It wouldn't cut down on the number of VLANs, but it would probably halve the number of ACLs you need for VLANs that reach across all 3 firewalls.
Of course, there's no way for me to know if this makes sense in your environment.
Why do you have both inbound and outbound access lists? You should try to catch traffic as close to the source as possible. That would mean only inbound access lists, halving your total number of ACL's. This would help keep the scope down. When having only one possible access list per flow, your ASA will become easier to maintain and more importantly: easier to troubleshoot when things go wrong.
Also, do all VLAN's have to go past a firewall to reach each other? This severely limits throughput. Remember: an ASA is a firewall, not a (good) router.
private vlans? Another alternative might be breaking business units intoVRFs. Either of those might help manage some of the explosion of ACL requirements. Honestly though, it's hard to comment on this question because so much depends on the business and technical reasons for your existing design
