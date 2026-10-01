---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-148
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [21234, 21395]
sha256: 5f9824ef0a2feadeb17521df1de03688f0e5f2904db6f55ba872697f8768f6c1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 Figure 4-30 SRLG networking




                 Generally, core nodes on a backbone network (such as P1, P2, and P3 in this
                 example) are connected by transport network devices. The upper part of Figure
                 4-30 shows the abstract topology, and the lower part shows the actual topology.
                 NE1 is a transport network device. P1, P2, and P3 share the links marked yellow. If
                 a fault occurs on a shared link, both the primary and TE FRR bypass tunnels are
                 affected, causing an FRR protection failure. An SRLG can be configured to prevent
                 the TE FRR bypass tunnel from sharing a link with the primary CR-LSP, ensuring
                 that TE FRR properly protects the primary TE tunnel.

                 An SRLG is a set of links that share the same risk. If one of the links fails, other
                 links in the group may fail. Protection fails even if other links in the group function
                 as the hot standby or bypass CR-LSPs for the failed link.

Implementation
                 The SRLG link attribute is represented using a specific value. Links with the same
                 SRLG link attribute value belong to the same SRLG.

                 SRLG information is advertised to all nodes in an MPLS TE domain through IGP TE,
                 enabling all the nodes to acquire the SRLG information of all links within the
                 domain. The CSPF algorithm uses the SRLG attribute together with other
                 constraints, such as bandwidth, to calculate a path.

                 TE SRLG works in either of the following modes:

                 ●      Strict mode: The SRLG attribute is a necessary constraint used by CSPF to
                        calculate a hot-standby CR-LSP or TE FRR bypass path.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            355
MPLS Configuration
MPLS Configuration                                                                  4 MPLS TE Configuration


                 ●      Preferred mode: The SRLG attribute is an optional constraint used by CSPF to
                        calculate a hot-standby path or TE FRR bypass path. For example, if hot
                        standby is configured for a tunnel and CSPF fails to calculate a backup path
                        based on the SRLG constraint for the first time, CSPF will no longer take the
                        SRLG constraint into account in the second attempt to calculate a backup
                        path.


Application Scenarios
                 SRLGs apply to networks with CR-LSP hot standby or TE FRR configured.


Benefits
                 An SRLG restricts the selection of hot-standby paths and TE FRR bypass paths,
                 preventing the primary and backup paths from being established on links with the
                 same risk.

4.19.2 Configuring an MPLS TE SRLG

Prerequisites
                 Before configuring an MPLS TE SRLG, complete the following task:

                 ●      Configure a dynamic MPLS TE tunnel.


Context
                 SRLGs are mainly used in CR-LSP hot standby and TE FRR scenarios to enhance TE
                 tunnel reliability. SRLG configuration includes:

                 ●      Configure an SRLG-based path calculation mode globally.
                 ●      Configure SRLG attributes.
                 ●      (Optional) Delete all member interfaces of all SRLGs.

                 Perform the following operations based on actual requirements:


Procedure
                 ●      Configure an SRLG-based path calculation mode globally.

                        Perform the following steps on the ingress of a primary tunnel with CR-LSP
                        hot standby configured or on the PLR of a TE FRR bypass tunnel.

                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS view.
                             mpls

                        c.   Configure an SRLG-based path calculation mode.
                             mpls te srlg path-calculation [ strict | preferred ]




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             356
MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration


                                     NOTE

                                    ● If strict is configured, CSPF always uses an SRLG as a constraint when
                                      calculating a bypass CR-LSP or backup CR-LSP path.
                                    ● If preferred is configured, CSPF uses an SRLG as a constraint when calculating
                                      a bypass CR-LSP or backup CR-LSP path for the first time. However, if
                                      calculation fails, CSPF no longer uses the SRLG as a constraint.
                 ●      Configure SRLG attributes.

                        Perform the following steps on an interface of an SRLG member link.

                        a.   Enter the system view.
                             system-view

                        b.   Enter the view of an MPLS TE-enabled link interface.
                             interface interface-type interface-number

                        c.   Switch the interface working mode from Layer 2 to Layer 3.
                             undo portswitch

                             Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                             S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                             be switched from Layer 2 mode to Layer 3 mode using the undo
                             portswitch command.Determine whether to perform this step based on
                             the current interface mode.
                        d.   Add the interface to an SRLG.
                             mpls te srlg srlg-number

                             In a CR-LSP hot standby or TE FRR scenario, you need to configure SRLG
                             attributes for the outbound tunnel interface of the MPLS TE tunnel
                             ingress or PLR, as well as for other member links in the SRLG to which
                             the interface belongs. When adding a link to an SRLG, you only need to
                             configure SRLG attributes on any outbound interface of the link.
                 ●      (Optional) Delete all member interfaces of all SRLGs.

                        To delete the SRLG configuration on all interfaces of an MPLS TE tunnel node
                        in batches, perform the following steps on the node.

                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS view.
                             mpls

                        c.   Delete all member interfaces of all SRLGs from the TE node.
                             undo mpls te srlg all-config

                                     NOTE

                                    Running the undo mpls te srlg all-config command does not delete the mpls te
                                    srlg path-calculation configuration in the MPLS view.

                 ----End

Verifying the Configuration
                 ●      Run the display mpls te srlg { srlg-number | all } command to check SRLG
                        configuration and SRLG member interface information.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      357

