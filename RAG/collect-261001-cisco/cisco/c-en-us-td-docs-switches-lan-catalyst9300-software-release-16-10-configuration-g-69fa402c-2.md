---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c.md
source_anchor: ""
source_lines: [80, 164]
sha256: d9915afcd07caa118f7a451a1b2ca94d82716ba2b0c13ef5eeee17ac4cb8d032
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c

Half Entry—represents a mapping between the local and global address/ports and is maintained in the translation database of
NAT module. A half entry may be created statically or dynamically based on the configured NAT rule.
Full Entry/Flow entry—represents a unique flow corresponding to a given session. In addition to the local to global mapping,
it also maintains the destination information which fully qualifies the given flow. A Full entry is always created dynamically
and maintained in the translation database of NAT module.
Types of NAT
You can configure NAT such that it will advertise only a single address for your entire network to the outside world. Doing
this effectively hides the internal network from the world, giving you some additional security.
The types of NAT include:
Static address translation (static NAT)—Allows one-to-one mapping between local and global addresses.
Dynamic address translation (dynamic NAT)—Maps unregistered IP addresses to registered IP addresses from a pool of registered
IP addresses.
Overloading / PAT—Maps multiple unregistered IP addresses to a single registered IP address (many to one) using different
Layer 4 ports. This method is also known as Port Address Translation (PAT). By using overloading, thousands of users can be
connected to the Internet by using only one real global IP address.
Using NAT to Route Packets to the Outside Network (Inside Source Address Translation)
You can translate unregistered IP addresses into globally unique IP addresses when communicating outside your network.
You can configure static or dynamic inside source address translation as follows:
Static translation establishes a one-to-one mapping between the inside local address and an inside global address. Static
translation is useful when a host on the inside must be accessible by a fixed address from the outside. Static translation
can be enabled by configuring a static NAT rule as explained in the Configuring Static Translation of Inside Source Addresses section.
Dynamic translation establishes a mapping between an inside local address and a pool of global addresses dynamically. Dynamic
translation can be enabled by configuring a dynamic NAT rule and the mapping is established based on the result of the evaluation
of the configured rule at run-time. You can employ an Access Control List (ACL), both Standarad and Extended ACLs, to specify
the inside local address. The inside global address can be specified through an address pool or an interface. Dynamic translation
is enabled by configuring a dynamic rule as explained in the Dynamic translation section section.
The figure below illustrates a device that is translating a source address inside a network to a source address outside the
network.
The following process describes the inside source address translation, as shown in the figure above:
The user at host 10.1.1.1 opens a connection to Host B in the outside network.
NAT module intercepts the corresponding packet and attempts to translate the packet.
The following scenarios are possible based on the presence or absence of a matching NAT rule:
If a matching static translation rule exists, the packet gets translated to the corresponding inside global address. Otherwise,
the packet is matched against the dynamic translation rule and in the event of a successful match, it gets translated to the
corresponding inside global address. The NAT module inserts a fully qualified flow entry corresponding to the translated packet,
into its translation database. This facilitates fast translation and forwarding of the packets corresponding to this flow,
in either direction.
The packet gets forwarded without any address translation in the absence of a successful rule match.
The packet gets dropped in the event of failure to obtain a valid inside global address even-though we have a successful rule
match.
Note
If an ACL is employed for dynamic translation, NAT evaluates the ACL and ensures that only the packets that are permitted
by the given ACL are considered for translation.
The device replaces the inside local source address of host 10.1.1.1 with the inside global address of the translation, 203.0.113.2,
and forwards the packet.
4. Host B receives the packet and responds to host 10.1.1.1 by using the inside global IP destination address (DA) 203.0.113.2
The response packet from host B would be destined to the inside global address and the NAT module intercepts this packet and
translates it back to the corresponding inside local address with the help of the flow entry that has been setup in the translation
database.
Host 10.1.1.1 receives the packet and continues the conversation. The device performs Steps 2 to 5 for each packet that it
receives.
Outside Source Address Translation
You can translate the source address of the IP packets that travel from outside of the network to inside the network. This
type of translation is usually employed in conjunction with inside source address translation to interconnect overlapping
networks.
You can conserve addresses in the inside global address pool by allowing a device to use one global address for many local
addresses and this type of NAT configuration is called overloading or port address translation. When overloading is configured,
the device maintains enough information from higher-level protocols (for example, TCP or UDP port numbers) to translate the
global address back to the correct local address. When multiple local addresses map to one global address, the TCP or UDP
port numbers of each inside host distinguish between the local addresses.
The figure below illustrates a NAT operation when an inside global address represents multiple inside local addresses. The
TCP port numbers act as differentiators.
The device performs the following process in the overloading of inside global addresses, as shown in the figure above. Both
Host B and Host C believe that they are communicating with a single host at address 203.0.113.2. Whereas, they are actually
communicating with different hosts; the port number is the differentiator. In fact, many inside hosts can share the inside
global IP address by using many port numbers.
The user at host 10.1.1.1:1723 opens a connection to Host B and the user at host 10.1.1.2:1723 opens a connection to Host
C.
NAT module intercepts the corresponding packets and attempts to translate the packets.
Based on the presence or absence of a matching NAT rule the following scenarios are possible:
If a matching static translation rule exists, then it takes precedence and the packets are translated to the corresponding
global address. Otherwise, the packets are matched against dynamic translation rule and in the event of a successful match,
they are translated to the corresponding global address. NAT module inserts a fully qualified flow entry corresponding to
the translated packets, into its translation database, to facilitate fast translation and forwarding of the packets corresponding
to this flow, in either direction.
The packets get forwarded without any address translation in the absence of a successful rule match.
The packets get dropped in the event of failure to obtain a valid inside global address even though we have a successful rule
match.
As this is a PAT configuration, transport ports help translate multiple flows to a single global address. (In addition to
source address, the source port is also subjected to translation and the associated flow entry maintains the corresponding
translation mappings.)
The device replaces inside local source address/port 10.1.1.1/1723 and 10.1.1.2/1723 with the corresponding selected global
address/port 203.0.113.2/1024 and 203.0.113.2/1723 respectively and forwards the packets.
Host B receives the packet and responds to host 10.1.1.1 by using the inside global IP address 203.0.113.2, on port 1024.
Host C receives the packet and responds to host 10.1.1.2 using the inside global IP adress 203.0.113.2, on port 1723.
