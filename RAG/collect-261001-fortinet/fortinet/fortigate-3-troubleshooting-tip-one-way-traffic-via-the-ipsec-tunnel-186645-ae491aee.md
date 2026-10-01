---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-one-way-traffic-via-the-ipsec-tunnel-186645-ae491aee
title: "fortigate-3-troubleshooting-tip-one-way-traffic-via-the-ipsec-tunnel-186645-ae491aee"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-one-way-traffic-via-the-ipsec-tunnel-186645-ae491aee.md
source_anchor: ""
source_lines: [1, 4]
sha256: 9ba6e8d4b7f213ff51535498c37026fb62990d76d76f05e4cdd52fad1ec789a4
---

# fortigate-3-troubleshooting-tip-one-way-traffic-via-the-ipsec-tunnel-186645-ae491aee

Troubleshooting Tip: One-way traffic via the IPsec tunnel
| Description | This article describes how to troubleshoot one-way traffic over the IPSec tunnel between 2 FortiGates. | 
| Scope | FortiGate. | 
| Solution | Topology:   The machine on subnet 10.122.0.0/20 can reach (ping) devices on subnet 10.171.0.0/20, but not the other way around. It is necessary to check:   get router info routing-table all get router info routing-table database  diagnose vpn ike gateway list name <tunnel name> diagnose vpn tunnel list name <tunnel name>    Example:  diagnose debug flow filter addr 10.171.0.10 10.122.0.10 and <----- This will capture traffic from and to these 2 addresses. diagnose debug flow filter proto 1 <----- ICMP protocol. diagnose debug flow show function-name enable diagnose debug flow show iprope enable diagnose debug flow trace start 100 diagnose debug enable  To disable:  diagnose debug disable diagnose debug reset The debug output will show how FortiGate processes the traffic. Possible reasons:    Packet capture on FGT-B:  Disable NPU offloading in phase 1 of the IPsec tunnels and capture IKE and ESP traffic on both devices. config vpn ipsec phase1-interface     edit <tunnel name>         set npu-offload disable     next end Note: The above NPU setting may not be required on the FortiGate VM devices. Flush the tunnel: diagnose vpn ike gateway flush name <tunnel name> diagnose vpn tunnel flush <tunnel name> Capture IPsec traffic: diagnose sniffer packet any 'host <Remote gateway IP of the tunnel endpoint>' 6 0 l To get the capture in the Wireshark format, run it from the FortiGate GUI: Network -> Diagnostics. If no NAT traversal is detected (no UPD encapsulated packets) and only one-way ESP traffic is visible, then check the NAT-T setting in phase 1 and try enabling NAT-T and flushing the tunnel. If this doesn't solve the problem, then use Forced NAT-T. This will allow FortiGate to use UDP 4500 encapsulation for data packets even if NAT-T is not detected. The NAT-T detection varies from tunnel to tunnel, even if the tunnels terminate on the same device. Each tunnel may use a different path, and intermediate devices sometimes block ESP (protocol 50) packets. By enabling/forcing NAT-T, the ESP packets are encapsulated in UDP packets. It is recommended to use the same NAT-T setting on both tunnel endpoints. In case packets are already UDP 4500 encapsulated, and one tunnel endpoint doesn't receive these packets, it is recommended to consult the ISP for any blocks. Related documents: Technical Tip: Source IP for self-originating IPsec tunnel traffic Debugging the packet flow |
