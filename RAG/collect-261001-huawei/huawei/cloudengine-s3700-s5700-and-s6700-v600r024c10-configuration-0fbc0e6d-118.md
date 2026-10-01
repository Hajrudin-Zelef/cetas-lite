---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-118
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["compute", "copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [16750, 16905]
sha256: 60b1297c606be472c54e08e684155b3c5ff41a1461a6963926e52c75aad2385f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        #
                        return
                 ●      LSR4
                        #
                        sysname LSR4
                        #
                        vlan batch 300
                        #
                        mpls lsr-id 4.4.4.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 00.0005.0000.0000.0004.00
                         traffic-eng level-2
                        #
                        interface Vlanif300
                         ip address 10.1.3.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      LSR5
                        #
                        sysname LSR5
                        #
                        vlan batch 400 500
                        #
                        mpls lsr-id 5.5.5.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 00.0005.0000.0000.0005.00
                         traffic-eng level-2
                        #
                        interface Vlanif400
                         ip address 10.1.4.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.1.5.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      282
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration

                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 5.5.5.9 255.255.255.255
                         isis enable 1
                        #
                        return



4.10 Configuring RSVP-TE Authentication

4.10.1 Understanding RSVP-TE Authentication
                 RSVP-TE authentication leverages key authentication to ensure packet security and
                 prevent spoofing attacks. In addition, it provides the authentication enhancement
                 mechanism by adding the authentication lifetime, handshake, and message
                 window features, preventing authentication performance deterioration (for
                 example, neighbor relationships may be repeatedly established and deleted during
                 network flapping), packet disorder, and replay attacks. This greatly improves the
                 security of RSVP-TE.


Related Concepts
                 ●      Raw IP: RSVP-TE uses raw IP to transmit protocol packets. Similar to UDP, raw
                        IP is unreliable, meaning that there is no control mechanism available to
                        determine whether raw IP packets have been received.
                 ●      Spoofing attack: The peer end establishes a neighbor relationship with the
                        local end without authorization, or forges RSVP-TE packets to establish an
                        RSVP-TE neighbor relationship with the local end for initiating attacks (for
                        example, malicious reservation of a large amount of bandwidth).
                 ●      Replay attack: The peer end repeatedly sends outdated packets with sequence
                        numbers smaller than the sequence number currently saved on the local end.
                        When the sequence number of an RSVP-TE packet is smaller than the
                        currently saved sequence number of the peer end, the RSVP-TE authentication
                        relationship is terminated, and the established CR-LSP is torn down.


Context
                 RSVP-TE uses raw IP to transmit protocol packets. Because raw IP has no security
                 mechanism and is prone to attacks, RSVP authentication can be configured to
                 prevent spoofing attacks through key authentication. The authentication
                 relationship is terminated upon the receipt of an outdated packet.

                 This type of key authentication cannot prevent replay attacks or solve the problem
                 of neighbor relationship termination caused by RSVP-TE packet disorder. By
                 contrast, RSVP-TE authentication enhancement can be used to solve these
                 problems and also enhance the capability of verifying neighbor relationship
                 validity in unfavorable network environments such as network congestion.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            283
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


RSVP-TE Authentication Fundamentals
                 ●      Key authentication
                        RSVP-TE authentication protects nodes against spoofing attacks by verifying
                        keys in packets exchanged between neighboring nodes. The same key must
                        be configured on both nodes involved in authentication. A node uses this key
                        and the HMAC-MD5 algorithm to compute a digest for a packet to be sent.
                        The digest is used as an integrity object and then sent to the peer node along
                        with the packet. After receiving the packet, the peer node uses the same key
                        and algorithm to compute a digest for the packet, and compares the
                        computed digest with the one carried in the packet. If they are the same, the
                        packet is accepted; if they are different, the packet is discarded.



                            NOTICE

                        HMAC-MD5 authentication provides low security. To ensure higher security,
                        you are advised to use keychain authentication and a high-security algorithm,
                        such as HMAC-SHA-256.

