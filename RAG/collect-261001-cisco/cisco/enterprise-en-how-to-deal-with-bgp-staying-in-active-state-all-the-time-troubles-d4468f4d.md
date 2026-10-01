---
id: collect-261001-cisco/cisco/enterprise-en-how-to-deal-with-bgp-staying-in-active-state-all-the-time-troubles-d4468f4d
title: "enterprise-en-how-to-deal-with-bgp-staying-in-active-state-all-the-time-troubles-d4468f4d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-how-to-deal-with-bgp-staying-in-active-state-all-the-time-troubles-d4468f4d.md
source_anchor: ""
source_lines: [1, 33]
sha256: de4887a2c8b38b2d2778a0f58b5b5b9b0f8c544bf2eb364eb737a2725f23feff
---

# enterprise-en-how-to-deal-with-bgp-staying-in-active-state-all-the-time-troubles-d4468f4d

Hello everyone,
Today, I will introduce how to deal with BGP staying in Active state all the time.
BGP stays in the Active state all the time, indicating that TCP failed to establish a connection and the timer expired.
In the connect state, a TCP connection is initiated proactively, and in the active state, a TCP connection is waiting passively.
If the connect Retry timer expires, BGP restarts the Connect Retry timer and attempts to establish a TCP connection with the peer. In this case, BGP remains in the Connect state. In the Active state, BGP does not initiate a TCP connection.
Neighbor assignment error
The peer device does not have a route to the local device.
The two ends use loopback interfaces to establish a neighbor relationship. The source is not specified on both ends.
The local configuration is correct, and the local end does not receive TCP packets from the peer end.
1. Run the display current-configuration configuration bgp command to check the BGP configurations on both ends.
[*HUAWEI]bgp 100
[*HUAWEI-bgp]router-id 10.3.3.3
[*HUAWEI-bgp]peer 10.1.1.1 as-number 100
[*HUAWEI-bgp]peer 10.1.1.1 connect-interface Loopback 0
Check whether the peer x.x.x.x connect-interface command on both ends is configured with the interface IP address or loopback interface. The two ends should be the same to avoid an interface address at one end and a loopback address at the other end.
2. Run display ip routing table on the remote device to check whether there is a route to the local device. If not, you need to check the reason why the peer does not learn the local route, such as route filtering.
3. When the two parties use the loopback interface to establish a BGP neighbor relationship, the source is not specified.
Use the peer x.x.x.x connect-interface command to specify the source interface or source IP address.
<HUAWEI> system-view
[~HUAWEI] interface LoopBack 0
[*HUAWEI-LoopBack0] ip address 10.1.1.1 32
[*HUAWEI-LoopBack0] quit
[*HUAWEI] bgp 100
[*HUAWEI-bgp] peer 10.1.1.1 as-number 200
[*HUAWEI-bgp] peer 10.1.1.1 connect-interface LoopBack 0
4. The local configuration is correct, and no TCP packets sent by the peer end are received. You need to check the BGP state machine and configuration of the peer device.
That is all I want to share with you!
Thanks for sharing
Very interesting!
Good share
Thanks for sharing
Good one
Thanks for sharing
