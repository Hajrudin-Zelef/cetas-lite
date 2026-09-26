---
id: etape7-phasea-os/00-os/a42-decision-guide-which-distribution-for-which-role
title: "A42 — Decision guide: which distribution for which role"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["Oracle", "United States"]
dates: ["2025-05-13", "2025-05-20", "2026-06-11", "2026-07-11", "2026-09-22", "2029-03-01", "2029-05-01"]
keywords: ["distribution", "advisory", "benchmark", "pricing", "research"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [670, 735]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: d1b0b70bdbda58562c53dfda6d451ea422ab3e99cc3cca14a09662cb2375ef42
---

# A42 — Decision guide: which distribution for which role

## A42 — Decision guide: which distribution for which role

**Proxmox VE hypervisor hosts.** Debian — non-negotiable in practice: Proxmox VE *is* Debian 13 (v9) / Debian 12 (v8) with Proxmox repositories. Run the matching Debian base; do not substitute Ubuntu or RHEL-family on the host.

**General-purpose VMs and containers (homelab/small business).** Ubuntu 24.04 LTS (or 26.04 LTS for new builds from late 2026): largest package ecosystem, best cloud-init support, free 5-machine Pro tier for ESM, Landscape available. Debian 13 is the equally valid minimalist alternative with zero subscription surface.

**Regulated / compliance-mandated workloads (FIPS, STIG, CIS, CC).** RHEL 9/10 with appropriate subscriptions; Ubuntu Pro 24.04 as the Debian-family alternative; CIQ RLC Pro Hardened for a Rocky-based FIPS story; TuxCare-backed AlmaLinux where AlmaLinux is already standardized. Debian and community rebuilds are excluded where FIPS/STIG is contractual.

**CentOS 7/8 refugee estates.** Rocky 9 or AlmaLinux 9 via `migrate2rocky` / `almalinux-deploy` (in-place, low-risk); then plan 9 → 10 via ELevate/Leapp before 2027–2030. If hardware predates 2015, AlmaLinux (x86_64-v2 build) is mandatory.

**Oracle-adjacent estates (Oracle DB, OCI-heavy).** Oracle Linux with Premier support: Ksplice zero-downtime patching for database hosts, support synergy with Oracle products, UEK tuned for Oracle workloads.

**Kubernetes / container node fleets.** Fedora CoreOS (or RHEL CoreOS under OpenShift), openSUSE MicroOS, or minimal Ubuntu/Debian images — not full RHEL/Debian installs. See 7B for the container layer.

**Longest support horizon without migration.** RHEL 9/10 + ELS (13 years), Ubuntu LTS + Legacy add-on (12–15 years), Debian + Freexian ELTS (10 years). All paid except Debian's community LTS window.

---

## A43 — Gaps, conflicts and non-comparable figures register

1. **Debian 12 regular-support end:** 2026-06-11 vs 2026-07-11 — unresolved; both recorded (A4). Operational impact: low (LTS covers both dates' aftermath identically).
2. **RHEL 10 GA day:** 2025-05-13 vs 2025-05-20 — month firm, day unverified (A13).
3. **AlmaLinux 8 security-support end:** 2029-03-01 (endoflife.date) vs 2029-05-01 (AlmaLinux wiki) — unresolved (A22).
4. **Ubuntu 26.04 specifics** (kernel 7.0, sudo-rs default, GNOME 50, lifecycle to 2036/2038): secondary press only; Canonical primary source not fetched — flagged throughout (A7).
5. **Market-share figures** (W3Techs vs Stack Overflow vs OpenLogic vs CentOS-alternatives survey): different populations and methods — presented side by side, must not be merged (A36).
6. **RHEL list pricing:** SKU-dense and frequently repackaged; figures are order-of-magnitude `[secondary]` (A16, A41).
7. **Canonical Landscape standalone pricing:** no current Canonical-published figure found; third-party ~$15/host/month marked `[unverified]` (A11, A41).
8. **Rocky 10 / AlmaLinux 10 current minor** at 2026-09-22: not verified from project pages in this pass (A18, A21).
9. **Ubuntu 24.04 FIPS 140-3 completion:** in NIST queue per Canonical discussions — do not claim certified (A33).
10. **RHEL 10 STIG V1R1:** confirmed via secondary source only; Red Hat's own tooling was built against a pre-publication draft — verify profile currency (A34).
11. **Oracle Linux 10 lifecycle dates:** derived from Oracle's lifetime support policy, not a per-release published EOL — marked `[unverified]` (A2).
12. **TuxCare ESU/ELS current pricing** for Rocky/Alma 8/9: quote-based at research time (A27).

---

## A44 — Glossary

- **ABI compatibility:** guarantee that compiled binaries keep working across releases — AlmaLinux's post-2023 promise vs RHEL.
- **AppStream:** RHEL-family repository of versioned application streams (e.g., multiple PostgreSQL majors on one OS major).
- **backports:** newer packages rebuilt for a stable release (Debian Backports).
- **bootc:** project distributing bootable Linux as OCI container images.
- **CIS Benchmark:** consensus hardening checklist from the Center for Internet Security.
- **deb822:** the structured multi-line format for APT source entries (`.sources` files).
- **DISA STIG:** US Defense Information Systems Agency Security Technical Implementation Guide — mandatory hardening for DoD systems.
- **DSA/DLA:** Debian Security Advisory / Debian LTS Advisory.
- **EAL:** Evaluation Assurance Level (Common Criteria).
- **ELS/ESU:** Extended Life (Cycle) Support / Extended Security Updates — paid post-EOL patching.
- **EUS:** Extended Update Support — pinning a RHEL minor with backported security fixes.
- **FIPS 140-3:** US federal cryptographic-module validation standard.
- **GA kernel vs HWE:** Ubuntu's fixed original kernel vs the rolling Hardware Enablement kernel.
- **Ignition:** declarative first-boot provisioning for Fedora CoreOS.
- **Ksplice/kpatch/Livepatch/KernelCare:** live kernel patching systems (Oracle/Red Hat/Canonical/TuxCare).
- **Leapp/ELevate:** major-version in-place upgrade frameworks (Red Hat / AlmaLinux).
- **LTS/ELTS:** Debian Long-Term Support / Extended LTS (commercial, Freexian).
- **MAC:** mandatory access control (SELinux, AppArmor).
- **RHCK:** Red Hat Compatible Kernel (Oracle Linux's RHEL-identical kernel option).
- **rpm-ostree:** hybrid image/package system behind Fedora CoreOS atomic updates.
- **SCAP/USG:** machine-readable compliance automation (OpenSCAP content / Ubuntu Security Guide).
- **UEK:** Unbreakable Enterprise Kernel (Oracle Linux's optimized kernel).
- **USN/RHSA:** Ubuntu Security Notice / Red Hat Security Advisory.
- **x86_64-v2/v3:** x86-64 microarchitecture feature levels — v3 excludes pre-~2015 CPUs.

---

## A45 — Source index (verbatim URLs, in order of first use)

