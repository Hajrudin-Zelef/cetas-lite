---
id: collect-261001-general-networking/general-networking/r-networking-comments-7iom20-bgp-configuration-gotchas-2f54b089
title: "r-networking-comments-7iom20-bgp-configuration-gotchas-2f54b089"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/r-networking-comments-7iom20-bgp-configuration-gotchas-2f54b089.md
source_anchor: ""
source_lines: [1, 44]
sha256: 1004987e98ac7d013ba56e0c2bc9e0aee097faab7b2a9b575a97d414abd88191
---

# r-networking-comments-7iom20-bgp-configuration-gotchas-2f54b089

BGP configuration gotchas 
        
    I've taken over a small ISP with its own ASN and have taken it from being single-homed to multi-homed. I get what BGP is basically but missed a few important details for actually using it at the start.
One is that the "generic" case on our BGP software is actually IPv4 only and there is a completely separate section for IPv6 for almost every single setting. Based on what I've seen, I think this varies based on the router software and some explicitly call out IPv4 in their configuration.
Another is that changing a route-map requires explicitly issuing an outbound soft reset for it to take effect even though the router delays applying changes until they are explicitly committed. We aren't using Cisco, but the outbound soft reset came up in Cisco related documentation and I thought that was worth trying. I don't know if this is generally true to all routers.
Is the information about Cisco and BGP generally applicable enough that I should try some sort of Cisco course to pick up these sorts of details? Where else should I go to find out what I don't know about using BGP on a day to day basis?
If you have any personal tips to share that would also be appreciated.
Section des commentaires
Please buy and read the first 6 chapters of TCP/IP Routing Volume II by Jeff Doyle.
Don't be scared by the "CCIE" title on it, Jeff has a very unique way of introducing concepts and then configuration examples (Albeit on Cisco IOS).
He goes over common gotchas, such as refresh routes after a route policy change (which is most likely required across vendors because it is not part of the BGP RFC).
The first two chapters, in particular, do a very good job of going over newly multi-homed networks and the design considerations that go into determining routing policies.
https://www.amazon.com/Routing-TCP-IP-Professional-Development/dp/1587054701/ref=dp_ob_title_bk
EDIT: Changed link to the second edition, which was just published this year.
The old version of that book, which has DEC listed as hosting a network exchange, happened to be lying around the house. I'll check it out, thanks.
Hard to give any specific advice tbh, sounds like you have the basics down.
The main thing is to make sure you have policies (route-maps) on all peerings. Everything announced/received should be deliberate, you don't want to leave anything to "just happen" without any filters in place.
Other things I'd recommend using Team Cymru BOGON peering and uRPF if you can:
http://www.team-cymru.org/bogon-reference.html
Setting up remote-triggered blackholed is a good idea too if you can, and your upstreams support it (by using community string).
Off topic, but woild you recommend that book for learning multicast routing?
The same book I recommended above covers Cisco implementation of multicast, although I havent read that portion as of yet.
Yep. They are different protocols, so that get different address families. Some platforms do allow you to peer both over a single BGP session, but besides that you still need separate configs for both.
If both peers support route refresh then you don't need to clear the peer. All current Cisco platforms support this, not sure about other vendors.
Most of the BGP related Cisco topics are vendor agnostic, except for the configuration. However, you'd probably be better off finding a course or documentation from the vendor you are using.
I believe route-refresh still requires the peer to be cleared "in" or "out" in order to trigger the refresh.
See these for example: https://ccieblog.co.uk/bgp/route-refresh-capability-vs-soft-reconfiguration https://networklessons.com/bgp/bgp-route-refresh-capability/
http://jaluther.blogspot.ca/2012/04/bgp-route-refresh-capability.html?m=1
From the last one:
This triggers the route-refresh.
I update route maps all the time without resetting the peer.
RFC 2918 doesn't specifically mention how a ROUTE-REFRESH message is triggered. In my experience, it something is changed in a filtering policy, the router sends the ROUTE-REFRESH message immediately, triggering the peer to update it's RIB.
Soft reconfiguration is a different approach which address a copy of the received routing table and is not defined by an RFC. Using the "soft in" command will force the router to re-evaluate the full received route table against inbound filters. The problem with this approach is that it uses more memory as it has to store all route even if they are followed on ingress.
Pretty much all implementations support route-refresh at this stage.
Why are you using route map? Outbound traffic? Why?
To my best knowledge, filtering using a route map instead of a prefix list allows us to attach a community to the set of routes we're advertising. There may be another way to attach a community but that's the way I found.
If you're asking why are we filtering our route announcements at all - without limiting our announcements to just our own networks, if one of our upstreams made a mistake we might find ourselves receiving traffic being sent from one upstream to the other. That would not end well.
You probably have a command along the lines of 'advertised-routes' somewhere that will show you the routes you're advertising.
How many /24 do you have? Why are you using community?
I understand why the route map but simple is best and if you are using route map only to receive traffic and leaving upload on best path then you should be using prefix list.
Even then I'd use a prefix list for out filter and route map for in filter.
Having said all that, I do not know your needs or your design, the questions help me understand, though a simple "explain your needs and info" should've been said.
You can also use route map in conjunction with prefix list.
Multiprotocol BGP extensions solve some of this. An IPv4 peering session can support IPv6 routes and the use of multiprotocol extensions on EBGP peers isn't unusual. There's a good explanation of IPv6 routes on IPv4 sessions in MPLS in the SDN Era.
