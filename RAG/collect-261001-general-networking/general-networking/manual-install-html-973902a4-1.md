---
id: collect-261001-general-networking/general-networking/manual-install-html-973902a4-1
title: "manual-install-html-973902a4"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "memory"]
source: docs/RAG/collect-261001-general-networking/manual-install-html-973902a4.md
source_anchor: ""
source_lines: [1, 82]
sha256: 6f72c6d4953bbced84746f81676439694d33436beefc5e1c87bdcc81112e6d7c
---

# manual-install-html-973902a4

Initial Installation & Configuration
Note
Just looking for how to invoke the installer? When the live environment has been started, log in with user installer and password opnsense.
Architecture
The software setup and installation of OPNsense® is available for the x86-64 microprocessor architecture only.
Embedded vs Full
OPNsense offers two Image types with all major releases: embedded and full images. The Embedded Image is intended for environments where preinstalling the storage media is required due to a lack of local resources on the firewall like storage, and/or console access (VGA/Serial). The image is tailored to reduce write cycles as well, but the image can be used anywhere. Another reason for the Embedded Image is to eliminate the need for local console access for installing OPNsense. Installation is managed by prewriting the image to a storage device, installing the storage device, and booting the system.
Full Images provide installation tools like OPNsense Importer, Live Environment, and Installer. Full Images are released to support different console/hardware installation requirements.
Both image types can be installed and run from virtual disks (VM), SD memory cards, USB disks, solid-state disks (SSD), or hard disk drives (HDD).
The main differences between embedded and full images are:
| Embedded | Full | 
|---|---|
| Writes to RAM disk | Writes to local disk | 
| No log data retention after reboot | Log data retention after reboot | 
| Not intended for local disk writes | Suitable for disk writes. | 
| Embedded only use, SWAP file is optional | Can enable RAM disk for embedded mode. | 
The embedded image stores logging and cache data in memory only, while full versions will keep the data stored on the local drive. A full version can mimic the behavior of an embedded version by enabling RAM disks, which is especially useful for SD memory card installations.
Warning
See the chapter Hardware Sizing & Setup for further information on hardware requirements prior to an install.
Installation Images
Depending on your hardware and use case, different installation options are available:
| Type | Description | Image Type | 
|---|---|---|
| dvd | ISO image boots into a live environment in VGA-only mode with UEFI support | Full | 
| vga | USB image boots into a live environment in VGA-only mode with UEFI support | Full | 
| serial | USB image boots into live environment running in serial console (115200) mode only with UEFI support | Full | 
| nano | Image for preinstalling onto >=4 GB USB drives, SD, or CF cards for use with embedded devices running in serial console (115200) mode with secondary VGA support (no kernel messages though) | Embedded | 
Note
All Full Image types can run both OPNsense Importer before booting into the Live environment and also run Installer once booted into the Live environment.
Warning
Flash memory cards will only tolerate a limited number of writes and re-writes. For the Nano image, /var/log and /tmp are memory disks by default to prolong the lifetimes of CF/SD cards and SSDs in special cases.
To enable non-embedded versions: Go to , change the setting, then reboot. Consider enabling an external syslog server as well.
Image Filename Composition
Note
Please be aware that the latest installation media does not always correspond with the latest released version available. OPNsense installation images are provided on a scheduled basis with major release versions in January and July. More information on our release schedule is available from our package repository, see README. We are encouraged to update OPNsense after installation to be on the latest release available, see Update Page.
Download and Verification
The OPNsense distribution can be downloaded from one of our mirrors.
OpenSSL is used for image file verification. 4 files are needed for verification process:
- The SHA-256 checksum file (<filename>.sha256)
- The bzip-compressed image file (<filename>.<image>.bz2)
- The signature file for the uncompressed image file (<filename>.<image>.sig)
- The OpenSSL public key (<filename>.pub)
Use one of the OPNsense mirrors to download these files:
- Go to the bottom of OPNsense download page.
- Click one of the available mirrors closest to your location.
- Download one of each file mentioned above for your Image type.
The OpenSSL public key (.pub) is required to verify against. Although the file is available on the mirror’s repository, you should not trust the copy there. Download it, open it up, and verify the public key matches the one from other sources. If it does not, the mirror may have been hacked, or you may be the victim of a man-in-the-middle attack. Some other sources to get the public key from include:
- https://pkg.opnsense.org/releases/mirror/README
- https://forum.opnsense.org/index.php?board=11.0
- https://opnsense.org/blog/
- https://github.com/opnsense/changelog/tree/master/community
- https://pkg.opnsense.org (/<FreeBSD:<version>:<architecture>/<release version>/sets/changelog.txz)
Note
Only major release announcements for images contain the public key, and update release announcements will not. i.e. 22.1 will have a copy of the public key in the release announcement, but 22.1.9 will not.
Once you download all the required files and verify that the public key matches the public key found in one of the alternate sources listed above, you can be relatively confident that the key has not been tampered with. To verify the downloaded image, run the following commands (substituting the filenames in brackets for the files you downloaded):
openssl sha256 OPNsense-<filename>.bz2
Match the checksum command output with the checksum values in the file OPNsense-<version>-OpenSSL-checksums-amd64.sha256.
If the checksums don’t match, redownload your image file.
If the checksums match, continue with the verification commands.
openssl base64 -d -in OPNsense-<filename>.<image>.sig -out /tmp/image.sig
openssl dgst -sha256 -verify OPNsense-<filename>.pub -signature /tmp/image.sig OPNsense-<filename>.<image>
Warning
Make sure to unpack the image using bunzip2 before verifying. Our signatures are generated before compressing them
(as of OPNsense version 24.1)
If the output of the second command is “Verified OK”, your image file was verified successfully, and it is safe to install from it. Any other outputs, and you may need to check your commands for errors, or the image file may have been compromised.
Installation Media
Now that you have downloaded and verified the installation image from above, you must unpack the image file before you can write the image to disk. For Unix-like OSes use the following command:
bzip2 -d OPNsense-<filename>.bz2
For Windows use an application like 7zip.  The .bz2 will
be removed from the end of the filename after command/application completes.
After unpacking the image you can create the installation media. The easiest method to install OPNsense is to use the USB “vga” Image. If your target platform has a serial console interface choose the “serial” image. If you need to know more about using the serial console interface, consult the serial access how-to.
Write the image to a USB flash drive (>=1 GB) or hard disk, using either dd for Unix-like OSes or, for Windows, physdiskwrite, Etcher, or Rufus.
FreeBSD
dd if=OPNsense-##.#.##-[Type]-[Architecture].img of=/dev/daX bs=16k
Where X = the device number of your USB flash drive (check dmesg)
OpenBSD
dd if=OPNsense-##.#.##-[Type]-[Architecture].img of=/dev/rsd6c bs=16k
The device must be the ENTIRE device (in Windows/DOS language: the ‘C’ partition), and a raw I/O device (the ‘r’ in front of the device “sd6”), not a block mode device.
Linux
sudo dd if=OPNsense-##.#.##-[Type]-[Architecture].img of=/dev/sdX bs=16k
where X = the IDE device name of your USB flash drive (check with hdparm -i /dev/sdX) (ignore the warning about trailing garbage - it’s because of the digital signature)
macOS
