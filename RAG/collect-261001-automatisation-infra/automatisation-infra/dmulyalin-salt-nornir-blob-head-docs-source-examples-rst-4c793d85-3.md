---
id: collect-261001-automatisation-infra/automatisation-infra/dmulyalin-salt-nornir-blob-head-docs-source-examples-rst-4c793d85-3
title: "apply logging configuration using jinja2 template"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "throughput"]
source: docs/RAG/collect-261001-automatisation-infra/dmulyalin-salt-nornir-blob-head-docs-source-examples-rst-4c793d85.md
source_anchor: ""
source_lines: [437, 517]
sha256: 5be16af4ec01a806e469ed2e849d59508759e6d89ca88d28e351a35aa7407275
---

# apply logging configuration using jinja2 template

            Ethernet0/0.107            10.1.107.7      YES NVRAM  up                    up
            Ethernet0/0.117            10.1.117.7      YES NVRAM  up                    up
            Ethernet0/0.2000           192.168.217.7   YES NVRAM  up                    up
            Ethernet0/1                unassigned      YES NVRAM  administratively down down
            Ethernet0/2                unassigned      YES NVRAM  administratively down down
            Ethernet0/3                unassigned      YES NVRAM  administratively down down
            Loopback0                  10.0.0.7        YES NVRAM  up                    up
[root@localhost /]# tree /var/salt-nornir/nrp1/files/
├── show_commands_output__11_July_2021_07_11_26__IOL1.txt
├── show_commands_output__11_July_2021_07_11_26__IOL2.txt
├── tf_aliases.json
[root@localhost /]# cat /var/salt-nornir/nrp1/files/show_commands_output__11_July_2021_07_11_26__IOL1.txt
*12:05:06.633 EET Sun Feb 14 2021
Interface                  IP-Address      OK? Method Status                Protocol
Ethernet0/0                unassigned      YES NVRAM  up                    up
Ethernet0/0.102            10.1.102.10     YES NVRAM  up                    up
Ethernet0/0.107            10.1.107.10     YES NVRAM  up                    up
Ethernet0/0.2000           192.168.217.10  YES NVRAM  up                    up
Ethernet0/1                unassigned      YES NVRAM  up                    up
Ethernet0/2                unassigned      YES NVRAM  up                    up
Ethernet0/3                unassigned      YES NVRAM  administratively down down
Loopback0                  10.0.0.10       YES NVRAM  up                    up
Loopback100                1.1.1.100       YES NVRAM  up                    up
Salt-Norir hcache functionality allows to cache devices output using in-memory
Nornir inventory, that in return allows to refer to cached data within jinja2 templates
used by nr.cfg function.
This example demonstrates salt cli command to perform these tasks:
- save device logging configuration into hcache
- use nr.cfg_gen to verify cached logging configuration
- use nr.cfg to re-apply cached logging configuration to device
Save devices output into hcache:
salt nrp1 nr.cli "show run | inc logging" hcache="log_config"
salt nrp1 nr.cfg_gen '{{ host.log_config["show run | inc logging"] }}'
salt nrp1 nr.cfg '{{ host.log_config["show run | inc logging"] }}'
NAPALM example to configure SSH jumphost in ~/.ssh/config:
host *.*
    ProxyCommand ssh -W %h:%p myuser@jumphost.company.com -i /run/secrets/ssh_key -o StrictHostKeyChecking=no -o ControlPath=/dev/shm/cm-%r@%h:%p -o ControlMaster=auto -o ControlPersist=10m -o IdentitiesOnly=yes
host *
UserKnownHostsFile /dev/null
IdentitiesOnly yes
IPQoS=throughput
StrictHostKeyChecking no
Netmiko and Scrapli example to configure SSH jumphost in ~/.ssh/config:
host jumphost
user myuser
hostname jumphost.company.com
ControlPath /dev/shm/cm-%r@%h:%p
ControlMaster auto
ControlPersist 10m
host *.*
ProxyJump jumphost
IdentityFile /run/secrets/ssh_key
StrictHostKeyChecking no
host *
StrictHostKeyChecking no
UserKnownHostsFile /dev/null
IdentitiesOnly yes
IPQoS=throughput
Salt-Nornir Pillar to use ~/.ssh/config configuration:
hosts:
  core-rtr-1:
    hostname: 192.168.1.10
    username: GENERIC_USERNAME
    password: GENERIC_PASSWORD
    connection_options:
      scrapli:
        platform: scrapli_platform
        extras
          auth_strict_key: False
          ssh_config_file: "/home/user_name/ssh/config"
          transport_options:
            open_cmd: ["-o", "KexAlgorithms=+diffie-hellman-group-exchange-sha1", "-o", "Ciphers=+aes256-cbc"]
      napalm:
        platform: napalm_platform
        extras:
          optional_args:
            ssh_config_file: "/home/user_name/ssh/config"
      netmiko:
        platform: netmiko_platform
        extras:
          ssh_config_file: "/home/user_name/ssh/config"
