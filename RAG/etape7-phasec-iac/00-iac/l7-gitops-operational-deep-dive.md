---
id: etape7-phasec-iac/00-iac/l7-gitops-operational-deep-dive
title: "L7. GitOps operational deep dive"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2026-08-27"]
keywords: ["agents", "aws", "cost", "governance", "guardrails", "memory", "parameters", "pricing"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [478, 526]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: 054bc87ba73db6185cdd6b85569ce25916f0ab14813983050b91bca095143865
---

# L7. GitOps operational deep dive

### L7. GitOps operational deep dive

- **Argo CD core objects**: `Application` (source repo + destination cluster/namespace + sync policy), `ApplicationSet` (templated multi-cluster/multi-env application generation), `AppProject` (RBAC scoping) [unverified — long-standing documented model].
- **Sync mechanics**: automated vs manual sync; sync waves and resource hooks order apply/delete sequences; SSA (server-side apply) default in v3 reduces controller memory and resolves field-ownership conflicts [secondary](http://dev.to/saaro_net/gitops-with-argocd-2026-cluster-pause-predelete-hooks-and-the-future-of-kubernetes-deployments-5bce).
- **Flux controllers**: source-controller (git/OCI/helm repos), kustomize-controller, helm-controller, notification-controller, image-reflector/image-automation controllers [unverified — documented architecture].
- **Image automation (Flux advantage)**: image-reflector scans registries, image-automation writes new tags back to Git — registry-driven deploys without a CI step committing to Git; Argo CD needs the Image Updater companion for the equivalent [secondary](https://devtoollab.com/blog/best-gitops-tools).
- **Secrets in GitOps**: neither engine should see plaintext secrets — standard patterns are ESO (sync from Vault/cloud secret stores into cluster Secrets), Sealed Secrets (one-way encrypted, controller-decrypted), or SOPS-encrypted manifests decrypted by KSOPS/Flux SOPS integration [unverified — established patterns].
- **Progressive delivery**: Argo Rollouts v1.10.0 (2026-08-27) adds canary/blue-green/mirror analysis on top of Argo CD-managed workloads [secondary](https://devtoollab.com/blog/best-gitops-tools).

### L8. Secrets deep dive

- **Vault auth methods**: Kubernetes auth, AWS/GCP/Azure IAM auth, LDAP, OIDC/JWT, AppRole, TLS certs — identity-based access is the Vault-native alternative to long-lived tokens [unverified — documented model].
- **Vault storage**: integrated Raft storage is the standard HA backend for self-managed clusters; HCP Vault Dedicated abstracts this with managed upgrades/backups/monitoring [secondary](https://infisical.com/blog/hashicorp-vault-pricing).
- **ESO architecture**: `SecretStore`/`ClusterSecretStore` (provider connection: Vault, AWS Secrets Manager, GCP Secret Manager, Azure Key Vault…) + `ExternalSecret` (what to sync, refresh interval, target Secret template). Refresh-driven sync keeps cluster Secrets converged with the provider [unverified — documented model].
- **SOPS**: encrypts values inside YAML/JSON/ENV/TFvars with age or PGP data keys wrapped by KMS; `.sops.yaml` creation rules scope keys per path — the Git-native pattern for IaC repos that must stay committable [unverified].
- **Vault vs OpenBao decision**: choose OpenBao when MPL 2.0 governance is a hard requirement (post-BSL); choose Vault/HCP when the HashiCorp ecosystem (HCP Terraform, Boundary, Sentinel) is already standardized [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).

### L9. Policy-as-code deep dive

- **OPA/Rego**: general-purpose policy engine; `conftest` runs OPA policies against Terraform plans, Kubernetes manifests, Dockerfiles in CI — the plan-time gate for OpenTofu/Terraform pipelines [unverified — established tooling].
- **Gatekeeper**: Kubernetes admission webhook running OPA; `ConstraintTemplate` (Rego) + `Constraint` (parameters) model; audit mode reports violations without blocking [unverified].
- **Kyverno**: policies as Kubernetes YAML (validate/mutate/generate), no Rego required; lower adoption barrier for cluster admins; 2026 version not re-verified (gap G6) [unverified].
- **Sentinel**: HCP Terraform/Enterprise policy language (plan-time); HCP Vault Dedicated Standard adds Sentinel for secrets governance [secondary](https://infisical.com/blog/hashicorp-vault-pricing).
- Reference enforcement chain: `conftest`/`tflint` in CI → Sentinel/OPA run tasks on the TACO → Kyverno/Gatekeeper admission in-cluster → drift detection post-apply [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e).

### L10. Cost-control and FinOps touchpoints in IaC

- **Cost estimation in plans**: HCP Terraform historically included cost estimation; 2026 community reports note estimation features being restricted on lower tiers — verify current tier gating before relying on it [secondary](https://dev.to/abhishek_gupta_pinpo/hcp-terraforms-free-tier-is-gone-what-aws-teams-should-actually-do-next-3fg5).
- **env0 / Spacelift / Infracost**: third-party cost-visibility layers that comment plan cost deltas on PRs; Infracost-style PR comments are the de-facto FinOps gate for Terraform/OpenTofu [unverified — established pattern].
- **Pulumi policy packs** and **OpenTofu dynamic prevent_destroy (1.12)** give guardrails that prevent expensive accidents (e.g. deleting production databases) independent of billing tier [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).

### L11. Platform-engineering reference stack (2026 synthesis)

| Layer | Default pick (2026) | Alternative |
|---|---|---|
| Config management | Ansible (core 2.21 / community 14) | Salt/Puppet (out of scope) |
| Provisioning | Terraform 1.x (incumbent) / OpenTofu 1.12 (open) | Pulumi 3.x, Crossplane |
| TACO / runs | HCP Terraform / Spacelift / env0 | Atlantis, GitLab components |
| Image build | Packer (BUSL-1.1) | osbuild/Image Builder |
| First-boot | cloud-init 26.2 | distro agents |
| GitOps | Argo CD 3.5.x | Flux 2.9.x, Fleet |
| Secrets | Vault/OpenBao + ESO/SOPS | HCP Vault Secrets |
| Policy | OPA/Gatekeeper/Kyverno + Sentinel | conftest in CI |
| State backend | HCP / S3 (+native lock, tofu 1.10+) | Pulumi Cloud |

*End of supplement. All version/date claims above carry provenance tags; items marked [unverified] were not re-checked against primary sources in this pass and should be verified before downstream use.*

---

