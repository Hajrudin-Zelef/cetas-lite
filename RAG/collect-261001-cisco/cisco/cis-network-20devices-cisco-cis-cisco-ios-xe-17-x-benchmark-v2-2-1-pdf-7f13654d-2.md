---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-2
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [58, 115]
sha256: 5c9d80ed0ca1dd4e9c818e08edef0bd5febdeb2a71625b8ea3a14a330cd6fdbb
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 3 
Internal Only - General 
1.1.4 Set 'login authentication for 'line vty' (Automated) ........................................................... 24 
1.1.5 Set 'login authentication for 'ip http' (Automated) ............................................................ 26 
1.1.6 Set 'aaa accounting' to log all privileged use commands using 'commands 15' 
(Automated) .............................................................................................................................. 28 
1.1.7 Set 'aaa accounting connection' (Automated) ................................................................. 30 
1.1.8 Set 'aaa accounting exec' (Automated) ........................................................................... 32 
1.1.9 Set 'aaa accounting network' (Automated) ...................................................................... 34 
1.1.10 Set 'aaa accounting system' (Automated) ..................................................................... 36 
1.2 Access Rules .................................................................................................................................. 38 
1.2.1 Set 'privilege 1' for local users (Manual) .......................................................................... 39 
1.2.2 Set 'transport input ssh' for 'line vty' connections (Automated) ....................................... 41 
1.2.3 Set 'no exec' for 'line aux 0' (Automated) ........................................................................ 43 
1.2.4 Create 'access-list' for use with 'line vty' (Manual) .......................................................... 45 
1.2.5 Set 'access-class' for 'line vty' (Automated) ..................................................................... 47 
1.2.6 Set 'exec-timeout' to less than or equal to 10 minutes for 'line aux 0' (Automated) ........ 49 
1.2.7 Set 'exec-timeout' to less than or equal to 10 minutes 'line console 0' (Automated)  ....... 51 
1.2.8 Set 'exec-timeout' to less than or equal to 10 minutes 'line vty' (Automated).................. 53 
1.2.9 Set 'http Secure-server' limit (Automated) ....................................................................... 55 
1.2.10 Set 'exec-timeout' to less than or equal to 10 min on 'ip http' (Automated) ................... 57 
1.3 Banner Rules .................................................................................................................................. 59 
1.3.1 Set the 'banner-text' for 'banner exec' (Automated) ........................................................ 60 
1.3.2 Set the 'banner-text' for 'banner login' (Automated) ........................................................ 62 
1.3.3 Set the 'banner-text' for 'banner motd' (Automated) ........................................................ 64 
1.3.4 Set the 'banner-text' for 'webauth banner' (Automated) .................................................. 66 
1.4 Password Rules ............................................................................................................................. 68 
1.4.1 Set 'password' for 'enable secret' (Automated) ............................................................... 69 
1.4.2 Enable 'service password-encryption' (Automated) ......................................................... 71 
1.4.3 Set 'username secret' for all local users (Automated) ..................................................... 73 
1.5 SNMP Rules .................................................................................................................................... 75 
1.5.1 Set 'no snmp-server' to disable SNMP when unused (Manual) ...................................... 76 
1.5.2 Unset 'private' for 'snmp-server community' (Automated) ............................................... 78 
1.5.3 Unset 'public' for 'snmp-server community' (Automated) ................................................. 80 
1.5.4 Do not set 'RW' for any 'snmp-server community' (Manual) ............................................ 82 
1.5.5 Set the ACL for each 'snmp-server community' (Manual) ............................................... 84 
1.5.6 Create an 'access-list' for use with SNMP (Manual) ........................................................ 86 
1.5.7 Set 'snmp-server host' when using SNMP (Automated) .................................................. 88 
1.5.8 Set 'snmp-server enable traps snmp' (Automated) ......................................................... 90 
1.5.9 Set 'priv' for each 'snmp-server group' using SNMPv3 (Automated) ............................... 92 
1.5.10 Require 'aes 128' as minimum for 'snmp-server user' when using SNMPv3 (Manual) . 94 
2 Control Plane ................................ ................................ ................................ .....................95 
2.1 Global Service Rules ..................................................................................................................... 95 
2.1.1 Setup SSH ................................................................................................................................. 95 
2.1.1.1 Configure Prerequisites for the SSH Service .................................................................. 96 
2.1.1.1.1 Set the 'hostname' (Automated) ................................................................................ 97 
2.1.1.1.2 Set the 'ip domain-name' (Automated) ...................................................................... 99 
2.1.1.1.3 Set 'modulus' to greater than or equal to 2048 for 'crypto key generate rsa' (Manual)
 ................................................................................................................................................ 101 
2.1.1.1.4 Set 'seconds' for 'ip ssh timeout' for 60 seconds or less (Automated)  .................... 103 
2.1.1.1.5 Set maximum value for 'ip ssh authentication-retries' (Automated) ........................ 105 
2.1.1.2 Set version 2 for 'ip ssh version' (Manual) .................................................................. 107 
2.1.2 Set 'no cdp run' (Manual) ............................................................................................... 109 
2.1.3 Set 'no ip bootp server' (Manual) ................................................................................... 111 
2.1.4 Set 'no service dhcp' (Automated) ................................................................................. 113 
2.1.5 Set 'service tcp-keepalives-in' (Automated) ................................................................... 115 
2.1.6 Set 'service tcp-keepalives-out' (Automated) ................................................................ 117

