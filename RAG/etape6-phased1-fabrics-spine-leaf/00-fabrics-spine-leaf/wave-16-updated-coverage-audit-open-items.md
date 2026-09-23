---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-16-updated-coverage-audit-open-items
title: "Wave 16 — Updated coverage audit & open items"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Google", "Nvidia", "xAI"]
dates: ["2026-03", "2026-09-02"]
keywords: ["acquisition", "alignment", "benchmarks", "claude", "cost", "disclosure", "gpu", "gpus", "hyperscaler", "license", "nvidia", "optics"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [426, 482]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: f91497a609c21e19a07d639bf94f05ab4cc8ad45ccfb9e3286d90d935878ac3f
---

# Wave 16 — Updated coverage audit & open items

## Wave 16 — Updated coverage audit & open items

**Scope-to-wave map:** Clos/spine-leaf principles + oversubscription + scaling math (Waves 1, 7) · failure domains & cabling (Wave 1) · 3-tier vs spine-leaf & migration (Wave 2) · Cisco (Waves 3, 8) · Arista (Waves 3, 9) · NVIDIA (Waves 3, 10) · Dell (Waves 3, 11) · Aruba/HPE (Waves 3, 11) · Juniper (Waves 3, 12) · whitebox/SONiC (Waves 3, 15) · underlay essentials: numbered/unnumbered, BFD, cabling plans (Waves 4, 10–11) · Tomahawk 6 silicon (Waves 7, 13) · hyperscaler deployments (Waves 5, 14) · economics (Wave 15) · decision checklist (Wave 6).

**Provenance audit:** all new factual claims carry tags; vendor scale/cost claims tagged [vendor-reported] or [secondary]; marketplace prices flagged as indicative.

**Open items / [gap]s (updated):** (1) Google Jupiter-class current fabric specifics — retained; (2) Azure SONiC current fleet counts — retained; (3) HPE/Juniper post-acquisition NOS roadmap — retained; (4) independent 800G fabric benchmarks — retained; (5) Dell Enterprise SONiC license pricing — retained; (6) Cisco ACI per-release max leaf counts (only APIC 5.2.x verified guides consulted) — new; (7) xAI Colossus exact topology — retained as [unverified]; (8) Tomahawk 6 independent power/thermal figures — retained.

*End of Phase D1 (16 waves). Deep BGP underlay and EVPN-VXLAN protocol coverage continues in sibling Phase D files.*

---

## Wave 17 — Cisco 2026 update: Silicon One G300, Nexus One, Hyperfabric, and a fabric security note

### 17.1 Silicon One G300 (announced Feb 2026)

At Cisco Live EMEA 2026 Cisco announced **Silicon One G300**: 102.4 Tb/s switching silicon, framed as 64×1600GbE, powering new Cisco Nexus 9000 and Cisco 8000 systems for AI scale-out [official/vendor-reported — https://www.sdxcentral.com/news/next-gen-switches-servers-everything-unveiled-at-cisco-live-emea-2026/ and https://www.eweek.com/news/cisco-silicon-one-g300-chip/]. Alongside: deep-buffer **800G P200-based switches** for scale-across, a new **1.6T pluggable optics** line for AI scale-out (plus 800G units claimed at 50% lower power than retimed modules [vendor-reported]), and 102.4T liquid-cooled systems. Customer voices in the announcement: du (UAE) citing 100Tbps capacity needs toward the 1.6T era; Sharon AI selecting **Nexus Hyperfabric** aligned to NVIDIA NCP design standards; Cirrascale on G300 + deep-buffer 800G [vendor-reported].

### 17.2 Nexus One and Unified Fabric (management plane)

**Nexus One** is Cisco's unified fabric management platform: single control plane across silicon, systems, optics and software with GPU-deep telemetry; from March 2026 it analyzes telemetry in place via **Splunk integration** (positioned for sovereign cloud) [official — https://www.sdxcentral.com/news/next-gen-switches-servers-everything-unveiled-at-cisco-live-emea-2026/]. **Nexus Hyperfabric** (cloud-managed, validated stacks incl. GPU/storage) targets enterprise AI deployment simplicity [vendor-reported].

### 17.3 Security note: CVE-2026-20212 (Silicon One Nexus 9000)

On 2026-09-02 Cisco PSIRT disclosed **CVE-2026-20212** (CVSS 9.8): unauthenticated root RCE via TCP 43210/43211 in the default L3 VRF on Silicon One-based Nexus 9000s — the switches carrying GPU-to-GPU RDMA in large AI fabrics; no known exploitation at disclosure [secondary — https://forkast.news/cisco-nexus-9000-silicon-one-rce-exposes-ai-data-center-fabric-to-root-compromise/]. **Fabric-design implication:** AI backend fabrics are now explicitly called out as high-value attack surfaces; out-of-band management isolation and control-plane ACL hygiene belong in every AI fabric design review. Tracked as an open item for the security phase.

---

## Wave 18 — AI backend fabric topologies: rail-only vs rail-optimized, ROD vs RUD

### 18.1 The rail concept

AI training traffic is dominated by same-rank GPU communication (allreduce rings): GPU *i* on every node talks to GPU *i* on other nodes. **Rails** group same-index GPUs: GPU0 of every node → Rail 0, GPU1 → Rail 1, etc. Each rail is a congestion/failure domain [secondary — https://github.com/atom00blue/ai-stack-silicon-to-tokens/blob/HEAD/D06-rack-cluster-network/D06.02-network-topology/D06.02%20Network%20Topology.md and https://github.com/ziwon/ai-data-center-network/blob/HEAD/ai-data-center-network/chap03/README.md].

| Design | Topology | Scale fit | Trade-off |
|---|---|---|---|
| **Rail-only** | One leaf (or leaf set) per rail, no spine | 32–256 GPUs | Full rail isolation, lowest cost; no any-to-any |
| **Rail-optimized** | Rail leaves + spine tier (Clos) | 10,000+ GPUs | Single-hop intra-rail (~600 ns), 3-hop cross-rail (<2 µs); spine cost/complexity |
| Standard fat-tree/Clos | Any-to-any | General | 3-hop for all GPU traffic regardless of pattern |

Canonical cabling rule: **NIC-N → Leaf-N** [secondary]. NCCL topology-aware scheduling maps ring/tree collectives onto rails, reducing PFC crosstalk between rails [secondary]. A published 1,024-GPU Clos reference (Cisco blueprint): 128 servers × 8 GPUs, 10 leaves + 4 spines, ~4.5 µs end-to-end [secondary — same source].

### 18.2 ROD vs RUD (GPU-to-leaf attachment styles)

| Design | Model | Benefit | Cost |
|---|---|---|---|
| **ROD** (rail-optimized design) | 1 GPU/NIC → 1 dedicated rail/leaf | Predictable paths, 1-hop rail-local, fault isolation by rail, NCCL alignment | More cables, more leaf ports |
| **RUD** (rail-up design / shared) | Multiple GPUs → 1 leaf | Simpler cabling, lower cost | Shared fate, path complexity |

[secondary — https://github.com/ziwon/ai-data-center-network/blob/HEAD/ai-data-center-network/chap03/README.md]

### 18.3 Operational classification (field lesson)

A rail fabric can be: full-mesh (any NIC reaches any NIC), rail-only (cross-rail unroutable), or **IP-routable but RDMA rail-only** — cross-rail ping works but RoCEv2 RDMA fails. Field guidance: **never classify from ping alone**; confirm at the RDMA layer (QP INIT→RTR failures, ETIMEDOUT) [secondary — https://github.com/rocm/mori/blob/HEAD/.claude/skills/cluster-network-topology/SKILL.md].

---

