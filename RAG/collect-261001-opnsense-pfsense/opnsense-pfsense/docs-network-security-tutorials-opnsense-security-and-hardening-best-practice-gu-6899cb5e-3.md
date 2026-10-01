---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e-3
title: "docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e.md
source_anchor: ""
source_lines: [323, 500]
sha256: b0b694839db94937eec902d1478cdab88eac0802c16998ece34d2bcc405b9a29
---

# docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e

By following the steps you can disable ssh access on OPNsense:

1. 
Navigate to the **System** >**Settings** >**Administration** .
2. 
Uncheck `Enable Secure Shell` option in the**Secure Shell** pane.

**Figure 12.** *Disable SSH Connections*

## Regularly Backup and Protect Backup Files

A software or hardware failure, a human-caused incident, or even natural disasters such as floods, fires, earthquakes, or tornadoes might result in a total system wipeout. And sometimes unfortunate events occur when we least anticipate them or are least prepared for them. The data backup makes it accessible in the event of data loss or corruption. You can only retrieve data from a previous time period if you have a backup.

To safeguard system data and ensure its accessibility, you must regularly back up your OPNsense. To harden your OPNsense security you need to protect your backup files.

### How to Backup/Restore Configuration Manually?

You may manually backup your OPNsense configuration by following the steps given below:

1. 
Navigate to the **System** >**Configuration** >**Backups** .
2. 
You may enable **Encrypt this configuration file** option. If you select this option, fill in the`Password` and`Confirmation` fields with a strong password.
3. 
Click **Download configuration** to download the configuration file.

**Figure 13.** *Backup OPNsense Configuration Manually*

You may manually restore your OPNsense configuration from a backup file by following the steps given below:

1. 
Navigate to the **System** >**Configuration** >**Backups** .
2. 
Select the **Restore Area** , such as`All` to store the full configuration.
3. 
Click **Choose File** to select the backup file.
4. 
You may select `Configuration file is encrypted.` and type the encryption password used during the backup operation.

**Figure 14.** *Restoring OPNsense Configuration Manually*

Full or incremental restores from backups are possible. Due to the interdependence of configuration parts, it is best to restore the whole configuration rather than just individual parts.

You may see the changes that have been made to your settings and even go back to a previous version by following the steps given below:

1. 
Navigate the **System** >**Configuration** >**History** menu.
2. 
Pick the older configuration using the left column of radio choices, then select the updated configuration using the right column.

**Figure 15.** *Viewing OPNsense Configuration Differences*

1. Click **View differences** button to see the changes between the two.

**Figure 16.** *Viewing OPNsense Configuration History*

Choosing the number of backups to preserve from this option is useful when a greater auditability standard is needed.

### How to Backup Configuration Automatically?

You can use a cronjob script to automate the backup process

1. Install the `os-api-backup` plugin.

**Figure 17.** *Installing os-api-backup plugin*

1. Go to **System** >**Access** >**Groups** in the OPNsense WebUI.

**Figure 18.** *Navigating System > Access > Groups*

1. 
Click **Add** button with`+` to create a new group with restricted rights.
2. 
Type a **Group name** , like backup.
3. 
Type a descriptive name in the **Description** field.

**Figure 19.** *Adding backup API Group*

1. 
Click **Save** .
2. 
Click **Edit** button with a pen icon for editing newly created group.
3. 
Click the **Edit** button with a pen icon on Assigned Privileges pane.

**Figure 20.** *Editing Assigned Priveleges of backup API Group*

1. Type `Backup API` in the search field and tick "Backup API" in the list. Leave the rest of the options unchecked.

**Figure 21.** *Saving Backup API System Privileges*

1. 
Click **Save** . This will save the privilege settings and take you back to the Group settings.
2. 
Clicking **Save** to activate the settings.
3. 
Go to **System** >**Access** >**Users** in the OPNsense WebUI to create a new user.
4. 
Click **Add** button with a`+` icon to add a user.
5. 
Set the **Username** , such as backup
6. 
Set a strong password.

**Figure 22.** *Adding Backup User*

1. Set group membership of the user by selecting the `backup` group.

**Figure 23.** *Setting group membership*

1. 
Click **Save & go back** button.
2. 
Edit the newly created backup user by clicking on the **Edit** button with a pen icon.
3. 
Scroll down to the **API keys** section.

**Figure 24.** *Adding Backup API Key*

1. 
Click **Add** button with a`+` icon to add a new key. This will create an api key and open an dialog box..
2. 
Click **Save** to download the tiny text file containing an API key and a secret, which you should keep somewhere safe.

**Figure 25.** *Downloading Backup API Key*

1. 
Click **Save and go back** .
2. 
Write the following script in a remote Linux or BSD-based machine using your preferred text editor:

```
#!/bin/bash
KEY="api_key"
SECRET="api_secret"
HOST="opnsense_hostname"
PATH="/path/to/backups"
curl -s -k -u $KEY:$SECRET https://$HOST/api/backup/backup/download \
-o $PATH/opnsense-config-$(date +%Y%m%d).xml
find $PATH/ -type f -name '*.xml' -mtime +30 -exec rm {} \;
```
Do not forget to update the variable values according to your settings.

This script will save backup files with the name "opnsense-config-yyyymmdd.xml" and delete anything older than 30 days.

You may try to use another advanced backup script given at `https://codeberg.org/SWEETGOOD/andersgood-opnsense-scripts/src/branch/main/backup-opnsense-via-api.sh` as well.

1. Set up a cron job to execute this on the schedule you like, and you're done.

## Configure OPNsense for High Availability

For hardware failover, OPNsense uses the Common Address Redundancy Protocol or CARP. A failover group is set up with two or more firewalls. If a main interface breaks or the primary goes offline totally, the secondary becomes operational. Using high-availability capability of OPNsense produces a completely redundant firewall with smooth failover. While switching to the backup network, customers experience minimal disruption to their network connections. You may easily configure high availability on your OPNsense.

## Disable root Access for the WebGUI

The default login after installing OPNsense is the root user. Because the root user has complete access to files and processes, it is typically not recommended to log in as `root`. Users of Linux, for example, are prompted to create a distinct user account after installation. The user can then use `sudo` to boost their privileges and carry out administrative activities. Theoretically, even if the user account is compromised, the root account is still safe (assuming there is no privilege escalation vulnerability being exploited or the password has been discovered). There is an option inside OpenSSH to prevent root access to the SSH server. It disables immediate log in as the root user as a security measure. OPNsense, which is based on FreeBSD (HardenedBSD, to be precise), does not deviate from this advice.

To disable the root access you may follow the next steps listed below:

1. 
Navigate to the **System** >**Settings** >**Administration** .
2. 
Uncheck `Permit User Login` option in the**Secure Shell** pane.

**Figure 26.** *Disabling root access*

## Restrict Access to Management Portal

Using a management network to gain administrative access to your network devices has numerous advantages:

- 
**Reliability:** When management traffic is separate from production or business traffic, management access can be preserved during production network reconfiguration because it won't have to compete for resources.
- 
**Simpler policies:** By separating the management from the production traffic, policies are more easily implemented. Policies with specific purposes are easier to understand and troubleshoot.
- 
**Security:** When network devices' management access is on a different network, accessing those devices on the production network is more challenging.

