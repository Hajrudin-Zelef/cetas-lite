---
id: collect-261001-general-networking/general-networking/questions-2147-can-one-configure-several-static-eigrp-neighbors-on-one-interface-a1547a1b
title: "questions-2147-can-one-configure-several-static-eigrp-neighbors-on-one-interface-a1547a1b"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2021-01-03"]
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-2147-can-one-configure-several-static-eigrp-neighbors-on-one-interface-a1547a1b.md
source_anchor: ""
source_lines: [1, 112]
sha256: c64ce193a18f985eb69c7b7cd99ebcffc0affe1caa5c61dd0294c923131bbd49
---

# questions-2147-can-one-configure-several-static-eigrp-neighbors-on-one-interface-a1547a1b

Can one configure several static EIGRP Neighbors on one interface and are there any considerations apart from disabling split-horizon that one should consider?
- 
        what's the consideration for using static neighbor instead of multicast discovery?sdaffa23fdsf– sdaffa23fdsf2013-06-29 07:33:32 +00:00Commented Jun 29, 2013 at 7:33
- 
            
            
- 
        2in some networks, like nonbroadcast multiaccess networks (NBMA) like frame relay you are not able to discover neighbor, so you use neighbor command in router eigrp config modeMilan Prpic– Milan Prpic2013-06-29 11:21:20 +00:00Commented Jun 29, 2013 at 11:21
- 
        1in addition to this comment there may also be concrete reasons concering security and specific loop prevention scenarios.Max– Max2013-07-01 08:56:16 +00:00Commented Jul 1, 2013 at 8:56
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you can post and accept your own answer.Ron Maupin– Ron Maupin ♦2021-01-03 05:52:57 +00:00Commented Jan 3, 2021 at 5:52
2 Answers 2
Yes, that works fine.
R1#sh run | s eigrp
router eigrp 100
network 123.123.123.0 0.0.0.255
no auto-summary
neighbor 123.123.123.3 FastEthernet0/0
neighbor 123.123.123.2 FastEthernet0/0
R1#sh ip eigrp nei
IP-EIGRP neighbors for process 100
H   Address                 Interface       Hold Uptime   SRTT   RTO  Q  Seq
                                            (sec)         (ms)       Cnt Num
1   123.123.123.3           Fa0/0             10 00:01:04   47   282  0  6
0   123.123.123.2           Fa0/0             10 00:01:21   49   294  0  6
R1#sh ip eigrp nei detail 
IP-EIGRP neighbors for process 100
H   Address                 Interface       Hold Uptime   SRTT   RTO  Q  Seq
                                            (sec)         (ms)       Cnt Num
1   123.123.123.3           Fa0/0             14 00:01:41   47   282  0  6
   Static neighbor
   Restart time 00:01:13
   Version 12.4/1.2, Retrans: 1, Retries: 0
0   123.123.123.2           Fa0/0             10 00:01:57   49   294  0  6
   Static neighbor
   Restart time 00:01:13
   Version 12.4/1.2, Retrans: 0, Retries: 0
R1#debug eigrp packets hello
EIGRP Packets debugging is on
    (HELLO)
R1#
EIGRP: Received HELLO on FastEthernet0/0 nbr 123.123.123.3
AS 100, Flags 0x0, Seq 0/0 idbQ 0/0 iidbQ un/rely 0/0 peerQ un/rely 0/0
EIGRP: Sending HELLO on FastEthernet0/0 nbr 123.123.123.3
AS 100, Flags 0x0, Seq 0/0 idbQ 0/0 iidbQ un/rely 0/0
EIGRP: Sending HELLO on FastEthernet0/0 nbr 123.123.123.2
AS 100, Flags 0x0, Seq 0/0 idbQ 0/0 iidbQ un/rely 0/0
R1#
EIGRP: Received HELLO on FastEthernet0/0 nbr 123.123.123.2
AS 100, Flags 0x0, Seq 0/0 idbQ 0/0 iidbQ un/rely 0/0 peerQ un/rely 0/0
you may also want to consider adding 'no ip next-hop-self eigrp X' to the interface pointing to those static neighbors. This will enable "Third Party Next-Hop", which will allow the spoke routers to send traffic directly to each other instead of via hub. Assuming R1<->R2 and R1<->R3 (loop0 - 3.3.3.3)
 Before:
R2#sh ip eigrp topology 3.3.3.3 255.255.255.255
IP-EIGRP (AS 1): Topology entry for 3.3.3.3/32
  State is Passive, Query origin flag is 1, 1 Successor(s), FD is 2560051456
  Routing Descriptor Blocks:
  10.0.0.1 (FastEthernet0/0), from 10.0.0.1, Send flag is 0x0
      Composite metric is (2560051456/2560025856), Route is External
      Vector metric:
        Minimum bandwidth is 1 Kbit
        Total delay is 2010 microseconds
        Reliability is 1/255
        Load is 1/255
        Minimum MTU is 1
        Hop count is 2
      External data:
        Originating router is 3.3.3.3  
        AS number of route is 0
        External protocol is Connected, external metric is 0
        Administrator tag is 0 (0x00000000)
R2#sh ip route 3.3.3.3
Routing entry for 3.3.3.3/32
  Known via "eigrp 1", distance 170, metric 2560051456, type external
  Redistributing via eigrp 1
  Last update from 10.0.0.1 on FastEthernet0/0, 00:04:04 ago
  Routing Descriptor Blocks:
  * 10.0.0.1, from 10.0.0.1, 00:04:04 ago, via FastEthernet0/0
      Route metric is 2560051456, traffic share count is 1
      Total delay is 2010 microseconds, minimum bandwidth is 1 Kbit
      Reliability 1/255, minimum MTU 1 bytes
      Loading 1/255, Hops 2
After:
R2#sh ip route 3.3.3.3
Routing entry for 3.3.3.3/32
  Known via "eigrp 1", distance 170, metric 2560051456, type external
  Redistributing via eigrp 1
  Last update from 10.0.0.3 on FastEthernet0/0, 00:00:14 ago
  Routing Descriptor Blocks:
  * 10.0.0.3, from 10.0.0.1, 00:00:14 ago, via FastEthernet0/0
      Route metric is 2560051456, traffic share count is 1
      Total delay is 2010 microseconds, minimum bandwidth is 1 Kbit
      Reliability 1/255, minimum MTU 1 bytes
      Loading 1/255, Hops 2
R2#sh ip eigrp topology 3.3.3.3 255.255.255.255
IP-EIGRP (AS 1): Topology entry for 3.3.3.3/32
  State is Passive, Query origin flag is 1, 1 Successor(s), FD is 2560051456
  Routing Descriptor Blocks:
  10.0.0.3 (FastEthernet0/0), from 10.0.0.1, Send flag is 0x0
      Composite metric is (2560051456/2560025856), Route is External
      Vector metric:
        Minimum bandwidth is 1 Kbit
        Total delay is 2010 microseconds
        Reliability is 1/255
        Load is 1/255
        Minimum MTU is 1
        Hop count is 2
      External data:
        Originating router is 3.3.3.3
        AS number of route is 0
        External protocol is Connected, external metric is 0
        Administrator tag is 0 (0x00000000)
