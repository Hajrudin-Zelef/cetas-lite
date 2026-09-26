---
id: etape7-phasec-iac/00-iac/b5-crossplane
title: "B5. Crossplane"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2026-09"]
keywords: ["apache", "aws", "cost", "governance", "license", "pricing", "research"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [131, 162]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: 6f0aea54decb68630b6e859eea2d60198b84027be74eb8dd805733196a888fbd
---

# B5. Crossplane

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

