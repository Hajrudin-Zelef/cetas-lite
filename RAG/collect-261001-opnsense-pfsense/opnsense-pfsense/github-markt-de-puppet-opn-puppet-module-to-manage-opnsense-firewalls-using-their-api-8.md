---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api-8
title: "Step 1: remove the old entry"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-markt-de-puppet-opn-puppet-module-to-manage-opnsense-firewalls-using-their-api.md
source_anchor: ""
source_lines: [1545, 1597]
sha256: 7b1be85519385552bed7fbbd59b2a3d2f841b355bb578e857605f5b6cc0c9d12
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
  manage_resources => true,
}
```
Singleton types (`opn_acmeclient_settings`, `opn_haproxy_settings`, `opn_hasync`, `opn_node_exporter`, `opn_puppet_agent`, `opn_zabbix_agent`, `opn_zabbix_proxy`) are excluded from exported resources because they represent per-device global settings that should only be managed on the management server.

The same parameters can also be configured via Hiera with deep merge:

```
# Client node Hiera data (e.g. role/webserver.yaml)
opn::client::firewall_aliases:
  webserver_ips:
    devices:
      - 'opnsense01.example.com'
    config:
      type: 'host'
      content: "%{facts.networking.ip}"
      description: "%{facts.networking.fqdn} - Web server IPs"
      enabled: '1'
opn::client::haproxy_servers:
  web01:
    devices:
      - 'opnsense01.example.com'
    config:
      address: "%{facts.networking.ip}"
      port: '8080'
      description: "%{facts.networking.fqdn} - Web backend"
      enabled: '1'
```
Migrating to this module is pretty simple: Use the `resource` command to convert the current OPNsense system configuration into Puppet code.

```
$ puppet resource opn_firewall_alias
$ puppet resource opn_zabbix_agent
$ puppet resource opn_zabbix_agent_userparameter
```
Note that this requires a valid puppet-opn configuration to work as expected. Hence enabling this module without additional parameters is the first step.

Classes and parameters are documented in REFERENCE.md.

All default values can be found in the `data/` directory.

Please use the GitHub issues functionality to report any bugs or requests for new features. Feel free to fork and submit pull requests for potential contributions.

BSD-2-Clause
