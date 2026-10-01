---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-323
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [48112, 48280]
sha256: c9e19e7e1842f477c6f9ed07d43fe9bb51b4a622ebc1667632d300cc59b1de78
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 4 Configure VPLS.
                    # Configure PE1.
                    [PE1] vsi s1 static
                    [PE1-vsi-s1] pwsignal ldp
                    [PE1-vsi-s1-ldp] vsi-id 10
                    [PE1-vsi-s1-ldp] peer 3.3.3.3
                    [PE1-vsi-s1-ldp] quit
                    [PE1-vsi-s1] quit
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] portswitch
                    [PE1-10GE1/0/1] port link-type trunk
                    [PE1-10GE1/0/1] quit
                    [PE1] interface 10ge 1/0/1.1
                    [PE1-10GE1/0/1.1] shutdown
                    [PE1-10GE1/0/1.1] dot1q termination vid 10
                    [PE1-10GE1/0/1.1] l2 binding vsi s1
                    [PE1-10GE1/0/1.1] undo shutdown
                    [PE1-10GE1/0/1.1] quit

                    # Configure PE2.
                    [PE2] vsi s1 static
                    [PE2-vsi-s1] pwsignal ldp
                    [PE2-vsi-s1-ldp] vsi-id 10
                    [PE2-vsi-s1-ldp] peer 3.3.3.3
                    [PE2-vsi-s1-ldp] quit
                    [PE2-vsi-s1] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] portswitch
                    [PE2-10GE1/0/1] port link-type trunk
                    [PE2-10GE1/0/1] quit
                    [PE2] interface 10ge 1/0/1.1
                    [PE2-10GE1/0/1.1] shutdown
                    [PE2-10GE1/0/1.1] dot1q termination vid 10
                    [PE2-10GE1/0/1.1] l2 binding vsi s1
                    [PE2-10GE1/0/1.1] undo shutdown
                    [PE2-10GE1/0/1.1] quit

                    # Configure PE3.
                    [PE3] vsi s1 static
                    [PE3-vsi-s1] pwsignal ldp
                    [PE3-vsi-s1-ldp] vsi-id 10
                    [PE3-vsi-s1-ldp] peer 1.1.1.1
                    [PE3-vsi-s1-ldp] peer 2.2.2.2
                    [PE3-vsi-s1-ldp] quit
                    [PE3-vsi-s1] quit
                    [PE3] interface 10ge 1/0/1
                    [PE3-10GE1/0/1] portswitch
                    [PE3-10GE1/0/1] port link-type trunk
                    [PE3-10GE1/0/1] quit
                    [PE3] interface 10ge 1/0/1.1
                    [PE3-10GE1/0/1.1] shutdown
                    [PE3-10GE1/0/1.1] dot1q termination vid 10
                    [PE3-10GE1/0/1.1] l2 binding vsi s1
                    [PE3-10GE1/0/1.1] undo shutdown
                    [PE3-10GE1/0/1.1] quit

         Step 5 Configure ERPS on PE1, PE2, CE1, and CE2.
                    # Configure PE1.
                    [PE1] erps ring 1
                    [PE1-erps-ring1] control-vlan 100


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   775
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                    [PE1-erps-ring1] protected-instance 1
                    [PE1-erps-ring1] version v2
                    [PE1-erps-ring1] sub-ring
                    [PE1-erps-ring1] quit
                    [PE1] stp region-configuration
                    [PE1-mst-region] instance 1 vlan 10 100
                    [PE1-mst-region] quit
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] portswitch
                    [PE1-10GE1/0/1] port link-type trunk
                    [PE1-10GE1/0/1] undo port trunk allow-pass vlan 1
                    [PE1-10GE1/0/1] stp disable
                    [PE1-10GE1/0/1] erps ring 1
                    [PE1-10GE1/0/1] erps vpls-subinterface enable
                    [PE1-10GE1/0/1] quit

                    # Configure PE2.
                    [PE2] erps ring 1
                    [PE2-erps-ring1] control-vlan 100
                    [PE2-erps-ring1] protected-instance 1
                    [PE2-erps-ring1] version v2
                    [PE2-erps-ring1] sub-ring
                    [PE2-erps-ring1] quit
                    [PE2] stp region-configuration
                    [PE2-mst-region] instance 1 vlan 10 100
                    [PE2-mst-region] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] portswitch
                    [PE2-10GE1/0/1] port link-type trunk
                    [PE2-10GE1/0/1] undo port trunk allow-pass vlan 1
                    [PE2-10GE1/0/1] stp disable
                    [PE2-10GE1/0/1] erps ring 1
                    [PE2-10GE1/0/1] erps vpls-subinterface enable
                    [PE2-10GE1/0/1] quit

                    # Configure CE1.
                    <Device> system-view
                    [Device] sysname CE1
                    [CE1] erps ring 1
                    [CE1-erps-ring1] control-vlan 100
                    [CE1-erps-ring1] protected-instance 1
                    [CE1-erps-ring1] version v2
                    [CE1-erps-ring1] sub-ring
                    [CE1-erps-ring1] quit
                    [CE1] stp region-configuration
                    [CE1-mst-region] instance 1 vlan 10 100
                    [CE1-mst-region] quit
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] portswitch
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] undo port trunk allow-pass vlan 1
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [CE1-10GE1/0/1] stp disable
                    [CE1-10GE1/0/1] erps ring 1
                    [CE1-10GE1/0/1] quit
                    [CE1] interface 10ge 1/0/2
                    [CE1-10GE1/0/2] portswitch
                    [CE1-10GE1/0/2] port link-type trunk
                    [CE1-10GE1/0/2] undo port trunk allow-pass vlan 1
                    [CE1-10GE1/0/2] port trunk allow-pass vlan 10
                    [CE1-10GE1/0/2] stp disable
                    [CE1-10GE1/0/2] erps ring 1
                    [CE1-10GE1/0/2] quit

                    # Configure CE2.
                    <Device> system-view
                    [Device] sysname CE2
                    [CE2] erps ring 1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   776
VPN Configuration
VPN Configuration                                                                        6 VPLS Configuration

                    [CE2-erps-ring1] control-vlan 100
                    [CE2-erps-ring1] protected-instance 1
                    [CE2-erps-ring1] version v2
                    [CE2-erps-ring1] sub-ring
                    [CE2-erps-ring1] quit
                    [CE2] stp region-configuration
                    [CE2-mst-region] instance 1 vlan 10 100
                    [CE2-mst-region] quit
                    [CE2] interface 10ge 1/0/1
                    [CE2-10GE1/0/1] portswitch
                    [CE2-10GE1/0/1] port link-type trunk
                    [CE2-10GE1/0/1] undo port trunk allow-pass vlan 1
                    [CE2-10GE1/0/1] port trunk allow-pass vlan 10
                    [CE2-10GE1/0/1] stp disable
                    [CE2-10GE1/0/1] erps ring 1
                    [CE2-10GE1/0/1] quit
                    [CE2] interface 10ge 1/0/2
                    [CE2-10GE1/0/2] portswitch
                    [CE2-10GE1/0/2] port link-type trunk
                    [CE2-10GE1/0/2] undo port trunk allow-pass vlan 1
                    [CE2-10GE1/0/2] port trunk allow-pass vlan 10
                    [CE2-10GE1/0/2] stp disable
                    [CE2-10GE1/0/2] erps ring 1 rpl owner
                    [CE2-10GE1/0/2] quit

         Step 6 Verify the configuration.

