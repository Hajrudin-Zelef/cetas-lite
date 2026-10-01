---
id: collect-261001-automatisation-infra/automatisation-infra/dmulyalin-salt-nornir-blob-head-docs-source-examples-rst-4c793d85-1
title: "apply logging configuration using jinja2 template"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/dmulyalin-salt-nornir-blob-head-docs-source-examples-rst-4c793d85.md
source_anchor: ""
source_lines: [1, 253]
sha256: ebfa98f5473a404bea1f510d4b7c99bf26bcc774506f86eed60c39bd46d74d41
---

# apply logging configuration using jinja2 template

With NAPALM plugin:
salt nrp1 nr.cfg "logging host 1.1.1.1" "ntp server 1.1.1.2" plugin=napalm
With Netmiko plugin:
salt nrp1 nr.cfg "logging host 1.1.1.1" "ntp server 1.1.1.2" plugin=netmiko
With Scrapli plugin:
salt nrp1 nr.cfg "logging host 1.1.1.1" "ntp server 1.1.1.2" plugin=scrapli
With PyATS plugin:
salt nrp1 nr.cfg "logging host 1.1.1.1" "ntp server 1.1.1.2" plugin=pyats
Make sure that device configured accordingly and NAPALM or Scrapli or Netmiko can interact with it, e.g. for NAPALM SCP server enabled on Cisco IOS or Scrapli library installed on minion machine or hosts' connection options specified in inventory, e.g.:
hosts:
  IOL1:
    hostname: 192.168.217.10
    platform: ios
    groups: [lab]
    connection_options:
      scrapli:
        platform: cisco_iosxe
        port: 22
        extras:
          ssh_config_file: True
          auth_strict_key: False
Data defined in pillar file /etc/salt/pillar/nrp1.sls:
hosts:
  IOL1:
    hostname: 192.168.217.10
    platform: ios
    groups: [lab]
    data:
      syslog: ["1.1.1.1", "2.2.2.2"]
  IOL2:
    hostname: 192.168.217.7
    platform: ios
    groups: [lab]
    data:
      syslog: ["1.1.1.2", "2.2.2.1"]
groups:
  lab:
    username: nornir
    password: nornir
Combine data with template file /etc/salt/template/nr_syslog_cfg.j2:
hostname {{ host.name }}
!
{%- for server in host.syslog %}
logging host {{ server }}
{%- endfor %}
Hosts' data injected in templates under host variable.
First, optional, generate configuration using nr.cfg_gen function without applying it to devices:
[root@localhost /]# salt nrp1 nr.cfg_gen filename=salt://templates/nr_syslog_cfg.j2
nrp1:
    ----------
    IOL1:
        ----------
        Rendered salt://templates/nr_syslog_cfg.j2 config:
            hostname IOL1
            !
            logging host 1.1.1.1
            logging host 2.2.2.2
    IOL2:
        ----------
        Rendered salt://templates/nr_syslog_cfg.j2 config:
            hostname IOL2
            !
            logging host 1.1.1.2
            logging host 2.2.2.1
[root@localhost /]#
If configuration looks ok, can apply it to devices:
[root@localhost /]# salt nrp1 nr.cfg filename=salt://templates/nr_syslog_cfg.j2 plugin=netmiko
nrp1:
    ----------
    IOL1:
        ----------
        netmiko_send_config:
            ----------
            changed:
                True
            diff:
            exception:
                None
            failed:
                False
            result:
                configure terminal
                Enter configuration commands, one per line.  End with CNTL/Z.
                IOL1(config)#hostname IOL1
                IOL1(config)#!
                IOL1(config)#logging host 1.1.1.1
                IOL1(config)#logging host 2.2.2.2
                IOL1(config)#end
    IOL2:
        ----------
        netmiko_send_config:
            ----------
            changed:
                True
            diff:
            exception:
                None
            failed:
                False
            result:
                IOL2#configure terminal
                IOL2(config)#hostname IOL2
                IOL2(config)#!
                IOL2(config)#logging host 1.1.1.2
                IOL2(config)#logging host 2.2.2.1
                IOL2(config)#end
                IOL2#
Verify configuration applied:
[root@localhost /]# salt nrp1 nr.cli "show run | inc logging"
nrp1:
    ----------
    IOL1:
        ----------
        show run | inc logging:
            logging host 1.1.1.1
            logging host 2.2.2.2
    IOL2:
        ----------
        show run | inc logging:
            logging host 1.1.1.2
            logging host 2.2.2.1
Sample Salt Master configuration excerpt defining base environment pillar and states location,
file /etc/salt/master snippet:
...
file_roots:
  base:
    - /etc/salt
    - /etc/salt/states
pillar_roots:
  base:
    - /etc/salt/pillar
...
Define data in pillar file /etc/salt/pillar/nrp1.sls:
hosts:
  IOL1:
    hostname: 192.168.217.10
    platform: ios
    groups: [lab]
    data:
      syslog: ["1.1.1.1", "2.2.2.2"]
  IOL2:
    hostname: 192.168.217.7
    platform: ios
    groups: [lab]
    data:
      syslog: ["1.1.1.2", "2.2.2.1"]
groups:
  lab:
    username: nornir
    password: nornir
Jinja2 template used with state to configure syslog servers, file salt://templates/nr_syslog_cfg.j2
same as absolute path /etc/salt/template/nr_syslog_cfg.j2:
hostname {{ host.name }}
!
{%- for server in host.syslog %}
logging host {{ server }}
{%- endfor %}
SaltStack State file /etc/salt/states/nr_cfg_syslog_and_ntp_state.sls content:
# apply logging configuration using jinja2 template
configure_logging:
  nr.cfg:
    - filename: salt://templates/nr_syslog_cfg.j2
    - plugin: netmiko
# apply NTP servers configuration using inline commands
configure_ntp:
  nr.task:
    - plugin: nornir_netmiko.tasks.netmiko_send_config
    - config_commands: ["ntp server 7.7.7.7", "ntp server 7.7.7.8"]
# save configuration using netmiko_save_config task plugin
save_configuration:
  nr.task:
    - plugin: nornir_netmiko.tasks.netmiko_save_config
Run state.apply command to apply state to devices:
[root@localhost /]# salt nrp1 state.apply nr_cfg_syslog_and_ntp_state
nrp1:
----------
          ID: configure_logging
    Function: nr.cfg
      Result: True
     Comment:
     Started: 12:45:41.339857
    Duration: 2066.863 ms
     Changes:
              ----------
              IOL1:
                  ----------
                  netmiko_send_config:
                      ----------
                      changed:
                          True
                      diff:
                      exception:
                          None
                      failed:
                          False
                      result:
                          configure terminal
                          Enter configuration commands, one per line.  End with CNTL/Z.
                          IOL1(config)#hostname IOL1
                          IOL1(config)#!
                          IOL1(config)#logging host 1.1.1.1
                          IOL1(config)#logging host 2.2.2.2
                          IOL1(config)#end
              IOL2:
                  ----------
                  netmiko_send_config:
                      ----------
                      changed:
                          True
                      diff:
                      exception:
                          None
                      failed:
                          False
                      result:
                          configure terminal
                          Enter configuration commands, one per line.  End with CNTL/Z.
                          IOL2(config)#hostname IOL2
                          IOL2(config)#!
                          IOL2(config)#logging host 1.1.1.2
                          IOL2(config)#logging host 2.2.2.1
                          IOL2(config)#end
                          IOL2#
----------
          ID: configure_ntp
    Function: nr.task
      Result: True
     Comment:
     Started: 12:45:43.407745
    Duration: 717.144 ms
     Changes:
              ----------
              IOL1:
                  ----------
                  nornir_netmiko.tasks.netmiko_send_config:
                      IOL1#configure terminal
                      IOL1(config)#ntp server 7.7.7.7
                      IOL1(config)#ntp server 7.7.7.8
                      IOL1(config)#end
              IOL2:
                  ----------
                  nornir_netmiko.tasks.netmiko_send_config:
                      configure terminal
                      Enter configuration commands, one per line.  End with CNTL/Z.
                      IOL2(config)#ntp server 7.7.7.7
                      IOL2(config)#ntp server 7.7.7.8
                      IOL2(config)#end
                      IOL2#
----------
          ID: save_configuration
    Function: nr.task
      Result: True
     Comment:
