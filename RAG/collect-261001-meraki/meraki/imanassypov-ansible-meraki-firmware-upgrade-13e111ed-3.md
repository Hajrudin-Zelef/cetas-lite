---
id: collect-261001-meraki/meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed-3
title: "List all networks in an organization"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed.md
source_anchor: ""
source_lines: [184, 287]
sha256: d78c3b747c177836fdd02b6663b498de177d8d1ebad64d54413e54ec19546880
---

# List all networks in an organization

- Authentication (injects the API key header on every request)
- Pagination (fetches all pages of results automatically)
- Rate limiting (respects Meraki's API rate limits, retries when throttled)
- Error handling (raises Python exceptions with descriptive messages)
pip3 install -r python/requirements.txt
The requirements.txt file specifies:
meraki       # Meraki Dashboard API Python library
PyYAML       # YAML file parsing (for reading data-model/ files)
File: python/1.0_check_available_firmware_by_tag.py
Purpose: Read-only. Discovers all networks matching your tags and prints current and available firmware versions for each device type (MX appliance, MS switch, MR access point).
The script follows this logic:
1. Read tags from data-model/firmware/firmware_targets.yaml
   (or from --tags argument, or MERAKI_NETWORK_TAGS env var)
2. Authenticate with the Meraki SDK
3. List all organizations visible to the API key
4. For each org: list all networks
5. Filter: keep only networks whose tags overlap with your target tags
6. For each matching network: query /networks/{id}/firmwareUpgrades
7. Print a formatted summary showing:
   - Current firmware version and ID for each device type
   - All available versions with their IDs and release types
The version IDs in the output are what you copy into data-model/firmware/firmware_targets.yaml when you want to schedule an upgrade.
# Auto-read tags from data-model/firmware/firmware_targets.yaml (recommended)
python python/1.0_check_available_firmware_by_tag.py
# Specify tags explicitly
python python/1.0_check_available_firmware_by_tag.py --tags Cisco-Lab
# Multiple tags
python python/1.0_check_available_firmware_by_tag.py --tags Cisco-Lab FTD══════════════════════════════════════════════════════════════════
Network : Meraki Core
Org     : Igor M
Tags    : Cisco-Lab
──────────────────────────────────────────────────────────────────
APPLIANCE (MX)
  Current   : MX 18.211.5.2 (ID: 4625)
  Available :
    - MX 19.2.8 [stable]      →  ID: 15806
    - MX 19.2.4 [stable]      →  ID: 15506
──────────────────────────────────────────────────────────────────
SWITCH (MS)
  Current   : MS 15.21.1 (ID: 3101)
  Available :
    - MS 17.2.2 [stable]      →  ID: 6016
──────────────────────────────────────────────────────────────────
WIRELESS (MR)
  Current   : MR 30.7.1 (ID: 11440)
  Available :
    - MR 32.1.7 [stable]      →  ID: 15763
    - MR 31.5.2 [stable]      →  ID: 13875
══════════════════════════════════════════════════════════════════
From this output, you can see that:
- To upgrade MX to the latest stable, use appliance_version_id: "15806"
- To upgrade MS to the latest stable, use switch_version_id: "6016"
- To upgrade MR to the latest stable, use wireless_version_id: "15763"
File: python/2.0_schedule_firmware_upgrade_by_tag.py
Purpose: Reads data-model/firmware/firmware_targets.yaml and enforces the declared firmware state. For each network, it compares the current version against the target and schedules an upgrade only where they differ.
Idempotent means you can run the same operation multiple times and get the same result — no duplicates, no errors, no unintended side effects.
Script 2.0 is idempotent. If a network is already on the target firmware version, the script skips it and prints COMPLIANT. If you run it 10 times, only the first run (when the firmware actually differs) does anything.
# Normal run — evaluate compliance and schedule upgrades
python python/2.0_schedule_firmware_upgrade_by_tag.py
# Dry run — show the plan without making any API changes
python python/2.0_schedule_firmware_upgrade_by_tag.py --check
Always run with --check first on a new environment.
══════════════════════════════════════════════════════════════════
COMPLIANCE EVALUATION (DRY RUN — no changes will be made)
══════════════════════════════════════════════════════════════════
Network : Meraki Core  |  Org: Igor M
  Appliance : UPGRADE NEEDED   Current: MX 18.211.5.2 (4625) → Target: 15806
  Switch    : UPGRADE NEEDED   Current: MS 15.21.1 (3101)    → Target: 6016
  Wireless  : COMPLIANT        Current: MR 32.1.7 (15763)   — no action needed
──────────────────────────────────────────────────────────────────
Scheduled at : next maintenance window (upgrade_datetime not set)
══════════════════════════════════════════════════════════════════
Ansible expresses automation as a playbook — a YAML file listing tasks to run in order, each calling a module. A module is a self-contained piece of logic that knows how to do one specific thing: create a user, copy a file, call an API endpoint.
The key insight is that Ansible is designed for humans to read. A well-written playbook is close to plain English:
- name: Get all Meraki organizations
  cisco.meraki.organizations_info:
    meraki_api_key: "{{ meraki_api_key }}"
  register: org_list
- name: Filter networks matching our tags
  ansible.builtin.set_fact:
    target_networks: "{{ all_networks | selectattr('tags', 'intersect', network_tags) }}"
Anyone on your team can read this, understand what it does, and propose a change — even without deep Python knowledge.
The cisco.meraki collection is distributed through Ansible Galaxy. Install it with:
ansible-galaxy collection install -r ansible/collections/requirements.yml
This downloads the cisco.meraki collection (version ≥ 2.18.0) and makes its modules available to Ansible. The requirements.yml file pins the version so your automation is reproducible.
File: ansible/1.0_check_available_firmware_by_tag.yml
Purpose: Mirrors python/1.0_* exactly — same data, different tool.
# From the repository root
ansible-playbook ansible/1.0_check_available_firmware_by_tag.yml \
  -i ansible/inventory/meraki.yml \
  -e '{"network_tags": ["Cisco-Lab"]}' \
  --vault-password-file .vault_pass
Important: Always use the JSON dict format (-e '{"network_tags": [...]}') for list variables in Ansible. The shorthand key=value format passes the value as a string, not a list, and the tag filter will not work correctly.
Unlike traditional Ansible that SSH into devices, these playbooks connect only to localhost. All Meraki API calls happen over HTTPS from your machine to api.meraki.com. Ansible is acting as an HTTP client, not a remote executor.
Your machine (localhost)
      │
      │  HTTPS port 443
      ▼
api.meraki.com  →  Your Meraki Org → Networks → Devices
No SSH keys. No device credentials. No firewall rules needed on the devices themselves.
File: ansible/2.0_schedule_firmware_upgrade_by_tag.yml
Purpose: Mirrors python/2.0_*. Reads data-model/firmware/firmware_targets.yaml, compares current state to desired state, and schedules upgrades only where needed.
ansible-playbook ansible/2.0_schedule_firmware_upgrade_by_tag.yml \
  -i ansible/inventory/meraki.yml \
