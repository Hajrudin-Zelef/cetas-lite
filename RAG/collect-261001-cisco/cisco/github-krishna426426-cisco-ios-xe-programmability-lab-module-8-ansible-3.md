---
id: collect-261001-cisco/cisco/github-krishna426426-cisco-ios-xe-programmability-lab-module-8-ansible-3
title: "sudo pip install ansible"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/github-krishna426426-cisco-ios-xe-programmability-lab-module-8-ansible.md
source_anchor: ""
source_lines: [422, 544]
sha256: 073527309670fc482b49e3d536ed9b916278d3de7b53f63ba6f065625f0b15c7
---

# sudo pip install ansible

Step 3.  To configure telemetry subscriptions using NETCONF, on the Windows
host, open **Sublime Text** from the Start menu, create a new file
and enter the following into the window:

```
---
- name: test Ansible connection netconf on Cisco IOS XE
  hosts: c9300
  vars:
      ansible_connection: netconf
      ansible_network_os: default
  gather_facts: no
  tasks:
  - name: establish subscription
    netconf_config:
      xml: |
          <config>
            <mdt-config-data xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-mdt-cfg">
              <mdt-subscription>
                <subscription-id>501</subscription-id>
                <base>
                  <stream>yang-push</stream>
                  <encoding>encode-kvgpb</encoding>
                  <source-address>10.1.1.5</source-address>
                  <period>1000</period>
                  <xpath>/process-cpu-ios-xe-oper:cpu-usage/cpu-utilization/five-seconds</xpath>
                </base>
                <mdt-receivers>
                  <address>10.1.1.3</address>
                  <port>57500</port>
                  <protocol>grpc-tcp</protocol>
                </mdt-receivers>
              </mdt-subscription>
              <mdt-subscription>
                <subscription-id>502</subscription-id>
                <base>
                  <stream>yang-push</stream>
                  <encoding>encode-kvgpb</encoding>
                  <source-address>10.1.1.5</source-address>
                  <period>1000</period>
                  <xpath>/if:interfaces-state/interface[name="GigabitEthernet1/0/24"]/statistics</xpath>
                </base>
                <mdt-receivers>
                  <address>10.1.1.3</address>
                  <port>57500</port>
                  <protocol>grpc-tcp</protocol>
                </mdt-receivers>
              </mdt-subscription>
            </mdt-config-data>
          </config>
```
Save the file as **z:\ansible\Telemetry.yaml**

You should see the following:

Step 4.  Now run the playbook in the Ubuntu Server with the following command
and provide the password **Cisco123**:

```
auto@automation:~/python/ansible$ ansible-playbook Telemetry.yaml -u admin -k
SSH password:  Cisco123
```
Connect to c9300 in the Windows desktop and verify interface description on gi1/0/1 have been applied.

You should see the following output:

Step 1.  To upgrade a code on IOS XE switch and copy the code from a ubuntu VM to switch, on the Windows host, open **Sublime Text** from the Start menu, create a new file and enter the following into the window:

```
---
- name: upgrade on Cisco IOS XE switches
  hosts: c9300
  gather_facts: no
  connection: network_cli
################################################################################ #
# Step 1: Define necessary Variables and use these variable on below tasks.
# ################################################################################
  vars:
    SERVER_IP: 10.1.1.3
    SERVER_IP_USER: username
    SERVER_IP_PASS: password
    IOS_PATH: /var/www/html/
    IOS_FILE: cat9k_iosxe.17.02.01.SPA.bin
    BOOT_LOC: flash
################################################################################ #
# Step 2: Define task to copy the code from ubuntu VM to 9300
# ################################################################################
  tasks:
  - name: COPY THE CODE FROM UBUNTU VM TO 9300
    ios_command:
      commands:
        - command: 'copy http://{{ SERVER_IP}}/{{ IOS_FILE }} {{ BOOT_LOC }}:{{ IOS_FILE }}'
          check_all: False
          prompt:
            - "Destination filename [{{ IOS_FILE }}]?"
            - "Do you want to over write? [confirm]"
          answer: 
            - "\r"
            - "n"
    vars:
      ansible_command_timeout: 600
    tags:
      - COPY_CODE
################################################################################ #
# Step 3: Define task to install the code and then reload
# ################################################################################
  - name: INSTALL THE CODE AND THEN RELOAD THE SWITCH
    ios_command:
      commands:
        - wr mem
        - install add file {{ BOOT_LOC }}:{{ IOS_FILE }} activate commit prompt-level none
    vars:
      ansible_command_timeout: 1200
    tags:
      - REBOOT_SWITCH
```
Save the file as **z:\ansible\switch_upgrade.yaml**

You should see the following:

Now run the playbook on the ubuntu server and the whole process will take around 20 minutes for copying the file from ubuntu to the switch and then execute upgrade on the switch.

This module has shown how to start with ansible and create different playbooks for configuring Access lists, VRFs and how to use NETCONF module to configure the switch using XML code. Also shown a playbook for executing an upgrade on the switch.
