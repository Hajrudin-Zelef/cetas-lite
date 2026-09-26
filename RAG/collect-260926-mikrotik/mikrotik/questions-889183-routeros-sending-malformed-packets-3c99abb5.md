---
id: collect-260926-mikrotik/mikrotik/questions-889183-routeros-sending-malformed-packets-3c99abb5
title: "questions-889183-routeros-sending-malformed-packets-3c99abb5"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-889183-routeros-sending-malformed-packets-3c99abb5.md
source_anchor: ""
source_lines: [1, 37]
sha256: 81474cc89cc66f3d7ff50db47d7f2d83872f4e8db2f7346c538ef7c96c5579b9
---

# questions-889183-routeros-sending-malformed-packets-3c99abb5

I have a Routerboard 951G-2HnD which runs latest stable:
RouterOS v6.50.5
Firmware v3.41
The board acts as the WiFi AP in WPA2-PSK mode. Recently I have noticed unhealthy amount of traffic sent from AP to clients for no good reason. Example of Wireshark capture:
    Frame 11: 14 bytes on wire (112 bits), 14 bytes captured (112 bits) on interface 0
    Interface id: 0 (wlan0)
    Encapsulation type: Ethernet (1)
    Arrival Time: XX XX, 2017 18:24:01 XXX
    [Time shift for this packet: 0.000000000 seconds]
    Epoch Time: XX.XX seconds
    [Time delta from previous captured frame: 0.073480897 seconds]
    [Time delta from previous displayed frame: 0.073480897 seconds]
    [Time since reference or first frame: 14.088600317 seconds]
    Frame Number: 11
    Frame Length: 14 bytes (112 bits)
    Capture Length: 14 bytes (112 bits)
    [Frame is marked: False]
    [Frame is ignored: False]
    [Protocols in frame: eth:ethertype:ip]
Ethernet II, Src: Routerbo_XX:XX:cd (XX:XX:XX:XX:XX:cd), Dst: Tp-LinkT_XX:XX:d1 (XX:XX:XX:XX:XX:d1)
    Destination: Tp-LinkT_XX:XX:d1 (XX:XX:XX:XX:XX:d1)
        Address: Tp-LinkT_XX:XX:d1 (XX:XX:XX:XX:XX:d1)
        .... ..0. .... .... .... .... = LG bit: Globally unique address (factory default)
        .... ...0 .... .... .... .... = IG bit: Individual address (unicast)
    Source: Routerbo_55:43:cd (XX:XX:XX:XX:XX:cd)
        Address: Routerbo_55:43:cd (XX:XX:XX:XX:XX:cd)
        .... ..0. .... .... .... .... = LG bit: Globally unique address (factory default)
        .... ...0 .... .... .... .... = IG bit: Individual address (unicast)
    Type: IPv4 (0x0800)
[Malformed Packet: IPv4]
    [Expert Info (Error/Malformed): Malformed Packet (Exception occurred)]
        [Malformed Packet (Exception occurred)]
        [Severity level: Error]
        [Group: Malformed]
Packet:
0000  98 de d0 15 2a d1 e4 8d 8c 55 43 cd 08 00         ....*....UC...
What it could be?
