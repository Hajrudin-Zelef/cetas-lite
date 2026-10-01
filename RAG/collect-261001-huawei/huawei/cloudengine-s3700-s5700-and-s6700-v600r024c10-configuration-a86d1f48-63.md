---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-63
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [7542, 7653]
sha256: 8a602fc26c870906a5850bdada2d1dfb0cd3b7306abf93ce4a861416bcfd2ed3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   Configure the      macsec capability         Meanings of MACsec capability
                   MACsec             capability-value          values:
                   capability                                   ● 2: Both integrity check and
                   value.                                         confidentiality check are
                                                                  supported. Encryption can only
                                                                  start from the packet header.
                                                                ● 3: Both integrity check and
                                                                  confidentiality check are
                                                                  supported. The encryption offset
                                                                  can be 0 bytes (encryption starting
                                                                  from the packet header), 30 bytes,
                                                                  or 50 bytes.
                                                                By default, the MACsec capability
                                                                value is 3.
                                                                If the local and peer device types are
                                                                different and the peer device supports
                                                                only MACsec capability 2, you need to
                                                                set the MACsec capability value of
                                                                the local device to be the same as
                                                                that of the peer device.
                                                                When the MACsec capability value is
                                                                set to 2, setting the MACsec
                                                                encryption offset to 30 or 50 using
                                                                the macsec confidentiality-offset
                                                                command does not take effect. In this
                                                                case, the MACsec encryption offset is
                                                                always 0 (encryption starting from
                                                                the packet header).

                   Configure          macsec policy { must-     Security policy description:
                   security           secure | should-          ● must-secure: Discard all non-
                   policies on        secure }                    cipher packets after MKA session
                   MACsec.                                        negotiation starts. After the
                                                                  negotiation succeeds, packets are
                                                                  encrypted and decrypted normally.
                                                                ● should-secure: Permits non-cipher
                                                                  packets before MKA session
                                                                  negotiation succeeds. After the
                                                                  negotiation succeeds, non-cipher
                                                                  packets are discarded and packets
                                                                  are encrypted and decrypted
                                                                  normally.
                                                                By default, the MACsec security policy
                                                                is must-secure.



                  ----End



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                          140
Security Configuration
Security Configuration                                                           8 MACsec Configuration




8.7 Verifying the Configuration

Procedure
         Step 1 Run the display macsec port ability [ slot slot-id ] command to check whether
                the port on the corresponding board supports MACsec.
                  Only the S6750-H series devices support this command.
         Step 2 Run the display mac-security-profile configuration [ name profile-name ]
                command to check the MACsec profile configuration.
         Step 3 Run the display mka interface { interface-name | interface-type interface-
                number } [ verbose ] command to check the MACsec configuration and MKA
                session information on a specified interface.
         Step 4 Run the display macsec statistics interface { interface-name | interface-type
                interface-number } command to check statistics about data packets protected by
                MACsec on a specified interface.

                  ----End


8.8 Example for Configuring Device-to-Device MACsec
Networking Requirements
                  In Figure 8-3, DeviceA and DeviceB are connected and exchange important
                  information, which needs to be protected.

                  Figure 8-3 Network diagram of device-to-device MACsec
                          NOTE

                         In this example, interface 1 represents 10GE 1/0/1.




Configuration Roadmap
                  When MACsec is configured on two devices, the configuration roadmap is as
                  follows:
                  1.     Configure the key server priority. In this example, DeviceA is configured as the
                         key server.
                  2.     Set the encryption mode to normal, indicating that both data encryption and
                         integrity check are enabled.
                  3.     Set the CKN and CAK for MKA session negotiation to
                         f1c3b2a4d6d9a7c5b4e1ab56dc21ed79ac97be533671dcab2678ac55cf71aced
                         and ab2145369adcadef69512347adceb210, respectively.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          141
Security Configuration
Security Configuration                                                                   8 MACsec Configuration


                          NOTE




