---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9-5
title: "1. Install this collection and its upstream dependency into a playbook-adjacent"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9.md
source_anchor: ""
source_lines: [642, 808]
sha256: 685d92136e1676218f8a4f6f51e8e0247fe5cb665a3421e215dc45dc63288345
---

# 1. Install this collection and its upstream dependency into a playbook-adjacent

1. Network aliases (one per zone) and a `cluster_dns` host alias, created and
applied before the savepoint.
2. Savepoint create on the filter controller.
3. Flattened matrix rows applied as rules with explicit `sequence` ,`quick: true` ,`direction: in` . DHCP allow rules (seq 10–12, UDP 67) run first. Isolation
blocks (seq 100–103) run before allows (seq 200–205) on the same interface.
WAN egress (seq 400–403) last.
4. Reload, verify API reachable, cancel rollback (commit). If verification fails or the run is interrupted, the 60-second server timer auto-reverts.

```
sequenceDiagram
    participant R as opn_firewall
    participant FW as OPNsense
    R->>FW: create zone aliases (reload:false)
    R->>FW: reload aliases
    R->>FW: savepoint create (filter)
    R->>FW: apply matrix rules (reload:false)
    R->>FW: reload rules
    R->>FW: list rules (verify reachable)
    alt verified
        R->>FW: cancel_rollback (commit)
    else not reached in time
        FW-->>FW: auto-revert filter section
    end
```
    The collection does not create masquerade NAT rules. OPNsense's automatic outbound NAT already masquerades internal zones to the WAN address. For HA deployments, the operator switches to manual outbound NAT with the CARP VIP as the translation target.

🔬 **Deeper — automation namespace.** All rules are in `Firewall → Automation`
(MVC namespace), isolated from `Firewall → Rules` (legacy namespace).
Automation rules process before legacy rules on the same interface.


CARP VIPs, HA configuration sync, and explicit sync trigger. Gated by
`opnsense_ha_enabled`. Creates a CARP VIP per zone gateway with a unique VHID
(`opnsense_carp_vhid_base + vlan_tag`), configures hasync with
`pfsync_version: 1400` (OPNsense 24.7+), and on the master triggers
`core/hasync_status/restart_all` to push configuration to the backup.

The per-node identity (`opnsense_carp_advskew`, 0 for master, higher for backup)
comes from inventory `host_vars`; everything else — the VIP addresses, VHIDs, and
sync items — is derived from the netspec and the HA variables.

```
graph TB
    subgraph "Master (advskew 0)"
        M["OPNsense A"]
    end
    subgraph "Backup (advskew > 0)"
        B["OPNsense B"]
    end
    VIP(["CARP VIP per zone<br/>e.g. 10.10.0.1 (mgmt)<br/>VHID = base + vlan"])
    CLIENTS["zone clients<br/>(gateway = VIP)"]
    CLIENTS --> VIP
    VIP -. elected .-> M
    VIP -. standby .-> B
    M <-->|"pfsync state<br/>(SYNC zone, 10.40.0.0/24)"| B
    M -->|"config sync<br/>(hasync / XMLRPC)"| B
    style M fill:#2d5016,color:#fff
    style VIP fill:#1f3a5f,color:#fff
```
    
⚠️ HA requires a second node and the dedicated SYNC segment. Enable only when the peer and sync zone are present.

Withdraws edge DHCP/DNS from the legacy cluster segment. Evidence-gated: queries
active Kea leases via `kea/leases4/search` and asserts `stats.active == 0`.

**Blast radius.** This role removes service from the live LAN. The evidence
gate is a safety, not a formality.

⚠️ 

SSH/root API key mint. Renders a PHP mint script to the box, runs it as root
(the script uses OPNsense's auth model to create and persist a key), removes
the script, parses the result, and writes the key/secret to the deployment's
`.env`. Username sanitized to `[a-zA-Z0-9._-]` at template render time.
Idempotent and non-accumulating: without `force` it does nothing when a key
exists; with `force` it mints a new key and deletes the prior keys for the
user, so exactly one key remains. Full lifecycle in §12.

| Likely cause | Check | Fix | 
|---|---|---|
| Credentials not in environment | confirm `OPNSENSE_API_KEY` /`SECRET` exported | `direnv allow` , or source`.env` | 
| Key not on the box | mint or rotate (§12) | re-run bootstrap with `force` | 
| `lookup('env')` empty | ansible not inheriting shell env | use `direnv exec <dir> ansible-playbook ...` | 

💡 Prove a credential independently:
`curl -k -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" https://<fw>/api/core/firmware/status`


Collections not installed. Run:
`ansible-galaxy collection install -r requirements.yml -p collections`

Same as §10.2. Confirm `collections/ansible_collections/containercraft/opnsense`
exists.

`lookup('env', ...)` reads the ansible process environment. Launch from the
configured deploy shell or use `direnv exec`.

The savepoint auto-reverted because verification did not complete in time. Re-run the role.

The built-in `home.arpa` static local zone is intercepting queries before the
forward zone. Verify `insecuredomain` and `privatedomain` are set to `home.arpa`
in `unbound.advanced`. The built-in zone serves `NS localhost` and
`SOA localhost` — if these appear in responses, the hardening raw API task did
not apply correctly.

Verify: (1) `phase.zone_dhcp` is `true` (phase is `migrating` or later),
(2) `opnsense_enforce_dnsmasq_disabled` is `true` and Dnsmasq is off,
(3) interface IPs are assigned on zone gateways (manual step),
(4) DHCP allow rules (seq 10–12) exist in the firewall,
(5) Kea `socket_type` is `udp` (if `raw`, pf rules do not gate DHCP traffic).

- **One change per iteration.** Run a single tag, or flip a single toggle.
- **Dry-run before every apply.**`--check --diff` shows exactly what changes.
- **Stay in `coexist` until interfaces have IPs and switches are configured.**

```
graph LR
    SCOPE["pick ONE scope<br/>(a tag or a toggle)"] --> DRY["--check --diff"]
    DRY --> READ{"diff correct?"}
    READ -->|no| FIX["adjust values / data"]
    FIX --> DRY
    READ -->|yes| APPLY["apply (same tag)"]
    APPLY --> NEXT["next scope"]
    NEXT --> SCOPE
    style DRY fill:#1f3a5f,color:#fff
    style APPLY fill:#2d5016,color:#fff
```
    ```
# 0. Prove connectivity and credentials.
./stargate.yml --tags connect
# 1. System baseline (CARP tunables; identity stays off).
./stargate.yml --tags system --check --diff
./stargate.yml --tags system
# 2. Interfaces (zone VLANs + blackhole).
./stargate.yml --tags interfaces --check --diff
./stargate.yml --tags interfaces
# 3. DNS (hardened Unbound, DNSBL, ACLs, forward zone).
./stargate.yml --tags dns --check --diff
./stargate.yml --tags dns
# 4. Firewall (DHCP allow + isolation + allows + egress, savepoint-wrapped).
./stargate.yml --tags firewall --check --diff
./stargate.yml --tags firewall
# 5. Assign interface IPs in the OPNsense web UI (manual step — see SPIKEBUSTING.md Phase 3).
# 6. Configure switches with VLAN trunks (manual step — see SPIKEBUSTING.md Phase 4).
# 7. Advance phase and enable zone DHCP.
#    Set OPNSENSE_NETWORK_PHASE=migrating and opnsense_enforce_dnsmasq_disabled=true.
./stargate.yml --tags dhcp --check --diff
./stargate.yml --tags dhcp
```
A deployment's values file disables every optional capability so a full run is
minimal. Each toggle is a single iteration: set it, re-run that role's tag
with `--check --diff`, read the diff, apply.

```
# vars/<deploy>_values.yml (capability toggles, all disabled to start)
opnsense_apply_identity: false        # flip for opn_system identity/WAN
opnsense_prepare_trunk_ports: false   # flip for a deliberate trunk cutover
opnsense_assign_interfaces: false     # flip once per-interface IPs are handled
opnsense_enforce_dnsmasq_disabled: false  # flip when bringing up Kea
opnsense_ha_enabled: false            # flip when the HA peer exists
opnsense_require_zero_leases: true    # safety gate — leave on
```
The deployment includes a `SPIKEBUSTING.md` ceremony document with checkbox
verification for every phase, including manual steps (interface IP assignment
in the web UI, switch VLAN configuration). The ceremony covers the full path
from `coexist` through `migrating` with zone isolation verification.

