---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api-7
title: "Step 1: remove the old entry"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api.md
source_anchor: ""
source_lines: [1247, 1544]
sha256: ed0d152cd698b8667a0acb065c54088f53b2e31dd0cab70f02a882b27ec58bb1
---

# Step 1: remove the old entry

```
# Setup device configs.
class { 'opn::config':
  devices => {
    'opnsense01.example.com' => {
      'url'        => 'https://opnsense01.example.com/api',
      'api_key'    => 'key',
      'api_secret' => 'secret',
    },
  },
}
opn_cron { 'Daily haproxy reload@opnsense01.example.com':
  ensure => present,
  config => {
    'command'  => 'haproxy reload',
    'minutes'  => '0',
    'hours'    => '3',
    'days'     => '*',
    'months'   => '*',
    'weekdays' => '*',
    'enabled'  => '1',
  },
}
opn_gateway { 'TEST_GW@opnsense01.example.com':
  ensure => present,
  config => {
    'interface'       => 'lan',
    'ipprotocol'      => 'inet',
    'gateway'         => '192.168.123.1',
    'descr'           => 'TEST Gateway',
    'monitor_disable' => '1',
  },
}
opn_plugin { 'os-haproxy@opnsense01.example.com':
  ensure => present,
}
opn_firewall_alias { 'http_ports@opnsense01.example.com':
  ensure => present,
  config => {
    'type'        => 'port',
    'content'     => "80\n443",
    'description' => 'HTTP(S) ports',
    'enabled'     => '1',
  },
}
opn_firewall_category { 'web@opnsense01.example.com':
  ensure => present,
  config => { 'color' => '0088cc' },
}
opn_firewall_group { 'dmz_ifaces@opnsense01.example.com':
  ensure => present,
  config => {
    'members' => 'em1,em2',
    'descr'   => 'DMZ interfaces',
  },
}
opn_firewall_rule { 'Allow SSH from mgmt@opnsense01.example.com':
  ensure => present,
  config => {
    'action'          => 'pass',
    'interface'       => 'lan',
    'protocol'        => 'tcp',
    'source_net'      => 'mgmt_hosts',
    'destination_port'=> '22',
    'enabled'         => '1',
  },
}
opn_user { 'jdoe@opnsense01.example.com':
  ensure => present,
  config => {
    'password'    => 'plaintextpassword',
    'description' => 'John Doe',
  },
}
opn_group { 'vpn_users@opnsense01.example.com':
  ensure => present,
  config => { 'description' => 'VPN Users' },
}
opn_haproxy_server { 'web01@opnsense01.example.com':
  ensure => present,
  config => {
    'address'     => '10.0.0.1',
    'port'        => '80',
    'description' => 'Web server 01',
    'enabled'     => '1',
  },
}
opn_haproxy_backend { 'web_pool@opnsense01.example.com':
  ensure => present,
  config => {
    'mode'        => 'http',
    'description' => 'Web backend pool',
    'enabled'     => '1',
  },
}
opn_haproxy_frontend { 'http_in@opnsense01.example.com':
  ensure => present,
  config => {
    'bind'        => '0.0.0.0:80',
    'mode'        => 'http',
    'description' => 'HTTP listener',
    'enabled'     => '1',
  },
}
opn_haproxy_settings { 'opnsense01.example.com':
  ensure => present,
  config => {
    'general' => {
      'enabled' => '1',
    },
  },
}
opn_hasync { 'opnsense01.example.com':
  ensure => present,
  config => {
    'pfsyncenabled'   => '1',
    'pfsyncinterface' => 'lan',
    'synchronizetoip' => '10.0.0.2',
    'username'        => 'root',
    'password'        => 'secret',
  },
}
opn_node_exporter { 'opnsense01.example.com':
  ensure => present,
  config => {
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
    'zfs'           => '1',
  },
}
opn_puppet_agent { 'opnsense01.example.com':
  ensure => present,
  config => {
    'enabled'           => '1',
    'fqdn'              => 'puppet.example.com',
    'environment'       => 'production',
    'runinterval'       => '30m',
    'runtimeout'        => '1h',
    'usecacheonfailure' => '1',
  },
}
opn_route { 'Server network@opnsense01.example.com':
  ensure => present,
  config => {
    'network'  => '10.0.0.0/24',
    'gateway'  => 'Wan_DHCP',
    'disabled' => '0',
  },
}
opn_snapshot { 'pre-upgrade@opnsense01.example.com':
  ensure => present,
  active => true,
  config => {
    'note' => 'Snapshot before upgrade',
  },
}
opn_syslog { 'Central syslog@opnsense01.example.com':
  ensure => present,
  config => {
    'transport' => 'udp4',
    'hostname'  => 'syslog.example.com',
    'port'      => '514',
    'enabled'   => '1',
  },
}
opn_dhcrelay_destination { 'DHCP Servers@opnsense01.example.com':
  ensure => present,
  config => {
    'server' => '10.0.0.1,10.0.0.2',
  },
}
opn_dhcrelay { 'LAN IPv4 Relay@opnsense01.example.com':
  ensure => present,
  config => {
    'interface'   => 'lan',
    'destination' => 'DHCP Servers',
    'enabled'     => '1',
  },
}
opn_trust_ca { 'Internal CA@opnsense01.example.com':
  ensure => present,
  config => {
    'action'     => 'internal',
    'key_type'   => 'RSA',
    'digest'     => 'SHA256',
    'lifetime'   => '3650',
    'country'    => 'DE',
    'commonname' => 'Internal CA',
  },
}
opn_trust_cert { 'web.example.com@opnsense01.example.com':
  ensure => present,
  config => {
    'action'     => 'internal',
    'caref'      => '<ca-uuid>',
    'key_type'   => 'RSA',
    'digest'     => 'SHA256',
    'lifetime'   => '365',
    'commonname' => 'web.example.com',
  },
}
opn_trust_crl { 'Internal CA@opnsense01.example.com':
  ensure => present,
  config => {
    'descr'     => 'CRL for Internal CA',
    'lifetime'  => '9999',
    'crlmethod' => 'internal',
  },
}
opn_tunable { 'kern.maxproc@opnsense01.example.com':
  ensure => present,
  config => {
    'value'       => '4096',
    'description' => 'Maximum number of processes',
  },
}
opn_zabbix_proxy { 'opnsense01.example.com':
  ensure => present,
  config => {
    'enabled'    => '1',
    'server'     => 'zabbix.example.com',
    'serverport' => '10051',
    'hostname'   => 'opnsense01-proxy',
  },
}
opn_zabbix_agent { 'opnsense01.example.com':
  ensure => present,
  config => {
    'local' => {
      'hostname' => 'opnsense01.example.com',
    },
    'settings' => {
      'main' => {
        'enabled'    => '1',
        'serverList' => 'zabbix.example.com',
        'listenPort' => '10050',
      },
    },
  },
}
opn_zabbix_agent_userparameter { 'custom.uptime@opnsense01.example.com':
  ensure => present,
  config => {
    'command'      => '/usr/bin/uptime',
    'enabled'      => '1',
    'acceptParams' => '0',
  },
}
opn_zabbix_agent_alias { 'ping@opnsense01.example.com':
  ensure => present,
  config => {
    'sourceKey'    => 'icmpping',
    'enabled'      => '1',
    'acceptParams' => '0',
  },
}
```
Exported resources allow application servers (client nodes) to declare OPNsense resources that are collected and applied by the management server. This requires PuppetDB to be set up.

**Client node** — use the `opn::client` class to export resources. Each resource item must include a `devices` key listing the target OPNsense device names. You can use facts like `$facts['networking']['fqdn']` and `$facts['networking']['ip']` in config values to identify the origin node.

```
class { 'opn::client':
  firewall_aliases => {
    'webserver_ips' => {
      'devices' => ['opnsense01.example.com'],
      'config'  => {
        'type'        => 'host',
        'content'     => $facts['networking']['ip'],
        'description' => "${facts['networking']['fqdn']} - Web server IPs",
        'enabled'     => '1',
      },
    },
  },
  haproxy_servers => {
    'web01' => {
      'devices' => ['opnsense01.example.com', 'opnsense02.example.com'],
      'config'  => {
        'address'     => $facts['networking']['ip'],
        'port'        => '8080',
        'description' => "${facts['networking']['fqdn']} - Web backend",
        'enabled'     => '1',
      },
    },
  },
}
```
**Management server** — enable collection with `manage_resources => true` on the `opn` class. Collected resources automatically depend on the per-device credential file.

