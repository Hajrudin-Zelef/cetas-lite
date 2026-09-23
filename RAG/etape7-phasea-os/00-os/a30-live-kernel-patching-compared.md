---
id: etape7-phasea-os/00-os/a30-live-kernel-patching-compared
title: "A30 — Live kernel patching compared"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["Oracle", "United States"]
dates: []
keywords: ["agents", "benchmark", "benchmarks", "distribution", "pricing"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [459, 518]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: b8c95919287a28b65e5d8e22d7e34bf35e3473857320d9bbce2273f2e9228ede
---

# A30 — Live kernel patching compared

## A30 — Live kernel patching compared

| Solution | Vendor | Distributions | Model | Notes |
|---|---|---|---|---|
| **Canonical Livepatch** | Canonical | Ubuntu LTS | Included with Ubuntu Pro (free tier incl.) | Patches running kernel for CVEs without reboot; covers the GA/HWE kernels of supported LTS |
| **kpatch** | Red Hat | RHEL | Included with subscription | Same-host coexistence with other livepatch agents unsupported |
| **Ksplice** | Oracle | Oracle Linux (UEK + RHCK) | Requires Premier support | Extends beyond kernel to **glibc and OpenSSL userspace** zero-downtime patching `[vendor-reported]` |
| **KernelCare** | TuxCare | RHEL-family, Ubuntu, Debian, others | Commercial per-node | Must **not** be combined with Livepatch or kpatch on the same host `[vendor-reported]` |

**Operational limitation (all vendors).** Live patching cannot safely implement changes that alter kernel data structures; such CVEs still require a reboot, and vendors document which CVEs are livepatch-eligible per release `[secondary]`. Live patching reduces reboot frequency — it does not eliminate reboots. Plan maintenance windows accordingly; do not promise "zero reboots ever" on the basis of any of these products.

**Pricing note.** A widely circulated older Canonical datasheet quotes Livepatch-related pricing that is **historical only** — it must not be presented as 2026 pricing `[secondary]`. Current Livepatch entitlement is bundled in Ubuntu Pro tiers (A10).

---

## A31 — Security patching cadence and embargoes

**Debian.** The Debian Security Team issues advisories (DSA) for stable; LTS security via the LTS team (DLA advisories). Turnaround is volunteer-driven; critical CVEs are typically patched within days, but there is no SLA `[secondary]`.

**Ubuntu.** Canonical publishes USNs (Ubuntu Security Notices) with CVE tracking; ESM updates continue for 10 years on LTS. Canonical participates in the upstream embargo process (`distros` list) so supported releases get coordinated-day patches `[secondary]`.

**RHEL/Rocky/Alma/Oracle.** Red Hat publishes RHSA/RHBA/RHEA errata with severity ratings and CVSS; fixes are backported to the fixed kernel/userspace of each supported major/minor. Rocky and AlmaLinux rebuild the errata with a short lag (typically hours to days) `[secondary]`; Oracle publishes its own errata for both UEK and RHCK. None of the rebuilds participate in the embargoed `distros` list under their own name to the same extent — in practice their patches land slightly after Red Hat's, which is the residual risk of the rebuild model.

**Key takeaway for patching SLAs.** Only paid subscriptions (RHEL, Ubuntu Pro+Support, CIQ RLC, TuxCare, Oracle Premier) carry contractual CVE response commitments. Community rebuilds and plain Debian/Ubuntu inherit upstream speed without a contract.

---

## A32 — CIS hardening per distribution

| Distribution | CIS Benchmark | Automation | Source |
|---|---|---|---|
| Ubuntu 24.04 | CIS Ubuntu 24.04 Benchmark v1.0.0 | **Ubuntu Security Guide (USG)** — Canonical's tool that generates audit/remediation content from CIS and DISA profiles | `[official]` |
| Ubuntu 22.04/20.04 | CIS benchmarks published | USG + `usg` CLI | `[official]` |
| RHEL 10 | CIS RHEL 10 Benchmark (v1.0.0 published 2025) `[secondary]` | **scap-security-guide** (ships in RHEL as `scap-security-guide` package) + OpenSCAP | `[official]` |
| RHEL 9/8 | CIS RHEL 9/8 Benchmarks | scap-security-guide + Ansible remediations | `[official]` |
| Rocky/Alma 9/10 | Same CIS RHEL benchmarks apply | scap-security-guide works (profile names reference RHEL); community **ansible-lockdown** roles | `[secondary]` |
| Debian 12/13 | CIS Debian 12 Benchmark exists; Debian 13 benchmark `[unverified]` in this pass | Community ansible-lockdown roles; no vendor tooling | `[secondary]` |
| Oracle Linux 9/10 | CIS Oracle Linux benchmarks published | scap-security-guide (Oracle ships profiles) | `[secondary]` |

**USG vs OpenSCAP.** Canonical's USG is the Ubuntu-native equivalent of Red Hat's OpenSCAP/scap-security-guide workflow: both produce machine-readable audit results and automated remediation for CIS and STIG profiles. USG is tied to Ubuntu Pro for the compliance profiles; OpenSCAP content ships in RHEL itself.

---

## A33 — FIPS 140-3 validation status

| Distribution | Status (Sept 2026) |
|---|---|
| RHEL 9 | FIPS 140-3 validated modules (kernel crypto, OpenSSL, NSS, GnuTLS, libgcrypt) — Red Hat maintains the CMVP certificate list `[official]` |
| RHEL 10 | FIPS 140-3 validation in progress; Red Hat publishes "FIPS compliance" documentation per release — confirm current certificate numbers on the NIST CMVP site before contractual claims `[official]`/`[unverified]` |
| Ubuntu 24.04 | FIPS 140-3 modules **assessed and in the NIST queue awaiting final CMVP certification** as of Canonical's 2025/2026 compliance discussions; Ubuntu Pro FIPS certified images for 20.04/22.04 exist (FIPS 140-2/140-3) `[secondary]` — **do not claim 24.04 FIPS 140-3 validation is complete** without a current NIST/Canonical confirmation |
| Ubuntu 22.04/20.04 | FIPS 140-2 certified modules; 140-3 modules available via Pro `[official]` |
| Rocky 10/9 | Via **CIQ RLC Pro Hardened**: CIQ reports FIPS 140-3 certificates **#5117, #5116, #5113, #5095** `[vendor-reported]` |
| AlmaLinux 9/10 | Via **TuxCare**: 5 FIPS-validated modules reported for AlmaLinux `[vendor-reported]` |
| Debian 12/13 | **No FIPS-validated modules** — no vendor pursues CMVP validation for Debian; disqualifies Debian where FIPS is contractual |
| Oracle Linux | Oracle maintains FIPS-validated modules for UEK/RHCK on supported releases; confirm certificate numbers per release `[vendor-reported]` |

**Why this matters.** FIPS 140-3 is a procurement checkbox in US federal and many regulated-industry contracts. The practical shortlist for FIPS-mandated Linux in 2026 is **RHEL, Ubuntu Pro, CIQ RLC Pro Hardened, TuxCare-covered AlmaLinux, Oracle Linux** — Debian and community Rocky/Alma are excluded unless a vendor wrapper is purchased.

---

