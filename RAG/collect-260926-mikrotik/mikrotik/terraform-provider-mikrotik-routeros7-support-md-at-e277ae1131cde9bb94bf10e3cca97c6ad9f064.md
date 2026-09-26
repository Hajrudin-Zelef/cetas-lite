---
id: collect-260926-mikrotik/mikrotik/terraform-provider-mikrotik-routeros7-support-md-at-e277ae1131cde9bb94bf10e3cca97c6ad9f064
title: "Start RouterOS 7 container"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/terraform-provider-mikrotik-routeros7-support-md-at-e277ae1131cde9bb94bf10e3cca97c6ad9f06449-lkolo-p.md
source_anchor: ""
source_lines: [1, 196]
sha256: 27994b068fbfee071b5fb6fbf500a2a6b63ce52307b78c1642b34c3dfe8cf69b
---

# Start RouterOS 7 container

This document describes the updates made to support MikroTik RouterOS 7.x.

This provider has been updated to fully support RouterOS 7.x with new resources and improved compatibility.

- **go-routeros** : Updated from v0.0.0-20210123142807 to v3.0.1 (latest)
  - Better RouterOS 7 API compatibility
  - Improved error handling
  - Performance improvements
  - **Requires Go 1.21+** (for`log/slog` support)

- **Go version** : Updated from 1.18 to 1.21-1.23
- Test matrix now includes RouterOS 7.14.3, 7.16.2, 7.17
- Removed RouterOS 6.x testing (legacy support)
- All tests run against RouterOS 7.x by default

Replaces the deprecated `mikrotik_bgp_instance` and `mikrotik_bgp_peer` resources.

**Features:**

- Unified BGP configuration
- Template support for configuration reuse
- Enhanced filtering with `input.filter` and`output.filter`
- Built-in BFD support
- MPLS and VPN (VPNv4/VPNv6) support

**Example:**

```
resource "mikrotik_bgp_connection" "isp" {
  name           = "isp-connection"
  as             = 65001
  remote_address = "10.0.0.1"
  remote_as      = 65000
  router_id      = "192.168.1.1"
  
  address_family = "ip,ipv6"
  use_bfd        = true
  multihop       = false
  
  input_filter   = "bgp-in"
  output_filter  = "bgp-out"
}
```
Allows configuration reuse across multiple BGP connections.

**Features:**

- Reusable configuration blocks
- Default values for BGP parameters
- Filtering templates
- Route reflection configuration

**Example:**

```
resource "mikrotik_bgp_template" "default" {
  name           = "default"
  as             = 65001
  hold_time      = "3m"
  keepalive_time = "1m"
  
  output_default_originate = "if-installed"
  use_bfd                  = true
}
```
Pre-connection-tracking firewall rules for performance optimization.

**Features:**

- Process packets before connection tracking
- Improved performance for high-throughput scenarios
- DDoS mitigation capabilities
- Connection state bypass

**Example:**

```
resource "mikrotik_firewall_raw" "fastpath" {
  chain       = "prerouting"
  action      = "accept"
  src_address = "192.168.0.0/16"
  dst_address = "10.0.0.0/8"
  comment     = "Internal traffic fast path"
}
```
Improved VLAN interface with RouterOS 7 features.

**Features:**

- Hardware acceleration support
- Service tag support (Q-in-Q)
- Improved MTU handling
- Better integration with bridge VLAN filtering

**Example:**

```
resource "mikrotik_interface_vlan7" "management" {
  name            = "vlan-mgmt"
  vlan_id         = 10
  interface       = "ether1"
  use_service_tag = false
  mtu             = 1500
}
```
Hardware-accelerated VLAN filtering on bridges.

**Features:**

- Hardware offload on supported devices
- Ingress filtering
- Frame type control
- PVID mode configuration

**Example:**

```
resource "mikrotik_bridge_vlan_filtering" "main" {
  bridge            = "bridge1"
  vlan_filtering    = true
  ingress_filtering = true
  frame_types       = "admit-only-vlan-tagged"
  pvid_mode         = "secure"
}
```
All existing resources continue to work with RouterOS 7:

- ✅ `mikrotik_interface_wireguard` - Full RouterOS 7 support
- ✅ `mikrotik_interface_wireguard_peer` - Full RouterOS 7 support
- ✅ `mikrotik_dhcp_server` - Full RouterOS 7 support
- ✅ `mikrotik_dhcp_server_network` - Full RouterOS 7 support
- ✅ `mikrotik_firewall_filter` - Full RouterOS 7 support
- ✅ `mikrotik_ip_address` - Full RouterOS 7 support
- ✅ `mikrotik_interface_list` - Full RouterOS 7 support
- ✅ `mikrotik_bridge` - Full RouterOS 7 support
- ✅ `mikrotik_bridge_port` - Full RouterOS 7 support

The following resources are deprecated in RouterOS 7 but maintained for backward compatibility:

- ⚠️ `mikrotik_bgp_instance` - Use`mikrotik_bgp_connection` instead
- ⚠️ `mikrotik_bgp_peer` - Use`mikrotik_bgp_connection` instead
- ⚠️ `mikrotik_wireless_interface` - CAPsMAN v2 recommended for RouterOS 7

See MIGRATION_ROUTEROS7.md for detailed migration instructions.

RouterOS 7 provides several performance enhancements:

1. **Fast Track** : Improved fast-path processing
2. **Hardware Offloading** : Better use of switch chip capabilities
3. **Connection Tracking** : Optimized for higher throughput
4. **Bridge VLAN Filtering** : Hardware-accelerated on supported devices
5. **WireGuard** : Native kernel implementation

All resources are tested against:

- RouterOS 7.14.3 (stable)
- RouterOS 7.16.2 (stable)
- RouterOS latest (experimental)

Run tests locally:

```
# Start RouterOS 7 container
make routeros ROUTEROS_VERSION=7.16.2
# Run tests
export MIKROTIK_HOST=127.0.0.1:8728
export MIKROTIK_USER=admin
export MIKROTIK_PASSWORD=""
make testacc
```
1. **BGP Template Inheritance** : Complex template inheritance may require explicit configuration
2. **VLAN Filtering** : Some older hardware may not support all filtering modes
3. **Container Resources** : Not yet implemented (planned for future release)

Future enhancements planned:

-  `/container` resource support
-  `/routing/filter` resource (new routing filters in v7)
-  `/routing/ospf` v3 resources (new OSPF implementation)
-  `/routing/rip` v2 resources (new RIP implementation)
- CAPsMAN v2 resources (new wireless management)
- ZeroTier integration resources

Contributions for RouterOS 7 features are welcome! Please:

1. Test against RouterOS 7.x
2. Include unit tests
3. Update documentation
4. Follow existing code patterns

See CONTRIBUTING.md for details.

| Provider Version | RouterOS 6.x | RouterOS 7.x | 
|---|---|---|
| < 1.0.0 | ✅ Full |  | 
| >= 1.0.0 |  | ✅ Full |
