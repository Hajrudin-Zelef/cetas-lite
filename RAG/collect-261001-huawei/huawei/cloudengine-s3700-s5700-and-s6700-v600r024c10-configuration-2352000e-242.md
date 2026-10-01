---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-242
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2008-07-24", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [35542, 35617]
sha256: c1c299db269ea34764dc67423f5b59a1cbe310412e9ef599e5cdeffb013da51e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                               5 VPWS Configuration

                     BFD for PW                : available
                       BFD sessionIndex            : 257        BFD state : up
                     manual fault              : not set
                     active state           : inactive
                     forwarding entry            : existent
                     link state           : up
                     local VC MTU               : 4470          remote VC MTU        : 4470
                     Local VCCV         : cw alert lsp-ping bfd
                     Remote VCCV           : cw alert lsp-ping bfd
                     local control word          : enable        remote control word : enable
                     tunnel policy name             : --
                     traffic behavior name : --
                     PW template name                 : 1to3
                     primary or secondary : secondary
                     VC tunnel/token info : 1 tunnels/tokens
                       NO.0 TNL type : lsp , TNL ID : 0x1002006
                     create time             : 0 days, 1 hours, 17 minutes, 42 seconds
                     up time               : 0 days, 1 hours, 15 minutes, 55 seconds
                     last change time            : 0 days, 1 hours, 15 minutes, 55 seconds
                     VC last up time : 2008-07-24 12:31:31
                     VC total up time: 0 days, 2 hours, 12 minutes, 51 seconds
                     CKey              : 17
                     NKey               : 18
                    reroute policy            : delay 30 s, resume 10 s
                    reason of last reroute : --
                    time of last reroute : -- days, -- hours, -- minutes, -- seconds
                    delay timer ID             : --          residual time :--
                    resume timer ID               : --        residual time :--

                    # Shut down 10GE1/0/2 on PE1 to simulate a failure on the primary PW. In this
                    case, CE1 cannot use the primary address to ping 10.1.1.2 of CE2. After the
                    secondary PW takes over services, CE1 can use the secondary IP address to ping
                    10.1.2.2 of CE2 successfully.
                    <CE1> ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                       Request time out
                       Request time out
                       Request time out
                       Request time out
                       Request time out
                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       0 packet(s) received
                       100.00% packet loss
                    <CE1> ping 10.1.2.2
                     PING 10.1.2.2: 56 data bytes, press CTRL_C to break
                       Reply from 10.1.2.2: bytes=56 Sequence=1 ttl=255 time=140 ms
                       Reply from 10.1.2.2: bytes=56 Sequence=2 ttl=255 time=160 ms
                       Reply from 10.1.2.2: bytes=56 Sequence=3 ttl=255 time=160 ms
                       Reply from 10.1.2.2: bytes=56 Sequence=4 ttl=255 time=160 ms
                       Reply from 10.1.2.2: bytes=56 Sequence=5 ttl=255 time=160 ms
                     --- 10.1.2.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 140/156/160 ms

                    # Run the display mpls l2vc interface command on PEs again to check the PW
                    status. The command outputs show that the VC state field displays down and the
                    BFD for PW field displays unavailable for the primary PW. The VC state field
                    displays up, the BFD for PW field displays available, and the BFD state field
                    displays up for the secondary PW.
                    <PE1> display mpls l2vc interface 10ge 1/0/1
                     *client interface  : 10GE1/0/1 is up
                      session state    : down
                      AC status        : up


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                  564
VPN Configuration
VPN Configuration                                                                                 5 VPWS Configuration

