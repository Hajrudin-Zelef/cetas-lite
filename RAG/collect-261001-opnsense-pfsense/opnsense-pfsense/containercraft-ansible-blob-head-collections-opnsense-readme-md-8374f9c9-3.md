---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9-3
title: "1. Install this collection and its upstream dependency into a playbook-adjacent"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/containercraft-ansible-blob-head-collections-opnsense-readme-md-8374f9c9.md
source_anchor: ""
source_lines: [296, 451]
sha256: 2be6479549b5006d4043da7e58cc58b6da814a6b1f31f5d6a396283e7c45ed93
---

# 1. Install this collection and its upstream dependency into a playbook-adjacent

```
# Mint an API key over SSH as root, write it to the deployment's .env.
ansible-playbook bootstrap.yml \
  -i inventory/bootstrap.yml \
  -e ansible_user=<wheel-user> \
  -e opn_bootstrap_apikey_user=<api-key-user> \
  --ask-pass --ask-become-pass
# Rotate an existing key (mints new, deletes prior keys for the user).
ansible-playbook bootstrap.yml \
  -i inventory/bootstrap.yml \
  -e ansible_user=<wheel-user> \
  -e opn_bootstrap_apikey_user=<api-key-user> \
  -e opn_bootstrap_apikey_force=true \
  --ask-pass --ask-become-pass
```
💡 `--check` and `--tags connect` never write to the firewall and are safe to
run at any time. They are the two commands an operator runs most.


Configuration flows from three sources, lowest precedence first: role defaults
(generic), then the environment and the deployment's values file (site-specific),
then command-line `-e` overrides (one-off). Secrets always come from the
environment, never from a committed file.

| Variable | Environment | Default | Purpose | 
|---|---|---|---|
| `opnsense_firewall_address` | `OPNSENSE_FIREWALL` | empty | API host/IP | 
| `opnsense_api_key` | `OPNSENSE_API_KEY` | empty | API key | 
| `opnsense_api_secret` | `OPNSENSE_API_SECRET` | empty | API secret | 
| `opnsense_api_port` | `OPNSENSE_API_PORT` | `443` | API port | 
| `opnsense_ssl_verify` | `OPNSENSE_SSL_VERIFY` | `false` | validate TLS cert | 
| `opnsense_api_timeout` | — | `30` | request timeout (seconds) | 
| `opnsense_api_retries` | — | `2` | retry count on transient errors | 

`opnsense_api_args` is a convenience mapping of the connection values, used by
tasks that call modules outside the credential-injecting action group (the `raw`
module and `ansible.builtin.uri`).

| Variable | Environment | Default | Purpose | 
|---|---|---|---|
| `network_phase` | `OPNSENSE_NETWORK_PHASE` | `coexist` | migration phase | 
| `opnsense_managed_tag` | `OPNSENSE_MANAGED_TAG` | `ansible-managed` | description prefix on managed objects; also the idempotency key | 
| `opnsense_reload` | — | `false` | per-object reload discipline | 

| Variable | Default | Purpose | 
|---|---|---|
| `opnsense_apply_identity` | `false` | run the initialsetup wizard (off by default) | 
| `opnsense_hostname` | `opnsense` | hostname | 
| `opnsense_domain` | `home.arpa` | domain | 
| `opnsense_timezone` | `UTC` | timezone | 
| `opnsense_wan_ipv4_type` | `dhcp` | `dhcp` \|`static` \|`pppoe` | 
| `opnsense_wan_block_private` | `true` | block RFC1918 on WAN | 
| `opnsense_wan_block_bogons` | `true` | block bogons on WAN | 
| `opnsense_tunables` | CARP hardening | sysctl name→value map | 

The default tunables configure CARP HA observability:
`net.inet.carp.senderr_demotion_factor: 240` (demote master on NIC send errors)
and `net.inet.carp.log: 2` (log state transitions). IPv4/IPv6 forwarding is
already enabled by OPNsense default and is not configured explicitly.

**Blast radius.** `opnsense_apply_identity: true` runs OPNsense's
first-boot wizard, which is the only API path to WAN configuration in current
releases and re-applies the wizard's settings. Enable it deliberately and
confirm the WAN settings match the upstream link first.

⚠️ 

| Variable | Default | Purpose | 
|---|---|---|
| `opnsense_dns_hardening` | (dict) | Unbound `advanced.*` settings via raw API | 
| `opnsense_dns_acl_default_action` | `refuse` | ACL default for unknown sources | 
| `opnsense_dns_acls` | per-zone allow | explicit ACL entries | 
| `opnsense_dns_dnsbl_enabled` | `true` | enable DNSBL threat feeds | 
| `opnsense_dns_dnsbl_type` | `[atf, hgz011]` | Abuse.ch ThreatFox + Hagezi TI | 
| `opnsense_dns_cache` | (dict) | cache sizing for bare metal | 

The hardening dict sets identity hiding (`hideidentity`, `hideversion`), DNSSEC
enforcement (`dnssecstripped`, `belownxdomain`), cache poisoning resistance
(`aggressivensec`, `unwantedreplythreshold: 10000000`), privacy
(`qnameminstrict: 0`), rebinding protection (`privateaddress` with full RFC 1918

- special-use ranges, `privatedomain: home.arpa` ), DNSSEC bypass for the
internal zone (`insecuredomain: home.arpa` ), and observability
(`extendedstatistics` ,`logservfail` ,`valloglevel: 1` ).

The cache dict sizes `msgcachesize: 256m`, `rrsetcachesize: 512m`,
`outgoingrange: 8192`, `numqueriesperthread: 4096`, `infracachenumhosts: 20000`.

| Variable | Default | Purpose | 
|---|---|---|
| `opnsense_enforce_dnsmasq_disabled` | `false` | disable default Dnsmasq DHCP before enabling Kea | 
| `opnsense_dhcp_socket_type` | `udp` | UDP sockets (pf-gated); `raw` bypasses pf | 
| `opnsense_dhcp_general_hardening` | (dict) | Kea general settings via raw API | 
| `opnsense_dhcp_subnet_overrides` | (dict per zone) | per-subnet allocator, client-ID, lifetime | 

The general hardening sets `service_sockets_max_retries: 5` and
`service_sockets_retry_wait_time: 5000` so Kea retries interface binding on
startup rather than silently failing.

The subnet overrides configure per-zone policies:

| Zone | `allocator` | `match-client-id` | `valid_lifetime` | 
|---|---|---|---|
| mgmt | `random` | `1` (true) | `7200` (2h) | 
| iot | `random` | `0` (false) | `3600` (1h) | 
| dmz | `random` | `1` (true) | `3600` (1h) | 

`allocator: random` prevents predictable address assignment.
`match-client-id: 0` on the IoT subnet uses MAC-based identification because
IoT devices frequently generate unstable client IDs. Shorter lifetimes on IoT
and DMZ reduce the window for stale address assignments after device removal.

All subnets use `auto_options: false` with explicit `gateway`, `dns`, and
`domain` values. This prevents the ordering race condition where Kea
auto-collects router/DNS options from an interface that does not yet have an IP
assigned, and prevents HA failover from auto-collecting the backup node's
physical IP instead of the CARP VIP.

| Variable | Default | Purpose | 
|---|---|---|
| `opnsense_prepare_trunk_ports` | `false` | remove trunk parents from an existing bridge | 
| `opnsense_bridge_description` | empty | bridge to reconfigure (when preparing) | 
| `opnsense_bridge_keep_members` | `[]` | members to retain | 
| `opnsense_assign_interfaces` | `false` | assign VLAN devices to interface slots | 
| `opnsense_interface_assignments` | `[]` | the assignments (zone, device) | 

**Blast radius.** `opnsense_prepare_trunk_ports` mutates an existing bridge —
a live-network change. It is opt-in for a deliberate, sequenced cutover only.

⚠️ 

| Variable | Default | Purpose | 
|---|---|---|
| `opnsense_require_zero_leases` | `true` | require zero active leases before withdrawing service | 

| Variable | Environment | Default | Purpose | 
|---|---|---|---|
| `opnsense_carp_password` | `OPNSENSE_CARP_PASSWORD` | empty | CARP VHID group password | 
| `opnsense_carp_vhid_base` | `OPNSENSE_CARP_VHID_BASE` | `1` | base VHID; per-zone VHID = base + VLAN | 
| `opnsense_carp_advskew` | — | `0` | 0 = master; host_vars override for backup | 
| `opnsense_hasync_username` | `OPNSENSE_HASYNC_USERNAME` | `root` | config-sync user | 
| `opnsense_hasync_password` | `OPNSENSE_HASYNC_PASSWORD` | empty | config-sync password | 
| `opnsense_hasync_pfsyncversion` | — | `1400` | pfsync protocol (OPNsense 24.7+) | 
| `opnsense_hasync_syncitems` | — | aliases, rules, nat, virtualip, dhcpd, unbound | sections to sync | 

| Variable | Default | Purpose | 
|---|---|---|
| `opn_bootstrap_apikey_user` | `ansible` | OPNsense user to mint the key for | 
| `opn_bootstrap_apikey_force` | `false` | rotate: mint new, delete prior keys | 
| `opn_bootstrap_apikey_env_file` | `{{ playbook_dir }}/.env` | where to write the minted pair | 

The network's shape is declared as data in `shared/netspec/`, and the roles are
pure functions of that data. To change the network, an operator edits data, not
tasks. There are three files: the zones and trunks (`zones.yml`), the phase map
(`phases.yml`, covered in §4), and the inter-zone policy (`firewall_matrix.yml`).

