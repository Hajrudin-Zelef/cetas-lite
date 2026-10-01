---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-1
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [1, 99]
sha256: 93db7dd11ea166a8162683632873569caf95daccbaab6fd806af219955306db7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

QoS Configuration


QoS Configuration

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
QoS Configuration
QoS Configuration                                                                                                                                                             Contents




                                                                                                                                                        Contents

1 About This Document.............................................................................................................1
2 Overview of QoS...................................................................................................................... 5
3 MQC Configuration................................................................................................................. 9
3.1 Overview of MQC.................................................................................................................................................................... 9
3.2 Understanding MQC............................................................................................................................................................ 10
3.3 Configuration Precautions for MQC............................................................................................................................... 12
3.4 Configuring a Traffic Classifier......................................................................................................................................... 12
3.5 Configuring a Traffic Behavior..........................................................................................................................................17
3.6 Configuring a Traffic Policy............................................................................................................................................... 19
3.7 Applying a Traffic Policy..................................................................................................................................................... 19
3.8 Verifying the Configuration............................................................................................................................................... 21
3.9 Maintaining MQC................................................................................................................................................................. 22

4 Packet Filtering Configuration........................................................................................... 23
4.1 Overview of Packet Filtering............................................................................................................................................. 23
4.2 Configuration Precautions for Packet Filtering...........................................................................................................23
4.3 Configuring MQC-based Packet Filtering..................................................................................................................... 23
4.4 Example for Configuring MQC-based Packet Filtering............................................................................................ 27
4.5 Example for Configuring Access Control Based on Source MAC Addresses..................................................... 29

5 Traffic Statistics Collection Configuration...................................................................... 33
5.1 Overview of Traffic Statistics Collection........................................................................................................................33
5.2 Configuration Precautions for Traffic Statistics Collection..................................................................................... 34
5.3 Configuring MQC-based Traffic Statistics Collection................................................................................................ 34
5.4 Example for Configuring MQC-based Traffic Statistics Collection.......................................................................37

6 Re-marking Configuration.................................................................................................. 40
6.1 Overview of Re-marking.................................................................................................................................................... 40
6.2 Re-marking..............................................................................................................................................................................40
6.3 Configuration Precautions for Re-marking.................................................................................................................. 42
6.4 Configuring MQC-based Priority Re-marking............................................................................................................. 42
6.5 Example for Configuring Re-marking to Distinguish Users................................................................................... 45
6.6 Example for Configuring Re-marking to Distinguish Services.............................................................................. 49

7 Redirection Configuration................................................................................................... 53

Issue 01 (2025-03-03)                                   Copyright © Huawei Technologies Co., Ltd.                                                                                          ii
QoS Configuration
QoS Configuration                                                                                                                                                           Contents

