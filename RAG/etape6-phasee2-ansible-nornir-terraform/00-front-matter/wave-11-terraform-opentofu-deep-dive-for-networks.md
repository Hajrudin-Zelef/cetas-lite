---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-11-terraform-opentofu-deep-dive-for-networks
title: "Wave 11 — Terraform/OpenTofu deep dive for networks"
domain: front-matter
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "benchmark", "license"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [202, 229]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: 2f4447c981cb62b665d9d7765c2ba6801bf5bfd966ec8ab7f8d1d032715ce517
---

# Wave 11 — Terraform/OpenTofu deep dive for networks

## Wave 11 — Terraform/OpenTofu deep dive for networks

- **State as source of truth:** `terraform.tfstate` records real-world resource bindings; remote backends (Terraform Cloud, S3+GCS+Azurerm, OpenTofu native S3 locking without DynamoDB) enable team collaboration and CI [secondary — OpenTofu comparisons].
- **State encryption (OpenTofu 1.7+):** encrypt state at rest with passphrase or KMS (AWS KMS, GCP KMS, OpenBao) — important because network provider state often contains credentials/secrets [secondary — https://scalr.com/learning-center/what-is-opentofu].
- **Ephemeral values (OpenTofu 1.11+):** mark sensitive inputs ephemeral so they never persist in state or plan files — designed for API tokens/keys in network providers [secondary].
- **Import workflows:** `import` blocks with `for_each` (OpenTofu 1.7+) bulk-import brownfield devices/VLANs/VRFs into management without manual `terraform import` per resource [secondary].
- **`removed` blocks:** declare intent to stop managing a resource; on next apply it's forgotten (not destroyed) — safe decommissioning path for decommissioned sites [secondary].
- **Provider `for_each` (OpenTofu 1.9+):** instantiate provider configurations per region/site dynamically — one module, N sites with different controller endpoints [secondary].
- **Module registry:** OCI-based registry (OpenTofu 1.10+) alongside classic registry; ~3,900 providers claimed (conflict C1, Wave 8) [secondary].
- **CI pipeline (real pattern):** GitHub Actions → `fmt` → `init` → `validate` → `tflint` → `checkov` (policy) → `terraform-docs` → `plan` artifact on PR → manual approval → `apply` on merge [secondary — https://github.com/joycemwangi/hybrid-cloud-infrastructure-as-code-with-terraform].
- **Safe planning without devices:** `terraform_data`/`null_resource` with triggers in "demo mode" lets teams validate module logic in CI without touching real controllers [secondary — joycemwangi repo].
- **Drift detection:** scheduled `plan -refresh-only` (or `tofu plan -refresh-only`) jobs surface out-of-band changes; alert when drift detected [secondary].
- **Workspaces vs directories:** per-environment state via workspaces or separate directories+backends; directories preferred for prod isolation (blast radius) [secondary].
- **Apstra provider workflow:** HCL declares blueprint intent (ASN, VNI, rack types) → `terraform apply` pushes to Apstra → Apstra renders device configs → devices stream telemetry back; Day-2 changes = HCL edits [official — Juniper solution brief].
- **NetBox-as-code workflow:** `e-breuninger/netbox` 4.1.0 manages devices/IPs/prefixes/VLANs/cables; Terraform Cloud VCS-driven runs; NetBox webhooks then feed Ansible/Nornir inventory — full "define once, consume everywhere" loop [secondary — simonpainter blog].
- **Homelab pattern (xiiisins):** NetBox + Terraform provider in a GitOps homelab; mirrors enterprise flow at small scale [secondary — https://github.com/xiiisins/homelab/blob/HEAD/docs/services/netbox.md].
- **IPAM-as-code (Infoblox):** next-available-network allocation at apply; DNS A records created/removed with the resource lifecycle; extensible attributes as tags/filters [secondary — mohflo blog].
- **Palo Alto caution:** community tutorial uses API-key auth in provider config and inline secrets — treat as pedagogical only; production must use Vault/env-injected credentials, never committed keys [secondary — packetswitch tutorial; flagged].
- **F5OS provider ops:** Go 1.25 minimum (v1.12.0+), LAG resources (LACP/STATIC), high test coverage; fits ADC-as-code pipelines [official — https://github.com/F5Networks/terraform-provider-f5os/releases].
- **Multi-provider stack (real):** one Terraform layer combining Cisco ACI + Juniper Mist + FortiManager + F5 BIG-IP + NetBox providers, version-controlled, PR-approved applies per environment [secondary — simonpainter "bringing it all together"].
- **Terraform vs OpenTofu 2026 decision factors:** license (BSL vs MPL-2.0), registry lag (providers trail 6–12 weeks on OpenTofu per one benchmark), HCP Terraform backend (Terraform-only), state encryption/ephemerals (OpenTofu-only), corporate stewardship (HashiCorp/IBM vs Linux Foundation/CNCF) [secondary — comparisons; **C4**: Terraform 1.9 "last MPL" claim unverified].
- **Migration evidence:** teams switching binary only (`opentofu = "1.12.6"` in mise.toml, HCL unchanged) report zero-config migration for network modules not using terraform-specific data sources [secondary — smana ADR].
- **Policy as code:** OPA/Conftest or checkov rules on plans (e.g., "no 0.0.0.0/0 in security policies", "all VLANs must have descriptions") before apply [secondary].

---

## Wave 12 — Python libraries deep dive

