---
id: etape7-phasec-iac/00-iac/part-15
title: "Step 7 Phase C — IaC & Platform Automation (part 15)"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: []
dates: ["2026-03-31", "2026-09-22"]
keywords: ["compute", "packaging", "research"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [741, 750]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: 7deaf610419da833ea2eaabccec2c76d71b2de70b1ed359171f2561e3e226f6a
---

# Step 7 Phase C — IaC & Platform Automation (part 15)

- **Step 6 Phase E2** covered Ansible/Terraform/OpenTofu for *network-device* automation (network resource modules, NAPALM/Nornir, Batfish network validation, NetBox-as-source-of-truth for networks). This file covers the *platform/compute* angle; treat the two as companions, not duplicates.
- **Step 7 Phase A** (OS: Debian/Ubuntu/RHEL/Rocky) covers the operating systems these tools manage; **Phase B** (containers/Kubernetes) covers the clusters Argo CD/Flux deploy to; **Phase D** (Proxmox/backup) covers the virtualization layer Terraform/OpenTofu and cloud-init commonly provision.
- For RAG ingestion: the highest-signal tables are M1 (timeline), M2 (HCP Terraform tiers), M8 (Argo CD vs Flux), O (version pins), and B6 (IaC decision matrix). The conflicts list (Section I) should be indexed as caveat metadata, not facts.
- Prices in this file are point-in-time (2026-09-22 cutoff) and vendor packaging changes frequently — especially HCP Terraform (restructured 2026-03-31), Pulumi Cloud, and the TACO vendors.
- Suggested v2 follow-ups: named enterprise IaC case studies; 2026 market-share figures; verification of the gap list G1–G11 against primary sources.
- Methodology note: all web research was read-only; no logins, purchases, or form submissions were performed.
- File stats: 16 top-level sections (A–P), 17 H2 headings, one H1, zero fenced code blocks (indented blocks used for examples).
- Line-count target: ≥750 lines met at final verification (2026-09-22).

*Final end of Step 7 Phase C file. Total sections A–P. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*
