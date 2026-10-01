---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-10
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [2271, 2486]
sha256: c765fddaffacf5155d9385487c4c871626cfc7395e5d5f31715ad0e054bb62a0
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

      Configuration of a backup static route on a router.


Topology




                   Figure 4.1 Lab topology for static and default routes


                                       HUAWEI TECHNOLOGIES                 Page51
Scenario

Assume that you are a network administrator of a company that contains a single
administrative domain and within the administrative domain, multiple networks
have been defined, for which currently no method of routing exists.
Since the network scale is small, with only a few networks, static routes and default
routes are to be used to implement interwork communication. The network
addressing is to be applied as shown in Figure 4.1.
If a password is requested, and unless otherwise stated, please use the password:
huawei


Tasks


Step 1 Perform basic system and IP address configuration.

Configure the device names and IP addresses for R1, R2, and R3.
<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R1
[R1]interface GigabitEthernet 0/0/0
[R1-GigabitEthernet0/0/0]ip address 10.0.13.1 24
[R1-GigabitEthernet0/0/0]quit
[R1]interface GigabitEthernet 0/0/1
[R1-GigabitEthernet0/0/1]ip address 10.0.12.1 24
[R1-GigabitEthernet0/0/1]quit
[R1]interface LoopBack 0
[R1-LoopBack0]ip address 10.0.1.1 24




Run the display current-configuration command to check the configuration.
<R1>display ip interface brief
Interface                          IP Address/Mask       Physical   Protocol
......output omitted......
GigabitEthernet0/0/0              10.0.13.1/24           up          up
GigabitEthernet0/0/1              10.0.12.1/24           up          up
GigabitEthernet0/0/2              unassigned             up          down
LoopBack0                         10.0.1.1/24            up          up(s)


                                                   HUAWEI TECHNOLOGIES         Page52
......output omitted......


<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R2
[R2]interface GigabitEthernet 0/0/1
[R2-GigabitEthernet0/0/1]ip address 10.0.12.2 24
[R2-GigabitEthernet0/0/1]quit
[R2]interface GigabitEthernet0/0/2
[R2-GigabitEthernet0/0/2]ip add 10.0.23.2 24
[R2-GigabitEthernet0/0/2]quit
[R2]interface LoopBack0
[R2-LoopBack0]ip address 10.0.2.2 24




<R2>display ip interface brief
Interface                            IP Address/Mask        Physical    Protocol
......output omitted......
GigabitEthernet0/0/0              unassigned                up           down
GigabitEthernet0/0/1              10.0.12.2/24              up           up
GigabitEthernet0/0/2              10.0.23.2/24              up           up
LoopBack0                         10.0.2.2/24               up           up(s)
......output omitted......


<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R3
[R3]interface GigabitEthernet 0/0/0
[R3-GigabitEthernet0/0/0]ip address 10.0.13.3 24
[R3-GigabitEthernet0/0/0]quit
[R3]interface GigabitEthernet0/0/2
[R3-GigabitEthernet0/0/2]ip address 10.0.23.3 24
[R3-GigabitEthernet0/0/2]quit
[R3]interface LoopBack 0
[R3-LoopBack0]ip address 10.0.3.3 24


<R3>display ip interface brief
Interface                            IP Address/Mask        Physical    Protocol
......output omitted......
GigabitEthernet0/0/0              10.0.13.3/24         up          up
GigabitEthernet0/0/1              unassigned           up          down
GigabitEthernet0/0/2              10.0.23.3/24         up          up



                                                   HUAWEI TECHNOLOGIES             Page53
LoopBack0                             10.0.3.3/24         up        up(s)
......output omitted......


Use the ping command to test network connectivity from R1.
<R1>ping 10.0.12.2
  PING 10.0.12.2: 56 data bytes, press CTRL_C to break
     Reply from 10.0.12.2: bytes=56 Sequence=1 ttl=255 time=30 ms
     Reply from 10.0.12.2: bytes=56 Sequence=2 ttl=255 time=30 ms
     Reply from 10.0.12.2: bytes=56 Sequence=3 ttl=255 time=30 ms
     Reply from 10.0.12.2: bytes=56 Sequence=4 ttl=255 time=30 ms
     Reply from 10.0.12.2: bytes=56 Sequence=5 ttl=255 time=30 ms


  --- 10.0.12.2 ping statistics ---
     5 packet(s) transmitted
     5 packet(s) received
     0.00% packet loss
     round-trip min/avg/max = 30/30/30 ms


<R1>ping 10.0.13.3
  PING 10.0.13.2: 56 data bytes, press CTRL_C to break
     Reply from 10.0.13.3: bytes=56 Sequence=1 ttl=255 time=6 ms
     Reply from 10.0.13.3: bytes=56 Sequence=2 ttl=255 time=2 ms
     Reply from 10.0.13.3: bytes=56 Sequence=3 ttl=255 time=2 ms
     Reply from 10.0.13.3: bytes=56 Sequence=4 ttl=255 time=2 ms
     Reply from 10.0.13.3: bytes=56 Sequence=5 ttl=255 time=2 ms


  --- 10.0.13.3 ping statistics ---
     5 packet(s) transmitted
     5 packet(s) received
     0.00% packet loss
     round-trip min/avg/max = 2/2/6 ms


Use the ping command to test network connectivity from R2


<R2>ping 10.0.23.3
  PING 10.0.23.3: 56 data bytes, press CTRL_C to break
     Reply from 10.0.23.3: bytes=56 Sequence=1 ttl=255 time=31 ms
     Reply from 10.0.23.3: bytes=56 Sequence=2 ttl=255 time=31 ms
     Reply from 10.0.23.3: bytes=56 Sequence=3 ttl=255 time=41 ms
     Reply from 10.0.23.3: bytes=56 Sequence=4 ttl=255 time=31 ms
     Reply from 10.0.23.3: bytes=56 Sequence=5 ttl=255 time=41 ms



                                                    HUAWEI TECHNOLOGIES     Page54
 --- 10.0.23.3 ping statistics ---
   5 packet(s) transmitted
   5 packet(s) received
   0.00% packet loss
   round-trip min/avg/max = 31/35/41 ms




Step 2 Test connectivity

Use the ping command to test network connectivity from R2 to neworks
10.0.13.0/24 and 10.0.3.0/24
<R2>ping 10.0.13.3
 PING 10.0.13.3: 56 data bytes, press CTRL_C to break
   Request time out
   Request time out
   Request time out
   Request time out
   Request time out


 --- 10.0.13.3 ping statistics ---
   5 packet(s) transmitted
   0 packet(s) received
   100.00% packet loss


<R2>ping 10.0.3.3
 PING 10.0.3.3: 56 data bytes, press CTRL_C to break
   Request time out
   Request time out
   Request time out
   Request time out
   Request time out


 --- 10.0.3.3 ping statistics ---
   5 packet(s) transmitted
   0 packet(s) received
100.00% packet loss


If R2 wishes to communicate with the network segment 10.0.3.0, a route destined for
this network segment must be configured on R2, and routes destined for the R2
interface must be configured on R3.


                                              HUAWEI TECHNOLOGIES        Page55
The preceding test result shows that R2 cannot communicate with 10.0.3.3 and
10.0.13.3.


