---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1699284-bad-packet-loss-between-unifi-switch-and-unifi-dream-machine-p-0cce64f2
title: "questions-1699284-bad-packet-loss-between-unifi-switch-and-unifi-dream-machine-p-0cce64f2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1699284-bad-packet-loss-between-unifi-switch-and-unifi-dream-machine-p-0cce64f2.md
source_anchor: ""
source_lines: [1, 53]
sha256: 941e3c03b9193921f03730824c863864c4eda0552066d41805efb5b3ef890e3a
---

# questions-1699284-bad-packet-loss-between-unifi-switch-and-unifi-dream-machine-p-0cce64f2

I'm pulling my hair out on this one. I'm seeing really bad network performance (bad packet loss) via one of my switches. In short, this issue makes VoIP impossible to use, causes certain applications to hang intermittently, and is really frustrating.
All devices are fully up to date with the latest software/firmware.
As a test, I am running iperf3, and I appear to be losing packets between my Switch 24 PoE and the UDM Pro. This test is run between a client connected to the UDP Pro, and the server running on the Switch 24:
    Accepted connection from 10.0.0.22, port 58682
[  5] local 10.0.0.5 port 5201 connected to 10.0.0.22 port 53440
[ ID] Interval           Transfer     Bitrate         Jitter    Lost/Total Datagrams
[  5]   0.00-1.00   sec   104 KBytes   852 Kbits/sec  0.437 ms  0/13 (0%)
[  5]   1.00-2.00   sec  80.0 KBytes   655 Kbits/sec  0.393 ms  8/18 (44%)
[  5]   2.00-3.00   sec   112 KBytes   917 Kbits/sec  3.333 ms  0/14 (0%)
[  5]   3.00-4.00   sec  80.0 KBytes   655 Kbits/sec  2.022 ms  8/18 (44%)
[  5]   4.00-5.00   sec   112 KBytes   917 Kbits/sec  1.793 ms  0/14 (0%)
[  5]   5.00-6.00   sec  80.0 KBytes   655 Kbits/sec  1.255 ms  8/18 (44%)
[  5]   6.00-7.00   sec   112 KBytes   917 Kbits/sec  0.992 ms  0/14 (0%)
[  5]   7.00-8.00   sec  40.0 KBytes   328 Kbits/sec  1.144 ms  13/18 (72%)
[  5]   8.00-9.00   sec   112 KBytes   917 Kbits/sec  0.959 ms  0/14 (0%)
[  5]   9.00-10.00  sec  80.0 KBytes   656 Kbits/sec  0.844 ms  8/18 (44%)
- - - - - - - - - - - - - - - - - - - - - - - - -
[ ID] Interval           Transfer     Bitrate         Jitter    Lost/Total Datagrams
[  5]   0.00-10.01  sec   912 KBytes   746 Kbits/sec  0.844 ms  45/159 (28%)  receiver
As you can see, the loss is around 28%, which is awful for a short hop on a local network.
When both client and server are connected just to the UDM Pro, I get this result:
Accepted connection from 10.0.0.22, port 58784
[  5] local 10.0.0.10 port 5201 connected to 10.0.0.22 port 59176
[ ID] Interval           Transfer     Bitrate         Jitter    Lost/Total Datagrams
[  5]   0.00-1.00   sec   120 KBytes   983 Kbits/sec  0.408 ms  0/15 (0%)
[  5]   1.00-2.00   sec   128 KBytes  1.05 Mbits/sec  0.300 ms  0/16 (0%)
[  5]   2.00-3.00   sec   128 KBytes  1.05 Mbits/sec  0.556 ms  0/16 (0%)
[  5]   3.00-4.00   sec   128 KBytes  1.05 Mbits/sec  0.410 ms  0/16 (0%)
[  5]   4.00-5.00   sec   128 KBytes  1.05 Mbits/sec  0.377 ms  0/16 (0%)
[  5]   5.00-6.00   sec   128 KBytes  1.05 Mbits/sec  0.503 ms  0/16 (0%)
[  5]   6.00-7.00   sec   128 KBytes  1.05 Mbits/sec  8.344 ms  0/16 (0%)
[  5]   7.00-8.00   sec   128 KBytes  1.05 Mbits/sec  3.167 ms  0/16 (0%)
[  5]   8.00-9.00   sec   128 KBytes  1.05 Mbits/sec  1.454 ms  0/16 (0%)
[  5]   9.00-10.00  sec   128 KBytes  1.05 Mbits/sec  0.628 ms  0/16 (0%)
- - - - - - - - - - - - - - - - - - - - - - - - -
[ ID] Interval           Transfer     Bitrate         Jitter    Lost/Total Datagrams
[  5]   0.00-10.01  sec  1.24 MBytes  1.04 Mbits/sec  0.628 ms  0/159 (0%)  receiver
So there's clearly no issues when just using the UDP Pro.
When doing a icmp test/ping, it seems to drop every other packet, so it looks like:
64 bytes from 10.0.0.251: icmp_seq=1481 ttl=64 time=7.697 ms
Request timeout for icmp_seq 1482
64 bytes from 10.0.0.251: icmp_seq=1483 ttl=64 time=1.348 ms
Request timeout for icmp_seq 1484
64 bytes from 10.0.0.251: icmp_seq=1485 ttl=64 time=2.794 ms
Request timeout for icmp_seq 1486
64 bytes from 10.0.0.251: icmp_seq=1487 ttl=64 time=1.298 ms
Request timeout for icmp_seq 1488
64 bytes from 10.0.0.251: icmp_seq=1489 ttl=64 time=1.959 ms
Request timeout for icmp_seq 1490
64 bytes from 10.0.0.251: icmp_seq=1491 ttl=64 time=9.319 ms
Request timeout for icmp_seq 1492
64 bytes from 10.0.0.251: icmp_seq=1493 ttl=64 time=1.525 ms
Any ideas or suggestions would be very helpful - thanks in advance.
