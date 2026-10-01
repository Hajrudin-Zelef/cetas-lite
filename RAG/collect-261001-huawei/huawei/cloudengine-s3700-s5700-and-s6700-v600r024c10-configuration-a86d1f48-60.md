---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-60
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [7206, 7341]
sha256: 5ab97439a1dc14e37b18f13b59702bbdc52d2e122157bcbe109ac0541c473d58
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  priority is elected as the key server. If the two devices have the same priority, their
                  Secure Channel Identifiers (SCIs) are compared. An SCI consists of the MAC
                  address of an interface and the last two bytes of the interface index. The device
                  with a smaller SCI value is elected as the key server.
                  To ensure secure transmission of service data on the network, MACsec provides
                  the data encryption and integrity check functions. You can specify an encryption
                  mode to selectively enable data encryption and integrity check. Three encryption
                  modes are available:
                  ●      None: Neither data encryption nor integrity check is performed.
                  ●      Normal: Both data encryption and integrity check are performed.
                  ●      Integrity-only: Integrity check is performed, and data encryption is not
                         performed.
                          NOTE

                         After a MACsec connection is established, the server can switch the encryption algorithm,
                         and the client also switches the encryption algorithm accordingly. However, when a Huawei
                         device functions as a server and a non-Huawei device that does not support encryption
                         algorithm switching functions as a client, encryption algorithm switching will cause the
                         interruption of MACsec negotiation.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create and configure a MACsec profile.
                  1.     Create a MACsec profile and enter the MACsec profile view.
                         mac-security-profile name profile-name

                  2.     (Optional) Configure the MKA key server priority.
                         mka keyserver priority priority

                         By default, the MKA key server priority is 16.
                  3.     (Optional) Configure the MACsec encryption mode.
                         macsec mode { none | normal | integrity-only }

                         By default, the MACsec encryption mode is normal.

                               NOTE

                              When configuring MACsec on a network where data traffic is being transmitted, you
                              can set the encryption mode on both ends to none. After MKA session negotiation is
                              successful, change the encryption mode on both ends to normal. This shortens the
                              traffic interruption time.

         Step 3 Return to the system view.
                  quit

         Step 4 Enter the view of the interface to which the MACsec profile is to be applied.
                  interface interface-type interface-number

         Step 5 Configure the CKN and CAK.
                  mka cak-mode static ckn string-name cak string-key

                  By default, no CKN or CAK is configured.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                  133
Security Configuration
Security Configuration                                                                    8 MACsec Configuration


                  A CKN is the name of a CAK. To ensure that an MKA session can be successfully
                  established between two devices, configure the same CKN and CAK on their
                  connected interfaces.
         Step 6 Apply the MACsec profile to the interface.
                  mac-security-profile profile-name

                          NOTE

                         The MACsec function takes effect only after the MACsec profile is applied to an interface.
                         For details about interfaces that support MACsec, see Configuration Precautions for
                         MACsec.

         Step 7 (Optional) Enable the MACsec module of the device to convert the byte order of
                XPN salt values.
                  macsec xpn-salt reverse

                  When network devices of other brands implement the MACsec function, the salt
                  values delivered in XPN mode use the network sequence or host sequence. You
                  can configure this command to implement compatibility among different
                  implementation modes.

                  ----End


8.6 Modifying MACsec Parameters

Context
                  You can create and configure a MACsec profile, or modify parameters in an
                  existing MACsec profile.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the MACsec profile view.
                  mac-security-profile name profile-name

         Step 3 Configure MACsec parameters.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      134
Security Configuration
Security Configuration                                                        8 MACsec Configuration


                  Table 8-5 Configuring MACsec parameters
                   Operation        Command                   Description

                   Configure an     mka cryptographic-        The device uses an MKA key
                   MKA key          algorithm { aes-          generation algorithm to generate the
                   generation       cmac-128 | aes-           KEK, ICK, and SAK based on the CAK
                   algorithm.       cmac-256 | sm4-           and CKN.
                                    cmac-128 }                By default, the MKA key generation
                                                              algorithm is AES-CMAC-128.
                                                              After generating the SAK, the key
                                                              server sends the encrypted SAK to the
                                                              peer device through an MKA packet.
                                                              Upon receipt of the MKA packet, the
                                                              peer device checks the integrity of the
                                                              packet. If the check fails, the peer
                                                              device drops the packet. If the check
                                                              succeeds, the peer device decrypts the
                                                              packet to obtain the original SAK. In
                                                              addition, ensure that the key
                                                              generation algorithms configured on
                                                              both ends are the same. Otherwise,
                                                              the negotiation may fail.
                                                              The sm4-cmac-128 algorithm is
                                                              supported only on the S6750-H and
                                                              S6780-H and S5755-S.

