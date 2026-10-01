---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-19
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "cost", "distribution", "preemption"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [2415, 2592]
sha256: d21a357a35871f6836ac1b833e8865de02db81982f0abf6ce825d3787a2579bd
---

# A line starting with the # sign is comments.

Link Attributes
MPLS TE link attributes describe bandwidth resources, route costs, and link reliability. The link
attributes are as follows:
l Total link bandwidth: is physical link bandwidth.
l Maximum reservable bandwidth: is the maximum bandwidth that a link can reserve for an
MPLS TE tunnel to be established. The maximum reservable bandwidth must be lower
than or equal to the total link bandwidth.
l TE metric: is the TE cost of a link. To better control path calculation for an MPLS TE
tunnel, MPLS TE provides the TE metric so that an IGP route can be selected independently.
By default, a link uses the IGP metric as the TE metric.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
49

l SRLG: is a set of links which are likely to fail concurrently when sharing a physical resource
(for example, an optical fiber). Links in an SRLG share the same risk of faults. If one link
fails, other links in the SRLG also fail.
An SRLG enhances CR-LSP reliability on an MPLS TE network enabled with CR-LSP
hot standby or TE FRR. For more information about the SRLG, see SRLG.
l Link administrative group: is also called link color. A link administrative group is a 32-bit
vector, with each bit set to a specified value that is associated with a desired meaning. For
example, a link administrative group attribute can be configured to describe link bandwidth,
a performance parameter or a management policy. The policy can be a traffic type (multicast
for example) or a flag indicating that an MPLS TE tunnel passes over the link. The link
administrative group attribute is used together with affinities to control the paths for
tunnels.
Tunnel Attributes
LSPs in an MPLS TE tunnel are constraint-based routed LSPs (CR-LSPs). These constraints are
tunnel attributes.
Unlike Label Distribution Protocol (LDP) LSPs that are established using routing information,
CR-LSPs are established based on bandwidth and path constraints in addition to routing
information:
l Bandwidth constraint: is the tunnel bandwidth.
l Path constraint: includes the explicit path, priority and preemption, route pinning, affinity
attribute, and hop limit.
The mechanism for establishing and managing these constraints is called Constraint-based
Routing (CR). The device supports the following CRs:
l Tunnel bandwidth
the values are planned based on services that are to pass through a tunnel. The configured
bandwidth is reserved on each node through which a tunnel passes.
l Explicit path
An explicit path is used to establish a CR-LSP. Nodes to be included or excluded are
specified on this path. Explicit paths are classified into the following types:
– Strict explicit path
A strict explicit path includes specified nodes through which a CR-LSP must pass. The
next hop must be directly connected to the previous hop. By specifying a strict explicit
path, the most accurate path is provided for a CR-LSP.
Figure 3-4 Strict explicit path
Strict explicit path
LSRB LSRD
LSRC LSRE
LSRF LSRA
 Explicit path
 LSRB Strict
 LSRC Strict
 LSRE Strict
 LSRD Strict
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
50

For example, a CR-LSP is set up between LSRA and LSRF on the network shown in
Figure 3-4. LSRA is the ingress, and LSRF is the egress. "X Strict" specifies the LSR
that the CR-LSP must travel through. For example, "LSRB Strict" indicates that the
CR-LSP must travel through LSRB, and the previous hop of LSRB must be LSRA.
"LSRC Strict" indicates that the CR-LSP must travel through LSRC, and the previous
hop of LSRC must be LSRB. The procedure repeats. A path with each node specified
is provided for the CR-LSP.
– Loose explicit path
A loose explicit path contains specified nodes through which a CR-LSP must pass. Other
routers that are not specified can also exist on the CR-LSP.
Figure 3-5 Loose explicit path
Loose explicit path
LSRB LSRD
LSRC LSRE
LSRF LSRA
 Explicit path
 LSRD Loose
 
For example, a CR-LSP is set up over a loose explicit path between LSRA and LSRF
on the network shown in Figure 3-5. LSRA is the ingress, and LSRF is the egress.
"LSRD Loose" indicates that the CR-LSP must pass through LSRD and LSRD and
LSRA may not be directly connected. This means that other LSRs may exist between
LSRD and LSRA.
l Priorities and preemption
They are used to allow TE tunnels to be established preferentially to transmit important
services, preventing random resource competition during tunnel establishment.
CR-LSPs use setup and holding priorities to determine whether to preempt resources. A
new CR-LSP and an established CR-LSP compete for resources by comparing the priorities.
The new path can succeed in preemption when its setup priority is higher than the holding
priority of the established path. The priority value ranges from 0 to 7. A smaller value
allows for a higher priority. The setup priority must be lower than or equal to the holding
priority for a tunnel.
If there is no path meeting the bandwidth requirement of a desired CR-LSP, a device can
tear down an established CR-LSP and use the bandwidth assigned to that CR-LSP to
establish a desired CR-LSP. This is called preemption. The following preemption modes
are supported:
– Hard preemption: A CR-LSP with a higher priority can directly preempt resources
assigned to a CR-LSP with a lower priority. Some traffic is dropped on the CR-LSP
with a lower priority during the hard preemption process.
– Soft preemption: The Make-Before-Break mechanism applies. A CR-LSP with a
higher priority has to wait until traffic over a lower-priority CR-LSP switches to another
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
51

CR-LSP before the higher-priority CR-LSP preempts bandwidth assigned to the lower-
priority CR-LSP.
The priority and preemption attributes are used in conjunction to determine resource
preemption among tunnels. If multiple CR-LSPs are to be established, CR-LSPs with high
priorities can be established by preempting resources. If resources (such as bandwidth) are
insufficient, a CR-LSP with a higher setup priority can preempt resources of an established
CR-LSP with a lower holding priority.
As shown in Figure 3-6, there are two TE tunnels on the network. The link bandwidth
allocation is shown in Figure 3-6 and the links have the same metric value.
– Tunnel 0/0/1: established over the path LSRA -> LSRB -> LSRE. Its bandwidth is 100
Mbit/s, and its setup and holding priority values are 0.
– Tunnel 0/0/2: established over the path LSRC -> LSRB -> LSRF. Its bandwidth is 100
Mbit/s, and its setup and holding priority values are 7.
Figure 3-6 Before a link fault occurs
100M
LSRE
100M
100M 100M
1G
1G
LSRA
LSRB
LSRC LSRD
LSRF
Tunnel 0/0/1
Tunnel 0/0/2
1G
Path1
Path2
Path of Tunnel 0/0/2
Path of Tunnel 0/0/1
 
If the link between LSRB and LSRE fails, LSRA recalculates a path LSRA -> LSRB ->
LSRF -> LSRE for tunnel 0/0/1. The link between LSRB and LSRF is shared by tunnel
0/0/1 and tunnel 0/0/2, but has insufficient bandwidth for these two tunnels. As a result,
preemption is triggered, as shown in Figure 3-7.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
52

Figure 3-7 After preemption is triggered
100M
LSRE
100M
100M 100M
1G
1G
LSRA
LSRB
LSRC LSRD
LSRF
Tunnel 0/0/1
Tunnel 0/0/2
1G
Path3
Path2
Old path of Tunnel 0/0/2
New path of Tunnel 0/0/1
Link fault
Preemption
Occur 
Path4
New path of Tunnel 0/0/2
 
