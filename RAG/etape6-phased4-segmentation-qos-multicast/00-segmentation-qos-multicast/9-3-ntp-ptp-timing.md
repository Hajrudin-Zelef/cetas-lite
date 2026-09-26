---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/9-3-ntp-ptp-timing
title: "9.3 NTP/PTP timing"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [246, 256]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 1478c0219be500e7c5399336006e738d123a0bcd3df9ae334cac67489dbf7d36
---

# 9.3 NTP/PTP timing

### 9.3 NTP/PTP timing
- NTP: hierarchical stratum model, ms-level accuracy — sufficient for logging/correlation; every fabric device + hypervisor should sync to redundant stratum-1/2 sources `[secondary]` — standard practice.
- PTP (IEEE 1588): sub-microsecond sync via boundary clocks / transparent clocks on switches; required for financial trading (MiFID II timestamping), 5G fronthaul, and some AI-collective telemetry `[secondary]` — widely documented; Arista EOS and NX-OS both document PTP boundary/transparent clock modes `[official]` (per-vendor docs, not deep-linked here).
- DC relevance: PTP in the underlay for telemetry correlation; NTP everywhere as baseline; never mix unauthenticated NTP across security zones without NTS/authentication where required `[secondary]`.
- Gap: per-platform PTP profile support (G.8275.1 vs default E2E/P2P) matrix not compiled — open item `[unverified]`.

### Wave 9 verification
- Sources: 5. Cisco 17.5.1/17.6.1 DHCP-relay version boundary captured verbatim. DNS/PTP sections are practice-level; flagged where no authoritative doc was consulted.

---

