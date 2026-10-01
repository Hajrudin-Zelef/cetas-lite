---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api-4
title: "Step 1: remove the old entry"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api.md
source_anchor: ""
source_lines: [571, 844]
sha256: 49a961d081bca469f5471e344d6549053144b86225459762a852d807a18cca1d
---

# Step 1: remove the old entry

```
class { 'opn':
  devices => { ... },
  haproxy_settings => {
    'opnsense01.example.com' => {
      'ensure' => 'present',
      'config' => {
        'general' => {
          'enabled' => '1',
          'stats'   => {
            'enabled' => '0',
          },
        },
        'maintenance' => {
          'cronjobs' => {
            # Add a reference to an existing cron job.
            'syncCertsCron' => 'HAProxy: sync certificates',
          },
        },
      },
    },
  },
}
```
IPsec resources are managed via the `ipsec_*` parameters of the `opn` class or directly via the corresponding `opn_ipsec_*` types. After any IPsec change, Puppet calls `ipsec/service/reconfigure` at most once per device per run.

The module manages the full IPsec connection hierarchy: connections, local/remote authentication, child SAs, pools, pre-shared keys, key pairs, VTI entries, and global settings.

Relation fields in `opn_ipsec_child`, `opn_ipsec_local`, and `opn_ipsec_remote` accept connection descriptions and key pair names which are automatically resolved to UUIDs.

The `privateKey` field in `opn_ipsec_keypair` and the `Key` field in `opn_ipsec_presharedkey` are excluded from idempotency checks because they contain secret material.

```
class { 'opn':
  devices => { ... },
  ipsec_connections => {
    'site-to-site' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'version'      => '2',
        'proposals'    => 'aes256-sha256-modp2048',
        'local_addrs'  => '0.0.0.0/0',
        'remote_addrs' => '198.51.100.1',
        'enabled'      => '1',
      },
    },
  },
  ipsec_children => {
    'child-lan' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'connection' => 'site-to-site',
        'mode'       => 'tunnel',
        'local_ts'   => '10.0.0.0/24',
        'remote_ts'  => '10.0.1.0/24',
        'enabled'    => '1',
      },
    },
  },
  ipsec_settings => {
    'opnsense01.example.com' => {
      'ensure' => 'present',
      'config' => {
        'general' => {
          'enabled' => '1',
        },
      },
    },
  },
}
```
KEA DHCP is the modern DHCP service in OPNsense, supporting both DHCPv4 and DHCPv6
with high availability, prefix delegation, and per-subnet option data. After any
change to KEA types, Puppet calls `kea/service/reconfigure` once per device.

The KEA Control Agent is managed as a singleton resource per device via the
`kea_ctrl_agents` parameter or directly via the `opn_kea_ctrl_agent` type.

```
class { 'opn':
  kea_ctrl_agents => {
    'opnsense01.example.com' => {
      'ensure' => 'present',
      'config' => {
        'general' => {
          'enabled'   => '1',
          'http_host' => '127.0.0.1',
          'http_port' => '8000',
        },
      },
    },
  },
}
```
DHCPv4 global settings are managed as a singleton resource per device via the
`kea_dhcpv4s` parameter or directly via the `opn_kea_dhcpv4` type. The `general`,
`lexpire` and `ha` sections are supported.

```
class { 'opn':
  kea_dhcpv4s => {
    'opnsense01.example.com' => {
      'ensure' => 'present',
      'config' => {
        'general' => {
          'enabled'          => '1',
          'interfaces'       => 'lan',
          'valid_lifetime'   => '4000',
          'fwrules'          => '1',
          'dhcp_socket_type' => 'raw',
        },
        'lexpire' => {
          'reclaim_timer_wait_time' => '10',
        },
        'ha' => {
          'enabled' => '0',
        },
      },
    },
  },
}
```
DHCPv4 subnets are managed via `kea_dhcpv4_subnets` or the `opn_kea_dhcpv4_subnet`
type. The resource title is the subnet CIDR (e.g. `192.168.1.0/24`). The provider
uses a search+get pattern to fetch full subnet details including option_data.

```
class { 'opn':
  kea_dhcpv4_subnets => {
    '192.168.1.0/24' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'description'             => 'LAN DHCP',
        'option_data_autocollect' => '1',
        'option_data'             => {
          'routers'             => '192.168.1.1',
          'domain_name_servers' => '8.8.8.8,8.8.4.4',
          'domain_name'         => 'example.com',
        },
        'pools' => '192.168.1.100 - 192.168.1.200',
      },
    },
  },
}
```
DHCPv4 reservations are managed via `kea_dhcpv4_reservations` or the
`opn_kea_dhcpv4_reservation` type. The resource title is the reservation
description. Reservations autorequire their parent subnet. The `subnet` field
accepts a subnet CIDR which is resolved to a UUID via the IdResolver.

```
class { 'opn':
  kea_dhcpv4_reservations => {
    'Web Server' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'subnet'     => '192.168.1.0/24',
        'hw_address' => 'AA:BB:CC:DD:EE:FF',
        'ip_address' => '192.168.1.10',
        'hostname'   => 'webserver',
      },
    },
  },
}
```
DHCPv4 HA peers are managed via `kea_dhcpv4_peers` or the `opn_kea_dhcpv4_peer`
type. The resource title is the peer name.

```
class { 'opn':
  kea_dhcpv4_peers => {
    'primary-node' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'role' => 'primary',
        'url'  => 'http://10.0.0.1:8000',
      },
    },
  },
}
```
DHCPv6 global settings are managed as a singleton resource per device via the
`kea_dhcpv6s` parameter or directly via the `opn_kea_dhcpv6` type. The `general`,
`lexpire` and `ha` sections are supported.

```
class { 'opn':
  kea_dhcpv6s => {
    'opnsense01.example.com' => {
      'ensure' => 'present',
      'config' => {
        'general' => {
          'enabled'    => '1',
          'interfaces' => 'lan',
        },
        'lexpire' => {
          'reclaim_timer_wait_time' => '10',
        },
        'ha' => {
          'enabled' => '0',
        },
      },
    },
  },
}
```
DHCPv6 subnets are managed via `kea_dhcpv6_subnets` or the `opn_kea_dhcpv6_subnet`
type. The resource title is the subnet CIDR (e.g. `fd00::/64`). The provider
uses a search+get pattern to fetch full subnet details.

```
class { 'opn':
  kea_dhcpv6_subnets => {
    'fd00::/64' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'description' => 'LAN DHCPv6',
        'interface'   => 'lan',
        'option_data' => {
          'dns_servers' => 'fd00::1',
        },
        'pools' => 'fd00::100 - fd00::200',
      },
    },
  },
}
```
DHCPv6 reservations are managed via `kea_dhcpv6_reservations` or the
`opn_kea_dhcpv6_reservation` type. Reservations autorequire their parent subnet.

```
class { 'opn':
  kea_dhcpv6_reservations => {
    'Mail Server' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'subnet'     => 'fd00::/64',
        'ip_address' => 'fd00::10',
        'duid'       => '01:02:03:04:05:06',
        'hostname'   => 'mailserver',
      },
    },
  },
}
```
DHCPv6 prefix delegation pools are managed via `kea_dhcpv6_pd_pools` or the
`opn_kea_dhcpv6_pd_pool` type. PD pools autorequire their parent subnet.

```
class { 'opn':
  kea_dhcpv6_pd_pools => {
    'Customer PD Pool' => {
      'devices' => ['opnsense01.example.com'],
      'ensure'  => 'present',
      'config'  => {
        'subnet'        => 'fd00::/64',
        'prefix'        => 'fd00:1::/48',
        'prefix_len'    => '48',
        'delegated_len' => '64',
      },
    },
  },
}
```
DHCPv6 HA peers are managed via `kea_dhcpv6_peers` or the `opn_kea_dhcpv6_peer`
type. The resource title is the peer name.

