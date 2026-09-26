---
id: collect-260926-mikrotik/mikrotik/questions-739201-mikrotik-traffic-flow-netflow-octets-counter-wrap-be739cfd
title: "questions-739201-mikrotik-traffic-flow-netflow-octets-counter-wrap-be739cfd"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-739201-mikrotik-traffic-flow-netflow-octets-counter-wrap-be739cfd.md
source_anchor: ""
source_lines: [1, 114]
sha256: 7f18fedcde1d3c5dbf9db9dc8957bba75c764bf475c48f7fa294e6c7f5730b0f
---

# questions-739201-mikrotik-traffic-flow-netflow-octets-counter-wrap-be739cfd

I am using Traffic Flow with pmacct (nfacct) to do IP Accounting.
I've noticed that if a flow exceeds ~4GBytes in less than a minute (which is my active-flow-timeout) the exported flow Octets counter wraps around losing a significant amount of total data measured.
I believe this is an issue with the Octet counter being 32bit unsigned and if the traffic is over that threshold (4294967296) then the exporter wraps around the counter without first sending the flow to the collector (I am not sure how other vendors handle this).
This is quite serious since it results in very wrong traffic totals!
Here is my traffic flow configuration:
/ip traffic-flow
set active-flow-timeout=1m cache-entries=1k enabled=yes interfaces=sfp1
/ip traffic-flow target
add dst-address=X.X.X.X v9-template-refresh=60 v9-template-timeout=1m
And here are a couple of flow captures from wireshark.
Flow 3
    [Duration: 59.590000000 seconds (switched)]
    Packets: 5700194
    Octets: 4255323704
    InputInt: 16
    OutputInt: 0
    SrcAddr: 31.X.X.254
    DstAddr: 185.X.X.254
    Protocol: UDP (17)
    IP ToS: 0x00
    SrcPort: 2043 (2043)
    DstPort: 2299 (2299)
    NextHop: 185.X.X.X
    DstMask: 0
    SrcMask: 0
    TCP Flags: 0x00
    Destination Mac Address: Routerbo_XX:XX:XX (d4:ca:6d:XX:XX:XX)
    Post Source Mac Address: 00:00:00_00:00:00 (00:00:00:00:00:00)
    Post NAT Source IPv4 Address: 31.X.X.254
    Post NAT Destination IPv4 Address: 185.X.X.254
    Post NAPT Source Transport Port: 0
    Post NAPT Destination Transport Port: 0
Second capture:
Flow 3
    [Duration: 59.590000000 seconds (switched)]
    Packets: 5532208
    Octets: 4003344704
    InputInt: 16
    OutputInt: 0
    SrcAddr: 31.X.X.254
    DstAddr: 185.X.X.254
    Protocol: UDP (17)
    IP ToS: 0x00
    SrcPort: 2043 (2043)
    DstPort: 2299 (2299)
    NextHop: 185.X.X.X
    DstMask: 0
    SrcMask: 0
    TCP Flags: 0x00
    Destination Mac Address: Routerbo_XX:XX:XX (d4:ca:6d:XX:XX:XX)
    Post Source Mac Address: 00:00:00_00:00:00 (00:00:00:00:00:00)
    Post NAT Source IPv4 Address: 31.X.X.254
    Post NAT Destination IPv4 Address: 185.X.X.254
    Post NAPT Source Transport Port: 0
    Post NAPT Destination Transport Port: 0
At the time of those captures, a bandwidth test (UDP, 1500bytes, 1Gbit, receive) was running for quite some time.
So running at 1gbit for 60seconds (active-flow-timeout) it should have measured at least ~7864320000 Octets (~7.3GB)
If I reduce the bandwidth test to 460mbit then the exported flows seem to report the traffic properly since the Octets counter does not exceed the 32bit unsigned maximum.
Though I see quite a lot of overhead and I wonder why that is.
At 460mbit sustained traffic, in 60seconds it should measure ~3617587200 octets (=3.36GB).
But instead it measured 4269160500 (=3.9GB)
I am not sure where the extra ~600MB came from.
Flow 6
    [Duration: 59.590000000 seconds (switched)]
    Packets: 2846107
    Octets: 4269160500
    InputInt: 16
    OutputInt: 0
    SrcAddr: 31.X.X.254
    DstAddr: 185.X.X.254
    Protocol: UDP (17)
    IP ToS: 0x00
    SrcPort: 2058 (2058)
    DstPort: 2314 (2314)
    NextHop: 185.X.X.X
    DstMask: 0
    SrcMask: 0
    TCP Flags: 0x00
    Destination Mac Address: Routerbo_0d:95:72 (d4:ca:6d:XX:XX:XX)
    Post Source Mac Address: 00:00:00_00:00:00 (00:00:00:00:00:00)
    Post NAT Source IPv4 Address: 31.X.X.254
    Post NAT Destination IPv4 Address: 185.X.X.254
    Post NAPT Source Transport Port: 0
    Post NAPT Destination Transport Port: 0
But if I increase the bandwidth test to 480mbit for example, then the exported flow has its counter wrapped around losing a significant amount of data (ie: ~4GBytes of data)
Flow 3
    [Duration: 59.590000000 seconds (switched)]
    Packets: 2865308
    Octets: 2994704 <-- Only 2.8MB?! Even with 64byte packets,
                    based on the measured packets above, 
                    it should have measured > 174MBytes of data!
    InputInt: 16
    OutputInt: 0
    SrcAddr: 31.X.X.254
    DstAddr: 185.X.X.254
    Protocol: UDP (17)
    IP ToS: 0x00
    SrcPort: 2055 (2055)
    DstPort: 2311 (2311)
    NextHop: 185.X.X.X
    DstMask: 0
    SrcMask: 0
    TCP Flags: 0x00
    Destination Mac Address: Routerbo_0d:95:72 (d4:ca:6d:XX:XX:XX)
    Post Source Mac Address: 00:00:00_00:00:00 (00:00:00:00:00:00)
    Post NAT Source IPv4 Address: 31.X.X.254
    Post NAT Destination IPv4 Address: 185.X.X.254
    Post NAPT Source Transport Port: 0
    Post NAPT Destination Transport Port: 0
The above tests were made on a CCR1036-8G-2S+ running version 6.32.1 (I cannot upgrade since this is a production system).
Doing the same tests on a x86 installation (running 6.29 - also cannot upgrade because it's in production) the results are even worse!
There it appears that the Octets counter wraps around at 2147483647 which suggests that either in versions < 6.32.1 or in non Tilera builds the Octets counter is 32bit Signed.
The whole situation is pretty much the same with when you monitor a Gbit interface with v1 SNMP (32bit counters). The solution in SNMP is very simple. Use SNMP v2 that supports 64bit counters. But I cannot find any solution for Netflow.
Can anyone else confirm this issue? Does anyone know a workaround for it? Is this a limitation of the netflow protocol or a bug in RouterOS? How do other vendors handle this (I don't have any other equipment at the moment to test this out) ?
