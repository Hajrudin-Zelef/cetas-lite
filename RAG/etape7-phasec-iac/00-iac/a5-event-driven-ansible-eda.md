---
id: etape7-phasec-iac/00-iac/a5-event-driven-ansible-eda
title: "A5. Event-Driven Ansible (EDA)"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["research"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [63, 77]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: b70c0da66ddf5b951efd50b89ae6c6e520caf250a23e7b37ae4c0a8c3a7dd16d
---

# A5. Event-Driven Ansible (EDA)

### A5. Event-Driven Ansible (EDA)

- EDA pairs event sources with **rulebooks** (`ansible-rulebook` CLI) so automation triggers on events (webhooks, Kafka, monitoring alerts) instead of schedules [secondary](https://github.com/fitbeard/automation-platform/blob/HEAD/README.md).
- The EDA server component is at **1.2.12** in the AAP 2.6-era component set [secondary](https://github.com/fitbeard/automation-platform/blob/HEAD/README.md).
- EDA decision environments (container images with `ansible-rulebook` + event-source plugins) mirror the EE pattern from job execution [secondary](https://github.com/fitbeard/automation-platform/blob/HEAD/README.md).
- 2026 maturity note: EDA is positioned for closed-loop remediation (alert → rulebook → playbook → ticket update), but independent 2026 adoption numbers were not found in this research pass — flagged as a gap.

### A6. Ansible adoption signals

- Ansible holds a 9.1/10-style rating on TrustRadius 2026 comparisons vs Microsoft System Center and others [secondary](https://www.trustradius.com/compare-products/microsoft-system-center-vs-red-hat-ansible-automation-platform).
- Community health: the AWX operator and AWX repos show continuous release activity through 2026 (operator 2.19.x, Helm chart 3.2.1) [official](https://github.com/ansible/awx-operator/releases).
- No independently verified 2026 market-share figure for Ansible vs Puppet/Chef/Salt was found in this pass — flagged as a gap. (Puppet/Foreman are out of scope for this file; see Step 7 Phase A for OS-level tooling overlap.)

---

