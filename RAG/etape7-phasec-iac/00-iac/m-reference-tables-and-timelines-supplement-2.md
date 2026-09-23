---
id: etape7-phasec-iac/00-iac/m-reference-tables-and-timelines-supplement-2
title: "M. Reference tables and timelines (supplement 2)"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: []
dates: ["2023-08-10", "2024-01-09", "2024-04-30", "2025-05-06", "2025-06-24", "2025-07-21", "2025-11-03", "2025-12-09", "2026-03-12", "2026-03-31", "2026-05-14", "2026-05-31", "2026-06-02", "2026-07-28", "2026-08-04", "2026-08-31", "2026-09-14"]
keywords: ["acquisition", "cost", "license", "open source", "pricing"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [527, 581]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: 75201e9b65e6fb11009f047d82367a4f44d57d93c43ea68f8ea1d8f285e1d6ee
---

# M. Reference tables and timelines (supplement 2)

## M. Reference tables and timelines (supplement 2)

### M1. IaC licensing and market timeline (2023–2026)

| Date | Event | Provenance |
|---|---|---|
| 2023-08-10 | HashiCorp adopts BSL 1.1 for all future releases (Terraform, Packer, Vault, Consul, Nomad, Boundary, Waypoint) | [official](https://www.hashicorp.com/en/blog/hashicorp-adopts-business-source-license) |
| 2023-08 | OpenTF fork announced (later renamed OpenTofu) in response to BSL | [secondary](https://www.theregister.com/tag/hashicorp) |
| 2024-01-09 | OpenTofu 1.6 GA (first stable, Linux Foundation) | [secondary](https://github.com/ahamed-x/endoflife.date/blob/HEAD/products/opentofu.md) |
| 2024-04 | HCP Terraform rebrand (ex-Terraform Cloud) | [secondary](https://www.env0.com/blog/terraform-cloud) |
| 2024-04-30 | OpenTofu 1.7: client-side state encryption | [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/) |
| 2025-02 | IBM completes $6.4B HashiCorp acquisition | [secondary](https://www.env0.com/blog/terraform-cloud) |
| 2025-05-06 | Argo CD v3.0 GA (OCI-native, SSA) | [secondary](https://github.com/ncr38/argo-cd/blob/HEAD/docs/developer-guide/release-process-and-cadence.md) |
| 2025-06-24 | OpenTofu 1.10: OCI registry, S3 native locking | [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/) |
| 2025-07-21 | ansible-core 2.19 released | [secondary](https://endoflife.date/ansible-core) |
| 2025-08 | HCP Vault Dedicated "Starter" tier discontinued; renamed Development/Essentials/Standard | [secondary](https://infisical.com/blog/hashicorp-vault-pricing) |
| 2025-11-03 | ansible-core 2.20 released | [secondary](https://endoflife.date/ansible-core) |
| 2025-12-09 | OpenTofu 1.11: ephemeral values | [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/) |
| 2026-03-12 | awx-operator Helm chart v3.2.1 | [official](https://forum.ansible.com/t/release-announcement-awx-operator-helm-chart-v3-2-1/45518) |
| 2026-03-31 | HCP Terraform legacy free plan ends; enhanced free tier (500 resources) | [secondary](https://github.com/robhunter/agentdeals/issues/68) |
| 2026-05-14 | OpenTofu 1.12: dynamic prevent_destroy | [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/) |
| 2026-05-31 | ansible-core 2.21 released; 2.18 EOL | [secondary](https://endoflife.date/ansible-core) |
| 2026-06-02 | Ansible community package 14 (core 2.21) | [secondary](https://github.com/endoflife-date/endoflife.date/blob/HEAD/products/ansible.md) |
| 2026-07-28 | cloud-init 26.2 released | [official](https://launchpad.net/cloud-init/+download) |
| 2026-08-04 | Argo CD v3.5 GA | [secondary](https://github.com/ppapapetrou76/argo-cd/blob/HEAD/docs/developer-guide/release-process-and-cadence.md) |
| 2026-08-31 | Flux v2.9.5 | [secondary](https://devtoollab.com/blog/best-gitops-tools) |
| 2026-09-14 | Argo CD v3.5.3 (current patch) | [secondary](http://eosl.date/eol/product/argo-cd/) |

### M2. HCP Terraform tier feature matrix (2026)

| Capability | Free (500 res) | Essentials $0.10 | Standard $0.47 | Premium $0.99 | TFE (custom) |
|---|---|---|---|---|---|
| Remote state, VCS workflows | Yes | Yes | Yes | Yes | Yes |
| Unlimited users/workspaces | Yes | Yes | Yes | Yes | Yes |
| Private registry | Limited | 10 modules | Unlimited | Unlimited | Unlimited |
| No-code provisioning | — | — | Yes | Yes | Yes |
| Audit logging | — | — | Yes | Yes | Yes |
| Sentinel / OPA policies | — | — | Yes (Sentinel) | Yes | Yes |
| Module revocation, Waypoint actions | — | — | — | Yes | Yes |
| SAML SSO | — | — | — | — | Yes |
| Air-gapped / self-hosted | — | — | — | — | Yes |

Compiled from [secondary](https://www.env0.com/blog/terraform-cloud) and [secondary](https://github.com/bobikenobi12/bb-thesis-2026/blob/HEAD/spec/mvp/competitors/terraform-cloud.md); verify against the live pricing page — tiers changed twice in 2025–2026.

### M3. AWX vs Ansible Automation Platform (2026)

| Dimension | AWX (upstream) | AAP (Red Hat) |
|---|---|---|
| License/cost | Free, open source | Subscription (per managed node; aggregators cite $5k–$14k/yr ladders [secondary]) |
| Deployment | AWX Operator on K8s (self-managed) | Operator or OpenShift, supported topologies |
| Content | Galaxy community collections | Certified collections + automation hub |
| EDA | Community EDA components | Supported EDA server (1.2.12 in 2.6-era set [secondary]) |
| Support | Community | Red Hat SLA |
| Reference versions | Operator 2.19.1 / AWX 24.6.1 [official] | 2.6-era bundle: AWX 26.0.0 [secondary] |

