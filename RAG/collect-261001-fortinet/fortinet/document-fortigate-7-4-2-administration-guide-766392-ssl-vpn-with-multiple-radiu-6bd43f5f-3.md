---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-2-administration-guide-766392-ssl-vpn-with-multiple-radiu-6bd43f5f-3
title: "diagnose sniffer packet any 'port 1812' 4 0 l"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2020-05-15"]
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-2-administration-guide-766392-ssl-vpn-with-multiple-radiu-6bd43f5f.md
source_anchor: ""
source_lines: [140, 159]
sha256: 6f3c7d5efd93d9f28c4d860e3f1c6305daa6a906853833e83cc465e6969d28a7
---

# diagnose sniffer packet any 'port 1812' 4 0 l

  - Select All Other Users/Groups and click Edit.
  - From the Portal dropdown, select web-access.
  - Click OK.
- Create a web portal for PrimarySecondaryGroup.			
  - Under Authentication/Portal Mapping, click Create New.
  - Click Users/Groups and select dualPrimaryGroup.
  - From the Portal dropdown, select full-access.
  - Click OK.
To configure SSL VPN firewall policy:
- Go to Policy & Objects > Firewall Policy.
- Click Create New to create a new policy, or double-click an existing policy to edit it.
			Name Enter a name for the policy. Incoming Interface SSL-VPN tunnel interface (ssl.root) Outgoing interface Set to the local network interface so that the remote user can access the internal network. For this example, select port3. Source In the Address tab, select SSLVPN_TUNNEL_ADDR1 In the User tab, select dualPrimaryGroup Destination Select the internal protected subnet 192.168.20.0. Schedule always Service All Action Accept NAT Enable
- Configure any remaining firewall and security options as required.
- Click OK.
To configure SSL VPN using the CLI:
- Configure the internal interface and firewall address:		config system interface edit "port3" set vdom "root" set ip 192.168.20.5 255.255.255.0 set alias "internal" next end config firewall address edit "192.168.20.0" set uuid cc41eec2-9645-51ea-d481-5c5317f865d0 set subnet 192.168.20.0 255.255.255.0 next end
- Configure the RADIUS server:config user radius edit "win2k16" set server "192.168.20.6" set secret <secret> next edit "fac" set server "192.168.2.71" set secret <secret> next end
- Add the RADIUS user to the user group:config user group edit "dualPrimaryGroup" set member "win2k16" "fac" next end
- Configure SSL VPN settings:config vpn ssl settings set servercert "server_certificate" set tunnel-ip-pools "SSLVPN_TUNNEL_ADDR1" set source-interface "port1" set source-address "all" set default-portal "web-access" config authentication-rule edit 1 set groups "dualPrimaryGroup" set portal "full-access" next end end
- Configure one SSL VPN firewall policy to allow remote users to access the internal network:config firewall policy edit 1 set name "sslvpn-radius" set srcintf "ssl.root" set dstintf "port3" set srcaddr "all" set dstaddr "192.168.20.0" set groups "dualPrimaryGroup" set action accept set schedule "always" set service "ALL" set nat enable next end To verify the connection:User fackeith is a member of the FortiAuthenticator server only. User radkeith is a member of both the NPS server and the FortiAuthenticator server, but has different passwords on each server. Case 1: Connect to the SSLVPN tunnel using FortiClient with user FacAdmin:# diagnose sniffer packet any 'port 1812' 4 0 l interfaces=[any] filters=[port 1812] 2020-05-15 17:21:31.217985 port3 out 192.168.20.5.11490 -> 192.168.20.6.1812: udp 118 2020-05-15 17:21:31.218091 port1 out 192.168.2.5.11490 -> 192.168.2.71.1812: udp 118 2020-05-15 17:21:31.219314 port3 in 192.168.20.6.1812 -> 192.168.20.5.11490: udp 20 <-- access-reject 2020-05-15 17:21:31.219519 port3 out 192.168.20.5.11490 -> 192.168.20.6.1812: udp 182 2020-05-15 17:21:31.220219 port3 in 192.168.20.6.1812 -> 192.168.20.5.11490: udp 42 2020-05-15 17:21:31.220325 port3 out 192.168.20.5.11490 -> 192.168.20.6.1812: udp 119 2020-05-15 17:21:31.220801 port3 in 192.168.20.6.1812 -> 192.168.20.5.11490: udp 20 2020-05-15 17:21:31.236009 port1 in 192.168.2.71.1812 -> 192.168.2.5.11490: udp 20 <--access-accept Access is denied by the NPS server because the user does not exist. However, access is accepted by FortiAuthenticator. The end result is the authentication is successful. # get vpn ssl monitor SSL VPN Login Users: Index User Group Auth Type Timeout From HTTP in/out HTTPS in/out 0 fackeith dualPrimaryGroup 2(1) 292 192.168.2.202 0/0 0/0 SSL VPN sessions: Index User Group Source IP Duration I/O Bytes Tunnel/Dest IP 0 fackeith dualPrimaryGroup 192.168.2.202 149 70236/4966 10.212.134.200 Case 2: Connect to the SSLVPN tunnel using FortiClient with user radkeith:# diagnose sniffer packet any 'port 1812' 4 0 l interfaces=[any] filters=[port 1812] 2020-05-15 17:26:07.335791 port1 out 192.168.2.5.17988 -> 192.168.2.71.1812: udp 118 2020-05-15 17:26:07.335911 port3 out 192.168.20.5.17988 -> 192.168.20.6.1812: udp 118 2020-05-15 17:26:07.337659 port3 in 192.168.20.6.1812 -> 192.168.20.5.17988: udp 20 <--access-accept 2020-05-15 17:26:07.337914 port3 out 192.168.20.5.17988 -> 192.168.20.6.1812: udp 182 2020-05-15 17:26:07.339451 port3 in 192.168.20.6.1812 -> 192.168.20.5.17988: udp 228 2020-05-15 17:26:08.352597 port1 in 192.168.2.71.1812 -> 192.168.2.5.17988: udp 20 <--access-reject There is a password mismatch for this user on the Secondary RADIUS server. However, even though the authentication was rejected by FortiAuthenticator, it was accepted by Windows NPS. Therefore, the end result is authentication successful. # get vpn ssl monitor SSL VPN Login Users: Index User Group Auth Type Timeout From HTTP in/out HTTPS in/out 0 radkeith dualPrimaryGroup 2(1) 290 192.168.2.202 0/0 0/0 SSL VPN sessions: Index User Group Source IP Duration I/O Bytes Tunnel/Dest IP 0 radkeith dualPrimaryGroup 192.168.2.202 142 64875/4966 10.212.134.200
