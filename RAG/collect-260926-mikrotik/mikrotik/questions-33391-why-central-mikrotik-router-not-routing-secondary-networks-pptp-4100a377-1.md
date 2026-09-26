---
id: collect-260926-mikrotik/mikrotik/questions-33391-why-central-mikrotik-router-not-routing-secondary-networks-pptp-4100a377-1
title: "questions-33391-why-central-mikrotik-router-not-routing-secondary-networks-pptp--4100a377"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-33391-why-central-mikrotik-router-not-routing-secondary-networks-pptp--4100a377.md
source_anchor: ""
source_lines: [1, 119]
sha256: e65f4a1ec76ba0ebcd6e0cd4458903d4a52562fe751e0568100fa51700af942f
---

# questions-33391-why-central-mikrotik-router-not-routing-secondary-networks-pptp--4100a377

I have 3 offices:
1.2.1.0/24 main1      office
1.2.2.0/24 secondary2 office
1.2.3.0/24 secondary3 office
There are three Mikrotik 951Ui-2HnD (current-firmware: 3.18) in each office:
main1:      LAN: 1.2.1.1 | WAN: XXX.XXX.XXX.115
secondary2: LAN: 1.2.2.1 | WAN: XXX.XXX.XXX.112
secondary3: LAN: 1.2.3.1 | WAN: XXX.XXX.XXX.116
There are no any /ip firewall rules yet in each mikrotiks.
Main1 mikrotik configuration
/interface pptp-server server
set enabled=yes
/ip address
add address=1.2.1.1/24 interface=bridge-local network=1.2.1.0
add address=XXX.XXX.XXX.115/29 interface=ether1-gateway network=XXX.XXX.XXX.112
/ip route
add distance=1 gateway=XXX.XXX.XXX.113
add distance=2 dst-address=1.2.2.0/24 gateway=172.16.1.2
add distance=2 dst-address=1.2.3.0/24 gateway=172.16.1.3
/ppp secret
add local-address=172.16.1.1 name=secondary3 password=123 profile=pptp-in remote-address=172.16.1.3 service=pptp
add local-address=172.16.1.1 name=secondary2 password=123 profile=pptp-in remote-address=172.16.1.2 service=pptp
Secondary2 mikrotik configuration
/interface pptp-client
add add-default-route=no allow=mschap1,mschap2 connect-to=XXX.XXX.XXX.115 dial-on-demand=no disabled=no keepalive-timeout=60 max-mru=1450 \
    max-mtu=1450 mrru=disabled name=pptp-out1 password=123 profile=default-encryption user=secondary2
/ip address
add address=1.2.2.1/24 interface=bridge-local network=1.2.2.0
add address=XXX.XXX.XXX.117/29 interface=ether1-gateway network=XXX.XXX.XXX.112
/ip route
add distance=1 gateway=XXX.XXX.XXX.113
add distance=1 dst-address=1.2.1.0/24 gateway=172.16.1.1
add distance=1 dst-address=1.2.3.0/24 gateway=172.16.1.1
Secondary3 mikrotik configuration
interface pptp-client
add add-default-route=no allow=mschap1,mschap2 connect-to=XXX.XXX.XXX.115 dial-on-demand=no disabled=no keepalive-timeout=60 max-mru=1450 \
    max-mtu=1450 mrru=disabled name=pptp-out1 password=123 profile=default-encryption user=secondary3
/ip address
add address=1.2.3.1/24 interface=bridge-local network=1.2.3.0
add address=XXX.XXX.XXX.116/29 interface=ether1-gateway network=XXX.XXX.XXX.112
/ip route
add distance=1 gateway=XXX.XXX.XXX.113
add distance=1 dst-address=1.2.1.0/24 gateway=172.16.1.1
add distance=1 dst-address=1.2.2.0/24 gateway=172.16.1.1
Secondary2 see's Main1, but not see's secondary3
[secondary2 ] /ping 1.2.1.1 ... ok
[secondary2 ] /ping 1.2.3.1 ... timeout
[secondary2 ] /tool traceroute 1.2.3.1
  # ADDRESS                          LOSS SENT    LAST
  1 172.16.1.1                         0%    1     2ms
  2                                  100%    1 timeout
  ...
[secondary2 ] /ip address print  
  Flags: X - disabled, I - invalid, D - dynamic 
  #   ADDRESS            NETWORK         INTERFACE                                                                                          
  0   1.2.2.1/24         1.2.2.0         bridge-local                                                                                       
  1   XXX.XXX.XXX.117/29   XXX.XXX.XXX.112   ether1-gateway                                                                                     
  2 D 172.16.1.2/32      172.16.1.1      pptp-out1  
[secondary2 ] /ip route print
  Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, 
  B - blackhole, U - unreachable, P - prohibit 
   #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
   0 A S  0.0.0.0/0                          XXX.XXX.XXX.113             1
   1 A S  1.2.1.0/24                         172.16.1.1                1
   2 ADC  1.2.2.0/24         1.2.2.1         bridge-local              0
   3 A S  1.2.3.0/24                         172.16.1.1                1
   4 ADC  XXX.XXX.XXX.112/29   XXX.XXX.XXX.117   ether1-gateway            0
   5 ADC  172.16.1.1/32      172.16.1.2      pptp-out1                 0
Secondary3 see's Main1, but not see's secondary2
[secondary3 ] /ping 1.2.1.1 ... ok
[secondary3 ] /ping 1.2.2.1 ... timeout
[secondary3 ] /tool traceroute 1.2.2.1
  # ADDRESS                          LOSS SENT    LAST
  1 172.16.1.1                         0%    1     2ms
  2                                  100%    1 timeout
  ...
[secondary3 ] /ip address print  
 Flags: X - disabled, I - invalid, D - dynamic 
 #   ADDRESS            NETWORK         INTERFACE                                                                                          
 0   1.2.3.1/24         1.2.3.0         bridge-local                                                                                       
 1   XXX.XXX.XXX.116/29   XXX.XXX.XXX.112   ether1-gateway                                                                                     
 2 D 172.16.1.3/32      172.16.1.1      pptp-out1
[secondary3 ] /ip route print
 Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, 
 B - blackhole, U - unreachable, P - prohibit 
  #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
  0 A S  0.0.0.0/0                          XXX.XXX.XXX.113             1
  1 A S  1.2.1.0/24                         172.16.1.1                1
  2 A S  1.2.2.0/24                         172.16.1.1                1
  3 ADC  1.2.3.0/24         1.2.3.1         bridge-local              0
  4 ADC  XXX.XXX.XXX.112/29   XXX.XXX.XXX.116   ether1-gateway            0
  5 ADC  172.16.1.1/32      172.16.1.3      pptp-out1                 0
Main1 see's both
[main1 ] /ping 1.2.2.1 ... ok (mikrotik secondary2)
[main1 ] /ping 1.2.2.2 ... ok (bd_server secondary2)
[main1 ] /ping 1.2.3.1 ... ok (mikrotik secondary3)
[main1 ] /ping 1.2.3.2 ... ok (bd_server secondary3)
[main1 ] /tool traceroute 1.2.3.2
  # ADDRESS                          LOSS SENT    LAST     AVG    BEST   WORST STD-DEV STATUS                                               
  1 172.16.1.3                         0%    3   0.6ms     0.7     0.6     0.8     0.1                                                      
  2 1.2.3.2                            0%    3   0.6ms     0.8     0.6     1.3     0.3
[main1 ] /ip address print
  Flags: X - disabled, I - invalid, D - dynamic 
  #   ADDRESS            NETWORK         INTERFACE                                                                                          
  0   1.2.1.1/24         1.2.1.0         bridge-local                                                                                       
  1   XXX.XXX.XXX.115/29   XXX.XXX.XXX.112   ether1-gateway                                                                                     
  2 D 172.16.1.1/32      172.16.1.3      <pptp-secondary3>                                                                                      
  3 D 172.16.1.1/32      172.16.1.2      <pptp-secondary2>   
[main1 ] /ip firewall filter<SAFE> /ip route print
  Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, 
  B - blackhole, U - unreachable, P - prohibit 
  #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
  0 A S  0.0.0.0/0                          XXX.XXX.XXX.113             1
  1 ADC  1.2.1.0/24         1.2.1.1         bridge-local                  0
  2 A S  1.2.2.0/24                         172.16.1.2                    2
  3 A S  1.2.3.0/24                         172.16.1.3                    2
  4 ADC  XXX.XXX.XXX.112/29   XXX.XXX.XXX.115   ether1-gateway                0
  5 ADC  172.16.1.2/32      172.16.1.1      <pptp-secondary2>             0
  6 ADC  172.16.1.3/32      172.16.1.1      <pptp-secondary3>             0
