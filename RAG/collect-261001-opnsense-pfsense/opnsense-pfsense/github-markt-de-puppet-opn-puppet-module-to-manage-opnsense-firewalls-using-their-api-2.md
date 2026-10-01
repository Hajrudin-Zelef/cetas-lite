---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api-2
title: "Step 1: remove the old entry"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api.md
source_anchor: ""
source_lines: [157, 339]
sha256: 46df2c52140374b4be5f89adfe63b09aec271fd92c035cb00e849c9350cad5ca
---

# Step 1: remove the old entry

```
class { 'opn':
  devices => {
    'opnsense01.example.com' => {
      'url'        => 'https://opnsense01.example.com/api',
      'api_key'    => 'OPNSENSE_API_KEY',
      'api_secret' => 'OPNSENSE_API_SECRET',
    },
  },
  additional_devices => {
    'opnsense-remote01.example.com' => {
      'url'        => 'https://opnsense-remote01.example.com/api',
      'api_key'    => 'REMOTE_API_KEY',
      'api_secret' => 'REMOTE_API_SECRET',
    },
  },
  # This Zabbix proxy is only applied to the remote device
  zabbix_proxies => {
    'opnsense-remote01.example.com' => {
      'devices' => ['opnsense-remote01.example.com'],
      'config'  => {
        'enabled' => '1',
      },
    },
  },
  # This firewall alias is only applied to opnsense01 (the default device),
  # NOT to opnsense-remote01
  firewall_aliases => {
    'webservers' => {
      'config' => {
        'type'    => 'host',
        'content' => '10.0.0.1,10.0.0.2',
        'enabled' => '1',
      },
    },
  },
}
```
A device name must not appear in both `devices` and `additional_devices` — the module raises an error if names overlap.

Most `opn_*` resource types manage **lists of entries** (firewall aliases, HAProxy servers, Zabbix UserParameters, etc.). Their resource title uses the format `identifier@device_name`, where the part before `@` is the **unique identifier** that OPNsense uses to track the entry — for example the alias name, the server description, or the Zabbix UserParameter key.

This identifier is set from the resource title, not from the `config` hash. Specifying the identifier field inside `config` has no effect — the title value always takes precedence and is what gets written to OPNsense.

As a consequence, **renaming an identifier requires two steps**: declare the old resource with `ensure => absent` and add a new resource with the new identifier in the title.

```
# Step 1: remove the old entry
opn_haproxy_server { 'web01@opnsense01.example.com':
  ensure => absent,
}
# Step 2: add the new entry with the renamed identifier
opn_haproxy_server { 'new-web01@opnsense01.example.com':
  ensure => present,
  config => {
    'address' => '10.0.0.1',
    'port'    => '80',
    'enabled' => '1',
  },
}
```
The same pattern applies to all list-based types: `opn_acmeclient_account`, `opn_acmeclient_action`, `opn_acmeclient_certificate`, `opn_acmeclient_validation`, `opn_cron`, `opn_dhcrelay_destination`, `opn_dhcrelay`, `opn_firewall_alias`, `opn_firewall_category`, `opn_firewall_group`, `opn_firewall_rule`, `opn_user`, `opn_group`, all `opn_haproxy_*` types, all `opn_ipsec_*` list types, all `opn_openvpn_*` types, `opn_snapshot`, `opn_syslog`, `opn_trust_ca`, `opn_trust_cert`, `opn_trust_crl`, `opn_tunable`, `opn_zabbix_agent_userparameter`, and `opn_zabbix_agent_alias`.

**Singleton resources** (`opn_acmeclient_settings`, `opn_haproxy_settings`, `opn_hasync`, `opn_ipsec_settings`, `opn_node_exporter`, `opn_puppet_agent`, `opn_zabbix_proxy`, `opn_zabbix_agent`) are different: their title is the device name itself, and the entire `config` hash is written to the OPNsense API on every change.

**Important:** Singleton resources always exist in the OPNsense API — the API always returns their current configuration, even when all values are at defaults. Because of this, `ensure => absent` will trigger a `destroy` action on **every** Puppet run (resetting the config and calling reconfigure each time). To disable a singleton service, use `ensure => present` with `'enabled' => '0'` instead. This is idempotent and only triggers a change when the current state differs from the desired state.

ACME Client resources (accounts, actions, certificates, validations, settings) are managed via the `acmeclient_*` parameters or directly via the corresponding `opn_acmeclient_*` types. The plugin `os-acme-client` must be installed on the device.

Certificate relation fields (`account`, `validationMethod`, `restartActions`) and validation relation fields (`http_haproxyFrontends`) accept names which are automatically resolved to UUIDs. Settings relation fields (`UpdateCron`, `haproxyAclRef`, `haproxyActionRef`, `haproxyServerRef`, `haproxyBackendRef`) are also resolved by name.

Only changes to `opn_acmeclient_settings` trigger `acmeclient/service/reconfigure`. The other four types do not trigger a reconfigure.

Note that some items in Acme Client must have `enabled=1` set, otherwise they cannot be used/referenced by other items.

```
class { 'opn':
  devices => { ... },
  plugins => {
    'os-acme-client' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
    },
  },
  acmeclient_accounts => {
    'le-account' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'ca'      => 'letsencrypt',
        'email'   => 'admin@example.com',
        'enabled' => '1',
      },
    },
  },
  acmeclient_actions => {
    'restart_haproxy' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'type'                    => 'configd_generic',
        'configd_generic_command' => 'haproxy restart',
        'enabled'                 => '1',
      },
    },
  },
  acmeclient_validations => {
    'http-01' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'method'       => 'http01',
        'http_service' => 'haproxy',
        'enabled'      => '1',
      },
    },
  },
  acmeclient_certificates => {
    'web.example.com' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'altNames'         => 'www.example.com',
        'account'          => 'le-account',
        'validationMethod' => 'http-01',
        'restartActions'   => 'restart_haproxy',
        'enabled'          => '1',
      },
    },
  },
  acmeclient_settings => {
    'opnsense01.example.com' => {
      'ensure' => 'present',
      'config' => {
        'environment' => 'stg',
        'autoRenewal' => '1',
      },
    },
  },
}
```
Cron jobs can be managed via the `cron_jobs` parameter or directly via the `opn_cron` type. The job **description** is used as the resource identifier and must be unique per device. After any change, Puppet calls `cron/service/reconfigure` once per device.

```
class { 'opn':
  devices => { ... },
  cron_jobs => {
    'Reload HAProxy' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'command'  => 'haproxy reload',
        'minutes'  => '0',
        'hours'    => '3',
        'days'     => '*',
        'months'   => '*',
        'weekdays' => '*',
        'enabled'  => '1',
      },
    },
    'HAProxy: sync certificates' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'command'  => 'haproxy cert_sync_bulk',
        'minutes'  => '0',
        'hours'    => '1',
        'days'     => '*',
        'months'   => '*',
        'weekdays' => '*',
        'enabled'  => '1',
        'origin'   => 'HAProxy',
      },
    },
  },
}
```
DHCP Relay resources (destinations and relays) are managed via the `dhcrelay_destinations` and `dhcrelays` parameters or directly via the corresponding `opn_dhcrelay_destination` and `opn_dhcrelay` types.

Destinations are named groups of DHCP server IPs. Relays are per-interface relay instances referencing a destination. The `destination` config key on a relay accepts the destination name, which is automatically resolved to the corresponding UUID by the provider.

Important: Relay instances have no name/description field in the OPNsense API. The resource title is a freeform label (e.g. `"LAN IPv4 Relay@opnsense01"`), and the provider matches existing API resources by the `interface` value from the config hash. Each interface can only have one relay per device.

