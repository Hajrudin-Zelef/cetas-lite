---
id: etape7-phasec-iac/00-iac/e-secrets-management
title: "E. Secrets management"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2025-08"]
keywords: ["apache", "aws", "governance", "license", "pricing"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [234, 278]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: 8cbb69b4f179f9033168642b9f102da93966cf966959df7ed575a46037c3826c
---

# E. Secrets management

## E. Secrets management

### E1. HashiCorp Vault (now IBM Vault)

- Vault is the reference secrets engine (dynamic secrets, encryption-as-a-service, PKI, identity-based auth); covered by the BSL 1.1 change for new releases [secondary](https://community.ops.io/eriklz/tidy-cloud-aws-hashicorp-goes-bsl-5222).
- Aggregators now list it as **"IBM Vault (formerly HashiCorp Vault)"** [secondary](https://www.trustradius.com/compare-products/hashicorp-vault-vs-oracle-database).
- **HCP Vault Dedicated** pricing, 2026 (via Infisical's pricing guide citing HashiCorp's official pricing page) [secondary](https://infisical.com/blog/hashicorp-vault-pricing):

| Cluster size | Development | Essentials | Standard |
|---|---|---|---|
| Extra Small | **$0.62/hr** (~$450/mo) | — | — |
| Small | — | **$1.58/hr** (~$1,152/mo) | **$1.84/hr** (~$1,345/mo) |
| Medium | — | **$3.16/hr** (~$2,307/mo) | **$3.69/hr** (~$2,694/mo) |
| Large | — | **$7.49/hr** (~$5,468/mo) | **$9.41/hr** (~$6,870/mo) |

- Plus **$72.92/month per client** (unique app/service/user authenticating) on Essentials and Standard tiers [secondary](https://infisical.com/blog/hashicorp-vault-pricing).
- Tier notes: Development = single node, max 25 clients, no HA, no SLA, Silver support w/o Sev-1; Essentials = 99.9% SLA, audit streaming, backup/restore; Standard = + performance replication, Sentinel policies, control groups, Gold support. The old "Starter" tier was discontinued **August 2025**; Standard/Plus naming replaced by Development/Essentials/Standard [secondary](https://infisical.com/blog/hashicorp-vault-pricing).
- **HCP Vault Secrets** (serverless secrets, per-secret metering) official flex pricing [official](https://www.hashicorp.com/en/pricing/consumption-table):
  - Standard edition: first 1–100 secrets **$0.0006944/hr/secret** (Silver) / $0.0007246 (Gold); 101–200 half rate; 201+ quarter rate.
  - Plus edition: first 1–5,999 secrets **$0.0013014/hr/secret** (Silver); volume breaks at 6k/20k/50k.
  - API calls: **$0.00001/API call/month** (both editions).
- Self-managed Vault Enterprise remains quote-based [secondary](https://www.trustradius.com/compare-products/hashicorp-vault-vs-oracle-database).
- Certification note: Vault Associate exam (003) is **$70.50**, online proctored, 1 hour, credential valid 2 years [official](https://github.com/hashicorp/dev-portal/blob/HEAD/src/content/certifications/exam-faqs/vault-associate-003.mdx).

### E2. Open-source and Kubernetes-native alternatives

- **OpenBao**: community fork of Vault (Linux Foundation) preserving MPL 2.0 governance after the BSL change; supports the state-encryption KMS role for OpenTofu [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).
- **SOPS** (Secrets OPerationS, Mozilla): encrypts YAML/JSON/ENV/TFvars files with age/PGP/KMS; the Git-native secrets pattern for IaC repos — encrypt the file, commit it, decrypt in CI. [unverified — not re-checked in this pass; include only as a known-pattern pointer]
- **External Secrets Operator (ESO)**: syncs secrets from external providers (Vault, AWS Secrets Manager, etc.) into Kubernetes Secrets; the standard answer to "no secrets in Git" for GitOps flows. [unverified — version not re-checked in this pass]
- **Sealed Secrets** (Bitnami): one-way encrypted Secrets for GitOps; simpler than ESO, no external provider needed. [unverified — version not re-checked in this pass]
- Pattern guidance: never store plaintext secrets in state-adjacent Git repos; prefer ESO/SOPS + Vault/OpenBao backends; OpenTofu ephemeral values (1.11+) and client-side state encryption (1.7+) close the state-file leak vector [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).

### E3. Secrets-tooling comparison

| Tool | Model | License (2026) | Best fit |
|---|---|---|---|
| HCP Vault Dedicated | Managed clusters | Commercial (BSL-era) | Enterprises wanting managed Vault |
| Vault self-managed | Self-hosted | BSL 1.1 (new releases) | Existing HashiCorp estates |
| OpenBao | Self-hosted | MPL 2.0 (LF) | Open-governance requirement |
| SOPS | File encryption in Git | Apache 2.0 [unverified] | Git-native IaC repos |
| External Secrets Operator | K8s sync from providers | Apache 2.0 [unverified] | GitOps secret injection |
| HCP Vault Secrets | Serverless per-secret | Commercial | Low-volume, no-cluster overhead |

---

