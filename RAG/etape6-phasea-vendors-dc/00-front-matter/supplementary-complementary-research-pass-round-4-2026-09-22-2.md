---
id: etape6-phasea-vendors-dc/00-front-matter/supplementary-complementary-research-pass-round-4-2026-09-22-2
title: "Supplementary / Complementary Research Pass — round 4 (2026-09-22)"
domain: front-matter
role: reference
task: reference
actors: ["Nvidia"]
dates: ["2019-04-28", "2020-11-29", "2021-06-30", "2021-12-31", "2025-10", "2026-01", "2026-03", "2026-03-10", "2026-08", "2026-09", "2026-09-22"]
keywords: ["research", "advisory", "cyber", "ethernet", "gpu", "memory", "nvidia"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1178, 1232]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 198f177761dbcce5675fb52c36ef9f46424a2229a049b3a0e8d2170bc5594ad9
---

# Supplementary / Complementary Research Pass — round 4 (2026-09-22)

## Supplementary / Complementary Research Pass — round 4 (2026-09-22)

**Scope note:** Three supplementary passes (§§A–M, §§N–X, "round 2" §§N–S) and a round-3 pass (§§AA–II) already exist. This round-4 pass adds only material not present in the base §§1–8 or in earlier passes. Nothing above was modified. Read the base report and earlier passes first; the provenance legend from the base applies here.

### JJ. NVIDIA UFM 2026 security bulletins — new risk context for IB fabrics

Base §2.3 names UFM Enterprise 6.26.1 but carries no 2026 security detail. An August 2026 NVIDIA security bulletin covered **UFM Enterprise and DGX Spark** [secondary — securityonline.info, ~Aug 26, 2026]:

- **10 CVEs** (5 High, 5 Medium); no confirmed active exploitation; no public PoCs at the time of the bulletin [secondary].
- Highest severity: **CVE-2026-24170, CVSS 8.8** — improper authentication in the UFM Enterprise web-interface authorization component; an authenticated user could send crafted HTTP requests to cause improper authentication, leading to code execution / privilege escalation [secondary].
- CVE-2026-24169 (8.0): plugin-management API code injection by authenticated attacker [secondary].
- CVE-2026-24262 / CVE-2026-24263 / CVE-2026-47626 (8.2): heap/out-of-bounds/NULL-pointer issues [secondary].
- CVE-2026-24167 / CVE-2026-24168 (6.8): command injection via improper input handling in user management and diagnostic tools; CVE-2026-24166: hard-coded cryptographic key in session management [secondary].
- Remediation: apply NVIDIA's latest UFM security updates [secondary]. Treat the CVE inventory as a snapshot, not an audit — the vendor advisory is the authoritative source.

Background: the UFM family has three members — **UFM Enterprise** (monitoring, management, performance optimization, configuration checks, secure cable management), **UFM Cyber-AI** (AI-based anomaly detection / failure prediction), and **UFM Telemetry** (real-time telemetry streamed to on-prem/cloud databases). Operators quoted in coverage include NCI Australia and the Ohio Supercomputer Center [secondary — TechPowerUp, 2025].

### KK. NVIDIA DOCA SDK 3.x — 2026 data-plane release train

Round-3 §DD covered DOCA Platform Framework (DPF) orchestration (v26.x), not the **DOCA SDK data-plane line**. The DOCA SDK 3.x train [official — NVIDIA DOCA docs]:

- **DOCA-OFED 3.2.0 GA — October 2025**; **3.3.0 — January 2026** [official — docs.nvidia.com DOCA general-support page, updated Mar 2, 2026].
- **DOCA 3.2.0** archive includes GPUNetIO changes: CPU shared library no longer depends on CUDA RT; DPDK dependency removed from GPUNetIO samples; Ethernet functions reworked as open-source CUDA inline header functions (Share-QP send on block/warp scopes, CPU proxy execution mode); Comm Channel API for DPU↔GPU communication; dmabuf GPU-memory support in all GPUNetIO Verbs/Ethernet samples [official — DOCA 3.2.0 archive docs].
- **DOCA HBN (BlueField networking app)** releases on NGC: **3.2.1.1** (Feb 4, 2026), **3.2.1.2** (Aug 21, 2026), **3.2.2** (Apr 24, 2026), **3.2.3** (Jul 21, 2026), **3.3.0** (Feb 26, 2026), **3.4.0** (Jun 9, 2026), **3.5.0** (Aug 31, 2026) [official — NGC catalog].
- SNAP virtio-fs service release notes reference versions built against **DOCA 2.8.0, 3.1.0 and 3.2.0**, last updated March 2026 [official].
- Do not conflate with DPF (§DD): DPF v26.x is the Kubernetes provisioning/orchestration layer; DOCA SDK 3.x is the data-plane SDK and OFED stack [gap — DPF vs DOCA SDK version-number correlation not documented].

### LL. Mellanox heritage: legacy switch EOL detail (extends round-3 §CC)

Round-3 §CC covered ConnectX adapter EOL. The switch side of the Mellanox lineage [official — NVIDIA network.nvidia.com EOL notices]:

- **SB7700 (Switch-IB EDR 36-port, 100G QSFP28)**: manufacture discontinue / EOL announced **April 28, 2019** (LCR-000453); short-depth variant had no replacement option; replacement for standard-depth = **MSB7800-ES2F2 (SB7800, Switch-IB-2)**; last supported firmware rel-11_1910_0618, last supported software 3.7.1134 [official].
- **SB7800 (Switch-IB-2 EDR)**: EOL announced **November 29, 2020** (LCR-000694); last-time-buy **June 30, 2021**; last ship **December 31, 2021**; reason: low market interest [official].
- **SX6036 (FDR 36-port)**: community homelab documentation (Sept 2026) notes the platform is **end-of-life and no longer receiving firmware updates** — it must be staged to its last image via incremental upgrades (30–60 min per image) [secondary — GitHub soliddowant/infra-mk3 docs].
- 40G passive-copper cable lines (MC22101xx, MC23091xx, MC26091xx, MC33091xx, MCP1700, MCP2104) were EOL'd in 2021-era notices [official — LCR-000916].
- Installed-base implication: EDR-era (100G) Mellanox fabrics are ~5 years past EOL; migration path is Quantum-2 (NDR) / Quantum-X800 (XDR) per base §2.3.

### MM. Aruba AOS-CX 2026 security bulletins — campus-to-DC exposure

Not covered elsewhere in this report. Two 2026 HPE security bulletins affected the full AOS-CX fleet:

**March 2026 — HPESBNW05027** (initial release March 10, 2026) [secondary — imfht.com mirror; SecurityWeek, Mar 2026]:
- **CVE-2026-23813 (CVSS 9.8)**: unauthenticated remote **administrator-password reset** via the web management interface — full switch takeover.
- CVE-2026-23814 (8.8), CVE-2026-23815/23816 (7.2): authenticated remote command injection; CVE-2026-23817 (6.5): unauthenticated open redirect in the login page (credential-harvesting vector).
- Fixed versions: **10.17.1001, 10.16.1030, 10.13.1161, 10.10.1180** [secondary].
- Affected hardware: CX 4100i, 6000, 6100, 6200(F), 6300, 6400, 8320, 8325, 8360, 8400, 9300, 10000 — i.e. the entire campus + DC portfolio [secondary].

**September 2026 — HPESBNW05134** [secondary — techtimes.com, socradar.io, cvetodo.com, redlegg.com; all ~Sep 4, 2026]:
- **CVE-2026-73749 (CVSS 9.8)**: multiple buffer-overflow defects (bundled under one CVE) in an AOS-CX daemon; unauthenticated remote attacker sends crafted packets → **RCE with elevated privileges**.
- CVE-2026-73752: unauthenticated adjacent-network arbitrary file write via an AOS-CX API endpoint (persistent-backdoor path).
- CVE-2026-73751 / 73753 / 73750: low-privileged-user → arbitrary OS commands / privilege escalation / DoS; CVE-2026-73781: stored XSS; CVE-2026-73780: CSRF on certificate-authenticated sessions.
- Total scope: **>150 defects collapsed into 34 CVEs** in one bulletin [secondary — cvetodo.com].
- Fixed versions: **10.18.1002, 10.17.1030, 10.16.1060, 10.13.1190, 10.10.1181**; **the 10.10.x branch is End of Maintenance** and only receives fixes for internally discovered critical issues [secondary].
- HPE workaround if immediate patching is impossible: isolate CLI/web management interfaces to a dedicated L2 segment/VLAN, enforce L3+ firewall policies, disable HTTP(S) on SVIs/routed ports, enable accounting [secondary].

