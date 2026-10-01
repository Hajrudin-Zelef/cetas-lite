---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-62
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [7445, 7541]
sha256: 16f60b1d2ff8301917b8d7b0f0a0990f9d41f276fd08dd798dc0bcfe07af33cf
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   Configure the    macsec include-sci        An SCI identifies the source of a
                   device to add                              packet. It consists of the MAC address
                   an SCI to the                              of an interface and the last two bytes
                   MACsec frame                               of the interface index. When the
                   header.                                    device is connected to a non-Huawei
                                                              device and MACsec is configured on
                                                              them, the MACsec frame header may
                                                              need to contain the SCI to identify
                                                              the source of packets because the
                                                              implementation of MACsec on the
                                                              device is different from that on the
                                                              non-Huawei device. Currently, the
                                                              device supports only device-to-device
                                                              MACsec. The two MACsec-enabled
                                                              interfaces on both devices only
                                                              establish a session with each other. In
                                                              this case, the MACsec frame header
                                                              does not need to contain an SCI.
                                                              By default, the MACsec frame header
                                                              contains an SCI.
                                                              The settings of whether the MACsec
                                                              frame header contains an SCI must
                                                              be the same at both ends. That is, the
                                                              MACsec frame header contains an SCI
                                                              or does not contain an SCI at both
                                                              ends.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                          138
Security Configuration
Security Configuration                                                         8 MACsec Configuration


                   Operation        Command                   Description

                   Configure the    mka timer mka-life        After session negotiation is complete
                   MKA session      life-time                 and a secure channel is established,
                   timeout                                    the two devices exchange MKA
                   period.                                    protocol packets to ensure that the
                                                              session is alive. The MKA protocol
                                                              defines an MKA session keepalive
                                                              timer that specifies the timeout
                                                              period of an MKA session. The local
                                                              device starts the timer after receiving
                                                              MKA protocol packets from the
                                                              remote device.
                                                              ● If the local device receives
                                                                subsequent MKA protocol packets
                                                                within the timeout period, it
                                                                restarts the timer.
                                                              ● If the local device does not receive
                                                                subsequent MKA protocol packets
                                                                within the timeout period, it
                                                                considers the session insecure,
                                                                deletes the session, and performs
                                                                MKA session negotiation again.
                                                              By default, the MKA session timeout
                                                              period is 6 seconds.

                   Configure the    mka timer sak-life        When MACsec is used for secure
                   SAK timeout      life-time                 communication, an SAK is used to
                   period.                                    encrypt and decrypt data packets. To
                                                              improve the security of data packets,
                                                              an SAK needs to be replaced when
                                                              the number of data packets
                                                              encrypted using the SAK reaches a
                                                              certain value or the time for using the
                                                              SAK exceeds a certain period. In this
                                                              case, the key server generates and
                                                              advertises a new SAK. You can adjust
                                                              the SAK timeout period based on site
                                                              requirements.
                                                              By default, the SAK timeout period is
                                                              3600 seconds.
                                                              If the local end is not the key server,
                                                              the SAK timeout period advertised by
                                                              the key server is used. If the local end
                                                              is the key server, the locally
                                                              configured SAK timeout period is
                                                              used, and the local end advertises the
                                                              SAK timeout period to the remote
                                                              end.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           139
Security Configuration
Security Configuration                                                           8 MACsec Configuration


                   Operation          Command                   Description

