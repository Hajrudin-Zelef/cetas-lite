---
id: collect-261001-cisco/cisco/taterbrown-cisco-secure-config-raw-3b6d38cac26f112286dfafc1abd1bf3fba229d1e-cisc-f83b173f-2
title: "Cisco Auditor Tool"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/taterbrown-cisco-secure-config-raw-3b6d38cac26f112286dfafc1abd1bf3fba229d1e-cisc-f83b173f.md
source_anchor: ""
source_lines: [207, 343]
sha256: 4917ad9a8842cea91b6395cff88f1c8ee444e230669de2edaf81c9e58d55ecec
---

# Cisco Auditor Tool

service tcp-keepalives-in 
service tcp-keepalives-out 
 
Enable Memory & CPU Threshold Notifications 
 
memory free low-watermark processor 204800 
memory free low-watermark io 204800

memory reserve critical 20480 
 
process cpu threshold type total rising 80 interval 60 falling 70 interval 60 
process cpu statistics limit entry-percentage 80 size 60 
 
memory reserve console 4096 
 
exception memory ignore overflow io 
exception memory ignore overflow processor 
 
exception crashinfo maximum files 32 
 
Enable Secure Copy & IOS Software Resilient 
 
ip scp server enable 
copy scp://usernam@10.10.100.250/home/tater/file.txt flash: 
 
configuration mode exclusive auto 
 
secure boot-image 
secure boot-config 
 
Disable Unused Ports & Apply Port Security 
# global command to recovery from port security violation 
errdisable recovery cause psecure-violation 
 
int range fa0/1 - 48 
    switchport port-security maximum 2 
    switchport port-security mac-address sticky 
    switchport port-security violation shutdown

switchport port-security port-security aging time 15 
    shut 
 
Secure STP operation 
#Enable BPDU guard on access ports 
 
spanning-tree bpduguard enable 
 
Prevent VLAN Hopping 
# OPTION 1: Change the Native VLAN and then Prune the Native VLAN (CDP, PAgp and DTP 
will still function) 
 
switchport trunk encapsulation dot1q 
switchport trunk native vlan 800 
switchport trunk allowed vlan remove 800 
switchport mode trunk 
 
# OPTION 2: Force the switch to tag the native VLAN (global cmd, must be done on all 
switches) 
 
vlan dot1q tag native 
 
OSPF Authentication 
#Interface that OSPF is enabled on 
interface FastEthernet0/1   
description to KeyWest   
ip address 200.120.45.253 255.255.255.252   
no ip directed-broadcast   
bandwidth 512   
ip ospf authentication message-digest   
ip ospf message-digest-key 1 md5 <password>

#The other router interface that peers with Key West 
interface FastEthernet0/1   
description to Miami  
ip address 200.120.45.254 255.255.255.252   
no ip directed-broadcast   
bandwidth 512   
ip ospf authentication message-digest   
ip ospf message-digest-key 1 md5 <password> 
 
Bogon Address ACL 
#Denies invalid IP address space from being routed by core switch 
#Apply ACL to egress port of the core switch connecting to the perimeter router/firewall 
 
ip access-list extended 102 
 remark BOGON ADDRESSES 
deny ip 0.0.0.0 0.255.255.255 any  
deny ip 10.0.0.0 0.255.255.255 any 
deny ip 100.64.0.0 0.63.255.255 any 
deny ip 127.0.0.0 0.255.255.255 any 
deny ip 169.254.0.0 0.0.255.255 any 
deny ip 172.16.0.0 0.15.255.255 any 
deny ip 192.0.0.0 0.0.0.255 any 
deny ip 192.0.2.0 0.0.0.255 any 
deny ip 192.168.0.0 0.0.255.255 any 
deny ip 198.18.0.0 0.1.255.255 any 
deny ip 198.51.100.0 0.0.0.255 any 
deny ip 203.0.113.0 0.0.0.255 any 
deny ip 224.0.0.0 31.255.255.255 any 
permit ip any any

RADIUS Authentication 
#Configure AAA lines 
aaa authentication login default local enable 
aaa authentication login aaa_login group radius local 
aaa authorization exec default local 
aaa authorization network default group radius local 
 
#Create Local Device User Accounts 
#Accounts must match how the user account is created within LDAP (e.g., LDAP account 
tater.tot must be tater.tot on the device) 
username mr.robot privilege 15 secret 9  
username tater.tot privilege 15 secret 9  
 
#Setup Radius Server 
radius server DUO-Radius 
 address ipv4 x.x.x.x auth-port 1812 acct-port 1813 
 non-standard 
 key 7  
 
#Configure VTY Lines 
line vty 0 4 
 access-class 101 in 
 login authentication aaa_login 
 length 0 
 transport input ssh 
 transport output ssh

Dynamic Trunking Protocol (DTP) 
#Disable DTP on access ports 
#Add “switchport nonegotiate” to disable DTP on the interface 
# Haven’t found a global command to disable yet 
 
int range gig 0/1 - 48 
   switchport mode access 
   switchport nonegotiate  
   spanning-tree portfast edge
