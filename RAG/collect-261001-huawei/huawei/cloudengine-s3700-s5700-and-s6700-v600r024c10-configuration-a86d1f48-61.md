---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-61
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [7342, 7444]
sha256: 4b36a4cf603f3d3e30e21c63bd92aa0af00a67dce0dd00658634d435705417fb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   Configure an     macsec cipher-suite       After session negotiation is complete,
                   encryption       { gcm-aes-128 | gcm-      the devices at both ends use the SAK
                   algorithm for    aes-256 | gcm-aes-        to encrypt and decrypt data packets
                   MACsec data      xpn-128 | gcm-aes-        for encrypted communication.
                   packets.         xpn-256 | gcm-            By default, the encryption algorithms
                                    sm4-128 | gcm-sm4-        supported by different devices vary.
                                    xpn-128 } *               The actually used algorithm is
                                                              automatically negotiated by the
                                                              devices at both ends in descending
                                                              order of encryption strength.
                                                              The gcm-sm4-128 and gcm-sm4-
                                                              xpn-128 algorithms are supported
                                                              only on the S6750-H and S6780-H
                                                              and S5755-S.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           135
Security Configuration
Security Configuration                                                         8 MACsec Configuration


                   Operation        Command                   Description

                   Configure the    macsec                    The MACsec encryption offset
                   MACsec           confidentiality-offset    indicates from which byte behind the
                   encryption       offset-value              MACsec tag field a data frame is
                   offset.                                    encrypted. The protocol provides
                                                              three encryption offsets: 0 bytes, 30
                                                              bytes, and 50 bytes. For some
                                                              applications (such as load balancing)
                                                              that need to identify IPv4/IPv6
                                                              headers, packet headers must not be
                                                              encrypted. In this case, you need to
                                                              configure the encryption offset.
                                                              By default, the MACsec encryption
                                                              offset is 0 bytes.
                                                              If the local end is not the key server,
                                                              the encryption offset advertised by
                                                              the key server is used. If the local end
                                                              is the key server, the locally
                                                              configured encryption offset is used,
                                                              and the local end advertises the
                                                              offset to the remote end.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            136
Security Configuration
Security Configuration                                                        8 MACsec Configuration


                   Operation        Command                   Description

                   Configure the    macsec replay-            To prevent malicious users from
                   MACsec replay    window window-size        repeatedly sending captured data
                   protection                                 packets, the receiver discards
                   window size.                               duplicate or out-of-order data
                                                              packets by default. In some cases,
                                                              however, data packets are reordered
                                                              during transmission because their
                                                              sending priorities are different. As a
                                                              result, the data packets are out of
                                                              order when they reach the receiver. To
                                                              ensure that these out-of-order data
                                                              packets can be received normally,
                                                              configure the replay protection
                                                              window size.
                                                              Assume that the replay protection
                                                              window size configured on the device
                                                              is a. If the device receives a packet
                                                              with the sequence number x, the
                                                              sequence number of the next packet
                                                              that is allowed to be received must be
                                                              greater than or equal to (x + 1 - a).
                                                              Set an appropriate replay protection
                                                              window size based on the data
                                                              packet forwarding path on the
                                                              network. If data packets may be
                                                              forwarded multiple times, there is a
                                                              high probability that many of them
                                                              become out of order. To address this
                                                              issue, you are advised to increase the
                                                              replay protection window size.
                                                              Otherwise, decrease the window size.
                                                              By default, the MACsec replay
                                                              protection window size is 0.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                         137
Security Configuration
Security Configuration                                                        8 MACsec Configuration


                   Operation        Command                   Description

