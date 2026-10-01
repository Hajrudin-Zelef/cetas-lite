---
id: collect-261001-meraki/meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed-2
title: "List all networks in an organization"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed.md
source_anchor: ""
source_lines: [64, 183]
sha256: 77254362a96d3c92059c9b542fb9997afb91700c30282cf107703ed11c9fc7cc
---

# List all networks in an organization

| 5 | Shared data model that serves both firmware tools (Ansible + Python) from a single YAML file | Python, Ansible | data-model/firmware/firmware_targets.yaml | 
| 6 | CI/CD automation — auto-apply Terraform when the data model changes, using a self-hosted runner | GitHub Actions | .github/workflows/nac-apply.yml | 
| 7 | Network tag–based targeting — operate on any number of networks without listing individual IDs | Python, Ansible | All scripts | 
ansible-meraki-firmware-upgrade/
│
├── README.md                           ← You are here
│
├── data-model/                         ← Shared YAML configuration (the source of truth)
│   ├── firmware/
│   │   └── firmware_targets.yaml       ← Target firmware versions per network (Ansible + Python)
│   └── nac/
│       ├── networks_nac.yaml           ← SSID + PSK declarations (Terraform NaC only)
│       └── defaults_nac.yaml           ← Org-wide SSID defaults merged by NaC module
│
├── python/                             ← Python scripting approach
│   ├── 1.0_check_available_firmware_by_tag.py   ← Read-only: discover firmware versions
│   ├── 2.0_schedule_firmware_upgrade_by_tag.py  ← Schedule firmware upgrades
│   └── requirements.txt               ← pip dependencies (meraki, PyYAML)
│
├── ansible/                            ← Ansible playbook approach
│   ├── 1.0_check_available_firmware_by_tag.yml  ← Read-only: discover firmware versions
│   ├── 2.0_schedule_firmware_upgrade_by_tag.yml ← Schedule firmware upgrades
│   ├── README.md                       ← Detailed Ansible reference documentation
│   ├── collections/
│   │   └── requirements.yml           ← cisco.meraki Galaxy collection dependency
│   ├── inventory/
│   │   └── meraki.yml                 ← Ansible inventory (localhost only)
│   └── group_vars/all/
│       ├── vars.yml                   ← Non-sensitive shared variables
│       └── vault.yml                  ← API key (Ansible Vault encrypted)
│
├── network-as-code/                    ← Terraform Network as Code approach
│   ├── terraform.tf                   ← Provider configuration (CiscoDevNet/meraki ~> 1.12)
│   ├── main.tf                        ← NaC module declaration with methodology comments
│   ├── vault-envconsul.sh             ← CI/CD wrapper: pull PSKs from HashiCorp Vault
│   ├── .envrc.example                 ← Template for local environment variables
│   └── .gitignore                     ← Excludes .terraform/, state files, .envrc
│
├── .github/
│   └── workflows/
│       └── nac-apply.yml              ← GitHub Actions: auto terraform apply on data model push
│
├── diagrams/                           ← Diagrams (Mermaid PNGs + Draw.io sources)
│   ├── iac-abstraction.png            ← IaC abstraction layer diagram
│   ├── iac-abstraction.mmd            ← Diagram source (Mermaid)
│   ├── repo-workflow.png              ← Data flow diagram
│   ├── repo-workflow.mmd              ← Diagram source (Mermaid)
│   ├── 1.0-meraki-admin-logical-flow.drawio
│   ├── 2.0-pb1-module-mapping.drawio
│   └── 3.0-pb2-module-mapping.drawio
│
├── cisco-workflows/                    ← Reference Meraki workflow templates
│   ├── Appliance - Available Firmware.txt
│   └── Schedule Firmware Upgrade for Networks by Tag.json
│
├── ansible.cfg                         ← Ansible configuration (inventory path, settings)
├── .envrc                             ← direnv environment variables (gitignored)
└── .gitignore                         ← Root gitignore
This section walks you through everything you need to set up your environment and run your first command. Take it one step at a time.
Before you begin, make sure the following are installed and available in your terminal.
| Tool | Minimum Version | How to Check | Install Guide | 
|---|---|---|---|
| Git | 2.x | git --version | git-scm.com | 
| Python | 3.9 | python3 --version | python.org | 
| A Meraki API key | — | Meraki Dashboard → Profile → API access | Meraki docs | 
| Tool | How to Install | 
|---|---|
| meraki Python SDK | pip3 install meraki | 
| PyYAML | pip3 install PyYAML | 
| Tool | Minimum Version | How to Install | 
|---|---|---|
| Ansible Core | 2.15 | pip3 install ansible | 
| cisco.meraki collection | 2.18.0 | ansible-galaxy collection install -r ansible/collections/requirements.yml | 
| Tool | Minimum Version | How to Check | Install Guide | 
|---|---|---|---|
| Terraform | 1.8.0 | terraform version | developer.hashicorp.com | 
| direnv (recommended) | any | direnv version | brew install direnv | 
git clone https://github.com/imanassypov/ansible-meraki-firmware-upgrade.git
cd ansible-meraki-firmware-upgrade
Log into the Meraki Dashboard. Click your username in the top-right corner → My profile → scroll down to API access → Generate new API key.
Copy the key somewhere safe. You will only see it once.
Security note: Treat your API key like a password. It grants full read/write access to your Meraki organization. Never paste it into a command directly in a shared terminal, never commit it to Git.
The safest way to provide your API key to all three tools in this repository is via an environment variable. All tools — the Python SDK, the Ansible cisco.meraki collection, and the Terraform CiscoDevNet/meraki provider — read this same variable automatically.
export MERAKI_API_KEY="your-api-key-here"
If you want this to persist across terminal sessions, add the export line to your shell profile (~/.zshrc, ~/.bash_profile), or use direnv (recommended — see Section 12).
Run a quick check with Python to confirm the key works:
python3 -c "
import meraki, os
dashboard = meraki.DashboardAPI(os.environ['MERAKI_API_KEY'], suppress_logging=True)
orgs = dashboard.organizations.getOrganizations()
print('Connected! Organizations visible to this key:')
for o in orgs:
    print(f'  {o[\"name\"]} (ID: {o[\"id\"]})')
"
Expected output:
Connected! Organizations visible to this key:
  Igor M (ID: 443964)
If you see AuthorizationError, double-check that the API key was copied correctly and that API access is enabled in the Meraki Dashboard.
Both the Python scripts and Ansible playbooks use network tags to find which networks to operate on. Tags are set in the Meraki Dashboard:
- Log into dashboard.meraki.com
- Select your organization
- Go to Network-wide → General for each network
- Under Tags, add a tag (e.g., Cisco-Lab ,branch ,production )
Tags are case-sensitive. The scripts will only operate on networks that carry a tag you specify.
Choose the approach that matches your background and the task you want to accomplish:
| If you want to… | Start with… | 
|---|---|
| Check what firmware versions are available | Section 7 (Python) or Section 8 (Ansible) | 
| Schedule firmware upgrades automatically | Section 7.2 or Section 8.2 | 
| Manage SSID configuration as code | Section 9 (Terraform NaC) | 
| Understand how the data model works | Section 10 | 
| Set up CI/CD automation | Section 11 | 
Python is the Swiss Army knife of network automation. The Meraki Python SDK (pip install meraki) wraps the entire Meraki Dashboard REST API in simple Python objects. Instead of constructing HTTP requests by hand, you call methods:
import meraki, os
dashboard = meraki.DashboardAPI(os.environ['MERAKI_API_KEY'])
# List all networks in an organization
networks = dashboard.organizations.getOrganizationNetworks(organizationId="443964")
# Get firmware information for a specific network
firmware = dashboard.networks.getNetworkFirmwareUpgrades(networkId="L_650207196201616147")
The SDK handles:
