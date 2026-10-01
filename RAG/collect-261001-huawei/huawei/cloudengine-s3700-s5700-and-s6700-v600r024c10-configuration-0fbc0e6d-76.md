---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-76
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [10696, 10840]
sha256: 463e3da5b57cff19a75e1c06d98168e0a8e11cbe1d561097ccb4c3eb7747a031
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        d.   (Optional) Exclude a specified peer from authentication.
                             authentication exclude peer peer-id

                             By default, after LDP keychain authentication is enabled for all LDP peers,
                             keychain authentication takes effect on all LDP peers. Perform this step
                             to disable the device from authenticating a specified LDP peer.
                 ●      Configure LDP keychain authentication for a UDP connection.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   Enable keychain authentication for Targeted Hello packets of a specified
                             LDP peer and specify a keychain name.
                             authentication udp-remote key-chain peer peer-id name keychain-name

                             After the configuration is successful, the configured keychain
                             authentication takes effect on the specified peer. If the authentication
                             fails, the LDP session cannot be established.
                                   NOTE

                                 This command supports only the keychain authentication using a strong
                                 encryption algorithm (SHA-256, HMAC-SHA-256, HMAC-SHA-384, HMAC-
                                 SHA-512, AES-128-CMAC, or SM3) but not a weak encryption algorithm.


                 ----End

3.23.4 Configuring LDP GTSM

Context
                 The GTSM checks TTL values to verify packets and defends devices against attacks.
                 LDP peers that are configured with the GTSM and a valid TTL range check TTLs in
                 LDP messages exchanged between them. If the TTL in an LDP message is not in
                 the valid range, this LDP message is considered invalid and discarded. GTSM
                 defends against CPU utilization attacks initiated using a large number of forged
                 packets and protects upper-layer protocols.

                 Perform the following configuration at both ends of an LDP peer relationship.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                180
MPLS Configuration
MPLS Configuration                                                                    3 MPLS LDP Configuration


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 Configure LDP GTSM.
                 gtsm peer ip-address valid-ttl-hops hops

                 If the value of hops is set to the maximum number of valid hops permitted by
                 GTSM, when the TTL values carried in the packets sent by an LDP peer are within
                 the range [255 – hops + 1, 255], the packets are accepted; otherwise, the packets
                 are discarded.
                 By default, GTSM is not configured for any LDP peer.

                 ----End

3.23.5 Verifying the Configuration
Procedure
                 ●      Run the display mpls ldp session verbose command to check the
                        configurations of LDP MD5 authentication and LDP keychain authentication.
                 ●      Run the display gtsm statistics { slot-id | all } command to check GTSM
                        statistics on a specified slot or all slots.
                 ----End


3.24 Checking MPLS Network Connectivity Using Ping/
Tracert

3.24.1 Checking MPLS Network Connectivity Using Ping
Prerequisites
                         NOTE

                        This feature is supported only by the following series: S6780-H, S6750-H, S6750E-S, S6750-
                        S, S5755E-H, S5755-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2
                 ●      The MPLS network has been correctly configured before MPLS ping is used to
                        check the connectivity of an LSP.
                 ●      The undo lspv mpls-lsp-ping echo disable command has been run to enable
                        a device to respond to MPLS Echo Request messages.
                 ●      Both the initiator and responder send LSP ping packets to the main control
                        unit for processing. If a large number of packets are sent to the main control
                        unit, the CPU usage of the main control unit surges, affecting normal device
                        running. To prevent this problem, you can run the lspv mpls-lsp-ping cpu-
                        defend cpu-defend command to limit the rate at which MPLS Echo Request
                        messages are sent to the main control unit.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     181
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration


Context
                 On an MPLS network, the MPLS control plane for establishing label switched
                 paths (LSPs) cannot detect data forwarding failures in LSPs, complicating network
                 maintenance. To address this, MPLS ping and tracert can be used to detect LSP
                 faults and quickly locate faulty nodes. MPLS ping is mainly used to check network
                 connectivity and host reachability.

                 MPLS ping checks the LSP status by sending MPLS Echo Request and Reply
                 messages. The two types of messages are encapsulated into UDP packets and
                 transmitted using port 3503. The receiver identifies the MPLS Echo Request and
                 Reply messages based on the UDP port number.

                 An MPLS Echo Request message carries information about the forwarding
                 equivalence class (FEC) for an LSP to be checked. The MPLS Echo Request
                 message is forwarded along the LSP with other packets belonging to this FEC. This
                 procedure enables the LSP connectivity to be checked. MPLS Echo Request
                 messages are forwarded to the destination using MPLS, whereas MPLS Echo Reply
                 messages are forwarded to the source using IP.

                 To prevent the egress from forwarding a received MPLS Echo Request message to
                 other nodes, you can set the destination IP address to 127.0.0.0/8 (the local
                 loopback address) and the TTL value to 1 in the IP header of the message.

                 In Figure 3-36, an LSP destined for Device D is configured on Device A. The MPLS
                 ping process on Device A is as follows:

                 1.     Device A checks whether the LSP is established. If the LSP does not exist, an
                        error message is displayed, and the ping process ends. If the LSP is
                        established, the ping process continues.
                 2.     Device A constructs an MPLS Echo Request message, with destination IP
                        address 127.0.0.0/8 and TTL value 1 in the IP header. Then, Device A adds an
                        LSP label to the MPLS Echo Request message and sends the message to
                        Device B.
                 3.     Transit nodes Device B and Device C forward this MPLS Echo Request message
                        using MPLS.
                        If either transit node fails to forward the MPLS Echo Request message, this
                        message is discarded.
                 4.     If the forwarding is successful, the MPLS Echo Request message reaches the
                        egress (Device D), which returns an MPLS Echo Reply message.

                 Figure 3-36 MPLS network




