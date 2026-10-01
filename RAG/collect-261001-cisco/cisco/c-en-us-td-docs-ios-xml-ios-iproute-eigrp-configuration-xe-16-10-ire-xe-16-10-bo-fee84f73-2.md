---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-iproute-eigrp-configuration-xe-16-10-ire-xe-16-10-bo-fee84f73-2
title: "c-en-us-td-docs-ios-xml-ios-iproute-eigrp-configuration-xe-16-10-ire-xe-16-10-bo-fee84f73"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-iproute-eigrp-configuration-xe-16-10-ire-xe-16-10-bo-fee84f73.md
source_anchor: ""
source_lines: [100, 129]
sha256: 4f25cbf94789da98018f437df151add83143a47678813894b31908daeb229a96
---

# c-en-us-td-docs-ios-xml-ios-iproute-eigrp-configuration-xe-16-10-ire-xe-16-10-bo-fee84f73

The following example shows you how to configure the customer edge (CE) device in the network to advertise local routes to
the Enhanced Interior Gateway Routing Protocol (EIGRP) Route Reflectors.
The following table provides release information about the feature or features described in this module. This table lists
only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise,
subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco
Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
Table 1. Feature Information for
EIGRP Over the Top
Feature
Name
Releases
Feature
Information
EIGRP Over
the Top
The EIGRP
Over the Top feature enables a single end-to-end routing domain between two or
(EIGRP) more Enhanced Interior Gateway Routing Protocol sites that are
connected using a private or public WAN connection. EIGRP OTP also supports the
propagation of SGT over L3 network.
The
following commands were introduced or modified:
remote-neighbor (EIGRP),
neighbor
(EIGRP),
cts propagate sgt
,and
show ip eigrp
neighbors .
