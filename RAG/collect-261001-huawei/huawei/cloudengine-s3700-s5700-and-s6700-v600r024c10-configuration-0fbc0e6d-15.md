---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-15
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["agent", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [1453, 1595]
sha256: c8558e0efae9a51b3955dae0917f1757e6b5bceaa67193a4c9cf4514f1814392
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        ● You must also run the snmp-agent trap enable feature-name mpls_lspm trap-name
                          { hwmplsresourcethresholdexceed | hwmplsresourcethresholdexceedclear }
                          command to enable the MPLS resource alarming function. Otherwise, when the LDP
                          resource usage reaches the upper threshold or falls below the lower threshold, such an
                          alarm is not generated or cleared, respectively.
                        ● After the snmp-agent trap enable feature-name mpls_lspm trap-name
                          { hwmplsresourcetotalcountexceed | hwmplsresourcetotalcountexceedclear }
                          command is run to enable the device to generate and clear an alarm related to the total
                          number of LDP resources:
                           –      An alarm is generated when the total number of LDP LSPs reaches the upper limit
                                  supported by the device.
                           –      The alarm is cleared when the total number of LDP LSPs falls below 95% of the
                                  upper limit.

                 ----End

2.6.4 Configuring Alarm Thresholds for Other TE Resources
Context
                 To facilitate device operation and maintenance, configure alarm thresholds for
                 other TE resources. This enables the device to report an alarm when the usage of
                 automatic bypass TE tunnel interfaces or automatic primary tunnels reaches the
                 upper limit and to clear the alarm when the usage falls below the lower limit.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Configure alarm thresholds for automatic bypass tunnel interfaces.
                 mpls autobypass-tunnel-number threshold-alarm upper-limit upper-limit-value lower-limit lower-limit-
                 value

                 In the preceding command:
                 ●       upper-limit-value specifies an alarm triggering threshold for the TE resource
                         usage.
                 ●       lower-limit-value specifies an alarm clearing threshold for the TE resource
                         usage.
                 ●       upper-limit-value must be greater than lower-limit-value.
                 By default, the alarm triggering threshold for the TE resource usage is 80%, and
                 the alarm clearing threshold for the TE resource usage alarm is 75%.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       23
MPLS Configuration
MPLS Configuration                                                                     2 Basic MPLS Configuration


                         NOTE

                        ● You must also run the snmp-agent trap enable feature-name mpls_lspm trap-name
                          { hwmplsresourcethresholdexceed | hwmplsresourcethresholdexceedclear }
                          command to enable the MPLS resource alarming function. Otherwise, when the TE
                          resource usage reaches the upper limit or falls below the lower limit, such an alarm is
                          not generated or cleared, respectively.
                        ● After the snmp-agent trap enable feature-name mpls_lspm trap-name
                          { hwmplsresourcetotalcountexceed | hwmplsresourcetotalcountexceedclear }
                          command is run to enable the device to generate and clear an alarm related to the total
                          number of TE resources:
                           –    An alarm is generated when the total number of used TE resources reaches the
                                upper limit supported by the device.
                           –    The alarm is cleared when the total number of used TE resources falls below 95%
                                of the upper limit.

                 ----End

2.6.5 Configuring Alarm Thresholds for RSVP LSPs

Context
                 To facilitate device operation and maintenance, configure alarm thresholds for
                 RSVP LSPs. This enables the device to report an alarm when the RSVP LSP usage
                 reaches the upper limit and to clear the alarm when the usage falls below the
                 lower limit.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Configure alarm thresholds for RSVP LSPs.
                 mpls rsvp-lsp-number threshold-alarm upper-limit upper-limit-value lower-limit lower-limit-value

                 In the preceding command:

                 ●      upper-limit-value specifies an alarm triggering threshold for the RSVP LSP
                        usage. If the proportion of established RSVP LSPs to total RSVP LSPs
                        supported reaches the specified threshold, an RSVP LSP alarm is generated.
                 ●      lower-limit-value specifies an alarm clearing threshold for the RSVP LSP
                        usage. If the proportion of established RSVP LSPs to total RSVP LSPs
                        supported falls below the specified threshold, the RSVP LSP alarm is
                        generated.
                 ●      upper-limit-value must be greater than lower-limit-value.

                 By default, the alarm triggering threshold for the RSVP LSP usage is 80%, and the
                 alarm clearing threshold for the RSVP LSP usage is 75%.



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          24
MPLS Configuration
MPLS Configuration                                                                 2 Basic MPLS Configuration


                         NOTE

                        ● You must also run the snmp-agent trap enable feature-name mpls_lspm trap-name
                          { hwmplslspthresholdexceed | hwmplslspthresholdexceedclear } command to enable
                          the RSVP LSP alarming function. Otherwise, when the RSVP LSP usage reaches the
                          upper limit or falls below the lower limit, such an alarm is not generated or cleared,
                          respectively.
                        ● After the snmp-agent trap enable feature-name mpls_lspm trap-name
                          { hwmplslsptotalcountexceed | hwmplslsptotalcountexceedclear } command is run
                          to enable the device to generate and clear alarms related to the total number of LDP
                          LSPs:
                           –    An alarm is generated when the total number of used RSVP LSPs reaches the
                                upper limit supported by the device.
                           –    The alarm is cleared when the total number of used RSVP LSPs falls below 95% of
                                the upper limit.

                 ----End




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    25
MPLS Configuration
MPLS Configuration                                                       3 MPLS LDP Configuration




                                   3           MPLS LDP Configuration


