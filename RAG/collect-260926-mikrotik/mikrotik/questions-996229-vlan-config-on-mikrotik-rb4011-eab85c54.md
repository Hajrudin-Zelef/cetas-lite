---
id: collect-260926-mikrotik/mikrotik/questions-996229-vlan-config-on-mikrotik-rb4011-eab85c54
title: "dec/19/2019 14:17:44 by RouterOS 6.46"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-996229-vlan-config-on-mikrotik-rb4011-eab85c54.md
source_anchor: ""
source_lines: [1, 75]
sha256: d69b1a388ed023f81b41369f71caf126f0ee02e1ac79950e9d0817e5c1866005
---

# dec/19/2019 14:17:44 by RouterOS 6.46

I'm trying to configure the router to have two subnets 192.168.1.0/24 and 192.168.2.0/24. The first subnet is the default and all ports should use this as the default IP range, the second subnet should only be active when plugged into ETH2 Port on the router. Block traffic from the two subnets to talk to each other, but allow both to the internet.
I'd also like to allow tagging to VLAN ID 2 on all ports other than ETH2.
The issue is that I have configured the VLAN, the ip's, the dhcp server, and the tagged ports. But I'm not pulling the proper IP from the port. I have included the .rsc file I exported from the router. Any help is greatly appreciated.
# dec/19/2019 14:17:44 by RouterOS 6.46
# software id = PLDE-DYLT
#
# model = RB4011iGS+5HacQ2HnD
# serial number = B3A30A573840
/interface bridge add admin-mac=74:4D:28:5C:66:E9 auto-mac=no comment=defconf name=bridge
/interface bridge add name=bridgeVLAN1
/interface bridge add name=bridgeVLAN2 pvid=2 vlan-filtering=yes
/interface wireless
# no supported channel and secondary channel combination
set [ find default-name=wlan1 ] band=5ghz-a/n/ac channel-width=20/40/80mhz-XXXX disabled=no distance=indoors frequency=auto installation=indoor mode=ap-bridge secondary-channel=auto ssid=MikroTik-5C66F2 wireless-protocol=802.11
/interface wireless set [ find default-name=wlan2 ] band=2ghz-b/g/n channel-width=20/40mhz-XX disabled=no distance=indoors frequency=auto installation=indoor mode=ap-bridge ssid=MikroTik-0E1DE9 wireless-protocol=802.11
/interface vlan add interface=bridgeVLAN1 name=vlan1 vlan-id=1
/interface vlan add interface=bridgeVLAN2 name=vlan2 vlan-id=2
/interface ethernet switch port set 0 default-vlan-id=0
/interface ethernet switch port set 1 default-vlan-id=0
/interface ethernet switch port set 2 default-vlan-id=0
/interface ethernet switch port set 3 default-vlan-id=0
/interface ethernet switch port set 4 default-vlan-id=0
/interface ethernet switch port set 5 default-vlan-id=0
/interface ethernet switch port set 6 default-vlan-id=0
/interface ethernet switch port set 7 default-vlan-id=0
/interface ethernet switch port set 8 default-vlan-id=0
/interface ethernet switch port set 9 default-vlan-id=0
/interface ethernet switch port set 10 default-vlan-id=0
/interface ethernet switch port set 11 default-vlan-id=0
/interface list add comment=defconf name=WAN
/interface list add comment=defconf name=LAN
/interface wireless security-profiles set [ find default=yes ] supplicant-identity=MikroTik
/ip pool add name=dhcp ranges=192.168.1.10-192.168.1.254
/ip pool add name=pool2 ranges=192.168.2.1-192.168.2.254
/ip dhcp-server add address-pool=dhcp disabled=no interface=bridge name=defconf
/ip dhcp-server add address-pool=pool2 disabled=no interface=ether1 name=serverVLAN2
/interface bridge port add bridge=bridge comment=defconf interface=ether2
/interface bridge port add bridge=bridgeVLAN2 comment=defconf interface=ether3 pvid=2
/interface bridge port add bridge=bridge comment=defconf interface=ether4
/interface bridge port add bridge=bridge comment=defconf interface=ether5
/interface bridge port add bridge=bridge comment=defconf interface=ether6
/interface bridge port add bridge=bridge comment=defconf interface=ether7
/interface bridge port add bridge=bridge comment=defconf interface=ether8
/interface bridge port add bridge=bridge comment=defconf interface=ether9
/interface bridge port add bridge=bridge comment=defconf interface=ether10
/interface bridge port add bridge=bridge comment=defconf interface=sfp-sfpplus1
/interface bridge port add bridge=bridge comment=defconf interface=wlan1
/interface bridge port add bridge=bridge comment=defconf interface=wlan2
/ip neighbor discovery-settings set discover-interface-list=LAN
/interface bridge vlan add bridge=bridgeVLAN1 tagged=bridgeVLAN1 untagged=ether4 vlan-ids=1
/interface bridge vlan add bridge=bridgeVLAN2 tagged=bridgeVLAN2 untagged=ether3 vlan-ids=2
/interface list member add comment=defconf interface=bridge list=LAN
/interface list member add comment=defconf interface=ether1 list=WAN
/ip address add address=192.168.1.1/24 comment=defconf interface=ether2 network=192.168.1.0
/ip address add address=192.168.2.1/24 interface=vlan2 network=192.168.2.0
/ip dhcp-client add comment=defconf disabled=no interface=ether1
/ip dhcp-server network add address=192.168.1.0/24 comment=defconf gateway=192.168.1.1 netmask=24
/ip dhcp-server network add address=192.168.2.0/24 dns-server=8.8.8.8,8.8.4.4 gateway=192.168.1.1 netmask=24
/ip dns set allow-remote-requests=yes
/ip dns static add address=192.168.1.1 comment=defconf name=router.lan
/ip firewall filter add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
/ip firewall filter add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
/ip firewall filter add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
/ip firewall filter add action=accept chain=input comment="defconf: accept to local loopback (for CAPsMAN)" dst-address=127.0.0.1
/ip firewall filter add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN
/ip firewall filter add action=accept chain=forward comment="defconf: accept in ipsec policy" ipsec-policy=in,ipsec
/ip firewall filter add action=accept chain=forward comment="defconf: accept out ipsec policy" ipsec-policy=out,ipsec
/ip firewall filter add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related
/ip firewall filter add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
/ip firewall filter add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
/ip firewall filter add action=drop chain=forward comment="defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
/ip firewall nat add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN
/system clock set time-zone-name=America/New_York
/tool mac-server set allowed-interface-list=LAN
/tool mac-server mac-winbox set allowed-interface-list=LAN
