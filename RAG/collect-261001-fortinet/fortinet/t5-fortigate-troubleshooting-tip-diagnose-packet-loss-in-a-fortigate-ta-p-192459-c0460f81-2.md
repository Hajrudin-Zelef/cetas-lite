---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-diagnose-packet-loss-in-a-fortigate-ta-p-192459-c0460f81-2
title: "t5-fortigate-troubleshooting-tip-diagnose-packet-loss-in-a-fortigate-ta-p-192459-c0460f81"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-diagnose-packet-loss-in-a-fortigate-ta-p-192459-c0460f81.md
source_anchor: ""
source_lines: [165, 218]
sha256: 310e9e6b0466354a33cff998a4aa694a5a84a01fa583890367eb922fe2584bb9
---

# t5-fortigate-troubleshooting-tip-diagnose-packet-loss-in-a-fortigate-ta-p-192459-c0460f81

```
Endeavour-kvm96 # get sys performance status 
CPU states: 0% user 1% system 0% nice 99% idle 0% iowait 0% irq 0% softirq 
CPU0 states: 0% user 1% system 0% nice 99% idle 0% iowait 0% irq 0% softirq 
Memory: 2040052k total, 966408k used (47.4%), 663468k free (32.5%), 410176k freeable (20.1%) 
Average network usage: 38 / 3 kbps in 1 minute, 41 / 5 kbps in 10 minutes, 40 / 5 kbps in 30 minutes
Maximal network usage: 56 / 11 kbps in 1 minute, 409 / 110 kbps in 10 minutes, 409 / 144 kbps in 30 minutes
Average sessions: 43 sessions in 1 minute, 24 sessions in 10 minutes, 24 sessions in 30 minutes
Maximal sessions: 47 sessions in 1 minute, 48 sessions in 10 minutes, 51 sessions in 30 minutes
Average session setup rate: 0 sessions per second in last 1 minute, 0 sessions per second in last 10 minutes, 0 sessions per second in last 30 minutes
Maximal session setup rate: 0 sessions per second in last 1 minute, 16 sessions per second in last 10 minutes, 23 sessions per second in last 30 minutes
Average NPU sessions: 0 sessions in last 1 minute, 0 sessions in last 10 minutes, 0 sessions in last 30 minutes
Maximal NPU sessions: 0 sessions in last 1 minute, 0 sessions in last 10 minutes, 0 sessions in last 30 minutes
Virus caught: 0 total in 1 minute
IPS attacks blocked: 0 total in 1 minute
Uptime: 4 days, 8 hours, 5 minutes
```

1. **WAD drops:** If a Web proxy or Explicit proxy is configured on the FortiGate and is suspected to be dropping packets, run the following diagnose debug commands to trace the drop reasons due to WAD processing.


**Caution**: **The below WAD debugs are quite verbose, even with a specific filter applied, it could print a lot of debugs. It is preferable to run the following debugs in a change window:**

```
diagnose wad debug enable category all
diagnose wad debug enable level verbose
diagnose wad debug display pid enable
diagnose wad filter src 172.16.0.1   
diagnose wad filter dst 4.2.2.2      
diagnose debug enable
...........
diagnose debug disable
```
**In this example:**

- src 172.16.0.1 is the client IP address.
- dst 4.2.2.2 is the server IP address.


More details regarding wad drops are discussed here: Technical Tip: How to debug the web proxy and explicit proxy.**Further troubleshooting:**

Identifying whether the issue is with the FortiGate or not requires running the packet capture on the ingress and egress interfaces of the firewall for the destination IP. Initiate the ICMP packets from the source machine to the destination and observe the packet loss. Refer to Technical Tip: Useful filters for sniffer packet capture.


Isolating the issue to the FortiGate will require bypassing any other device in the path. If packet loss is going to the internet, then make sure the FortiGate is directly connected to the ISP, and the testing PC is directly connected to a FortiGate interface.


If there is an alternative interface that has a route to the same destination, redirect traffic through that interface using a policy route for forwarding traffic to see any difference in packet loss. See Technical Tip: Configuring the Firewall Policy Routes.


When initiating pings from the FortiGate CLI, use the 'ping-option' command to specify the interface, then initiate the ping to the destination. For more information about ping-option, refer to Troubleshooting Tip: Using PING options from the FortiGate CLI.


**Related articles:**
