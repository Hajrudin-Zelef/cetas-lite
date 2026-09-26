---
id: etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22
title: "Wave 25 — Key numbers at a glance (2026-09-22)"
domain: github-workflows-network-ci-yml-sketch
role: deep-dive
task: reference
actors: []
dates: ["2026-07", "2026-08-19", "2026-09-22", "2026-11-30", "2027-05", "2028-04-01"]
keywords: ["guardrails", "mcp", "research", "sandbox"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [676, 731]
section: ".github/workflows/network-ci.yml (sketch)"
sha256: 42410a70edb8acc0d6582f425480102bdc61c898439cd64afd0cde8bdfd91180
---

# Wave 25 — Key numbers at a glance (2026-09-22)

## Wave 25 — Key numbers at a glance (2026-09-22)

- ansible-core: 2.21.4 current; 2.19 EOL 2026-11-30; 2.22 GA Nov 2026 [official].
- ansible-dev-tools 26.1.0; molecule/ansible-lint/ansible-navigator 26.x CalVer [official/secondary].
- cisco.ios 11.5.0; cisco.iosxr 12.4.0; arista.eos 12.1.2; junipernetworks.junos 11.1.1 → juniper.device [official].
- AWX 24.6.1 (last formal); AAP 2.5 (Mar 2026); EDA 1.1.17 [official/secondary].
- Nornir 3.6.0 (py≥3.10); nornir-utils 0.3.0; nornir-infrahub 1.2.0; nornir_salt 0.23.3 [official/independent].
- OpenTofu 1.12.6 (2026-08-19); MPL-2.0; CNCF Sandbox since Apr 2025 [independent].
- Netmiko 4.8.0; Scrapli 2026.6.13rc14; ncclient 0.7.0; pyATS 25.1 (+142 parsers, +78 APIs) [official/independent].
- ntc-templates 500+; Genie ~4,500+ parsers / ~2,400+ APIs (historical totals) [secondary/official].
- Provider registry ~3,900 providers claimed (conflict C1) [secondary].
- CVEs Sept 2026: CVE-2026-11332, CVE-2026-16493 (ansible-galaxy) [official/secondary].
- Deprecation deadlines: arista.eos `src` auto-templating → 2028-03; junipernetworks.junos redirects → 2028-04-01 / Ansible 14 [official].
- Infrahub 1.11.0; infrahub-mcp 1.1.7; NetBox provider e-breuninger/netbox 4.1.0 [secondary].
- F5OS provider 1.12.0 (Go 1.25+); panos 1.11.0 (tutorial); Infoblox provider 2.5/2.6 (2024) [official/secondary].

---

## Wave 26 — FAQ (field questions answered from research)

- **"Which ansible-core should I run in Sept 2026?"** 2.21.x (supported to Nov 2027) or 2.20.x (to May 2027); leave 2.19 before its Nov 2026 EOL; 2.18 is EOL [official].
- **"Is junipernetworks.junos dead?"** Deprecated; works via redirects to juniper.device until 2028-04-01, but may drop from Ansible 14 — migrate FQCNs now [official].
- **"Can I manage Cisco IOS with Terraform?"** Only via controllers/APIs (Catalyst Center, ACI, Meraki); CLI-only IOS needs Ansible/Nornir [secondary].
- **"OpenTofu or Terraform for a new network IaC project?"** OpenTofu for MPL-2.0 + state encryption + ephemerals; Terraform for HCP Terraform backend and day-one provider releases [secondary].
- **"Do I still need Netmiko if I have Nornir?"** Nornir orchestrates; Netmiko/Scrapli/NAPALM do the device talking — they compose, not compete [secondary].
- **"Ansible or Nornir for 5,000 switches?"** Either scales with tuning (Ansible: persistent connections + fact cache + serial; Nornir: thread pool + filters); pick by team skill and controller needs [secondary].
- **"How do I handle secrets?"** Vault dynamic secrets; Ansible Vault for static; OpenTofu ephemerals + encrypted state; `no_log: true` [secondary].
- **"What validates a change worked?"** pyATS/Genie diffs, NAPALM getters pre/post, `tofu plan -refresh-only` drift jobs, compliance playbooks on schedule [secondary].
- **"Where does telemetry fit?"** gNMI streaming (pygnmi/ansible-gnmi) feeds monitoring and EDA rulebooks; config automation consumes the same data models — see Phase E3 [secondary].
- **"Is AWX dead?"** No — 24.6.1 was the last formal release during a refactor; operator and community builds continue; AAP is the supported path [secondary].
- End of file.

---

## Wave 27 — Adoption signals & community health (2026)

- **Ansible network mindshare:** Red Hat's validated network content program (network.base/network.interfaces) signals continued enterprise investment; Cisco maintains first-party collections (ios/nxos/iosxr/aci/meraki/dnac/ise) plus validated Catalyst Center IaC [vendor-reported/official].
- **Collection release cadence:** cisco.ios 11.4.2 → 11.5.0 within a month (June–July 2026); arista.eos 12.1.0 → 12.1.2 rapid patches — vendors actively tracking ansible-core changes [official].
- **Nornir steadiness:** core at 3.6.0 with plugins evolving independently (nornir-utils ruff/uv modernization, nornir-infrahub Infrahub coupling) — the "no bundled plugins since 3.0" model keeps core stable while the ecosystem moves [official/independent].
- **Scrapli momentum:** CalVer pre-releases through mid-2026, ssh2 transport, active repo (days-old commits at research date) — the speed-focused alternative to Netmiko keeps gaining [independent].
- **OpenTofu traction:** CNCF Sandbox (Apr 2025), practitioner ADRs switching binary-only, registry parity claims — the fork is past the "will it survive" question [independent/secondary].
- **NetBox-as-code pattern:** multiple independent 2026 sources (blogs, homelabs) converge on e-breuninger/netbox + Terraform Cloud + GitHub Actions — a de-facto reference stack [secondary].
- **Infrahub rise:** 1.11.0 schema contracts, nornir-infrahub 1.2.0, infrahub-mcp for AI assistants — positioned as the Git-like SoT alternative to NetBox [secondary].
- **pyATS endurance:** 25.1 with Python 3.13 and 4,500+ parsers — Cisco's test framework remains the validation backbone for Cisco-heavy shops [official].
- **EDA maturation:** from AAP 2.4 GA to 1.1.17 with throttling guardrails and documented event-storm patterns — event-driven network remediation is production-real [official/secondary].
- **Container-native automation:** EE images, awx-operator, universal NetDevOps Docker images, Containerlab CI — 2026 network automation runs in containers by default [secondary].
- **AI-assisted NetOps:** infrahub-mcp exposes SoT to AI assistants; validated playbooks lower the authoring bar — automation authorship is broadening beyond CLI experts [secondary].
- **Skills gap note:** demand for Jinja2/Python/HCL/Git fluency in network teams keeps rising; NornFlow and validated content explicitly target lowering that bar [secondary].
- **Risk watch 2026–2027:** ansible-core 2.19 EOL (Nov 2026), 2.22 GA (Nov 2026), AAP container-only shift (2.7+), juniper collection deprecation, OpenTofu provider lag — plan upgrades against these dates [official/secondary].
- **What "done" looks like:** SoT in NetBox/Infrahub → platforms via Terraform/OpenTofu → devices via Ansible/Nornir → validation via pyATS → events via EDA → everything in Git with CI gates [secondary].
- End of file.

---

## Wave 28 — Quick-reference commands (2026 tooling)

