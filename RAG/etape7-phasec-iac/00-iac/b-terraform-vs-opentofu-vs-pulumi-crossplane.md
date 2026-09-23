---
id: etape7-phasec-iac/00-iac/b-terraform-vs-opentofu-vs-pulumi-crossplane
title: "B. Terraform vs OpenTofu vs Pulumi (+ Crossplane)"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2023-08", "2023-08-10", "2024-04", "2024-07-29", "2025-01-09", "2025-02", "2025-06-24", "2025-12-09", "2026-03-31", "2026-05-11", "2026-05-14", "2026-06", "2026-08-19", "2026-09", "2027-02-01"]
keywords: ["acquisition", "apache", "aws", "cost", "governance", "license", "memory", "open source", "pricing", "research"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [78, 162]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: c2bfc662e43c2b9633f135392ee978aa291d3930c7ad12c653b5b09372846dbb
---

# B. Terraform vs OpenTofu vs Pulumi (+ Crossplane)

## B. Terraform vs OpenTofu vs Pulumi (+ Crossplane)

### B1. Terraform under HashiCorp/IBM (BSL era)

- HashiCorp moved Terraform (and Packer, Vault, Consul, Nomad, Boundary, Waypoint) from MPL 2.0 to **BSL 1.1** on 2023-08-10 for all future releases; the BSL bars hosted/embedded competitive use, converting to full FOSS after the change date [official](https://www.hashicorp.com/en/blog/hashicorp-adopts-business-source-license).
- Providers, SDKs and CDKTF stayed open source (MPL 2.0) [secondary](https://community.ops.io/eriklz/tidy-cloud-aws-hashicorp-goes-bsl-5222).
- IBM's **$6.4B** acquisition of HashiCorp closed in **February 2025** [secondary](https://www.env0.com/blog/terraform-cloud).
- Current-generation Terraform cited in September 2026 comparisons: **1.9–1.15** under BSL 1.1 [secondary](https://shattered.io/terraform-vs-pulumi-vs-opentofu-2026/). Exact latest 1.x patch not re-verified in this pass.
- Terraform remains the default among existing users due to the provider ecosystem; the license change is the main driver of OpenTofu evaluation [secondary](https://github.com/nilesuan/pdlc/blob/HEAD/techstacks/09-iac.md).

### B2. OpenTofu (Linux Foundation fork)

- Fork of Terraform 1.5.x announced August 2023, renamed OpenTofu, hosted by the Linux Foundation under **MPL 2.0** [secondary](https://github.com/nilesuan/pdlc/blob/HEAD/techstacks/09-iac.md).
- Release train (2026 status) [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/):

| Release | Released | Latest patch (2026-07/08) | Support status |
|---|---|---|---|
| 1.12 | 2026-05-14 | 1.12.6 (2026-08-19) | Supported until ~2027-02-01 |
| 1.11 | 2025-12-09 | 1.11.14 (2026-08-19) | EOL reached 2026-08-19 |
| 1.10 | 2025-06-24 | 1.10.10 (2026-05-11) | Supported (patch line) |
| 1.9 | 2025-01-09 | 1.9.4 | EOL reached 2026-05-14 |
| 1.8 | 2024-07-29 | 1.8.11 | EOL reached 2025-12-09 |

- Support policy: patches for the three most recent releases (same model as Terraform) [secondary](https://github.com/ahamed-x/endoflife.date/blob/HEAD/products/opentofu.md).
- OpenTofu-exclusive features shipped by 2026 (Terraform's open binary lacks equivalents) [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/):
  - **Client-side state encryption** (1.7): AES-GCM encryption of state/plan files; keys via passphrase, AWS KMS, GCP KMS, or OpenBao. `enforced = true` refuses to touch unencrypted state without an explicit migration fallback [secondary](https://github.com/christosgalano/christosgalano.github.io/blob/HEAD/_drafts/opentofu-vs-terraform-2026.md).
  - **Early variable evaluation** (1.8): variables/locals usable in `backend` and `module source` blocks.
  - **Provider iteration with `for_each`** (1.9): multi-region/account without duplicated provider blocks.
  - **`-exclude` flag** (1.9): inverse of `-target`.
  - **OCI registry support** (1.10): providers/modules from OCI container registries.
  - **S3 state locking without DynamoDB** (1.10): native S3 conditional writes.
  - **Ephemeral values** (1.11): in-memory-only secrets never persisted to state.
  - **Dynamic `prevent_destroy`** (1.12): lifecycle guards wired to variables (verified behavior on v1.12.4 in a community test) [secondary](https://github.com/christosgalano/christosgalano.github.io/blob/HEAD/_drafts/opentofu-vs-terraform-2026.md).
- Adoption drivers cited: open governance, CI/cloud vendors shipping Terraform-compatible tooling without BSL constraints, regulated shops needing state encryption [secondary](https://github.com/nilesuan/pdlc/blob/HEAD/techstacks/09-iac.md).
- **Conflict flagged:** eosl.date (crawled Sept 2026) lists 1.10.10 as "the latest" and recommends upgrading *to* 1.10.10 — this contradicts the 1.12.x line documented above and appears stale; prefer the 1.12.6 figure from release-tracker sources [secondary](http://eosl.date/eol/product/opentofu/).

### B3. HCP Terraform pricing (2026)

- Terraform Cloud → **HCP Terraform** rebrand: April 2024 [secondary](https://www.env0.com/blog/terraform-cloud).
- Pricing model: **RUM (Resources Under Management)** — billed on peak hourly managed-resource count across state files [secondary](https://github.com/bobikenobi12/bb-thesis-2026/blob/HEAD/spec/mvp/competitors/terraform-cloud.md).
- Tiers as of 2026 (all per managed resource / month) [secondary](https://www.env0.com/blog/terraform-cloud):
  - **Free** — up to **500 managed resources**, unlimited users, VCS integration, remote state. Note: the *legacy* free plan ended **2026-03-31**; existing orgs were moved to an enhanced free tier [secondary](https://github.com/robhunter/agentdeals/issues/68).
  - **Essentials — $0.10**/resource/month: remote state, VCS, projects, secure vars, private registry (10 modules).
  - **Standard — $0.47**/resource/month: + unlimited modules, no-code provisioning, Waypoint templates, team notifications, audit logs.
  - **Premium — $0.99**/resource/month: + module revocation, Waypoint actions.
  - **Terraform Enterprise (self-managed)** — custom/sales quote; no resource limits, audit logging, SAML SSO, air-gapped deployment [secondary](https://www.env0.com/blog/terraform-cloud-pricing-a-complete-guide).
- Rule-of-thumb math (third-party): ~1,000 resources ≈ **$470/mo** on Standard; 10,000 ≈ **$4,700/mo**; resource counts often run 30–50% above naive estimates (every IAM policy, SG rule, S3 lifecycle config counts) [secondary](https://github.com/bobikenobi12/bb-thesis-2026/blob/HEAD/spec/mvp/competitors/terraform-cloud.md).
- New HCP accounts receive a **$500 IBM/HashiCorp cloud credit** [secondary](https://github.com/bobikenobi12/bb-thesis-2026/blob/HEAD/spec/mvp/competitors/terraform-cloud.md).
- Governance: Sentinel (HashiCorp's proprietary PaC language) plus OPA support via Run Tasks on HCP Terraform [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e).
- **Conflict/ambiguity flagged:** one source (env0, earlier 2026) describes the free tier being replaced by a **$500 trial credit** consumed at tier rates (500 resources on Essentials ≈ 10 months of credit; on Standard ≈ 2.1 months; on Premium ≈ 1 month), while later sources (HashiCorp blog via agentdeals, June 2026) describe an **enhanced usage-based free tier** (500 resources, unlimited users) with legacy users auto-migrated on 2026-03-31. Both may be true at different dates; verify against the current HCP Terraform pricing page before quoting [secondary](https://www.env0.com/blog/terraform-cloud-pricing-a-complete-guide) vs [secondary](https://github.com/robhunter/agentdeals/issues/68).

### B4. Pulumi

- Positioning: IaC in general-purpose languages (TypeScript, Python, Go, C#, Java, YAML); Apache 2.0 license [secondary](https://shattered.io/terraform-vs-pulumi-vs-opentofu-2026/).
- Pulumi **3.x** is the current generation cited in September 2026 comparisons [secondary](https://shattered.io/terraform-vs-pulumi-vs-opentofu-2026/).
- Pulumi Cloud pricing (2026, via Spacelift's pricing survey and community comparisons) [secondary](https://spacelift.io/blog/pulumi-pricing):
  - **Individual** — free; automatic state management, unlimited updates, secret management; 500 deployment minutes/month free (Spacelift's figure) [secondary](https://spacelift.io/blog/pulumi-pricing).
  - **Team** — **$0.37/resource/month + $0.50/secret/month**; 150k free credits/month (≈ 200 resources free); up to 10 users, concurrency 5 [secondary](https://dev.to/mechcloud/iac-tools-comparison-2026-mechcloud-terraform-and-pulumi-325c).
  - **Enterprise** — **$1.10/resource/month + $0.75/secret/month**; unlimited users, RBAC, SAML/SSO; AWS Marketplace listing cited at **$32,850/yr** [secondary](https://spacelift.io/blog/pulumi-pricing).
  - **Business Critical** — self-hosting, org-wide/compliance/remediation policies, audit-log export, 24×7 support; Azure Marketplace listing cited at **$50,000/yr** [secondary](https://spacelift.io/blog/pulumi-pricing).
  - **Pulumi Deployments** (managed execution): **$0.01/deployment-minute** extra [secondary](https://dev.to/mechcloud/iac-tools-comparison-2026-mechcloud-terraform-and-pulumi-325c).
- Third-party budget guidance (Vendr transaction data, 2026): small teams $6k–$18k/yr; mid-size Enterprise $50k–$100k/yr; large Enterprise $120k–$300k/yr; Business Critical $150k–$500k+/yr [secondary](https://www.vendr.com/marketplace/pulumi).
- Note: beneath the language front-end, Pulumi still uses a state-file architecture mapping code to cloud resources [secondary](https://dev.to/mechcloud/iac-tools-comparison-2026-mechcloud-terraform-and-pulumi-325c).

### B5. Crossplane

- Crossplane extends Kubernetes with CRDs to provision cloud infrastructure from the K8s API; positioned as the "control-plane" alternative to CLI-driven IaC [secondary](https://dev.to/mechcloud_academy/iac-tool-pricing-comparison-terraform-crossplane-and-pulumi-3oe5).
- Crossplane itself is free/open (CNCF project); cost accrues via the managed control planes or operators running it.
- **Gap:** exact 2026 release line (v1.x vs v2) and CNCF maturity status were not re-verified in this research pass — flagged `[unverified]`, verify against crossplane.io before quoting.

### B6. Decision matrix (IaC engine choice, 2026)

| Dimension | Terraform (HashiCorp/IBM) | OpenTofu (Linux Foundation) | Pulumi | Crossplane |
|---|---|---|---|---|
| License | BSL 1.1 | MPL 2.0 | Apache 2.0 | Apache 2.0 (CNCF) |
| Language | HCL | HCL (drop-in for ≤1.5-era) | TS/Python/Go/C#/Java/YAML | YAML via K8s CRDs |
| State encryption (client-side) | No (backend-side only) | Yes (AES-GCM, v1.7+) | Backend-side | etcd/K8s secrets |
| Provider ecosystem | Largest | Reuses Terraform providers | Large (incl. Crosswalk) | Provider-based |
| Managed platform | HCP Terraform ($0.10–$0.99/res/mo) | Any TACO (Spacelift/env0/…) | Pulumi Cloud (from free) | Self-operated |
| Best fit | Incumbent estates, HCP stack | Regulated/open-governance, CI vendors | Developer-centric teams | K8s-native platform teams |

- The 2026 market has genuinely diverged: each ships features the others lack, so the choice is now architectural (license, state handling, control plane) rather than purely ecosystem-driven [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).

---

