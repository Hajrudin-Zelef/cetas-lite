---
id: collect-260926-mikrotik/mikrotik/wireguard-site-to-site-routing-with-a-central-server-5
title: "2023-11-08 21:11:31 by RouterOS 7.11.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-site-to-site-routing-with-a-central-server.md
source_anchor: ""
source_lines: [590, 615]
sha256: 1c11288bce9ed4ff3647159857dcc6fd4385971fabfa53b16f74c9a432d8b0ff
---

# 2023-11-08 21:11:31 by RouterOS 7.11.2

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
set time-zone-name=Europe/Prague
/system note
set show-at-login=no
/tool mac-server
set allowed-interface-list=none
```
