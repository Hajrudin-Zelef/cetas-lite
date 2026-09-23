---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-9-ansible-networking-internals-validated-patterns
title: "Wave 9 — Ansible networking internals & validated patterns"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2026-07", "2028-03"]
keywords: ["embedding", "guardrails", "latency", "lean"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [156, 201]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: 05c52ffa3f36b2b4542976d5524a5ada3046657aaa3eec12d2df00675368563a
---

# Wave 9 — Ansible networking internals & validated patterns

## Wave 9 — Ansible networking internals & validated patterns

- **Connection plugins:** `ansible.netcommon.network_cli` (SSH CLI with persistent connection), `ansible.netcommon.netconf` (NETCONF over SSH, used by Junos and IOS-XR), `ansible.netcommon.httpapi` (REST — eAPI for EOS, NX-API, Meraki/DNAC cloud APIs) [secondary — field playbooks; skills docs].
- **Persistent connections** multiplex CLI/NETCONF sessions per host across tasks in a play, reducing SSH handshake overhead at scale; enable with `ansible_connection: ansible.netcommon.network_cli` + `ansible_network_cli_ssh_type: paramiko|libssh` [secondary].
- **Facts modules** (`ios_facts`, `eos_facts`, `junos_facts`, `nxos_facts`, `iosxr_facts`) gather structured device facts into `ansible_net_*` variables (hostname, version, serial, model, interfaces, neighbors) usable in Jinja2 and inventory [official — collection docs].
- **Resource modules** (e.g., `ios_interfaces`, `ios_l3_interfaces`, `ios_acls`, `ios_vrf_global`, `iosxr_bgp_global`, `eos_config`) implement state-driven config with `state: merged|replaced|overridden|deleted|rendered|parsed|gathered` semantics; `rendered` outputs CLI without applying (dry-run friendly) [official — collection docs; arista.eos changelog].
- **cli_parse:** `ansible.utils.cli_parse` parses raw CLI output into structured data using `ntc_templates` or `pyats` parsers — key for brownfield discovery and compliance checks [secondary — Red Hat validated content].
- **Red Hat validated network content** collections: `network.base` (platform-agnostic roles: build brownfield inventory, persist facts, deploy rendered configs) and `network.interfaces`; roles: Resource Manager (state management), gather network resource facts, persist, deploy, configure, remediate [vendor-reported — https://www.redhat.com/pt-br/blog/accelerating-your-network-automation-journey-ansible-network-validated-content].
- **Multi-vendor playbook pattern (real):** inventory groups by `ansible_network_os`; group_vars set connection plugin per OS; roles per vendor; Jinja2 templates for common services (banner, SNMP, ACLs, NTP); `serial` batching for rolling changes [secondary — https://github.com/tamersaid2022/ansible-network-playbooks].
- **ansible-navigator 26.x**: TUI to develop/run playbooks inside Execution Environments; integrates with ansible-lint and molecule in the dev-tools bundle [secondary — https://forum.ansible.com/t/release-announcement-ansible-dev-tools-v26-1-0/45073].
- **Execution Environments:** container images (ansible-builder 3.1.x, EE v3 schema) containing ansible-core + pinned collections + Python deps; best practice: pin base image by digest, build in CI, scan for CVEs (Sept 2026 CVEs affect galaxy installs — see Wave 1) [secondary — skills docs; AlmaLinux changelog].
- **Check mode caveat:** network resource modules support check mode unevenly; always test check-mode behavior of new collection versions in a lab before production use [secondary].
- **ansible-lint 26.x / molecule 26.x:** lint playbooks/roles in CI; molecule tests roles against containerized or virtual network OS images (drivers: docker, delegated to Containerlab/EVE-NG topologies) [secondary — dev-tools announcement; skills docs].
- **AWX/AAP inventory sources:** sync inventory from NetBox/Infrahub/custom scripts into controller; credentials types for network_cli (SSH key/password), Vault, cloud [secondary — AWX docs community].
- **Surveys & RBAC in AAP:** parameterize job templates (VLAN ID, site) with surveys; RBAC limits who can launch network change jobs — approval nodes in workflows for change windows [secondary].
- **EDA for networks (rulebook sketch):** event source (webhook from monitoring) → condition (interface down, debounce with `once_within`) → action (`run_job_template` running a remediation playbook or embedded tasks) [official — Red Hat EDA blog].
- **Event storm guardrails:** EDA ships default throttling; rulebooks should add `once_within`/`once_after` on flapping-prone events to avoid flooding the controller with remediation jobs [secondary — ansible-automator SKILL.md].
- **Idempotency rule of thumb:** prefer resource/state modules; reserve `*_config`/`*_command` for one-offs; `shell`/`command` never idempotent by themselves [secondary — ansible-automator SKILL.md].
- **Upgrade hygiene:** read collection changelogs (cisco.ios 11.x deprecated-import cleanups for ansible-core 2.23; arista.eos `src` auto-templating removal March 2028) before bumping pins; run `--check` after upgrades [official changelogs].
- **Telemetry note:** RHEL-packaged ansible-core 2.16.3 (July 2026) added telemetry for RHEL (RHEL-219532) — relevant for regulated environments auditing automation tooling [official — AlmaLinux git changelog].
- **Deprecation watch:** `junipernetworks.junos` → migrate FQCNs to `juniper.device.*`; `arista.eos` `eos_config src` auto-Jinja removal 2028-03; `ansible.builtin` warnings API changes (use `AnsibleModule.warn()`) [official changelogs].
- **Ansible for security policy:** community playbooks manage Palo Alto/Fortinet/Check Point (check_point.mgmt 6.4.0, community.fortios) alongside network OS collections for firewall-adjacent automation [secondary — env dump; rogerperkin tools list].

---

## Wave 10 — Nornir architecture & usage patterns

- **Core model:** Nornir = inventory (hosts/groups with data) + tasks (Python functions) + runners (threaded execution) + results (structured per-host). No DSL — plain Python, versioned like software [secondary — Nornir docs/book].
- **Runners:** `threaded` runner (default) with `num_workers` tuning; `serial` runner for ordered changes (e.g., core switches before leaves) [secondary — Nornir docs].
- **Processors:** transform/aggregate results across hosts (e.g., build a compliance report dict from all `napalm_get` facts) [secondary — Nornir docs].
- **Filtering:** `nr.filter(site="dc1", role="spine")` selects subsets; dynamic inventory from NetBox/Infrahub via plugins (`nornir-infrahub`) keeps the filter keys fresh [official — https://github.com/opsmill/nornir-infrahub/blob/HEAD/CHANGELOG.md].
- **Task pattern (canonical):** define `def backup_config(task): r = task.run(task=napalm_cli, commands=["show running-config"])`; `nr.run(task=backup_config)`; inspect `result.failed_hosts` [secondary — 2026 Cisco Press book].
- **nornir_napalm tasks:** `napalm_get` (getters: facts, interfaces, bgp_neighbors, lldp_neighbors), `napalm_cli`, `napalm_configure` (merge/replace with diff + commit/rollback), `napalm_install_config` [secondary — book; NAPALM docs].
- **nornir_netmiko tasks:** `netmiko_send_command`, `netmiko_send_config`, `netmiko_save_config`, `netmiko_commit` (for Junos-style commit) [secondary].
- **nornir_scrapli tasks:** `send_command`, `send_config`, `send_configs_from_file`; benefits: typed API, fast ssh2 transport, unit tests against virtual devices [official — https://github.com/scrapli/nornir_scrapli].
- **NornFlow (0.9.0):** layers declarative YAML workflows over Nornir — CLI (`nornflow run`), variable precedence (environment → global → domain → workflow → CLI → runtime), hooks (pre/post task), failure strategies (skip-failed, fail-fast, run-all), Jinja2 filters; aimed at teams wanting Ansible-like UX without leaving Nornir [independent — https://pypi.org/project/nornflow/0.9.0/].
- **NorFab (via nornir_salt 0.23.3):** Salt-based network automation fabric using Nornir plugins; proxy-minion model for event-driven device interaction [independent — https://pypi.org/project/nornir_salt/0.23.3/].
- **nornir-srl 0.2.1:** ready-made Nornir tasks for Nokia SR Linux in Containerlab (bgp-peers, bgp-rib, lldp, mac table, sub-interfaces) + `fcli` containerized CLI [independent — https://pypi.org/project/nornir-srl/0.2.1/].
- **Scaling guidance:** threads are I/O-bound friendly (SSH latency dominates); for 1k+ devices, batch with filters, stagger `num_workers`, and prefer connection reuse (scrapli persistent) over per-task reconnects [secondary].
- **Credentials pattern:** pull from Vault/Infrahub/NetBox at runtime (never hardcode); nornir-infrahub maps credential stores into inventory data [official — nornir-infrahub changelog].
- **When Nornir wins:** custom logic (diffing, multi-device transactions, conditional rollbacks), Python-native teams, embedding automation inside larger apps; **when Ansible wins:** declarative playbooks, controller/AAP RBAC/approvals, large operator teams without Python depth [secondary — community comparisons].
- **Testing Nornir:** pytest tasks against Containerlab topologies; record fixtures of `napalm_get` outputs for unit tests [secondary].
- **Community health 2026:** core stable (3.6.0), plugin ecosystem maintained separately (nornir-utils, nornir-infrahub, nornir_scrapli active); no bundled plugins since 3.0 keeps core lean [independent — libraries.io; GitHub activity].

---

