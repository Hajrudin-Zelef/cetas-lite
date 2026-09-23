---
id: etape7-phasea-os/00-os/a2-master-lifecycle-table
title: "A2 — Master lifecycle table"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["Intel", "Oracle"]
dates: ["2014-06-10", "2019-05-07", "2019-08-06", "2020-04-23", "2021-03-30", "2021-06-21", "2021-08-14", "2021-12-31", "2022-04-21", "2022-05-18", "2022-05-26", "2022-06-30", "2022-07-14", "2023-06-10", "2024-04-25", "2024-05-01", "2024-05-31", "2024-06-30", "2025-05-20", "2025-05-27", "2025-05-31", "2025-06-11", "2025-08-09", "2025-11-24", "2026-04-23", "2026-06", "2026-06-11", "2026-07-11", "2026-08-31", "2026-09-12", "2026-12", "2027-05-31", "2028-06-30", "2028-08-09", "2029-03-01", "2029-05-01", "2029-05-31", "2029-06-30", "2030-05-31", "2030-06-30", "2031-06-30", "2032-05-31", "2033-05-31", "2033-06-30", "2035-05-31", "2035-06-30", "2036-05-31", "2039-05-31"]
keywords: ["cost", "distribution", "intel", "pricing"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [46, 97]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: 5a06b4619c3c6d87c5c94002248340e3bdbc7dab12995cd0154dd00c8f2ed96d
---

# A2 — Master lifecycle table

## A2 — Master lifecycle table

Dates below come from vendor lifecycle pages and endoflife.date's machine-readable data unless noted. "EOL" means end of all vendor/community security support, including paid extensions where offered.

| Release | GA | Regular/full support ends | Maintenance/LTS ends | Extended (paid) ends | Notes |
|---|---|---|---|---|---|
| Debian 13 "Trixie" | 2025-08-09 `[secondary]` | 2028-08-09 `[secondary]` | 2030-06-30 (Debian LTS) `[secondary]` | 2035-06-30 (Freexian ELTS) `[secondary]` | ELTS is commercial, via Freexian |
| Debian 12 "Bookworm" | 2023-06-10 `[secondary]` | 2026-06-11 or 2026-07-11 — conflict, see A9 | 2028-06-30 (LTS) `[secondary]` | 2033-06-30 (ELTS) `[secondary]` | Latest point release 12.15 (2026-07-11) `[secondary]` |
| Debian 11 "Bullseye" | 2021-08-14 `[secondary]` | 2024-06-30 `[secondary]` | 2026-08-31 (LTS) `[secondary]` | 2031-06-30 (ELTS) `[secondary]` | **LTS ended 2026-08-31 — 3 weeks before this writing** |
| Ubuntu 26.04 LTS | 2026-04-23 `[secondary]` | 2031-04 (standard) `[secondary]` | 2036-04 (Pro/ESM) `[secondary]` | 2038-04 (Legacy add-on) `[unverified]` | Legacy 15-yr window documented for 20.04; assumed same structure |
| Ubuntu 24.04 LTS | 2024-04-25 `[secondary]` | 2029-04 (standard) `[official]` | 2034-04 (Pro/ESM) `[official]` | 2036-04 (Legacy add-on) `[vendor-reported]` | 24.04.5 (Sept 2026) ships HWE kernel 7.0 `[secondary]` |
| Ubuntu 22.04 LTS | 2022-04-21 `[secondary]` | 2027-04 (standard) `[official]` | 2032-04 (Pro/ESM) `[official]` | 2034-04 (Legacy add-on) `[vendor-reported]` | |
| Ubuntu 20.04 LTS | 2020-04-23 `[secondary]` | 2025-05-31 (standard) `[official]` | 2030-04 (Pro/ESM) `[official]` | 2032-04 (Legacy add-on) `[vendor-reported]` | Free personal Pro tier still covers ESM |
| RHEL 10 | 2025-05-20 (or 05-13) `[secondary]` | 2030-05-31 (full support) `[secondary]` | 2035-05-31 (maintenance) `[secondary]` | 2039-05-31 (ELS) `[unverified]` | ELS add-on pricing quote-based |
| RHEL 9 | 2022-05-18 `[secondary]` | 2027-05-31 (full support) `[official]` | 2032-05-31 (maintenance) `[official]` | 2036-05-31 (ELS) `[secondary]` | Current minor: 9.8 (June 2026) `[secondary]` |
| RHEL 8 | 2019-05-07 `[secondary]` | 2024-05-31 (full support) `[official]` | 2029-05-31 (maintenance) `[official]` | 2033-05-31 (ELS) `[secondary]` | 8.10 was the final minor release `[secondary]` |
| RHEL 7 | 2014-06-10 `[secondary]` | 2019-08-06 (full) `[secondary]` | 2024-06-30 (maintenance) `[official]` | 2029-05-31 (ELS) `[secondary]` | TuxCare ELS4 covers 7 until 2029-06-30 `[vendor-reported]` |
| Rocky 9 | 2022-07-14 `[official]` | 2027-05-31 (active) `[official]` | 2032-05-31 (EOL) `[official]` | via CIQ RLC Pro LTS `[vendor-reported]` | Mirrors RHEL lifecycle |
| Rocky 10 | 2025-06-11 `[official]` | 2030-05-31 (active) `[official]` | 2035-05-31 (EOL) `[official]` | via CIQ RLC Pro LTS `[vendor-reported]` | x86_64-v3 minimum |
| Rocky 8 | 2021-06-21 `[official]` | 2024-05-31 (active) `[official]` | 2029-05-31 (EOL) `[official]` | via CIQ/TuxCare ELS `[vendor-reported]` | |
| AlmaLinux 9 | 2022-05-26 `[official]` | 2027-05-31 (active) `[official]` | 2032-05-31 (security) `[official]` | via TuxCare ESU/ELS `[vendor-reported]` | |
| AlmaLinux 10 | 2025-05-27 `[official]` | 2030-05-31 (active) `[official]` | 2035-05-31 (security) `[official]` | via TuxCare ESU/ELS `[vendor-reported]` | 10.1 "Heliotrope Lion" 2025-11-24 `[official]` |
| AlmaLinux 8 | 2021-03-30 `[official]` | 2024-05-01 (active) `[official]` | 2029-03-01 or 2029-05-01 — conflict, see A25 `[secondary]` | via TuxCare ELS `[vendor-reported]` | |
| Oracle Linux 9 | 2022-06-30 `[secondary]` | Premier support: 2032-06 (approx 10 yr) `[unverified]` | Extended: 2035-06 `[unverified]` | — | Oracle lifetime support policy applies |
| Oracle Linux 10 | 2025-06 (GA) `[secondary]` | Premier support: ~2035 `[unverified]` | — | — | Ships UEK 8.1 + RHCK 6.12 |

**CentOS context (why this table looks the way it does):** CentOS Linux 8 reached EOL 2021-12-31; CentOS Linux 7 reached EOL 2024-06-30 `[official]`. CentOS Stream 8 EOL was 2024-05-31; CentOS Stream 9 EOL is expected around 2027 (aligned with RHEL 9 full-support end) `[secondary]`. RHEL 11 is expected around 2028 on the historical ~3-year cadence `[secondary]`.

---

## A3 — Debian 13 "Trixie": the current stable

**Release.** Debian 13 "Trixie" was released 2025-08-09 `[secondary]`, superseding Debian 12 as the stable distribution. The latest point release at observation time is **13.7, dated 2026-09-12** — ten days before this writing `[secondary]`. Regular security support runs until **2028-08-09**; Debian LTS coverage then extends to **2030-06-30**, and Freexian commercial ELTS to **2035-06-30** `[secondary]`.

**Kernel and toolchain.** Trixie ships **Linux 6.12 LTS** `[secondary]`, GCC 14 (default), Python 3.13, and **APT 3.0** with a redesigned solver interface `[secondary]`. The 6.12 kernel is significant for infrastructure: it is an upstream LTS branch (supported by Greg Kroah-Hartman until December 2026 at the earliest per kernel.org policy — exact LTS EOL for 6.12 has moved with announcements, verify before quoting) `[secondary]`.

**Security hardening highlights** `[secondary]`:

- **64-bit `time_t` on all architectures**, including 32-bit ports — the long-planned Y2038 mitigation is complete in Trixie.
- **Intel CET (Control-flow Enforcement Technology) and ARM Pointer Authentication (PAC)** enabled where the hardware supports them.
- Reproducible-build coverage continued to increase; Trixie set a new record for reproducible packages `[secondary]`.
- `systemd` 257, OpenSSL 3.5, and `run0` (systemd's privilege-escalation tool) available as a sudo alternative `[secondary]`.
- `/tmp` is a tmpfs by default on new installations (change introduced in Bookworm, retained) `[secondary]`.

**Architectures.** Trixie supports amd64, arm64, armel, armhf, i386, mips64el, ppc64el, riscv64 and s390x `[secondary]`. Note the riscv64 port is present in Trixie, and Proxmox-relevant x86-64/amd64 remains the primary server target.

**Package management.** APT 3.0 is the headline change: a new, faster dependency solver and a modernized CLI, while remaining backward compatible with `apt-get` scripting interfaces `[secondary]`. Sources now default to the **deb822 format** (`/etc/apt/sources.list.d/*.sources`) rather than one-line `sources.list` entries `[secondary]`.

**Debian 13 for servers — assessment.** Debian's value proposition is unchanged: zero licensing cost, conservative stable package set, no vendor lock-in, and the distribution Proxmox VE itself is built on `[secondary]`. Its weaknesses for enterprise estates are also unchanged: no vendor SLA (unless buying Freexian LTS/ELTS), no FIPS-validated crypto modules, no DISA STIG, and CIS hardening is community/ansible-lockdown based rather than vendor-shipped.

---

