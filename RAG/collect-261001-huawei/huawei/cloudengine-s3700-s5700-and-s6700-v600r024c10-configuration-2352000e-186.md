---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-186
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [27095, 27228]
sha256: 90656e99d45cdea35a4274d93ca5f7182c99e31f46a60b1d797babadf6ececfe
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    BGP VPWS uses VPN targets to control the advertisement or acceptance of VPN
                    routes, which improves networking flexibility. BGP VPWS assigns VC labels from
                    label blocks. A label block is allocated to each CE in advance. The size of the label
                    block allocated to a CE determines the number of connections that the CE can
                    establish with other CEs. BGP VPWS allows allocation of additional labels for
                    future capacity expansion. PEs obtain the inner labels of packets based on these
                    label blocks and then transmit packets based on the inner labels. BGP VPWS has
                    good scalability and supports both local and remote connections.

                    Figure 5-7 BGP VPWS topology




Basic Concepts
                    BGP VPWS can use label blocks to allocate labels to multiple connections at the
                    same time. The CE range specified for a CE indicates the number of connections
                    that can be established between this CE and other CEs. Only one label block can
                    be allocated to a CE at one time. The label block size equals the CE range.
                    Additional label allocation may waste labels in the short term, but will reduce the
                    configuration workload during future VPN capacity expansion. For example, an
                    enterprise VPN has 10 CEs, and the number may increase to 20 in the future.
                    Here, the CE range can be set to 20, with 10 labels reserved for new CEs.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           430
VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration


Implementation
                    Packet forwarding in BGP VPWS is similar to that in LDP VPWS. The two types of
                    VPWS both use two labels in compliance with standards, but they use different
                    signaling protocols to exchange these labels: LDP VPWS uses extended LDP,
                    whereas BGP VPWS uses Multiprotocol Extensions for BGP (MP-BGP).

                    Figure 5-8 BGP VPWS connection modes




                    On the network shown in Figure 5-8, six CEs (CE1 to CE6) access VPN1. To enable
                    these CEs to communicate, full-mesh connections must be established between
                    them. That is, a CE must establish a VC with each of the other CEs. To establish
                    these connections, perform the following configurations on PE1, PE2, and PE3:
                    1.   Create VPN1 on each PE and create CEs connected to the PE. For example, on
                         PE1, create CE1, CE2, and CE3.
                    2.   Allocate label blocks to CEs. Here, each CE needs to be connected to five CEs.
                         Therefore, a label block containing at least five labels must be allocated to
                         each CE.
                    3.   On each PE, specify peer CE IDs and PE interfaces connecting to local CEs.
                    Like CCC VPWS, BGP VPWS also supports local connections, and PEs serve as
                    switches. It is easy to use BGP VPWS to establish full-mesh connections.

VC Label Calculation
                    A label block is a consecutive range of labels. BGP VPWS uses MP-BGP as the
                    signaling protocol to transmit label block information. A label block is a
                    consecutive range of labels.
                    For a clear description of a label block, several parameters are defined: Label Base
                    (LB) (indicating the initial label of the label block), Label Range (LR) (indicating
                    the label block size), and Label-Block Offset (LO), as shown in Figure 5-9.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             431
VPN Configuration
VPN Configuration                                                                   5 VPWS Configuration


                    Figure 5-9 Calculating VC labels




                    During the configuration of a CE on a PE, the LR of the label block needs to be
                    specified, and the LB is automatically allocated by the PE. This label block is
                    transmitted to other PEs as a network layer reachable information (NLRI) entry
                    through BGP. When the configuration of the CE is deleted or the connection
                    between the CE and PE becomes invalid, the label block is deleted. BGP then sends
                    a Label Withdraw message for notification. The following assumes that when BGP
                    VPWS is deployed, CE1 needs to establish two VCs with other remote CEs. In this
                    case, the LR cannot be smaller than 2. To allow for future capacity expansion, the
                    LR can be set to 10.
                    The labels may be insufficient as VCs increase, regardless of the LR. If this
                    situation occurs, the LR needs to be redefined for a larger label space. However,
                    the data of the label block is transmitted through the BGP NLRI, and this label
                    block has been used to calculate VC labels and forward data. To protect the
                    original VC connection, a new label block is allocated to this CE and advertised as
                    a new NLRI through BGP. In this situation, the label space of a CE may consist of
                    multiple label blocks. The LO defines the relationship between multiple labels and
                    identifies the total size of the label blocks preceding a label block. For example, if
                    the LR of the first label block is 100 and the LO is 0, and the LR of the second
                    label block is 50, then the LO of the second label block is 100 and the LO of the
                    third label block if any is 150. The LO is used to calculate VC labels. As such, a
                    label block can be defined by three parameters: LB, LR, and LO.
                    A CE ID can be used to:
                    ●   Uniquely identify a CE in a VPN. In a VPN, CE IDs must be unique. A CE ID is
                        carried in each NLRI; different label blocks can therefore be associated with
                        their corresponding CEs.
                    ●   Calculate VC labels. If the LR of the local CE is x and the local CE needs to be
                        connected to the remote CE with the CE ID being y, the condition (x > y) must
                        be met. If this condition is not met, increase the value of x.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              432
VPN Configuration
VPN Configuration                                                                    5 VPWS Configuration


                    Figure 5-10 Calculating label blocks




                    Table 5-3 Calculating label blocks

                     Item                    Description      Item             Description

                     Label block allocated   Lm               Label block      Ln
                     by PE1 to CE-m                           allocated
                                                              by PE2 to
                                                              CE-n

                     LO of Lm                LOm              LO of Ln         LOn

                     LB of Lm                LBm              LB of Ln         LBn

                     LR of Lm                LRm              LR of Ln         LRn




