---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-672292-how-can-i-use-cloud-init-nocloud-with-opnsense-21-54ff8af5
title: "How can I use cloud-init NoCloud with OPNsense 21?"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-672292-how-can-i-use-cloud-init-nocloud-with-opnsense-21-54ff8af5.md
source_anchor: ""
source_lines: [1, 39]
sha256: a40ca1f42b77083bbe1c9df0e5d539752c974fa5119e7d35d61753def193443c
---

# How can I use cloud-init NoCloud with OPNsense 21?

*Score : 1 | Source : https://unix.stackexchange.com/questions/672292/how-can-i-use-cloud-init-nocloud-with-opnsense-21*

I'm new to OPNSense (also to FreeBSD in general) and I'm interested to use cloud-init to configure at least LAN (vtnet0) with static ip address, root password and eventually running custom scripts (or shell commands) in a OPNsense VM created with Qemu to apply a custom configuration.
I saw that opnsense github repo has cloud-init port, so I installed it with:
pkg install net/cloud-init
Then I added cidata.iso image to Qemu as required by cloud-init NoCloud with user-data and meta-data. I already tested those files on ubuntu server 21 and kali linux. They are correct at least on those OSs ;)
I found the cdrom as /dev/cd0 and mounted with
mkdir -p /media/cdrom
mount -t cd9660 /dev/cd0 /media/cdrom
I also edited /etc/fstab appending this line:
/dev/cd0 /media/cdrom cd9660 ro 0 0
to auto mount the cdrom at boot.
Finally, I created (because unexisting) /etc/rc.conf with this content:
cloudinit_enable="YES"
and restarted my OPNsense VM.
What I expect now is that cloud-init will start automatically at boot.
However this doesn't happen, probably because I have to configure something.
If I run cloud-init init via terminal it throws error:
stages.py[WARNING]: Failed to rename devices: Unexpected error while running command.
Command [`ip`, `-6`, `addr`,`show`, `permanent`, `scope`,`global`]
Exit code: -
Reason: [Errno 2] No such file or directory: b`ip`
Stdout: -
Stderr: -
No `init` modules to run under section `cloud_init_modules`
On both Kali Linux and Ubuntu Server it works easely.
I have some questions about this:
I already posted this question here, but I didn't receive answers.

---

### Reponse — score 3

I'm working on expanding cloud-init support to BSDs*. The errors should be fixed by now since at least https://github.com/canonical/cloud-init/pull/1779
And by now, Vultr is fairly well supported.
If something doesn't work in net/cloud-init there's a great chance that it works in net/cloud-init-devel. And if it doesn't work there, I'm very happy to fix reported issues. cloud-init has moved their bug tracker from Launchpad to GitHub, which might make it easier to contribute.
*mostly FreeBSD, since this is a FreeBSD Foundation sponsored project, and I'm most familiar with FreeBSD)
