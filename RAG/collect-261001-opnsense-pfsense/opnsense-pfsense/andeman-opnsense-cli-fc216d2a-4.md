---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/andeman-opnsense-cli-fc216d2a-4
title: "update linked item with names"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/andeman-opnsense-cli-fc216d2a.md
source_anchor: ""
source_lines: [599, 764]
sha256: 34b7c020d82baee780f7101570ceda482c1716359e8650bd11d9099c73659c9e
---

# update linked item with names

```
Usage: opn-cli nodeexporter config [OPTIONS] COMMAND [ARGS]...
  Manage nodeexporter config
Options:
  -h, --help  Show this message and exit.
Commands:
  edit  Edit configuration
  show  Show configuration
```
```
Usage: opn-cli syslog destination [OPTIONS] COMMAND [ARGS]...
  Manage syslog destination
Options:
  -h, --help  Show this message and exit.
Commands:
  create  Create a new destination
  delete  Delete destination
  list    Show all destination
  show    Show details for destination
  update  Update a destination.
```
```
Usage: opn-cli syslog stats list [OPTIONS]
  Show syslog statistics
Options:
  --search TEXT                   Search for this string
  -o, --output [cols|table|json|json_filter|plain|yaml]
                                  Specifies the Output format.  [default:
                                  table]
  -c, --cols TEXT                 Which columns should be printed? Pass empty
                                  string (-c ) to show all columns  [default: #
                                  ,Description,SourceName,SourceId,SourceInsta
                                  nce,State,Type,Number]
  -h, --help                      Show this message and exit.
```
```
Usage: opn-cli openvpn [OPTIONS] COMMAND [ARGS]...
  Export OpenVPN configuration.
Options:
  -h, --help  Show this message and exit.
Commands:
  accounts   Show all accounts for an OpenVPN server.
  download   Download client config for chosen OpenVPN server and account.
  providers  Show all available OpenVPN servers.
  templates  Show all available export templates.
```
```
Usage: opn-cli plugin [OPTIONS] COMMAND [ARGS]...
  OPNsense plugins management
Options:
  -h, --help  Show this message and exit.
Commands:
  install    Install plugin by name
  installed  Show installed plugins.
  list       Show all available plugins.
  lock       Lock plugin.
  reinstall  Reinstall plugin by name.
  show       Show plugin details.
  uninstall  Uninstall plugin by name.
  unlock     Unlock plugin.
```
```
Usage: opn-cli unbound host [OPTIONS] COMMAND [ARGS]...
  Manage host overrides
Options:
  -h, --help  Show this message and exit.
Commands:
  create  Create a new host override
  delete  Delete a host override
  list    Show all hosts overrides
  show    Show details for host override
  update  Update a host override
```
```
Usage: opn-cli unbound alias [OPTIONS] COMMAND [ARGS]...
  Manage unbound host alias overrides
Options:
  -h, --help  Show this message and exit.
Commands:
  create  Create a new alias
  delete  Delete alias
  list    Show all alias
  show    Show details for alias
  update  Update an alias.
```
```
Usage: opn-cli unbound domain [OPTIONS] COMMAND [ARGS]...
  Manage unbound domain overrides
Options:
  -h, --help  Show this message and exit.
Commands:
  create  Create a new domain
  delete  Delete domain
  list    Show all domain
  show    Show details for domain
  update  Update a domain.
```
```
# host
opn-cli unbound host create --hostname '*' --domain example.com --rr A --server 192.168.1.254
opn-cli unbound host create --hostname 'mailin' --domain example.com --rr MX --mxprio 10 --mx mail.example.com
# alias
opn-cli unbound alias create --host '*|example.com|A|||192.168.1.254' --hostname another --domain example.com
opn-cli unbound alias create --host 'mailin|example.com|MX|10|mail.example.com|' --hostname mail03 --domain example.com
# domain
opn-cli unbound domain create --domain rockin.com --server 192.168.56.3
```
This feature needs the opnsense plugin os-api-backup. This plugin was deprecated with the OPNsense 24.1 release. So if you are running OPNsense 24.1 or higher use the Configbackup command.

```
$ opn-cli plugin install os-api-backup
```
Download a backup of the OPNsense system configuration to the current directory:

```
$ opn-cli apibackup backup download
successfully saved to: ./config.xml
```
Or specify a path and filename:

```
$ opn-cli apibackup backup download -p /tmp/config_backup.xml
successfully saved to: /tmp/config_backup.xml
```
The plugin "os-api-backup" was discontinued in OPNsense Version 24.1, because the core API provides the same functionality. This command provides the exact same functionality than Apibackup but uses the OPNsense Core API-Endpoint. So if you are running a OPNsense Instance version 24.1 or higher use this command for configuration backups instead of "apibackup".

Download a backup of the OPNsense system configuration to the current directory:

```
$ opn-cli configbackup backup download
successfully saved to: ./config.xml
```
Or specify a path and filename:

```
$ opn-cli configbackup backup download -p /tmp/config_backup.xml
successfully saved to: /tmp/config_backup.xml
```
Requirements:

- vagrant
- virtualbox

This will install opn-cli, add it to your PATH and setup an opnsense vm with vagrant:

```
# setup 
scripts/create_test_env
# test opn-cli
opn-cli plugin list
# teardown 
scripts/remove_test_env
```
```
# lint your code
scripts/lint
# execute all unit tests
scripts/unit_tests
# execute a single test
scripts/unit_tests opnsense_cli/tests/command/tests/test_format.py::TestFormatter
# execute tests with coverage report
scripts/coverage
```
Please use the GitHub issues functionality to report any bugs or requests for new features. Feel free to fork and submit pull requests for potential contributions.

All contributions must pass all existing tests, new features should provide additional unit/acceptance tests.
