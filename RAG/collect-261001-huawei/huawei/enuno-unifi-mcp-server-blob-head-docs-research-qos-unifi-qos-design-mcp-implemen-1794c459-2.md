---
id: collect-261001-huawei/huawei/enuno-unifi-mcp-server-blob-head-docs-research-qos-unifi-qos-design-mcp-implemen-1794c459-2
title: "enuno-unifi-mcp-server-blob-head-docs-research-qos-unifi-qos-design-mcp-implemen-1794c459"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["mcp", "agents", "energy", "ethernet", "latency", "throughput", "voice"]
source: docs/RAG/collect-261001-huawei/enuno-unifi-mcp-server-blob-head-docs-research-qos-unifi-qos-design-mcp-implemen-1794c459.md
source_anchor: ""
source_lines: [68, 135]
sha256: 0e91f8355e764d12468c0da095e9ee3ba5b724cde99cda4cc136ebdbeea5a3e5
---

# enuno-unifi-mcp-server-blob-head-docs-research-qos-unifi-qos-design-mcp-implemen-1794c459

  - Enable QoS tagging on voice network: UniFi Network → Settings → Networks → [Voice VLAN] → Advanced → QoS Priority = High
  - Configure QoS rule: Objective = Prioritize, Source = Network (Voice VLAN), Destination = Any
  - If WAN <300 Mbps: Enable Smart Queues at 85% of tested rate to prevent bufferbloat from bulk uploads
  - On switches: Set port priority to Critical for phone ports; trust DSCP EF 46 from phones
- DSCP Handling: Preserve EF 46 from endpoints; remark untrusted devices to CS0
- Validation: 10 concurrent calls + 50 Mbps upload → MOS >4.0, zero packet loss
- Intent: Stable HD video for remote work; minimize freezes and "your connection is unstable"
- Target: Remote workers, executive home offices, hybrid meeting rooms
- Implementation:
  - Use Trusted Devices VLAN for work laptops and room systems
  - QoS rule: Objective = Prioritize, Source = Device Group (video endpoints), Destination = App (Zoom, Teams, Webex)
  - Match DSCP AF41 from conferencing apps; if missing, remark at switch ingress
  - Smart Queues optional at 90% rate for links <200 Mbps; omit for gigabit fiber
- Bandwidth: Reserve 3 Mbps per HD stream upstream; 1.5 Mbps for 720p
- Validation: 3 concurrent meetings + bulk download → <2% frame loss, latency <100 ms
- Intent: Low latency, jitter, and packet loss for cloud gaming platforms
- Target: Residential gaming, esports venues, gaming lounges
- Implementation:
  - Create Gaming Device Group (by MAC or static IP)
  - QoS rule: Objective = Prioritize, Source = Device Group, Destination = IP ranges (Stadia, GeForce Now, XCloud)
  - Enable Smart Queues at 85–90% of WAN rate; fq_codel automatically prioritizes small packets
  - On switches: Set gaming ports to High priority; disable EEE (Energy Efficient Ethernet) to reduce latency
- DSCP: Mark as EF 46 if platform supports it; trust endpoint markings
- Validation: Gaming latency <50 ms, jitter <10 ms while saturating link with downloads
- Intent: Ensure 4K streaming quality without buffering; deprioritize background bulk
- Target: Media rooms, residential entertainment networks
- Implementation:
  - No Smart Queues (streaming benefits from high throughput, not shaping)
  - QoS rule: Objective = Limit, Source = Any, Destination = App (Windows Update, iCloud Backup, OneDrive) → cap at 50% of downstream during peak hours
  - Use Schedule to allow unlimited bulk during off-peak (e.g., 2 AM–6 AM)
  - VLAN optional; prioritize by application signature rather than network segment
- DSCP: Use AF41 for streaming devices if switch supports remarking; otherwise rely on endpoint defaults
- Validation: 4K stream maintains 25 Mbps sustained; background updates do not cause buffering
- Intent: Contain backup/sync traffic to prevent WAN saturation and latency spikes
- Target: Remote offices with cloud backup, MSP-managed endpoints
- Implementation:
  - Create Backup VLAN (e.g., VLAN 900) or Device Group for backup agents
  - QoS rule: Objective = Limit, Source = Network (Backup VLAN), Destination = Any → 30 Mbps down, 10 Mbps up (adjust per link)
  - Mark traffic as DSCP CS1 (Scavenger) at switch ingress to ensure de-prioritization if it exits via alternate path
  - Disable Smart Queues; shaping already applied via rate limit
- Schedule: Allow higher limits during maintenance windows (e.g., weekends)
- Validation: Bulk transfer capped at defined rate; ping to 8.8.8.8 remains <20 ms during backup
- Intent: Isolate guest traffic; prevent guest devices from impacting business or primary traffic
- Target: Guest Wi-Fi, IoT networks, untrusted devices
- Implementation:
  - Guest VLAN with Hotspot or WPA2-Enterprise
  - QoS rule: Objective = Limit, Source = Network (Guest VLAN), Destination = Any → 5–10 Mbps per client, 50 Mbps aggregate
  - Enable Port Isolation on switches to prevent guest-to-guest communication
  - No DSCP trust; remark all guest traffic to CS0 at switch ingress
  - Disable inter-VLAN routing to guest VLAN (layer-2 only)
- Validation: Guest speedtest shows capped rate; business VLAN throughput unaffected
Smart Queues Throughput Ceiling: Legacy USG/ER-X platforms experience drastic throughput reduction (to ~280 Mbps) when Smart Queues enable software-forwarding paths. Modern UXG-Pro/UDM-Pro handle SQM with <1% CPU overhead, but rates >300 Mbps still render SQM unnecessary.14520
QoS Rules and Offload: Enabling any QoS rule disables hardware offload on pre-UDM platforms. Validate gateway CPU capacity under peak load; consider upgrade if sustained throughput drops below business requirements.217
Switch Model Limitations: Not all UniFi switches support DSCP remarking in current firmware. Verify feature support (Settings → Profiles → Ethernet Ports → Manual → Enable QoS) before designing trust-boundary policies.8
Smart Queues + QoS Rules: Concurrent use is supported but requires careful ordering. Smart Queues shape first; QoS rules apply within shaped bandwidth. Avoid double-shaping (e.g., Smart Queues at 90 Mbps + QoS rule limiting device to 80 Mbps) as it creates unpredictable behavior.2212
Per-SSID Bandwidth Profiles: Wi-Fi bandwidth limits (Settings → WiFi → [SSID] → Bandwidth Profile) operate independently of gateway QoS. Use these for airtime fairness, not WAN shaping. Combining SSID limits with Smart Queues can lead to unintended starvation.2312
DSCP Trust and WMM: UniFi APs automatically map DSCP to WMM queues. However, if QoS is disabled on the SSID or AP radio, DSCP is ignored and all traffic falls into Best Effort. Ensure QoS is enabled on all SSIDs carrying prioritized traffic.109
- Over-subscribing priority queues: Marking video, gaming, voice, and critical apps all as "Critical" results in none receiving true priority. Limit high-priority classes to ≤3.197
- Mismatched DSCP trust: Phones mark EF 46, but switch remarks to CS0; or switch trusts DSCP but gateway QoS rule matches on port only. Align trust boundaries end-to-end.178
- Smart Queues on gigabit fiber: Unnecessary latency injection; policy QoS alone suffices. Disable Smart Queues for >300 Mbps circuits.51
- Ignoring SIP ALG: SIP ALG on UniFi gateways breaks SIP signaling for most modern platforms. Disable under Settings → Internet → [WAN] → Advanced → SIP ALG.16
- Setting Smart Queues to 100% ISP rate: Fails to absorb ISP bufferbloat; always set 10–20% below tested rate.1312
Continuous Monitoring: Use UniFi Insights → Real-Time Traffic to verify QoS rule hits and bandwidth consumption. Graph Smart Queue drops and latency metrics if available.
Bufferbloat Testing: Monthly waveform.bufferbloat.net tests; target A+ grade with <20 ms latency increase under load. If grade drops, adjust Smart Queue rates downward by 5% increments.2412
VoIP MOS Scoring: Deploy synthetic call probes (e.g., ThousandEyes, PingPlotter) to measure MOS, packet loss, and jitter during peak hours. Target MOS >4.0, loss <0.5%, jitter <30 ms.16
Change Control: Version profiles in Git; apply via UniFi API or Settings → Advanced → Backup/Restore. Document WAN speed, Smart Queue rates, and test results per site.
For MCP server rollouts, standardize on the six core profiles defined in Section 4.2, selecting the appropriate profile based on site function (office, residential, gaming lounge). Always shape at the WAN edge using Smart Queues for links <300 Mbps or policy QoS for higher bandwidth. Preserve DSCP EF 46 for voice end-to-end; trust endpoint markings where possible. Disable SIP ALG and restrict port forwards to required ranges only. Validate each deployment with bufferbloat and synthetic VoIP testing, and maintain versioned profile documentation in infrastructure-as-code repositories.
The profile library provides a foundation for repeatable, supportable QoS policies across diverse UniFi deployments, balancing latency sensitivity for real-time applications with bandwidth availability for bulk transfers.
