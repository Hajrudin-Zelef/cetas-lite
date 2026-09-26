---
id: collect-260926-mikrotik/mikrotik/community-routeros-readme-md-at-3c2f48bfde5067283c6ec6072c919ea4bc608753-ansible-collectio
title: "community-routeros-readme-md-at-3c2f48bfde5067283c6ec6072c919ea4bc608753-ansible-collections-communi"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "licenses"]
source: docs/RAG/lot-mikrotik/RouterOS/community-routeros-readme-md-at-3c2f48bfde5067283c6ec6072c919ea4bc608753-ansible-collections-communi.md
source_anchor: ""
source_lines: [1, 152]
sha256: 68ef0b62c5b12c1ef35cfbb16847784c00e0791d9eec6989e2d3d3b2ee22949d
---

# community-routeros-readme-md-at-3c2f48bfde5067283c6ec6072c919ea4bc608753-ansible-collections-communi

Provides modules for Ansible to manage MikroTik RouterOS instances.

You can find documentation for the modules and plugins in this collection here.

We follow Ansible Code of Conduct in all our interactions within this project.

If you encounter abusive behavior violating the Ansible Code of Conduct, please refer to the policy violations section of the Code of Conduct for information on how to raise a complaint.

- 
Join the Ansible forum: 
  - Get Help: get help or help others.Please add appropriate tags if you start new discussions, for example the `routeros` tag.
  - Posts tagged with 'routeros': subscribe to participate in RouterOS related conversations.
  - Social Spaces: gather and interact with fellow enthusiasts.
  - News & Announcements: track project-wide announcements including social events.
- Get Help: get help or help others.Please add appropriate tags if you start new discussions, for example the 
- 
The Ansible Bullhorn newsletter: used to announce releases and important changes.

For more information about communication, see the Ansible communication guide.

Tested with the current ansible-core 2.15, ansible-core 2.16, ansible-core 2.17, ansible-core 2.18, ansible-core 2.19, ansible-core 2.20, ansible-core 2.21, ansible-core 2.22 releases, and the current development version of ansible-core. Ansible 2.9, ansible-base 2.10, and ansible-core versions before 2.15.0 are not supported.

The exact requirements for every module are listed in the module documentation.

The collection supports the `network_cli` connection.

Please note that `community.routeros.api` module does **not** support Windows jump hosts!

Browsing the **latest** collection documentation will show docs for the *latest version released in the Ansible package*, not the latest version of the collection released on Galaxy.

Browsing the **devel** collection documentation shows docs for the *latest version released on Galaxy*.

We also separately publish **latest commit** collection documentation which shows docs for the *latest commit in the `main` branch*.

If you use the Ansible package and do not update collections independently, use **latest**. If you install or update this collection directly from Galaxy, use **devel**. If you are looking to contribute, use **latest commit**.

- `community.routeros.api`
- `community.routeros.api_facts`
- `community.routeros.api_find_and_modify`
- `community.routeros.api_info`
- `community.routeros.api_modify`
- `community.routeros.command`
- `community.routeros.facts`

You can find documentation for the modules and plugins in this collection here.

See Ansible Using collections for general detail on using collections.

There are two approaches for using this collection. The `command` and `facts` modules use the `network_cli` connection and connect with SSH. The `api` module connects with the HTTP/HTTPS API.

The terminal-based modules in this collection (`community.routeros.command` and `community.routeros.facts`) do not support arbitrary symbols in router's identity. If you are having trouble connecting to your device, please make sure that your MikroTik's identity contains only alphanumeric characters and dashes. Also, the `community.routeros.command` module does not support nesting commands and expects every command to start with a forward slash (`/`). Running the following command will produce an error.

```
- community.routeros.command:
    commands:
      - /ip
      - print
```
Example inventory `hosts` file:

```
[routers]
router ansible_host=192.168.1.1
[routers:vars]
ansible_connection=ansible.netcommon.network_cli
ansible_network_os=community.routeros.routeros
ansible_user=admin
ansible_ssh_pass=test1234
```
Example playbook:

```
---
- name: RouterOS test with network_cli connection
  hosts: routers
  gather_facts: false
  tasks:
    - name: Run a command
      community.routeros.command:
        commands:
          - /system resource print
      register: system_resource_print
    - name: Print its output
      ansible.builtin.debug:
        var: system_resource_print.stdout_lines
    - name: Retrieve facts
      community.routeros.facts:
    - ansible.builtin.debug:
        msg: "First IP address: {{ ansible_net_all_ipv4_addresses[0] }}"
```
Example playbook:

```
---
- name: RouterOS test with API
  hosts: localhost
  gather_facts: false
  vars:
    hostname: 192.168.1.1
    username: admin
    password: test1234
  module_defaults:
    group/community.routeros.api:
      hostname: "{{ hostname }}"
      password: "{{ password }}"
      username: "{{ username }}"
      tls: true
      force_no_cert: false
      validate_certs: true
      validate_cert_hostname: true
      ca_path: /path/to/ca-certificate.pem
  tasks:
    - name: Get "ip address print"
      community.routeros.api:
        path: ip address
      register: print_path
    - name: Print the result
      ansible.builtin.debug:
        var: print_path.msg
    - name: Change IP address to 192.168.1.1 for interface bridge
      community.routeros.api_find_and_modify:
        path: ip address
        find:
          interface: bridge
        values:
          address: "192.168.1.1/24"
    - name: Retrieve facts
      community.routeros.api_facts:
    - ansible.builtin.debug:
        msg: "First IP address: {{ ansible_net_all_ipv4_addresses[0] }}"
```
We're following the general Ansible contributor guidelines; see Ansible Community Guide.

If you want to clone this repositority (or a fork of it) to improve it, you can proceed as follows:

1. Create a directory `ansible_collections/community` ;
2. In there, checkout this repository (or a fork) as `routeros` ;
3. Add the directory containing `ansible_collections` to your ANSIBLE_COLLECTIONS_PATH.

See Ansible's dev guide for more information.

See the collection's changelog.

We plan to regularly release minor and patch versions, whenever new features are added or bugs fixed. Our collection follows semantic versioning, so breaking changes will only happen in major releases.

This collection is primarily licensed and distributed as a whole under the GNU General Public License v3.0 or later.

See LICENSES/GPL-3.0-or-later.txt for the full text.

Parts of the collection are licensed under the BSD 2-Clause license.

All files have a machine readable `SPDX-License-Identifier:` comment denoting its respective license(s) or an equivalent entry in an accompanying `.license` file. Only changelog fragments (which will not be part of a release) are covered by a blanket statement in `REUSE.toml`. This conforms to the REUSE specification.
