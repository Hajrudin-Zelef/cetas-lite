---
id: collect-260926-mikrotik/mikrotik/wireguard-site-to-site-routing-with-a-central-server-4
title: "2023-11-08 21:11:31 by RouterOS 7.11.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2023-11-08"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-site-to-site-routing-with-a-central-server.md
source_anchor: ""
source_lines: [415, 589]
sha256: ee5cafcd2c3e0f2b0ddb687d4d3ac7238961a1b9bcf8232233caa3a8ae0f42cf
---

# 2023-11-08 21:11:31 by RouterOS 7.11.2

```
# 2023-11-08 21:11:31 by RouterOS 7.11.2
# software id = **ELIDED**
#
# model = RBD25GR-5HPacQD2HPnD
# serial number = **ELIDED**
/interface bridge
add admin-mac= **ELIDED** auto-mac=no comment=defconf name=bridge
/interface lte
# SIM not inserted
set [ find default-name=lte1 ] allow-roaming=no band=""
/interface wireless
# managed by CAPsMAN
# channel: 2412/20-Ce/gn(26dBm), SSID: Wifi3, local forwarding
set [ find default-name=wlan1 ] disabled=no ssid=MikroTik
# managed by CAPsMAN
# channel: 5180/20-Ceee/ac(12dBm), SSID: Wifi4, local forwarding
set [ find default-name=wlan2 ] disabled=no ssid=MikroTik
/interface wireguard
add listen-port=13231 mtu=1420 name=wireguard1
/caps-man security
add authentication-types=wpa2-psk comment=defconf disable-pmkid=yes \
    encryption=aes-ccm group-encryption=aes-ccm name=capSec
/caps-man configuration
add channel.band=2ghz-b/g/n .control-channel-width=20mhz .extension-channel=\
    XX comment=defconf datapath.client-to-client-forwarding=yes \
    .local-forwarding=yes distance=indoors installation=indoor name=cfg-2ghz \
    security=capSec ssid=Wifi3
add channel.band=5ghz-a/n/ac .control-channel-width=20mhz .extension-channel=\
    XXXX comment=defconf datapath.client-to-client-forwarding=yes \
    .local-forwarding=yes distance=indoors installation=indoor name=\
    cfg-5ghz-ac security=capSec ssid=Wifi4
add channel.band=5ghz-a/n .control-channel-width=20mhz .extension-channel=XX \
    comment=defconf datapath.client-to-client-forwarding=yes \
    .local-forwarding=yes distance=indoors installation=indoor name=\
    cfg-5ghz-an security=capSec ssid=Wifi4
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
add authentication-types=wpa2-psk comment=defconf disable-pmkid=yes mode=\
    dynamic-keys name=wpsSync supplicant-identity=MikroTik
/interface wireless
set [ find default-name=wlan3 ] band=5ghz-a/n/ac channel-width=20/40mhz-XX \
    disabled=no hide-ssid=yes mode=ap-bridge security-profile=wpsSync ssid=\
    Backhaul
/ip pool
add name=default-dhcp ranges=192.168.184.20-192.168.184.200
/ip dhcp-server
add address-pool=default-dhcp interface=bridge lease-time=1m name=defconf
/caps-man manager
set enabled=yes
/caps-man manager interface
set [ find default=yes ] forbid=yes
add comment=defconf disabled=no interface=bridge
/caps-man provisioning
add action=create-dynamic-enabled comment=defconf hw-supported-modes=gn \
    master-configuration=cfg-2ghz name-format=prefix name-prefix=2ghz
add action=create-dynamic-enabled comment=defconf hw-supported-modes=ac \
    master-configuration=cfg-5ghz-ac name-format=prefix name-prefix=5ghz-ac
add action=create-dynamic-enabled comment=defconf hw-supported-modes=an \
    master-configuration=cfg-5ghz-an name-format=prefix name-prefix=5ghz-an
/interface bridge port
add bridge=bridge ingress-filtering=no interface=wlan3
/ip neighbor discovery-settings
set discover-interface-list=LAN
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1 list=WAN
add comment=defconf interface=lte1 list=WAN
add interface=ether2 list=LAN
add interface=wireguard1 list=LAN
/interface wireguard peers
add allowed-address=192.168.224.0/24,192.168.16.0/22,192.168.188.0/24 \
    endpoint-address=[PUBLIC_IP] endpoint-port=13231 interface=wireguard1 \
    persistent-keepalive=5s public-key="..."
/interface wireless cap
# 
set bridge=bridge caps-man-addresses=127.0.0.1 enabled=yes interfaces=\
    wlan1,wlan2
/ip address
add address=192.168.184.1/24 comment=defconf interface=bridge network=\
    192.168.184.0
add address=192.168.224.5/24 interface=wireguard1 network=192.168.224.0
/ip dhcp-client
add interface=ether1
add add-default-route=no interface=ether2
/ip dhcp-server network
add address=192.168.184.0/24 comment=defconf dns-server=192.168.184.1 \
    gateway=192.168.184.1
/ip dns
set allow-remote-requests=yes
/ip dns static
add address=192.168.184.1 comment=defconf name=centrala.utechov
add address=192.168.184.2 name=garaz.utechov
add address=192.168.184.3 name=chodba.utechov
/ip firewall filter
add action=accept chain=input comment="Accept established,related,untracked" \
    connection-state=established,related,untracked
add action=drop chain=input comment="Drop invalid" connection-state=invalid
add action=accept chain=input comment="Simply allow all ICMP" protocol=icmp
add action=accept chain=input comment="Local loopback for CAPsMAN" \
    dst-address=127.0.0.1
add action=accept chain=input comment="WinBox from LAN" dst-port=8291 \
    in-interface-list=LAN protocol=tcp
add action=accept chain=input comment="Bandwith test" dst-port=2000 protocol=\
    tcp
add action=drop chain=input comment="Drop everything not listed above"
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
/ip route
add disabled=no distance=1 dst-address=192.168.16.0/22 \
    gateway=wireguard1 pref-src="" routing-table=main scope=30 \
    suppress-hw-offload=no target-scope=10
add disabled=no dst-address=192.168.188.0/24 gateway=\
    wireguard1 routing-table=main suppress-hw-offload=no
/ip service
set telnet disabled=yes
set ftp disabled=yes
set www disabled=yes
set ssh disabled=yes
set api disabled=yes
set api-ssl disabled=yes
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
add action=accept chain=input comment="defconf: accept UDP traceroute" port=\
    33434-33534 protocol=udp
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
