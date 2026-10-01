---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-97
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [13643, 13776]
sha256: 46b8fb3d12128822e3210071506ac925b12f84f7243f4ba171761a5fdb3fcbee
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                             IS-IS TE uses the sub-TLV of TLV 22 to carry TE attributes. Therefore, the
                             IS-IS wide metric attribute must be enabled. By default, IS-IS sends and
                             receives only the packets carrying a route metric that is expressed in
                             narrow mode.
                        d.   Enable IS-IS TE.
                             traffic-eng [ level-1 | level-2 | level-1-2 ]

                             If no IS-IS level is specified, IS-IS TE takes effect for both Level-1 and
                             Level-2.
                        e.   (Optional) Set the types of sub-TLVs that carry the DiffServ-aware Traffic
                             Engineering (DS-TE) attribute.
                             te-set-subtlv { bw-constraint bw-constraint-value | lo-multiplier lo-multiplier-value |
                             unreserve-bw-sub-pool unreserve-bw-sub-pool-value }*

                             There are no unified standards for sub-TLVs that carry DS-TE attributes in
                             non-IETF mode. To ensure interconnection between devices of different
                             vendors, you need to manually configure TLV values for these sub-TLVs.
                             After the command is run, the TEDB is regenerated, which causes the

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                            231
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                           reestablishment of TE tunnels. TLVs can be sent only if sub-TLVs are
                           configured for them.

                 ----End

4.7.3 Configuring CSPF

Context
                 To calculate CR-LSP paths that satisfy specified constraints, configure CSPF on the
                 tunnel ingress. CSPF extends the shortest path first (SPF) algorithm and is able to
                 calculate the shortest path meeting MPLS TE requirements.

                 Specifically, the CSPF algorithm uses the link attributes maintained in the TEDB to
                 prune the nodes and links that do not meet the tunnel attributes, and then uses
                 the SPF algorithm to find the shortest path to the destination of a tunnel.

                 For example, on the network shown in Figure 4-8, except the blue links and the
                 links marked a specific bandwidth value, all the links are black and have the
                 bandwidth of 100 Mbit/s. The lower part of Figure 4-8 shows the topology pruned
                 by CSPF when an MPLS TE tunnel that meets the following constraints needs to be
                 established: The destination is LSR5, the tunnel bandwidth is 80 Mbit/s, the
                 affinity (a 32-bit vector that describes the links to be used by a TE tunnel; detailed
                 in 4.16.1 Understanding the Link Administrative Group and Affinity Attributes
                 of MPLS TE) is black, and the transit node is LSR7.


                 Figure 4-8 Prune process of CSPF




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                            232
MPLS Configuration
MPLS Configuration                                                                                 4 MPLS TE Configuration


                 Then, CSPF calculates a path in the same way as SPF does. Figure 4-9 shows the
                 calculation result.

                 Figure 4-9 CSPF calculation result




                 During CSPF path calculation, if multiple paths have the same weight, you can
                 configure tie-breaking to select the optimal path.
                 The path is selected based on the ratio of the available remaining bandwidth to
                 the maximum reservable bandwidth. The following tie-breaking policies are
                 available:
                 ●      Most-fill: The path with a larger percentage of the available remaining
                        bandwidth to the maximum reservable bandwidth is preferred. That is, the
                        path with lower bandwidth usage is preferred.
                 ●      Least-fill: The path with a smaller percentage of the available remaining
                        bandwidth to the maximum reservable bandwidth is preferred. That is, the
                        path with higher bandwidth usage is preferred.
                 ●      Random: The device selects a path at random. This mode allows LSPs to be
                        evenly distributed among links, regardless of their bandwidth.

                         NOTE

                        The most-fill and least-fill modes take effect only when the bandwidth usage difference
                        between two links exceeds 10%. For example, if the bandwidth usage of link A is 50% and
                        the bandwidth usage of link B is 45% (a 5% difference), the most-fill and least-fill modes
                        do not take effect and the random mode is used.

                 Perform the following configuration on the ingress of an MPLS TE tunnel.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Enable CSPF on the local node.
                 mpls te cspf

         Step 4 (Optional) Configure the preferred IGP, process, and area for CSPF-based path
                calculation.
                 mpls te cspf preferred-igp { isis [ process-id [ level-1 | level-2 ] ] | ospf [ process-id [ area area-id ] ] }

                 By default, OSPF is preferred.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                    233
MPLS Configuration
MPLS Configuration                                                                             4 MPLS TE Configuration


         Step 5 (Optional) Configure CSPF to calculate the shortest path from all IGP processes
                and areas.
                 mpls te cspf multi-instance shortest-path [ preferred-igp { isis | ospf } [ process-id ] ]

                         NOTE

                        This command is mutually exclusive with the mpls te cspf preferred-igp command run in
                        Step 4. If both of them are run, the latest configuration overrides the previous one.

         Step 6 (Optional) Enable the longest match function.
                 mpls te cspf loose-explicit-path longest-match

                 After this command is run, the path with the largest number of hops is
                 preferentially selected in multi-segment explicit path scenarios. If multiple explicit
                 paths are qualified and have the same number of hops, the path with the smallest
                 metric value is preferentially selected.
         Step 7 (Optional) Disable the CSPF-specific optimization mode.
                 mpls te cspf optimize-mode disable

