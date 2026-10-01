---
id: collect-261001-fortinet/fortinet/support-forum-92-ikev2-ipsec-troubleshooting-tips-needed-on-fortios-7-2-8-163325-d374f5e8
title: "support-forum-92-ikev2-ipsec-troubleshooting-tips-needed-on-fortios-7-2-8-163325-d374f5e8"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "Nvidia"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/support-forum-92-ikev2-ipsec-troubleshooting-tips-needed-on-fortios-7-2-8-163325-d374f5e8.md
source_anchor: ""
source_lines: [1, 36]
sha256: c5388002941206bad931ac6fe898bfd243047b8cbb9799f5c24076f78094d96a
---

# support-forum-92-ikev2-ipsec-troubleshooting-tips-needed-on-fortios-7-2-8-163325-d374f5e8

IKEv2 IPSec Troubleshooting Tips needed on FortiOS 7.2.8
Hello,
in a transition from another manufacturers router to Fortigate I have to transpose IKEv2 Dial-In Client configurations into FortiOS 7.2.8. Although the necessary parameters should match. The connections do not work. Therefore I tried to troubleshoot the IKE connections with the following CLI commands:
diagnose debug application ike -1
diagnose vpn ike log-filter src-addr4 <remote WAN-IP>
alternatively: diagnose vpn ike log-filter dst-addr4 <local WAN-IP>
diagnose debug console timestamp enable
diagnose debug enable
But there are only empty lines. I double-checked that the dial-up client could reach the Fortigate and successfully pcap´ed.
Although I am new to FortiOS I would bet that there is a debug command to have a live view / monitor the setup negotiations of IKE phase 1 and IPSec phase 2 connections. But I have no clue how start these. Your hints would be appreciated.
The Admin-Guide https://docs.fortinet.com/document/fortigate/7.2.8/administration-guide/834425/understanding-vpn-related-logs "Understanding VPN related Logs" mentions these logs, but I don´t know where to find them.
Any suggestions?
Here is a sample IKEv2 configuration. The NCP VPN software client f. Windows is used with it.
config vpn ipsec phase1-interface
edit "SAMPLE"
set type dynamic
set interface "port10"
set ike-version 2
set peertype one
set net-device disable
set mode-cfg enable
set proposal aes256-sha256
set dpd on-idle
set dhgrp 14
set peerid "xy@xyz.com"
set assign-ip-from dhcp
set dns-mode auto
set ipv4-split-include "INTRANET Subnet"
set psksecret ENC <Encrypted PSK>
set dpd-retryinterval 60
config vpn ipsec phase2-interface
edit "SAMPLE"
set phase1name "SAMPLE"
set proposal aes256-sha256
set dhgrp 14
next
