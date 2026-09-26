---
id: collect-260926-mikrotik/mikrotik/github-simda-buru-chr-xx-mikrotik-cloud-hosted-router-chr-on-proxmox-ve
title: "github-simda-buru-chr-xx-mikrotik-cloud-hosted-router-chr-on-proxmox-ve"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "mit license"]
source: docs/RAG/lot-mikrotik/chr-proxmox/github-simda-buru-chr-xx-mikrotik-cloud-hosted-router-chr-on-proxmox-ve.md
source_anchor: ""
source_lines: [1, 33]
sha256: 947747ab03aa5d2b5fd21a4c300bf687b10d986015eb461bade2a3a0556f7cf2
---

# github-simda-buru-chr-xx-mikrotik-cloud-hosted-router-chr-on-proxmox-ve

**MikroTik Cloud Hosted Router (CHR)** on **Proxmox VE**

A simple shell script to automate the installation of **MikroTik Cloud Hosted Router (CHR)** on **Proxmox VE**.

This script helps you deploy CHR quickly and efficiently from the command line without manual configuration.

| File | Description | 
|---|---|
| `install-chr.sh` | Main shell script to install CHR on Proxmox | 
| `install-chr-7.sh` | Main shell script to install CHR on Proxmox | 

There are two ways to run the script, depending on your preference:

This method is ideal if you want to execute the script instantly without cloning the repository.

`bash <(curl -s https://raw.githubusercontent.com/simda-buru/CHR-XX/main/install-chr.sh)``bash <(curl -s https://raw.githubusercontent.com/simda-buru/CHR-XX/main/install-chr-7.sh)`
This method is better if you want to inspect or modify the script beforehand.

1. Clone the repository:

```
git clone https://github.com/simda-buru/CHR-XX.git
cd CHR-Proxmox
```
1. Make the script executable and run it:

```
chmod +x install-chr.sh
./install-chr.sh
```
This project is licensed under the MIT License.

You are free to use, modify, and distribute this script. Please note that it is provided as-is without any warranty.
