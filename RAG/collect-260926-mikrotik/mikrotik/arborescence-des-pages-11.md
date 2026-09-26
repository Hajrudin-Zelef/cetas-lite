---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-11
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-11.md
source_anchor: ""
source_lines: [1, 55]
sha256: feab55f6a084a8c2184edbe45ac4aed07cdc4f448aa2f2a86cdab0cc094483c3
---

# Summary

The Bandwidth Tester can be used to measure the throughput to another MikroTik router (either wired or wireless) and thereby help to discover network "bottlenecks".

The TCP test uses the standard TCP protocol with acknowledgments and follows the TCP algorithm on how many packets to send according to latency, dropped packets, and other features in the TCP algorithm. Please review the TCP protocol for details on its internal speed settings and how to analyze its behavior. Statistics for throughput are calculated using the entire size of the TCP data stream. As acknowledgments are an internal working of TCP, their size and usage of the link are not included in the throughput statistics. Therefore this statistic is not as reliable as the UDP statistic when estimating throughput.

The UDP tester sends 110% or more packets than currently reported as received on the other side of the link. To see the maximum throughput of a link, the packet size should be set for the maximum MTU allowed by the links which is usually 1500 bytes. There is no acknowledgment required by UDP; this implementation means that the closest approximation of the throughput can be seen.

- Up to RouterOS version 6.44beta39 Bandwidth Test used only single CPU core and reached its limits when core was 100% loaded.
- Bandwidth Test uses all available bandwidth (by default) and may impact network usability.

- Bandwidth Test uses a lot of resources. If you want to test real throughput of a router, you should run bandwidth test through the tested router not from or to it. To do this you need at least 3 routers connected in chain: the Bandwidth Server, the router being tested and the Bandwidth Client.

- If you use UDP protocol then Bandwidth Test counts IP header+UDP header+UDP data. In case if you use TCP then Bandwidth Test counts only TCP data (TCP header and IP header are not included).

# Bandwidth Test Server

| Property | Description | 
|---|---|
| **allocate-udp-ports-from** (*integer 1000..64000* ; Default:**2000** ) | Beginning of UDP port range | 
| **authenticate** (*yes \| no* ; Default:**yes** ) | Communicate only with authenticated clients | 
| **enabled** (*yes \| no* ; Default:**yes** ) | Defines whether bandwidth server is enabled or not | 
| **max-sessions** (*integer 1..1000* ; Default:**100** ) | Maximal simultaneous test count | 

**Example**

Bandwidth Server:

Active sessions:

To enable **bandwidth-test** server without client authentication:

# Bandwidth Test Client

| Property | Description | 
|---|---|
| **address** (*IP address \| IPv6 prefix[%interface]* ; Default:) | IP address of host | 
| **direction** (*both \| receive \| transmit* ; Default: **receive** ) | Direction of data flow | 
| **duration** (*time* ; Default:  ) | Duration of the test | 
| **interval** (*time: 20ms..5s* ; Default: **1s** ) | Delay between reports (in seconds) | 
| **local-tx-speed** (*integer 0..18446744073709551615* ; Default: ) | Transfer test maximum speed (bits per second) | 
| **local-udp-tx-size** (*integer: 28..64000* ) | Local transmit packet size in bytes | 
| **password** (*string* ; Default:**""** ) | Password for the remote user | 
| **protocol** (*udp \| tcp* ; Default:**udp** ) | Protocol to use | 
| **random-data** (*yes \| no* ; Default:**no** ) | If random-data is set to yes, the payload of the bandwidth test packets will have incompressible random data stream so that links that use data compression will not distort the results (this is CPU intensive and random-data should be set to no for low speed CPUs) | 
| **remote-tx-speed** (*integer 0..18446744073709551615* ; Default: ) | Receive test maximum speed (bits per second) | 
| **remote-udp-tx-size** (*integer: 28..64000* ) | Remote transmit packet size in bytes | 
| **connection-count** (*integer 1..255* ; Default:) | Number of TCP connections to use | 
| **user** (*string* ; Default:**""** ) | Remote user | 

**Example**

To run 15-second long bandwidth-test to the **10.0.0.32** host sending and receiving **1000**-byte UDP packets and using username **admin** to connect:

**Link-local IPv6 example**:
