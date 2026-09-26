---
id: collect-260926-mikrotik/mikrotik/questions-1192990-linux-bond0-with-802-3ad-not-receiving-lacp-response-from-mikr-472ab60c
title: "questions-1192990-linux-bond0-with-802-3ad-not-receiving-lacp-response-from-mikr-472ab60c"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1192990-linux-bond0-with-802-3ad-not-receiving-lacp-response-from-mikr-472ab60c.md
source_anchor: ""
source_lines: [1, 66]
sha256: 6a026d613a0ed1fc7a03fe90e0bc1b3f5efce546933c09fef39e3489697cd534
---

# questions-1192990-linux-bond0-with-802-3ad-not-receiving-lacp-response-from-mikr-472ab60c

I’m trying to set up an 802.3ad (LACP) bond on Linux with four interfaces (ens3f0, ens3f1, ens2f0, ens2f1) connected to a MikroTik switch. The bond is configured with
- mode=802.3ad,
- miimon=100,
- lacp_rate=slow
(to match MikroTik’s 30s setting), and xmit_hash_policy=layer2. The bond itself comes up and the links show as 10Gbps full duplex, but only one of the slaves actually becomes active at a time. When I check /proc/net/bonding/bond0, there’s no partner system information at all, which suggests the Linux side isn’t receiving any LACP negotiation from the MikroTik. The kernel logs back this up with warnings like
No 802.3ad response from the link partner.
So right now the bond technically works, but it’s only running on a single interface instead of aggregating all four? This is what i've tried so far. Appreciate any inputs.
The sequence of my commands
sudo ip link delete bond0 //i tried to set up the bond twice
sudo ip link add bond0 type bond
mode 802.3ad
miimon 100
lacp_rate slow
xmit_hash_policy layer2
sudo ip link set ens3f0 down
sudo ip link set ens3f0 master bond0
sudo ip link set ens3f0 up
sudo ip link set ens3f1 down
sudo ip link set ens3f1 master bond0
sudo ip link set ens3f1 up
sudo ip link set ens2f0 down
sudo ip link set ens2f0 master bond0
sudo ip link set ens2f0 up
sudo ip link set ens2f1 down
sudo ip link set ens2f1 master bond0
sudo ip link set ens2f1 up
sudo ip addr add 10.xx.xx.xx/24 dev bond0
sudo ip link set bond0 up
sudo ip route add default via 10.xx.xx.xx
my /proc/net/bonding/bond0 looks something like this:
Bonding Mode: IEEE 802.3ad Dynamic link aggregation
Transmit Hash Policy: layer2
MII Status: up
MII Polling Interval: 100 ms
Up Delay: 0
Down Delay: 0
LACP rate: slow
Bonding State: Active Aggregator
Actor Churn State: none
Partner Churn State: none
Actor Churned Count: 0
Partner Churned Count: 0
Here's my syst log:
Sep 22 23:23:21 storinator kernel: **bond**0: Enslaving ens2f0 as a backup interface with a down link
Sep 22 23:23:21 storinator kernel: **bond**0: Removing slave ens2f0
Sep 22 23:23:21 storinator kernel: **bond**0: Releasing backup interface ens2f0
Sep 22 23:23:49 storinator kernel: **bond**0: Enslaving ens2f1 as a backup interface with a down link
Sep 22 23:23:49 storinator kernel: **bond**0: Removing slave ens2f1
Sep 22 23:23:49 storinator kernel: **bond**0: Releasing backup interface ens2f1
Sep 22 23:25:13 storinator kernel: **bond**0: Enslaving ens3f0 as a backup interface with a down link
Sep 22 23:25:16 storinator kernel: **bond**0: Enslaving ens3f1 as a backup interface with a down link
Sep 22 23:25:16 storinator kernel: **bond**0: Removing slave ens3f1
Sep 22 23:25:16 storinator kernel: **bond**0: Releasing backup interface ens3f1
Sep 22 23:25:17 storinator kernel: **bond**0: link status definitely up for interface ens3f0, 10000 Mbps full duplex
Sep 22 23:25:17 storinator kernel: **bond**0: Warning: No 802.3ad response from the link partner for any adapters in the **bond**
Sep 22 23:25:17 storinator kernel: **bond**0: first active interface up!
on my Mikrotik
[admin@MikroTik] > /interface bonding
[admin@MikroTik] /interface bonding> print
Flags: X - disabled, R - running
0 R name="bonding1" mtu=1500 mac-address=DC:2C:6E:1B:BF:AD arp=enabled arp-timeout=auto slaves=combo3,combo4 mode=802.3ad primary=none link-monitoring=mii arp-interval=100ms arp-ip-targets=""
mii-interval=100ms down-delay=0ms up-delay=0ms lacp-rate=30secs transmit-hash-policy=layer-2 min-links=0
1 R name="bonding2" mtu=1500 mac-address=DC:2C:6E:1B:BF:A3 arp=enabled arp-timeout=auto slaves=ether1,ether2,ether3,ether4 mode=802.3ad primary=none link-monitoring=mii arp-interval=100ms
arp-ip-targets="" mii-interval=100ms down-delay=0ms up-delay=0ms lacp-rate=30secs transmit-hash-policy=layer-2 min-links=0
[admin@MikroTik] /interface bonding>
tcpdump -e -n -i ens2f0? (If not - does a different host connected in its place see them? Could it be that your NIC's firmware eats them, as part of some kind of aggregation offload feature?)
