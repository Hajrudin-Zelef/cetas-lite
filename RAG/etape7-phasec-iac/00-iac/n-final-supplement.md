---
id: etape7-phasec-iac/00-iac/n-final-supplement
title: "N. Final supplement"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: []
dates: ["2026-03-12", "2026-03-31", "2026-06-18", "2026-07-28", "2026-08-19", "2026-08-27", "2026-08-31", "2026-09", "2026-09-08", "2026-09-14", "2026-09-22", "2026-11-03", "2026-11-30", "2027-05-31", "2027-11-30"]
keywords: ["apache", "aws", "compute", "packaging", "pricing", "research"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [664, 740]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: 0402411565ef2cedc805ebaedda55d738112e912cf4cae244a179891d3d8cca1
---

# N. Final supplement

## N. Final supplement

### N1. ansible-core changelog notes (2.19 → 2.21 era)

- The 2.19–2.21 releases continued the established 6-month cadence (May/November); 2.21's headline platform change is the managed-node Python floor at 3.9 and PowerShell 7 support on Windows targets [secondary](https://endoflife.date/ansible-core).
- Operators should note the EOL cliff: 2.19 goes EOL 2026-11-30, leaving 2.20 (EOL 2027-05-31) and 2.21 (EOL 2027-11-30) as the supported lines through 2027 [secondary](https://endoflife.date/ansible-core).
- The community package numbering (11→14) tracks core directly: community 14 = core 2.21, 13 = 2.20, 12 = 2.19, 11 = 2.18 [secondary](https://github.com/endoflife-date/endoflife.date/blob/HEAD/products/ansible.md).
- Detailed per-release changelogs live in the ansible/ansible changelog and the `ansible-core` porting guides; behavioral deltas between 2.19 and 2.21 were not itemized in this pass — consult the porting guides before upgrading fleets [unverified].

### N2. HCL language features (Terraform/OpenTofu shared surface)

- `moved` blocks (refactoring without destroy/recreate), `import` blocks (declarative import), and `removed` blocks (declarative decommissioning) are the modern refactoring primitives shared by both CLIs [unverified — documented language features].
- `terraform_remote_state` / `tftest` patterns for cross-stack references; OpenTofu's `tofu test` framework (`.tftest.hcl`) brings unit-style assertions to modules on both sides of the fork [unverified].
- Provider-defined functions and the `terraform_data` resource round out the 1.x-era language surface cited in September 2026 comparisons (Terraform 1.9–1.15) [secondary](https://shattered.io/terraform-vs-pulumi-vs-opentofu-2026/).

### N3. Costed HCP Terraform examples (2026 RUM math)

Using third-party rule-of-thumb figures (Standard $0.47/res/mo) [secondary](https://github.com/bobikenobi12/bb-thesis-2026/blob/HEAD/spec/mvp/competitors/terraform-cloud.md):

| Footprint | Resources | Essentials/mo | Standard/mo | Premium/mo |
|---|---|---|---|---|
| Lab / tutorial | 100 | $10 | $47 | $99 |
| Small prod (1 EKS + net + IAM) | 1,000 | $100 | $470 | $990 |
| Mid org, multi-env | 10,000 | $1,000 | $4,700 | $9,900 |
| Large estate | 50,000 | $5,000 | $23,500 | $49,500 |

- The $500 new-account credit covers ~10 months at 500 resources on Essentials, ~2.1 months on Standard, ~1 month on Premium [secondary](https://www.env0.com/blog/terraform-cloud-pricing-a-complete-guide).
- Free tier (500 resources, unlimited users) absorbs the lab row entirely — the 2026-03-31 restructuring mainly hit teams in the 500–2,000 resource band [secondary](https://github.com/robhunter/agentdeals/issues/68).

### N4. Ansible inventory and variables (platform angle)

- Inventory sources: static INI/YAML, dynamic inventory plugins (aws_ec2, azure_rm, gcp_compute, constructed, advanced_host_list), and AWX/AAP synced inventories (cloud, VM, custom scripts) [unverified — documented plugin set].
- Variable precedence (simplified): role defaults < inventory < play vars < host_vars/group_vars < extra vars (`-e` always wins) — the extra-vars-wins rule is what makes CI-driven overrides safe [unverified — documented behavior].
- Vault-encrypted variables (`ansible-vault encrypt_string`) are the legacy secrets pattern; 2026 best practice moves secret *material* to Vault/ESO and keeps only references in Ansible vars [unverified — established practice].

### N5. Pulumi ESC and Crossplane compositions (brief)

- **Pulumi ESC** (Environments, Secrets, Configuration): centralized config/secrets store feeding stacks and non-Pulumi consumers (CI, kubectl) — Pulumi's answer to Vault-lite for IaC-native teams [unverified — 2026 status not re-checked].
- **Crossplane compositions**: XRDs (composite resource definitions) + Compositions let platform teams expose simplified self-service APIs ("a database") backed by composed managed resources — the Kubernetes-native platform-engineering pattern [unverified].
- Both are alternatives to writing raw HCL/YAML per environment; neither was depth-verified in this pass (gaps G4/G5).

### N6. What to verify before downstream use (recap of highest-risk claims)

1. AAP official current version and per-node price book (aggregator $5k/$10k/$14k figures are `[secondary]`).
2. HCP Terraform free-tier current state (credit vs enhanced free tier — changed twice in 2026).
3. Exact latest Terraform 1.x patch; Packer latest version.
4. Kyverno/Crossplane/ESO/SOPS/Sealed Secrets/driftctl 2026 versions.
5. Spacelift/env0/Scalr current packaging and list prices.
6. OpenTofu 1.12.x latest patch (1.12.6 seen 2026-08-19; may be newer by 2026-09-22).

*End of Step 7 Phase C — IaC & Platform Automation. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*

---

## O. Version pin quick-reference (2026-09-22)

| Component | Current / latest seen | EOL / note | Provenance |
|---|---|---|---|
| ansible-core | 2.21.4 (2026-09-08) | 2.19 EOL 2026-11-30 | [secondary](https://endoflife.date/ansible-core) |
| Ansible community pkg | 14.4.0 (2026-09-08) | 13 EOL 2026-06-18 | [secondary](https://github.com/endoflife-date/endoflife.date/blob/HEAD/products/ansible.md) |
| AWX / operator | 24.6.1 / 2.19.1 | Helm chart 3.2.1 (2026-03-12) | [official](https://github.com/ansible/awx-operator/releases) |
| AAP bundle | 2.6-era (AWX 26.0.0, EDA 1.2.12) | Verify vs Red Hat notes | [secondary](https://github.com/fitbeard/automation-platform/blob/HEAD/README.md) |
| OpenTofu | 1.12.6 (2026-08-19) | 1.11 EOL 2026-08-19 | [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/) |
| Terraform | 1.9–1.15 line (BSL 1.1) | IBM-owned since 2025-02 | [secondary](https://shattered.io/terraform-vs-pulumi-vs-opentofu-2026/) |
| Pulumi | 3.x generation | Apache 2.0 | [secondary](https://shattered.io/terraform-vs-pulumi-vs-opentofu-2026/) |
| Packer | BUSL-1.1 | Latest ver. not pinned | [official](https://github.com/hashicorp/packer/blob/HEAD/README.md) |
| cloud-init | 26.2 (2026-07-28) | CalVer YY.M | [official](https://launchpad.net/cloud-init/+download) |
| Argo CD | 3.5.3 (2026-09-14) | v3.6 GA 2026-11-03 | [secondary](http://eosl.date/eol/product/argo-cd/) |
| Flux | 2.9.5 (2026-08-31) | ControlPlane-maintained | [secondary](https://devtoollab.com/blog/best-gitops-tools) |
| Argo Rollouts | 1.10.0 (2026-08-27) | — | [secondary](https://devtoollab.com/blog/best-gitops-tools) |
| HCP Terraform | RUM: $0.10/$0.47/$0.99 | Free: 500 resources | [secondary](https://www.env0.com/blog/terraform-cloud) |
| HCP Vault Dedicated | $0.62–$9.41/hr + $72.92/client/mo | Starter EOL 2025-08 | [secondary](https://infisical.com/blog/hashicorp-vault-pricing) |

---

## P. Cross-references and usage notes

