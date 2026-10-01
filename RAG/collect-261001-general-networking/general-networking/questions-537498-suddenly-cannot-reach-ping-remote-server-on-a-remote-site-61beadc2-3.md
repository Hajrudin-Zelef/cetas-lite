---
id: collect-261001-general-networking/general-networking/questions-537498-suddenly-cannot-reach-ping-remote-server-on-a-remote-site-61beadc2-3
title: "questions-537498-suddenly-cannot-reach-ping-remote-server-on-a-remote-site-61beadc2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-537498-suddenly-cannot-reach-ping-remote-server-on-a-remote-site-61beadc2.md
source_anchor: ""
source_lines: [117, 145]
sha256: cb0e0ae4772490aa49f966ef778fd2f4a14a308c83c14c58e49bc59c0cac1741
---

# questions-537498-suddenly-cannot-reach-ping-remote-server-on-a-remote-site-61beadc2

tab=255 vf=0 scope=254 type=2 proto=2 prio=0 0.0.0.0/0.0.0.0/0->64.15.130.52/32 pref=64.15.130.56 gwy=0.0.0.0 dev=3(wan1)
tab=255 vf=0 scope=253 type=3 proto=2 prio=0 0.0.0.0/0.0.0.0/0->10.10.10.255/32 pref=10.10.10.1 gwy=0.0.0.0 dev=4(dmz)
tab=255 vf=0 scope=253 type=3 proto=2 prio=0 0.0.0.0/0.0.0.0/0->127.0.0.0/32 pref=127.0.0.1 gwy=0.0.0.0 dev=7(root)
tab=255 vf=0 scope=254 type=2 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.254.254/32 pref=192.168.254.254 gwy=0.0.0.0 dev=5(internal)
tab=255 vf=0 scope=253 type=3 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.250.255/32 pref=192.168.250.254 gwy=0.0.0.0 dev=5(internal)
tab=255 vf=0 scope=254 type=2 proto=2 prio=0 0.0.0.0/0.0.0.0/0->127.0.0.1/32 pref=127.0.0.1 gwy=0.0.0.0 dev=7(root)
tab=255 vf=0 scope=253 type=3 proto=2 prio=0 0.0.0.0/0.0.0.0/0->64.15.130.48/32 pref=64.15.130.56 gwy=0.0.0.0 dev=3(wan1)
tab=255 vf=0 scope=254 type=2 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.250.254/32 pref=192.168.250.254 gwy=0.0.0.0 dev=5(internal)
tab=255 vf=0 scope=253 type=3 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.254.255/32 pref=192.168.254.254 gwy=0.0.0.0 dev=5(internal)
tab=255 vf=0 scope=254 type=2 proto=2 prio=0 0.0.0.0/0.0.0.0/0->127.0.0.0/8 pref=127.0.0.1 gwy=0.0.0.0 dev=7(root)
PING succesfull on a server
diagnose sniffer packet any "host 192.168.253.23" 4
23.232067 internal in 192.168.254.15 -> 192.168.253.23: icmp: echo request
23.232329 HOCH-KAN out 192.168.254.15 -> 192.168.253.23: icmp: echo request
23.248800 HOCH-KAN in 192.168.253.23 -> 192.168.254.15: icmp: echo reply
23.248932 internal out 192.168.253.23 -> 192.168.254.15: icmp: echo reply
PING failed on a server
diagnose sniffer packet any "host 192.168.253.18" 4
8.212249 internal in 192.168.254.15 -> 192.168.253.18: icmp: echo request
8.212479 wan1 out 64.15.130.56 -> 192.168.253.18: icmp: echo request
10.508155 internal in 192.168.254.15.1113 -> 192.168.253.18.139: syn 1271941747
10.508436 wan1 out 64.15.130.56.42334 -> 192.168.253.18.139: syn 1271941747
11.706287 internal in 192.168.254.15.1112 -> 192.168.253.18.445: syn 341420858
11.706540 wan1 out 64.15.130.56.42332 -> 192.168.253.18.445: syn 341420858
Why the route taken is different for the server on the same network ? I don't use any RIP, OSPF, BGP routing. No policy routing. Juste a static route between VPNs. Nothing is showing a dynamic route for 192.168.253.23 and the Fortigate decide to route it into the wan1 interface instead.
Anything I could check next time it happens ?
Thank in advance
And sorry if is not fully clear, french is my mother language
S.
