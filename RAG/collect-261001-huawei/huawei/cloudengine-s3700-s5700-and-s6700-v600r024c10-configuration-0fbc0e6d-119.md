---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-119
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [16906, 17016]
sha256: 2c76fa09e9d994f97870d72f60bc302c345e00e3fae8a6689f5ea5c1993bc842
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      Authentication lifetime
                        The authentication lifetime determines how long an RSVP-TE neighbor
                        relationship can persist. When network flapping occurs, the neighbor
                        relationship may be repeatedly established and deleted. Each time the
                        neighbor relationship is established, the handshake process needs to be
                        authenticated, delaying CR-LSP establishment. To prevent performance
                        deterioration caused by authentication in flapping situations, RSVP-TE
                        introduces the authentication lifetime mechanism. This mechanism provides
                        the following functions:
                        –   If no CR-LSP exists between RSVP-TE neighbors, the RSVP-TE neighbor
                            relationship can be retained until the RSVP-TE authentication lifetime
                            expires. The RSVP-TE authentication lifetime does not affect the status of
                            existing CR-LSPs.
                        –   Ceaseless RSVP-TE authentication can be prevented. Assume that devices
                            A and B establish an RSVP-TE authentication relationship. If the key
                            carried in the packet that A sends to B is damaged due to packet
                            tampering, B discards the packet after receiving it and detecting the
                            incorrect key. This process repeats, but the authentication relationship
                            between the two devices cannot be torn down. To prevent this problem,
                            configure an RSVP-TE authentication lifetime. If a valid RSVP-TE packet is
                            received within the lifetime, the RSVP-TE authentication lifetime is reset.
                            However, if no valid packet is received when the lifetime expires, the
                            RSVP-TE authentication relationship is torn down.
                 ●      Handshake mechanism
                        The handshake mechanism enables a device to maintain the RSVP-TE
                        authentication status when it receives outdated packets. When the device
                        receives a packet from the peer device for the first time or receives a
                        disordered packet, it sends a handshake packet to the peer device to
                        resynchronize their sequence number windows.
                        After two RSVP-TE neighboring nodes authenticate each other, they exchange
                        handshake packets. If they accept the packets, they record a successful

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            284
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                        handshake. When receiving an outdated packet, the local end processes it as
                        follows:
                        –   Discards the packet if the packet indicates that the handshake
                            mechanism is not enabled on the peer end.
                        –   Discards the packet if the packet indicates that the handshake
                            mechanism is enabled on the peer end and the local end records a
                            successful handshake with the peer end. If the local end does not have
                            such a record, this packet is the first one received from the peer end. In
                            this case, the local end needs to start a handshake process with the peer
                            end.
                 ●      Message window
                        The message window saves the sequence numbers of RSVP-TE packets sent
                        from an RSVP-TE neighbor. To prevent replay attacks, a device allocates a 64-
                        bit monotonically increasing sequence number to each packet and then sends
                        the packet carrying the sequence number to the peer end. After receiving the
                        packet, the peer end checks not only the digest but also whether the
                        sequence number is within an allowable window. If the sequence number of
                        the packet is less than the lower limit defined in the window, the packet is
                        considered as a replay attack packet and therefore discarded.

RSVP-TE Key Management
                 RSVP-TE supports the following key management modes:
                 ●      HMAC-MD5
                        An HMAC-MD5 key is entered in either ciphertext or simple text for an RSVP-
                        TE interface or neighbor. This key management mode has the following
                        characteristics:
                        –   Protocols require their respective keys, meaning that keys cannot be
                            shared.
                        –   An interface or neighbor can be configured with only one key. A key can
                            be changed only through reconfiguration.


                            NOTICE

                        HMAC-MD5 authentication provides low security. To ensure higher security,
                        you are advised to use keychain authentication and a high-security algorithm,
                        such as HMAC-SHA-256.

                 ●      Keychain
                        Keychain is an enhanced encryption algorithm. It allows you to define a group
                        of passwords to form a password string. You can specify the encryption and
                        decryption algorithms and validity period for each password. The system
                        selects a valid password for each packet based on existing configurations.
                        Using the encryption and decryption algorithms matching the password, the
                        system encrypts a packet before sending it and decrypts a packet after
                        receiving it, respectively. In addition, the system automatically switches to a
                        new valid password based on the password validity period, which minimizes
                        password decryption risks if the password is not changed for a long time.
                        This key management mode has the following characteristics:

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           285
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                        –   The password of keychain authentication, the encryption and decryption
                            algorithms, and the password validity period can be configured separately
                            on a keychain configuration node. A keychain configuration node requires
                            at least one password as well as encryption and decryption algorithms to
                            be specified.
                        –   Keychain configurations can be referenced by various protocols,
                            implementing centralized management and sharing of keys.

RSVP-TE Authentication Mode
                 RSVP-TE keys can be configured for interfaces or neighbors. Figure 4-16 shows
                 three RSVP-TE key authentication modes that are currently supported.

                 Figure 4-16 RSVP-TE key authentication




