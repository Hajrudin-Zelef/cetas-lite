---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-ssl-vpn-user-authentication-issues-when-firewall-is-in-815808e7
title: "fortigate-3-technical-tip-ssl-vpn-user-authentication-issues-when-firewall-is-in-815808e7"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-ssl-vpn-user-authentication-issues-when-firewall-is-in-815808e7.md
source_anchor: ""
source_lines: [1, 24]
sha256: b09e0899ef0f8ef847f0da9f850c2bbf9f433440c502b0a70e53a664102eb562
---

# fortigate-3-technical-tip-ssl-vpn-user-authentication-issues-when-firewall-is-in-815808e7

Technical Tip: SSL VPN user authentication issues when firewall is in NGFW policy-based mode
Description
This article describes some common errors when configuring SSL VPN on a FortiGate running NGFW policy-based mode.
Scope
FortiGate.
Solution
Issue 1: SSL VPN is unreachable.
In this example, TCP port 20443 has been configured in SSL VPN settings on the external interface, and the TCP SYN packet reaches the firewall, but it does not respond.
The Local-In Policy list shows no open TCP or UDP 20443 port:
- Verify SSL VPN is enabled and fully configured in VPN -> SSL-VPN Settings.
- Verify an SSL Inspection & Authentication Rule referencing the SSL VPN interface exists:
After configuring both items, the SSL VPN will run and listen on the intended port for connection attempts.
Issue 2: Unexpected 'Permission Denied' when using the intended VPN user and password
Verify the intended VPN user or group is referenced in at least one of the following locations:
- SSL VPN Authentication/Portal Mapping
- SSL Inspection & Authentication Policy.
Issue 3: The User can connect to the VPN, but has no access to the intended resources.
- Verify any Central SNAT policies to ensure the correct setting is applied (Source NAT should usually be disabled for tunnel traffic)
- Verify SSL VPN split tunnel routing configuration. If inheriting from firewall policy, note this is based on destination addresses in 'SSL Inspection & Authentication'.
- Verify appropriate 'Security Policy' and 'SSL Inspection & Authentication' policies both exist.
Apart from the requirement to configure appropriate policies in 'Security Policy' and 'SSL Inspection & Authentication', SSL VPN troubleshooting for NGFW policy-based mode is otherwise very similar to profile-based mode.
Related articles:
Troubleshooting Tip: SSL VPN Troubleshooting
Technical Tip: A quick guide to FortiGate SSL VPN authentication and common issues and misunderstandings
