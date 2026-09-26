---
id: collect-260926-mikrotik/mikrotik/routeros-scripts-doc-backup-email-md-at-b5acba6d53b4432bfe272a32a439c07063731ff6-eworm-de-
title: "routeros-scripts-doc-backup-email-md-at-b5acba6d53b4432bfe272a32a439c07063731ff6-eworm-de-routeros-s"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/routeros-scripts-doc-backup-email-md-at-b5acba6d53b4432bfe272a32a439c07063731ff6-eworm-de-routeros-s.md
source_anchor: ""
source_lines: [1, 42]
sha256: fdd1ce8c38b708ad8514b93469dac80b563c9c8004934d2154e9f7c9a845dc5a
---

# routeros-scripts-doc-backup-email-md-at-b5acba6d53b4432bfe272a32a439c07063731ff6-eworm-de-routeros-s

ℹ️ **Info**: This script can not be used on its own but requires the base
installation. See main README for details.


This script sends binary backup (`/system/backup/save`) and complete
configuration export (`/export terse show-sensitive`) via e-mail.

Just install the script and the required module:

```
$ScriptInstallUpdate mod/notification-email,backup-email;
```
Also make sure you configure sending notifications via e-mail.

The configuration goes to `global-config-overlay`, these are the parameters:

- `BackupFileNameDate` : whether to add date & time in filenames
- `BackupSendBinary` : whether to send binary backup
- `BackupSendExport` : whether to send configuration export
- `BackupSendGlobalConfig` : whether to send`global-config-overlay`
- `BackupPassword` : password to encrypt the backup with
- `BackupRandomDelay` : delay up to amount of seconds when run from scheduler

ℹ️ **Info**: Copy relevant configuration from
`global-config` (the one without `-overlay`) to
your local `global-config-overlay` and modify it to your specific needs.


Just run the script:

```
/system/script/run backup-email;
```
Creating a scheduler may be an option:

```
/system/scheduler/add interval=1w name=backup-email on-event="/system/script/run backup-email;" start-time=09:15:00;
```
- Upload backup to Mikrotik cloud
- Save configuration to fallback partition
- Send notifications via e-mail
- Upload backup to server
