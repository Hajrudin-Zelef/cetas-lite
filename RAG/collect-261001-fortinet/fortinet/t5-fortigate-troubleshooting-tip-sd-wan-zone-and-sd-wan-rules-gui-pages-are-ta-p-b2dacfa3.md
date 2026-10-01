---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-sd-wan-zone-and-sd-wan-rules-gui-pages-are-ta-p-b2dacfa3
title: "t5-fortigate-troubleshooting-tip-sd-wan-zone-and-sd-wan-rules-gui-pages-are-ta-p-b2dacfa3"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-sd-wan-zone-and-sd-wan-rules-gui-pages-are-ta-p-b2dacfa3.md
source_anchor: ""
source_lines: [1, 5]
sha256: 5530d6b1334b7222d7bbcaae31f5ced3ff4204e3dd098d3a9203de1b59b45419
---

# t5-fortigate-troubleshooting-tip-sd-wan-zone-and-sd-wan-rules-gui-pages-are-ta-p-b2dacfa3

| Description | This article describes how to work around the issue where, after a firmware upgrade to FortiOS 7.0.12 or 7.2.5, the **Network -> SD-WAN -> SD-WAN Zone** and**Network -> SD-WAN -> SD-WAN** **Rules** GUI pages get stuck loading.   The **Network -> SD-WAN -> SD-WAN Performance SLAs** GUI page is not affected.  There is a known GUI issue on FortiOS 7.0.12 and 7.2.5 in which the **Network -> SD-WAN- > SD-WAN Zone** and**Network -> SD-WAN -> SD-WAN** **Rules** GUI pages will not load under the following conditions:    Note that this is a cosmetic issue only and does not affect the operations of SD-WAN or ADVPN.  This issue is documented in the FortiOS 7.0.12 and 7.2.5 release notes under bug ID 907041 and is scheduled to be fixed in FortiOS 7.0.13 and 7.2.6. | 
| Scope | FortiGate v7.0.12 or v7.2.5. | 
| Solution | To confirm that this issue is being encountered, it is possible to check FortiGate’s crash logs by running:   **diagnose debug crashlog read**  There will be a HTTPSD crash generated each time the **Network -> SD-WAN -> SD-WAN Zone** or**Network -> SD-WAN > SD-WAN** **Rules** GUI pages are accessed and the issue occurs. To workaround the issue, it is possible to disable the shortcut tunnel(s) by navigating to **Dashboard -> Network -> IPsec Monitor** , 'right-clicking' the shortcut tunnel (which is designated by the suffix _# in the VPN tunnel name), and bringing the entire shortcut tunnel down:   When there are no shortcut tunnels active, the **Network -> SD-WAN- > SD-WAN Zone** and**Network -> SD-WAN -> SD-WAN** **Rules** pages will load properly:    **Note:** If there is active spoke-to-spoke traffic on the network, the shortcut tunnel will re-establish shortly after manually bringing it down. It may also not be feasible to bring down all ADVPN shortcut tunnels. In both cases, if modifications need to be made to the SD-WAN Zones or SD-WAN Rules configuration, make the changes in the CLI instead.  **Related document:** FortiGate 7.2.5 known issues. | 

The Fortinet Security Fabric brings together the concepts of convergence and consolidation to provide comprehensive cybersecurity protection for all users, devices, and applications and across all network edges.
