---
id: etape6-phasee4-validation-cicd/00-front-matter/e4-7-testing-frameworks
title: "E4.7 — Testing frameworks"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2026-05", "2026-09"]
keywords: ["apache", "license", "throughput"]
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [245, 310]
section: "Phase E4 — Network Validation, Observability & CI/CD"
sha256: 35668bdd8a216bd480e70d002d7681916665bcc383da7d6d6ae98e0d77814c6e
---

# E4.7 — Testing frameworks

## E4.7 — Testing frameworks

### E4.7.1 pyATS (Cisco)

- pyATS (Python Automated Test Systems): Cisco's Python test automation framework, originally internal engineering tooling, now public under Apache License 2.0 `[official]` — https://developer.cisco.com/docs/pyats/pyats-on-devnet/.
- Scale claim: "de-facto test framework for Cisco engineers... running millions of CI/CD, sanity, regression, scale, HA, and solution tests monthly"; used by thousands of network engineers/developers worldwide `[official]` (Cisco DevNet).
- Standardizes: topology definition/modeling, programmatic device interaction (connection libraries), test-script definition/execution, test runs and reporting `[official]` (DevNet).
- Components: core framework, Genie libraries (parsers, Ops, APIs, triggers & verifications), Unicon (connection), Clean (device cleanup/recovery), Easypy (runner), aetest, robot integration (`pyats.robot` / `genie.libs.robot`) `[official]` — package list from Cisco Live 2026 lab (see below).
- Version in use 2026: `pyats[full]==26.5` installed via pip in Cisco Live 2026 NX-OS DevOps lab (LTRDCN-3903) with genie 26.5, Python 3.11.14 `[official]` — https://nxos-devops.ciscolive.com/lab/pod26/pyats/setup.
- Release cadence: monthly-ish minor releases documented on DevNet (24.7, 24.8 examples: dozens of new parsers/APIs per release; Cat8200/Cat8300/Cat9800-CL/IR1800 Clean support; SONiC ops support) `[official]` — https://developer.cisco.com/docs/pyats/24-7/ and https://developer.cisco.com/docs/pyats/24-8/.
- Multi-vendor note: DevNet describes "ready-to-use, multi-vendor device libraries (e.g., parsers, APIs)" `[official]`; in practice strongest on Cisco platforms, with community/third-party parser extensions `[unverified]`.
- Community tooling: pyATS web app (FastAPI + React, testbed management, Genie Learn snapshots for 15+ features, config diff) exists as third-party project `[secondary]` — https://github.com/stormbliss/pyats-web-app.
- Genie Learn: operational-state snapshots (config, BGP, OSPF, interfaces) used for drift/change validation `[secondary]` (pyats-web-app).

### E4.7.2 Robot Framework

- Robot Framework: generic keyword-driven acceptance-test framework, Python, Apache 2.0; originated at Nokia Networks (2005), open-sourced 2008 `[secondary]` — http://en.wikipedia.org/wiki/Robot_Framework.
- Stable release 7.5 on 14 September 2026 `[secondary]` (Wikipedia, crawled ~Sept 2026).
- Network use: Cisco Live 2026 session BRKATO-1009/CTF-2002 (GitHub `sandipcbr/brkato-1009-2026`) runs Robot Framework + pyATS side by side for DC fabric testing: `bgp_neighbor_health.robot` (neighbor state, prefix counts, uptime), `bgp_failover_convergence.robot` (link-failure trigger, convergence timing, reroute validation), `interface_validation_restconf.robot` (RESTCONF interface state/MTU/speed), BDD-style Gherkin failover tests `[secondary]` — https://github.com/sandipcbr/brkato-1009-2026.
- Security practice in that repo: credentials via environment variables, never hardcoded; HTTPS/NETCONF recommended; isolated test environment `[secondary]` (same repo).
- CI/CD: example GitHub Actions workflow runs `robot tests/` with secrets-injected credentials `[secondary]` (same repo).
- Ecosystem synergy: `netascode/nac-test` 2.0.0 adds experimental pyATS integration alongside Robot tests, with `--pyats`/`--robot` flags, combined HTML dashboard, merged `xunit.xml` for CI/CD, diagnostic collection flag, fail-fast auth validation; note: pyATS tests unavailable on Windows (WSL2 recommended), macOS needs Python 3.12+ `[secondary]` — https://github.com/netascode/nac-test/blob/HEAD/CHANGELOG.md.
- Robot Framework Browser library 20.4.0 (Aug 19, 2026): Playwright-based web testing; requires RF 7.1.1+, Python 3.10+ `[secondary]` — https://github.com/marketsquare/robotframework-browser/blob/HEAD/docs/releasenotes/Browser-20.4.0.md (web-testing adjacent, relevant for dashboard/portal validation).

### E4.7.3 pytest and Python-native patterns

- pytest is the de-facto runner for Python network test suites (pyATS aetest scripts, Nornir-based checks, custom Netmiko/Scrapli assertions) `[secondary]` — evidenced by Cisco Live 2026 repo running `python -m pytest tests/bgp_neighbor_health.py -v` alongside Robot `[secondary]` (brkato-1009-2026).
- nac-test 2.0.0: `SSHTestBase.parse_output()` became async (breaking change) — reflects async SSH testing patterns `[secondary]` (nac-test CHANGELOG).
- xUnit/JUnit XML output (`xunit.xml`) is the interchange format feeding CI dashboards (GitHub Actions, GitLab, Jenkins) `[secondary]` (nac-test CHANGELOG).

### E4.7.4 Synthetic monitoring (context)

- Synthetic monitoring (active probes: ping, traceroute, TWAMP, iPerf-style throughput, synthetic transactions) complements config validation and observability; vendor implementations live in ThousandEyes (Cisco), Arista DANZ/CloudVision, Juniper Paragon — deep-dive out of scope for E4, noted here as the third validation pillar alongside pre-change verification and post-change observability `[unverified]` (no 2026 primary sources pulled; flag for follow-up).

---

## E4.8 — CI/CD and GitOps for networks

### E4.8.1 Canonical pipeline pattern

- Plan-on-PR / apply-on-merge: `plan` runs read-only on every pull request; `apply` only after merge to a long-lived environment branch; plan output retained as artifact; plan summary posted to the PR; CODEOWNERS-gated reviewers per environment; manual approval gate before production apply; policy-as-code (Sentinel/OPA/Conftest) runs before apply `[secondary]` — https://github.com/timstewart-dynatrace/best-practice-notebooks (AUTOM-07 CI/CD integration recipe, 2026 reference).
- Secrets: three places to handle — pipeline storage, in flight, at rest in state; short-lived credentials pulled at pipeline start (e.g., Vault integration) `[secondary]` (same source).
- Drift detection: scheduled `plan` runs reporting diffs to a channel `[secondary]` (same source).

### E4.8.2 GitHub Actions / GitLab CI / Jenkins in network automation

- GitHub Actions is the most-cited CI runner in 2026 network-automation examples: Containerlab labs in Actions, Robot/pyATS suites in Actions, nac-test xUnit ingestion `[secondary]` (multiple sources above).
- Reference end-to-end pattern (2026): push → Actions (lint/unit, SAST, build, SBOM, sign) → config repo update → GitOps sync (ArgoCD-style) → canary/analysis → observability/DORA metrics `[secondary]` — https://github.com/austinxyz/job-preparation/blob/HEAD/skills/tech/infra/CI-CD%20Pipeline%20Engineering.md.
- Network-specific adaptation: the "deploy" step is a config push to devices (via Ansible/Netmiko/NAPALM/Terraform providers) gated by pre-change validation (Batfish) and followed by post-change observability checks (SuzieQ) — this pre/post pairing is explicitly recommended in the NetBox Labs × SuzieQ integration narrative `[official]` — https://netboxlabs.com/blog/netbox-labs-stardust-systems-network-observability-partnership/.
- GitLab CI and Jenkins remain in use (GitLab equivalent pipelines documented in AUTOM-07); no 2026 share data found `[unverified]`.

### E4.8.3 Change management workflows

- Branching: feature branches for config changes, environment branches (dev/staging/prod), required reviewers, branch protection `[secondary]` (AUTOM-07 recipe).
- Pre-change validation gates: Batfish differential analysis (reachability identical current vs planned; ACL changes provably collateral-free) `[official]` (Batfish README); Forward Predict deterministic pre-change verification `[vendor-reported]` (Computer Weekly, May 2026).
- Post-change verification: SuzieQ 90-second polling for drift `[vendor-reported]`; IP Fabric intent checks `[vendor-reported]`; pyATS/Genie Learn snapshots diffed for change validation `[secondary]`.
- Approval/audit: plan artifacts retained for forensics; signed pipelines (cosign/SLSA) in mature shops `[secondary]` (AUTOM-07).

### E4.8.4 Config diff testing

- Genie Learn snapshots + diff: operational-state capture before/after, diffed for drift/change validation `[secondary]` (pyats-web-app).
- Batfish change analysis questions: end-to-end reachability identical across configs; planned ACL/firewall changes provably correct `[official]` (Batfish README).
- `netlab` validation tests with `wait` parameter and color-coded output (ipspace netlab 1.7.1) show the lab-as-CI pattern maturing `[secondary]` — https://github.com/ipspace/netlab/blob/HEAD/docs/release/1.7.md.

---

