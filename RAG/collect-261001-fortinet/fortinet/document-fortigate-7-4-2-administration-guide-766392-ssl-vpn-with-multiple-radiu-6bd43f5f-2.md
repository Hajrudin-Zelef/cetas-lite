---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-2-administration-guide-766392-ssl-vpn-with-multiple-radiu-6bd43f5f-2
title: "diagnose sniffer packet any 'port 1812' 4 0 l"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2020-05-15"]
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-2-administration-guide-766392-ssl-vpn-with-multiple-radiu-6bd43f5f.md
source_anchor: ""
source_lines: [45, 139]
sha256: 7787dd0fb94b0e3b23f0f4ec729aac1c70a3b19fce931b1827309606a8867b9f
---

# diagnose sniffer packet any 'port 1812' 4 0 l

  - Select All Other Users/Groups and click Edit.
  - From the Portal dropdown, select web-access.
  - Click OK.
- Create a web portal for PrimarySecondaryGroup.			
  - Under Authentication/Portal Mapping, click Create New.
  - Click Users/Groups and select PrimarySecondaryGroup.
  - From the Portal dropdown, select full-access.
  - Click OK.
To configure SSL VPN firewall policy:
- Go to Policy & Objects > Firewall Policy.
- Click Create New to create a new policy, or double-click an existing policy to edit it and configure the following settings:Name Enter a name for the policy. Incoming Interface SSL-VPN tunnel interface (ssl.root) Outgoing interface Set to the local network interface so that the remote user can access the internal network. For this example, select port3. Source In the Address tab, select SSLVPN_TUNNEL_ADDR1 In the User tab, select PrimarySecondaryGroup Destination Select the internal protected subnet 192.168.20.0. Schedule always Service All Action Accept NAT Enable
- Configure any remaining firewall and security options as required.
- Click OK.
To configure SSL VPN using the CLI:
- Configure the internal interface and firewall address:			config system interface edit "port3" set vdom "root" set ip 192.168.20.5 255.255.255.0 set alias "internal" next end config firewall address edit "192.168.20.0" set uuid cc41eec2-9645-51ea-d481-5c5317f865d0 set subnet 192.168.20.0 255.255.255.0 next end
- Configure the RADIUS server:		config user radius edit "PrimarySecondary" set server "192.168.20.6" set secret <secret> set secondary-server "192.168.2.71" set secondary-secret <secret> next end
- Add the RADIUS user to the user group:
			config user group edit "PrimarySecondaryGroup" set member "PrimarySecondary " next end
- Configure SSL VPN settings:			config vpn ssl settings set servercert "server_certificate" set tunnel-ip-pools "SSLVPN_TUNNEL_ADDR1" set source-interface "port1" set source-address "all" set default-portal "web-access" config authentication-rule edit 1 set groups "PrimarySecondaryGroup " set portal "full-access" next end end
- Configure one SSL VPN firewall policy to allow remote users to access the internal network:			config firewall policy edit 1 set name "sslvpn-radius" set srcintf "ssl.root" set dstintf "port3" set srcaddr "all" set dstaddr "192.168.20.0" set groups "PrimarySecondaryGroup" set action accept set schedule "always" set service "ALL" set nat enable next end
To verify the connection:
User radkeith is a member of both the NPS server and the FAC server.
When the Primary server is up, it will connect to the SSL VPN tunnel using FortiClient.
# diagnose sniffer packet any 'port 1812' 4 0 l
interfaces=[any]
filters=[port 1812]
2020-05-15 16:26:50.838453 port3 out 192.168.20.5.2374 -> 192.168.20.6.1812: udp 118
2020-05-15 16:26:50.883166 port3 in 192.168.20.6.1812 -> 192.168.20.5.2374: udp 20
2020-05-15 16:26:50.883374 port3 out 192.168.20.5.2374 -> 192.168.20.6.1812: udp 182
2020-05-15 16:26:50.884683 port3 in 192.168.20.6.1812 -> 192.168.20.5.2374: udp 228
The access request is sent to the Primary NPS server 192.168.20.6, and the connection is successful.
# get vpn ssl monitor
SSL VPN Login Users:
Index   User            Group                   Auth Type      Timeout         From     HTTP in/out    HTTPS in/out
0       radkeith       PrimarySecondaryGroup   2(1)            285     192.168.2.202          0/0     0/0
SSL VPN sessions:
Index   User       Group                   Source IP      Duration        I/O Bytes       Tunnel/Dest IP
0       radkeith   PrimarySecondaryGroup   192.168.2.202   62              132477/4966    10.212.134.200
When the Primary server is down, and the Secondary server is up, the connection is made to the SSLVPN tunnel again:
# diagnose sniffer packet any 'port 1812' 4 0 l
interfaces=[any]
filters=[port 1812]
2020-05-15 16:31:23.016875 port3 out 192.168.20.5.7989 -> 192.168.20.6.1812: udp 118
2020-05-15 16:31:28.019470 port3 out 192.168.20.5.7989 -> 192.168.20.6.1812: udp 118
2020-05-15 16:31:30.011874 port1 out 192.168.2.5.23848 -> 192.168.2.71.1812: udp 118
2020-05-15 16:31:30.087564 port1 in 192.168.2.71.1812 -> 192.168.2.5.23848: udp 20
Access request is sent to the Primary NPS server 192.168.20.6, but there was no response. RADIUS authentication falls through to the Secondary FortiAuthenticator 192.168.2.71, and the authentication was accepted. The VPN connection is established.
# get vpn ssl monitor
SSL VPN Login Users:
Index   User            Group                  Auth Type      Timeout         From     HTTP in/out    HTTPS in/out
0       radkeith        PrimarySecondaryGroup   2(1)            287     192.168.2.202        0/0        0/0
SSL VPN sessions:
Index   User            Group                    Source IP      Duration        I/O Bytes       Tunnel/Dest IP
0       radkeith        PrimarySecondaryGroup   192.168.2.202   48              53544/4966     10.212.134.200
Authenticating to two RADIUS servers concurrently
There are times where users are located on separate RADIUS servers. This may be the case when migrating from an old server to a new one for example. In this scenario, a Windows NPS server and a FortiAuthenticator are configured in the same User Group. The access-request is sent to both servers concurrently. If FortiGate receives an access-accept from either server, authentication is successful.
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
To configure the first RADIUS server:
- Go to User & Authentication > RADIUS Servers and click Create New.
- Set Name to win2k16.
- Leave Authentication method set to Default. The PAP, MS-CHAPv2, and CHAP methods will be tried in order.
- Under Primary Server, set IP/Name to 192.168.20.6 and Secret to the shared secret configured on the RADIUS server.
- Click Test Connectivity to test the connection to the server, and ensure that Connection status is Successful.
- Click OK.
To configure the second RADIUS server:
- Go to User & Authentication > RADIUS Servers and click Create New.
- Set Name to fac.
- Leave Authentication method set to Default. The PAP, MS-CHAPv2, and CHAP methods will be tried in order.
- Under Primary Server, set IP/Name to 192.168.2.71 and Secret to the shared secret configured on the RADIUS server.
- Click Test Connectivity to test the connection to the server, and ensure that Connection status is Successful.
- Click OK.
To configure the user group:
- Go to User & Authentication > User Groups and click Create New.
- In the Name field, enter dualPrimaryGroup..
- In the Remote Groups area, click Add, and from the Remote Server dropdown, select fac.
- Click Add again. From the Remote Server dropdown select win2k16 and click OK.
- Click OK, and then click OK again.
To configure the SSL VPN settings:
- Go to VPN > SSL-VPN Settings.
- From the Listen on Interface(s) dropdown select port1.
- In the Listen on Port field enter 10443.
- Optionally, from the Server Certificate dropdown, select the authentication certificate if you have one for this SSL VPN portal.
- Under Authentication/Portal Mapping, set the default portal web-access.
			
