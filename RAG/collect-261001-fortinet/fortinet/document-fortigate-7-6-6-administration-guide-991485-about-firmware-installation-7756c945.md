---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-6-administration-guide-991485-about-firmware-installation-7756c945
title: "document-fortigate-7-6-6-administration-guide-991485-about-firmware-installation-7756c945"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-6-administration-guide-991485-about-firmware-installation-7756c945.md
source_anchor: ""
source_lines: [1, 23]
sha256: 3d222fa435c178073f9612e3f7f126f796bcbead8c088e1199a085b888758e1d
---

# document-fortigate-7-6-6-administration-guide-991485-about-firmware-installation-7756c945

About firmware installations
About firmware installations
Fortinet periodically releases new FortiGate firmware to include new features and resolve important issues. After successful registration of your FortiGate unit, firmware updates are available from FortiGuard and from the Fortinet Customer Service & Support website. See Registering FortiGate for more information.
|  | If the Firmware & General Updates (FMWR) license is invalid or the FortiGate appliance has reached End of Support (EOS), the FortiGate will enforce an automatic update to the latest patch of the current minor version. A valid FMWR license is required to update to the next minor or major version; without it, the FortiGate cannot be upgraded to these versions. See Required firmware upgrades for FortiGate appliances with invalid support contracts or that have reached EOES. | 
Installing a new firmware image replaces the current antivirus and attack definitions, along with the definitions included with the firmware release that is being installing. After you install new firmware, make sure that the antivirus and attack definitions are up to date.
|  | It is recommended to back up your configuration before making any firmware changes. You will be prompted to back up your configuration as part of the upgrade process. See also Configuration backups and reset. | 
Before you install any new firmware, follow the below steps:
- 
                                                    Understand the maturity level of the current and target firmware releases to help you determine whether to upgrade. See Firmware maturity levels. See also Selected availability (SA) versions.
- 
                                                    Review the Release Notes for a new firmware release.
- 
                                                    Review the Supported Upgrade Paths.
- 
                                                    Download a copy of the currently installed firmware, in case you need to revert to it. See Downloading a firmware image and Downgrading individual device firmware for details.
- 
                                                    Have a plan in place in case there is a critical failure, such as the FortiGate not coming back online after the update. This could include having console access to the device (Connecting to the CLI), ensuring that your TFTP server is working (Installing firmware from system reboot), and preparing a USB drive (Restoring from a USB drive).
- 
                                                    Back up the current configuration, including local certificates. The upgrade process prompts you to back up the current configuration. See also Configuration backups and reset for details.
- 
                                                    Test the new firmware until you are satisfied that it applies to your configuration. See Testing a firmware version and Using controlled upgrades for details.
Installing new firmware without reviewing release notes or testing the firmware may result in changes to settings and unexpected issues.
|  | Only FortiGate admin users and administrators whose access profiles contain system read and write privileges can change the FortiGate firmware. |
