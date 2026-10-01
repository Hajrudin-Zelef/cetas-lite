---
id: collect-261001-meraki/meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed-1
title: "List all networks in an organization"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed.md
source_anchor: ""
source_lines: [1, 63]
sha256: de658926d8133e395027a10324742ac13f35b8604dd205f8743453ce5ceff174
---

# List all networks in an organization

Imagine you are a network engineer responsible for 50 branch offices, each with a Cisco Meraki security appliance, a stack of switches, and a wireless access point. Every week there are routine tasks: updating firmware, rotating Wi-Fi passwords, adjusting security policies. Done manually — logging into each dashboard, clicking through menus, verifying the change — this work is:
- Slow. An experienced engineer might configure 10 devices per hour.
- Error-prone. One missed checkbox, one typo in a VLAN ID, and you have an incident.
- Undocumented. What was the network configured to before you changed it? Who changed it? Why?
- Unscalable. 50 sites today, 500 sites tomorrow — the work scales linearly with headcount.
Infrastructure as Code (IaC) solves all of these problems.
Infrastructure as Code is the practice of defining and managing your network (or system) configuration in plain text files, and using software tools to apply those files to your infrastructure automatically.
Instead of clicking through a dashboard, you write a file that says:
network: Branch-Office-01
firmware:
  wireless_version_id: "15763"   # MR 32.1.7 stable
  upgrade_datetime: "2026-07-01T02:00:00Z"
Then you run a command — ansible-playbook, python script.py, or terraform apply — and the tool reads that file, talks to the Meraki Dashboard API, and makes the configuration match exactly what you declared. No clicking. No manual steps. No forgetting.
| Traditional "ClickOps" | Infrastructure as Code | 
|---|---|
| Changes are invisible — no record of what changed or when | Every change is a Git commit — full history, author, timestamp, and reason | 
| Documentation is a separate, always-out-of-date Word document | The configuration is the documentation | 
| Reproducing a configuration on a new device requires tribal knowledge | Re-run the playbook or terraform apply on any device | 
| Testing means hoping you did not break anything | You can run --check (dry run) to preview changes before applying them | 
| One engineer can manage ~50 devices | One engineer with automation can manage 5,000 devices | 
| A mistake is hard to reverse | A rollback is git revert followed by re-applying | 
IaC delivers the most value in environments that share some of these characteristics:
- Scale — more than a handful of devices or networks to manage
- Repetition — the same configuration pattern applied across many sites
- Compliance requirements — auditors need proof that configurations match a documented standard
- High rate of change — frequent firmware updates, regular password rotations, dynamic policy changes
- Team environments — multiple people making changes to the same infrastructure
- Disaster recovery — the ability to rebuild a configuration from scratch quickly
Even in small environments (5–10 devices), adopting IaC early builds habits and skills that pay dividends as the network grows.
One of the most powerful concepts in software engineering — and increasingly in network engineering — is abstraction: the idea that you should only have to think about what you want, not how to achieve it.
When you drive a car, you turn the steering wheel — you do not think about the hydraulic fluid pressure in the power steering rack. The steering wheel is an abstraction over a complex mechanical system. IaC works the same way.
The diagram below shows the four layers of abstraction in this repository:
Diagram source: diagrams/iac-abstraction.mmd
Layer 1 — Human Intent. You write a YAML file, a Python script, or an Ansible playbook that expresses what you want in plain, readable language. You do not need to know anything about HTTP headers, JSON payloads, or API rate limits.
Layer 2 — Automation Engine. A tool (Python SDK, Ansible collection, or Terraform provider) reads your intent and translates it into the correct API calls. It handles authentication, pagination, error retries, and idempotency — meaning it only makes a change if a change is actually needed.
Layer 3 — Meraki Dashboard API. The Meraki REST API is the authoritative interface to your Meraki organization. Every operation — reading network state, scheduling firmware upgrades, configuring SSIDs — happens through this API. The automation engines call it on your behalf.
Layer 4 — Physical Infrastructure. Your actual Meraki devices (MX appliances, MS switches, MR access points) receive configuration changes pushed by the Meraki Dashboard cloud. You never connect directly to the devices.
The practical benefit of this layered model is that you — the network engineer — work entirely at Layer 1. You never need to:
- Write an HTTP PUT request by hand
- Remember that firmware upgrades use a different endpoint format than SSID configuration
- Handle API pagination when an organization has more than 1,000 networks
- Worry about whether to use GET beforePUT to check current state
The automation engine handles all of that. Your job is to declare intent clearly.
When multiple engineers use the same tool to apply the same data model, configuration drift becomes impossible. There is no way for one engineer to configure Site A differently from Site B if both sites are defined in the same YAML file and applied with the same playbook. The abstraction enforces consistency as a side effect.
This repository demonstrates three distinct IaC approaches for Meraki network management. Each one sits at a different point on the spectrum from flexibility to declarative intent:
 FLEXIBLE ◄────────────────────────────────────────► DECLARATIVE
   │                        │                               │
Python SDK            Ansible Playbooks           Network as Code
                                                    (Terraform)
 "I write the          "I describe the              "I declare the
  logic"                steps"                       desired state"
Python gives you a full programming language. You write loops, conditionals, data transformations, and error handling exactly the way you want. The Meraki Python SDK wraps the REST API in simple Python function calls.
Best for: Custom logic, one-off data collection, complex multi-step workflows that do not fit neatly into a tool's model.
Ansible is a task automation tool built around the concept of playbooks — ordered lists of tasks expressed in YAML. Each task calls a module (a pre-built piece of logic) that knows how to talk to a specific system. The cisco.meraki Ansible Galaxy collection provides modules for the Meraki Dashboard API.
Best for: Multi-step workflows that must run in order, teams already using Ansible for server or network automation, situations where you want human-readable automation that non-developers can review.
Terraform is a declarative tool: you describe the desired final state of your network in a data model (YAML files), and Terraform figures out what API calls are needed to reach that state. The netascode/nac-meraki Terraform module reads your YAML and translates it into Meraki API operations.
Best for: Long-lived configuration that must always match a declared state, SSID management, policy enforcement, teams familiar with Terraform from cloud infrastructure management.
| # | Use Case | Approach | Files | 
|---|---|---|---|
| 1 | Discover available firmware versions across networks by tag | Python, Ansible | python/1.0_* ,ansible/1.0_* | 
| 2 | Schedule firmware upgrades across networks by tag, only where current ≠ target | Python, Ansible | python/2.0_* ,ansible/2.0_* | 
| 3 | Manage SSID configuration (name, PSK, encryption, visibility) declaratively | Network as Code | network-as-code/ ,data-model/nac/ | 
| 4 | PSK rotation without storing secrets in files using environment variable injection | Network as Code | network-as-code/networks_nac.yaml ,!env tags | 
