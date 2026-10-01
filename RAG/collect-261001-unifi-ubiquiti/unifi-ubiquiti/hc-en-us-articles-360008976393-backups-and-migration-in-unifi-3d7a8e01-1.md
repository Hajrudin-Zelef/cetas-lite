---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-360008976393-backups-and-migration-in-unifi-3d7a8e01-1
title: "hc-en-us-articles-360008976393-backups-and-migration-in-unifi-3d7a8e01"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-360008976393-backups-and-migration-in-unifi-3d7a8e01.md
source_anchor: ""
source_lines: [1, 60]
sha256: 9e86dbf64fb6e10051e932925f79438aa529dd8311f3be92ebfb4ceb758cb6be
---

# hc-en-us-articles-360008976393-backups-and-migration-in-unifi-3d7a8e01

Backups and Migration in UniFi
UniFi simplifies backups and migration for all devices that host UniFi, including Cloud Gateways, CloudKeys, and Network Video Recorders (UNVR, ENVR). This ensures seamless transitions between devices while preserving settings and configurations.
UniFi provides both complete system backups and Network-specific backups. Most users will not need network-specific backups; they are advanced methods typically used by users with a Gen1 CloudKey, or those self-hosting the UniFi Network Server on a Windows, macOS, or Linux machine.
System Config Backups
System Config Backups include all OS, application, and device configurations. These backups can be created and restored in Settings > Control Plane > Console. This is the most common form of backups for UniFi Cloud Gateways and CloudKeys.
Automated Cloud Backups
The simplest way to ensure your data is backed up is to enable Automated System Backups:
- Go to Settings > Control Plane > Backups.
- Check System Backups.
If enabled, a cloud backup will be generated automatically each week and prior to each major update. Manual backups and restorations are managed in the Console's Control Plane.
Only the Cloud Gateway Owner has permissions to manage these backups, because they are saved directly to their UI Account. Visit account.ui.com/backups to view all backups associated with your account.
Offline Backups
For users who cannot use cloud backups, a local offline backup file can be used. You can download it in Settings > Control Plane > Backups.
Restoring From a Backup
Restoring from a backup allows you to revert to a previously stable configuration, ensuring continuity and minimizing disruptions. This is especially useful after a system failure, misconfiguration, or when setting up a new device.
To restore from a backup:
- Go to Settings > Control Plane > Backups and click Restore.
- Choose your most recent Cloud Backup or upload a backup file.
By default, “Restore All Applications and Settings” is enabled. To perform a partial restoration, uncheck it and select the specific apps, settings and admin configurations you need. This is especially useful when restoring to a different device—such as transferring Protect settings from a UDM Pro backup to a UNVR. Note that cloud backups can be downloaded from your UI Account Backups page.
Migration & Cloud Gateway Transfers
System Config Backups enable easy migration of devices and settings when moving from an old UniFi Cloud Gateway or console to a new one - such as when upgrading from an older or smaller device (e.g., UniFi Dream Machine Pro) to a larger or newer one (e.g., Enterprise Fortress Gateway).
This process is the only way to transfer device management from one Cloud Gateway, CloudKey or UNVR to another—other than performing a factory reset, which is generally undesirable.
To migrate backups to a new device:
- On the original device, go to Settings > Control Plane > Backups.
- Restore or Download a local backup.
- Replace the original Console or Cloud Gateway with the new one (downstream network connections should remain unchanged)
- Restore the backup when prompted during the setup process, or afterwards from Settings > Control Plane > Backups
Note: If migrating to a different model Cloud Gateway, CloudKey or UNVR, the new device must be running UniFi OS 3.1 or higher.
Migration Considerations: Migrating Protect
Migrating to Protect to a new UniFi Console follows the same procedure listed above. After restoring your configuration backup, all connected cameras will automatically begin recording on the new system. However, some key considerations apply:
- 
Existing recordings cannot be transferred between UniFi Consoles.
  - They remain accessible on the original Host, until you remove or reformat the storage.
- If you intend to reuse the hard drives in the new Console, they must be reformatted. This will remove any previously stored recordings.
- Cameras will appear as “adopted but offline” on the original Console, allowing access to old recordings while they are recording new footage on the new Console.
Network-Only Backups
All users with a dedicated UniFi Cloud Gateway or UniFi Console are encouraged to use the comprehensive System Config Backup. For rare cases, UniFi does offer Network-only Backups, which contain UniFi Network settings and device configurations.
These backups are generally only necessary for Self-Hosted UniFi Network Servers or other older installations.
These backups can be restored within UniFi OS in Settings > Control Plane > Backups.
Below are walkthroughs for some specific cases where Network backups would be necessary.
Migration from a UniFi Network Server to UniFi OS Server, Official Hosting, or a Cloud Gateway / Cloud Key
For users who cannot use cloud Network backups, a local offline Network backup file can be used.
This method is most applicable when replacing a self-hosted UniFi Network Server with a UniFi OS Server instance, Official Hosting, or when migrating from UniFI Network Server to a Cloud Gateway / Cloud Key with a single site.
- Ensure that your new Network Application is up to date.
- Download the Network Application backup file (*.unf) from your original Network Application (Settings > Control Plane > Backups). It is also recommended to copy the SSH username and password from Devices > Device Updates & Settings > Device SSH Settings, in case any devices need help later when connecting to the new instance of UniFi Network.
- Finish setting up the new UniFi OS Server instance, Console, or Cloud Gateway (leave all other UniFi devices connected as they were previously).
- Restore the backup file to the new Network Application in Settings > Control Plane > Backups.
- Wait for all devices to appear online. Once they appear online, shut down the old UniFI Network instance and/or select each device in the old instance and click "Forget" to ensure they do not reconnnect to the old console.
  - If any device appears offline or “Managed by Another Console,” try the below steps starting with (5) to migrate with Layer 3 Adoption. If that fails, factory reset the device and re-adopt it to the new console.
Migrating with Layer 3 Adoption
Layer 3 adoption is often used when managing UniFi devices at a remote location, using DNS or DHCP entries. If you have previously adopted a device to a self-hosted UniFi Network Server or CloudKey using DNS, DHCP Option 43, or Layer 3 adoption commands to connect across VLANs or the internet, follow these steps:
- Ensure that your new Network Application is up to date.
- Download the Network Application backup file (*.unf) from your original Network Application (Settings > Control Plane > Backups).
- Connect the new Network host, Console, or Cloud Gateway (leave all other UniFi devices connected as they were previously) and remove the old one.
- Restore the backup file to the new Network Application in Settings > Control Plane > Backups.
- Once restoration is complete, enable the Override Inform Host toggle on the original Network Application and enter the IP address of the new console. This tells your devices where they should connect.
  - If migrating to Official UniFi Hosting, use the Inform URL from the Network Application dashboard. Remove the initial http:// and the ending :8080/inform when pasting. Remember to update any DNS entries at your other sites.
- Wait for all devices to appear as online in the new Network Application. If a device appears offline or 'Managed by Another Console,' factory reset and re-adopt it.
- Forget the migrated devices from the original Network Application through the devices' settings and shut down your original UniFi Network instance or host.
Exporting Individual Sites from Older UniFi Hosting Options
