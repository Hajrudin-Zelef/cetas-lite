---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-14
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["agent", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [1320, 1452]
sha256: 7e870a6ba3424a3a625fd39993c2c6936b3c4a966f91fa8bc2a4982ea95f13bd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 In the preceding command:
                 ●      upper-limit-value specifies an upper threshold at which the device generates
                        an LDP LSP usage (proportion of established LDP LSPs to total LDP LSPs
                        supported) alarm.
                 ●      lower-limit-value specifies a lower threshold below which the device clears an
                        LDP LSP usage alarm.
                 ●      upper-limit-value must be greater than lower-limit-value.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         20
MPLS Configuration
MPLS Configuration                                                                    2 Basic MPLS Configuration


                 By default, the upper alarm threshold for the LDP LSP usage is 80%, and the lower
                 alarm threshold for the LDP LSP usage alarm is 75%.

                         NOTE

                        ● You must also run the snmp-agent trap enable feature-name mpls_lspm trap-name
                          { hwmplslspthresholdexceed | hwmplslspthresholdexceedclear } command to enable
                          the LDP LSP alarming function. Otherwise, when the LDP LSP usage reaches the upper
                          limit or falls below the lower limit, such an alarm is not generated or cleared,
                          respectively.
                        ● After the snmp-agent trap enable feature-name mpls_lspm trap-name
                          { hwmplslsptotalcountexceed | hwmplslsptotalcountexceedclear } command is run
                          to enable the device to generate and clear alarms related to the total number of LDP
                          LSPs:
                           –    An alarm is generated when the total number of LDP LSPs reaches the upper limit
                                supported by the device.
                           –    The alarm is cleared when the total number of LDP LSPs falls below 95% of the
                                upper limit.

                 ----End

2.6.2 Configuring Alarm Thresholds for Dynamic Labels

Context
                 When a certain number of dynamic labels are consumed, new dynamic label
                 requests may fail to be satisfied due to insufficient resources. As a result, the
                 module that applies for dynamic labels cannot run properly. To facilitate device
                 operation and maintenance, configure alarm thresholds for dynamic labels. This
                 enables the device to generate an alarm when the dynamic label usage reaches
                 the upper threshold and to clear the alarm when the dynamic label usage falls
                 below the lower threshold.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Configure alarm thresholds for dynamic labels.
                 mpls dynamic-label-number threshold-alarm upper-limit upper-limit-value lower-limit lower-limit-value

                 In the preceding command:

                 ●      upper-limit-value specifies an upper threshold at which the device generates a
                        dynamic label usage alarm. You are advised not to specify a value greater
                        than 95% for this parameter.
                 ●      lower-limit-value specifies a lower threshold below which the device clears a
                        dynamic label usage alarm.
                 ●      upper-limit-value must be greater than lower-limit-value.
                 By default, the upper threshold is 80%, and the lower threshold is 70%.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         21
MPLS Configuration
MPLS Configuration                                                                     2 Basic MPLS Configuration


                         NOTE

                        ● You must also run the snmp-agent trap enable feature-name mpls_lspm trap-name
                          { hwMplsDynamicLabelThresholdExceed |
                          hwMplsDynamicLabelThresholdExceedClear } command to enable the dynamic label
                          alarming function. Otherwise, when the dynamic label usage reaches the upper limit or
                          falls below the lower limit, such an alarm is not generated or cleared, respectively.
                        ● After the snmp-agent trap enable feature-name mpls_lspm trap-name
                          { hwMplsDynamicLabelTotalCountExceed |
                          hwMplsDynamicLabelTotalCountExceedClear } command is run to enable the device
                          to generate and clear alarms related to the total number of dynamic labels:
                           –    An alarm is generated when the total number of dynamic labels reaches the upper
                                limit supported by the device.
                           –    The alarm is cleared when the total number of dynamic labels falls below 95% of
                                the upper limit.

                 ----End

2.6.3 Configuring Alarm Thresholds for Other LDP Resources

Context
                 To facilitate device operation and maintenance, configure alarm thresholds for
                 other LDP resources (such as remote LDP adjacencies and outgoing segments),
                 enabling the device to report an alarm when the usage of a specified type of LDP
                 resource reaches the upper threshold and to clear the alarm when the usage falls
                 below the lower threshold.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Configure alarm thresholds for other LDP resources as needed.
                 ●      Configure alarm thresholds for remote LDP adjacencies.
                        mpls remote-adjacency-number threshold-alarm upper-limit upper-limit-value lower-limit lower-
                        limit-value
                 ●      Configure alarm thresholds for outgoing segments.
                        mpls outsegment-number threshold-alarm upper-limit upper-limit-value lower-limit lower-limit-
                        value

                 In the preceding command:

                 ●      upper-limit-value specifies an upper threshold at which the device generates
                        an LDP resource usage alarm.
                 ●      lower-limit-value specifies a lower threshold below which the device clears an
                        LDP resource usage alarm.
                 ●      upper-limit-value must be greater than lower-limit-value.
                 By default, the upper alarm threshold is 80%, and the lower alarm threshold is
                 75%.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          22
MPLS Configuration
MPLS Configuration                                                                    2 Basic MPLS Configuration


                         NOTE

