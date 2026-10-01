---
id: collect-261001-meraki/meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed-4
title: "List all networks in an organization"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed.md
source_anchor: ""
source_lines: [288, 412]
sha256: af99342657a281197415b97e6e6d286e261dc723d9eef4f4e1d42bc8b049a78b
---

# List all networks in an organization

  -e '{"network_tags": ["Cisco-Lab"]}' \
  --vault-password-file .vault_passansible-playbook ansible/2.0_schedule_firmware_upgrade_by_tag.yml \
  -i ansible/inventory/meraki.yml \
  -e '{"network_tags": ["Cisco-Lab"]}' \
  --vault-password-file .vault_pass \
  --check
Your Meraki API key grants full access to your organization. Never commit it in plain text.
Ansible provides Ansible Vault — an encryption feature built into Ansible that encrypts files so they can be safely committed to Git. The key is decrypted in memory at runtime — it never touches the disk unencrypted.
# 1. Create a vault password file (strong random password)
echo "$(openssl rand -base64 32)" > .vault_pass
chmod 600 .vault_pass
# 2. Add to gitignore (critical — never commit this file)
echo ".vault_pass" >> .gitignore
# 3. Encrypt the vault file that contains your API key
ansible-vault encrypt ansible/group_vars/all/vault.yml \
  --vault-password-file .vault_pass
After encryption, vault.yml looks like this (safe to commit):
$ANSIBLE_VAULT;1.1;AES256
30333534343937636662323135663063...
For a complete Ansible reference, including all playbook variables, idempotency details, and troubleshooting, see the detailed ansible/README.md.
If Python scripting is imperative ("do step 1, then step 2, then step 3") and Ansible is procedural ("run these tasks in order"), then Terraform is declarative:
You describe the desired final state. Terraform figures out how to get there.
You do not write code that says "create this SSID, then set the PSK, then set the encryption mode." You write a YAML file that says "I want this SSID to exist with these properties" — and Terraform calls the Meraki API to make that true, whatever the current state happens to be.
The netascode/nac-meraki Terraform module is the engine. It reads your YAML data model files and translates them into Meraki API calls using the CiscoDevNet/meraki Terraform provider.
The execution chain:
terraform apply
   │
   ├── reads network-as-code/main.tf
   │
   ├── invokes module "netascode/nac-meraki/meraki"
   │    └── scans yaml_directories: ["../data-model/nac"]
   │
   ├── reads data-model/nac/defaults_nac.yaml   ← org-wide defaults
   ├── reads data-model/nac/networks_nac.yaml   ← your SSID declarations
   │
   ├── merges defaults + network config → unified model
   │
   └── calls CiscoDevNet/meraki provider
        └── Meraki Dashboard API
             ├── GET  /organizations/{id}/networks   (look up existing networks)
             └── PUT  /networks/{id}/wireless/ssids/{n}  (configure SSIDs)
The SSID configuration lives in data-model/nac/networks_nac.yaml. This file is the single source of truth for SSID state. Here is what it looks like:
meraki:
  domains:
    - name: "meraki.local"
      organizations:
        - name: "Igor M"
          managed: false          # Org already exists — do not create
          networks:
            - name: "Meraki Core"
              managed: false      # Network already exists — do not create
              tags: [Cisco-Lab]
              wireless:
                ssids:
                  - name: "Test-WiFi"
                    ssid_number: "0"
                    enabled: true
                    auth_mode: "psk"
                    psk: !env MERAKI_CISCO_LAB_PSK    # ← secret, never stored in file
                    encryption_mode: "wpa"
                    wpa_encryption_mode: "WPA2 only"
                    visible: true
Two things worth noting:
managed: false — This tells the NaC module that the organization and network already exist in the Meraki Dashboard. Terraform will look them up (using data sources) rather than trying to create them. This is the correct setting for most real-world environments where the org and networks were created manually or by another process.
!env MERAKI_CISCO_LAB_PSK — The !env YAML tag is a NaC module feature. At runtime, the module substitutes this tag with the value of the named environment variable. The PSK is never written to any file, never appears in terraform plan output, and never gets committed to Git.
cd network-as-code
terraform init
terraform init downloads the netascode/nac-meraki module and the CiscoDevNet/meraki provider from the Terraform Registry. This only needs to be done once (or when you upgrade provider versions).
Copy the example file and populate it:
cp network-as-code/.envrc.example network-as-code/.envrc
Edit .envrc:
export MERAKI_API_KEY=your-api-key-here
export MERAKI_CISCO_LAB_PSK=your-wifi-password-here
export MERAKI_FTD_PSK=your-other-wifi-password-here
If you use direnv, run direnv allow in the network-as-code/ directory. Variables will be injected automatically whenever you cd into that directory.
Always plan before applying:
terraform plan
Terraform shows exactly what it will do. You should see something like:
  # module.meraki.meraki_wireless_ssid.networks_wireless_ssids["meraki.local/Igor M/Meraki Core/Test-WiFi"] will be updated in-place
  ~ resource "meraki_wireless_ssid" "networks_wireless_ssids" {
      ~ psk      = (sensitive value)
        name     = "Test-WiFi"
        # ...
    }
Plan: 0 to add, 1 to change, 0 to destroy.
If the plan shows unexpected deletions or creations, investigate before applying.
terraform apply
Terraform prompts for confirmation. Review the plan summary, type yes, and Terraform calls the Meraki API to apply the changes.
Rotating a Wi-Fi password is now a one-line operation:
# Direct (local dev)
export MERAKI_CISCO_LAB_PSK=NewSecurePassword123
terraform apply
# Via HashiCorp Vault (CI/CD)
vault kv put secret/meraki/cisco-lab-ssid psk=NewSecurePassword123
./network-as-code/vault-envconsul.sh terraform apply
The PSK change appears in the Meraki Dashboard within seconds of terraform apply completing.
You may notice that data-model/ has two subdirectories:
data-model/
├── firmware/     ← Read by Python and Ansible ONLY
│   └── firmware_targets.yaml
└── nac/          ← Read by Terraform ONLY
    ├── networks_nac.yaml
    └── defaults_nac.yaml
Both firmware_targets.yaml and networks_nac.yaml use the same meraki: root key in their YAML structure. If Terraform were to read both files, it would merge them into a single configuration tree — with unpredictable results. By placing them in separate directories and configuring Terraform to only scan data-model/nac/, the two data models are kept completely isolated. Python and Ansible never see data-model/nac/, and Terraform never sees data-model/firmware/.
The data model is the definition of what you want your network to look like — written in YAML, stored in data-model/, and consumed by your automation tools. It is the bridge between human intent and machine action.
Diagram source: diagrams/repo-workflow.mmd
Think of the data model as a configuration contract: "Networks tagged Cisco-Lab should be running MR firmware version 15763, with SSID Test-WiFi on slot 0, using WPA2 with the PSK in MERAKI_CISCO_LAB_PSK." Any tool that reads the contract and talks to the Meraki API will produce the same outcome.
This file defines the desired firmware end state for each network. Python scripts and Ansible playbooks read it.
meraki:
  domains:
    - organizations:
        - name: "Igor M"
          networks:
            - name: "Meraki Core"
              tags: [Cisco-Lab]
              firmware:
                description: "Cisco-Lab — all device families stable firmware"
                upgrade_datetime: ""           # "" = next maintenance window
                upgrade_strategy: "minimizeClientDowntime"
                target:
                  appliance_version_id: "15806"  # MX 19.2.8 stable
                  switch_version_id: "6016"       # MS 17.2.2 stable
                  wireless_version_id: "15763"    # MR 32.1.7 stable
| Field | Required | Description | 
|---|---|---|
