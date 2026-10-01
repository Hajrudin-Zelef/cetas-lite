---
id: collect-261001-huawei/huawei/enuno-unifi-mcp-server-blob-head-docs-research-qos-unifi-qos-design-mcp-implemen-1794c459-1
title: "enuno-unifi-mcp-server-blob-head-docs-research-qos-unifi-qos-design-mcp-implemen-1794c459"
domain: huawei
role: reference
task: reference
actors: ["SpaceX"]
dates: []
keywords: ["mcp", "latency", "throughput", "voice"]
source: docs/RAG/collect-261001-huawei/enuno-unifi-mcp-server-blob-head-docs-research-qos-unifi-qos-design-mcp-implemen-1794c459.md
source_anchor: ""
source_lines: [1, 67]
sha256: 8bd5dfd7077e5ed521ef35da0206bc8172d3633023c7ff7d33fff15eabfcb2f4
---

# enuno-unifi-mcp-server-blob-head-docs-research-qos-unifi-qos-design-mcp-implemen-1794c459

Executive Summary
This report synthesizes vendor guidance, practitioner best practices, and real-world tuning data to deliver a production-ready QoS framework for UniFi Network (9.x+) environments. The architecture addresses the critical decision points between Smart Queues (SQM), policy-based QoS rules, and DSCP-only designs, providing reusable profile templates optimized for multi-tenant MSP deployments. Key findings include: Smart Queues should be limited to WAN links under 300 Mbps at 80–95% of measured line rate; DSCP EF 46 marking for RTP voice must be preserved end-to-end; and over-prioritization across all traffic classes negates QoS benefits. The recommended profile library covers six core traffic classes with specific DSCP mappings, bandwidth allocations, and UniFi implementation paths.
UniFi Network provides three primary QoS mechanisms that operate at different layers of the stack. Understanding their interaction is essential for coherent policy design.
Smart Queues implement fq_codel (Fair Queuing with Controlled Delay) at the WAN edge to combat bufferbloat. The mechanism:123
- Shapes egress traffic to configured up/down rates
- Automatically prioritizes small, latency-sensitive packets (DNS, VoIP, gaming)
- Requires disabling hardware offload on legacy platforms (USG series) but runs efficiently on modern gateways (UDM/UXG series)4
- Recommended only for ISP connections <300 Mbps; counterproductive on high-bandwidth links due to unnecessary queuing delays51
Policy-based rules (Network 9.x) allow explicit Prioritize, Limit, or Prioritize and Limit actions on traffic matching source/destination criteria. Key capabilities:6
- Match by device, network (VLAN), application, IP address/port, domain, or region
- Enforce bandwidth caps per direction (Mbps)
- Higher granularity than Smart Queues but requires manual classification
- Performance impact scales with rule complexity; test throughput ceilings after activation7
DSCP values (0–63) mark packets for per-hop behavior across switches and access points:8
- EF (46): Expedited Forwarding for VoIP RTP (strict priority)
- AF41 (34): Assured Forwarding for interactive video
- CS0 (0): Default best-effort
- UniFi switches and APs automatically map DSCP to WMM access categories (AC_VO, AC_VI, AC_BE, AC_BK) when QoS is enabled on the port/profile910
Port-level QoS assignment complements gateway policies:
- Port Priority (High/Critical) elevates all traffic on a switch port
- Voice VLANs isolate phone traffic and enable LLDP-MED auto-configuration11
- Trust boundaries must be defined: either preserve endpoint DSCP or remark at switch ingress for untrusted devices8
| WAN Speed | Primary Use Case | Recommended Approach | Rationale | 
|---|---|---|---|
| <100 Mbps | Residential, small office, VoIP-heavy | Smart Queues at 85–90% of tested rate | Bufferbloat dominates; automatic fairness reduces latency spikes 1213 | 
| 100–300 Mbps | SMB, mixed-use environments | Smart Queues at 80–95% rate OR QoS rules for voice + Smart Queues for bulk | 权衡：Smart Queues add latency but simplify management; policy QoS preserves throughput for non-congested links 12 | 
| >300 Mbps | High-performance, MSP backbone | Policy QoS only (no Smart Queues) | SQM overhead unnecessary; shaping at non-bottleneck is ineffective 15 | 
| Variable (cellular, shared fiber) | Remote sites, pop-up offices | Adaptive Smart Queues at 80% of minimum observed rate + DSCP trust | Accommodates link variability; prevents queue overruns during troughs 142 | 
Table: WAN speed-based QoS strategy selection
Community consensus and fq_codel best practices recommend setting Smart Queue rates 10–20% below measured ISP throughput to account for variable overhead and upstream ISP buffering:12152
- Stable fiber/coax: 90–95% of sustained speedtest results
- Variable cable/DSL: 80–85% of off-peak measurements
- Cellular/Starlink: 70–80% of median observed rates
- Always verify with real-world tests (speedtest, bufferbloat.net) after enabling1612
Preserve end-to-end DSCP markings to leverage hardware queueing in switches and APs:178
- Trust endpoint markings for enterprise phones and video conferencing systems (most mark RTP correctly)
- Remark at switch for untrusted IoT or guest devices to prevent DSCP spoofing
- Avoid mid-path remarking unless converting between domains (e.g., ISP-facing EF remapping)
- Use VLANs as primary trust boundary: separate voice, video, guest, and bulk traffic into distinct broadcast domains with consistent DSCP policies1816
Disable SIP ALG on UniFi gateways for 95% of modern VoIP platforms (SIP ALG routinely breaks signaling). Restrict port forwards to required UDP/TCP ranges only; avoid blanket "DMZ" configurations. Use IP Group objects in QoS rules to target PBX/trunk IPs precisely.16
Apply bandwidth limits exclusively where congestion occurs—typically WAN upstream on asymmetric links. LAN-side shaping between VLANs is rarely necessary on gigabit UniFi switches and introduces unnecessary complexity.1916
Assigning Critical priority to multiple traffic classes creates contention and negates QoS benefits. Limit the number of high-priority classes to 2–3 per deployment (e.g., voice + gaming, or voice + video). Use Limit objectives for bulk traffic rather than deprioritizing into low queues.719
Enabling QoS rules disables hardware offload on legacy platforms, capping throughput to ~250–300 Mbps on USG/ER-X devices. Modern UXG/UDM series incur <1% CPU overhead. Always conduct post-implementation speed tests to validate gateway sizing.204
Test QoS efficacy under load:
- Bufferbloat: Use waveform.bufferbloat.net or DSL Reports speedtest (target A+ grade)
- VoIP: Concurrent calls + bulk upload (e.g., iCloud backup) → measure MOS score and packet loss
- Gaming: Saturate link with downloads while monitoring ping/jitter to game servers
- Video: 4K stream + software update download → verify playback continuity1216
Maintain a profile registry (YAML/JSON) with:
- Profile name, version, target environment
- WAN speed range, Smart Queue rates, QoS rule UUIDs
- DSCP mappings, VLAN IDs, port profiles
- Validation test results and known limitations
The following six profiles provide reusable archetypes for MCP server deployments. Each profile specifies intent, traffic class focus, UniFi implementation details, and recommended bandwidth allocations.
| Profile Name | Traffic Class Focus | Smart Queues | DSCP Tags | VLAN Strategy | Bandwidth Limits | Primary UniFi Knobs | 
|---|---|---|---|---|---|---|
| Voice-First | VoIP RTP, SIP signaling | Optional (<300 Mbps) | EF 46 (RTP), CS3 (SIP) | Dedicated voice VLAN | Reserve 30% upstream for voice | QoS rule (Prioritize voice VLAN), Smart Queues at 85% | 
| Video-Conferencing | Zoom, Teams, WebRTC | Optional (<300 Mbps) | AF41 34, EF 46 (if phone) | Trusted devices VLAN | Min 3 Mbps per concurrent HD stream | QoS rule (Prioritize AF41), DSCP trust on switch | 
| Cloud-Gaming | Stadia, GeForce Now, XCloud | Recommended (<300 Mbps) | EF 46 | Gaming device group | Min 15 Mbps down, 5 Mbps up | QoS rule (Prioritize gaming IPs), Smart Queues | 
| Streaming-Media | Netflix, YouTube 4K | Not recommended | AF41 34, AF31 26 | Media player VLAN | Limit bulk to 50% during peak | QoS rule (Limit Windows Update, etc.) | 
| Bulk-Backup | iCloud, OneDrive, Backblaze | Not recommended | CS1 8 (Scavenger) | Separate backup VLAN | Cap at 30% of WAN, de-prioritize | QoS rule (Limit backup VLAN to 30 Mbps) | 
| Guest-Best-Effort | Guest Wi-Fi, IoT | Not recommended | CS0 0 (default) | Guest VLAN only | Hard limit 5–10 Mbps per client | QoS rule (Limit guest network), port isolation | 
Table: Core QoS profile library for MCP deployments
- Intent: Ensure toll-quality VoIP regardless of background load
- Target: SMB offices, call centers, home offices with hosted PBX
- Implementation:
  - Create dedicated Voice VLAN (e.g., VLAN 110)
