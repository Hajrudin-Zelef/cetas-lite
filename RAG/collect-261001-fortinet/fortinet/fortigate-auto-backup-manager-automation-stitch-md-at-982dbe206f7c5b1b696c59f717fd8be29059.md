---
id: collect-261001-fortinet/fortinet/fortigate-auto-backup-manager-automation-stitch-md-at-982dbe206f7c5b1b696c59f717fd8be29059
title: "fortigate-auto-backup-manager-automation-stitch-md-at-982dbe206f7c5b1b696c59f717fd8be29059d107-crysi"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-auto-backup-manager-automation-stitch-md-at-982dbe206f7c5b1b696c59f717fd8be29059d107-crysi.md
source_anchor: ""
source_lines: [1, 101]
sha256: a91c01965ea25239518cfe762267bda41c2a61f90df50c68e88270696df45c85
---

# fortigate-auto-backup-manager-automation-stitch-md-at-982dbe206f7c5b1b696c59f717fd8be29059d107-crysi

This document describes the configuration of automated daily backups on FortiGate using **Automation Stitch**, integrated with an external backup processing system.

The solution ensures:

- Automated configuration backups
- Centralized storage via FTP
- Integration with external rotation and notification systems

1. 
FortiGate generates a configuration backup
2. 
Backup is uploaded to an FTP server
3. 
Windows-based automation: 
  - Performs backup rotation (cleanup)
  - Sends email notification

Navigate to:

**Security Fabric → Automation → Actions**

- 
Click **Create New**
- 
Set: 
  - **Action Type:** CLI Script
  - **Name:**`Auto_Backup_Config`
- 
Add the following script:

`execute backup config ftp backup.conf <FTP_IP> <USERNAME> <PASSWORD>`
- Click **OK / Save**

Navigate to:

**Security Fabric → Automation → Triggers**

- 
Click **Create New**
- 
Set: 
  - **Name:**`Daily_Backup_Trigger`
  - **Type:** Scheduled
  - **Time:** 09:00 AM (Daily)
- 
Click **OK / Save**

Navigate to:

**Security Fabric → Automation → Stitches**

- 
Click **Create New**
- 
Set: 
  - **Name:**`Daily_Backup_Stitch`
  - **Trigger:**`Daily_Backup_Trigger`
  - **Action:**`Auto_Backup_Config`
- 
Enable the stitch

FortiGate uploads backup files to a centralized FTP server.

- 
Backup generated on FortiGate
- 
File transferred to FTP server
- 
External system accesses backup for: 
  - Rotation
  - Reporting

- Centralized storage
- Easy automation integration
- Separation of concerns (Firewall vs Processing)

```
FortiGate → FTP Server → Windows Scheduler → Rotation → Email Notification
```
- Do not expose FTP credentials in public repositories
- Restrict access to FTP server
- Use secure password policies
- Store sensitive data outside scripts (environment variables preferred)

- Fully automated daily backup lifecycle
- Zero manual intervention
- Improved disaster recovery readiness
- Clean and maintainable infrastructure

Replace the following placeholders with your actual values:

- `<FTP_IP>`
- `<USERNAME>`
- `<PASSWORD>`

Never commit real credentials to version control systems.

- Secure FTP (SFTP/FTPS) support
- Cloud storage integration
- Backup integrity validation
- Real-time alerting system
