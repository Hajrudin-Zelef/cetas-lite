---
id: collect-260926-mikrotik/mikrotik/wireguard-site-to-site-routing-with-a-central-server-3
title: "2023-11-08 21:11:31 by RouterOS 7.11.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-site-to-site-routing-with-a-central-server.md
source_anchor: ""
source_lines: [298, 414]
sha256: c43c94426ab52dba022f5bae2e81b2dc9a6d831eab36f07fe52f7ae1102d1648
---

# 2023-11-08 21:11:31 by RouterOS 7.11.2

add action=accept chain=forward comment="everyone has internet access" \
    in-interface=!ether2 out-interface=ether2 src-address=192.168.0.0/16
add action=accept chain=forward comment="everyone uses HTTP on pokladnice" \
    dst-address=192.168.16.61 dst-port=80 out-interface=friendly-vlan \
    protocol=tcp
add action=accept chain=forward comment="allow SyncThing" dst-port=22000 \
    out-interface=friendly-vlan protocol=tcp
add action=accept chain=forward comment="friendly talk to friendly" \
    dst-address=192.168.16.0/22 in-interface=friendly-vlan out-interface=\
    friendly-vlan src-address=192.168.16.0/22
add action=accept chain=forward comment="friendly talk to unfriendly" \
    dst-address=192.168.96.0/22 in-interface=friendly-vlan out-interface=\
    unfriendly-vlan src-address=192.168.16.0/22
add action=accept chain=forward comment="manage APs" dst-address=\
    192.168.8.0/23 in-interface=handling-vlan out-interface=backbone-vlan \
    src-address=192.168.4.0/28
add action=accept chain=forward comment="wireguard to friendly" dst-address=\
    192.168.16.0/22 in-interface=wireguard1 out-interface=friendly-vlan \
    src-address=192.168.224.0/24
add action=accept chain=forward comment="wireguard to unfriendly" disabled=yes \
    dst-address=192.168.96.0/22 in-interface=wireguard1 out-interface=\
    unfriendly-vlan src-address=192.168.224.0/24
add action=accept chain=forward comment="TEMPORARY: slatina to friendly" \
    dst-address=192.168.16.0/22 in-interface=wireguard1 out-interface=\
    friendly-vlan src-address=192.168.188.0/24
add action=accept chain=forward comment="TEMPORARY: slatina to unfriendly" \
    dst-address=192.168.96.0/22 in-interface=wireguard1 out-interface=\
    unfriendly-vlan src-address=192.168.188.0/24
add action=accept chain=forward comment="friendly to wireguard" in-interface=\
    friendly-vlan out-interface=wireguard1 src-address=192.168.16.0/22
add action=drop chain=forward comment=\
    "skip log: devices not responding to DHCP" src-address=169.254.0.0/16
add action=drop chain=forward comment="deny the rest" log=yes log-prefix="\?"
add action=accept chain=input
add action=accept chain=input comment="accept established,related" \
    connection-state=established,related
add action=drop chain=input comment="drop invalid" connection-state=invalid
add action=drop chain=input comment="incoming from internet" \
    connection-state=new in-interface=ether2
add action=drop chain=input comment="incoming without public IP" \
    in-interface=ether2 log=yes log-prefix=!public src-address-list=\
    not_in_internet
add action=accept chain=input comment="wide open for management" \
    in-interface=handling-vlan log=yes log-prefix=mgmt
add action=jump chain=input comment="separate chain for ICMP" jump-target=\
    icmp protocol=icmp
add action=accept chain=input comment="DNS server" dst-port=53 protocol=udp \
    src-address=192.168.0.0/16
add action=accept chain=input comment="DHCP server incl. broadcast" dst-port=\
    67 protocol=udp
add action=accept chain=input comment="infrastructure Mikrotik" dst-port=5678 \
    in-interface=backbone-vlan protocol=udp src-address=192.168.8.0/23 \
    src-port=5678
add action=accept chain=input comment="bandwith test server" dst-port=2000 \
    protocol=tcp src-port=""
add action=accept chain=input comment="configure via WiFi [TEMPORARY]" \
    src-address=192.168.16.0/22
add action=drop chain=input comment=\
    "skip log: HIKvision discovery is not used" dst-address=255.255.255.255 \
    dst-port=7989 in-interface=unfriendly-vlan protocol=udp src-address=\
    192.168.96.96
add action=drop chain=input comment=\
    "skip log: devices not responding to DHCP" src-address=169.254.0.0/16
add action=drop chain=input comment="deny the rest" log=yes log-prefix="\?"
add action=accept chain=icmp comment="echo reply" icmp-options=0:0 protocol=\
    icmp
add action=accept chain=icmp comment="net unreachable" icmp-options=3:0 \
    protocol=icmp
add action=accept chain=icmp comment="host unreachable" icmp-options=3:1 \
    protocol=icmp
add action=accept chain=icmp comment=\
    "host unreachable fragmentation required" icmp-options=3:4 protocol=icmp
add action=accept chain=icmp comment="allow echo request" icmp-options=8:0 \
    protocol=icmp
add action=accept chain=icmp comment="allow parameter bad" icmp-options=12:0 \
    protocol=icmp
add action=accept chain=icmp comment="allow time exceed" icmp-options=11:0 \
    protocol=icmp
add action=drop chain=icmp comment="deny the rest" log=yes log-prefix="\?"
/ip firewall nat
add action=masquerade chain=srcnat out-interface=ether2
/ip route
add comment=utechov disabled=no dst-address=192.168.184.0/24 gateway=\
    wireguard1 routing-table=main suppress-hw-offload=no
add comment=slatina disabled=no dst-address=192.168.188.0/24 gateway=\
    wireguard1 routing-table=main suppress-hw-offload=no
/ip service
set telnet disabled=yes
set ftp disabled=yes
set www disabled=yes
set ssh disabled=yes
set api disabled=yes
set api-ssl disabled=yes
/routing bfd configuration
add disabled=no
/system logging
add action=email disabled=yes topics=info,firewall
add action=email topics=error
add action=email topics=critical
add action=email topics=warning
/system note
set show-at-login=no
/system ntp client
set enabled=yes
/system ntp client servers
add address=time.google.com
add address=tik.cesnet.cz
add address=tak.cesnet.cz
/system routerboard settings
set boot-os=router-os
/tool sniffer
set filter-interface=allinall-bridge filter-operator-between-entries=and \
    memory-scroll=no
```

Main router on site B:

