---
id: etape7-phasea-os/00-os/a27-tuxcare-extended-lifecycle-and-live-patching-vendor
title: "A27 — TuxCare: extended lifecycle and live patching vendor"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["Oracle"]
dates: ["2026-09", "2029-06-30"]
keywords: ["distribution", "pricing", "research"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [399, 458]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: 3bc22f5f0a97161e683ec04aac027a30a5fa66a156710b622fd62b6cadb427d3
---

# A27 — TuxCare: extended lifecycle and live patching vendor

## A27 — TuxCare: extended lifecycle and live patching vendor

**Who TuxCare is.** The enterprise-support arm of CloudLinux, selling extended-lifecycle and live-patching services for RHEL-family distributions (including the rebuilds) `[vendor-reported]`.

**Product lines relevant here** `[vendor-reported]`:

- **ELS (Extended Lifecycle Support):** vendor-grade security patching for EOL distributions — covers CentOS 7 (ELS4 to 2029-06-30), AlmaLinux/Rocky 8 and 9 after their community EOLs.
- **ESU (Extended Security Updates):** for *still-supported but specific* minors — observed covering **AlmaLinux 9.2, 9.6 and Rocky 9.6** (per TuxCare docs, September 2026), i.e., the EUS-like pinning that the community rebuilds don't offer natively.
- **KernelCare:** live kernel patching (see A32), compatible with RHEL, CentOS, Rocky, AlmaLinux, Oracle Linux, Ubuntu and Debian — but **must not run alongside Canonical Livepatch or kpatch on the same host** `[vendor-reported]`.
- **FIPS for AlmaLinux:** TuxCare reports **5 FIPS 140-2/140-3 validated cryptographic modules** for AlmaLinux `[vendor-reported]`.

**Pricing observed:** TuxCare's CentOS 7 ELS listed at **$4.25/month or $42.50/year per system** `[secondary]` (date-labeled; current ELS/ESU pricing for Rocky/Alma 8/9 was quote-based at research time — gap flagged).

**Assessment.** TuxCare is the "insurance policy" vendor for the rebuild ecosystem: it backfills the two things community rebuilds lack — long-tail EOL patching and FIPS validation. Any Rocky/Alma compliance story should name TuxCare explicitly.

---

## A28 — Kernel version master table

| Distribution / release | Kernel at GA | Kernel notes |
|---|---|---|
| Debian 13 | 6.12 LTS `[secondary]` | Upstream LTS branch |
| Debian 12 | 6.1 LTS `[secondary]` | Upstream LTS branch |
| Debian 11 | 5.10 LTS `[secondary]` | Upstream LTS branch |
| Ubuntu 26.04 | 7.0 `[secondary]` | Non-LTS upstream; HWE will roll forward |
| Ubuntu 24.04 (GA / HWE) | 6.8 GA / 7.0 HWE (24.04.5) `[secondary]` | HWE rolls at point releases |
| Ubuntu 22.04 (GA / HWE) | 5.15 GA / 6.8 HWE (22.04.5) `[secondary]` | |
| Ubuntu 20.04 (GA / HWE) | 5.4 GA / 5.15 HWE `[secondary]` | |
| RHEL 10 | 6.12.0 (Red Hat tree) `[secondary]` | Backport-heavy; not comparable to upstream 6.12 |
| RHEL 9 | 5.14.0 (Red Hat tree) `[secondary]` | Backport-heavy; not comparable to upstream 5.14 |
| RHEL 8 | 4.18.0 (Red Hat tree) `[secondary]` | |
| Rocky 10 / AlmaLinux 10 | 6.12 family `[official]` | Track RHEL kernel |
| Rocky 9 / AlmaLinux 9 | 5.14 family `[official]` | Track RHEL kernel |
| Oracle Linux 10 | UEK 8.1 (6.12-based) default; RHCK 6.12 alt `[secondary]` | UEK diverges from RHEL kernel |
| Oracle Linux 9 | UEK 7 (5.15-based) default; RHCK 5.14 alt `[secondary]` | 9.6 moved default UEK to 6.12-based UEK 8 |
| Proxmox VE 9 (host) | 6.14.8-2 (Ubuntu-25.04-based) `[secondary]` | Proxmox builds its own kernel on Debian 13 |
| Proxmox VE 8 (host) | 6.8 (Ubuntu-based) `[secondary]` | On Debian 12 base |

**Reading guide.** Red Hat-family kernels (RHEL/Rocky/Alma/RHCK) keep a fixed major version for the entire major release and backport drivers and fixes — a RHEL 9 "5.14" kernel from 2026 supports far newer hardware than upstream 5.14 did in 2021. Debian/Ubuntu kernels track upstream LTS branches more closely. Never compare a RHEL-family kernel version number directly against a Debian/Ubuntu one for hardware-support conclusions.

---

## A29 — Package management compared

| Distribution | Manager | Format | Notes |
|---|---|---|---|
| Debian 12/13 | APT 3.0 (13) / 2.x (12) | .deb | deb822 `.sources` format default in 13; `apt` CLI stable |
| Ubuntu | APT (2.x generation) | .deb | Snaps for desktop/IoT; server stays deb-first |
| RHEL 10 | DNF 5 | .rpm | Major DNF rewrite: faster, smaller, unified `dnf5` CLI |
| RHEL 9 | DNF 4 (libdnf) | .rpm | `yum` shim retained |
| Rocky/Alma/Oracle | DNF (matching RHEL major) | .rpm | Module streams (AppStream) for parallel package versions |
| Proxmox VE | APT (Debian base) | .deb | Proxmox repos (`pve-no-subscription` / enterprise) layered on Debian |

**Operational notes.**

- **Unattended upgrades:** Debian/Ubuntu — `unattended-upgrades` (security-only by default on server); RHEL family — `dnf-automatic` with configurable reboot behavior. Both are the standard mechanism for auto-patching fleets (pairs with Landscape on Ubuntu, or Ansible on either).
- **Version pinning:** `apt-mark hold` / apt pinning (Debian/Ubuntu); `versionlock` plugin (RHEL family). Essential for kernel holds on hypervisors.
- **Module/AppStream streams (RHEL family):** allow e.g. PostgreSQL 15/16 or Node.js 20/22 in parallel on one major — the answer to "Debian stable is too old for app X" without leaving the major release.
- **Backports (Debian):** `bookworm-backports` / `trixie-backports` provide newer kernels and selected packages rebuilt for stable — the sanctioned way to get newer hardware support on Debian stable.

