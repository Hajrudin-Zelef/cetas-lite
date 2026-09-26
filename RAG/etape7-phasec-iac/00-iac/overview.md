---
id: etape7-phasec-iac/00-iac/overview
title: "Step 7 Phase C — IaC & Platform Automation"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: []
dates: ["2023-08-10", "2024-04", "2024-05-20", "2024-11-04", "2025-02", "2025-06", "2025-07-21", "2025-11-03", "2025-11-30", "2026-03-12", "2026-04-30", "2026-05-31", "2026-06-02", "2026-06-18", "2026-09-08", "2026-09-09", "2026-09-22", "2026-11-30", "2027-05-31", "2027-11-30"]
keywords: ["acquisition", "apache", "benchmarks", "compute", "governance", "license", "pricing", "research", "training"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [1, 62]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: a0f884070ce5468b2bb663de879e701732a0bcf2f53ddc520454e09108f90e78
---

# Step 7 Phase C — IaC & Platform Automation

> **Scope:** Infrastructure-as-Code and platform automation from the *platform/VM/container* angle: Ansible ecosystem, Terraform/OpenTofu/Pulumi/Crossplane, Packer, cloud-init, golden-image pipelines, GitOps (Argo CD, Flux), secrets management (Vault, SOPS, ESO), policy as code (Sentinel, OPA/Gatekeeper, Kyverno), state backends, drift detection, and TACO platforms (Spacelift, env0, Scalr). Network-device automation (Ansible network modules, NAPALM, Nornir, Batfish for network validation) was already covered in **Step 6 Phase E2** — this file cross-references it and does not duplicate it.
>
> **Research date / cutoff:** 2026-09-22. All "current" statements are relative to this date.
>
> **Method:** Read-only web research (search + page fetch). One writer, append-only build. No workspace files other than this one were touched.
>
> **Provenance legend:** `[official]` = vendor/project documentation or announcement. `[vendor-reported]` = vendor marketing or vendor-supplied figures. `[independent]` = third-party benchmarks, analyst data, or reproducible tests. `[secondary]` = aggregators, blogs, deal-trackers, training repos citing primary sources. `[unverified]` = claim seen in one place only or not re-verified in this pass.

## Landscape map (2026)

- The IaC market in 2026 is a three-way split: Terraform (HashiCorp/IBM, BSL 1.1) remains the incumbent; OpenTofu (Linux Foundation, MPL 2.0) is the open-governance fork shipping CLI features Terraform lacks; Pulumi (Apache 2.0) owns the general-purpose-language niche [secondary](https://shattered.io/terraform-vs-pulumi-vs-opentofu-2026/).
- The HashiCorp BSL 1.1 relicense of 2023-08-10 still shapes 2026 decisions: it covers Terraform, Packer, Vault, Consul, Nomad, Boundary, Waypoint (and Vagrant) for all future releases; APIs, SDKs and "almost all other libraries" stayed MPL 2.0 [official](https://www.hashicorp.com/en/blog/hashicorp-adopts-business-source-license).
- IBM completed its **$6.4 billion** acquisition of HashiCorp in **February 2025**; Terraform Cloud was rebranded **HCP Terraform** in April 2024 [secondary](https://www.env0.com/blog/terraform-cloud).
- Ansible remains the dominant agentless config-management tool; Red Hat's commercial vehicle is the Ansible Automation Platform (AAP), with AWX as the upstream [secondary](https://www.g2.com/products/red-hat-ansible-automation-platform/pricing).
- GitOps is the dominant Kubernetes deployment model; Argo CD (CNCF graduated 2022) leads on adoption, Flux continues as the lightweight alternative [secondary](https://devtoollab.com/blog/best-gitops-tools).

---

## A. Ansible ecosystem

### A1. ansible-core releases (2026 status)

Release train (data via endoflife.date, last updated 2026-09-09) [secondary](https://endoflife.date/ansible-core):

| Release | Released | EOL | Latest patch (as of 2026-09-08) | Control-node Python | Managed-node Python |
|---|---|---|---|---|---|
| 2.21 | 2026-05-31 | 2027-11-30 | 2.21.4 | 3.12–3.14 | 3.9–3.14 (WinRM PS 5.1–7) |
| 2.20 | 2025-11-03 | 2027-05-31 | 2.20.9 | 3.12–3.14 | 3.9–3.14 |
| 2.19 | 2025-07-21 | 2026-11-30 | 2.19.13 | 3.11–3.13 | 3.8–3.13 |
| 2.18 | 2024-11-04 | 2026-05-31 (EOL reached) | 2.18.19 | 3.11–3.13 | 3.8–3.13 |
| 2.17 | 2024-05-20 | 2025-11-30 (EOL reached) | 2.17.14 | 3.10–3.12 | 3.7–3.12 |

- **ansible-core 2.21** is the current feature release (released 2026-05-31) [secondary](https://endoflife.date/ansible-core).
- **ansible-core 2.18 reached EOL on 2026-05-31**; 2.19 goes EOL 2026-11-30 — operators on either should be migrating to 2.20/2.21 in H2 2026 [secondary](https://endoflife.date/ansible-core).
- The **Ansible community package** (batteries-included: core + Galaxy collections) is on major **14**, released 2026-06-02, bundling ansible-core 2.21, latest **14.4.0** on 2026-09-08; community 13 (core 2.20) went EOL 2026-06-18; community 12 (core 2.19) and 11 (core 2.18) also listed [secondary](https://github.com/endoflife-date/endoflife.date/blob/HEAD/products/ansible.md).
- Practical note: managed-node Python floor keeps rising — 2.21 requires Python ≥3.9 on targets, which finally drops EL7-era 3.6 targets from support [secondary](https://endoflife.date/ansible-core).

### A2. AWX (upstream) and AAP 2.6 component versions

- **AWX** is the open-source upstream of Ansible controller/AAP: web UI, REST API, job scheduling, RBAC, inventory sync with cloud providers [secondary](https://medium.com/@btech-engineering/ansible-awx-infrastructure-automation-on-top-of-kubernetes-9c81986131c4).
- Since AWX 18.0 the supported deployment method is the **AWX Operator on Kubernetes** (k3s minimum cited: 4 cores / 8 GB RAM single-node lab) [secondary](https://medium.com/@btech-engineering/ansible-awx-infrastructure-automation-on-top-of-kubernetes-9c81986131c4).
- AWX Operator **2.19.1** shipped with **AWX 24.6.1** (June 2025) [official](https://github.com/ansible/awx-operator/releases).
- The community **awx-operator Helm chart** reached **v3.2.1** on 2026-03-12 [official](https://forum.ansible.com/t/release-announcement-awx-operator-helm-chart-v3-2-1/45518).
- A third-party AAP 2.6 source-bundle mirror lists the downstream component set as: AWX **26.0.0**, EDA server **1.2.12**, EDA UI 2.6.13, platform gateway 2.6.x, AWX operator 2.6 [secondary](https://github.com/fitbeard/automation-platform/blob/HEAD/README.md). Treat the exact AAP 2.6 mapping as `[secondary]` — Red Hat's official AAP release notes should be consulted before quoting it as fact.
- AWX remains free/open; the commercial AAP adds support, certified content, and the automation hub/mesh [secondary](https://www.trustradius.com/compare-products/jfrog-artifactory-vs-red-hat-ansible-automation-platform).

### A3. Ansible Automation Platform pricing (2026 signals)

- Aggregator pricing pages (G2, updated 2026-04-30) list AAP in three "Tower" editions: **Basic $5,000/yr**, **Enterprise $10,000/yr**, **Premium $14,000/yr**, with a free trial noted [secondary](https://www.g2.com/products/red-hat-ansible-automation-platform/pricing).
- TrustRadius 2026 comparison pages repeat the same $5,000/$10,000/$14,000 per-year ladder [secondary](https://www.trustradius.com/compare-products/microsoft-system-center-vs-red-hat-ansible-automation-platform).
- **Caveat:** Red Hat's official AAP pricing is historically **per managed node** (Starter/Standard/Premium tiers with node bands), not the flat "Tower" ladder above; aggregator figures may reflect legacy Ansible Tower SKUs or simplified listings. Do not budget from these numbers without a Red Hat quote — flagged as `[secondary]`, unverified against Red Hat's current price book.
- G2 reviewers (366 reviews cited) praise the agentless architecture and human-readable YAML; AAP is positioned for the enterprise segment (48.5% of reviews) [secondary](https://www.g2.com/compare/pliant-vs-red-hat-ansible-automation-platform).

### A4. Galaxy collections, Execution Environments, automation mesh

- **Ansible Galaxy** (`galaxy.ansible.com`) is the community collection hub; AAP adds the certified **automation hub** [secondary](https://galaxy.ansible.com).
- **Execution Environments (EEs)** — container images bundling ansible-core, collections and Python deps, built with `ansible-builder` — are the standard way to run automation consistently across controller, mesh hops and CI [secondary](https://github.com/fitbeard/automation-platform/blob/HEAD/README.md).
- **Automation mesh** replaces the old isolated-node model for scaling execution capacity across networks without direct controller-to-target connectivity [secondary](https://medium.com/@btech-engineering/ansible-awx-infrastructure-automation-on-top-of-kubernetes-9c81986131c4).
- Network-focused collections and modules (e.g. `ansible.netcommon`, vendor network collections) are covered in **Step 6 Phase E2**; this file covers the platform/compute-side collections only.

