---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-136
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [15063, 15179]
sha256: 243089a99fbaa2a86d140b5ce2575933c0eb7e1618f80c681302ba32fd42fb09
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Context
                    By default, devices are preconfigured with service awareness signature databases.
                    The Huawei Security Center (https://isecurity.huawei.com/) periodically releases
                    the latest service awareness signature database. To help devices effectively identify
                    applications, update devices' service awareness signature databases to the latest
                    version.
                    Devices support the update of the service awareness signature databases used by
                    an AC and AP, as described in Table 13-4.

                    Table 13-4 Introduction to service awareness signature databases
                     Signature            Signature            Description
                     Database Name        Database Alias

                     SA signature         SA-SDB               The service awareness signature
                     database                                  database used by a device that does
                                                               not support the native AC function
                                                               describes the signatures of common
                                                               applications/protocols on the network.
                                                               These signatures are used by
                                                               application identification to identify
                                                               applications/protocols in traffic.
                                                               NOTE
                                                                This signature database is supported only
                                                                on the S6730E-H-V2, S6730-H-V2, S5732-
                                                                H-V2, S5755E-H, S5755-H, S6750E-S,
                                                                S5755-S, and S6750-S.

                     APSA signature       APSA-SDB             In a native AC scenario, the device
                     database                                  functions as an AC to manage the AP,
                                                               which installs and updates this
                                                               database through the AC. The
                                                               procedure is as follows:
                                                               1. The AC downloads or updates the
                                                                  service awareness signature
                                                                  database used by the AP.
                                                               2. The AC traverses the service
                                                                  awareness signature database
                                                                  support and version of the online
                                                                  AP, and then transfers the latest
                                                                  signature database file to the AP.
                                                               3. The AP loads the service awareness
                                                                  signature database.
                                                               NOTE
                                                                This signature database is supported only
                                                                on the S6730E-H-V2, S6730-H-V2, S5732-
                                                                H-V2, S5755E-H, and S5755-H.




                    This section provides only the basic steps for updating a signature database online.
                    For details about offline signature database update and other related update

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               277
QoS Configuration
QoS Configuration                                                              13 Experience Assurance Configuration


                    information, see "Signature Database Update Configuration" in CLI Configuration
                    Guide > System Management Configuration.

Procedure
         Step 1 Configure the DNS server to ensure that the device can correctly parse the domain
                name of the security center.
                    system-view
                    dns resolve
                    dns server ip-address

         Step 2 Update a service awareness signature database.

                    The device supports online and offline updates of service awareness signature
                    databases. Online updates are classified as automatic or manual online update.

                    ●    Automatically update a service awareness signature database online.
                         a.    Enable the automatic update function.
                               update schedule { sa-sdb | apsa-sdb } enable

                         b.    Configure an automatic update time.
                               update schedule sa-sdb { daily | weekly { mon | tue | wed | thu | fri | sat | sun } } time
                               update schedule apsa-sdb { hourly minute | { daily | weekly { mon | tue | wed | thu | fri | sat
                               | sun } } time }

                               The scheduled update time can be manually configured.
                               It is recommended that a service awareness signature database be
                               updated once a week.
                    ●    Manually update a service awareness signature database online.
                         update online { sa-sdb | apsa-sdb }

                         After the command is executed, the device immediately queries and
                         downloads the signature database, and then updates the local signature
                         database to the latest version.
                    ●    Update a service awareness signature database offline.
                         a.    Log in to Huawei Security Center (isecurity.huawei.com) and download
                               the latest service awareness signature database file.
                         b.    Upload the signature database file to the device.
                               The signature database file is in .zip format and does not need to be
                               decompressed. You can upload the file to the root directory (flash:/) of
                               the device through FTP, SFTP, or TFTP. For details, see "File System
                               Management Configuration" in CLI Configuration Guide > Basic
                               Configuration.
                         c.    Use the signature database file uploaded to the root directory for offline
                               update.
                               update local { sa-sdb | apsa-sdb } file file-name

                    ----End

Verifying the Configuration
                    ●    Run the display update configuration command to check whether the
                         configuration for the signature database update is correct.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                              278
QoS Configuration
QoS Configuration                                                                   13 Experience Assurance Configuration


