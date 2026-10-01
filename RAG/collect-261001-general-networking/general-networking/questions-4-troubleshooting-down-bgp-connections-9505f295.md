---
id: collect-261001-general-networking/general-networking/questions-4-troubleshooting-down-bgp-connections-9505f295
title: "questions-4-troubleshooting-down-bgp-connections-9505f295"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["ethernet", "full-duplex"]
source: docs/RAG/collect-261001-general-networking/questions-4-troubleshooting-down-bgp-connections-9505f295.md
source_anchor: ""
source_lines: [1, 48]
sha256: 735ac6e97f9d89e6c130984e30e7adc424331411ed18a8ddbb473a5a5a788d5d
---

# questions-4-troubleshooting-down-bgp-connections-9505f295

Our network experienced a short outage when one of our BGP routes went down for a short time yesterday. Thankfully our connections failed over to our secondary BGP route after a few minutes, and the primary route became operational after a shut/no shut on the ISP side.
We're running 2 stacked (backplane) Cisco 3750e switches running iOS 12.2 58.
In my conversation with our ISP, they couldn't give any definitive answers to the cause. Is there anything that we can do to pinpoint the cause on our end to avoid this issue in the future?
Log at the time of error
172258: May  6 14:43:06: %BGP-5-ADJCHANGE: neighbor xxx.xxx.12.34 Down BGP Notification sent
172259: May  6 14:43:06: %BGP-3-NOTIFICATION: sent to neighbor xxx.xxx.12.34 4/0 (hold time expired) 0 bytes
172260: May  6 14:43:06: %BGP_SESSION-5-ADJCHANGE: neighbor xxx.xxx.12.34 IPv4 Multicast topology base removed from session  BGP Notification sent
172261: May  6 14:43:06: %BGP_SESSION-5-ADJCHANGE: neighbor xxx.xxx.12.34 IPv4 Unicast topology base removed from session  BGP Notification sent
Log when ISP did a shut/no shut to reset BGP on their side
172542: May  6 15:04:15: %LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet2/0/49, changed state to down
172543: May  6 15:04:16: %LINK-3-UPDOWN: Interface GigabitEthernet2/0/49, changed state to down
172544: May  6 15:04:16: %PIM-5-NBRCHG: neighbor xxx.xxx.12.34 DOWN on interface GigabitEthernet2/0/49 non DR
172545: May  6 15:04:16: %PIM-5-NBRCHG: neighbor xxx.xxx.12.34 UP on interface GigabitEthernet2/0/49 
172546: May  6 15:04:16: %PIM-5-DRCHG: DR change from neighbor 0.0.0.0 to xxx.xxx.12.35 on interface GigabitEthernet2/0/49
172547: May  6 15:04:18: %LINK-3-UPDOWN: Interface GigabitEthernet2/0/49, changed state to up
172548: May  6 15:04:19: %LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet2/0/49, changed state to up
Log when the BGP connection finally went from idle to Up
172828: May  6 15:27:33: %BGP-5-ADJCHANGE: neighbor xxx.xxx.12.34 Up
BGP interface on our end (note: no CRC, drops, collisions reported...)
GigabitEthernet2/0/49 is up, line protocol is up (connected)
Hardware is Gigabit Ethernet, address is xxxx.xxxx
Internet address is xxx.xxx.12.35/31
MTU 1500 bytes, BW 1000000 Kbit/sec, DLY 10 usec,
reliability 255/255, txload 1/255, rxload 3/255
Encapsulation ARPA, loopback not set
Keepalive not set
Full-duplex, 1000Mb/s, link type is auto, media type is 1000BaseLX SFP
input flow-control is off, output flow-control is unsupported
ARP type: ARPA, ARP Timeout 04:00:00
Last input 00:00:09, output 00:00:12, output hang never
Last clearing of "show interface" counters never
Input queue: 0/75/52/0 (size/max/drops/flushes); Total output drops: 0
Queueing strategy: fifo
Output queue: 0/40 (size/max)
5 minute input rate 14536000 bits/sec, 1655 packets/sec
5 minute output rate 1010000 bits/sec, 640 packets/sec
413176726 packets input, 428902543141 bytes, 0 no buffer
Received 143495 broadcasts (0 IP multicasts)
0 runts, 0 giants, 0 throttles
0 input errors, 0 CRC, 0 frame, 0 overrun, 0 ignored
0 watchdog, 139275 multicast, 0 pause input
0 input packets with dribble condition detected
125748632 packets output, 42915625632 bytes, 0 underruns
0 output errors, 0 collisions, 0 interface resets
0 unknown protocol drops
0 babbles, 0 late collision, 0 deferred
0 lost carrier, 0 no carrier, 0 pause output
0 output buffer failures, 0 output buffers swapped out
