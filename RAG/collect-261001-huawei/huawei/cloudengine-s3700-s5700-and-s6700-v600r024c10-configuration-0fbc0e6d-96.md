---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-96
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [13490, 13642]
sha256: 7beaf481c764f21e190fbd3a1651674abcbee4c331f86592b1a045c752fae2fa
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

         Step 7 Enter the MPLS TE link interface view.
                 interface interface-type interface-number

         Step 8 Switch the interface working mode from Layer 2 to Layer 3.
                 undo portswitch

                 Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                 S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                 Layer 2 mode to Layer 3 mode using the undo portswitch command.

                 Determine whether to perform this step based on the current interface mode.

         Step 9 Enable MPLS on the interface.
                 mpls

        Step 10 Enable MPLS TE on the interface.
                 mpls te

        Step 11 Enable RSVP-TE on the interface.
                 mpls rsvp-te

                         NOTE

                        A secondary IP address cannot be configured on an RSVP-TE-enabled interface. If a
                        secondary IP address is configured for such an interface, a tunnel may fail to be established.

        Step 12 (Optional) Configure link bandwidth.

                 Plan link bandwidth before you perform this procedure. The reserved bandwidth
                 must be higher than or equal to the bandwidth required by MPLS TE traffic.

                 In real-world applications, link bandwidth attributes only need to be configured on
                 the outbound interfaces that reside on a TE tunnel link and have specific
                 bandwidth requirements.

                 ●      Configure the maximum reservable link bandwidth.
                        mpls te bandwidth max-reservable-bandwidth max-bw-value

                        This command sets the bandwidth reserved for an MPLS TE tunnel on the
                        local link. By default, the maximum reservable bandwidth is not set. With the
                        default configuration, when the ingress of an MPLS TE tunnel initiates a
                        request to establish a CR-LSP with bandwidth constraints, the required CR-LSP
                        bandwidth may exceed the maximum reservable link bandwidth. If this is the
                        case, the CR-LSP fails to be established.
                 ●      Configure BC bandwidth for the link.
                        mpls te bandwidth bc0 bc0-bw-value


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       229
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


                         NOTE

                        ● Ensure that the maximum reservable link bandwidth does not exceed the actual physical
                          link bandwidth. You are advised to set the maximum reservable link bandwidth to be
                          less than or equal to 80% of the actual physical link bandwidth.
                        ● Ensure that the BC0 bandwidth does not exceed the maximum reservable link
                          bandwidth.

                 ----End

4.7.2 Configuring IGP TE

Context
                 All nodes, especially ingresses on an MPLS TE network, need to gather information
                 about link resource distribution so as to determine the paths and nodes that a
                 dynamic MPLS TE tunnel passes. In MPLS TE, the information about link resource
                 distribution is advertised by the information advertisement component. This
                 component is a TE extension upon an IGP. It automatically collects information for
                 advertisement and floods the information to other nodes within an MPLS TE area.
                 As a result, all nodes in the area maintain consistent information in the TEDB.

                 The information advertisement component can be either an IS-IS TE or OSPF TE
                 one.

                 ●      IS-IS TE
                        IS-IS TE uses the sub-type-length-value (sub-TLV) of the Extended IS
                        Reachability TLV (22) to carry TE attributes. By default, an IS-IS process does
                        not support TE. To support IS-IS TE, the IS-IS wide metric must be enabled.
                        The IS-IS wide metric supports the wide, compatible, and wide-compatible
                        metric types. By default, IS-IS sends and receives only the packets carrying a
                        route metric that is expressed in narrow mode.
                 ●      OSPF TE
                        OSPF TE uses Opaque Type 10 LSAs to carry TE attributes. By default, an
                        OSPF area does not support MPLS TE. To support OSPF TE, the OSPF Opaque
                        capability must be enabled. TE LSAs can be generated when at least one
                        OSPF neighbor is in the full state.

                 Information flooding is triggered by the establishment of an MPLS TE tunnel or by
                 one of the following conditions:
                 ●      A link is activated or deactivated.
                 ●      An LSP fails to be established when no adequate bandwidth can be reserved.
                 ●      Link attributes, such as the administrative group attribute or affinity attribute,
                        change.

                         NOTE

                        If neither OSPF TE nor IS-IS TE is configured, no TE LSA or TE LSP exists on the network,
                        and no TEDB can be generated. In this case, CR-LSPs are generated based on IGP routes,
                        rather than being calculated by the CSPF algorithm.
                        TE tunnels cannot be established across areas. In an inter-area scenario, you must configure
                        explicit paths and specify the inbound and outbound interfaces for a TE tunnel; otherwise,
                        the tunnel fails to be established.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    230
MPLS Configuration
MPLS Configuration                                                                             4 MPLS TE Configuration


                 Determine the TE information advertisement mode based on the IGP used on the
                 backbone network. Perform the following configuration on each node of an MPLS
                 TE tunnel.

Procedure
                 ●      Configure OSPF TE.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the OSPF view.
                             ospf [ process-id ]

                        c.   Enable the OSPF Opaque capability.
                             opaque-capability enable

                        d.   (Optional) Enable the device to advertise an MPLS LSR ID to multiple
                             areas.
                             advertise mpls-lsr-id

                                    NOTE

                                  Perform this step on an ABR only when it connects to multiple OSPF areas.
                        e.   Enter the OSPF area view.
                             area area-id

                        f.   Enable MPLS TE in the OSPF area.
                             mpls-te enable [ standard-complying ]

                 ●      Configure IS-IS TE.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the IS-IS view.
                             isis [ process-id ]

                        c.   Configure the IS-IS wide metric attribute.
                             cost-style { narrow | wide | wide-compatible | { compatible | narrow-compatible } [ relax-
                             spf-limit ]}

