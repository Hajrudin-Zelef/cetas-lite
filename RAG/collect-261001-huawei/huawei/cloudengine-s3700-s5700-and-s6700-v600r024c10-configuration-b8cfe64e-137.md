---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-137
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [15180, 15304]
sha256: 4a42259e1b0864fbfd7b60339c25ba636d7414f53c9d52157aeef69a6fccc775
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    ●    After the signature database is updated, run the display engine information
                         command to check the running status of the ASE and version information of
                         the signature database. Ensure that the signature database has been updated
                         to the latest version.

13.5.2 Configuring MQC-based Experience Assurance

Context
                    After the experience assurance function is configured, the device re-marks the
                    priority of packets matching an application-based policy to control network traffic.

Procedure
         Step 1 Enable the SA function.
                    1.   Enter the interface view.
                         interface interface-type interface-number
                    2.   Enable the SA function on the interface.
                         sa enable
                         By default, the SA function is disabled. After this function is enabled, the
                         device can identify application traffic passing through the interface.
         Step 2 (Optional) Configure an application identification whitelist.
                    sa whitelist acl { acl-number | acl-name }

                    By default, no application identification whitelist is configured.

                          NOTE

                         An advanced ACL must have been created and configured.

         Step 3 Configure a traffic classifier.
                    1.   Create a traffic classifier and enter the traffic classifier view, or enter the view
                         of an existing traffic classifier.
                         traffic classifier classifier-name [ type { and | or } ]
                    2.   Configure application identification.
                         if-match application application-name [ time-range time-name ]
                    3.   Exit the traffic classifier view.
                         quit

         Step 4 Configure a traffic behavior.
                    1.   Create a traffic behavior and enter the traffic behavior view, or enter the view
                         of an existing traffic behavior.
                         traffic behavior behavior-name
                    2.   Configure a matching rule as required.
                         –      Re-marking 802.1p values of VLAN packets
                                remark 8021p 8021p-value
                         –      Re-marking DSCP values of IP packets
                                remark dscp { dscp-name | dscp-value }

         Step 5 Configure a traffic policy.
                    1.   Create a traffic policy and enter the traffic policy view, or enter the view of an
                         existing traffic policy.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                       279
QoS Configuration
QoS Configuration                                                            13 Experience Assurance Configuration

                         traffic policy policy-name

                    2.   Bind a traffic behavior and a traffic classifier to the traffic policy.
                         classifier classifier-name behavior behavior-name [ precedence precedence-value ]

                    3.   Exit the traffic policy view.
                         quit

         Step 6 Apply the traffic policy to an interface.
                    1.   Enter the interface view.
                         interface interface-type interface-number

                    2.   Apply the traffic policy to the interface.
                         traffic-policy policy-name { inbound | outbound }

                    3.   Exit the interface view.
                         quit

                    ----End

13.5.3 (Optional) Configuring Global Parameters for
Experience Assurance
Context
                    In the system view, you can configure global parameters for the experience
                    assurance function. Table 13-5 lists the global parameters that can be configured.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure global parameters for experience assurance.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                  280
QoS Configuration
QoS Configuration                                                  13 Experience Assurance Configuration


                    Table 13-5 Global parameters for experience assurance
                     Operation               Command                Description                        Ap
                                                                                                       plic
                                                                                                       ati
                                                                                                       on
                                                                                                       Sce
                                                                                                       nar
                                                                                                       io

                     Enable the              sa application-        After the application              To
                     application             statistic enable       identification statistics          obt
                     identification                                 collection function is             ain
                     statistics collection                          enabled, you can run the           pac
                     function.                                      display sa flow-table              ket
                                                                    command to query the               stat
                                                                    traffic statistics of identified   istic
                                                                    applications.                      s of
                                                                                                       an
                                                                                                       app
                                                                                                       lica
                                                                                                       tion
                                                                                                       ,
                                                                                                       use
                                                                                                       this
                                                                                                       fun
                                                                                                       ctio
                                                                                                       n.

