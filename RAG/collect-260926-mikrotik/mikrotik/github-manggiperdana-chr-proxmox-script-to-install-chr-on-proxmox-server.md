---
id: collect-260926-mikrotik/mikrotik/github-manggiperdana-chr-proxmox-script-to-install-chr-on-proxmox-server
title: "github-manggiperdana-chr-proxmox-script-to-install-chr-on-proxmox-server"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "mit license"]
source: docs/RAG/lot-mikrotik/chr-proxmox/github-manggiperdana-chr-proxmox-script-to-install-chr-on-proxmox-server.md
source_anchor: ""
source_lines: [1, 30]
sha256: 8b1b2af92468632c6baf285767888bc0e1e524089016ed9da28200267353680b
---

# github-manggiperdana-chr-proxmox-script-to-install-chr-on-proxmox-server

A simple shell script to automate the installation of **MikroTik Cloud Hosted Router (CHR)** on **Proxmox VE**.

This script helps you deploy CHR quickly and efficiently from the command line without manual configuration.

| File | Description | 
|---|---|
| `chr-install.sh` | Main shell script to install CHR on Proxmox | 

There are two ways to run the script, depending on your preference:

This method is ideal if you want to execute the script instantly without cloning the repository.

`bash <(curl -s https://raw.githubusercontent.com/manggiperdana/CHR-Proxmox/main/chr-install.sh)`
This method is better if you want to inspect or modify the script beforehand.

1. Clone the repository:

```
git clone https://github.com/manggiperdana/CHR-Proxmox.git
cd CHR-Proxmox
```
1. Make the script executable and run it:

```
chmod +x chr-install.sh
./chr-install.sh
```
This project is licensed under the MIT License.

You are free to use, modify, and distribute this script. Please note that it is provided as-is without any warranty.
