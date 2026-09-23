---
id: etape7-phasea-os/00-os/overview
title: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["Oracle"]
dates: ["2021-12-31", "2024-06-30", "2025-05", "2025-05-13", "2025-05-20", "2025-05-27", "2025-06", "2025-06-11", "2025-08-09", "2026-04", "2026-04-23", "2026-06-11", "2026-07-11", "2026-09", "2026-09-22", "2027-05-31", "2028-06-30"]
keywords: ["benchmark", "distribution", "pricing", "research"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [1, 45]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: db3dca249c4f4e9b7e019bb92b79cbd292ede3d147020d1074930ea0fafb9f54
---

# Step 7 — Phase A: Server Operating Systems (Linux OS Layer)

## A0 — Scope, observation date, method and provenance legend

**Scope.** This file documents the Linux operating-system layer of the infrastructure stack: the major server distributions evaluated for Proxmox hosts, virtual machines, containers and bare-metal nodes — Debian, Ubuntu, RHEL, Rocky Linux, AlmaLinux — plus a secondary comparison of Oracle Linux, and the minimal/immutable variants (Fedora CoreOS, openSUSE MicroOS, bootc-based images) that compete with them in container-first roles.

**Observation date:** 2026-09-22. All version, kernel and lifecycle claims are current as of that date unless a different observation date is given inline.

**Method.** Web research conducted 2026-09-22 against primary vendor documentation (Debian Project, Canonical, Red Hat, Rocky Linux, AlmaLinux, Oracle), machine-readable lifecycle data (endoflife.date), security authority documents (DISA, NIST, CIS) and reputable secondary press. Every factual claim below carries one of five provenance tags:

| Tag | Meaning |
|---|---|
| `[official]` | Primary source: the vendor's own documentation, release notes, blog or price list |
| `[vendor-reported]` | Claim made by a commercial vendor about its own product (not independently verified) |
| `[independent]` | Independent survey, benchmark or technical analysis from a third party |
| `[secondary]` | Trade press / community documentation repeating primary claims, not cross-verified |
| `[unverified]` | Could not be confirmed against a reliable source at research time; treat with caution |

**Known gaps (read before use).** (1) Debian 12's regular-support handoff date is reported as both 2026-06-11 and 2026-07-11 by different sources — unresolved, both recorded. (2) RHEL 10's exact GA day appears as 2025-05-13 in one Red Hat regional page and 2025-05-20 in press summaries — the month is firm, the day is flagged. (3) Ubuntu 26.04 (April 2026) details are from secondary press; first-party Canonical release notes could not be fetched in this pass. (4) Public "server OS market share" figures are methodologically incompatible across sources; they are presented side by side but must not be compared as a single ranking. (5) Canonical Landscape standalone pricing could not be confirmed against current Canonical documentation; only third-party aggregators quote a figure.

**No invented identifiers.** No URLs, version numbers, kernel revisions, prices or dates in this file were invented; identifiers not found in a source are marked `[unverified]` or omitted.

---

## A1 — Executive snapshot: the server Linux landscape in September 2026

The server Linux market in late 2026 is shaped by three post-CentOS facts. First, **CentOS Linux 8 died 2021-12-31** and **CentOS Linux 7 died 2024-06-30** `[official]`, pushing the former "free RHEL" base into three camps: RHEL paid subscriptions, the two community RHEL rebuilds (Rocky Linux and AlmaLinux), and migrations to Ubuntu/Debian `[independent]`. Second, the rebuilds have matured into commercial-grade platforms with paid support ecosystems (CIQ for Rocky, TuxCare ESU/ELS covering both, Oracle as an alternative) `[vendor-reported]`. Third, Canonical's Ubuntu is the only distribution with a **10-year free-ish extended maintenance window plus a 2-year paid "Legacy" add-on** (Ubuntu Pro) `[vendor-reported]`, while the RHEL family offers the longest paid lifecycles (13 years for RHEL 9/10 with ELS) `[official]`.

**Current major versions (September 2026):**

| Distribution | Current major | Codename | GA date | GA kernel | Status |
|---|---|---|---|---|---|
| Debian | 13 | Trixie | 2025-08-09 `[secondary]` | 6.12 LTS `[secondary]` | Stable (current) |
| Ubuntu | 26.04 LTS | Resolute Raccoon | 2026-04-23 `[secondary]` | 7.0 `[secondary]` | Current LTS |
| RHEL | 10 | (none — Red Hat dropped release codenames; internal name "Coughlan" in press) `[secondary]` | May 2025 `[secondary]` | 6.12 `[secondary]` | Current |
| Rocky Linux | 10 | Red Quartz | 2025-06-11 `[official]` | 6.12 `[official]` | Current |
| AlmaLinux | 10 | Purple Lion | 2025-05-27 `[official]` | 6.12.0-55.9.1 `[secondary]` | Current |
| Oracle Linux | 10 | — | June 2025 `[secondary]` | UEK 8.1 (6.12-based) `[secondary]` | Current (secondary comparison) |

**Previous still-supported majors:** Debian 12 (LTS until 2028-06-30), Ubuntu 24.04 LTS (standard support until 2029), RHEL 9 (full support until 2027-05-31), Rocky 9 / AlmaLinux 9 (EOL 2032), Oracle Linux 9 `[official]`.

**Strategic read for infrastructure use.** For a Proxmox-based lab/self-host estate, Debian 12/13 is the native host OS (Proxmox VE is built on Debian) `[secondary]`; Ubuntu 24.04 LTS dominates general server/VM adoption `[independent]`; RHEL 9/10 and its rebuilds are the compliance-heavy choice (FIPS 140-3, DISA STIG, CIS) `[official]`; Rocky and AlmaLinux are functionally interchangeable for most server roles, with the choice decided by support vendor preference (CIQ vs community/TuxCare) and ISA baseline (AlmaLinux keeps an x86_64-v2 variant; Rocky requires x86_64-v3) `[official]`.

---

