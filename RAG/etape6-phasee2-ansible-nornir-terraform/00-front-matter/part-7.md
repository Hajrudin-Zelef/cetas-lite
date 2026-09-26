---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/part-7
title: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries (part 7)"
domain: front-matter
role: reference
task: reference
actors: []
dates: []
keywords: ["embedding", "latency", "lean"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [183, 201]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: 0a69f32b8d48e0c521dbf1c231bae3288904f460258b2d94d07f51231bc21e8c
---

# Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries (part 7)

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

