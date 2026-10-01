---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1714344-configure-opnsense-to-check-for-updates-and-install-them-autom-310f1a0b
title: "questions-1714344-configure-opnsense-to-check-for-updates-and-install-them-autom-310f1a0b"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1714344-configure-opnsense-to-check-for-updates-and-install-them-autom-310f1a0b.md
source_anchor: ""
source_lines: [1, 11]
sha256: 9725a659a1142c0e66453c067aaf3d8955acedd612f615e400a18befb05d54e4
---

# questions-1714344-configure-opnsense-to-check-for-updates-and-install-them-autom-310f1a0b

I run OPNsense as well as a few Linux systems with Webmin.

Webmin has an option to check for updates periodically and install them automatically (iirc this can be done for all updates, or just for the security-relevant ones). This ensures that the system is always up to date, or at least has all security fixes installed.

OPNsense, on the other hand, does not seem to have an option for that (correct me if I am wrong). However, if checking for updates and installing them can be done via the command line and without user input, it should be possible to run that as a cron job.

What command would I need to run for that? (Or is there an easier way?)

Btw, for a discussion of the pros and cons of automation, see https://security.stackexchange.com/q/183173/49551.

`opnsense-update`. However, if the system is up to date, it just prints`Nothing to do`, so I’ll have to wait for the next updates. There usually is one every 3 weeks or so, the last one was 12 days ago. github.com/opnsense/update has some information on the update tools bundled with OPNsense.`opnsense-update`clains there is nothing to do, whereas the update check in the web UI finds new packages to install. Even after that check (without actually installing the packages),`opnsense-update`still finds nothing. The man page also provides no hints to any command line options I might have missed.
