---
id: collect-261001-huawei/huawei/ansible-2-10-collections-community-network-ce-command-module-html-ee0f6d86
title: "Note: examples below use the following provider dict to handle"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/ansible-2-10-collections-community-network-ce-command-module-html-ee0f6d86.md
source_anchor: ""
source_lines: [1, 65]
sha256: e063b661b46cc0a1723868b83a542865aa247f4b9889783b5c894f4e58101053
---

# Note: examples below use the following provider dict to handle

community.network.ce_command – Run arbitrary command on HUAWEI CloudEngine devices.¶
Note
This plugin is part of the community.network collection (version 1.3.2).
To install it use: ansible-galaxy collection install community.network.
To use it in a playbook, specify: community.network.ce_command.
Synopsis¶
- Sends an arbitrary command to an HUAWEI CloudEngine node and returns the results read from the device. The ce_command module includes an argument that will cause the module to wait for a specific condition before returning or timing out if the condition is not met.
Parameters¶
| Parameter | Choices/Defaults | Comments | 
|---|---|---|
| commands                      string                                              / required                     |  | The commands to send to the remote HUAWEI CloudEngine device over the configured provider.  The resulting output from the command is returned. If the wait_for argument is provided, the module is not returned until the condition is satisfied or the number of retries has been exceeded. | 
| interval                      string                                                                  | Default: 1 | Configures the interval in seconds to wait between retries of the command.  If the command does not pass the specified conditional, the interval indicates how to long to wait before trying the command again. | 
| match                      string                                                                  | Default: "all" | The match argument is used in conjunction with the wait_for argument to specify the match policy.  Valid values are all orany .  If the value is set toall then all conditionals in the wait_for must be satisfied.  If the value is set toany then only one of the values must be satisfied. | 
| retries                      string                                                                  | Default: 10 | Specifies the number of retries a command should by tried before it is considered failed.  The command is run on the target device every retry and evaluated against the wait_for conditionals. | 
| wait_for                      string                                                                  |  | Specifies what to evaluate from the output of the command and what conditionals to apply.  This argument will cause the task to wait for a particular conditional to be true before moving forward.   If the conditional is not true by the configured retries, the task fails.  See examples. | 
Notes¶
Note
- Recommended connection is network_cli .
- This module also works with local connections for legacy playbooks.
Examples¶
# Note: examples below use the following provider dict to handle
#       transport and authentication to the node.
- name: CloudEngine command test
  hosts: cloudengine
  connection: local
  gather_facts: no
  vars:
    cli:
      host: "{{ inventory_hostname }}"
      port: "{{ ansible_ssh_port }}"
      username: "{{ username }}"
      password: "{{ password }}"
      transport: cli
  tasks:
  - name: "Run display version on remote devices"
    community.network.ce_command:
      commands: display version
      provider: "{{ cli }}"
  - name: "Run display version and check to see if output contains HUAWEI"
    community.network.ce_command:
      commands: display version
      wait_for: result[0] contains HUAWEI
      provider: "{{ cli }}"
  - name: "Run multiple commands on remote nodes"
    community.network.ce_command:
      commands:
        - display version
        - display device
      provider: "{{ cli }}"
  - name: "Run multiple commands and evaluate the output"
    community.network.ce_command:
      commands:
        - display version
        - display device
      wait_for:
        - result[0] contains HUAWEI
        - result[1] contains Device
      provider: "{{ cli }}"
Return Values¶
Common return values are documented here, the following are the fields unique to this module:
| Key | Returned | Description | 
|---|---|---|
| failed_conditions                    list                    / elements=string                     | failed | the conditionals that failed Sample: ['...', '...'] | 
| stdout                    list                    / elements=string                     | always | the set of responses from the commands Sample: ['...', '...'] | 
| stdout_lines                    list                    / elements=string                     | always | The value of stdout split into a list Sample: [['...', '...'], ['...'], ['...']] |
