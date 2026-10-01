---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-23
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "distribution", "parameters"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [3062, 3209]
sha256: 7eb6ddefe49be7f516bcea81f251a51f106fc7f73d3cb7340629c1efcd0fe187
---

# A line starting with the # sign is comments.

– The proportion of the bandwidth reserved for an MPLS TE tunnel to the available
bandwidth in the TEDB is greater than or equal to a specific threshold.
– The proportion of the bandwidth released by an MPLS TE tunnel to the available
bandwidth in the TEDB is greater than or equal to a specific threshold.
If either of the preceding conditions is met, an IGP floods link bandwidth information, and
the device updates the TEDB.
For example, the available bandwidth of a link is 100 Mbit/s and 100 TE tunnels, each with
bandwidth of 1 Mbit/s, are established over the link. The flooding threshold is 10%. The
Figure 3-13 shows the proportion of the bandwidth reserved for each MPLS TE tunnel to
the available bandwidth in the TEDB.
Bandwidth flooding is not performed when tunnels 1 to 9 are created. After tunnel 10 is
created, the bandwidth information (10 Mbit/s in total) on tunnels 1 to 10 is flooded. The
available bandwidth is 90 Mbit/s. Similarly, no bandwidth information is flooded after
tunnels 11 to 18 are created. After tunnel 19 is created, bandwidth information on tunnels
11 to 19 is flooded. The process repeats until tunnel 100 is established.
Figure 3-13 Proportion of the bandwidth reserved for each MPLS TE tunnel to the available
bandwidth in the TEDB
Second flooding
Available bandwidth
 80Mbit/s
First flooding
Available bandiwdth
 90Mbit/s
Original available 
bandwidth 100Mbit/s
2%
3%
4%
5%
6%
7%
8%
9%
10%
1.1%
2.2%
3.3%
4.4%
5.6%
6.7%
7.8%
8.9%
1.3%
2.5%
3.8%
10%
1%
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22......
 
Results Obtained After Information Advertisement
Every node creates a TEDB in an MPLS TE area after OSPF TE or IS-IS TE floods bandwidth
information.
After MPLS TE is deployed on a network, related resource information needs to be advertised
to each node. Each node collects information about link constraints and bandwidth usage in the
local area to form a database covering network link attributes and topology attributes. This
database is called TE Database (TEDB).
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
63

A node calculates the optimal path to another node in the MPLS TE area based on information
in the TEDB. MPLS TE then establishes a CR-LSP over this optimal path.
The TEDB and IGP link-state data base (LSDB) are independent of each other. Both TEDB and
LSDB contain information flooded by IGP, but they have different content and functions. In
addition to information contained in the LSDB, TEDB contains TE information. An IGP uses
information in an LSDB to calculate the shortest path, while MPLS TE uses information in a
TEDB to calculate the optimal path.
3.2.4 Path Calculation
MPLS TE uses constrained shortest path first (CSPF) to calculate the optimal path to a specified
node. CSPF, which is derived from SPF, is an algorithm that supports constraints.
CSPF Fundamentals
CSPF works based on the following parameters:
l Bandwidth of an LSP tunnel to be established, explicit path, setup priority, hold priority,
and affinity attribute, which are configured on the ingress node of a tunnel
l Traffic engineering database (TEDB)
NOTE
A TEDB can be generated only after Interior Gateway Protocol (IGP) TE is configured. On an IGP TE-
incapable network, CR-LSPs are established based on IGP routes, but not CSPF calculation results.
CSPF Calculation Process
CSPF checks the constraints for establishing an LSP to exclude the links that do not meet tunnel
attribute requirements in the TEDB, and then calculates the shortest path to the destination of
the tunnel by using SPF.
NOTE
When both OSPF TE and IS-IS TE are configured, CSPF attempts to use the OSPF TEDB to establish a
path for a CR-LSP. If a path is successfully calculated using OSPF TEDB information, CSPF completes
calculation and does not use the IS-IS TEDB to calculate a path. If path calculation fails, CSPF attempts
to use IS-IS TEDB information to calculate a path.
CSPF can be configured to use the IS-IS TEDB to calculate a CR-LSP path. If path calculation fails, CSPF
uses the OSPF TEDB to calculate a path.
CSPF calculates the shortest path to a destination. If there are several shortest paths with the
same metric, CSPF uses a tie-breaking policy to select one of them. The following tie-breaking
policies for selecting a path are available:
l Most-fill: selects a link with the highest proportion of used bandwidth to the maximum
reservable bandwidth, efficiently using bandwidth resources.
l Least-fill: selects a link with the lowest proportion of used bandwidth to the maximum
reservable bandwidth, evenly using bandwidth resources among links.
l Random: selects links randomly, allowing LSPs to be established evenly over links,
regardless of bandwidth distribution.
When several links have the same proportion of used bandwidth to the maximum reservable
bandwidth, the link discovered first is selected, irrespective of whether most-fill or least-fill is
configured.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
64

The network shown in Figure 3-14 illustrates the CSPF path calculation process. CSPF removes
links marked blue and links each with bandwidth of 50 Mbit/s based on tunnel constraints and
uses other links each with bandwidth of 100 Mbit/s to calculate a path for an MPLS TE tunnel
on the network shown in Figure 3-14. The constraints include the destination LSRE, bandwidth
of 80 Mbit/s, and a transit node LSRH.
Figure 3-14 Process of link removal
LSRA
LSRB LSRC LSRD
LSRE
LSRF LSRG LSRH
50
50
50
LSRA
LSRC LSRD
LSRF LSRG LSRH
LSRE
MPLS TE Tunnel 1/0/0:
Destination = LSRE
Bandwidth = 80 Mbit/s
Affinity = Black
LSRH Loose
Calculated topology
Blue Blue
Blue
CSPF calculates a path shown in Figure 3-15 in the same way SPF would calculate it.
Figure 3-15 CSPF calculation result
LSRD
LSRF LSRG LSRH
LSRELSRA
Differences Between CSPF and SPF
CSPF is dedicated to calculating MPLS TE paths. It has similarities with SPF but they have the
following differences:
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
65

