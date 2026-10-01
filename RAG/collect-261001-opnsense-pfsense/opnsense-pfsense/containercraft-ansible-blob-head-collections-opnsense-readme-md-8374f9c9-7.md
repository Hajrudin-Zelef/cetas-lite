---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9-7
title: "1. Install this collection and its upstream dependency into a playbook-adjacent"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9.md
source_anchor: ""
source_lines: [981, 1073]
sha256: 453969378e1fb77fb0b2e30ffb70355a83e2bf747a791742ed877a5532f8aed1
---

# 1. Install this collection and its upstream dependency into a playbook-adjacent

A capability is a boolean in `shared/netspec/phases.yml` under each phase, and a
`when:` gate in the consuming role. Define the capability for every phase so the
phase map stays total, then gate the role's task on `phase.<capability>`.

Give the role a complete `defaults/main.yml` with the connection surface, its own
toggles, and safe data fallbacks. Compose it into `site.yml` with an FQCN
reference and a tag. Match the existing roles' shape: deferred reloads,
description-keyed idempotency, check-mode guards on tasks that cannot dry-run,
and the savepoint pattern for anything that can sever connectivity.

```
ansible-lint collections/opnsense
ansible-galaxy collection build collections/opnsense --output-path build/
./stargate.yml --check --diff
```
- Lints clean at the production profile.
- Roles remain self-contained (every variable has a role default).
- No site identity leaks into the collection.
- Mutating tasks are check-mode guarded; connectivity-severing tasks are savepoint-wrapped.
- The changelog records the change.

The repository's workflows lint on pull request and publish on a version tag
(`<collection>-v<semver>`). The `version` field in `galaxy.yml` is a placeholder
the publish step overwrites from the git tag.

| Term | Meaning | 
|---|---|
| **Zone** | a logical network (usually a VLAN) with a subnet, gateway, and policy | 
| **Trunk** | a physical port carrying tagged zone VLANs to a switch/power domain | 
| **Blackhole VLAN** | a deny-all native VLAN per trunk; carries no addresses | 
| **Phase** | the monotonic migration state ( `coexist` …`converged` ) | 
| **netspec** | the network's shape declared as data in `shared/netspec/` | 
| **Firewall matrix** | inter-zone policy declared as a table of sequenced rows | 
| **Savepoint** | a server-side firewall checkpoint with a 60-second auto-revert timer | 
| **Managed tag** | the description prefix marking automation-owned objects; also the idempotency key | 
| **Action group** | the upstream module set into which `module_defaults` injects credentials | 
| **Deploy layer** | the site-specific tree ( `deploy/<name>/` ) that consumes the collection | 
| **Entrypoint** | the executable, shebang-driven playbook that carries a site's identity | 
| **CARP / pfsync** | the HA mechanisms: virtual IPs (CARP) and state sync (pfsync v1400) | 
| **Kea / Unbound** | the DHCP server and the DNS resolver OPNsense uses | 
| **DNSBL** | DNS-based blocklist; Unbound blocks queries for known-malicious domains | 
| **Raw API** | `oxlorg.opnsense.raw` module calling OPNsense endpoints not exposed by typed modules | 

| Variable | Consumed by | Default | 
|---|---|---|
| `OPNSENSE_FIREWALL` | connection | empty | 
| `OPNSENSE_API_KEY` /`OPNSENSE_API_SECRET` | connection | empty | 
| `OPNSENSE_API_PORT` | connection | 443 | 
| `OPNSENSE_SSL_VERIFY` | connection | false | 
| `OPNSENSE_NETWORK_PHASE` | posture | coexist | 
| `OPNSENSE_MANAGED_TAG` | marking | ansible-managed | 
| `OPNSENSE_APPLY_IDENTITY` | opn_system | false | 
| `OPNSENSE_HOSTNAME` /`OPNSENSE_DOMAIN` /`OPNSENSE_TIMEZONE` | opn_system | opnsense / home.arpa / UTC | 
| `OPNSENSE_WAN_IPV4_TYPE` | opn_system | dhcp | 
| `OPNSENSE_ENFORCE_DNSMASQ_DISABLED` | opn_dhcp | false | 
| `OPNSENSE_REQUIRE_ZERO_LEASES` | opn_edge_decommission | true | 
| `OPNSENSE_CARP_PASSWORD` /`OPNSENSE_CARP_VHID_BASE` | opn_ha | empty / 1 | 
| `OPNSENSE_HASYNC_USERNAME` /`OPNSENSE_HASYNC_PASSWORD` | opn_ha | root / empty | 

| Path | Purpose | 
|---|---|
| `collections/opnsense/roles/` | the nine roles | 
| `collections/opnsense/galaxy.yml` | collection manifest | 
| `collections/opnsense/meta/runtime.yml` | required Ansible version | 
| `collections/opnsense/changelogs/` | release history | 
| `shared/netspec/zones.yml` | zones and trunks | 
| `shared/netspec/phases.yml` | phase → capability map | 
| `shared/netspec/firewall_matrix.yml` | sequenced inter-zone policy | 
| `deploy/opnsense/stargate.yml` | executable site entrypoint | 
| `deploy/opnsense/site.yml` | generic orchestration | 
| `deploy/opnsense/bootstrap.yml` | API-key mint (SSH/root) | 
| `deploy/opnsense/SPIKEBUSTING.md` | validation ceremony with manual steps | 
| `deploy/opnsense/vars/*_values.yml` | per-site values | 
| `deploy/opnsense/inventory/` | API target and connection mode | 

Each role maps to one or more OPNsense API controllers via the upstream
`oxlorg.opnsense` modules. The authoritative mapping is the role's
`tasks/main.yml` — the module name encodes the controller. Roles that use
`oxlorg.opnsense.raw` name the controller explicitly in the task's `controller:`
parameter. The raw API pattern extends the collection's reach to model fields
not exposed by typed modules, including Unbound `advanced.*`, Kea
`dhcpv4/set_subnet/<uuid>`, and Unbound `acls.*` / `dnsbl.*`.

| Component | Requirement | 
|---|---|
| `ansible-core` | `>= 2.17` | 
| Python (controller) | `>= 3.12` | 
| `oxlorg.opnsense` | pinned exact version (see `galaxy.yml` dependencies) | 
| OPNsense | 26.1 or compatible release | 

*This document is the canonical reference for `containercraft.opnsense`. For the
collection's machine-readable metadata see `galaxy.yml` and `meta/runtime.yml`;
for release history see `changelogs/`.*
