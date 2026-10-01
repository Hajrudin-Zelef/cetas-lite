---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-1
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [1, 98]
sha256: 39af3423d10ef4f3e0b7afa568e6dce904d3d3cc9bbf0bd38ebb23bbbab80af9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

MPLS Configuration


MPLS Configuration

Issue    01
Date     2025-03-03




HUAWEI TECHNOLOGIES CO., LTD.
Copyright © Huawei Technologies Co., Ltd. 2025. All rights reserved.
No part of this document may be reproduced or transmitted in any form or by any means without prior
written consent of Huawei Technologies Co., Ltd.

Trademarks and Permissions

      and other Huawei trademarks are trademarks of Huawei Technologies Co., Ltd.
All other trademarks and trade names mentioned in this document are the property of their respective
holders.

Notice
The purchased products, services and features are stipulated by the contract made between Huawei and
the customer. All or part of the products, services and features described in this document may not be
within the purchase scope or the usage scope. Unless otherwise specified in the contract, all statements,
information, and recommendations in this document are provided "AS IS" without warranties, guarantees
or representations of any kind, either express or implied.

The information in this document is subject to change without notice. Every effort has been made in the
preparation of this document to ensure accuracy of the contents, but all statements, information, and
recommendations in this document do not constitute a warranty of any kind, express or implied.




Huawei Technologies Co., Ltd.
Address:       Huawei Industrial Base
               Bantian, Longgang
               Shenzhen 518129
               People's Republic of China

Website:       https://e.huawei.com




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                                  i
MPLS Configuration
MPLS Configuration                                                                                                                                                             Contents




                                                                                                                                                          Contents

1 About This Document.............................................................................................................1
2 Basic MPLS Configuration..................................................................................................... 5
2.1 Overview of MPLS................................................................................................................................................................... 5
2.2 Understanding MPLS............................................................................................................................................................. 6
2.2.1 Basic Concepts of MPLS..................................................................................................................................................... 6
2.2.2 LSP Establishment............................................................................................................................................................. 13
2.2.3 MPLS Forwarding.............................................................................................................................................................. 14
2.2.4 MPLS MTU........................................................................................................................................................................... 15
2.3 Configuration Precautions for MPLS Basics................................................................................................................. 15
2.4 Configuring the MPLS MTU.............................................................................................................................................. 15
2.5 Configuring MPLS TTL Processing Modes.................................................................................................................... 16
2.5.1 Understanding MPLS TTL Processing Modes........................................................................................................... 16
2.5.2 Configuring MPLS TTL Processing Modes.................................................................................................................19
2.6 Configuring Alarm Thresholds for MPLS Resources................................................................................................. 20
2.6.1 Configuring Alarm Thresholds for LDP LSPs............................................................................................................ 20
2.6.2 Configuring Alarm Thresholds for Dynamic Labels...............................................................................................21
2.6.3 Configuring Alarm Thresholds for Other LDP Resources.................................................................................... 22
2.6.4 Configuring Alarm Thresholds for Other TE Resources....................................................................................... 23
2.6.5 Configuring Alarm Thresholds for RSVP LSPs......................................................................................................... 24

3 MPLS LDP Configuration..................................................................................................... 26
3.1 Overview of MPLS LDP....................................................................................................................................................... 27
3.2 Understanding MPLS LDP.................................................................................................................................................. 28
3.2.1 Basic LDP Concepts........................................................................................................................................................... 28
3.2.2 LDP NSR................................................................................................................................................................................ 30
3.2.3 LDP MTU.............................................................................................................................................................................. 32
3.2.4 Coexistent Local and Remote LDP Session...............................................................................................................32
3.2.5 Distributing Labels to All Peers.................................................................................................................................... 33
3.3 Configuration Precautions for MPLS LDP..................................................................................................................... 34
3.4 Default Settings for MPLS LDP........................................................................................................................................ 34
3.5 Configuring Static LSPs....................................................................................................................................................... 35
3.5.1 Understanding Static LSPs............................................................................................................................................. 35
3.5.2 Configuring Static LSPs................................................................................................................................................... 36


Issue 01 (2025-03-03)                                    Copyright © Huawei Technologies Co., Ltd.                                                                                           ii
MPLS Configuration
MPLS Configuration                                                                                                                                                     Contents

