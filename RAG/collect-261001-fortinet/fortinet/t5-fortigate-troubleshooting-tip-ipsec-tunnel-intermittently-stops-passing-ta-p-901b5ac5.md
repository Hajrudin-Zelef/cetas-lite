---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-ipsec-tunnel-intermittently-stops-passing-ta-p-901b5ac5
title: "t5-fortigate-troubleshooting-tip-ipsec-tunnel-intermittently-stops-passing-ta-p--901b5ac5"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2026-03"]
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-ipsec-tunnel-intermittently-stops-passing-ta-p--901b5ac5.md
source_anchor: ""
source_lines: [1, 4]
sha256: cd2d37ad271465321fedf0b57d4db4d0d3e3606c660b5c58aa1140cd512db58d
---

# t5-fortigate-troubleshooting-tip-ipsec-tunnel-intermittently-stops-passing-ta-p--901b5ac5

Troubleshooting Tip: IPsec tunnel intermittently stops passing traffic after upgrading FortiGate 4xF/6xF series to FortiOS v7.6.4 (Known Issue)
| Description | This article describes a known issue and workaround where, after upgrading FortiGate 4xF/6xF series to FortiOS v7.6.4, IPsec traffic fails when the IPsec tunnel is up. | 
| Scope | FortiGate 4xF/6xF, FortiOS v7.6.4. | 
| Solution | This issue affects both user traffic and the FortiGate’s local-out traffic (such as BGP traffic or SD-WAN Performance SLA probes) sent out via the IPsec tunnel. Packet captures will show outbound traffic egressing the IPsec tunnel interface, but no inbound reply packets.  The issue is related to hardware-acceleration of IPsec traffic and may be temporarily resolved when the IPsec tunnel is rekeyed or manually flushed. As a workaround, disable hardware acceleration for the IPsec tunnel to prevent the issue from occurring (note that disabling IPsec npu-offload will flush the existing IPsec tunnel, resulting in a brief disruption):  config vpn ipsec phase1-interface edit <tunnel> set npu-offload disable next end  This is a Known Issue (tracked by Issue ID 1206506) and it is resolved in FortiOS v7.4.10, Bug ID -1206506 "Traffic disruption occurs when IPsec tunnel manager write sequence issue happens." and v7.6.5, Bug ID - 1206506 "Traffic disruption occurs when IPsec tunnel manager write sequence issue happens." Fix will be available also in FortiOS v8.0.0 which is estimated to be released in March 2026. To confirm a match to the issue, gather the following diagnostic command set multiple times (10 seconds apart) while the issue is actively occurring (i.e., capture when traffic is being dropped by the IPsec tunnel and not when IPsec is behaving normally):  execute time diagnose npu np6xlite dce 0 fnsysctl cat /proc/net/np6xlite_0/fos-perf fnsysctl cat /proc/net/np6xlite_0/ipsec-perf fnsysctl cat /proc/net/np6xlite_0/ipsec diagnose vpn ipsec status diagnose npu np6xlite sse-stats diagnose cp soc4 vpn-stats 0 diagnose vpn tunnel list fnsysctl cat /proc/net/np6xlite_0/ipsec-ob0 fnsysctl cat /proc/net/np6xlite_0/ipsec-ib0  diag debug application ike -1 diag debug console timestamp enable diag debug enable  After collecting the debugging output, disable the debug processes with the following commands:  diagnose debug disable diagnose debug reset  Note: super_admin administrator access is required to run the commands above, particularly the fnsysctl commands.  Once gathered, open a Fortinet TAC ticket and submit a file (or files) containing the output of the diagnostic commands to the ticket for further analysis. |
