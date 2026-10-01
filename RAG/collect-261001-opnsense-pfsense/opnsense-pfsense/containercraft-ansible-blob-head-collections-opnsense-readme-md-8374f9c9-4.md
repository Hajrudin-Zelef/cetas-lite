---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9-4
title: "1. Install this collection and its upstream dependency into a playbook-adjacent"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9.md
source_anchor: ""
source_lines: [452, 641]
sha256: 528dc95f6bf2a7b8b6b1e944d8648ec9e028fd1e1bdb3cfa208d81abe5c2a76c
---

# 1. Install this collection and its upstream dependency into a playbook-adjacent

```
graph TB
    subgraph "Trunks (physical ports carrying tagged VLANs)"
        QNAP["qnap — em2<br/>cluster/infra power domain"]
        POE["poe — em3<br/>AP/WiFi power domain"]
    end
    subgraph "Zones (tagged VLANs on the trunks)"
        MGMT["mgmt — vlan10<br/>10.10.0.0/24"]
        IOT["iot — vlan20<br/>10.20.0.0/24"]
        DMZ["dmz — vlan30<br/>10.30.0.0/24"]
        SYNC["sync — vlan4000<br/>10.40.0.0/24 (HA pfsync)"]
        CL["cluster — untagged<br/>10.0.0.0/24 (existing LAN)"]
    end
    MGMT --> QNAP
    MGMT --> POE
    IOT --> QNAP
    IOT --> POE
    DMZ --> QNAP
    DMZ --> POE
    SYNC --> QNAP
    CL --> QNAP
```
    A **trunk** is a physical OPNsense port carrying tagged zone VLANs to one switch
or power domain. A **zone** is a logical network that declares which trunks it
rides; the interfaces role creates the zone's VLAN sub-interface on each listed
trunk parent. This is how one zone can be present on two independent power
domains.

Each zone is a dictionary entry with a consistent shape:

```
# shared/netspec/zones.yml (one zone)
netspec:
  zones:
    mgmt:
      enabled: true
      vlan: 10
      tagged: true
      trunks: [qnap, poe]      # rides both power domains
      role: zone               # zone | cluster | transit | sync
      subnet: "10.10.0.0/24"
      gateway: "10.10.0.1"
      dhcp: true
      dhcp_pool: "10.10.0.100 - 10.10.0.200"
      description: "MGMT - infrastructure management"
```
| Zone field | Meaning | 
|---|---|
| `enabled` | include the zone in a run | 
| `vlan` | 802.1q tag (0/untagged for the existing flat LAN) | 
| `tagged` | carried as a tagged VLAN vs. the native/untagged segment | 
| `trunks` | which trunk parents the zone VLAN is created on | 
| `role` | `zone` (client subnet),`cluster` (existing LAN),`transit` (routed core),`sync` (HA state) | 
| `subnet` /`gateway` | the zone's network and its gateway address | 
| `dhcp` /`dhcp_pool` | whether Kea serves the zone, and the lease range | 

Top-level netspec fields carry shared values: `cluster_dns` (the resolver to
forward internal queries to), `cluster_domain` (the internal split-horizon zone),
`wan_interface`, `native_blackhole_vlan` (a deny-all native VLAN per trunk), and
`static_hosts` (Unbound reservations).

Inter-zone policy is a table, not procedural rules. Each row carries an explicit sequence number for deterministic pf ordering and is rendered into one firewall rule, matched on its description (which makes it idempotent — see §15).

```
# shared/netspec/firewall_matrix.yml (excerpts)
firewall_matrix:
  dhcp_allow:
    - {seq: 10, from: mgmt, to: any, action: pass, proto: udp, src: any, port: 67, desc: "dhcp server mgmt"}
  isolation:
    - {seq: 100, from: iot, to: mgmt, action: block, proto: any, desc: "iot deny mgmt"}
  allows:
    - {seq: 203, from: mgmt, to: cluster_dns, action: pass, proto: "TCP/UDP", port: 53, desc: "mgmt to dns"}
  wan_egress:
    - {seq: 400, from: mgmt, to: wan, action: pass, proto: any, desc: "mgmt to internet"}
```
| Group | Sequence range | Intent | 
|---|---|---|
| `dhcp_allow` | 10–19 | DHCP server allow (UDP 67); required because `fw_rules: false` on Kea | 
| `isolation` | 100–199 | explicit inter-zone denies; block rules before pass on the same interface | 
| `allows` | 200–299 | explicit permits (mgmt lateral, DNS port 53 per zone) | 
| `wan_egress` | 400–499 | per-zone egress to the internet | 

All rules use `quick: true` (pf first-match semantics) and `direction: in`
(filter on the source interface inbound).

```
graph LR
    MGMT["mgmt"]
    IOT["iot"]
    DMZ["dmz"]
    CL["cluster"]
    DNS(["cluster_dns:53"])
    WAN(["WAN / internet"])
    MGMT -->|pass| WAN
    IOT -->|pass| WAN
    DMZ -->|pass| WAN
    CL -->|pass| WAN
    MGMT -->|pass| IOT
    MGMT -->|pass| DMZ
    MGMT -->|pass| CL
    MGMT -->|pass 53| DNS
    IOT -->|pass 53| DNS
    DMZ -->|pass 53| DNS
    IOT -.->|block| MGMT
    IOT -.->|block| CL
    DMZ -.->|block| MGMT
    DMZ -.->|block| CL
    style WAN fill:#1f3a5f,color:#fff
    style DNS fill:#2d5016,color:#fff
```
    Solid edges are permits; dashed edges are explicit denies.

A deployment overrides the shared netspec without copying it. `netspec_overrides`
is deep-merged over `shared/netspec/zones.yml`, so a deployment sets only the
leaves that differ:

```
# vars/<deploy>_values.yml
netspec_overrides:
  cluster_dns: "10.0.0.51"
  zones:
    mgmt: {subnet: "10.10.0.0/24", gateway: "10.10.0.1"}
```
💡 The shared netspec is the single source of truth. Overrides should contain only true deviations; mirroring the defaults into an override silently decouples the deployment from future changes to the shared data.


The collection contains nine roles. One (`opn_bootstrap_apikey`) runs over SSH and
is invoked by its own playbook; the other eight are API-driven and composed by
`site.yml` in dependency order.

```
graph LR
    CONN["opn_connect"] --> SYS["opn_system"]
    SYS --> IF["opn_interfaces"]
    IF --> DNS["opn_dns"]
    DNS --> DHCP["opn_dhcp"]
    DHCP --> FW["opn_firewall"]
    FW --> HA["opn_ha"]
    HA --> DEC["opn_edge_decommission"]
    style CONN fill:#1f3a5f,color:#fff
    style HA fill:#444,color:#fff
    style DEC fill:#5c2d00,color:#fff
```
    Read-only API probe. Issues a single list request and asserts the response
carries data. Tagged `connect` and `always`.

Sysctl tunables via `core/tunables` GET-diff-POST (real idempotency). Identity
and WAN via `core/initialsetup` wizard (gated by `opnsense_apply_identity`).
Default tunables: `net.inet.carp.senderr_demotion_factor: 240`,
`net.inet.carp.log: 2`.

`opnsense_apply_identity: false` until a
deliberate identity change is intended.

⚠️ Identity/WAN run the wizard. Leave

VLAN sub-interfaces on trunk ports: blackhole VLAN 3999 per trunk, zone VLANs per netspec. Optional trunk-port preparation (bridge mutation) and interface slot assignment, both off by default.

🔬 **Deeper.** The OPNsense MVC API does not expose per-interface IP
configuration on opt* interfaces. After VLAN creation, the operator assigns
IPs via the OPNsense web UI before advancing to the `migrating` phase. This
is documented in the spikebusting ceremony (`SPIKEBUSTING.md`, Phase 3).


Hardened Unbound resolver with split-horizon forwarding and DNS-layer threat blocking:

1. **General settings** via`oxlorg.opnsense.unbound_general` : enabled, port 53,
DNSSEC, static mapping registration for Kea reservations.
2. **Security hardening** via raw API to`unbound.advanced.*` : identity hiding,
DNSSEC enforcement, rebinding protection, cache poisoning detection, cache
sizing, observability (extended statistics, SERVFAIL logging, DNSSEC
validation logging).
3. **ACL default action** set to`refuse` via raw API to`unbound.acls.*` . All
unknown sources are refused; explicit per-zone`allow` ACLs for mgmt, iot,
dmz, and cluster subnets.
4. **DNSBL** via raw API to`unbound.dnsbl.*` : Abuse.ch ThreatFox IOC and Hagezi
Threat Intelligence Feeds enabled by default.
5. **Forward zone** :`home.arpa` → CoreDNS at`10.0.0.51` , with`privatedomain: home.arpa` (allow RFC 1918 in responses) and`insecuredomain: home.arpa` (bypass DNSSEC for unsigned CoreDNS responses).
6. **Static host reservations** for operator-defined entries.

Kea DHCPv4 with per-zone scopes, phase-gated (`zone_dhcp`, off at `coexist`):

1. Dnsmasq disabled (conflicts with Kea on the same interfaces).
2. Kea enabled with `socket_type: udp` (pf-gated; raw bypasses pf),`fw_rules: false` (DHCP allow rules in the firewall matrix), and socket
retry configuration via raw API (`service_sockets_max_retries: 5` ).
3. Per-zone subnets with `auto_options: false` and explicit routers/dns/domain.
4. Per-subnet hardening via raw API: `allocator` ,`match-client-id` ,`valid_lifetime` set per zone after the oxlorg module creates the subnet.

Zone aliases, DHCP allow rules, and the inter-zone policy matrix — all savepoint-wrapped:

