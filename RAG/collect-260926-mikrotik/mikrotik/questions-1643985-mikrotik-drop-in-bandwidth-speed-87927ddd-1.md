---
id: collect-260926-mikrotik/mikrotik/questions-1643985-mikrotik-drop-in-bandwidth-speed-87927ddd-1
title: "apr/22/2021 19:32:24 by RouterOS 6.48"
domain: mikrotik
role: reference
task: reference
actors: ["Samsung"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/qos/questions-1643985-mikrotik-drop-in-bandwidth-speed-87927ddd.md
source_anchor: ""
source_lines: [1, 159]
sha256: 2e78900d3ec2397a1ac8933636207c04a0b03e4dfb492b8b66a447af623ca575
---

# apr/22/2021 19:32:24 by RouterOS 6.48

When I connect my laptop directly to my modem and run a speed test I get the proper 500Mb/s speed. However, when I connect my laptop to my RB750Gr3 (which is connected to the modem) my internet speed goes down to 200Mb/s. I tried with other laptops/devices with the same results. The problem persists regardless of Queues being enabled or disabled. Is there something I need to check on my Mikrotik so that I can obtain 500Mb/s on my devices connected to my router please? My configuration is the following:
# apr/22/2021 19:32:24 by RouterOS 6.48
# software id = 8GVC-967D
#
# model = RB750Gr3
# serial number = 8AFF091B1B69
/caps-man channel
add band=2ghz-onlyn control-channel-width=20mhz frequency=2462 name=channel1
/interface bridge
add name=bridge1
/interface ethernet
set [ find default-name=ether1 ] name=ether1-WAN
/interface pptp-client
add connect-to=server1.freevpn.me name=VPN-NAME password= user=\
    .me
/caps-man datapath
add bridge=bridge1 name=Bridge
/caps-man security
add authentication-types=wpa2-psk,wpa2-eap encryption=aes-ccm name=security1 \
    passphrase=
/caps-man configuration
add channel=channel1 country=malta datapath=Bridge mode=ap name=Config \
    security=security1 ssid=S
/interface list
add name=WAN
add include=all name=LAN
/interface wireless security-profiles
set [ find default=yes ] authentication-types=wpa2-psk eap-methods="" mode=\
    dynamic-keys supplicant-identity=MikroTik wpa2-pre-shared-key=\
    2WR133301567
/ip firewall layer7-protocol
add name=Facebook regexp="^.+(facebook).*\$"
add name=Youtube regexp=\
    "^.+(youtube.com | googlevideo.com | akamaihd.net).*\$"
add name=Discord regexp="^.+(discord).*\$"
/ip kid-control
add disabled=yes fri=18h-22h mon=18h-22h name="Kyle Schedule" sun=6h-22h thu=\
    18h-22h tue=18h-22h wed=18h-22h
/ip pool
add name=dhcp_pool ranges=192.168.2.10-192.168.2.90
/ip dhcp-server
add address-pool=dhcp_pool disabled=no interface=bridge1 name=dhcp3
/queue simple
add max-limit=16M/400M name="All traffic" target=192.168.2.0/24
add max-limit=15M/400M name=Unlimited parent="All traffic" priority=1/1 \
    target="192.168.2.196/32,192.168.2.197/32,192.168.2.183/32,192.168.2.184/3\
    2,192.168.2.252/32"
add max-limit=3M/150M name=Limited parent="All traffic" target=\
    192.168.2.90/32,192.168.2.89/32,192.168.2.13/32
/user group
set full policy="local,telnet,ssh,ftp,reboot,read,write,policy,test,winbox,pas\
    sword,web,sniff,sensitive,api,romon,dude,tikapp"
/caps-man manager
set enabled=yes
/caps-man manager interface
set [ find default=yes ] forbid=yes
add disabled=no interface=bridge1
/caps-man provisioning
add action=create-dynamic-enabled master-configuration=Config
/interface bridge port
add bridge=bridge1 interface=ether2
add bridge=bridge1 interface=ether3
add bridge=bridge1 interface=ether4
add bridge=bridge1 interface=ether5
/ip neighbor discovery-settings
set discover-interface-list=all
/interface detect-internet
set detect-interface-list=all internet-interface-list=all lan-interface-list=\
    all wan-interface-list=all
/interface list member
add interface=ether1-WAN list=WAN
add interface=ether2 list=LAN
add interface=bridge1 list=LAN
add list=LAN
/ip address
add address=192.168.2.1/24 interface=bridge1 network=192.168.2.0
/ip dhcp-client
add disabled=no interface=ether1-WAN
/ip dhcp-server lease
add address=192.168.2.191 client-id=1:d8:9c:67:62:db:9 comment=\
    "Elton JC HP Wifi" mac-address=D8:9C:67:62:DB:09 server=dhcp3
add address=192.168.2.188 client-id=1:ec:71:db:54:e1:52 comment="Front CCTV" \
    mac-address=EC:71:DB:54:E1:52 server=dhcp3
add address=192.168.2.178 client-id=1:ec:71:db:cf:28:71 comment="Living CCTV" \
    mac-address=EC:71:DB:CF:28:71 server=dhcp3
add address=192.168.2.173 comment="Security System" mac-address=\
    00:1F:08:04:92:7B server=dhcp3
add address=192.168.2.177 client-id=1:f4:81:39:e2:e9:39 comment=\
    "Canon Printer" mac-address=F4:81:39:E2:E9:39 server=dhcp3
add address=192.168.2.254 client-id=1:4:18:d6:9e:f5:9a comment="Ubiquiti AP" \
    mac-address=04:18:D6:9E:F5:9A server=dhcp3
add address=192.168.2.2 client-id=1:74:4d:28:72:fd:c3 comment="BedRoom CAP" \
    mac-address=74:4D:28:72:FD:C3 server=dhcp3
add address=192.168.2.90 client-id=1:d4:5d:64:4:29:8a comment=\
    "K gaming PC Lan" mac-address=D4:5D:64:04:29:8A server=dhcp3
add address=192.168.2.162 comment="Argus CCTV" mac-address=18:62:E4:37:97:DC \
    server=dhcp3
add address=192.168.2.114 mac-address=68:9A:87:82:82:EF server=dhcp3
add address=192.168.2.100 client-id=1:f4:91:1e:d1:b8:74 mac-address=\
    F4:91:1E:D1:B8:74 server=dhcp3
add address=192.168.2.118 client-id=1:66:f9:11:61:a8:69 comment="OnePlus 6" \
    mac-address=66:F9:11:61:A8:69 server=dhcp3
add address=192.168.2.109 client-id=1:c4:84:66:b7:10:9a mac-address=\
    C4:84:66:B7:10:9A server=dhcp3
add address=192.168.2.119 client-id=1:98:9:cf:5a:1c:f1 comment="OnePlus 7" \
    mac-address=98:09:CF:5A:1C:F1 server=dhcp3
add address=192.168.2.199 client-id=1:8c:85:90:78:bf:29 comment="Macbook Pro" \
    mac-address=8C:85:90:78:BF:29 server=dhcp3
add address=192.168.2.198 client-id=1:7c:d3:a:75:82:5d comment=\
    "work laptop wifi" mac-address=7C:D3:0A:75:82:5D server=dhcp3
add address=192.168.2.101 client-id=1:58:40:4e:ae:3e:66 comment="iPad" \
    mac-address=58:40:4E:AE:3E:66 server=dhcp3
add address=192.168.2.197 client-id=1:d0:c6:37:60:cb:10 comment=\
    "Work Laptop Wifi" mac-address=D0:C6:37:60:CB:10 server=dhcp3
add address=192.168.2.194 client-id=1:54:88:e:a0:dd:b3 comment=\
    "Samsung Living Rm TV Wifi" mac-address=54:88:0E:A0:DD:B3 server=dhcp3
add address=192.168.2.120 comment="Android Box FTP" mac-address=\
    00:11:6E:03:08:46 server=dhcp3
add address=192.168.2.196 client-id=1:94:5:bb:16:d1:4c comment="mac wired" \
    mac-address=94:05:BB:16:D1:4C server=dhcp3
add address=192.168.2.193 client-id=1:5e:51:c8:d7:2d:1f comment=\
    "Ipad2" mac-address=5E:51:C8:D7:2D:1F server=dhcp3
add address=192.168.2.192 client-id=1:24:4b:3:a7:c1:37 comment=\
    "Samsung Living Room TV Ethernet" mac-address=24:4B:03:A7:C1:37 server=\
    dhcp3
add address=192.168.2.110 client-id=1:3e:30:3e:f7:c6:be mac-address=\
    3E:30:3E:F7:C6:BE server=dhcp3
add address=192.168.2.89 client-id=1:dc:41:a9:1:30:1a comment="Surface" \
    mac-address=DC:41:A9:01:30:1A server=dhcp3
add address=192.168.2.250 client-id=1:74:ac:b9:6c:4c:c7 comment=\
    "Ubiquity Living rm" mac-address=74:AC:B9:6C:4C:C7 server=dhcp3
add address=192.168.2.202 client-id=1:54:4:a6:a6:db:f4 mac-address=\
    54:04:A6:A6:DB:F4 server=dhcp3
add address=192.168.2.190 client-id=1:c8:d9:d2:9c:4a:15 comment=\
    "Dongle Wired" mac-address=C8:D9:D2:9C:4A:15 server=dhcp3
add address=192.168.2.200 client-id=1:c8:d9:d2:7d:d3:d3 comment="JC Wired" \
    mac-address=C8:D9:D2:7D:D3:D3 server=dhcp3
/ip dhcp-server network
add address=192.168.2.0/24 dns-server=208.67.222.222,208.67.220.220 gateway=\
    192.168.2.1 netmask=24
/ip dns
set allow-remote-requests=yes
/ip firewall address-list
add address=0.0.0.0/8 comment="Self-Identification [RFC 3330]" list=Bogons
add address=10.0.0.0/8 comment="Private[RFC 1918] - CLASS A # Check if you nee\
    d this subnet before enable it" list=Bogons
add address=127.0.0.0/8 comment="Loopback [RFC 3330]" list=Bogons
add address=169.254.0.0/16 comment="Link Local [RFC 3330]" list=Bogons
add address=172.16.0.0/12 comment="Private[RFC 1918] - CLASS B # Check if you \
    need this subnet before enable it" list=Bogons
add address=192.0.2.0/24 comment="Reserved - IANA - TestNet1" list=Bogons
add address=192.88.99.0/24 comment="6to4 Relay Anycast [RFC 3068]" list=\
    Bogons
add address=198.18.0.0/15 comment="NIDB Testing" list=Bogons
add address=198.51.100.0/24 comment="Reserved - IANA - TestNet2" list=Bogons
add address=203.0.113.0/24 comment="Reserved - IANA - TestNet3" list=Bogons
add address=224.0.0.0/4 comment=\
    "MC, Class D, IANA # Check if you need this subnet before enable it" \
    list=Bogons
