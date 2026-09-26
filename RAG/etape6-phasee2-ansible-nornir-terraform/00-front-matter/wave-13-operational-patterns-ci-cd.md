---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-13-operational-patterns-ci-cd
title: "Wave 13 — Operational patterns & CI/CD"
domain: front-matter
role: reference
task: reference
actors: []
dates: []
keywords: ["mcp", "model context protocol"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [230, 269]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: 7d4087cac9a4f2f132b7e2ac33624cea819993d17c62b5e7e3579f22d7129678
---

# Wave 13 — Operational patterns & CI/CD

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

