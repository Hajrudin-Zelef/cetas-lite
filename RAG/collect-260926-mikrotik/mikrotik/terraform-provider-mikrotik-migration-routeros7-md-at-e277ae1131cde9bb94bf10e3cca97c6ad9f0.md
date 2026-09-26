---
id: collect-260926-mikrotik/mikrotik/terraform-provider-mikrotik-migration-routeros7-md-at-e277ae1131cde9bb94bf10e3cca97c6ad9f0
title: "Optional: Create a template for reusable configuration"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/terraform-provider-mikrotik-migration-routeros7-md-at-e277ae1131cde9bb94bf10e3cca97c6ad9f06449-lkolo.md
source_anchor: ""
source_lines: [1, 220]
sha256: fd86b2795cfaefae6b93facbed1b5cc02f6440eb5e87bfd53b0aa95b1159ad65
---

# Optional: Create a template for reusable configuration

This guide helps you migrate your Terraform MikroTik provider configuration from RouterOS 6 to RouterOS 7.

RouterOS 7 introduced significant architectural changes, particularly in the BGP implementation. The legacy BGP commands (`/routing/bgp/instance` and `/routing/bgp/peer`) have been replaced with a new, more flexible system.

```
resource "mikrotik_bgp_instance" "main" {
  name      = "main"
  as        = 65530
  router_id = "172.16.0.1"
  
  redistribute_connected = true
  redistribute_static    = true
}
resource "mikrotik_bgp_peer" "upstream" {
  name           = "upstream-peer"
  instance       = mikrotik_bgp_instance.main.name
  remote_address = "10.0.0.1"
  remote_as      = 65531
  ttl            = "255"
}
```
```
# Optional: Create a template for reusable configuration
resource "mikrotik_bgp_template" "default" {
  name              = "default-template"
  as                = 65530
  router_id         = "172.16.0.1"
  hold_time         = "3m"
  keepalive_time    = "1m"
  address_families  = "ip,ipv6"
  
  output_default_originate = "always"
}
# Create BGP connection
resource "mikrotik_bgp_connection" "upstream" {
  name           = "upstream-connection"
  as             = 65530
  remote_address = "10.0.0.1"
  remote_as      = 65531
  router_id      = "172.16.0.1"
  
  templates      = mikrotik_bgp_template.default.name
  
  ttl            = "255"
  address_family = "ip,ipv6"
  
  output_default_originate = "always"
  output_network           = "connected,static"
  
  use_bfd        = true
  multihop       = true
}
```
**Key Differences:**

- **Unified Resource** : BGP Instance and Peer are merged into`bgp_connection`
- **Templates** : New template system for configuration reuse
- **Input/Output Filters** : More granular control with`input.filter` and`output.filter`
- **Network Advertisement** : Use`output.network` instead of`redistribute_*` options
- **Enhanced Features** : Built-in BFD support, MPLS, VPNv4/v6

RouterOS 7 introduces the RAW firewall table for pre-connection-tracking processing:

```
resource "mikrotik_firewall_raw" "prerouting_accept" {
  chain            = "prerouting"
  action           = "accept"
  src_address      = "192.168.1.0/24"
  dst_address      = "10.0.0.0/8"
  protocol         = "tcp"
  dst_port         = "80,443"
  comment          = "Fast path for trusted internal traffic"
}
```
**Use Cases:**

- Bypass connection tracking for performance
- Early packet filtering before connection state
- DDoS mitigation
- Fast-path optimization

RouterOS 7 improves VLAN handling with hardware acceleration:

```
# Modern VLAN interface
resource "mikrotik_interface_vlan7" "vlan100" {
  name            = "vlan100"
  vlan_id         = 100
  interface       = "ether1"
  use_service_tag = false
  mtu             = 1500
  comment         = "Management VLAN"
}
# Bridge VLAN filtering (hardware accelerated)
resource "mikrotik_bridge_vlan_filtering" "main_bridge" {
  bridge             = "bridge1"
  vlan_filtering     = true
  ingress_filtering  = true
  frame_types        = "admit-only-vlan-tagged"
  pvid_mode          = "secure"
}
```
1. Backup your current configuration
2. Upgrade to RouterOS 7.x (latest stable recommended)
3. Test all functionality after upgrade

```
# Update your terraform configuration
terraform init -upgrade
```
Update your `terraform` block:

```
terraform {
  required_providers {
    mikrotik = {
      source  = "ddelnano/mikrotik"
      version = "~> 1.0"  # Use latest version supporting RouterOS 7
    }
  }
}
```
1. **Export existing BGP configuration** from RouterOS 6
2. **Create migration script** to convert to new format
3. **Test in staging** before production

Example migration script:

```
#!/bin/bash
# Convert old BGP instance to new connection format
OLD_INSTANCE="main"
OLD_PEER="upstream-peer"
NEW_CONNECTION="upstream-connection"
# Export old config
/routing bgp instance export file=old-bgp-instance
/routing bgp peer export file=old-bgp-peer
# After manual conversion, import new config
/routing bgp connection import file=new-bgp-connection
```
Take advantage of RouterOS 7 improvements:

1. **Fast Track** : Enable for improved throughput
2. **BFD** : Use for faster BGP convergence
3. **Hardware VLAN Filtering** : Enable on supported hardware
4. **IPv6** : Enhanced IPv6 support throughout

| Feature | RouterOS 6 | RouterOS 7 | Notes | 
|---|---|---|---|
| Legacy BGP (instance/peer) | ✅ Supported |  | Will be removed in future | 
| New BGP (connection/template) | ❌ Not Available | ✅ Supported | Recommended | 
| Firewall RAW Table | ❌ Not Available | ✅ Supported | Performance feature | 
| Hardware VLAN Filtering |  | ✅ Full Support | Better performance | 
| WireGuard | ❌ Not Available | ✅ Supported | Native VPN | 
| Container | ❌ Not Available | ✅ Supported | Run containers on router | 

After migration, verify your Terraform state:

```
terraform plan
# Should show no changes if migration is correct
```
Test critical paths:

```
# Test BGP connectivity
/routing bgp connection print
/routing bgp session print
# Test firewall rules
/ip firewall raw print
/ip firewall filter print
# Test VLAN functionality
/interface vlan print
/interface bridge vlan print
```
Always have a rollback plan:

1. Keep RouterOS 6 configuration backup
2. Document all changes
3. Test rollback procedure in staging

**Symptoms**: BGP connections show "idle" or "active" state

**Solutions**:

1. Check `router-id` is set correctly
2. Verify `remote.address` and`local.address`
3. Enable `multihop` if needed
4. Check firewall rules (port 179)

**Symptoms**: No traffic passing through VLAN interfaces

**Solutions**:

1. Ensure `vlan-filtering=yes` on bridge
2. Check `bridge vlan` table is configured
3. Verify `pvid` settings on bridge ports
4. Check `frame-types` settings

**Symptoms**: Throughput lower after migration

**Solutions**:

1. Enable FastTrack in `/ip firewall filter`
2. Use RAW table for early packet filtering
3. Review connection tracking settings
4. Check hardware offloading status

If you encounter issues:

1. Check the GitHub Issues
2. Join the Discord Community
3. Review MikroTik's official documentation
4. Post detailed logs and configuration when asking for help

- **RouterOS 6.x** : Legacy support (maintenance mode)
- **RouterOS 7.x** : Full support (actively developed)
- **Migration Period** : Both versions supported during transition

Plan your migration within 6-12 months for best experience and support.
