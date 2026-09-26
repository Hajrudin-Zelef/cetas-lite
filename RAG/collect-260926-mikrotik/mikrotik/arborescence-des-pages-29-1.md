---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-29-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-29.md
source_anchor: ""
source_lines: [1, 93]
sha256: 6639ea13f571dee41f913f353387d3538671e39654b28fa6fed87cefcc1e08aa
---

# Overview

MikroTik devices are preinstalled with RouterOS, so installation is usually not needed, except in the case where installing RouterOS on a bare metal x86 PC or a virtual machine via CHR images. The upgrade procedure on already installed devices is straightforward.

# Upgrading

## Version numbering

RouterOS versions are numbered sequentially when a period is used to separate sequences, it does *not* represent a decimal point, and the sequences do *not* have positional significance. An identifier of 2.5, for instance, is not "two and a half" or "halfway to version three", it is the fifth second-level revision of the second first-level revision. Therefore v5.2 is older than v5.18, which is newer.

RouterOS versions are released in several "release chains": Long term, Stable, Testing, and Development. When upgrading RouterOS, you can choose a release chain from which to install the new packages.

- **Long term** : Released rarely, and includes only the most critical fixes, upgrades within one number branch do not contain new features. When a**Stable** release has been out for a while and seems to be stable enough, it gets promoted into the long-term branch, replacing an older release, which is then moved to the archive. This consecutively adds new features.
- **Stable** : Released every few months, including all tested new features and fixes.
- **Testing** : Released every few weeks, only undergoes basic internal testing, and should not be used in production.
- **Development** : Released when necessary. Includes raw changes and is available for software enthusiasts for testing new features.

## Standard upgrade

The package upgrade feature connects from the router to the MikroTik download servers via HTTPS (since 7.23) and checks if there is a newer RouterOS version in the selected release channel. This menu can also be used for downgrading, if you change the channel to one that offers an older, but more stable release. Note the feature connects from the router, not your computer, so the router itself needs HTTPS connectivity to MikroTik servers. Make sure TCP port 443 is allowed in other firewalls that might be in front of this router.

After clicking the *Check For Updates* button in QuickSet or in the System → Packages menu, the *Check For Updates* window will open with the current or the latest changelog (if a newer version exists). If newer version exists, buttons *Download* and *Download&Install* will appear. By clicking the *Download* button the newest version will be downloaded (manual device reboot is required), by clicking *Download&Install*, download will start, and after a successful download, the device will be rebooted.

The versions offered will depend on the selected release channel. Not all versions might be available. It will not be possible to upgrade from an older version to the latest version in one go, when using check-for-updates approach. For example, if running RouterOS v6.x, even selecting the major release upgrade channel, called "Upgrade", you will only see v7.12.1 as the available version. You must first upgrade to that intermediate version and only then newer releases will be available in the channels. This intermediate step can be done using check for updates too, but you will simply have to repeat check for updates after the first update to the intermediate version.

If custom packages are installed, the downloader will take that into account and download all necessary packages.

# Settings

**Sub-menu:** `/system package update`

| Property | Description | 
|---|---|
| **channel (development \| long-term  \| stable \| testing )** | Upgrade channel to use when checking for new versions. See above. | 
| **check-certificate** (no \| yes \| yes-without-crl; Default:**yes** ) | Whether and how to validate the server SSL certificate. Recommended to always use "yes". | 
| **ip-version** (auto \| ipv4 \|ipv6; Default:**auto** ) | IPv4 or IPv6 preference | 
| **mode** (http \| https; Default: https ) | You can use http in case your network blocks https, or there is another reason to use plain http, but it is suggested to use HTTPS | 

It is strongly recommended to upgrade the bootloader after RouterOS update. To upgrade the bootloader, execute command "*/system routerboard upgrade*" in CLI, followed by a reboot. Alternatively, navigate to the GUI System → RouterBOARD menu and click the "Upgrade" button, then reboot the device.

You can **automate** the upgrade process by running a script in the system scheduler. This script queries the MikroTik upgrade servers for new versions, if the response received says "New version is available", the script then issues the upgrade command below. Important note, this will not work, if you are running it for the first time on a release that is older. It might not see latest versions as available, if you are running v6.x, you would first have to manually select the "Upgrade" channel to do a major release upgrade to v7.12.1 intermediate version, and only afterwards newer v7 releases will be visible in the upgrade channels. 

## Manual upgrade 


You can upgrade RouterOS in the following ways:

- WinBox – drag and drop files to the Files menu
- WebFig - upload files from the Files menu
- FTP - upload files to the root directory

It is strongly recommended to upgrade the bootloader after upgrading RouterOS. To upgrade the bootloader, execute command "*/system routerboard upgrade*" in CLI, followed by a reboot. Alternatively, navigate to the GUI System → RouterBOARD menu and click the "Upgrade" button, then reboot the device.

### Manual upgrade process

- First step - visit www.mikrotik.com and head to the Software page, then choose the architecture of the system you have the RouterOS installed on (system architecture can be found in System → Resource section);
- Download the **routeros *(main)*** and extra packages that are installed on a device;
- Upload packages to a device using one of the previously mentioned methods:

**Menu:** */system/package/update install **ignore-missing*** command allows upgrading only the RouterOS main package, while omitting packages that are either missing or not uploaded during a manual upgrade process.

#### Using WinBox

Choose your system type, and download the upgrade package. Connect to your router with WinBox, Select the downloaded file with your mouse, and drag it to the Files menu. If some files are already present, make sure to put the package in the root menu, not inside the hotspot folder! The upload will start.

After it finishes - reboot the device. The New version number will be seen in the Winbox Title and in the Packages menu

#### Using FTP

- Open your favorite SFTP program (in this case it is Filezilla), select the package, and upload it to your router (demo2.mt.lv is the address of my router in this example). note that in the image I'm uploading many packages, but in your case - you will have one file that contains them all
- if you wish, you can check if the file is successfully transferred onto the router (optional):

- reboot your router for the upgrade process to begin:

[admin@MikroTik] >/system reboot
Reboot, yes? [y/N]: y

- after the reboot, your router will be up to date, you can check it in this menu:

[admin@MikroTik] >/system package print

- if your router did not upgrade correctly, make sure you check the **log**

[admin@MikroTik] >/log print without-paging

## RouterOS local upgrade

**Sub-menu:** `system/package/local-update/`

You can upgrade one or multiple MikroTik routers within your local network by using one device which have all needed packages. Feature is available from **7.17beta3** version in (system > packages local update) and will replace (system > auto update) feature. Here is a simple example with 3 routers (the same method works on networks with infinite numbers of routers):

Place needed packages under Files menu, on your main router:

