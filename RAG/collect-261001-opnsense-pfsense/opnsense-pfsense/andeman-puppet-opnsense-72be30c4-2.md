---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/andeman-puppet-opnsense-72be30c4-2
title: "node: opnsense.example.com"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/andeman-puppet-opnsense-72be30c4.md
source_anchor: ""
source_lines: [222, 313]
sha256: 76371d310f52d028f892bcd624db3cc48311534b89f699a535d33288186c6622
---

# node: opnsense.example.com

```
# node: client1.example.com
class { 'opnsense::client::haproxy':
  servers  => {
    "client1.example.com" => {
      "devices"     => ["opnsense.example.com"],
      "description" => "client test server",
      "address"     => "client1.example.com",
      "port"        => "443",
      "enabled"     => ture,
    },
  },
  backends => {
    "web_backend" => {
      "devices"        => ["opnsense.example.com"],
      "description"    => "test backend",
      "mode"           => "http",
      "linked_servers" => ["server1", "server2"],
      "enabled"        => false,
    }
  },
  frontends => {
    "web_frontend" => {
      "devices"           => ["opnsense.example.com"],
      "description"       => "test frontend",
      "bind"              => "127.0.0.1:9000",
      "ssl_enabled"       => false,
      "default_backend"   => "localhost_backend",
      "enabled"           => true,
    }
  },
}
```
When connecting to the OPNsense API, this module will tell opn-cli to use the system-wide installed CA certificates to verify the SSL connection. However, this will only work when using a valid certificate for the OPNsense WebUI.

If the OPNsense WebUI still uses the pre-installed self-signed certificate, then it is possible to use the OPNsense CA certificate for SSL verification:

```
class { 'opnsense':
  use_system_ca => false,
  ca_file       => '/root/.opn-cli/ca.pem',
  ca_content    => '-----BEGIN CERTIFICATE-----
AAAAAABBBBBBBBBCCCCCCCCCCDDDDDDDDDDDEEEEEEEEEEEFFFFFFFFFGGGGGGGG
-----END CERTIFICATE-----'
}
```
The OPNsense CA certificate can be downloaded from `System: Trust: Authorities` on the OPNsense firewall.

You find more examples in the examples folder.

Types and providers are documented in REFERENCE.md.

For an extensive list of supported operating systems, see metadata.json

CI/CD is done via Github Actions.

You need to install the following requirements to setup the local development environment:

```
scripts/create_test_env 
```
Unit testing uses pdk

```
scripts/unit_tests
```
Acceptance testing uses puppet litmus.

```
scripts/acceptance_tests
```
```
scripts/remove_test_env
```
First prepare the release with:

```
./scripts/release_prep
```
This will set the version in `metadata.json`, create `REFERENCE.md` and  `CHANGELOG.md`.

Then commit the changes and push them to the repository.

Ensure that the following secrets are set in the github repository:

- FORGE_API_KEY (your puppet forge api key)

Please use the GitHub issues functionality to report any bugs or requests for new features. Feel free to fork and submit pull requests for potential contributions.

All contributions must pass all existing tests, new features should provide additional unit/acceptance tests.

See Changelog.
