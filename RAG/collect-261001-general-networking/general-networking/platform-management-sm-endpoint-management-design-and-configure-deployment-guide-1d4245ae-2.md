---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-design-and-configure-deployment-guide-1d4245ae-2
title: "platform-management-sm-endpoint-management-design-and-configure-deployment-guide-1d4245ae"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-design-and-configure-deployment-guide-1d4245ae.md
source_anchor: ""
source_lines: [33, 88]
sha256: 77b4ffac7e8d10c2f54326126383e92e1f6fd18fbb04292d3ad32f883db6a676
---

# platform-management-sm-endpoint-management-design-and-configure-deployment-guide-1d4245ae

Note: Consider DEP settings which "Skip" the option "Restore from backup", as iCloud backups that were captured while the device was enrolled into a previous MDM may attempt to restore the device back into the previous MDM and cause issues. It is recommended to set up iOS devices as new devices and, if desired, sign into iCloud after the device sets up to restore the iCloud data (such as iCloud photos, iCloud app data, etc). That way you can still restore the important iCloud backed up data without restoring the entire device image.
Unenrolling from a Previous MDM Provider
macOS Devices not enrolled in DEP
- Option 1:
    
  - Login to your previous MDM provider's portal and send an Erase Device command to the devices. This will factory erase and unenroll them.
- Option 2:
    
  - Restore macOS to factory settings to unenroll the device from the previous MDM provider.
- Option 3 (no factory erase):
    
  - Take devices 1-by-1 and go to System Preferences > Profiles and find the management profile for the previous MDM provider. Delete it and the device is now unenrolled from the previous MDM provider. If the previous MDM provider has an agent program running on macOS you should also uninstall this agent application. View your previous MDM provider's documentation for more information on uninstalling their agent. If you are unsure, a device factory erase is always a good option to confirm the previous MDM provider's agent is removed.
macOS Devices enrolled in DEP
- Option 1:
    
  - Login to your previous MDM provider's portal and send an Erase Device command to the devices. This will factory erase and unenroll them. When devices run through the macOS Setup Assistant, they will automatically be enrolled into Meraki Systems Manager via the DEP settings assigned to the device.
- Option 2:
    
  - Restore macOS to factory settings to unenroll the device from the previous MDM provider. When devices set up again through the macOS Setup Assistant they will automatically be enrolled into Meraki Systems Manager via the DEP settings assigned to the device.
- Option 3 (no factory erase):
    
  - If you would prefer not to factory erase the macOS devices follow the "not in DEP" options above. However, it is recommended to factory erase when using DEP so the DEP settings can become "Pushed" and applied to the device.
Note: A factory erase is required for the devices to receive your organization's assigned DEP settings.
iOS devices not in DEP
- Option 1:
    
  - Login to your previous MDM provider's portal and send an Erase Device command to the devices. This will factory erase and unenroll them.
- Option 2:
    
  - Factory erase device via Apple Configurator (and then enroll devices using the Apple Configurator Manual Enrollment steps)
- Option 3:
    
  - Factory erase device via iTunes.
- Option 4:
    
  - Take devices 1-by-1 and go to Settings > General > Reset > Erase all content and settings.
- Option 5 (no factory erase):
    
  - Take devices 1-by-1 and go to Settings > General > Profile & Device Management and remove the previous MDM's enrollment profile. This requires that the MDM profile was not set as unremovable.
iOS devices in DEP
- Option 1:
    
  - Login to your previous MDM provider's portal and send an Erase Device command to the devices. This will factory erase them so they are now unenrolled. DEP is the most efficient method because devices will automatically enroll into Meraki Systems Manager during the initial iOS Setup Assistant.
- Option 2 (no factory erase):
    
  - If you would prefer not to factory erase the iOS devices, follow the "not in DEP" options above. However, it is recommended to factory erase when using DEP so the DEP settings can become "Pushed" and applied to the device.
Note: A factory erase is required for the devices to receive your organization's assigned DEP settings.
Note: If the previous MDM provider does not support Activation Lock Bypass, be sure that end users sign out of their Find My iPhone/iPad iCloud accounts prior to the factory erase/unenroll actions so devices aren't locked to their iCloud account.
Enrolling into Meraki Systems Manager
Now that the migration has been set up, and the devices have been removed from the previous MDM provider's management: it is time to enroll into Meraki Systems Manager. Follow the guides and videos below for your specific device's enrollment steps. If using DEP: devices will automatically enroll upon their initial iOS/macOS Setup Assistant.
For more information, refer to the following guides:
After the Migration
Shortly after migration, you can navigate to Systems Manager > Monitor > Client list and view your enrolled devices. Compare this list with the inventory exported from the previous MDM provider to ensure consistency. Be sure to scope apps that your end users need and install any configuration profiles. You can restore the data backed up from the previous MDM provider as well, such as:
If using Apple IDs/iCloud to backup end user's information, this is a good time to have end users sign into their iCloud or Managed Apple ID to restore any cloud-hosted app data, photos, etc.
Note: iCloud backups for iOS devices may pose a problem when migrating from a previous MDM provider into Systems Manager, as the iCloud backup would also be capturing the device's enrollment from the previous MDM provider. It is recommended to set up iOS devices as new devices and, if desired, sign into iCloud after the device sets up to sync iCloud data (such as iCloud photos, iCloud app data, etc). That way you can restore iCloud data without restoring the entire device image.
Note: Apple does not support a way to sync iCloud after the initial setup. This can only be performed by 3rd party apps.
