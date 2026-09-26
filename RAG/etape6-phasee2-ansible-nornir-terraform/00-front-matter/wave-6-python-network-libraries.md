---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-6-python-network-libraries
title: "Wave 6 — Python network libraries"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2026-03-12", "2026-04-16", "2026-06-13", "2026-07-03"]
keywords: ["apache", "mcp"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [108, 138]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: 1eda0a9d88c67882c6fe54b5a84ad0ac57b674c1afbfb4352dfe972c8099bbec
---

# Wave 6 — Python network libraries

## Wave 6 — Python network libraries

- **Netmiko 4.8.0** (PyPI, Sept 2026): multi-vendor SSH CLI library built on Paramiko; extras `netmiko[par4]` (Paramiko 4 compat) and `netmiko[bulk-encrypt]`; platform coverage documented in PLATFORMS file; `use_textfsm=True` auto-parses with ntc-templates [independent — https://libraries.io/pypi/netmiko].
- **NAPALM**: repo **updated 2026-04-16** (still active, Apache-2.0); cross-vendor abstraction (get_facts, get_config, load_merge/replace_candidate_config, commit/rollback); drivers for IOS/IOS-XE/NX-OS/IOS-XR/EOS/Junos; napalm-logs normalizes syslogs to OpenConfig/IETF YANG models [secondary — https://github.com/napalm-automation].
- **Scrapli 2026.6.13rc14** (2026-06-13): CalVer; **Python 3.10+**; SSH/Telnet/NETCONF transports; wheels for manylinux/musllinux/macOS arm64; speed-focused with optional `ssh2` transport; fully typed; `scrapli-netconf` for NETCONF [independent — https://pypi.org/project/scrapli/2026.6.13rc14/].
- **ncclient 0.7.0** (2026-03-12): PEP517 build; NETCONF client with call-home (RFC 8071), YANG 1.1 action support, Nokia SR OS device handler, lazy-loaded transports, libssh SSH transport option, TLS support, Paramiko dependency update [secondary — https://github.com/ncclient/ncclient/releases; https://www.freshports.org/net-mgmt/py-ncclient/].
- **pygnmi** (akarneliuk, pure-Python gNMI client): 0.8.x line with IPv6 support, master arbitration for Set, auto-encoding detection via Capabilities() (json priority over json_ietf), `pygnmicli` tool; repo docs last updated ~2022 — **maintenance status flagged as gap, see Wave 8** [secondary — https://github.com/akarneliuk/pygnmi].
- **pyATS 25.1** (Cisco DevNet): Python 3.13 support; Genie +142 parsers / +78 APIs; Ixia ngpf traffic-gen support; SMU hot-install stages; testbed-wide parallel methods (connect/disconnect/execute/configure/parse); `pyats clean` with 44+ stages; Health Check; Genie total ~4,500+ parsers / 2,400+ APIs historically [official — https://developer.cisco.com/docs/pyats/25-1/].
- **Parsing:** `textfsm` + **ntc-templates** (500+ multi-vendor templates) remain the standard CLI-parsing pair [secondary — docker_network_automation].
- **Source-of-truth clients:** `pynetbox` for NetBox REST API; `infrahub-sdk` (1.16+/1.20+) for Infrahub GraphQL; `infrahub-mcp 1.1.7` (2026-07-03) exposes Infrahub to AI assistants via MCP (stdio/streamable-http, branch-isolated writes, FastMCP 3.x) [secondary — https://pypi.org/project/infrahub-mcp/].
- **Infrahub** (OpsMill, AGPLv3): schema-driven infra source of truth with Git-like branching, proposed changes, generators; **v1.11.0** (Aug 2026) with write/read schema contracts; **1.0** added OIDC SSO + permission framework; observability via Alloy/Loki/Prometheus/Tempo/Grafana Helm charts [secondary — https://github.com/opsmill/infrahub/releases/tag/infrahub-v1.11.0].
- Testing: `pytest` + **pytest-ansible** for Ansible module/plugin unit tests; **molecule 26.x** for role testing (multi-driver); Containerlab/EVE-NG for lab topologies [secondary].
- The 2026 NetDevOps Docker pattern bundles: netmiko, scrapli, napalm, nornir (orchestration), ncclient, scrapli-netconf, pygnmi (model-driven), textfsm+ntc-templates (parsing), pynetbox (SoT), pytest (testing) — covering CLI, APIs, NETCONF, gNMI, orchestration, parsing, SoT, testing [secondary — https://github.com/andersonmavi30/docker_network_automation/blob/HEAD/README.md].

---

## Wave 7 — Cross-cutting patterns

- **Idempotency:** prefer native/resource modules (which track state and report changes) over `shell`/`command` modules — shell tasks execute blindly and break idempotency guarantees [secondary — ansible-automator SKILL.md]. NAPALM config workflows show diffs before commit and support rollback [secondary]. OpenTofu `plan`/`apply` is natively declarative; use `removed` blocks for clean resource retirement [secondary].
- **Dry-run / check mode:** run Ansible playbooks in **check mode** first to catch deprecated parameter errors after collection upgrades; Terraform/OpenTofu `plan` (-refresh-only for drift) is the equivalent gate; NAPALM `load_*_candidate_config` + diff before `commit_config` [secondary — multiple skills docs; community guides].
- **Pinning:** pin collection versions in `requirements.yml` (e.g., `cisco.ios: ">=11.0.0,<12.0.0"`); pin EE base images by digest, not `latest`; pin OpenTofu/Terraform versions (e.g., mise.toml) — unpinned deps cause silent breakage on upstream releases [secondary — chrishuffman5 ansible-network SKILL.md; skills docs].
- **Jinja2 templating:** resource modules + Jinja2 templates remain the standard for config generation; arista.eos `eos_config` now prefers the `content` parameter with `ansible.builtin.template` lookup over deprecated `src` auto-templating [official — arista.eos CHANGELOG]. NornFlow adds rich Jinja2 filters on top of Nornir [independent].
- **Secrets handling:** Ansible Vault encrypts variables (vault password = root of trust; never store it in source control); `community.hashi_vault` 6.2.0 integrates HashiCorp/OpenBao Vault; OpenTofu **ephemeral values** (1.11+) keep secrets out of state; state encryption at rest (OpenTofu 1.7+) via passphrase/KMS [secondary — skills docs; OpenTofu docs].
- **Inventory from source of truth:** NetBox as inventory via `e-breuninger/netbox` Terraform provider or `pynetbox`; Infrahub as GraphQL-backed Nornir inventory (`nornir-infrahub`); Ansible `network.base` validated content builds brownfield inventory from live device facts into host_vars [secondary — simonpainter blog; opsmill; Red Hat validated content].
- **Performance at scale:** Ansible — fact caching (Redis/JSON), `serial` batching, persistent connections for network_cli; Nornir — threaded task dispatch, num_workers tuning; Scrapli ssh2 transport for raw speed; EDA debounce rules for event storms [secondary — skills docs; nornir docs].
- **Multi-vendor playbook pattern (real example):** inventory groups keyed by `ansible_network_os` (cisco.ios.ios, cisco.nxos.nxos, junipernetworks.junos.junos) with `ansible_connection: ansible.netcommon.network_cli` (or netconf for Junos), roles per vendor, Jinja2 templates for banner/snmp/acl, compliance baseline in files/ [secondary — https://github.com/tamersaid2022/ansible-network-playbooks].

---

## Wave 8 — Verification log, gaps & conflicts

