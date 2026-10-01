---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-iproute-bfd-configuration-xe-16-9-irb-xe-16-9-book-b-f8d685b9
title: "c-en-us-td-docs-ios-xml-ios-iproute-bfd-configuration-xe-16-9-irb-xe-16-9-book-b-f8d685b9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-iproute-bfd-configuration-xe-16-9-irb-xe-16-9-book-b-f8d685b9.md
source_anchor: ""
source_lines: [1, 124]
sha256: d0ef3013e39381013d9603543ee8214cc0f2cd7d38d346448916a362f1af9644
---

# c-en-us-td-docs-ios-xml-ios-iproute-bfd-configuration-xe-16-9-irb-xe-16-9-book-b-f8d685b9

IP Routing: BFD Configuration Guide, Cisco IOS XE Fuji 16.9.x
Bias-Free Language
Bias-Free Language
The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
The BFD Multihop Support for IPv4 Static Routes feature enables detection of IPv4 network failure between paths that are
not directly connected. If a Bidirectional Forwarding Detection (BFD) session is up (that is, the next-hop destination is
reachable), IPv4 static routes that are associated with IPv4 static BFD configuration are added to a routing table. If the
BFD session is down, the routing table removes all associated static routes from the routing table.
This feature is applicable on different kinds of interfaces such as physical, subinterface, and virtual tunnels and across
intra-area and interarea topologies.
Your software release may not support all the features documented in this module. For the latest caveats and feature information,
see Bug Search Tool and the release notes for your platform and software release. To find information about the features documented in this module,
and to see a list of the releases in which each feature is supported, see the feature information table.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature
Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
Prerequisites for BFD Multihop Support for IPv4 Static Routes
The BFD destination for which an IPv4 static route has to be configured must be reachable by all devices.
The configured device must have at least one static route with the next-hop destination as a BFD destination for an associated
session. If not, the BFD session is not created on the device.
Information About BFD Multihop Support for IPv4 Static Routes
In Bidirectional Forwarding Detection for IPv4 (BFDv4) associated mode, an IPv4 static route is automatically associated
with an IPv4 static BFDv4 multihop destination address if the static route next hop exactly matches the static BFDv4 multihop
destination address.
The state of the BFDv4 session is used to determine whether the associated IPv4 static routes are added in the IPv4 routing
information base (RIB). For example, static routes are added in the IPv4 RIB only if the BFDv4 multihop destination is reachable,
and the static routes are removed from the IPv4 RIB if the BFDv4 multihop destination subsequently becomes unreachable.
BFDv4 Unassociated Mode
In Bidirectional Forwarding Detection for IPv4 (BFDv4), an IPv4 static BFD multihop destination can be configured in unassociated
mode. In unassociate mode, a BFD neighbor is not associated with a static route, and the BFD sessions are requested if the
IPv4 static BFD is configured.
Unassociated mode is useful in the following scenario:
Absence of an IPv4 static route—This scenario occurs when a static route is on device A, and device B is the next hop. In
associated mode, you must create both a static BFD multihop destination address and a static route on both devices to bring
up the BFDv4 session from device B to device A. Specifying the static BFD multihop destination in unassociated mode on device
B avoids the need to configure an unwanted static route.
How to Configure BFD Multihop Support for IPv4 Static Routes
Specify a BFD destination address which is same as the IPv4 static route next hop or gateway address.
Configure a BFD map and a BFD multihop template for an interface on the device. The destination address and source address
configured for a BFD map must match the BFD static multihop configuration and the source address must be a valid IP address
configured for an interface in the routing table.
SUMMARY STEPS
enable
configure terminal
ip route prefix mask ip-address
ip route static bfd multihop-destination-address multihop-source-address
ip route static bfd multihop-destination-address multihop-source-addressunassociate
end
DETAILED STEPS
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configure terminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
ip route prefix mask ip-address
Example:
Device(config)# ip route 192.0.2.0 255.255.255.0 10.1.1.2
Configures an IPv4 static route that BFD multihop uses to monitor static routes.
Step 4
ip route static bfd multihop-destination-address multihop-source-address
Example:
Device(config)# ip route static bfd 192.0.2.1 10.1.1.1
Configures the static IPv4 BFD multihop to be associated with a static IPv4 route.
Step 5
ip route static bfd multihop-destination-address multihop-source-addressunassociate
Example:
Device(config)# ip route static bfd 192.0.2.1 10.1.1.1 unassociate
(Optional) Configures the static IPv4 BFD multihop to be associated with a static IPv4 route in unassociated mode.
Step 6
end
Example:
Device(config)# end
Exits global configuration mode and returns to privileged EXEC mode.
Verifying BFD Multihop Support for IPv4 Static Routes
The following show commands can be used to verify IPv4 static routes for BFD multihop:
SUMMARY STEPS
show bfd neighbor
show ip static route bfd
DETAILED STEPS
Step 1
show bfd neighbor
Displays a line-by-line listing of existing BFD adjacencies.
Step 2
show ip static route bfd
Displays information about the IPv4 static BFD configured parameters.
Configuration Examples for BFD Multihop Support for IPv4 Static Routes
The Cisco Support website provides extensive online resources,
including documentation and tools for troubleshooting and
resolving technical issues with Cisco products and technologies.
To receive security and technical information about your
products, you can subscribe to various services, such as the
Product Alert Tool (accessed from Field Notices), the Cisco
Technical Services Newsletter, and Really Simple Syndication
(RSS) Feeds.
Access to most tools on the Cisco Support website requires a
Cisco.com user ID and password.
Feature Information for BFD Multihop Support for IPv4 Static Routes
The following table provides release information about the feature or features described in this module. This table lists
only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise,
subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco
Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
Table 1. Feature Information for BFD Multihop Support for IPv4 Static Routes
Feature Name
Releases
Feature Information
BFD Multihop Support for IPv4 Static Routes
Cisco IOS XE Release 3.9S
The BFD Multihop Support for IPv4 Static Routes feature enables detection of IPv4 network failure between paths that are not
directly connected. If a Bidirectional Forwarding Detection (BFD) session is up (that is, the next-hop destination is reachable),
IPv4 static routes that are associated with IPv4 static BFD configuration are added to a routing table. If the BFD session
is down, the routing table removes all associated static routes from the routing table.
The following commands were modified:
ip route static bfd
and
show ip static route bfd .
