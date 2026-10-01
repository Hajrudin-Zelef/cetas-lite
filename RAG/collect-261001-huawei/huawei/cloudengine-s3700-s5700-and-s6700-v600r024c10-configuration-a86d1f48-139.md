---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-139
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "license"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [18076, 18257]
sha256: cb38dd39e2f534d2320237646187d141b5f02cee94ee4971413f304ae0869fc9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  These packets consume device resources, causing the device performance to
                  deteriorate or even device breakdown. In this case, to prevent potential malicious
                  attacks, you can disable an EVC Layer 2 sub-interface from broadcasting received
                  packets to other EVC Layer 2 sub-interfaces in the same BD.

                  This function applies to networks without user changes or networks with static
                  MAC address-based forwarding paths.

                          NOTE

                         This configuration is supported only by the S6780-H, S5755-H, S5755E-H, S5755-S, S6750-
                         H, S6730-H-V2, S6730E-H-V2, S6750-S, S6750E-S and S5732-H-V2 series.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the BD view.
                  bridge-domain bd-id

         Step 3 Disable an interface from forwarding unknown multicast packets to other
                interfaces in the same BD.
                  unknown-multicast discard

                  ----End



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  334
Security Configuration
Security Configuration                                                23 Layer 2 Traffic Suppression Configuration




23.4 Configuring VLAN-based Layer 2 Traffic
Suppression
                  VLAN-based traffic suppression prevents excessive traffic from burdening the
                  network.

Context
                  Traffic suppression can be implemented only on a Layer 2 interface.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the interface view.
                  interface interface-type interface-number

         Step 3 Switch the interface working mode to Layer 2.
                  portswitch

                  This step is supported only on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                  S6750E-S, S6750-S, S5755-H, S5755E-H, S5755-S and S5732-H-V2. Determine
                  whether to perform this step based on the current interface working mode.
         Step 4 Return to the system view.
                  quit

         Step 5 Create a specified VLAN.
                  vlan batch { vlan-id1 [ to vlan-id2 ] } &<1-10>

         Step 6 Add interfaces to the VLAN.
                  port interface-type { interface-number1 [ to interface-number2 ] } &<1-10>

         Step 7 Disable interfaces in the VLAN from forwarding unknown multicast packets.
                  unknown-multicast discard

                  ----End




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   335
Security Configuration
Security Configuration                                                 24 Trusted System Configuration




                24                   Trusted System Configuration


                  24.1 Overview of Trusted Systems
                  24.2 Configuration Precautions for Trusted Systems
                  24.3 Digital Signature of Software Packages
                  24.4 Secure Boot
                  24.5 Trusted Boot


24.1 Overview of Trusted Systems
Definition
                  To enhance system trustworthiness, you can utilize the trusted system to
                  efficiently oversee and regulate digital signature, remote attestation, and secure
                  boot functions on a device.
                  A trusted system is one in which the hardware and software are functioning as
                  intended and designed. To ensure a system is trustworthy, it is essential that the
                  software integrity is of the highest level, preventing any unauthorized
                  modifications or intrusions.

Purpose
                  The proper functioning of communication networks is crucial for global
                  communications service providers, enterprises, and governments. Also, the
                  integrity of data and IT infrastructure is fundamental to upholding network
                  security and fostering user trust. In today's dynamic threat landscape,
                  safeguarding networks against intrusion, forgery, and tampering by malicious
                  actors has become increasingly important.
                  A communication device comprises several embedded computer systems, which
                  are susceptible to virus attacks, tampering, and exploitation of vulnerabilities by
                  malicious actors through the use of Trojan horses. Once an untrusted device
                  accesses the network, the entire network may be compromised. Therefore, to
                  create a trusted network environment, it is essential to ensure that all connected
                  devices are trustworthy. By utilizing digital signature, remote attestation, and

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                            336
Security Configuration
Security Configuration                                                   24 Trusted System Configuration


                  secure boot, users can guarantee the trustworthiness of all network-connected
                  devices.

Benefits
                  Trusted system functions include digital signature, remote attestation, and secure
                  boot, which bring the following security benefits:

                  ●      Trust in the trusted boot platform can be established by utilizing device
                         hardware capabilities and initial boot code.
                  ●      Secure boot allows only trusted devices to boot.


24.2 Configuration Precautions for Trusted Systems
Licensing Requirements
                  Trusted system configuration is not under license control.

Hardware Requirements
                  All products support the trusted system function.

Feature Requirements
                  None


24.3 Digital Signature of Software Packages
Definition
                  The digital signature mechanism ensures the validity and integrity of software
                  packages, which ultimately enhances the security and availability of installed
                  software.

Purpose
                  After a software package is released, the subsequent phases of transfer, download,
                  storage, and installation pose potential security risks, including the possibility of
                  components being replaced or tampered with. Before a software package is
                  released, a digital signature is embedded within it, and this signature is verified
                  during the loading process onto a device. The software package is considered
                  complete and trusted, and applications can be installed, only after the verification
                  succeeds.

Benefits
                  When you set up a patch or system software package for the next boot or install a
                  patch, the digital signature is verified to ensure the integrity of the system
                  software package or patch.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           337
Security Configuration
Security Configuration                                                 24 Trusted System Configuration




24.4 Secure Boot
Context
                  A communication device comprises several embedded computer systems, which
                  are susceptible to virus attacks, tampering, and exploitation of vulnerabilities by
                  malicious actors through the use of Trojan horses.
                  The secure boot function prohibits system boot in case of any damage to the
                  system software, improving system security and reliability.

