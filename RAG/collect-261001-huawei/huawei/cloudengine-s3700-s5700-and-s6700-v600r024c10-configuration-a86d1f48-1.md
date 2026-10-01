---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-1
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [1, 97]
sha256: 5c79fd8b79c063c19a887a414f96899bcddb5d3180cfc49b80507674da31d0a0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Security Configuration


Security Configuration

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
Security Configuration
Security Configuration                                                                                                                                          Contents




                                                                                                                                            Contents

1 About This Document.............................................................................................................1
2 Overview of Security.............................................................................................................. 5
3 Local Attack Defense Configuration................................................................................... 8
3.1 Overview of Local Attack Defense.................................................................................................................................... 8
3.2 Configuration Precautions for Local Attack Defense................................................................................................10
3.3 Default Settings for Local Attack Defense................................................................................................................... 10
3.4 Configuring CPU Attack Defense.....................................................................................................................................13
3.4.1 Understanding CPU Attack Defense........................................................................................................................... 13
3.4.2 Configuring the CPCAR Value....................................................................................................................................... 14
3.4.3 (Optional) Configuring Adaptive Adjustment of the Default CPCAR Value for Protocol Packets........ 15
3.4.4 (Optional) Configuring Packet Loss Monitoring for Protocol Packet Rate Limiting.................................. 17
3.4.5 (Optional) Configuring a Filter.....................................................................................................................................17
3.4.6 (Optional) Configuring Host Attack Defense.......................................................................................................... 18
3.4.7 Verifying the Configuration............................................................................................................................................18
3.4.8 Example for Configuring CPU Attack Defense........................................................................................................ 19
3.5 Configuring Port Attack Defense..................................................................................................................................... 21
3.5.1 Understanding Port Attack Defense........................................................................................................................... 21
3.5.2 Configuring Port Attack Defense................................................................................................................................. 22
3.5.3 Verifying the Configuration............................................................................................................................................24
3.6 Configuring User-Level Rate Limiting............................................................................................................................ 24
3.6.1 Understanding User-Level Rate Limiting.................................................................................................................. 24
3.6.2 Configuring User-Level Rate Limiting........................................................................................................................ 25
3.6.3 (Optional) Configuring Packet Loss Monitoring for User-Level Rate Limiting............................................ 26
3.6.4 Verifying the Configuration............................................................................................................................................26
3.7 Configuring Attack Source Tracing................................................................................................................................. 26
3.7.1 Understanding Attack Source Tracing........................................................................................................................ 27
3.7.2 Configuring Attack Source Tracing.............................................................................................................................. 27
3.7.3 Verifying the Configuration............................................................................................................................................30
3.7.4 Example for Configuring Attack Source Tracing..................................................................................................... 30
3.8 Configuring Defense Against Malformed Packet Attacks....................................................................................... 33
3.8.1 Understanding Defense Against Malformed Packet Attacks............................................................................. 33
3.8.2 Configuring Defense Against Malformed Packet Attacks................................................................................... 34


Issue 01 (2025-03-03)                               Copyright © Huawei Technologies Co., Ltd.                                                                                ii
Security Configuration
Security Configuration                                                                                                                                                   Contents

