---
id: collect-261001-cisco/cisco/github-krishna426426-cisco-ios-xe-programmability-lab-module-8-ansible-1
title: "sudo pip install ansible"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention", "open source", "parameters"]
source: docs/RAG/collect-261001-cisco/github-krishna426426-cisco-ios-xe-programmability-lab-module-8-ansible.md
source_anchor: ""
source_lines: [1, 202]
sha256: 70ba2827f1b4316c18f47449ad34efc632c55245ead643c9455371d59022c744
---

# sudo pip install ansible

Ansible

Server Installation

Environment setup

Create an inventory file

Test device connectivity

Ansible Documentation

Ansible Playbooks

Build an IOS XE CLIs Playbook

Build a Playbook for IOS XE commands

Using NETCONF in Ansible

Build a Playbook for upgrade on IOS XE switches

Conclusion

Ansible is an open source configuration management platform.

Configuration Management is the practice of automating software provisioning, configuration management, and application deployment.

This practice has been widely used across IT systems management by many organizations large and small for over a decade. Having been proven successful with servers and applications, in recent years, it has extended to the network as well.

With these tools you can define and enforce configuration related to system level operations (e.g., authentication, logging, image), interface level configuration (e.g., VLAN, QoS, Security), routing configurations (e.g., OSPF or BGP), and much more.

Ansible is based on an "agent-less" architecture, so there is no need to install anything onto the Cisco IOS XE device to get started.

Ansible Core modules for Cisco IOS and IOS XE devices have been available since release 2.3 so there is no need for anything outside of a base Ansible installation.

Both CLI-based and NETCONF-based configuration are supported. The Ansible NETCONF implementation changed in the 2.5 release, leveraging the newly introduced NETCONF connection instead of the usual local connection.

In this lab we are going to use Ansible 2.5.4 release.

The picture below describes the main Ansible terminology:

Ansible provides off-the-shelf modules to manage a variety of servers, application and network devices. In this lab we are going to mainly use the Ansible IOS modules.

Ansible can be installed on any computer running the most popular Linux distributions like Red Hat, CentOS, Fedora, Debian, or Ubuntu, using either the OS package manager or via the Python package manager (pip).

For instance, you can install latest Ansible on an Ubuntu server with:

```
# sudo pip install ansible
```
In this lab there is no Internet connectivity and Ansible has been already pre-installed in the Ubuntu Server.

All the Cisco IOS XE modules are included in Ansible Core so no additional effort is required to begin automating your Cisco IOS XE devices.

Step 1.  Open the Ubuntu server PuTTY session by clicking on **Ubuntu** on
the desktop. You can verify if Ansible is installed in our lab and
related version with:

Note: If you are using an existing PuTTY window, be sure to return to
the home directory by typing **cd** into the Ubuntu server.

```
auto@automation:~/ ansible --version
```
You should see the following:

Create an Ansible configuration file

The Ansible configuration file stores the default configs used by all Playbooks.

Step 2.  On the Windows host, open **Sublime Text 3** from the Start menu,
create a new file and enter the following into the window:

```
[defaults]
inventory = ./hosts
host_key_checking = False
roles_path = ./
remote_user = admin
deprecation_warnings=False
```
Save the file as **z:\ansible\ansible.cfg**

These ensure that Ansible:

- Uses the inventory **hosts** file in the local directory
- Disables host checking to automatically add hosts to
**known_hosts** file
- Sets the roles path to the local directory
- Set the default user for device connections to **admin**
- Turn off any **deprecation warnings**

The inventory file is where the devices under management are listed. Devices can be grouped, and a single device can be included in multiple groups.

Step 3.  On the Windows host, open **Sublime Text 3** from the Start menu,
create a new file and enter the following into the window:

```
[csr1000v]
10.1.1.2
[c9300]
10.1.1.5
[c9800]
10.1.1.6
[ios:children]
csr1000v
c9300
c9800
[ios:vars]
ansible_connection=network_cli
ansible_network_os=ios
[ubuntu]
10.1.1.3
```
Save the file as **z:\ansible\hosts**

In the above hosts file, we have three groups called **csr100v**,
**c9300** and **c9800**, with one device each. Groups can be nested like
the **ios-xe** group which includes all three groups.

The default device connection is set to **network_cli** (that is CLIs
over SSH), and Operating System type to ios for all the device in the
**ios-xe** group, that is, for all the devices in this lab.

Defaults can be overridden in playbooks and at the command line.

We’ll see some override examples later in the lab.

To make sure the Ansible server can reach all the devices, go to the Ubuntu server, move to the ansible directory:

```
auto@automation:~/ cd /home/auto/python/ansible
```
and use the following Ansible command and provide the device SSH
password **admin**

```
auto@automation:~/python/ansible\$ ansible ios -m ping -u admin -k
SSH password: Cisco123
```
Where:

- **ios-xe** is the device group
- **ping** is the Ansible module to test device connectivity
- **-u** is the option to provide the device SSH username
- **-k** is the option to provide the device password at run time

You should see the following:

Thorough documentation for all Cisco IOS XE modules can be found on the Ansible website (http://docs.ansible.com/ansible/latest/modules/ios_vrf_module.html) or alternatively from the terminal, by utilizing the inbuilt documentation tool.

1. Test the documentation tool with the **ios_config** module:

```
auto@automation:~/python/ansible$ ansible-doc ios_config
```
You should see the following:

The documentation tool provides a description of the given module, all the mandatory and optional parameters, as well as some useful examples to start playing with the module.

Exit out of the documentation screen when finished by pressing "q".

An Ansible Playbook is a repeatable standard config. Playbooks are written in YAML, a common encoding format, which is very easy to read and uses indentation for creating the script hierarchy. Pay extra attention to the indentation!!!

In the next steps, you will build and execute several Ansible Playbooks

Step 5. Write a Playbook to configure VRFs

In this initial playbook we will provision a number of VRFs across all
devices and remove (**purge**) any other VRF configured on the device.
We will use the Ansible module called **ios_vrf** to automate this
task.

On the Windows host, open **Sublime Text 3** from the Start menu, create
a new file and enter the following into the window:

```
---
- name: configure vrfs and remove any other vrf configured
  hosts: ios
  gather_facts: no
  tasks:
    - name: configure vfrs and purge all others
      ios_vrf:
        vrfs:
            - red
            - blue
            - yellow
        purge: yes
```
**Note:** Please note carefully that the file begins with three hyphens!
(---)

Save the file as **z:\ansible\vrf.yaml**

Sublime will autodetect the YAML format and display the file like in the picture below:

If the colors do not match those shown in the screen shot, most likely something is wrong with the file indentation. Double check before proceeding to the next step.

Step 6. Let’s analyze the playbook top to bottom:

