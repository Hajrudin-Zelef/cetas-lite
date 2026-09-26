---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-17-1
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-17.md
source_anchor: ""
source_lines: [1, 74]
sha256: 7f92ea565a3766b71d17c009c237cf434a42a8f6cb1dda7881d7184c5349e73a
---

# Introduction

Netinstall is a tool for installing and reinstalling MikroTik devices running RouterOS.

Always reinstall device using Netinstall if you suspect that your device is not working properly.

The tool is available for Windows with a graphical interface and for Linux command line tool.

Be careful. Netinstall re-formats the system's drive, all configuration and user files will be lost.

Netinstall does not erase the RouterOS license key, nor does it reset RouterBOOT related settings, for example, CPU frequency is not changed during reinstalling the device.

## Device installation workflow using Netinstall:

- prepare and start Netinstall host on same L2 network segment where re-installable device will be connected, preferable dedicated NIC and hub/switch to avoid IP/DHCP/BOOTP conflict caused by other devices on same network
- boot device from Netinstall by holding Reset button till device is detected by Netinstall. As alternative device RouterBOOT can be configured for network boot, during which device will boot to Netinstall
- Install device using specified packages and configuration scripts
- During next boot device applies configuration parameters specified in Netinstall

## Device configuration scripts:

**Configure script** - installs on device custom default configuration script which replaces with RouterOS supplied default configuration script. The script can access factory set passwords with read-only variables `$defconfPassword` and `$defconfWifiPassword` starting from RouterOS 7.10beta8.

**Mode script** - puts on device one‑time custom script to run on the device’s first boot after installation. This script can be used to configure **device‑mode** and **protected‑routerboot**  and other settings during device initial configuration.

Configure and Mode scripts have a maximum execution timeout of 120 seconds.

Configure script will be kept during RouterOS updates and used after further configuration resets till device will be reinstalled with new script or removed if no script provided.

Mode script feature requires RouterOS and Netinstall version 7.22 or newer.

Configure and Mode scripts are regular/import file, accepts valid MikroTik RouterOS CLI commands.

For more reference see:

# Netinstall for Windows

The available actions and parameters are as follows:

| Controls/Actions | Description | 
|---|---|
| **Netbooting** | Allows to Enable, Disable and configure IP addressing for Netinstall build-in in BOOTP and TFPT server required for netbooting MikroTik device using Netinstall client | 
| **Shutdown** | Performs shutdown of selected MiktoTik device booted to Netinstall | 
| **Reboot** | Performs reboot of selected MiktoTik device booted to Netinstall | 
| **Install**  | Performs MiktoTik device Netinstall using selected options | 
| **Cancel** | Cancel current in progress installation | 
| **Routers/Drives** | Device selection for Netinstall. For device to appear in list device must be started to Netinstall mode | 
| **Packages** | Allows to specify location where RouterOS packages files are located and create package sets for next installations | 

| Netinstall parameters | Description | 
|---|---|
| **Auto reboot** | Set action what is performed after device installation is successfully completed | 
| **Software ID** | Displays current device Software ID | 
| **Key** | Allows to specify and install new license key during installation. By default Netinstall will keep detected license key from device. | 
| **Keep old configuration** | Instruct Netinstall to read and restore device core configuration (/export; /users) after device installation. This options does not keep user files, containers etc. | 
| **Keep branding** | Instruct Netinstall to keep device branding package if it present on device. Factory installed branding packages can not be discarded and always will be kept. | 
| **Apply default config** | If set, after device installation and reboot device will apply default configuration script. `/system/default-configuration/script/print` | 
| **Configure script** | Allows during device installation to install on device custom default configuration script which replaces with RouterOS supplied default configuration script. `/system/default-configuration/custom-script/print` This script will be kept during RouterOS updates and used after further configuration resets till device will be reinstalled with new script or removed if no script provided. | 
| **Mode script** | Specifies a one‑time custom script to run on the device’s first boot after installation. Use this script  to configure **device‑mode** and**protected‑routerboot** and other settings during device deployment. Mode script executes before any custom or default configuration scripts. Upon completion, script is automatically removed from the device. If the script modifies the device-mode, the device will be reboot immediately after execution. This feature requires RouterOS and Netinstall version 7.22 or newer. | 
| **IP address Gateway Baudrate** | Netinstall will autocreate **Configure script** with specified parameters for initial device configuration using IPv4 connectivity or serial console. Parameter available only when device is Installed to empty configuration, eg no other configuration option options specified. | 

# Netinstall for Linux

The Netinstall Linux version `netinstall-cli` is a command line tool, which offers nearly the same parameters as the Windows counterpart. 

The tool requires root privileges and must be run as root or use sudo.


Command line:

`netinstall-cli` [-r] [-e] [-b] [-m [-o]] [-f] [-k <keyfile>] [-s <userscript>] [-sm <modescript>] [--mac <mac address>] {-i <interface> | -a <client-ip>} [PACKAGES]

Available actions and parameters are as follows:

