---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-90
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [11070, 11227]
sha256: b5150df7e41ecdffc8770f2c7e2bb12c12c91af445aa870296d2d81a835de2d1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

         Step 3 Enter the certificate attribute group view. If no certificate attribute group exists,
                create one first.
                  pki certificate attribute-group group-name

         Step 4 Configure certificate attribute conditions.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      203
Security Configuration
Security Configuration                                                                             10 PKI Configuration


                   To...                                                Run...

                   Configure the start time and end time                attribute id validity from begintime
                   of the certificate validity period                   begindate to endtime enddate
                   Configure the FQDN                                   attribute id alt-subject-name fqdn
                                                                        { ctn | equ | nctn | nequ } attribute-
                                                                        value
                   Configure the certificate IP address                 attribute id alt-subject-name ip { ctn
                                                                        | equ | nctn | nequ } ip-address

                   Configure the certificate issuer name                attribute id issuer-name dn { ctn |
                                                                        equ | nctn | nequ } attribute-value

                   Configure the certificate subject name               attribute id subject-name dn { ctn |
                                                                        equ | nctn | nequ } attribute-value



         Step 5 Return to the system view.
                  quit

         Step 6 Enter the certificate attribute-based access control policy view. If no certificate
                attribute-based access control policy exists, create one first.
                  pki certificate access-control-policy name policy-name

                  By default, no certificate attribute-based access control policy is created.

         Step 7 Configure a certificate attribute-based control rule.
                  rule id { permit | deny } group-name

                  By default, no certificate attribute-based control rule is configured.

         Step 8 Configure the description of the certificate attribute-based access control policy.
                  description description

                  By default, a certificate attribute-based access control policy does not have a
                  description.

         Step 9 Adjust the sequence of certificate attribute-based control rules.
                  pki certificate access-control-policy [ policy-name policy-name ] rule move rule-id1 { before | after }
                  rule-id2

                  When you adjust the sequence of such rules, the sequence of rule IDs (rule-id) is
                  unchanged, but the rule contents are swapped. For example:

                  The certificate attribute-based access control policy a has the following rules:
                  pki certificate access-control-policy name a
                   rule 5 permit test1
                   rule 20 permit test2

                  After the pki certificate access-control-policy policy-name a rule move 20
                  before 5 command is executed, the rules are changed as follows:
                  pki certificate access-control-policy name a
                   rule 5 permit test2
                   rule 20 permit test1

        Step 10 Exit the certificate attribute-based access control policy view.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                             204
Security Configuration
Security Configuration                                                               10 PKI Configuration

                  quit

                  ----End

Verifying the Configuration
                  ●      Run the display pki certificate access-control-policy all command in the
                         user view to check information of all the access control policies.
                  ●      Run the display pki certificate attribute-group all command in the user
                         view to check information about all certificate attribute groups.

10.9.3 Configuring a Certificate Whitelist to Implement Access
Control

Prerequisites
                  Certificate whitelist files need to be stored in advance under the flash:/pki/public
                  directory of the device.

Context
                  A certificate whitelist contains common names (CNs) or serial numbers (SNs) of
                  base station certificates. When the local device receives a certificate authentication
                  request from the peer device and if the CN or SN of the peer certificate is
                  whitelisted on the local device, the certificate authentication succeeds.

                  To make the PKI certificate whitelist check function take effect, import certificate
                  whitelist files to the device memory.

Procedure
                  ●      Import a certificate whitelist file to the device memory.
                         system-view
                         pki import whitelist filename file-name

                  ●      Delete a certificate whitelist file.
                         system-view
                         pki delete whitelist filename file-name

                  ----End

Verifying the Configuration
                  Run the display pki whitelist { all | filename file-name } command to check the
                  content of certificate whitelist files on the device.

10.9.4 Example for Configuring Certificate Attribute-based
Filtering to Implement Access Control

Networking Requirements
                  In Figure 10-16, DeviceA functions as the gateway of network C (an enterprise
                  network). DeviceA and those on networks A and B authenticate each other using

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          205
Security Configuration
Security Configuration                                                                         10 PKI Configuration


                  certificates. Devices on networks A and B can set up connections with DeviceA and
                  access resources on network C after successful certificate authentication.

                  A certificate attribute-based access control policy is required to allow certificates
                  with the following certificate attributes to pass authentication:

                  ●      The certificate issuer name is networkb_ca.
                  ●      The certificate subject name is cert_ca.

                  Figure 10-16 Network diagram of configuring certificate attribute-based filtering
                  to implement access control




                          NOTE

                         This example provides only the configurations related to certificate attribute-based access
                         control policies.


Configuration Roadmap
                  The configuration roadmap is as follows:

                  1.     Create a certificate attribute group and specify the attributes issuer name and
                         subject name.
                  2.     Create a certificate attribute-based access control policy and allow certificates
                         matching the specified attributes to pass authentication.

