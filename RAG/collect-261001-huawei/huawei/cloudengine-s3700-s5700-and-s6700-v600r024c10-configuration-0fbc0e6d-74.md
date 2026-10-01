---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-74
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [10435, 10559]
sha256: e6e76d396ee4a30a851e05175821e2e0caead7636602a4f680b46be5c0eb447f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 LDP MD5 authentication identifies LDP message modification by generating
                 unique message digest for the same information segment. It is stricter than
                 common TCP checksum-based authentication.
                 You can configure either LDP MD5 authentication or LDP keychain authentication
                 to match your scenario:
                 ●      The MD5 algorithm is easy to configure and generates a single password,
                        which can only be changed manually. MD5 authentication applies to
                        networks requiring short-period encryption.
                 ●      Keychain authentication involves a set of passwords, which can be
                        automatically switched based on the configuration. However, keychain
                        authentication is complex to configure and applies to networks requiring high
                        security.
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
                 ●      Configure LDP MD5 authentication for a single LDP peer.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   Enable MD5 authentication and set an authentication password.
                             md5-password { plain | cipher } peer-lsr-id password

                                    NOTE

                                  ● The password must be at least eight characters long and contain at least two
                                    of the following character types: uppercase letters, lowercase letters, digits,
                                    and special characters.
                                  ● For security purposes, you are advised to configure a ciphertext password and
                                    change the password periodically.
                                  ● For security purposes, MD5 is not recommended. If MD5 is required, run the
                                    install feature-software WEAKEA command first to install the weak security
                                    algorithm/protocol feature package WEAKEA.

                             A password can be set either in non-ciphertext or ciphertext. These two
                             modes differ in how passwords are recorded in the configuration file. In
                             non-ciphertext mode, a user-defined password is directly recorded in the
                             configuration file; however, in ciphertext mode, a password is encrypted
                             using a special algorithm and then recorded in the configuration file.
                             By default, LDP MD5 authentication is not performed between LDP peers.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       176
MPLS Configuration
MPLS Configuration                                                                      3 MPLS LDP Configuration


                                 NOTICE

                             ● If you configure a password in non-ciphertext, the password will be
                               saved in the same way in the configuration file. The non-ciphertext
                               mode has high security risks, and therefore the ciphertext mode is
                               recommended. To ensure device security, change the password
                               periodically.
                             ● Configuring LDP MD5 authentication causes an LDP session to be
                               reestablished and deletes the LSP associated with the original LDP
                               session.

                 ●      Configure LDP MD5 authentication in a batch for a specified peer group.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   Enable MD5 authentication for LDP peers in a specified LDP peer group
                             and set a password.
                             md5-password { plain | cipher } peer-group ip-prefix-name password

                             An IP prefix list can be specified using ip-prefix-name to define the range
                             of peer IP addresses in a peer group. Ensure that the IP prefix list
                             specified by ip-prefix-name has been configured before performing this
                             step.

                                   NOTE

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
                 ●      Configure LDP MD5 authentication for all LDP peers.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   Enable MD5 authentication for all LDP peers and set a password.
                             md5-password { plain | cipher } all password




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   177
MPLS Configuration
MPLS Configuration                                                                  3 MPLS LDP Configuration


                                   NOTE

