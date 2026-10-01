---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/andeman-puppet-opnsense-72be30c4-1
title: "node: opnsense.example.com"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-opnsense-pfsense/andeman-puppet-opnsense-72be30c4.md
source_anchor: ""
source_lines: [1, 221]
sha256: d01e78066d7acbcdc1595a4831b8825e0ec33dd4b692277a0ee8d99b70dd1a26
---

# node: opnsense.example.com

The opnsense module configures OPNsense firewalls.

It allows administrators to manage an OPNsense firewall directly via the sysutils/puppet-agent opnsense plugin and/or manage multiple firewalls from a bastion host running a puppet-agent with opn-cli installed.

The main target of module is to enable GitOps for your network security policies. Developers could submit pull request for new firewall rules and loadbalancer configurations and the network or ops team could review it and deploy it to a pre production environment for testing and verification. If everything passes, you could deploy it to production.

You can automate the following with the module:

- plugins
- firewall aliases
- firewall rules
- haproxy servers
- haproxy backends
- haproxy frontends
- prometheus nodeexporter
- syslog destinations
- static routes

If you want to manage your firewall directly with a puppet-agent running on the device.

OPNsense plugins:

- sysutils/puppet-agent
- os-firewall for managing firewall rules
- os-haproxy for managing haproxy rules

```
Menu->Firmware->Plugins
Install plugin: sysutils/puppet-agent
```
If you want a bastion hosts running a puppet-agent which could manage multiple firewalls via https API calls.

- opn-cli >= 1.6.0

```
$packages = [
    'python3',
    'python3-pip',
]
$pip_packages = [
    'opn-cli',
]
package { $packages:
    ensure => present,
}
-> package { $pip_packages:
    ensure   => latest,
    provider => 'pip3',
}
```
If you want to manage an OPNsense firewall, you need to supply credentials and connection information for the device.

To create an api_key and api_secret see: https://docs.opnsense.org/development/how-tos/api.html#creating-keys.

**If you want to use ssl verification (recommended):**`

To download the default self-signed cert, open the OPNsense web gui and go to System->Trust->Certificates. Search for the name: "Web GUI SSL certificate" and press the "export user cert" button.

If you use a ca signed certificate, go to System->Trust->Authorities and press the "export CA cert" button to download the ca.

Save the cert or ca and make sure the puppet agent is able to read it.

```
include opnsense
```
You can manage multiple opnsense firewalls with this module.

In the following example a single OPNsense firewall running a puppet agent is manged which allows clients to export configuration via exported resources (manage_resources => true):

```
# node: opnsense.example.com
class { 'opnsense':
  manage_resources => true,
  devices => {
    'opnsense.example.com' => {
      'url'        => 'https://127.0.0.1/api',
      'api_key'    => 'your_api_key',
      'api_secret' => 'your_api_secret',
      'ssl_verify' => true,
      'timeout'    => 60,
      'ca'         => '~/.opn-cli/ca.pem',
      'plugins'    => {
        'os-helloworld' => {}
      },
      nodeexporter => {
        enabled        => true,
        listen_address => '192.168.1.1',
        listen_port    => '9200',
        cpu            => false,
        exec           => false,
        filesystem     => false,
        loadavg        => false,
        meminfo        => false,
        netdev         => false,
        time           => false,
        devstat        => false,
        interrupts     => true,
        ntp            => true,
        zfs            => true,
      },
      "ensure"      => "present"      
    }
  },
  firewall => {
    aliases => {
      'my_http_ports_local' => {
        'devices'     => ['opnsense.example.com'],
        'type'        => 'port',
        'content'     => ['80', '443'],
        'description' => 'example local http ports',
        'enabled'     => true,
        'ensure'      => present
      },
    },
    rules => {
      'allow all from lan' => {
        'devices'   => ['opnsense.example.com'],
        'sequence'  => '1',
        'action'    => 'pass',
        'interface' => ['lan']
      }
    }
  },
  syslog => {
    destinations => {
      'syslogger 1' => {
        devices     => ['opnsense.example.com'],
        enabled     => true,
        transport   => 'tcp4',
        program     => 'ntp,ntpdate',
        level       => ['crit', 'alert', 'emerg'],
        facility    => ['ntp'],
        hostname    => 'syslog.example.com',
        certificate => '',
        port        => '10514',
        rfc5424     => true,
        ensure      => present,
      },
    },
  },
  route => {
    static => {
      'static route 1' => {
        devices    => ['opnsense.example.com'],
        network    => '10.0.0.98/24',
        gateway    => 'WAN_DHCP',
        disabled   => false,
        ensure     => 'present',
      },
    },
  },
  haproxy => {
    servers => {
      "server1" => {
        "devices"     => ["opnsense.example.com"],
        "description" => "first local server",
        "address"     => "127.0.0.1",
        "port"        => "8091",
      },
      "server2" => {
        "devices"   => ["opnsense.example.com"],
        "description" => "second local server",
        "address"     => "127.0.0.1",
        "port"        => "8092",
      },
    },
    backends => {
      "localhost_backend" => {
        "devices"        => ["opnsense.example.com"],
        "description"    => "local server backend",
        "mode"           => "http",
        "linked_servers" => ["server1", "server2"],
      }
    },
    frontends => {
      "localhost_frontend" => {
        "devices"           => ["opnsense.example.com"],
        "description"       => "local frontend",
        "bind"              => "127.0.0.1:8090",
        "ssl_enabled"       => false,
        "default_backend"   => "localhost_backend",
      }
    },
  },
}
```
This feature use exported resources. You need to enable catalog storage and searching (storeconfigs) on your primary puppet server.

Here the client (client1.example.com) is exporting it´s security configuration to the firewall (opnsense.example.com) defined above:

```
# node: client1.example.com
class { 'opnsense::client::firewall':
  aliases => {
    'client1_example_com' => {
      'devices'     => ['opnsense.example.com'],
      'type'        => 'host',
      'content'     => ['client1.example.com'],
      'description' => 'client.example.com alias',
      'enabled'     => true,
      'ensure'      => present
    },
  },
  rules => {
    'allow https from lan to client1.example.com' => {
      'devices'          => ['opnsense.example.com'],
      'sequence'         => '100',
      'action'           => 'pass',
      'interface'        => ['lan'],
      'protocol'         => 'TCP',
      'destination_net'  => 'client1_example_com',
      'destination_port' => 'https',
      'ensure'           => present
    },
  }
}
```
This feature use exported resources. You need to enable catalog storage and searching (storeconfigs) on your primary puppet server.

Here the client (client1.example.com) is exporting it´s haproxy configuration to the firewall (opnsense.example.com) defined above:

