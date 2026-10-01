---
id: collect-261001-general-networking/general-networking/questions-1745-inbound-bgp-load-balancing-from-same-isp-router-899a7741
title: "questions-1745-inbound-bgp-load-balancing-from-same-isp-router-899a7741"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/collect-261001-general-networking/questions-1745-inbound-bgp-load-balancing-from-same-isp-router-899a7741.md
source_anchor: ""
source_lines: [1, 32]
sha256: 487b40e2a955d324f58d29c61688c57b8b92a56d2f15a89d1fcc425d62725caf
---

# questions-1745-inbound-bgp-load-balancing-from-same-isp-router-899a7741

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
10
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have two BGP routers for DIA with an issue with one provider in that the two links to this provider land on the same SP router. My ethernet handoffs via two separate MetroE links from another provider on behalf on my main provider as the one provider already had fiber to the bldg. If someone could also clear up the ISP terminology when one provider transports services for another, I'd appreciate it. The two circuits terminate L3 with the same SP router, so my two routers each peer with the same provider router. I'm assigned PA space from this SP.
I have no issues with outbound load-balancing (or load-sharing as I guess would be more technically accurate). Outbound, I do ECLB at the firewall which picks one of the two edge routers based on a srcip/dstip hash.
This particular carrier -- forget about the carrier just providing transport -- is not load-balancing inbound traffic from their one router across the two links to my two routers, and this is the direction where we could use the combined 5x50Mb BW that we have contracted. The SP sees equal paths to us for the same advertised network and essentially just the first path they learn is what becomes the bestpath.
I've listed what I'm considering as my options below to get traffic over both circuits, and would like to know what the experts here think is best particularly if you're familiar with typical SP SOPs. Since I have a contract, changing the contract at this time is not an option to have this rebuilt some other way.
Allowing maximum-paths 2 in the SP network fixes this, but this applies to all their BGP customers on the same router that I don't think they will allow. At least one option that will work involves static routes, but that’s not what I would prefer.
Below are the options I have considered in my order of preference.
Allow BGP maximum-paths 2 on SP router (affects all BGP customers homed there) so /24 is used when advertised on both circuits
Split my /24 in half and advertise separate /25s over each link along with the /24. The SP has recently stated that a non-documented community could be used for them to accept > /24 prefixes. This requires manipulating NAT on my firewall to use global addresses in both /25’s as most traffic now is destined back to us on just a few addresses in the lower /25.
SP static routes to /24 to force load-balancing w/BGP /24 (floating route).
SP static routes to /25s to force ECLB w/BGP /24 prefix (in RIB, but not used unless failure of /25s).
I think advertising the /25s in BGP is the best option which I only found out recently is possible with the SP undocumented community, but are there other options I haven't considered or concerns about out-of-order packets with some of these choices?
This is sort of the reverse load-balancing problem most people ask with BGP.
I would ask for 'maximum-paths' (it's usually called ECMP in standards and documents, not ECLB). And if ECMP is non-starter, then fallback to your /25 plan.
Other acronyms that I couldn't immediately figure out were DIA (dedicated internet access?) and SOP (standard operating procedure?). I'm not sure if these are really so universal acronyms that they should be used in stackexchange without at least hovertext to resolve them.
Talking about terminology you've really already said transport which I think is what that type of service is most commonly referred to as. Sometimes folks will also refer to it as a "tether".
You have another option of moving the interface on your side to the same router and then setting up eBGP multihop*** between you and your provider. Yes, you lose redundancy on your side, but on the opposite side it's going to the same router anyway, so that's kind of a moot point. This also eliminates the need for the provider to do eBGP multipath, in case they're not willing to do that with you (but I've found that in general most providers are fine with turning this on, if it's not turned on already).
If that doesn't seem plausible then announcing two /25's is likely your best bet, unless of course your provider is willing to turn on maximum-paths (again, assuming they haven't done so already).
***Doing load balancing with eBGP multihop in your scenario will involve the following:
Move the second ISP connection to one router on your side.
Both you and the provider configure update-source Loopback0 on each of your sessions - you don't have to use Lo0 if you don't want to, as long as you guys agree on which address to static-route to.
Configure 2x /32 static routes to each others' loopbacks via the connected interfaces (or next-hop IP's) - this is how the load balancing works, as it's really just ECMP.
Configure ebgp-multihop 2 on each others' sessions (You want to keep this number as small as possible to avoid TCP session hijacking).
Voila, load balancing. This also scales per interface, as adding a new port will involve adding another interface and a static route.
The optimum solution here is BGP ECMP via maximum-paths 2 - however, I will say that "accepting routes > /24" requiring a community tag sounds like epic stupidity - assuming you're a customer of theirs, they should accept whatever you give them up to a maximum number of prefixes and simply filter outbound according to whatever agreements they have with their other peers. one of the upstreams I have puts me in a similar boat with regard to not doing ECMP, I always expect to be allowed to announce any prefix size I desire to a transit provider and have them use it.
So, given that they don't sound terribly competent, to balance traffic, do not withdraw your /24 prefix - instead, leave it be and ensure that your /25 routes are tagged with NO_EXPORT in addition to whatever it is that your provider allegedly requires so that your /25s won't accidentally leak out of their AS (not that it's likely to get very far if it does).
One last note - make sure that this "undocumented community" isn't in fact a blackhole community, because that'd be... y'know, bad.
