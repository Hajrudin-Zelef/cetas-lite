---
id: collect-261001-cisco/cisco/taterbrown-cisco-secure-config-raw-3b6d38cac26f112286dfafc1abd1bf3fba229d1e-cisc-f83b173f-1
title: "Cisco Auditor Tool"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2021-07"]
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/taterbrown-cisco-secure-config-raw-3b6d38cac26f112286dfafc1abd1bf3fba229d1e-cisc-f83b173f.md
source_anchor: ""
source_lines: [1, 206]
sha256: 6421e4f13051817a6a06c62f5510b7da2f29cbad8607ca3480f028b775c8d9e2
---

# Cisco Auditor Tool

Cisco IOS Security Hardening & Best Practices 
 
 
Date: 8 July 2021

Table of Contents 
Best Practices & Tools ................................ ................................ ................................ ............... 3 
Hardening ................................ ................................ ................................ ................................ .. 3 
Create Local User Admin Account ................................ ................................ .......................... 3 
Configure Management IP ................................ ................................ ................................ ...... 3 
Restrict and Secure Remote Mgmt Access ................................ ................................ ............ 4 
Restrict Console Access ................................ ................................ ................................ ......... 4 
Configure SSH Options ................................ ................................ ................................ .......... 5 
Enable Secure Login Checking ................................ ................................ ..............................  5 
Enable Logging ................................ ................................ ................................ ...................... 5 
Enable Configuration Change Notification & Logging ................................ ............................. 6 
Disable Log to Console or Monitor Sessions ................................ ................................ .......... 6 
Enable NTP Server................................ ................................ ................................ ................. 6 
NTP Authentication................................ ................................ ................................ ................. 7 
Restrict SNMP Access ................................ ................................ ................................ ........... 7 
Disable Unused Services................................ ................................ ................................ ........ 7 
Enable Login Banner ................................ ................................ ................................ .............. 8 
Enable Keepalives TCP Sessions ................................ ................................ .......................... 8 
Enable Memory & CPU Threshold Notifications ................................ ................................ ...... 8 
Enable Secure Copy & IOS Software Resilient ................................ ................................ ....... 9 
Disable Unused Ports & Apply Port Security ................................ ................................ .......... 9 
Secure STP operation ................................ ................................ ................................ ...........10 
Prevent VLAN Hopping ................................ ................................ ................................ .........10 
OSPF Authentication ................................ ................................ ................................ .............10 
Bogon Address ACL ................................ ................................ ................................ ..............11 
RADIUS Authentication ................................ ................................ ................................ .........12 
Dynamic Trunking Protocol (DTP) ................................ ................................ .........................13

Best Practices & Tools 
 
# Cisco Auditor Tool 
https://github.com/cisco-config-analysis-tool/ccat 
 
# How to use CCAT 
https://www.blackhillsinfosec.com/how-to-use-ccat-an-analysis-tool-for-cisco-configuration-files/ 
 
# Cisco Security Best Practices 
https://www.cisco.com/c/en/us/support/docs/ip/access-lists/13608-21.html 
 
Hardening 
 
Create Local User Admin Account 
 
username netadmin privilege 15 secret 1111  
enable secret 2222  
service password-encryption  
aaa new-model  
aaa authentication login enable 
aaa local authentication attempts max-fail 10 
aaa session-id common 
 
Configure Management IP 
 
# Loopback interfaces are always up so use loopback interface for SSH remote mgmt access 
int lo0 
     ip add 10.10.100.250 255.255.255.255

Restrict and Secure Remote Mgmt Access 
 
ip access-list extended SSHAccess 
 permit tcp host 192.168.20.48 any eq 22 log 
 permit tcp host 192.168.20.2 any eq 22 log 
 deny tcp any any log 
! 
line vty 0 4 
 access-class SSHAccess in 
 transport input ssh 
 exec-timeout 15 
 
OR 
 
ip access-list standard 1 
 remark SSH ACCESS   
permit x.x.x.x 
   permit x.x.x.x 
! 
line vty 0 4 
 access-class 1 in 
 transport input ssh 
 exec-timeout 15 0 
 
Restrict Console Access 
 
line con 0  
   exec-timeout 15 
   no privilege level 15

Configure SSH Options 
 
ip domain-name 1.com  
crypto key generate rsa modulus 2048  
ip ssh version 2  
ip ssh time-out 30  
ip ssh logging events  
ip ssh maxstartups 10  
ip ssh authentication-retries 5 
ip ssh server algorithm encryption aes256-ctr  (sets aes256-ctr as the only SSH cipher to be used) 
ip ssh server algorithm mac hmac-sha1  (sets hmac-sha1 as the only SSH integrity cipher to be used) 
 
 Enable Secure Login Checking 
 
# Helps prevent dictionary attack/brute force 
login block-for 300 attempts 5 within 120 
login delay 2 
login on-failure log 
login on-success log 
 
Enable Logging 
 
logging buffered 16000 informational  
logging 10.10.10.5  (note: syslog server IP) 
logging source-interface Loopback 0 
service timestamps debug datetime msec localtime show-timezone 
service timestamps log datetime msec localtime show-timezone

Enable Configuration Change Notification & Logging 
#Enables configuration management and history of changes on the device 
 
archive 
    log config 
    logging enable 
    logging size 200 
    hidekeys 
    notify syslog 
 
sh archive log config all 
 (OUTPUT) 
 idx sess user@line Logged command 
 1 1 console@console |access-list 199 permit icmp host 10.10.10.10 host 20.20.20.20 
 2 1 console@console |crypto map NiStTeSt1 10 ipsec-manual 
 3 1 console@console |match address 199 
 4 1 console@console |set peer 20.20.20.20 
 5 1 console@console |exit 
 6 1 console@console |no access-list 199 
 7 1 console@console |no crypto map NiStTeSt1 
 8 2 netadmin@console |crypto key generate rsa modulus ***** 
 9 0 netadmin@vty0 |!exec: enable 
 
Disable Log to Console or Monitor Sessions 
 
no logging console 
no logging monitor 
 
Enable NTP Server 
 
clock timezone CST -6 0

clock summer-time CDT recurring 
ntp server x.x.x.x 
 
NTP Authentication 
 
#On the switch/router that will be the master NTP server 
ntp master 3 
ntp authenticate 
ntp authentication-key 1 md5 <password> 
 
#On the client switches/routers to receive NTP from the NTP server 
ntp server x.x.x.x  key 1 
ntp authenticate 
ntp authentication-key 1 md5 <password> 
ntp trusted-key 1 
 
Restrict SNMP Access 
 
ip access-list standard ACL-SNMP 
    permit 10.10.100.6 
    deny any log 
snmp-server community T@s9aMon RO ACL-SNMP 
snmp-server location AL 
snmp-server contact tater@1.com 
 
Disable Unused Services 
 
no ip http server 
no ip http secure-server 
no service dhcp

no cdp run 
no lldp run global 
no ip bootp server 
no ip domain-lookup  
no ip source-route 
 
Enable Login Banner 
#Change this banner message to reflect your organization’s preferred warning banner 
 
banner login #  
 UNAUTHORIZED ACCESS TO THIS DEVICE IS PROHIBITED! You must have explicit 
permission to access or configure this system.  
 All activities performed on this system may be logged, and violations of this policy may result in 
disciplinary action and may be reported to law enforcement.  
 Use of this system shall constitute consent to monitoring. 
 #  
banner motd #  
 AUTHORIZED ACCESS ONLY! If you are not an authorized user, disconnect IMMEDIATELY! 
All connections are monitored and recorded. 
 # 
 
Enable Keepalives TCP Sessions 
 
