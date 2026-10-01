---
id: collect-261001-general-networking/general-networking/bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget-3
title: "bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget.md
source_anchor: ""
source_lines: [177, 220]
sha256: fd6fda6e0e777190ad78d7b0d61feb57ebca98ace5b3e1e73862dfbf90e1c01c
---

# bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget

- The route reflector sends to every other peer what it receives from a route reflector client or an external BGP peer.
- The route reflector only sends what it receives from non-route reflector clients to other clients and external BGP peers.

Once you've established these rules, it's time to examine the BGP sessions in your network. Check every BGP router on the way, and ensure they don't violate the route reflector rules. The BGP prefixes propagate from every edge router to all other routers using these rules. This is where thorough network documentation is essential.

Another common reason an IP prefix doesn't propagate across your network is that the external subnets on the edge of your network are not advertised to your core routers.

The IP address of the next-hop router doesn't change when a BGP router sends an IP prefix to an internal BGP neighbor. Thus, the IP next hop of an external route is always the IP address of a router one hop beyond the edge of your AS.

Network administrators must insert the IP subnets that connect edge routers to their external neighbors into their internal routing protocol, such as OSPF or IS-IS. Otherwise, some internal BGP routers decide if the BGP next hop is unreachable and ignore the IP prefix. It appears in the BGP table, but the router can't use or propagate it to other BGP peers.

**Is the prefix sent to external neighbors?**

As the last step in troubleshooting BGP route propagation, check whether the IP prefixes transported across your network are announced to your external BGP peers. This article explains the techniques for troubleshooting outbound BGP route propagation.

**Is the traffic traversing the network?**

Even if your BGP route propagation works well, the IP packets might not be able to traverse your network. These are pure IP networks, so the process could change if you add MPLS.

The most common cause of a black hole in your network is a router in the transit path that doesn't run BGP and doesn't know how to route the received IP packet toward the destination network. IP routing works hop by hop. Even though the ingress edge router knows which egress edge router to use and how to get there, it can't pass that information to the intermediate routers. All of them must also run BGP.

To identify a black hole in your network, perform a traceroute from your customer's network to an internet destination. The last router that responds to the traceroute is one hop before the black hole.

Even though all core routers in your network must run BGP, the internal BGP sessions don't have to follow the network's physical structure. For example, you could have a few central routers acting as route reflectors for all BGP routers in your network.

BGP is a critical but complex protocol. Many factors complicate its configuration and troubleshooting, including the following:

- Peering agreements that define the routing policies and security settings for route exchanges between separate networks.
- Security settings, including filters, authentication or other controls.
- Areas of the network that are beyond your control, such as other ISPs or partner networks.
- Custom maximum transmission unit settings that don't match.

Begin working with BGP after determining it's necessary. It's a common protocol for large, routed networks and ISPs, but most small organizations never need BGP. If your organization establishes a need for BGP, address some prerequisites before you begin configuration:

- Create an accurate network map.
- Confirm you have the correct identifiers, including IP addresses, router IDs and AS numbers. You need these values to configure BGP, so document them beforehand.
- Address any basic connectivity, router security or network congestion challenges.
- Confirm you have the correct authentication information for all devices.

Don't underestimate the complexity of deploying and managing a BGP-enabled network. BGP is a complex protocol, but you can grasp this essential internet routing protocol by understanding the configuration steps and creating trustworthy network documentation.

**Editor's note:** *This article was originally written by Ivan Pepelnjak and expanded by Damon Garn to include more BGP troubleshooting information.*

*Damon Garn owns Cogspinner Coaction and provides freelance IT writing and editing services. He has written multiple CompTIA study guides, including the Linux+, Cloud Essentials+ and Server+ guides, and contributes extensively to Informa TechTarget Editorial, The New Stack and CompTIA Blogs.*
