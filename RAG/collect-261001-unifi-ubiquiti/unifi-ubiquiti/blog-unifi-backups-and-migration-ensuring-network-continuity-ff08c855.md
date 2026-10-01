---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-backups-and-migration-ensuring-network-continuity-ff08c855
title: "blog-unifi-backups-and-migration-ensuring-network-continuity-ff08c855"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-backups-and-migration-ensuring-network-continuity-ff08c855.md
source_anchor: ""
source_lines: [1, 116]
sha256: d168e627804de76fa6aab4e1a634edb3354c8621e191f25cc636bf65a7c51309
---

# blog-unifi-backups-and-migration-ensuring-network-continuity-ff08c855

Hello tech enthusiasts! Juan David here, your Tech Lead Support and Ubiquiti Certified Trainer at Flytec. Today, we’re diving deep into the world of **UniFi backups and migration**, essential processes for maintaining a reliable network infrastructure. Whether you're upgrading your UniFi Console, migrating to a new environment, or just securing your configurations, this guide has you covered.

UniFi offers two primary backup types:

1. **System Config Backups** : Comprehensive backups for dedicated UniFi Consoles like UniFi Dream Machines (UDMs), CloudKeys (CKs), and UniFi Network Video Recorders (UNVRs).
2. **Network Application Backups** : For Gen1 CloudKey users or those self-hosting the UniFi Network Server on Windows/macOS/Linux machines. These backups save UniFi Network settings and device configurations and can be created in the Backups section of the Network Settings.

**Automated Cloud Backups**

If **Remote Management** is enabled, UniFi simplifies the process by generating **automatic cloud backups** weekly and before major updates. These backups are securely stored in your **UI Account**, accessible at account.ui.com/backups.

**Key Features:**

- **Cloud Backups** : Automatically created, ensuring you always have an up-to-date configuration saved.
- **Manual Backups** : Manage these via the Console’s Control Plane for additional control.

**UniFi Console Migration**

Moving to a new UniFi Console? Backups are the key to a seamless migration. Here’s how to do it:

**Steps for Migration:**

1. **Create a Backup** :
  - Generate a cloud backup or download a local backup file from the Control Plane.
2. **Replace the Console** :
  - Swap the original UniFi Console with the new one. Keep all network connections unchanged.
3. **Set Up the New Console** :
  - Configure the new console as instructed during the setup process.
4. **Restore the Backup** :
  - Use the backup file during setup or later via the Control Plane.

**Requirements**:

- The new console must run **UniFi OS 3.1 or higher** .
- Both consoles must be of the same model for optimal compatibility.

**Advanced Scenarios**

UniFi Network Video Recorders (UNVR) can be migrated to the Enterprise Network Video Recorder (ENVR). After restoring the backup, connected cameras will automatically record on the new system. However, note:

- **Recording Retention** : Existing recordings cannot be migrated but remain accessible on the original UNVR.
- **HDD Reformatting** : Reformat the HDDs on the ENVR for new recordings.

**Migrating from UniFi Cloud Hosting to Local Hosting**

Some users may want to switch from UniFi's Official Hosting to self-hosted local infrastructure for better control:

1. **Download the Backup** : Access your UniFi Cloud account and download the latest backup.
2. **Set Up the Local Server** : Install the UniFi Network Application on your local hardware (e.g., Windows, macOS, or Linux server).
3. **Restore the Backup** : Upload the backup file to the self-hosted application.
4. **Redirect Inform Addresses** : Update your devices' inform addresses to point to the local server.

*This scenario is ideal for users who want to eliminate recurring cloud hosting costs while retaining full control over their network.*

**Migration from Self-Hosting to a UniFi Console**

If you're currently self-hosting the UniFi Network application and want to migrate to a UniFi Console, follow these steps:

- 1.Create a backup file from your self-hosted UniFi Network Server.
- 2.Set up the new UniFi Console and ensure it's running the same UniFi Network version as your backup.
- 3.Upload the backup file to the new console via Settings > System > Restore Backup.
- 4.Re-adopt any devices as necessary and verify the configurations.

*This migration provides an opportunity to centralize management and benefit from the enhanced features of a UniFi Console.*

**Migration Methods**

**Standard Migration**

Best for users managing devices on the same network/VLAN as their UniFi Console.

1. Update the new Network Application.
2. Download the UniFi backup file from the original application.
3. Connect the new UniFi Console.
4. Restore the backup file.
5. Wait for devices to appear online in the new application.

**Note**: If devices are "Managed by Another Console," factory reset them and re-adopt them into the new application.

**Layer 3 Device Migration**

For devices on different networks/VLANs:

1. Follow steps 1–4 from the Standard Migration process.
2. Enable the **Override Inform Host** option in the original application. Enter the IP address of the new console to redirect devices.
3. Monitor devices as they reconnect to the new application.

**Multi-Site Migration**

For UniFi Consoles managing multiple sites, you can migrate individual sites to a new console. Here’s how:

1. Export the site from the original application via **Settings > System > Site Management** .
2. Import the site to the new application using the **Site Switcher** .
3. Complete the guided walkthrough to migrate devices to the new site.

**Tip**: Enable **Multi-Site Management** in the settings if the option isn’t visible.

**FAQ: Your Backup and Migration Questions Answered**

**Q: What happens during a failover?****A:** During a failover, the shadow gateway (if configured) automatically takes over, minimizing downtime.

**Q: Can I migrate recording data between UNVR and ENVR?****A:** No, recordings cannot be transferred. However, they remain accessible on the original UNVR.

**Q: What if devices appear offline after migration?****A:** Factory reset the devices and re-adopt them into the new console.

**Q: Do I need multiple IP addresses for migration?****A:** No, a single IP address is sufficient.

**Final Thoughts**

Backing up and migrating your UniFi infrastructure ensures continuity, minimizes downtime, and protects your hard work. Whether you're upgrading your hardware or simply safeguarding your settings, these processes are straightforward with UniFi's intuitive tools.

Need help with your UniFi setup or migration? Flytec Computers has you covered with expert advice and support.

Stay connected,

**Juan David**, your Tech Lead Support at Flytec!
