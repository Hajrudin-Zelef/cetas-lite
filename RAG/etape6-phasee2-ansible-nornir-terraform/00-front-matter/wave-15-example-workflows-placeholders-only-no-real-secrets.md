---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-15-example-workflows-placeholders-only-no-real-secrets
title: "Wave 15 — Example workflows (placeholders only, no real secrets)"
domain: front-matter
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "benchmark", "embedding", "license"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [270, 287]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: 2ed27c718c2178eadb4988689135ef906f18c988be2ceb3ece701d561c579e1e
---

# Wave 15 — Example workflows (placeholders only, no real secrets)

- **Ansible:** choose for declarative config management, operator-friendly YAML, AAP controller (RBAC/approvals/scheduling), large existing playbook investments, multi-vendor resource modules; watch collection deprecation notices (junipernetworks.junos, arista.eos `src`) [official changelogs; secondary].
- **Nornir:** choose for Python-native teams, custom multi-device logic, embedding in applications, fastest parallel ad-hoc operations; add NornFlow for YAML UX; accept smaller module library vs Ansible [secondary].
- **Terraform/OpenTofu:** choose for declarative lifecycle of API-driven platforms (Apstra, Meraki, ACI, Infoblox, NetBox objects, F5OS), GitOps with plan/apply gates, state-tracked inventory; less suited to CLI-only legacy devices [secondary].
- **pyATS/Genie:** choose for Cisco-centric testing/validation, parser library reuse, pre/post change diffs; complements (not replaces) config tools [official — Cisco DevNet].
- **Raw Python (Netmiko/Scrapli/NAPALM):** choose for libraries inside custom apps, one-off migrations, maximum control; wrap with Nornir for scale [secondary].
- **Hybrid reference architecture (2026 community):** NetBox/Infrahub (SoT) → Terraform (provision platforms/IPAM) → Ansible/Nornir (device config) → pyATS (validate) → EDA (react to events); Git as the audit spine [secondary — simonpainter; Red Hat; opsmill].
- **Scale notes:** Ansible — persistent connections + fact caching + serial batching; Nornir — thread tuning + connection reuse; OpenTofu — per-env state, targeted applies; all — lab validation before fleet rollout [secondary].
- **Skill matrix:** Ansible (YAML, low Python need) < Nornir (Python required) < raw libraries (Python advanced); Terraform (HCL, API mental model); pyATS (Python + Cisco CLIs) [secondary].
- **Licensing 2026:** Ansible/AAP (GPLv3+/Red Hat subscription), Nornir (Apache-2.0), OpenTofu (MPL-2.0), Terraform (BSL 1.1), Infrahub (AGPLv3), NetBox (Apache-2.0) — license posture matters for SaaS/embedded use [independent].
- **Risk register:** collection deprecations (Wave 8), provider registry lag on OpenTofu (6–12 weeks per one benchmark), AWX upstream release pause (use operator/community builds), pygnmi maintenance uncertainty [secondary].

---

## Wave 15 — Example workflows (placeholders only, no real secrets)

### 15a. Ansible — multi-vendor VLAN push (resource modules, check-first)

```yaml
