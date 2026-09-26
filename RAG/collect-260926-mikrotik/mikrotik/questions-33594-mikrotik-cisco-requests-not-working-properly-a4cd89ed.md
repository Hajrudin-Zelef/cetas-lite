---
id: collect-260926-mikrotik/mikrotik/questions-33594-mikrotik-cisco-requests-not-working-properly-a4cd89ed
title: "questions-33594-mikrotik-cisco-requests-not-working-properly-a4cd89ed"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "voice"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-33594-mikrotik-cisco-requests-not-working-properly-a4cd89ed.md
source_anchor: ""
source_lines: [1, 98]
sha256: d2715714bb3b6412e83e4fa14b9b8c5ea192775607de5f9a49ad4c73433a448d
---

# questions-33594-mikrotik-cisco-requests-not-working-properly-a4cd89ed

We've got a problem with our WiFi company network. We bought a Cisco Switch and Wireless LAN Controller (Cisco 2500 WLC) plus eight Access Point (Aironet 2700) to offer clients seamless roaming between APs. These APs work like charm (I think it's just because they are from the same vendor), but our network was at first stages built around Mikrotik devices and now we are heading several problems.
Short about our topology: We want two WiFi networks (two SSIDs) so we made two VLANs (20,30). The traffic from Access Points is packed inside one agregated management VLAN (10). Because of agregated link the WLAN Controller is connected to network using one LAN connection. At this point everything works and everything blinks :) (and inside this Cisco "subnetwork" I am able to ping these Cisco devices each other).
Switch:
version 15.0
no service pad
service timestamps debug datetime msec
service timestamps log datetime msec
no service password-encryption
service sequence-numbers
!
hostname sw-01
!
boot-start-marker
boot-end-marker
!
!
aaa new-model
!
!
aaa authentication login default group radius local
aaa authorization console
aaa authorization exec default group radius local
aaa accounting exec default start-stop group radius
aaa accounting connection default start-stop group radius
aaa accounting system default start-stop group radius
!
!
aaa session-id common
switch 1 provision ws-c2960s-48ts-l
!
ip dhcp pool 172.18.106.0
 description Pool for first WiFi network
 network 172.18.106.0 255.255.254.0
 default-router 172.18.106.19
 dns-server 8.8.8.8
!
ip dhcp pool 172.18.113.0
 description Pool for second WiFi network
 network 172.18.113.0 255.255.255.0
 default-router 172.18.113.19
 dns-server 8.8.8.8
!
spanning-tree mode pvst
spanning-tree extend system-id
no spanning-tree vlan 1-4094
!
!
vlan internal allocation policy ascending
!
interface GigabitEthernet1/0/1
 description Connection WLAN Controller switch
 switchport mode trunk
!
interface GigabitEthernet1/0/4
 description AP connection to WLAN Controller 
 switchport access vlan 10
 switchport mode access
 switchport voice vlan 10
 spanning-tree portfast
!
interface GigabitEthernet1/0/13
 description Connection WLAN Controller and Mikrotik 
 switchport mode trunk
!
interface Vlan20
 ip address 172.18.106.10 255.255.254.0
!
interface Vlan30
 ip address 172.18.113.10 255.255.255.0
!
ip default-gateway 172.18.105.19
ip http server
ip http secure-server
!
logging history informational
logging facility local6
logging host XX.XX.XX.XX
!
snmp-server community public RO
!
ntp logging
ntp server xx.xx.xx.xx
end
BUT:
- first problem is the DHCP server. When it is configured on Cisco switch all WiFi clients get their IP with no problems. But when it is configured on Mikrotik the WiFi client gets the IP but when looking into Mikrotik ARP tables they have the same MAC address. The MAC address of WLAN Controller.
- when I try to ping Mikrotik router from Cisco switch it works. When I try to ping Cisco switch from Mikrotik it also works. But when I try to ping Mikrotik from WiFi client the client sends ARP request but is unable to get any reply from Mikrotik at all (all clients have IP address and default gateway set properly). And when I try to ping client from Mikrotik router it also doesn't work but after 2 or 3 minutes it starts work!
This is madness and we had to misconfigure something or some sort of service/protocol does not works properly between Mikrotik and Cisco. Have you ever seen something like this.
Thank you for your help!
PARTIAL SOLUTION:
Well, we found that we had several problems under one roof. And we overcame it by using brute force techniques (very time-consuming...)
- all AP should be connected through Trunk BUT with a native VLAN configured. In our case it is trunk with native VLAN 10.
- we changed ARP cache timeout on Cisco Switch to 20 sec.
- we made DHCP server on Mikrotik BUT we also had to configure Cisco Switch as DHCP Relay Agent.
- the most important thing is that we change AP mode to flexconnect technology, so all APs are autonomous but we lost advanced roaming feature provided by Wireless LAN Controller. Now, all WiFi clients manage their connection, itself. This is very buggy because of bad implementation of roaming in today's WiFi chips.
- we flashed the WLC to newest version...
- we turned off VTP prunning on Cisco switch. Maybe this was the main cause of that undefined behavior of ping.
So simple...
172.18.106.19and172.18.113.19? Why does the switch have a different gateway,172.18.105.19, but there is no VLAN with that network?
