---
id: collect-261001-cisco/cisco/ansible-cisco-ios-command-example-network-automation-blog
title: "ansible-cisco-ios-command-example-network-automation-blog"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/ansible-cisco-ios-command-example-network-automation-blog.md
source_anchor: ""
source_lines: [1, 88]
sha256: a37a3e4988e59bfc5e1f8e2cf79d50a10ee86115b2ed9a3992cd94ee6f266a99
---

# ansible-cisco-ios-command-example-network-automation-blog

In this Ansible IOS_Command Example post I will share a simple Ansible Playbook example that will write an ACL to a Cisco Switch.

This playbook will work equally well on any Cisco device and once you grasp the concepts of using the IOS_Command Module and also the IOS_Config module you can configure whatever you want.

You will soon see that Ansible Network Automation skills are very good to know to make your life as a network engineer a lot easier in the new world of network automation.

The IOS Command module is documented here – http://docs.ansible.com/ansible/latest/ios_command_module.html

There are the two parameters that I will highlight and I am using in my example playbook

- **authorize** – yes / no – This parameter enters privileged mode on the Cisco device if code says yes
- **waitfor** – List of conditions to wait for before progressing with the Playbook

## Ansible IOS_Config Module

The IOS Config module is documented here – http://docs.ansible.com/ansible/latest/ios_config_module.html

### Ansible Playbook

So let’s get into the playbook

```
--
- hosts: 3560-SW1
  gather_facts: false
  connection: local
  tasks:
   - name: Configure ACL on Cisco Switch
     ios_config:
       authorize: yes
       lines:
         - access-list 99 permit 172.16.1.100
   - name: Verify ACL is present
     ios_command:
       commands:
         - sh access-l
       waitfor:
         - result[0] contains 'permit 172.16.1.100'
   - name: SAVE CONFIG
     ios_config:
       authorize: yes
       save: yes
```
This is a very simple play that is very self explanatory just by reading the Jinja2 code

It will configure a line in an access list using the IOS_Config module it then verifies if that line is there using the ios_command module and finally saves the config using the IOS_Config module again.

**Note:** This playbook was written in Ansible 2.3 – the latest version 2.4 has another parameter in the IOS_Config module called **save_when**

This can be set to always / never / modified.

So the config can be saved only when it was modified.

The playbook is available on my GitHub – https://github.com/rogerperkin/playbooks

If you are serious about learning Network Automation check out my course below.

**Read my other Ansible posts:** 

Fabian martinez

Hi Roger,

Does your course cover setting up ansible on a mac? I recently installed ansible on my mac and would love some help on getting this going.

Thanks

Roger Perkin

Fabian,

My course covers using Ansible for Network Automation, it does not matter if you are running it from a Mac or a Linux host all the playbooks and operation and folder structures will be the same.

All you need to operate Ansible is to be able to SSH to your target devices.

HTH

kiborg

I made a script for backup, but i am having trouble with ssh to cisco switch.

I’ve tried many different option but no success. I am able to ssh from my ansible server and I have key stored in .ssh/known_hosts but my log says “connecting to host returned an error”

“Bad authentication type”…” Any idea how to resolve this?

Roger Perkin

Can you post your playbook, and also run the playbook with -vvv at the end and post that as well
