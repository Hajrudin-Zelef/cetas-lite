---
id: collect-260926-mikrotik/mikrotik/installing-mikrotik-image-on-pnetlab-or-eve-ng-simple-and-easy-undercode-testing
title: "installing-mikrotik-image-on-pnetlab-or-eve-ng-simple-and-easy-undercode-testing"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/lot-mikrotik/forum/misc/installing-mikrotik-image-on-pnetlab-or-eve-ng-simple-and-easy-undercode-testing.md
source_anchor: ""
source_lines: [1, 77]
sha256: 7d89b4ffdfcd5da3be7249e9b9b365f5d14e534ba130a95bc497a41401a86dca
---

# installing-mikrotik-image-on-pnetlab-or-eve-ng-simple-and-easy-undercode-testing

## Listen to this Post

In this tutorial, we’ll guide you through installing a MikroTik image on PNETLab or EVE-NG, essential tools for network simulation and cybersecurity labs.

## Steps to Install MikroTik on PNETLab/EVE-NG:

## 1. Download MikroTik CHR Image:

- Get the MikroTik Cloud Hosted Router (CHR) image from the official site:
wget https://download.mikrotik.com/routeros/CHR-7.11.2.img.zip -O chr.img.zip
- Extract the image:
unzip chr.img.zip

## 2. Upload to PNETLab/EVE-NG:

- Use SCP to transfer the `.img` file to your lab server:
scp chr.img root@your-eve-ng-server:/opt/unetlab/addons/qemu/mikrotik-chr/
- Set correct permissions:
chmod 777 /opt/unetlab/addons/qemu/mikrotik-chr/chr.img

## 3. Create a New Lab in EVE-NG:

- Add a new node, select QEMU, and choose the MikroTik CHR image.
- Configure CPU/RAM as needed (e.g., 1 vCPU, 512MB RAM).

## 4. Start and Configure MikroTik:

- Boot the device and log in via console (default credentials: `admin` /no password).
- Set up interfaces and IPs:
/ip address add address=192.168.1.1/24 interface=ether1

## You Should Know:

- Verify Image Integrity:

## Use `sha256sum` to ensure the image isn’t corrupted:

sha256sum chr.img

- Automate with Scripts:

## For mass deployments, automate using:

```
for node in {1..5}; do 
cp chr.img /opt/unetlab/addons/qemu/mikrotik-chr/node$node.img
done
```
- Troubleshooting Tips:
- If EVE-NG doesn’t detect the image, restart the service:
/etc/init.d/unetlab restart
- Use `lsmod | grep kvm` to verify KVM support for QEMU.
- 
MikroTik CLI Cheatsheet: /system reboot Reboot device /interface print List interfaces /ip firewall nat add chain=srcnat out-interface=ether1 action=masquerade Enable NAT

## What Undercode Say:

MikroTik on EVE-NG/PNETLab is a game-changer for network engineers and pentesters. Here are extra Linux/Windows commands to enhance your lab:

- Linux Networking:
ip link set eth0 up Bring interface up tcpdump -i eth0 port 8291 Capture MikroTik Winbox traffic
- 
Windows PowerShell: Test-NetConnection 192.168.1.1 -Port 22 Check SSH access
- 
EVE-NG Management: virsh list --all List all VMs

Expected Output: A fully functional MikroTik router in your virtual lab, ready for routing, firewall testing, or malware analysis.

Reference: MikroTik CHR Download | EVE-NG Docs

## References:

Reported By: Matheuscazzi Instalando – Hackers Feeds

Extra Hub: Undercode MoN

Basic Verification: Pass ✅
