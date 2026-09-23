---
id: etape7-phasec-iac/00-iac/l-deep-dives-supplement
title: "L. Deep dives (supplement)"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "cost", "distribution", "governance", "memory", "open source", "packaging", "pricing"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [431, 477]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: 96fa5b4be8bb65a38b8b639d9088875f75cb1c691206e936a3c266f2bc6bde0a
---

# L. Deep dives (supplement)

## L. Deep dives (supplement)

### L1. Ansible toolchain beyond the core

- **ansible-lint**: the official linter for playbooks/roles; rules enforce naming, idempotency and FQCN (fully-qualified collection names) usage. In 2026-era content, lint-clean playbooks are the baseline for automation hub certification [unverified — version not re-checked in this pass].
- **Molecule**: role testing framework (lint → create → converge → verify with testinfra/ansible); still the reference for testing Ansible roles in CI before publishing to Galaxy or automation hub [unverified].
- **ansible-navigator**: TUI/CLI for running playbooks inside Execution Environments locally — the dev-loop companion to EE-based production runs [unverified].
- **Support lifecycle distinction**: ansible-core community support follows the 6-month release / ~12-month EOL windows in the table (A1); Red Hat AAP subscriptions add multi-year support windows for the same content via certified collections — one reason enterprises pay for AAP rather than running AWX [secondary](https://www.trustradius.com/compare-products/jfrog-artifactory-vs-red-hat-ansible-automation-platform).
- **Notable collections (platform angle)**: `ansible.builtin` (core), `community.general` (the broad utility collection), `amazon.aws` / `community.aws`, `azure.azcollection`, `google.cloud`, `community.docker`, `kubernetes.core` (k8s module coverage used with Argo CD-adjacent workflows), `community.crypto`, `ansible.posix`. Exact 2026 collection versions were not pinned in this pass — verify on Galaxy before quoting [secondary](https://github.com/endoflife-date/endoflife.date/blob/HEAD/products/ansible.md).
- **Idempotency contract**: Ansible modules are expected to be idempotent (report `changed=false` on converged state); check mode (`--check`) and diff mode (`--diff`) are the dry-run primitives used in CI gates [unverified — long-standing documented behavior].

### L2. Terraform/OpenTofu operational deep dive

- **Migration Terraform → OpenTofu**: the documented path is `tofu init` against existing state (HCL compatibility for ≤1.5-era language); provider ecosystem is shared via the OpenTofu registry plus OCI-registry distribution (1.10+) [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).
- **State encryption operations (OpenTofu 1.7+)**: configured in the `terraform { encryption { … } }` block; methods include PBKDF2 passphrase, AWS KMS, GCP KMS, OpenBao; `enforced = true` blocks writes to unencrypted state, with an explicit `fallback` stanza for migration windows [secondary](https://github.com/christosgalano/christosgalano.github.io/blob/HEAD/_drafts/opentofu-vs-terraform-2026.md).
- **Ephemeral values (1.11+)**: `ephemeral` resources/variables exist only in memory during a run — designed for short-lived credentials (e.g. Vault-issued DB passwords) that must never land in state [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).
- **HCP Terraform vs Terraform Enterprise (2026)**: HCP is SaaS-only (no self-hosting); TFE is the self-managed offering with air-gapped deployment, custom pricing, no resource limits, audit logging and SAML SSO [secondary](https://www.env0.com/blog/terraform-cloud-pricing-a-complete-guide).
- **Sentinel vs OPA on HCP**: Sentinel is HashiCorp-proprietary and tier-gated; OPA is supported as an alternative via Run Tasks for teams standardizing on open policy tooling [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e).
- **RUM billing mechanics**: billed hourly on *peak* managed-resource count; every IAM policy, SG rule and S3 lifecycle configuration counts as a resource — counts typically land 30–50% above naive estimates [secondary](https://github.com/bobikenobi12/bb-thesis-2026/blob/HEAD/spec/mvp/competitors/terraform-cloud.md).

### L3. Terragrunt and Atlantis (the missing middle layer)

- **Terragrunt** (Gruntwork): thin wrapper adding DRY multi-module orchestration, remote-state bootstrapping and dependency ordering on top of Terraform/OpenTofu; the standard answer to "Terraform at monorepo scale" before adopting a TACO. [unverified — 2026 version/support status not re-checked in this pass]
- **Atlantis** (open source, now under Digger's stewardship per community reports): Terraform pull-request automation — `atlantis plan/apply` as PR comments, with locking and policy checks. [unverified — 2026 status not re-checked]
- Positioning vs TACO: Terragrunt/Atlantis are self-hosted building blocks; Spacelift/env0/Scalr are managed platforms absorbing the same concerns (runs, policies, drift, cost estimation) with commercial support [secondary](https://www.env0.com/blog/terraform-cloud-pricing-a-complete-guide).

### L4. TACO pricing signals (2026)

- Spacelift's own 2026 Pulumi pricing survey confirms the per-resource meter is now the industry norm for IaC SaaS (Pulumi Team $0.37/res/mo; HCP Terraform Essentials $0.10 → Premium $0.99) [secondary](https://spacelift.io/blog/pulumi-pricing).
- Spacelift positions itself on workflow-based pricing "not tied to resource count under management" as the differentiator vs HCP Terraform's RUM model [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e).
- Exact 2026 list prices for Spacelift, env0 and Scalr tiers were not captured — flagged as gap G7; all three publish pricing pages that change packaging frequently, so quote from primary sources only.

### L5. Packer operational notes

- Packer templates are HCL2 (`*.pkr.hcl`); a build combines **sources** (builders: amazon-ebs, azure-arm, googlecompute, vsphere-iso, qemu, docker…), **provisioners** (shell, ansible, powershell, file) and **post-processors** (manifest, compress, vagrant) [official](https://github.com/hashicorp/packer/blob/HEAD/README.md).
- The **ansible provisioner** is the standard bridge: Packer builds the image, Ansible configures it, cloud-init finishes per-instance identity at boot — the golden-image triangle referenced in C3 [official](https://github.com/hashicorp/packer/blob/HEAD/README.md).
- Plugin model: builders/provisioners ship as external plugins (post-1.7 architecture), installable via `packer init` from the plugin registry [official](https://github.com/hashicorp/packer/blob/HEAD/README.md).

### L6. cloud-init module and datasource reference

- Datasources (selection): EC2, Azure, GCE, OpenStack, CloudStack, MAAS, NoCloud, ConfigDrive, OVF, Vultr, Hetzner, OpenNebula, LXD, WSLEnv [official](https://docs.cloud-init.io/_/downloads/en/latest/pdf/).
- Frequently used `cloud-config` modules: `users`/`groups`, `ssh_authorized_keys`, `packages`, `runcmd`, `write_files`, `bootcmd`, `mounts`, `disk_setup`/`fs_setup`, `resolv_conf`, `ntp`/`timesyncd`, `ca_certs`, `ansible` (pull mode), `phone_home`, `final_message` [official](https://docs.cloud-init.io/_/downloads/en/latest/pdf/).
- Network config v1/v2 renderers: ENI, netplan, sysconfig, NetworkManager, networkd — 26.2 added route-metric support in the NetworkManager renderer [official](https://launchpad.net/cloud-init/+download).
- Debugging: `cloud-init status --wait`, `cloud-init analyze boot`, `/var/log/cloud-init*.log` [official](https://docs.cloud-init.io/_/downloads/en/latest/pdf/).

---

