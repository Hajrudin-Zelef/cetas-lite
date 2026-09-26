---
id: collect-260926-rattrapage/rattrapage/questions-821873-nat-smtp-to-a-specific-public-ip-on-a-mikrotik-bc79ecce
title: "dec/20/2016 21:15:40 by RouterOS 6.30.4"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-rattrapage/ai-llm/questions-821873-nat-smtp-to-a-specific-public-ip-on-a-mikrotik-bc79ecce.md
source_anchor: ""
source_lines: [1, 131]
sha256: 63387d05ef1bf26cc75dfb68e3823dcdc14fd76382812902b2d5d9724a23b651
---

# dec/20/2016 21:15:40 by RouterOS 6.30.4

We had some peeps who supplied us a Mikrotik with our fibre connection, but have since moved the contract to another provider, and they don't want to help me configure the device. I desperately need a hand, I am out of my depth.
We have some public IPs 154.117.185.242 - all normal traffic, surfing, guest wifi on this one 154.117.185.243 - our web server :80 is on this one 192.168.10.157 :80 154.117.185.244 154.117.185.245 - another web server :80 on this one 192.168.10.9 :80 154.117.185.246
For some reason 154.117.185.243 has been majorly blaklisted, I think someone had a virus. So I want to move outgoing SMTP (port 25) from 154.117.185.242 to 154.117.185.244 so that people can reliably recieve their transactional mail.
I tried to create this rule: srcnat src.add: 192.168.10.157 protocol tcp port 25 -> action src-nat to address 154.117.185.244 port 25 (I wish I could attach a screen grab)
But that did nothing.
To be honest, I don't even know where to begin creating this rule. I really need help.
I assume that response communication from the net will also have to have a rule to go back to 192.168.10.157:25 - but if that's neccessary I think for that I can pretty much copy one of the webserver rules and change port numbers. It's the rule for the outgoing stuff that I am really stumped with.
Thanks
Steve
Edit - added export:
# dec/20/2016 21:15:40 by RouterOS 6.30.4
# software id = UUL8-1EL2
#
/ip firewall filter
add action=drop chain=forward src-address=192.168.10.54
/ip firewall nat
add action=masquerade chain=srcnat
add action=dst-nat chain=dstnat dst-address=154.117.185.243 dst-port=80 protocol=tcp \
    to-addresses=192.168.10.157 to-ports=80
add action=dst-nat chain=dstnat dst-address=154.117.185.243 dst-port=81 protocol=tcp \
    to-addresses=192.168.10.241 to-ports=81
add action=dst-nat chain=dstnat dst-address=154.117.185.243 dst-port=82 protocol=tcp \
    to-addresses=192.168.10.242 to-ports=82
add action=dst-nat chain=dstnat dst-address=154.117.185.245 dst-port=21-80 protocol=tcp \
    to-addresses=192.168.10.9 to-ports=21-80
add action=src-nat chain=srcnat protocol=tcp src-address=192.168.10.157 src-port=25 \
    to-addresses=154.117.185.245 to-ports=25
and the whole config:
# dec/22/2016 12:53:00 by RouterOS 6.30.4
# software id = UUL8-1EL2
#
/interface ethernet
set [ find default-name=ether1 ] comment="Bitco Fibre"
set [ find default-name=ether2 ] comment=Internal
set [ find default-name=ether3 ] comment=unused
set [ find default-name=ether5 ] comment="Guests Wifi"
set [ find default-name=ether6 ] comment=unused
/ip neighbor discovery
set ether1 comment="Bitco Fibre"
set ether2 comment=Internal
set ether3 comment=unused
set ether5 comment="Guests Wifi"
set ether6 comment=unused
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/ip pool
add name=dhcp_pool1 ranges=192.168.22.100-192.168.22.254
/ip dhcp-server
add address-pool=dhcp_pool1 disabled=no interface=ether5 lease-time=1d10m \
    name=dhcp1
/snmp community
set [ find default=yes ] addresses=154.66.208.0/24
/ip address
add address=154.117.185.243/22 interface=ether1 network=154.117.184.0
add address=192.168.22.2/24 interface=ether5 network=192.168.22.0
add address=192.168.10.1/24 interface=ether2 network=192.168.10.0
add address=154.117.185.245/22 interface=ether1 network=154.117.184.0
add address=154.117.185.242/22 interface=ether1 network=154.117.184.0
add address=154.117.185.244/22 interface=ether1 network=154.117.184.0
/ip dhcp-server network
add address=192.168.22.0/24 dns-server=41.79.80.34,8.8.8.8 gateway=\
    192.168.22.2
/ip dns
set allow-remote-requests=yes servers=41.79.80.34,8.8.8.8
/ip firewall filter
add action=drop chain=forward disabled=yes src-address=192.168.10.109
add action=drop chain=forward disabled=yes src-address=192.168.10.28
add action=drop chain=forward src-address=192.168.10.54
add action=drop chain=forward disabled=yes src-address=192.168.22.116
add action=drop chain=forward comment="BLOCK SPAMMERS OR INFECTED USERS" \
    dst-port=25 protocol=tcp src-address-list=SPAMMER
add action=add-src-to-address-list address-list=SPAMMER address-list-timeout=\
    23h59m59s chain=forward comment=\
    "Detect and add-list SMTP virus or spammers" connection-limit=10,32 \
    dst-port=25 limit=10,5 protocol=tcp
/ip firewall mangle
add action=mark-routing chain=prerouting dst-port=25 new-routing-mark=\
    "Webserver SMTP" passthrough=no protocol=tcp src-address=192.168.10.157
/ip firewall nat
add action=masquerade chain=srcnat
add action=dst-nat chain=dstnat comment="webserver port 80" dst-address=\
    154.117.185.243 dst-port=80 protocol=tcp to-addresses=192.168.10.157 \
    to-ports=80
add action=dst-nat chain=dstnat comment=CCTV dst-address=154.117.185.243 \
    dst-port=81 protocol=tcp to-addresses=192.168.10.241 to-ports=81
add action=dst-nat chain=dstnat comment="unused CCTV" dst-address=\
    154.117.185.243 dst-port=82 protocol=tcp to-addresses=192.168.10.242 \
    to-ports=82
add action=dst-nat chain=dstnat comment=xmpie dst-address=154.117.185.245 \
    dst-port=80 protocol=tcp to-addresses=192.168.10.9 to-ports=21-80
/ip route
add comment="send web SMTP through 244" distance=1 gateway=154.117.185.244 \
    routing-mark="Webserver SMTP" scope=255
add distance=1 gateway=154.117.185.217
/ip service
set telnet address=192.168.10.0/24
set ftp address=192.168.10.0/24 disabled=yes
set www address=192.168.10.0/24
set ssh address=192.168.10.0/24
set www-ssl address=192.168.10.0/24
set api address=192.168.10.0/24
set winbox address=192.168.10.0/24
set api-ssl address=192.168.10.0/24
/lcd
set time-interval=hour
/snmp
set enabled=yes
/system clock
set time-zone-name=Africa/Johannesburg
/system identity
set name=Bowens
/system routerboard settings
set protected-routerboot=disabled
/system script
add name=SPAMMERS owner=admin source=":log error \\\"----------Users detected \
    like \\\
    \n    SPAMMERS -------------\\\";\
    \n\\n:foreach i in \\[/ip firewall address-list find \\\
    \n    list=spammer\\] do={:set usser \\[/ip firewall address-list get \\\$\
    i \\\
    \n    address\\];\
    \n\\n:foreach j in=\\[/ip hotspot active find address=\\\$usser\\] \\\
    \n    do={:set ip \\[/ip hotspot active get \\\$j user\\];\
    \n\\n:log error \\\$ip;\
    \n\\n:log \\\
    \n    error \\\$usser} };\" policy=ftp,read,write,policy,test,winbox "
/tool graphing interface
add allow-address=192.168.10.0/24
/tool romon port
add
/ip route exporttoo
