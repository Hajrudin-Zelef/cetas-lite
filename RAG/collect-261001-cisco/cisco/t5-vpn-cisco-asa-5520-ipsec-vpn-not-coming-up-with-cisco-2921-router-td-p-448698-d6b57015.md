---
id: collect-261001-cisco/cisco/t5-vpn-cisco-asa-5520-ipsec-vpn-not-coming-up-with-cisco-2921-router-td-p-448698-d6b57015
title: "t5-vpn-cisco-asa-5520-ipsec-vpn-not-coming-up-with-cisco-2921-router-td-p-448698-d6b57015"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-cisco/t5-vpn-cisco-asa-5520-ipsec-vpn-not-coming-up-with-cisco-2921-router-td-p-448698-d6b57015.md
source_anchor: ""
source_lines: [1, 133]
sha256: e2f6aa6e10bfc5dcba957812c3dc4bb19074bb9c41acc791e248ae27080d01ad
---

# t5-vpn-cisco-asa-5520-ipsec-vpn-not-coming-up-with-cisco-2921-router-td-p-448698-d6b57015

CISCO ASA 5520 IPSEC VPN NOT COMING UP with CISCO 2921 Router
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-15-2021 08:09 AM
I am facing the issue of creating an IPSEC VPN between CISCO ASA 5520 and CISCO 2921 ROUTER.So can any one help me in resolving this issue as i am new to CISCO ASA side. The configuration for sides are pasted below, as there is nothing on the debug outputs as neither the phase 1 is coming up but on Sophos XG firewall instaed of this CISCO ASA 5520 THE IPSEC both Phases have gone up.As PPPOE Internet Link is working fine with the IP address
Cisco ASA 5520 Configuration:-
interface GigabitEthernet0/0
nameif outside
security-level 0
pppoe client vpdn group TO_ICE
ip address 101.53.X.X 255.255.255.255 pppoe
crypto isakmp policy 10
authentication pre-share
encryption 3des
hash md5
group 2
lifetime 28800
object-group network LOCAL
network-object 192.168.16.0 255.255.255.0
object-group network REMOTE
network-object 172.17.0.0 255.255.252.0
access-list ICE_TO_DP extended permit ip object-group LOCAL object-group REMOTE
crypto ipsec transform-set TSET esp-3des esp-md5-hmac
crypto ipsec security-association lifetime seconds 28800
crypto ipsec security-association lifetime kilobytes 4608000
tunnel-group 124.29.X.X type ipsec-l2l
tunnel-group 124.29.X.X ipsec-attributes
pre-shared-key *
crypto map CMAP 10 match address ICE_TO_DP
crypto map CMAP 10 set peer 124.29.X.X
crypto map CMAP 10 set transform-set TSET
crypto map CMAP 10 set reverse-route
crypto map CMAP interface outside
crypto isakmp enable outside
CISCO 2921 ROUTER:-
crypto isakmp policy 10
encr 3des
hash md5
authentication pre-share
group 2
lifetime 28800
crypto isakmp key cis***** address 101.53.X.X
crypto ipsec transform-set tset esp-3des esp-md5-hmac
mode tunnel
crypto map VPN_Remote_Sites 60 ipsec-isakmp
description ** VPN Tunnel from DP to NEW ICE Office **
set peer 101.53.X.X
set transform-set tset
match address For-ICE
IP access list Extended For-ICE
permit ip 172.17.0.0 0.0.3.255 192.168.16.0 0.0.0.255
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2021 01:55 AM
Based on your config, most of it looks fine. I don't see that you actually applied this crypto map to appropriate interface, nor do I see routing for both peer addresses and crypto domain.
Have you tried enabling some debugs in order to understand at which phase are you failing?
On ASA side, you could enable following debugs for phase 1:
debug crypto condition peer X.X.X.X (in order to reduce debugs only to specific remote address)
debug crypto ikev1 200 (or isakmp, depending on SW version)
For phase 2, you could use:
debug crypto ipsec 200
Based on these inputs, you should be able to see at which phase are you failing. You can disable debugs with 'undebug all' command.
BR,
Milos
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2021 11:52 AM
There is nothing on the debug on the CISCO ASA end. I tried everything with debugs but no output on those as the issue seems to be on the CISCO ASA end as when we built the IPSEC between the Cisco 2921 Router and Sophos XG the IPSEC tunnel both phases are up , but without changing anything on the Cisco 2921 configuration the ISAKMP PHASE 1 and Phase 2 are not making a connection.
The most of the stuff are related to the Ikev1 and 1Kev2 but in this Cisco ASA there is not Ikev enable and only isakmp is working and as i am scractching my heads as nothing seems to be working and i dont find anything on the google.
The router End is working as we have 6 IPSECs with the router ipsec and with Sophos the same ASA WAN ips are makiing ipsec both phase1 and phase 2 connections.
For the routing part, Can you help What route i need to give on router end and on the Cisco ASA end ? As i think the Proxy ids matching in the ACL that will be called in the Transform is enough for the IPSEC to work or may be i am wrong.
Really need help for the making of IPSEC TUNNEL.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-23-2021 07:47 AM
Now the IPSEC is stuck in this MM_WAIT_MSG2 after i configured it with the ASDM, i am confirm that the issue on this CISCO ASA 5520 side as when we replaced the SOPHOS XG firewall both the IPSEC Phase 1 and Phase 2 got up and dont stuck in any of the states.Even on the Sophos XG firewall with the PPPOE Connection with the Static Public ip and without any route the IPSEC got up at first go and in the cisco ASA 5520 side even with the default route pointing towards the Gateway address the IPSEC got stuck in this state.
So i was reading that this MM_NO_STATE issue is related to the route issue and on one forum that it is due to the filtering of the UDP Port 500 towards the ISP on the outside interface of the Firewall.So what is the issue exactly i dont know ?
ICE-ASA# show crypto isakmp sa
Active SA: 1
Rekey SA: 0 (A tunnel will report 1 Active and 1 Rekey SA during rekey)
Total IKE SA: 1
1 IKE Peer: 124.29.X.X
Type : user Role : initiator
Rekey : no State : MM_WAIT_MSG2
124.29.X.X 101.53.X.X MM_NO_STATE 0 ACTIVE (deleted)
Also the Web VPN is working over the same PPPOE ISL Link so is it the ISP issue or the Configuraiton issue on the Cisco ASA for the IPSEC Site to site VPN .
ICE-ASA# show vpn-sessiondb webvpn
Session Type: WebVPN
Username : wa***** Index : 3
Public IP : 110.93.X.X
Protocol : Clientless
License : SSL VPN
Encryption : AES128 Hashing : SHA1
Bytes Tx : 133771 Bytes Rx : 25608
Group Policy : DfltGrpPolicy Tunnel Group : DefaultWEBVPNGroup
Login Time : 13:18:43 UTC Sat Oct 16 2021
Duration : 0h:00m:28s
NAC Result : Unknown
VLAN Mapping : N/A VLAN : none
ICE-ASA#
