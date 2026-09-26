---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-44-1
title: "Types of NAT:"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-44.md
source_anchor: ""
source_lines: [1, 78]
sha256: 438729bdb3befad62a9263ea0e7b3c48949745d385ee3dd779fb52e809f71ecb
---

# Types of NAT:

Network Address Translation is an Internet standard that allows hosts on local area networks to use one set of IP addresses for internal communications and another set of IP addresses for external communications. A LAN that uses NAT is ascribed as a *natted* network. For NAT to function, there should be a NAT gateway in each *natted* network. The NAT gateway (NAT router) performs IP address rewriting on the way while packets travel from/to LAN. In RouterOS NAT is supported for IPv4. RouterOS does not support NAT64. 

Nat matches only the first packet of the connection, connection tracking remembers the action and performs on all other packets belonging to the same connection.

Whenever NAT rules are changed or added, the connection tracking table should be cleared otherwise NAT rules may seem to be not functioning correctly until the connection entry expires.

# Types of NAT:

There are two types of NAT:

- **source NAT or srcnat.** This type of NAT is performed on__packets that are originated__ from a natted network. A NAT router replaces the private source address of an IP packet with a new public IP address as it travels through the router. A reverse operation is applied to the reply packets traveling in the other direction.
- **destination NAT or dstnat.** This type of NAT is performed on__packets that are destined__ for the natted network. It is most commonly used to make hosts on a private network to be accessible from the Internet. A NAT router performing dstnat replaces the destination IP address of an IP packet as it travels through the router toward a private network.

Since RouterOS v7 the firewall NAT has two new *INPUT* and *OUTPUT* chains which are traversed for packets delivered to and sent from applications running on the local machine:

- **input** - used to process packets__entering the router__ through one of the interfaces with the destination IP address which is one of the router's addresses. Packets passing through the router are not processed against the rules of the input chain.
- **output** - used to process packets that__originated from the router__ and leave it through one of the interfaces. Packets passing through the router are not processed against the rules of the output chain.

## Destination NAT

Network address translation works by modifying network address information in the packet's IP header. Let`s take a look at the common setup where a network administrator wants to access an office server from the internet.

We want to allow connections from the internet to the office server whose local IP is 10.0.0.3. In this case, we have to configure a destination address translation rule on the office gateway router:

The rule above translates: when an incoming connection requests TCP port 22 with destination address 172.16.16.1, use the *dst-nat* action and depart packets to the device with local IP address 10.0.0.3 and port 22.

To allow access only from the PC at home, we can improve our *dst-nat* rule with *"src-address=192.168.88.1"* which is a Home`s PC public (this example) IP address. It is also considered to be more secure!

## Source NAT

If you want to hide your local devices behind your public IP address received from the ISP, you should configure the source network address translation (masquerading) feature of the MikroTik router. 

Let`s assume you want to hide both the office computer and server behind the public IP 172.16.16.1, the rule will look like the following one:

Now your ISP will see all the requests coming with IP 172.16.16.1 and they will not see your LAN network IP addresses.

### Masquerade

Firewall NAT `action=masquerade` is a unique subversion of `action=srcnat`*,* it was designed for specific use in situations when public IP can randomly change, for example, DHCP server changes assigned IP or PPPoE tunnel after disconnect gets a different IP, in short - **when public IP is dynamic**.

Every time when interface disconnects and/or its IP address changes, the router will clear all masqueraded connection tracking entries related to the interface, this way improving system recovery time after public IP change. If `srcnat` is used instead of `masquerade`*,* connection tracking entries remain and connections can simply resume after a link failure.

Unfortunately, this can lead to some issues with unstable links when the connection gets routed over different links after the primary link goes down. In such a scenario following things can happen:

- on disconnect, all related connection tracking entries are purged;
- next packet from every purged (previously masqueraded) connection will come into the firewall as *new* , and, if a primary interface is not back, a packet will be routed out via an alternative route (if you have any) thus creating a new masqueraded connection;
- the primary link comes back, routing is restored over the primary link, so packets that belong to existing connections are sent over the primary interface without being masqueraded, that way leaking local IPs to a public network.

To work around this situation **blackhole** route can be created as an alternative to the route that might disappear on disconnect.

Hosts behind a NAT-enabled router do not have true end-to-end connectivity. Therefore some Internet protocols might not work in scenarios with NAT. Services that require the initiation of TCP connection from outside the private network or stateless protocols such as UDP, can be disrupted.

To overcome these limitations RouterOS includes a number of so-called NAT helpers, that enable NAT traversal for various protocols. When `action=srcnat` is used instead, connection tracking entries remain and connections can simply resume.

Though Source NAT and masquerading perform the same fundamental function: mapping one address space into another one, the details differ slightly. Most noticeably, masquerading chooses the source IP address for the outbound packet from the IP bound to the interface through which the packet will exit.

### CGNAT (NAT444)

To combat IPv4 address exhaustion, a new RFC 6598 was deployed. The idea is to use shared 100.64.0.0/10 address space inside the carrier's network and perform NAT on the carrier's edge router to a single public IP or public IP range.

Because of the nature of such a setup, it is also called NAT444, as opposed to a NAT44 network for a 'normal' NAT environment, three different IPv4 address spaces are involved.

CGNAT configuration on RouterOS does not differ from any other regular source NAT configuration:

Where:

- 2.2.2.2 - public IP address,
- public_if - interface on provider's edge router connected to the internet

The advantage of NAT444 is obvious, fewer public IPv4 addresses are used. But this technique comes with major drawbacks:

- The service provider router performing CGNAT needs to maintain a state table for all the address translations: this requires a lot of memory and CPU resources.
- Console gaming problems. Some games fail when two subscribers using the same outside public IPv4 address try to connect to each other.
- Tracking users for legal reasons means extra logging, as multiple households go behind one public address.
- Anything requiring incoming connections is broken. While this already was the case with regular NAT, end-users could usually still set up port forwarding on their NAT router. CGNAT makes this impossible. This means no web servers can be hosted here, and IP Phones cannot receive incoming calls by default either.
- Some web servers only allow a maximum number of connections from the same public IP address, as a means to counter DoS attacks like SYN floods. Using CGNAT this limit is reached more often and some services may be of poor quality.
- 6to4 requires globally reachable addresses and will not work in networks that employ addresses with a limited topological span.

