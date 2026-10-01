---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-joakimmelin-opnsense-backup-simple-script-for-backing-up-opnsense-configuration-to-
title: "github-joakimmelin-opnsense-backup-simple-script-for-backing-up-opnsense-configuration-to-another-se"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/github-joakimmelin-opnsense-backup-simple-script-for-backing-up-opnsense-configuration-to-another-se.md
source_anchor: ""
source_lines: [1, 16]
sha256: 9db86b1608e59235d7b837c121581e9e8f8fdd2aa23265d97c5a64632fcb1311
---

# github-joakimmelin-opnsense-backup-simple-script-for-backing-up-opnsense-configuration-to-another-se

A simple script for backing up OPNSense configuration to another server. Feel free to copy and/or modify as needed.

The script is supposed to run every 24 hours and creates a unique file name every time (`<hostname>-config-<date>.xml`).

Requires the following:

1. Install Bash (`pkg install bash` )
2. Make sure login via SSH works without password (i.e with SSH key)

The script is pretty self-explanatory. Save it where ever you want to run it but /usr/local/etc is probably best.

Create a new entry in root's crontab:

`*       *       2       *       *       (/usr/local/etc/backup.sh) >/dev/null 2>&1`

Make sure to modify the script and include hostname, user and target directory for where the backups will be sent.
