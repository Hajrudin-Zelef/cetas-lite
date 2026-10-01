---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-121
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [17158, 17272]
sha256: c31909157a5eb3a3986075dcf1b23565baea8e6fa0c7dafa74c381cfc5db39d3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              288
MPLS Configuration
MPLS Configuration                                                                          4 MPLS TE Configuration


         Step 2 Enter the interface view or MPLS RSVP-TE neighbor view as required.
                 ●      Enter the MPLS TE link interface view.
                        interface interface-type interface-number

                        RSVP-TE authentication configured in the interface view takes effect only on
                        the current interface and has the lowest priority.
                              NOTE

                        For an Ethernet interface, you also need to run the undo portswitch command to switch
                        the interface to Layer 3 mode.
                        Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S, S6750E-S,
                        S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from Layer 2 mode to
                        Layer 3 mode using the undo portswitch command. Determine whether to run this
                        command based on the current interface mode.
                 ●      Enter the MPLS RSVP-TE neighbor view.
                        mpls
                        mpls rsvp-te peer peer-addr

                        –    If peer-addr is set to the interface address of the neighbor and is different
                             from the LSR ID of the neighbor, the authentication function is configured
                             based on the interface address of the neighbor. This configuration mode
                             enables the authentication function to take effect only on this interface
                             of the neighbor, delivering high security. Therefore, it has the highest
                             priority.
                        –    If peer-addr is set to the LSR ID of the neighbor, the authentication
                             function is configured based on this LSR ID. This configuration mode
                             enables the authentication function to take effect on the entire device,
                             without the need to specify interfaces.
                              NOTE

                        If the LSR ID of the peer device is used as the neighbor address, CSPF must be enabled on
                        the device that needs to be configured with RSVP-TE authentication.
                        If interface authentication, neighboring node authentication, and neighbor interface
                        address authentication are all configured for the same neighbor, neighbor interface address
                        authentication preferentially takes effect, followed by neighboring node authentication and
                        interface authentication, in descending order.

         Step 3 Configure an RSVP-TE authentication key.
                 mpls rsvp-te authentication { { cipher | plain } auth-key | keychain keychain-name }

                 Configure HMAC-MD5 or keychain authentication based on the selected
                 parameter.
                 ●      cipher: Configure HMAC-MD5 authentication and enable the device to display
                        the authentication key in ciphertext.
                 ●      plain: Configure HMAC-MD5 authentication and enable the device to display
                        the authentication key in plaintext.
                 ●      keychain: Configure keychain authentication and enable the device to
                        reference a globally configured keychain.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    289
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                        NOTICE

                 If you select the plain mode in authentication key configuration, the password is
                 saved in the configuration file in plaintext. Because this mode has high security
                 risks, you are advised to select the cipher mode. For security purposes, change the
                 password periodically.
                 It is recommended that the password contain uppercase letters, lowercase letters,
                 digits, and special characters.
                 HMAC-MD5 authentication provides low security. To ensure higher security, you
                 are advised to use keychain authentication and a high-security algorithm, such as
                 HMAC-SHA-256.
                 The configuration must be completed on the two directly connected interfaces or
                 neighboring nodes within three refresh periods. Otherwise, the involved session
                 goes down.

         Step 4 (Optional) Configure RSVP-TE authentication enhancement as required. For
                details, see Table 4-13.
                 RSVP-TE authentication enhancement adds the authentication lifetime, handshake,
                 and message window features based on the original authentication functions. To
                 configure RSVP-TE authentication enhancement, perform one or more steps listed
                 in Table 1.

                 Table 4-13 Configuring RSVP-TE authentication enhancement
                  Operation                      Command                     Purpose

                  Configure an RSVP-TE           mpls rsvp-te                By default, the RSVP-TE
                  authentication lifetime.       authentication lifetime     authentication lifetime is
                                                 lifetime                    30 minutes.
                                                                             If no CR-LSP exists
                                                                             between RSVP-TE
                                                                             neighbors, the RSVP-TE
                                                                             neighbor relationship
                                                                             can be retained until the
                                                                             RSVP-TE authentication
                                                                             lifetime expires. The
                                                                             RSVP-TE authentication
                                                                             lifetime does not affect
                                                                             the status of existing CR-
                                                                             LSPs.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           290
MPLS Configuration
MPLS Configuration                                                          4 MPLS TE Configuration


                  Operation                   Command                     Purpose

