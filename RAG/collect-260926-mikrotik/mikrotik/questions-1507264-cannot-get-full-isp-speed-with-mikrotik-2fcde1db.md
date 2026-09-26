---
id: collect-260926-mikrotik/mikrotik/questions-1507264-cannot-get-full-isp-speed-with-mikrotik-2fcde1db
title: "questions-1507264-cannot-get-full-isp-speed-with-mikrotik-2fcde1db"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-1507264-cannot-get-full-isp-speed-with-mikrotik-2fcde1db.md
source_anchor: ""
source_lines: [1, 128]
sha256: 4fa92572aef9c08e0f219dda5b9babfcc7139c77664807a06c785270091ad3c6
---

# questions-1507264-cannot-get-full-isp-speed-with-mikrotik-2fcde1db

I have CRS109-8G-1S-2HnD as gateway and wAP ac as WiFi hotspot which is connected to gateway via cat5e ethernet using DHCP. My ISP gives me 1gigabit internet speed. The problem is i can get only 350mbits internet speed while connected to wAP ac 5Ghz hotspot from my Macbook Pro. If i connect directly to ISP ethernet then i get about 900mbits.
Both RouterOS versions are 6.46.
On gateway i have no firewall rules except NAT and FastTrack rules. On hotspot i have no firewall rules at all.
Here's my gateway config:
/interface bridge
add admin-mac=E4:8D:8C:9D:E8:CD auto-mac=no comment=defconf name=bridge
/interface wireless
set [ find default-name=wlan1 ] band=2ghz-b/g/n channel-width=20/40mhz-XX disabled=no distance=indoors frequency=auto installation=indoor mode=ap-bridge ssid=gateway \
    wireless-protocol=802.11
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface wireless security-profiles
set [ find default=yes ] authentication-types=wpa-psk,wpa2-psk mode=dynamic-keys supplicant-identity=MikroTik wpa-pre-shared-key=**** wpa2-pre-shared-key=****
/ip pool
add name=dhcp ranges=192.168.88.10-192.168.88.254
/ip dhcp-server
add address-pool=dhcp disabled=no interface=bridge name=defconf
/interface bridge port
add bridge=bridge comment=defconf interface=ether3
add bridge=bridge comment=defconf interface=ether4
add bridge=bridge comment=defconf interface=ether5
add bridge=bridge comment=defconf interface=ether6
add bridge=bridge comment=defconf interface=ether7
add bridge=bridge comment=defconf interface=ether8
add bridge=bridge comment=defconf interface=sfp1
add bridge=bridge comment=defconf interface=wlan1
/ip neighbor discovery-settings
set discover-interface-list=LAN
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1 list=WAN
add interface=ether2 list=WAN
/ip address
add address=192.168.88.1/24 comment=defconf interface=ether3 network=192.168.88.0
/ip dhcp-client
add comment=defconf default-route-distance=2 interface=ether1 use-peer-dns=no
add default-route-distance=4 disabled=no interface=ether2 use-peer-dns=no
/ip dhcp-server network
add address=192.168.88.0/24 comment=defconf gateway=192.168.88.1
/ip dns
set allow-remote-requests=yes servers=1.1.1.1
/ip dns static
add address=192.168.88.1 name=router.lan
/ip firewall filter
add action=fasttrack-connection chain=forward connection-state=established,related
add action=accept chain=forward connection-state=established,related
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN
/lcd interface pages
set 0 interfaces=wlan1
/system clock
set time-zone-name=Europe/Moscow
/system identity
set name="Floor#0(GW)"
/tool mac-server
set allowed-interface-list=LAN
/tool mac-server mac-winbox
set allowed-interface-list=LAN
Here's my WiFi hotspot config:
/interface bridge
add comment=defconf fast-forward=no name=bridge
/interface wireless
set [ find default-name=wlan1 ] band=2ghz-b/g/n country=russia disabled=no distance=indoors frequency=2457 frequency-mode=regulatory-domain mode=ap-bridge radio-name=tik ssid=\
    home-sweet-home wireless-protocol=802.11 wmm-support=enabled wps-mode=disabled
set [ find default-name=wlan2 ] band=5ghz-a/n/ac channel-width=20/40/80mhz-Ceee country="united states" disabled=no distance=indoors frequency=auto mode=ap-bridge ssid=\
    home-sweet-home-5ghz wireless-protocol=802.11
/interface ethernet
set [ find default-name=ether1 ] speed=100Mbps
/interface list
add exclude=dynamic name=discover
add name=mactel
add name=mac-winbox
/interface wireless security-profiles
set [ find default=yes ] authentication-types=wpa2-psk eap-methods="" group-key-update=30m mode=dynamic-keys supplicant-identity=MikroTik wpa-pre-shared-key=**** \
    wpa2-pre-shared-key=****
/ip hotspot profile
set [ find default=yes ] html-directory=flash/hotspot
/ip pool
add name=default-dhcp ranges=192.168.88.10-192.168.88.254
/ip dhcp-server
add address-pool=default-dhcp authoritative=after-2sec-delay interface=bridge name=defconf
/system logging action
add email-start-tls=yes email-to=**** name=email target=email
/interface bridge port
add bridge=bridge comment=defconf interface=wlan1
add bridge=bridge comment=defconf interface=wlan2
add bridge=bridge hw=no interface=ether1
/ip neighbor discovery-settings
set discover-interface-list=discover
/interface list member
add interface=wlan1 list=discover
add interface=wlan2 list=discover
add interface=bridge list=discover
add interface=bridge list=mactel
add interface=bridge list=mac-winbox
/ip address
add address=192.168.88.1/24 comment=defconf disabled=yes interface=bridge network=192.168.88.0
/ip dhcp-client
add comment=defconf disabled=no interface=bridge
/ip dhcp-server network
add address=192.168.88.0/24 comment=defconf gateway=192.168.88.1 netmask=24
/ip dns
set allow-remote-requests=yes
/ip dns static
add address=192.168.88.1 name=gateway.local
/ip service
set telnet disabled=yes
set ftp disabled=yes
set ssh disabled=yes
set api disabled=yes
set api-ssl disabled=yes
/ip ssh
set allow-none-crypto=yes forwarding-enabled=remote
/system clock
set time-zone-name=Europe/Moscow
/system identity
set name=Floor#1
/system logging
add disabled=yes topics=wireless,debug
/system ntp client
set enabled=yes primary-ntp=85.21.78.23 secondary-ntp=185.68.101.10
/tool e-mail
set address=smtp.yandex.ru from=**** password="****" port=587 start-tls=yes user=****
/tool mac-server
set allowed-interface-list=mactel
/tool mac-server mac-winbox
set allowed-interface-list=mac-winbox
