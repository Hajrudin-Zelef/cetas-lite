---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/meta-utility-collections
title: "Meta/utility collections"
domain: front-matter
role: reference
task: actor-profile
actors: ["Meta"]
dates: ["2024-07", "2026-03", "2026-03-25", "2026-05", "2028-03", "2028-04-01"]
keywords: ["aws", "guardrails", "memory"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [39, 69]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: 6e37092dcfa809867f854dc111c1329505976803660f90ae19fe4e8638fc5d7a
---

# Meta/utility collections

- **arista.eos v12.1.2** (Sept 2026): v12.1.1 standardized action plugin naming with the `eos_` prefix (runtime routing for backward compat), removed deprecated action plugins for deleted modules, bumped `ansible.netcommon` to **>=8.5.2**; v12.1.0 added `content` parameter to `eos_config` (pre-rendered templates); `src` auto-Jinja2-processing deprecated, removal planned **March 2028** [official — https://github.com/ansible-collections/arista.eos/blob/HEAD/CHANGELOG.rst].
- **junipernetworks.junos**: repo moved to **Juniper/ansible-junos-stdlib**; **v11.1.1** current (Sept 2026); v11.0.0 requires ansible.netcommon >=8.1.0 and ansible-core >=2.16 [official — https://github.com/juniper/ansible-junos-stdlib/blob/HEAD/ansible_collections/junipernetworks/junos/CHANGELOG.rst].
- Migration path: `junipernetworks.junos` → **`juniper.device`** namespace (e.g., `juniper.device.junos_config`); redirects emit deprecation warnings, removal after **2028-04-01** (v11.1.0 extended the window ~2 years) [official — juniper changelog]. juniper.device observed at **2.0.1** (Jan 2026 field report) [secondary].
- **Deprecation warning:** the Ansible 11 porting guide states the `junipernetworks.junos` collection **has been deprecated** and will be **removed from Ansible 14** if no one resumes maintenance [official — https://github.com/ansible/ansible-documentation/blob/HEAD/docs/docsite/rst/porting_guides/porting_guide_11.rst].
- A Galaxy release-process issue noted junipernetworks.junos **11.1.0 was released to Galaxy but not git-tagged** (repo-management requirement violation; risk of removal from the community package) [secondary — https://github.com/ansible-collections/junipernetworks.junos/issues/589].

### Meta/utility collections

- **ansible.netcommon** is the shared transport/facts layer (network_cli, netconf, httpapi connections); 8.x line current (8.2.0, 8.5.2 required by arista.eos 12.1.x); **ansible.utils** and **ansible.network** (meta collection) complete the stack [official changelogs; secondary field reports].
- **community.network 5.1.0** observed in May 2026 stacks [secondary].
- Community stack snapshot (May 2026, one user's listing): amazon.aws 9.4.0, ansible.posix 1.6.2, ansible.utils 5.1.2, arista.eos 10.1.1, awx.awx 24.6.1, check_point.mgmt 6.4.0, community.general 10.6.0, community.hashi_vault 6.2.0, cisco.ucs 1.16.0, cisco.mso 2.10.0, community.ciscosmb 1.0.10 — useful as a real-world pinned-set example [secondary — juniper.device issue env dump].

---

## Wave 3 — AWX, Ansible Automation Platform & Event-Driven Ansible

- **AWX 24.6.1** (July 2024) was the last formal upstream release; upstream AWX releases paused for a **major refactor** while the devel branch stays active; awx-operator ~2.12.x still ships for Kubernetes deploys [secondary — https://github.com/iuliandita/skills/blob/HEAD/skills/ansible/SKILL.md].
- Community builds track newer state: fitbeard/automation-platform builds **AWX 26.0.0** images (Sept 2026) with gateway 2.6, EDA server 1.2.12, awx-operator 2.6 [secondary — https://github.com/fitbeard/automation-platform/blob/HEAD/README.md].
- **Ansible Automation Platform 2.6** (Oct 2025) was the **last RPM-installable** release; **AAP 2.7+ is containerized-only** [secondary — skills doc]. AAP **2.5** (March 25, 2026 refresh) components: automation controller **4.6.27**, automation hub **4.10.13**, Event-Driven Ansible **1.1.17**, Receptor **1.6.4** [official — https://docs.redhat.com/en/documentation/red_hat_ansible_automation_platform/2.5/html-single/release_notes/release_notes].
- AAP 2.5 release notes (March 2026) fixed job-fact race conditions (AAP-69263), cancel propagation to dependent workflow jobs (AAP-68974), project sync deletion races (AAP-71407), Jinja2 errors in project_update.yml with newer ansible-core (AAP-68784) [official — Red Hat docs].
- **Event-Driven Ansible (EDA)**: GA since AAP 2.4; EDA controller orchestrates rulebooks across event sources (monitoring/observability tools); actions via `run_job_template` or embedded playbooks; event throttling with `once_within` (reactive) and `once_after` (passive) conditions plus default throttling guardrails [official — https://www.redhat.com/es/blog/event-driven-ansible-is-here].
- EDA edge case: **event storms** from flapping alerts require debounce logic in rulebook conditions to avoid runbook flooding [secondary — https://github.com/azzindani/kea/blob/HEAD/knowledge/skills/ansible-automator/SKILL.md].
- **ansible-navigator** (26.x CalVer): TUI for running/inspecting playbooks inside Execution Environments [secondary]. Execution Environments are built with **ansible-builder 3.1.x** (EE definition v3); community pattern: Ansible via `pipx`, Python libraries isolated in `/opt/venv` inside the image; pin EE base images by digest, never `latest` [secondary — https://github.com/andersonmavi30/docker_network_automation; skills docs].
- Large-inventory guidance: enable fact caching (Redis/JSON) and `serial` batching for 10k+ hosts to avoid controller memory exhaustion [secondary — skills docs].
- A community OpenAPI spec effort curates **277 of 631** upstream AWX/AAP v2 API operations for common CRUD (inventories, job templates, workflows, credentials, EEs, schedules) [secondary — https://github.com/itential/assets/blob/HEAD/Ansible/AWX%20(AAP)/README.md].
- Red Hat maintains **Ansible network validated content** (`network.base`, `network.interfaces` collections): platform-agnostic roles (Resource Manager, build-brownfield-inventory, gather/persist/deploy/configure) covering Arista EOS, Cisco IOS-XR/IOS/IOS-XE, Junos, NX-OS [vendor-reported — https://www.redhat.com/pt-br/blog/accelerating-your-network-automation-journey-ansible-network-validated-content].

---

## Wave 4 — Nornir ecosystem

