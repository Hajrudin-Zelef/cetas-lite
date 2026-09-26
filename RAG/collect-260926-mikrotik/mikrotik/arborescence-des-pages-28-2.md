---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-28-2
title: "arborescence-des-pages-28"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["nand"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-28.md
source_anchor: ""
source_lines: [98, 167]
sha256: 43e5f061c1527e097b2d17e2dcab6cac0352d61359826ce65fb6ac48a701fb4f
---

# arborescence-des-pages-28

The "Check installation" function ensures the integrity of the RouterOS system by verifying the readability and correct placement of files. Its primary purpose is to confirm the health and status of your NAND/Flash storage.

**Menu:** */system/package/update install ignore-missing* command allows upgrading only the RouterOS main package, while omitting packages that are either missing or not uploaded during a manual upgrade process.

## Auto install

It is also possible to **automatically** install packages after uploading them to the router with FTP or SFTP. The package file must be named with the extension *.auto.npk. Once the file will be uploaded, router will automatically go into reboot in order to install the package.

".auto.npk" in the filename is mandatory for a package to be automatically installed.

## Local Update

Instead of connecting directly to MikroTik servers you can upload package files to one of your local RouterOS device, and use it as a local package server.

**Menu:** */system package* *local-update*

| Command | Description | 
|---|---|
| **download-all** | Downloads all compatible (matching device architecture) packages that are available on the local package server. Downloaded packages are saved in root directory. | 
| **download** | Downloads specific compatible (matching device architecture) packages that are available on the local package server. Downloaded packages are saved in root directory. | 
| **refresh** | Refreshes/checks the list of available compatible (matching device architecture) packages on the local package server. | 

Server from which to get the package can be defined in *system/package/local-update/ update-package-source/*

***update-package-source*** properties list:

| Property | Description | 
|---|---|
| **address** (IPv4 a*ddress [IPv4]/IPv6 address [IPv6];*  Default: ) | Address of the local package server. | 
| **user** (string; Default: ) | Username that is used for accessing the local package server. | 
| **password** (string; Default: )*sensitive* | Password that is used for accessing the local package server. | 

Also, you can mirror packages (for all architectures) from your main local package server using *system/package/local-update/mirror/*Downloaded packages saved into 

*packs*folder in root directory.

**properties list:**

*mirror*
| Property | Description | 
|---|---|
| **primary-server** (IPv4 a*ddress [IPv4]/IPv6 address [IPv6];*  Default: ) | Address of the primary local package server. | 
| **secondary-server** (IPv4 a*ddress [IPv4]/IPv6 address [IPv6];* Default: ) | Address of the secondary local package server. | 
| **user** (string; Default: ) | Username that is used for accessing the local package server. | 
| **password** (string; Default: )*sensitive* | Password that is used for accessing the local package server. | 
| **check-interval** (time [HH:MM:SS]; Default: 24:00:00) | Time interval at which device checks the local package server for new packages availability, if new package/packages is located begins package download. (only downloads the packages that are not already present on the device) | 
| **enabled** (yes \| no; default: no) | Whether to enable or no the periodical check and download of packages from the local package server. | 

**Menu:** */system package* *local-update mirror* 

| Command | Description | 
|---|---|
| **force-check** | Checks the local package server for new packages availability, if new package/packages is located begins package download. (only downloads the packages that are not already present on the device) | 

 

## Examples

### Listing packages


*zerotier* package is disabled, but installed; *iot* package is available on the server, but has not been downloaded to the router and enabled; *dude* package is scheduled for uninstall.

Uninstall package

Disable package

Downgrade

Cancel uninstall or disable action
