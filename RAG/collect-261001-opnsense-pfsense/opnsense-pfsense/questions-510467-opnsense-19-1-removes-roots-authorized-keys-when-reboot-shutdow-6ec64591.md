---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-510467-opnsense-19-1-removes-roots-authorized-keys-when-reboot-shutdow-6ec64591
title: "questions-510467-opnsense-19-1-removes-roots-authorized-keys-when-reboot-shutdow-6ec64591"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-510467-opnsense-19-1-removes-roots-authorized-keys-when-reboot-shutdow-6ec64591.md
source_anchor: ""
source_lines: [1, 11]
sha256: 3d059d52309d5adce9567f449be76026f2c61b0997592697c1e5f16e62a5876e
---

# questions-510467-opnsense-19-1-removes-roots-authorized-keys-when-reboot-shutdow-6ec64591

I have an OPNsense 19.1 virtual firewall, which is based on FreeBSD 11.2 HBSD (HardenedBSD).

I added it to Ansible Server inventory (with the right configurations), and once copied the PK to that BSD /root/.ssh/authorized_keys, I can connect to it (yes, running commands as root, it is in an isolated training environment).

The issue is that **when I reboot or turn off and on the machine, that authorized_keys file disappears**.

The file permissions are 600 (-rw-------).

Do you know why this happens and how to solve it? As far as I could see, in a normal FreeBSD I think there's no problem in having authorized_keys file in the root account...

`tmpfs`? Run`mount`without arguments to check.
