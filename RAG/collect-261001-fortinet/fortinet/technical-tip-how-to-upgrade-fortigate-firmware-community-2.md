---
id: collect-261001-fortinet/fortinet/technical-tip-how-to-upgrade-fortigate-firmware-community-2
title: "technical-tip-how-to-upgrade-fortigate-firmware-community"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-fortinet/technical-tip-how-to-upgrade-fortigate-firmware-community.md
source_anchor: ""
source_lines: [153, 269]
sha256: 22e6bc88532c13741784c5fec0631e3233b7111df557131fccf19514437e0c96
---

# technical-tip-how-to-upgrade-fortigate-firmware-community

```
This operation will replace the current firmware version!
Do you want to continue? (y/n)
```

Type y. The FortiGate unit will upload the firmware image file, upgrade to the new firmware version, and restart. This process takes a few minutes and reconnects to the CLI.


**Updating the firmware on FortiGate. Browse to support.fortinet.com and log in.**

1. Go to **Downloads -> Firmware Images -> FortiGate -> Vr _ -> MR_ -> Patch _** and view the list for the image file matching the device model.
2. Back up the FortiGate Config by going to the menu tabs on the left of the interface window. 
  1. Go to **System -> Dashboard -> Status -> System Information -> System Config -> Backup** .
  2. Select 'Backup' and allow the browser to save the file to a secure location.
  3. Load the firmware and reboot by going to the menu tabs on the left of the interface window.
  4. Go to **System -> Dashboard -> Status -> System Information -> Firmware Version -> Update** .
  5. In the 'Upgrade From' field, choose 'Local Hard Disk'.
  6. Browse to the location of the saved firmware, downloaded in step 2, by pressing the 'Browse' button.
  7. Take note of the 'Upgrade Partition' (this cannot be altered here).
  8. To boot the FortiGate firmware, ensure that the 'Boot the New Firmware' box is selected. This option is not available on earlier firmware.
  9. Press **OK** . The FortiGate will reboot.


**Upgrading from the Details window:**

Load the firmware and reboot by going to the menu tabs on the left of the interface window. Go to **System -> Dashboard -> Status -> System Information -> Firmware Version -> Details**.

- Select the partition to upload the firmware (it is best practice to select the non-active partition for fallback reasons).
- Select **Upload** at the top.
- In the 'Upgrade From' field, choose 'Local Hard Disk'.
- Browse to the location of the saved firmware downloaded in step 2 above by pressing the 'Browse button'.
- Take note of the 'Upgrade Partition' (this cannot be altered here).
- To boot the firmware, ensure that the 'Boot the New Firmware' box is selected. This option is not available on earlier firmware.
- If it is not desirable to boot immediately to the new firmware, deselect the 'Boot the New Firmware' box.
- Press **OK** .


The FortiGate will reboot.


**Upload and Boot to Firmware at a later time or Boot to Previous Firmware.**

Loading the other partition can be useful to downgrade quickly to the previous working firmware.

However, there are a few considerations to be aware of:

- This is not available on Virtual Machines.
- In the case of an HA cluster, this will not be synchronized, so the commands must be run on each member individually.


In the CLI, use the following commands.

To list partitions and check if they are active:


`diagnose sys flash list`   

To indicate what partition to boot from the next time the device reboots (Partition 1 is the primary and Partition 2 is the secondary):


`execute set-next-reboot <primary|secondary>`

To reboot the FortiGate:


`execute reboot`

If the FortiGuard License for **Firmware and General Updates** is renewed and is not reflected on FortiGate, then execute the following command to synchronize license information with the portal:


`execute update-now`

If the device is in an HA cluster, both devices should have a valid license to fetch firmware upgrades from FortiGuard.

Once the command is executed, if firmware updates are available, it can be seen on the device, or the top-right notification count will be highlighted.


**To debug the upgrade process:**


Run these commands in console connection to debug the upgrade process:


```
diagnose debug duration 0
diagnose debug kernel level 7
diagnose debug console timestamp enable
diagnose debug enable
```

It will show the details of the upgrade process and any errors/failures if any.

**Note**: 

When the FortiGate is about to reboot for the firmware upgrade, FortiGate will send an SNMP shutdown event.

```
Image verification OK!
snmpd: received shutdown signal
snmpd: shutdown
```

**Note**: From v.6.2, with models 40F, 60F, 70F, 80F, and 100F (variants) supports sharing a single license for both clusters: Single FortiGuard license for FortiGate A-P HA cluster.


More information on loading the secondary partition can be found in the following articles:


**Compatibility note:**It is necessary to perform any possible upgrade to compatible versions of other Fortinet products while considering a FortiGate upgrade.

- Technical Tip: Compatibility Matrix guide between FortiGate and other Fortinet products
- Compatibility Tool (for FortiManager)
- FortiAP and FortiOS 7.x Compatibility Matrix
- FortiLink Compatibility


**Related article:**
