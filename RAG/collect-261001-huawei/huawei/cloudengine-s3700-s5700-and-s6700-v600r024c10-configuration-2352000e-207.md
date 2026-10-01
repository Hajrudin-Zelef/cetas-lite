---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-207
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [30389, 30511]
sha256: c6fec53edd25ae562f2b3bc2cc599a825f6bb5d791b5c9e425ae48525f2902e6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Verifying the Configuration
                    # Check SVC L2VPN connection information on PEs. The command output shows
                    that a static L2VC has been established.
                    The following example uses the command output on PE1.
                    <PE1> display mpls static-l2vc interface vlanif10
                     *Client Interface      : Vlanif10 is up
                      AC Status           : up
                      VC State           : up
                      VC ID            :0
                      VC Type            : VLAN
                      Destination          : 3.3.3.9
                      Transmit VC Label : 100
                      Receive VC Label          : 200
                      Label Status         :0
                      Token Status           :0
                      Control Word            : Disable
                      VCCV Capability          : alert ttl lsp-ping bfd
                      active state       : active
                      TTL Value            :1
                      Link State         : up
                      Tunnel Policy         : --
                      PW Template Name               : --
                      Main or Secondary : Main
                      load balance type : flow
                      Access-port          : false
                      VC tunnel info         : 1 tunnels
                        NO.0 TNL Type           : ldp          , TNL ID : 0x0000000001004c8b43
                      Create time          : 0 days, 0 hours, 3 minutes, 24 seconds
                      UP time            : 0 days, 0 hours, 1 minutes, 59 seconds
                      Last change time           : 0 days, 0 hours, 1 minutes, 59 seconds
                      VC last up time         : 2024/08/07 14:40:59
                      VC total up time         : 0 days, 0 hours, 1 minutes, 59 seconds
                      CKey             : 449
                      NKey              : 16777512
                      BFD for PW             : unavailable

                    # Check information about the interface used by the L2VPN connection. The
                    command output shows that the VC type is static-vc and the VC status is up.
                    The following example uses the command output on PE1.
                    <PE1> display l2vpn ccc-interface vc-type static-vc up
                    Total ccc-interface of SVC VC: 1
                    up (1), down (0)


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                         486
VPN Configuration
VPN Configuration                                                                                    5 VPWS Configuration

                    Interface             Encap Type                State      VC Type
                    Vlanif10              vlan                 up           static-vc

                    # CE1 and CE2 can ping each other successfully.
                    <CE1> ping 10.10.1.2
                     PING 10.10.1.2: 56 data bytes, press CTRL_C to break
                       Reply from 10.10.1.2: bytes=56 Sequence=1 ttl=255 time=46 ms
                       Reply from 10.10.1.2: bytes=56 Sequence=2 ttl=255 time=91 ms
                       Reply from 10.10.1.2: bytes=56 Sequence=3 ttl=255 time=74 ms
                       Reply from 10.10.1.2: bytes=56 Sequence=4 ttl=255 time=88 ms
                       Reply from 10.10.1.2: bytes=56 Sequence=5 ttl=255 time=82 ms
                     --- 10.10.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 46/76/91 ms


Configuration Scripts
                    ●    CE1
                         #
                         sysname CE1
                         #
                         vlan batch 10
                         #
                         interface Vlanif10
                          ip address 10.10.1.1 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 10
                         #
                         return
                    ●    PE1
                         #
                         sysname PE1
                         #
                         vlan batch 10 20
                         #
                         mpls lsr-id 1.1.1.9
                         mpls
                         #
                         mpls l2vpn
                         #
                         mpls ldp
                         #
                         mpls ldp remote-peer 3.3.3.9
                          remote-ip 3.3.3.9
                         #
                         interface Vlanif10
                          mpls static-l2vc destination 3.3.3.9 transmit-vpn-label 100 receive-vpn-label 200
                         #
                         interface Vlanif20
                          ip address 10.1.1.1 255.255.255.0
                          mpls
                          mpls ldp
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 10
                         #
                         interface 10GE1/0/2
                          port link-type trunk
                          port trunk allow-pass vlan 20
                         #
                         interface LoopBack1
                          ip address 1.1.1.9 255.255.255.255


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                         487
VPN Configuration
VPN Configuration                                                             5 VPWS Configuration

