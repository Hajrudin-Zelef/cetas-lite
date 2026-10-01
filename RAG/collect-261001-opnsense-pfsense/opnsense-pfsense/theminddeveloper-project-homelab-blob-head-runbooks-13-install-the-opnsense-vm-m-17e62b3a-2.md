---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/theminddeveloper-project-homelab-blob-head-runbooks-13-install-the-opnsense-vm-m-17e62b3a-2
title: "1 · the VM is running and set to autostart"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/theminddeveloper-project-homelab-blob-head-runbooks-13-install-the-opnsense-vm-m-17e62b3a.md
source_anchor: ""
source_lines: [208, 224]
sha256: c5d86cad13a87503f897a87f8b3055c6d50d5d103be13d49dc8674df25f57c30
---

# 1 · the VM is running and set to autostart

| Symptom | Cause | Fix | 
|---|---|---|
| Boots the installer again | ISO still attached, SeaBIOS fell back to CD | `qm set 200 --ide2 none,media=cdrom` | 
| "Failed to connect to server" on the console | browser has not accepted the hosting node's certificate | open the console from the node hosting the guest, or visit its `:8006` once | 
| GUI at `.60` unreachable | "Block private networks" / "Block bogon networks" ticked | untick both; if locked out, runbook 17 | 
| GUI unreachable, checkboxes already correct | address in `config.xml` but not on the interface | `qm reboot 200` | 
| "This IP address conflicts with another interface" | `.60` was assigned to the wrong leg first | move the wrong one to something else, then reassign | 
| Config importer appears at boot | a key was pressed during boot | press Enter on an empty line | 
| Everything you configured is gone after a reboot | you were in live mode | check for `/dev/gpt/rootfs` , reinstall properly | 

```
qm stop 200
qm destroy 200
```
The bridge and the containers are untouched. If containers have already been
moved into the DMZ, move them back to `vmbr0` first or they lose all
connectivity.
