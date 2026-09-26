---
id: collect-260926-mikrotik/mikrotik/github-hreskiv-chr-eve-ng-repository-contain-lab-s-for-mikrotik-trainings-and-helpers-for-
title: "On your EVE-NG server:"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "training"]
source: docs/RAG/lot-mikrotik/chr-proxmox/github-hreskiv-chr-eve-ng-repository-contain-lab-s-for-mikrotik-trainings-and-helpers-for-installing.md
source_anchor: ""
source_lines: [1, 69]
sha256: d23f3602687964a927b9d72cbc856fe480a4c338d2c88c8b9af51f24cf61f38a
---

# On your EVE-NG server:

Bash script that automates adding **MikroTik Cloud Hosted Router (CHR)** images into EVE-NG. Downloads, converts, and deploys CHR images — ready to use in seconds.

```
# On your EVE-NG server:
curl -O https://raw.githubusercontent.com/hreskiv/chr-eve-ng/master/chr-eve.sh
chmod +x chr-eve.sh
sudo ./chr-eve.sh install --version 7.20
```
The CHR node is now available in EVE-NG — add it to any topology.

- Downloads CHR images from the official MikroTik CDN
- Supports **all release channels** : stable, long-term, testing, RC (e.g.`7.20rc5` )
- Converts raw images to `qcow2` format automatically
- Detects image format (raw/qcow2) and handles each correctly
- Colored, step-by-step terminal output
- Runs `fixpermissions` automatically
- `--dry-run` mode to preview actions without making changes
- Install from a **local file** (no download needed)

- **EVE-NG** (Community or Pro) — script runs on the EVE-NG server itself
- `bash` ,`curl` ,`unzip` ,`qemu-img`
- Root privileges (`sudo` )

```
sudo ./chr-eve.sh install --version 7.20
sudo ./chr-eve.sh install --version 7.20rc5
sudo ./chr-eve.sh install --version 7.20 --force       # overwrite existing
sudo ./chr-eve.sh install --version 7.19.4 --local /tmp/chr-7.19.4.img  # from local file
sudo ./chr-eve.sh install --version 7.20 --dry-run     # preview only
sudo ./chr-eve.sh install --version 7.20 --name my-chr # custom directory name
```
| Flag | Description | 
|---|---|
| `--version` | **(required)** RouterOS version — e.g.`7.20` ,`7.11rc1` ,`6.49.17` | 
| `--local` | Path to a local `.img` file (skip download) | 
| `--name` | Custom directory name under `/opt/unetlab/addons/qemu/` | 
| `--force` | Overwrite if the version is already installed | 
| `--dry-run` | Show what would happen without making changes | 

`sudo ./chr-eve.sh list````
sudo ./chr-eve.sh remove --version 7.19.4
sudo ./chr-eve.sh remove --name mikrotik-7.19.4   # by directory name
```
Images are placed in:

```
/opt/unetlab/addons/qemu/mikrotik-<version>/hda.qcow2
```
In EVE-NG, add a node of type **MikroTik** and select the installed version from the dropdown.

|  | Supported | 
|---|---|
| EVE-NG | Community & Pro | 
| RouterOS | v6.x and v7.x | 
| Channels | stable, long-term, testing, RC, beta | 
| Host OS | Ubuntu 20.04 / 22.04 (EVE-NG base) | 

The `Labs/` folder contains ready-to-use EVE-NG lab topologies for MikroTik training courses.

`CHR-EVE-slides.pdf` — presentation slides covering CHR deployment in EVE-NG.

**Ihor Hreskiv** — MikroTik Certified Trainer

- mtik.pl — MikroTik training (Poland, Kraków)
- mtik.tech — MikroTik training (Ukraine, online)
- YouTube PL · YouTube UA
- LinkedIn · GitHub

This project is provided as-is. See LICENSE for details.
