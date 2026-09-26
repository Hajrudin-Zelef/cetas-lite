---
id: collect-260926-mikrotik/mikrotik/questions-540967-how-do-i-configure-routing-for-an-ipsec-tunnel-between-openswan-ba174fc8
title: "service ipsec status"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-540967-how-do-i-configure-routing-for-an-ipsec-tunnel-between-openswan-ba174fc8.md
source_anchor: ""
source_lines: [1, 36]
sha256: 983724c9f45d364efadda9ab7bd84ed239ced72668b062552e83281137d3c21f
---

# service ipsec status

I am trying to create a site-to-site VPN between a Linux router that runs openswan and shorewall (host A, serving subnet 10.10.0.0/16) and a MikroTek RouterBoard running RouterOS 6.3 (host B, serving 192.168.88.0/24).
The IPSEC tunnel itself seems to be up, host A says:
# service ipsec status
IPsec running  - pluto pid: 4292
pluto pid 4292
1 tunnels up
some eroutes exist
and:
#ipsec auto --status
<SNIP>
000 #2: "office-connect":500 STATE_QUICK_I2 (sent QI2, IPsec SA established); EVENT_SA_REPLACE in 27422s; newest IPSEC; eroute owner; isakmp#1; idle; import:admin initiate
000 #2: "office-connect" esp.65bcd1d@<REDACTED> esp.c8d18ebd@<REDACTED> tun.0@<REDACTED> tun.0@<REDACTED> ref=0 refhim=4294901761
000 #1: "office-connect":500 STATE_MAIN_I4 (ISAKMP SA established); EVENT_SA_REPLACE in 2348s; newest ISAKMP; lastdpd=72s(seq in:0 out:0); idle; import:admin initiate
while on host B:
/ip ipsec remote-peers print
 0 local-address=<REDACTED> remote-address=<REDACTED> state=established side=responder established=11m26s 
and:
/ip ipsec policy print
Flags: T - template, X - disabled, D - dynamic, I - inactive 
 0    src-address=192.168.88.0/24 src-port=any dst-address=10.10.0.0/16 dst-port=any protocol=all action=encrypt level=require ipsec-protocols=esp tunnel=yes 
      sa-src-address=<REDACTED> sa-dst-address=<REDACTED> proposal=Office-Connect priority=0
I followed the instrutions on http://www.shorewall.net/IPSEC-2.6.html to configure shorewall and http://wiki.mikrotik.com/wiki/Manual:IP/IPsec to setup a NAT Bypass rule.
However, I can't actually get any packets through the tunnel, on A:
# ping -c4 192.168.88.1
PING 192.168.88.1 (192.168.88.1) 56(84) bytes of data.
--- 192.168.88.1 ping statistics ---
4 packets transmitted, 0 received, 100% packet loss, time 3016ms
on B:
/ping count=4 10.10.0.1
HOST                                     SIZE TTL TIME  STATUS                                                                                               
10.10.0.1                                               timeout                                                                                              
10.10.0.1                                               timeout                                                                                              
10.10.0.1                                               timeout                                                                                              
10.10.0.1                                               timeout                                                                                              
    sent=4 received=0 packet-loss=100% 
I am at a bit of a loss how to proceed, my experience in networking is not too great. So I'd be helpful for any hints, even just how to debug this problem. I'll gladly provide additional configuration samples or log output if required. Many thanks!
