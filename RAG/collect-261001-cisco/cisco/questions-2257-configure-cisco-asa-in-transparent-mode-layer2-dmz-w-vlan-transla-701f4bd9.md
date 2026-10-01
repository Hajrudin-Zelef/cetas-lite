---
id: collect-261001-cisco/cisco/questions-2257-configure-cisco-asa-in-transparent-mode-layer2-dmz-w-vlan-transla-701f4bd9
title: "questions-2257-configure-cisco-asa-in-transparent-mode-layer2-dmz-w-vlan-transla-701f4bd9"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2013-02-19"]
keywords: ["ethernet", "research"]
source: docs/RAG/collect-261001-cisco/questions-2257-configure-cisco-asa-in-transparent-mode-layer2-dmz-w-vlan-transla-701f4bd9.md
source_anchor: ""
source_lines: [1, 133]
sha256: d7e55a917449599ad2b8e898bc7c8f65de785f224f57d4793aa85c5c15a6705a
---

# questions-2257-configure-cisco-asa-in-transparent-mode-layer2-dmz-w-vlan-transla-701f4bd9

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
15
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am in the middle of a project to migrate some existing switched ethernet dot1q trunks behind an ASA firewall... these trunks have five vlans each (numbered 51 - 55). This is a simple drawing of the original layer2 service...
One of the requirements is to have an ASA firewall context per Vlan in the original dot1q trunk. This means I wind up using a BVI to bridge the new INSIDE interface to the DMZ interface in each FW context. Due to other constraints, I wind up with a FW config like this (I am summarizing out all the context stuff to simplify the question)...
Cisco ASAs in transparent mode wind up using two different vlan IDs to connect a single layer2 vlan service. Connect the two vlans through interface BVI1; the bridge-group 1 configuration on each physical interface makes the connection between Vlan51 and Vlan951 in the config above.
Assume ASA:Gi0/2 connects to 4507:Gi1/2... Notice what happens to the DMZ interface... the ASA DMZ Vlan is 951, and this connects via a dot1q trunk to the DMZ switch (the Cat4507). I need to connect D1 to switchport 4507:Gi1/1, but I must deliver the Vlan951-955 services to D1 as dot1q Vlan51-55 on 4507:Gi1/1. In other words, the Vlan BVI wrangling I had to do on the ASA messes up the Vlan numbering in my original service definition.
Sadly, I cannot easily renumber the Vlans on D1. The perfect solution would be to somehow translate Vlan951 on 4507:Gi1/2 to Vlan51 on 4507:Gi1/1. Cisco has a feature called vlan mapping, but it seems to require QinQ... all my services are simple dot1q... the 4500 vlan mapping docs are not clear about how they handle simple dot1q encapsulation.
I know I can translate the vlans in the 4500 via loopback cable, but this burns two extra ports per Vlan... a total of ten extra ports for all the Vlans in the service (v51 - v55).
Question
Refer to the diagram below.
How can I translate all Vlans numbered 95x on the 4507:Gi1/2 dot1q trunk to be numbered Vlan5x on 4507:Gi1/1 dot1q? I need to use the fewest number of ports consumed for "translation overhead". Please include configs for all ports that your answer requires.
I am open to vlan mapping, if someone can explain how it would work in this topology...
I don't have SUP7 to test, but it works on SUP6 and SUP32, I would presume SUP7 retains this functionality.
I've tested between JNPR M320 <-> SUP32, and 'vlan mapping JNPR SUP32' works just fine.
There is no need for QinQ, what the QinQ option does is it adds top tag to one particularly tag. So switchport vlan mapping 1042 dot1q-tunnel 42 would map incoming [1042] stack to [42 1042] stack.
As opposed to switchport vlan mapping 1042 42 which maps incoming dot1q Vlan [1042] to dot1q Vlan [42].
JNPR M320 config:
{master}[edit interfaces ge-0/1/0 unit 1042]
user@m320# show
vlan-id 1042;
family inet {
address 10.42.42.1/24;
}
{master}[edit interfaces ge-0/1/0 unit 1042]
user@m320# run show interfaces ge-0/1/0
Physical interface: ge-0/1/0, Enabled, Physical link is Up
Interface index: 135, SNMP ifIndex: 506
Description: B: SUP32 ge5/1
Link-level type: Flexible-Ethernet, MTU: 9192, Speed: 1000mbps, BPDU Error: None,
MAC-REWRITE Error: None, Loopback: Disabled, Source filtering: Disabled, Flow control: Disabled,
Auto-negotiation: Enabled, Remote fault: Online
Device flags : Present Running
Interface flags: SNMP-Traps Internal: 0x4000
CoS queues : 8 supported, 8 maximum usable queues
Current address: 00:12:1e:d5:90:7f, Hardware address: 00:12:1e:d5:90:7f
Last flapped : 2013-02-19 09:14:29 UTC (19w6d 21:12 ago)
Input rate : 4560 bps (5 pps)
Output rate : 6968 bps (4 pps)
Active alarms : None
Active defects : None
Interface transmit statistics: Disabled
SUP32 config:
SUP32#show run int giga5/1
Building configuration...
Current configuration : 365 bytes
!
interface GigabitEthernet5/1
description F: M320 ge-0/1/0
switchport
switchport trunk encapsulation dot1q
switchport mode trunk
switchport nonegotiate
switchport vlan mapping enable
switchport vlan mapping 1042 42
mtu 9216
bandwidth 1000000
speed nonegotiate
no cdp enable
spanning-tree portfast edge trunk
spanning-tree bpdufilter enable
end
SUP32#show ru int vlan42
Building configuration...
Current configuration : 61 bytes
!
interface Vlan42
ip address 10.42.42.2 255.255.255.0
end
SUP32#sh int GigabitEthernet5/1 vlan mapping
State: enabled
Original VLAN Translated VLAN
------------- ---------------
1042 42
SUP32#sh int vlan42
Vlan42 is up, line protocol is up
Hardware is EtherSVI, address is 0005.ddee.6000 (bia 0005.ddee.6000)
Internet address is 10.42.42.2/24
MTU 1500 bytes, BW 1000000 Kbit, DLY 10 usec,
reliability 255/255, txload 1/255, rxload 1/255
Encapsulation ARPA, loopback not set
Keepalive not supported
ARP type: ARPA, ARP Timeout 04:00:00
Last input 00:00:09, output 00:01:27, output hang never
Last clearing of "show interface" counters never
Input queue: 0/75/0/0 (size/max/drops/flushes); Total output drops: 0
Queueing strategy: fifo
Output queue: 0/40 (size/max)
5 minute input rate 0 bits/sec, 0 packets/sec
5 minute output rate 0 bits/sec, 0 packets/sec
L2 Switched: ucast: 17 pkt, 1920 bytes - mcast: 0 pkt, 0 bytes
L3 in Switched: ucast: 0 pkt, 0 bytes - mcast: 0 pkt, 0 bytes mcast
L3 out Switched: ucast: 0 pkt, 0 bytes mcast: 0 pkt, 0 bytes
38 packets input, 3432 bytes, 0 no buffer
Received 21 broadcasts (0 IP multicasts)
0 runts, 0 giants, 0 throttles
0 input errors, 0 CRC, 0 frame, 0 overrun, 0 ignored
26 packets output, 2420 bytes, 0 underruns
0 output errors, 0 interface resets
0 output buffer failures, 0 output buffers swapped out
And
SUP32#ping 10.42.42.1
Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 10.42.42.1, timeout is 2 seconds:
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 1/1/1 ms
SUP32#sh arp | i 10.42.42.1
Internet 10.42.42.1 12 0012.1ed5.907f ARPA Vlan42
SUP32#show mac address-table dynamic address 0012.1ed5.907f
Legend: * - primary entry
age - seconds since last seen
n/a - not available
vlan mac address type learn age ports
------+----------------+--------+-----+----------+--------------------------
Active Supervisor:
* 450 0012.1ed5.907f dynamic Yes 0 Gi5/1
* 50 0012.1ed5.907f dynamic Yes 0 Gi5/1
* 40 0012.1ed5.907f dynamic Yes 0 Gi5/1
* 42 0012.1ed5.907f dynamic Yes 5 Gi5/1
user@m320# run ping 10.42.42.2 count 2
PING 10.42.42.2 (10.42.42.2): 56 data bytes
64 bytes from 10.42.42.2: icmp_seq=0 ttl=255 time=0.495 ms
64 bytes from 10.42.42.2: icmp_seq=1 ttl=255 time=0.651 ms
--- 10.42.42.2 ping statistics ---
2 packets transmitted, 2 packets received, 0% packet loss
round-trip min/avg/max/stddev = 0.495/0.573/0.651/0.078 ms
{master}[edit interfaces ge-0/1/0 unit 1042]
user@m320# run show arp no-resolve |match 10.42.42.2
00:05:dd:ee:60:00 10.42.42.2 ge-0/1/0.1042 none
Nice thing about the Brocade is that you can convert any tag to another tag, double-tag to another double-tag, double-tag to single-tag, and single-tag to double-tag
