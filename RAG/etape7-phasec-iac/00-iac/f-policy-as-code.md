---
id: etape7-phasec-iac/00-iac/f-policy-as-code
title: "F. Policy as code"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: regulation
actors: []
dates: []
keywords: ["agent", "cost", "governance", "packaging", "pricing"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [279, 325]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: a9f904b63bb9b3a4b824ef3d774b02938ddef8880130a09d661209613d314587
---

# F. Policy as code

## F. Policy as code

### F1. The 2026 policy landscape

- **Sentinel** (HashiCorp): proprietary PaC language for HCP Terraform/Terraform Enterprise plan enforcement; available on Standard/Premium tiers (audit logs on Standard+) [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e). Also supported in HCP Vault Dedicated Standard tier [secondary](https://infisical.com/blog/hashicorp-vault-pricing).
- **OPA (Open Policy Agent) / Gatekeeper**: the open alternative; HCP Terraform supports OPA via Run Tasks for teams preferring open tooling [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e). Gatekeeper enforces OPA policies as Kubernetes admission control.
- **Kyverno**: Kubernetes-native policy engine (policies as YAML, no Rego learning curve); CNCF project. Exact 2026 version/graduation status not re-verified in this pass — flagged `[unverified]`.
- Governance-platform framing (2026): best-in-class setups combine the IaC platform's native policy (Sentinel or OPA run tasks) with cluster admission policy (Kyverno/Gatekeeper) so both `terraform plan` and `kubectl apply` are gated [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e).

### F2. Policy enforcement points in a platform pipeline

1. **Plan time** (Terraform/OpenTofu/Pulumi): Sentinel/OPA run tasks or `conftest` (OPA) in CI fail the run before apply.
2. **Admission time** (Kubernetes): Gatekeeper/Kyverno validate or mutate manifests delivered by Argo CD/Flux.
3. **Runtime/audit time**: drift detection (see G3) + Vault Sentinel/control groups for secrets access.
- Cost note: on HCP Terraform, governance features are tier-gated — audit logs need Standard; Sentinel needs Standard/Premium [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e).

---

## G. State backends, TACO platforms, drift detection

### G1. State backends

- **HCP Terraform**: remote state with locking/versioning included in all tiers (Free: unlimited state history) [secondary](https://github.com/atnaszurc/iac-bootcamp/blob/HEAD/TF-300-advanced/TF-305-workspaces-remote-state/4-hcp-terraform-state/README.md).
- **S3 + native locking**: OpenTofu 1.10+ supports S3 state locking via native conditional writes — no DynamoDB table required [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).
- **Pulumi Cloud**: automatic state management; self-managed backends (S3/Azure Blob/GCS) available [secondary](https://dev.to/mechcloud/iac-tools-comparison-2026-mechcloud-terraform-and-pulumi-325c).
- Security baseline: encrypt state at rest (OpenTofu 1.7+ client-side AES-GCM, or backend SSE-KMS); state files routinely contain secrets [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).

### G2. TACO — Terraform Automation and Collaboration platforms

TACO vendors (the "Terraform Cloud alternatives" layer) relevant in 2026:

- **Spacelift**: policy-driven orchestration for Terraform/OpenTofu/Pulumi/Kubernetes; usage/workflow-based pricing rather than per-resource RUM [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e).
- **env0**: Terraform/OpenTofu automation with cost-visibility tooling; publishes the 2026 HCP Terraform pricing analyses cited in this file [secondary](https://www.env0.com/blog/terraform-cloud).
- **Scalr**: Terraform/OpenTofu remote operations, often cited in RUM-pricing escape analyses [secondary](https://www.env0.com/blog/terraform-cloud-pricing-a-complete-guide).
- **Terramate**: orchestration/change-detection layer on top of Terraform/OpenTofu stacks [secondary](https://www.env0.com/blog/terraform-cloud-pricing-a-complete-guide).
- **GitLab OpenTofu CI/CD components**: GitLab's component catalog ships OpenTofu full-pipeline components (v3.16.0 seen with OpenTofu 1.11.x) — `include: component: gitlab.com/components/opentofu/full-pipeline` [secondary](https://newreleases.io/project/gitlab/components/opentofu/release/3.16.0).
- **Gap:** current per-vendor list pricing (Spacelift/env0/Scalr 2026 tiers) was not captured in this pass — flagged; these vendors price by seats/workflows/minutes and change packaging frequently.

### G3. Drift detection

- HashiCorp's own drift-detection tooling for Terraform ("which engineers tweaked the settings") has been covered in the trade press [secondary](https://www.theregister.com/tag/hashicorp).
- Native `terraform plan`/`tofu plan` refresh remains the baseline drift signal; scheduled-plan TACO features (Spacelift drift detection, env0) productize it [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e).
- **driftctl** (Snyk-era open tool): status in 2026 not re-verified — flagged `[unverified]`.
- Kubernetes-side drift: Argo CD's sync status + diff is the de-facto drift detector for GitOps-managed clusters [secondary](https://devtoollab.com/blog/best-gitops-tools).

---

