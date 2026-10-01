---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9-2
title: "1. Install this collection and its upstream dependency into a playbook-adjacent"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9.md
source_anchor: ""
source_lines: [144, 295]
sha256: eb010776680669ce3b2b3df1773923568958635b589c36275170b1f47b73d2d0
---

# 1. Install this collection and its upstream dependency into a playbook-adjacent

- **Authentication is an API key/secret pair** , sent on every request. There is
no session, no SSH key, no password to the box for normal operation.
- **A change is two steps: configure, then reload.** Roles create or update many
objects with`reload: false` , then issue a single`reload` per subsystem. This
is faster and avoids partial-apply churn.
- **Most operations are check-mode safe.** Because they are API reads and writes,
Ansible can predict and diff them without touching the box.
- **Connection resilience is built in.** The`module_defaults` block sets`api_timeout: 30` and`api_retries: 2` , tolerating up to two transient
connection failures and a 30-second overall timeout per request.

A firewall configured through its API needs an API key before it can be
configured. Minting that first key writes `config.xml`, which is root-gated and
cannot be done through the API. Exactly one role — `opn_bootstrap_apikey` — breaks
the API-only rule: it connects over SSH, escalates to root, mints a key, and
writes it where the API-driven roles can read it. After bootstrap, everything
returns to the API model. This chicken-and-egg resolution is detailed in §12.

Some OPNsense model fields are not exposed by the upstream `oxlorg.opnsense`
collection's typed modules. The collection reaches these fields through the
`oxlorg.opnsense.raw` module, which calls any OPNsense API endpoint directly.
This pattern is used for:

- Unbound hardening parameters at `unbound.advanced.*` (identity hiding, DNSSEC
hardening, rebinding protection, cache sizing, DNSBL, ACL default action)
- Kea general settings at `dhcpv4.general.*` (socket retry configuration)
- Kea per-subnet settings at `dhcpv4/set_subnet/<uuid>` (allocator, client-ID
matching, per-zone lease lifetimes)

🔬 **Deeper.** The collection consumes the `oxlorg.opnsense` collection for the
actual API module implementations. `containercraft.opnsense` orchestrates;
`oxlorg.opnsense` speaks the protocol. The boundary between them is described in
§16, and it is the reason the roles can stay declarative.


An edge router is rarely built on a green field. It usually replaces an existing
router while the network stays up. The collection models this as a **monotonic
phase**, a single variable that advances in one direction and gates what each role
is allowed to do.

```
stateDiagram-v2
    [*] --> coexist
    coexist --> migrating
    migrating --> hardening
    hardening --> converged
    converged --> [*]
    coexist: coexist — new edge stands up beside the live LAN
    migrating: migrating — clients move; zone DHCP/DNS active
    hardening: hardening — posture tightened; HA prepared
    converged: converged — legacy edge services withdrawn
```
    The phase is not a feature flag set; it is a promise about what the **existing
network** experiences at each step. The guiding rule is *never strand a client*.

| Phase | Legacy segment edge DHCP/DNS | New zones | Zone DHCP | HA | Decommission | 
|---|---|---|---|---|---|
| `coexist` | untouched | created, no clients | off | off | off | 
| `migrating` | still served | active | on | off | off | 
| `hardening` | still served (fallback) | active | on | prepared | off | 
| `converged` | withdrawn | active | on | active | on | 

A phase resolves to a small set of capability booleans (defined in
`shared/netspec/phases.yml`), and roles gate their riskier tasks on those
booleans rather than checking the phase string directly:

```
# shared/netspec/phases.yml (excerpt)
phase_defaults:
  coexist:
    zone_dhcp: false
    edge_serves_legacy_lan: true
    decommission_legacy: false
  converged:
    zone_dhcp: true
    edge_serves_legacy_lan: false
    decommission_legacy: true
```
**Blast radius.** The phase is monotonic by intent. Advancing to `converged`
instructs `opn_edge_decommission` to remove edge DHCP/DNS from the legacy
segment. That role is additionally gated on measured evidence (zero active
leases) so the withdrawal cannot strand a device that has not yet moved. See
§9.8.

⚠️ 

💡 The default phase is `coexist` — the safe, non-disruptive posture. A fresh
run with no configuration touches no existing client traffic.


The collection contains **reusable roles and nothing site-specific**. A separate
**deploy layer** supplies the site's identity — its inventory, its credentials,
its network data, and an executable entrypoint. Understanding this split explains
why the roles have no subnets in them and where an operator's own values belong.

```
graph TB
    subgraph "collections/opnsense — GENERAL PURPOSE (published)"
        R["roles/<br/>opn_connect, opn_system, opn_interfaces,<br/>opn_dns, opn_dhcp, opn_firewall,<br/>opn_ha, opn_edge_decommission,<br/>opn_bootstrap_apikey"]
        D["each role's defaults/main.yml<br/>(sane, generic, env-overridable)"]
        G["galaxy.yml, meta/runtime.yml,<br/>changelogs/"]
    end
    subgraph "deploy/opnsense — SITE SPECIFIC (private)"
        E["stargate.yml<br/>(executable entrypoint — box identity)"]
        S["site.yml<br/>(generic orchestration)"]
        V["vars/stargate_values.yml<br/>(this site's overrides)"]
        I["inventory/ + .envrc + ansible.cfg"]
    end
    N["shared/netspec/<br/>zones, phases, firewall_matrix<br/>(network data)"]
    E --> S
    S --> R
    V -.overrides.-> D
    S -.loads.-> N
```
    | Layer | Lives in | Contains | Site identity? | 
|---|---|---|---|
| **Collection** | `collections/opnsense/` | reusable roles, defaults, galaxy metadata | none | 
| **Deploy** | `deploy/opnsense/` | entrypoint, inventory, values, env | confined here | 
| **Network data** | `shared/netspec/` | zones, phases, firewall matrix | site topology as data | 

The boundary is enforced by a single discipline: **every variable a role needs has
a default inside that role.** A consumer can install the collection from a galaxy
server and run its roles with sane defaults, overriding through the environment or
playbook variables. The example deployment in this repository is one such
consumer; it is not part of the published collection.

The commands below assume a configured deploy tree. In a deploy tree the
executable entrypoint (`./stargate.yml`) is equivalent to `ansible-playbook site.yml` with the site's values pre-selected; both forms are shown where they
differ.

| Goal | Command | 
|---|---|
| Confirm API + credentials (read-only) | `./stargate.yml --tags connect` | 
| Dry-run everything with diffs | `./stargate.yml --check --diff` | 
| Apply one concern | `./stargate.yml --tags <tag>` | 
| Apply everything | `./stargate.yml` | 
| List the tasks a run would execute | `./stargate.yml --list-tasks` | 
| Syntax check only | `./stargate.yml --syntax-check` | 

Each role contributes one or more tags so a run can be scoped to a single
concern. Tags compose; `--tags dns,dhcp` runs both.

| Tag | Role | Effect | 
|---|---|---|
| `connect` | opn_connect | read-only API reachability probe | 
| `system` | opn_system | CARP tunables and (gated) identity/WAN | 
| `interfaces` ,`vlans` | opn_interfaces | trunk VLANs, blackhole VLAN, assignment | 
| `dns` ,`unbound` | opn_dns | hardened Unbound, DNSBL, ACLs, split-horizon forward | 
| `dhcp` ,`kea` | opn_dhcp | Kea DHCP per zone with per-subnet hardening | 
| `firewall` ,`nat` | opn_firewall | sequenced aliases and zone rules with DHCP allow | 
| `ha` | opn_ha | CARP VIPs, hasync (pfsync v1400), config sync trigger (gated) | 
| `decommission` | opn_edge_decommission | legacy-segment withdrawal (gated) | 

