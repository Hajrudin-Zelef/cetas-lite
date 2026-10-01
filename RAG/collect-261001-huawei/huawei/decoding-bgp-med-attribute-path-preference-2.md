---
id: collect-261001-huawei/huawei/decoding-bgp-med-attribute-path-preference-2
title: "Understanding BGP MED and BGP Deterministic MED"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["reasoning"]
source: docs/RAG/collect-261001-huawei/decoding-bgp-med-attribute-path-preference.md
source_anchor: ""
source_lines: [205, 458]
sha256: f299b569cd3928f4b0042eae5372bf31c69316132d939e609ad5f5a6b0e0eff5
---

# Understanding BGP MED and BGP Deterministic MED

7.1.4.  MEDs and Temporal Route Selection

Some implementations have hooks to apply temporal behavior in MED-based best path selection. That is, all things being equal up to MED consideration, preference would be applied to the "oldest" path, without preference for the lower MED value. The reasoning for this is that "older" paths are presumably more stable, and thus preferable. However, temporal behavior in route selection results in non-deterministic behavior, and as such, may often be undesirable.


Rack1R1#**show ip bgp 24.1.1.0/24**

BGP routing table entry for 24.1.1.0/24, version 4

Paths: (4 available, best #4, table Default-IP-Routing-Table)

Flag: 0x820

Advertised to update-groups:

2

200 400

54.1.12.2 from 54.1.12.2 (2.2.2.2)

Origin incomplete, metric 200, localpref 100, valid, external

200 400

54.1.13.3 from 54.1.13.3 (3.3.3.3)

Origin incomplete, metric 300, localpref 100, valid, external

200 400

54.1.14.4 from 54.1.14.4 (4.4.4.4)

Origin incomplete, metric 400, localpref 100, valid, external

300 400

54.1.15.5 from 54.1.15.5 (**5.5.5.5**)

Origin incomplete, metric 500, localpref 100, valid, external,**best**

Rack1R1#

First off it’s important to understand that the paths are compared in pairs starting with the newest path and comparing it with the second newest.  The winning path between the first and second is then compared to the third and in our case the winner of that comparison is finally compared with the fourth and final path.  On R1 for the 24.1.1.0/24 network, R2’s and R3’s paths are compared first.  Everything in the BGP best path decision algorithm is the same down to MED (weight, local preference, AS path, etc).  Since the advertisements by R2 and R3 are in the same AS the MED is compared and R2 wins since it has a MED of 200 as opposed to R3’s MED of 300.  Next R2 is then compared to the third oldest entry which is R4’s.  R2 and R4 are in the same AS so R2 wins based upon the lower MED value. Finally R2 is compared with R5. Everything is equal but the MED, router ID and age of the advertisement.  Since R2 and R5 are in different ASes and the **bgp always-compare-med** isn’t enabled, MED isn’t compared.  Additionally we do not have **bgp bestpath compare-routerid** enabled which leads the R1 to select the oldest advertisement.  Since R5 is listed below R2 we know that it is older and in turn wins out due to being the older advertisement and is installed as the best path to reach the 24.1.1.0/24 network.

As we can see the MED comparison between the paths advertised by AS 200 did not happen as intended by AS 200. AS 200 was setting the MED so that AS 100 will use R2 as the ingress point into AS 200. This is only because R5’s advertisement was second to the oldest that in turn broke the MED comparison between the AS 200 routers (R2, R3 and R4).

Ideally we want the MED compared between advertisements from the same AS irrespective of their age.  This is where the **bgp deterministic-med** router configuration command is useful.  When this command is enabled the router will group all paths from the same AS and compare them together before comparing them to paths from different ASes.  Lets enable the command on R1.  We should see that R2 is selected as the preferred path between R2, R3 and R4 but this will mean that once R2 is compared to R5, R5 will be installed since it is an older advertisement.

Rack1R1#**show run | sec router bgp**

router bgp 100

no synchronization

bgp router-id 1.1.1.1

bgp log-neighbor-changes

**bgp deterministic-med**

neighbor 54.1.12.2 remote-as 200

neighbor 54.1.13.3 remote-as 200

neighbor 54.1.14.4 remote-as 200

neighbor 54.1.15.5 remote-as 300

no auto-summary

Rack1R1#**show ip bgp 24.1.1.0/24**

BGP routing table entry for 24.1.1.0/24, version 5

Paths: (4 available, best #4, table Default-IP-Routing-Table)

Flag: 0x820

Advertised to update-groups:

2

200 400

54.1.12.2 from 54.1.12.2 (2.2.2.2)

Origin incomplete, metric 200, localpref 100, valid, external

200 400

54.1.13.3 from 54.1.13.3 (3.3.3.3)

Origin incomplete, metric 300, localpref 100, valid, external

200 400

54.1.14.4 from 54.1.14.4 (4.4.4.4)

Origin incomplete, metric 400, localpref 100, valid, external

300 400

54.1.15.5 from 54.1.15.5 (**5.5.5.5**)

Origin incomplete, metric 500, localpref 100, valid, external,**best**

Rack1R1#

If we want to have R2 selected as best we can clear the BGP neighbor relationship with R5 which will in turn cause R5’s paths to be cleared out. Once the neighbor relationship with R5 comes back up and R5 advertised the 24.1.1.0/24 path, it will be the newest advertisement and in turn be listed at the top.

Rack1R1#**clear ip bgp 54.1.15.5**

Rack1R1#

%BGP-5-ADJCHANGE: neighbor 54.1.15.5 Down User reset

Rack1R1#

%BGP-5-ADJCHANGE: neighbor 54.1.15.5 Up

Rack1R1#

Now as expected R2 was finally selected as the best path.

Rack1R1#**show ip bgp 24.1.1.0**

BGP routing table entry for 24.1.1.0/24, version 6

Paths: (4 available, best #2, table Default-IP-Routing-Table)

Flag: 0x820

Advertised to update-groups:

2

300 400

54.1.15.5 from 54.1.15.5 (5.5.5.5)

Origin incomplete, metric 500, localpref 100, valid, external

200 400

54.1.12.2 from 54.1.12.2 (**2.2.2.2**)

Origin incomplete, metric 200, localpref 100, valid, external,**best**

200 400

54.1.13.3 from 54.1.13.3 (3.3.3.3)

Origin incomplete, metric 300, localpref 100, valid, external

200 400

54.1.14.4 from 54.1.14.4 (4.4.4.4)

Origin incomplete, metric 400, localpref 100, valid, external

Rack1R1#

Of course to always ensure R2 is selected in our network as the best path we could also use the **bgp always-compare-med command** to compare MED between different ASes but this command is normally not used in the real world unless MED policies are standardized between neighboring ASes.

Rack1R1#**show run | sec router bgp**

router bgp 100

no synchronization

bgp router-id 1.1.1.1

**bgp always-compare-med**

bgp deterministic-med

neighbor 54.1.12.2 remote-as 200

neighbor 54.1.13.3 remote-as 200

neighbor 54.1.14.4 remote-as 200

neighbor 54.1.15.5 remote-as 300

no auto-summary

Rack1R1#

Rack1R1#**clear ip bgp ***

%BGP-5-ADJCHANGE: neighbor 54.1.12.2 Down User reset

%BGP-5-ADJCHANGE: neighbor 54.1.13.3 Down User reset

%BGP-5-ADJCHANGE: neighbor 54.1.14.4 Down User reset

%BGP-5-ADJCHANGE: neighbor 54.1.15.5 Down User reset

Rack1R1#

%BGP-5-ADJCHANGE: neighbor 54.1.12.2 Up

%BGP-5-ADJCHANGE: neighbor 54.1.13.3 Up

%BGP-5-ADJCHANGE: neighbor 54.1.14.4 Up

%BGP-5-ADJCHANGE: neighbor 54.1.15.5 Up

Rack1R1#**show ip bgp 24.1.1.0** 

BGP routing table entry for 24.1.1.0/24, version 4

Paths: (4 available, best #2, table Default-IP-Routing-Table)

Flag: 0x10860

Advertised to update-groups:

2

300 400

54.1.15.5 from 54.1.15.5 (5.5.5.5)

Origin incomplete, metric 500, localpref 100, valid, external

200 400

54.1.12.2 from 54.1.12.2 (**2.2.2.2**)

Origin incomplete, metric 200, localpref 100, valid, external,**best**

200 400

54.1.13.3 from 54.1.13.3 (3.3.3.3)

Origin incomplete, metric 300, localpref 100, valid, external

200 400

54.1.14.4 from 54.1.14.4 (4.4.4.4)

Origin incomplete, metric 400, localpref 100, valid, external

Rack1R1#

If BGP Deterministic MED is used, it should be enabled on all BGP speaking devices within an AS to ensure a consistent policy regarding the use of MEDs.

We should now have a better understanding of how MED is used in the BGP route selection process and the BGP route selection process is general.

My next post will be in regards to the Two Rate Three Color Marker (trTCM) as defined in RFC 2698 and implemented in the Cisco IOS. Also I hope to see many of you in my new RS Bootcamps.

Read more related articles to BGP Metric:
