---
id: etape6-phased3-bgp-ha/00-bgp-ha/3-5-aruba-cx-vsx-virtual-switching-extension-aos-cx
title: "3.5 Aruba CX VSX (Virtual Switching Extension) — AOS-CX"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: []
dates: ["2026-05"]
keywords: ["agent", "benchmark", "benchmarks"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [191, 236]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: 536dd6c08f7d454986814bf119131f4329c7e394dafd777a90e13f6978d55766
---

# 3.5 Aruba CX VSX (Virtual Switching Extension) — AOS-CX

### 3.5 Aruba CX VSX (Virtual Switching Extension) — AOS-CX

- **Architecture:** VSX pairs two CX switches: **one logical device at L2, independent devices at L3**. Inter-Switch Link (**ISL**, a LAG of 2–8 directly-connected same-speed ports, typically ≥100G) carries the ISLP protocol synchronizing LACP, MAC, ARP, STP, and DHCP state; the **VSX keepalive** is a UDP probe (default port **7678**, configurable) over a separate L3 path — best practice is the OOBM port or a routed port in a custom VRF, or a loopback over redundant routed paths [official — AOS-CX 10.16 VSX Guide; Aruba DC validated design guide, May 2026].
- **Split-brain behavior** [official — AOS-CX 10.14/10.16 VSX guides]:
  - **ISL down, both switches up:** keepalive confirms both are alive; the configured **primary keeps its VSX LAGs up**, the **secondary forces its VSX LAG links down** with a reason code. On ISL recovery: MAC/ARP sync from primary → secondary, then a **link-up delay timer** expires before the secondary re-enables its VSX LAGs.
  - **ISL down AND peer down:** the surviving switch (either role) keeps its VSX LAGs up; on peer return, ISL comes up first, tables sync, delay timer, then LAGs re-enable.
  - Warns that keepalive + ISL on the same line card is a single point of failure risking split-brain.
- **VSX LAG specifics:** L2-only (no routed mode), LACP active or static, standard timers/hashing/LACP fallback supported; no extra encapsulation on the ISL so no MTU penalty, though jumbo MTU tuning is recommended if endpoints need it [official — VSX technical whitepaper].
- **Active-gateway:** Aruba's anycast-gateway equivalent for VSX (covered in §4.2).
- **Config sync:** `vsx-sync` replicates management-plane config from primary to secondary, reducing human error [official].
- Sources: https://arubanetworking.hpe.com/techdocs/AOS-CX/10.16/PDF/vsx.pdf, https://www.arubanetworks.com/techdocs/VSG/docs/040-dc-design/Media/PDF/Aruba_VSG_Data-Center-Design.pdf

### 3.6 Juniper MC-LAG — Junos (QFX/EX)

- **Architecture:** MC-LAG peers run **ICCP** (Inter-Chassis Control Protocol, RFC 7275) over a dedicated link plus an **ICL** (inter-chassis link, an ae interface) that protects MC-AE interfaces; each MC-AE interface gets an `mc-ae-id` matching on both peers and per-chassis IDs [official — Juniper MC-LAG user guides].
- **Liveness:** ICCP `liveness-detection` with BFD (recommended minimum-receive-interval **2000 ms**, multiplier 4 in the documented example); `session-establishment-hold-time` recommended **50 s**; optional `backup-liveness-detection` via a backup peer IP for sub-second detection of peer reboot [official].
- **Failure matrix** (documented ICCP failure scenarios) [official — "Configuring Multichassis Link Aggregation on a QFX Series Switch"]:
  - ICCP down + backup liveness inactive/not configured → standby peer's LACP system ID reverts to default (peers diverge; downstream sees two systems).
  - ICCP down + backup liveness active → LACP system ID also reverts (per the published table).
  - ICCP up + ICL down → standby MC-AE LACP state moves to standby/waiting.
  - Split-brain state brings the MC-LAG link fully down if primary members are also down; auto-recovers when ICCP adjacency returns.
- **Status control:** `status-control active/standby` per MC-AE plus `events iccp-peer-down prefer-status-control-standby` decides which peer forwards when ICCP drops [official].
- **Core isolation (QFX5100):** if a leaf loses all spine connectivity, EVPN core isolation brings LACP bundles down (Mux → Waiting) to stop traffic ingress; **enabled by default** on QFX5100; conflicts with `minimum-links` config (documented JTAC mismatch case) [official — Juniper support portal].
- Sources: https://www.juniper.net/documentation/en_US/junos/information-products/pathway-pages/mc-lag/multichassis-link-aggregation-groups.pdf, https://supportportal.juniper.net/s/article/QFX-Mismatched-LACP-bundle-state-when-core-isolation-and-minimum-links-config-are-enabled

### 3.7 Cross-vendor MC-LAG comparison

| Aspect | Cisco vPC | Arista MLAG | Cumulus CLAG | Dell VLT | Aruba VSX | Juniper MC-LAG |
|---|---|---|---|---|---|---|
| Control protocol | vPC CFSoE | MLAG agent (TCP) | clagd | VLT protocol | ISLP + keepalive | ICCP (RFC 7275) |
| Keepalive default | 1 s | 2–4 s (train-dependent) | via backup IP | 30 s | UDP 7678 | ICCP + BFD 2 s×4 |
| Keepalive path | separate routed link/VRF | separate VRF mgmt (dual-primary detect) | backup IP, vrf mgmt | OOB mgmt recommended | OOBM / routed port / loopback | backup-peer-ip |
| Split behavior | secondary shuts member ports | association dissolves → independent | backup IP keeps domain | secondary shuts VLT ports | secondary shuts VSX LAGs | standby LACP ID reverts |
| L3 model | peer-gateway for FHRP | VARP (anycast) | anycast (VRRP) | VLT peer routing | active-gateway | VRRP / anycast |
| ISSU support | yes (NX-OS) | MLAG ISSU | [unverified] | [unverified] | yes (VSX design goal) | [unverified] |

- **Gap:** sub-second convergence claims for individual MC-LAG failovers are vendor-reported and scenario-specific; no independent head-to-head benchmark found.

### 3.8 Wave-3 verification notes and open items

- **Verified:** vPC timers and failure matrix; Arista heartbeat semantics and dual-primary detection; Cumulus peerlink/backup-IP mechanics; Dell VLT liveliness and role behavior; Aruba VSX keepalive/split-brain flows; Juniper ICCP failure table and core isolation.
- **Open:** ISSU support matrix for Cumulus/Dell/Juniper MC-LAG; independent failover benchmarks.
- *End of Wave 3.*

---

