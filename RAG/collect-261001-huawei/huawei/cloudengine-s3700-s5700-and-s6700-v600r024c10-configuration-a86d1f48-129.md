---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-129
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2019-12-10", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [16575, 16711]
sha256: a5ffb0df3ef472572f06d3b6fb4762ade68614e61d58c09dc82d0e424c81dbcc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  Figure 15-7 Keychain networking diagram
                          NOTE

                         In this example, interface1 represents VLANIF1.




                  To complete the configuration, you need the following data:

                  ●      keychain name
                  ●      Acceptance tolerance of a keychain
                  ●      Values of TCP Kind and TCP Alg ID fields in the TCP Enhanced Authentication
                         Option
                  ●      Key ID in a keychain
                  ●      Key authentication algorithm and key string
                  ●      Send lifetime and accept lifetime of a key

Precautions
                  ●      NTP and BGP must be first configured.
                  ●      The keychain names configured on DeviceA and DeviceB must be the same.
                  ●      DeviceA and DeviceB must have the same time mode configured for the
                         keychains.
                  ●      The key IDs in the keychains configured on DeviceA and DeviceB must be the
                         same. When multiple keys are configured, the same number of keys with the
                         same IDs must be configured on both ends.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        303
Security Configuration
Security Configuration                                                                  15 Keychain Configuration


                  ●      For the same key, the same authentication algorithm and key string must be
                         configured on DeviceA and DeviceB.
                  ●      For the same key, the send lifetime and accept lifetime configured on DeviceA
                         and DeviceB must match. For example, the accept lifetime configured on
                         DeviceB must include the send lifetime configured on DeviceA to prevent
                         packet loss. Similarly, the accept lifetime configured on DeviceA also must
                         include the send lifetime configured on DeviceB.
                  ●      If multiple keys are configured in a keychain, only one of them can be
                         configured as the default send key.

Configuration Roadmap
                  1.     Create a keychain.
                  2.     Configure the key in the keychain and set the authentication algorithm of the
                         key ID to hmac-sha-256.
                  3.     Configure keychain authentication for BGP.

Procedure
         Step 1 Create a keychain.
                  # Configure DeviceA.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] keychain huawei mode absolute
                  [DeviceA-keychain-huawei] receive-tolerance 10
                  [DeviceA-keychain-huawei] tcp-kind 182
                  [DeviceA-keychain-huawei] tcp-algorithm-id hmac-sha-256 17
                  [DeviceA-keychain-huawei] quit

                  # Configure DeviceB.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceB
                  [DeviceB] keychain huawei mode absolute
                  [DeviceB-keychain-huawei] receive-tolerance 10
                  [DeviceB-keychain-huawei] tcp-kind 182
                  [DeviceB-keychain-huawei] tcp-algorithm-id hmac-sha-256 17
                  [DeviceB-keychain-huawei] quit

         Step 2 Configure a key in the keychain.
                  # Configure DeviceA.
                  [DeviceA] keychain huawei
                  [DeviceA-keychain-huawei] key-id 1
                  [DeviceA-keychain-huawei-keyid-1] algorithm hmac-sha-256
                  [DeviceA-keychain-huawei-keyid-1] key-string cipher YsHsjx_202207
                  [DeviceA-keychain-huawei-keyid-1] send-time 12:00 2019-12-10 to 15:00 2019-12-10
                  [DeviceA-keychain-huawei-keyid-1] receive-time 12:00 2019-12-10 to 15:00 2019-12-10
                  [DeviceA-keychain-huawei-keyid-1] default send-key-id
                  [DeviceA-keychain-huawei-keyid-1] quit
                  [DeviceA-keychain-huawei] key-id 2
                  [DeviceA-keychain-huawei-keyid-2] algorithm hmac-sha-256
                  [DeviceA-keychain-huawei-keyid-2] key-string cipher YsHsjx_202206
                  [DeviceA-keychain-huawei-keyid-2] send-time 15:05 2019-12-10 to 18:00 2019-12-10
                  [DeviceA-keychain-huawei-keyid-2] receive-time 15:05 2019-12-10 to 18:00 2019-12-10
                  [DeviceA-keychain-huawei-keyid-2] quit
                  [DeviceA-keychain-huawei] quit

                  # Configure DeviceB.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                    304
Security Configuration
Security Configuration                                                                  15 Keychain Configuration

                  [DeviceB] keychain huawei
                  [DeviceB-keychain-huawei] key-id 1
                  [DeviceB-keychain-huawei-keyid-1] algorithm hmac-sha-256
                  [DeviceB-keychain-huawei-keyid-1] key-string cipher YsHsjx_202207
                  [DeviceB-keychain-huawei-keyid-1] send-time 12:00 2019-12-10 to 15:00 2019-12-10
                  [DeviceB-keychain-huawei-keyid-1] receive-time 12:00 2019-12-10 to 15:00 2019-12-10
                  [DeviceB-keychain-huawei-keyid-1] default send-key-id
                  [DeviceB-keychain-huawei-keyid-1] quit
                  [DeviceB-keychain-huawei] key-id 2
                  [DeviceB-keychain-huawei-keyid-2] algorithm hmac-sha-256
                  [DeviceB-keychain-huawei-keyid-2] key-string cipher YsHsjx_202206
                  [DeviceB-keychain-huawei-keyid-2] send-time 15:05 2019-12-10 to 18:00 2019-12-10
                  [DeviceB-keychain-huawei-keyid-2] receive-time 15:05 2019-12-10 to 18:00 2019-12-10
                  [DeviceB-keychain-huawei-keyid-2] quit
                  [DeviceB-keychain-huawei] quit

         Step 3 Configure keychain authentication for BGP.
                  # Configure DeviceA.
                  [DeviceA] vlan batch 1
                  [DeviceA] interface vlanif 1
                  [DeviceA-Vlanif1] ip address 192.168.1.1 24
                  [DeviceA-Vlanif1] quit
                  [DeviceA] bgp 1
                  [DeviceA-bgp] router-id 1.1.1.1
                  [DeviceA-bgp] peer 192.168.1.2 as-number 1
                  [DeviceA-bgp] peer 192.168.1.2 keychain huawei
                  [DeviceA-bgp] quit
                  [DeviceA] quit

                  # Configure DeviceB.
                  [DeviceA] vlan batch 2
                  [DeviceB] interface vlanif 2
                  [DeviceB-Vlanif2] ip address 192.168.1.2 24
                  [DeviceB-Vlanif2] quit
                  [DeviceB] bgp 1
                  [DeviceB-bgp] router-id 2.2.2.2
                  [DeviceB-bgp] peer 192.168.1.1 as-number 1
                  [DeviceB-bgp] peer 192.168.1.1 keychain huawei
                  [DeviceB-bgp] quit
                  [DeviceB] quit

                  ----End

