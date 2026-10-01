---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-83
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["compute", "copyright", "parameters", "preemption"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [11638, 11742]
sha256: 06d30544a1d71aba096aa1ea4749f2877c9ad4e60644a216e39ae669199d95ed
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Tunnel Attribute
                 An MPLS TE tunnel consists of several CR-LSPs. The constraints for these LSPs are
                 determined by the attributes of the tunnel. The constraints are categorized as
                 bandwidth constraints and path constraints.
                 ●      Bandwidth constraint: The bandwidth of a tunnel should be planned and
                        configured based on services that will be transmitted through it. When the
                        tunnel is established, the configured bandwidth is reserved on each node
                        along the tunnel path, implementing bandwidth assurance.
                 ●      Path constraints: consist of many factors such as the metric type, affinity
                        attribute, explicit path, priority and preemption, and hop limit.
                        –   Metric type: You can select either the IGP metric or TE metric for
                            calculating MPLS TE paths. By default, the TE metric type is used.
                        –   Affinity attribute: a 32-bit vector that specifies the links required by a TE
                            tunnel. This attribute is configured on the ingress of a tunnel and must
                            be used in conjunction with the administrative group attribute. After a
                            tunnel is assigned an affinity, a device evaluates this affinity against the
                            link administrative group attribute during the process of link selection.
                            Depending on the evaluation result, the device determines whether to
                            choose a link that possesses the specified attributes. For details about
                            affinity attributes, see Affinity Attribute.
                        –   Priorities and preemption are used to prioritize the establishment of TE
                            tunnels for transmitting critical services, thereby preventing competition
                            for resources during the tunnel setup process. Setup priorities and holding
                            priorities of tunnels determine whether a new tunnel can preempt the
                            resources currently occupied by existing tunnels. If the setup priority of a
                            new CR-LSP exceeds the holding priority of an existing CR-LSP, the new
                            CR-LSP can preempt the resources of the existing CR-LSP. The priority
                            ranges from 0 to 7. The value 7 indicates the lowest priority. The setup
                            priority of a tunnel cannot be higher than the holding priority of this
                            tunnel.
                        –   Explicit path
                            An explicit path is a CR-LSP that is established by manually specifying the
                            nodes to traverse or bypass. Explicit paths are divided into two categories:
                            strict and loose. For details about explicit paths, see 4.17.1
                            Understanding MPLS TE Explicit Paths.
                        –   Hop limit


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                 198
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                          A hop limit acts as a criterion for path selection during the establishment
                          of a CR-LSP. Similar to the administrative group and affinity attributes, a
                          hop limit defines the number of hops that a CR-LSP allows.


4.2.2 MPLS TE Tunnel Types
                 Depending on the mode of CR-LSP establishment, MPLS TE tunnels are
                 categorized into two types: static and dynamic.


Static MPLS TE Tunnel
                 Static MPLS TE tunnels are created using static CR-LSPs. To do this, manually
                 configure label forwarding and resource reservation parameters on each node
                 involved to create a static CR-LSP, and then associate the static CR-LSP to a tunnel
                 interface. A static MPLS TE tunnel is then created.


Dynamic MPLS TE Tunnel
                 Dynamic MPLS TE tunnels are created using dynamic CR-LSPs. RSVP-TE serves as
                 the signaling protocol for their creation, and therefore dynamic MPLS TE tunnels
                 are also referred to as RSVP-TE tunnels. Their establishment involves a series of
                 protocol components, including the information advertisement, path calculation,
                 and path establishment components. For details about these components, see
                 Table 4-1.


                 Table 4-1 Components involved for establishing dynamic MPLS TE tunnels

                  Compone         Description
                  nt Name

                  Informatio      The information advertisement component can be either an IS-IS
                  n               TE or OSPF TE one. It collects the attributes of each link on each
                  advertisem      node, floods the information, and advertises the information to
                  ent             other nodes on the MPLS TE network to generate traffic
                  componen        engineering databases (TEDBs). These databases are then utilized
                  t               to calculate the optimal path for a CR-LSP.

                  Path            The path computation component uses the constrained shortest
                  calculation     path first (CSPF) algorithm to compute the shortest path that
                  componen        meets the tunnel attribute requirements based on the link
                  t               attributes recorded in the TEDB. Evolving from the shortest path
                                  first (SPF) algorithm, CSPF excludes nodes and links that do not
                                  satisfy tunnel attributes from the topology and then uses SPF to
                                  calculate a path.

                  Path            The path establishment component is used by the RSVP-TE
                  establishm      signaling protocol to dynamically reserve resources allocate labels
                  ent             on each node, thereby establishing a dynamic CR-LSP. This
                  componen        approach eliminates the need for hop-by-hop configuration,
                  t               making it suitable for large-scale networks.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            199
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


