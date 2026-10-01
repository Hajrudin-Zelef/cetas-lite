---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api-5
title: "Step 1: remove the old entry"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api.md
source_anchor: ""
source_lines: [845, 1085]
sha256: 31774af8e74cc2fc0d33c1a957fc83dcfd2437142dab03c4b3007dc1c306a487
---

# Step 1: remove the old entry

```
class { 'opn':
  kea_dhcpv6_peers => {
    'primary-node' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'role' => 'primary',
        'url'  => 'http://[fd00::1]:8000',
      },
    },
  },
}
```
Prometheus Node Exporter settings are managed as a singleton resource per device via the `node_exporters` parameter or directly via the `opn_node_exporter` type. The `node_exporters` hash is keyed by **device name**. After any change, Puppet calls `nodeexporter/service/reconfigure` once per device. The plugin `os-node_exporter` must be installed on the device.

```
class { 'opn':
  devices => { ... },
  plugins => {
    'os-node_exporter' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
    },
  },
  node_exporters => {
    'opnsense01.example.com' => {
      'ensure' => 'present',
      'config' => {
        'enabled'       => '1',
        'listenaddress' => '0.0.0.0',
        'listenport'    => '9100',
        'cpu'           => '1',
        'exec'          => '1',
        'filesystem'    => '1',
        'loadavg'       => '1',
        'meminfo'       => '1',
        'netdev'        => '1',
        'time'          => '1',
        'devstat'       => '1',
        'interrupts'    => '0',
        'ntp'           => '0',
        'zfs'           => '1',
      },
    },
  },
}
```
OpenVPN resources are managed via the `openvpn_*` parameters of the `opn` class or directly via the corresponding `opn_openvpn_*` types. After any OpenVPN change, Puppet calls `openvpn/service/reconfigure` at most once per device per run.

The module manages OpenVPN instances, static keys, and client-specific overrides (CSOs).

The `servers` field in `opn_openvpn_cso` and the `tls_key` field in `opn_openvpn_instance` accept instance/key descriptions which are automatically resolved to UUIDs.

The `password` field in `opn_openvpn_instance` is excluded from idempotency checks because it contains secret material.

```
class { 'opn':
  devices => { ... },
  openvpn_statickeys => {
    'my-tls-auth-key' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'key'  => '-----BEGIN OpenVPN Static key-----',
        'mode' => 'auth',
      },
    },
  },
  openvpn_instances => {
    'my-openvpn-server' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'role'               => 'server',
        'proto'              => 'udp',
        'port'               => '1194',
        'server'             => '10.8.0.0/24',
        'tls_key'            => 'my-tls-auth-key',
        'cert'               => 'my-openvpn-server-cert',
        'verify_client_cert' => 'required',
        'enabled'            => '1',
      },
    },
  },
  openvpn_csos => {
    'client1' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'servers'        => 'my-openvpn-server',
        'tunnel_network' => '10.8.1.0/24',
        'enabled'        => '1',
      },
    },
  },
}
```
Plugins can be managed via the `plugins` parameter of the `opn` class or directly using the `opn_plugin` type. The `devices` key controls which firewalls the plugin is deployed to. If `devices` is omitted, the plugin is applied to all devices defined in `$devices`.

```
class { 'opn':
  devices => { ... },
  plugins => {
    'os-haproxy' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
    },
    'os-acme-client' => {
      'devices' => ['opnsense01.example.com', 'opnsense02.example.com'],
      'ensure'  => 'present',
    },
  },
}
```
Puppet Agent settings are managed as a singleton resource per device via the `puppet_agents` parameter or directly via the `opn_puppet_agent` type. The `puppet_agents` hash is keyed by **device name**. After any change, Puppet calls `puppetagent/service/reconfigure` once per device. The plugin `os-puppet-agent` must be installed on the device.

```
class { 'opn':
  devices => { ... },
  plugins => {
    'os-puppet-agent' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
    },
  },
  puppet_agents => {
    'opnsense01.example.com' => {
      'ensure' => 'present',
      'config' => {
        'enabled'           => '1',
        'fqdn'              => 'puppet.example.com',
        'environment'       => 'production',
        'runinterval'       => '30m',
        'runtimeout'        => '1h',
        'usecacheonfailure' => '1',
      },
    },
  },
}
```
Static routes can be managed via the `routes` parameter or directly via the `opn_route` type. The route **description** (`descr`) is used as the resource identifier. After any change, Puppet calls `routes/routes/reconfigure` once per device.

```
class { 'opn':
  devices => { ... },
  routes => {
    'Server network' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'network'  => '10.0.0.0/24',
        'gateway'  => 'LAN_GW',
        'disabled' => '0',
      },
    },
  },
}
```
ZFS snapshots can be managed via the `snapshots` parameter or directly via the `opn_snapshot` type. The snapshot **name** is used as the resource identifier. The `active` property controls whether a snapshot is the active boot target.

```
class { 'opn':
  devices => { ... },
  snapshots => {
    'pre-upgrade' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'active' => true,
        'note'   => 'Snapshot before upgrade',
      },
    },
  },
}
```
Syslog remote destinations can be managed via the `syslog_destinations` parameter or directly via the `opn_syslog` type. The destination **description** is used as the resource identifier. After any change, Puppet calls `syslog/service/reconfigure` once per device.

```
class { 'opn':
  devices => { ... },
  syslog_destinations => {
    'Central syslog' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'transport' => 'udp4',
        'hostname'  => 'syslog.example.com',
        'port'      => '514',
        'enabled'   => '1',
      },
    },
  },
}
```
Trust Certificate Authorities can be managed via the `trust_cas` parameter or directly via the `opn_trust_ca` type. The CA **description** (`descr`) is used as the resource identifier.

Many fields (action, key_type, digest, etc.) are only relevant during initial creation and are ignored during idempotency checks.

```
class { 'opn':
  devices => { ... },
  trust_cas => {
    'Internal CA' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'action'     => 'internal',
        'key_type'   => 'RSA',
        'digest'     => 'SHA256',
        'lifetime'   => '3650',
        'country'    => 'DE',
        'commonname' => 'Internal CA',
      },
    },
  },
}
```
Trust certificates can be managed via the `trust_certs` parameter or directly via the `opn_trust_cert` type. The certificate **description** (`descr`) is used as the resource identifier. Like CAs, volatile fields are only used during creation.

```
class { 'opn':
  devices => { ... },
  trust_certs => {
    'web.example.com' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'action'     => 'internal',
        'caref'      => '<ca-uuid>',
        'key_type'   => 'RSA',
        'digest'     => 'SHA256',
        'lifetime'   => '365',
        'commonname' => 'web.example.com',
      },
    },
  },
}
```
Certificate Revocation Lists can be managed via the `trust_crls` parameter or directly via the `opn_trust_crl` type. The hash key is the **CA description** that the CRL belongs to. The provider resolves the CA description to the internal `caref` identifier automatically.

