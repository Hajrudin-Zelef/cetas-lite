---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-11-terraform-opentofu-deep-dive-for-networks
title: "Wave 11 — Terraform/OpenTofu deep dive for networks"
domain: front-matter
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["apache", "aws", "benchmark", "embedding", "license", "mcp", "model context protocol"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [202, 287]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: 40974c62cd0ebb7d28e601599e2bd550ef9b703397fbf1083f75aefd0afb060c
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

- **Paramiko/libssh:** Netmiko defaults to Paramiko; `netmiko[par4]` extra tracks Paramiko 4; scrapli offers `ssh2` transport (libssh2) for speed; ncclient 0.7.0 added libssh transport option alongside Paramiko [independent — libraries.io; ncclient releases].
- **Netmiko device types:** broad `cisco_ios`, `cisco_xe`, `cisco_nxos`, `cisco_xr`, `arista_eos`, `juniper_junos`, `hp_comware`, `nokia_sros`, `paloalto_panos`, `fortinet` etc. (see PLATFORMS); `use_textfsm=True` + `ntc_templates` dir auto-parses [independent].
- **TextFSM/ntc-templates:** 500+ community templates turning `show` output into dicts/lists; index file maps (platform, command) → template [secondary — docker_network_automation].
- **NAPALM getters (selected):** `get_facts`, `get_interfaces`, `get_lldp_neighbors`, `get_bgp_neighbors`, `get_network_instances`, `get_config` (running/candidate/startup), `get_probes_config/results`, `ping`, `traceroute`; config methods: `load_merge_candidate`, `load_replace_candidate`, `compare_config` (diff), `commit_config`, `discard_config`, `rollback` [secondary — NAPALM docs].
- **napalm-logs:** syslog listener normalizing messages to OpenConfig/IETF YANG models — feeds EDA-style event pipelines [secondary — napalm-automation org].
- **Scrapli variants:** `scrapli` (SSH/Telnet generic driver + network drivers), `scrapli-netconf` (NETCONF), `scrapli-community` (extra platforms); async support via `asyncssh`; 2026.6.x requires Python 3.10+ [independent — PyPI].
- **ncclient 0.7.0 ops:** call-home (RFC 8071) for devices initiating NETCONF to the controller (NAT/firewall friendly); YANG 1.1 `action` operations; device handlers (e.g., Nokia SR OS quirks) [secondary — ncclient releases].
- **pygnmi ops:** `pygnmicli` for ad-hoc Get/Set/Subscribe; Set with master arbitration for redundant controllers; auto-encoding detection via Capabilities (prefers `json` over `json_ietf` when both offered); telemetry subscriptions feed streaming pipelines (see Phase E3) [secondary — https://github.com/akarneliuk/pygnmi].
- **pyATS/Genie testing:** `pyats run job` executes test scripts; `genie parse "show ip bgp summary" --os iosxe` produces structured data offline; pyats clean wipes/reloads lab devices through 44+ stages; 25.1 added Ixia ngpf traffic generation [official — https://developer.cisco.com/docs/pyats/25-1/].
- **pyATS learning:** `genie learn` snapshots device state (routing, interfaces, platform) — diff learns across change windows to detect unintended changes [official — Cisco DevNet].
- **pynetbox:** Pythonic NetBox REST client (`nb.dcim.devices.all()`, `nb.ipam.prefixes.create()`); used to build dynamic inventories and to push discovered data back to NetBox [secondary].
- **infrahub-sdk:** GraphQL client for Infrahub (queries, mutations, branches, proposed changes); underpins nornir-infrahub and custom sync jobs [secondary — opsmill].
- **infrahub-mcp 1.1.7:** exposes Infrahub schema/data to AI assistants via Model Context Protocol (stdio + streamable HTTP); branch-isolated writes; FastMCP 3.x; schema write contracts for validation [secondary — https://pypi.org/project/infrahub-mcp/].
- **Dependency hygiene 2026:** nornir-utils moved poetry→uv and pylama/black→ruff; scrapli ships manylinux/musllinux wheels; pin with `uv.lock`/`requirements.txt` hashes in automation images [official — nornir-utils changelog; PyPI].
- **Universal NetDevOps image (v3.0.0, 2026):** Rocky Linux 10 base; Ansible via pipx; `/opt/venv` for Python libs; covers CLI (netmiko/scrapli/napalm), model-driven (ncclient/pygnmi), parsing (textfsm/ntc-templates), SoT (pynetbox), testing (pytest) [secondary — https://github.com/andersonmavi30/docker_network_automation].

---

## Wave 13 — Operational patterns & CI/CD

- **Pipeline stages (network CI/CD):** lint (ansible-lint/tflint/ruff) → syntax/unit (molecule/pytest) → render/dry-run (ansible `--check`, `tofu plan`, resource-module `state=rendered`) → lab deploy (Containerlab/EVE-NG virtual devices) → peer review (PR) → approval gate (AAM workflow/AAP) → staged prod rollout (canary site → fleet) → post-change validation (pyATS diff, NAPALM getters) → commit state/backups [secondary — community guides].
- **Pre/post checks:** snapshot `get_facts`/`get_bgp_neighbors`/interface counters before change; re-run after; fail pipeline on unexpected diff [secondary — pyATS learn/diff docs].
- **Rollback strategies:** NAPALM `rollback()`/discard candidate; Ansible re-apply previous rendered config; OpenTofu `apply` of previous known-good commit (state history); device-local `configure replace` checkpoints on IOS-XR/NX-OS [secondary].
- **Change windows:** AAP workflow approval nodes; EDA rulebooks paused during freeze (disable rulebook via API); Terraform applies gated on schedule [secondary].
- **Inventory from SoT (Ansible):** NetBox inventory plugin (`ansible-inventory -i netbox.yml --list`) groups by site/role/tenant; Infrahub via custom inventory plugin or nornir-infrahub for Nornir [secondary].
- **Inventory from SoT (Nornir):** `nornir-infrahub` inventory plugin (v1.2.0) queries Infrahub GraphQL; hostname/schema mapping configurable; idempotent file transfer tasks included [official — https://github.com/opsmill/nornir-infrahub/blob/HEAD/CHANGELOG.md].
- **Secrets:** Ansible Vault for playbook secrets (vault password via env/CI secret); `community.hashi_vault` lookup for dynamic secrets; OpenTofu ephemeral values + encrypted state; never log secrets (`no_log: true` on tasks handling credentials) [secondary].
- **Brownfield onboarding:** `network.base` validated role builds inventory from live facts → render intended config (`state=rendered`) → diff vs running → staged deploy; OpenTofu `import` blocks for bulk brownfield adoption [vendor-reported; secondary].
- **Compliance as code:** scheduled playbooks/tofu runs assert intended state (ACLs, banners, NTP, SNMP) and remediate or alert on drift; pyATS/Genie parsers power the assertions [secondary].
- **Backup pattern:** nightly Nornir/Ansible jobs pull running configs to Git (per-device files); diff alerts on unexpected changes; restores via config-replace [secondary].
- **GitOps for networks:** Git = desired state; PR = change request; CI = validation; merge = approved apply (Terraform Cloud VCS runs / AWX project sync / Argo-style pullers); full audit trail in git history [secondary — simonpainter; Red Hat].
- **Multi-team scaling:** AAP organizations/teams/RBAC; Terraform per-team state with remote-state data sources for cross-team references (e.g., net team outputs consumed by security team) [secondary].
- **Observability of automation:** log every job (AWX job output, CI logs) to central logging; EDA event throttling metrics; OpenTofu experimental OpenTelemetry traces [secondary — OpenTofu docs].
- **Lab strategy:** Containerlab (vrnetlab images: vr-eos, vr-nxos, vr-xrv9k, vr-sros, vr-junos) for CI; EVE-NG for manual validation; pyATS clean for reset between runs [secondary].
- **Documentation:** auto-generate docs from SoT (NetBox/Infrahub → Markdown diagrams); `terraform-docs` for modules; ansible `ansible-doc` for roles [secondary].

---

## Wave 14 — Selection guide: which tool when

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
