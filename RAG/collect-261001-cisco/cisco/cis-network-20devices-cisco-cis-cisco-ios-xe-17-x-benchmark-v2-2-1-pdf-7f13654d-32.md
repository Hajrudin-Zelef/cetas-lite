---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-32
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [5602, 5780]
sha256: 268d18cb38309f4de992d3e332f5f70c92514c3d1699610a89f8691e8e082fee
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 209 
Internal Only - General 
Recommendation Set 
Correctly 
Yes No 
2.3.2 Set 'ip address' for 'ntp server'   
2.4.1 Create a single 'interface loopback'   
2.4.2 Set AAA 'source-interface'   
2.4.3 Set 'ntp source' to Loopback Interface   
2.4.4 Set 'ip tftp source-interface' to the Loopback Interface   
3.1.1 Set 'no ip source-route'   
3.1.2 Set 'no ip proxy-arp'   
3.1.3 Set 'no interface tunnel'   
3.1.4 Set 'ip verify unicast source reachable-via'   
3.2.1 Set 'ip access-list extended' to Forbid Private Source 
Addresses from External Networks   
3.2.2 Set inbound 'ip access-group' on the External Interface   
3.3.1.1 Set 'key chain'   
3.3.1.2 Set 'key'   
3.3.1.3 Set 'key-string'   
3.3.1.4 Set 'address-family ipv4 autonomous-system'   
3.3.1.5 Set 'af-interface default'   
3.3.1.6 Set 'authentication key-chain'   
3.3.1.7 Set 'authentication mode md5'   
3.3.1.8 Set 'ip authentication key-chain eigrp'   
3.3.1.9 Set 'ip authentication mode eigrp'   
3.3.2.1 Set 'authentication message-digest' for OSPF area   
3.3.2.2 Set 'ip ospf message-digest-key md5'   
3.3.3.1 Set 'neighbor password'  

Page 210 
Internal Only - General 
Appendix: CIS Controls v7 IG 3 Mapped 
Recommendations 
Recommendation Set 
Correctly 
Yes No 
1.1.1 Enable 'aaa new-model'   
1.1.2 Enable 'aaa authentication login'   
1.1.3 Enable 'aaa authentication enable default'   
1.1.4 Set 'login authentication for 'line vty'   
1.1.5 Set 'login authentication for 'ip http'   
1.1.6 Set 'aaa accounting' to log all privileged use commands 
using 'commands 15'   
1.1.7 Set 'aaa accounting connection'   
1.1.8 Set 'aaa accounting exec'   
1.1.9 Set 'aaa accounting network'   
1.1.10 Set 'aaa accounting system'   
1.2.2 Set 'transport input ssh' for 'line vty' connections   
1.2.3 Set 'no exec' for 'line aux 0'   
1.2.4 Create 'access-list' for use with 'line vty'   
1.2.5 Set 'access-class' for 'line vty'   
1.2.6 Set 'exec-timeout' to less than or equal to 10 minutes for 
'line aux 0'   
1.2.7 Set 'exec-timeout' to less than or equal to 10 minutes 
'line console 0'   
1.2.8 Set 'exec-timeout' to less than or equal to 10 minutes 
'line vty'   
1.3.1 Set the 'banner-text' for 'banner exec'   
1.3.2 Set the 'banner-text' for 'banner login'   
1.3.3 Set the 'banner-text' for 'banner motd'   
1.3.4 Set the 'banner-text' for 'webauth banner'   
1.4.1 Set 'password' for 'enable secret'   
1.4.2 Enable 'service password-encryption'   
1.4.3 Set 'username secret' for all local users  

Page 211 
Internal Only - General 
Recommendation Set 
Correctly 
Yes No 
1.5.1 Set 'no snmp-server' to disable SNMP when unused   
1.5.2 Unset 'private' for 'snmp-server community'   
1.5.3 Unset 'public' for 'snmp-server community'   
1.5.4 Do not set 'RW' for any 'snmp-server community'   
1.5.5 Set the ACL for each 'snmp-server community'   
1.5.6 Create an 'access-list' for use with SNMP   
1.5.7 Set 'snmp-server host' when using SNMP   
1.5.8 Set 'snmp-server enable traps snmp'   
1.5.9 Set 'priv' for each 'snmp-server group' using SNMPv3   
1.5.10 Require 'aes 128' as minimum for 'snmp-server user' 
when using SNMPv3   
2.1.1.1.3 Set 'modulus' to greater than or equal to 2048 for 'crypto 
key generate rsa'   
2.1.1.1.4 Set 'seconds' for 'ip ssh timeout' for 60 seconds or less   
2.1.2 Set 'no cdp run'   
2.1.3 Set 'no ip bootp server'   
2.1.4 Set 'no service dhcp'   
2.1.5 Set 'service tcp-keepalives-in'   
2.1.6 Set 'service tcp-keepalives-out'   
2.1.7 Set 'no service pad'   
2.2.1 Set 'logging enable'   
2.2.2 Set 'buffer size' for 'logging buffered'   
2.2.3 Set 'logging console critical'   
2.2.4 Set IP address for 'logging host'   
2.2.5 Set 'logging trap informational'   
2.2.6 Set 'service timestamps debug datetime'   
2.2.7 Set 'logging source interface'   
2.2.8 Set 'login success/failure logging'   
2.3.1.1 Set 'ntp authenticate'   
2.3.1.2 Set 'ntp authentication-key'   
2.3.1.3 Set the 'ntp trusted-key'   
2.3.1.4 Set 'key' for each 'ntp server'  

Page 212 
Internal Only - General 
Recommendation Set 
Correctly 
Yes No 
2.3.2 Set 'ip address' for 'ntp server'   
2.4.1 Create a single 'interface loopback'   
2.4.2 Set AAA 'source-interface'   
2.4.3 Set 'ntp source' to Loopback Interface   
2.4.4 Set 'ip tftp source-interface' to the Loopback Interface   
3.1.1 Set 'no ip source-route'   
3.1.2 Set 'no ip proxy-arp'   
3.1.3 Set 'no interface tunnel'   
3.1.4 Set 'ip verify unicast source reachable-via'   
3.2.1 Set 'ip access-list extended' to Forbid Private Source 
Addresses from External Networks   
3.2.2 Set inbound 'ip access-group' on the External Interface   
3.3.1.1 Set 'key chain'   
3.3.1.2 Set 'key'   
3.3.1.3 Set 'key-string'   
3.3.1.4 Set 'address-family ipv4 autonomous-system'   
3.3.1.5 Set 'af-interface default'   
3.3.1.6 Set 'authentication key-chain'   
3.3.1.7 Set 'authentication mode md5'   
3.3.1.8 Set 'ip authentication key-chain eigrp'   
3.3.1.9 Set 'ip authentication mode eigrp'   
3.3.2.1 Set 'authentication message-digest' for OSPF area   
3.3.2.2 Set 'ip ospf message-digest-key md5'   
3.3.3.1 Set 'neighbor password'  

Page 213 
Internal Only - General 
Appendix: CIS Controls v7 Unmapped 
Recommendations 
Recommendation Set 
Correctly 
Yes No 
 No unmapped recommendations to CIS Controls v7  

Page 214 
Internal Only - General 
Appendix: CIS Controls v8 IG 1 Mapped 
Recommendations 
Recommendation Set 
Correctly 
Yes No 
1.1.6 Set 'aaa accounting' to log all privileged use commands 
using 'commands 15'   
1.1.8 Set 'aaa accounting exec'   
1.1.9 Set 'aaa accounting network'   
1.1.10 Set 'aaa accounting system'   
1.2.2 Set 'transport input ssh' for 'line vty' connections   
1.2.3 Set 'no exec' for 'line aux 0'   
1.2.6 Set 'exec-timeout' to less than or equal to 10 minutes for 
'line aux 0'   
1.2.7 Set 'exec-timeout' to less than or equal to 10 minutes 
'line console 0'   
1.2.8 Set 'exec-timeout' to less than or equal to 10 minutes 
'line vty'   
1.3.1 Set the 'banner-text' for 'banner exec'   
1.3.2 Set the 'banner-text' for 'banner login'   
1.3.3 Set the 'banner-text' for 'banner motd'   
1.3.4 Set the 'banner-text' for 'webauth banner'   
1.4.1 Set 'password' for 'enable secret'   
1.5.1 Set 'no snmp-server' to disable SNMP when unused   
1.5.2 Unset 'private' for 'snmp-server community'   
1.5.3 Unset 'public' for 'snmp-server community'   
1.5.4 Do not set 'RW' for any 'snmp-server community'   
1.5.9 Set 'priv' for each 'snmp-server group' using SNMPv3   
1.5.10 Require 'aes 128' as minimum for 'snmp-server user' 
when using SNMPv3   
2.1.1.1.4 Set 'seconds' for 'ip ssh timeout' for 60 seconds or less   
2.1.2 Set 'no cdp run'   
2.1.3 Set 'no ip bootp server'  

