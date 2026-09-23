---
id: etape6-phasee4-validation-cicd/00-front-matter/e4-9-event-driven-automation
title: "E4.9 — Event-driven automation"
domain: front-matter
role: reference
task: reference
actors: ["AWS"]
dates: ["2019-10"]
keywords: ["agentic", "apache", "aws", "consumer", "mcp", "open source", "research"]
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [311, 375]
section: "Phase E4 — Network Validation, Observability & CI/CD"
sha256: 73eacb2b82aa13883b52e1c952d003db8f7df62bc57eb937e4c29bf01f9f2199
---

# E4.9 — Event-driven automation

## E4.9 — Event-driven automation

### E4.9.1 StackStorm

- StackStorm ("IFTTT for Ops"): open-source event-driven automation — rules engine, workflow engine, 160+ integration packs with 6000+ actions, ChatOps; installer at docs.stackstorm.com `[official]` — https://github.com/StackStorm/st2.
- History: created 2013 as a DevOps startup; acquired by Brocade (2016); Extreme Networks acquired Brocade's data-center assets including StackStorm (2017); Extreme transitioned StackStorm to the Linux Foundation (October 2019) `[secondary]` — https://www.sdxcentral.com/news/extreme-gifts-stackstorm-to-linux-foundation/.
- 2026 status: repository active (crawled ~22 days before research); community metrics cited: ~6,446 GitHub stars, 781 forks, Apache-2.0, Python `[secondary]` — https://www.opensourcealternative.to/project/StackStorm. Last-commit recency suggests maintenance mode rather than rapid feature development `[unverified]` (single third-party snapshot).
- Typical network use: alert → automated remediation workflow; common apps include automated remediation, continuous deployment, ChatOps, automated security response; users include tier-one service providers, retailers, media companies `[secondary]` (SDxCentral, Extreme press).
- Extreme Workflow Composer: commercial platform powered by StackStorm `[secondary]` (Extreme/LF press release PDF).

### E4.9.2 Event-Driven Ansible (Red Hat)

- Ansible EDA: rulebooks (YAML) bind event sources → conditions → actions (`run_playbook`, `run_job_template`, `run_workflow_template`); `ansible-rulebook` CLI; EDA Controller in Ansible Automation Platform (AAP) for enterprise scale `[secondary]` — https://github.com/cfoxrhel/ansible-skills/blob/HEAD/skills/ansible-eda-rulebook/SKILL.md.
- Event sources in `ansible.eda` collection (v2.11.0): Alertmanager, AWS CloudTrail, AWS SQS, Azure Service Bus, file_watch, generic, journald, Kafka, pg_listener (PostgreSQL NOTIFY), range, tick, url_check, webhook `[secondary]` — https://www.heinlein-support.de/sites/default/files/media/documents/2026-05/SLAC2026_Self-Healing_mit_Event_Driven_Ansible_Ren%C3%A9_Koch.pdf (SLAC 2026 slides).
- Event Streams (Red Hat, 2025): enhancement routing one webhook endpoint to many rulebook activations; horizontally-scaled activations; credential-required endpoints (HashiCorp Vault/CyberArk integration); recommended path: webhooks for dev/test → Event Streams for production → Kafka/SQS/Service Bus for guaranteed delivery, persistence, replay, high volume `[secondary]` — https://www.redhat.com/de/blog/event-driven-ansible-simplified-event-routing-event-streams.
- Network event-handling patterns: Alertmanager webhook → EDA rulebook → condition on `event.alert.labels.alertname` → `run_workflow_template` with extracted extra_vars; throttling (`once_within: 3 hours`, `group_by_attributes`) to dedupe; alert remapping and conditional routing `[secondary]` (ansible-skills repo).
- EDA ports (self-managed): 5000/tcp webhook, 5001/tcp AlertManager, 5002/tcp Insights `[secondary]` — https://github.com/zaskan/event-driven-automation.
- AI angle (2026): community projects combine EDA rulebooks with MCP/LLM playbook generation (ansible-aiops demo: EDA rulebook listens on webhooks/Kafka, falls back to MCP template finder; Red Hat Code Assistant generates playbooks from events) `[secondary]` — https://github.com/iamgini/ansible-aiops; treat as experimental `[unverified]`.

### E4.9.3 Webhooks and Kafka as network event buses

- Webhook pattern: monitoring (Prometheus Alertmanager, Checkmk, Dynatrace) POSTs JSON to an automation endpoint; rulebooks/actions parse and remediate `[secondary]` (multiple sources above).
- Kafka pattern: high-volume, multi-consumer, replayable event streams; used for network telemetry/alert pipelines and as EDA source `[secondary]` (Red Hat EDA blog).
- MQTT/file-watch/journald cover IoT, log-driven, and host-level triggers `[secondary]` (SLAC 2026 slides).

### E4.9.4 StackStorm vs EDA positioning (synthesis)

- StackStorm: standalone, workflow-centric, 160+ packs, ChatOps-native; LF-governed; broad IT scope `[analysis]`.
- Event-Driven Ansible: Ansible-native, rulebook-centric, AAP-integrated, Event Streams for production routing; strongest where Ansible content already exists `[analysis]`.
- No 2026 independent comparison found `[unverified]`.

---

## E4.10 — Selection guidance (synthesis)

### E4.10.1 Validation tool selection

| Need | Candidates | Notes |
|---|---|---|
| Pre-change config correctness proof | Batfish | Free, open source, offline; partial EVPN |
| Pre-change production impact | Forward Predict | Commercial; deterministic twin claim [vendor-reported] |
| Observed-state drift/compliance | SuzieQ | OSS + Enterprise; 90s polling [vendor-reported] |
| Intent verification at scale | IP Fabric | Commercial; 160+ checks, cloud twin |
| Agentic AI on network truth | Forward AI | Commercial; GA Apr 2026 planned [unverified] |
| Service-provider closed loop | Cisco Crosswork | Cisco-centric; mature portfolio |

### E4.10.2 Lab platform selection

- Containerlab: best CI/CD fit, container-native, YAML-as-code, free `[analysis]`.
- EVE-NG Pro: largest VM-based labs (1024 nodes), team features, commercial `[analysis]`.
- GNS3: free, broad appliance library, split GUI/server `[analysis]`.
- CML: Cisco-centric, API-first, MCP server for AI tooling, free tier per 2.8 notes `[analysis]`.

### E4.10.3 Testing framework selection

- pyATS: Cisco-heavy shops, Genie parsers, millions-of-tests pedigree `[analysis]`.
- Robot Framework: keyword-driven readability, mixed stacks, CI-friendly reports; pairs with pyATS (nac-test) `[analysis]`.
- pytest: Python-native custom suites, async SSH patterns `[analysis]`.

### E4.10.4 Pipeline pattern (recommended minimum)

1. Config change in Git (branch) → 2. lint + Batfish differential analysis in CI → 3. ephemeral lab (Containerlab/CML) spin-up with Robot/pyATS tests → 4. human review (plan artifact) → 5. merge → 6. staged apply with approval gate → 7. post-change SuzieQ/IP Fabric verification + drift watch → 8. event-driven remediation (EDA/StackStorm) on anomalies. `[analysis]` synthesis of sources above.

---

