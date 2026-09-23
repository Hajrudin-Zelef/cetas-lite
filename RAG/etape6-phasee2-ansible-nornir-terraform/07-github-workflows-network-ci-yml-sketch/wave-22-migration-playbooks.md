---
id: etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/wave-22-migration-playbooks
title: "Wave 22 — Migration playbooks"
domain: github-workflows-network-ci-yml-sketch
role: deep-dive
task: reference
actors: []
dates: ["2028-03", "2028-04-01"]
keywords: ["apache", "license", "open source"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [608, 675]
section: ".github/workflows/network-ci.yml (sketch)"
sha256: b423980a96607102a713e1ac2d56c7b4645688f5d393033b152949b248c8b97f
---

# Wave 22 — Migration playbooks

## Wave 22 — Migration playbooks

- **junipernetworks.junos → juniper.device:** 1) inventory FQCN usage (`grep -r junipernetworks.junos`); 2) install `juniper.device` (>=2.0.0); 3) rewrite FQCNs (`juniper.device.junos_config` etc.); 4) run in lab with `--check`; 5) complete before 2028-04-01 redirect removal / Ansible 14 drop [official — juniper changelog; porting guide].
- **arista.eos `src` → `content`:** replace `eos_config: src=template.j2` with `content: "{{ lookup('ansible.builtin.template', 'template.j2') }}"`; deadline March 2028 [official — arista.eos CHANGELOG].
- **cisco.ios 10.x → 11.x:** requires ansible.netcommon >=8.1.0 and ansible-core >=2.19-era; last 10.x line (10.1.1) stays for ansible-core ≤2.18 estates [official — cisco.ios CHANGELOG].
- **Terraform → OpenTofu:** 1) audit for `terraform`-specific data sources/modules; 2) pin `opentofu = "1.12.6"` (mise); 3) `tofu init` (registry-compatible); 4) `tofu plan` diff vs terraform plan — expect no changes; 5) migrate state backend config; 6) switch CI images; 7) adopt state encryption + ephemeral values post-migration [secondary — smana ADR; OpenTofu docs].
- **AWX 24.6.1 → community builds/operator:** upstream formal releases paused; options: awx-operator 2.12.x, fitbeard/automation-platform 26.0.0 images, or AAP subscription; test job parity in staging [secondary].
- **Brownfield → GitOps:** 1) build inventory from live facts (network.base roles); 2) render intended configs; 3) diff and review; 4) import into OpenTofu (`import` blocks) or commit rendered configs; 5) enable drift detection; 6) cut over change process to PRs [secondary].
- **CLI scripts → Nornir/Ansible:** wrap Netmiko/Scrapli snippets as Nornir tasks or Ansible modules; add inventory from SoT; add CI; retire snowflake scripts per device-type [secondary].

---

## Wave 23 — Open items carried forward

- Confirm cisco.nxos current major line (9.x vs 12.0.0) against Galaxy API.
- Confirm pygnmi latest release and maintenance status.
- Confirm Infoblox Terraform provider 2026 releases.
- Verify Terraform 1.9 "last MPL" claim against HashiCorp release notes.
- Locate 2026 NetDevOps adoption survey numbers.
- Re-check junipernetworks.junos Galaxy tagging status (issue #589).
- Track ansible-core 2.22 GA (Nov 2026) and AAP 2.7 container-only migration notes.
- End of file.

---

## Wave 24 — Head-to-head comparison matrices

### Ansible vs Nornir vs Terraform/OpenTofu

| Dimension | Ansible | Nornir | Terraform / OpenTofu |
|---|---|---|---|
| Paradigm | Declarative playbooks (YAML) [secondary] | Imperative Python tasks [secondary] | Declarative HCL + state [secondary] |
| Learning curve | Low (YAML), moderate (Jinja2) [secondary] | Requires Python [secondary] | HCL + state mental model [secondary] |
| Multi-vendor network modules | Rich (cisco.ios/arista.eos/juniper.device…) [official] | Via NAPALM/Netmiko/Scrapli plugins [secondary] | Via API providers (Apstra, Meraki, NetBox…) [secondary] |
| CLI-only legacy devices | Yes (network_cli) [secondary] | Yes (Netmiko/Scrapli) [secondary] | No — needs API provider [secondary] |
| Controller/RBAC | AAP/AWX (approvals, schedules, RBAC) [official] | DIY (wrap in app/CI) [secondary] | Terraform Cloud / Spacelift / CI [secondary] |
| Event-driven | EDA rulebooks [official] | Custom (napalm-logs + runner) [secondary] | No native eventing [secondary] |
| Dry-run | `--check` (module-dependent) [secondary] | Manual (diff before commit) [secondary] | `plan` (first-class) [secondary] |
| State tracking | No persistent state [secondary] | No persistent state [secondary] | `tfstate` (encrypted in OpenTofu) [secondary] |
| Scale pattern | Persistent connections, fact cache, serial [secondary] | Thread pool, connection reuse [secondary] | Parallel provider calls, targeted apply [secondary] |
| License (2026) | GPLv3+ / AAP subscription [independent] | Apache-2.0 [independent] | BSL 1.1 / MPL-2.0 [independent] |
| Best fit | Operator teams, compliance playbooks [secondary] | Python teams, custom logic [secondary] | Platform teams, API-driven infra [secondary] |

### Python library selection

| Library | Transport | Strength | Watch-out |
|---|---|---|---|
| Netmiko 4.8.0 | SSH (Paramiko) | Broadest platform coverage, textfsm auto-parse [independent] | Screen-scraping brittleness [secondary] |
| Scrapli 2026.6.x | SSH/Telnet/NETCONF | Speed (ssh2), typing, tests [independent] | Pre-release CalVer line [independent] |
| NAPALM | SSH/API per driver | Unified getters, diff/commit/rollback [secondary] | Driver maintenance varies [secondary] |
| ncclient 0.7.0 | NETCONF | Call-home, YANG 1.1 actions [secondary] | XML verbosity [secondary] |
| pygnmi | gNMI | Streaming telemetry + Set [secondary] | Maintenance status unclear [unverified] |
| pyATS/Genie 25.1 | SSH/API | 4,500+ parsers, test framework [official] | Cisco-centric [official] |

### Controller / platform selection

| Platform | Model | 2026 note |
|---|---|---|
| AWX upstream | Open source controller | 24.6.1 last formal; refactor in progress [secondary] |
| awx-operator | K8s deployment | ~2.12.x current [secondary] |
| AAP 2.5 | Red Hat subscription | Controller 4.6.27, EDA 1.1.17 (Mar 2026) [official] |
| AAP 2.6/2.7+ | Red Hat subscription | 2.6 last RPM; 2.7+ container-only [secondary] |
| Terraform Cloud | SaaS runs | Terraform-only backend [secondary] |
| Spacelift/env0/Scalr | SaaS/adjacent | OpenTofu-compatible [secondary] |
| NetBox / Infrahub | SoT | NetBox Apache-2.0; Infrahub AGPLv3 [independent] |

---

