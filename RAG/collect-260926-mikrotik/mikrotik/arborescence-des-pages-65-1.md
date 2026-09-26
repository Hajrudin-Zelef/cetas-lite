---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-65-1
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "nand", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-65.md
source_anchor: ""
source_lines: [1, 125]
sha256: 18d1e00838041e1c82121652413fbe66c5e76c9e4ab01b0600c6edbf99daf38e
---

# Summary

**Sub-menu:** `/disk`

This menu will list all attached storage devices, presuming that they are supported and in working condition. This is especially useful for RouterBOARD devices with SD/CF/USB/SATA/NVMe slots and x86 systems with additional dedicated storage drives - as the built-in storage is quite small, an external drive comes in very handy when you want a big User Manager database, proxy cache or possibly SMB shares on your router.

You can add as many external or secondary drives as you want, and select any number of them for each of the mentioned feature usages. For example, User Manager could be used on 3 disks, one of them would be the active database, and the rest would be backups. You can then add a fourth disk, copy the active data to it - unmount - unplug it - and move to another server, to keep using the actual database. This means migration and backup are made easy!

Disks carry names where they are physically connected.

**Always use `/disk eject` before physically removing any disks from your RouterOS device to prevent data loss!**

# Properties

| Property | Description | 
|---|---|
| **eject** () | Safely unmounts (ejects) drive of your selection by using "slot" that is assigned to it. After issuing this command it can be removed from host device. | 
| **format** () | Command to initiate disk formatting process. Contains additional properties of its own. Such as "file-system" and "label".  | 
| **trim** | Discards the unused data blocks for performance (fstrim equivalent). Some NVMe enclosures might not support disk trimming. | 
| **reset-counters**  | Resets disk (slot) statistics | 
| **monitor-traffic**  | Check real time disk performance and health stats | 
| **test**  | allows performing performance tests of selected device (Available from RouterOS 7.16)  | 
| **mount-read-only** | Sets the mounted disk in read only mode when set to *yes* . | 
| **mount-point-template** | Sets the mounting point for the file system. It is possible to set the mount point as the following parameters based on the disk:  Additionally, it is possible to combine multiple variables to create a single mount point: | 

# Flags

| Property | Description | 
|---|---|
| **X - disabled** | Disabled device | 
| **E - empty** | Empty slot | 
| **B - BLOCK-DEVICE** | The "B - BLOCK-DEVICE"- Flag means that this device works using blocks for input/output operations. In the context of RouterOS, its distinction is crucial, as it helps determine whether a device is functioning as a data carrier or simply providing information about the disk layout structure. This difference becomes important when considering the extender with the device behind it. If a device is marked with the letter "B", this indicates its ability to be used as storage or memory. In contrast, devices that do not have a "B" mark are designed primarily to understand the structure of the disk. This allows to quickly recognize the presence of a PCIe or SAS expander, as well as detect the presence of drives in the first expander. In addition, it allows you to estimate the speed of the connection to which each device is connected. However, the most notable benefit of the "B" flag is its ability to instantly indicate whether a device can be formatted or used for RAID purposes. | 
| **M - mounted** | Mounted partition | 
| **F - formatting** | The device is currently in the formatting process | 
| **p - partition** | The device has a partition | 
| **f - raid-member-failed** | These options are used with the ROSE package. | 
| **r - raid-member** |  | 
| **c - encrypted** |  | 
| **g - guid-partition-table** |  | 
| **t - nvme-tcp-export** |  | 
| **i - iscsi-export** |  | 
| **s - smb-export** |  | 
| **n - nfs-export** |  | 
| **O - tcg-opal-self-encryption-enabled** |  | 
| **o - tcg-opal-self-encryption-supported** |  | 

# Settings

| Property | Description | 
|---|---|
| **auto-smb-sharing** (yes \| no; Default: ) | Enables dynamic SMB shares when new disk/partition item is added in "/disk" | 
| **auto-smb-user** (list of strings; Default: )  | Default value for smb-sharing/smb-user setting, when new disk/partition item is added in "/disk" | 
| **auto-media-share** (yes \| no; Default: ) | Enables DLNA dynamically when new disk/partition item is added in "/disk" | 
| **auto-media-interface** (list of strings; Default: ) | Interface that will be used in dynamic instance for ip/media when new disk/partition item is added in "/disk" | 
| **default-mount-point-template** (string, Default: ) | Sets the default mount point template for each item added in "/disk" | 

Notes

With "auto-smb-sharing=yes" and "/ip smb share enabled=auto" SMB server gets enabled when a storage device is physically plugged in

# Examples

## Formatting attached storage unit - Simple

1. Disk is attached, and already mounted automatically by the system.

2. Formatting the disk, in either of two supported file-systems (ext4 or fat32).

3. It's done! Drive is formatted and should be automatically mounted after formatting process is finished.

## Formatting attached storage unit - Detailed

Let us presume that you have added a storage device to your device that is running RouterOS. System will try to automatically mount it and in such case if storage is formatted in a supported file-system and partition record, it will be found in "/files" menu moments after you plugged it in to the host device.

If not, here is what you have to do.

1. Do a quick print of disk menu, to make sure that router sees the attached storage.

We can here see that system sees one storage drive and also that it is formatted with a known file-system type.

When running file menu print-out we also see that is mounted.

2. To formatting drive - we issue command with previously know id or name(slot) and with desired file-system (ext4 or fat32), we can also assign label to device as I did in this example and make mbr partition table

**Note:** In printout, you can see that there is a progress percentage counter in formatting process. For larger storage drives, it might take longer for this process to finish, so be patient.

## Creating multiple disk partitions

If multiple GPT partitions are needed format drive without partition table and add them manually:

**Note:** Slot (partition or disk name) is assumed automatically, but can be overwritten by using slot parameter. 

If partition size is not used all available space will be used from last partition.

To offset partition start "partition-offset" parameter can be used.

## Web-Proxy cache configuration example

Enter proxy cache path under IP -> Proxy menu and web proxy store is automatically created in files menu. If a non-existent directory path is used, an additional sub-directory is also created automatically.

## Log on disk configuration example

When configuring logging on disk make sure that you create directories in which you want to store the log files manually, as non-existent directories will NOT be automatically created in this case.

**Note:** Logging topics such as firewall, web-proxy and some other topics that tend to save a large amount or rapid printing of logs on system NAND disk might cause it to wear out faster, so using some attached storage or remote logging is recommended in this case or save data in RAM folder

## Allocate RAM to folder

It is possible to add folders linked to RAM. Folders will be emptied on reboot or power loss.

RAM will be filled up to tmpfs-max-size and if this variable in not provided - up to 1/2 from available RAM.

## Test disk performance


Disk performance tests may slowly degrade disk health

On write tests all files and file systems on disks will be destroyed

Starting from 7.16 to run disk performance tests. Disks has to be disabled or without mountable file system (unformatted).

Check available disks, if disk is already mounted - disable it.

## Swap space

