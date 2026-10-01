---
id: collect-261001-huawei/huawei/enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226-2
title: "enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "throughput", "voice"]
source: docs/RAG/collect-261001-huawei/enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226.md
source_anchor: ""
source_lines: [83, 181]
sha256: f4a4d3098dfa49e2b633800ecbdce651c161a8b74dccb0efca7cdd6acb688f62
---

# enuno-unifi-mcp-server-blob-head-docs-research-qos-qos-design-and-implementation-f7401226

- Do not over-prioritize: Gaming typically generates sparse traffic flows; fq_codel inherently prioritizes sparse flows[web: 105]
- Smart Queues sufficient: For sub-300 Mbps links, SQM alone dramatically improves gaming latency[web: 10][web: 57]
- Explicit prioritization: If needed, create QoS rules targeting gaming devices or applications with Prioritize objective
- DSCP marking: Consider marking gaming traffic as EF (46) or AF41 if endpoint supports it[web: 11]
Explicit de-prioritization of bulk transfers (backups, software updates, P2P) is often more effective than trying to prioritize everything else:
- Limit objective: Apply bandwidth limits (e.g., 75% of WAN) to file transfer and P2P categories[web: 5]
- Schedule restrictions: Apply rules during business hours only; allow full bandwidth overnight
- Background class: Mark backup traffic as AF1x or CS1 for background handling
Principle: QoS shaping is effective only when applied at the actual congestion point—typically the WAN uplink/downlink.
- Avoid double shaping: Do not apply Smart Queues at the gateway AND rate limits at the AP for the same traffic
- ISP equipment: If ISP modem buffers heavily, shaping must occur before traffic hits that device
- Internal links: LAN-side shaping is rarely needed unless specific inter-VLAN bottlenecks exist
| Guideline | Recommendation | Rationale | 
|---|---|---|
| Rate cap | 85–95% of tested WAN speed | Creates artificial bottleneck at gateway where fq_codel manages queues[web: 10][web: 14] | 
| Hardware threshold | Not recommended > 300 Mbps (vendor guidance) | CPU-intensive; newer hardware can exceed this[web: 45] | 
| Validation | A+ bufferbloat score | Confirms latency under load is < 5ms[web: 132] | 
| PPPoE consideration | Test carefully | PPPoE + SQM doubles CPU load; may reduce throughput significantly[web: 10] | 
| Traffic Class | DSCP Value | PHB | Use Case | 
|---|---|---|---|
| Voice (RTP) | 46 | EF | VoIP media streams | 
| Video conferencing | 34 | AF41 | Zoom, Teams, WebRTC video | 
| Call signaling | 24 | CS3 | SIP, H.323, SRTP signaling | 
| Streaming media | 26 | AF31 | 4K OTT, adaptive bitrate | 
| Network control | 48 | CS6 | Routing protocols, SNMP | 
| Best effort | 0 | CS0 | Default web, email | 
| Bulk/background | 10 | AF11 | Backups, updates, sync | 
End-to-end consistency: Configure switches to trust DSCP at ingress ports connected to trusted devices (phones, APs, known endpoints)[web: 55][web: 46].
Recommendation: Disable SIP ALG in all deployments unless explicitly required by VoIP vendor[web: 34][web: 9].
UniFi path: Settings → Application Firewall → Firewall Rules → Conntrack Modules → Disable SIP[web: 40]
SIP ALG rewrites SIP headers and often corrupts RTP stream negotiation, causing one-way audio, dropped calls, and registration failures.
| Approach | Pros | Cons | Recommendation | 
|---|---|---|---|
| Class-based (VLAN, DSCP, device group) | Deterministic, scalable, DPI-independent | Requires endpoint marking or manual classification | Preferred for enterprise | 
| Application-based (DPI signatures) | Automatic identification, no endpoint config | DPI CPU overhead, signature lag, encrypted traffic challenges | Supplement only | 
Guidance: Use VLAN and DSCP as primary classification; supplement with DPI-based rules for specific applications where endpoint marking is impossible[web: 9].
The following profiles represent reusable archetypes for IaC-style deployment templates. Each profile specifies intent, target environment, and UniFi implementation details.
| Profile Name | Traffic Class | Target Environment | Primary Mechanisms | Key UniFi Settings | 
|---|---|---|---|---|
| VOICE-FIRST | VoIP | Any with desk phones | Voice VLAN, DSCP EF, switch priority | VLAN + Port Profile + QoS rule (prioritize) | 
| VIDEO-COLLAB | Video conferencing | Remote work, hybrid office | DSCP AF41, QoS prioritize | QoS rule on Zoom/Teams/Meet | 
| GAMING-INTERACTIVE | Gaming, remote desktop | Home office, SMB | Smart Queues, sparse flow priority | SQM enabled; optional prioritize rule | 
| STREAMING-MEDIA | 4K OTT, adaptive streaming | Residential, hospitality | Best effort or AF31 | WiFi speed limit; no explicit priority | 
| BULK-BACKGROUND | Backups, updates, sync | All environments | De-prioritize, rate limit | QoS rule (limit) on file transfer category | 
| GUEST-BESTEFFORT | Guest WiFi | Retail, hospitality, office | Isolated VLAN, rate limit | SSID isolation + WiFi speed limit | 
Intent: Ensure voice calls experience < 150ms latency, < 30ms jitter, and < 1% packet loss regardless of concurrent traffic.
Target Environment: SMB office, MSP client sites, any deployment with dedicated VoIP phones.
Implementation:
| Component | Configuration | 
|---|---|
| VLAN | Dedicated voice VLAN (e.g., VLAN 16); isolate from data | 
| Switch port profile | Voice VLAN assignment + LLDP-MED enabled | 
| DSCP handling | Trust DSCP on voice ports; phones mark EF (46) for RTP, CS3 (24) for signaling | 
| Port priority | High or Critical on voice ports | 
| Gateway QoS | Prioritize rule targeting voice VLAN or phone IP group | 
| Smart Queues | Enable if WAN < 300 Mbps | 
| SIP ALG | Disabled | 
| Firewall | Allow UDP 5060 (SIP), 10000–20000 (RTP) inbound to voice VLAN | 
Validation: Place concurrent bulk transfers and VoIP test calls; measure MOS score > 4.0.
Intent: Maintain smooth video conferencing (Zoom, Teams, Meet, WebRTC) during network congestion.
Target Environment: Remote workforce, hybrid office, education.
Implementation:
| Component | Configuration | 
|---|---|
| QoS rule | Prioritize objective; destination: App category "Video Conferencing" | 
| DSCP | Mark video as AF41 (34) if endpoint supports | 
| Bandwidth burst | Short (allows brief spikes for video keyframes) | 
| Smart Queues | Enable if WAN < 300 Mbps | 
| WiFi | Ensure WMM enabled (default); consider dedicated SSID for corporate devices | 
Note: Most video conferencing applications adapt bitrate dynamically; prioritization helps during congestion but bandwidth headroom is equally important[web: 80].
Intent: Minimize latency and jitter for real-time gaming and remote desktop applications.
Target Environment: Home office, residential, creative studios with remote workstations.
Implementation:
| Component | Configuration | 
|---|---|
| Smart Queues | Primary mechanism; fq_codel inherently prioritizes sparse flows | 
| Rate cap | 90% of tested WAN speed for stable connections; 85% for variable | 
| QoS rule (optional) | Prioritize specific gaming consoles or PCs if contention persists | 
| DSCP | Mark gaming traffic EF (46) if router supports | 
| WiFi | Prefer 5 GHz; minimize AP hop count; enable band steering | 
Validation: Run bufferbloat tests during gaming sessions; target A+ grade.
Intent: Deliver consistent 4K streaming without impacting higher-priority traffic.
Target Environment: Residential, hospitality, common areas.
Implementation:
| Component | Configuration | 
|---|---|
| Priority | Best effort (CS0) or AF31 if explicit marking desired | 
| QoS rule | None or Limit (cap at 50–75% WAN if contention exists) | 
| WiFi speed limit | Per-client cap if shared environment (e.g., 50 Mbps) | 
| Smart Queues | Enable; streaming adapts well to available bandwidth | 
Rationale: Adaptive bitrate streaming (HLS, DASH) handles congestion gracefully; explicit prioritization rarely needed.
Intent: Explicitly de-prioritize backups, software updates, and sync traffic to prevent impact on interactive applications.
Target Environment: All environments; particularly effective when combined with voice/video profiles.
Implementation:
| Component | Configuration | 
|---|---|
| QoS rule | Limit objective; destination: Categories "File Transfer," "Peer-to-Peer," "Software Updates" | 
| Bandwidth cap | 50–75% of WAN during business hours | 
| Schedule | Restrict limits to 08:00–18:00 weekdays; full bandwidth overnight | 
