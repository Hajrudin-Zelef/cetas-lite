---
id: collect-260926-mikrotik/mikrotik/capsman-can-t-get-20mhz-channels-on-2-4ghz-2
title: "2024-12-18 16:44:27 by RouterOS 7.16.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2024-12-18"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/capsman-can-t-get-20mhz-channels-on-2-4ghz.md
source_anchor: ""
source_lines: [17, 184]
sha256: 997ee38962ef27dae1050f7366b4b32115d31a0459c7bd72567785acafbe179a
---

# 2024-12-18 16:44:27 by RouterOS 7.16.2

```
# 2024-12-18 16:44:27 by RouterOS 7.16.2
# software id = UZ54-YG8N
#
# model = E50UG
# serial number = XXXXXXXXX
/interface bridge
add admin-mac=F4:1E:57:7A:XX:XX auto-mac=no comment=defconf name=bridge
/interface pppoe-client
add add-default-route=yes dial-on-demand=yes disabled=no interface=ether1 \
name="Exetel PPPoE" user=fake@qld.exetel.com.au
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface wifi configuration
add channel.band=2ghz-n .frequency=2412,2437,2462 .secondary-frequency=\
disabled .width=20mhz disabled=no name="2G Subsettings"
add channel.band=5ghz-ac .frequency=5180,5260,5500,5660 .width=20/40/80mhz \
disabled=no name="5G Subsettings"
/interface wifi steering
add disabled=no name=steering1 neighbor-group=dynamic-jazzy-ap-blahblah rrm=\
yes wnm=yes
/interface wifi configuration
add country=Australia disabled=no mode=ap name="Jasmine Home" \
security.authentication-types=wpa2-psk,wpa3-psk .wps=disable ssid=\
jazzy-ap steering=steering1 steering.neighbor-group=\
dynamic-jazzy-ap-fd22ae5e .rrm=yes .wnm=yes
/ip pool
add name=default-dhcp ranges=192.168.88.10-192.168.88.254
/ip dhcp-server
add address-pool=default-dhcp interface=bridge name=defconf
/disk settings
set auto-media-interface=bridge auto-media-sharing=yes auto-smb-sharing=yes
/interface bridge port
add bridge=bridge comment=defconf interface=ether2
add bridge=bridge comment=defconf interface=ether3
add bridge=bridge comment=defconf interface=ether4
add bridge=bridge comment=defconf interface=ether5
/ip neighbor discovery-settings
set discover-interface-list=LAN
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1 list=WAN
/interface wifi capsman
set enabled=yes package-path="" require-peer-certificate=no upgrade-policy=\
none
/interface wifi provisioning
add action=create-dynamic-enabled disabled=no master-configuration=\
"Jasmine Home" name-format=%I
add action=create-dynamic-enabled disabled=no master-configuration=\
"2G Subsettings" supported-bands=2ghz-n
add action=create-dynamic-enabled disabled=no master-configuration=\
"5G Subsettings" supported-bands=5ghz-ac
/ip address
add address=192.168.88.1/24 comment=defconf interface=bridge network=\
192.168.88.0
/ip dhcp-client
add comment=defconf interface=ether1
/ip dhcp-server network
add address=192.168.88.0/24 comment=defconf dns-server=192.168.88.1 gateway=\
192.168.88.1
/ip dns
set allow-remote-requests=yes servers=9.9.9.11,149.112.112.11
/ip dns static
add address=192.168.88.1 comment=defconf name=router.lan type=A
/ip firewall filter
add action=accept chain=input comment=\
"defconf: accept established,related,untracked" connection-state=\
established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=\
invalid
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=accept chain=input comment=\
"defconf: accept to local loopback (for CAPsMAN)" dst-address=127.0.0.1
add action=drop chain=input comment="defconf: drop all not coming from LAN" \
in-interface-list=!LAN
add action=accept chain=forward comment="defconf: accept in ipsec policy" \
ipsec-policy=in,ipsec
add action=accept chain=forward comment="defconf: accept out ipsec policy" \
ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" \
connection-state=established,related hw-offload=yes
add action=accept chain=forward comment=\
"defconf: accept established,related, untracked" connection-state=\
established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" \
connection-state=invalid
add action=drop chain=forward comment=\
"defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat \
connection-state=new in-interface-list=WAN
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" \
ipsec-policy=out,none out-interface-list=WAN
add action=masquerade chain=srcnat src-address=192.168.88.0/24
/ipv6 firewall address-list
add address=::/128 comment="defconf: unspecified address" list=bad_ipv6
add address=::1/128 comment="defconf: lo" list=bad_ipv6
add address=fec0::/10 comment="defconf: site-local" list=bad_ipv6
add address=::ffff:0.0.0.0/96 comment="defconf: ipv4-mapped" list=bad_ipv6
add address=::/96 comment="defconf: ipv4 compat" list=bad_ipv6
add address=100::/64 comment="defconf: discard only " list=bad_ipv6
add address=2001:db8::/32 comment="defconf: documentation" list=bad_ipv6
add address=2001:10::/28 comment="defconf: ORCHID" list=bad_ipv6
add address=3ffe::/16 comment="defconf: 6bone" list=bad_ipv6
/ipv6 firewall filter
add action=accept chain=input comment=\
"defconf: accept established,related,untracked" connection-state=\
established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=\
invalid
add action=accept chain=input comment="defconf: accept ICMPv6" protocol=\
icmpv6
add action=accept chain=input comment="defconf: accept UDP traceroute" \
dst-port=33434-33534 protocol=udp
add action=accept chain=input comment=\
"defconf: accept DHCPv6-Client prefix delegation." dst-port=546 protocol=\
udp src-address=fe80::/10
add action=accept chain=input comment="defconf: accept IKE" dst-port=500,4500 \
protocol=udp
add action=accept chain=input comment="defconf: accept ipsec AH" protocol=\
ipsec-ah
add action=accept chain=input comment="defconf: accept ipsec ESP" protocol=\
ipsec-esp
add action=accept chain=input comment=\
"defconf: accept all that matches ipsec policy" ipsec-policy=in,ipsec
add action=drop chain=input comment=\
"defconf: drop everything else not coming from LAN" in-interface-list=\
!LAN
add action=accept chain=forward comment=\
"defconf: accept established,related,untracked" connection-state=\
established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" \
connection-state=invalid
add action=drop chain=forward comment=\
"defconf: drop packets with bad src ipv6" src-address-list=bad_ipv6
add action=drop chain=forward comment=\
"defconf: drop packets with bad dst ipv6" dst-address-list=bad_ipv6
add action=drop chain=forward comment="defconf: rfc4890 drop hop-limit=1" \
hop-limit=equal:1 protocol=icmpv6
add action=accept chain=forward comment="defconf: accept ICMPv6" protocol=\
icmpv6
add action=accept chain=forward comment="defconf: accept HIP" protocol=139
add action=accept chain=forward comment="defconf: accept IKE" dst-port=\
500,4500 protocol=udp
add action=accept chain=forward comment="defconf: accept ipsec AH" protocol=\
ipsec-ah
add action=accept chain=forward comment="defconf: accept ipsec ESP" protocol=\
ipsec-esp
add action=accept chain=forward comment=\
"defconf: accept all that matches ipsec policy" ipsec-policy=in,ipsec
add action=drop chain=forward comment=\
"defconf: drop everything else not coming from LAN" in-interface-list=\
!LAN
/system clock
set time-zone-autodetect=no time-zone-name=Australia/Queensland
/system note
set show-at-login=no
/system ntp client
set enabled=yes
/system ntp client servers
add address=0.au.pool.ntp.org
add address=1.au.pool.ntp.org
add address=2.au.pool.ntp.org
/tool mac-server
set allowed-interface-list=LAN
/tool mac-server mac-winbox
set allowed-interface-list=LAN
```
