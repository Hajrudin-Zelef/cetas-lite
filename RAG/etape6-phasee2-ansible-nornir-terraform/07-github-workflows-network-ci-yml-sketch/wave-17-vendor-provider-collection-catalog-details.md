---
id: etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details
title: "Wave 17 — Vendor provider & collection catalog details"
domain: github-workflows-network-ci-yml-sketch
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2026-01-08", "2026-03-12", "2026-03-25", "2026-04-16", "2026-05", "2026-05-31", "2026-06-01", "2026-06-13", "2026-06-14", "2026-06-22", "2026-07-03", "2026-07-23", "2026-07-30", "2026-08-19", "2026-09-08", "2026-09-10", "2026-09-22", "2028-04-01"]
keywords: ["aws", "compute", "mcp", "research"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [503, 554]
section: ".github/workflows/network-ci.yml (sketch)"
sha256: a3017948790e3cef3961978d8ff325d74ca2ae9700aae3119af9f47bfeec69a6
---

# Wave 17 — Vendor provider & collection catalog details

## Wave 17 — Vendor provider & collection catalog details

- **Cisco ACI provider (CiscoDevNet/aci):** manages tenants, bridge domains, EPGs, contracts, L3Outs as Terraform resources; pairs with `cisco.aci` Ansible collection (2.11.0) for mixed IaC/playbook estates [secondary — simonpainter blog].
- **Juniper Mist provider (Juniper/mist):** manages Mist cloud orgs, sites, WLANs, device profiles; Terraform is the documented GitOps path for Mist at scale [secondary — simonpainter blog].
- **FortiManager provider (fortinetdev/fortimanager):** policy packages, ADOMs, device groups; complements Ansible `community.fortios` for device-level tasks [secondary].
- **Cisco Meraki provider + cisco.meraki 2.20.10:** dashboard API as code (networks, VLANs, SSIDs, firewall rules); Meraki's cloud API makes it one of the most Terraform-friendly network platforms [secondary].
- **cisco.dnac 6.31.3 / catalyst-center-ansible-iac 1.0.0:** Catalyst Center Day 0/1/2/N via validated playbooks (site hierarchy, LAN automation, SD-Access fabrics, virtual networks) — Cisco's endorsed IaC path for campus/branch [official — catalyst-center-ansible-iac CHANGELOG].
- **cisco.ise 2.10.0:** ISE policy/endpoint automation; pairs with Catalyst Center AAA integration playbooks [secondary].
- **cisco.asa 6.1.0:** ASA firewall object/policy automation for legacy estates [secondary].
- **cisco.ucs 1.16.0 / cisco.mso 2.10.0 / cisco.intersight 2.0.20:** compute + multi-site orchestration collections rounding out Cisco's IaC story [secondary].
- **check_point.mgmt 6.4.0:** Check Point management API automation [secondary].
- **community.network 5.1.0:** long-tail platforms (EdgeOS, RouterOS, fortios legacy, etc.) for brownfield estates [secondary].
- **ansible.posix 1.6.2 / community.general 10.6.0 / amazon.aws 9.4.0:** supporting cast in network stacks (firewalld/sysctl on Linux jump hosts, AWS for cloud-adjacent networking) [secondary].
- **awx.awx 24.6.1:** manage AAP/AWX itself as code (inventories, job templates, credentials) — controller configuration in Git [secondary].
- **community.hashi_vault 6.2.0:** Vault KV lookups and dynamic secrets inside playbooks [secondary].
- **Provider auth patterns:** env-var-injected tokens (never committed), Vault dynamic credentials, OIDC for cloud-adjacent providers; the panos tutorial's committed API key is the anti-pattern [secondary].
- **Version pinning matrix (real-world May 2026 stack):** ansible 13.1.0 / ansible-core 2.20.x; arista.eos 10.1.1; cisco.ios 11.x; juniper migration in progress — shows a healthy estate pins per-collection ranges, not latest [secondary — env dump].
- **Collection dependency chains:** cisco.ios 11.x → ansible.netcommon >=8.1.0; arista.eos 12.1.x → ansible.netcommon >=8.5.2; junipernetworks.junos 11.x → ansible.netcommon >=8.1.0 + ansible-core >=2.16 — upgrading one collection can force netcommon/core upgrades [official changelogs].
- **Last-compatible versions:** cisco.ios 10.1.1 for ansible-core ≤2.18; teams on older cores must pin accordingly [official — cisco.ios CHANGELOG].
- **Galaxy vs git tags:** junipernetworks.junos 11.1.0 shipped to Galaxy untagged — automation that pins by git tag would miss it; prefer Galaxy version pins [secondary — issue #589].
- **ansible-gnmi v3.0.0 scope:** Cisco IOS-XE/IOS-XR/NX-OS tested; other platforms via `platform='auto'`; removed unmaintained nokia_sros/arista_eos profiles — check platform support before adopting for non-Cisco [official — commit d759860].
- **ansible-nd 1.5.0:** Nexus Dashboard backup/restore + ND 4.x — DC controller state management alongside ACI [official — CHANGELOG].

---

## Wave 18 — 2026 release timeline (network automation)

- 2026-01-08: ansible-dev-tools **26.1.0** (builder, lint, navigator, molecule CalVer) [official].
- 2026-03-12: ciscodevnet/ansible-nd **1.5.0** (backup/restore, ND 4.x) [official].
- 2026-03-12: ncclient **0.7.0** (libssh transport, call-home, YANG 1.1 actions) [secondary].
- 2026-03-25: AAP **2.5** component refresh (controller 4.6.27, hub 4.10.13, EDA 1.1.17, Receptor 1.6.4) [official].
- 2026-04-16: NAPALM org repos updated (project active) [secondary].
- 2026-05: OpenTofu **1.12.0**; community stacks on ansible 13.x / ansible-core 2.20 [secondary].
- 2026-05-31: ansible-core **2.21** GA (Python 3.12–3.14 control) [official].
- 2026-06-01: ciscodevnet/ansible-gnmi **v3.0.0** (breaking: Cisco-only tested profiles) [official].
- 2026-06-13: scrapli **2026.6.13rc14** pre-release (Python 3.10+) [independent].
- 2026-06-14: nornir_salt **0.23.3** [independent].
- 2026-06-22: cisco.ios **11.4.2** (bugfix); catalyst-center-ansible-iac **v1.0.0** GA [official].
- 2026-06: OpenTofu comparisons note provider lag 6–12 weeks; 22%/27% faster plan/apply claim vs Terraform 1.9 [secondary].
- 2026-07-03: infrahub-mcp **1.1.7** [secondary].
- 2026-07-23: cisco.ios **11.5.0** + cisco.iosxr **12.4.0** (deprecated-import cleanup) [official].
- 2026-07-30: nornir-infrahub **1.2.0** (idempotent file transfer, py3.10–3.14) [official].
- 2026-08: Infrahub **1.11.0** (schema contracts); smana ADR adopts OpenTofu 1.12.6 [secondary].
- 2026-08-19: OpenTofu **1.12.6** stable [independent].
- 2026-09-08: ansible-core **2.20.9** + **2.21.4** (security) [official].
- 2026-09-10: SUSE fixes CVE-2026-16493 in ansible-core [secondary].
- 2026-09-22: arista.eos **12.1.2**, junipernetworks.junos **11.1.1**, Netmiko **4.8.0**, Nornir **3.6.0**, nornir-utils **0.3.0** current [official/independent].
- Planned: ansible-core **2.22** GA Nov 2026; arista.eos `src` auto-templating removal Mar 2028; junipernetworks.junos redirect removal 2028-04-01 (or Ansible 14) [official].
- AAP 2.5 (Mar 2026) was last documented minor in the 2.5 line at research time; AAP 2.6 (Oct 2025) last RPM-installable; 2.7+ container-only [official/secondary].

---

