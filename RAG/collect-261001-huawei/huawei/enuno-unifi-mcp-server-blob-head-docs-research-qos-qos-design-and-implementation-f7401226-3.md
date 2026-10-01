---
id: collect-261001-huawei/huawei/enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226-3
title: "enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226"
domain: huawei
role: reference
task: reference
actors: []
dates: ["2025-11-15"]
keywords: ["containment", "latency", "liability", "throughput", "voice"]
source: docs/RAG/collect-261001-huawei/enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226.md
source_anchor: ""
source_lines: [182, 265]
sha256: 4a3e1a1734bf3c0d671a36f09008081841ea1fbc0f70ce3716151ad95080498c
---

# enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226

| DSCP | Mark as AF11 (10) or CS1 (8) if endpoint supports | 
Example rule:
- Objective: Prioritize and Limit
- Download/Upload: 75% of WAN
- Bandwidth burst: Short
- Destination: App categories (File Transfer, Peer-to-Peer, Software Updates)[web: 5]
Intent: Provide functional guest access without impacting production traffic or enabling abuse.
Target Environment: Retail, hospitality, enterprise visitor networks.
Implementation:
| Component | Configuration | 
|---|---|
| VLAN | Isolated guest VLAN; enable "Isolate Network" | 
| WiFi | Dedicated SSID with WiFi speed limit (e.g., 20–50 Mbps per client) | 
| Client isolation | Device Isolation (ACL) at switch level | 
| mDNS | Disabled for guest network | 
| Proxy ARP | Enabled to reduce broadcast traffic | 
| Traffic restrictions | Block file transfer and P2P categories | 
| QoS | No prioritization; best effort only | 
| Captive portal | Optional; consider for liability acknowledgment | 
Aggregate bandwidth (optional): Use QoS Limit rule on entire guest VLAN to cap total bandwidth (e.g., 25–50% of WAN)[web: 110][web: 113].
| Gateway | IDS/IPS Throughput | Smart Queues Impact | QoS Rule Impact | 
|---|---|---|---|
| USG | 85 Mbps | Significant; ~300 Mbps max | Not recommended | 
| USG Pro 4 | 250 Mbps | ~400 Mbps max | Hardware offload disabled | 
| UDM | 850 Mbps | Minimal | Throughput reduction likely | 
| UDM Pro | 3.5 Gbps | ~5% CPU overhead | 25–45% reduction > 1 Gbps[web: 35] | 
| UDM Pro Max | 5 Gbps | Minimal impact | Significant reduction at high rates | 
| UDM SE | 3.5 Gbps | Supported | Hardware offload disabled | 
| UCG-Fiber | 5 Gbps | Excellent results[web: 10] | Test required | 
| UCG-Ultra | 1 Gbps | Supported | Limited headroom | 
| UCG-Max | 2.3 Gbps | Supported | Test required | 
| UXG-Pro | 3+ Gbps | Dedicated gateway resources | Similar to UDM Pro | 
- Smart Queues: Available on all UniFi gateways; requires UniFi Network 5.x+
- QoS rules (Policy Engine): Requires UniFi Network 9.x+; zone-based firewall recommended for 9.4+[web: 4][web: 96]
- Pro AV profiles: Requires Network 8.4.59+ and supported switches (Pro/Enterprise lines) with firmware 7.1.26+[web: 18]
- DSCP switch QoS: Limited to Pro/Enterprise switches; not available on standard Lite/Flex models[web: 84]
| Misconfiguration | Symptom | Resolution | 
|---|---|---|
| Double shaping | Excessive latency, reduced throughput | Shape only at WAN edge; remove AP/VLAN limits if gateway shaping active | 
| DSCP trust mismatch | Priority not honored; voice quality issues | Verify trust at ingress ports; check end-to-end DSCP handling | 
| Over-prioritization | Nothing actually prioritized; all traffic equal | Prioritize only truly critical classes; de-prioritize bulk instead | 
| SIP ALG enabled | One-way audio, dropped calls, registration failures | Disable SIP ALG; verify NAT traversal works without it | 
| Smart Queues on PPPoE | Severe throughput reduction | Test carefully; PPPoE encapsulation overhead compounds CPU load | 
| QoS rules on high-speed WAN | 25–45% throughput loss | Use DSCP-only approach for > 1 Gbps; avoid gateway QoS rules | 
| WiFi speed limit vs. QoS confusion | Bandwidth capped but no prioritization | Speed limits are not QoS; use QoS rules for priority | 
Order of operations:
- WiFi speed limits apply per-client at the AP (air interface)
- QoS rules evaluate at the gateway for WAN-bound traffic
- Smart Queues shape aggregate WAN traffic after all other processing
Guidance:
- WiFi speed limits reduce what reaches the gateway; useful for guest containment
- QoS rules classify and queue traffic independent of WiFi limits
- Smart Queues provide aggregate fairness; effective even without explicit rules
- Avoid redundant limits: If QoS rule limits a VLAN to 100 Mbps, per-client WiFi limits within that VLAN may be unnecessary
| Test | Tool/Method | Target Result | 
|---|---|---|
| Bufferbloat | waveform.com/tools/bufferbloat | A+ grade (< 5ms latency under load) | 
| VoIP quality | Synthetic call generator, MOS scoring | MOS > 4.0 under concurrent load | 
| Concurrent transfers | iperf3 + latency monitoring | Voice/gaming latency stable during bulk transfers | 
| Real-world validation | Video call + backup running simultaneously | No perceivable quality degradation | 
Flent testing (advanced):[web: 50][web: 139]
flent rrul -p all_scaled -H flent-fremont.bufferbloat.net
Generates comprehensive latency/throughput charts under Real-time Response Under Load (RRUL) test.
Naming conventions for profiles:
| Element | Convention | Example | 
|---|---|---|
| QoS rule | [ENV]-[CLASS]-[ACTION] | SMB-VOICE-PRIORITIZE | 
| VLAN | [SITE]-[PURPOSE]-[ID] | HQ-VOICE-16 | 
| Switch port profile | [CLASS]-PORT | VOIP-PORT | 
| WiFi speed limit | [LIMIT]-[TARGET] | 50M-GUEST | 
Version control concepts:
- Maintain profile definitions in version-controlled JSON/YAML (UniFi API v0.1 supports limited export)[web: 99][web: 102]
- Document baseline configurations per site type (home office, SMB, enterprise)
- Track changes with dated notes: "2025-11-15: Increased SQM cap 900→950 Mbps after ISP upgrade"
- Traffic Flows: UniFi Network 9.x provides real-time flow visibility (Insights → Flows)[web: 112]
- DPI statistics: Deep Packet Inspection identifies application traffic; useful for validating rule matches[web: 85][web: 126]
- Gateway CPU: Monitor when Smart Queues or QoS rules active; sustained > 80% indicates hardware limit
- Bufferbloat regression: Periodic testing recommended after firmware updates or ISP changes
UniFi provides a layered QoS architecture suitable for deployments ranging from home offices to multi-site MSP templates. The key design decisions center on:
- WAN bandwidth determines primary mechanism: Smart Queues for sub-300 Mbps; policy-based QoS with caution at mid-range speeds; DSCP/VLAN-only for high-throughput environments
- Hardware offload tradeoffs are real: QoS rules disable offloading with significant throughput implications above 1 Gbps
- End-to-end consistency matters: DSCP marking, trust boundaries, and switch queue mapping must align across the entire path
- Validation is essential: Bufferbloat testing, MOS scoring, and real-world concurrent load tests confirm QoS effectiveness
The profile library provided offers a foundation for reusable, IaC-ready templates that can be adapted to specific deployment requirements while maintaining consistent design principles across environments.
