---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-29-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "ethernet", "license"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-29.md
source_anchor: ""
source_lines: [94, 138]
sha256: 0cd628483eace94927110bae1b7417ba05e3d30b329fbd6f4f8bffdae83ad63f
---

# Overview

**Optional**, you can set mirror device between main one, if not needed, skip this step:

- Choose Local Package Sources and enable Mirror device. Set Primary Server where the packages are located, 10.155.136.50. Check Interval **minimum** setting can be set to 00:07:12, at which device will connect using Winbox to a main device and check for packages.
If new packages are available, it will begin to download, please note download process is slow and may require some time when large amount of files are used. In case some failures, download will resume on next Check.

- New "packs" folder is created, where mirror device will store packages:

- Add new package source on device which will be updated, in this example we use mirror device 10.155.136.71:

- Once you click Refresh in Local Update packages tab, device using Winbox will try to connect to source and check if there are new packages.

- Choose packages and click download, after download completes device will be needed to reboot for update.

- Use system/package/local-update/refresh to automate this in your scripts and tools fetch url= can be used to download packages from our web page, for example: tool/fetch url=https://download.mikrotik.com/routeros/7.16.1/routeros-7.16.1-arm.npk

## RouterOS upgrade using Dude

## License issues

When upgrading from older versions, there could be issues with your license key. Possible scenarios:

- When upgrading from RouterOS v2.8 or older, the system might complain about an expired upgrade time. To override this, use Netinstall to upgrade. Netinstall will ignore old license restrictions and will upgrade
- When upgrading to RouterOS v4 or newer, the system will ask you to update the license to a new format. To do this, ensure your Winbox PC (not the router) has a working internet connection without any restrictions to reach www.mikrotik.com and click "update license" in the license menu.

# Netinstall

NetInstall is a widely-used installation tool for RouterOS. It runs on Windows systems or via a command-line tool, netinstall-cli, on Linux, or through Wine (with superuser permissions required).

The NetInstall utilities can be downloaded from the MikroTik download section.

NetInstall is also used to re-install RouterOS in cases where a previous installation has failed, been damaged, or where access passwords have been lost.

To use NetInstall, your device must support booting from Ethernet, with a direct Ethernet connection between the NetInstall computer and the target device. All RouterBOARDs support PXE network booting, which can be enabled in the RouterOS "routerboard" menu (if RouterOS is accessible) or in the bootloader settings using a serial console cable.

**Note:** For RouterBOARD devices without a serial port or RouterOS access, you can activate PXE booting using the Reset button.

NetInstall can also directly install RouterOS onto a disk (USB/CF/IDE/SATA) connected to the NetInstall Windows machine. Once installed, simply transfer the disk to the Router machine and boot from it.

**Attention!** Do not try to install RouterOS on your system drive. Action will format your hard drive and wipe out your existing OS.

# CD Install

# RouterOS Package Types

Information about RouterOS packages can be found here
