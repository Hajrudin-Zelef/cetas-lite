---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api-1
title: "Step 1: remove the old entry"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api.md
source_anchor: ""
source_lines: [1, 156]
sha256: c8d6efe51a34dc465d8505a0099fc211cf54e9894cd955e466fa084c29796991
---

# Step 1: remove the old entry

A Puppet module to manage OPNsense firewalls via the OPNsense REST API. It is meant as a replacement for puppet-opnsense.

This module provides the following resource types for one or more OPNsense devices:

| Type | Manages | 
|---|---|
| `opn_acmeclient_account` | ACME Client accounts | 
| `opn_acmeclient_action` | ACME Client automation actions | 
| `opn_acmeclient_certificate` | ACME Client certificates | 
| `opn_acmeclient_settings` | ACME Client global settings (singleton per device) | 
| `opn_acmeclient_validation` | ACME Client validation methods | 
| `opn_cron` | Cron jobs | 
| `opn_dhcrelay_destination` | DHCP Relay destinations | 
| `opn_dhcrelay` | DHCP Relay instances | 
| `opn_firewall_alias` | Firewall aliases | 
| `opn_firewall_category` | Firewall categories | 
| `opn_firewall_group` | Firewall interface groups | 
| `opn_firewall_rule` | Firewall filter rules (new GUI) | 
| `opn_group` | Local groups | 
| `opn_haproxy_acl` | HAProxy ACLs (conditions) | 
| `opn_haproxy_action` | HAProxy actions (rules) | 
| `opn_haproxy_backend` | HAProxy backend pools | 
| `opn_haproxy_cpu` | HAProxy CPU affinity / thread binding | 
| `opn_haproxy_errorfile` | HAProxy error files | 
| `opn_haproxy_fcgi` | HAProxy FastCGI applications | 
| `opn_haproxy_frontend` | HAProxy frontend listeners | 
| `opn_haproxy_group` | HAProxy user-list groups | 
| `opn_haproxy_healthcheck` | HAProxy health checks | 
| `opn_haproxy_lua` | HAProxy Lua scripts | 
| `opn_haproxy_mailer` | HAProxy mailers | 
| `opn_haproxy_mapfile` | HAProxy map files | 
| `opn_haproxy_resolver` | HAProxy DNS resolvers | 
| `opn_haproxy_server` | HAProxy backend servers | 
| `opn_haproxy_settings` | HAProxy global settings (singleton per device) | 
| `opn_haproxy_user` | HAProxy user-list users | 
| `opn_hasync` | HA sync / CARP settings (singleton per device) | 
| `opn_ipsec_child` | IPsec child SAs (Swanctl) | 
| `opn_ipsec_connection` | IPsec connections (Swanctl) | 
| `opn_ipsec_keypair` | IPsec key pairs (Swanctl) | 
| `opn_ipsec_local` | IPsec local authentication (Swanctl) | 
| `opn_ipsec_pool` | IPsec address pools (Swanctl) | 
| `opn_ipsec_presharedkey` | IPsec pre-shared keys (Swanctl) | 
| `opn_ipsec_remote` | IPsec remote authentication (Swanctl) | 
| `opn_ipsec_settings` | IPsec global settings (singleton per device) | 
| `opn_ipsec_vti` | IPsec VTI entries (Swanctl) | 
| `opn_kea_ctrl_agent` | KEA Control Agent settings (singleton per device) | 
| `opn_kea_dhcpv4` | KEA DHCPv4 global settings (singleton per device) | 
| `opn_kea_dhcpv4_peer` | KEA DHCPv4 HA peers | 
| `opn_kea_dhcpv4_reservation` | KEA DHCPv4 reservations | 
| `opn_kea_dhcpv4_subnet` | KEA DHCPv4 subnets | 
| `opn_kea_dhcpv6` | KEA DHCPv6 global settings (singleton per device) | 
| `opn_kea_dhcpv6_pd_pool` | KEA DHCPv6 prefix delegation pools | 
| `opn_kea_dhcpv6_peer` | KEA DHCPv6 HA peers | 
| `opn_kea_dhcpv6_reservation` | KEA DHCPv6 reservations | 
| `opn_kea_dhcpv6_subnet` | KEA DHCPv6 subnets | 
| `opn_node_exporter` | Prometheus Node Exporter settings (singleton per device) | 
| `opn_openvpn_cso` | OpenVPN client-specific overrides | 
| `opn_openvpn_instance` | OpenVPN instances | 
| `opn_openvpn_statickey` | OpenVPN static keys | 
| `opn_gateway` | Routing gateways | 
| `opn_plugin` | Firmware plugins / packages | 
| `opn_puppet_agent` | Puppet Agent settings (singleton per device) | 
| `opn_route` | Static routes | 
| `opn_snapshot` | ZFS snapshots | 
| `opn_syslog` | Syslog remote destinations | 
| `opn_trust_ca` | Trust Certificate Authorities | 
| `opn_trust_cert` | Trust certificates | 
| `opn_trust_crl` | Trust Certificate Revocation Lists | 
| `opn_tunable` | System tunables (sysctl) | 
| `opn_user` | Local users | 
| `opn_zabbix_agent` | Zabbix Agent settings (singleton per device) | 
| `opn_zabbix_agent_alias` | Zabbix Agent Alias entries | 
| `opn_zabbix_agent_userparameter` | Zabbix Agent UserParameter entries | 
| `opn_zabbix_proxy` | Zabbix Proxy settings (singleton per device) | 

- No external tools required — only Ruby's built-in HTTP library
- Simple, uniform provider implementation — low maintenance overhead
- Validation delegated to the OPNsense API — no duplication of API logic in providers
- Integrated UUID resolver for ModelRelationField and CertificateField references
- Automatic reload/reconfigure after configuration changes (once per device per run)
- Config passthrough — the `config` hash is sent to the API as-is, new API fields work without code changes
- Custom fact (`opnsense` ) exposes version and installed plugins on OPNsense hosts

One or more OPNsense firewalls with API access enabled. API credentials can be created in OPNsense under **System → Access → Users → (User) → API keys**.

The `opn` class manages API credential files for one or more OPNsense devices. These files are written to `$config_dir` and are read by the `opn_*` providers at catalog application time.

```
class { 'opn':
  devices => {
    'localhost' => {
      'url'        => 'https://localhost/api',
      'api_key'    => 'OPNSENSE_API_KEY',
      'api_secret' => 'OPNSENSE_API_SECRET',
      'ssl_verify' => false,
    },
  },
  plugins => {
    'os-helloworld' => {
      'devices' => ['localhost'],
      'ensure'  => 'present',
    },
  },
  firewall_aliases => {
    'alias_test001' => {
      'devices' => ['localhost'],
      'ensure'  => 'present',
      'config'  => {
        'type'        => 'host',
        'content'     => '192.168.1.1',
        'description' => 'Test alias',
        'enabled'     => '1',
      },
    },
  },
}
```
Note that for some parameters, OPNsense expects a newline as separator. In these cases the value must be provided as `"value1\nvalue2"`, as demonstrated in some examples below. One of the main goals of this module is code simplification, so this is not done by the provider. However, when using Hiera, there's a better solution available than using `\n`:

```
opn::firewall_aliases:
  geoip_example:
    ensure: 'present'
    config:
      enabled: '1'
      type: 'geoip'
      description: 'List of countries'
      content: |-
        AO
        BF
        BI
        BJ
        BW
```
When managing more than one OPNsense firewall, add each device to the `devices` hash. Resource titles use the `resource_name@device_name` format to identify which device a resource belongs to. All resources that do not specify an explicit `devices` array are applied to **every** device in the `devices` hash.

```
class { 'opn':
  devices => {
    'opnsense01.example.com' => {
      'url'        => 'https://opnsense01.example.com/api',
      'api_key'    => 'OPNSENSE_API_KEY',
      'api_secret' => 'OPNSENSE_API_SECRET',
      'ssl_verify' => true,
    },
    'opnsense02.example.com' => {
      'url'        => 'https://opnsense02.example.com/api',
      'api_key'    => 'OPNSENSE_API_KEY',
      'api_secret' => 'OPNSENSE_API_SECRET',
      'ssl_verify' => true,
    },
  },
}
```
If a remote device only needs API credentials (e.g. for targeted resources like Zabbix Proxy) but should **not** receive all default resources, use `additional_devices` instead. Devices in `additional_devices` get a credential file created, but are **not** included in the default device list for resource iteration. You can still target them explicitly via the `devices` array on individual resources.

