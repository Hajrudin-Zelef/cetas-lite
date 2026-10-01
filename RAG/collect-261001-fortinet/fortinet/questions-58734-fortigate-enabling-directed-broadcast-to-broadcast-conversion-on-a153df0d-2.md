---
id: collect-261001-fortinet/fortinet/questions-58734-fortigate-enabling-directed-broadcast-to-broadcast-conversion-on-a153df0d-2
title: "questions-58734-fortigate-enabling-directed-broadcast-to-broadcast-conversion-on-a153df0d"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2020-07-21"]
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-58734-fortigate-enabling-directed-broadcast-to-broadcast-conversion-on-a153df0d.md
source_anchor: ""
source_lines: [47, 87]
sha256: e7028d2eb2d2a736e76ab1747a2549105670e8ac696acde55f0cacf4a81095e5
---

# questions-58734-fortigate-enabling-directed-broadcast-to-broadcast-conversion-on-a153df0d

So at least, something is happening. But these packets are (at layer 2) not real broadcasts, but they're being sent to DstMac 00:00:00:00:00:00 (where I'd expect ff:ff:ff:ff:ff:ff)
filters=[host 192.168.115.1]
4.492989 WAN1 in 192.168.115.1 -> 192.168.10.255: icmp: echo request
0x0000   0000 0000 0001 0000 0000 0000 0800 4500        ..............E.
0x0010   0054 2614 0000 4001 5544 c0a8 7301 c0a8        .T&...@.UD..s...
0x0020   0aff 0800 8984 5700 0000 13a2 c25c 0000        ......W......\..
0x0030   0000 3a7c 0700 0000 0000 0000 0000 0000        ..:|............
0x0040   0000 0000 0000 0000 0000 0000 0000 0000        ................
0x0050   0000 0000 0000 0000 0000 0000 0000 0000        ................
0x0060   0000                                           ..
4.493277 internal1 out 192.168.115.1 -> 192.168.10.255: icmp: echo request
0x0000   0000 0000 0000 0009 0f09 0b01 0800 4500        ..............E.
0x0010   0054 2614 0000 3f01 5644 c0a8 7301 c0a8        .T&...?.VD..s...
0x0020   0aff 0800 8984 5700 0000 13a2 c25c 0000        ......W......\..
0x0030   0000 3a7c 0700 0000 0000 0000 0000 0000        ..:|............
0x0040   0000 0000 0000 0000 0000 0000 0000 0000        ................
0x0050   0000 0000 0000 0000 0000 0000 0000 0000        ................
0x0060   0000                                           ..
This behaviour is seen with or without any of the multicast config bits in place, and with or without the narrow unicast firewall policy.
Adding set broadcast-forward enable to the egress interface does not change the DstMAC address being used in the egress packet. Still, some systems on the local subnet seem to react to DstMAC 00:00:00:00:00:00 and send their ping replies. these of course are out-of-state to the firewall and get dropped - no harm in that.
I'll have the server team try WoL with the given configuration - if that won't work, we'll try setting a static ARP entry mapping 192.168.10.255 to ff:ff:ff:ff:ff:ff
[/ADDON-1]:
[ADDON-2 2020-07-21]:
Yes, it took a while for the Systems Managament people to get back to the topic and eventually find some time to send some WoL Magic Packets down the WAN.
Testing was done on a Fortigate 100E with FortiOS 6.0.8.
See Lukas' answer below for a config example. No form of broadcast-forward enable was needed. I am aware that zac67's answer says the same, but includes broadcast-forward enable.
This is what the directed broadcast looked like when it left the FG100 into the given LAN/Subnet. You'll note the proper broadcast destination address (ffff.ffff.ffff).
87.273480 port7 -- 192.168.35.9.54525 -> 192.168.37.255.7: udp 102
0x0000   ffff ffff ffff 0009 0f09 2012 0800 4500        ..............E.
0x0010   0082 185b 0000 7e11 59b7 c0a8 2309 c0a8        ...[..~.Y...#...
0x0020   25ff d4fd 0007 006e 5494 ffff ffff ffff        %......nT.......
0x0030   f8b4 6a25 9dd7 f8b4 6a25 9dd7 f8b4 6a25        ..j%....j%....j%
0x0040   9dd7 f8b4 6a25 9dd7 f8b4 6a25 9dd7 f8b4        ....j%....j%....
0x0050   6a25 9dd7 f8b4 6a25 9dd7 f8b4 6a25 9dd7        j%....j%....j%..
0x0060   f8b4 6a25 9dd7 f8b4 6a25 9dd7 f8b4 6a25        ..j%....j%....j%
0x0070   9dd7 f8b4 6a25 9dd7 f8b4 6a25 9dd7 f8b4        ....j%....j%....
0x0080   6a25 9dd7 f8b4 6a25 9dd7 f8b4 6a25 9dd7        j%....j%....j%..
Hint: the FG100E showed similar behaviour as the FG60E from earlier tests.
With diag sniffer packet any <someFilterString>, the destination MAC was shown as 0000.0000.0000, but diag sniffer packet port7 <someFilterString> showed ffff.ffff.ffff. That's not quite what one would expect, and extends troubleshooting unnecessarily.
[/ADDON-2]:
index=8 ifname=internal1 255.255.255.255 ff:ff:ff:ff:ff:ff state=00000040 ...exec ping 192.168.1.255), the packets are leaving the box with DstMAc addressff:ff:ff:ff:ff:ff.
