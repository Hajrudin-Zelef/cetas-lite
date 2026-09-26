---
id: collect-260926-mikrotik/mikrotik/questions-1643985-mikrotik-drop-in-bandwidth-speed-87927ddd-2
title: "apr/22/2021 19:32:24 by RouterOS 6.48"
domain: mikrotik
role: reference
task: reference
actors: ["Samsung"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/qos/questions-1643985-mikrotik-drop-in-bandwidth-speed-87927ddd.md
source_anchor: ""
source_lines: [160, 268]
sha256: 63cb722891404c1fb215a4762c60c33178f53020f7ef6b08443f1278be0e0781
---

# apr/22/2021 19:32:24 by RouterOS 6.48

add address=www.youtube.com list="Block youtube"
add address=googlevideo.com list="Block youtube"
add address=v16a.tiktokcdn.com list="Block tiktok"
add address=p16-tiktokcdn-com.akamaized.net list="Block tiktok"
add address=log.tiktokv.com list="Block tiktok"
add address=ib.tiktokv.com list="Block tiktok"
add address=api-h2.tiktokv.com list="Block tiktok"
add address=v16m.tiktokcdn.com list="Block tiktok"
add address=api.tiktokv.com list="Block tiktok"
add address=v19.tiktokcdn.com list="Block tiktok"
add address=mon.musical.ly list="Block tiktok"
add address=api2-16-h2.musical.ly list="Block tiktok"
add address=api2.musical.ly list="Block tiktok"
add address=log2.musical.ly list="Block tiktok"
add address=api2-21-h2.musical.ly list="Block tiktok"
add address=192.168.2.101 disabled=yes list=VPN
add address=240.0.0.0/4 comment=Reserved list=Bogons
add address=192.168.2.177-192.168.2.188 list="Allow WAN"
add address=192.168.2.118/31 list="Allow WAN"
add address=192.168.2.190-192.168.2.202 list="Allow WAN"
add address=192.168.2.173 list="Allow WAN"
add address=192.168.2.254-192.168.2.250 list="Allow WAN"
add address=192.168.2.101 list="Allow WAN"
add address=192.168.2.89 list="Allow WAN"
add address=192.168.2.118 list="Allow Lan"
add address=192.168.2.16 list="Allow Lan"
/ip firewall filter
# inactive time
add action=drop chain=forward comment="Disable ALL WAN" out-interface=\
    ether1-WAN src-address-list="!Allow Lan" time=\
    21h30m-7h,sun,mon,tue,wed,thu,fri,sat
add action=drop chain=forward comment="Disable Selective WAN" disabled=yes \
    out-interface=ether1-WAN src-address-list="!Allow WAN"
add action=drop chain=forward comment="Disable WAN on DHCP with time" \
    disabled=yes out-interface=ether1-WAN src-address=\
    192.168.2.3-192.168.2.99 time=20h-17h,sun,mon,tue,wed,thu,fri,sat
add action=drop chain=forward comment="Tiktok drop" dst-address-list=\
    "Block tiktok" log=yes log-prefix=tk protocol=tcp
add action=accept chain=input disabled=yes port=69 protocol=udp
add action=accept chain=forward disabled=yes port=69 protocol=udp
add action=drop chain=forward comment="defconf: drop invalid" \
    connection-state=invalid
add action=drop chain=input comment="DNS from outside drop UDP" dst-port=53 \
    in-interface=ether1-WAN protocol=udp
add action=drop chain=input comment="DNS from outside drop TCP" dst-port=53 \
    in-interface=ether1-WAN protocol=tcp
add action=drop chain=forward comment=\
    "defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat \
    connection-state=new in-interface=ether1-WAN
add action=drop chain=forward comment="Drop to bogon list" dst-address-list=\
    Bogons
add action=accept chain=forward comment="defconf: accept established,related" \
    connection-state=established,related
add action=accept chain=input comment="Allow ping" dst-limit=\
    30,30,dst-address/1m40s limit=30,30:packet protocol=icmp
add action=accept chain=input comment="Accept established" connection-state=\
    established
add action=accept chain=input comment="Accept related" connection-state=\
    related
add action=drop chain=input comment="Drop the rest" in-interface=ether1-WAN
add action=fasttrack-connection chain=forward comment="Fasttrack DNS TCP" \
    dst-port=53 protocol=tcp
add action=fasttrack-connection chain=forward comment="Fasttrack DNS UDP" \
    dst-port=53 protocol=udp
/ip firewall mangle
add action=mark-connection chain=prerouting comment=\
    "Facebook -created automatically Layer 7" connection-mark=no-mark \
    dst-port=53 layer7-protocol=*1 new-connection-mark=youtube_conn \
    passthrough=yes protocol=udp
add action=mark-routing chain=prerouting disabled=yes new-routing-mark=vpn \
    passthrough=yes src-address-list=VPN
/ip firewall nat
add action=redirect chain=dstnat comment="Proxy redirect" disabled=yes \
    dst-port=80 protocol=tcp to-ports=8080
add action=masquerade chain=srcnat disabled=yes out-interface=VPN-NAME
add action=masquerade chain=srcnat comment=Masquerade ipsec-policy=out,none \
    out-interface-list=WAN
add action=dst-nat chain=dstnat dst-port=53 log-prefix=elt protocol=udp \
    src-address=192.168.2.118 to-addresses=8.8.8.8 to-ports=53
add action=dst-nat chain=dstnat dst-port=53 protocol=tcp src-address=\
    192.168.2.118 to-addresses=8.8.8.8
/ip kid-control device
add mac-address=44:D8:84:31:BA:15 name=kyle-iphone user="kSchedule"
add mac-address=A8:5E:45:63:DF:95 name=kyle-gaming user="kSchedule"
add mac-address=00:E0:33:2D:B8:2F name=Kyle-samsung-pc user="kSchedule"
add mac-address=4C:63:71:E3:32:1D name=kyle-xaomi user="kSchedule"
add mac-address=D4:5D:64:04:29:8A name=kyle-gaming-lan user="kSchedule"
add mac-address=B4:B6:76:79:B9:4F name="kyle Samsung PC" user="kSchedule"
/ip proxy
set cache-administrator=anon@gmail.com cache-on-disk=yes cache-path=\
    disk1/webproxy
/ip route
add distance=1 gateway=VPN-NAME routing-mark=vpn
add disabled=yes distance=1 dst-address=192.168.0.1/32 gateway=ether1-WAN
/ip service
set telnet disabled=yes
set ftp disabled=yes
set api disabled=yes
set api-ssl disabled=yes
/ip ssh
set allow-none-crypto=yes forwarding-enabled=remote
/ip upnp
set enabled=yes
/system clock
set time-zone-name=Europe/Malta
/system watchdog
set watchdog-timer=no
/tool bandwidth-server
set enabled=no
