---
id: collect-260926-mikrotik/mikrotik/community-routeros
title: "Community.Routeros"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/community-routeros.md
source_anchor: ""
source_lines: [1, 45]
sha256: f33c818f1f5e6e283a8ab3ca685e7d5c14b8fb31ed3769c903f14360c1bc50d8
---

# Community.Routeros

Collection version 3.21.0

Modules and plugins for MikroTik RouterOS

**Authors:**

- Egor Zaitsev (github.com/heuels)
- Nikolay Dachev (github.com/NikolayDachev)
- Felix Fontein (github.com/felixfontein)

**Supported ansible-core versions:**

- 2.15.0 or newer

- Forum: Ansible Forum: General usage and support questions.
- Forum: Ansible Forum: Discussions about RouterOS.
- Matrix room `#users:ansible.im` : General usage and support questions.
- IRC channel `#ansible` (Libera network):
General usage and support questions.

These are the plugins in the community.routeros collection:

### Modules

- api module – Ansible module for RouterOS API
- api_facts module – Collect facts from remote devices running MikroTik RouterOS using the API
- api_find_and_modify module – Find and modify information using the API
- api_info module – Retrieve information from API
- api_modify module – Modify data at paths with API
- command module – Run commands on remote devices running MikroTik RouterOS
- facts module – Collect facts from remote devices running MikroTik RouterOS

### Cliconf Plugins

- routeros cliconf – Use routeros cliconf to run command on MikroTik RouterOS platform

### Filter Plugins

- join filter – Join a list of arguments to a command
- list_to_dict filter – Convert a list of arguments to a dictionary
- quote_argument filter – Quote an argument
- quote_argument_value filter – Quote an argument value
- split filter – Split a command into arguments
