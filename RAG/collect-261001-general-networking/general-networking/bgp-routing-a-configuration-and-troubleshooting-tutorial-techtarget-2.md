---
id: collect-261001-general-networking/general-networking/bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget-2
title: "bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget.md
source_anchor: ""
source_lines: [88, 176]
sha256: 3d6e4caf82e883ac7c57faaadc383272a2fcd91e7621fefaaa1154fb8edb698d
---

# bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget

First, identify the source of the connectivity problem you're debugging. BGP might be involved if a customer reports they have limited or no internet connectivity beyond your network. Follow these quick steps to find the source:

1. Perform a traceroute from a workstation on the problematic LAN. If the trace reaches the first BGP-speaking router or gets beyond your network's edge, it's likely a BGP issue.
2. Check whether the BGP-speaking router advertises a default route into your network. Without a default route, other routers in your network can't reach internet destinations.

If you can't access a LAN-attached workstation, perform the traceroute from the customer-premises router. Be sure to use the router's LAN address as the source IP address in the traceroute packets.

### Adjacent BGP router issue troubleshooting

BGP must establish a TCP session between adjacent BGP routers before they can exchange routes. The first step is to check the status of the BGP sessions between the routers.

Administrators configure the BGP neighbors manually. Potential configuration errors include the following:

- **Neighbor IP address mismatch.** The destination IP address configured on one BGP neighbor must match the source IP address -- or the IP address of the directly connected interface -- configured on the other.
- **AS number mismatch.** The neighbor AS number configured on one side of the BGP session must match the neighbor's actual AS number.

You could also have a problem with packet filters deployed on the BGP-speaking router. These filters must send packets to and from TCP port 179.

### BGP route propagation troubleshooting

If your users want to receive traffic from the internet, the IP prefix assigned to your network must be visible throughout the internet. The following steps outline how to create a visible IP prefix:

1. Network administrators insert the IP prefix into their BGP routing table.
2. The BGP router advertises the IP prefix to its BGP neighbors.
3. Other networks propagate the IP prefix throughout the internet.

**Is the route inserted into BGP?**

Most routing protocols automatically insert directly connected IP subnets into their routing tables or databases. Due to security requirements, BGP is an exception. It originates an IP prefix only if you manually configure it. For example, Cisco routers configure advertised IP prefixes with the network statement.

Another option is route redistribution, which enables the network to direct traffic from a different protocol. However, this is highly discouraged in the internet environment.

Furthermore, BGP announces a configured IP prefix only if the IP routing table has a matching route to avoid attracting unrouteable traffic. You can generate the matching IP route through route summarization, but it's usually best to configure a static route that points to a null interface or its equivalent.

To check whether your IP prefix is in your BGP routing table, use a BGP show command. These commands vary depending on the type of routers you use. For example, type show ip bgp prefix mask on a Cisco router.

**Is the route advertised to your neighbors?**

The BGP router announces IP prefixes in the BGP table to all neighbors by default. You must implement output and input filters to modify this default behavior to comply with security and routing policy requirements.

If you have applied output filters toward your BGP neighbors, you must check whether these filters enable the BGP router to propagate the IP prefix to the external neighbors. The command to display routes advertised to a BGP neighbor on a Cisco router is show ip bgp neighbor ip-address advertised.

**Is the route visible throughout the internet?**

Even if you successfully announce your IP prefix to your BGP neighbors, the router might not propagate it throughout the internet. It's hard to figure out what the BGP router propagates beyond the boundaries of your network, but BGP looking glasses tools can help with this process. With these tools, you can inspect BGP tables at various points throughout the internet and check whether your IP prefix has made it to those destinations.

A few factors could block your IP prefix from the internet. The most common one is BGP route flap dampening: If an IP prefix flaps -- meaning it disappears and reappears -- too often in a short period, the prefix gets blocked for an extended period of time, up to an hour by default. This might happen if you clear your BGP sessions or change a configuration. If your IP prefix is dampened, you must wait it out.

You could also have an invalid or missing entry in the IP routing registries. Another possibility is that the upstream ISPs might have inbound filters. These problems are beyond the scope of technical BGP troubleshooting, so your upstream ISP can help you fix the problem.

### Advanced BGP troubleshooting

So far, this article has addressed identifying whether a routing problem is a BGP problem, troubleshooting BGP sessions, and troubleshooting IP route origination and propagation.

Next, let's focus on a more advanced scenario: transit ISP networks.

To establish end-to-end connectivity across a service provider network, the ISP must receive customers' IP prefixes via BGP and announce them to other ISPs. The exact process must happen in reverse; the ISP must announce the default route to the customer.

Network-wide BGP troubleshooting includes the following steps:

1. Receive the IP prefix.
2. Propagate the IP prefix across your network.
3. Send the IP prefix to external BGP neighbors at the other edge of the network.

**Have you received the prefix?**

Identifying inbound BGP problems can be the most difficult part of troubleshooting. Potential reasons an IP prefix isn't in your BGP table are as follows:

- The neighbor is not sending the prefix.
- Your inbound filters are blocking the prefix.

The debugging facility on your edge router is the only tool to help you identify the problem. You typically don't have access to the other BGP neighbor. When debugging, you must recognize that a BGP neighbor can send several hundred thousand routes. Ensure the debugging output produced by the troubleshooting session doesn't overwhelm the router.

Furthermore, the BGP router sends prefixes only when they change, not periodically like with Routing Information Protocol updates or OSPF link-state advertisement floods. Your debugging tool does not show you an IP prefix until it has changed or you've cleared the BGP session with your neighbor.

Some BGP routers can store a separate copy of all routes a neighbor sends in a parallel BGP table. To enable this functionality on Cisco Internetworking Operating System, configure soft-reconfiguration in for a BGP neighbor.

With the parallel per-neighbor table, you can pinpoint what the neighbor has sent or the content of the parallel table. You can also see the routes that have passed your input filters or the contents of the main BGP table. However, the parallel per-neighbor table consumes a large amount of memory.

**Is the IP prefix propagated across your network?**

Even when an edge router receives an IP prefix via BGP, it might not propagate to the other end of your network. First, an internal BGP -- or BGP within a single AS -- requires a full mesh of BGP sessions among all routers. Every router between every pair of edge routers must run BGP, or the network might drop the traffic. This means the number of BGP sessions could become excessively large.

This illustrates the BGP sessions needed in a small four-router network.

Two tools -- BGP route reflectors and BGP confederations -- can help you keep the number of sessions to a reasonable level. Route reflectors are most commonly used.

BGP route reflector rules are as follows:

