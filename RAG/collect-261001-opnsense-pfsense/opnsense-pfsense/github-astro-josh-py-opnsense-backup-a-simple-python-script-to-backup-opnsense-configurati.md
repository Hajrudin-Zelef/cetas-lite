---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-astro-josh-py-opnsense-backup-a-simple-python-script-to-backup-opnsense-configurati
title: "github-astro-josh-py-opnsense-backup-a-simple-python-script-to-backup-opnsense-configurations-with-t"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/github-astro-josh-py-opnsense-backup-a-simple-python-script-to-backup-opnsense-configurations-with-t.md
source_anchor: ""
source_lines: [1, 21]
sha256: 214cc3454fe468fb41c5fe60a7f4071f8d0273a8ed616097cd7e91d0aea6c2b1
---

# github-astro-josh-py-opnsense-backup-a-simple-python-script-to-backup-opnsense-configurations-with-t

A simple python script used in conjunction with the os-api-backup plugin on OPNsense to remotely backup OPNsense configuration.

- OPNSense instance with the `os-api-backup` plugin installed.
- Machine with Python >3.6 and requests library.

1. Install os-api-backup plugin on OPNSense instance.
2. On OPNSense web page UI, navigate to System -> Access -> Groups and click the "+" icon to add a new group.
3. Give group a name such as "backup_group" and save.
4. Click the edit icon for the newly created group and then under "Assigned Privileges" click the edit icon and then check the "Backup API" option in list, leave all other options unchecked.
5. Click save for the Assigned Privileges and click save again to save group settings.
6. Navigate to System -> Access -> Users and click the "+" icon to add a new user, name user whatever you want.
7. Under "Group Membership", add the user to the "backup_group".
8. Click the edit icon for the newly created user. Under "API keys", and click the "+" icon to create a API key to use for the script.
9. A text file will be downloaded containing API key and secret needed for script to run.
10. Click save to save adding API key for user.

`python opnsense_backup/get_opnsense_config.py --key <user_api_key> --secret <user_api_secret> --host <ip_or_fqdn_of_opnsense> --path <optional path to save file>`

`get-opnsense-config --key <user_api_key> --secret <user_api_secret> --host <ip_or_fqdn_of_opnsense> --path <optional path to save file>`

The script will save a configuration file of the name `opnsense-config-YYYY-MM-DD.xml` to path specified or the current working directory if none provided.
