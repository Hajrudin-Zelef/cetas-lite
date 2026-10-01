---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-130
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2019-12-10", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [16712, 16833]
sha256: 7bc39d6f371884447e828ece7a41569ffcdf7ff87aeeac56a20c6e7e7cb54393
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Verifying the Configuration
                  Using DeviceA as an example, check whether keychain authentication is
                  successfully configured for BGP.
                  ●      Run the display keychain keychain-name command to check the key ID in
                         the Active state.
                         <DeviceA> display keychain huawei
                          Keychain Information:
                          ----------------------
                          Keychain Name                 : huawei
                            Timer Mode                : Absolute
                            Receive Tolerance(min) : 10
                            Digest Length            : 32
                            Time Zone               : LMT
                            TCP Kind              : 182
                            TCP Algorithm IDs           :
                             HMAC-MD5                   :5
                             HMAC-SHA1-12                 :2
                             HMAC-SHA1-20                 :6
                             MD5                 :3
                             SHA1                :4


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 305
Security Configuration
Security Configuration                                                                        15 Keychain Configuration

                           HMAC-SHA-256             : 17
                           SHA-256             :8
                           SM3               :9
                           HMAC-SHA-384             : 11
                           HMAC-SHA-512             : 12
                         Number of Key ID         :2
                         Active Send Key ID      :1
                         Active Receive Key ID    : 01
                         Default send Key ID      :1

                         Key ID Information:
                         ----------------------
                         Key ID                  :1
                           Key string             : ******
                           Algorithm                : HMAC-SHA-256
                           SEND TIMER                  :
                            Start time            : 2019-12-10 12:00
                            End time               : 2019-12-10 15:00
                            Status              : Active
                           RECEIVE TIMER                 :
                            Start time            : 2019-12-10 12:00
                            End time               : 2019-12-10 15:00
                            Status              : Active

                         Key ID                :2
                          Key string            : ******
                          Algorithm               : HMAC-SHA-256
                          SEND TIMER                 :
                           Start time           : 2019-12-10 15:05
                           End time              : 2019-12-10 18:00
                           Status             : Inactive
                          RECEIVE TIMER                :
                           Start time           : 2019-12-10 15:05
                           End time              : 2019-12-10 18:00
                           Status             : Inactive
                  ●      Run the display bgp peer ipv4-address verbose command to verify that the
                         authentication type configured for the BGP peer is Keychain(huawei).
                         <DeviceA> display bgp peer 192.168.1.2 verbose
                                BGP Peer is 192.168.1.2, remote AS 1
                                Type: IBGP link
                                BGP version 4, Remote router ID 2.2.2.2
                                Update-group ID: 3
                                BGP current state: Established, Up for 00h27m26s
                                BGP current event: RecvKeepalive
                                BGP last state: OpenConfirm
                                BGP Peer Up count: 2
                                Received total routes: 0
                                Received active routes total: 0
                                Advertised total routes: 0
                                Port: Local - 58168      Remote - 179
                                Configured: Connect-retry Time: 32 sec
                                Configured: Min Hold Time: 0 sec
                                Configured: Active Hold Time: 180 sec Keepalive Time:60 sec
                                Received : Active Hold Time: 180 sec
                                Negotiated: Active Hold Time: 180 sec Keepalive Time:60 sec
                                Peer optional capabilities:
                                Peer supports bgp multi-protocol extension
                                Peer supports bgp route refresh capability
                                Peer supports bgp 4-byte-as capability
                                Address family IPv4 Unicast: advertised and received
                          Received: Total 34 messages
                                       Update messages             1
                                       Open messages              1
                                       KeepAlive messages           32
                                       Notification messages        0
                                       Refresh messages           0
                          Sent: Total 33 messages
                                       Update messages             1
                                       Open messages              1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       306
Security Configuration
Security Configuration                                                              15 Keychain Configuration

                                     KeepAlive messages          31
                                     Notification messages       0
                                     Refresh messages           0
                         Authentication type configured: Keychain(huawei)
                         Last keepalive received: 2019-12-10 10:12:29+00:00
                         Last keepalive sent : 2019-12-10 10:12:04+00:00
                         Last update received: 2019-12-10 09:45:14+00:00
                         Last update sent : 2019-12-10 09:45:14+00:00
                         No refresh received since peer has been configured
                         No refresh sent since peer has been configured
                         Minimum route advertisement interval is 15 seconds
                         Optional capabilities:
                         Route refresh capability has been enabled
                         4-byte-as capability has been enabled
                         Peer Preferred Value: 0
                         Routing policy configured:
                         No routing policy is configured


