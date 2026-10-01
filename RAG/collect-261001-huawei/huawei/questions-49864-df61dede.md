---
id: collect-261001-huawei/huawei/questions-49864-df61dede
title: "questions-49864-df61dede"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2018-04-12"]
keywords: []
source: docs/RAG/collect-261001-huawei/questions-49864-df61dede.md
source_anchor: ""
source_lines: [1, 65]
sha256: 932a7b5898db21e4c92bda036ed2dfd3b622036851976a74e8bc0ad2c8edc456
---

# questions-49864-df61dede

I'm making a test to create a PORT AGGREGATION between my RB750 and Huawei NE20.
Rb750:
Flags: X - disabled, R - running 
 0  R name="bonding1" mtu=1500 mac-address=D4:CA:6D:1D:6F:66 arp=enabled 
      arp-timeout=auto slaves=ether1,ether2 mode=802.3ad primary=none 
      link-monitoring=mii arp-interval=100ms arp-ip-targets="" 
      mii-interval=100ms down-delay=0ms up-delay=0ms lacp-rate=30secs 
      transmit-hash-policy=layer-2-and-3 min-links=0 
NE20:
Eth-Trunk1 current state : UP (ifindex: 73)
Line protocol current state : UP
Link quality grade : GOOD
Description:
Switch Port, PVID :  200, TPID : 8100(Hex), Hash arithmetic : According to flow, Maximal BW: 200Mbps, Current BW: 200Mbps, The Maximum Transmit Unit is 1500
Internet protocol processing : disabled
IP Sending Frames' Format is PKTFMT_ETHNT_2, Hardware address is 487b-6b4a-cc5c
Current system time: 2018-04-12 15:32:33
Physical is ETH_TRUNK
    Last 300 seconds input rate 111 bits/sec, 0 packets/sec
    Last 300 seconds output rate 5291 bits/sec, 8 packets/sec
    Input: 15012450 packets,22684847541 bytes
           14971431 unicast,30620 broadcast,10399 multicast
           10 errors,0 drops
    Output:795001 packets,252032412 bytes
           170137 unicast,540101 broadcast,84763 multicast
           0 errors,0 drops
    Last 300 seconds input utility rate:  0.01%
    Last 300 seconds output utility rate: 0.01%
----------------------------------------------------------
PortName                      Status              Weight
----------------------------------------------------------
GigabitEthernet0/3/28         UP                  1
GigabitEthernet0/3/29         UP                  1
----------------------------------------------------------
The Number of Ports in Trunk : 2
The Number of UP Ports in Trunk : 2
Well, apparently I got to establish a connection between them, another configuration were done.
VLAN 200 set in Eth-Trunk:
   #
    interface Eth-Trunk1
     portswitch
     port link-type access
     port default vlan 200
     mode lacp-static
    #
VLAN 200 set in bonding interface in Rb750
 #    NAME                   MTU ARP             VLAN-ID INTERFACE               
 0 R  VLAN-200              1500 enabled             200 bonding1 
Using a BTest I got this result. You can realize the traffic is comming to one interface.
                         name:  ether1   ether2 bonding1
        rx-packets-per-second:       1       14       15
           rx-bits-per-second:  480bps  8.5kbps  9.0kbps
     fp-rx-packets-per-second:       1       11        6
        fp-rx-bits-per-second:  480bps  5.9kbps  3.8kbps
          rx-drops-per-second:       0        0        0
         rx-errors-per-second:       0        0        0
        tx-packets-per-second:       0    8 191    8 191
           tx-bits-per-second:    0bps 99.4Mbps 99.4Mbps
     fp-tx-packets-per-second:       0    8 171        0
        fp-tx-bits-per-second:    0bps 98.9Mbps     0bps
          tx-drops-per-second:       0        0        0
    tx-queue-drops-per-second:       0      882        0
         tx-errors-per-second:       0        0        0
-- [Q quit|D dump|C-z pause]
Note: I intent to deploy this using CCR1036 - Ne20, establishing 2GB. Rb750 is just to test.
