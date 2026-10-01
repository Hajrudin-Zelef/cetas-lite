---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-898110-edgeos-bgp-routes-received-but-not-being-injected-8c19f5ee
title: "questions-898110-edgeos-bgp-routes-received-but-not-being-injected-8c19f5ee"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["decode"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-898110-edgeos-bgp-routes-received-but-not-being-injected-8c19f5ee.md
source_anchor: ""
source_lines: [1, 37]
sha256: bc2410471b933f08874ebd08668316999e2f05446d29c048aa85200a81a64ead
---

# questions-898110-edgeos-bgp-routes-received-but-not-being-injected-8c19f5ee

I'm trying to set up BGP over IPSEC with Google Cloud Platform. The IPSEC tunnel is successfully established, and GCP reports the BGP session as up. EdgeMax router is receiving the routes, but isn't injecting them into the RIB.
dan@router:~$ show ip bgp neighbors 169.254.0.1 received-routes
BGP table version is 7, local router ID is 169.254.0.2
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal
Origin codes: i - IGP, e - EGP, ? - incomplete
    Network          Next Hop            Metric    LocPrf       Weight Path
*>  10.142.0.0/20    169.254.0.1          100                   0       65000 ?
Total number of prefixes 1
dan@router:~$ show ip route
Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP
       O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       > - selected route, * - FIB route, p - stale info
IP Route Table for VRF "default"
S    *> 0.0.0.0/0 [210/0] via 71.186.153.1, eth0
C    *> <external address> is directly connected, eth0
C    *> 127.0.0.0/8 is directly connected, lo
C    *> 169.254.0.0/30 is directly connected, vti0
S    *> 169.254.0.1/32 [1/0] is directly connected, vti0
C    *> 192.168.0.0/24 is directly connected, eth1
S    *> 192.168.42.0/24 [1/0] via 192.168.0.29, eth1
Initially the local router ID was the address of an interface I had removed, so I've updated that to be 169.254.0.2, but that hasn't helped.
My log looks like this, repeated every few seconds:
Feb 20 16:01:30 router BGP[1679]:  BGP-6: [RIB] Scanning BGP Network Routes...
Feb 20 16:01:38 router BGP[1679]:  BGP-6: 169.254.0.1-Outgoing [FSM] Keep-alive-Timer Expiry
Feb 20 16:01:38 router BGP[1679]:  BGP-6: 169.254.0.1-Outgoing [FSM] State: Established Event: 11
Feb 20 16:01:38 router BGP[1679]:  BGP-6: 169.254.0.1-Outgoing [ENCODE] Msg-Hdr: Type 4
Feb 20 16:01:38 router BGP[1679]:  BGP-6: 169.254.0.1-Outgoing [ENCODE] Keepalive: 806 KAlive msg(s) sent
Feb 20 16:01:41 router BGP[1679]:  BGP-6: 169.254.0.1-Outgoing [FSM] AS-Origination Timer Expiry
Feb 20 16:01:41 router BGP[1679]:  BGP-6: 169.254.0.1-Outgoing [FSM] State: Established Event: 33
Feb 20 16:01:41 router BGP[1679]:  BGP-6: 169.254.0.1-Outgoing [DECODE] Msg-Hdr: type 4, length 19
Feb 20 16:01:41 router BGP[1679]:  BGP-6: 169.254.0.1-Outgoing [DECODE] Msg-Hdr: Requesting immediate Read (0)
Feb 20 16:01:41 router BGP[1679]:  BGP-6: 169.254.0.1-Outgoing [DECODE] KAlive: Received!
Feb 20 16:01:41 router BGP[1679]:  BGP-6: 169.254.0.1-Outgoing [FSM] State: Established Event: 26
The full configuration is here: https://gist.github.com/danellis/82b471665cbe3fb7e8d788e03b8ab257
Hopefully there's just one small thing I'm missing. Can anyone see what it is?
