---
id: collect-260926-mikrotik/mikrotik/questions-799993-download-rate-or-data-limitation-on-mikrotik-hotspot-dont-apply-a8f5de58
title: "aug/30/2016 16:51:44 by RouterOS 6.34.6"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/qos/questions-799993-download-rate-or-data-limitation-on-mikrotik-hotspot-dont-apply-a8f5de58.md
source_anchor: ""
source_lines: [1, 80]
sha256: 8b80f997300a64f2366081c858bed359cf09cd9a2a630ca51664d487d3eaa6b2
---

# aug/30/2016 16:51:44 by RouterOS 6.34.6

I wonder to know has anyone had a problem of limitation on user in hotspot. (RB3011 - even tested on RB951G)
After I config the hotspot and give permission to users based on the username and password user can connect to the internet but when I add a limitation for user profile in term of download rate or data transfer that each user allowed to used, non of them applied to them though in the queue I can see the limited parameter.
I have directly added the user profile on the mikrotik.
I would be appreciate if you help me, what is the problem?
Thanks
# aug/30/2016 16:51:44 by RouterOS 6.34.6
# software id = ETTV-J0EC
#
/interface bridge
add admin-mac=6C:3B:6B:1C:77:33 auto-mac=no name=Bridge
/interface ethernet
set [ find default-name=ether2 ] name=ether2-master
set [ find default-name=ether3 ] disabled=yes master-port=ether2-master
set [ find default-name=ether4 ] disabled=yes master-port=ether2-master
set [ find default-name=ether5 ] disabled=yes master-port=ether2-master
set [ find default-name=ether6 ] name=ether6-master
set [ find default-name=ether7 ] disabled=yes master-port=ether6-master
set [ find default-name=ether8 ] disabled=yes master-port=ether6-master
set [ find default-name=ether9 ] disabled=yes master-port=ether6-master
set [ find default-name=ether10 ] disabled=yes master-port=ether6-master
set [ find default-name=sfp1 ] disabled=yes
/ip neighbor discovery
set ether1 discover=no
/ip hotspot profile
add dns-name=hotspot.ghadirhotel.ir hotspot-address=192.168.100.1 name="Hotspot Profile" use-radius=yes
/ip hotspot user profile
add name="Limited 10M 64K" rate-limit=64K/64K
/ip pool
add name=Pool ranges=192.168.100.100-192.168.100.250
/ip dhcp-server
add address-pool=Pool disabled=no interface=Bridge name=DHCP
/ip hotspot
add address-pool=Pool addresses-per-mac=1 disabled=no interface=Bridge name=Hotspot profile="Hotspot Profile"
/interface bridge port
add bridge=Bridge interface=ether2-master
add bridge=Bridge interface=ether6-master
add bridge=Bridge disabled=yes interface=sfp1
/ip address
add address=192.168.100.1/24 interface=Bridge network=192.168.100.0
add address=192.168.1.10/24 disabled=yes interface=ether1 network=192.168.1.0
/ip dhcp-client
add dhcp-options=*FFFFFFFF,*FFFFFFFF disabled=no interface=ether1
/ip dhcp-server network
add address=192.168.100.0/24 gateway=192.168.100.1
/ip dns
set allow-remote-requests=yes servers=8.8.8.8,4.2.2.4
/ip dns static
add address=192.168.100.1 name=router
/ip firewall filter
add action=passthrough chain=unused-hs-chain comment="place hotspot rules here" disabled=yes
add chain=input comment="defconf: accept ICMP" protocol=icmp
add chain=input comment="defconf: accept establieshed,related" connection-state=established,related
add action=drop chain=input comment="defconf: drop all from WAN" in-interface=ether1
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related
add chain=forward comment="defconf: accept established,related" connection-state=established,related
add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
add action=drop chain=forward comment="defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface=ether1
/ip firewall nat
add action=passthrough chain=unused-hs-chain comment="place hotspot rules here" disabled=yes
add action=masquerade chain=srcnat comment="defconf: masquerade" out-interface=ether1
add action=masquerade chain=srcnat comment="masquerade hotspot network" src-address=192.168.100.0/24
/ip hotspot ip-binding
add mac-address=F0:BF:97:5C:33:44 server=Hotspot type=bypassed
/ip hotspot user
add name=Admin password=admin
add name=test password=test profile="Limited 10M 64K"
/radius
add address=127.0.0.1 secret=admin service=hotspot
/radius incoming
set accept=yes
/system clock
set time-zone-name=Asia/Tehran
/system routerboard settings
set protected-routerboot=disabled
/tool mac-server
set [ find default=yes ] disabled=yes
add interface=Bridge
/tool mac-server mac-winbox
set [ find default=yes ] disabled=yes
add interface=Bridge
