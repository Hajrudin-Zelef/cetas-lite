---
id: collect-260926-mikrotik/mikrotik/wireguard-site-to-site-routing-with-a-central-server-1
title: "2023-11-08 21:11:31 by RouterOS 7.11.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-site-to-site-routing-with-a-central-server.md
source_anchor: ""
source_lines: [1, 130]
sha256: 44834081562892b272b158a631371cde6c06a063556560677f4730d6d6733bc1
---

# 2023-11-08 21:11:31 by RouterOS 7.11.2

Beginner’s mistake? I’m new here and the post seemed long enough.

Here are the full scripts:



Central server:

```
/interface ethernet
set [ find default-name=ether1 ] advertise=\
    10M-half,10M-full,100M-half,100M-full,1000M-half,1000M-full
set [ find default-name=ether2 ] advertise=\
    10M-half,10M-full,100M-half,100M-full,1000M-half,1000M-full
/interface wireguard
add listen-port=13231 mtu=1420 name=wireguard1
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/routing bgp template
set default disabled=no output.network=bgp-networks
/routing ospf instance
add disabled=no name=default-v2
/routing ospf area
add disabled=yes instance=default-v2 name=backbone-v2
/snmp community
set [ find default=yes ] addresses=0.0.0.0/0
/interface wireguard peers
add allowed-address=\
    192.168.224.2/32,192.168.8.0/23,192.168.16.0/22,192.168.96.0/22 \
    interface=wireguard1 persistent-keepalive=5s \
    public-key="..."
add allowed-address=192.168.224.16/32 \
    interface=wireguard1 persistent-keepalive=5s public-key=\
    "..."
add allowed-address=192.168.224.97/32 interface=\
    wireguard1 public-key="..."
add allowed-address=192.168.224.3/32,192.168.188.0/24 \
    interface=wireguard1 persistent-keepalive=1m \
    public-key="..."
add allowed-address=192.168.224.4/32,192.168.240.0/24 \
    interface=wireguard1 persistent-keepalive=1m \
    public-key="..."
add allowed-address=192.168.224.5/32,192.168.184.0/24 \
    interface=wireguard1 persistent-keepalive=5s public-key=\
    "..."
/ip address
add address=[PUBLIC_IP]/24 interface=ether1 network=89.221.222.0
add address=192.168.224.1/24 interface=wireguard1 network=192.168.224.0
/ip dns
set servers=46.28.108.2,31.31.72.3
/ip firewall address-list
add address=0.0.0.0/8 comment=RFC6890 list=not_in_internet
add address=172.16.0.0/12 comment=RFC6890 list=not_in_internet
add address=192.168.0.0/16 comment=RFC6890 list=not_in_internet
add address=10.0.0.0/8 comment=RFC6890 list=not_in_internet
add address=169.254.0.0/16 comment=RFC6890 list=not_in_internet
add address=127.0.0.0/8 comment=RFC6890 list=not_in_internet
add address=224.0.0.0/4 comment=Multicast list=not_in_internet
add address=198.18.0.0/15 comment=RFC6890 list=not_in_internet
add address=192.0.0.0/24 comment=RFC6890 list=not_in_internet
add address=192.0.2.0/24 comment=RFC6890 list=not_in_internet
add address=198.51.100.0/24 comment=RFC6890 list=not_in_internet
add address=203.0.113.0/24 comment=RFC6890 list=not_in_internet
add address=100.64.0.0/10 comment=RFC6890 list=not_in_internet
add address=240.0.0.0/4 comment=RFC6890 list=not_in_internet
add address=192.88.99.0/24 comment=RFC3068 list=not_in_internet
/ip firewall filter
add action=accept chain=input comment="accept established,related" \
    connection-state=established,related
add action=drop chain=input comment="drop invalid" connection-state=invalid
add action=jump chain=input comment="own chain for ICMP" jump-target=icmp \
    protocol=icmp
add action=accept chain=input comment="WinBox from WireGuard" dst-port=8291 \
    in-interface=wireguard1 protocol=tcp src-address=192.168.16.0/22
add action=accept chain=input comment="Bandwith test server" dst-port=2000 \
    in-interface=wireguard1 protocol=tcp
add action=accept chain=input comment="WireGuard server" dst-port=13231 \
    protocol=udp
add action=accept chain=input comment="WireGuard traffic" src-address=\
    192.168.224.0/24
add action=drop chain=input
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
/ip route
add disabled=no dst-address=0.0.0.0/0 gateway=89.221.222.1
add disabled=no distance=1 dst-address=192.168.16.0/22 gateway=wireguard1 \
    pref-src="" routing-table=main scope=30 suppress-hw-offload=no \
    target-scope=10
add disabled=no distance=1 dst-address=192.168.188.0/24 gateway=wireguard1 \
    pref-src="" routing-table=main scope=30 suppress-hw-offload=no \
    target-scope=10
add disabled=no distance=1 dst-address=192.168.96.0/22 gateway=wireguard1 \
    pref-src="" routing-table=main scope=30 suppress-hw-offload=no \
    target-scope=10
add disabled=no distance=1 dst-address=192.168.8.1/23 gateway=wireguard1 \
    routing-table=main suppress-hw-offload=no
add disabled=no dst-address=192.168.184.0/24 gateway=wireguard1 \
    routing-table=main suppress-hw-offload=no
/ip service
set telnet disabled=yes
set ftp disabled=yes
set www disabled=yes
set ssh disabled=yes
set api disabled=yes
set api-ssl disabled=yes
/system clock
set time-zone-name=Europe/Prague
/system ntp client
set enabled=yes
/system ntp client servers
add address=time.google.com
add address=tik.cesnet.cz
add address=tak.cesnet.cz
```

Main router on Site A:

