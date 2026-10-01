---
id: collect-261001-cisco/cisco/using-ip-sla-to-change-routing-ciscozine-1
title: "using-ip-sla-to-change-routing-ciscozine"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/using-ip-sla-to-change-routing-ciscozine.md
source_anchor: ""
source_lines: [1, 138]
sha256: 289d7025e5ef0ae04e87c534c7c5ce48cfeaa89ba918394e52f30b02663eb714
---

# using-ip-sla-to-change-routing-ciscozine

Cisco IP SLAs is a part of Cisco IOS that allows Cisco **customers to analyze IP service levels for IP applications and services** by using active traffic monitoring for measuring network performance. With Cisco IOS IP SLAs, **service provider** customers **can measure and provide service level agreements**, and **enterprise** customers can **verify service levels**, verify outsourced service level agreements, and understand network performance. Cisco IOS IP SLAs can perform network assessments, verify quality of service (QoS), ease the deployment of new services, and assist with network troubleshooting.

IP SLAs collects a unique subset of **these** performance **metrics**:

- Delay (both round-trip and one-way)
- Jitter (directional)
- Packet loss (directional)
- Packet sequencing (packet ordering)
- Path (per hop)
- Connectivity (directional)
- Server or website download time


**Example** Suppose that the Ciscozine router has two different links, one is the main connection (red link) and the other one (blue link) is the backup connection; the question is: how can I enable the backup link if the main connection goes down? In general, the best solution for this scenario is to use the dynamic routing protocol, but what can I do if I can’t use them? The solution is the IP SLA.

**1. Define the ip sla operation.** The Ciscozine router will send an ICMP request to 172.16.255.2 (the Ciscozine default gateway) every 10 second with a timeout of 1000ms and a threshold value of 500ms.

Ciscozine(config)#ip sla 1
Ciscozine(config-ip-sla)#icmp-echo 172.16.255.2 source-interface FastEthernet1/0
Ciscozine(config-ip-sla-echo)#timeout 5000
Ciscozine(config-ip-sla-echo)#frequency 10
Ciscozine(config-ip-sla-echo)#threshold 500

**2. Start the ip sla.** It is possible schedule the SLA operation in different ways but in this tutorial I want to start the ip SLA operation immediately and forever. Notice that the “1”  refers to “ip sla 1” command.

Ciscozine(config)#ip sla schedule 1 start-time now life forever

**3. Track the state of IP SLA.** Every IP SLAs operation maintains an operation return-code value. This return code is interpreted by the tracking process. The return code may return OK, OverThreshold, and several other return codes. Two aspects of an IP SLAs operation can be tracked: state and reachability.

| **Tracking** | **Return Code** | **Track State** | 
| State | OK (all other return codes) | Up Down | 
| Reachability | OK or over threshold (all other return codes) | Up Down | 

In this case, I prefer to use the “reachability”, so the “track state” will be down only in case of a ICMP timeout.

Ciscozine(config)#track 10 ip sla 1 reachability

**Remember:** If you want to use the “state”, remember that the “track state” will be down also if the the threshold is reached.

**Note:** with Cisco IOS Release 12.4(20)T, 12.2(33)SXI1, 12.2(33)SRE and Cisco IOS XE Release 2.4, the track rtr command is replaced by the track ip sla command. See the track ip sla command for more information.

**4. Define the tracked route.** At the end, I must delete the old default gateway entry, add the default gateway with the track feature (notice that the number 10 represents the track object defined in the previous step) and insert a default route with a distance administrative less “strong” . Hence if the track status is down the last route will be used to forward all the traffic (notice that the number 5 define the administrative distance).

Ciscozine(config)#ip route 0.0.0.0 0.0.0.0 172.16.255.2 track 10
Ciscozine(config)#no ip route 0.0.0.0 0.0.0.0 172.16.255.2
Ciscozine(config)#ip route 0.0.0.0 0.0.0.0 172.16.255.6 5

**5. Check the IP SLA.**

Now that I have defined the IP SLA object, I can check some useful informations when the main link (red link) is UP or NOT.

**Red link UP**

To display information about the IP route track table:

Ciscozine#show ip route track-table
 ip route 0.0.0.0 0.0.0.0 172.16.255.2 track 10 state is [up]
Ciscozine#

To display information about the IP routing table:

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
To display information about IP SLA

```
Ciscozine#show ip sla statistics 
IPSLAs Latest Operation Statistics
IPSLA operation id: 1
        Latest RTT: 24 milliseconds
Latest operation start time: 10:40:43 UTC Wed May 8 2013
Latest operation return code: OK
Number of successes: 3
Number of failures: 0
Operation time to live: Forever
Ciscozine#
```
To display information about the track object

```
Ciscozine#show track
Track 10
  IP SLA 1 reachability
  Reachability is Up
    12 changes, last change 00:33:48
  Latest operation return code: OK
  Latest RTT (millisecs) 24
  Tracked by:
    STATIC-IP-ROUTING 0
Ciscozine#
```

**Red link DOWN**

First of all, I ping the web server (192.168.1.10) that it is on the headquarter, then I unplug the Ciscozine fastethernet1/0 cable

Ciscozine#ping 192.168.1.10 repeat 200
Type escape sequence to abort.
Sending 200, 50-byte ICMP Echos to 192.168.1.10, timeout is 2 seconds:
!!!!!!!!!!!!!!!!!!!!!!!!!!!!....
May 8 03:23:30.082: %TRACKING-5-STATE: 10 ip sla 1 reachability Up->Down.!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
Success rate is 97 percent (195/200), round-trip min/avg/max = 20/66/248 ms
Ciscozine#

As you can see, when the cable is unplugged there are three different events:

1. there are some timeout
2. the tracking state goes down (May 8 03:23:30.082: %TRACKING-5-STATE: 10 ip sla 1 reachability Up->Down)
3. the route tracked goes down and the backup default route (ip route 0.0.0.0 0.0.0.0 172.16.255.6 5) is up.

So the output of the previous show commands will be:

The track object is down:

Ciscozine#show ip route track-table
 ip route 0.0.0.0 0.0.0.0 172.16.255.2 track 10 state is [down]

The default route is 172.16.255.6 (the backup connection):

