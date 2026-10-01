---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-macos-iphone-users-never-connect-to-ssl-vpn-usin-63fb1ab4
title: "fortigate-3-troubleshooting-tip-macos-iphone-users-never-connect-to-ssl-vpn-usin-63fb1ab4"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-macos-iphone-users-never-connect-to-ssl-vpn-usin-63fb1ab4.md
source_anchor: ""
source_lines: [1, 4]
sha256: e5264e8d78cac7ec84575e45ae01816457f9bd223dba5927dc2717d281c26096
---

# fortigate-3-troubleshooting-tip-macos-iphone-users-never-connect-to-ssl-vpn-usin-63fb1ab4

Troubleshooting Tip: MacOS/iPhone users never connect to SSL VPN using FortiClient
| Description | This article describes how to fix an issue where, even with the right credentials, users are unable to connect to the VPN their system either shows an endless connecting error or states the VPN connection is down. | 
| Scope | FortiGate, FortiClient | 
| Solution | When the user is trying to connect to the VPN, check the following two places:   VPN logs:   SSL VPN debugging:  diagnose debug application sslvpn -1 diagnose debug application fnbamd diagnose debug enable   sslvpn_dtls_timeout_check:312 waiting for client hello timeout   The MacOS and iPhone (free) versions of FortiClient have no option to enable DTLS. All newer versions of FortiGate have it enabled for better performance. This causes FortiGate to wait for the FortiClient to make the DTLS connection (which is not enabled), leading to a failure that brings down the whole tunnel.  Make sure to disable the DTLS option on FortiGate, test out the connection, and also monitor the SSL VPN performance. To disable DTLS on SSL VPN, run the following commands:  config vpn ssl setting set dtls-tunnel disable end  This has been enabled by default since 5.4.  If assistance is needed, contact Fortinet support.  Related article: Technical Tip: Using DTLS to improve SSL VPN performance. |
