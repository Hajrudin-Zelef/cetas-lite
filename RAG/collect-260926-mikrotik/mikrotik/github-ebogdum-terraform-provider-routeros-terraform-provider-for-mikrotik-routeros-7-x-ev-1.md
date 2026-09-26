---
id: collect-260926-mikrotik/mikrotik/github-ebogdum-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros-7-x-ev-1
title: "1. Build the provider binary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["lora"]
source: docs/RAG/lot-mikrotik/RouterOS/github-ebogdum-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros-7-x-every-menu-e.md
source_anchor: ""
source_lines: [1, 188]
sha256: fb99c0d04fcea7f5a0c633ca417c82ade503a980d0a4afcbea6636c34c323d61
---

# 1. Build the provider binary

Manage MikroTik **RouterOS 7.x** devices as code with Terraform. Complete
coverage of every device menu -- **420 menus, 182 resources, 276 data sources,
75 actions, 3289 properties** -- generated from a schema validated
property-by-property against a live router.

Manage one MikroTik router or an entire fleet from a single Terraform configuration. Apply firewall rules in deterministic order, import existing device state, detect and correct out-of-band drift, and ship secrets safely through the plugin framework's sensitive-value handling.

**Keywords:** terraform-provider-mikrotik, terraform-provider-routeros,
mikrotik terraform, routeros terraform, mikrotik infrastructure as code,
routeros REST API, mikrotik automation, network as code, mikrotik CHR,
mikrotik fleet management, terraform firewall mikrotik, mikrotik IPsec
terraform.


| Feature | This provider | 
|---|---|
| RouterOS API | REST (HTTPS) | 
| Menu coverage | 420 menus (every menu surfaced over REST) | 
| Resources / data sources / actions | 182 / 276 / 75 | 
| Multi-router from one provider block | Yes (named map, no provider aliases) | 
| Deterministic firewall ordering | Yes ( `position` integer; stable across destroy/recreate) | 
| Lockout safety guards | Firewall, user, user-group, mac-server | 
| Sensitive field redaction | 28 properties marked sensitive | 
| Out-of-band drift detection | Yes (verified end-to-end) | 
| Terraform import | Yes ( `<router>/<.id>` format) | 
| Schema source | Live device + WebFig skin files + Confluence docs + per-property device validation | 
| Plugin framework | terraform-plugin-framework (v1.19.0) | 
| Minimum Terraform | 1.4 | 
| Minimum RouterOS | 7.1 (REST API requirement) | 

```
terraform {
  required_providers {
    routeros = {
      source  = "ebogdum/routeros"
      version = "~> 2.0"
    }
  }
}
provider "routeros" {
  host     = "https://192.0.2.1"
  username = "admin"
  password = var.routeros_password
  insecure = true # set to false in production with a real cert
}
resource "routeros_ip_address" "lan" {
  address   = "192.168.88.1/24"
  interface = "bridge1"
  comment   = "Managed by Terraform"
}
resource "routeros_ip_firewall_filter" "allow_established" {
  chain            = "input"
  action           = "accept"
  connection_state = "established,related"
  position         = 100
  comment          = "Managed by Terraform"
  lockout_ack      = true
}
```
`terraform init && terraform apply` and you are done.

```
# 1. Build the provider binary
make build
# 2. Tell Terraform to use the local binary instead of the registry
cat > ~/.terraformrc <<EOF
provider_installation {
  dev_overrides {
    "ebogdum/routeros" = "$(pwd)/bin"
  }
  direct {}
}
EOF
```
Manage every router in your network from a single Terraform configuration. No provider aliases, no duplicate blocks -- just a named map.

```
provider "routeros" {
  routers = {
    core = {
      host     = "https://10.0.0.1"
      username = "admin"
      password = var.core_password
      insecure = true
    }
    edge_se = {
      host     = "https://10.0.1.1"
      username = "admin"
      password = var.edge_se_password
      insecure = true
    }
    edge_nw = {
      host     = "https://10.0.2.1"
      username = "admin"
      password = var.edge_nw_password
      insecure = true
    }
  }
}
# Push the same identity policy to every router with one resource.
resource "routeros_system_identity" "label" {
  for_each = toset(["core", "edge_se", "edge_nw"])
  router   = each.key
  name     = each.key
}
# Cross-router data flow: feed core's address-list into edge_se's firewall.
data "routeros_ip_firewall_address_list" "core_lan" {
  router = "core"
}
resource "routeros_ip_firewall_filter" "edge_to_core" {
  router           = "edge_se"
  chain            = "forward"
  action           = "accept"
  src_address_list = data.routeros_ip_firewall_address_list.core_lan.records[0].list
  position         = 200
  comment          = "Allow core LAN inbound"
}
```
Omit `router =` to target the default -- the entry named `default`, or the
first router in sorted order if there is no `default`.

Every menu the REST API surfaces is mapped. Categories include:

- **IP** : address, route, pool, ARP, DHCP server/client/relay, DNS, firewall
(filter, NAT, mangle, raw, address-list, layer7-protocol, service-port),
IPsec (incl. policy-group, QKD keys), hotspot (incl. service-port), proxy,
service, neighbor (incl. discovery-settings), kid-control, traffic-flow
(incl. IPFIX), media, cloud (incl. advanced), DNS (static, forwarders,
adlist), settings.
- **IPv6** : address, route, pool, DHCP client/server/relay (incl. relay
options), firewall family, neighbor discovery (incl. default prefix),
settings.
- **Interface** : bridge (port, vlan, msti, mst-config, settings, filter, nat,
calea, host), bonding, vlan, vrrp, wireguard (peers), wifi (wave2 +
legacy), 6to4, eoip, eoipv6, gre, gre6, ipip, ipipv6, l2tp client/server,
list (members), macsec, macvlan, ovpn client/server, ppp, pppoe, pptp,
sstp, veth, vrrp, vti.
- **Routing** : BGP (connection, template), OSPF (instance, area,
interface-template, static-neighbor), RIP (instance, interface, neighbor),
ISIS (instance, interface-template), filter (rule, select-rule, chain),
BFD, RPKI (session), IGMP-Proxy, PIM-SM, table, rule.
- **System** : identity, clock, NTP client/server, scheduler, script,
package (incl. local-update mirror), resource (incl. IRQ affinity, USB
settings), health settings, routerboard (settings and the mode/WPS/reset
button bindings), backup, logging (action), users (user, group, ssh-keys,
aaa), notes, history, watchdog, leds.
- **Tools** : bandwidth-test, netwatch, fetch, sniffer, traffic-generator
(with packet templates), email, sms, romon, mac-server, snmp-get,
snmp-walk.
- **Certificate** : full lifecycle (sign, export, import, SCEP, ACME, CRL).
- **Queue** : simple, tree, type, interface.
- **SNMP** : community, send-trap.
- **Container** ,**LCD** ,**IoT/LoRa** ,**disk** ,**environment** ,**RADIUS** .

Every resource has a matching data source for querying existing state.

The provider refuses to apply changes that would obviously sever management
access. Each guard is conservative and per-resource -- set
`lockout_ack = true` on the specific resource to override.

| Resource | What it refuses | 
|---|---|
| `routeros_ip_firewall_filter` ,`routeros_ipv6_firewall_filter` | `chain=input | 
| `routeros_user` | Deleting or disabling the last admin account | 
| `routeros_user_group` | Removing `api` ,`web` ,`winbox` ,`ssh` ,`password` ,`policy` , or`write` from the`full` group | 
| `routeros_tool_mac_server` | Emptying `allowed-interface-list` (kills MAC-Winbox recovery) | 

Sensitive properties (passwords, secrets, keys, OTP, SIM PIN -- 28 total
across the schema) are marked `Sensitive: true`. Terraform redacts them in
plan output, state files, and CLI display.

The provider's Read implementation queries the device directly on every plan. If a row is changed outside Terraform (via Winbox, CLI, or another tool), the next plan shows the drift and the following apply reconciles it.

This is verified by an explicit acceptance test (`TestAccDriftCorrectionIPAddress`)
that:

1. Creates a row through Terraform.
2. Mutates that row's `comment` field via raw REST, outside Terraform.
3. Re-applies the original config.
4. Asserts the field is restored.

Every collection-style menu (`/ip/address`, `/ip/firewall/filter`, ...) is a
resource. Every singleton menu (`/system/identity`, `/ip/dns`, ...) is a
resource that maps to a single device-wide value.

Every collection-style menu is also a data source -- query existing rows by
property filter, project specific fields via `proplist`.

