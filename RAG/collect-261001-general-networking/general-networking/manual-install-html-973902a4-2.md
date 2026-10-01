---
id: collect-261001-general-networking/general-networking/manual-install-html-973902a4-2
title: "manual-install-html-973902a4"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "memory"]
source: docs/RAG/collect-261001-general-networking/manual-install-html-973902a4.md
source_anchor: ""
source_lines: [83, 172]
sha256: 6e609a3ab5a832c77cc929f4bab2fc966b33eb0bd10e3f528a253b916a0d9d03
---

# manual-install-html-973902a4

sudo dd if=OPNsense-##.#.##-[Type]-[Architecture].img of=/dev/rdiskX bs=64k
where r = raw device, and where X = the disk device number of your CF card (check Disk Utility) (ignore the warning about trailing garbage - it’s because of the digital signature)
Windows
physdiskwrite -u OPNsense-##.#.##-[Type]-[Architecture].img
(use v0.3 or later!)
System Boot Preparation
After preparing the installation media, we need to make sure we can access the console (either via keyboard and [virtual]monitor or serial connectivity). Next we need to know how to access the boot menu or the system BIOS (UEFI) to boot from the installation media. In most cases, it will be a function (F#), Del, or ESC key that needs to pressed immediately after powering on (or rebooting) the system. Usually within the first 2 to 3 seconds from powering up.
Tip
OPNsense devices from the OPNsense shop use <ESC> to enter the BIOS and boot selection
options.
Note
Serial connectivity settings for DECXXXX devices can be found here
Installation Instructions
Install Instructions
OPNsense installation boot process allows us to run several optional configuration steps. The boot process was designed to always boot into the live environment, allowing us to access the GUI or even SSH directly. If a timeout was missed, restart the boot procedure.
OPNsense Importer
All Full Images have the OPNsense Importer feature that offers flexibility in recovering failed firewalls, testing new releases without overwriting the current installation by running the new version in memory with the existing configuration or migrating configurations to new hardware installations. Using Importer is slightly different between previous installs with existing configurations on disk vs new installations/migrations.
For systems that have OPNsense installed and the configuration intact, here is the process:
- Boot the system with installation media
- Press any key when you see “Press any key to start the configuration importer”. 
  - If you see the login prompt, you have passed the importer and will need to reboot.
- Type the device name of the existing drive that contains the configuration and press enter.
- If Importer is successful, the boot process will continue into the live environment using the stored configuration on disk.
- If Importer was unsuccessful, we will be returned to the device selection prompt. Confirm the device name is correct and try again. Otherwise, there may be possible disk corruption and restoring from backup.
At this point, the system will boot up with a fully functional firewall in live environment using existing configuration but will not overwrite the previous installation. Use this feature for safely previewing or testing upgrades.
For new installations or migrations, follow this process:
- We must have a 2nd USB drive formatted with FAT or FAT32 File system. 
  - Preferable non-bootable USB drive.
- Create a conf directory on the root of the USB drive
- Place an unencrypted <downloaded backup>.xml into /conf and rename the file to config.xml ( /conf/config.xml )
- Put both the Installation media and the 2nd USB drive into the system and power up / reboot.
- Boot the system from the OPNsense Installation media via Boot Menu or BIOS (UEFI).
- Press any key when you see: “Press any key to start the configuration importer”
- Type the device name of the 2nd USB Drive, e.g. da0 or nvd0 , and press Enter. 
  - If Importer is successful, the boot process will continue into the Live environment using the configuration stored on the USB drive.
  - If unsuccessful, the importer will error and return to the device selection prompt. In this case, repeat steps 1-3 again.
Live Environment
After booting with an OPNsense Full Image (DVD, VGA, Serial), the firewall will be in the Live environment with and without the use of OPNsense Importer. We can interact with the Live environment via Local Console, GUI (HTTPS), or SSH.
By default, we can log into the shell using the user root with the password
opnsense to operate the live environment via the local console.
The GUI is accessible at https://192.168.1.1/ using Username:
root Password: opnsense by default (unless a previous configuration was imported).
Using SSH we can access the firewall at IP 192.168.1.1 . Both the root and installer users are available with the password specified above.
Note
The installation media is read-only, which means your current live configuration will be lost after reboot.
Continue to OPNsense Installer to install OPNsense to the local storage device.
OPNsense Installer
Note
To invoke the installer, log in with user installer and password opnsense
After successfully booting up with the OPNsense Full Image (DVD, VGA, Serial),
the firewall will be at the Live Environment’s login: prompt.  To start the
installation process, login with the user installer and password opnsense.
If Importer was used to import an existing configuration, the installer and root
user password will be the root password from the imported configuration.
If the installer user does not work, log in as user root and select: 8) Shell
from the menu and type opnsense-installer.  The opnsense-importer can also
be run this way should you need to rerun the import.
The installer can always be run to clone an existing system, even for Nano images. This can be useful for creating live backups for later recovery.
Tip
The installer can also be started from an internal host using SSH.  The default IP
address is 192.168.1.1
Attention
When installing an appliance with SD or MMC flash card storage, ensure you select the UFS filesystem
and choose mmcsd0 in the Disk Selection. Do not choose da0 as that is the USB drive.
As an example, this is applicable to the DEC677 Desktop Security Appliance.
The installation process involves the following steps:
- Keymap selection - The default configuration should be fine for most occasions.
- Install (UFS|ZFS) - Choose UFS or ZFS filesystem. ZFS is in most cases the best option as it is the most reliable option, but it does require enough capacity (a couple of gigabytes at least).
- Partitioning (ZFS) - Choose a device type. The default option (stripe) is usually acceptable when using a single disk.
- Disk Selection (ZFS) - Select the Storage device e.g. da0 ornvd0
- Last Chance! - Select Yes to continue with partitioning and to format the disk. However, doing so will destroy the contents of the disk.
- Continue with recommended swap (UFS) - Yes is usually fine here unless the install target is very small (< 16GB)
- Select Root Password - Change and confirm the new root password
- Select Complete Install - Exits the installer and reboots the machine. The system is now installed and ready for initial configuration.
Warning
You will lose all files on the installation disk. If another disk is to be used then choose a Custom installation instead of the Quick/Easy Install.
Nano Image
To use the nano image follow this process:
- Create the system disk by using the Nano image. See Installation Media how to write the nano image to disk.
- Install the system disk drive into the system.
- Configure the system (BIOS) to boot from this disk.
- After the system boots, the firewall is ready to be configured.
Using the nano image for embedded systems, your firewall is already up and running. The configuration settings to enable Memory Disks (RAM disks) that minimize write cycles to relevant partitions by mounting these partitions in system memory and reporting features are disabled by default.
Initial Configuration
After installation, the system will prompt you for the interface assignment. If you ignore this, then the default settings will be applied. Installation ends with the login prompt.
By default you have to log in to enter the console.
Welcome message
* * * Welcome to OPNsense [OPNsense 15.7.25 (amd64/OpenSSL) on OPNsense * * *
WAN (em1)     -> v4/DHCP4: 192.168.2.100/24
LAN (em0)     -> v4: 192.168.1.1/24
