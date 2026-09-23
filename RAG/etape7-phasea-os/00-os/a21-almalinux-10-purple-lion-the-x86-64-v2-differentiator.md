---
id: etape7-phasea-os/00-os/a21-almalinux-10-purple-lion-the-x86-64-v2-differentiator
title: "A21 — AlmaLinux 10 \"Purple Lion\": the x86_64-v2 differentiator"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: []
dates: ["2021-03-30", "2022-05-26", "2023-06", "2024-05-01", "2025-05-27", "2025-11-24", "2027-05-31", "2029-03-01", "2029-05-01", "2030-05-31", "2032-05-31", "2035-05-31"]
keywords: ["governance", "pricing"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [306, 351]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: c26d3800fbc8933e204dff7b01d8c7d298c92cbbfb2a4d183046789233bd1884
---

# A21 — AlmaLinux 10 "Purple Lion": the x86_64-v2 differentiator

## A21 — AlmaLinux 10 "Purple Lion": the x86_64-v2 differentiator

**Release.** AlmaLinux 10.0 "Purple Lion" was released **2025-05-27** `[official]` — notably *before* Rocky 10, the fastest major rebuild release of the cycle. Kernel **6.12.0-55.9.1** at GA `[secondary]`.

**The ISA differentiator.** Unlike Rocky 10 (x86_64-v3 only), **AlmaLinux 10 ships both a standard x86_64 build and an x86_64-v2 compatibility build** `[official]`. For estates with pre-2015 hardware (or virtualized environments masking CPU flags), AlmaLinux 10 is the only RHEL-10-class option that runs — this single fact decides Rocky-vs-Alma for legacy-hardware estates.

**Compatibility policy (post-2023).** After Red Hat restricted RHEL source publication in June 2023, AlmaLinux redefined its goal as **ABI/application compatibility with RHEL** rather than byte-identical bug-for-bug rebuilds `[official]`. The outdated "1:1 binary clone" description should not be used without this qualification.

**Lifecycle.** Active support to **2030-05-31**, security support to **2035-05-31** `[official]`.

**Current minor.** **AlmaLinux 10.1 "Heliotrope Lion"** was released **2025-11-24** `[official]`; further 2026 minors follow the RHEL cadence (exact current minor at observation time `[unverified]` in this pass).

---

## A22 — AlmaLinux 9 and 8

**AlmaLinux 9** (GA 2022-05-26, kernel 5.14): active support to **2027-05-31**, security support to **2032-05-31** `[official]`. The most widely deployed AlmaLinux major, and the most common CentOS migration landing zone alongside Rocky 9.

**AlmaLinux 8** (GA 2021-03-30, kernel 4.18): active support ended 2024-05-01; security support ends **2029-03-01 per endoflife.date vs 2029-05-01 per AlmaLinux's own wiki** — conflict flagged, both recorded `[secondary]`. AlmaLinux 8 → 9 upgrades are supported via ELevate (A28).

---

## A23 — AlmaLinux Kitten 10: preview, not production

**AlmaLinux Kitten 10** is a **development preview** built from CentOS Stream 10 sources, published so the community and SIGs can prepare for the next AlmaLinux major `[official]`. It is explicitly not a production release and carries no lifecycle guarantees. Its existence demonstrates the project's post-2023 build pipeline (Stream-based, with independent QA) — relevant context for evaluating rebuild trustworthiness.

---

## A24 — Rocky vs AlmaLinux: decision factors

| Factor | Rocky Linux | AlmaLinux | Source |
|---|---|---|---|
| Governance | Rocky Enterprise Software Foundation (nonprofit) | AlmaLinux OS Foundation (nonprofit) | `[official]` |
| Commercial backer | CIQ | (community; TuxCare provides ELS/ESU) | `[official]`/`[vendor-reported]` |
| x86_64 ISA baseline (v10) | **x86_64-v3** (no pre-2015 CPUs) | **x86_64-v2** option available | `[official]` |
| riscv64 (v10) | Yes | `[unverified]` in this pass | `[official]` |
| Release lag after RHEL | Days to ~2 weeks | Days (10.0 beat Rocky by ~2 weeks) | `[official]` |
| Minor cadence | 2/year, previous minor vaulted | Tracks RHEL cadence | `[official]` |
| FIPS 140-3 | Via CIQ RLC Pro Hardened (certs #5117/#5116/#5113/#5095) | Via TuxCare (5 validated modules, AlmaLinux) `[vendor-reported]` | `[vendor-reported]` |
| Migration tooling | `migrate2rocky` | `almalinux-deploy`, ELevate | `[official]` |
| Lifecycle (v9/v10) | Identical (mirrors RHEL) | Identical (mirrors RHEL) | `[official]` |

**Bottom line.** For new deployments on modern hardware the two are interchangeable; choose on **support-vendor preference** (CIQ's public node pricing vs community/TuxCare) and **hardware age** (AlmaLinux for x86_64-v2). Avoid religious debates — the technical delta is small.

---

