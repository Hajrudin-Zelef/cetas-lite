---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-128
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2019-12-10", "2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [16398, 16574]
sha256: fb7d9ef8a2f1769560f03cebff6cd0d513bc7a108b493322697f735ea2eb9474
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         NLPID       IPV4
                         AREA ADDR 10
                         INTF ADDR 192.168.1.2
                         INTF ADDR 192.168.2.2
                         NBR ID      0000.0000.0002.01 COST: 10
                         NBR ID      0000.0000.0003.01 COST: 10
                         IP-Internal 192.168.1.0  255.255.255.0 COST: 10
                         IP-Internal 192.168.2.0  255.255.255.0 COST: 10

                         0000.0000.0002.01-00 0x0000007e 0xa767         305         55     0/0/0
                          SOURCE      0000.0000.0002.01
                          NLPID     IPV4
                          NBR ID     0000.0000.0002.00 COST: 0
                          NBR ID     0000.0000.0001.00 COST: 0

                         0000.0000.0003.00-00 0x0000020c 0xd59e     322             68     0/0/0
                          SOURCE        0000.0000.0003.00
                          NLPID       IPV4
                          AREA ADDR 10
                          INTF ADDR 192.168.2.1
                          NBR ID      0000.0000.0003.01 COST: 10
                          IP-Internal 192.168.2.0   255.255.255.0 COST: 10

                         0000.0000.0003.01-00 0x0000007e 0xcc3f        322         55      0/0/0
                          SOURCE      0000.0000.0003.01
                          NLPID     IPV4
                          NBR ID     0000.0000.0003.00 COST: 0
                          SOURCE      0000.0000.0002.01
                          NLPID     IPV4
                          NBR ID     0000.0000.0002.00 COST: 0
                          NBR ID     0000.0000.0001.00 COST: 0

                         0000.0000.0003.00-00 0x0000020c 0xd59e     322             68     0/0/0
                          SOURCE        0000.0000.0003.00
                          NLPID       IPV4
                          AREA ADDR 10
                          INTF ADDR 192.168.2.1
                          NBR ID      0000.0000.0003.01 COST: 10
                          IP-Internal 192.168.2.0   255.255.255.0 COST: 10

                         0000.0000.0003.01-00 0x0000007e 0xcc3f        322         55      0/0/0
                          SOURCE      0000.0000.0003.01
                          NLPID     IPV4
                          NBR ID     0000.0000.0003.00 COST: 0
                          NBR ID     0000.0000.0002.00 COST: 0

                         Total LSP(s): 5
                            *(In TLV)-Leaking Route, *(By LSPID)-Self LSP, +-Self LSP(Extended),
                                 ATT-Attached, P-Partition, OL-Overload


Configuration Scripts
                  ●      DeviceA
                         #
                         sysname DeviceA
                         #
                         vlan batch 1
                         #
                         keychain huawei mode absolute
                          receive-tolerance 10
                          #
                          key-id 1
                           algorithm hmac-sha-256
                           key-string cipher %+%#)teP2/_7j#@>|r-p:jgDgyKC%=80dRNA,;Cjwwv~%+%#
                           send-time 12:00 2019-12-10 to 18:00 2019-12-10
                           receive-time 12:00 2019-12-10 to 18:00 2019-12-10
                           default send-key-id
                         #
                         isis 1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                        301
Security Configuration
Security Configuration                                                                 15 Keychain Configuration

                          is-level level-1
                          network-entity 10.0000.0000.0001.00
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 1
                         #
                         interface vlanif 1
                          ip address 192.168.1.1 255.255.255.0
                          isis enable 1
                          isis authentication-mode keychain huawei
                         #
                         return

                  ●      DeviceB
                         #
                         sysname DeviceB
                         #
                         vlan batch 1 2
                         #
                         keychain huawei mode absolute
                          receive-tolerance 10
                          #
                          key-id 1
                           algorithm hmac-sha-256
                           key-string cipher %+%#$V_<R'XnL6F&H`P2DLn#IE7-+'~ks9~\acM<OSf)%+%#
                           send-time 12:00 2019-12-10 to 18:00 2019-12-10
                           receive-time 12:00 2019-12-10 to 18:00 2019-12-10
                           default send-key-id
                         #
                         isis 1
                          is-level level-1
                          network-entity 10.0000.0000.0002.00
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 1
                         #
                         interface vlanif 1
                          ip address 192.168.1.2 255.255.255.0
                          isis enable 1
                          isis authentication-mode keychain huawei
                         #
                         interface 10GE1/0/2
                          port link-type trunk
                          port trunk allow-pass vlan 2
                         #
                         interface vlanif 2
                          ip address 192.168.2.2 255.255.255.0
                          isis enable 1
                          isis authentication-mode keychain huawei
                         #
                         return

                  ●      DeviceC
                         #
                         sysname DeviceC
                         #
                         vlan batch 2
                         #
                         keychain huawei mode absolute
                          receive-tolerance 10
                          #
                          key-id 1
                           algorithm hmac-sha-256
                           key-string cipher %+%#v@>@B\eP.Ruug(%b,;fS!5}]GV:rLU3(]U'zd9|>%+%#
                           send-time 12:00 2019-12-10 to 18:00 2019-12-10
                           receive-time 12:00 2019-12-10 to 18:00 2019-12-10
                           default send-key-id
                         #


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                302
Security Configuration
Security Configuration                                                        15 Keychain Configuration

                         isis 1
                          is-level level-1
                          network-entity 10.0000.0000.0003.00
                         #
                         interface 10GE1/0/2
                          port link-type trunk
                          port trunk allow-pass vlan 2
                         #
                         interface vlanif 2
                          ip address 192.168.2.1 255.255.255.0
                          isis enable 1
                          isis authentication-mode keychain huawei
                         #
                         return


15.5.6 Example for Configuring Keychain Authentication for
BGP
Networking Requirements
                  In Figure 15-7, DeviceA and DeviceB communicate with each other through BGP.

                  To ensure the stability and security of BGP connections, configure a keychain to
                  provide dynamic security authentication for BGP.

