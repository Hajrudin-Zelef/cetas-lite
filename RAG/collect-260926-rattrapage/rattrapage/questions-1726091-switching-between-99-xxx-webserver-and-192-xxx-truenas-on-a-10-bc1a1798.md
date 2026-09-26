---
id: collect-260926-rattrapage/rattrapage/questions-1726091-switching-between-99-xxx-webserver-and-192-xxx-truenas-on-a-10-bc1a1798
title: "questions-1726091-switching-between-99-xxx-webserver-and-192-xxx-truenas-on-a-10-bc1a1798"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-rattrapage/servers-reviews/questions-1726091-switching-between-99-xxx-webserver-and-192-xxx-truenas-on-a-10-bc1a1798.md
source_anchor: ""
source_lines: [1, 32]
sha256: af5f277ab70a802980379a6efb42d5604b70cb64b30eb0c9801b4ef4e7d6c4f5
---

# questions-1726091-switching-between-99-xxx-webserver-and-192-xxx-truenas-on-a-10-bc1a1798

Iperf3 from my webserver (99.xxx.9) to TrueNAS (192.168.1.80) is only 150Mb/s but they are on the same 10G Mikrotik switch so I expect a few Gb/s. I suspect my routing config on my webserver is not correct but I could be wrong since if I power off my Netgear GS106 there is no connectivity between my webserver and Truenas.
Iperf 192.xxx.80 to 192.xxx.90 is 3Gb/s as expected since old computers on 10G fiber.
Iperf 99.xxx.9 to 99.xxx.10 is 9.4Gb/s as expected since net computers on 10G fiber.
My network:
                          Microtik      NETGEAR        ATT PACE
                           CRS305       GS106          5268AC
Ubuntu 99.xxx.9 -----10G-----+
Ubuntu 99.xxx.10 ----10G-----+------1G----+------1G------+
Ubuntu 192.xxx.80 ---10G-----+            +              +
Ubuntu 192.xxx.90 ---10G-----+            +              +--- internet
                                          +              +
Windows 192.xxx.x -----------------1G-----+              +
                                                         +
Windows 192.xxx.x --------------------1G-----------------+
The AT&T Pace 5268AC does NAT for all the 192.xxx devices to the internet.
The AT&T Pace 5268AC has HTTP/HTTPS pinhole configured for the public IP address 99.xxx.9
The routing for 99.xxx.10 is
$ route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         99.153.71.14    0.0.0.0         UG    100    0        0 enp2s0f0
99.153.71.8     0.0.0.0         255.255.255.248 U     100    0        0 enp2s0f0
192.168.1.0     0.0.0.0         255.255.255.0   U     10     0        0 enp2s0f0
The Microtik is a switch, so should only use mac address. It has factory default config.
Truenas is configured with static route in GUI but this appears to be not propaged to the "route" command.
root@truenas 192.168.1.80 [~]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.1.254   0.0.0.0         UG    0      0        0 enp2s0f0
172.16.0.0      0.0.0.0         255.255.0.0     U     0      0        0 kube-bridge
192.168.1.0     0.0.0.0         255.255.255.0   U     0      0        0 enp2s0f0
I would appreciate help figuring out what I got wrong or diagnosing suggestions.
