---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-73
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [10299, 10434]
sha256: e4c50f180cce7915d4f1b4481282af33109a10059cce72a4a2d8fce1680b78ff
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

         Step 4 Disable LDP interface flapping suppression.
                 suppress-flapping interface disable

                 ----End


3.23 Configuring LDP Security Features

3.23.1 Understanding LDP Security Features
LDP MD5 Authentication
                 Message-Digest Algorithm 5 (MD5) is a digest algorithm defined in relevant
                 standards. MD5 is typically used to calculate a message digest to prevent message
                 spoofing. The MD5 message digest is a unique result calculated using an
                 irreversible character string conversion algorithm. If the message has been
                 modified during transmission, the receive end can determine this the moment the
                 message arrives by checking if the newly calculated digest is different from the
                 carried digest.
                 LDP MD5 authentication identifies LDP message modification by generating
                 unique message digest for the same information segment. It is stricter than
                 common TCP checksum-based authentication.
                 LDP MD5 authentication is performed before LDP messages are sent over TCP. A
                 unique message digest is added following the TCP header in a message. The
                 message digest is calculated using the MD5 algorithm based on the TCP header,
                 LDP message, and user-defined password.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                        173
MPLS Configuration
MPLS Configuration                                                                  3 MPLS LDP Configuration


                 When receiving the message, the receive end obtains the TCP header, message
                 digest, and LDP message. MD5 generates the message digest based on the TCP
                 header, LDP message and the locally saved password. Then, it compares the
                 calculated message digest with the message digest carried in the LDP message. If
                 they differ, the receive end interprets the LDP message as having been tampered
                 with.

                 A password can be set either in non-ciphertext or ciphertext. These two modes
                 differ in how passwords are recorded in the configuration file. In non-ciphertext
                 mode, a user-defined password is directly recorded in the configuration file;
                 however, in ciphertext mode, a password is encrypted using a special algorithm
                 and then recorded in the configuration file.

                 Characters set by users are used in digest calculation, regardless of whether a
                 password is recorded in non-ciphertext or ciphertext mode. Although the
                 conversion algorithms of the non-ciphertext and ciphertext are proprietary to
                 different vendors, those proprietary algorithms become transparent to other
                 vendors by this way.

                         NOTE

                        As MD5 is insecure, you are advised to use a more secure authentication mode.


LDP Keychain Authentication
                 Keychain is an enhanced encryption algorithm. Similar to MD5, it calculates a
                 digest for a piece of information to prevent LDP packets from being tampered
                 with.

                 Keychain allows users to define a group of passwords to form a password string.
                 Each password is assigned encryption and decryption algorithms, such as MD5
                 and secure hash algorithm-1 (SHA-1), and a validity period. The system selects a
                 valid password based on the user configuration before sending or receiving a
                 packet. Within the validity period of the password, the system uses the encryption
                 algorithm matching the password to encrypt the packet before sending it. The
                 system also uses the decryption algorithm matching the password to decrypt the
                 packet after accepting it. In addition, the system automatically switches to a new
                 valid password based on the password validity period, which minimizes password
                 decryption risks if the password is not changed for a long time.

                 The password of keychain authentication, the encryption and decryption
                 algorithms, and the expiration period of the password can be configured
                 separately on a keychain configuration node. A keychain configuration node at
                 least requires one password and has the encryption and decryption algorithms
                 specified.

                 To reference a keychain configuration node, specify a peer IP address and a node
                 name in the MPLS-LDP view. The keychain configuration node is then used to
                 encrypt an LDP session. Different peers can reference the same keychain
                 configuration node.

LDP GTSM
                 LDP Generalized TTL Security Mechanism (GTSM) is the application of GTSM in
                 LDP. GTSM determines whether a packet is valid by checking its TTL. This protects

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              174
MPLS Configuration
MPLS Configuration                                                        3 MPLS LDP Configuration


                 devices from attacks. GTSM for LDP involves applying GTSM to LDP messages
                 between adjacent devices or devices close to each other (based on the number of
                 next hops). A TTL value range is preset. The LDP messages with TTLs not within
                 the specified value range are interpreted as attack messages and discarded.

                 GTSM is used to protect the TCP/IP-based control plane against CPU utilization
                 attacks, for example, CPU overload attacks. GTSM for LDP is used to verify all
                 types of LDP packets to prevent LDP from suffering CPU utilization attacks when
                 LDP receives and processes a large number of forged packets.


                 Figure 3-35 Networking topology of LDP GTSM




                 As shown in Figure 3-35, LSR1 to LSR5 are core devices on the backbone network.
                 When LSR-A is indirectly connected to the core devices through other devices, LSR-
                 A may forge LDP packets among the LSRs (LSR1 to LSR5) to launch attacks.

                 After LSR-A accesses the network through another device, the TTL value carried in
                 the forged packet cannot be forged. This is the prerequisite for GTSM.

                 A GTSM policy is configured on LSR1 through LSR5 separately and is used to verify
                 packets reaching possible neighbors. For example, on LSR5, the valid number of
                 hops is set to 1 or 2, and the valid TTL is set to 254 or 255 for packets sent from
                 LSR2. The forged packet sent by LSR-A to LSR5 through multiple intermediate
                 devices contains a TTL value that is out of the preset TTL range. LSR5 discards the
                 forged packet and prevents the attack.

3.23.2 Configuring LDP MD5 Authentication

Context
                 MD5 authentication can be configured for a TCP connection over which an LDP
                 session is established to improve security. Two peers of an LDP session can be
                 configured with different authentication modes but must be configured with the
                 same password.

Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                          175
MPLS Configuration
MPLS Configuration                                                                      3 MPLS LDP Configuration


