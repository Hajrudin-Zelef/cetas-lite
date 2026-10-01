---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-140
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [18258, 18354]
sha256: bafdc27de69fc70006d46b896afd1ebfb2181bf930ae050cfd3884b00a1705ea
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Related Concepts
                  A trusted system is one in which the hardware and software are functioning as
                  intended and designed. To ensure a system is trustworthy, it is essential that the
                  software integrity is of the highest level, preventing any unauthorized
                  modifications or intrusions.

Fundamentals
                  Most boards that support secure boot have it enabled by default to ensure a
                  trusted system. However, for certain boards, manual activation of secure boot is
                  required.
                  Secure boot establishes a root of trust (RoT) for the platform by utilizing device
                  hardware capabilities, unalterable initial boot code, and signature verification keys.
                  As shown in Figure 24-1, during the system boot process, the trust root, BIOS,
                  Bootloader, OS kernel, and system software package are booted in that order, with
                  each level measuring the file trustworthiness of the next level. If the file
                  trustworthiness of one component fails the verification, the component cannot be
                  booted.
                  If the BIOS verification step or any preceding steps fail, the device must be
                  returned to the manufacturer for resolution. If the BIOS verification passes, but
                  any subsequent verification fails, you can enter the BootROM menu by pressing
                  Ctrl+B. From there, you can replace the system software package and reboot the
                  device to resolve the issue.

                  Figure 24-1 Secure boot process




Benefits
                  The secure boot function brings the following security benefits:

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                             338
Security Configuration
Security Configuration                                                      24 Trusted System Configuration


                  ●      Trust the software system of the device that can be properly booted.
                  ●      Prohibit the boot of the device on which the software system is detected as
                         untrustworthy.


24.5 Trusted Boot
                          NOTE

                         This feature is supported only by S6730-H24X6C-TV2, S6730-H48X6C-TV2, S6730-H48X6CZ-
                         TV2, S6730-H28X6CZ-TV2, S6730-H48Y6C-TV2, S5732-H48UM4Y2CZ-TV2, S5732-
                         H24UM4Y2CZ-TV2, S5732-H44S4X6QZ-TV2, S5732-H24S4X6QZ-TV2, S5735I-S24T4XE-T-V2,
                         S5735I-S24U4XE-T-V2, S5735I-S8T4XN-T-V2, S5735-L10T4X-TA-V2, S5735-L24P4XE-TA-V2,
                         S5735-L48T4XE-TA-V2, S5735-L8P2T4X-TA-V2, S5735-L14P2S-TQA-V2, S5755-H24UM4Y2CZ-
                         T, S5755-H48UM4Y2CZ-T, S5755-H24UTM4X4Y2C-T and S5755-H48UTM4X4Y2C-T.


Background
                  Communication devices consist of multiple embedded computer systems. The
                  software that runs on these devices may be vulnerable to viruses, modified by
                  attackers, or implanted with Trojan horses.
                  The trusted boot function promptly detects issues that affect the trusted status of
                  the system, helping improve the security and reliability of the system.

Related Concepts
                  Trusted system: A trusted system indicates that system hardware and software are
                  running as designed. The prerequisite for a trusted system is that the system
                  software integrity is high and free of intrusion or unauthorized modification.

Basic Principles
                  The trusted boot function establishes an RoT for the trusted boot platform based
                  on the hardware capabilities of a device and an initial boot code.
                  During the boot process, the system establishes a complete trust chain from the
                  RoT, BIOS, and BootLoader, to the OS kernel and system software package, with
                  each level measuring the boot phase of the next level. The measurement results
                  are irrevocably saved to the HTM. This implementation ensures:
                  ●      Setup and transmission of the trust chain.
                  ●      Recording of the system's trusted status.

Benefits
                  This feature offers the following security benefits:
                  ●      Software integrity measurement
                         Measures the integrity of the software during the boot process, establishes
                         and transfers a chain of trust, and records the system's trusted status.
                  ●      Trusted status query
                         Provides query of the trusted status of the system.
                  ●      Trusted status alarm

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               339
Security Configuration
Security Configuration                                                  24 Trusted System Configuration


                         Generates an alarm if the trusted status of the system is abnormal.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          340

