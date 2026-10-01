---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-ip-routing-b-ip-routing-m-ip6-route-ospf-79490805-1
title: "c-en-us-td-docs-routers-ios-config-17-x-ip-routing-b-ip-routing-m-ip6-route-ospf-79490805"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-ip-routing-b-ip-routing-m-ip6-route-ospf-79490805.md
source_anchor: ""
source_lines: [1, 83]
sha256: 3b43d2498c0a9caefd89ea7b52a58ca8f7a91a546731f1b47fbec591b12bc177
---

# c-en-us-td-docs-routers-ios-config-17-x-ip-routing-b-ip-routing-m-ip6-route-ospf-79490805

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
Open Shortest Path First version 3 (OSPFv3) is an IPv4 and IPv6 link-state routing protocol that supports IPv6 and IPv4 unicast
address families (AFs).
Complete the OSPFv3 network strategy and planning for your IPv6 network. For example, you must decide whether multiple areas
are required.
Enable IPv6 unicast routing.
Enable IPv6 on the interface.
Restrictions for IPv6 Routing: OSPFv3
When running a dual-stack IP network with OSPF version 2 for IPv4 and OSPFv3, be careful when changing the defaults for commands
used to enable OSPFv3. Changing these defaults may affect your OSPFv3 network, possibly adversely.
OSPFv3 is a routing protocol for IPv4 and IPv6. It is a link-state protocol, as opposed to a distance-vector protocol. Think
of a link as being an interface on a networking device. A link-state protocol makes its routing decisions based on the states
of the links that connect source and destination machines. The state of a link is a description of that interface and its
relationship to its neighboring networking devices. The interface information includes the IPv6 prefix of the interface, the
network mask, the type of network it is connected to, the devices connected to that network, and so on. This information is
propagated in various type of link-state advertisements (LSAs).
A device’s collection of LSA data is stored in a link-state database. The contents of the database, when subjected to the
Dijkstra algorithm, result in the creation of the OSPF routing table. The difference between the database and the routing
table is that the database contains a complete collection of raw data; the routing table contains a list of shortest paths
to known destinations via specific device interface ports.
OSPFv3, which is described in RFC 5340, supports IPv6 and IPv4 unicast AFs.
Comparison of OSPFv3 and OSPF Version 2
Much of OSPF version 3 is the same as in OSPF version 2. OSPFv3, which is described in RFC 5340, expands on OSPF version
2 to provide support for IPv6 routing prefixes and the larger size of IPv6 addresses.
In OSPFv3, a routing process does not need to be explicitly created. Enabling OSPFv3 on an interface will cause a routing
process, and its associated configuration, to be created.
In OSPFv3, each interface must be enabled using commands in interface configuration mode. This feature is different from
OSPF version 2, in which interfaces are indirectly enabled using the device configuration mode.
When using a nonbroadcast multiaccess (NBMA) interface in OSPFv3, you must manually configure the device with the list of
neighbors. Neighboring devices are identified by their device ID.
In IPv6, you can configure many address prefixes on an interface. In OSPFv3, all address prefixes on an interface are included
by default. You cannot select some address prefixes to be imported into OSPFv3; either all address prefixes on an interface
are imported, or no address prefixes on an interface are imported.
Unlike OSPF version 2, multiple instances of OSPFv3 can be run on a link.
OSPF automatically prefers a loopback interface over any other kind, and it chooses the highest IP address among all loopback
interfaces. If no loopback interfaces are present, the highest IP address in the device is chosen. You cannot tell OSPF to
use any particular interface.
LSA Types for OSPFv3
The following list describes LSA types, each of which has a different purpose:
Device LSAs (Type 1)—Describes the link state and costs of a device’s links to the area. These LSAs are flooded within an
area only. The LSA indicates if the device is an Area Border Router (ABR) or Autonomous System Boundary Router (ASBR), and
if it is one end of a virtual link. Type 1 LSAs are also used to advertise stub networks. In OSPFv3, these LSAs have no address
information and are network-protocol-independent. In OSPFv3, device interface information may be spread across multiple device
LSAs. Receivers must concatenate all device LSAs originated by a given device when running the SPF calculation.
Network LSAs (Type 2)—Describes the link-state and cost information for all devices attached to the network. This LSA is
an aggregation of all the link-state and cost information in the network. Only a designated device tracks this information
and can generate a network LSA. In OSPFv3, network LSAs have no address information and are network-protocol-independent.
Interarea-prefix LSAs for ABRs (Type 3)—Advertises internal networks to devices in other areas (interarea routes). Type 3
LSAs may represent a single network or a set of networks summarized into one advertisement. Only ABRs generate summary LSAs.
In OSPFv3, addresses for these LSAs are expressed as
prefix,
prefix length instead of
address,
mask. The default route is expressed as a prefix with length 0.
Interarea-device LSAs for ASBRs (Type 4)—Advertises the location of an ASBR. Devices that are trying to reach an external
network use these advertisements to determine the best path to the next hop. Type 4 LSAs are generated by ABRs on behalf of
ASBRs.
Autonomous system external LSAs (Type 5)—Redistributes routes from another autonomous system, usually from a different routing
protocol into OSPFv3. In OSPFv3, addresses for these LSAs are expressed as
prefix,
prefixlength instead of
address,
mask. The default route is expressed as a prefix with length 0.
Link LSAs (Type 8)—Have local-link flooding scope and are never flooded beyond the link with which they are associated. Link
LSAs provide the link-local address of the device to all other devices attached to the link, inform other devices attached
to the link of a list of prefixes to associate with the link, and allow the device to assert a collection of Options bits
to associate with the network LSA that will be originated for the link.
Intra-Area-Prefix LSAs (Type 9)—A device can originate multiple intra-area-prefix LSAs for each device or transit network,
each with a unique link-state ID. The link-state ID for each intra-area-prefix LSA describes its association to either the
device LSA or the network LSA and contains prefixes for stub and transit networks.
An address prefix occurs in almost all newly defined LSAs. The prefix is represented by three fields: PrefixLength, PrefixOptions,
and Address Prefix. In OSPFv3, addresses for these LSAs are expressed as
prefix,
prefixlength instead of
address,
mask. The default route is expressed as a prefix with length 0. Type 3 and Type 9 LSAs carry all prefix (subnet) information that,
in OSPFv2, is included in device LSAs and network LSAs. The Options field in certain LSAs (device LSAs, network LSAs, interarea-device
LSAs, and link LSAs) has been expanded to 24 bits to provide support for OSPFv3.
In OSPFv3, the sole function of the link-state ID in interarea-prefix LSAs, interarea-device LSAs, and autonomous-system
external LSAs is to identify individual pieces of the link-state database. All addresses or device IDs that are expressed
by the link-state ID in OSPF version 2 are carried in the body of the LSA in OSPFv3.
The link-state ID in network LSAs and link LSAs is always the interface ID of the originating device on the link being described.
For this reason, network LSAs and link LSAs are now the only LSAs whose size cannot be limited. A network LSA must list all
