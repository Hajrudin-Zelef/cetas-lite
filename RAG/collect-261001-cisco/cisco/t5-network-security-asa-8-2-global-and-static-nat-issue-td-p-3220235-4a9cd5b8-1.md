---
id: collect-261001-cisco/cisco/t5-network-security-asa-8-2-global-and-static-nat-issue-td-p-3220235-4a9cd5b8-1
title: "t5-network-security-asa-8-2-global-and-static-nat-issue-td-p-3220235-4a9cd5b8"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-network-security-asa-8-2-global-and-static-nat-issue-td-p-3220235-4a9cd5b8.md
source_anchor: ""
source_lines: [1, 191]
sha256: 229430b8f2f2271243eac7b6e00fb773e59054fe15044c2952bb660144951640
---

# t5-network-security-asa-8-2-global-and-static-nat-issue-td-p-3220235-4a9cd5b8

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-21-2017 07:37 AM - edited 02-21-2020 06:48 AM
Im setting up a home lab and having a real time with natting. My global statement seems to override my static statements causing a drop in my inbound traffic to a plex server.
Right now in the below outputs i kinda have the static statements reversed i think just because i tried to place them on the outside interface to separate them from the global statement but im pretty sure its all wrong. Any and all advise is very much appreciated.
Below are some show and packet trace inputs, a couple of notes I currently have a permit ip any any for troubleshooting reasons and I included a couple of packet trace inputs, one that has my publick ip as the source and one that just has any public IP.
access-list OUTSIDE_access_in extended permit ip any any
access-list OUTSIDE_access_in extended permit udp any host 10.10.X.X eq 20122 log
access-list OUTSIDE_access_in extended permit tcp any host 10.10.X.X eq 5050 log
access-list OUTSIDE_access_in extended permit tcp any host 10.10.X.X eq 8989 log
access-list OUTSIDE_access_in extended permit tcp any host 10.10.X.X eq 8080 log
access-list OUTSIDE_access_in extended permit tcp any host 10.10.X.X eq 32400 log
access-list OUTSIDE_access_in extended deny ip any any
access-list LAB_access_in extended permit ip any any
access-list LAB_access_in extended deny ip any any
access-list WIRELESS_access_in extended permit ip any any
access-list WIRELESS_access_in extended deny ip any any
access-list plex_access_in extended permit ip any any
access-list plex_access_in extended deny ip any any
access-list OUTSIDE_acess_in extended permit udp any any
pager lines 24
mtu Outside 1500
mtu WIRELESS 1500
mtu LAB 1500
mtu Plex 1500
ip local pool vpn_users 10.10.X.2-10.10.X.254 mask 255.255.255.0
no failover
icmp unreachable rate-limit 1 burst-size 1
no asdm history enable
arp timeout 14400
global (Outside) 10 interface
nat (Plex) 10 10.10.X.0 255.255.255.0
static (Outside,Plex) tcp 10.10.X.X 32400 159.118.X.X 32400 netmask 255.255.255.255
static (Outside,Plex) tcp 10.10.X.X 8989 159.118.X.X 8989 netmask 255.255.255.255
static (Outside,Plex) tcp 10.10.X.X 8080 159.118.X.X 8080 netmask 255.255.255.255
static (Outside,Plex) tcp 10.10.X.X 5050 159.118.X.X 5050 netmask 255.255.255.255
static (Outside,Plex) udp 10.10.X.X 22 159.118.X.X 20122 netmask 255.255.255.255
access-group OUTSIDE_access_in in interface Outside
access-group WIRELESS_access_in in interface WIRELESS
access-group LAB_access_in in interface LAB
access-group plex_access_in in interface Plex
timeout xlate 3:00:00
NAT policies on Interface Outside:
 match tcp Outside host 159.118.X.X eq 32400 Plex any
 static translation to 10.10.X.X/32400
 translate_hits = 0, untranslate_hits = 17
 match tcp Outside host 159.118.X.X eq 8989 Plex any
 static translation to 10.10.X.X/8989
 translate_hits = 0, untranslate_hits = 0
 match tcp Outside host 159.118.X.X eq 8080 Plex any
 static translation to 10.10.X.X/8080
 translate_hits = 0, untranslate_hits = 0
 match tcp Outside host 159.118.X.X eq 5050 Plex any
 static translation to 10.10.X.X/5050
 translate_hits = 0, untranslate_hits = 0
 match udp Outside host 159.118.X.X eq 20122 Plex any
 static translation to 10.10.X.X/22
 translate_hits = 0, untranslate_hits = 0
NAT policies on Interface Plex:
 match ip Plex 10.10.X.X255.255.255.0 Outside any
 dynamic translation to pool 10 (159.118.X.X [Interface PAT])
 translate_hits = 538, untranslate_hits = 43
 match ip Plex 10.10.X.X 255.255.255.0 WIRELESS any
 dynamic translation to pool 10 (No matching global)
 translate_hits = 0, untranslate_hits = 0
 match ip Plex 10.10.X.X 255.255.255.0 LAB any
 dynamic translation to pool 10 (No matching global)
 translate_hits = 0, untranslate_hits = 0
 match ip Plex 10.10.X.X 255.255.255.0 Plex any
 dynamic translation to pool 10 (No matching global)
 translate_hits = 0, untranslate_hits = 0
Scott-ASA5510-EDGE(config)# packet-tracer input outside udp 159.118.X.X 20122 10.1$
Phase: 1
Type: ACCESS-LIST
Subtype:
Result: ALLOW
Config:
Implicit Rule
Additional Information:
 Forward Flow based lookup yields rule:
 in id=0xab90ae00, priority=1, domain=permit, deny=false
 hits=11924721, user_data=0x0, cs_id=0x0, l3_type=0x8
 src mac=0000.0000.0000, mask=0000.0000.0000
 dst mac=0000.0000.0000, mask=0100.0000.0000
Phase: 2
Type: ROUTE-LOOKUP
Subtype: input
Result: ALLOW
Config:
Additional Information:
in 10.10.X.X 255.255.255.0 Plex
Phase: 3
Type: ACCESS-LIST
Subtype:
Result: DROP
Config:
Implicit Rule
Additional Information:
 Forward Flow based lookup yields rule:
 in id=0xab928108, priority=500, domain=permit, deny=true
 hits=18, user_data=0x6, cs_id=0x0, reverse, flags=0x0, protocol=0
 src ip=159.118.X.X, mask=255.255.255.255, port=0
 dst ip=0.0.0.0, mask=0.0.0.0, port=0, dscp=0x0
Result:
input-interface: Outside
input-status: up
input-line-status: up
output-interface: Plex
output-status: up
output-line-status: up
Action: drop
Drop-reason: (acl-drop) Flow is denied by configured rule
Scott-ASA5510-EDGE(config)# packet-tracer input outside udp 1.1.1.1 20122 10.1$
Phase: 1
Type: ACCESS-LIST
Subtype:
Result: ALLOW
Config:
Implicit Rule
Additional Information:
 Forward Flow based lookup yields rule:
 in id=0xab90ae00, priority=1, domain=permit, deny=false
 hits=11946670, user_data=0x0, cs_id=0x0, l3_type=0x8
 src mac=0000.0000.0000, mask=0000.0000.0000
 dst mac=0000.0000.0000, mask=0100.0000.0000
Phase: 2
Type: ROUTE-LOOKUP
Subtype: input
Result: ALLOW
Config:
Additional Information:
in 10.10.X.X 255.255.255.0 Plex
Phase: 3
Type: ACCESS-LIST
Subtype: log
Result: ALLOW
Config:
access-group OUTSIDE_access_in in interface Outside
access-list OUTSIDE_access_in extended permit ip any any
Additional Information:
 Forward Flow based lookup yields rule:
 in id=0xacbc0a38, priority=12, domain=permit, deny=false
 hits=3838, user_data=0xa8b3fb80, cs_id=0x0, flags=0x0, protocol=0
 src ip=0.0.0.0, mask=0.0.0.0, port=0
 dst ip=0.0.0.0, mask=0.0.0.0, port=0, dscp=0x0
Phase: 4
Type: IP-OPTIONS
Subtype:
Result: ALLOW
Config:
Additional Information:
 Forward Flow based lookup yields rule:
 in id=0xab90d538, priority=0, domain=inspect-ip-options, deny=true
 hits=40865, user_data=0x0, cs_id=0x0, reverse, flags=0x0, protocol=0
 src ip=0.0.0.0, mask=0.0.0.0, port=0
 dst ip=0.0.0.0, mask=0.0.0.0, port=0, dscp=0x0
Phase: 5
Type: NAT
Subtype: rpf-check
Result: DROP
Config:
nat (Plex) 10 10.10.X.X 255.255.255.0
 match ip Plex 10.10.X.X 255.255.255.0 Outside any
 dynamic translation to pool 10 (159.118.X.X [Interface PAT])
 translate_hits = 701, untranslate_hits = 48
Additional Information:
 Forward Flow based lookup yields rule:
 out id=0xacb91d98, priority=1, domain=nat-reverse, deny=false
 hits=49, user_data=0xaba3dd70, cs_id=0x0, flags=0x0, protocol=0
 src ip=0.0.0.0, mask=0.0.0.0, port=0
 dst ip=10.10.30.0, mask=255.255.255.0, port=0, dscp=0x0
Result:
input-interface: Outside
input-status: up
input-line-status: up
output-interface: Plex
output-status: up
output-line-status: up
Action: drop
Drop-reason: (acl-drop) Flow is denied by configured rule
Solved! Go to Solution.
- Labels:
- 
						
							
		
