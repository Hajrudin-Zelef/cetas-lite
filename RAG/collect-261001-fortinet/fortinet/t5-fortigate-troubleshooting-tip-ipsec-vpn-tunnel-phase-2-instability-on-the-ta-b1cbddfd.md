---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-ipsec-vpn-tunnel-phase-2-instability-on-the-ta-b1cbddfd
title: "t5-fortigate-troubleshooting-tip-ipsec-vpn-tunnel-phase-2-instability-on-the-ta--b1cbddfd"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-ipsec-vpn-tunnel-phase-2-instability-on-the-ta--b1cbddfd.md
source_anchor: ""
source_lines: [1, 4]
sha256: 1399d2a2315e181ba262f64e9e78bab798b1753ba44b4c08240418a4d708bd61
---

# t5-fortigate-troubleshooting-tip-ipsec-vpn-tunnel-phase-2-instability-on-the-ta--b1cbddfd

Troubleshooting Tip: IPsec VPN tunnel phase 2 instability on the NP6xlite platform
| Description | This article describes an issue with IPsec VPN Tunnel Phase 2 instability on the NP6xlite platform. | 
| Scope | FortiGate. | 
| Solution | Firmware:     Troubleshooting:  Perform np6xlite debugging:  diagnose npu np6xlite dce DROP_IPSEC0_ENGINB0:0000000000000683[80] DROP_IPSEC0_ENGINB1:0000000000000002[81]  IKE debugging (shown invalid ESP 4 (replay) SPI from the tunnel):  diagnose vpn ike log filter name "XXXX" diagnose debug application ike -1 diagnose debug enable ike V=root:0:XXXX: invalid ESP 4 (replay) SPI 3fe65c76 seq 00000000:00a02e94 7 Y.Y.Y.Y->Z.Z.Z.Z:0 ike V=root:0:XXXX: invalid ESP 4 (replay) SPI 3fe65c76 seq 00000000:00a02eac 7 Y.Y.Y.Y->Z.Z.Z.Z:0 ike V=root:0:XXXX: invalid ESP 4 (replay) SPI 3fe65c76 seq 00000000:00a02eca 7 Y.Y.Y.Y->Z.Z.Z.Z:0 ike V=root:0:XXXX: invalid ESP 4 (replay) SPI 3fe65c76 seq 00000000:00a02ede 7 Y.Y.Y.Y->Z.Z.Z.Z:0  Workaround:   config vpn ipsec phase2-interface edit "XXXX" set replay disable end diagnose vpn ike gateway filter name "XXXX" diagnose vpn ike gateway flush |
