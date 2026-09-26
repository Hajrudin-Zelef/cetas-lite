---
id: collect-260926-mikrotik/mikrotik/questions-1376419-unstable-10gbit-connection-of-mikrotik-css326-24g-2srm-with-mi-c485cf8a
title: "nov/18/2018 14:20:22 by RouterOS 6.43.4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-1376419-unstable-10gbit-connection-of-mikrotik-css326-24g-2srm-with-mi-c485cf8a.md
source_anchor: ""
source_lines: [1, 113]
sha256: 1c07df45ae28181a534575764498196ac3db79e6e0ed749f879996c791656e6a
---

# nov/18/2018 14:20:22 by RouterOS 6.43.4

I use one ASUS XG-C100C network card in each of two PCs that are connected to Mikrotik CSS326-24G-2S+RM (each via Mikrotik S+RJ10 SFP).
One PC is very close (2 meter Cat5E cable), the other is far (15 meter Cat5E cable). I get brilliant 10Gbit speed on both links. Everything works fine until I even slightly touch SFP+ module or RJ-45 connector. If not touched, links work brilliant at least for a month without problems (have not tested longer yet). If touched, link immediately goed down and then random interruptions start to occur. Next interruption can happen in several minutes or in several hours - unpredictable.
I use latest RouterOS and firmware. The only thing that seems to help - I try to insert connectors with SFP+ modules more strongly (firmly) into the SFP+ ports, but this seems strange and unstable, because even a slight touch can break connection again.
What am I doing wrong? Any suggestions?
Mikrotik CSS326-24G-2S+RM config:
# nov/18/2018 14:20:22 by RouterOS 6.43.4
# software id = SUYY-DNX3
#
# model = CRS326-24G-2S+
# serial number = 763C08FD32BB
/interface bridge
add admin-mac=CC:2D:E0:D3:EC:44 auto-mac=no comment=defconf name=bridge
/interface ethernet
set [ find default-name=ether1 ] speed=100Mbps
set [ find default-name=ether2 ] speed=100Mbps
set [ find default-name=ether3 ] speed=100Mbps
set [ find default-name=ether4 ] speed=100Mbps
set [ find default-name=ether5 ] speed=100Mbps
set [ find default-name=ether6 ] speed=100Mbps
set [ find default-name=ether7 ] speed=100Mbps
set [ find default-name=ether8 ] speed=100Mbps
set [ find default-name=ether9 ] speed=100Mbps
set [ find default-name=ether10 ] speed=100Mbps
set [ find default-name=ether11 ] speed=100Mbps
set [ find default-name=ether12 ] speed=100Mbps
set [ find default-name=ether13 ] speed=100Mbps
set [ find default-name=ether14 ] speed=100Mbps
set [ find default-name=ether15 ] speed=100Mbps
set [ find default-name=ether16 ] speed=100Mbps
set [ find default-name=ether17 ] speed=100Mbps
set [ find default-name=ether18 ] speed=100Mbps
set [ find default-name=ether19 ] speed=100Mbps
set [ find default-name=ether20 ] speed=100Mbps
set [ find default-name=ether21 ] speed=100Mbps
set [ find default-name=ether22 ] speed=100Mbps
set [ find default-name=ether23 ] speed=100Mbps
set [ find default-name=ether24 ] speed=100Mbps
set [ find default-name=sfp-sfpplus1 ] advertise=10M-half,10M-full,100M-half,100M-full,1000M-half,1000M-full,2500M-full,5000M-full,10000M-full rx-flow-control=auto tx-flow-control=auto
set [ find default-name=sfp-sfpplus2 ] speed=10Gbps
/interface list
add name=WAN
add name=LAN
/interface bridge port
add bridge=bridge interface=ether1
add bridge=bridge interface=ether2
add bridge=bridge interface=ether3
add bridge=bridge interface=ether4
add bridge=bridge interface=ether5
add bridge=bridge interface=ether6
add bridge=bridge interface=ether7
add bridge=bridge interface=ether8
add bridge=bridge interface=ether9
add bridge=bridge interface=ether10
add bridge=bridge interface=ether11
add bridge=bridge interface=ether12
add bridge=bridge interface=ether13
add bridge=bridge interface=ether14
add bridge=bridge interface=ether15
add bridge=bridge interface=ether16
add bridge=bridge interface=ether17
add bridge=bridge interface=ether18
add bridge=bridge interface=ether19
add bridge=bridge interface=ether20
add bridge=bridge interface=ether21
add bridge=bridge interface=ether22
add bridge=bridge interface=ether23
add bridge=bridge interface=ether24
add bridge=bridge interface=sfp-sfpplus1
add bridge=bridge interface=sfp-sfpplus2
/interface list member
add interface=ether1 list=WAN
add interface=ether2 list=LAN
add interface=ether3 list=LAN
add interface=ether4 list=LAN
add interface=ether5 list=LAN
add interface=ether6 list=LAN
add interface=ether7 list=LAN
add interface=ether8 list=LAN
add interface=ether9 list=LAN
add interface=ether10 list=LAN
add interface=ether11 list=LAN
add interface=ether12 list=LAN
add interface=ether13 list=LAN
add interface=ether14 list=LAN
add interface=ether15 list=LAN
add interface=ether16 list=LAN
add interface=ether17 list=LAN
add interface=ether18 list=LAN
add interface=ether19 list=LAN
add interface=ether20 list=LAN
add interface=ether21 list=LAN
add interface=ether22 list=LAN
add interface=ether23 list=LAN
add interface=ether24 list=LAN
add interface=sfp-sfpplus1 list=LAN
add interface=sfp-sfpplus2 list=LAN
/ip address
add address=192.168.9.9/8 comment=defconf interface=ether2 network=192.0.0.0
/ip cloud
set ddns-enabled=yes
/ip dns
set servers=8.8.8.8
/ip route
add distance=1 gateway=192.168.9.1
/system clock
set time-zone-name=Europe/Moscow
/system leds
set 0 leds=user-led type=off
add interface=ether1 leds=ether1-led type=interface-activity
/system logging
add action=remote disabled=yes topics=error,warning,critical,info
/system routerboard settings
set boot-os=router-os silent-boot=no
