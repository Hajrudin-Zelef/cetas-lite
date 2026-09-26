---
id: collect-260926-mikrotik/mikrotik/questions-1021200-routeros-shell-command-result-with-ansible-is-different-ca0215f4
title: "questions-1021200-routeros-shell-command-result-with-ansible-is-different-ca0215f4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-1021200-routeros-shell-command-result-with-ansible-is-different-ca0215f4.md
source_anchor: ""
source_lines: [1, 17]
sha256: acced497d4771ee78bb6ca2c8d27d20898265b14e18bdd06a1dd986b06706179
---

# questions-1021200-routeros-shell-command-result-with-ansible-is-different-ca0215f4

You will see, I have a little problem with the execution of an Ansible task, since I am giving it the same command that I work with in the Mikrotik shell, but the result is a little different and I just don't understand why.
You will see, I am trying to write a few concrete subnets to make permissive passage with SSH from those subnets, a "Available From" of all life, to allow connection only from the subnets that I indicate.
The execution of the command from the Mikrotik shell is as follows:
/ip service set ssh address=10.0.0.0/8,172.16.0.0/12,19
2.168.0.0/16
And the result is the following:
- hosts: Mikrotik-Routers_TestSenseLlum
  remote_user: admin
  gather_facts: no
  connection: network_cli
  tasks:
    - name: Allow ssh only inside the private networks
      routeros_command:
        commands: /ip service set ssh address=10.0.0.0/8,172.16.0.0/12,192.168.0.0/16
The result is wrong, only enter the first subnet, 10, but the others, do not enter them.
I don't know exactly what the problem is, has something similar happened to someone?
commands: "/ip service set ssh address=10.0.0.0/8,172.16.0.0/12,192.168.0.0/16"or evencommands: /ip service set ssh address="10.0.0.0/8,172.16.0.0/12,192.168.0.0/16"and even eventuallycommands: '/ip service set ssh address="10.0.0.0/8,172.16.0.0/12,192.168.0.0/16"'. I suspect your comma seperated ips are interpreted as an old style list asignation to anaddressoption in the task...
