---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/overview
title: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2025-11-03", "2026-01-08", "2026-03-12", "2026-05-31", "2026-06-01", "2026-06-22", "2026-07", "2026-07-23", "2026-09-08", "2026-09-10", "2026-09-22", "2026-11-30", "2027-05", "2027-05-31", "2027-11-30", "2028-05"]
keywords: ["benchmarks", "compute", "gpu", "research"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [1, 38]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: ba87eea400665651ecfc0ca28423f42eddcb99bdeb173ed2622aeae3e2aca786
---

# Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries

## Wave 0 — Scope, method, provenance

- **Scope:** Network automation frameworks as of 2026-09-22: Ansible (core, community package, network collections, AWX/AAP, Event-Driven Ansible, ansible-navigator), Nornir ecosystem, Terraform/OpenTofu and network providers, Python network libraries (Netmiko, NAPALM, Scrapli, ncclient, pygnmi, pyATS/Genie), and cross-cutting patterns (idempotency, dry-run, Jinja2 templating, secrets handling, inventory from source of truth).
- **Date:** 2026-09-22. **Method:** read-only web research (browser_search, browser_open on search-result URLs); no live-browser visits; nothing sent externally. No identifiers guessed; every URL below is verbatim from search results.
- **Provenance legend:** `[official]` = vendor/project documentation or release notes; `[vendor-reported]` = vendor blog/marketing claims; `[independent]` = third-party reviews, benchmarks, registries; `[secondary]` = community docs, skills repos, forum posts, GitHub READMEs not from the project itself; `[unverified]` = single-source or unconfirmable claims. Gaps, conflicts and unverified items are registered in Wave 8.

---

## Wave 1 — ansible-core releases & lifecycle (2026)

- ansible-core **2.21** was released **2026-05-31**; latest patch **2.21.4** released **2026-09-08**; support runs through **2027-11-30**. Control node Python 3.12–3.14; managed node Python 3.9–3.14; managed PowerShell 5.1–7 [official — https://docs.ansible.com/projects/ansible-core/devel/reference_appendices/release_and_maintenance.html; http://eosl.date/eol/product/ansible-core/].
- ansible-core **2.20** (released 2025-11-03) latest **2.20.9** (2026-09-08), support through **2027-05-31**; **2.19** latest 2.19.12, End of Support **2026-11-30** (plan upgrade path now); **2.18** reached EOL **2026-05-31** and receives no more security updates [official — http://eosl.date/eol/product/ansible-core/; https://github.com/endoflife-date/endoflife.date/blob/HEAD/products/ansible-core.md].
- Upstream roadmap: **2.22** targets GA **Nov 2026** (Critical May 2027, Security Nov 2027, EOL May 2028), control node Python 3.13–3.15; **2.24** (Nov 2027) and **2.25** planned with PowerShell 7.8 LTS overlap; next multi-LTS after that is 2.28 [official — https://github.com/ansible/ansible-documentation/blob/HEAD/docs/docsite/rst/reference_appendices/release_and_maintenance.rst].
- Community `ansible` package: **13.x** depends on ansible-core **2.20**; versions 13.0/13.1.0 observed in the field (e.g., Homebrew installs python 3.14 with jinja 3.1.6, pyyaml 6.0.3) [secondary — https://github.com/iuliandita/skills/blob/HEAD/skills/ansible/SKILL.md; GitHub issue reports].
- ansible-dev-tools **v26.1.0** released **2026-01-08**: bundles ansible-builder, ansible-core, ansible-creator, ansible-dev-environment, ansible-lint, ansible-navigator (TUI), ansible-sign, molecule, pytest-ansible [official — https://forum.ansible.com/t/release-announcement-ansible-dev-tools-v26-1-0/45073].
- Tooling moved to CalVer in 2026: molecule **26.x**, ansible-lint **26.x**, ansible-navigator **26.x**, ansible-builder **3.1.x** (Execution Environment definition v3) [secondary — https://github.com/iuliandita/skills/blob/HEAD/skills/ansible/SKILL.md].
- **Security (Sept 2026):** CVE-2026-11332 — ansible-galaxy role installs could allow arbitrary git configuration injection via malformed role requirements; fixed in RHEL-packaged ansible-core 2.16.3-4 (RHEL-194126) [official — AlmaLinux/RHEL changelog via https://git.almalinux.org]. CVE-2026-16493 — argument injection in `ansible-galaxy collection install` via git clone; fixed by SUSE update SUSE-SU-2026:23704-1 (2026-09-10), ansible-core-2.18.3-160000.4.1 [secondary — https://linuxsecurity.com/advisories/suse/suse-2026-23704-1-important-for-ansible-core].
- RHEL-packaged ansible-core 2.16.3 ships with telemetry additions for RHEL (RHEL-219532, July 2026) [official — AlmaLinux git changelog].

---

## Wave 2 — Ansible network collections (2026 versions)

### Cisco collections

- **cisco.ios 11.5.0** released **2026-07-23**: cleaned up deprecated ansible-core imports and `warnings` in `exit_json` (via `emit_warnings`/`AnsibleModule.warn()`) for newer ansible-core compatibility; fixed `ios_acls` multi-word option handling; **11.4.2** released 2026-06-22 (IPv6 ACL remark CLI syntax bugfixes, ios_user SSH key matching) [official — https://forum.ansible.com/t/cisco-ios-11-5-0-release/46076; http://forum.ansible.com/t/release-announcement-cisco-ios-11-4-2/45971; https://github.com/ansible-collections/cisco.ios/blob/HEAD/CHANGELOG.rst].
- cisco.ios **11.0.0** bumped `ansible.netcommon` dependency to **>=8.1.0**; last version compatible with ansible-core ≤2.18 is cisco.ios **10.1.1** [official — cisco.ios CHANGELOG]. 11.1.0 added ios_config prompt-answering in config mode and chassis_id in ansible_net_neighbors LLDP dict; 11.2.x fixed ios_vrf_global idempotency for multiple import/export route-targets and ios_hsrp_interfaces parsing [official — cisco.ios CHANGELOG].
- **cisco.iosxr 12.4.0** released **2026-07-23**: deprecated-import cleanup across modules for ansible-core 2.23, `bgp_global` neighbor shutdown parsing/idempotency fix (`no shutdown` handling) [official — https://forum.ansible.com/t/cisco-ios-xr-12-4-0-release/46077; https://github.com/ansible-collections/cisco.iosxr/blob/HEAD/CHANGELOG.rst].
- **cisco.nxos**: field bug reports reference version **12.0.0** (June–July 2026) — e.g., a BGP address-family VRF attribute leak issue on 12.0.0 [secondary — https://github.com/ansible-collections/cisco.nxos/issues/1082]. CHANGELOG HEAD shows the 9.x line (9.2.x with `nxos_vrf_global`, `nxos_vrf_address_family` resource modules) — **the 9.x vs 12.x lineage is a registered gap, see Wave 8** [official — https://github.com/ansible-collections/cisco.nxos/blob/HEAD/CHANGELOG.rst].
- **cisco.aci 2.11.0**, **cisco.meraki 2.20.10**, **cisco.dnac 6.31.3**, **cisco.asa 6.1.0**, **cisco.ise 2.10.0**, **cisco.intersight 2.0.20** observed in a May-2026 community stack listing; **cisco.intersight 2.19.0** adds OAuth2 Bearer-token auth, AI storage/GPU/confidential-compute/BIOS tuning policies [secondary — juniper.device issue env dump; official — https://github.com/dsoper2/ansible-intersight/blob/HEAD/CHANGELOG.md].
- **ciscodevnet/ansible-nd 1.5.0** (2026-03-12): backup/restore options, ND 4.x support [official — https://github.com/ciscodevnet/ansible-nd/blob/HEAD/CHANGELOG.rst].
- **cisco-en-programmability/catalyst-center-ansible-iac v1.0.0** GA **2026-06-22**: Cisco-validated playbooks for Catalyst Center Day 0/1/2/N (users & roles, ISE/AAA integration, site hierarchy, LAN automation, SD-Access fabric sites/zones/transits, virtual networks, L2/L3 gateways) [official — https://github.com/cisco-en-programmability/catalyst-center-ansible-iac/blob/HEAD/CHANGELOG.rst].
- **ciscodevnet/ansible-gnmi v3.0.0** (2026-06-01): BREAKING — removed nokia_sros and arista_eos placeholder platform profiles; tested platform support is Cisco IOS XE / IOS XR / NX-OS; non-Cisco devices use platform='auto' [official — https://github.com/ciscodevnet/ansible-gnmi/commit/d759860f574ec9d9bd75b7].

### Arista, Juniper

