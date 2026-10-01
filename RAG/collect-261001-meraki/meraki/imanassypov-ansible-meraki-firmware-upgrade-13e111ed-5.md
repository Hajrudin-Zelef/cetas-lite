---
id: collect-261001-meraki/meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed-5
title: "List all networks in an organization"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed.md
source_anchor: ""
source_lines: [413, 525]
sha256: da06b77a10cb7044c5b27c9a2fe4ac2afe075bde3252604632b79161dfda67ce
---

# List all networks in an organization

| name | Yes | Exact network name (case-sensitive) as shown in the Meraki Dashboard | 
| tags | Yes | Network tags — used for informational grouping in the data model | 
| firmware.description | No | Human-readable label shown in script output | 
| firmware.upgrade_datetime | No | ISO 8601 UTC timestamp. Empty string = next maintenance window | 
| firmware.upgrade_strategy | No | minimizeClientDowntime (default) orminimizeUpgradeTime | 
| firmware.target.appliance_version_id | No | MX target version ID. Empty string = do not touch | 
| firmware.target.switch_version_id | No | MS target version ID. Empty string = do not touch | 
| firmware.target.wireless_version_id | No | MR target version ID. Empty string = do not touch | 
Version IDs come from running python/1.0_* or ansible/1.0_*. They are stable numeric strings that Meraki uses internally to identify specific firmware builds.
This file defines the desired SSID and wireless configuration. Terraform reads it exclusively.
See Section 9.3 for a full explanation.
For a complete, step-by-step walkthrough of wiring GitHub to the local
runner — including registration, the .env file, a line-by-line workflow
explanation, the Terraform state nuance, and the real issues we hit — see the
dedicated guide: README-GITHUB-ACTIONS.md.
When you run terraform apply locally, the SSID change happens immediately. But what if you want changes to happen automatically whenever a team member edits the data model and pushes to Git? That is what CI/CD (Continuous Integration / Continuous Deployment) provides.
The workflow file .github/workflows/nac-apply.yml defines a GitHub Actions pipeline that:
- Triggers whenever a file in data-model/nac/ is pushed to themain branch
- Runs on a self-hosted runner — your laptop or a dedicated machine — so it has access to local Terraform state and does not need secrets stored in GitHub
- Executes: terraform init →terraform plan →terraform apply
on:
  workflow_dispatch: {}         # Manual trigger from GitHub UI
  push:
    branches: [main]
    paths: ["data-model/nac/**"]    # Only triggers on NaC data model changes
jobs:
  terraform-apply:
    runs-on: self-hosted
    defaults:
      run:
        working-directory: network-as-code
    steps:
      - uses: actions/checkout@v4
      - run: terraform init -input=false
      - run: terraform plan -input=false -out=tfplan
      - run: terraform apply -input=false tfplan
A self-hosted runner is a background process on your machine that listens for jobs from GitHub and executes them locally.
- Go to your repository on GitHub → Settings → Actions → Runners → New self-hosted runner
- Follow the instructions to download and configure the runner
- Set required environment variables in ~/actions-runner/.env :
MERAKI_API_KEY=your-api-key-here
MERAKI_CISCO_LAB_PSK=your-psk-here
MERAKI_FTD_PSK=your-other-psk-here
PATH=/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
Note: The PATH entry is required because the runner's shell does not inherit your interactive shell's PATH. Without it, terraform will not be found even if it is installed.
cd ~/actions-runner
./run.sh
You should see:
√ Connected to GitHub
Listening for Jobs
To survive reboots without manual intervention, install the runner as a launchd service:
cd ~/actions-runner
./svc.sh install
./svc.sh start
The pipeline triggers automatically on any push to main that changes a file under data-model/nac/. It can also be triggered manually from the GitHub UI:
GitHub repository → Actions → NaC — Terraform apply → Run workflow
Never commit credentials to Git. Not in YAML files, not in comments, not in example files. Not even "just for testing."
This repository is designed so that all credentials — the Meraki API key and Wi-Fi PSKs — are provided exclusively through environment variables. No credential ever touches a file that could be committed.
direnv is a shell extension that automatically loads and unloads environment variables when you change directories. It reads from a .envrc file in each directory.
# Install (macOS)
brew install direnv
# Add to your shell profile (~/.zshrc or ~/.bash_profile)
eval "$(direnv hook zsh)"   # for zsh
eval "$(direnv hook bash)"  # for bash
Then create .envrc at the repository root:
# .envrc (this file is gitignored)
export MERAKI_API_KEY=your-api-key-here
export MERAKI_NETWORK_TAGS=Cisco-Lab,FTD
And in network-as-code/.envrc:
# network-as-code/.envrc (also gitignored)
export MERAKI_API_KEY=your-api-key-here
export MERAKI_CISCO_LAB_PSK=your-psk-here
export MERAKI_FTD_PSK=your-other-psk-here
Allow direnv to load each file once:
direnv allow .
direnv allow network-as-code/
After that, variables are injected automatically whenever you cd into the directory.
The self-hosted GitHub Actions runner reads environment variables from ~/actions-runner/.env. This file lives on your machine only — never in the Git repository.
# ~/actions-runner/.env
MERAKI_API_KEY=your-api-key-here
MERAKI_CISCO_LAB_PSK=your-psk-here
MERAKI_FTD_PSK=your-other-psk-here
PATH=/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
For team environments where the API key and PSKs are shared among multiple engineers, HashiCorp Vault is the recommended secret store. The network-as-code/vault-envconsul.sh script wraps any Terraform command to pull PSK values from Vault at runtime using envconsul:
# Store PSK in Vault
vault kv put secret/meraki/cisco-lab-ssid psk=NewSecurePassword
# Apply with PSK injected from Vault
export VAULT_ADDR=https://vault.example.com
export VAULT_TOKEN=<token>
./network-as-code/vault-envconsul.sh terraform apply
| Symptom | Likely Cause | Resolution | 
|---|---|---|
| MERAKI_API_KEY not set | Environment variable missing | Run export MERAKI_API_KEY=your-key-here or use direnv | 
| AuthorizationError (401) | API key invalid or expired | Regenerate the key in Meraki Dashboard → Profile → API access | 
| 403 Forbidden on a specific org | Key does not have access to that org | Check org-level API access in the Meraki Dashboard | 
| No networks found matching tags | Tag misspelled, or networks not tagged | Tags are case-sensitive. Verify tags in Meraki Dashboard → Network-wide → General | 
| Symptom | Likely Cause | Resolution | 
|---|---|---|
| ModuleNotFoundError: No module named 'meraki' | Dependencies not installed | pip3 install -r python/requirements.txt | 
| Script exits immediately with no output | Tags derived from data model but file not found | Run python 1.0_* --tags YourTag to specify tags directly | 
| No networks found | Network not tagged in Meraki Dashboard | Log into dashboard.meraki.com, go to Network-wide → General, and add the tag | 
| Symptom | Likely Cause | Resolution | 
|---|---|---|
| ERROR! A required module is missing: cisco.meraki | Collection not installed | ansible-galaxy collection install -r ansible/collections/requirements.yml | 
| fatal: network_tags is empty | Tag variable not passed | Use -e '{"network_tags": ["YourTag"]}' | 
| ERROR! Attempting to decrypt but no vault secrets found | --vault-password-file missing | Add --vault-password-file .vault_pass to the command | 
| Tag filter returns 0 networks despite correct tag | List passed as string, not array | Always use JSON dict syntax: -e '{"network_tags": ["tag"]}' | 
For the complete Ansible troubleshooting reference, see ansible/README.md.
| Symptom | Likely Cause | Resolution | 
|---|---|---|
| Error: No valid credential sources found | MERAKI_API_KEY not set | export MERAKI_API_KEY=your-key-here | 
| Error: could not find network "Meraki Core" | managed: false lookup fails — network name mismatch | Check exact network name in Meraki Dashboard (case-sensitive) | 
| PSK not updating on Meraki | !env variable not set or set to wrong name | Run echo $MERAKI_CISCO_LAB_PSK to verify the variable is exported | 
