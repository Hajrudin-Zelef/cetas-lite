---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/andeman-opnsense-cli-fc216d2a-2
title: "update linked item with names"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-opnsense-pfsense/andeman-opnsense-cli-fc216d2a.md
source_anchor: ""
source_lines: [158, 363]
sha256: 8502ff10b3e7b435393ba32900703b8a98a7df076c0915ac0bf0b2689d409193
---

# update linked item with names

```
$ opn-cli plugin installed -o json_filter -c name,version
[{"name": "os-firewall", "version": "1.0_2"}, {"name": "os-virtualbox", "version": "1.0_1"},
```
Show the output separated by space.

```
$ opn-cli  plugin installed -o plain 
os-firewall 1.0_2 Firewall API supplemental package N/A
os-iperf 1.0_1 Connection speed tester N/A
os-virtualbox 1.0_1 VirtualBox guest additions N/A
os-zabbix-agent 1.8_2 Zabbix monitoring agent N/A
```
Show yaml output.

```
$ opn-cli plugin installed -o yaml 
- name: os-firewall
  version: '1.0_2'
  comment: Firewall API supplemental package
  locked: N/A
- name: os-iperf
  version: '1.0_1'
  comment: Connection speed tester
  locked: N/A
- name: os-virtualbox
  version: '1.0_1'
  comment: VirtualBox guest additions
  locked: N/A
- name: os-zabbix-agent
  version: '1.8_2'
  comment: Zabbix monitoring agent
  locked: N/A
```
To assist the rapid development of new opn-cli commands, the code generator generates scaffolding code which could be
used as a **starting point** for opnsense core or plugin modules.

```
$  opn-cli new -h
Usage: opn-cli new [OPTIONS] COMMAND [ARGS]...
Generate scaffolding code
Options:
-h, --help  Show this message and exit.
Commands:
command  Generate code for a new command
```
This generates all necessary code to implement api calls to a core module.

For a list of all plugin api endpoints see:

Or:

```
$ opn-cli new api list --module-type core
```
Example that generates code for the core module 'cron':

```
$ opn-cli new api core cron
```
This generates a class for every controller of the core module. Every class contains methodes to call all api endpoints of the corresponding controller.

Please move the file from the output dir to the destination folders under opnsense_cli/. The default output path is opnsense_cli/output/api/core/ .

This generates all necessary code to implement api calls to a plugin module.

For a list of all plugin api endpoints see:

Or:

```
$ opn-cli new api list --module-type plugin
```
Example that generates code for the plugin module 'haproxy':

```
$ opn-cli new api plugin haproxy
```
This generates a class for every controller of the plugin module. Every class contains methodes to call all api endpoints of the corresponding controller.

Please move the file from the output dir to the destination folders under opnsense_cli/. The default output path is opnsense_cli/output/api/plugins/ .

This generates all necessary code and tests to implement a new command from a core module.

For a list of core modules see:

Search model.xml files here:

To add help texts, you need to specify a form.xml for the module. For Unbound dot the form.xml url is:

Make sure to pass text/plain content from raw.githubusercontent.com instead of github.com.

Examples:

```
$ opn-cli new command core unbound dot --tag dots \
-m https://raw.githubusercontent.com/opnsense/core/master/src/opnsense/mvc/app/models/OPNsense/Unbound/Unbound.xml \
-f https://raw.githubusercontent.com/opnsense/core/master/src/opnsense/mvc/app/controllers/OPNsense/Unbound/forms/dialogDot.xml
```
This generates a command class, facade class and a integration test.

Please move them from the output dir to the destination folders under opnsense_cli/, import the files in "cli.py" and register the commands groups and commands there.

The facade use API classes which should be generated with "opn-cli new api plugin" command. Make sure to remove all unnecessary API classes and methods to have a proper code coverage.

After some tweaks you should be able to use the new command.

This generates all necessary code and tests to implement a new command from a plugin module.

For a list of plugin modules see:

Search for model.xml and form.xml files here:

Make sure to pass text/plain content from raw.githubusercontent.com instead of github.com.

Examples:

```
$ opn-cli new command plugin haproxy server --tag servers \
-m https://raw.githubusercontent.com/opnsense/plugins/master/net/haproxy/src/opnsense/mvc/app/models/OPNsense/HAProxy/HAProxy.xml \
-f https://raw.githubusercontent.com/opnsense/plugins/master/net/haproxy/src/opnsense/mvc/app/controllers/OPNsense/HAProxy/forms/dialogServer.xml
```
This generates a command class, facade class and a integration test.

Please move them from the output dir to the destination folders under opnsense_cli/, import the files in "cli.py" and register the commands groups and commands there.

The facade use API classes which should be generated with "opn-cli new api plugin" command. Make sure to remove all unnecessary API classes and methods to have a proper code coverage.

After some tweaks you should be able to use the new command.

This generates files needed for a custom resource type for puppet from an exiting opn-cli command.

Takes the name of the opn-cli command and the namevar as arguments.

The namevar is the column name that puppet will use to find the uuid of the resource.

**Examples:**

Generate files for the opn-cli command "opn-cli route static" and use the column "descr" as the namevar:

```
$ opn-cli new puppet resource-type route static descr
```
If you want to link items with options, you could link them with uuids or with their names.

If the name exists mulitple times, the uuid from the first match is used.

**Example:**

```
$ opn-cli haproxy server list -c 'uuid,name'
+--------------------------------------+----------+
|                 uuid                 |   name   |
+--------------------------------------+----------+
| 162e7c70-dbea-4813-8676-33c506e1b1e2 | server1  |
| a293d23f-5fa5-46e5-b724-47f0031d8e9b | server2  |
+--------------------------------------+----------+
$ opn-cli haproxy backend list -c 'uuid,name,Servers'
+--------------------------------------+--------+-----------------+
|                 uuid                 |  name  |     Servers     |
+--------------------------------------+--------+-----------------+
| 54def7b3-93a8-4ea9-8858-22131948fb91 | pool1  |                 |
| c30cf40b-e026-4316-ad81-32affe0c1d85 | pool2  |                 |
+--------------------------------------+--------+-----------------+
# update linked item with names
$ opn-cli haproxy backend update 54def7b3-93a8-4ea9-8858-22131948fb91 --linkedServers server1,server2
# or update linked items with uuids
$ opn-cli haproxy backend update c30cf40b-e026-4316-ad81-32affe0c1d85 --linkedServers a293d23f-5fa5-46e5-b724-47f0031d8e9b
+--------------------------------------+--------+-----------------+
|                 uuid                 |  name  |     Servers     |
+--------------------------------------+--------+-----------------+
| 54def7b3-93a8-4ea9-8858-22131948fb91 | pool1  | server1,server2 |
| c30cf40b-e026-4316-ad81-32affe0c1d85 | pool2  |     server2     |
+--------------------------------------+--------+-----------------+
# delete linked items
$ opn-cli haproxy backend update c30cf40b-e026-4316-ad81-32affe0c1d85 --linkedServers ''
+--------------------------------------+--------+-----------------+
|                 uuid                 |  name  |     Servers     |
+--------------------------------------+--------+-----------------+
| 54def7b3-93a8-4ea9-8858-22131948fb91 | pool1  | server1,server2 |
| c30cf40b-e026-4316-ad81-32affe0c1d85 | pool2  |                 |
+--------------------------------------+--------+-----------------+
```
This feature needs the opnsense plugin os-firewall.

```
$ opn-cli plugin install os-firewall
```
See: https://wiki.opnsense.org/manual/aliases.html

```
Usage: opn-cli firewall alias [OPTIONS] COMMAND [ARGS]...
  Manage OPNsense firewall aliases.
  See: https://wiki.opnsense.org/manual/aliases.html
Options:
  -h, --help  Show this message and exit.
Commands:
  create  Create a new alias.
  delete  Delete an alias
  list    Show all aliases
  show    Show details for alias
  table   Show pf table entries for alias
  update  Update an alias.
```
See: https://docs.opnsense.org/manual/firewall.htm

