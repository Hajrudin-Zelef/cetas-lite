---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-131
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2019-12-10", "2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [16834, 17001]
sha256: 39768546365b83e27945cc156ba65b8f963032acfd1bfd7aa06525274760b78d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Configuration Scripts
                  ●      DeviceA
                         #
                         sysname DeviceA
                         #
                         vlan batch 1
                         #
                         keychain huawei mode absolute
                          receive-tolerance 10
                          tcp-kind 182
                          tcp-algorithm-id hmac-sha-256 17
                          #
                          key-id 1
                           algorithm hmac-sha-256
                           key-string cipher %+%#1h29-c>>[H,XTu>Q}##;"}JOQOK#c>TD6>~d-BaJ%+%#
                           send-time 12:00 2019-12-10 to 15:00 2019-12-10
                           receive-time 12:00 2019-12-10 to 15:00 2019-12-10
                           default send-key-id
                          #
                          key-id 2
                           algorithm hmac-sha-256
                           key-string cipher %+%#^<Sn.IK2iK'N%[VnMhv-I)|C4d<K$F$a.6%jEN@K%+%#
                           send-time 15:05 2019-12-10 to 18:00 2019-12-10
                           receive-time 15:05 2019-12-10 to 18:00 2019-12-10
                         #
                         interface vlanif 1
                          ip address 192.168.1.1 255.255.255.0
                         #
                         bgp 1
                          router-id 1.1.1.1
                          peer 192.168.1.2 as-number 1
                          peer 192.168.1.2 keychain huawei
                          #
                          ipv4-family unicast
                           peer 192.168.1.2 enable
                         #
                         return
                  ●      DeviceB
                         #
                         sysname DeviceB
                         #
                         vlan batch 2
                         #
                         keychain huawei mode absolute
                          receive-tolerance 10
                          tcp-kind 182
                          tcp-algorithm-id hmac-sha-256 17
                          #
                          key-id 1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                             307
Security Configuration
Security Configuration                                                             15 Keychain Configuration

                           algorithm hmac-sha-256
                           key-string cipher %+%#p8cb/;OMFES0Wx@PY^"Ka{6q2MB;oG|[ZO-_]u}&%+%#
                           send-time 12:00 2019-12-10 to 15:00 2019-12-10
                           receive-time 12:00 2019-12-10 to 15:00 2019-12-10
                           default send-key-id
                          #
                          key-id 2
                           algorithm hmac-sha-256
                           key-string cipher %+%#&Yq4=s*P:L<"8iG-|o1ZB*Qi0qCn%N{Y3a&Z-zuD%+%#
                           send-time 15:05 2019-12-10 to 18:00 2019-12-10
                           receive-time 15:05 2019-12-10 to 18:00 2019-12-10
                         #
                         interface vlanif 2
                          ip address 192.168.1.2 255.255.255.0
                         #
                         bgp 1
                          router-id 2.2.2.2
                          peer 192.168.1.1 as-number 1
                          peer 192.168.1.1 keychain huawei
                          #
                          ipv4-family unicast
                           peer 192.168.1.1 enable
                         #
                         return




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              308
Security Configuration
Security Configuration                                                         16 DIM Configuration




                                         16                DIM Configuration


                  16.1 Overview of DIM
                  16.2 Configuration Precautions for DIM
                  16.3 Default Settings for DIM
                  16.4 Enabling DIM


16.1 Overview of DIM

Definition
                  Dynamic Integrity Measurement (DIM) is a code integrity measurement
                  technology. It measures the code segment integrity of the running process
                  memory and detects malicious attacks on the running process memory in a timely
                  manner. The running process memory includes process code segment, shared
                  library code segment, kernel, and kernel module.

Purpose
                  Malicious attacks on the running process memory are highly covert and cannot be
                  detected through common methods such as trusted boot and file integrity check.
                  In addition, more and more malware is used to attack the memory. Therefore, it is
                  a must to improve the security detection capability of the device to detect
                  malicious attacks on the running process memory in a timely manner. DIM can
                  detect whether the running process memory is tampered with and whether the
                  memory is changed due to injection attacks. Trusted boot only ensures that the
                  device is trusted or detects untrusted boot behavior during device startup. File
                  integrity check can detect whether static files are tampered with. The combination
                  of DIM, trusted boot and file integrity check can ensure the reliability of the
                  system in the entire running process.

Fundamentals
                  DIM can measure the running process memory whose content (process code
                  segment, shared library code segment, kernel, and kernel module) remains

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                         309
Security Configuration
Security Configuration                                                                 16 DIM Configuration


                  unchanged. The principle of the DIM is to periodically or proactively detect a
                  measurement object, calculate its actual hash value, and match the hash value
                  with the baseline value preset in the software package. If the two values are
                  inconsistent, the measurement object may be tampered with. Then the system
                  saves the hash value to the PCRs of the HTM and records the corresponding
                  stored measurement log (SML). The baseline value preset in the software package
                  stores the reference hash value of the measurement object. The remote
                  attestation (RA) server periodically initiates RA challenges to the device, obtains
                  the PCR values and measurement logs generated during DIM on the device,
                  verifies them based on the baseline file imported to the RA server, and generates a
                  DIM trust report for the device.


16.2 Configuration Precautions for DIM

16.3 Default Settings for DIM
                  Table 16-1 describes the default settings for DIM.

                  Table 16-1 Default settings for DIM

                   Parameter                                         Default Setting

                   DIM function                                      Disabled




16.4 Enabling DIM

Context
                  The DIM function must be used together with RA. The DIM result is obtained by
                  the RA server, which then generates the DIM trust report of the device.

                  You can enable the DIM function in either of the following ways:

