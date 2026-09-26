---
id: etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-5-terraform-opentofu-for-networks
title: "Wave 5 — Terraform / OpenTofu for networks"
domain: front-matter
role: reference
task: reference
actors: ["AWS"]
dates: ["2025-04", "2026-05", "2026-06", "2026-06-12", "2026-06-14", "2026-07-30", "2026-08-19", "2026-08-21", "2026-08-30", "2026-09-22"]
keywords: ["aws", "benchmark", "sandbox"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [70, 95]
section: "Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries"
sha256: a1532b630ba6272817c5ebacd6ba02826cde4ab30decbfe579f3f39d068c5761
---

# Wave 5 — Terraform / OpenTofu for networks

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

