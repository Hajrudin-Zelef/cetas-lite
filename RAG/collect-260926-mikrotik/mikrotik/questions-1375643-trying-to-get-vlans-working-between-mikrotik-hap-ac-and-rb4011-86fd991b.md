---
id: collect-260926-mikrotik/mikrotik/questions-1375643-trying-to-get-vlans-working-between-mikrotik-hap-ac-and-rb4011-86fd991b
title: "nov/15/2018 22:53:31 by RouterOS 6.43.4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-1375643-trying-to-get-vlans-working-between-mikrotik-hap-ac-and-rb4011-86fd991b.md
source_anchor: ""
source_lines: [1, 138]
sha256: b73ba2569a5e58bfe1fe66109337b86f40db3885ae14b923bf250a25e8e7411c
---

# nov/15/2018 22:53:31 by RouterOS 6.43.4

Here is the network diagram:
I have set up 3 vlans: vlan1: 192.168.9.0/24 vlan11-lan: 192.168.11.0/24 vlan22-guest: 192.168.22.0/24
vlan1 is working perfectly. Every host in 192.168.9.0/24 network is pinging each other.
vlan11-lan and vlan22-guest do not work correctly:
- HAP AC can ping RB4011 (192.168.11.1 or 192.168.22.1)
- RB4011 can ping HAP AC (192.168.11.2 or 192.168.22.2)
- PC5 and PC2 cannot ping each other and cannot ping RB4011 (192.168.11.1) or HAP AC (192.168.11.2). They do not even get MAC addresses of each other in their ARP tables.
- PC3 cannot ping RB4011 (192.168.22.1) or HAP AC (192.168.22.2). Also it does not get their MAC addresses in its ARP table.
I disabled firewall completely, but yet no success. I have masquerade going out from internet interface (eth1 on RB4011iGS+).
RB4011iGS+ config:
# nov/15/2018 22:53:31 by RouterOS 6.43.4
# software id = WP4U-Z565
#
# model = RB4011iGS+
# serial number = 968A09187F4C
/interface bridge
add admin-mac=B8:69:F4:92:25:57 auto-mac=no comment=defconf name=bridge vlan-filtering=yes
/interface ethernet
set [ find default-name=ether1 ] l2mtu=1598
set [ find default-name=ether2 ] l2mtu=1598
set [ find default-name=ether3 ] l2mtu=1598
set [ find default-name=ether4 ] l2mtu=1598
set [ find default-name=ether5 ] l2mtu=1598
set [ find default-name=ether6 ] l2mtu=1598
set [ find default-name=ether7 ] l2mtu=1598
set [ find default-name=ether8 ] l2mtu=1598
set [ find default-name=ether9 ] l2mtu=1598
set [ find default-name=ether10 ] l2mtu=1598
/interface vlan
add interface=ether10 name=vlan11-lan vlan-id=11
add interface=ether10 name=vlan22-guest vlan-id=22
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/ip pool
add name=dhcp ranges=192.168.9.50-192.168.9.254
/ip dhcp-server
add address-pool=dhcp disabled=no interface=bridge name=defconf
/interface bridge port
add bridge=bridge interface=ether10
add bridge=bridge interface=ether5 pvid=11
add bridge=bridge interface=ether6
/ip neighbor discovery-settings
set discover-interface-list=LAN
/interface bridge vlan
add bridge=bridge tagged=ether10 vlan-ids=11
add bridge=bridge tagged=ether10 vlan-ids=22
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1 list=WAN
add list=LAN
/ip address
add address=192.168.100.2/24 interface=ether1 network=192.168.100.0
add address=192.168.22.1/24 interface=vlan22-guest network=192.168.22.0
add address=192.168.9.1/24 interface=bridge network=192.168.9.0
add address=192.168.11.1/24 interface=vlan11-lan network=192.168.11.0
/ip cloud
set ddns-enabled=yes
/ip dhcp-server network
add address=192.168.9.0/24 gateway=192.168.9.1 netmask=24
/ip dns
set allow-remote-requests=yes
/ip dns static
add address=192.168.9.1 name=router.lan
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN
add action=netmap chain=dstnat dst-port=3484 in-interface=ether1 protocol=tcp to-addresses=192.168.9.4 to-ports=3306
add action=netmap chain=dstnat dst-port=443 in-interface=ether1 protocol=tcp to-addresses=192.168.9.6 to-ports=3389
add action=masquerade chain=srcnat dst-port=80 protocol=tcp src-address=192.168.9.0/24
add action=netmap chain=dstnat dst-port=80 in-interface=ether1 protocol=tcp to-addresses=192.168.9.4 to-ports=80
/ip route
add distance=1 gateway=192.168.100.1
/ip traffic-flow
set cache-entries=32k interfaces=local
/system clock
set time-zone-name=Europe/Moscow
/system identity
set name=RB4011
/system routerboard settings
set silent-boot=no
/tool mac-server
set allowed-interface-list=LAN
/tool mac-server mac-winbox
set allowed-interface-list=LAN
/tool sniffer
set filter-interface=ether10
HAP AC config:
# nov/15/2018 22:47:07 by RouterOS 6.43.2
# software id = R9TC-1I4K
#
# model = RouterBOARD 962UiGS-5HacT2HnT
# serial number = 6737065A9A5D
/interface bridge
add admin-mac=6C:3B:6B:11:EB:C1 auto-mac=no name=bridge vlan-filtering=yes
/interface wireless
set [ find default-name=wlan1 ] band=2ghz-b/g/n channel-width=20/40mhz-Ce disabled=no distance=indoors frequency=auto mode=ap-bridge ssid=MikroTik-11EBC7 wireless-protocol=802.11
set [ find default-name=wlan2 ] band=5ghz-a/n/ac channel-width=20/40/80mhz-Ceee disabled=no distance=indoors frequency=auto mode=ap-bridge ssid=MikroTik-11EBC6 wireless-protocol=802.11
/interface vlan
add interface=ether1 name=vlan11-lan vlan-id=11
add interface=ether1 name=vlan22-guest vlan-id=22
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface wireless security-profiles
set [ find default=yes ] authentication-types=wpa-psk,wpa2-psk mode=dynamic-keys supplicant-identity=MikroTik wpa-pre-shared-key=1234567123 wpa2-pre-shared-key=1234567123
/ip hotspot profile
set [ find default=yes ] html-directory=flash/hotspot
/interface bridge port
add bridge=bridge interface=ether1
add bridge=bridge interface=ether3
add bridge=bridge interface=ether4 pvid=11
add bridge=bridge interface=ether5 pvid=22
/interface bridge vlan
add bridge=bridge tagged=ether1 vlan-ids=11
add bridge=bridge tagged=ether1 vlan-ids=22
/interface list member
add comment=defconf interface=bridge list=LAN
add interface=sfp1 list=WAN
/ip address
add address=192.168.22.2/24 interface=vlan22-guest network=192.168.22.0
add address=192.168.9.2/24 interface=bridge network=192.168.9.0
add address=192.168.11.2/24 interface=vlan11-lan network=192.168.11.0
/ip dns
set allow-remote-requests=yes
/ip dns static
add address=192.168.9.2 name=router.lan
/ip route
add distance=1 gateway=192.168.9.1
/system clock
set time-zone-name=Europe/Moscow
/system identity
set name=HAP_AC
/system routerboard settings
set silent-boot=no
/tool sniffer
set filter-interface=ether1 filter-ip-address=!192.168.13.2/32
