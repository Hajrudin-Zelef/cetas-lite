---
id: collect-261001-cisco/cisco/questions-16698-cisco-static-nat-with-route-maps-e6ea4b4b
title: "questions-16698-cisco-static-nat-with-route-maps-e6ea4b4b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-16698-cisco-static-nat-with-route-maps-e6ea4b4b.md
source_anchor: ""
source_lines: [1, 19]
sha256: e196f4ec88ce44e768ae60cb44664ce0c2c358723b701960f5f47a771c059466
---

# questions-16698-cisco-static-nat-with-route-maps-e6ea4b4b

Update:
Seems that route maps only match on IP addresses, not on the ports. Had another situation this week on a different device, model and software version. Ended up changing the NAT statements to:
ip nat inside source static tcp 192.168.1.20 3389 x.x.x.x 3389
I then restricted access based on an ACL rather than a route map. Would have been nice to define conditional NATing, but it seems it just doesn't work.
So we have a pretty standard NAT box setup to offer a hosted NAT solution for a number of customers.
Here is the basic topology:
Cisco IOS Software, C2900 Software (C2900-UNIVERSALK9-M), Version 15.2(4)M3, RELEASE SOFTWARE (fc2)
The issue I have is to do with the route-map section of the NAT statements.
ip nat inside source static tcp 10.1.10.201 22 x.x.x.x 22 vrf Customer1-vrf route-map Customer1-portforwarding extendable
ip access-list extended Customer1-forwarding-acl
 permit tcp host 10.1.10.201 host a.a.a.a eq 22
 permit tcp host 10.1.10.201 host b.b.b.b eq 22
route-map Customer1-portforwarding permit 10
 match ip address Customer1-forwarding-acl
I believe I have the understanding of the route-map correct. It is meant to designate what is allowed to NAT and what isn't. I am basically trying to set it up to allow only translations from specific public source addresses. It doesn't seem to be doing that. It seems to be permitting translations from any public address.
I changed the ACL completely to a 'deny ip any any' statement and it still permits. I'm at a bit of a loss. It appears the route-map isn't doing anything.
Any help would be greatly appreciated!
Cheers,
H
