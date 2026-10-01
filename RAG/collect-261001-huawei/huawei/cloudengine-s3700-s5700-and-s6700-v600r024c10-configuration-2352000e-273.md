---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-273
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["agi", "copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [40171, 40284]
sha256: cfafc8eb44e1745c33eb82bf72b65193cd24af0bc01839a637e4150ca4deb373
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Implementation
                    BGP AD VPLS uses BGP signaling for automatic VPLS member discovery,
                    simplifying the configuration and conserving labels. BGP AD VPLS combines the
                    advantages of LDP VPLS and BGP VPLS.
                    BGP AD VPLS-enabled devices use extended BGP Update messages carrying VPLS
                    member information to automatically discover VPLS members and use LDP FEC
                    129 to negotiate and establish VPLS PWs. In this way, VPLS members are
                    automatically discovered and VPLS PWs are automatically established.
                    1.     Automatically discovering VPLS members
                           Automatically discovering VPLS members is the first phase of PW
                           establishment. Figure 6-16 shows the process for automatically discovering
                           VPLS members using BGP.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                              643
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                         Figure 6-16 Automatically discovering VPLS members




                         The process for automatically discovering VPLS members is as follows:
                         a.   After having parameters such as the VPLS ID, RD, RT, and VSI ID
                              configured, PE1 and PE2 encapsulate these parameters into a BGP
                              Update message and send the message as a BGP AD message to all peer
                              PEs in the BGP domain.
                         b.   After a PE receives a BGP Update message from its peer, the PE filters the
                              BGP AD message based on the RT policy configured on it. If the message
                              matches the RT policy, the PE obtains remote VSI information from the
                              message and compares this information with the locally configured
                              information.

                              ▪   If the VPLS IDs of VSIs on PEs at both ends are the same, the two
                                  VSIs belong to the same VPLS domain and can have only one PW
                                  established between them.

                              ▪   If the VPLS IDs of VSIs on PEs at both ends are different, the two
                                  VSIs belong to different VPLS domains and can have no PW
                                  established between them.
                    2.   Automatically deploying a VPLS PW
                         After VPLS member discovery is complete, LDP FEC 129 is used to negotiate
                         and establish a PW. Figure 6-17 shows the process for automatically
                         establishing a VPLS PW.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            644
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration


                        Figure 6-17 Process for automatically establishing a VPLS PW




                        The process for automatically deploying a VPLS PW is as follows:
                        a.   After an LDP session is established between PE1 and PE2, PE1 and PE2
                             advertise an LDP Label Mapping (FEC 129) message carrying information
                             including the AGI, SAII, TAII, and label to each other.
                                  NOTE

                             After VPLS member discovery, BGP AD VPLS proactively triggers LDP to establish an
                             LDP session based on service requirements. If this LDP session is not needed after
                             VPLS services are withdrawn, BGP AD VPLS proactively triggers LDP to tear down this
                             LDP session. This method reduces the LDP session topology maintenance workload,
                             improves system resource utilization, and improves network performance.
                        b.   After a PE receives an LDP Label Mapping (FEC 129) message from its
                             peer, the PE obtains information including the VPLS ID, PW type, MTU,
                             and TAII from the message, and compares this information with its local
                             VSI information. If the information matches, negotiation succeeds, and
                             the requirements of establishing a PW are met, the PE establishes a PW
                             with its peer.


Application Scenarios
                    As VPLS technologies are used more widely and VPLS networks grow in scale,
                    VPLS configurations on networks increase accordingly. BGP AD VPLS is introduced
                    to simplify network configurations, enable automatic service deployment, and
                    reduce OPEX.


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                    645
VPN Configuration
VPN Configuration                                                                6 VPLS Configuration


                    BGP AD VPLS combines the advantages of BGP VPLS and LDP VPLS. BGP AD VPLS
                    uses extended BGP messages to automatically discover VPLS members and uses
                    LDP FEC 129 to negotiate and establish a PW for automatic VPLS PW deployment.
                    BGP AD VPLS also has the following advantages over other VPLS modes:
                    ●   As described in Table 6-6, BGP AD VPLS saves local label resources but
                        involves complex PW establishment and depends on LDP signaling.
                    ●   As described in Table 6-7, BGP AD VPLS uses existing BGP sessions to discover
                        members in a VPLS domain when new nodes are added to the VPLS domain
                        or a new VPLS domain is deployed. In this way, PWs can be established
                        between the local and remote devices without member information being
                        explicitly configured. This simplifies PW configurations on PEs.
                    With automatic VPLS member discovery and automatic PW deployment, BGP AD
                    VPLS reduces the VPLS network configuration workload, implements automatic
                    service deployment, and reduces customers' OPEX.

                    Table 6-6 Comparison between BGP AD VPLS and BGP VPLS
                     VPLS Mode                  Advantage                   Disadvantage

