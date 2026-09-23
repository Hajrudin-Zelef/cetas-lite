---
id: etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/wave-19-glossary-concept-map
title: "Wave 19 — Glossary & concept map"
domain: github-workflows-network-ci-yml-sketch
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [555, 607]
section: ".github/workflows/network-ci.yml (sketch)"
sha256: d0c4fc1256676dcfbf0c2d24c04fec0a193a7df381f1e79b56a56d44ae589aa4
---

# Wave 19 — Glossary & concept map

## Wave 19 — Glossary & concept map

- **Idempotence:** repeated runs converge to same state without side effects; resource/state modules and plan/apply are idempotent by design; raw CLI tasks are not [secondary].
- **Check mode / dry-run:** Ansible `--check`, OpenTofu `plan`, NAPALM `compare_config`, resource-module `state=rendered` — validate before mutating [secondary].
- **Source of truth (SoT):** NetBox/Infrahub hold intended inventory; automation reads from and writes back to it [secondary].
- **Brownfield:** existing manually-built networks onboarded via fact-gathering, rendered diffs, and import blocks [secondary].
- **Desired-state / GitOps:** Git holds intent; CI validates; merge triggers apply; state files track reality [secondary].
- **Drift:** out-of-band changes detected by refresh-only plans or compliance scans [secondary].
- **Execution Environment:** container image with pinned ansible-core + collections + Python deps for reproducible runs [secondary].
- **Rulebook:** EDA YAML mapping event sources → conditions → actions [official].
- **Ephemeral values:** OpenTofu 1.11+ secrets that never persist in state [secondary].
- **Proposed change:** Infrahub's Git-like review flow for infra data changes [secondary].
- **Parsers:** textfsm/ntc-templates (regex templates), Genie parsers (Python, Cisco-focused), `cli_parse` (Ansible-native) [secondary].
- **Transports:** SSH CLI (Netmiko/Scrapli/network_cli), NETCONF (ncclient/netconf), REST (httpapi), gNMI (pygnmi/ansible-gnmi) [secondary].
- **Telemetry tie-in:** streaming telemetry (gNMI) feeds EDA/monitoring; config automation consumes the same models — see Phase E3 [secondary].
- **RBAC/approvals:** AAP controller gates who runs what and when; Terraform Cloud run approvals; Infrahub proposed changes [secondary].
- **State locking:** prevents concurrent applies corrupting state; OpenTofu native S3 locking (no DynamoDB) [secondary].

---

## Wave 20 — Collection module catalog (notable modules, 2026)

- **cisco.ios resource modules:** `ios_acls` (multi-word option handling fixed 11.5.0), `ios_vrf_global` (idempotency for multiple import/export route-targets fixed 11.2.x), `ios_hsrp_interfaces` (parsing fixed 11.2.x), `ios_user` (SSH key matching fixed 11.4.2), `ios_config` (prompt-answering in config mode added 11.1.0), `ios_facts`, `ios_command`, `ios_banner`, `ios_snmp_server`, `ios_ntp_global`, `ios_logging_global`, `ios_hostname` [official — cisco.ios CHANGELOG].
- **cisco.iosxr resource modules:** `iosxr_bgp_global` (neighbor shutdown parsing/idempotency fixed 12.4.0), `iosxr_config`, `iosxr_facts`, `iosxr_command`, `iosxr_interfaces`, `iosxr_l3_interfaces`, `iosxr_ospfv2`, `iosxr_netconf` (NETCONF subsystem handling) [official — cisco.iosxr CHANGELOG].
- **cisco.nxos resource modules:** `nxos_vrf_global`, `nxos_vrf_address_family` (9.2.x), `nxos_bgp_global`, `nxos_config`, `nxos_facts`, `nxos_command`, `nxos_vlans`, `nxos_interfaces`, `nxos_nxapi` [official — cisco.nxos CHANGELOG].
- **arista.eos modules:** `eos_config` (new `content` param 12.1.0; `src` auto-Jinja deprecated → removal 2028-03), `eos_facts`, `eos_command`, `eos_vlans`, `eos_interfaces`, `eos_l3_interfaces`, `eos_bgp_global`, `eos_acls`, action plugins renamed with `eos_` prefix 12.1.1 (runtime redirect keeps old names working) [official — arista.eos CHANGELOG].
- **juniper.device modules (migration target):** `junos_config` (load merge/replace/override, confirmed-commit), `junos_facts`, `junos_command`, `junos_rpc`, `junos_netconf`, `junos_package` (Junos upgrades), `junos_scp` — replaces `junipernetworks.junos.*` FQCNs one-for-one [official — juniper changelog].
- **cisco.aci modules (2.11.0):** `aci_tenant`, `aci_bd`, `aci_epg`, `aci_contract`, `aci_l3out`, `aci_vlan_pool`, `aci_rest` (raw API escape hatch) [secondary].
- **cisco.meraki modules (2.20.10):** `meraki_network`, `meraki_vlan`, `meraki_ssid`, `meraki_mx_l3_firewall`, `meraki_organization`, inventory plugin for Dashboard [secondary].
- **cisco.dnac modules (6.31.3):** `dnac_device_*`, `dnac_site_*`, `dnac_provision_*`, `dnac_event_*` — Catalyst Center intent APIs [secondary].
- **ansible.netcommon plugins:** `network_cli`, `netconf`, `httpapi`, `libssh` connection plugins; `cli_parse`/`restconf_parse`/`xmltojson` filter plugins; `network_resource` action plugin backing resource modules [official].
- **ansible.utils:** `cli_parse`, `validate` (JSON-schema validation of vars), `fact_diff`, `ipaddr`/`ipv6` filters, `keep_keys`, `to_paths` — the data-shaping toolkit for templates and assertions [secondary].
- **community.hashi_vault 6.2.0:** `hashi_vault` lookup, `vault_kv2_get`, auth methods (token, approle, userpass, jwt) — dynamic secrets in playbooks [secondary].
- **awx.awx 24.6.1:** `controller_job_template`, `controller_inventory`, `controller_credential`, `controller_workflow_job_template`, `controller_schedule`, `execution_environment` — manage the controller itself [secondary].
- **Module deprecation hygiene:** run `ansible-doc --list` per collection after upgrades; grep playbooks for removed modules before bumping major pins; CI with `--check` catches signature changes early [secondary].

---

## Wave 21 — Hardening checklist for network automation (2026)

- Patch automation hosts: ansible-core CVEs Sept 2026 (CVE-2026-11332 galaxy git-config injection; CVE-2026-16493 galaxy argument injection) — upgrade to 2.20.9/2.21.4 or distro-fixed builds [official/secondary].
- Pin and hash dependencies: `requirements.yml` version ranges, EE digests, `uv.lock`/hashes for Python images; unpinned pulls caused silent breakage in community reports [secondary].
- Secrets: Vault/CI-secret-injected tokens; Ansible Vault for static secrets; `no_log: true` on credential tasks; OpenTofu ephemeral values + encrypted state; never commit keys (panos tutorial anti-pattern) [secondary].
- Least privilege: dedicated automation service accounts per domain; read-only roles for discovery/compliance jobs; write roles gated by AAP approvals [secondary].
- Transport security: prefer SSH with keys over passwords; NETCONF/TLS where supported; verify host keys (`host_key_checking` discipline) [secondary].
- Controller hardening: AAP behind SSO/OIDC; RBAC per team; audit logging to SIEM; EDA webhook sources authenticated [secondary].
- Change safety: check-mode/plan gates; canary sites; pre/post pyATS diffs; rollback plans; freeze windows enforced by disabling schedules/rulebooks [secondary].
- Supply chain: verify collection signatures (`ansible-sign`), scan EEs for CVEs, mirror Galaxy internally for air-gapped estates [secondary].
- State protection: remote encrypted state; state locking; per-environment state separation; backup state before migrations [secondary].
- Telemetry awareness: RHEL ansible-core 2.16.3 added telemetry (RHEL-219532) — review data-sharing posture in regulated environments [official].

---

