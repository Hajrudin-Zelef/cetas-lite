---
id: collect-260926-mikrotik/mikrotik/mikrotik-cloud-router
title: "*Versions this guide is based on:*"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mikrotik-cloud-router.md
source_anchor: ""
source_lines: [1, 18]
sha256: 86d5d6d03bbf6bff89cff13868e842c9c86413edb17f62614e0ff4f8fb80b5d0
---

# *Versions this guide is based on:*

| EVE Image Foldername | Downloaded Filename | Version | vCPUs | RAM | Console | 
|---|---|---|---|---|---|
| mikrotik-6.40 | chr-6.40.4.img.zip | 6.4.0 | 1 | 256 | Telnet | 

NIC Order tested and working for versions:

**6.44.5, 6.48.6, 7.2.3, 7.3.1, 7.4.1, 7.5.0, 7.6.0**

*NON working version 6.49.7 and 7.1.X Issues with different structure*

| Instructions | 
|---|
| Other versions should also be supported following bellow’s procedure. 1. Download Mikrotik Cloud router image chr-6.40.4.img.zip from: https://mikrotik.com/download 2. Using any archivator program unzip it to get image file chr-6.40.4.img Accordingly our image naming table: http://www.eve-ng.net/index.php/documentation/images-table 3. SSH to your EVE and create directory for Mikrotik node. `mkdir /opt/unetlab/addons/qemu/mikrotik-6.40.4/`4. Upload the chr-6.40.4.img image to the created directory using for example FileZilla or WinSCP. 5. Go to newly created folder and convert the disk to the qcow2 format: ``` cd /opt/unetlab/addons/qemu/mikrotik-6.40.4/
 mv chr-6.40.4.img hda.qcow2
  ``` 6. Fix permissions: ``` /opt/unetlab/wrappers/unl_wrapper -a fixpermissions
  ``` Default username is **admin** with no password**.** |
