---
id: collect-260926-mikrotik/mikrotik/questions-1051669-how-to-access-subnet-from-open-vpn-server-ubuntu-vpn-client-ro-cc7d9371
title: "questions-1051669-how-to-access-subnet-from-open-vpn-server-ubuntu-vpn-client-ro-cc7d9371"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vpn/questions-1051669-how-to-access-subnet-from-open-vpn-server-ubuntu-vpn-client-ro-cc7d9371.md
source_anchor: ""
source_lines: [1, 43]
sha256: e8494938ccb81dcbc41a57a5bc66786984e432d0fc465f3570ec005b647fa0f8
---

# questions-1051669-how-to-access-subnet-from-open-vpn-server-ubuntu-vpn-client-ro-cc7d9371

Here's the setup:
[OpenVPN server]    --- WAN --- [RouterOS client]   --- [Local subnet client]
10.5.0.0                        10.5.0.14               10.10.10.2
$ route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         123.45.128.1    0.0.0.0         UG    0      0        0 eth0
10.5.0.0        10.5.0.2        255.255.255.0   UG    0      0        0 tun0
10.5.0.2        0.0.0.0         255.255.255.255 UH    0      0        0 tun0
10.18.0.0       0.0.0.0         255.255.0.0     U     0      0        0 eth0
123.45.128.0    0.0.0.0         255.255.240.0   U     0      0        0 eth0
$ ifconfig
eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 123.45.67.89  netmask 255.255.240.0  broadcast 123.45.123.255
        inet6 fe80::b47f:49ff:fe42:8567  prefixlen 64  scopeid 0x20<link>
        ether b6:7f:49:42:85:67  txqueuelen 1000  (Ethernet)
        RX packets 106602  bytes 58005294 (58.0 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 129890  bytes 57533013 (57.5 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 296  bytes 89008 (89.0 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 296  bytes 89008 (89.0 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
tun0: flags=4305<UP,POINTOPOINT,RUNNING,NOARP,MULTICAST>  mtu 1500
        inet 10.5.0.1  netmask 255.255.255.255  destination 10.5.0.2
        inet6 fe80::8d8b:2759:671:2327  prefixlen 64  scopeid 0x20<link>
        unspec 00-00-00-00-00-00-00-00-00-00-00-00-00-00-00-00  txqueuelen 100  (UNSPEC)
        RX packets 1015  bytes 69353 (69.3 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 2248  bytes 325055 (325.0 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
$sudo ufw status
OpenSSH                    ALLOW       Anywhere
OpenSSH (v6)               ALLOW       Anywhere (v6)
$ ping 10.5.0.14
reachable, 0% packet loss
❔ How can I make route from [OpenVPN server] to [Local subnet]?
e.g. ping 10.10.10.2
