---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-ssl-vpn-login-error-due-to-tunnel-interface-ta-813fd64f
title: "t5-fortigate-troubleshooting-tip-ssl-vpn-login-error-due-to-tunnel-interface-ta--813fd64f"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-ssl-vpn-login-error-due-to-tunnel-interface-ta--813fd64f.md
source_anchor: ""
source_lines: [1, 4]
sha256: ff664ea368e2ee85b4ea53ea92f06100189b7585a47bdbe8758bb1e2613ff80b
---

# t5-fortigate-troubleshooting-tip-ssl-vpn-login-error-due-to-tunnel-interface-ta--813fd64f

Troubleshooting Tip: SSL VPN login error due to tunnel interface being down
| Description | This article describes how to fix an error that occurs with SSL VPN login where the user is informed that the tunnel interface is down. | 
| Scope | FortiGate 6.X and 7.X | 
| Solution | SSL VPN login error due to tunnel Interface down.    config system interface edit ssl.root show config system interface edit "ssl.root" set vdom "root" set allowaccess fabric set status down -> [Tunnel status is down] set type tunnel set snmp-index 4 next end   The VPN interface will appear greyed out (as shown in the screenshot above) when it is administratively down.    Note: If 'Fabric' is enabled for allowaccess, the tunnel status can not be changed.  Related articles: |
