---
id: collect-260926-mikrotik/mikrotik/openvpn-ping-issues-2
title: "aug/27/2018 16:26:29 by RouterOS 6.42.7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/openvpn-ping-issues.md
source_anchor: ""
source_lines: [297, 351]
sha256: 617360b68319eda4026919309f0b22386199f0d557f9aec9c5d053b9e2eefe1e
---

# aug/27/2018 16:26:29 by RouterOS 6.42.7

add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN

/ip route

add distance=1 gateway=lte1

/routing ospf network

add area=area255 network=10.255.255.0/24

add area=area255 network=192.168.88.0/24

/system clock

set time-zone-autodetect=no time-zone-name=America/Kentucky/Louisville

/system logging

add topics=debug,!ospf

/system ntp client

set enabled=yes server-dns-names=0.pool.ntp.org,1.pool.ntp.org,2.pool.ntp.org,3.pool.ntp.org

/system routerboard settings

set silent-boot=no

/tool mac-server

set allowed-interface-list=LAN

/tool mac-server mac-winbox

set allowed-interface-list=LAN



**Client Routes**

Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, B - blackhole, U - unreachable, P - prohibit

# DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE

0 A S  0.0.0.0/0                          lte1                      1

1 ADC  10.255.255.1/32    10.255.255.254  ovpn-DCMain               0

2 ADC  33.208.199.255/32  33.208.199.255  lte1                      0

3 ADo  192.168.1.0/24                     10.255.255.1            110

4 ADC  192.168.88.0/24    192.168.88.1    bridge                    0

Any ideas?
