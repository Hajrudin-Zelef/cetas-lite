---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-15-services-deep-dive-dhcpv6-slaac-dns-ntp-nts-ptp-prof
title: "Wave 15 — Services deep-dive: DHCPv6/SLAAC, DNS, NTP/NTS, PTP profiles"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["agents"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [426, 485]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 8e88026df5445f553a072eeaffb16317b52b5f2842ae0e0dc5637d41e83a53e9
---

# Wave 15 — Services deep-dive: DHCPv6/SLAAC, DNS, NTP/NTS, PTP profiles

## Wave 15 — Services deep-dive: DHCPv6/SLAAC, DNS, NTP/NTS, PTP profiles

### 15.1 DHCPv6 & SLAAC
- SLAAC (RFC 4862): hosts auto-configure from router advertisements — no server; fine for infrastructure addressing, insufficient where DNS/option control is needed `[official]` — RFC 4862.
- DHCPv6 (RFC 8415): stateful address + option assignment; relay agents (like DHCPv4) forward toward servers; in EVPN fabrics the same first-hop-relay and VRF-awareness rules apply as DHCPv4 `[official]` — RFC 8415; Cisco 17.6.1+ EVPN DHCP handling covers v4+v6 (Wave 9).
- Design note: many DCs run SLAAC for link/infra + DHCPv6 for options, or DHCPv4-only dual-stack — there is no single standard pattern; document per site `[secondary]`.

### 15.2 DNS hardening
- Anycast DNS for resilience; DNSSEC validation on resolvers; response-rate limiting; keep resolver infrastructure in the shared-services VRF with explicit leaking `[secondary]` — operator practice.
- Gap flagged in Wave 9 stands: deployment-prevalence data not found `[unverified]`.

### 15.3 NTP/NTS
- NTP stratum hierarchy; `ntp authenticate` / NTS (RFC 8915, Network Time Security) for authenticated time where compliance demands it `[official]` — RFC 8915.
- Redundant sources, diverse paths; monitor offset/jitter; alert on stratum changes — time jumps break Kerberos, logs, and trading timestamps `[secondary]` — operator practice.

### 15.4 PTP profiles
- IEEE 1588-2008/2019 defines the protocol; PROFILES define the subset: Default (E2E/P2P delay), Telecom G.8275.1 (phase/time sync with full timing support, boundary clocks), G.8275.2 (partial timing support), Power (C37.238), AVB/TSN (802.1AS) `[official]` — IEEE 1588; ITU-T G.8275.x.
- DC use: Default P2P/E2E profile on capable switches for telemetry correlation; G.8275.1 only where carrier-grade phase sync is required `[secondary]`.
- Boundary clock vs transparent clock: boundary terminates and regenerates time per hop (scales, adds servo error per hop); transparent corrects residence time in-flight (simpler, needs HW timestamping on every box) `[official]` — IEEE 1588 concepts.
- Gap from Wave 9 stands: per-platform profile support matrix not compiled `[unverified]`.

### Wave 15 verification
- Sources: 6 (RFC 4862, 8415, 8915, IEEE 1588, ITU-T G.8275.x). Practice-level claims labeled; gaps carried forward.

---

## Final verification (post Waves 10–15)

- Line count, tail integrity, and Markdown sanity verified by script after the last append.
- All new claims carry provenance tags; RFC references are to stable IETF/IEEE/ITU documents.
- No identifiers guessed; URLs verbatim from search results only.
- Open items from Waves 1–9 carried forward unchanged; new items: UEC deployment timing (`[unverified]`), mVPN-vs-TRM scoping note.

*End of Phase D4 — all 15 waves complete, append-only.*

## Wave 16 — Segmentation in practice: compliance, DMZ, service insertion

### 16.1 Compliance-driven segmentation (PCI-DSS, etc.)
- PCI-DSS requires cardholder-data environments to be isolated from the rest of the network; in DC fabrics this maps to dedicated tenant VRFs + firewalled boundaries + no route leaking except explicitly approved service prefixes `[secondary]` — PCI-DSS v4.0 scoping guidance; operator practice.
- Audit-scoping benefit: a well-segmented fabric shrinks the systems "in scope" for audit, which is frequently what funds the segmentation program `[secondary]` — ronutz/arsenal (Wave 3).
- Evidence artifacts auditors ask for: VRF/RT design doc, ACL/policy exports, firewall rule reviews, flow logs showing denied inter-segment traffic, change history `[secondary]` — operator practice; exact artifact lists vary by QSA.

### 16.2 DMZ and internet-edge segmentation
- Classic three-tier: outside / DMZ / inside, each in its own VRF or firewall zone; in EVPN fabrics the DMZ is a tenant VRF with border-leaf exit toward the edge firewalls `[secondary]` — enizaksoy/imamassypov lab topologies show dual-DMZ designs with per-location shared-services subnets (Wave 9).
- Internet egress consolidation: one shared Internet VRF on border leaves, tenant VRFs leak default toward it — avoids per-tenant edge hardware `[secondary]` — enizaksoy lab (Waves 2, 11).

### 16.3 Service insertion / service chaining
- Forcing inter-segment traffic through inspection: PBR or VRF-aware static routes steer flows to firewall/IDS clusters; return traffic must be symmetric or the firewall drops it — the #1 service-insertion bug `[secondary]` — operator practice.
- In ACI, service graphs (L4–L7 insertion) automate this; in EVPN-VXLAN fabrics it is manual PBR/VRF-leak design `[official]/[secondary]` — Cisco ACI docs; operator practice.
- Load-balancer/ADC insertion follows the same pattern (one-armed vs inline) — one-armed needs source-NAT or DSR awareness to keep return paths correct `[secondary]`.

### 16.4 Monitoring the segmented fabric
- Denied-flow telemetry (firewall denies, ACL hit counts, NSX DFW logs) is the highest-signal detection source precisely because the fabric is segmented — a denied east-west flow is an anomaly worth alerting `[secondary]` — ronutz/arsenal (Wave 3).
- ERSPAN/mirror sessions, sFlow/IPFIX exports, and telemetry (gNMI) feed the SOC; mirror destination placement must respect VRF boundaries or captures go dark `[secondary]` — operator practice.

### Wave 16 verification
- Sources: 5. Compliance section is practice-level; PCI-DSS cited as the framework, not quoted.

---

