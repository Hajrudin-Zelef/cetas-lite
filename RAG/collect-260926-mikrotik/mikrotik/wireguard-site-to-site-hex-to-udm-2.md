---
id: collect-260926-mikrotik/mikrotik/wireguard-site-to-site-hex-to-udm-2
title: "UDM Pro Site C"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-site-to-site-hex-to-udm.md
source_anchor: ""
source_lines: [110, 188]
sha256: 5e18feee0cb9937ac6274713e32607d7f7396d5e374757934bb6233d40ffa5fd
---

# UDM Pro Site C

```
# Hex site A
/interface wireguard
add listen-port=51820 mtu=1420 name=212-Wireguard
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1 list=WAN
add interface=212-Wireguard list=LAN
add interface=212-Wireguard list=WAN
/interface wireguard peers
add allowed-address=192.168.0.0/24 comment=355 endpoint-address=\
    SITE-C.dyndns.org endpoint-port=51820 interface=212-Wireguard \
    persistent-keepalive=1h47m44s public-key=\
    "4HEOBxxxxxx"
add allowed-address=192.168.88.0/24 comment=371 \
    endpoint-address=SITE-B.dydns.org endpoint-port=51820 interface=\
    212-Wireguard persistent-keepalive=30m public-key=\
    "zoZtxxxxxxxx"
/ip address
add address=192.168.2.2/24 comment=defconf interface=bridge network=\
    192.168.2.0
add address=192.168.30.2/24 interface=ether3 network=192.168.30.0
add address=10.10.10.1/24 interface=212-Wireguard network=10.10.10.0
/ip firewall address-list
add address=jrs212.dyndns.org list=WAN
add address=192.168.2.0/24 list=LAN
/ip firewall filter
add action=accept chain=input comment="NEW defconf: accept ICMP" protocol=\
    icmp
add action=accept chain=input in-interface=212-Wireguard log=yes
add action=accept chain=forward log=yes out-interface=212-Wireguard
add action=accept chain=input log=yes protocol=udp src-port=51820
add action=accept chain=forward in-interface=212-Wireguard
add action=accept chain=forward dst-address=192.168.2.0/24 in-interface=\
    212-Wireguard log=yes
add action=accept chain=input comment=\
    "NEW defconf: accept established,related,untracked" connection-state=\
    established,related,untracked
add action=drop chain=input comment="NEW defconf: drop invalid" \
    connection-state=invalid
add action=accept chain=input comment=NEW in-interface-list=LAN
add action=drop chain=input comment="NEW drop all else"
add action=fasttrack-connection chain=forward comment=\
    "NEW defconf: fasttrack" connection-state=established,related hw-offload=\
    yes
add action=accept chain=forward comment=\
    "NEW defconf: accept established,related, untracked" connection-state=\
    established,related,untracked
add action=accept chain=forward comment="NEW allow port forwarding" \
    connection-nat-state=dstnat log=yes
add action=accept chain=forward in-interface-list=LAN out-interface-list=WAN
add action=drop chain=forward comment="NEW defconf: drop invalid" \
    connection-state=invalid
add action=drop chain=forward comment=NEW
/ip firewall mangle
add action=mark-connection chain=prerouting comment=\
    "Mark connection for hairpin NAT" dst-address-list=WAN \
    new-connection-mark="Hairpin NAT" passthrough=yes src-address-list=LAN
/ip firewall nat
add action=masquerade chain=srcnat comment="Hairpin NAT" connection-mark=\
    "Hairpin NAT"
add action=masquerade chain=srcnat comment="NEW defconf: masquerade" \
    out-interface-list=WAN
add action=dst-nat chain=dstnat dst-address=192.168.2.176 dst-port=8123 log=\
    yes protocol=tcp to-addresses=192.168.2.176
add action=src-nat chain=srcnat comment="new 8123" disabled=yes dst-address=\
    192.168.2.176 dst-port=8123 protocol=tcp to-addresses=192.168.2.176
add action=src-nat chain=srcnat comment="new 5800" disabled=yes dst-port=5800 \
    protocol=tcp to-addresses=192.168.2.22
add action=src-nat chain=srcnat comment="new 5900" disabled=yes dst-port=5900 \
    protocol=tcp to-addresses=192.168.2.22
add action=dst-nat chain=dstnat comment="PORT FWD:  8123" dst-address-list=\
    WAN dst-port=8123 protocol=tcp to-addresses=192.168.2.176 to-ports=8123
/ip route
add disabled=yes dst-address=0.0.0.0/0 gateway=192.168.2.1
add disabled=no distance=1 dst-address=192.168.88.0/24 gateway=212-Wireguard \
    pref-src="" routing-table=main scope=30 suppress-hw-offload=no \
    target-scope=10
```
