---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-2-administration-guide-766392-ssl-vpn-with-multiple-radiu-6bd43f5f-1
title: "diagnose sniffer packet any 'port 1812' 4 0 l"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-2-administration-guide-766392-ssl-vpn-with-multiple-radiu-6bd43f5f.md
source_anchor: ""
source_lines: [1, 44]
sha256: 272d32544ef72a0864503d7631c1a87dbb6e67e0678db4acd6725cbb215b478f
---

# diagnose sniffer packet any 'port 1812' 4 0 l

SSL VPN with multiple RADIUS servers
SSL VPN with multiple RADIUS servers
When configuring two or more RADIUS servers, you can configure a Primary and Secondary server within the same RADIUS server configurations for backup purposes. You can also configure multiple RADIUS servers within the same User Group to service the access request at the same time.
|  | A tertiary server can be configured in the CLI. | 
Sample topology
Sample configurations
- Configure a Primary and Secondary server for backup
- Authenticating to two RADIUS servers concurrently
Configure a Primary and Secondary server for backup
When you define a Primary and Secondary RADIUS server, the access request will always be sent to the Primary server first. If the request is denied with an Access-Reject, then the user authentication fails. However, if there is no response from the Primary server after another attempt, the access request will be sent to the Secondary server.
In this example, you will use a Windows NPS server as the Primary server and a FortiAuthenticator as the Secondary server. It is assumed that users are synchronized between the two servers.
To configure the internal and external interfaces:
- Go to Network > Interfaces.
- Edit the port1 interface and set IP/Network Mask to 192.168.2.5/24.
- Edit the port2 interface and set IP/Network Mask to 192.168.20.5/24.
- Click OK.
To create a firewall address:
- Go to Policy & Objects > Addresses and select Address.
- Click Create new.
- Set Name to 192.168.20.0.
- Leave Type as Subnet
- Set IP/Netmask to 192.168.20.0/24.
- Click OK.
To add the RADIUS servers:
- Go to User & Authentication > RADIUS Servers and click Create New.
- Set Name to PrimarySecondary.
- Leave Authentication method set to Default. The PAP, MS-CHAPv2, and CHAP methods will be tried in order.
- Under Primary Server, set IP/Name to 192.168.20.6 and Secret to the shared secret configured on the RADIUS server.
- Click Test Connectivity to test the connection to the server, and ensure that Connection status is Successful.
- Under Secondary Server, set IP/Name to 192.168.2.71 and Secret to the shared secret configured on the RADIUS server.
- Click Test Connectivity to test the connection to the server, and ensure that Connection status is Successful.
- Click OK.
To configure the user group:
- Go to User & Authentication > User Groups and click Create New.
- In the Name field, enter PrimarySecondaryGroup.
- In the Remote Groups area, click Add, and from the Remote Server dropdown, select PrimarySecondary.
- Click OK, and then click OK again.
To configure the SSL VPN settings:
- Go to VPN > SSL-VPN Settings.
- From the Listen on Interface(s) dropdown select port1.
- In the Listen on Port field enter 10443.
- Optionally, from the Server Certificate dropdown, select the authentication certificate if you have one for this SSL VPN portal.
- Under Authentication/Portal Mapping, set the default portal web-access.
			
