---
id: collect-260926-mikrotik/mikrotik/v7-1-mpls-vpls-with-ipv4-ipv6-ldp-question
title: "TRANSPORT    LOCAL-TRANSPORT  PEER           ADDRESSES"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/v7-1-mpls-vpls-with-ipv4-ipv6-ldp-question.md
source_anchor: ""
source_lines: [1, 128]
sha256: 29190c43e10f8f7e58068af1c689e963bc8e91bc9b77d1566b71f56ecee77afb
---

# TRANSPORT    LOCAL-TRANSPORT  PEER           ADDRESSES

I ran into a similar issue. In 7.2rc7, the IPv4 afi works by itself and the ipv6 afi does too. But if you use the IPv4 and IPv6 afi together, LDP becomes non-functional and targeted LDP for VPLS stops working for IPv4 and IPv6.

Submitted this as a bug to MikroTik (SUP-78710)

**if the ip afi is set, ldp becomes operational and VPLS works**

```
[zuul@ccr2004-2splus-01.test.lab.ipa] > mpls/export 
/mpls interface
add input=yes interface=vlan3100 mpls-mtu=1530
add input=yes interface=vpls-101-ipv6
add input=yes interface=vpls-101-ipv4
/mpls ldp
add afi=ip disabled=no lsr-id=100.127.1.1 transport-addresses=100.127.1.1
/mpls ldp interface
add afi=ipv6 disabled=yes interface=vlan3100 transport-addresses=200:127:1::1
add afi=ip disabled=no interface=vlan3100 transport-addresses=100.127.1.1
[zuul@ccr2004-2splus-01.test.lab.ipa] > mpls/ldp/neighbor/print 
Flags: D, I - INACTIVE; O, T - THROTTLED; t - SENDING-TARGETED-HELLO; v - VPLS; p - PASSIVE
Columns: TRANSPORT, LOCAL-TRANSPORT, PEER, ADDRESSES
#       TRANSPORT    LOCAL-TRANSPORT  PEER           ADDRESSES   
0 DOtvp 100.127.1.2  100.127.1.1      100.127.1.2:0  10.255.44.15
                                                     100.126.1.9 
                                                     100.127.1.2 
                                                     198.18.0.2  
                                                     203.0.113.2 
1 DO  p 100.127.1.3  100.127.1.1      100.127.1.3:0  10.255.44.23
                                                     100.126.1.2 
                                                     100.126.1.10
                                                     100.127.1.3
```

**if the ipv6 afi is set, ldp becomes operational and VPLS works**

```
[zuul@ccr2004-2splus-01.test.lab.ipa] > mpls/export 
/mpls interface
add input=yes interface=vlan3100 mpls-mtu=1530
add input=yes interface=vpls-101-ipv6
add input=yes interface=vpls-101-ipv4
/mpls ldp
add afi=ipv6 disabled=no lsr-id=100.127.1.1 transport-addresses=200:127:1::1
/mpls ldp interface
add afi=ipv6 disabled=no interface=vlan3100 transport-addresses=200:127:1::1
add afi=ip disabled=yes interface=vlan3100 transport-addresses=100.127.1.1
[zuul@ccr2004-2splus-01.test.lab.ipa] > 
[zuul@ccr2004-2splus-01.test.lab.ipa] > mpls/ldp/neighbor/print
Flags: D, I - INACTIVE; O, T - THROTTLED; t - SENDING-TARGETED-HELLO; v - VPLS; p - PASSIVE
Columns: TRANSPORT, LOCAL-TRANSPORT, PEER, ADDRESSES
#       TRANSPORT     LOCAL-TRANSPORT  PEER           ADDRESSES                         
0 DOtvp 200:127:1::2  200:127:1::1     100.127.1.2:0  200:126:2::2                      
                                                      200:127:1::2                      
                                                      203:0:113::2                      
                                                      fc00:da7a:44:1:de2c:6eff:fe8a:990d
                                                      fe80::87:6fff:fe98:ee32           
                                                      fe80::1458:daff:fe81:387e         
                                                      fe80::9472:c1ff:fe70:168f         
                                                      fe80::de2c:6eff:fe8a:990d         
                                                      fe80::de2c:6eff:fe8a:990f         
1 DO  p 200:127:1::3  200:127:1::1     100.127.1.3:0  200:126:1::3                      
                                                      200:126:2::3                      
                                                      200:127:1::3                      
                                                      fe80::1458:daff:fe81:387e         
                                                      fe80::9472:c1ff:fe70:168f         
                                                      fe80::de2c:6eff:fe7a:d0e6         
                                                      fe80::de2c:6eff:fe7a:d0f6
```

**if the ip and ipv6 afi is set, ldp for IPv6 becomes operational but ldp for IPv4 does not. After rebooting the router, LDPv6 neighbors will become operational, but targeted LDP for VPLS over the LDPv6 neighbors no longer functions.**

```
[zuul@ccr2004-2splus-01.test.lab.ipa] > mpls/export 
# apr/03/2022 08:24:25 by RouterOS 7.2rc7
# software id = XF0U-LU9N
#
# model = CCR2004-16G-2S+
# serial number = HBJ07YRKGA0
/mpls interface
add input=yes interface=vlan3100 mpls-mtu=1530
add input=yes interface=vpls-101-ipv6
add input=yes interface=vpls-101-ipv4
/mpls ldp
add afi=ip,ipv6 disabled=no lsr-id=100.127.1.1 transport-addresses=100.127.1.1,200:127:1::1
/mpls ldp interface
add afi=ipv6 disabled=no interface=vlan3100 transport-addresses=200:127:1::1
add afi=ip disabled=no interface=vlan3100 transport-addresses=100.127.1.1
[zuul@ccr2004-2splus-01.test.lab.ipa] > 
[zuul@ccr2004-2splus-01.test.lab.ipa] > mpls/ldp/neighbor/print
Flags: D, I - INACTIVE; O, T - THROTTLED; t - SENDING-TARGETED-HELLO; v - VPLS; p - PASSIVE
Columns: TRANSPORT, LOCAL-TRANSPORT, PEER, ADDRESSES
#       TRANSPORT     LOCAL-TRANSPORT  PEER           ADDRESSES                         
0 DOtvp 200:127:1::2  200:127:1::1     100.127.1.2:0  200:126:2::2                      
                                                      200:127:1::2                      
                                                      203:0:113::2                      
                                                      fc00:da7a:44:1:de2c:6eff:fe8a:990d
                                                      fe80::87:6fff:fe98:ee32           
                                                      fe80::1458:daff:fe81:387e         
                                                      fe80::9472:c1ff:fe70:168f         
                                                      fe80::de2c:6eff:fe8a:990d         
                                                      fe80::de2c:6eff:fe8a:990f         
1 D tv  100.127.1.2                                                                     
2 DO  p 200:127:1::3  200:127:1::1     100.127.1.3:0  200:126:1::3                      
                                                      200:126:2::3                      
                                                      200:127:1::3                      
                                                      fe80::1458:daff:fe81:387e         
                                                      fe80::9472:c1ff:fe70:168f         
                                                      fe80::de2c:6eff:fe7a:d0e6         
                                                      fe80::de2c:6eff:fe7a:d0f6    
                                                      
                                                      
                                                      
#### After reboot 
[zuul@ccr2004-2splus-01.test.lab.ipa] > mpls/ldp/neighbor/print 
Flags: D, I - INACTIVE; O, W, T - THROTTLED; t - SENDING-TARGETED-HELLO; v - VPLS; p - PASSIVE
Columns: TRANSPORT, LOCAL-TRANSPORT, PEER, ADDRESSES
#       TRANSPORT     LOCAL-TRANSPORT  PEER           ADDRESSES                
0 DWtvp 100.127.1.2   100.127.1.1      100.127.1.2:0                           
1 D tv  200:127:1::2                                                           
2 DO  p 200:127:1::3  200:127:1::1     100.127.1.3:0  200:126:1::3             
                                                      200:126:2::3             
                                                      200:127:1::3             
                                                      fe80::1458:daff:fe81:387e
                                                      fe80::9472:c1ff:fe70:168f
                                                      fe80::de2c:6eff:fe7a:d0e6
                                                      fe80::de2c:6eff:fe7a:d0f6
```

The lab topology is below
