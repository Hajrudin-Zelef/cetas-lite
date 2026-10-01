---
id: collect-261001-cisco/cisco/github-krishna426426-cisco-ios-xe-programmability-lab-module-8-ansible-2
title: "sudo pip install ansible"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/github-krishna426426-cisco-ios-xe-programmability-lab-module-8-ansible.md
source_anchor: ""
source_lines: [203, 421]
sha256: 8cfdc82bd3609f2dcf837cc9c4342bd0ceb706aa4da9d41b1d3e0cf591747cd7
---

# sudo pip install ansible

- All the YAML files start with the triple dash characters.
- **name** : name is assigned to the playbook “configure vrfs…”
- **hosts** : define the group of devices, remember**ios-xe** group
include both the cat3k and cat9k. By default, the Playbook executes
against all the devices in the given group
- **gather_facts** : facts (HW and SW info) can be collected from the
managed device. In this playbook it is disabled to speed up the
playbook execution but you can turn it on replacing**no** with**yes**
- **tasks** : there is a single task in this example and it’s named
“configure vfrs and purge all others”
- **ios_vrf** : is the Ansible module provided to manage VRFs on a
Cisco device running IOS
- **vrfs** : is the**ios_vrf** module option to provide the list of
VRFs to be configured, in this case the VRFs are**red** ,**blue** and**yellow**
- **purge:** is the optional**ios_vrf** module parameter. If set to**yes** , all VRFs other than the one listed above are removed from
the device

By default, Ansible will use the inventory hosts file located in
**/etc/ansible/hosts**. However, a different hosts file can be specified
using the **-i** flag at runtime or defined in the **ansible.cfg** file.
We set the inventory to the local **./hosts** file in the ansible.cfg
file in previously.

Step 7.  Now run the playbook in the Ubuntu Server with the following command
and provide the password **Cisco123**:

```
auto@automation:~/python/ansible$ ansible-playbook vrf.yaml -u admin -k
SSH password: Cisco123
```
You should see the following:

Let’s quickly analyze the output:

- both the playbook and task name (PLAY and TASK) are shown during the playbook execution
- the execution was successful (OK=1) for all the devices
- the configuration of both devices has been changed (changed=1)

Step 8. Now connect to one of the devices using the links in the Windows desktop (c9300) and verify that the VRFs have been configured.

You should see the following:

```
c9300# sh run | include vrf def
vrf definition Mgmt-vrf
vrf definition blue
vrf definition red
vrf definition yellow
```
Ansible provides many more modules to configure other IOS XE features like VLANs, users, static routes, L2 and L3 interfaces and so on.

For IOS XE features not yet supported by Ansible or for configurations
based on a known list of configuration CLIs, Ansible provides a module
named **ios_config**.

Step 9.  On the Windows host, open **Sublime Text 3** from the Start menu,
create a new file and enter the following into the window:

```
---
- name: configure a set of IOS XE CLIs using the Ansible ios_config module 
  hosts: ios
  gather_facts: no
  tasks:
    - name: Configure ntp server
      ios_config:
        lines:
          - ntp server 171.68.38.65
          - ntp server 1.2.3.4
    - name: Configure acl
      ios_config:
        lines:
            - 10 permit ip host 1.1.1.1 any log
            - 20 permit ip host 2.2.2.2 any log
            - 30 permit ip host 3.3.3.3 any log
            - 40 permit ip host 4.4.4.4 any log
            - 50 permit ip host 5.5.5.5 any log
        parents: ip access-list extended AnsibleTest
        before: no ip access-list extended AnsibleTest
        match: exact
```
Save the file as **z:\ansible\config.yaml**

You should see the following:

The first task configures a given list of NTP servers and the second task an extended ACL.

Notice the single ACL entries are applied after having removed the ACL
with the option **before** and entered in the ACL config mode using the
option **parents**.

Step 10. Return to the Ubuntu server and execute the playbook:

```
auto@automation:~/python/ansible$ ansible-playbook config.yaml -u admin -k
SSH password:  Cisco123
```
Connect to one of the devices using the link in the Windows desktop and verify both the NTP and ACL configurations have been applied.

You should see the following output:

```
c9300# sh run | i ntp
ntp server 171.68.38.65
ntp server 1.2.3.4
c9300#
c9300# sh run | sec ip acc
ip access-list extended AnsibleTest
 permit ip host 1.1.1.1 any log
 permit ip host 2.2.2.2 any log
 permit ip host 3.3.3.3 any log
 permit ip host 4.4.4.4 any log
 permit ip host 5.5.5.5 any log
```
In the previous step we used the module for IOS XE configuration CLIs.

What about IOS XE exec CLIs? You need to use the **ios_command** module
instead.

Step 11.  On the Windows host, open **Sublime Text 3** from the Start menu,
create a new file and enter the following into the window:

```
---
- name: run commands on Cisco IOS XE devices
  hosts: ios
  gather_facts: no
  tasks:
    - name: show version and ip interfaces brief
      ios_command:
        commands:
            - show version
            - show ip interface brief
```
Save the file as z:\ansible\commands.yaml

You should see the following:

In the Playbook above we have only one task to execute two IOS XE exec commands.

Step 12.  Now run the playbook in the Ubuntu Server with the following command
and provide the password **Cisco123**:

```
auto@automation:~/python/ansible$ ansible-playbook commands.yaml -u admin -k
SSH password:  Cisco123
```
You should see the following:

The Playbook was executed successfully (OK=1) but…where is the CLI output???

Step 13.  Try again but this time add a **-v** option at the end.

```
auto@automation:~/python/ansible$ ansible-playbook commands.yaml -u admin -k -v
```
You should see the CLIs output like in the figure below.

If you check the output carefully, you’ll see that the output is stored
in two variables named **stdout** and **stdout_lines**. The first
variable stores the output in a single string and the single CLI output
lines are separated by \n, that is, Carriage Return (Enter) characters
while the second variable is a list of strings, each storing a single
CLI output line.

The **-v** option we added in the second playbook execution is the
Ansible option to provide a more verbose output, very useful to debug
Playbooks executions and failures. The more v's you provide, the more
verbosity you get. For example, adding **-vv** will provide more verbose
output than **-v**.

NETCONF connection is also available in Ansible. Using “netconf-config” module in ansible allows the user to send a configuration XML file to a networking device and detects if there was a configuration change.

Step 1.  To configure an interface description using NETCONF, on the Windows
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
  - name: set Management interface description
    netconf_config:
      xml: |
        <config>
          <interfaces xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces">
            <interface>
              <name>GigabitEthernet1/0/1</name>
              <description>Managed by Ansible using netconf connection</description>
            </interface>
          </interfaces>
        </config>        
```
Save the file as **z:\ansible\netconf-description.yaml**

You should see the following:

- In this playbook we are using “**vars** ” to define
ansible_connection and ansible_network_os to use locally instead
of changing in the ansible.cfg.
- **netconf_config** is the module in ansible to configure on the
network device using NETCONF connection

Step 2.  Now run the playbook in the Ubuntu Server with the following command
and provide the password **Cisco123**:

```
auto@automation:~/python/ansible$ ansible-playbook netconf-description.yaml -u admin -k
SSH password:  Cisco123
```
Connect to c9300 in the Windows desktop and verify interface description on gi1/0/1 have been applied.

You should see the following output:

