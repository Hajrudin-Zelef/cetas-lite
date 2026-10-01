---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-34
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2025-05-22"]
keywords: ["benchmark", "benchmarks"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [5946, 6055]
sha256: 10dccaad0d45613d656c9e17604a85abe5208937e71ec8dbe1585ffbd6371471
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 220 
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

Page 221 
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

Page 222 
Internal Only - General 
Appendix: CIS Controls v8 Unmapped 
Recommendations 
Recommendation Set 
Correctly 
Yes No 
 No unmapped recommendations to CIS Controls v8  

Page 223 
Internal Only - General 
Appendix: Change History 
Date Version Changes for this version 
Oct 29, 2024 2.2.1 Benchmark Cisco IOS XE 
17 - Assessment failing (ip 
ssh authenticatio-retries) 
due to regular expression 
not matching configuration 
(Ticket 22102) 
May 22, 2025 2.2.1 Set 'no ip bootp server' 
failing validation (Ticket 
21935) 
May 22, 2025 2.2.1 "ip ssh authentication-
retries" failed validation 
(Ticket 21934) 
May 22, 2025 2.2.1 Set 'seconds' for 'ip ssh 
timeout' for 60 seconds or 
less false positive (Ticket 
21918) 
May 22, 2025 2.2.1 Set 'exec-timeout' to less 
than or equal to 10 min on 
'ip http' false positive (Ticket 
21917) 
May 22, 2025 2.2.1 "transport input none" can 
be entered in configuration 
but doesn't appear in config 
so can't be validated (Ticket 
21915) 
May 22, 2025 2.2.1 SCAP for routers and 
switches supporting the 
latest benchmarks. (Ticket 
21575)
