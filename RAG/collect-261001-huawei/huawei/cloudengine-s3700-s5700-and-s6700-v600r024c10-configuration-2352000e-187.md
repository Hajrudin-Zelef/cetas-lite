---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-187
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [27229, 27332]
sha256: 08d6fd554d0198b33087405e6c2ea93f8657f6e294b7bc5529aa59d823dc1a69
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    The following assumes that PE1 and PE2 establish a VC with CE-m and CE-n,
                    respectively, and that they belong to the same VPN (VPN-x).

                    PE1 receives a label block LBn/LRn/LOn from PE2.

                    1.   PE1 checks whether the encapsulation type of CE-n received from PE2 is the
                         same as that of CE-m. If not, PE1 stops subsequent processing.
                    2.   PE1 checks whether the CE IDs (m and n) are the same. If so, PE1 reports an
                         error and stops subsequent processing.
                    3.   If CE-m has multiple label blocks, PE1 checks whether these label blocks meet
                         the condition: LOm ≤ n < LOm + LRm. If this condition is not met, PE1 reports
                         an error and stops subsequent processing.
                    4.   PE1 checks whether all the label blocks of CE-n meet the condition: LOn ≤ m
                         < LOn + LRn. If this condition is not met, PE1 reports an error and stops
                         subsequent processing.
                    5.   PE1 checks whether the outer tunnel between PE-m and PE-n is established
                         normally. If not, PE1 stops the process. The outer tunnel is assumed as an LSP
                         with the label Z.
                    6.   PE1 allocates an inner label (LBn + m - LOn), namely, the outgoing label of
                         the VC, to CE-n, and allocates an inner label (LBm + n - LOm), namely, the
                         incoming label of the VC, to CE-m.
                    7.   The label of the outer tunnel from PE2 to PE1 is Z.
                    8.   After the inner and outer labels are calculated and the VC is up, Layer 2
                         packets can be transmitted.

                    The following example describes the process of allocating CE label blocks.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           433
VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration


                    Assume that PEs exchange label block information using BGP; all public network
                    LSPs are up, and only VC labels need to be calculated. The ID of CE1 is 1; the ID of
                    CE2 is 2; and so on.

                    Figure 5-11 Calculating VC labels




                    In Figure 5-11, label blocks are allocated to CE1 and CE3 as follows:

                    1.   PE1 allocates a label block (LB/LR/LO = 1000/5/0) to CE1 and receives the
                         label block (LB/LR/LO = 1010/2/0) allocated to CE3 from PE2. According to
                         the preceding calculation rule, the incoming and outgoing VC labels can be
                         calculated.
                    2.   PE1 is also connected to CE2. Following the allocation of the label block to
                         CE1, PE1 allocates a label block (LB/LR/LO = 1005/50/0) to CE2. CE2 does not
                         establish any connection with other CEs. The label block is allocated for future
                         capacity expansion. Therefore, PE1 does not calculate labels for CE2.
                    3.   PE1 determines whether the label block allocated to CE1 meets the
                         requirement: LOm ≤ n < LOm + LRm. Because LOm is 0, n is 3, and LRm is 5,
                         this requirement is met. PE1 also determines whether this label block meets
                         the requirement: LOn ≤ m < LOn + LRn. Because LOn is 0, m is 1, and LRn is 2,
                         this requirement is also met.
                    4.   PE1 calculates the incoming and outgoing VC labels. The outgoing VC label
                         (inner label of CE3) is LBn + m - LOn = 1010 + 1 - 0 = 1011. The incoming VC
                         label (inner label of CE1) is LBm + n - LOm = 1000 + 3 - 0 = 1003.
                    5.   PE2 determines whether the label block allocated by PE1 to CE1 meets the
                         requirement: LOm ≤ n < LOm + LRm. Because LOm is 0, n is 1, and LRm is 2,
                         this requirement is met. PE2 also determines whether this label block meets
                         the requirement: LOn ≤ m < LOn + LRn. Because LOn is 0, m is 3, and LRn is 5,
                         this requirement is also met.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           434
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


                    6.   PE2 calculates the incoming and outgoing VC labels. The outgoing VC label
                         (inner label of CE1) is LBn + m - LOn = 1000 + 3 - 0 = 1003. The incoming VC
                         label (inner label of CE3) is LBm + n - LOm = 1010 + 1 - 0 = 1011.
                         In Figure 5-12, if CE13 is added to the VPN, a VC needs to be established
                         between CE13 and CE1, and PE3 allocates a label block LB/LR/LO = 1000/4/0
                         to CE13.

                         Figure 5-12 Calculating VC labels after a CE is added




                         Whether the ID of CE13 is correct is judged in a similar way. PE1 determines
                         whether the label block allocated to CE1 meets the requirement: LOm ≤ n <
                         LOm + LRm. Because LOm is 0, n is 13, and LRm is 5, this requirement is not
                         met. PE1 determines whether the label block allocated to CE1 meets the
                         requirement: LOn ≤ m < LOn + LRn. Because LOn is 0, m is 1, and LRn is 4,
                         this requirement is met.
                         The CE ID 13 is greater than the value (LO + LR) of CE1; therefore, the
                         outgoing VC label (inner label of CE1) cannot be calculated. The LR of CE1
                         needs to be changed to 15 on PE1 in this example, and another label block
                         needs to be to CE1 with the LR as 10. Following the allocation of the label
                         block to CE2, PE1 allocates a second label block (LB/LR/LO = 1055/10/5) to
                         CE1. PE1 re-determines whether this label block meets the requirement: LOm
                         ≤ n < LOm + LRm. Because LOm is 5, n is 13, and LRm is 10, this requirement

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          435
VPN Configuration
VPN Configuration                                                                   5 VPWS Configuration


