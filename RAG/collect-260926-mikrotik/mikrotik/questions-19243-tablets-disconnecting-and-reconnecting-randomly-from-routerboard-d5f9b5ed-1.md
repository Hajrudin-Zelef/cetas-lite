---
id: collect-260926-mikrotik/mikrotik/questions-19243-tablets-disconnecting-and-reconnecting-randomly-from-routerboard-d5f9b5ed-1
title: "questions-19243-tablets-disconnecting-and-reconnecting-randomly-from-routerboard-d5f9b5ed"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-19243-tablets-disconnecting-and-reconnecting-randomly-from-routerboard-d5f9b5ed.md
source_anchor: ""
source_lines: [1, 74]
sha256: 2462a92a9ba94b3c2d5a66268f99696bade13aa76ea8aa5491f4fb801547acc7
---

# questions-19243-tablets-disconnecting-and-reconnecting-randomly-from-routerboard-d5f9b5ed

We have a product that considers installing on a client’s room a Mikrotik RouterBoard 951G 2HnD hosting a wifi network with WPA2 PSK security connected to our client’s network and an Android tablet connected to the Mikrotik’s wifi network. This solution has been installed on dozens of our clients’ sites and we have never had any issues before.
Two weeks ago the all the rooms (four) inside one of our client’s building started to have wifi problems. After working great for several hours, each tablet would disconnect (¿or be dropped?) from the RouterBoard 951G 2HnD’s wifi network and then immediately connect. After the first disconnection (or being dropped) and reconnection the tablet would keep disconnecting (or being dropped) and reconnecting from the wifi for several minutes, which would end in a final disconnection and then no more reconnection intents.
The solution worked great for months but after we changed the tablet this started happening. At the beginning we thought this was because we had changed the tablet but after trying with another Android tablets and an Android cell phone even, the problem would still occur. We also replaced the RouterBoard 951G 2HnD with one that was working great at another office and it also started to show this unusual behaviour.
Finally, we took one of the four room’s Android tablet and RouterBoard 951G 2HnD, and installed them on one of our laboratories without modifying any configuration on both of them. After almost 24 hours the problem has not replicated.
I thought activating the Wireless Debug Logs would us show more relevant information and help us focus our investigation on the right direction but it didn’t. Next are the log records:
jan/01/1970 20:03:17 interface,info ether1_Conexion_Hacia_LAN link up (speed 100M, full duplex) 
jan/01/1970 20:03:17 interface,info ether2_Conexion_Hacia_terminal_VC link up (speed 100M, full duplex) 
jan/01/1970 20:04:00 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jan/01/1970 20:04:01 dhcp,info dhcp1_My_Companys_Name assigned 192.168.89.2 to 0C:B3:19:29:08:13 
jan/01/1970 20:04:09 system,error,critical router was rebooted without proper shutdown 
jan/01/1970 20:04:20 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, received deauth: sending station leaving (3) 
jan/01/1970 20:04:29 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jan/01/1970 20:07:28 interface,info ether4_Conexion_Hacia_dispositivos_internos link up (speed 100M, full duplex) 
jan/01/1970 20:07:41 system,info,account user admin logged in via winbox 
jan/01/1970 20:17:48 system,info,account user admin logged in via winbox 
jan/01/1970 20:17:50 system,info,account user admin logged out via winbox 
jan/01/1970 20:19:01 system,info,account user admin logged in via local 
jan/01/1970 20:19:28 system,info nat rule added by admin 
jan/01/1970 20:19:31 system,info nat rule added by admin 
jan/01/1970 20:19:31 system,info nat rule added by admin 
jan/01/1970 20:19:31 system,info nat rule added by admin 
jan/01/1970 20:22:51 interface,info ether4_Conexion_Hacia_dispositivos_internos link down 
jan/01/1970 20:22:52 system,info,account user admin logged out via winbox 
jan/01/1970 20:22:52 system,info,account user admin logged out via local 
jan/01/1970 20:33:20 interface,info ether2_Conexion_Hacia_terminal_VC link down 
jan/01/1970 20:33:23 interface,info ether2_Conexion_Hacia_terminal_VC link up (speed 100M, full duplex) 
jan/01/1970 23:59:08 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: reassociating 
jan/01/1970 23:59:08 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, ok 
jan/01/1970 23:59:08 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jan/02/1970 00:02:09 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: reassociating 
jan/02/1970 00:02:09 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, ok 
jan/02/1970 00:02:09 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jan/02/1970 00:05:29 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, received deauth: class 3 frame received (7) 
jan/04/1970 15:20:45 system,info,account user admin logged in from 10.252.163.33 via winbox 
jan/04/1970 15:37:52 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jan/04/1970 23:03:37 system,info,account user admin logged in from 10.252.163.33 via winbox 
jan/04/1970 23:03:48 system,info log rule added by admin 
jan/04/1970 23:04:07 system,info,account user admin logged out from 10.252.163.33 via winbox 
jan/04/1970 23:07:23 system,info SNTP client configuration changed by admin 
jun/15 15:53:43 system,info SNTP client configuration changed by admin 
jun/15 18:09:25 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, received deauth: class 3 frame received (7) 
jun/15 18:09:41 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
jun/15 18:09:41 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 in local ACL, accept 
jun/15 18:09:41 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jun/15 18:15:29 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
jun/15 18:15:29 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: reassociating 
jun/15 18:15:29 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, ok 
jun/15 18:15:29 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 in local ACL, accept 
jun/15 18:15:29 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jun/15 18:17:11 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, received deauth: class 3 frame received (7) 
jun/15 18:17:15 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
jun/15 18:17:15 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 in local ACL, accept 
jun/15 18:17:15 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jun/15 18:17:20 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, received deauth: class 3 frame received (7) 
jun/15 18:17:20 wireless,info wlan1_My_Companys_Name: data from unknown device 0C:B3:19:29:08:13, sent deauth 
jun/15 18:17:20 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
jun/15 18:17:20 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 in local ACL, accept 
jun/15 18:17:20 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jun/15 18:23:12 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
jun/15 18:23:12 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: reassociating 
jun/15 18:23:12 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, ok 
jun/15 18:23:12 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 in local ACL, accept 
jun/15 18:23:12 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jun/15 18:23:37 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
jun/15 18:23:37 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: reassociating 
jun/15 18:23:37 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, ok 
jun/15 18:23:37 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 in local ACL, accept 
jun/15 18:23:37 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jun/15 18:25:42 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
jun/15 18:25:42 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: reassociating 
jun/15 18:25:42 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, ok 
jun/15 18:25:42 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 in local ACL, accept 
jun/15 18:25:42 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jun/15 18:27:00 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
