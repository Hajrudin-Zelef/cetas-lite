---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/6ixfalls-terraform-provider-opnsense-0e0cdd56-1
title: "1. Configure Terraform to use the provider"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/6ixfalls-terraform-provider-opnsense-0e0cdd56.md
source_anchor: ""
source_lines: [1, 155]
sha256: 19c79be14ef7a4bc799d3fedbd8e7163d24c789e9c32b78a3745a62edb410e13
---

# 1. Configure Terraform to use the provider

This Terraform provider enables management of various configs and resources within OPNsense®.

Warning

This provider is under active development and makes no guarantee of stability. Breaking changes to resource and data source schemas will occur as needed until v1.0. **It is not recommended to use this provider in production environments.**

```
# 1. Configure Terraform to use the provider
terraform {
  required_providers {
    opnsense = {
      source  = "browningluke/opnsense"
      version = "~> 0.16"
    }
  }
}
# 2. Configure the OPNsense provider with API credentials
provider "opnsense" {
  uri        = "https://opnsense.example.com"
  # Either reference the API credentials literally
  api_key    = "<api key>"
  api_secret = "<api password>"
  # Or specify them with environment variables
  # export OPNSENSE_API_KEY="<api key>"
  # export OPNSENSE_API_SECRET="<api key>"
}
# 3. Create resources - example: firewall rule
resource "opnsense_firewall_filter" "allow_https" {
  enabled     = true
  description = "Allow inbound HTTPS traffic"
  interface = {
    interface = ["wan"]
  }
  filter = {
    action    = "pass"
    direction = "in"
    protocol  = "TCP"
    source = {
      net = "any"
    }
    destination = {
      net  = "192.168.1.100"
      port = "https"
    }
    log = true
  }
}
```
Version 1.0 will be released once the provider achieves feature-parity with the **Core** OPNsense API and all resources have comprehensive acceptance tests (see Current API Coverage). Plugin resources will be added as requested (at a lower priority than requests for Core resources). There is no Plugin API converage requirement for v1.

v1 represents the first release where resource and data source schemas will be guaranteed to be stable, and breaking changes to these schemas will be forbidden. Any updates to these schemas will following appropriate SemVer conventions. Until v1.0 is reached, **schemas are subject to change as needed** to improve usability and align with best practices. Users should always check the release notes when upgrading between pre-v1.0 versions to understand any breaking changes that may affect their configurations.

- **Terraform Registry Documentation** - Full resource and data source reference
- **Examples** - Working examples for all resources

Interested in contributing? Please see our Contributing Guide for development setup, testing requirements, and guidelines.

This provider is actively expanding to cover the OPNsense API. The tables below contain the current status of said coverage.

- ✅ = Fully implemented
- 🚧 = Missing acceptance tests
- ❌ = Not implemented

| Module/Controller/Resource | Resource | Data Source | 
|---|---|---|
| `Auth/Group` | ❌ | ❌ | 
| `Auth/Priv` | ❌ | ❌ | 
| `Auth/User` | ❌ | ❌ | 
| `Captiveportal/Settings` | ❌ | ❌ | 
| `Captiveportal/Service/Template` | ❌ | ❌ | 
| `Captiveportal/Settings/Zone` | ❌ | ❌ | 
| `Core/Hasync` | ❌ | ❌ | 
| `Core/Snapshots` | ❌ | ❌ | 
| `Core/Tunables` | ❌ | ❌ | 
| `Cron/Job` | ❌ | ❌ | 
| `Dhcrelay/Settings` | ❌ | ❌ | 
| `Dhcrelay/Settings/Dest` | ❌ | ❌ | 
| `Dhcrelay/Settings/Relay` | ❌ | ❌ | 
| `Diagnostics/Interface` | ❌ | 🚧 | 
| `Diagnostics/Lvtemplate` | ❌ | ❌ | 
| `Diagnostics/Lvtemplate/Item` | ❌ | ❌ | 
| `Dnsmasq/Settings` | ❌ | ❌ | 
| `Dnsmasq/Settings/Boot` | ❌ | ❌ | 
| `Dnsmasq/Settings/Domain` | ❌ | ❌ | 
| `Dnsmasq/Settings/Host` | ✅ | ✅ | 
| `Dnsmasq/Settings/Option` | ❌ | ❌ | 
| `Dnsmasq/Settings/Range` | ❌ | ❌ | 
| `Dnsmasq/Settings/Tag` | ❌ | ❌ | 
| `Firewall/Alias` | ✅ | ✅ | 
| `Firewall/Category` | 🚧 | 🚧 | 
| `Firewall/Filter` | ✅ | ✅ | 
| `Firewall/Group` | ❌ | ❌ | 
| `Firewall/NPTv6` | ❌ | ❌ | 
| `Firewall/Source NAT` | ✅ | ✅ | 
| `Firewall/One-to-One NAT` | ✅ | ✅ | 
| `Interfaces/Bridge` | ❌ | ❌ | 
| `Interfaces/Gif` | ❌ | ❌ | 
| `Interfaces/Gre` | ❌ | ❌ | 
| `Interfaces/Lagg` | ❌ | ❌ | 
| `Interfaces/Loopback` | ❌ | ❌ | 
| `Interfaces/Neighbor` | ❌ | ❌ | 
| `Interfaces/Overview` |  | ✅ | 
| `Interfaces/Vip` | ✅ | ✅ | 
| `Interfaces/Vlan` | ✅ | ✅ | 
| `Interfaces/Vxlan` | ❌ | ❌ | 
| `Ipsec/Settings` | ❌ | ❌ | 
| `Ipsec/Connections/Local` | ✅ | ❌ | 
| `Ipsec/Connections/Remote` | ✅ | ❌ | 
| `Ipsec/Connections/Child` | ✅ | ❌ | 
| `Ipsec/Connections/Connection` | ✅ | ❌ | 
| `Ipsec/KeyPairs` | ❌ | ❌ | 
| `Ipsec/ManualSpd` | ❌ | ❌ | 
| `Ipsec/Pools` | ❌ | ❌ | 
| `Ipsec/Psk` | ✅ | ❌ | 
| `Ipsec/Vti` | ✅ | ❌ | 
| `Kea/CtrlAgent` | ❌ | ❌ | 
| `Kea/Dhcpv4/Peer` | ✅ | ✅ | 
| `Kea/Dhcpv4/Reservation` | ✅ | ✅ | 
| `Kea/Dhcpv4/Subnet` | ✅ | ✅ | 
| `Kea/Dhcpv6/PD Pool` | ✅ | ✅ | 
| `Kea/Dhcpv6/Peer` | ✅ | ✅ | 
| `Kea/Dhcpv6/Reservation` | ✅ | ✅ | 
| `Kea/Dhcpv6/Subnet` | ✅ | ✅ | 
| `Monit/Settings` | ❌ | ❌ | 
| `Monit/Settings/Alert` | ❌ | ❌ | 
| `Monit/Settings/Service` | ❌ | ❌ | 
| `Monit/Settings/Test` | ❌ | ❌ | 
| `Openvpn/Client Overwrites` | ✅ | ✅ | 
| `Openvpn/Instances` | ✅ | ✅ | 
| `Openvpn/Instances/Static Key` | ✅ | ✅ | 
| `Openvpn/Instances/Generate Key` | ✅ (ephemeral) |  | 
| `Routes/Route` | ✅ | ✅ | 
| `Routing/Gateway` | ❌ | ❌ | 
| `Syslog/Settings` | ❌ | ❌ | 
| `Syslog/Settings/Destination` | ❌ | ❌ | 
| `Trafficshaper/Pipe` | ❌ | ❌ | 
| `Trafficshaper/Queue` | ❌ | ❌ | 
| `Trafficshaper/Rule` | ❌ | ❌ | 
| `Trust/Settings` | ❌ | ❌ | 
| `Trust/CA` | ❌ | ❌ | 
| `Trust/Cert` | ❌ | ❌ | 
| `Unbound/Settings` | ✅ | ✅ | 
| `Unbound/Settings/Domain Override` | ✅ | ✅ | 
| `Unbound/Settings/Forward` | ✅ | ✅ | 
| `Unbound/Settings/Host Alias` | ✅ | ✅ | 
| `Unbound/Settings/Host Override` | ✅ | ✅ | 
| `Unbound/Settings/ACL` | ❌ | ❌ | 
| `Wireguard/Settings` | ❌ | ❌ | 
| `Wireguard/Client` | 🚧 | 🚧 | 
| `Wireguard/Server` | 🚧 | 🚧 | 
| `Wireguard/Generate Key Pair` | ❌ | ❌ | 
| `Wireguard/Generate PSK` | ❌ | ❌ | 

The following is a non-exhaustive list of the plugin APIs OPNsense supports. The table shows those which are 'highest priority'. Please open a feature request to indicate interest for any plugin not listed here.

