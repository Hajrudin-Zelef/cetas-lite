---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-1
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [1, 98]
sha256: 086bb6184009cf545f0ec54bd45f6a361ddf7cca3501e88993b96299ac8e9243
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration


VPN Configuration

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
VPN Configuration
VPN Configuration                                                                                                                                                              Contents




                                                                                                                                                         Contents

1 About This Document.............................................................................................................1
2 GRE Configuration................................................................................................................... 5
2.1 Overview of GRE...................................................................................................................................................................... 5
2.2 Understanding GRE................................................................................................................................................................ 6
2.2.1 GRE Fundamentals.............................................................................................................................................................. 6
2.2.2 Keepalive Detection.......................................................................................................................................................... 10
2.3 Configuration Precautions for GRE................................................................................................................................. 11
2.4 Configuring a GRE Tunnel..................................................................................................................................................11
2.4.1 Configuring a Tunnel Interface..................................................................................................................................... 11
2.4.2 (Optional) Enabling the Keepalive Function........................................................................................................... 12
2.4.3 Configuring Tunnel Routes.............................................................................................................................................13
2.4.4 Verifying the Configuration............................................................................................................................................14
2.4.5 Example for Configuring an IPv4 over IPv4 GRE Tunnel..................................................................................... 15
2.4.6 Example for Configuring an IPv6 over IPv4 GRE Tunnel..................................................................................... 19
2.4.7 Example for Enlarging the Operation Scope of a Network with a Hop Limit............................................. 24
2.5 Maintaining GRE................................................................................................................................................................... 30
2.5.1 Monitoring the Operating Status of GRE.................................................................................................................. 30
2.5.2 Enabling GRE Traffic Statistics Collection and Checking GRE Traffic Statistics............................................31
2.5.3 Resetting Keepalive Message Statistics on Tunnel Interfaces............................................................................31

3 IPv4 L3VPN Configuration...................................................................................................32
3.1 Overview of IPv4 L3VPN.................................................................................................................................................... 32
3.2 Understanding IPv4 L3VPN............................................................................................................................................... 33
3.2.1 Basic Concepts of IPv4 L3VPN...................................................................................................................................... 33
3.2.2 VPN NSR............................................................................................................................................................................... 39
3.2.3 Inter-AS VPN....................................................................................................................................................................... 40
3.2.4 Label Allocation Modes of IPv4 L3VPN over MPLS............................................................................................... 41
3.3 Configuration Precautions for IPv4 L3VPN.................................................................................................................. 44
3.4 Default Settings for IPv4 L3VPN...................................................................................................................................... 44
3.5 Configuring Mutual Access Between Local IPv4 L3VPNs........................................................................................44
3.5.1 Configuring an IPv4 VPN Instance on a PE.............................................................................................................. 44
3.5.2 Binding an Interface to an IPv4 VPN Instance........................................................................................................ 46
3.5.3 Verifying the Configuration............................................................................................................................................48


Issue 01 (2025-03-03)                                    Copyright © Huawei Technologies Co., Ltd.                                                                                          ii
VPN Configuration
VPN Configuration                                                                                                                                                   Contents

