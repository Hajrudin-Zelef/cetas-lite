---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-ip-routing-b-ip-routing-m-ip6-route-ospf-79490805-2
title: "c-en-us-td-docs-routers-ios-config-17-x-ip-routing-b-ip-routing-m-ip6-route-ospf-79490805"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-ip-routing-b-ip-routing-m-ip6-route-ospf-79490805.md
source_anchor: ""
source_lines: [84, 174]
sha256: 3f701e0ed8028d568ae630cb7e7e91fd4212918fc70fcaf4bb0e2db467807a20
---

# c-en-us-td-docs-routers-ios-config-17-x-ip-routing-b-ip-routing-m-ip6-route-ospf-79490805

devices connected to the link, and a link LSA must list all of the address prefixes of a device on the link.
Load Balancing in OSPFv3
When a device learns multiple routes to a specific network via multiple routing processes (or routing protocols), it installs
the route with the lowest administrative distance in the routing table. Sometimes the device must select a route from among
many learned via the same routing process with the same administrative distance. In this case, the device chooses the path
with the lowest cost (or metric) to the destination. Each routing process calculates its cost differently and the costs may
need to be manipulated in order to achieve load balancing.
OSPFv3 performs load balancing automatically in the following way. If OSPFv3 finds that it can reach a destination through
more than one interface and each path has the same cost, it installs each path in the routing table. The only restriction
on the number of paths to the same destination is controlled by the
maximum-paths command. The default maximum paths is 16, and the range is from 1 to 64.
Addresses Imported into OSPFv3
When importing the set of addresses specified on an interface on which OSPFv3 is running into OSPFv3, you cannot select specific
addresses to be imported. Either all addresses are imported, or no addresses are imported.
OSPFv3 Customization
You can customize OSPFv3 for your network, but you likely will not need to do so. The defaults for OSPFv3 are set to meet
the requirements of most customers and features. If you must change the defaults, refer to the IPv6 command reference to find
the appropriate syntax.
Caution
Be careful when changing the defaults. Changing defaults will affect your OSPFv3 network, possibly adversely.
Force SPF in OSPFv3
When the
process keyword is used with the
clearipv6ospf command, the OSPFv3 database is cleared and repopulated, and then the SPF algorithm is performed. When the
force-spf keyword is used with the
clearipv6ospf command, the OSPFv3 database is not cleared before the SPF algorithm is performed.
Once you have
completed step 3 and entered OSPFv3 router configuration mode, you can perform
any of the subsequent steps in this task as needed to configure OSPFv3 Device
configuration.
Enters router
configuration mode for the IPv4 or IPv6 address family.
Note
The previous syntax for enabling OSPF for IPv6 was:
ipv6 router ospf process-id
The new recommended syntax is:
router ospfv3process-id
It is recommended to use the new syntax for configuring OSPF for both IPv4 and IPv6. When migrating from older configurations,
enter the new router ospfv3process-id command in global configuration mode. This automatically converts all existing OSPFv3 configurations at global and interface-level,
including ipv6 ospf commands to the updated CLI syntax. You do not have to manually reconfigure each OSPFv3 enabled interface.
Device# show ospfv3 database
OSPFv3 Device with ID (172.16.4.4) (Process ID 1)
Device Link States (Area 0)
ADV Device Age Seq# Fragment ID Link count Bits
172.16.4.4 239 0x80000003 0 1 B
172.16.6.6 239 0x80000003 0 1 B
Inter Area Prefix Link States (Area 0)
ADV Device Age Seq# Prefix
172.16.4.4 249 0x80000001 FEC0:3344::/32
172.16.4.4 219 0x80000001 FEC0:3366::/32
172.16.6.6 247 0x80000001 FEC0:3366::/32
172.16.6.6 193 0x80000001 FEC0:3344::/32
172.16.6.6 82 0x80000001 FEC0::/32
Inter Area Device Link States (Area 0)
ADV Device Age Seq# Link ID Dest DevID
172.16.4.4 219 0x80000001 50529027 172.16.3.3
172.16.6.6 193 0x80000001 50529027 172.16.3.3
Link (Type-8) Link States (Area 0)
ADV Device Age Seq# Link ID Interface
172.16.4.4 242 0x80000002 14 PO4/0
172.16.6.6 252 0x80000002 14 PO4/0
Intra Area Prefix Link States (Area 0)
ADV Device Age Seq# Link ID Ref-lstype Ref-LSID
172.16.4.4 242 0x80000002 0 0x2001 0
172.16.6.6 252 0x80000002 0 0x2001 0
Device# show ospfv3 neighbor
OSPFv3 Device with ID (10.1.1.1) (Process ID 42)
Neighbor ID Pri State Dead Time Interface ID Interface
10.4.4.4 1 FULL/ - 00:00:39 12 vm1
OSPFv3 Device with ID (10.2.1.1) (Process ID 100)
Neighbor ID Pri State Dead Time Interface ID Interface
10.5.4.4 1 FULL/ - 00:00:35 12 vm1
Example: Forcing SPF Configuration
The following example shows how to trigger SPF to redo the SPF and repopulate the routing tables:
The Cisco Support and Documentation website provides online resources to download documentation, software, and tools. Use
these resources to install and configure the software and to troubleshoot and resolve technical issues with Cisco products
and technologies. Access to most tools on the Cisco Support and Documentation website requires a Cisco.com user ID and password.
The following table provides release information about the feature or features described in this module. This table lists
only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise,
subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature
Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
Table 1. Feature Information for IPv6 Routing: OSPFv3
Feature Name
Releases
Feature Information
IPv6 Routing: OSPFv3
Cisco IOS Release 15.2(6)E
OSPF version 3 for IPv6 expands on OSPF version 2 to provide support for IPv6 routing prefixes and the larger size of IPv6
addresses.
Table 2. Feature Information for IPv6 Routing: OSPFv3
