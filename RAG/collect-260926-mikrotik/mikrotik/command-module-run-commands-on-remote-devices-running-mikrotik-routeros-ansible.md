---
id: collect-260926-mikrotik/mikrotik/command-module-run-commands-on-remote-devices-running-mikrotik-routeros-ansible
title: "command-module-run-commands-on-remote-devices-running-mikrotik-routeros-ansible"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/command-module-run-commands-on-remote-devices-running-mikrotik-routeros-ansible.md
source_anchor: ""
source_lines: [1, 55]
sha256: c0fcb302fb33d9e3a97e63cc395252a82ac9c8c52b7dbdf2e9106e2ef3266489
---

# command-module-run-commands-on-remote-devices-running-mikrotik-routeros-ansible

Note

This module is part of the community.routeros collection (version 3.9.0).

You might already have this collection installed if you are using the `ansible` package. It is not included in `ansible-core`. To check whether it is installed, run `ansible-galaxy collection list`.

To install it, use: `ansible-galaxy collection install community.routeros`.

To use it in a playbook, specify: `community.routeros.command`.

| Parameter | Comments | 
|---|---|
| **commands** list / elements=string / required | List of commands to send to the remote RouterOS device over the configured provider. The resulting output from the command is returned. If the argument is provided, the module is not returned until the condition is satisfied or the number of retries has expired.**wait_for** | 
| **interval** integer | Configures the interval in seconds to wait between retries of the command. If the command does not pass the specified conditions, the interval indicates how long to wait before trying the command again. **Default:**`1` | 
| **match** string | The argument is used in conjunction with the**match** argument to specify the match policy. Valid values are**wait_for**`all` or`any` . If the value is set to`all` then all conditionals in the wait_for must be satisfied. If the value is set to`any` then only one of the values must be satisfied. **Choices:**  | 
| **retries** integer | Specifies the number of retries a command should by tried before it is considered failed. The command is run on the target device every retry and evaluated against the conditions.**wait_for** **Default:**`10` | 
| **wait_for** list / elements=string | List of conditions to evaluate against the output of the command. The task will wait for each condition to be true before moving forward. If the conditional is not true within the configured number of retries, the task fails. See examples. | 

| Attribute | Support | Description | 
|---|---|---|
| **check_mode** | **Support:** **none** Before community.routeros 3.0.0, the module claimed to support check mode. It simply executed the command in check mode. | Can run in `check_mode` and return changed status prediction without modifying target. | 
| **diff_mode** | **Support:** **none** | Will return details on what has changed (or possibly needs changing in `check_mode` ), when in diff mode. | 
| **idempotent** | **Support:** N/A Whether the executed command is idempotent depends on the command. | When run twice in a row outside check mode, with the same arguments, the second invocation indicates no change. This assumes that the system controlled/queried by the module has not changed in a relevant way. | 
| **platform** | **Platform:****RouterOS** | Target OS/families that can be operated against. | 

```
---
- name: Run command on remote devices
  community.routeros.command:
    commands: /system routerboard print
- name: Run command and check to see if output contains routeros
  community.routeros.command:
    commands: /system resource print
    wait_for: result[0] contains MikroTik
- name: Run multiple commands on remote nodes
  community.routeros.command:
    commands:
      - /system routerboard print
      - /system identity print
- name: Run multiple commands and evaluate the output
  community.routeros.command:
    commands:
      - /system routerboard print
      - /interface ethernet print
    wait_for:
      - result[0] contains x86
      - result[1] contains ether1
```
 Common return values are documented here, the following are the fields unique to this module:

| Key | Description | 
|---|---|
| **failed_conditions** list / elements=string | The list of conditionals that have failed. **Returned:** failed **Sample:**`["...", "..."]` | 
| **stdout** list / elements=string | The set of responses from the commands. **Returned:** always apart from low level errors (such as action plugin) **Sample:**`["...", "..."]` | 
| **stdout_lines** list / elements=string | The value of stdout split into a list. **Returned:** always apart from low level errors (such as action plugin) **Sample:**`[["...", "..."], ["..."], ["..."]]` |
