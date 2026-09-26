---
id: collect-260926-mikrotik/mikrotik/openvpn-ping-issues-1
title: "aug/27/2018 16:26:29 by RouterOS 6.42.7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/openvpn-ping-issues.md
source_anchor: ""
source_lines: [1, 296]
sha256: 811a64deff784c2fea148abbb532f352ade3b0d02a6d269238cbe413ac136b9b
---

# aug/27/2018 16:26:29 by RouterOS 6.42.7

I feel like i have been through every tutorial and forum post out there on fixing this issue.  Just like countless other posts, I am not able to ping a computer on our office lan subnet from a computer on our remote subnet over an openvpn connection.  I can ping the gateways from either side, but the packets will just not go through.  Hopefully one of you kind souls can spot an issue in our configuration.

**Server Side**

# aug/27/2018 16:26:29 by RouterOS 6.42.7

# software id = 6A1G-CH30


# model = RouterBOARD wAP R-2nD

# serial number = 870E0760D839

/interface lte

set [ find ] mac-address=B6:9E:9D:2E:30:96 name=lte1

/interface bridge

add admin-mac=64:D1:54:7D:CE:72 auto-mac=no comment=defconf name=bridge

/interface wireless

set [ find default-name=wlan1 ] band=2ghz-b/g/n channel-width=20/40mhz-Ce country="united states" disabled=no distance=indoors frequency=auto mode=ap-bridge ssid=DCMRouter wireless-protocol=802.11

/interface list

add comment=defconf name=WAN

add comment=defconf name=LAN

/interface lte apn

set [ find default=yes ] apn=wap.tracfone

/interface wireless security-profiles

set [ find default=yes ] authentication-types=wpa-psk,wpa2-psk mode=dynamic-keys supplicant-identity=MikroTik

/ip hotspot profile

set [ find default=yes ] html-directory=flash/hotspot

/ip pool

add name=dhcp ranges=192.168.1.120-192.168.1.130

add name=pool-ovpn ranges=10.255.255.2-10.255.255.254

/ip dhcp-server

add address-pool=dhcp disabled=no interface=bridge name=defconf

/ppp profile

add local-address=10.255.255.1 name=profile1 remote-address=pool-ovpn

/routing ospf area

add area-id=0.0.0.255 name=area255

/interface bridge port

add bridge=bridge comment=defconf interface=ether1

add bridge=bridge comment=defconf interface=wlan1

/ip neighbor discovery-settings

set discover-interface-list=LAN

/interface list member

add comment=defconf interface=bridge list=LAN

add comment=defconf interface=lte1 list=WAN

/interface ovpn-server server

set certificate=SERVER default-profile=profile1 enabled=yes netmask=32

/ip address

add address=192.168.1.111/24 comment=defconf interface=ether1 network=192.168.1.0

/ip dhcp-server network

add address=192.168.1.0/24 comment=defconf gateway=192.168.1.111 netmask=24

/ip dns

set allow-remote-requests=yes

/ip dns static

add address=192.168.1.111 name=router.lan

/ip firewall filter

add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked

add action=accept chain=input dst-port=1194 protocol=tcp

add action=accept chain=input protocol=ospf

add action=accept chain=forward src-address=10.255.255.254

add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid

add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp

add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN

add action=accept chain=forward comment="defconf: accept in ipsec policy" ipsec-policy=in,ipsec

add action=accept chain=forward comment="defconf: accept out ipsec policy" ipsec-policy=out,ipsec

add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related

add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked

add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid

add action=drop chain=forward comment="defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN

/ip firewall nat

add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN

/ppp secret

add name=DCRRouter1 profile=profile1

/routing ospf network

add area=area255 network=10.255.255.0/24

add area=area255 network=192.168.1.0/24

/system clock

set time-zone-autodetect=no time-zone-name=America/Kentucky/Louisville

/system logging

add topics=debug

/system ntp client

set enabled=yes server-dns-names=0.pool.ntp.org,1.pool.ntp.org,2.pool.ntp.org,3.pool.ntp.org

/system routerboard settings

set silent-boot=no

/tool mac-server

set allowed-interface-list=LAN

/tool mac-server mac-winbox

set allowed-interface-list=LAN

**Server Routes**

Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, B - blackhole, U - unreachable, P - prohibit

# DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE

0 ADS  0.0.0.0/0                          lte1                      2

1 ADC  10.255.255.254/32  10.255.255.1             0

2 ADC  100.84.9.172/32    100.84.9.172    lte1                      0

3 ADC  192.168.1.0/24     192.168.1.111   bridge                    0

4 ADo  192.168.88.0/24                    10.255.255.254          110

**Client Config**

# aug/27/2018 16:36:28 by RouterOS 6.42.7

# software id = GXG5-3WJ6


# model = RouterBOARD wAP R-2nD

# serial number = 7B7307D42EF2

/interface lte

set [ find ] mac-address=46:7F:C0:A2:6C:B6 name=lte1

/interface ovpn-client

add connect-to=notreallyourname.ddns.net mac-address=02:3D:FD:9E:80:DD name=ovpn-DCMain user=DCRRouter1

/interface bridge

add admin-mac=64:D1:54:7D:C7:66 auto-mac=no comment=defconf name=bridge

/interface wireless

set [ find default-name=wlan1 ] band=2ghz-b/g/n channel-width=20/40mhz-Ce country="united states" disabled=no distance=indoors frequency=auto mode=ap-bridge ssid=DCRRouter1 wireless-protocol=802.11

/interface list

add comment=defconf name=WAN

add comment=defconf name=LAN

/interface lte apn

set [ find default=yes ] apn=wap.tracfone

/interface wireless security-profiles

set [ find default=yes ] authentication-types=wpa-psk,wpa2-psk mode=dynamic-keys supplicant-identity=MikroTik

/ip hotspot profile

set [ find default=yes ] html-directory=flash/hotspot

/ip pool

add name=dhcp ranges=192.168.88.10-192.168.88.254

/ip dhcp-server

add address-pool=dhcp disabled=no interface=bridge name=defconf

/routing ospf area

add area-id=0.0.0.255 name=area255

/interface bridge port

add bridge=bridge comment=defconf interface=ether1

add bridge=bridge comment=defconf interface=wlan1

/ip neighbor discovery-settings

set discover-interface-list=LAN

/interface list member

add comment=defconf interface=bridge list=LAN

add comment=defconf interface=lte1 list=WAN

/ip address

add address=192.168.88.1/24 comment=defconf interface=ether1 network=192.168.88.0

/ip dhcp-server network

add address=192.168.88.0/24 comment=defconf gateway=192.168.88.1

/ip dns

set allow-remote-requests=yes

/ip dns static

add address=192.168.88.1 name=router.lan

/ip firewall filter

add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked

add action=accept chain=input protocol=ospf

add action=accept chain=forward src-address=10.255.255.0/24

add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid

add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp

add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN

add action=accept chain=forward comment="defconf: accept in ipsec policy" ipsec-policy=in,ipsec

add action=accept chain=forward comment="defconf: accept out ipsec policy" ipsec-policy=out,ipsec

add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related

add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked

add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid

add action=drop chain=forward comment="defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN

/ip firewall nat

