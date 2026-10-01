---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-132
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [17002, 17139]
sha256: 0a6dfbdd05a8f9d6bfe9a0c7acf49d027dd47c24fba25ac41a3e1a03386ab5ac
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  ●      Periodically enabling the DIM function: In this mode, the device enables the
                         DIM function at a specified time every day to measure the running process
                         memory.
                  ●      Immediately enabling the DIM function: In this mode, the DIM function is
                         disabled after the measurement is complete. To perform DIM on the running
                         process memory again, re-execute the command for immediately enabling
                         the DIM function.

Procedure
                  ●      Enable the DIM function periodically.
                         system-view
                         trustem
                         dynamic-integrity-measurement daily time-value


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            310
Security Configuration
Security Configuration                                                           16 DIM Configuration


                         If the dynamic-integrity-measurement daily command is run more than
                         once, the latest configuration overrides the previous one.
                         The DIM function is triggered only when the CPU usage of the device falls
                         below 85%. If the CPU usage is greater than or equal to 85%, the DIM
                         function is triggered when the CPU usage falls below 85%.
                  ●      Enable the DIM function immediately.
                         system-view
                         trustem
                         start dynamic-integrity-measurement right-now

                  ----End




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          311
Security Configuration
Security Configuration                                            17 Remote Attestation Configuration




      17                 Remote Attestation Configuration


                  17.1 Overview of RA
                  17.2 Understanding RA
                  17.3 Configuration Precautions for Remote Attestation
                  17.4 Default Settings for RA
                  17.5 Configuring RA


17.1 Overview of RA
Definition
                  Trusted computing uses the Hardware Trust Module (HTM) as the root of trust
                  and measures the booted components level by level throughout the entire boot
                  process (from device power-on and BIOS boot, to GRUB and operating system
                  kernel loading). It then saves the measurement result in the Platform Configure
                  Registers (PCRs) of the HTM chip, and records the Storage Measurement Log
                  (SML).
                  Remote Attestation (RA) is one of the key technologies used in trusted computing.
                  RA calculates the actual status of a device based on the PCR values and SML
                  obtained during device boot, and compares the actual status with the reference
                  values in the baseline file to determine whether the device is trusted.

Purpose
                  Malicious actors might tamper with or replace communication devices and the
                  system software they run, compromising security if such devices are connected to
                  networks. In order to build a trusted network environment, each device must be
                  therefore verified to ensure its identity before being connected to the network. RA,
                  which is deployed independent of devices, enables customers to remotely audit
                  each device's trust status, improving network security by preventing untrusted
                  devices from accessing the network.




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                           312
Security Configuration
Security Configuration                                               17 Remote Attestation Configuration




17.2 Understanding RA
RA Fundamentals
                  The remote attestation (RA) system consists of the baseline file, RA server, RA
                  client, and CA. The administrator manages the entire RA system, including
                  downloading the baseline file and applying for and uploading the CA certificate,
                  as shown in Figure 1.


                  Figure 17-1 RA system architecture




                  ●      Baseline file: provided for the RA server. The baseline file contains the
                         reference values of measurement objects and serves as the reference baseline
                         for RA challenges. It is used to verify the SML obtained from the RA client for
                         verifying the trust status of the RA client. Baseline files are protected by
                         digital signatures and uploaded to the Huawei technical support website.
                         Baseline file name: product-name_version_RABASE.tar.gz.
                         The SML is generated by the RA client and records the actual hash value and
                         measurement sequence of a measurement object.
                  ●      RA server: provides an RA web UI for users to log in. The RA server, as the
                         core component of RA, sends a challenge request to the RA client, collects the
                         PCR value and SML of the RA client, and verifies the trust status of the RA
                         client based on the baseline file.
                  ●      RA client: refers to a network device whose trust status is to be verified. The
                         RA client is equipped with an HTM chip and supports trusted boot. It responds
                         to the challenge request from the RA server and sends its PCR values and SML
                         to the RA server.



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           313
Security Configuration
Security Configuration                                                17 Remote Attestation Configuration


                  ●      CA: creates and issues certificates. It is an authoritative organization trusted
                         by users. To authenticate the RA client and prevent it from being forged, the
                         CA issues an attestation key (AK) certificate to the RA client.
                  The RA client's AK certificate can be categorized into two types: initial attestation
                  key (IAK) and local attestation key (LAK).
                  ●      IAK: an AK set in the HTM chip before delivery. It is used to sign the data
                         (such as the PCR value) generated by the HTM. An IAK certificate is
                         configured on the device before delivery. Therefore, you do not need to apply
                         for a CA certificate.
                  ●      LAK: an AK created locally by a user after the device is delivered. It is used to
                         sign the data (such as the PCR value) generated by the HTM. The user can
                         then apply for an LAK certificate from the CA through the PKI.

