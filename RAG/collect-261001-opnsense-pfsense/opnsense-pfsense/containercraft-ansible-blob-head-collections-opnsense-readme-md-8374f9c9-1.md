---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9-1
title: "1. Install this collection and its upstream dependency into a playbook-adjacent"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9.md
source_anchor: ""
source_lines: [1, 143]
sha256: d061a37d1e354c75852efefd83f9350d6cbdac28c3e7869e3d019392f5ae1e42
---

# 1. Install this collection and its upstream dependency into a playbook-adjacent

A general-purpose Ansible collection that configures an OPNsense edge router entirely through its REST API. It builds VLAN interfaces, a hardened Unbound DNS resolver with DNSBL threat intelligence, Kea DHCP with per-subnet security policy, a sequenced zone-to-zone firewall with savepoint rollback, and CARP high availability — and it sequences those changes through a monotonic migration model so a live network can be re-architected without an outage.

The collection ships reusable roles only. A companion *deploy* layer (an
inventory, a values file, and an executable entrypoint) turns those roles into a
running deployment. The two layers are kept strictly separate so the collection
stays publishable and reusable while each site's identity lives in its own deploy
tree.

This document serves every reader from a first-time hobbyist to a collection
maintainer. It is ordered by **how often a section is revisited**, not by
difficulty: the material an operator returns to daily sits near the top, and the
deep, read-once architecture and contributor material sits lower.

```
graph TD
    A["§1-2 Orientation & Quickstart<br/>read once, but first"] --> B["§3-5 Mental models<br/>read early, then settle"]
    B --> C["§6-9 Operational reference<br/>revisited constantly"]
    C --> D["§10-12 Troubleshooting & runbooks<br/>revisited under pressure"]
    D --> E["§13-16 Architecture & internals<br/>read deeply, revisit when changing"]
    E --> F["§17-19 Contributing & appendices<br/>maintainer depth, occasional"]
    style C fill:#2d5016,color:#fff
    style D fill:#5c2d00,color:#fff
```
    Each section opens at the lowest level of background it can and adds depth toward the end. A reader can stop at the first heading that answers the question at hand. Three callout markers flag content by audience:

**Blast radius** — an operation that changes a live network; read before running.

⚠️ 

💡 **Shortcut** — a faster or safer path for routine work.


🔬 **Deeper** — internals and rationale for maintainers and SMEs; skippable for operation.


| Reader | Read first | Then | Reference | 
|---|---|---|---|
| Hobbyist / student | §1, §2 | §3, §4 | §6, §7 | 
| Intern / junior operator | §1–§4 | §7, §10, §11 | §6, §9 | 
| Senior operator | §3–§5 | §8, §9, §11, §12 | §10, §19 | 
| Staff / principal engineer | §5, §13, §14 | §15, §16 | §9, §17 | 
| Maintainer / SME | §14–§17 | §18, §19 | all | 

`containercraft.opnsense` automates an OPNsense firewall the way a person would
automate a cloud provider: by calling an API. Every change — a VLAN, a firewall
rule, a DHCP scope, a DNS hardening parameter — is an HTTP request to the
firewall's REST endpoint. There is no configuration file pushed to the box and,
with one deliberate exception, no SSH session. The firewall is treated as a
service with an API, and the roles are clients of that service.

The collection encodes four things that raw API calls do not:

- **A data model.** The network's shape — its zones, subnets, trunk ports, and
inter-zone policy — is declared once as data. Roles render that data into API
calls. No role contains a literal subnet, VLAN tag, or IP address.
- **A migration model.** Real edge routers are replaced while in service. A
monotonic phase variable (`coexist → migrating → hardening → converged` ) gates
which services run where, so the existing LAN keeps working until clients have
moved.
- **An idempotent, recoverable apply.** Re-running converges to the same state.
Firewall changes are wrapped in a server-side savepoint that auto-reverts if
connectivity is lost, so a bad rule cannot lock the operator out.
- **A hardened security posture.** The DNS resolver hides its identity and
version, validates DNSSEC, enforces rebinding protection with private-address
filtering, runs DNSBL threat intelligence feeds (Abuse.ch ThreatFox, Hagezi),
sets the ACL default to refuse with explicit per-zone allows, and sizes caches
for bare-metal performance. The DHCP server uses UDP sockets (pf-gated, not
raw), disables option auto-collection to prevent HA race conditions, and applies
per-subnet allocator, client-ID, and lifetime policies via the raw API. The
firewall uses explicit sequence numbering, quick-match semantics, and includes
DHCP server allow rules that are required when Kea's automatic firewall rules
are disabled.

- It is **not** a replacement for OPNsense's own configuration. It manages the
objects it creates (tagged with a managed-by marker) and leaves everything else
alone.
- It is **not** a one-shot installer. It is designed to be run repeatedly, a
little at a time, advancing through phases.
- It does **not** implement masquerade/outbound NAT. OPNsense's automatic
outbound NAT already masquerades internal zones to the WAN address; the
collection relies on that default rather than duplicating it (see §9.6 and §16).
- It does **not** set per-interface IP addresses on opt* interfaces. The OPNsense
MVC API does not expose this; it is a manual step in the web UI between the
VLAN creation phase and the DHCP activation phase.
- It is **not** tied to any one site. The collection is general-purpose; the
example network used throughout this document belongs to a deploy layer, not to
the collection.

This section takes a reader from nothing to a safe, read-only dry run against a firewall. It assumes the firewall is reachable and an operator account exists on it.

- `ansible-core >= 2.17` and Python`>= 3.12` on the controller.
- Network reachability to the firewall's HTTPS API.
- An OPNsense user that can hold an API key. A dedicated service user is recommended over a human login.

```
# 1. Install this collection and its upstream dependency into a playbook-adjacent
#    collections/ directory (the deploy layer does this automatically; shown here
#    for a standalone install).
ansible-galaxy collection install -r requirements.yml -p collections
# 2. Provide credentials. The deploy layer reads them from the environment.
export OPNSENSE_FIREWALL=10.0.0.1
export OPNSENSE_API_KEY=...        # see §12 to mint these
export OPNSENSE_API_SECRET=...
# 3. Confirm the API answers and the key authenticates (read-only).
ansible-playbook site.yml --tags connect
# 4. Dry-run the whole configuration. Nothing is written; the would-change set
#    and per-object diffs are printed.
ansible-playbook site.yml --check --diff
# 5. Apply one concern at a time once the dry run looks correct.
ansible-playbook site.yml --tags dns
```
💡 In a configured deploy tree, the entrypoint replaces `ansible-playbook site.yml` with a single executable: `./stargate.yml --tags connect`. See §11.


`--check` and `--tags connect`.

⚠️ Command 5 writes to the firewall. Until an operator has read §4 (phases) and §11 (the iterative runbook), it is safest to stay on

`--tags connect` ends with `OPNsense API reachable at <address>` and
`failed=0`. If it does not, jump to §10.1.

The single most important idea in this collection: **the firewall is configured
by calling its REST API, not by editing it.** Internalizing this explains the
structure of every role.

```
graph LR
    OP["Operator"] -->|ansible-playbook| C["Ansible controller<br/>(connection: local)"]
    C -->|HTTPS REST<br/>key + secret| FW["OPNsense API<br/>https://firewall/api/..."]
    FW --> CFG["config.xml<br/>(applied on reload)"]
    subgraph "controller side"
        C
    end
    subgraph "firewall side"
        FW
        CFG
    end
```
    The controller runs the modules locally (`connection: local`) and each module
issues HTTPS requests to the firewall. The firewall validates, stores, and — on a
`reload` call — applies the change. This has several consequences that shape the
whole collection:

