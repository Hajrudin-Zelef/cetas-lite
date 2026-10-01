---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api-6
title: "Step 1: remove the old entry"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api.md
source_anchor: ""
source_lines: [1086, 1246]
sha256: d7f578013244f2b2a67a8f9860a4ce74382eea0a056bb1858667a9cf1b3230c7
---

# Step 1: remove the old entry

Each CA can have at most one CRL. The `set` endpoint creates or updates the CRL.

```
class { 'opn':
  devices => { ... },
  trust_crls => {
    'Internal CA' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'descr'     => 'CRL for Internal CA',
        'lifetime'  => '9999',
        'crlmethod' => 'internal',
      },
    },
  },
}
```
System tunables (sysctl variables) can be managed via the `tunables` parameter or directly via the `opn_tunable` type. The **tunable name** (e.g. `kern.maxproc`) is used as the resource identifier. After any change, Puppet calls `core/tunables/reconfigure` once per device.

```
class { 'opn':
  devices => { ... },
  tunables => {
    'kern.maxproc' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'value'       => '4096',
        'description' => 'Maximum number of processes',
      },
    },
  },
}
```
Local OPNsense users can be managed via the `users` parameter or directly via the `opn_user` type.

```
class { 'opn':
  devices => { ... },
  users => {
    'jdoe' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'password' => 'plaintextpassword',
        'descr'    => 'John Doe',
        'email'    => 'jdoe@example.com',
        # The 'uid' attribute is optional. Be aware that OPNsense will
        # change the uid if it is already taken by another user.
        #'uid'     => '2001',
      },
    },
  },
}
```
The Zabbix Agent configuration is also a singleton resource per device. It requires the `os-zabbix-agent` plugin.

The `zabbix_agents` hash is keyed by **device name**. The `config` hash is passed to the `opn_zabbix_agent` resource type. Its structure mirrors the nested OPNsense ZabbixAgent model (`settings.main`, `settings.tuning`, `settings.features`, `local`).

UserParameter and Alias entries are managed separately via `zabbix_agent_userparameters` and `zabbix_agent_aliases`. All changes to agent resources trigger a single `zabbixagent/service/reconfigure` call per device per Puppet run.

The hash key in `zabbix_agent_userparameters` and `zabbix_agent_aliases` is the **Zabbix key** (the identifier sent to OPNsense). It must not be repeated inside the config — any `key` value there is ignored. See Resource identifiers for the general explanation and rename pattern.

```
class { 'opn':
  devices => {
    'opnsense01.example.com' => {
      'url'        => 'https://opnsense01.example.com/api',
      'api_key'    => 'OPNSENSE_API_KEY',
      'api_secret' => 'OPNSENSE_API_SECRET',
    },
  },
  plugins => {
    'os-zabbix-agent' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
    },
  },
  zabbix_agents => {
    'opnsense01.example.com' => {
      'ensure' => 'present',
      'config' => {
        'local' => {
          'hostname' => 'opnsense01.example.com',
        },
        'settings' => {
          'main' => {
            'enabled'    => '1',
            'serverList' => 'zabbix.example.com',
            'listenPort' => '10050',
          },
          'features' => {
            'enableActiveChecks'   => '1',
            'activeCheckServers'   => 'zabbix.example.com',
            'enableRemoteCommands' => '0',
          },
        },
      },
    },
  },
  zabbix_agent_userparameters => {
    'custom.uptime' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'command'      => '/usr/bin/uptime',
        'enabled'      => '1',
        'acceptParams' => '0',
      },
    },
  },
  zabbix_agent_aliases => {
    'ping' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'sourceKey'    => 'icmpping',
        'enabled'      => '1',
        'acceptParams' => '0',
      },
    },
  },
}
```
The Zabbix Proxy configuration is a singleton resource — one per OPNsense device. It requires the `os-zabbix-proxy` plugin to be installed. Use `opn_plugin` to install it before applying the settings.

The `zabbix_proxies` hash is keyed by **device name** (not by a `name@device` title), since only one Zabbix Proxy configuration exists per device. The `config` hash is passed to the `opn_zabbix_proxy` resource type.

After any change, Puppet calls `zabbixproxy/service/reconfigure` once to apply the new configuration.

```
class { 'opn':
  devices => {
    'opnsense01.example.com' => {
      'url'        => 'https://opnsense01.example.com/api',
      'api_key'    => 'OPNSENSE_API_KEY',
      'api_secret' => 'OPNSENSE_API_SECRET',
    },
  },
  plugins => {
    'os-zabbix-proxy' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
    },
  },
  zabbix_proxies => {
    'opnsense01.example.com' => {
      'ensure' => 'present',
      'config' => {
        'enabled'    => '1',
        'server'     => 'zabbix.example.com',
        'serverport' => '10051',
        'hostname'   => 'opnsense01-proxy',
      },
    },
  },
}
```
All types can also be used directly without the `opn` wrapper class.

