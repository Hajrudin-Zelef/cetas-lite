---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible
title: "Wave 3 — AWX, Ansible Automation Platform & Event-Driven Ansible"
domain: front-matter
role: reference
task: reference
actors: ["AWS"]
dates: ["2024-07", "2025-04", "2026-03", "2026-03-25", "2026-05", "2026-06", "2026-06-12", "2026-06-14", "2026-07-30", "2026-08-19", "2026-08-21", "2026-08-30", "2026-09-22"]
keywords: ["aws", "benchmark", "guardrails", "memory", "sandbox"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [53, 107]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: 1281ee9873a12294a60c370ab2c535986af74a10185ce94e7cb0a6775d75ca82
---

# Wave 3 — AWX, Ansible Automation Platform & Event-Driven Ansible

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

- **Nornir 3.6.0** current on PyPI (Sept 2026): pure-Python, multi-threaded, inventory + task dispatch; **requires Python 3.10+**; since 3.0.0 **no bundled plugins** — installed separately via pip [independent — https://libraries.io/pypi/nornir].
- **nornir-utils v0.3.0** (Sept 2026): Python ≥3.10, migrated from poetry to **uv**, pylama/black replaced with **ruff**, tests on Python 3.10–3.14 [official — https://github.com/nornir-automation/nornir-utils/blob/HEAD/CHANGELOG.md].
- **nornir_napalm** (0.5.0 in 2026 Cisco Press material): wraps NAPALM drivers inside Nornir tasks — initialize inventory → define task calling `napalm_cli`/`napalm_configure` → run across devices in parallel → structured dict results; supports IOS/IOS-XE/Junos/EOS via NAPALM drivers [secondary — 2026 Cisco Press "Hands-On Cisco Automation with Python", Graziani/Iliesiu].
- **nornir_scrapli**: scrapli's official Nornir plugin; key features: fully-typed, unit + functional tests against virtual routers, optional `ssh2` transport for speed; repo active (updated 8 days before 2026-09-22) [official — https://github.com/scrapli/nornir_scrapli].
- **nornir_netmiko**: plugin wiring Netmiko connections into Nornir tasks (documented in the same 2026 Cisco Press book alongside nornir_napalm) [secondary].
- **nornir-infrahub v1.2.0** (2026-07-30): Nornir inventory plugin backed by Infrahub; depends on `nornir>=3.5.0,<4`, `infrahub-sdk>=1.20.1,<2`, pydantic v2; added idempotent file upload/download task plugins (SHA-1 checksum), hostname schema mapping, Python 3.10–3.14 support; dropped Python 3.9 [official — https://github.com/opsmill/nornir-infrahub/blob/HEAD/CHANGELOG.md].
- **nornir_salt 0.23.3** (2026-06-14): Nornir plugins for SaltStack Salt-Nornir Proxy Minion and **NorFab** [independent — https://pypi.org/project/nornir_salt/0.23.3/].
- **NornFlow** (pre-release, 0.9.0 on PyPI): workflow orchestration on top of Nornir — declarative YAML interface, CLI, multi-level variable precedence (environment/global/domain/workflow/CLI/runtime), Jinja2 filters, hooks, failure strategies (skip-failed, fail-fast, run-all), standardized project layout [independent — https://pypi.org/project/nornflow/0.9.0/].
- **nornir-srl 0.2.1**: containerized `fcli` CLI for Nokia SR Linux via Containerlab topologies — Nornir tasks for bgp-peers, bgp-rib, lldp, mac table, sub-interfaces [independent — https://pypi.org/project/nornir-srl/0.2.1/].
- Community adoption signal: a 2026 "universal NetDevOps" Docker image (v3.0.0, Rocky Linux 10 base) ships Nornir alongside Netmiko/Scrapli/NAPALM as the orchestration layer, with Ansible via pipx for declarative playbooks [secondary — https://github.com/andersonmavi30/docker_network_automation].

---

## Wave 5 — Terraform / OpenTofu for networks

### OpenTofu 2026 status

- **OpenTofu 1.12.6** released **2026-08-19** (stable); 1.12.2 (2026-06-12), 1.12.0 (May 2026), 1.8 (early 2026) preceded it; MPL-2.0 licensed, Linux Foundation project, **CNCF Sandbox since April 2025** [independent — https://en.wikipedia.org/wiki/OpenTofu; https://scalr.com/learning-center/what-is-opentofu].
- Fork history: HashiCorp relicensed Terraform MPL-2.0 → **BUSL 1.1** (Aug 2023); OpenTF fork started, renamed **OpenTofu**, Linux Foundation (Sept 2023); first stable **1.6.0** (Jan 2024), compatible with Terraform 1.5.x [independent — Wikipedia; Scalr].
- Divergent features (June 2026 comparison): native **state encryption** (1.7+, passphrase/KMS via AWS/GCP KMS/OpenBao), **provider-defined functions** (1.7), **loopable imports** (`for_each` in import blocks, 1.7), **`removed` blocks**, `-exclude` flag (1.9), **provider `for_each`** (1.9), OCI registry for modules/providers (1.10+), native **S3 state locking without DynamoDB**, experimental OpenTelemetry tracing, **ephemeral values** (secrets never touch state, 1.11), **dynamic `prevent_destroy`** (1.12) [secondary — https://medium.com/@satyajitdas0033/...terraform-vs-opentofu; https://scalr.com/learning-center/what-is-opentofu].
- Provider ecosystem: OpenTofu Registry mirrors Terraform Registry — **~3,900+ providers, ~23,600+ modules** claimed by one June-2026 comparison (other sources cite 3,000+ / 3,200+ providers — **conflict registered, see Wave 8**) [secondary].
- Operational caveats: providers lag Terraform by **6–12 weeks** (AWS provider 4 releases behind in Jan 2026, per one practitioner's benchmark); no official Terraform Cloud/Enterprise equivalent (Spacelift/env0/Scalr used instead); some modules using `terraform`-specific data sources can break silently on provider signature changes; one 2026 benchmark claims 22% faster plan / 27% faster apply vs Terraform 1.9 [secondary — http://kubaik.github.io/pulumi-vs-terraform-vs-opentofu-in-2026/].
- Terraform side: **1.9.x BSL** current line; Terraform 1.9.0 described as "last open-source version under MPL" by one source (2026) — **flag as unverified**; HCP Terraform/Terraform Enterprise backends work with Terraform only; HashiCorp acquired by **IBM (2025)** [secondary — Medium comparisons].
- Practitioner adoption: platform teams pin `opentofu = "1.12.6"` in mise.toml with **no terraform entry**, HCL unchanged (only the binary pin) — e.g., smana/cloud-native-ref ADR (2026-08-21, verified 2026-08-30) [secondary — https://github.com/smana/cloud-native-ref/blob/HEAD/website/content/docs/decisions/0014-opentofu-over-terraform.md].
### Network providers

- **Juniper Apstra Terraform provider** (official): configures Apstra intent-based DC fabrics via the Apstra Go SDK APIs; pushes HCL-defined intent through Apstra to multivendor devices (Juniper, Cisco, Dell, Arista); supports Day 0–Day 2+ provisioning, GitOps/CI-CD patterns [official — https://Www.juniper.net/content/dam/www/assets/solution-briefs/us/en/network-automation/configuring-apstra-through-terraform.pdf].
- **NetBox provider `e-breuninger/netbox` v4.1.0**: manages physical/logical network inventory (devices, IPs, prefixes, VLANs, cables) from Terraform — "documentation-first" pattern where NetBox is the queryable source of truth written by Terraform; used with Terraform Cloud VCS-driven workflows + GitHub Actions CI [secondary — https://github.com/simonpainter/www.simonpainter.com/blob/HEAD/blog/netbox-terraform.md; https://github.com/xiiisins/homelab/blob/HEAD/docs/services/netbox.md].
- **F5 `F5Networks/terraform-provider-f5os`**: v1.12.0 introduced breaking Go 1.25 minimum, `f5os_lag` LACP/STATIC modes, 75%+ coverage enforcement, govulncheck; handles F5OS device config as code [official — https://github.com/F5Networks/terraform-provider-f5os/releases].
- **Palo Alto `PaloAltoNetworks/panos`** (v1.11.0 in community tutorials): firewall interfaces, zones, address objects, security policies via API-key auth [secondary — https://www.packetswitch.co.uk/palo-alto-automation-with-terraform/].
- **Infoblox Terraform provider v2.5/v2.6** (2024): NIOS DDI resources (networks, A records, DNS zones), extensible-attribute filtering, next-available-network allocation; enables IPAM-as-code (IP allocated at apply time, DNS records created/removed with resources) [secondary/vendor-reported — https://ddi.mohflo.net/index.php/2024/06/11/whats-new-in-infoblox-terraform-provider-v2-5-guiding-to-infoblox-terraform-provider-v2-5-new-features/; http://www.infoblox.com/resources/solution-notes/optimize-your-hybrid-multi-cloud-infrastructure-with-infrahub-plug-in-for-terraform].
- Registry-notable vendor providers (Terraform Registry): **CiscoDevNet/aci** (Cisco ACI), **Juniper/mist** (Mist cloud), **fortinetdev/fortimanager** (FortiManager), **Cisco Meraki** provider — commonly combined under one Terraform layer for LAN/SD-WAN/DC [secondary — https://github.com/simonpainter/www.simonpainter.com/blob/HEAD/blog/bringing-it-all-together.md].
- Community reference architecture: one Terraform layer + multiple vendor providers, version-controlled in Git, approved PRs auto-applying to environments — network infra "always in a known consistent state" [secondary — simonpainter blog].
- **State management for networks:** remote state backends (Terraform Cloud, S3 with native locking in OpenTofu 1.10+), per-environment state separation, state encryption (OpenTofu 1.7+) for credentials in state; drift detection via `tofu plan -refresh-only` [secondary — OpenTofu docs/comparisons].
- **GitOps workflows:** GitHub Actions CI running `terraform fmt + init + validate`, tflint, checkov, terraform-docs; demo-mode `null_resource` with triggers for safe planning without touching devices; PR-driven apply pipelines [secondary — https://github.com/joycemwangi/hybrid-cloud-infrastructure-as-code-with-terraform].

---

