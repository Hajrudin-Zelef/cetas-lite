---
id: etape7-phasea-os/00-os/a34-disa-stig-coverage
title: "A34 — DISA STIG coverage"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["AWS", "Oracle", "United States"]
dates: ["2026-04", "2026-08"]
keywords: ["benchmarks", "distribution", "open source", "packaging"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [519, 571]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: 3e2e015e838651702283c830e7072abcf4dd333b7644edc2994e15f980b5ba90
---

# A34 — DISA STIG coverage

## A34 — DISA STIG coverage

| Distribution | STIG status |
|---|---|
| RHEL 10 | **DISA STIG V1R1** published (secondary confirmation 2026); Red Hat's `rhel10-stig` Ansible role was initially built against an *expected* STIG before DISA published — verify profile currency before use `[secondary]` |
| RHEL 9 | DISA STIG **V2R8** (April 2026) — mature, with SCAP content `[secondary]` |
| RHEL 8 | DISA STIG V8Rx — mature `[secondary]` |
| Ubuntu 24.04 | **DISA Ubuntu 24.04 STIG published**; Canonical's USG consumes DISA STIG profiles for automated hardening `[official]` |
| Ubuntu 22.04 | DISA STIG V2Rx published `[secondary]` |
| Rocky/Alma 9/10 | No DISA-published STIG under their own names; the **RHEL STIG applies technically** (same binaries/packaging) but formal compliance mapping is the deployer's responsibility — auditors vary on accepting this |
| Debian | No DISA STIG |

**STIG vs CIS.** STIGs are mandatory for US DoD systems; CIS benchmarks are the commercial equivalent. A distribution with both (RHEL, Ubuntu LTS) has the widest compliance addressability.

---

## A35 — Common Criteria, SELinux and AppArmor

**Common Criteria.** RHEL 9 holds **Common Criteria EAL4+** certification (with the evaluated configuration documented by Red Hat) `[secondary]`; RHEL 10 evaluation status should be confirmed per release. Ubuntu 22.04 achieved EAL2; Ubuntu 24.04 CC status could not be confirmed in this pass `[unverified]`. Debian, Rocky and AlmaLinux hold no CC certifications.

**Mandatory access control:**

- **SELinux (enforcing, targeted policy)** is default on RHEL, Rocky, AlmaLinux and Oracle Linux. It is the deeper, more complex MAC system — and the one STIG/CIS content is written against in the RHEL family.
- **AppArmor** is default on Ubuntu and Debian (Debian enables it by default since Buster). Simpler profile model; Ubuntu ships profiles for many services out of the box.
- Operational impact: container platforms (Docker/Podman/Kubernetes) interact with both; SELinux volume-label issues (`:z`/`:Z` flags) are the classic RHEL-family container gotcha, while AppArmor profile denials are the Ubuntu equivalent.

---

## A36 — Adoption and market share: read the caveats first

Public "Linux server market share" numbers are **not comparable across sources** — they measure different populations (public web servers vs developer surveys vs vendor surveys) with different methods. The figures below are presented side by side precisely so they are **not** merged into a single ranking.

**W3Techs (web-server fingerprinting, August 2026)** `[independent]`:

- Of surveyed websites: Ubuntu **15.1%**, Debian **5.9%**, CentOS **1.2%**, AlmaLinux **0.3%**, Amazon Linux **0.2%**, Rocky Linux **0.2%**.
- **76.8% of Linux-based sites do not reveal their distribution** — the single largest "segment" is unknown. Any share claim built on this data describes at most a quarter of the web.

**Stack Overflow Developer Survey 2025 (developer usage, personal + professional, overlapping)** `[independent]`:

- Ubuntu: **27.78%** personal / **27.70%** professional use. Debian: **11.39%** / **10.42%**. These are *usage* percentages among developers, not server market share.

**OpenLogic State of Open Source 2025 (enterprise survey)** `[independent]`:

- Paid Linux in use: Ubuntu **38%**, RHEL **28%**, CentOS (legacy) **22%**, Debian **15%**, Rocky **12%**, AlmaLinux **11%**, CentOS Stream **8%** (multiple selections allowed — sums exceed 100%).

**CentOS-alternatives migration survey (secondary compilation)** `[secondary]`:

- Post-CentOS destinations: Ubuntu **38%**, RHEL **28%**, CentOS legacy **22%**, Rocky **12%**, AlmaLinux **11%**. Survey-specific; do not compare with the figures above.

**What can actually be concluded.** Ubuntu is the most widely deployed general server Linux; Debian is second; the RHEL family (paid + rebuilds) dominates regulated enterprise; Rocky and AlmaLinux split the rebuild community roughly evenly with Rocky slightly ahead in most surveys. Everything more precise than that is survey artifact.

---

