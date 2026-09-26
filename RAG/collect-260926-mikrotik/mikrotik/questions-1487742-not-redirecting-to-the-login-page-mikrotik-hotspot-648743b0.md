---
id: collect-260926-mikrotik/mikrotik/questions-1487742-not-redirecting-to-the-login-page-mikrotik-hotspot-648743b0
title: "sep/30/2019 12:41:41 by RouterOS 6.45.6"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-1487742-not-redirecting-to-the-login-page-mikrotik-hotspot-648743b0.md
source_anchor: ""
source_lines: [1, 142]
sha256: 2785bb9edcd0db1191aa994595d322461f3bfbc48abafa2a276b4398f6d061b9
---

# sep/30/2019 12:41:41 by RouterOS 6.45.6

I'm using the hap lite router, and I created a hotspot. Everything worked just fine!
The problem is when I'm connected to the WiFi from an device, it does not redirect automatically to ask for user and login pass, u have to manually set it on the URL to the specific dns name. so is there a way that i can fix it?
Note: this happen to https and http sites. the hotspot users are on the subnet 10.10.4.0/24
this is my configuration file script :
# sep/30/2019 12:41:41 by RouterOS 6.45.6
# software id = 4L42-C3TX
#
# model = RB941-2nD
# serial number = A1C30A560349
/interface bridge
add admin-mac=74:4D:28:40:EF:59 auto-mac=no comment=defconf name=bridge
add name=bridge-hotspot
/interface wireless
set [ find default-name=wlan1 ] band=2ghz-b/g/n channel-width=20/40mhz-Ce \
disabled=no distance=indoors frequency=auto mode=ap-bridge ssid=\
"Hot Spot Nasa" wireless-protocol=802.11
/interface ethernet
set [ find default-name=ether1 ] advertise=\
10M-half,10M-full,100M-half,100M-full,1000M-half,1000M-full
set [ find default-name=ether2 ] advertise=\
10M-half,10M-full,100M-half,100M-full,1000M-half,1000M-full
set [ find default-name=ether3 ] advertise=\
10M-half,10M-full,100M-half,100M-full,1000M-half,1000M-full
set [ find default-name=ether4 ] advertise=\
10M-half,10M-full,100M-half,100M-full,1000M-half,1000M-full
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface wireless security-profiles
set [ find default=yes ] authentication-types=wpa-psk,wpa2-psk eap- 
methods="" \
supplicant-identity=MikroTik wpa-pre-shared-key=xxxxxxxxxxxx\
wpa2-pre-shared-key=xxxxxxxxxx
add authentication-types=wpa-psk,wpa2-psk mode=dynamic-keys name=profile \
supplicant-identity=MikroTik wpa-pre-shared-key=xxxxxxxxxxxxx \
wpa2-pre-shared-key=xxxxxxxxxxxxx
/ip hotspot profile
add dns-name=hot.spot hotspot-address=10.10.4.1 login-by=http-chap,mac- 
cookie \
name=hsprof1 use-radius=yes
/ip pool
add name=dhcp ranges=10.10.3.2-10.10.3.253
add name=hs-pool-9 ranges=10.10.4.2-10.10.4.254
/ip dhcp-server
add address-pool=dhcp disabled=no interface=bridge name=defconf
add address-pool=hs-pool-9 disabled=no interface=bridge-hotspot lease- 
time=1h \
name=dhcp1
/ip hotspot
add address-pool=hs-pool-9 addresses-per-mac=1 disabled=no interface=\
bridge-hotspot name=hotspot1 profile=hsprof1
/interface bridge filter
# no interface
add action=drop chain=forward in-interface=*8
# no interface
add action=drop chain=forward out-interface=*8
/interface bridge port
add bridge=bridge comment=defconf interface=ether2
add bridge=bridge comment=defconf interface=ether3
add bridge=bridge comment=defconf interface=ether4
add bridge=bridge-hotspot comment=defconf interface=wlan1
add bridge=bridge interface=*8
/ip neighbor discovery-settings
set discover-interface-list=LAN
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1 list=WAN
/interface wireless access-list
add ap-tx-limit=40000000
/ip address
add address=10.10.3.1/24 comment=defconf interface=ether2 network=10.10.3.0
add address=192.168.1.2/24 interface=ether1 network=192.168.1.0
add address=10.10.4.1/24 interface=bridge-hotspot network=10.10.4.0
/ip dhcp-client
add comment=defconf dhcp-options=hostname,clientid interface=ether1
/ip dhcp-server network
add address=10.10.3.0/24 comment=defconf gateway=10.10.3.1 netmask=24
add address=10.10.4.0/24 comment="hotspot network" gateway=10.10.4.1
/ip dns
set allow-remote-requests=yes servers=1.1.1.1,1.0.0.1
/ip dns static
add address=10.10.3.1 name=router.lan
/ip firewall filter
add action=passthrough chain=unused-hs-chain comment=\
"place hotspot rules here" disabled=yes
add action=accept chain=input comment=\
"defconf: accept established,related,untracked" connection-state=\
established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection- 
state=\
invalid
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=drop chain=input comment="defconf: drop all not coming from LAN" 
\
in-interface-list=!LAN
add action=accept chain=forward comment="defconf: accept in ipsec policy" \
ipsec-policy=in,ipsec
add action=accept chain=forward comment="defconf: accept out ipsec policy" \
ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" \
connection-state=established,related
add action=accept chain=forward comment=\
"defconf: accept established,related, untracked" connection-state=\
established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" \
connection-state=invalid
add action=drop chain=forward comment=\
"defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat \
connection-state=new in-interface-list=WAN
/ip firewall nat
add action=passthrough chain=unused-hs-chain comment=\
"place hotspot rules here" disabled=yes
add action=masquerade chain=srcnat comment="defconf: masquerade" \
ipsec-policy=out,none out-interface-list=WAN
add action=dst-nat chain=dstnat comment=OpenVPN dst-address=192.168.1.2 \
dst-port=443 protocol=tcp to-addresses=10.10.3.231 to-ports=443
add action=masquerade chain=srcnat comment="masquerade hotspot network" \
src-address=10.10.4.0/24
/ip hotspot user
add name=admin password=xxxxxxxxxxxxxxxx
add comment="mark account" name=mark password=xxxxxxxxxxxx server=hotspot1
add comment="welcome jimaras how are you \?" name=Jimaras password=\
xxxxxxxxxxxxx server=hotspot1
add comment=welcome name=john password=xxxxxxxxxxxxxx server=hotspot1
add comment="asus laptop" name=babas password=xxxxxxxxxxx server=hotspot1
add name=Asus password=xxxxxxxxxxxx server=hotspot1
add name=laptopMark password=xxxxxxxxxxxxx server=hotspot1
/ip route
add distance=1 gateway=192.168.1.254
add distance=1 dst-address=10.10.10.0/24 gateway=10.10.3.231
/ip ssh
set allow-none-crypto=yes forwarding-enabled=remote
/radius
add address=127.0.0.1 secret=xxxxxxxxxx service=hotspot
/radius incoming
set accept=yes
/system clock
set time-zone-name=Europe/Athens
/tool mac-server
set allowed-interface-list=LAN
/tool mac-server mac-winbox
set allowed-interface-list=LAN
