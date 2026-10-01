---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-20
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "parameters", "preemption"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [2593, 2753]
sha256: e53b4ff0c7d2ed27bc1000eb3aeae835470417df050bfc882daf65ca3a116733
---

# A line starting with the # sign is comments.

The process of establishing a new path for tunnel 0/0/1 is as follows:
1. After MPLS TE path calculation, Path messages are forwarded over the path LSRA
-> LSRB -> LSRF -> LSRE and Resv messages are forwarded over the path LSRE -
> LSRF -> LSRB -> LSRA.
2. After receiving Resv messages from LSRF, LSRB triggers preemption when it detects
that bandwidth is insufficient for resource reservation. The preemption process differs
in two preemption modes.
– In hard preemption, LSRB directly tears down Path2 for tunnel 0/0/2 because the
priority of tunnel 0/0/1 is higher than that of tunnel 0/0/2. LSRB sends a PathTear
message to LSRF, requiring LSRF to remove Path2 information. In addition,
LSRB sends a ResvTear message to LSRC, requiring LSRC to delete the node
reservation state. In this case, some traffic on tunnel 0/0/2 is lost.
– In soft preemption, LSRB sends a ResvTear message to LSRC and establishes a
new path Path4 on the condition that LSRB and LSRC do not tear down Path2.
After Path4 is established and traffic is switched to it, LSRB tears down Path2 of
tunnel 0/0/2.
l Route pinning
Any changes in the network topology or tunnel attributes may cause an established CR-
LSP to be reestablished, leading to the following issues:
– The reestablished CR-LSP may be over a path that is different from the original one,
causing management difficulties.
– Traffic must switch from the original CR-LSP to the new one, causing traffic loss.
Route pinning can be used to resolve the preceding problems. Route pinning helps an
established CR-LSP remain over a path regardless of route changes. This function improves
service traffic continuity and reliability.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
53

l Affinity attribute
An affinity is a 32-bit vector, configured on the ingress of a tunnel. It must be used together
with a link administrative group attribute.
After a tunnel is configured with an affinity, a device compares the affinity with the
administrative group value during link selection to determine whether a link with specified
attributes is selected or not. The device implements two AND operations, one between a
32-bit mask and each affinity, and one between the 32-bit mask and the administrative
group value. If the two AND operations yield the same results, the path is selected. If the
results are different, the path is not selected. The following rules apply:
– If some bits in a mask are 1s, at least one bit in the administrative group is 1 and the
corresponding bit in the affinity must be 1. If some bits in the affinity are 0s, the
corresponding bits in the administrative group cannot be 1.
For example, an affinity is 0x0000FFFF and its mask is 0xFFFFFFFF. The higher-order
16 bits in the administrative group of available links are 0 and at least one of the lower-
order 16 bits is 1. This means the administrative group attribute ranges from 0x00000001
to 0x0000FFFF.
– If some bits in a mask are 0s, the corresponding bits in the administrative group are not
compared with the affinity bits.
For example, an affinity is 0xFFFFFFFF and its mask is 0xFFFF0000. At least one of
the higher-order 16 bits in an administrative group attribute is 1 and the lower-order 16
bits can be 0s and 1s. This means that the administrative group attribute ranges from
0x00010000 to 0xFFFFFFFF.
NOTE
Understand specific comparison rules before deploying devices of different vendors because the
comparison rules vary with the vendor.
A network administrator can use the link administrative group and affinities to control the
paths over which MPLS TE tunnels are established.
l Hop limit
Hop limit is a condition for path selection during CR-LSP establishment. Similar to the
administrative group and affinity attributes, a hop limit defines the number of hops that a
CR-LSP allows.
3.2.2 Implementation
Figure 3-8 shows the MPLS TE implementation framework.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
54

Figure 3-8 MPLS TE implementation framework
Traffic forwarding
Incoming 
packets
Outgoing 
packets
IS-IS/OSPF routing
Information 
advertisement
IGP route 
selection
LSDB
LSP route 
selection
TEDB Signaling 
protocol
Path 
establishment
Upstream 
nodes
Downstream 
nodesLocal nodes
Protocol packet exchanging
Data packet forwarding
Internal information 
processsing
Path 
establishment
Information 
advertisement
Information advertisement, path calculation, path establishment, and traffic forwarding are used
for establishing an MPLS TE tunnel. The information advertisement function is used to collect
TE related information using IGP. The path calculation function is used for calculating paths
based on information collected. The path establishment function establishes paths by exchanging
packets between upstream and downstream nodes using signaling protocols. The traffic
forwarding function imports data packets to the MPLS TE tunnel and forwards the packets.
Table 3-1 details the functions.
Table 3-1 MPLS TE implementation process
N
o.
Function Description
1 Informatio
n
Advertise
ment
Extends an IGP to advertise TE information, in addition to routing
information. TE information includes the maximum link bandwidth,
maximum reservable bandwidth, reserved bandwidth, and link colors.
Every node collects TE information about all nodes in a local area and
generates a traffic engineering database (TEDB).
2 Path
Calculatio
n
Runs Constraint Shortest Path First (CSPF) and uses TEDB data to
calculate a path that satisfies specific constraints. CSPF evolves from the
Shortest Path First (SPF) protocol. CSPF excludes nodes and links that
do not satisfy specific constraints and uses the same algorithm that SPF
supports to calculate a path.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
55

N
o.
Function Description
3 Path
Establish
ment
Establishes the following types of CR-LSPs:
l Static CR-LSP: set up by manually configuring labels and bandwidth,
irrespective of signaling protocols or path calculation.
Setting up a static CR-LSP consumes few resources because no MPLS
control packets are exchanged between two ends of the CR-LSP. The
static CR-LSP cannot be adjusted dynamically in a changing network
topology; therefore, the static CR-LSP is not widely used.
l Dynamic CR-LSP: set up using RSVP-TE signaling. RSVP-TE
carries parameters, such as the tunnel bandwidth, explicit path, and
affinities.
There is no need to manually configure each hop along a dynamic
CR-LSP. Dynamic CR-LSPs apply to large-scale networks.
You can use the RSVP authentication mechanism to improve
security and reliability during path establishment.
4 Traffic
Forwardin
g
Imports traffic to the MPLS TE tunnel and forwards traffic through the
tunnel. The first three functions help establish an MPLS TE tunnel. This
function forwards traffic after traffic is imported to the tunnel.
 
