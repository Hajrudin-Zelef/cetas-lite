---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-ssl-vpn-connection-blocked-by-local-in-policy-18-e8ad2e5c
title: "fortigate-3-troubleshooting-tip-ssl-vpn-connection-blocked-by-local-in-policy-18-e8ad2e5c"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-ssl-vpn-connection-blocked-by-local-in-policy-18-e8ad2e5c.md
source_anchor: ""
source_lines: [1, 4]
sha256: cd785931d355b1d5134af00c2fad7edd1dddfb8f1c6d7a5110400d2bcf1c55c3
---

# fortigate-3-troubleshooting-tip-ssl-vpn-connection-blocked-by-local-in-policy-18-e8ad2e5c

Troubleshooting Tip: SSL VPN connection blocked by local-in policy
| Description | This article describes recommendations on how to resolve cases where the SSL VPN connection is being attempted but gets blocked by the local-in policy, even though the SSL VPN setup is configured and enabled. In this scenario, the FortiGate is supposed to open the port that is configured for the SSL VPN: either the default 443 or the port that gets defined on the SSL VPN settings by the admin.  SSL VPN connections can be blocked by the FortiGate for different reasons, depending on config and restrictions. | 
| Scope | FortiGate, SSL VPN. | 
| Solution |   Starting from FortiOS v7.6.0, it is possible to configure local-in policy via the GUI. For guidance in configuring local-in policy, refer to the article: Technical Tip: Creating a Local-In policy (IPv4 and IPv6) on GUI.  Another way to determine this is to check under the Log & report -> Local traffic page. The traffic can be seen for the VPN connection requests being 'denied' by policy 0. Run debug flow in the firewall CLI:  diagnose debug reset diagnose debug flow filter saddr x.x.x.x                <-- Source IP address of test machine. diagnose debug flow filter port <port/service>          <-- Destination port/service. diagnose debug flow show function-name enable diagnose debug console timestamp enable diagnose debug flow trace start 1000 diagnose debug enable  To stop the debug, run the following command:  diagnose debug disable diagnose debug reset  Running debug flow will also show that the SSL VPN connection is dropped by iprope check (local-in-policy).  Note: As explained in Technical Tip: Upcoming changes on SSL VPN modes starting from v7.6.3. Starting from v7.6.3, the SSL VPN tunnel mode will no longer be supported, and SSL VPN web mode will be called 'Agentless VPN'. |
