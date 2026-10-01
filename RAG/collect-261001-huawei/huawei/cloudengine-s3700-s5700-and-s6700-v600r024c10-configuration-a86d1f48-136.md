---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-136
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [17565, 17740]
sha256: 848004f5a69b78c217cbde9c501f969181bf1536e167173cf0e2b615d509427c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Verifying the Configuration
                  Run the display hips state command in any view to check the status of each HIPS
                  detection module.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                        322
Security Configuration
Security Configuration                                                          19 FIPS Configuration




                                          19                FIPS Configuration


                  19.1 Overview of FIPS
                  19.2 Configuration Precautions for FIPS
                  19.3 Enabling the FIPS Mode


19.1 Overview of FIPS
Definition
                  Federal Information Processing Standards (FIPS) are security requirements
                  standards issued by the National Institute of Standards and Technology (NIST) for
                  cryptographic modules. The device supports FIPS 140-2, which is the latest version
                  of FIPS. In this document, FIPS refers to FIPS 140-2.


Purpose
                  FIPS specifies the security requirements that a cryptographic module in a security
                  system must meet to ensure the confidentiality and integrity of information
                  protected by the module. You can configure FIPS to ensure that the device passes
                  FIPS certification, improving device security.


19.2 Configuration Precautions for FIPS

19.3 Enabling the FIPS Mode
Context
                  With FIPS mode enabled, the device has stricter security requirements and checks
                  whether the algorithms used by cryptographic modules comply with FIPS. This
                  ensures the proper running of cryptographic modules.


Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                           323
Security Configuration
Security Configuration                                                                       19 FIPS Configuration


                          NOTE

                         Enabling the FIPS mode on the device will clear the configuration file for next startup and
                         immediately restart the device. Therefore, exercise caution when enabling the FIPS mode.
                         After the FIPS mode is enabled, the weak security algorithm/protocol feature package can
                         be installed on the device, but does not take effect. That is, the weak security algorithm/
                         protocol cannot be used.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable the FIPS mode.
                  fips-mode enable

                  By default, the FIPS mode is disabled on the device.

                  ----End

Verifying the Configuration
                  ●      Run the display fips-mode command in the system view to check whether
                         the FIPS mode is enabled on the device.
                  ●      Run the display fips-mode algorithm self-check command in the system
                         view to check whether the algorithms provided by cryptographic modules of
                         the device comply with FIPS.
                  ●      Run the display fips-mode finite-state command in the system view to
                         check the historical records of FIPS mode status changes.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       324
Security Configuration
Security Configuration                                                          20 GTSM Configuration




                                     20                  GTSM Configuration


                  20.1 Overview of GTSM
                  20.2 Configuration Precautions for GTSM
                  20.3 Enabling GTSM
                  20.4 (Optional) Configuring an Action to Process Packets That Do Not Match a
                  GTSM Policy


20.1 Overview of GTSM
Definition
                  The Generalized TTL Security Mechanism (GTSM), a protection mechanism based
                  on the time to live (TTL) value, checks whether the TTL value in the IP packet
                  header is within a pre-defined range and discards invalid packets to protect
                  TCP/IP-based control plane protocols from CPU overload attacks.

Purpose
                  An attacker simulates a routing protocol and continuously sends packets to a
                  device. If the device cannot determine the validity of packets, the device is busy
                  processing attack packets, resulting in CPU overload attacks.
                  In this case, a method is required to check the validity of packets. GTSM can be
                  used to check the TTL value in the packet header.
                  The main function of TTL values is to prevent IP packets from being circulated
                  over a network in infinite loops. The maximum TTL value is 255. The TTL value is
                  decremented by 1 each time a packet passes through a hop. Since the number of
                  hops between any two routing neighbors is limited by the network scale and
                  structure, the TTL values of protocol packets exchanged between the devices are
                  confined to a particular range.
                  Based on network conditions, the TTL value range used between routing
                  neighbors can be predefined. In this way, devices can check the validity of TTL
                  values in packets to determine packet validity and filter out invalid packets (attack
                  packets).

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                                325
Security Configuration
Security Configuration                                                         20 GTSM Configuration


Fundamentals
                  As shown in Figure 20-1:
                  ●      IGP connections are established to implement connectivity between devices.
                  ●      A BGP connection is established between DeviceA and DeviceB.
                  ●      An attacker remotely accesses the network from the Internet, simulates BGP
                         negotiation packets, and continuously sends packets to DeviceA or DeviceB.
                  When DeviceA and DeviceB negotiate a BGP peer relationship, one path can be
                  selected from three for packet forwarding. The number of hops (including the last
                  hop) on the selected path through which the packets pass may be 3, 5, or 6. That
                  is, a maximum of six hops are possible.
                  In this situation, GTSM can be used to predefine the TTL value range as [255 – 6 +
                  1, 255], that is, [250, 255]. Remote BGP attack packets whose TTL values are not
                  within the specified range are considered invalid and are dropped.

                  Figure 20-1 GTSM Attack Defense Fundamentals




                  Although planning and configuring TTL value ranges become complex on a
                  complex network, you can define a generally appropriate range based on network
                  conditions to allow GTSM to filter out attack packets as many as possible, since
                  their TTL values are out of the specified range.


20.2 Configuration Precautions for GTSM




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         326
Security Configuration
Security Configuration                                                                         20 GTSM Configuration




20.3 Enabling GTSM
Context
                  Without GTSM enabled on a device, the device sends protocol packets directly to
                  the control plane. After GTSM is enabled for a protocol:

