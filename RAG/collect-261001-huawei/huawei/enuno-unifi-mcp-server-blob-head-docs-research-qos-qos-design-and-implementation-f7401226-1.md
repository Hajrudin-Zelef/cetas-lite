---
id: collect-261001-huawei/huawei/enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226-1
title: "enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "throughput", "voice"]
source: docs/RAG/collect-261001-huawei/enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226.md
source_anchor: ""
source_lines: [1, 82]
sha256: 2e2152e814cb89d7caa7634be44f8f6069e2089f47e41266c5fe24b428b0f079
---

# enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226

This report provides a comprehensive technical framework for implementing Quality of Service (QoS) and Smart Queue Management (SQM) on UniFi infrastructure. It addresses the critical design decisions network engineers face when deploying QoS across mixed-use environments—including VoIP, video conferencing, gaming, streaming, and bulk transfers—and presents a reusable library of QoS profile archetypes suitable for IaC-style templating across multiple deployments.
Key Recommendations Summary:
| Deployment Scenario | Primary Mechanism | Rate Cap | Key Consideration | 
|---|---|---|---|
| WAN < 300 Mbps | Smart Queues (fq_codel) | 85–95% of tested line rate | Eliminates bufferbloat; ~5% CPU overhead | 
| WAN 300 Mbps – 1 Gbps | Policy-based QoS or SQM (situational) | N/A or 90–95% | Hardware offload disabled with QoS rules; test throughput impact | 
| WAN > 1 Gbps | DSCP/CoS-only + VLAN segmentation | None at gateway | QoS rules cause 25–45% throughput reduction[web: 1][web: 35] | 
| Dedicated voice/AV | VLAN + DSCP EF 46 + switch port priority | Per-application | End-to-end marking consistency required | 
Smart Queues implement the fq_codel (FlowQueue Controlled Delay) algorithm at the WAN edge[web: 10][web: 12]. This mechanism:
- Isolates flows using a hashing scheme and schedules them via Deficit Round-Robin (DRR)
- Automatically manages queue depth to minimize latency under load
- Provides implicit priority for sparse (low-bandwidth) flows over bulk transfers
- Shifts the bottleneck from ISP equipment to the local gateway, where intelligent queue management can occur[web: 10]
Ubiquiti explicitly recommends Smart Queues only for connections below 300 Mbps[web: 5][web: 45]. On modern hardware (UCG-Fiber, UDM-Pro, UDM-SE, UXG-Pro), testing demonstrates effective operation at gigabit speeds with minimal CPU overhead, though the 300 Mbps warning persists in the UI[web: 10][web: 51].
fq_codel vs. CAKE: UniFi currently implements fq_codel exclusively; CAKE (Common Applications Kept Enhanced) is not available despite community feature requests[web: 20]. CAKE offers additional capabilities including per-host fairness, ACK filtering, and simplified link-layer compensation, but requires ~30% more CPU resources[web: 108].
UniFi Network 9.x introduced gateway-based QoS rules via the Policy Engine, offering three objectives[web: 1]:
| Objective | Function | Use Case | 
|---|---|---|
| Prioritize | Moves traffic to higher-priority queue | Voice, video conferencing, gaming | 
| Limit | Enforces maximum up/down bandwidth | Backup traffic, software updates, guest networks | 
| Prioritize and Limit | Combines priority with bandwidth ceiling | Guaranteed bandwidth for critical apps with caps | 
Critical Constraint: Enabling any QoS rule disables hardware offloading on the gateway, reducing throughput by 24–45% for traffic exceeding 1 Gbps[web: 1][web: 31][web: 35]. This applies globally—not just to traffic matched by rules.
UniFi switches support Layer 2/3 QoS through port-level configuration[web: 1][web: 18]:
- Match Type: DSCP value, IP Precedence, or None (all traffic)
- Remark Traffic: Optionally rewrite DSCP or CoS on egress
- Queue Assignment: Map matched traffic to queues 0–7 (0 = lowest, 7 = highest priority)
For Pro/Enterprise switches, Pro AV profiles provide pre-configured DSCP/CoS mappings for Dante, Q-SYS, SDVoE, NDI, AES67, and Shure environments[web: 18][web: 94].
WMM (Wi-Fi Multimedia) provides airtime-based prioritization using four access categories[web: 117]:
| Access Category | 802.1p Priority | Traffic Type | DSCP Mapping | 
|---|---|---|---|
| AC_VO (Voice) | 6, 7 | VoIP, real-time | EF (46), CS6, CS7 | 
| AC_VI (Video) | 4, 5 | Video streaming, conferencing | AF41 (34), CS4, CS5 | 
| AC_BE (Best Effort) | 0, 3 | Default web, email | CS0 (0), AF2x | 
| AC_BK (Background) | 1, 2 | Bulk transfers, updates | CS1, AF1x | 
UniFi APs honor DSCP markings and map EF (46) traffic to WMM voice priority by default[web: 77]. WiFi Speed Limits are per-client bandwidth caps applied at the AP level—distinct from QoS prioritization[web: 70].
A trust boundary defines where the network honors or overwrites incoming DSCP/CoS markings[web: 55]. Best practices:
- Trust controlled devices (IP phones, managed endpoints) at the access layer
- Do not trust PC traffic passively—remark or assign default markings
- Maintain consistent marking end-to-end; avoid gratuitous re-marking in the core[web: 55]
┌─────────────────────────────────────────────────────────────┐
│                    WAN Bandwidth Assessment                 │
└─────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┼───────────────┐
              │               │               │
         < 300 Mbps     300 Mbps–1 Gbps    > 1 Gbps
              │               │               │
              ▼               ▼               ▼
     ┌────────────────┐ ┌───────────────┐ ┌─────────────────────┐
     │ Smart Queues   │ │ Hybrid/Test   │ │ DSCP/VLAN Only      │
     │ (Primary)      │ │ SQM or Policy │ │ No gateway shaping  │
     └────────────────┘ └───────────────┘ └─────────────────────┘
              │               │               │
              ▼               ▼               ▼
     Additional controls if  Measure HW      Trust DSCP at switches;
     needed: policy-based    offload impact  shape only at true
     QoS for specific apps   before deploy   bottleneck (ISP modem)
Methodology for determining optimal SQM rates:
- Baseline measurement: Run multiple speed tests without SQM at different times; record consistent achievable rates
- Initial configuration: Set down/up rates to 85–90% of measured speeds[web: 5][web: 64]
- Bufferbloat validation: Test at waveform.com/tools/bufferbloat; target A+ grade (< 5ms added latency under load)[web: 132][web: 135]
- Iterative tuning: Increase rates in 2–5% increments until bufferbloat grade degrades; step back one increment
Reference tuning example (1 Gbps fiber):[web: 10]
| Stage | SQM Caps (Down/Up) | Achieved Speed | Latency Under Load | 
|---|---|---|---|
| Baseline (No SQM) | N/A | 1,055/960 Mbps | Download: 89ms, Upload: 7ms | 
| Conservative | 950/900 Mbps | 863/817 Mbps | 4ms/4ms | 
| Optimized | 1,000/930 Mbps | 909/844 Mbps | 4ms/4ms | 
For variable-rate connections (cable, DSL), use 80–85% to accommodate rate fluctuation[web: 64].
Recommended architecture for VoIP deployments:
- Dedicated VLAN: Create a voice-only VLAN (e.g., VLAN 16) with appropriate IP scope[web: 2][web: 22]
- LLDP-MED provisioning: Enable voice VLAN advertisement on switch ports; compliant phones auto-negotiate VLAN membership[web: 54][web: 59]
- DSCP marking: Phones should mark RTP as EF (46), signaling as CS3 (24)[web: 74][web: 9]
- Switch port priority: Assign voice ports to High or Critical priority; enable DSCP trust[web: 11]
- SIP ALG: Disable SIP ALG on the gateway—it typically breaks SIP signaling[web: 34][web: 9]
Acceptable VoIP thresholds:[web: 127][web: 131]
| Metric | Acceptable | Optimal | 
|---|---|---|
| One-way latency | < 150ms | < 80ms | 
| Jitter | < 30ms | < 10ms | 
| Packet loss | < 1% | < 0.1% | 
Gaming and real-time interactive applications benefit from low and consistent latency rather than raw bandwidth. Recommended approach:
