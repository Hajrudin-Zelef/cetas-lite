---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/part-4
title: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries (part 4)"
domain: front-matter
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [96, 107]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: f2d59a8cfbb9adf738d3aaa22d9290f7c234e3aeae553e943f5ad77ad2ac4e89
---

# Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries (part 4)

- **Juniper Apstra Terraform provider** (official): configures Apstra intent-based DC fabrics via the Apstra Go SDK APIs; pushes HCL-defined intent through Apstra to multivendor devices (Juniper, Cisco, Dell, Arista); supports Day 0–Day 2+ provisioning, GitOps/CI-CD patterns [official — https://Www.juniper.net/content/dam/www/assets/solution-briefs/us/en/network-automation/configuring-apstra-through-terraform.pdf].
- **NetBox provider `e-breuninger/netbox` v4.1.0**: manages physical/logical network inventory (devices, IPs, prefixes, VLANs, cables) from Terraform — "documentation-first" pattern where NetBox is the queryable source of truth written by Terraform; used with Terraform Cloud VCS-driven workflows + GitHub Actions CI [secondary — https://github.com/simonpainter/www.simonpainter.com/blob/HEAD/blog/netbox-terraform.md; https://github.com/xiiisins/homelab/blob/HEAD/docs/services/netbox.md].
- **F5 `F5Networks/terraform-provider-f5os`**: v1.12.0 introduced breaking Go 1.25 minimum, `f5os_lag` LACP/STATIC modes, 75%+ coverage enforcement, govulncheck; handles F5OS device config as code [official — https://github.com/F5Networks/terraform-provider-f5os/releases].
- **Palo Alto `PaloAltoNetworks/panos`** (v1.11.0 in community tutorials): firewall interfaces, zones, address objects, security policies via API-key auth [secondary — https://www.packetswitch.co.uk/palo-alto-automation-with-terraform/].
- **Infoblox Terraform provider v2.5/v2.6** (2024): NIOS DDI resources (networks, A records, DNS zones), extensible-attribute filtering, next-available-network allocation; enables IPAM-as-code (IP allocated at apply time, DNS records created/removed with resources) [secondary/vendor-reported — https://ddi.mohflo.net/index.php/2024/06/11/whats-new-in-infoblox-terraform-provider-v2-5-guiding-to-infoblox-terraform-provider-v2-5-new-features/; http://www.infoblox.com/resources/solution-notes/optimize-your-hybrid-multi-cloud-infrastructure-with-infrahub-plug-in-for-terraform].
- Registry-notable vendor providers (Terraform Registry): **CiscoDevNet/aci** (Cisco ACI), **Juniper/mist** (Mist cloud), **fortinetdev/fortimanager** (FortiManager), **Cisco Meraki** provider — commonly combined under one Terraform layer for LAN/SD-WAN/DC [secondary — https://github.com/simonpainter/www.simonpainter.com/blob/HEAD/blog/bringing-it-all-together.md].
- Community reference architecture: one Terraform layer + multiple vendor providers, version-controlled in Git, approved PRs auto-applying to environments — network infra "always in a known consistent state" [secondary — simonpainter blog].
- **State management for networks:** remote state backends (Terraform Cloud, S3 with native locking in OpenTofu 1.10+), per-environment state separation, state encryption (OpenTofu 1.7+) for credentials in state; drift detection via `tofu plan -refresh-only` [secondary — OpenTofu docs/comparisons].
- **GitOps workflows:** GitHub Actions CI running `terraform fmt + init + validate`, tflint, checkov, terraform-docs; demo-mode `null_resource` with triggers for safe planning without touching devices; PR-driven apply pipelines [secondary — https://github.com/joycemwangi/hybrid-cloud-infrastructure-as-code-with-terraform].

---

