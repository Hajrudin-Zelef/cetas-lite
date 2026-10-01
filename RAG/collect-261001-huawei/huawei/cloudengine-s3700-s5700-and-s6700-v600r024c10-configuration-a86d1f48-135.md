---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-135
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "disclosure"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [17400, 17564]
sha256: fbcd5bbf60571240dce307f0c6f085eb9f8e2d7075db7605ce39d29d76c0cad0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                             You are advised to import the two sets of certificates, each with one Huawei root
                             certificate and one Huawei level-2 CA certificate. The same CA certificate cannot be
                             imported repeatedly.
                             If the current CA server does not verify the IAK certificate, you do not need to import
                             the Huawei level-2 CA certificate.
                  12. Bind the RA to a PKI CMP session.
                         trustem
                         remote-attestation pki bind cmp-session session-name

                         By default, the RA is not bound to a PKI CMP session.
                         session-name must be the same as the CMP session name in step Step 5.4.
                         If the certificate application is successful, the device automatically applies to
                         the CA server for updating the certificate when the certificate validity period
                         exceeds the 50% by default.
         Step 6 (Optional) Manually update the LAK certificate for the HTM.
                  remote-attestation pki update-request { all | slot slotID }

                  When the RA is configured, if a PKI certificate becomes invalid (for example, the
                  certificate is revoked due to certificate information disclosure), you must update
                  the PKI certificate. After this command is executed, the device immediately applies
                  to the CA server for certificate update.

                  ----End

Verifying the Configuration
                  Run the display htm status { slot slot-id | all } command to check the HTM chip
                  status.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      319
Security Configuration
Security Configuration                                                           18 HIPS Configuration




                                        18                  HIPS Configuration


                  18.1 Overview of HIPS
                  18.2 Understanding HIPS
                  18.3 Configuration Precautions for HIPS
                  18.4 Default Settings for HIPS
                  18.5 Enabling HIPS


18.1 Overview of HIPS
Definition
                  The Host-based Intrusion Prevention System (HIPS) monitors a device's system for
                  intrusions and infections. Unlike the Intrusion Prevention System (IPS) — which
                  analyzes and processes the traffic passing through a device to protect devices and
                  users on the internal network — HIPS protects the device's system.

Purpose
                  The security of network devices, which are important components of ICT
                  infrastructure, directly affects the security of the entire network. Network devices
                  are prone to hacker attacks and intrusions because they are usually deployed in
                  front of servers and terminals. After intruding into a network device, a hacker can
                  further penetrate the network through the device. To prevent this, HIPS is
                  introduced to monitor the device's operating system. Once a suspected intrusion
                  or infection event is detected, HIPS immediately sends a log to prompt the
                  administrator to isolate and protect the device, preventing further intrusions and
                  compromising the security of other devices.


18.2 Understanding HIPS
                  After intruding into the underlying operating system of a device, a hacker
                  configures and modifies the system for long-term control and further penetration.
                  HIPS monitors the underlying operating system of the device in real time and

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                            320
Security Configuration
Security Configuration                                                            18 HIPS Configuration


                  provides detection modules described in Table 18-1. Once a suspicious event is
                  detected, HIPS immediately sends the corresponding log.

                  Table 18-1 HIPS detection modules

                   Name           Description

                   File           After adding the SUID/SGID permission bit for executable files, a
                   privilege      user can run high-risk commands even if the user logs in to the
                   escalation     system as a common user later. HIPS sends the corresponding log
                   detection      when it detects that the SUID/SGID permission bit is added for
                                  executable files.

                   Abnormal       After intruding into a device successfully, a hacker may modify an
                   shell          existing shell of the device to facilitate the establishment of a
                   detection      control channel for a reverse shell. HIPS sends the corresponding
                                  log when it detects that a shell is modified.

                   Rootkit        A rootkit is a tool used by hackers to hide their tracks and retain
                   detection      root access during attacks. HIPS sends the corresponding log when
                                  it detects any system file that has rootkit characteristics on the
                                  device.

                   Key file       After a successful intrusion, a hacker may modify key files or leave
                   tampering      malicious files. HIPS sends the corresponding log when it detects
                   detection      that a key file is tampered with or a suspicious file exists in a key
                                  path.

                   Unauthoriz     Each user has a user identity (UID), and UID 0 is reserved for the
                   ed root        root user, making a non-root account with UID 0 highly suspicious.
                   user           HIPS sends the corresponding log when it detects a non-root
                   detection      account with UID 0.




18.3 Configuration Precautions for HIPS

18.4 Default Settings for HIPS
                  Table 18-2 describes the default settings for HIPS.

                  Table 18-2 Default settings for HIPS

                   Parameter                                  Default Setting

                   HIPS                                       Enabled

                   File privilege escalation detection        Enabled
                   module

                   Abnormal shell detection module            Enabled

                   Rootkit detection module                   Enabled


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           321
Security Configuration
Security Configuration                                                          18 HIPS Configuration


                   Parameter                                  Default Setting

                   Key file tampering detection module        Enabled

                   Unauthorized root user detection           Enabled
                   module




18.5 Enabling HIPS
Context
                  After HIPS is enabled, the configuration and enabling status of each detection
                  module are determined by the HIPS policy file. The policy file content cannot be
                  modified on the device and all detection modules are enabled by default.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable HIPS.
                  hips enable

                  ----End

