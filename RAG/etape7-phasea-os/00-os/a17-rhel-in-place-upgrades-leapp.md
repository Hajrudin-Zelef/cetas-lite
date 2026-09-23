---
id: etape7-phasea-os/00-os/a17-rhel-in-place-upgrades-leapp
title: "A17 — RHEL in-place upgrades: Leapp"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["Oracle"]
dates: ["2021-06-21", "2022-07-14", "2024-05-31", "2025-05", "2025-06-11", "2026-09-22", "2027-05-31", "2029-05-31", "2030-05-31", "2032-05-31", "2035-05-31"]
keywords: ["pricing"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [256, 305]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: 01c572aedd104d13a3a679b1fca170a49e95f09ecdbc45e124230a4ee9fec974
---

# A17 — RHEL in-place upgrades: Leapp

## A17 — RHEL in-place upgrades: Leapp

**Leapp** is Red Hat's supported framework for **major-version in-place upgrades** (7→8, 8→9, 9→10) `[official]`. RHEL 9.6 → 10.0 is a supported Leapp path `[secondary]`; 8→10 directly is not — upgrades must step through majors.

**Relevance to rebuilds.** AlmaLinux's **ELevate** (A28) is built on Leapp with AlmaLinux actors; Rocky documents its own upgrade guidance separately. Oracle offers `leapp`-based tooling for Oracle Linux major upgrades `[secondary]`.

**Operational note.** Leapp upgrades are pre-flight-checked (inhibitors block the upgrade on known-bad configurations) but remain riskier than reprovisioning for stateless workloads. The 7C automation layer (image builds with Packer, cloud-init) is the preferred path for immutable-style estates; Leapp is the tool for pets.

---

## A18 — Rocky Linux 10 "Red Quartz": current community enterprise major

**Release.** Rocky Linux 10.0 "Red Quartz" was released **2025-06-11** `[official]`, tracking RHEL 10.0's May 2025 release with the project's characteristic short lag. Kernel **6.12** `[official]`.

**Architectures and ISA baseline.** Rocky 10 ships for **x86_64, aarch64, ppc64le, s390x and riscv64** `[official]` — notably including riscv64, which RHEL 10 does not ship as a supported architecture. **x86-64-v3 is the minimum ISA level** for the x86_64 build `[official]` — hardware older than ~2015 (pre-Haswell-class) cannot run Rocky 10's x86_64 binaries. This is the single most important Rocky-vs-AlmaLinux technical differentiator (A24).

**Release policy.** The project publishes **two minor releases per year**; when a new minor ships, the previous minor moves to the vault (no further updates) `[official]`. The `.10` minor of each major is the final one and enters a maintenance phase `[official]`. This mirrors RHEL's cadence without RHEL's EUS pinning (pinning a Rocky minor long-term is a CIQ RLC feature, not a community feature).

**Lifecycle.** Active support to **2030-05-31**; EOL **2035-05-31** `[official]`, mirroring RHEL 10. Community support is best-effort via forums/Mattermost; commercial SLAs come from CIQ (A20).

**Current minor (Sept 2026).** The project had shipped 10.x minors on its twice-yearly cadence; the exact current minor at observation time could not be verified from the project's release pages in this pass `[unverified]`.

---

## A19 — Rocky Linux 9 "Blue Onyx" and 8 "Green Obsidian"

**Rocky 9** (GA 2022-07-14, kernel 5.14): active support to **2027-05-31**, EOL **2032-05-31** `[official]`. The dominant Rocky version in production fleets as of 2026 — most CentOS 7/8 → Rocky migrations landed on Rocky 8 or 9.

**Rocky 8** (GA 2021-06-21, kernel 4.18): active support ended 2024-05-31, EOL **2029-05-31** `[official]`. Rocky 8 → 9 major upgrades are supported via the project's documented migrate path (Leapp-based tooling adapted for Rocky) `[secondary]`.

---

## A20 — CIQ: the commercial company behind Rocky

**Who CIQ is.** CIQ (ctrl IQ, Inc.), founded by Rocky Linux founder Gregory Kurtzer, is the primary commercial sponsor of Rocky Linux and sells enterprise support and hardened builds `[vendor-reported]`.

**Product ladder (list pricing, USD, per node/year, observed 2026-09-22)** `[vendor-reported]`:

| Product | Self-support | Standard | Premium |
|---|---|---|---|
| **RLC Pro** (supported Rocky) | $350 | $600 | $825 |
| **RLC Pro Hardened** (FIPS/STIG-hardened) | — | $775 | $1,000 |
| **RLC Pro AI** (AI-optimized stack) | — | $775 | $1,000 |

CIQ states RLC Pro includes long-term-support windows, **FIPS 140-3 validated cryptography**, direct engineering (bug fixes), indemnification and CVE response-time commitments `[vendor-reported]`. **RLC+** is the lower-priced tier below RLC Pro (feature comparison published by CIQ) `[vendor-reported]`.

**Assessment.** CIQ's pricing is public and node-based — simpler to budget than Red Hat's SKU matrix, and roughly 30–50% below RHEL Server list for comparable SLA tiers `[secondary]`. The trade-off: CIQ is a smaller vendor (counterparty risk), and RLC Pro is Rocky-specific — migrating away later means another OS migration.

---

