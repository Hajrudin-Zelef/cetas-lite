---
id: collect-260926-mikrotik/mikrotik/questions-908823-mikrotik-routing-vpn-packets-to-wrong-interface-5ff02213
title: "DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE"
domain: mikrotik
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "ethernet"]
source: docs/RAG/lot-mikrotik/forum/vpn/questions-908823-mikrotik-routing-vpn-packets-to-wrong-interface-5ff02213.md
source_anchor: ""
source_lines: [1, 52]
sha256: ff03121f5abdff10b67956c95d1f1496eecec106c5701ba92fadd7a61e87ac9b
---

# DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE

I'm trying connect my network to AWS VPC using Static IP AWS VPN. I've followed the instructions of AWS and configured correctly my MikroTik router and can ping the Ubuntu instance I've attached in this VPC. As I have 2 ISP in my MikroTik, I configured another VPN in AWS and decided use two interfaces in distinct subnets at my Ubuntu AWS instance:
$ ifconfig
eth0      Link encap:Ethernet  HWaddr 0A:72:28:AF:C7:CE  
          inet addr:192.168.254.4  Bcast:192.168.254.15  Mask:255.255.255.240
          inet6 addr: fe80::872:28ff:feaf:c7ce/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:9001  Metric:1
          RX packets:24297 errors:0 dropped:0 overruns:0 frame:0
          TX packets:21627 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:7074474 (6.7 MiB)  TX bytes:2362441 (2.2 MiB)
eth1      Link encap:Ethernet  HWaddr 0A:A2:0F:7C:75:FC  
          inet addr:192.168.254.20  Bcast:192.168.254.31  Mask:255.255.255.240
          inet6 addr: fe80::8a2:fff:fe7c:75fc/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:9001  Metric:1
          RX packets:24996 errors:0 dropped:0 overruns:0 frame:0
          TX packets:25227 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:2307467 (2.2 MiB)  TX bytes:5382812 (5.1 MiB)
Now, even I having two VPN correctly configured in boot sides, just one VPN works at time. For example: imagine VPN named aws-main is configured to work with 192.168.254.0/28 subnet and VPN aws-backup is configured to work with 192.168.254.16/28 subnet, if I want to ping the ip 192.168.254.4 I need disable the route #2, and if I want ping the ip 192.168.254.20 I need disable the route #1
#      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
1 A S  ;;; ISP 1
    0.0.0.0/0                          yyy.yyy.yyy.yyy           1
2   S  ;;; ISP 2
    0.0.0.0/0                          zzz.zzz.zzz.zzz           1
4 ADC  yyy.yyy.yyy.0/24    y.y.y.y        ether3                 0
6 ADC  zzz.zzz.zzz.0/30    z.z.z.z        ether2                 0
9 ADC  192.168.15.0/24    192.168.15.254  ether5                 0
If I try ping 192.168.254.20 when both routes #1 and #2 are enabled I can see this message in log:
backup-out srcnat in:(unknown 0) out:ether3, src-mac xxx, proto ICMP(type 8,code 0), 192.168.15.31 ->192.168.254.20, len 84
Mikrotik is routing the packet to ether3 instead of ether2 (the correct gateway for VPN aws-backup), I think this is the problem, but I don't know how to force it to correct interface.
/ip firewall nat> print
0    chain=srcnat action=accept src-address=192.168.15.0/24 
  dst-address=192.168.254.0/28 log=yes log-prefix="main-out" 
1    chain=srcnat action=accept src-address=192.168.15.0/24 
  dst-address=192.168.254.16/28 log=yes log-prefix="backup-out" 
2    chain=srcnat action=accept src-address=192.168.254.0/28 
  dst-address=192.168.15.0/24 log=yes log-prefix="main-in" 
3    chain=srcnat action=accept src-address=192.168.254.16/28 
  dst-address=192.168.15.0/24 log=yes log-prefix="backup-in" 
4    chain=srcnat action=masquerade log=no log-prefix="masquerade"
this are the policies:
0  A  ;;; AWS Tunnel 2 - BACKUP
   src-address=0.0.0.0/0 src-port=any dst-address=192.168.254.16/28 
   dst-port=any protocol=all action=encrypt level=require 
   ipsec-protocols=esp tunnel=yes sa-src-address=zzz.zzz.zzz.zzz 
   sa-dst-address=x.x.x.x proposal=AWS ph2-count=9 
2  A  ;;; AWS Tunnel 1 - MAIN
   src-address=0.0.0.0/0 src-port=any dst-address=192.168.254.0/28 
   dst-port=any protocol=all action=encrypt level=require 
   ipsec-protocols=esp tunnel=yes sa-src-address=yyy.yyy.yyy.yyy 
   sa-dst-address=K.K.K.K proposal=AWS ph2-count=8
At least for me, mikrotik is not routing the VPN packets using the correct subnet to correct gateway. Is this the problem? How can I solve this?
