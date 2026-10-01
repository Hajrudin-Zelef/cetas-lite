---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-122
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2019-12-10", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [15495, 15611]
sha256: 3ed92c85959ad2f0871572ba6a8149ef8babcf6e7c6d7fb636d20e8dcd69cd9f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Security Configuration
Security Configuration                                                                 15 Keychain Configuration


Key ID and Category
                  To better manage keys collectively on a device, the keychain defines key IDs to
                  facilitate differentiation. Keys on the device are classified into two types: the send
                  key and accept key.
                  ●      Send key: used by the device to encrypt packets before sending them.
                  ●      Accept key: used by the device to decrypt received packets.
                          NOTE

                         ● The send key of the local device must be the same as the accept key of the peer device.
                           In this way, the peer device can use the same algorithm and key string as those of the
                           local device to decrypt the received encrypted packets.
                         ● The send key and accept key on a device can be the same or different. In a certain
                           period of time, the key used for encrypting the packets to be sent must be within its
                           send lifetime, and that used for decrypting received packets must be within its accept
                           lifetime.


Lifetime of a Key
                  To better control the use of keys on a device, the keychain defines the key lifetime,
                  which determines the currently valid key, and its validity period and successor.
                  The lifetime of a key indicates the time interval within which the key is considered
                  valid.
                  ●      A key within the send lifetime is the currently valid send key.
                  ●      A key in the accept lifetime is the currently valid accept key.
                  Both the send lifetime and accept lifetime can be defined in either of the following
                  time modes:
                  ●      Absolute time mode: A key in a keychain is valid only within a specified time
                         interval, for example, 12:00-18:00 on December 10, 2019.
                  ●      Periodic time mode: As shown in Table 15-1, a key in a keychain becomes
                         valid periodically within a specified time interval.

                         Table 15-1 Periodic time mode

                          Period                                        Time Interval

                          Daily                                         Specified time interval every day, for
                                                                        example, 12:00-18:00.

                          Weekly                                        Specified days every week, for
                                                                        example, Monday, Wednesday,
                                                                        Friday, and Sunday.

                          Monthly                                       Specified dates every month, for
                                                                        example, from the third day to the
                                                                        eighth day.

                          Yearly                                        Specified months every year, for
                                                                        example, from June to September.



Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    285
Security Configuration
Security Configuration                                                                   15 Keychain Configuration


                          NOTE

                         ● Default send key: If no active send key is available within a specified time interval,
                           packets sent by the device cannot be encrypted and security authentication services
                           cannot be provided for applications. To prevent this, the device supports configuration of
                           the default send key, which is valid when no other active send keys are available within
                           a specified time interval.
                         ● Acceptance tolerance: When the send key of the peer device is updated, the accept key
                           of the local device must also be updated accordingly. Otherwise, the local device cannot
                           decrypt the received packets and will therefore discard them. When the keys of the two
                           ends are updated at the same time, due to a time difference in packet transmission, the
                           packets encrypted by the old key from the peer end may arrive at the local end
                           although the local end has started to use a key. In addition, the clocks of the two ends
                           on the network may be asynchronous. To address this, the device supports configuration
                           of the acceptance tolerance to ensure a smooth transition during key rollover. The
                           tolerance limit takes effect only on accept keys. After the tolerance limit is configured,
                           the actual lifetime of an accept key equals its original lifetime plus two tolerance limits,
                           one for the start time and the other for the end time of the accept key.


Key Set
                  A keychain is a set of keys of the same type.
                  In most cases, the same type refers to the same lifetime mode. For example, a key
                  that has a lifetime of specified months in every year is of a different type from a
                  key that has a lifetime of specified dates in every month. If these two keys are
                  placed in one keychain, time-based key rollover cannot be performed.
                  Based on service requirements, the device supports multiple keychains for
                  applications to flexibly choose from.
                  For example, as shown in Table 15-2, Key1 and Key3 are of the same type and
                  can be placed in KeychainA; Key2 and Key4 are of the same type and can be
                  placed in KeychainB; Key5 can only be placed independently in KeychainC; and
                  Key6 can only be placed independently in KeychainD.

                  Table 15-2 Example of key sets
                   Key                   Authenticati        Authenticati        Lifetime           Key Set
                                         on Algorithm        on Key String

                   Key1                  HMAC-               AbCdEfGh            12:00-15:00        KeychainA
                                         SHA1-20                                 on December
                                                                                 10, 2019

                   Key2                  HMAC-               HgFeDcBa            Every              KeychainB
                                         SHA1-20                                 Monday,
                                                                                 Wednesday,
                                                                                 Friday, and
                                                                                 Sunday

                   Key3                  HMAC-               AcEgHfDb            15:00-18:00        KeychainA
                                         SHA-256                                 on December
                                                                                 10, 2019




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       286

