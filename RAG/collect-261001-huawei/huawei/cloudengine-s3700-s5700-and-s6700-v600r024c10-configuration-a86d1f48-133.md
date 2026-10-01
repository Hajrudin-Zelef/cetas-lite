---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-133
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [17140, 17257]
sha256: 9ca3f93b441d9f29cd0bab326e96edae02f4576e11d902905338bb9248e7e31f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Working Process
                  In the trusted boot phase, the device's binary programs such as the BIOS,
                  BootLoader, and OS are measured during the boot process. A measurement is a
                  one-way hash calculation of measurement objects (such as software packages and
                  program files in software packages). The measurement result is saved in the PCRs
                  (PCR values) of the HTM chip, and the measurement sequence is also recorded in
                  the SML. The measurement result changes if either the measured file or
                  measurement sequence changes. RA calculates the actual status of a device based
                  on the PCR values and SML obtained during device boot, and compares the actual
                  status with the reference values in the baseline file to determine whether the
                  device is trusted.
                  In the trusted boot phase, the RA client calculates the hash value of each booted
                  component and saves the values to its PCRs. The RA server (challenger)
                  periodically sends a challenge to the RA client, requesting its PCR values and SML
                  saved during the trusted boot phase. Baseline files downloaded from https://
                  support.huawei.com are loaded on the RA server in advance. These baseline files
                  contain reference values of each measurement object. After receiving the PCR
                  values and SML from the RA client, the RA server uses the PCR value to verify the
                  SML. If the verification is successful, the RA server compares the actual hash
                  values of the measurement objects in the SML with the reference values in the
                  baseline file and determines the trust status of the device based on the
                  comparison result. If the actual hash values of all measurement objects in the SML
                  match those in the baseline file, the RA client is trusted. Otherwise, the RA client is
                  untrusted.

Verification Process
                  The AK certificate of the RA client can be IAK or LAK. If you use the LAK
                  certificate, you need to configure the RA client to obtain the LAK certificate first.
                  For details, see 17.5 Configuring RA.
                  The RA function between the RA server and RA client can be implemented only
                  after the administrator completes the following tasks:
                  1.     Downloads the baseline file, Huawei root certificate, and Huawei level-2 CA
                         certificate from the Huawei technical support website. If the RA client uses
                         the LAK certificate, you also need to obtain the CA certificate chain that issues
                         the LAK certificate.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             314
Security Configuration
Security Configuration                                                17 Remote Attestation Configuration


                  2.     Logs in to the RA server in web mode and uploads the files prepared in step 1
                         to the RA server.
                  3.     Adds the device information about the RA client to the RA server, including
                         the IP address and port number.
                  4.     Verifies the network connectivity between the RA client and RA server to
                         ensure that the former can be managed by the latter after being powered on.
                         The RA client and RA server communicate with each other through NETCONF.

                  When the RA client and RA server can reach each other over the network, the RA
                  server can use the RA function to verify the trust status of the RA client. Figure 2
                  illustrates the RA process.

                  Figure 17-2 RA process




                  1.     The RA server collects device information from the RA client, including
                         product, HTM vendor, and version information. The collected information is
                         used to manage the RA client and verify its trust status.
                  2.     The RA client sends its device information, including the product name and
                         version, device serial number, HTM specification version, HTM vendor, and
                         HTM status, to the RA server.
                  3.     The RA server sends a challenge request to the RA client and collects the PCR
                         value and AK certificate specified on the RA client.
                  4.     After receiving the challenge request from the RA server, the RA client
                         interacts with the HTM to obtain the corresponding PCR value and AK
                         certificate.
                         The RA client returns different certificate content to the RA server based on
                         the specified AK certificate type:
                         –   If the specified AK certificate type is IAK, the RA client returns the IAK
                             certificate to the RA server.
                         –   If the specified AK certificate type is LAK, the RA client returns the LAK
                             certificate and IAK certificate to the RA server.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                315
Security Configuration
Security Configuration                                                  17 Remote Attestation Configuration


                  5.     The RA client returns the AK certificate and PCR value to the RA server. The
                         RA server verifies the AK certificate and PCR value.
                         a.   Uses the trust certificate configured on the RA server to verify whether
                              the AK certificate is trusted. If the verification is successful, the RA server
                              proceeds to the next step. Otherwise, the RA server sets the RA client's
                              status to untrusted.
                         b.   Uses the public key in the AK certificate to verify the PCR value. If the
                              verification is successful, the RA server proceeds to the next step.
                              Otherwise, the RA server sets the RA client's status to untrusted.
                  6.     The RA server requests the SML recorded by the RA client.
                  7.     The RA client returns the recorded SML to the RA server, which then checks
                         whether the RA client is trusted based on the received SML:The RA server uses
                         the PCR value in step 5 to verify the SML. If the verification is successful, the
                         RA server compares the actual hash values of the measurement objects in the
                         SML with the reference values in the baseline files and determines the trust
                         status of the RA client based on the comparison result. If the actual hash
                         values of all measurement objects in the SML match those in the baseline file,
                         the RA client status is set to trusted. Otherwise, the RA client status is set to
                         untrusted.


17.3 Configuration Precautions for Remote Attestation

17.4 Default Settings for RA
                  Table 17-1 describes the default settings for RA.

                  Table 17-1 Default settings for RA

                   Parameter                                      Default Setting

                   RA                                             Disabled




