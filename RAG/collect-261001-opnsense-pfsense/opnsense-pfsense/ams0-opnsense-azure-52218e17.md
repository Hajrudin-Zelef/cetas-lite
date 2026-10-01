---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/ams0-opnsense-azure-52218e17
title: "ams0-opnsense-azure-52218e17"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/ams0-opnsense-azure-52218e17.md
source_anchor: ""
source_lines: [1, 25]
sha256: d2c81d2721e4d35381d38f0e862d29ba4019e17e09d90a18a7dcc7b831cf8b5d
---

# ams0-opnsense-azure-52218e17

This repository contains an Ansible setup to manage an OPNsense firewall using the `oxlorg.opnsense` collection.

1. **Ansible** : Ensure Ansible is installed on your machine.
2. **Python** : Python 3 is required.

1. 
**Install Dependencies** :
Install the required Ansible collection:ansible-galaxy collection install -r requirements.yml
2. 
**Configure Inventory** :
Edit`inventory/hosts.yml` and update the`ansible_host` with your OPNsense firewall's IP address or hostname.
3. 
**Configure Credentials** :
Edit`playbook.yml` and update the`opnsense_api_key` and`opnsense_api_secret` variables.*Note: For production, it is highly recommended to use Ansible Vault to encrypt these secrets.*To generate API keys in OPNsense: 
  - Go to **System > Access > Users** .
  - Edit the user (e.g., `root` or a dedicated automation user).
  - Scroll down to **API keys** and click the**+** button to generate a new key/secret pair.
4. Go to 

Run the playbook:

`ansible-playbook playbook.yml`
The `playbook.yml` currently contains a task to create a firewall alias named `ANSIBLE_TEST_ALIAS`.

For more information on the available modules and usage, refer to the Ansible OPNsense Collection Documentation.
