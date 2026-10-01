---
id: collect-261001-ia-llm/ia-llm/document-fortigate-7-2-3-administration-guide-183204-ssl-vpn-with-fortitoken-mob-c7e6d0fe-2
title: "SSL VPN with FortiToken mobile push authentication"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-ia-llm/document-fortigate-7-2-3-administration-guide-183204-ssl-vpn-with-fortitoken-mob-c7e6d0fe.md
source_anchor: ""
source_lines: [181, 220]
sha256: 37f9c569ec9f8295d756905428be0825b5ec415e3f6da1fd42c8d28b34244b00
---

# SSL VPN with FortiToken mobile push authentication

1. From a remote device, use a web browser to log into the SSL VPN web portal *http://172.20.120.123:10443* .
2. Log in using the *sslvpnuser1* credentials.The FortiGate pushes a login request notification through the FortiToken mobile application.
3. Check your mobile device and select *Approve* .When the authentication is approved, *sslvpnuser1* is logged into the SSL VPN portal.
4. On the FortiGate, go to *Dashboard > Network* and expand the*SSL-VPN* widget to verify the user's connection.

###### To see the results of tunnel connection:

1. Download FortiClient from www.forticlient.com.
2. Open the FortiClient Console and go to *Remote Access > Configure VPN* .
3. Add a new connection:
  1. Set the connection name.
  2. Set *Remote Gateway* to the IP of the listening FortiGate interface, in this example,*172.20.120.123* .
  3. Select *Customize Port* and set it to*10443* .
4. Save your settings.
5. Log in using the *sslvpnuser1* credentials and click*FTM Push* .The FortiGate pushes a login request notification through the FortiToken mobile application.
6. Check your mobile device and select *Approve* .When the authentication is approved, *sslvpnuser1* is logged into the SSL VPN tunnel.

###### To check the SSL VPN connection using the GUI:

1. Go to *Dashboard > Network* and expand the*SSL-VPN* widget to verify the user's connection.
2. Go to *Log & Report > Forward Traffic* to view the details of the SSL VPN traffic.

###### To check the web portal login using the CLI:

get vpn ssl monitor
SSL VPN Login Users:
 Index   User           Auth Type    Timeout   From           HTTP in/out   HTTPS in/out
 0       sslvpnuser1    1(1)         229       10.1.100.254   0/0           0/0
SSL VPN sessions:
 Index   User    Source IP      Duration        I/O Bytes       Tunnel/Dest IP 

###### To check the tunnel login using the CLI:

get vpn ssl monitor
SSL VPN Login Users:
 Index   User          Auth Type    Timeout   From           HTTP in/out  HTTPS in/out
 0       sslvpnuser1   1(1)         291       10.1.100.254   0/0          0/0
SSL VPN sessions:
 Index   User          Source IP      Duration   I/O Bytes      Tunnel/Dest IP 
 0       sslvpnuser1   10.1.100.254   9          22099/43228    10.212.134.200
