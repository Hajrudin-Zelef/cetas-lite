---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9-6
title: "1. Install this collection and its upstream dependency into a playbook-adjacent"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9.md
source_anchor: ""
source_lines: [809, 980]
sha256: 60096982d64e6068ba08f924701122e793e2be96fbc04d8b8fe7e5bfda40eac2
---

# 1. Install this collection and its upstream dependency into a playbook-adjacent

```
sequenceDiagram
    participant OP as Operator
    participant C as Controller
    participant FW as Firewall (SSH+root)
    participant ENV as deploy/.env
    OP->>C: ansible-playbook bootstrap.yml --ask-pass --ask-become-pass
    C->>FW: render mint script (template)
    C->>FW: run mint script as root
    Note over FW: createKey() generates key+secret,<br/>persists crypted secret to config.xml
    FW-->>C: STATUS=created, KEY, SECRET (returned once)
    C->>FW: remove mint script
    C->>ENV: write OPNSENSE_API_KEY / OPNSENSE_API_SECRET
    Note over ENV: subsequent API-driven runs read these
```
    | Condition | Behavior | Result | 
|---|---|---|
| key exists, no `force` | do nothing | `exists` ;`.env` unchanged | 
| no key, no `force` | mint one | `created` ;`.env` written | 
| `force` | mint new, then delete all prior keys for the user | `created` ; exactly one key remains | 

```
graph TD
    START["mint requested"] --> Q1{"user has a key?"}
    Q1 -->|no| MINT["add key, persist"]
    Q1 -->|yes| Q2{"force?"}
    Q2 -->|no| EXISTS["STATUS=exists<br/>(no change)"]
    Q2 -->|yes| ROTATE["add new key,<br/>delete prior keys"]
    MINT --> VERIFY["re-read; confirm count >= 1"]
    ROTATE --> VERIFY
    VERIFY --> CREATED["STATUS=created<br/>write .env"]
    style EXISTS fill:#444,color:#fff
    style CREATED fill:#2d5016,color:#fff
```
    The plaintext secret is returned exactly once at creation and is not retrievable
afterward (OPNsense stores it hashed). If a key exists but its secret was not
captured, the only way to obtain a usable secret is to rotate with `force`.

⚠️ A dedicated, least-privilege service user is recommended as the API-key holder, distinct from the SSH/wheel user used to run bootstrap.

```
sequenceDiagram
    participant E as ./stargate.yml
    participant S as site.yml (play)
    participant PRE as pre_tasks
    participant ROLES as roles (FQCN)
    participant FW as OPNsense API
    E->>S: import_playbook, vars: values_file=stargate_values.yml
    S->>PRE: load netspec, phases, matrix, run toggles
    PRE->>PRE: load operator values (this deploy)
    PRE->>PRE: merge netspec_overrides; resolve phase capabilities
    S->>ROLES: opn_connect ... opn_edge_decommission (gated)
    ROLES->>FW: configure objects (reload:false)
    ROLES->>FW: one reload per subsystem
    ROLES->>FW: verify; commit (firewall savepoint)
```
    1. **Entry.** The executable entrypoint selects the deployment's values file and
imports the generic`site.yml` .
2. **Data load.**`pre_tasks` load the netspec, phase map, firewall matrix, and
operator values, then deep-merge overrides and resolve phase capabilities.
3. **Role composition.** Eight API-driven roles in dependency order. Each
configures objects with deferred reloads, then issues a single reload. The
firewall role wraps its changes in a savepoint. The DNS role applies
hardening and DNSBL via raw API. The DHCP role applies per-subnet overrides
via raw API after the typed module creates the subnets.
4. **Convergence.** Re-running produces no changes once the firewall matches the
declared state.

Every variable a role reads has a default inside that role. A galaxy install works without the deploy tree.

The site's name, inventory, real network data, and credentials are not reusable
and must not leak into the published artifact. It lives in `deploy/opnsense/`:

| Deploy artifact | Role | 
|---|---|
| `stargate.yml` | executable entrypoint; the only place the box name appears | 
| `site.yml` | generic orchestration (no site specifics) | 
| `vars/<deploy>_values.yml` | this site's overrides | 
| `inventory/` | API target, connection mode | 
| `.envrc` ,`.env` ,`ansible.cfg` | environment, credentials, runtime config | 

```
# deploy/opnsense/stargate.yml
#!/usr/bin/env -S ansible-playbook --inventory=inventory/hosts.yml
- name: "ContainerCraft | stargate | OPNsense edge"
  import_playbook: site.yml
  vars:
    values_file: "{{ playbook_dir }}/vars/stargate_values.yml"
```
`import_playbook` propagates `vars:` to the imported plays but not `vars_files:`.
Setting `values_file` via `vars:` is the mechanism that lets a generic `site.yml`
load a site-specific values file without naming it. One executable per site, a
generic playbook shared by all sites.

Most objects are matched on their description, which carries the managed-by
prefix (`opnsense_managed_tag`). Re-running finds the existing object and updates
it in place rather than creating a duplicate. Changing a managed value is detected
as a change; running again with no change is a no-op. Composite keys are used
where a description is insufficient (for example, Unbound host records match on
hostname + domain + record type).

Firewall filter changes can lock out the operator if a rule is wrong. The
firewall role wraps its changes in a server-side savepoint: it applies the rules,
verifies the API is still reachable, and only then cancels the savepoint's
auto-rollback. If the verification does not complete within the 60-second timer
window, the filter section reverts automatically. The operator cannot be
permanently locked out by a bad apply. The savepoint covers `firewall/filter`
only — not interfaces, DHCP, or DNS.

Bulk object tasks set `reload: false` and each role issues a single reload per
subsystem at its end. This avoids applying a half-built configuration and reduces
churn. A run interrupted mid-role generally leaves staged but un-applied changes
rather than a partial live configuration.

```
graph TD
    A["declare desired state (data)"] --> B["match existing by description"]
    B -->|found| C["update in place if drifted"]
    B -->|absent| D["create"]
    C --> E["reload subsystem once"]
    D --> E
    E --> F["verify reachable"]
    F -->|ok| G["commit / converged"]
    F -->|fail| H["savepoint auto-revert"]
```
    `containercraft.opnsense` orchestrates; `oxlorg.opnsense` implements the API
modules.

| Concern | Owned by | 
|---|---|
| API protocol, request/response, module argument specs | `oxlorg.opnsense` | 
| Which objects to create and in what order | `containercraft.opnsense` roles | 
| The network's shape (zones, trunks, matrix) | `shared/netspec/` data | 
| Credential injection into module calls | `module_defaults` action group | 
| Fields not exposed by typed modules | `oxlorg.opnsense.raw` API calls | 

```
graph TB
    DATA["shared/netspec/ data<br/>(zones, trunks, matrix)"]
    ROLES["containercraft.opnsense roles<br/>(what to create, in what order)"]
    MD["module_defaults<br/>group/oxlorg.opnsense.all<br/>(injects credentials + timeout + retries)"]
    OXL["oxlorg.opnsense modules<br/>(speak the REST protocol)"]
    RAW["raw API tasks<br/>(hardening, per-subnet overrides, DNSBL, ACL default)"]
    API["OPNsense REST API"]
    DATA --> ROLES
    ROLES --> MD
    ROLES --> RAW
    MD --> OXL
    OXL --> API
    RAW -->|explicit opnsense_api_args| API
    style ROLES fill:#2d5016,color:#fff
    style OXL fill:#1f3a5f,color:#fff
    style RAW fill:#5c2d00,color:#fff
```
    The `module_defaults` group injects the connection arguments (including
`api_timeout` and `api_retries`) into all modules that upstream lists in its
action groups. The `raw` module and `ansible.builtin.uri` are not in that group,
so tasks using them pass the connection arguments explicitly via
`opnsense_api_args`.

🔬 **Deeper.** Upstream pins matter. The collection declares an exact dependency
(`oxlorg.opnsense == <version>`) because module argument specs and action-group
membership change between versions. Upgrading the dependency is a deliberate,
tested change.


A zone is data. Add an entry to `shared/netspec/zones.yml` with the standard
shape. No role changes required.

A rule is a row in `shared/netspec/firewall_matrix.yml` with a unique `seq` and
`desc`. The firewall role flattens all groups and applies each row.

