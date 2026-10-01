---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/o-x-l-ansible-opnsense-blob-head-docs-source-usage-1-install-rst-2e78ffbc
title: "stable version:"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/o-x-l-ansible-opnsense-blob-head-docs-source-usage-1-install-rst-2e78ffbc.md
source_anchor: ""
source_lines: [1, 15]
sha256: 0157dfc3da3447df2c70da49193f25b219bcbf448b24adf7b3c3f9f3da300c88
---

# stable version:

See the documentation on how to install Ansible.

If you DO NOT want to use Ansible - this fork provides you with a raw Python3 interface.

The httpx python module is used for API communications!

`python3 -m pip install --upgrade httpx````
# stable version:
ansible-galaxy collection install oxlorg.opnsense
# latest version:
ansible-galaxy collection install git+https://github.com/O-X-L/ansible-opnsense.git
# install to specific directory for easier development
cd $PLAYBOOK_DIR
ansible-galaxy collection install git+https://github.com/O-X-L/ansible-opnsense.git -p ./collections
```
