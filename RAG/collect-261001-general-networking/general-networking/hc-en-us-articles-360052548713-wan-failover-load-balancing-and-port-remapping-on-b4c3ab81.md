---
id: collect-261001-general-networking/general-networking/hc-en-us-articles-360052548713-wan-failover-load-balancing-and-port-remapping-on-b4c3ab81
title: "hc-en-us-articles-360052548713-wan-failover-load-balancing-and-port-remapping-on-b4c3ab81"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["cost", "distribution", "kill switch", "latency", "parameters"]
source: docs/RAG/collect-261001-general-networking/hc-en-us-articles-360052548713-wan-failover-load-balancing-and-port-remapping-on-b4c3ab81.md
source_anchor: ""
source_lines: [1, 51]
sha256: e4472374511bf2e0df83da6b6e087eea33ef7c19e64811c381062ebe3b420f01
---

# hc-en-us-articles-360052548713-wan-failover-load-balancing-and-port-remapping-on-b4c3ab81

WAN Failover, Load Balancing and Port Remapping on UniFi Gateways
UniFi Gateways support Multi-WAN configurations to improve internet reliability and performance. You can connect up to eight WAN interfaces, limited only by the total number of physical ports on your device (with at least one reserved for LAN), and assign each to either Failover or Load Balancing mode.
This article explains how these modes work, how WAN health is monitored, and how to set up port roles and backup connectivity—including LTE—for different use cases.
WAN Modes
WAN interfaces in UniFi Network can be assigned to one of two operational modes:
- Failover, which prioritizes a single WAN and uses others only when needed.
- Load Balancing, which distributes new connections across multiple WANs based on configured weights.
Both modes can be combined in flexible ways to match your network goals.
WAN Failover
Failover mode is designed to prioritize one active WAN connection, while others remain on standby. If the primary WAN becomes unreachable, UniFi automatically reroutes traffic to the next highest-priority WAN. Once the primary WAN recovers, traffic returns to it.
This is ideal for networks with a main ISP and one or more backup links used only during outages.
WAN Load Balancing
Load balancing distributes new outbound connections across multiple WAN interfaces based on configurable traffic splits (e.g., 80% WAN1, 20% WAN2). This maximizes total available bandwidth and can reduce load on a single provider.
Traffic distribution is handled using a connection-hashing algorithm that maps each session to a WAN based on its source and destination IP address. This ensures session stickiness—a given session stays on the same WAN until it ends, even if other WANs are available.
In the event of a WAN failure, load balancing automatically shifts affected traffic to the remaining online WANs.
Creating a Dedicated WAN for Specific Traffic
In addition to standard failover and load balancing, UniFi allows you to dedicate a specific WAN to selected traffic types—such as VoIP calls, guest WiFi, camera streams, or IoT devices—for performance, isolation, or cost control.
To do this, ensure the WAN is not configured for active load balancing. Set its role to Failover (if using Failover Only mode) or Backup (if using Load Balancing mode). Then, use Policy-Based Routing to create a rule that assigns the desired device(s) or VLAN as the Source and the dedicated WAN as the Interface.
Note: To allow failover for this traffic if the WAN goes down, disable Kill Switch in the policy configuration.
WAN Monitoring & Failover Conditions
To determine if a WAN is online, UniFi performs continuous health checks using three probes:
- Ping to ping.ui.com.
- DNS resolution of ui.com using Cloudflare (1.1.1.1).
- DNS resolution of ui.com using Google (8.8.8.8).
A WAN is marked down if any two out of three probes fail. This is the default “Auto” SLA and is sufficient for most networks.
Custom SLAs & Verification Servers
UniFi offers flexibility to define your own custom SLAs & Failover criteria, instead of the “Auto” policy described above. To do so, navigate to Settings > Internet > Create New SLA, and then input your Verification Servers, and specify the Server Conditions–in other words, how many verification servers must fail before triggering a failover event. Each verification server can have the following parameters specified:
- Type: Use Ping or DNS when sending traffic to the verification server.
- Server Address: The IP Address or hostname of the verification server.
- Ping (Query) Interval: The number of seconds between ping (DNS) probes to the server.
- Time Period: The interval over which probes are averaged when evaluating SLA thresholds for failover.
- 
SLA Targets: The thresholds that will trigger a failover event. SLA targets are calculated based on the above Ping (Query) Interval, and the Time Period used in averaging the proves.
 
  - Packet Loss Threshold: The percentage of packets sent to the verification server that are lost.
  - Latency Threshold: The time it takes, in ms, for traffic to reach the verification server.
  - Jitter Threshold: The variation in time it takes, in ms, to reach the verification server.
- Packet Loss Threshold: The percentage of packets sent to the verification server that are lost.
- Threshold Condition: Specifies the above threshold(s) that correspond to an SLA not being met, and thus this specific verification server “failing”.ore granular behavior, such as detecting degraded performance before a full outage occurs.
Remapping WAN and LAN Ports
UniFi Gateways let you reassign which physical ports are used for WAN or LAN, allowing you to match your environment’s cabling or hardware preferences (e.g., RJ45 vs. SFP).
To remap ports:
- Go to Settings > Internet.
- Make sure the target port is already physically connected and active.
- Assign the new WAN role to that port.
- If swapping WAN1 and WAN2, temporarily assign a third port to avoid conflicts.
- Apply changes.
If a WAN appears offline after remapping, double-check ISP settings such as VLAN tags, static IP, or PPPoE credentials.
Using UniFi LTE for Backup Internet
UniFi LTE devices can serve as backup WANs by connecting to a LAN port on your gateway. Because they don’t occupy a dedicated WAN port, they don’t count toward your gateway’s WAN limit, and can be placed where cellular signal is strongest.
Learn how to set up UniFi LTE Backup
