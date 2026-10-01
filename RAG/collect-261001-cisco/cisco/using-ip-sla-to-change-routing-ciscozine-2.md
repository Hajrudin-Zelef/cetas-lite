---
id: collect-261001-cisco/cisco/using-ip-sla-to-change-routing-ciscozine-2
title: "using-ip-sla-to-change-routing-ciscozine"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/using-ip-sla-to-change-routing-ciscozine.md
source_anchor: ""
source_lines: [139, 239]
sha256: 16313a9cc985032dca2cdb844e4f2b611b72d5c2264599645cb72a069f6bad5f
---

# using-ip-sla-to-change-routing-ciscozine

```
Ciscozine#show ip route 
Codes: L - local, C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route, H - NHRP, l - LISP
       + - replicated route, % - next hop override
Gateway of last resort is 172.16.255.6 to network 0.0.0.0
S*    0.0.0.0/0 [5/0] via 172.16.255.6
      172.16.0.0/16 is variably subnetted, 4 subnets, 2 masks
C        172.16.255.0/30 is directly connected, FastEthernet1/0
L        172.16.255.1/32 is directly connected, FastEthernet1/0
C        172.16.255.4/30 is directly connected, FastEthernet1/1
L        172.16.255.5/32 is directly connected, FastEthernet1/1
Ciscozine#
```
The return code is “Timeout”:

```
Ciscozine#show ip sla statistics 
IPSLAs Latest Operation Statistics
IPSLA operation id: 1
        Latest RTT: NoConnection/Busy/Timeout
Latest operation start time: 10:42:03 UTC Wed May 8 2013
Latest operation return code: Timeout
Number of successes: 8
Number of failures: 3
Operation time to live: Forever
Ciscozine#
```
The track object is down:

```
Ciscozine#show track
Track 10
  IP SLA 1 reachability
  Reachability is Down
    13 changes, last change 00:00:22
  Latest operation return code: Timeout
  Tracked by:
    STATIC-IP-ROUTING 0
Ciscozine#
```

**Red link again UP**

At this point I ping the web server (192.168.1.10) , then I reconnect the Ciscozine fastethernet1/0 cable

Ciscozine#ping 192.168.1.10 size 50 repeat 200
Type escape sequence to abort.
Sending 200, 50-byte ICMP Echos to 192.168.1.10, timeout is 2 seconds:
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
Success rate is 100 percent (200/200), round-trip min/avg/max = 12/71/172 ms
Ciscozine#

```
Ciscozine#show ip route 
Codes: L - local, C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route, H - NHRP, l - LISP
       + - replicated route, % - next hop override
Gateway of last resort is 172.16.255.2 to network 0.0.0.0
S*    0.0.0.0/0 [1/0] via 172.16.255.2
      172.16.0.0/16 is variably subnetted, 4 subnets, 2 masks
C        172.16.255.0/30 is directly connected, FastEthernet1/0
L        172.16.255.1/32 is directly connected, FastEthernet1/0
C        172.16.255.4/30 is directly connected, FastEthernet1/1
L        172.16.255.5/32 is directly connected, FastEthernet1/1
Ciscozine#
```
As you can see, when the main line comes up (now the default gateway is again 172.16.255.2), there isn’t a packet lost.


**References:**

nice explanation

thanks for this tutorial..

Thanks for this very nice IP SLA Explanation !

Thanks for this step-by-step clear explanation of the IP SLA .

how about default route on the server side? how the server(router) know the link has change to backup?

Good question lase. Also on the server side it is required the ip sla feature if you can’t use a dynamic protocol.

when our traffic is load balanced ,can we use 2 track object in same firewall if one line goes down and to push that traffic on another line ?

Hi Sagar, check this link:

https://community.cisco.com/t5/firewalls/using-multiple-outside-interface-on-asa-5520/td-p/1457255
