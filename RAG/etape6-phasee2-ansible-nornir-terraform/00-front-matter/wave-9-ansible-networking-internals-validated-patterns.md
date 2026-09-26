---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-9-ansible-networking-internals-validated-patterns
title: "Wave 9 — Ansible networking internals & validated patterns"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2026-03", "2026-06", "2026-07", "2026-09-22", "2026-11-30", "2028-03"]
keywords: ["benchmark", "benchmarks", "guardrails", "research"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [139, 182]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: e0f4437a01352606801fbb9409e4f8e72ed3986c30203352bdcc6a138aa06295
---

# Wave 9 — Ansible networking internals & validated patterns

- **Line count:** to be verified after writing (target ≥750).
- **Gaps (no 2026 data found):**
  - cisco.nxos version lineage: field reports show **12.0.0** (mid-2026) while CHANGELOG HEAD documents 9.x — could not confirm the current major line on 2026-09-22. **[unverified]**
  - pygnmi maintenance status: repo docs stale (~2022); latest version number not confirmed. **[unverified]**
  - Infoblox Terraform provider: latest found is v2.5/v2.6 (2024); no 2026 release confirmed. **[unverified]**
  - Terraform 1.9.0 described as "last open-source version under MPL; 2.0+ is BSL" by one Medium source — not confirmed against HashiCorp's own release notes. **[unverified]**
  - Ansible network automation performance-at-scale benchmarks (playbook runs vs Nornir vs pyATS at 1k+ devices): no head-to-head 2026 benchmark found. **[gap]**
  - Adoption statistics (NetDevOps survey numbers, Ansible-for-networks market share 2026): no survey located. **[gap]**
- **Conflicts:**
  - C1: OpenTofu provider-registry size — 3,900+ providers/23,600+ modules (June 2026 Medium) vs 3,000+ providers (Medium, June 2026) vs 3,200+ providers (March 2026). Treated as indicative, not exact.
  - C2: ansible-core 2.19 EOL — endoflife.date says EOL 2026-11-30 (still supported at research date); eosl.date agrees. Consistent.
  - C3: cisco.nxos 9.x vs 12.0.0 — see gap above.
- **Collection-deprecation risk:** `junipernetworks.junos` deprecated (removal from Ansible 14 if unmaintained) while Juniper redirects users to `juniper.device`; network teams should standardize on `juniper.device` FQCNs [official].
- **Round-trip:** all URLs used are verbatim from search-result listings; no SKUs/URLs invented. Facts without a confirmable 2026 source are tagged `[unverified]`.

---

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

