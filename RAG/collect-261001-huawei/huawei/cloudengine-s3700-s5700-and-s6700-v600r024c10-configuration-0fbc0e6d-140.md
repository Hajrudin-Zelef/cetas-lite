---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-140
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [20074, 20223]
sha256: 74bb0b432f32c1cb6ab1846ef030a2c446db548d5c7b02380c4a0d793e9286eb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Verifying the Configuration
                 ●      Run the display mpls te tunnel-interface command to check tunnel
                        interface information on the local node.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       336
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration




4.16 Configuring the Link Administrative Group and
Affinity Attributes of MPLS TE

4.16.1 Understanding the Link Administrative Group and
Affinity Attributes of MPLS TE
                 The link administrative group and affinity attributes of MPLS TE can be used by
                 network administrators to control the paths over which MPLS TE tunnels are
                 established.

Link Administrative Group
                 The link administrative group attribute, also called link color, is represented by a
                 32-bit vector. Each bit in this vector can be assigned a specific meaning or left
                 unassociated, depending on the configuration, such as link bandwidth, a
                 performance parameter (such as delay), or a management policy. The policy can
                 be a traffic type (multicast for example) or a flag indicating that a link is used by
                 an MPLS TE tunnel. A link administrative group attribute needs to work with
                 Affinity Attribute to control tunnel paths.

Affinity Attribute
                 The affinity attribute is a 32-bit vector that specifies the links required by a TE
                 tunnel. This attribute is configured on the ingress of a tunnel and must be used in
                 conjunction with the Link Administrative Group.
                 After a tunnel is assigned an affinity, a device evaluates this affinity against the
                 link administrative group attribute during the process of link selection. Depending
                 on the evaluation result, the device determines whether to choose a link that
                 possesses the specified attributes. The link selection criteria are as follows:
                 ●      The result of performing an AND operation between the IncludeAny affinity
                        and the link administrative group attribute is not 0.
                 ●      The result of performing an AND operation between the ExcludeAny affinity
                        and the link administrative group attribute is 0.
                 IncludeAny = the affinity attribute value ANDed with the subnet mask value;
                 ExcludeAny = (–IncludeAny) ANDed with the subnet mask value; the link
                 administrative group value = the link administrative group value ANDed with the
                 subnet mask value.
                 The following rules apply:
                 ●      If some bits in the mask are 1s, at least one bit in the link administrative
                        group attribute is 1 and its corresponding affinity bit must be 1. If all the bits
                        in the affinity attribute are 0s, the corresponding bit in the link administrative
                        group must not be 1.
                 ●      If a bit in a mask is 0, the corresponding bit in the link administrative group is
                        not checked.
                 Figure 4-26 uses a 16-digit attribute value as an example to describe how the
                 affinity attribute works with the link administrative group attribute.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              337
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


                 Figure 4-26 Affinity attribute




                 The masks of the affinity attribute determine the bits of the link administrative
                 group attribute to be checked by the device. In this example, the bits with the
                 mask of 1 are bits 11, 13, 14, and 16, indicating that these bits need to be
                 checked. The values of bit 11 in both the affinity and the link administrative group
                 attributes of the link are 0s (not 1). In addition, the values of bits 13 and 16 in
                 both the affinity and the administrative group attributes of the link are 1s.
                 Therefore, the link matches the affinity of the tunnel and can be selected for the
                 tunnel.

                         NOTE

                        Understand specific comparison rules before deploying devices of different vendors because
                        the comparison rules vary with vendors.


4.16.2 Configuring the Link Administrative Group and Affinity
Attributes of MPLS TE

Prerequisites
                 Before configuring the link administrative group and affinity attributes, complete
                 the following task:

                 ●      Configure a dynamic MPLS TE tunnel.


Context
                 The affinity attribute of an MPLS TE tunnel is crucial in defining the tunnel's link
                 attributes. Combined with the link administrative group attributes, it specifies
                 which links the tunnel can utilize.

                 Configure the affinity attribute on the ingress of an MPLS TE tunnel and the link
                 administrative group attribute on the outbound interface of each node on the
                 tunnel.

                         NOTE

                        ● The updated link administrative group attribute takes effect only on the LSPs that are
                          yet to be established, not those that have already been established.
                        ● Changing a tunnel's affinity attribute configuration affects its LSPs that have been
                          established, triggering path re-computation for the tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                       338
MPLS Configuration
MPLS Configuration                                                                           4 MPLS TE Configuration


         Step 2 Enter the view of an MPLS TE-enabled link interface.
                 interface interface-type interface-number

         Step 3 Switch the interface working mode from Layer 2 to Layer 3.
                 undo portswitch

                 Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                 S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                 Layer 2 mode to Layer 3 mode using the undo portswitch command.

                 Determine whether to perform this step based on the current interface mode.

         Step 4 Configure the link administrative group attribute.
                 mpls te link administrative group group-value

         Step 5 Return to the system view.
                 quit

         Step 6 Enter the MPLS TE tunnel interface view.
                 interface tunnel tunnel-number

         Step 7 Configure the affinity attribute for the tunnel.
                 mpls te affinity property properties [ mask mask-value ] { secondary | best-effort }

                 ----End


