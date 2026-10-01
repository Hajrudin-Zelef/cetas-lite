---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-75
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [10560, 10695]
sha256: 21003e29540637973de77d11d63f0f0a758cc3f6200e3dedb1c1be408186c31c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                                 ● The password is at least eight characters long and contains at least two of the
                                   following types: upper-case letters, lower-case letters, digits, and special
                                   characters (except the question mark (?) and space).
                                 ● For security purposes, you are advised to configure a ciphertext password and
                                   change the password periodically.
                                 ● For security purposes, MD5 is not recommended. If MD5 is required, run the
                                   install feature-software WEAKEA command first to install the weak security
                                   algorithm/protocol feature package WEAKEA.
                        d.   (Optional) Exclude a specified peer from authentication.
                             authentication exclude peer peer-id

                             By default, after LDP keychain authentication is enabled for all LDP peers,
                             keychain authentication takes effect on all LDP peers. Perform this step
                             to disable the device from authenticating a specified LDP peer.

                 ----End

3.23.3 Configuring LDP Keychain Authentication

Prerequisites
                 Before configuring LDP keychain authentication, complete the following task:

                 ●      Keychain Configuration


Context
                 To help improve LDP session security, you can configure keychain authentication
                 for a TCP connection over which an LDP session has been established or LDP
                 keychain authentication for Targeted Hello messages.

                 Keychain allows users to define a group of passwords to form a password string.
                 Each password is assigned encryption and decryption algorithms, such as MD5
                 and SHA-1, and a validity period. The system selects a valid password based on
                 the user configuration before sending or receiving a packet. Within the validity
                 period of the password, the system uses the encryption algorithm matching the
                 password to encrypt the packet before sending it. The system also uses the
                 decryption algorithm matching the password to decrypt the packet after accepting
                 it. In addition, the system automatically switches to a new valid password based
                 on the password validity period, which minimizes password decryption risks if the
                 password is not changed for a long time.

                 You can configure either LDP MD5 authentication or LDP keychain authentication
                 to match your scenario:
                 ●      The MD5 algorithm is easy to configure and generates a single password,
                        which can only be changed manually. MD5 authentication applies to
                        networks requiring short-period encryption.
                 ●      Keychain authentication involves a set of passwords, which can be
                        automatically switched based on the configuration. However, keychain
                        authentication is complex to configure and applies to networks requiring high
                        security.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   178
MPLS Configuration
MPLS Configuration                                                                      3 MPLS LDP Configuration


                         NOTE

                        LDP authentication configurations are prioritized in descending order: for a single peer, for a
                        specified peer group, and for all peers. Both keychain and MD5 authentication can be
                        configured. However, configurations with a higher priority override those with a lower
                        priority, and those with the same priority are mutually exclusive. For example, if MD5
                        authentication is configured for Peer1 and keychain authentication is configured for all LDP
                        peers, MD5 authentication takes effect on Peer1 and keychain authentication takes effect
                        on other peers.
                        As MD5 is insecure, you are advised to use a more secure authentication mode, such as
                        keychain authentication.

                 Perform the following configuration on the LSRs at both ends of an LDP session.

Procedure
                 ●      Configure LDP keychain authentication for a specified LDP peer based on the
                        TCP connection.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   Enable LDP keychain authentication and specify a keychain name.
                             authentication key-chain peer peer-id name keychain-name

                             By default, LDP keychain authentication is not performed between LDP
                             peers.


                                  NOTICE

                             ● Configuring LDP keychain authentication causes an LDP session to be
                               reestablished and deletes the LSP associated with the original LDP
                               session.
                             ● When the AES-128-CMAC algorithm is configured, a key string must
                               contain 16 characters.

                 ●      Configure LDP keychain authentication for LDP peers in a specified LDP peer
                        group based on the TCP connection.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   Enable LDP keychain authentication for LDP peers in a specified LDP peer
                             group and specify a keychain name.
                             authentication key-chain peer-group ip-prefix-name name keychain-name

                             An IP prefix list can be specified using ip-prefix-name to define the range
                             of peer IP addresses in a peer group. Ensure that the IP prefix list
                             specified by ip-prefix-name has been configured before performing this
                             step.
                        d.   (Optional) Exclude a specified peer from authentication.
                             authentication exclude peer peer-id


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       179
MPLS Configuration
MPLS Configuration                                                                    3 MPLS LDP Configuration


                             By default, after LDP keychain authentication is enabled for all LDP peers,
                             keychain authentication takes effect on all LDP peers. Perform this step
                             to disable the device from authenticating a specified LDP peer.
                 ●      Configure LDP keychain authentication for all LDP peers based on the TCP
                        connection.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   Enable LDP keychain authentication for all LDP peers and specify a
                             keychain name.
                             authentication key-chain all name keychain-name

