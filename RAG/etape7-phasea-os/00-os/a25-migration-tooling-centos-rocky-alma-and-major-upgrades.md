---
id: etape7-phasea-os/00-os/a25-migration-tooling-centos-rocky-alma-and-major-upgrades
title: "A25 — Migration tooling: CentOS → Rocky/Alma, and major upgrades"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["Oracle"]
dates: ["2025-06", "2025-09", "2025-12", "2026-08-01"]
keywords: ["agents", "open source", "pricing"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [352, 398]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: 04ddc652f64460e518ba1a80a4421a0d035417f44d042f064fecf97d2df6fc8d
---

# A25 — Migration tooling: CentOS → Rocky/Alma, and major upgrades

## A25 — Migration tooling: CentOS → Rocky/Alma, and major upgrades

**From CentOS/RHEL/Oracle to the rebuilds:**

- **Rocky:** `migrate2rocky.sh` converts CentOS 7/8/9, RHEL, AlmaLinux and Oracle Linux installations to Rocky in place `[official]`.
- **AlmaLinux:** `almalinux-deploy.sh` converts CentOS 7/8/9, RHEL, Rocky and Oracle Linux to AlmaLinux `[official]`; `almalinux-deploy` also handles Oracle Linux specifically.

**Major-version upgrades (9 → 10):**

- **ELevate** (AlmaLinux project): Leapp-based framework supporting upgrades between majors (8→9, and **9→10 support added in 2025**, Leapp version `0.22.0-3.elevate.1` observed) `[official]`. In December 2025 the ELevate work was **contributed upstream into the Leapp project itself** `[secondary]` — a sign of its maturity.
- **Rocky:** documents Leapp-based 8→9 and 9→10 upgrade paths adapted from RHEL's tooling `[secondary]`.
- **RHEL:** native Leapp (A17); 9.6 → 10.0 is a supported path.

**Migration evidence (adoption data).** Per the OpenLogic *State of Open Source 2025* survey, among organizations migrating off CentOS, **~40% still ran CentOS in some capacity** while Rocky (12%) and AlmaLinux (11%) were the leading landing zones, with Ubuntu (38% in a CentOS-alternatives survey) absorbing a large share of non-RHEL-family migrations `[independent]`. A separate secondary compilation claims **over 40% of CentOS users migrated to Rocky or AlmaLinux** `[secondary]`. These figures come from different surveys with different populations — **not comparable**, recorded side by side (see A36).

---

## A26 — Oracle Linux: the secondary comparison

**Position.** Oracle Linux is a RHEL rebuild with two differentiators: Oracle's **Unbreakable Enterprise Kernel (UEK)** as the default, and **Ksplice** zero-downtime patching bundled with Premier support `[official]`. It is free to download, use and distribute; paid support is optional but required for Ksplice and indemnification.

**Oracle Linux 10** (GA June 2025) `[secondary]`:

- **UEK 8.1** (based on upstream 6.12 LTS — kernel `6.12.0-101.33.4.3.el10uek` in the September 2025 OCI platform image) as the default kernel `[secondary]`.
- **Red Hat Compatible Kernel (RHCK) 6.12** as the alternative boot option `[secondary]`.
- Valkey replaces Redis in the application streams; post-quantum crypto (ML-KEM, ML-DSA) as technology preview; OpenSSH keystroke-obfuscation hardening `[secondary]`.

**UEK lineage** (for kernel planning) `[secondary]`:

| UEK release | Base kernel | Ships with |
|---|---|---|
| UEK 8 | 6.12 LTS | Oracle Linux 9 (from 9.6) and 10 |
| UEK 7 | 5.15 LTS | Oracle Linux 8 and 9 |
| UEK 6 | 5.4 LTS | Oracle Linux 7 and 8 |

**Support pricing (Oracle price list dated 2026-08-01, USD per physical CPU pair per year)** `[official]`:

| Tier | Price |
|---|---|
| Basic | $699 |
| Premier (includes Ksplice zero-downtime patching) | $1,399 |
| Premier Plus | $2,499 |

**Assessment.** Oracle Linux is the rational choice for Oracle-database-adjacent estates (support synergy, Ksplice for zero-downtime DB hosts) and a credible free RHEL alternative generally. Caveats: Oracle's support culture and audit reputation are widely discussed in the industry `[independent]`; UEK diverges from the RHEL kernel, so RHEL-targeted kernel tooling (some eBPF/security agents) must be validated against UEK.

---

