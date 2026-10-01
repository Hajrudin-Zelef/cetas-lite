---
id: collect-261001-huawei/huawei/decoding-bgp-med-attribute-path-preference-1
title: "Understanding BGP MED and BGP Deterministic MED"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["reasoning"]
source: docs/RAG/collect-261001-huawei/decoding-bgp-med-attribute-path-preference.md
source_anchor: ""
source_lines: [1, 204]
sha256: f1f2c8ebafa4d2efbdd314d5370aa7b538a9a7c7684810314c9467fb86d10975
---

# Understanding BGP MED and BGP Deterministic MED

The BGP MED attribute, commonly referred to as the BGP metric, provides a means to convey to a neighboring Autonomous System (AS) a preferred entry point into the local AS. BGP MED is a non-transitive optional attribute and thus the receiving AS cannot propagate it across its AS borders. However, the receiving AS may reset the metric value upon receipt, if it so desires.

Previous versions of BGP (v2 and v3) defined this attribute as the inter-AS metric (INTER_AS_METRIC) but in BGPv4 it is defined as the multi-exit discriminator (MULTI_EXIT_DISC). The MED is an unsigned 32 bit integer.  The MED value can be any from 0 to 4,294,967,295 (2^32-1) with a lower value being preferred.  Certain implementations of BGP will treat a path with a MED value of 4,294,967,295 as infinite and hence the path would be deemed unusable so the MED value will be reset to 4,294,967,294.  This rewriting of the MED value could lead to inconsistencies, unintended path selections or even churn. I’ll do a follow up article on how BGP MED can possibly cause an endless convergence loop in certain topologies.

Cisco’s BGP implementation automatically assigns the value of the MED attribute based on the IGP metric value for any locally originate prefixes. The reasoning behind this is when there are multiple peering points with a neighboring AS the neighboring AS can use this metric to determine the best entry point into the local AS. This is the case when the originating AS’s network uses a single IGP. When multiple IGPs are used (i.e. OSPF and IS-IS) the metric value automatically copied into BGP will not be comparable. In this situation the metric values should be manually set before sending to the neighboring AS.

The MED value by default will only be used in Cisco’s BGP Best Path selection algorithm when comparing paths from the same AS.  If comparison is desired between different ASes the **bgp always-compare-med** router configuration command can be used. Use this command with caution as different ASes can have different policies regarding the setting of the MED value or in the case of the MED automatically being set they could be using different IGPs. Additionally by default MED is not compared between sub-autonomous systems in a BGP confederation. To enable comparison between different sub-ASes within a confederation use the **bgp bestpath med confed** router configuration command.

As mentioned by default the MED values are compared for paths from the same AS but this presents a problem in the way BGP path comparison is done in the IOS. Lets first examine how the path comparison is done to get a better understanding of the BGP Deterministic MED command and why Cisco recommends it to be enabled.

Here is the topology that we will use for this scenario:

We will primarily look at the effects of BGP MED on the BGP best path decision process from R1’s perspective. In this network AS 400 is advertising the 24.1.1.0/24 network. R2, R3 and R4 are in AS 200 with R5 being in AS 300. R2 is setting the MED for this network to 200, R3 to 300, R4 to 400 and R5 to 500 when the 24.1.1.0/24 network is advertised to R1. R1’s BGP configuration is below:

Rack1R1# **show run | sec router bgp 100**

router bgp 100

no synchronization

bgp router-id 1.1.1.1

neighbor 54.1.12.2 remote-as 200

neighbor 54.1.13.3 remote-as 200

neighbor 54.1.14.4 remote-as 200

neighbor 54.1.15.5 remote-as 300

no auto-summary

Rack1R1#

The output of the show ip bgp on R1:

Rack1R1#**show ip bgp**

BGP table version is 2, local router ID is 1.1.1.1

Status codes: s suppressed, d damped, h history, * valid, > best, i - internal,

r RIB-failure, S Stale

Origin codes: i - IGP, e - EGP, ? - incomplete

Network Next Hop Metric LocPrf Weight Path

* 24.1.1.0/24 54.1.12.2 200 0 200 400 ?

* 54.1.14.4 400 0 200 400 ?

* 54.1.15.5 500 0 300 400 ?

*> 54.1.13.3 300 0 200 400 ?

Rack1R1#

Now lets look at the 24.1.1.0/24 network in a little more detail:

Rack1R1#**show ip bgp 24.1.1.0/24**

BGP routing table entry for 24.1.1.0/24, version 2

Paths: (4 available, best #4, table Default-IP-Routing-Table)

Advertised to update-groups:

2

200 400

54.1.12.2 from 54.1.12.2 (2.2.2.2)

Origin incomplete, metric 200, localpref 100, valid, external

200 400

54.1.14.4 from 54.1.14.4 (4.4.4.4)

Origin incomplete, metric 400, localpref 100, valid, external

300 400

54.1.15.5 from 54.1.15.5 (5.5.5.5)

Origin incomplete, metric 500, localpref 100, valid, external

200 400

54.1.13.3 from 54.1.13.3 (3.3.3.3)

Origin incomplete, metric 300, localpref 100, valid, external, best

Rack1R1#

As we can see R1 has selected R3’s (3.3.3.3) advertisement of the 24.1.1.0/24 as the best path.  The MED is 300 for this advertisement which isn’t the lowest of all advertisements from AS 200.  The advertisement from R2 is actually lower as it has a MED value of 200.  Remember that the lower MED value is preferred since this value is normally copied from the IGP metric and with IGPs the lower metric value is preferred.  Since the MED attribute is optional it may not be present in all paths. By default, the BGP process will assume the MED value of zero for such paths, which will make them more preferred during the selection based on metric. If you want to change this behavior, use the **bgp bestpath med missing-as-worst** router configuration command.

Let's look at how R1 ended up selecting R3 as the best path.  First off the router will order the paths from the newest to the oldest.  By default all factors in the BGP best path decision process being the same, the oldest path will be selected as best.  BGP reduces the amount of churn in the routing table.  To change this behavior and not use the oldest path as the best, the BGP router-ID can be used to determine the best path.  To enable this use the **bgp bestpath compare-routerid** router configuration command.

Below the **bgp bestpath compare-routerid** command is enabled on R1.  Now R1 has selected R2’s path as the best since it has the lowest BGP router ID.

Rack1R1#**show run | sec router bgp**

router bgp 100

no synchronization

bgp router-id 1.1.1.1

bgp bestpath compare-routerid

neighbor 54.1.12.2 remote-as 200

neighbor 54.1.13.3 remote-as 200

neighbor 54.1.14.4 remote-as 200

neighbor 54.1.15.5 remote-as 300

no auto-summary

Rack1R1#**show ip bgp 24.1.1.0/24**

BGP routing table entry for 24.1.1.0/24, version 3

Paths: (4 available, best #1, table Default-IP-Routing-Table)

Flag: 0x10840

Advertised to update-groups:

2

200 400

54.1.12.2 from 54.1.12.2 (**2.2.2.2**)

Origin incomplete, metric 200, localpref 100, valid, external,**best**

200 400

54.1.14.4 from 54.1.14.4 (4.4.4.4)

Origin incomplete, metric 400, localpref 100, valid, external

300 400

54.1.15.5 from 54.1.15.5 (5.5.5.5)

Origin incomplete, metric 500, localpref 100, valid, external

200 400

54.1.13.3 from 54.1.13.3 (3.3.3.3)

Origin incomplete, metric 300, localpref 100, valid, external

Rack1R1#

The **bgp bestpath compare-routerid** command is removed for the remainder of this scenario.  When the command is removed R3 is once again selected as best.

Rack1R1#**show ip bgp 24.1.1.0/24**

BGP routing table entry for 24.1.1.0/24, version 4

Paths: (4 available, best #4, table Default-IP-Routing-Table)

Flag: 0x10840

Advertised to update-groups:

2

200 400

54.1.12.2 from 54.1.12.2 (2.2.2.2)

Origin incomplete, metric 200, localpref 100, valid, external

200 400

54.1.14.4 from 54.1.14.4 (4.4.4.4)

Origin incomplete, metric 400, localpref 100, valid, external

300 400

54.1.15.5 from 54.1.15.5 (5.5.5.5)

Origin incomplete, metric 500, localpref 100, valid, external

200 400

54.1.13.3 from 54.1.13.3 (**3.3.3.3**)

Origin incomplete, metric 300, localpref 100, valid, external,**best**

Rack1R1#

Additional RFC 4277 (Experience with the BGP-4 Protocol) mentions the following in regards to selecting a path based upon the oldest path.

