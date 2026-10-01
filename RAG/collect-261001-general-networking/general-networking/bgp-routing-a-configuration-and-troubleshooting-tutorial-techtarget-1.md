---
id: collect-261001-general-networking/general-networking/bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget-1
title: "bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget.md
source_anchor: ""
source_lines: [1, 87]
sha256: 7884a8c6bac5f28e8f49c01436e39096c99552dd7b5c2216074e982804994383
---

# bgp-routing-a-configuration-and-troubleshooting-tutorial-techtarget

Getty Images/iStockphoto

Service providers and network professionals working with IP networks agree that Border Gateway Protocol is one of the most complex and difficult-to-configure internet protocols. However, its emphasis on security and scalability makes it essential to the internet.

This tutorial examines BGP functionality and offers configuration and troubleshooting options, especially for administrators at ISPs or large regional or global network deployments. Use these steps and commands to manage BGP-enabled routers so they can exchange information efficiently and securely within the context of large networks or ISPs.

BGP is the routing protocol that makes the internet work. The protocol exchanges routing table information among BGP-speaking routers, enabling traffic across autonomous and diverse network environments.

Because IP address allocation lacks a structured assignment hierarchy, most service providers' core network routers must exchange information about several hundred thousand IP prefixes. BGP can manage this, as it's a highly scalable and security-focused routing protocol.

ISPs exchange BGP routing information between business entities on the public internet. They use peering agreements to define the necessary security conditions that enable routing exchanges between neighboring ISPs. Administrators configure BGP to adhere to these agreements. All adjacent routers must also satisfy the routing policies. Practical BGP implementations provide a rich set of route filters that enable ISPs to defend their networks and control what they advertise to their competitors.

### Autonomous systems

In BGP terminology, an independent routing domain -- which almost always means an ISP network -- is called an *autonomous system* (AS). ISPs use BGP as the routing protocol of choice in two scenarios:

1. **External BGP.** Exchanges routing tables between different ISPs on the internet.
2. **Internal BGP.** Exchanges routing tables within an ISP's own network.

BGP also evaluates routing tables and information along multiple routes between ISPs, using a selection algorithm that determines the best path to direct traffic based on specified commands and attributes. However, the best path isn't always the shortest path.

### BGP attributes

Many other routing protocols focus only on finding the optimal path toward all known destinations. BGP can't take this simplistic approach because the peering agreements between ISPs often result in complex routing policies. To help network administrators implement these policies, BGP carries many attributes with each IP prefix, including the following:

- **AS path.** The complete path documenting the ASes through which a packet must travel to reach the destination.
- **Local preference.** The internal cost of a destination, which ensures AS-wide consistency.
- **Multi-exit discriminator.** A BGP attribute that gives adjacent ISPs the ability to prefer one peering point over another.
- **Communities.** A set of generic tags that can signal various administrative policies between BGP routers.

Because BGP considers these factors, it might direct traffic to take one path from the destination and another path on its return trip, resulting in asymmetric routing. This isn't usually problematic for basic routing, but could complicate routes with firewalls and VPNs.

### BGP convergence

BGP's design and implementation focus on security and scalability, which makes it harder to configure than other routing protocols. This complexity is even more apparent when configuring other routing policies and network commands. BGP is also one of the slowest converging routing protocols.

The slow BGP convergence dictates a two-protocol design of an ISP network:

- Network administrators often use an internal routing protocol, such as Open Shortest Path First (OSPF) or Intermediate System to Intermediate System (IS-IS), to achieve faster convergence for internal routes, including IP addresses of BGP routers.
- Administrators configure BGP on routers to exchange internet routes.

The fast convergence of OSPF or IS-IS helps bypass a failure within the core network. At the same time, BGP on top of an internal routing protocol meets scalability, security and policy requirements. In addition, if you migrate customer routes into BGP, the customer problems don't affect the stability of your core network. For example, link flaps between your router and a customer's router don't affect your infrastructure.

Because of BGP's inherent complexity, customers and small ISPs often deploy BGP only where needed, such as on peering points. They typically configure it on a minimal subset of core routers, the ones between the peering points, as shown here.

The BGP-speaking routers must also generate a default route into the internal routing protocol to attract the traffic for internet destinations unknown to other routers in your network.

Any organization that wants to achieve redundant internet access must have its own AS and exchange BGP information with its ISPs. You'll likely deploy BGP on core and edge routers, so plan to include BGP on those routers as part of your initial network design, as shown here.

Although you might not rely on BGP everywhere in the initial network deployment, a blueprint could help when you need to scale the BGP-speaking part of the network.

In addition, BGP requires a full mesh of internal BGP sessions between routers in the same AS. It's possible to use BGP route reflectors or confederations to make the network scalable.

Another reason to deploy BGP throughout your network is that MPLS-based VPNs, large-scale quality of service deployments or large-scale differentiated web caching implementations rely on BGP to transport the information they need.

BGP commands and configurations vary by router vendor, but the overall process is similar across devices. Before you begin, ensure you have accurate documentation of all IP addresses, subnet masks, AS references and other network settings. Configuring BGP is challenging enough without adding inaccurate network information. Ensure basic connectivity between all involved devices, too. Don't let a pesky packet filter disrupt the process.

Begin by setting up the router. Enable secure administrative access, update the firmware and review the peering agreements before beginning the configuration. You must specify the BGP neighbors, define routing policies and advertise your router's networks.

### A basic BGP configuration example

An example of the basic BGP configuration steps on a Cisco router is as follows:

1. Enter the BGP configuration mode: router bgp {AS}.
2. Define the router's BGP neighbors using IP addresses and the remote AS: neighbor {IP} remote-as {AS}.
3. Advertise your router's networks to neighbors: network {network-address} mask {subnet-mask}.
4. Verify connectivity using tools like ping.
5. Set a router ID: bgp router-id {router-ID}.
6. Confirm the settings: show ip bgp summary and show ip bgp.

You've now configured BGP between routers in two separate AS environments.

Review the documentation for your specific network devices for additional options and settings. However, the general information should be similar across vendors.

Potential pitfalls include the following:

- BGP flapping -- a router or interface status going up and down -- due to hardware or media issues.
- Network congestion.
- Older, out-of-date router OSes or configurations.
- Authentication issues.

Remember to configure logging and monitoring. Also, test any failover or redundancy configurations to ensure they function as intended.

A structured approach to BGP troubleshooting can lead you from the initial problem diagnosis to a fix. Here is a simple scenario with a single BGP-speaking router in the network.

Multihomed customers and small ISPs use similar designs that don't offer BGP connectivity to their customers.

