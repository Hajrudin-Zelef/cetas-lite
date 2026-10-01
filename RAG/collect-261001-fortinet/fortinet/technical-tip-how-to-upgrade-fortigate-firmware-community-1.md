---
id: collect-261001-fortinet/fortinet/technical-tip-how-to-upgrade-fortigate-firmware-community-1
title: "technical-tip-how-to-upgrade-fortigate-firmware-community"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/technical-tip-how-to-upgrade-fortigate-firmware-community.md
source_anchor: ""
source_lines: [1, 152]
sha256: a56f1b560ff5cb0cb5967c150dcd366eebda60f0cb98b0dca1cb3a119de855f1
---

# technical-tip-how-to-upgrade-fortigate-firmware-community

**Description**


This article describes how to upgrade FortiGate firmware, FortiGate administrators whose access profiles contain system configuration read and write privileges, and the FortiGate admin user that can change the FortiGate firmware.

Download the most recent firmware build from the Fortinet Technical Support website at Welcome to Fortinet Support.


**Scope**


FortiGate.


**Solution**


**Usage awareness and preparation checklist before the upgrade:**

1. Firmware Images Checksums Link: Download locally and verify the hash of the Firmware that will be flashed (Always upgrade the firmware from a local copy).
2. Changes should be performed during a maintenance window to avoid any unexpected problems.
3. If the plan is to upgrade, double-check the Upgrade path tool table and release notes concerning the model and the version.
4. Make sure to perform backups before starting the upgrade process. Refer to this article: Technical Tip: How to download FortiGate configuration file & Debug log from GUI.
5. Perform the backup using a super_admin account only; otherwise, the backup may not include the super_admin profile and users during restore. Refer to this article: Technical Tip: Prof_Admin admin profile will not be able to back up the Super_Admin.
6. Verify the Release Notes for the intended version and evaluate **special notices, upgrade information,** product integration and support,**resolved issues, known issues, and limitations** of the version.
7. Make sure to have physical access to the device during the upgrade, and do not perform upgrade procedures remotely or through FortiGate Cloud.
8. It is necessary to have a Console Cable (tested and connected) to access the unit through the console for any unexpected problems. See this article: Technical Tip: How to connect to the FortiGate console port.
9. Verify the compatibility with the installed FortiManager and FortiAnalyzer. This must be upgraded before the FortiGate can be fully compatible. The upgrade path is: **FortiManager -> FortiAnalyzer -> FortiGate** .

**v5.2.x and v5.4.x:**

To upgrade the firmware:

1. Log in to the web-based manager as the administrative user.
2. Go to **System -> Dashboard -> Status** and locate the System Information widget.
3. Besides **Firmware Version** , select Update.
4. On the next screen, select the 'Browse' or 'Upload Firmware' button.
5. Locate the file on the local computer and select the firmware image file.
6. Select the 'Backup config and upgrade' button to back up the configuration and start a firmware upgrade.


The FortiGate unit uploads the firmware image file, upgrades to the new firmware version, restarts, and displays the FortiGate login. This process takes a few minutes.


**v5.6.x, v6.0.x and v6.2.x:**

1. **Log in to the web-based manager as the administrative user.**
2. **Go to System -> Firmware -> Select the 'Browse' button** to locate the firmware image file.
3. Locate the file on the local computer and select the firmware image file.
4. Select the 'Backup config and upgrade' button to back up the configuration and start a firmware upgrade.
5. The FortiGate unit uploads the firmware image file, upgrades to the new firmware version, restarts, and displays the FortiGate login. This process takes a few minutes.

**v7.0.x:**

1. Log in to the web-based manager as the admin user.
2. Go to **System -> Firmware,** and there will be 4 tabs:**Latest, All Upgrades, All Downgrades, and File Upload** .
3. Select the option File  **upload** , and select the**Browse** button to locate the firmware image file.
4. Locate the file on the local computer and select the firmware image file.
5. Select the 'Backup config and upgrade' button to back up the configuration and start a firmware upgrade.
6. The FortiGate unit uploads the firmware image file, upgrades to the new firmware version, restarts, and displays the FortiGate login. This process takes a few minutes.
7. An alternative process is to go to **System -> Fabric Management** , select the FortiGate, and select the**Upgrade** option. Similarly, 4 tabs will appear: follow steps 3 to 6.

**v7.2.x, and v7.4.x:**

1. Log in to the web-based manager as the admin user.
2. Go to **System -> Fabric Management,**  select the FortiGate, and then select the option Upgrade. 4 tabs will appear:**Latest, All Upgrades, All Downgrades, and File Upload** .


**Note**: Starting from version 7.4.x, **System -> Fabric Management** has been replaced with **System -> Firmware and Registration**.


1. Select the **File Upload** option and select the**Browse** button to locate the firmware image file.
2. Locate the file on the local computer and select the firmware image file.
3. Select the 'Backup config and upgrade' button to back up the configuration and start the firmware upgrade.
4. The FortiGate unit uploads the firmware image file, upgrades to the new firmware version, restarts, and displays the FortiGate login. This process takes a few minutes.


**v7.6.x:**

1. Log in to the web-based manager as the admin user.
2. Go to **System -> Firmware and Registration,**  select the FortiGate, and then select the option Upgrade.



1. Select Two Options will appear ' **Full Fabric Upgrade'** and '**FortiGate only** ', and select the appropriate option.


1. In the example, ' **FortiGate only'** has been selected, and then select 'Next'.


**Note**: Even though the option reads '**FortiGate only – Upgrade the selected FortiGate only**', selecting this option in an HA setup actually upgrades the **FortiGate HA cluster**, not just the individual unit.


1. At the Next window, **Recommended, All Upgrades, All Downgrades, and File Upload** .
2. In the example, the 'All Upgrades' option has been selected.



1. Select the ' **Select'** option and proceed to the next Section.
2. In this section, select 'Immediate', or a scheduled update can be selected as well using the ' **Specify'** option.



1. Choose ' **Specify'** if the upgrade needs to be performed automatically at a later time.



1. Proceed to the ' **Review'** section, '**Confirm and Backup Config',**  and the Firewall will be upgraded to the target version.



**Note:**Some users on v7.6.2 may encounter an issue where uploading the firmware manually via File Upload may see a browser out of memory issue. In those cases, consider using the All Upgrades option instead to download firmware from FortiGuard servers:

**Upgrading firmware through the built-in FortiGate Management Station:**

`execute restore image management-station ?`

It will bring up a list similar to the one below:

```
07004000FIMG0026804009 v7.4 MR4-GA-M P9 b2829 (upgrade)      
07004000FIMG0026804007 v7.4 MR4-GA-M P7 b2731 (downgrade)
07004000FIMG0026804006 v7.4 MR4-GA-M P6 b2726 (downgrade)
07004000FIMG0026804005 v7.4 MR4-GA-M P5 b2702 (downgrade)
07004000FIMG0026804004 v7.4 MR4-GA-F P4 b2662 (downgrade)  
```
```
execute restore image management-station 07004000FIMG0026804009
Getting image 07004000FIMG0026804009 from Management station...
####################################################################################################
This operation will replace the current firmware version!
Do you want to continue? (y/n)y
```
**Upgrading the firmware through the CLI.**

Before starting, ensure a TFTP server is running and accessible to the FortiGate unit.

- **Step 1** : Copy the new firmware image file to the root directory of the TFTP server.
- **Step****2** : Log in to the CLI.
- **Step 3** : Make sure the FortiGate can connect to the TFTP server. Use the following command to ping the computer running the TFTP server. For example, if the IP address of the TFTP server is 192.168.1.168:


`execute ping 192.168.1.168`

Enter the following command to copy the firmware image from the TFTP server to the FortiGate unit:


`execute restore image tftp <filename> <tftp_ipv4>`
 

The FortiGate unit responds with the message:

