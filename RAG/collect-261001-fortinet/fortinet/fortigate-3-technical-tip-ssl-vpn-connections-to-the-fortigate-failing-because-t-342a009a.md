---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-ssl-vpn-connections-to-the-fortigate-failing-because-t-342a009a
title: "fortigate-3-technical-tip-ssl-vpn-connections-to-the-fortigate-failing-because-t-342a009a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-ssl-vpn-connections-to-the-fortigate-failing-because-t-342a009a.md
source_anchor: ""
source_lines: [1, 4]
sha256: 8b03912d074793642807ee496a29610dda3f3e25ab2e16d6413fc0b2585054b2
---

# fortigate-3-technical-tip-ssl-vpn-connections-to-the-fortigate-failing-because-t-342a009a

Technical Tip: SSL VPN connections to the FortiGate failing because the sslvpnd process has not started
| Description | This article describes a known behavior where SSL VPN users are unable to connect successfully because the sslvpnd process has not started. The following symptoms can be observed in this scenario:       | 
| Scope | FortiGate. | 
| Solution | This issue occurs if there are no active Firewall Policies on the FortiGate that have the 'SSL-VPN tunnel interface (ssl.root)' set in the Incoming Interface field (i.e., if no relevant Firewall Policies exist or if they are all administratively disabled).  Without an active Firewall Policy, the sslvpnd daemon will not be active and will not listen for/accept any incoming connections. Additionally, the SSL VPN debugs (diagnose debug application sslvpn -1) will not show any output.  Below is the output of the diagnose sys top command. Note that 'sslvpnd' is not in the running processes list.    Likewise, the output of diagnose sys tcpsock \| grep <SSL-VPN Port> will show that sslvpnd is not listening on the configured port:  FortiGate # diagnose sys tcpsock \| grep 443 0.0.0.0:443->0.0.0.0:0->state=listen err=0 socktype=2 rma=0 wma=0 fma=0 tma=0 inode=28368 process=293/wad 10.255.1.1:443->0.0.0.0:0->state=listen err=0 socktype=1 rma=0 wma=0 fma=0 tma=0 inode=31637 process=239/fltund  Resolving the Issue:  To resolve the issue, create at least one active firewall policy under Policy & Objects -> Firewall Policy to allow traffic from the SSL VPN tunnel interface (ssl.root) to another interface. Below is an example of a firewall policy allowing traffic from the SSL VPN tunnel interface to the LAN network behind port 5.    After creating the firewall policy, the sslvpnd daemon will be started, and users will be able to connect to the VPN.    FortiGate # diagnose sys tcpsock \| grep 443 0.0.0.0:443->0.0.0.0:0->state=listen err=0 socktype=3 rma=0 wma=0 fma=0 tma=0 inode=2641407 process=3148/sslvpnd 0.0.0.0:443->0.0.0.0:0->state=listen err=0 socktype=2 rma=0 wma=0 fma=0 tma=0 inode=28368 process=293/wad 10.255.1.1:443->0.0.0.0:0->state=listen err=0 socktype=1 rma=0 wma=0 fma=0 tma=0 inode=31637 process=239/fltund    Note: Starting v7.6.3, the SSL VPN tunnel mode will no longer be supported, and SSL VPN web mode will be called 'Agentless VPN' as explained in Upcoming changes on SSL VPN modes starting from v7.6.3.  Related article: Troubleshooting Tip: SSL VPN Troubleshooting |
