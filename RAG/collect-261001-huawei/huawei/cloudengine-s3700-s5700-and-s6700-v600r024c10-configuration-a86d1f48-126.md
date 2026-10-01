---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-126
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [16081, 16249]
sha256: 3e3dafaed6d7f710c878bdb9f7e10056db2f64f4ec2cd95dd27deceba37762b5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  Determine whether to perform this step based on the current interface working
                  mode.

         Step 4 Configure keychain authentication for RIP.
                  rip authentication-mode md5 nonstandard keychain keychain-name

         Step 5 Exit the interface view.
                  quit

                  ----End

15.5.4 Verifying the Keychain Configuration
Procedure
                  ●      Run the display keychain keychain-name command to check the keychain
                         configuration.
                  ●      Run the display keychain keychain-name key-id key-id command to check
                         the key configuration in the keychain.

                  ----End

15.5.5 Example for Configuring Keychain Authentication for
IS-IS
Networking Requirements
                  In Figure 15-6, DeviceA, DeviceB, and DeviceC communicate with each other
                  through IS-IS.

                  To ensure the stability and security of IS-IS connections, configure a keychain to
                  provide dynamic security authentication for IS-IS.

                  Figure 15-6 Keychain networking diagram
                          NOTE

                         In this example, interface1 and interface2 represent VLANIF1 and VLANIF2, respectively.




                  To complete the configuration, you need the following data:

                  ●      IS-IS process ID
                  ●      Network entity title (NET) of an IS-IS process
                  ●      keychain name

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        296
Security Configuration
Security Configuration                                                         15 Keychain Configuration


                  ●      Acceptance tolerance of a keychain
                  ●      Key ID in a keychain
                  ●      Key authentication algorithm and key string
                  ●      Send lifetime and accept lifetime of a key

Precautions
                  ●      NTP must be first configured.
                  ●      The configurations on both ends of keychain authentication must be
                         consistent. Take DeviceA and DeviceB as an example:
                         –   The keychain names configured on DeviceA and DeviceB must be the
                             same.
                         –   DeviceA and DeviceB must have the same time mode configured for the
                             keychains.
                         –   The key IDs in the keychains configured on DeviceA and DeviceB must be
                             the same. When multiple keys are configured, the same number of keys
                             with the same IDs must be configured on both ends.
                         –   For the same key, the same authentication algorithm and key string must
                             be configured on DeviceA and DeviceB.
                         –   For the same key, the send lifetime and accept lifetime configured on
                             DeviceA and DeviceB must match. For example, the accept lifetime
                             configured on DeviceB must include the send lifetime configured on
                             DeviceA to prevent packet loss. Similarly, the accept lifetime configured
                             on DeviceA also must include the send lifetime configured on DeviceB.
                  ●      If multiple keys are configured in a keychain, only one of them can be
                         configured as the default send key.

Configuration Roadmap
                  1.     Configure IS-IS.
                  2.     Create a keychain.
                  3.     Configure the key in the keychain and set the authentication algorithm of the
                         key ID to hmac-sha-256.
                  4.     Configure keychain authentication for IS-IS.

Procedure
         Step 1 Configure IS-IS.
                  # Configure DeviceA.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] isis 1
                  [DeviceA-isis-1] is-level level-1
                  [DeviceA-isis-1] network-entity 10.0000.0000.0001.00
                  [DeviceA-isis-1] quit
                  [DeviceA] interface 10ge 1/0/1
                  [DeviceA-10ge1/0/1] port link-type trunk
                  [DeviceA-10ge1/0/1] port trunk allow-pass vlan 1
                  [DeviceA-10ge1/0/1] quit
                  [DeviceA] vlan batch 1
                  [DeviceA] interface vlanif 1
                  [DeviceA-10GE1/0/1Vlanif1] ip address 192.168.1.1 24


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           297
Security Configuration
Security Configuration                                                       15 Keychain Configuration

                  [DeviceA-10GE1/0/1Vlanif1] isis enable 1
                  [DeviceA-10GE1/0/1Vlanif1] quit

                  # Configure DeviceB.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceB
                  [DeviceB] isis 1
                  [DeviceB-isis-1] is-level level-1
                  [DeviceB-isis-1] network-entity 10.0000.0000.0002.00
                  [DeviceB-isis-1] quit
                  [DeviceB] interface 10ge 1/0/1
                  [DeviceB-10ge1/0/1] port link-type trunk
                  [DeviceB-10ge1/0/1] port trunk allow-pass vlan 1
                  [DeviceB-10ge1/0/1] quit
                  [DeviceA] vlan batch 1 2
                  [DeviceB] interface vlanif 1
                  [DeviceB-10GE1/0/1Vlanif1] ip address 192.168.1.2 24
                  [DeviceB-10GE1/0/1Vlanif1] isis enable 1
                  [DeviceB-10GE1/0/1Vlanif1] quit
                  [DeviceB] interface 10ge 1/0/2
                  [DeviceB-10ge1/0/2] port link-type trunk
                  [DeviceB-10ge1/0/2] port trunk allow-pass vlan 2
                  [DeviceB-10ge1/0/2] quit
                  [DeviceB] interface vlanif 2
                  [DeviceB-10GE1/0/2Vlanif2] ip address 192.168.2.2 24
                  [DeviceB-10GE1/0/2Vlanif2] isis enable 1
                  [DeviceB-10GE1/0/2Vlanif2] quit

                  # Configure DeviceC.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceC
                  [DeviceC] isis 1
                  [DeviceC-isis-1] is-level level-1
                  [DeviceC-isis-1] network-entity 10.0000.0000.0003.00
                  [DeviceC-isis-1] quit
                  [DeviceC] interface 10ge 1/0/2
                  [DeviceC-10ge1/0/2] port link-type trunk
                  [DeviceC-10ge1/0/2] port trunk allow-pass vlan 2
                  [DeviceC-10ge1/0/2] quit
                  [DeviceA] vlan batch 2
                  [DeviceC] interface vlanif 2
                  [DeviceC-10GE1/0/2Vlanif2] ip address 192.168.2.1 24
                  [DeviceC-10GE1/0/2Vlanif2] isis enable 1
                  [DeviceC-10GE1/0/2Vlanif2] quit

         Step 2 Create a keychain.
                  # Configure DeviceA.
                  [DeviceA] keychain huawei mode absolute
                  [DeviceA-keychain-huawei] receive-tolerance 10
                  [DeviceA-keychain-huawei] quit

                  # Configure DeviceB.
                  [DeviceB] keychain huawei mode absolute
                  [DeviceB-keychain-huawei] receive-tolerance 10
                  [DeviceB-keychain-huawei] quit

                  # Configure DeviceC.
                  [DeviceC] keychain huawei mode absolute
                  [DeviceC-keychain-huawei] receive-tolerance 10
                  [DeviceC-keychain-huawei] quit

         Step 3 Configure a key in the keychain.
                  # Configure DeviceA.

